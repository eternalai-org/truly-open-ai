from abc import ABC, abstractmethod
from typing import Generic, List, Optional, TypeVar
from x_content.models import AutoAgentTask

T = TypeVar("T", bound=AutoAgentTask)


class StatusHandlerBase(ABC, Generic[T]):

    @abstractmethod
    def commit(self, state: T):
        raise NotImplementedError(
            "commit method not implemented; cls: {}".format(
                self.__class__.__name__
            )
        )

    @abstractmethod
    async def acommit(self, state: T):
        raise NotImplementedError(
            "acommit method not implemented; cls: {}".format(
                self.__class__.__name__
            )
        )

    @abstractmethod
    def get_undone(self) -> List[T]:
        raise NotImplementedError(
            "get_undone method not implemented; cls: {}".format(
                self.__class__.__name__
            )
        )

    @abstractmethod
    def get(self, _id: str, none_if_error: bool = False) -> Optional[T]:
        raise NotImplementedError(
            "get method not implemented; cls: {}".format(
                self.__class__.__name__
            )
        )
