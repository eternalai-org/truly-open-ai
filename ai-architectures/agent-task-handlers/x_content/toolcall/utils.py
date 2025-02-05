import traceback
import logging
from x_content.models import (
    ToolParamDtype,
    ToolDef,
    AdvanceToolDef,
    ToolParam,
)
from x_content import constants as const
from typing import List, Any, Union
import requests
import json
import asyncio
from x_content import wrappers
from x_content.wrappers.magic import sync2async

logger = logging.getLogger(__name__)
LIMIT_TOTAL_TOKENS_PER_OBSERVATION = 700 * 3


def is_float(xx: Any):
    try:
        float(xx)
        return True
    except ValueError:
        return False


def split_params(spec: dict, action_input):
    resp = {}

    for i, (k, v) in enumerate(spec.items()):
        if v["type"] == "number":
            resp[k] = int(action_input[i])
        else:  # v['type'] == "string":
            resp[k] = str(action_input[i])

    return resp


def make_response(content, success=True):
    return {
        "success": success,
        "content": content,
    }


def execute_http_toolcall(method, endpoint, params, headers=None):
    payload = dict(
        method=method,
        url=endpoint,
    )

    if headers is not None:
        payload["headers"] = headers

    if method == "GET":
        payload["params"] = params
    else:
        payload["json"] = params

    logger.info(f"Calling API with payload: {json.dumps(payload, indent=2)}")
    resp = requests.request(**payload)

    if resp.status_code != 200:
        pre_formated = (
            "Method: {}\n"
            "Endpoint: {}\n"
            "Params: \n<pre>{}</pre>\n"
            "Headers: \n<pre>{}</pre>\n"
            "Response: \n<pre>{}</pre>\n"
            "Code: {}\n"
        ).format(
            method,
            endpoint,
            json.dumps(params, indent=2),
            (
                json.dumps({k: "***" for k, _ in headers.items()}, indent=2)
                if isinstance(headers, dict)
                else "n/a"
            ),
            resp.text,
            resp.status_code,
        )

        task_id = params.get("request_id", None)

        wrappers.telegram.send_message(
            "junk_nofitications",
            f"Request {task_id} (dynamic toolcall): Failed to call API\n\n{pre_formated}",
            room=wrappers.telegram.TELEGRAM_ALERT_ROOM,
            fmt="HTML",
        )

        return make_response(
            f"Error: {resp.text} (Code: {resp.status_code})", False
        )

    try:
        resp_to_agent = resp.json()
    except json.JSONDecodeError:
        resp_to_agent = resp.text

    return make_response(resp_to_agent, True)


def value_cast(dtype: ToolParamDtype, value: str):
    if dtype == ToolParamDtype.STRING:
        return value

    if dtype == ToolParamDtype.NUMBER:
        return float(value)

    if dtype == ToolParamDtype.BOOLEAN:
        return value.lower() in ["true", "1"]

    return value


def parse_toolcall_params(params: List[ToolParam], inputs):
    no_default_params = [p for p in params if p.default_value is None]
    res = {}

    for p in params:
        if p.default_value is not None:
            res[p.name] = p.default_value

    res = {
        **res,
        **{
            k.name: value_cast(k.dtype, v)
            for k, v in zip(no_default_params, inputs)
        },
    }

    return res


def execute_advance_tool(
    tool: AdvanceToolDef, action_input: str, request_id: str = None
):
    inputs = [e.strip() for e in action_input.split("|")]
    no_default_params = [p for p in tool.params if p.default_value is None]

    if len(no_default_params) == 0:
        return [
            execute_http_toolcall(
                tool.method,
                tool.executor,
                {
                    **parse_toolcall_params(tool.params, []),
                    "request_id": request_id,
                },
                tool.headers,
            )
        ]

    try:
        n_params = len(no_default_params)
        res = []

        for i in range(0, len(inputs), n_params):
            l, r = i, i + n_params

            if r > len(inputs):
                break

            res.append(
                execute_http_toolcall(
                    tool.method,
                    tool.executor,
                    {
                        **parse_toolcall_params(tool.params, inputs[l:r]),
                        "request_id": request_id,
                    },
                    tool.headers,
                )
            )

            if not tool.allow_multiple:
                break

        if len(res) == 0:
            if n_params > 0:
                return [
                    "Something went wrong while executing the tool and no results were returned."
                ]

            missing_params = [p.name for p in tool.params[-len(inputs) :]]
            return [
                f"{tool.name} requires {len(tool.params)} parameters, but only {len(inputs)} provided. Missing {len(missing_params)} value(s) for: {missing_params}"
            ]

        return res

    except Exception as err:
        traceback.print_exc()
        logger.error(
            f"Error while executing {tool.name} with input {action_input}: {err}"
        )
        return [f"Error while executing {tool.name} with input {action_input}"]


async def execute_tool(
    tool: Union[ToolDef, AdvanceToolDef],
    action_input: str,
    request_id: str = None,
):
    if isinstance(tool, AdvanceToolDef):
        return await sync2async(execute_advance_tool)(
            tool, action_input, request_id
        )

    inputs = action_input.split("|")

    if len(tool.params) == 0:
        if asyncio.iscoroutinefunction(tool.executor):
            return [await tool.executor()]
        else:
            return [await sync2async(tool.executor)()]

    try:
        n_params = len(tool.params)
        res = []

        for i in range(0, len(inputs), n_params):
            if i + n_params > len(inputs):
                break

            l, r = i, i + n_params
            if asyncio.iscoroutinefunction(tool.executor):
                result = await tool.executor(*inputs[l:r])
            else:
                result = await sync2async(tool.executor)(*inputs[l:r])
            res.append(result)

            if not tool.allow_multiple:
                break

        if len(res) == 0:
            missing_params = [p.name for p in tool.params[-len(inputs) :]]
            return [
                f"{tool.name} requires {len(tool.params)} parameters, but only {len(inputs)} provided. Missing {len(missing_params)} value(s) for: {missing_params}"
            ]

        return res

    except Exception as err:
        traceback.print_exc()
        logger.error(
            f"Error while executing {tool.name} with input {action_input}: {err}"
        )
        return [f"Error while executing {tool.name} with input {action_input}"]
