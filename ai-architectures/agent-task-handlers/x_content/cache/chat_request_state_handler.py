from x_content.cache.base_state_handler import StatusHandlerBase
from x_content.constants import MissionChainState
from x_content.models import ChatRequest
from redis import Redis
import json
from typing import Optional, List
import traceback
from pydantic_core import from_json
from x_content import constants as const


class ChatRequestStateHandler(StatusHandlerBase[ChatRequest]):
    BASE_KEY = const.REDIS_CHAT_REQUEST_BASE_KEY

    def __init__(self, connection: Redis) -> None:
        self.redis_client = connection
        self.cache_expiry = 3 * 24 * 3600  # Cache entries for 3 * 24 hours

    def commit(self, state: ChatRequest):
        key = f"{self.BASE_KEY}:{state.id}"
        jsons = json.dumps(state.model_dump())
        self.redis_client.setex(key, self.cache_expiry, jsons)
        return state

    async def acommit(self, state: ChatRequest):
        return self.commit(state)

    def get_undone(self) -> List[ChatRequest]:
        keys = self.redis_client.keys(f"{self.BASE_KEY}:*")
        states = []

        for key in keys:
            try:
                state = ChatRequest.model_validate(
                    from_json(self.redis_client.get(key).decode("utf-8"))
                )
            except ValueError:
                continue

            if not state.is_done() and not state.is_error():
                states.append(state)

        return states

    def get(
        self, _id: str, none_if_error: bool = False
    ) -> Optional[ChatRequest]:
        key = f"{self.BASE_KEY}:{_id}"
        json_state: Optional[bytes] = self.redis_client.get(key)

        if json_state is not None:
            try:
                state = ChatRequest.model_validate(
                    from_json(json_state.decode("utf-8"))
                )

            except ValueError as err:
                traceback.print_exc()

                if none_if_error:
                    return None

                # Fallback to error state if JSON deserialization fails
                state = ChatRequest(
                    state=MissionChainState.ERROR,
                    id=_id,
                    system_message=f"{_id} JSON deserialization failed",
                )
        else:
            if none_if_error:
                return None

            # Return an error state if the state doesn't exist in Redis
            state = ChatRequest(
                state=MissionChainState.ERROR,
                id=_id,
                system_message=f"{_id} Not found",
            )

        # Update the cache expiry time
        self.redis_client.expire(key, self.cache_expiry)
        return state
