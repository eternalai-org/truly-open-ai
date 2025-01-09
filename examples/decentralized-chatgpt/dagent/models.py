from pydantic import BaseModel, Field
from typing import List, Dict, Optional, Callable, Any
from enum import Enum
import json 
import uuid
import time

def random_uuid() -> str:
    return str(uuid.uuid4().hex)

class Serializable(object):
    def __str__(self) -> str:
        return self.__repr__()

    def __repr__(self) -> str:
        return json.dumps(self.__dict__())

class ToolParamDtype(str, Enum):
    STRING = "string"
    NUMBER = "number"

class ToolParam(BaseModel):
    name: str
    default_value: Optional[str] = None
    dtype: ToolParamDtype
    description: str

class Tool(BaseModel):
    name: str
    description: str
    param_spec: List[ToolParam]
    executor: Callable
    
    def prototype(self):
        params_str = ', '.join([f"{param.name}: {param.dtype.value}" 
                                for param in self.param_spec])
        
        return f'{self.name}({params_str}) -> {ToolParamDtype.STRING.value}: Takes {len(self.param_spec)} parameters, {self.description}'

class ClassRegistration(BaseModel):
    name: str
    init_params: Optional[Dict[str, Any]] = {}

class InferenceState(str, Enum):
    EXECUTING = "executing"
    DONE = "done"
    ERROR = "error"

class ChainState(str, Enum):
    NEW = "new"
    RUNNING = "running"
    DONE = "done"
    ERROR = "error"

class Characteristic(BaseModel):
    bio: Optional[List[str]] = []
    lore: Optional[List[str]] = []
    knowledge: Optional[List[str]] = []
    interested_topics: Optional[List[str]] = []
    agent_personal_info: Optional[Dict[str, str]] = {}
    system_prompt: Optional[str] = None
    example_posts: Optional[List[str]] = []
    example_messages: Optional[List[str]] = []

class DAgentLog(BaseModel):
    # auto
    id: str = Field(default_factory=lambda: f"fun-{random_uuid()}")
    
    characteristic: Characteristic

    toolset_cfg: List[ClassRegistration]
    llm_cfg: ClassRegistration
    agent_builder_cfg: ClassRegistration
    character_builder_cfg: ClassRegistration

    # for response
    infer_receipt: Optional[str] = None
    state: ChainState = ChainState.NEW
    scratchpad: List[Dict[str, str]] = []

    system_message: str = "" # for error messages
    verbose: bool = True

    def is_done(self):
        return self.state == ChainState.DONE

    def is_error(self):
        return self.state == ChainState.ERROR

    def clone(self) -> dict:
        return dict(
            id=self.id,
            infer_receipt=self.infer_receipt,
            state=self.state,
            scratchpad=self.scratchpad,
            system_message=self.system_message,
            toolset_cfg=[e.model_dump() for e in self.toolset_cfg],
            llm_cfg=self.llm_cfg.model_dump(),
            agent_builder_cfg=self.agent_builder_cfg.model_dump(),
            character_builder_cfg=self.character_builder_cfg.model_dump(),
            characteristic=self.characteristic.model_dump()
        )

class Mission(BaseModel):
    system_reminder: str
    task: str

class OnChainData(BaseModel):
    assignment_addresses: List[str] = []
    submit_address: Optional[str] = None
    infer_tx: Optional[str] = None 
    propose_tx: Optional[str] = None
    seize_miner_tx: Optional[str] = None
    input_cid: Optional[str] = None
    output_cid: Optional[str] = None 

class DAgentResponse(BaseModel):
    content: str
    onchain_data: Optional[OnChainData] = None

class NonInteractiveDAgentLog(DAgentLog):
    mission: Mission
    
    def clone(self):
        return dict(
            **super().clone(),
            mission=self.mission.model_dump()
        )

class ChatSession(BaseModel):
    id: str = Field(default_factory=lambda: f"chat-{random_uuid()}")
    messages: List[Dict[str, str]] = []
    llm_cfg: ClassRegistration = ClassRegistration(**
        {
            "name": "EternalAIChatCompletion",
            "init_params": {
                "model_name": "neuralmagic/Meta-Llama-3.1-405B-Instruct-quantized.w4a16",
                "max_tokens": 1024,
                "model_kwargs": {},
                "temperature": 0.3,
                "max_retries": 2
            }
        }
    )

    toolsets_cfg: List[ClassRegistration] = []
    agent_builder_cfg: Optional[ClassRegistration] = None
    character_builder_cfg: Optional[ClassRegistration] = None

    last_execution: Optional[float] = time.time()

# TODO: there should be a cachable interface
class InferenceResult(Serializable):
    def __init__(self, id: str, state: InferenceState, result: Optional[str]=None, error: Optional[str]=None, onchain_data: Optional[OnChainData]=None):
        self.state = state
        self.result = result
        self.error = error    
        self.id = id  
        self.onchain_data = onchain_data

    def __dict__(self) -> dict:
        return {
            "id": self.id,
            "state": self.state,
            "result": self.result,
            "error": self.error,
            "tx_hash": self.tx_hash
        }

class TweetObject(Serializable):
    """Represents a tweet from Twitter."""

    def __init__(self, tweet_id, twitter_id, twitter_username, full_text, like_count=0, retweet_count=0, reply_count=0, impression_count=0, posted_at=None, **kwargs):
        """
        Initialize a new TweetObject.

        :param tweet_id: Unique identifier for the tweet.
        :param twitter_id: Unique identifier for the Twitter user.
        :param twitter_username: Username of the Twitter user.
        :param full_text: Full text of the tweet.
        :param like_count: Number of likes the tweet has received.
        :param retweet_count: Number of retweets the tweet has received.
        :param reply_count: Number of replies the tweet has received.
        :param impression_count: Number of times the tweet has been seen.
        :param posted_at: Timestamp of when the tweet was posted.
        """
        super().__init__()
        self.tweet_id = tweet_id
        self.twitter_id = twitter_id
        self.twitter_username = twitter_username
        self.like_count = like_count
        self.retweet_count = retweet_count
        self.reply_count = reply_count
        self.impression_count = impression_count
        self.full_text = full_text
        self.posted_at = posted_at

    def __dict__(self) -> dict:
        return {
            "tweet_id": self.tweet_id,
            "twitter_username": self.twitter_username,
            "impression_count": self.impression_count,
            "posted_at": self.posted_at,
            "reply_count": self.reply_count,
            "retweet_count": self.retweet_count,
            "like_count": self.like_count,
            "full_text": self.full_text
        }


class TwitterUserObject(Serializable):
    """Represents a Twitter user."""

    def __init__(self, twitter_id, twitter_username, name, followings_count = 0, followers_count = 0, is_blue_verified = False, followed=False, **kwargs):
        """
        Initialize a new TwitterUserObject.

        :param twitter_id: Unique identifier for the Twitter user.
        :param twitter_username: Username of the Twitter user.
        :param name: Display name of the Twitter user.
        :param followings_count: Number of users this user is following.
        :param followers_count: Number of followers this user has.
        :param is_blue_verified: Boolean indicating if the user is blue verified.
        :param followed: Boolean indicating if the user is followed by the current user.
        """
        super().__init__()
        self.twitter_id = twitter_id
        self.username = twitter_username
        self.name = name
        self.followings_count = followings_count
        self.followers_count = followers_count
        self.is_blue_verified = is_blue_verified
        self.followed = followed

    def __dict__(self) -> dict:
        return {
            "twitter_id": self.twitter_id,
            "username": self.username,
            "name": self.name,
            "followings_count": self.followings_count,
            "followers_count": self.followers_count,
            "is_blue_verified": self.is_blue_verified,
            "followed": self.followed
        }