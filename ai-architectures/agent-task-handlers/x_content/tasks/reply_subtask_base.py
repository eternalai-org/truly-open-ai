from x_content.llm.base import OpenAILLMBase
from x_content.models import AgentKnowledgeBase, ReasoningLog
from x_content.wrappers.api.twitter_v2.models.objects import ExtendedTweetInfo


from abc import ABC, abstractmethod


class ReplySubtaskBase(ABC):

    def __init__(
        self,
        llm: OpenAILLMBase,
        kn_base: AgentKnowledgeBase,
        log: ReasoningLog,
        tweet_info: ExtendedTweetInfo,
        *args,
        **kwargs,
    ):
        super().__init__()

        self.llm: OpenAILLMBase = llm
        self.kn_base: AgentKnowledgeBase = kn_base

        self.log = log
        self.tweet_info = tweet_info

    @abstractmethod
    async def run(self) -> dict:
        raise NotImplementedError(
            "run method not implemented; cls: {}".format(
                self.__class__.__name__
            )
        )
