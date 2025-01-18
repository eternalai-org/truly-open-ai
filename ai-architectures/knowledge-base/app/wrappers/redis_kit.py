import os
import redis
from typing import Union
from typing import Callable
import logging
from typing import Callable
import redis.commands
import pickle

logger = logging.getLogger(__name__)


def get_redis_client_connection() -> redis.Redis:
    return redis.Redis(
        host=os.getenv("REDIS_HOST", "localhost"),
        port=os.getenv("REDIS_PORT", "6379"),
        db=int(os.getenv("REDIS_DB", "0")),
        password=os.getenv("REDIS_PASSWORD", ""),
    )
    
def has_redis_connection() -> bool:
    try:
        redis_client = get_redis_client_connection()
        redis_client.ping()
        return True
    except:
        return False

__redis_connection = None


def reusable_redis_connection(strict=False) -> Union[redis.Redis, None]:
    global __redis_connection

    if __redis_connection is None:
        __redis_connection = get_redis_client_connection()

    # check if the connection is still alive
    still_alive = False
    for i in range(3):
        try:
            __redis_connection.ping()
            still_alive = True
            break
        except:
            __redis_connection = get_redis_client_connection()

    if not still_alive:
        if strict:
            raise Exception("Redis connection is not alive")

        return None

    return __redis_connection


class IRQueue(object):
    def __init__(self, collection_name: str):
        self.collection_name = collection_name
        self.redis_client = get_redis_client_connection()

    @property
    def qsize(self):
        return self.redis_client.llen(self.collection_name)

    def clear(self):
        self.redis_client.delete(self.collection_name)

    @property
    def is_empty(self):
        return self.qsize == 0


class RNormalQueue(IRQueue):

    def enqueue(self, data: Union[dict, str]) -> None:
        self.redis_client.rpush(self.collection_name, data)

    def dequeue(self, block=True) -> Union[dict, str]:
        if block:
            res = self.redis_client.blpop([self.collection_name])

            if res is not None:
                res = res[1].decode("utf-8")

        else:
            res = self.redis_client.lpop(self.collection_name)

            if res is not None:
                res = res.decode("utf-8")

        return res


class RPriorityQueue(IRQueue):
    DISTRIBUTED_DEQUEUE_SCRIPT = """
local item = redis.call('ZRANGE', KEYS[1], 0, 0)
if item[1] then
    redis.call('ZREM', KEYS[1], item[1])
    return item[1]
else
    return nil
end
"""

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.dequeue_script = self.redis_client.register_script(
            self.DISTRIBUTED_DEQUEUE_SCRIPT
        )

    def enqueue(self, data: Union[bytes, str], priority: int) -> None:
        self.redis_client.zadd(self.collection_name, {data: priority})

    def dequeue(self) -> Union[dict, str]:
        res = self.dequeue_script(keys=[self.collection_name])
        return res

    def peek(self):
        res = self.redis_client.zrange(self.collection_name, 0, 0)
        return res[0] if res else None

    @property
    def qsize(self):
        return self.redis_client.zcard(self.collection_name)


class RAdvancedPriorityQueue(RPriorityQueue):
    DISTRIBUTED_DEQUEUE_SCRIPT = """
local item = redis.call('ZRANGEBYSCORE', KEYS[1], ARGV[1], ARGV[2], 'LIMIT', ARGV[3], ARGV[4], 'WITHSCORES')
if item[1] then
    redis.call('ZREM', KEYS[1], item[1])
    return item
else
    return nil
end
"""

    def __init__(self, priority_factory: Callable, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.priority_factory = priority_factory

    def enqueue(self, data: str, priority: float = None) -> None:
        if priority is None:
            priority = self.priority_factory(data)

        self.redis_client.zadd(self.collection_name, {data: priority})

    def dequeue(
        self, min_score: float = float("-inf"), max_score: float = None
    ) -> tuple:
        if max_score is None:
            max_score = self.priority_factory(float("inf"))

        res = self.dequeue_script(
            keys=[self.collection_name], args=[min_score, max_score, 0, 1]
        )

        if res is not None:
            res[1] = float(res[1])
            res[0] = res[0].decode("utf-8")

        return res

    @property
    def qsize(self):
        return self.redis_client.zcard(self.collection_name)



def register_script(cli: redis.Redis, script) -> redis.commands.core.Script:
    return cli.register_script(script)

_reg = None
_atomic_script = """
local current_value = redis.call('GET', KEYS[1])
if current_value == false then
    redis.call('SET', KEYS[1], ARGV[1], 'EX', ARGV[2])
    return true
end
return false
"""

def atomic_check_and_set_flag(
    redis_client: redis.Redis, key: str, timeout: float, value: str = "1"
) -> bool:
    
    global _reg, _atomic_script

    if _reg is None:
        _reg = register_script(redis_client, _atomic_script)

    return _reg(keys=[key], args=[value, int(timeout)])


def distributed_scheduling_job(interval_seconds: float, strict=False):
    def decorator(func: Callable):
        def wrapper(*args, **kwargs):
            fn_name = func.__name__
            fn_module = func.__module__

            redis_cache_key = f"{fn_module}.{fn_name}:executed"

            redis_client = reusable_redis_connection(strict)
            do_execute = True

            if redis_client is not None:
                do_execute = atomic_check_and_set_flag(
                    redis_client, 
                    redis_cache_key, 
                    interval_seconds
                )

            redis_cache_key_result = f"{fn_module}.{fn_name}:result"
            if not do_execute:

                if redis_client is not None:
                    pickle_str = redis_client.get(redis_cache_key_result)

                    if pickle_str is not None:
                        return pickle.loads(pickle_str)

                return None

            res = func(*args, **kwargs)

            if redis_client is not None:
                pickle_str = pickle.dumps(res)
                redis_client.set(redis_cache_key_result, pickle_str, ex=interval_seconds)

            return res

        return wrapper

    return decorator

import asyncio

def cache_for(interval_seconds: float):
    
    def get_parameters_hash(*args, **kwargs):
        return hash((*args, *sorted(kwargs.items())))
    
    def decorator(func: Callable):
        fn_name = func.__name__
        fn_module = func.__module__
        redis_cache_key = f"{fn_module}.{fn_name}:result"

        def sync_wrapper(*args, **kwargs):
            nonlocal redis_cache_key

            key = f"{redis_cache_key}:{get_parameters_hash(*args, **kwargs)}"
            redis_client = reusable_redis_connection()

            if redis_client is not None:
                pickle_str = redis_client.get(key)

                if pickle_str is not None:
                    logger.info(f"Cache hit for {key}; call {fn_module}.{fn_name}(args={args}, kwargs={kwargs})")
                    return pickle.loads(pickle_str)

            res = func(*args, **kwargs)

            if redis_client is not None:
                pickle_str = pickle.dumps(res)
                redis_client.set(key, pickle_str, ex=interval_seconds)

            return res
        
        async def async_wrapper(*args, **kwargs):
            nonlocal redis_cache_key

            key = f"{redis_cache_key}:{get_parameters_hash(*args, **kwargs)}"
            redis_client = reusable_redis_connection()

            if redis_client is not None:
                pickle_str = redis_client.get(key)

                if pickle_str is not None:
                    logger.info(f"Cache hit for {key}; call {fn_module}.{fn_name}(args={args}, kwargs={kwargs})")
                    return pickle.loads(pickle_str)

            res = await func(*args, **kwargs)

            if redis_client is not None:
                pickle_str = pickle.dumps(res)
                redis_client.set(key, pickle_str, ex=interval_seconds)

            return res

        if asyncio.iscoroutinefunction(func):
            return async_wrapper

        return sync_wrapper

    return decorator