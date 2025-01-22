from . import constants as const
import logging

logger = logging.getLogger(__name__)

SECRET_TOKEN = const.API_SECRET_TOKEN
from fastapi import HTTPException, Header


from hashlib import sha256
from typing import Annotated

sha256_of_secret_token = sha256(SECRET_TOKEN.encode()).hexdigest()
logger.info(f"sha256_of_secret_token: {sha256_of_secret_token}")

async def verify_opencall_x_token(x_token: Annotated[str | None, Header()] = None):
    global sha256_of_secret_token

    if x_token != sha256_of_secret_token:
        raise HTTPException(status_code=401, detail="Unauthorized")

    return x_token

async def verify_x_token(x_token: Annotated[str | None, Header()] = None):
    if x_token != SECRET_TOKEN:
        raise HTTPException(status_code=401, detail="X-Token header invalid")

    return x_token

def verify_third_party_authorization_key(backend_url: str, headers: dict={}):
    async def wrapper(authorization: Annotated[str | None, Header()] = ""):
        if authorization != "":
            return True

        raise HTTPException(status_code=401, detail="Unauthorized")

    return wrapper