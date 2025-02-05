import logging
import os

logger = logging.getLogger(__name__)


def get_and_warn(key, default_value=None):
    if key in os.environ:
        return os.environ[key]

    logger.warning(f"Environment variable {key} not found")
    return default_value
