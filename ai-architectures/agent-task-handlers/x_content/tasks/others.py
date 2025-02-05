from x_content.constants import MissionChainState
from .base import MultiStepTaskBase
from .utils import a_move_state
from x_content.models import ReasoningLog


class FallbackTask(MultiStepTaskBase):

    async def process_task(self, log: ReasoningLog) -> ReasoningLog:
        return await a_move_state(
            log,
            MissionChainState.ERROR,
            "No handler found to process the task or task skipped due to something else! ehe",
        )
