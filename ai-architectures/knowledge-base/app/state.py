from .wrappers import redis_kit
from pydantic_core import from_json
from typing import Generic, List, TypeVar
from abc import ABC, abstractmethod
from .models import InsertInputSchema

_t = TypeVar('T')

class BaseModelHandler(Generic[_t], ABC):
    base_redis_key_prefix = f"model_{_t.__name__}:"
    default_expiry_time = 60 * 60 * 24 * 7

    def get(self, id: str) -> _t:
        return self.from_bytes(self.redis_client.get(self.key(id)))

    def insert(self, item: _t):
        return self.redis_client.set(
            self.key(self.id(item)), 
            self.to_bytes(item)
        )

    def delete(self, id: str):
        return self.redis_client.delete(self.key(id))
    
    def get_all(self) -> List[_t]:
        return [
            self.from_bytes(self.redis_client.get(key)) 
            for key in self.keys()
        ]

    @abstractmethod
    def to_bytes(self, item: _t) -> bytes:
        raise NotImplementedError

    def from_bytes(self, data: bytes) -> _t:
        raise NotImplementedError

    def key(self, id: str) -> str:
        return f"{self.base_redis_key_prefix}:{id}"

    def keys(self) -> List[str]:
        return self.redis_client.keys(f"{self.base_redis_key_prefix}*")

    def id(self, item: _t) -> str:
        return item.id if hasattr(item, 'id') else hash(self.to_bytes(item))

    def __init__(
        self, 
        redis_client=None
    ):
        super().__init__()
        self.redis_client = redis_client or redis_kit.reusable_redis_connection()

import json

class InsertionRequestHandler(BaseModelHandler[InsertInputSchema]):
    def to_bytes(self, item: InsertInputSchema) -> bytes:
        return json.dumps(item.model_dump()).encode('utf-8')

    def from_bytes(self, data: bytes) -> InsertInputSchema:
        return InsertInputSchema.model_validate(from_json(data))

from functools import lru_cache

@lru_cache(maxsize=1)
def get_insertion_request_handler():
    return InsertionRequestHandler()