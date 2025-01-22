from functools import partial
import json
import requests
import httpx
import time
import traceback
from typing import Callable
import asyncio
from typing import Union

import logging
logger = logging.getLogger(__name__)

def retry(func: Callable, max_retry = 5, first_interval = 10, interval_multiply = 1):
    def sync_wrapper(*args, **kwargs):
        interval = first_interval
        for iter in range(max_retry + 1):
            try:
                result = func(*args, **kwargs)
                return result
            except Exception as err:
                traceback.print_exc()
                logger.error(f"Function {func.__name__} failed with error '{err}'. Retry attempt {iter}/{max_retry}")
            time.sleep(interval)
            interval *= interval_multiply

        logger.error(f"Function {func.__name__} failed after all retry.")
        raise Exception(f"Function {func.__name__} failed after all retry.")
    
    async def async_wrapper(*args, **kwargs):
        interval = first_interval
        for iter in range(max_retry + 1):
            try:
                result = await func(*args, **kwargs)
                return result
            except Exception as err:
                traceback.print_exc()
                logger.error(f"Function {func.__name__} failed with error '{err}'. Retry attempt {iter}/{max_retry}")
            time.sleep(interval)
            interval *= interval_multiply

        logger.error(f"Function {func.__name__} failed after all retry.")
        raise Exception(f"Function {func.__name__} failed after all retry.")

    return async_wrapper if asyncio.iscoroutinefunction(func) else sync_wrapper


def helpful_raise_for_status(resp: requests.Response | httpx.Response):
    """
    More helpful raise for status that also log request info when request failed
    """
    try:
        resp.raise_for_status()
    except Exception as err:
        data = {
            "err": str(err),
            "url": str(resp.url),
        }

        try:
            data["resp"] = resp.json()
        except:
            pass

        logger.error(f"Http request failed: {json.dumps(data)}")
        raise err


# to fake the async to sync
from starlette.concurrency import run_in_threadpool

def sync2async(sync_func: Callable):
    async def async_func(*args, **kwargs):
        res = await run_in_threadpool(partial(sync_func, *args, **kwargs))
        if asyncio.iscoroutinefunction(sync_func):            
            res = await res
        return res
        
    return async_func

def get_response_content(response: requests.Response) -> Union[str, dict]:
    try:
        return response.json()
    except Exception as err:
        return response.text
