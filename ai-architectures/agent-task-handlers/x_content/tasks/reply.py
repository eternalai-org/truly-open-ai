import logging
import traceback
from x_content.constants import GAME_TASKS_WHITELIST, MissionChainState
from x_content.wrappers.api import twitter_v2
from x_content.wrappers.api.twitter_v2.models.response import ExtendedTweetInfosDto, Response, TweetsDto
from x_content.wrappers.postprocess import post_process_tweet
from x_content.tasks.utils import render_reply_thread_conversation
from x_content.models import ReasoningLog
from x_content.wrappers.postprocess import remove_emojis

from .base import MultiStepTaskBase
from .utils import (
    a_move_state, 
    create_twitter_auth_from_reasoning_log,
    is_create_game_tweet,
    is_create_game_tweet_id,
    is_create_token_tweet,
    render_reply_game_conversation
)
from x_content.llm import OnchainInferResult

logging.basicConfig(level=logging.INFO if not __debug__ else logging.DEBUG)
logger = logging.getLogger(__name__)

MINIMUM_REPLY_LENGTH = 32

from x_content.wrappers.magic import sync2async
import asyncio

class ReplyTask(MultiStepTaskBase):
    resumable = True

    async def process_task(self, log: ReasoningLog) -> ReasoningLog:
        is_game_agent = log.meta_data.twitter_username in GAME_TASKS_WHITELIST
        if log.state == MissionChainState.NEW:
            try:
                response: Response[ExtendedTweetInfosDto] = await sync2async(twitter_v2.get_recent_mentioned_tweets_by_username_v2)(
                    auth=create_twitter_auth_from_reasoning_log(log),
                    num_tweets=10,
                    max_num_tweets_in_conversation=1,
                    get_all=True,
                )
                if response.is_error():
                    raise Exception(response.error)
                mentioned_tweets = [x.to_dict() for x in response.data.tweet_infos]

                if is_game_agent:
                    # Remove tweets that is in subtree of game tweet
                    mentioned_tweets = [
                        x for x in mentioned_tweets
                        if not await sync2async(is_create_game_tweet_id)(x["conversation_id"])
                    ]
                else:
                    # Remove tweets that is in subtree of game tweet (except the game tweet itself)
                    mentioned_tweets = [
                        x for x in mentioned_tweets
                        if x["conversation_id"] != x["tweet_object"]["tweet_id"] and not await sync2async(is_create_game_tweet_id)(x["conversation_id"])
                    ]

                mentioned_tweets = [
                    x for x in mentioned_tweets 
                    if not await sync2async(is_create_token_tweet)(x)
                ]
            except Exception as err:
                traceback.print_exc()
                logger.error(f"[process_task] Error retrieving mentioned tweets: {err}")
                return await a_move_state(log, MissionChainState.ERROR, f"Error retrieving mentioned tweets: {err}")
            
            log.execute_info = {
                "tweets": mentioned_tweets,
                "task_result": [],
                "conversation": [],
            }

            if len(mentioned_tweets) == 0:
                return await a_move_state(log, MissionChainState.DONE, "Mentioned tweets not found")

            for idx in range(len(mentioned_tweets)):
                tweet_info = mentioned_tweets[idx]
                context_resp: Response[TweetsDto] = await sync2async(twitter_v2.get_full_context_from_a_tweet)(
                    tweet_info["tweet_object"], 
                    tweet_info["parent_tweet_id"]
                )
                if context_resp.is_error():
                    tweets_context = []
                else:
                    tweets_context = context_resp.data.tweets

                resp = await sync2async(twitter_v2.get_relevent_information_v2)(
                    self.kn_base,
                    tweets=tweets_context,
                )
                knowledge_v2 = {} if resp.is_error() else resp.data.to_dict()

                mentioned_tweets[idx].update(
                    full_context=[x.to_dict() for x in tweets_context],
                    knowledge=knowledge_v2,
                    knowledge_v2=knowledge_v2
                )

                if not is_game_agent and await sync2async(is_create_game_tweet)(mentioned_tweets[idx]["tweet_object"]):
                    log.execute_info["conversation"].append(
                        await sync2async(render_reply_game_conversation)(log, idx)
                    )
                else:
                    log.execute_info["conversation"].append(
                        await sync2async(render_reply_thread_conversation)(log, idx)
                    )

            return await a_move_state(log, MissionChainState.RUNNING, "Task started")

        if log.state == MissionChainState.RUNNING:
            futures = [
                asyncio.ensure_future(self.llm.agenerate(conversation_thread, temperature=0.7))
                for conversation_thread in log.execute_info["conversation"]
            ]

            for idx, infer in enumerate(asyncio.as_completed(futures)):

                try:
                    infer_result: OnchainInferResult = await infer
                except Exception as err:
                    logger.info(f"[{log.id}] Error while processing index {idx}: {err} (inference fails).")
                    continue

                result = await sync2async(remove_emojis)(infer_result.generations[0].message.content.strip('" '))
                result = await sync2async(post_process_tweet)(result)

                liked_tweet = log.execute_info['tweets'][idx]["tweet_object"]
                liked_tweet_id = liked_tweet["tweet_id"]

                if len(result) >= MINIMUM_REPLY_LENGTH:
                    await sync2async(twitter_v2.reply)(
                        auth=await sync2async(create_twitter_auth_from_reasoning_log)(log),
                        tweet_id=liked_tweet_id, 
                        reply_content=result, 
                        tx_hash=infer_result.tx_hash
                    )
                
                log.execute_info["task_result"].append({
                    "tweet_id": liked_tweet_id,
                    "reply_content": result,
                    "tx_hash": infer_result.tx_hash,
                })
                
                log = await self.commit_log(log)

            return await a_move_state(log, MissionChainState.DONE, "Replied to all mentioned tweets")

        return log
