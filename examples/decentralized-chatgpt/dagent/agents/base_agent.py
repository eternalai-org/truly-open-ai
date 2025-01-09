from typing import Any
from dagent.models import DAgentLog, NonInteractiveDAgentLog, ChainState, Mission, DAgentResponse, OnChainData

class InteractiveDAgentBase(object):
    def __init__(self, log: DAgentLog) -> None:
        self.log = log
        
    @property
    def id(self) -> str:
        return self.log.id

    def step(self, mission: Mission) -> DAgentResponse:
        resp = self.__call__(mission)
        
        assert resp.scratchpad[-1]['role'] == 'assistant'
        
        return DAgentResponse(
            content=resp.scratchpad[-1]['content'],
            onchain_data=(
                None if resp.scratchpad[-1].get('onchain_data') is None 
                else OnChainData.model_validate(resp.scratchpad[-1]['onchain_data'])
            )
        )

    def __call__(self, log: Mission) -> DAgentLog:
        raise NotImplementedError("You must implement this method in your subclass")

class NonInteractiveDAgentBase(object):
    def __init__(self, log: NonInteractiveDAgentLog) -> None:
        self.log = log
        
    @property
    def id(self) -> str:
        return self.log.id

    @property
    def state(self) -> ChainState:
        return self.log.state

    def step(self) -> NonInteractiveDAgentLog:
        if self.log.state == ChainState.NEW or self.log.state == ChainState.RUNNING:
            return self.__call__()

        return self.log

    def __call__(self) -> NonInteractiveDAgentLog:
        raise NotImplementedError("You must implement this method in your subclass")