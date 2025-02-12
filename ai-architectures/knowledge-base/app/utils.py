from hashlib import md5
from typing import Union
import asyncio
from functools import partial
from typing import Callable
from typing import Generator, AsyncGenerator
import tempfile
import os
from functools import wraps
from starlette.concurrency import run_in_threadpool
from app.models import EmbeddingModel, SimMetric
from asyncio import Semaphore as AsyncSemaphore


def get_content_checksum(data: Union[bytes, str]) -> str:
    if isinstance(data, str):
        _data = data.encode()
    else:
        _data = data

    return md5(_data, usedforsecurity=False).hexdigest()

def batching(data: Generator, batch_size = 1):
    for i in range(0, len(data), batch_size):
        yield data[i:i + batch_size]

async def async_batching(data: AsyncGenerator, batch_size=1):
    current_batch = []
    
    async for item in data:
        current_batch.append(item)
            
        if len(current_batch) == batch_size:
            yield current_batch
            current_batch = []
            
    if len(current_batch) > 0:
        yield current_batch

def get_hash(*items):
    return md5("".join(items).encode()).hexdigest()

def sync2async(sync_func: Callable):
    async def async_func(*args, **kwargs):
        return await run_in_threadpool(partial(sync_func, *args, **kwargs))
    return async_func

def limit_asyncio_concurrency(num_of_concurrent_calls: int):
    semaphore = AsyncSemaphore(num_of_concurrent_calls)

    def decorator(func: Callable):
        @wraps(func)
        async def wrapper(*args, **kwargs):
            async with semaphore:
                return await func(*args, **kwargs)                
        return wrapper
    return decorator

def random_payload(length: int) -> str:
    return os.urandom(length).hex()

def get_tmp_directory():
    return os.path.join(tempfile.gettempdir(), random_payload(20))

def is_async_func(func: Callable) -> bool:
    return asyncio.iscoroutinefunction(func)

def background_task_error_handle(handler: Callable):
    def decorator(func: Callable):
        @wraps(func)
        async def wrapper(*args, **kwargs):
            try:
                return await func(*args, **kwargs)
            except Exception as e:
                res = handler(*args, e, **kwargs)

                if is_async_func(handler):
                    return await res
      
        return wrapper
    return decorator


def estimate_ip_from_distance(distance, model_use: EmbeddingModel):
    if model_use.prefer_metric == SimMetric.COSINE:
        return 1.0 - distance

    if model_use.prefer_metric == SimMetric.L2:
        return 1.0 / (1.0 + distance)

    return distance

import aiofiles

async def iter_file(file_name: str):
    async with aiofiles.open(file_name, "rb") as f:
        while True:
            chunk = await f.read(1024 * 20)

            if not chunk:
                break

            yield chunk