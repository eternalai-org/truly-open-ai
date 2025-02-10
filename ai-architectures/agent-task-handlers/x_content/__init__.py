import os
import logging

logger = logging.getLogger(__name__)

from dotenv import load_dotenv

if not load_dotenv():
    logger.warning("No .env file found")

__version__ = os.getenv("GIT_TAG", "v0.3.62")

from . import api
