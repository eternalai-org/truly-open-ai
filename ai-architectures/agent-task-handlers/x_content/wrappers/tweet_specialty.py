import logging
from enum import Enum
from typing import List
from x_content import constants as const
from x_content.constants import AgentUsername
from x_content.wrappers import llm_tasks
from x_content.wrappers.api import twitter_v2
from x_content.wrappers.api.twitter_v2.main import get_tweet_info_from_tweet_id
from x_content.wrappers.api.twitter_v2.models.objects import ExtendedTweetInfo
from x_content.wrappers.postprocess import StringProcessor

logging.basicConfig(level=logging.INFO if not __debug__ else logging.DEBUG)
logger = logging.getLogger(__name__)


class TweetSpecialty(str, Enum):
    CREATE_GAME = "create_game"
    CREATE_GAME_SUBTREE = "create_game_subtree"
    TOKEN_ANALYSIS = "token_analysis"


def is_analyzing_token_tweet(tweet_info: ExtendedTweetInfo):
    try:
        result = llm_tasks.is_analyzing_token_conversation(
            tweet_info.tweet_object.full_text
        )
        return result
    except Exception as err:
        logger.error(
            f"[is_analyzing_token_tweet] Error analyzing conversation: {err}"
        )
        return None


def get_mentioned_usernames(tweet_info: ExtendedTweetInfo):
    """Mock function to get mentioned usernames.
    TODO:
    - Implement actual Twitter API call to get all mentioned usernames
    - Parse response and extract usernames from mentions
    - must not contain @nobullshit
    """
    # usernames = ["agent1", "agent2"]
    username_to_remove = AgentUsername.CRYPTOCOMIC_AI
    mentions = tweet_info.tweet_object.mentions
    tweet_info.tweet_object.mentions = [
        mention
        for mention in mentions
        if mention.username != username_to_remove
    ]
    usernames = [
        mention.username for mention in tweet_info.tweet_object.mentions
    ]
    return usernames


# Filter tweets that contain emoji game patterns like:
# Example 1: "🎮 Let's play a game! 🎲"
# Example 2: "🎯 Guess the number between 1-10 🎲"
# Example 3: "🎪 Riddle time! 🤔"
def is_create_game_tweet(tweet_info: ExtendedTweetInfo):
    tweet_object = tweet_info.tweet_object
    text = tweet_object.full_text
    agent_usernames = get_mentioned_usernames(tweet_info)

    has_two_or_more_usernames = len(agent_usernames) >= 2

    # contains_emoji = any(emoji in text for emoji in GAME_EMOJIS)
    emoji_count = sum(1 for emoji in const.GAME_EMOJIS if emoji in text)
    contains_two_emojis = emoji_count >= 2
    text = (
        StringProcessor(text)
        .remove_tags()
        .remove_mentions()
        .remove_urls()
        .remove_emojis()
        .strip_head_and_tail_white_string()
        .get_text()
    )
    # logger.info(f"[is_create_game_tweet] text after postprocess {text}")
    contains_keyword = any(
        keyword in text.lower() for keyword in const.GAME_KEYWORDS
    )

    return (
        contains_two_emojis and contains_keyword and has_two_or_more_usernames
    )


def is_create_game_tweet_id(tweet_id: str) -> bool:
    resp = twitter_v2.get_tweet_info_from_tweet_id(tweet_id)
    if resp.is_error():
        return False
    return is_create_game_tweet(resp.data.tweet_info)


def detect_tweet_specialties(
    tweet_info: ExtendedTweetInfo,
) -> List[TweetSpecialty]:
    if is_create_game_tweet(tweet_info):
        return [TweetSpecialty.CREATE_GAME]
    if is_create_game_tweet_id(tweet_info.conversation_id):
        return [TweetSpecialty.CREATE_GAME_SUBTREE]
    if is_analyzing_token_tweet(tweet_info):
        return [TweetSpecialty.TOKEN_ANALYSIS]
    return []


if __name__ == "__main__":
    resp = get_tweet_info_from_tweet_id("1882604450705231916")
    tweet_info = resp.data.tweet_info
    print(is_create_game_tweet(tweet_info))
