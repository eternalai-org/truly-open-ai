from datetime import datetime
from typing import List, Optional
from pydantic import BaseModel

from x_content.wrappers.api.twitter_v2.models.objects import ExtendedTweetInfo, TweetInfo, TweetObject, TwitterUserObject
from typing import TypeVar, Generic

T = TypeVar('T')
class Response(BaseModel, Generic[T]):
    error: Optional[str] = None
    data: Optional[T] = None

    def is_error(self):
        return self.error is not None


class TweetsDto(BaseModel):
    tweets: List[TweetObject]


class SearchTweetDto(BaseModel):
    optimized_query: str
    tweets: List[TweetObject]


class TweetInfosDto(BaseModel):
    tweet_infos: List[TweetInfo]


class ExtendedTweetInfosDto(BaseModel):
    tweet_infos: List[ExtendedTweetInfo]


class ExtendedTweetInfoDto(BaseModel):
    tweet_info: ExtendedTweetInfo


class SearchRecentTweetsDto(BaseModel):
    search_query: str
    tweets: List[TweetObject]


class TwitterUserObjectDto(BaseModel):
    user: TwitterUserObject


class UsernamesDto(BaseModel):
    usernames: List[str]


class StructuredInformationDto(BaseModel):
    knowledge: List[str]
    news: List[str]    

    def to_dict(self):
        return {
            "knowledge": self.knowledge,
            "news": self.news,
        }

class TwitterUsersDto(BaseModel):
    users: List[TwitterUserObject]


class GetRecentOwnTweetDto(BaseModel):
    tweets: List[TweetObject]
    tweet_count: int
    search_start: datetime
    search_end: datetime


class GenerateActionDto(BaseModel):
    success: bool


class InscribeTweetByIdDto(BaseModel):
    success: bool
    metadata: dict

