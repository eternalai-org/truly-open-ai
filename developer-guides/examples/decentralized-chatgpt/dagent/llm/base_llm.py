from dagent.utils import SimpleCacheMechanism
from dagent.models import InferenceResult
from typing import Optional
import uuid 

class AsyncChatCompletion(object):
    MAX_CACHE_ITEMS = 2048

    def __init__(self, *args, **kwargs):
        self._cache = SimpleCacheMechanism()

    def commit(self, result: InferenceResult):
        return self._cache.commit(result)

    def __call__(self, *args, **kwds) -> InferenceResult:
        raise NotImplementedError("This method should be implemented by the subclass")

    def get(self, id: str, default=None) -> Optional[InferenceResult]:
        return self._cache.get(id, default)

    def generate_uuid(self) -> str:
        return str(uuid.uuid4())