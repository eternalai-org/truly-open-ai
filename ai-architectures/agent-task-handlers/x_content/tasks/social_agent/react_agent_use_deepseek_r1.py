from .react_agent import ReactAgent
from json_repair import json_repair
import json
import re
from .react_agent import dynamic_system_reminder
from x_content.models import ReasoningLog, ToolDef, MissionChainState
from typing import List
from x_content.tasks.utils import a_move_state
from x_content.toolcall.utils import execute_tool
from x_content.wrappers.magic import sync2async
from x_content.llm import OnchainInferResult
from x_content import constants as const
import logging

logger = logging.getLogger(__name__)


def format_prompt_v2(log: ReasoningLog, tools: List[ToolDef]):
    template_prompt = """
You have access to the following tools to get information or take actions:
{tools}

Your knowledge is outdated for years. So, for the best outcome, you must take some actions, to get information, to update your knowledge before acting something.

Your answer must be a single JSON object with exactly three keys described as follows:
- name: must be one of {toolnames}.
- parameters: provide the necessary parameters for the chosen action.

OR with exactly one key as follows:
- conclusion: your conclusion.

About you:
{base_system_prompt}

Again, only shortly think about the task and answer in a single JSON!
"""

    tool_names = ", ".join([tool.name for tool in tools])
    base_tool_str = "\n".join(e.prototype() for e in tools)

    return template_prompt.format(
        tools=base_tool_str,
        toolnames=tool_names,
        base_system_prompt=log._system_prompt_,
    )


def parse_conversational_react_response(response: str, verbose=True) -> dict:
    response_pattern = re.compile(
        "<think>(.*?)</think>(.*?)", re.IGNORECASE | re.DOTALL
    )

    try:
        match = response_pattern.match(response)
        assert match is not None, "No match"
        thought = match.group(1)
        other_response = match.group(2)
        json_response = json_repair.repair_json(
            other_response, return_objects=True
        )
    except json.JSONDecodeError:
        return {}

    segment_pad = {}

    segment_pad.update({"thought": thought})

    if "conclusion" in json_response:
        segment_pad.update({"conclusion": json_response["conclusion"]})
        return segment_pad

    if "name" in json_response:
        segment_pad.update({"name": json_response["name"]})

        if "parameters" not in json_response:
            json_response["parameters"] = ""

    if "parameters" in json_response:
        segment_pad.update({"parameters": json_response["parameters"]})

    return segment_pad


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
        for k in ["thought", "name", "parameters", "conclusion"]:
            if k in item:
                assistant_message[k] = item[k]

        if len(assistant_message) > 0:
            assistant_resp = "<think>{}</think>{}".format(
                assistant_message.get("thought", ""),
                json.dumps(
                    {
                        k: v
                        for k, v in assistant_message.items()
                        if k != "thought"
                    }
                ),
            )

            conversation.append(
                {"role": "assistant", "content": assistant_resp}
            )

        response = {
            **user_message,
            "system reminder": dynamic_system_reminder(log, it + 1, tools),
        }

        conversation.append({"role": "user", "content": json.dumps(response)})

    return conversation


class ReactAgentUsingDeepSeekR1(ReactAgent):
    resumable = True

    async def is_react_complete(self, log: ReasoningLog) -> bool:
        return (
            len(log.scratchpad) > 0 and "conclusion" in log.scratchpad[-1]
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
                for k in ["name", "parameters", "observation"]
            ):
                for kk in ["name", "parameters", "observation"]:
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

        if "name" in pad:
            if "parameters" not in pad:
                log.scratchpad[-1].update(
                    action=pad["name"], action_input="Not found!"
                )

                return await a_move_state(
                    log, MissionChainState.ERROR, "Action input not found"
                )

            if "task" in log.scratchpad[-1]:
                return await a_move_state(
                    log, MissionChainState.ERROR, "No thought found!"
                )

            action = pad["name"]
            logger.info(
                f"[{log.id}][React-Iter {len(log.scratchpad)}] 🛠️ Action: {action}"
            )

            action_input = pad["parameters"]
            logger.info(
                f"[{log.id}][React-Iter {len(log.scratchpad)}] 📥 Action Input: {action_input}"
            )

            log.scratchpad[-1].update(action=action, action_input=action_input)

            observation = "Action not found!"

            # TODO: refactor this loop
            for tool in tools:
                if tool.name == action:

                    action_input_ordered = []

                    for p in tool.params:
                        if p in action_input:
                            action_input_ordered.append(
                                str(action_input.get(p, ""))
                            )

                    action_input_str = "|".join(action_input_ordered)

                    observation = await execute_tool(
                        tool, action_input_str, request_id=log.id
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

        if "conclusion" in pad:
            if any(
                k in log.scratchpad[-1]
                for k in ["name", "parameters", "observation"]
            ):
                log.scratchpad.append({})

            logger.info(
                f"[{log.id}][React-Iter {len(log.scratchpad)}] 🎯 Final Answer: {pad['conclusion']}"
            )
            log.scratchpad[-1].update({"conclusion": pad["conclusion"]})

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
