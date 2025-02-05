import logging

logger = logging.getLogger(__name__)

import os

def get_and_warn(key, default_value = None):
    if key in os.environ:
        return os.environ[key]

    logger.warning(f"Environment variable {key} not found")
    return default_value

GCP_MEME_IMAGES_FOLDER = get_and_warn("GCP_MEME_IMAGES_FOLDER")
GCP_BUCKET_NAME = get_and_warn("GCP_BUCKET_NAME")
TWITTER_API_URL = get_and_warn("TWITTER_API_URL", "").rstrip("/")
TELEGRAM_API_KEY = get_and_warn("TELEGRAM_API_KEY")

REDIS_HOST = get_and_warn("REDIS_HOST", "localhost")
REDIS_PORT: int = get_and_warn("REDIS_PORT") or 6379

if isinstance(REDIS_PORT, str):
    REDIS_PORT = int(REDIS_PORT)

REDIS_PASSWORD = get_and_warn("REDIS_PASSWORD", "")
REDIS_DB = get_and_warn("REDIS_DB", "0")

FETCH_NEWS_INTERVAL_SECONDS = get_and_warn("FETCH_NEWS_INTERVAL_SECONDS") or 60 * 60 * 2

if isinstance(FETCH_NEWS_INTERVAL_SECONDS, str):
    FETCH_NEWS_INTERVAL_SECONDS = int(FETCH_NEWS_INTERVAL_SECONDS)

TWITTER_API_KEY = get_and_warn("TWITTER_API_KEY")
SELF_HOSTED_LLAMA_405B_MODEL_IDENTITY = get_and_warn("SELF_HOSTED_LLAMA_405B_MODEL_IDENTITY")
SELF_HOSTED_MODEL_IDENTITY = get_and_warn("SELF_HOSTED_MODEL_IDENTITY")
BING_SEARCH_API_KEY = get_and_warn("BING_SEARCH_API_KEY")
RAG_API = get_and_warn("RAG_API", "").rstrip() 
IS_MOCKING_ACTION: bool = get_and_warn("IS_MOCKING_ACTION") or True

if isinstance(IS_MOCKING_ACTION, str):
    IS_MOCKING_ACTION = IS_MOCKING_ACTION.lower() in ["true", "1"]

RAG_SECRET_TOKEN = get_and_warn("RAG_SECRET_TOKEN")
SELF_HOSTED_LLAMA_405B_URL = get_and_warn("SELF_HOSTED_LLAMA_405B_URL", "").rstrip("/")

TELEGRAM_ROOM = get_and_warn("TELEGRAM_ROOM", "-4622942390")
TELEGRAM_BOTNAME = get_and_warn("TELEGRAM_BOTNAME")
TELEGRAM_ALERT_ROOM = get_and_warn("TELEGRAM_ALERT_ROOM", "1002454447988")
DISABLED_TELEGRAM_USERS = get_and_warn("DISABLED_TELEGRAM_USERS", "")
TELEGRAM_IO_NOTIFICATIONS_ROOM = get_and_warn("TELEGRAM_IO_NOTIFICATIONS_ROOM", "-4605594188")
TELEGRAM_TASK_IO_NOTIFICATION_ROOM = get_and_warn("TELEGRAM_TASK_IO_NOTIFICATION_ROOM", "-4605594188")
TELEGRAM_MESSAGE_LENGTH_LIMIT = 4096  # characters
TELEGRAM_MESSAGE_LIST_REDIS_KEY = "telegram_message_list"

DISABLE_JUNK_NOTIFICATIONS = get_and_warn("DISABLE_JUNK_NOTIFICATIONS") or False

if isinstance(DISABLE_JUNK_NOTIFICATIONS, str):
    DISABLE_JUNK_NOTIFICATIONS = DISABLE_JUNK_NOTIFICATIONS.lower() in ["true", "1"]

SELF_HOSTED_LLAMA_URL = get_and_warn("SELF_HOSTED_LLAMA_URL", "").rstrip("/")
DISABLE_FETCH_NEWS: bool = get_and_warn("DISABLE_FETCH_NEWS") or False 

if isinstance(DISABLE_FETCH_NEWS, str):
    DISABLE_FETCH_NEWS = DISABLE_FETCH_NEWS.lower() in ["true", "1"]

SELF_HOSTED_LLAMA_API_KEY = get_and_warn("SELF_HOSTED_LLAMA_API_KEY")

VISION_API_KEY = get_and_warn("VISION_API_KEY")
VISION_API_URL = get_and_warn("VISION_API_URL", "").rstrip("/")
VISION_API_MODEL = get_and_warn("VISION_API_MODEL")


DISABLE_LOG_FUNCTION_CALL: bool = get_and_warn("DISABLE_LOG_FUNCTION_CALL") or False

if isinstance(DISABLE_LOG_FUNCTION_CALL, str):
    DISABLE_LOG_FUNCTION_CALL = DISABLE_LOG_FUNCTION_CALL.lower() in ["true", "1"]


SELF_HOSTED_TEXT_TO_IMAGE_URL = get_and_warn("SELF_HOSTED_TEXT_TO_IMAGE_URL", "").rstrip("/")

MAX_TEXT_LENGTH = 10000
MIN_TEXT_LENGTH_TO_SUMMARIZE = 100

SELF_HOSTED_HERMES_70B_URL = get_and_warn("SELF_HOSTED_HERMES_70B_URL", "").rstrip("/")
SELF_HOSTED_HERMES_70B_MODEL_IDENTITY = get_and_warn("SELF_HOSTED_HERMES_70B_MODEL_IDENTITY")
 
DEFAULT_MAX_OUTPUT_TOKENS = 1024
DEFAULT_TEMPERATURE = 0.7
