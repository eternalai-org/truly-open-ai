import logging
import os

logger = logging.getLogger(__name__)

def get_env_and_warning(key: str, default=None):
    if key not in os.environ:
        logger.warning(f"{key} not found in environment")
        return default
    
    return os.getenv(key)

# TODO: break this file into smaller parts and assign to each package
ETERNAL_X_API = get_env_and_warning("ETERNAL_X_API", "").rstrip("/")
ETERNAL_X_API_APIKEY = get_env_and_warning("ETERNAL_X_API_APIKEY")
IS_SANDBOX = get_env_and_warning("IS_SANDBOX", "0") == "1"

ETERNALAI_URL = get_env_and_warning("ETERNALAI_URL", "").rstrip("/")
ETERNALAI_API_KEY = get_env_and_warning("ETERNALAI_API_KEY") 

ETERNAL_CHAIN_ID = os.getenv("ETERNAL_CHAIN_ID", "45762")
ETERNAL_MODEL_NAME = os.getenv("ETERNAL_MODEL_NAME", "unsloth/Llama-3.3-70B-Instruct-bnb-4bit")

# for trading, not available in the current version
CHAIN_ID=None
CONTRACT_ID=None 

AUTO_SERVICE_SLEEP_TIME = 10

DEFAULT_TOP_K = 3
DEFAULT_BIO_MAX_LENGTH = 20
DEFAULT_LORE_MAX_LENGTH = 20
DEFAULT_KNOWLEDGE_MAX_LENGTH = 30
DEFAULT_EXAMPLE_POSTS_MAX_LENGTH = 15
DEFAULT_INTERESTED_TOPICS_MAX_LENGTH = 10

APP_NAME = "dagent"