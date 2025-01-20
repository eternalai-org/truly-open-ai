from typing import Any, List, Optional
from langchain_openai import ChatOpenAI
from langchain.schema import BaseMessage, ChatResult

class OpenAILLMBase(ChatOpenAI):
    def __init__(self, *args, **kwargs):
        super(OpenAILLMBase, self).__init__(*args, **kwargs)

    async def agenerate(self,
        messages: List[BaseMessage],
        stop: Optional[List[str]] = None,
        **kwargs: Any,
    ) -> ChatResult:
        raise NotImplementedError("agenerate method not implemented; cls: {}".format(self.__class__.__name__))
    
    def get_info(self):
        return {
            "model": self.model_name,
            "temperature": self.temperature,
            "max_tokens": self.max_tokens,
            "top_p": 1.0,
            "presence_penalty": 0.1,
            "n": self.n,
            "logit_bias": None,
            "frequency_penalty": self.frequency_penalty,
            "seed": self.seed
        }