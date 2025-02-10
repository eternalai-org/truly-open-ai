from x_content.cache.chat_request_state_handler import ChatRequestStateHandler
from x_content.tasks.utils import a_move_state
from x_content.tasks.base import BaseTask, ChatTaskBase, MultiStepTaskBase
from json_repair import repair_json
from x_content.utils import parse_knowledge_ids
from x_content.wrappers import redis_wrapper
from x_content.wrappers.conversation import get_llm_result_by_model_name
from x_content.wrappers.knowledge_base.local import KBStore
from x_content.models import AgentKnowledgeBase, ChatRequest, ReasoningLog
from x_content.models import ReasoningLog, MissionChainState
from x_content.llm.base import OnchainInferResult
from x_content import constants as const
from x_content.services import chat_service


class ChatV2(ChatTaskBase):
    resumable = True

    async def process_task(self, request: ChatRequest) -> ReasoningLog:
        if request.state == MissionChainState.NEW:
            request.chat_result = await chat_service.get_chat(
                request, self.llm, self.kn_base
            )

            return await a_move_state(
                request, MissionChainState.DONE, "Task completed"
            )

        return request
