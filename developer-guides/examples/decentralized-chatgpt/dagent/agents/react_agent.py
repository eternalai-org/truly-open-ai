from .base_agent import NonInteractiveDAgentBase, InteractiveDAgentBase
from dagent.models import Mission, NonInteractiveDAgentLog
import logging
from dagent.models import ChainState, InferenceState

logger = logging.getLogger(__name__)

from dagent.registry import get_cls, RegistryCategory, register_decorator
from dagent.models import ClassRegistration, DAgentLog
from typing import List
from dagent.tools import ToolsetComposer
from dagent.llm import AsyncChatCompletion
import json

def format_prompt_v2(base_system_prompt: str, toolsets: ToolsetComposer):
    template_prompt = '''
You have access to the following toolset:

{tools}

Your reply to user's message must be a single JSON object with exact three keys described as follows.
thought: your own thought about the next step, reflecting your unique persona.
action: must be one of {toolnames}.
action_input: provide the necessary parameters for the chosen action, separating multiple parameters with the | character.

OR with exact two keys as follows.
thought: your final thought to conclude.
final_answer: your conclusion.

{base_system_prompt}

Again, only return a single JSON!
'''

    tool_names = ', '.join(toolsets.names)
    base_tool_str = toolsets.render_instruction()

    system_prompt = template_prompt.format(
        tools=base_tool_str,
        toolnames=tool_names,
        base_system_prompt=base_system_prompt
    )

    return system_prompt


def render_conversation(log: NonInteractiveDAgentLog, tool: ToolsetComposer):
    system_prompt = format_prompt_v2(log, tool)

    conversation = [
        {
            "role": "system",
            "content": system_prompt
        }
    ]

    for item in log.scratchpad:
        user_message = {}
        for k in ['task', 'observation']:
            if k in item:
                user_message[k] = item[k]

        assistant_message = {}
        for k in ['thought', 'action', 'action_input', 'final_answer']:
            if k in item:
                assistant_message[k] = item[k]

        if len(assistant_message) > 0:
            conversation.append({
                "role": "assistant",
                "content": json.dumps(assistant_message)
            })

        if isinstance(log, NonInteractiveDAgentLog):    
            conversation.append({
                "role": "user",
                "content": json.dumps({
                    **user_message,
                    "system_reminder": log.mission.system_reminder or "Please follow the instructions carefully"
                })
            })

        else:
            conversation.append({
                "role": "user",
                "content": json.dumps(user_message)
            })
            
    return conversation

def parse_conversational_react_response(response: str) -> dict:
    try:
        json_response = json.loads(response)
    except json.JSONDecodeError:
        return {}

    segment_pad = {}

    if "thought" in json_response:
        segment_pad.update({
            "thought": json_response["thought"]
        })

    if "final_answer" in json_response:
        segment_pad.update({
            "final_answer": json_response["final_answer"]
        })

        return segment_pad

    if "action" in json_response:
        segment_pad.update({
            "action": json_response["action"]
        })

        if "action_input" not in json_response:
            json_response["action_input"] = ""

    if "action_input" in json_response:
        segment_pad.update({
            "action_input": json_response["action_input"]
        })

    return segment_pad

def build_llm(cfg: ClassRegistration):
    _cls = get_cls(RegistryCategory.LLM, cfg.name)

    if _cls is None:
        logger.error(f"LLM class {cfg.name} not found")
        return None

    return _cls(**cfg.init_params)

def build_toolset(cfg: List[ClassRegistration]) -> ToolsetComposer:
    _cls = [get_cls(RegistryCategory.ToolSet, e.name) for e in cfg]

    for c, n in zip(_cls, cfg):
        if c is None:
            logger.warning(f"Toolset class {n.name} not found")

    _cls = [e for e in _cls if e is not None]

    if len(_cls) == 0:
        logger.error("No toolset class found")
        return None

    _obj = [e(**f.init_params) for e, f in zip(_cls, cfg)]
    return ToolsetComposer(_obj)

@register_decorator(RegistryCategory.NonInteractiveDAgent)
class ReactReasoningDAgent(NonInteractiveDAgentBase):
    SCRATCHPAD_LENGTH_LIMIT = 30

    def __init__(self, log: NonInteractiveDAgentLog, verbose=True, *args, **kwargs) -> None:
        super().__init__(log)

        character_builder_cfg = log.character_builder_cfg
        llm_cfg = log.llm_cfg
        toolsets_cfg = log.toolset_cfg
        self.verbose = verbose

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

    def __call__(self) -> NonInteractiveDAgentLog:
        log = self.log

        if log.state == ChainState.NEW:
            log.state = ChainState.RUNNING

            system_prompt = format_prompt_v2(self.base_system_prompt, self.toolsets)
            logger.info("🤖 System: " + system_prompt)
            logger.info("👨‍💻 Task: " + log.mission.task)
            logger.info("🔔 Reminder: " + log.mission.system_reminder)

            log.scratchpad = [
                {
                    "task": log.mission.task.replace('\n', ' ').strip(),
                }
            ]
            receipt = self.llm(render_conversation(log, self.toolsets))
            logger.info("Inference receipt: " + receipt.id)
            log.infer_receipt = receipt.id
            return log

        elif log.state == ChainState.RUNNING:
            result = self.llm.get(log.infer_receipt)
            if result.state == InferenceState.EXECUTING:
                return log

            if result.state == InferenceState.ERROR:
                data = log.clone()
                data.update(
                    state=ChainState.ERROR,
                    system_message=result.error
                )
                self.verbose and logger.error("Error in inference: " + result.error)
                return NonInteractiveDAgentLog(**data)

            # update the scratch pad
            message_response = result.result
            pad: dict = parse_conversational_react_response(message_response)

            if len(pad) == 0:
                data = log.clone()
                data.update(
                    state=ChainState.ERROR,
                    system_message="Invalid response from the agent message; Last message: {}".format(message_response)
                )
                return NonInteractiveDAgentLog(**data)

            if 'thought' in pad:
                if 'thought' in log.scratchpad[-1] and any(
                    k not in log.scratchpad[-1]
                    for k in ['action', 'action_input', 'observation']
                ):
                    for kk in ['action', 'action_input', 'observation']:
                        if kk not in log.scratchpad[-1]:
                            log.scratchpad[-1][kk] = "Not found!"

                    data = log.clone()
                    data.update(
                        state=ChainState.ERROR,
                        system_message="Thought found without action/action input/observation"
                    )

                    self.verbose and logger.error("Thought found without action/action input/observation")

                    return NonInteractiveDAgentLog(**data)
                else:
                    log.scratchpad.append({
                        "thought": pad['thought']
                    })

                    self.verbose and logger.info("🤔 Thought: " + pad['thought'])

            if 'action' in pad:
                if 'action_input' not in pad:
                    log.scratchpad[-1]['action'] = pad['action']
                    log.scratchpad[-1]['action_input'] = "Not found!"

                    data = log.clone()
                    data.update(
                        state=ChainState.ERROR,
                        system_message="Action input not found"
                    )

                    self.verbose and logger.error("Action input not found")
                    return NonInteractiveDAgentLog(**data)

                elif 'question' in log.scratchpad[-1]:
                    data = log.clone()
                    data.update(
                        state=ChainState.ERROR,
                        system_message="No thought found"
                    )

                    self.verbose and logger.error("No thought found")
                    return NonInteractiveDAgentLog(**data)

                action = pad['action']
                action_input = pad['action_input']

                self.verbose and logger.info("🛠️ Action: " + action)
                self.verbose and logger.info("🔧 Action input: " + action_input)

                observation = str(self.toolsets.execute(action, action_input))

                self.verbose and logger.info(f"🔍 Observation: {observation}")

                log.scratchpad[-1]['action'] = action
                log.scratchpad[-1]['action_input'] = action_input
                log.scratchpad[-1]['observation'] = observation
            if 'final_answer' in pad:
                if any(k in log.scratchpad[-1] for k in ['action', 'action_input', 'observation']):
                    log.scratchpad.append({})

                log.scratchpad[-1].update({
                    "final_answer": pad['final_answer']
                })

                self.verbose and logger.info("🏁 Final answer: " + pad['final_answer'])
                log.state = ChainState.DONE
                log.system_message = "Final answer found"

                return log

            if len(log.scratchpad) > self.SCRATCHPAD_LENGTH_LIMIT:
                data = log.clone()
                data.update(
                    state=ChainState.ERROR,
                    system_message="Scratchpad length exceeded"
                )

                self.verbose and logger.error("Scratchpad length exceeded, stop here!")
                return NonInteractiveDAgentLog(**data)

            receipt = self.llm(render_conversation(log, self.toolsets))
            log.infer_receipt = receipt.id
            return log

        else:
            data = log.clone()
            data.update(
                state=ChainState.ERROR,
                system_message="Invalid state {}".format(log.state)
            )
            self.verbose and logger.error("Invalid state {}".format(log.state))
            return NonInteractiveDAgentLog(**data)

@register_decorator(RegistryCategory.InteractiveDAgent)
class ReactChatDAgent(InteractiveDAgentBase):
    SCRATCHPAD_LENGTH_LIMIT = 30

    def __init__(self, log: NonInteractiveDAgentLog, verbose=True, *args, **kwargs) -> None:
        super().__init__(log)

        character_builder_cfg = log.character_builder_cfg
        llm_cfg = log.llm_cfg
        toolsets_cfg = log.toolset_cfg
        
        self.verbose = verbose

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

    def _react_step(self, log: DAgentLog, mission: Mission) -> DAgentLog:

        if log.state == ChainState.NEW:
            log.state = ChainState.RUNNING

            system_prompt = format_prompt_v2(self.base_system_prompt, self.toolsets)
            logger.info("🤖 System: " + system_prompt)
            logger.info("👨‍💻 Task: " + mission.task)
            logger.info("🔔 Reminder: " + mission.system_reminder)

            log.scratchpad = [
                {
                    "task": mission.task.replace('\n', ' ').strip(),
                }
            ]
            receipt = self.llm(render_conversation(log, self.toolsets))
            logger.info("Inference receipt: " + receipt.id)
            log.infer_receipt = receipt.id
            return log

        elif log.state == ChainState.RUNNING:
            result = self.llm.get(log.infer_receipt)
            if result.state == InferenceState.EXECUTING:
                return log

            if result.state == InferenceState.ERROR:
                data = log.clone()
                data.update(
                    state=ChainState.ERROR,
                    system_message=result.error
                )
                return DAgentLog(**data)

            # update the scratch pad
            message_response = result.result
            pad: dict = parse_conversational_react_response(message_response)

            if len(pad) == 0:
                data = log.clone()
                data.update(
                    state=ChainState.ERROR,
                    system_message="Invalid response from the agent message; Last message: {}".format(message_response)
                )
                return DAgentLog(**data)

            if 'thought' in pad:
                if 'thought' in log.scratchpad[-1] and any(
                    k not in log.scratchpad[-1]
                    for k in ['action', 'action_input', 'observation']
                ):
                    for kk in ['action', 'action_input', 'observation']:
                        if kk not in log.scratchpad[-1]:
                            log.scratchpad[-1][kk] = "Not found!"

                    data = log.clone()
                    data.update(
                        state=ChainState.ERROR,
                        system_message="Thought found without action/action input/observation"
                    )

                    self.verbose and logger.error("Thought found without action/action input/observation")

                    return DAgentLog(**data)
                else:
                    log.scratchpad.append({
                        "thought": pad['thought']
                    })

                    self.verbose and print("🤔 Thought: " + pad['thought'])

            if 'action' in pad:
                if 'action_input' not in pad:
                    log.scratchpad[-1]['action'] = pad['action']
                    log.scratchpad[-1]['action_input'] = "Not found!"

                    data = log.clone()
                    data.update(
                        state=ChainState.ERROR,
                        system_message="Action input not found"
                    )

                    self.verbose and logger.error("Action input not found")
                    return DAgentLog(**data)

                elif 'question' in log.scratchpad[-1]:
                    data = log.clone()
                    data.update(
                        state=ChainState.ERROR,
                        system_message="No thought found"
                    )

                    self.verbose and logger.error("No thought found")
                    return DAgentLog(**data)

                action = pad['action']
                action_input = pad['action_input']

                self.verbose and print("🛠️ Action: " + action)
                self.verbose and print("🔧 Action input: " + action_input)

                observation = str(self.toolsets.execute(action, action_input))

                self.verbose and print(f"🔍 Observation: {observation}")

                log.scratchpad[-1]['action'] = action
                log.scratchpad[-1]['action_input'] = action_input
                log.scratchpad[-1]['observation'] = observation
            if 'final_answer' in pad:
                if any(k in log.scratchpad[-1] for k in ['action', 'action_input', 'observation']):
                    log.scratchpad.append({})

                log.scratchpad[-1].update({
                    "final_answer": pad['final_answer']
                })

                self.verbose and print("🏁 Final answer: " + pad['final_answer'])
                log.state = ChainState.DONE
                log.system_message = "Final answer found"

                return log

            if len(log.scratchpad) > self.SCRATCHPAD_LENGTH_LIMIT:
                data = log.clone()
                data.update(
                    state=ChainState.ERROR,
                    system_message="Scratchpad length exceeded"
                )

                self.verbose and logger.error("Scratchpad length exceeded, stop here!")
                return DAgentLog(**data)

            receipt = self.llm(render_conversation(log, self.toolsets))
            log.infer_receipt = receipt.id
            return log

        else:
            data = log.clone()
            data.update(
                state=ChainState.ERROR,
                system_message="Invalid state {}".format(log.state)
            )
            self.verbose and logger.error("Invalid state {}".format(log.state))
            return DAgentLog(**data)

    def __call__(self, mission: Mission) -> DAgentLog:
        log_data = self.log.clone()
        log = DAgentLog(**log_data)
        
        chat_history = [
            {
                "role": "system",
                "content": format_prompt_v2(self.base_system_prompt, self.toolsets)
            },
            {
                "role": "user",
                "content": mission.task
            }
        ]

        while log.state not in [ChainState.DONE, ChainState.ERROR]:
            log = self._react_step(log, mission)

        verbose_response = ''
        
        if not self.verbose:
            refine_key = {
                'thought': '🤔 Thought',
                'action': '🛠️ Action',
                'action_input': '🔧 Action input',
                'observation': '🔍 Observation',
                'final_answer': '🏁 Answer'
            }

            for item in log.scratchpad:
                for k in ['thought', 'action', 'action_input', 'observation', 'final_answer']:
                    if k in item:
                        display_key = refine_key[k]
                        verbose_response += f"\n{display_key}: {item[k]}"

            chat_history.append({
                "role": "assistant",
                "content": verbose_response
            })

            log.scratchpad = chat_history
            return log
        
        else:
            chat_history.append(
                {
                    "role": "assistant",
                    "content": log.scratchpad[-1].get(
                        "final_answer", 
                        log.scratchpad[-1].get("thought", "Sorry, I am unable to provide a response")
                    )
                }
            )

            log.scratchpad = chat_history
            return log