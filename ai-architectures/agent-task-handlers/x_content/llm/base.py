from langchain_openai import ChatOpenAI
from typing import Any, List, Optional
from langchain.schema import BaseMessage, ChatResult


class OnchainInferResult(ChatResult):
    tx_hash: str = None
    receipt: str = None


class OpenAILLMBase(ChatOpenAI):

    def __init__(self, *args, **kwargs):
        super(OpenAILLMBase, self).__init__(*args, **kwargs)

    def generate(
        self,
        messages: List[BaseMessage],
        stop: Optional[List[str]] = None,
        **kwargs: Any,
    ) -> OnchainInferResult:
        raise NotImplementedError(
            "generate method not implemented; cls: {}".format(
                self.__class__.__name__
            )
        )

    async def agenerate(
        self,
        messages: List[BaseMessage],
        stop: Optional[List[str]] = None,
        **kwargs: Any,
    ) -> OnchainInferResult:
        raise NotImplementedError(
            "agenerate method not implemented; cls: {}".format(
                self.__class__.__name__
            )
        )

    def get_info(self):
        raise NotImplementedError(
            "get_info method not implemented; cls: {}".format(
                self.__class__.__name__
            )
        )
