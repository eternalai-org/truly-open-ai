from x_content.constants.utils import get_and_warn


REDIS_HOST = get_and_warn("REDIS_HOST", "localhost")
REDIS_PORT: int = get_and_warn("REDIS_PORT") or 6379

if isinstance(REDIS_PORT, str):
    REDIS_PORT = int(REDIS_PORT)

REDIS_PASSWORD = get_and_warn("REDIS_PASSWORD", "")
REDIS_DB = get_and_warn("REDIS_DB", "0")
