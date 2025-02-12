import logging
import traceback
from x_content.constants import MissionChainState, AgentUsername
from x_content import constants as const
from x_content.models import ReasoningLog
from x_content.tasks.base import MultiStepTaskBase
from x_content.tasks.reply_subtask_base import ReplySubtaskBase
from x_content.tasks.utils import a_move_state, create_twitter_auth_from_reasoning_log

from x_content.wrappers.api import twitter_v2
from x_content.wrappers.api.twitter_v2.models.response import (
    GenerateActionDto,
    Response,
    ExtendedTweetInfosDto,
)
from x_content.wrappers.game import (
    GameAPIClient,
    GameInfo,
    _get_game_redis_cache,
    GameStatus,
)
from x_content.wrappers.magic import sync2async
from x_content.wrappers.tweet_specialty import is_create_game_tweet

logging.basicConfig(level=logging.INFO if not __debug__ else logging.DEBUG)
logger = logging.getLogger(__name__)


def _filter_create_game_tweets(tweet_infos):
    logger.info(
        f"[_filter_create_game_tweets] Starting to filter {len(tweet_infos)} tweets for game creation criteria"
    )

    # Filter tweets that match game criteria using is_create_game_tweet()
    filtered_tweets = [
        tweet
        for tweet in tweet_infos
        if is_create_game_tweet(tweet["tweet_object"])
    ]

    logger.info(
        f"[_filter_create_game_tweets] Found {len(filtered_tweets)} tweets matching game creation criteria out of {len(tweet_infos)} total tweets"
    )
    return filtered_tweets


# Update the existing functions to use the Redis cache class
def is_create_game_pending(_tweet_id):
    try:
        game_redis = _get_game_redis_cache()
        status = game_redis.get_game_status(_tweet_id)
        logger.info(
            f"[is_create_game_pending] Retrieved game status for tweet {_tweet_id}: {status}"
        )
        return status == GameStatus.CREATE_PENDING or status == GameStatus.NONE
    except Exception as err:
        logger.error(
            f"[is_create_game_pending] Failed to check pending status for tweet {_tweet_id}: {err}"
        )
        return False


def _mark_create_game_running(_tweet_id):
    try:
        game_redis = _get_game_redis_cache()
        logger.info(
            f"[_mark_create_game_running] Attempting to mark game for tweet {_tweet_id} as running"
        )
        result = game_redis.set_game_status(
            _tweet_id, GameStatus.CREATE_RUNNING
        )
        logger.info(
            f"[_mark_create_game_running] Successfully marked game for tweet {_tweet_id} as running"
        )
        return result
    except Exception as err:
        logger.error(
            f"[_mark_create_game_running] Failed to mark game as running for tweet {_tweet_id}: {err}"
        )
        return False


def _mark_create_game_pending(_tweet_id):
    try:
        game_redis = _get_game_redis_cache()
        logger.info(
            f"[_mark_create_game_pending] Attempting to mark game for tweet {_tweet_id} as pending"
        )
        result = game_redis.set_game_status(
            _tweet_id, GameStatus.CREATE_PENDING
        )
        logger.info(
            f"[_mark_create_game_pending] Successfully marked game for tweet {_tweet_id} as pending"
        )
        return result
    except Exception as err:
        logger.error(
            f"[_mark_create_game_pending] Failed to mark game as pending for tweet {_tweet_id}: {err}"
        )
        return False


def _mark_create_game_done(_tweet_id):
    try:
        game_redis = _get_game_redis_cache()
        logger.info(
            f"[_mark_create_game_done] Attempting to mark game for tweet {_tweet_id} as done"
        )
        game_redis.add_running_game(_tweet_id, GameStatus.CREATE_DONE)
        logger.info(
            f"[_mark_create_game_done] Successfully marked game for tweet {_tweet_id} as done"
        )
    except Exception as err:
        logger.error(
            f"[_mark_create_game_done] Failed to mark game as done for tweet {_tweet_id}: {err}"
        )


def _try_acquire_create_game_lock(_tweet_id):
    # Initialize lock variable
    lock = None

    try:
        game_redis = _get_game_redis_cache()
        logger.info(
            f"[_try_acquire_create_game_lock] Attempting to acquire lock for tweet {_tweet_id} with {const.CREATE_GAME_LOCK_EXPIRY}s expiry"
        )
        # Attempt to acquire the lock with an expiry time
        lock = game_redis.acquire_create_game_lock(
            _tweet_id, const.CREATE_GAME_LOCK_EXPIRY  # e.g., 300 seconds
        )
        logger.info(
            f"[_try_acquire_create_game_lock] Lock acquisition {'successful' if lock else 'failed'} for tweet {_tweet_id}"
        )
        return lock  # Returns True if acquired, False if already locked

    except Exception as e:
        # Log any errors during lock acquisition
        logger.error(
            f"[_try_acquire_create_game_lock] Failed to acquire lock for tweet {_tweet_id}: {str(e)}"
        )
        return False

    finally:
        # Always execute this block, even if an exception occurs
        if not lock:  # If lock wasn't acquired
            try:
                logger.info(
                    f"[_try_acquire_create_game_lock] Attempting to release lock for tweet {_tweet_id}"
                )
                # Attempt to release the lock
                game_redis.release_create_game_lock(_tweet_id)
                logger.info(
                    f"[_try_acquire_create_game_lock] Successfully released lock for tweet {_tweet_id}"
                )
            except Exception as e:
                # Log any errors during lock release
                logger.error(
                    f"[_try_acquire_create_game_lock] Failed to release lock for tweet {_tweet_id}: {str(e)}"
                )


"""
Get list of agent usernames from tweet mentions
"""


def _get_agent_usernames(_tweet_object):
    username_to_remove = AgentUsername.CRYPTOCOMIC_AI
    try:
        unfiltered_mentions = _tweet_object.get("mentions", [])
        logger.info(
            f"[_get_agent_usernames] Found {len(unfiltered_mentions)} total mentions in tweet"
        )

        filtered_mentions = [
            mention
            for mention in unfiltered_mentions
            if mention["username"] != username_to_remove
        ]

        logger.info(
            f"[_get_agent_usernames] Filtered to {len(filtered_mentions)} valid agent mentions"
        )

        if len(filtered_mentions) < 2:
            raise Exception(
                f"Need at least 2 agents mentioned, found only {len(filtered_mentions)}"
            )
        agent_usernames = [
            mention["username"] for mention in filtered_mentions
        ]
        if not agent_usernames:
            raise Exception("No mentioned usernames found after filtering")
        logger.info(
            f"[_get_agent_usernames] Successfully extracted agent usernames: {agent_usernames}"
        )
        return agent_usernames, None
    except Exception as e:
        logger.error(
            f"[_get_agent_usernames] Failed to extract agent usernames: {e}"
        )
        return None, e


"""
Get mapping of agent usernames to wallet addresses
mapping = {
    "agent1": "0x123...abc",
    "agent2": "0x456...def"
}
"""


def _get_agent_wallet_mapping(game_info: GameInfo):
    try:
        if game_info is None or not game_info.agent_wallets:
            return {}, Exception("Game info or agent wallets is nil")

        wallet_mapping = {
            wallet.username: wallet.address
            for wallet in game_info.agent_wallets
        }
        logger.info(
            f"[_get_agent_wallet_mapping] Successfully mapped {len(wallet_mapping)} agent usernames to wallet addresses"
        )
        return wallet_mapping, None
    except Exception as e:
        logger.error(
            f"[_get_agent_wallet_mapping] Failed to create wallet mapping: {e}"
        )
        return None, e


def _create_wallet_tweet(wallet_mapping):
    """Create tweet text with wallet mapping info"""
    hours = const.GAME_DURATION // 3600
    minutes = (const.GAME_DURATION % 3600) // 60
    wallet_tweet = const.GAME_CREATED_TWEET.format(
        hours=hours, minutes=minutes
    )
    for username, wallet in wallet_mapping.items():
        wallet_tweet += f"{username}: {wallet}\n"
    logger.info(
        f"[_create_wallet_tweet] Created wallet announcement tweet with {len(wallet_mapping)} agent mappings"
    )
    return wallet_tweet


async def _post_wallet_tweet(log: ReasoningLog, tweet_id, wallet_tweet):
    """Post wallet mapping tweet as reply"""
    try:
        logger.info(
            f"[_post_wallet_tweet] Attempting to post wallet announcement tweet in reply to {tweet_id}"
        )
        resp: Response[GenerateActionDto] = await sync2async(
            twitter_v2.reply_multi_unlimited
        )(
            auth=create_twitter_auth_from_reasoning_log(log),
            tweet_id=tweet_id,
            reply_content=wallet_tweet,
        )
        if resp.is_error() or not resp.data.success:
            raise Exception(f"Failed to post wallet tweet")
        logger.info(
            f"[_post_wallet_tweet] Successfully posted wallet announcement tweet"
        )
        return None
    except Exception as e:
        logger.error(
            f"[_post_wallet_tweet] Failed to post wallet announcement tweet: {e}"
        )
        return e


async def _handle_create_game_request(log: ReasoningLog, _tweet_object):
    """Handle request to create a new game"""
    try:
        tweet_id = _tweet_object["tweet_id"]
        logger.info(
            f"[_handle_create_game_request] Beginning game creation process for tweet {tweet_id}"
        )

        # Try to acquire lock before processing
        if not _try_acquire_create_game_lock(tweet_id):
            logger.info(
                f"[_handle_create_game_request] Tweet {tweet_id} is locked by another process, skipping..."
            )
            return None, Exception(
                "Tweet is being processed by another instance"
            )

        # Get agent usernames
        agent_usernames, err = _get_agent_usernames(_tweet_object)
        if err:
            return None, err

        logger.info(
            f"[_handle_create_game_request] Initiating game creation with tweet_id={tweet_id} for {len(agent_usernames)} agents"
        )
        game_info, err = await sync2async(GameAPIClient.start_game)(
            tweet_id, agent_usernames, const.GAME_DURATION
        )
        if err:
            return None, err

        # Get wallet mapping
        wallet_mapping, err = _get_agent_wallet_mapping(game_info)
        if err:
            return None, err

        # Create and post wallet tweet
        wallet_tweet = _create_wallet_tweet(wallet_mapping)
        return game_info, await _post_wallet_tweet(log, tweet_id, wallet_tweet)

    except Exception as err:
        traceback.print_exc()
        logger.error(
            f"[_handle_create_game_request] Game creation failed for tweet {_tweet_object.get('tweet_id', 'unknown')}: {err}"
        )
        return {}, err


class CreateGameSubtask(ReplySubtaskBase):

    async def run(self) -> dict:
        tweet_id = self.tweet_info.tweet_object.tweet_id
        tweet_object = self.tweet_info.tweet_object.to_dict()

        if not is_create_game_pending(tweet_id):
            logger.info(
                f"[CreateGameSubtask.run] Game {tweet_id} is not in pending state, skipping"
            )
            raise Exception(
                f"Game {tweet_id} is not in pending state, skipping"
            )

        _mark_create_game_running(tweet_id)
        task_result, err = await _handle_create_game_request(
            self.log, tweet_object
        )
        if err is not None:
            logger.error(
                f"[CreateGameSubtask.run] Game creation failed for tweet {tweet_id}: {err}"
            )
            _mark_create_game_pending(tweet_id)
        else:
            logger.info(
                f"[CreateGameSubtask.run] Successfully created game for tweet {tweet_id}"
            )
            _mark_create_game_done(tweet_id)

        return {
            "tweet_id": tweet_id,
        }
