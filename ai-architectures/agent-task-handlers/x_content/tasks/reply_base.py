from abc import abstractmethod
import logging
import traceback
from typing import List, Optional
from x_content.constants import MissionChainState
from x_content.tasks.reply_subtask_base import ReplySubtaskBase
from x_content.wrappers.api import twitter_v2
from x_content.wrappers.api.twitter_v2.models.objects import ExtendedTweetInfo
from x_content.wrappers.api.twitter_v2.models.response import (
    ExtendedTweetInfosDto,
    Response,
)
from x_content.models import ReasoningLog
from x_content.wrappers.tweet_specialty import TweetSpecialty, detect_tweet_specialties

from x_content.tasks.base import MultiStepTaskBase
from x_content.tasks.utils import (
    a_move_state,
    create_twitter_auth_from_reasoning_log,
)

logging.basicConfig(level=logging.INFO if not __debug__ else logging.DEBUG)
logger = logging.getLogger(__name__)

MINIMUM_REPLY_LENGTH = 32

from x_content.wrappers.magic import sync2async
import asyncio


class ReplyTaskBase(MultiStepTaskBase):
    resumable = True

    @abstractmethod
    def get_subtask_cls(
        self, log: ReasoningLog, specialties: List[TweetSpecialty]
    ) -> Optional[ReplySubtaskBase]:
        raise NotImplementedError(
            "get_subtask_cls method not implemented; cls: {}".format(
                self.__class__.__name__
            )
        )

    async def process_task(self, log: ReasoningLog) -> ReasoningLog:
        if log.state == MissionChainState.NEW:
            try:
                logger.info(
                    f"[{self.__class__.__name__}.process_task] Initializing new reply task with log ID {log.id}"
                )
                response: Response[ExtendedTweetInfosDto] = await sync2async(
                    twitter_v2.get_recent_mentioned_tweets_by_username_v2
                )(
                    auth=create_twitter_auth_from_reasoning_log(log),
                    num_tweets=5,
                    replied=0,
                    max_num_tweets_in_conversation=3,
                    preserve_img=True,
                    get_all=True,
                )
                if response.is_error():
                    raise Exception(response.error)

                mentioned_tweets = response.data.tweet_infos
                # TODO: Remove debug
                # mentioned_tweets = [x for x in mentioned_tweets if x.tweet_object.tweet_id == "1879789038519169405"]
                logger.info(
                    f"[{self.__class__.__name__}.process_task] Retrieved {len(mentioned_tweets)} recent mentions for {log.meta_data.twitter_username}"
                )

                async def get_specialties(tweet_info: ExtendedTweetInfo):
                    return tweet_info, await sync2async(
                        detect_tweet_specialties
                    )(tweet_info)

                futures = [
                    asyncio.ensure_future(get_specialties(tweet_info))
                    for tweet_info in mentioned_tweets
                ]

                mentioned_data = []

                totals = len(futures)
                for idx, future in enumerate(asyncio.as_completed(futures)):
                    try:
                        tweet_info, specialties = await future
                    except Exception as err:
                        logger.info(
                            f"[{log.id}] Error while processing index {idx+1} (out of {totals}): {err} (getting tweet specialties fail)."
                        )
                        continue

                    mentioned_data.append(
                        {
                            "tweet_info": tweet_info.to_dict(),
                            "specialties": [
                                specialty.name for specialty in specialties
                            ],
                        }
                    )
            except Exception as err:
                traceback.print_exc()
                logger.error(
                    f"[process_task] Error retrieving mentioned tweets: {err}"
                )
                return await a_move_state(
                    log,
                    MissionChainState.ERROR,
                    f"Error retrieving mentioned tweets: {err}",
                )

            logger.info(
                f"[{self.__class__.__name__}.process_task] Finished detecting specialties for all {len(mentioned_tweets)} tweets"
            )

            log.execute_info = {
                "tweets": mentioned_data,
                "task_result": [],
                "task_failed": [],
                "conversation": [],
            }

            if len(mentioned_data) == 0:
                return await a_move_state(
                    log, MissionChainState.DONE, "Mentioned tweets not found"
                )

            return await a_move_state(
                log, MissionChainState.RUNNING, "Task started"
            )

        if log.state == MissionChainState.RUNNING:
            mentioned_data = log.execute_info["tweets"]

            subtasks: List[ReplySubtaskBase] = []
            for i in range(len(mentioned_data)):
                tweet_info = ExtendedTweetInfo.model_validate(
                    mentioned_data[i]["tweet_info"]
                )
                specialties = [
                    TweetSpecialty[specialty]
                    for specialty in mentioned_data[i]["specialties"]
                ]

                subtask_cls = self.get_subtask_cls(
                    log, specialties, tweet_info
                )
                if subtask_cls is not None:
                    subtasks.append(
                        subtask_cls(
                            llm=self.llm,
                            kn_base=self.kn_base,
                            log=log,
                            tweet_info=tweet_info,
                        )
                    )

            logger.info(
                f"[{self.__class__.__name__}.process_task] Start processing {len(subtasks)} subtasks"
            )

            async def execute_subtask(subtask: ReplySubtaskBase):
                return subtask.tweet_info, await subtask.run()

            futures = [
                asyncio.ensure_future(execute_subtask(subtask))
                for subtask in subtasks
            ]

            totals = len(futures)
            for idx, future in enumerate(asyncio.as_completed(futures)):
                try:
                    tweet_info, result = await future
                    log.execute_info["task_result"].append(result)
                    logger.info(
                        f"[{log.id}] Successfully processed subtask index {idx+1} (out of {totals}) (tweet_id={tweet_info.tweet_object.tweet_id})"
                    )
                except Exception as err:
                    traceback.print_exc()
                    log.execute_info["task_failed"].append(
                        {
                            "tweet_id": tweet_info.tweet_object.tweet_id,
                            "error": str(err),
                        }
                    )
                    logger.info(
                        f"[{log.id}] Error while processing index {idx+1} (out of {totals}): {err} (subtask fail) (tweet_id={tweet_info.tweet_object.tweet_id})"
                    )
                    continue

            n_success = len(log.execute_info["task_result"])
            n_failed = len(log.execute_info["task_failed"])

            logger.info(
                f"[{log.id}] Finished processing all {len(futures)} subtasks ({n_success} success, {n_failed} failed)"
            )

            if n_success == 0:
                return await a_move_state(
                    log,
                    MissionChainState.ERROR,
                    f"Failed to reply all mentioned tweets",
                )

            return await a_move_state(
                log,
                MissionChainState.DONE,
                f"Replied to mentioned tweets ({n_success} success, {n_failed} failed)",
            )

        return log
