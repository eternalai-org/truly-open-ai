from enum import Enum
from x_content.wrappers.constants import *
from x_content.wrappers.constants_game import *
import logging 

logger = logging.getLogger(__name__)

class AgentUsername(str, Enum):
    CRYPTOCOMIC_AI = "CryptoComic AI"
    TWEETER_NAME_CRYPTOCOMIC_AI = "YorelisDiaz1"
    
class ModelName(str, Enum):
    INTELLECT_10B = "PrimeIntellect/INTELLECT-1-Instruct"

class HTTPMethod(str, Enum):
    GET = "GET"
    POST = "POST"

class ToolSet(str, Enum):
    DEFAULT = "default"
    NOT_SPECIFIED = "not_specified"
    REPLY_NON_MENTIONS = "reply_non_mentions"
    REPLY_NON_MENTIONS_TRADITIONAL = "reply_non_mentions_traditional"
    FOLLOW = "follow"
    POST = "post"
    QUOTE_TWEET = "quote_tweet"
    TRADING = "trading"
    INSCRIBE_TWEET = "inscribe_tweet"
    ISSUE_TOKEN = "issue_token"
    INSCRIBE_REPLY = "inscribe_reply"

class AgentTask(str, Enum):
    REACT_AGENT = "react_agent"
    POST_V2 = "post_search_v2"
    POST_V3 = "post_search_v3"
    REPLY = "reply"
    QUOTE_TWEET = "quote_tweet"
    SHADOW_REPLY = "shadow_reply"
    TRADING = "trading"
    POST_TWEET = "post_tweet"
    POST_SEARCH = "post_search"
    CREATE_GAME = "create_game"
    JUDGE_GAME = "judge_game"
    DEFAULT = "default"

class MissionChainState(str, Enum):
    NEW = "new"
    RUNNING = "running"
    DONE = "done"
    ERROR = "error"

BACKEND_LLM_API = get_and_warn("BACKEND_LLM_API")
USE_RQ = get_and_warn("USE_RQ") or True

if isinstance(USE_RQ, str):
    USE_RQ = USE_RQ.lower() in ["true", "1"]

LLM_MODE = get_and_warn("LLM_MODE")
KN_BASE_MODE = get_and_warn("KN_BASE_MODE")
OPENAI_API_KEY = get_and_warn("OPENAI_API_KEY")
BACKEND_LLM_TOKEN = get_and_warn("BACKEND_LLM_TOKEN")


import json

SELF_HOSTED_MODELS = get_and_warn("SELF_HOSTED_MODELS")

if SELF_HOSTED_MODELS is not None and isinstance(SELF_HOSTED_MODELS, str):
    SELF_HOSTED_MODELS = json.loads(SELF_HOSTED_MODELS)

REDIS_LOG_BASE_KEY = get_and_warn("REDIS_LOG_BASE_KEY", "ReactAgentReasoningLog-state-1")
BACKEND_AUTH_TOKEN = get_and_warn("BACKEND_AUTH_TOKEN")
GCS_TWIN_BUCKET = get_and_warn("GCS_TWIN_BUCKET")
APP_ENV = get_and_warn("APP_ENV")
RUN_SERVICE_V2: bool = get_and_warn("RUN_SERVICE_V2") or False

if isinstance(RUN_SERVICE_V2, str):
    RUN_SERVICE_V2 = RUN_SERVICE_V2.lower() in ["true", "1"]

BACKEND_API = get_and_warn("BACKEND_API")
BACKEND_LLM_MODEL = get_and_warn("BACKEND_LLM_MODEL")

API_SECRET_TOKEN = get_and_warn("API_SECRET_TOKEN", 'supersecret')
SERVER_HOST: str = get_and_warn("SERVER_HOST", "0.0.0.0")
SERVER_PORT: int = get_and_warn("SERVER_PORT") or 8000

if isinstance(SERVER_PORT, str):
    SERVER_PORT = int(SERVER_PORT)

MINIMUM_POST_LENGTH = 32
DEFAULT_REACT_MAX_STEPS = 10
NUM_OF_TWEETS_TO_POST = 1

CREATE_GAME_PREFIX_PATTERN = "[CREATE_GAME]" # TODO: To be defined

REACT_MODELS_BLACKLIST = [
    ModelName.INTELLECT_10B,
]

REPLY_MODELS_BLACKLIST = [
    ModelName.INTELLECT_10B,
]

ALL_BLACKLIST = [
    "itsmechaseb",
]

GAME_TASKS_WHITELIST = [
    AgentUsername.CRYPTOCOMIC_AI,
]
