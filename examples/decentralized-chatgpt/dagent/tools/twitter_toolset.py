from typing import List
from dagent.models import Tool, ToolParam, ToolParamDtype
from dagent.registry import RegistryCategory, register_decorator
from . base_toolset import Toolset
from . import functional

@register_decorator(RegistryCategory.ToolSet)
class TwitterToolset(Toolset):
    TOOLSET_NAME = "twitter"
    PURPOSE = "to interact with Twitter API"

    TOOLS: List[Tool] = [
        Tool(
            name="get_user_info_by_username",
            description="Get user info by username",
            param_spec=[
                ToolParam(
                    name="username",
                    dtype=ToolParamDtype.STRING,
                    description="Twitter username to get info"
                )
            ],
            executor=functional.get_user_info_by_username
        ),
        Tool(
            name="get_engaged_tweets_by_topic",
            description="Get engaged tweets by topic",
            param_spec=[
                ToolParam(
                    name="topic",
                    dtype=ToolParamDtype.STRING,
                    description="Topic to search"
                )
            ],
            executor=functional.get_engaged_tweets_by_topic
        ),
        Tool(
            name="find_user",
            description="Find a twitter user",
            param_spec=[
                ToolParam(
                    name="query",
                    dtype=ToolParamDtype.STRING,
                    description="Query to search"
                )
            ],
            executor=functional.find_user
        ),
        Tool(
            name="get_recent_mentioned_tweets",
            description="Get recent mentioned tweets of a specific user",
            param_spec=[
                ToolParam(
                    name="username",
                    dtype=ToolParamDtype.STRING,
                    description="Username to search"
                )
            ],
            executor=functional.get_recent_mentioned_tweets
        ),
        Tool(
            name="get_tweets_by_username",
            description="Get the most recent tweets by username",
            param_spec=[
                ToolParam(
                    name="username",
                    dtype=ToolParamDtype.STRING,
                    description="Username to search"
                )
            ],
            executor=functional.get_tweets_by_username
        ),
        Tool(
            name="get_following_users_by_username",
            description="Get following users of a twitter account",
            param_spec=[
                ToolParam(
                    name="username",
                    dtype=ToolParamDtype.STRING,
                    description="Username to search"
                )
            ],
            executor=functional.get_following_users_by_username
        ),
        Tool(
            name="follow",
            description="Follow a twitter user",
            param_spec=[
                ToolParam(
                    name="username",
                    dtype=ToolParamDtype.STRING,
                    description="Username to follow"
                )
            ],
            executor=lambda username: functional.follow(username)
        ),
        Tool(
            name="reply",
            description="Post a reply to a tweet",
            param_spec=[
                ToolParam(
                    name="tweet_id",
                    dtype=ToolParamDtype.STRING,
                    description="Tweet ID to reply"
                ),
                ToolParam(
                    name="reply_content",
                    dtype=ToolParamDtype.STRING,
                    description="Reply content to post"
                )
            ],
            executor=lambda tweet_id, reply_content: functional.reply(tweet_id, reply_content)
        ),
        Tool(
            name="quote_tweet",
            description="Quote a tweet",
            param_spec=[
                ToolParam(
                    name="tweet_id",
                    dtype=ToolParamDtype.STRING,
                    description="Tweet ID to quote"
                ),
                ToolParam(
                    name="comment",
                    dtype=ToolParamDtype.STRING,
                    description="Content to post"
                )
            ],
            executor=lambda tweet_id, comment: functional.quote_tweet(tweet_id, comment)
        ),
        Tool(
            name="tweet",
            description="Post a tweet",
            param_spec=[
                ToolParam(
                    name="content",
                    dtype=ToolParamDtype.STRING,
                    description="Content to post"
                )
            ],
            executor=lambda content: functional.tweet(content)
        )
    ]
