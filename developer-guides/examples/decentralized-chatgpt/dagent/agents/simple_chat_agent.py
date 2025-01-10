from dagent.models import Mission
from .base_agent import InteractiveDAgentBase, DAgentLog
from dagent.registry import register_decorator, get_cls, RegistryCategory
from dagent.llm import AsyncChatCompletion
from dagent.tools import ToolsetComposer

@register_decorator(RegistryCategory.InteractiveDAgent)
class SimpleChatDAgent(InteractiveDAgentBase):
    def __init__(self, log: DAgentLog, max_conversation_length=30, *args, **kwargs) -> None:
        super().__init__(log)
        self.max_conversation_length = max_conversation_length + (1 - max_conversation_length % 2)
        
        character_builder_cfg = log.character_builder_cfg
        llm_cfg = log.llm_cfg 
        toolsets_cfg = log.toolset_cfg

        self.llm: AsyncChatCompletion = get_cls(
            RegistryCategory.LLM, llm_cfg.name
        )(**llm_cfg.init_params)

        self.character_builder = get_cls(
            RegistryCategory.CharacterBuilder, character_builder_cfg.name
        )(**character_builder_cfg.init_params)

        self.toolsets = ToolsetComposer([
            get_cls(RegistryCategory.ToolSet, e.name)(**e.init_params) 
            for e in toolsets_cfg
        ])
        
        self.base_system_prompt = self.character_builder(log.characteristic)
        self.log.scratchpad.append({
            'role': 'system',
            'content': self.base_system_prompt
        })
        
    def render_conversation(self) -> list:
        chat_history = []
        ignore_role = ['system-log']

        for message in self.log.scratchpad:
            if message['role'] not in ignore_role:
                chat_history.append(message)
                
        return chat_history

    def __call__(self, mission: Mission) -> DAgentLog:
        self.log.scratchpad.append({
            'role': 'user',
            'content': mission.task
        })

        receipt = self.llm(self.log.scratchpad)
        resp = self.llm.get(receipt.id)

        if resp.result is not None:
            self.log.scratchpad.append({
                'role': 'assistant',
                'content': resp.result,
                'onchain_data': resp.onchain_data.model_dump() if resp.onchain_data else None
            })
        else:
            self.log.scratchpad.pop()
            raise Exception('No response from LLM, please check the LLM service and try again.')

        return self.log
        