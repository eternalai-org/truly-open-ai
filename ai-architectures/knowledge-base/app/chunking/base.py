from abc import ABC, abstractmethod
from typing import List, Union

class BaseChunker(ABC):
    @abstractmethod
    def chunks(self, text: Union[str, List[str]]) -> List[str]:
        return text if isinstance(text, list) else [text]