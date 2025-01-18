from abc import ABC, abstractmethod
from typing import List, Union, Dict, Any

class VectorStoreBase(ABC):
    @abstractmethod
    def insert(self, data: List[Dict[str, Any]]):
        pass

    @abstractmethod
    def retrieve(self, data: List[Dict[str, Any]]):
        pass