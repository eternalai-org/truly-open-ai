from abc import abstractmethod, ABC
from typing import List
from x_content.constants import MissionChainState
from x_content.models import AgentKnowledgeBase, ReasoningLog
from x_content.llm import OpenAILLMBase
from x_content.cache import mission_state_handler
from x_content.wrappers import redis_wrapper
from x_content.toolcall.toolcall import ToolListWrapper, IToolCall
from .utils import a_move_state
import traceback

class MultiStepTaskBase(ABC):
    handler = mission_state_handler.MissionStateHandler(
        connection=redis_wrapper.reusable_redis_connection()
    )

    resumable = False

    def __init__(self, 
                 llm: OpenAILLMBase, 
                 kn_base: AgentKnowledgeBase,
                 toolcall: IToolCall = [],
                 *args, **kwargs
    ):
        super().__init__()

        self.llm: OpenAILLMBase = llm
        self.kn_base: AgentKnowledgeBase = kn_base
        self.toolcall: IToolCall = toolcall

    @abstractmethod
    async def process_task(self, log: ReasoningLog) -> ReasoningLog:
        raise NotImplementedError("process_task method not implemented; cls: {}".format(self.__class__.__name__))
    
    async def commit_log(self, log: ReasoningLog) -> ReasoningLog:
        return await self.handler.acommit(log)

    async def run(self, log: ReasoningLog) -> ReasoningLog:

        while True:
            try:
                log = await self.process_task(log) or log

            except Exception as err:
                traceback.print_exc()
                log = await a_move_state(log, MissionChainState.ERROR, f"Error: {err}")

            finally:
                await self.commit_log(log)

            if log.state in [MissionChainState.ERROR, MissionChainState.DONE]:
                break

        return log