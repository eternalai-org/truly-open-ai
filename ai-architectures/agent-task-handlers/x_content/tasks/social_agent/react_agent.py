import json
import logging
from x_content.constants import MissionChainState
from x_content.toolcall.utils import execute_tool
from x_content.tasks.utils import format_prompt_v2
from x_content.models import ReasoningLog, ToolDef, ToolLabel
from typing import List
from x_content import constants as const

from x_content.tasks.base import MultiStepTaskBase
from x_content.tasks.utils import a_move_state
from x_content.llm.base import OnchainInferResult

from x_content.wrappers.conversation import get_llm_result_by_model_name
from x_content.wrappers.magic import sync2async

import json_repair

logging.basicConfig(level=logging.INFO if not __debug__ else logging.DEBUG)
logger = logging.getLogger(__name__)

SCRATCHPAD_LENGTH_LIMIT = 10


def parse_conversational_react_response(response: str, verbose=True) -> dict:
    try:
        json_response = json_repair.repair_json(response, return_objects=True)
    except json.JSONDecodeError:
        return {}

    segment_pad = {}

    if "thought" in json_response:
        segment_pad.update({"thought": json_response["thought"]})

    if "final_answer" in json_response:
        segment_pad.update({"final_answer": json_response["final_answer"]})
        return segment_pad

    if "action" in json_response:
        segment_pad.update({"action": json_response["action"]})

        if "action_input" not in json_response:
            json_response["action_input"] = ""

    if "action_input" in json_response:
        segment_pad.update({"action_input": json_response["action_input"]})

    return segment_pad


def has_action(log: ReasoningLog, it: int, tools: List[ToolDef]) -> bool:
    for item in log.scratchpad[:it]:
        action_name = item.get("action")

        for tool in tools:
            if tool.label == ToolLabel.ACTION and action_name == tool.name:
                return True

    return False


def has_news(log: ReasoningLog, it: int, tools: List[ToolDef]) -> bool:

    for item in log.scratchpad[:it]:
        action_name = item.get("action")

        for tool in tools:
            if tool.label == ToolLabel.QUERY and action_name == tool.name:
                return True

    return False


def dynamic_system_reminder(
    log: ReasoningLog, it: int, tools: List[ToolDef]
) -> str:
    query_tools = [tool for tool in tools if tool.label == ToolLabel.QUERY]
    tool_name_str = ", ".join([tool.name for tool in query_tools])

    if (
        it
        <= int(
            log.meta_data.params.get(
                "react_max_steps", const.DEFAULT_REACT_MAX_STEPS
            )
        )
        * 3
        // 4
    ):
        return "Follow the instructions carefully and step-by-step do the task. Hint: use the following tools ({}) to get necessary information, news, update your context before taking any action!".format(
            tool_name_str
        )

    if len(query_tools) > 0 and not has_news(log, it, tools):
        return "Find the latest news or necessary information to the task by using one of the following tools: {}".format(
            tool_name_str
        )

    action_tools = [tool for tool in tools if tool.label == ToolLabel.ACTION]

    if len(action_tools) > 0 and not has_action(log, it, tools):
        tool_name_str = ", ".join([tool.name for tool in action_tools])
        return "Complete the task by taking an action: {}; or provide a final_answer".format(
            tool_name_str
        )

    return "Complete the task by making a final_answer"


def render_conversation(log: ReasoningLog, tools: List[ToolDef]):
    conversation = [
        {"role": "system", "content": format_prompt_v2(log, tools)}
    ]

    for it, item in enumerate(log.scratchpad):
        user_message = {}
        for k in ["task", "observation", "hot_news"]:
            if k in item:
                user_message[k] = item[k]

        assistant_message = {}
        for k in ["thought", "action", "action_input", "final_answer"]:
            if k in item:
                assistant_message[k] = item[k]

        if len(assistant_message) > 0:
            conversation.append(
                {"role": "assistant", "content": json.dumps(assistant_message)}
            )

        response = {
            **user_message,
            "system reminder": dynamic_system_reminder(log, it + 1, tools),
        }

        conversation.append({"role": "user", "content": json.dumps(response)})

    return conversation


class ReactAgent(MultiStepTaskBase):
    resumable = True

    async def is_react_complete(self, log: ReasoningLog) -> bool:
        return (
            len(log.scratchpad) > 0 and "final_answer" in log.scratchpad[-1]
        ) or len(log.scratchpad) > log.meta_data.params.get(
            "react_max_steps", const.DEFAULT_REACT_MAX_STEPS
        )

    async def process_task(self, log: ReasoningLog) -> ReasoningLog:
        tools = self.toolcall.get_tools(log.toolset)

        if log.state == MissionChainState.NEW:
            log.scratchpad = [
                {
                    "task": log.prompt.replace("\n", " ").strip(),
                }
            ]

            log.execute_info = {
                "tool_call_metadata": [],
                "conversation": [],
            }

            return await a_move_state(
                log, MissionChainState.RUNNING, "Task started"
            )

        if log.state == MissionChainState.RUNNING:
            while not await self.is_react_complete(log):
                conversation = await sync2async(render_conversation)(
                    log, tools
                )
                log.execute_info["conversation"].append(conversation)
                infer_result: OnchainInferResult = await self.llm.agenerate(
                    conversation, temperature=0.7
                )

                result = infer_result.generations[0].message.content
                result = get_llm_result_by_model_name(result, log.model)
                pad: dict = await sync2async(
                    parse_conversational_react_response
                )(result)

                if len(pad) == 0:
                    return await a_move_state(
                        log,
                        MissionChainState.ERROR,
                        "No response (or wrong response format) from the agent message; Last: {}; receipt: {}; tx-hash: {}".format(
                            result, infer_result.receipt, infer_result.tx_hash
                        ),
                    )

                log = await self.update_react_scratchpad(log, pad)

                if log.state in [
                    MissionChainState.DONE,
                    MissionChainState.ERROR,
                ]:
                    break

                log.scratchpad[-1].update(
                    {
                        "tx_hash": infer_result.tx_hash,
                    }
                )

                log = await self.commit_log(log)

        return log

    async def update_react_scratchpad(self, log: ReasoningLog, pad: dict):
        tools = self.toolcall.get_tools(log.toolset)

        if "thought" in pad:
            if "thought" in log.scratchpad[-1] and any(
                k not in log.scratchpad[-1]
                for k in ["action", "action_input", "observation"]
            ):
                for kk in ["action", "action_input", "observation"]:
                    if kk not in log.scratchpad[-1]:
                        log.scratchpad[-1][kk] = "Not found!"

                return await a_move_state(
                    log,
                    MissionChainState.ERROR,
                    "Action/Action Input/Observation of the last step not found",
                )

            log.scratchpad.append({"thought": pad["thought"]})

            logger.info(
                f"[{log.id}][React-Iter {len(log.scratchpad)}] 🤔 Thought: {pad['thought']}"
            )

        if "action" in pad:
            if "action_input" not in pad:
                log.scratchpad[-1].update(
                    action=pad["action"], action_input="Not found!"
                )

                return await a_move_state(
                    log, MissionChainState.ERROR, "Action input not found"
                )

            if "task" in log.scratchpad[-1]:
                return await a_move_state(
                    log, MissionChainState.ERROR, "No thought found!"
                )

            action = pad["action"]
            logger.info(
                f"[{log.id}][React-Iter {len(log.scratchpad)}] 🛠️ Action: {action}"
            )

            action_input = pad["action_input"]
            logger.info(
                f"[{log.id}][React-Iter {len(log.scratchpad)}] 📥 Action Input: {action_input}"
            )

            log.scratchpad[-1].update(action=action, action_input=action_input)

            observation = "Action not found!"

            # TODO: refactor this loop
            for tool in tools:
                if tool.name == action:

                    observation = await execute_tool(
                        tool, action_input, request_id=log.id
                    )
                    new_observation = []

                    for x in observation:
                        if isinstance(x, tuple):
                            x, metadata = x
                            log.execute_info["tool_call_metadata"].append(
                                metadata
                            )

                        new_observation.append(x)

                    observation = new_observation
                    break

            log.scratchpad[-1]["observation"] = observation
            logger.info(
                f"[{log.id}][React-Iter {len(log.scratchpad)}] 🔍 Observation: {observation}"
            )

        if "final_answer" in pad:
            if any(
                k in log.scratchpad[-1]
                for k in ["action", "action_input", "observation"]
            ):
                log.scratchpad.append({})

            logger.info(
                f"[{log.id}][React-Iter {len(log.scratchpad)}] 🎯 Final Answer: {pad['final_answer']}"
            )
            log.scratchpad[-1].update({"final_answer": pad["final_answer"]})

            log = await a_move_state(
                log, MissionChainState.DONE, "Final answer found"
            )

        if len(log.scratchpad) > log.meta_data.params.get(
            "react_max_steps", const.DEFAULT_REACT_MAX_STEPS
        ):
            log = await a_move_state(
                log, MissionChainState.DONE, "Scratchpad length exceeded"
            )

        return log
