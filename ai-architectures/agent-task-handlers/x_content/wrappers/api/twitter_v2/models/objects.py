from enum import Enum
from typing import List, Optional, Union, Dict
from pydantic import BaseModel, model_validator

import json
from x_content.wrappers.knowledge_base.base import KnowledgeBase


class TweetObject(BaseModel):
    tweet_id: str
    twitter_id: str
    twitter_username: str
    like_count: Optional[int] = 0
    retweet_count: Optional[int] = 0
    reply_count: Optional[int] = 0
    impression_count: Optional[int] = 0
    full_text: str
    posted_at: str
    media: List[str] = []
    reference: List[Dict[str, str]] = []  # TODO: change this

    @model_validator(mode="before")
    def validate_ref(cls, data: Dict[str, str]) -> Dict[str, str]:
        if "reference" not in data or data["reference"] is None:
            data["reference"] = []
        return data

    def __str__(self) -> str:
        return self.__repr__()

    def to_dict(self):
        """Serialize the object to a dictionary."""
        return {
            "tweet_id": self.tweet_id,
            "twitter_id": self.twitter_id,
            "twitter_username": self.twitter_username,
            "full_text": self.full_text,
            "posted_at": self.posted_at,
            "like_count": self.like_count,
            "retweet_count": self.retweet_count,
            "reply_count": self.reply_count,
            "impression_count": self.impression_count,
            "media": self.media,
            "reference": self.reference,
        }

    @classmethod
    def from_dict(cls, data: dict):
        """Deserialize a dictionary to an object."""
        return cls(
            tweet_id=data.get("tweet_id"),
            twitter_id=data.get("twitter_id"),
            twitter_username=data.get("twitter_username"),
            full_text=data.get("full_text"),
            posted_at=data.get("posted_at"),
            like_count=data.get("like_count", 0),
            retweet_count=data.get("retweet_count", 0),
            reply_count=data.get("reply_count", 0),
            impression_count=data.get("impression_count", 0),
            media=data.get("media", []),
            reference=data.get("reference", []),
        )

    def __repr__(self) -> str:
        return json.dumps(
            {
                "tweet_id": self.tweet_id,
                "twitter_username": self.twitter_username,
                "impression_count": self.impression_count,
                "posted_at": self.posted_at,
                "reply_count": self.reply_count,
                "retweet_count": self.retweet_count,
                "like_count": self.like_count,
                "full_text": self.full_text,
            }
        )


class TweetInfo(BaseModel):
    tweet_object: TweetObject
    parent_tweet_id: Optional[str] = None
    conversation_id: Optional[str] = None

    def to_dict(self):
        return {
            "tweet_object": self.tweet_object.to_dict(),
            "parent_tweet_id": self.parent_tweet_id,
            "conversation_id": self.conversation_id,
        }


class MentionData(BaseModel):
    start: int
    end: int
    username: str

    def to_dict(self):
        return {
            "start": self.start,
            "end": self.end,
            "username": self.username,
        }

    @classmethod
    def from_dict(cls, data: dict):
        """Deserialize a dictionary to an object."""
        return cls(
            start=data.get("start"),
            end=data.get("end"),
            username=data.get("username"),
        )


class ExtendedTweetObject(TweetObject):
    image_urls: List[str] = []
    mentions: List[MentionData] = []

    def to_dict(self):
        obj = super().to_dict()
        obj.update(
            {
                "image_urls": self.image_urls,
                "mentions": [x.to_dict() for x in self.mentions],
            }
        )
        return obj


class ExtendedTweetInfo(BaseModel):
    tweet_object: ExtendedTweetObject
    parent_tweet_id: Optional[str] = None
    conversation_id: str

    def to_dict(self):
        return {
            "tweet_object": self.tweet_object.to_dict(),
            "parent_tweet_id": self.parent_tweet_id,
            "conversation_id": self.conversation_id,
        }


class TwitterNews(BaseModel):
    reporter: str
    content: Optional[str] = None
    time_left: str
    time_right: str

    def to_dict(self):
        return {
            "reporter": self.reporter,
            "content": self.content,
            "time_left": self.time_left,
            "time_right": self.time_right,
        }

    def __str__(self) -> str:
        return self.__repr__()

    def __repr__(self) -> str:
        return json.dumps(self.to_dict())


class TwitterUserObject(BaseModel):
    twitter_id: str
    username: str
    name: str
    followings_count: int
    followers_count: int
    is_blue_verified: bool

    def to_dict(self):
        return {
            "username": self.username,
            # "name": self.name,
            "followers_count": int(self.followers_count),
            "followings_count": int(self.followings_count),
            "is_blue_verified": self.is_blue_verified,
        }

    def __str__(self) -> str:
        return self.__repr__()

    def __repr__(self) -> str:
        return json.dumps(self.to_dict())


class TwitterRequestAuthorization(BaseModel):
    twitter_id: str
    twitter_username: str
    request_id: str
    ref_id: str
    chain_id: Union[str, int]
    agent_contract_id: str
    knowledge_id: Optional[str] = ""
    task: Optional[str] = "not_defined"
    toolset: Optional[str] = "not_defined"
    kn_base: KnowledgeBase
    prompt: Optional[str] = ""

    def to_dict(self):
        return {
            "twitter_id": self.twitter_id,
            "twitter_username": self.twitter_username,
            "request_id": self.request_id,
            "ref_id": self.ref_id,
            "chain_id": self.chain_id,
            "agent_contract_id": self.agent_contract_id,
            "knowledge_id": self.knowledge_id,
        }

    def __str__(self) -> str:
        return self.__repr__()

    def __repr__(self) -> str:
        return json.dumps(self.to_dict())


class TweetType(str, Enum):
    POST = "post"
    REPLY = "reply"


class StructuredInformation(BaseModel):
    knowledge: List[str]
    news: List[str]

    def to_dict(self):
        return {
            "knowledge": self.knowledge,
            "news": self.news,
        }
