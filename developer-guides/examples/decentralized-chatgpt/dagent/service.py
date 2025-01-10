import queue
from .models import (
    DAgentLog, ChainState, ClassRegistration, 
    ChatSession, NonInteractiveDAgentLog, Mission, 
    Characteristic
)

import threading
import time 
import logging
import traceback
from . import constant as C
from .registry import get_cls, RegistryCategory 
from typing import Any, Callable, Union, Dict
import schedule
from .characters import DEFAULT_CHAT_COMPLETION_CHARACTER_BUILDER
from .agents import NonInteractiveDAgentBase
from .llm import AsyncChatCompletion
from singleton_decorator import singleton

logger = logging.getLogger(__name__)

@singleton
class AutoServiceProvider(object):
    SCRATCHPAD_LENGTH_LIMIT = 30
    CHAT_SESSION_TIMEOUT = 60 * 60 * 3 # 3 hours

    def __init__(self) -> None:
        self._que = queue.Queue() # a queue of NonInteractiveDAgent
        self._interactive_sessions: Dict[str, ChatSession] = {}
        self._sleep_time = C.AUTO_SERVICE_SLEEP_TIME

    def start(self):
        self._background_thread = threading.Thread(target=self._run, daemon=True)
        self._background_thread.start()

    def schedule(self, cfg: dict):
        def get_or_warning(d: dict, key: str, default: Any = None) -> Any:
            if key not in d:
                logger.warning(f"Key {key} not found in the config dict")
                return default if not callable(default) else default()

            return d[key]

        characteristic: dict = get_or_warning(cfg, "characteristic", {})
        assert isinstance(characteristic, dict), "Characteristic must be a dictionary with a system_prompt "\
            "or detailed information about the character"

        self._characteristic = characteristic
        missions = get_or_warning(cfg, "missions", [])

        for mission in missions:
            task: str = get_or_warning(mission, "task", "")
            system_reminder: str = get_or_warning(mission, "system_reminder", "")

            toolsets_cfg: dict = get_or_warning(mission, "toolset_cfg", {})
            toolsets_cfg = [ClassRegistration(**e) for e in toolsets_cfg]
            llm_cfg = ClassRegistration(**get_or_warning(mission, "llm_cfg", {}))
            agent_builder_cfg = ClassRegistration(**get_or_warning(mission, "agent_builder", {}))

            interval_minutes = int(get_or_warning(mission.get("scheduling"), "interval_minutes", None)) 
            character_builder_cfg = ClassRegistration(**get_or_warning(mission, "character_builder", {}))

            agent_cls = get_cls(
                RegistryCategory.NonInteractiveDAgent, 
                agent_builder_cfg.name
            )

            if interval_minutes is not None and interval_minutes > 0:
                logger.info("Scheduling a mission with interval %d minutes", interval_minutes)

                creator = lambda: agent_cls(
                    NonInteractiveDAgentLog(
                        mission=Mission(
                            task=task,
                            system_reminder=system_reminder,
                        ),
                        toolset_cfg=toolsets_cfg,
                        llm_cfg=llm_cfg,
                        agent_builder_cfg=agent_builder_cfg,
                        character_builder_cfg=character_builder_cfg,
                        characteristic=Characteristic(
                            **characteristic
                        )
                    ), 
                    **agent_builder_cfg.init_params
                )

                if C.IS_SANDBOX:
                    self.enqueue(creator)

                schedule.every(interval=interval_minutes).minutes.do(self.enqueue, creator)

    def enqueue(self, state: Union[DAgentLog, Callable]) -> DAgentLog:
        if callable(state):
            state = state()
        
        logger.info("Enqueueing a new state; ID: %s", state.id)
        self._que.put(state)
        return state

    def _run(self):
        logger.info("The service is running asynchronously in background")
        
        while True:            
            que_length = self._que.qsize()
            
            if que_length > 0:
                logger.info("Processing %d items in the queue", que_length)

            while not self._que.empty():
                agent: NonInteractiveDAgentBase = self._que.get()

                try:
                    log_state: NonInteractiveDAgentLog = agent.step()

                    if agent.state in [ChainState.DONE, ChainState.ERROR]:
                        if agent.state == ChainState.DONE:
                            logger.info(f"Mission {agent.id} is done")

                        else:
                            logger.error(f"Mission {agent.id} has failed")
                            
                        import json, os, datetime
                        os.makedirs("logs", exist_ok=True)
                        dtime_str = datetime.datetime.now().strftime("%Y-%m-%d-%H-%M-%S")
                        with open(f"logs/{dtime_str}_{agent.id}.json", "w") as f:
                            json.dump(log_state.model_dump(), f, indent=4)

                        continue 

                    self._que.put(agent)

                except Exception as err:
                    traceback.print_exc()

            to_be_removed_sessions = []

            for session_id, session in self._interactive_sessions.items():
                if time.time() - session.last_execution > self.CHAT_SESSION_TIMEOUT:
                    to_be_removed_sessions.append(session_id)

            for session_id in to_be_removed_sessions:
                logger.info(f"Removing chat session {session_id}")
                del self._interactive_sessions[session_id]

            time.sleep(self._sleep_time)