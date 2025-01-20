from abc import ABC, abstractmethod

class BaseQueryOptimizer(ABC):
    @abstractmethod
    def optimize(self, query: str) -> str:
        return query