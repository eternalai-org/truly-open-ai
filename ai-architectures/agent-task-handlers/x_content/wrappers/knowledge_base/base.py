from typing import Any, List, Optional
from pydantic import BaseModel

from x_content.models import AgentKnowledgeBase


class RelatedInformation(BaseModel):
    content: str
    score: float
    reference: Optional[str] = None


class KnowledgeBase(BaseModel):
    base_url: str
    api_key: str
    kbs: List[AgentKnowledgeBase]
    default_top_k: int = 5
    similarity_threshold: float = 0.5

    def retrieve(
        self,
        query: str,
        top_k: int = None,
        threshold: float = None,
        *args,
        **kwargs: Any,
    ) -> List[RelatedInformation]:
        raise NotImplementedError(
            "retrieve method not implemented; cls: {}".format(
                self.__class__.__name__
            )
        )

    async def aretrieve(
        self,
        query: str,
        top_k: int = None,
        threshold: float = None,
        *args,
        **kwargs: Any,
    ) -> List[RelatedInformation]:
        raise NotImplementedError(
            "aretrieve method not implemented; cls: {}".format(
                self.__class__.__name__
            )
        )

    def get_kn_ids(self):
        return [x.kb_id for x in self.kbs]

    def get_info(self):
        return {
            "base_url": self.base_url,
            "api_key": self.api_key,
            "kbs": self.kbs,
            "default_top_k": self.default_top_k,
            "similarity_threshold": self.similarity_threshold,
        }
