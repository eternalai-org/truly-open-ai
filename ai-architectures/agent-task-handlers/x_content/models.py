from pydantic import BaseModel, Field, model_validator
from typing import List, Dict, Any, Optional, Union, Callable
from enum import Enum
import uuid
import json
from hashlib import md5
import logging
from datetime import datetime
from x_content import constants as const
from x_content.constants import AgentTask, HTTPMethod, MissionChainState, ToolSet
from langchain.schema import ChatMessage

logger = logging.getLogger(__name__)


def random_uuid() -> str:
    return str(uuid.uuid4().hex)


class ReactAgentReasoningMeta(BaseModel):
    twitter_id: Optional[str] = ""
    twitter_username: Optional[str] = ""
    agent_contract_id: Optional[str] = ""
    chain_id: Optional[Union[int, str]] = ""  # tbd
    system_reminder: Optional[str] = (
        None  # set a reminder for to LLM for each step
    )
    params: Dict[str, Any] = {
        "quote_username": "cryptopunksbot",
        "react_max_steps": const.DEFAULT_REACT_MAX_STEPS,
    }
    ref_id: Optional[str] = ""
    knowledge_base_id: str = ""


class TokenInfo(BaseModel):
    name: Optional[str] = None
    symbol: Optional[str] = None
    address: Optional[str] = None
    chain: Optional[str] = None

    def json_system_prompt(self):
        res = {}

        if self.name is not None and self.name != "":
            res["name"] = self.name

        if self.symbol is not None and self.symbol != "":
            res["symbol"] = self.symbol

        if self.address is not None and self.address != "":
            res["address"] = self.address

        if self.chain is not None and self.chain != "":
            res["chain"] = self.chain

        return res


class AgentKnowledgeBase(BaseModel):
    chain_id: str
    kb_id: str


class AgentMetadata(BaseModel):
    persona: Optional[str] = None
    token_info: Optional[TokenInfo] = None
    liked_topics: Optional[List[str]] = []
    disliked_topics: List[str] = []
    style: Optional[Union[str, list]] = None
    names: List[str] = []
    kb_agents: Optional[List[AgentKnowledgeBase]] = []

    @model_validator(mode="before")
    def validate_kb_agents(cls, data: dict):
        if "kb_agents" not in data or data["kb_agents"] is None:
            data["kb_agents"] = []

        if isinstance(data["kb_agents"], str):
            try:
                data["kb_agents"] = json.loads(data["kb_agents"])
            except Exception as e:
                logger.warning(
                    f"Failed to parse kb_agents: {e}; Received: {data['kb_agents']}"
                )
                data["kb_agents"] = []

        if isinstance(data["kb_agents"], list):
            data["kb_agents"] = [
                AgentKnowledgeBase(**x) for x in data["kb_agents"]
            ]

        return data

    def json_system_prompt(self) -> dict:
        # res = {}

        # res['personality'] = self.persona

        # if len(self.liked_topics) > 0:
        #     res['what you really love'] = self.liked_topics

        # if len(self.disliked_topics) > 0:
        #     res['what you hate'] = self.disliked_topics

        # if self.style is not None:
        #     res['your style'] = self.style

        # if self.token_info is not None:
        #     xx = self.token_info.json_system_prompt()
        #     if len(xx) > 0:
        #         res['your own token'] = xx # self.token_info.json_system_prompt()

        return self.persona


def is_single_token(token):
    if " " in token or "\n" in token or "\t" in token:
        return False

    return True


def modified_json_stringnify(obj: dict, depth=0, indent=2) -> str:
    leading_space = " " * (depth * indent)

    if isinstance(obj, dict):
        res = []

        for k, v in obj.items():
            if not is_single_token(k):
                k = f'"{k}"'

            res.append(
                "{k}: {v}".format(
                    leading_space,
                    k=k,
                    v=modified_json_stringnify(v, depth + 1, indent),
                )
            )

        return "{" + ", ".join(res) + "}"

    if isinstance(obj, list):
        res = []

        for v in obj:
            res.append(modified_json_stringnify(v))

        if len(res) == 1:
            return f"{res[0]}"

        return "[" + ", ".join(res) + "]"

    if obj is None:
        return "null"

    if isinstance(obj, bool):
        return "true" if obj else "false"

    if isinstance(obj, str):

        obj = (
            obj.replace('\\"', '"')
            .replace("\n", " ")
            .replace("\t", " ")
            .replace("\r", " ")
            .replace('"', '\\"')
            .strip()
        )

        if not is_single_token(obj) or len(obj) > 20:
            return f'"{obj}"'

        return obj

    return str(obj)


class ToolParamDtype(str, Enum):
    STRING = "string"
    NUMBER = "number"
    BOOLEAN = "bool"


class ToolParam(BaseModel):
    name: str
    dtype: ToolParamDtype

    default_value: Optional[Union[str, bool, float]] = None
    description: Optional[str] = None

    @model_validator(mode="before")
    def validate_default(cls, data: dict):
        if data is None:
            return data

        if "default_value" in data and data["default_value"] is not None:
            if data["dtype"] == ToolParamDtype.BOOLEAN:
                data["default_value"] = data["default_value"] in [
                    "true",
                    "True",
                    "1",
                    "TRUE",
                    True,
                ]

            if data["dtype"] == ToolParamDtype.NUMBER:
                try:
                    data["default_value"] = float(data["default_value"])
                except Exception as e:
                    raise ValueError(
                        f"Invalid default_value value for number type: {data['default_value']}"
                    )

            if data["dtype"] == ToolParamDtype.STRING:
                data["default_value"] = str(data["default_value"])

        return data


class ToolLabel(str, Enum):
    QUERY = "query"
    ACTION = "action"


class ToolDef(BaseModel):
    name: str
    description: str
    params: List[ToolParam]
    executor: Optional[Union[Callable, str]] = None
    allow_multiple: bool = False
    label: ToolLabel

    def prototype(self):
        no_default_params = [
            param for param in self.params if param.default_value is None
        ]
        params_str = ", ".join(
            [
                f"{param.name}: {param.dtype.value}"
                for param in no_default_params
            ]
        )

        return f"{self.name}({params_str}): Takes {len(no_default_params)} parameters, {self.description}"


class AdvanceToolDef(ToolDef):
    headers: Optional[Dict[str, str]] = None
    method: HTTPMethod = HTTPMethod.GET
    label: ToolLabel = ToolLabel.QUERY


class AutoAgentTask(BaseModel):
    # auto
    id: str = Field(default_factory=lambda: f"fun-{random_uuid()}")

    # request
    meta_data: Optional[ReactAgentReasoningMeta] = None
    agent_meta_data: AgentMetadata = AgentMetadata()
    model: Optional[str] = ""

    # @deprecated
    system_prompt: Optional[str] = "You are a helpful assistant."

    # task status
    scratchpad: List[Dict[str, Any]] = []
    state: MissionChainState = MissionChainState.NEW
    system_message: str = ""  # for error messages

    created_at: str = Field(
        default_factory=lambda: datetime.now().strftime(
            "%Y-%m-%dT%H:%M:%S.%fZ"
        )
    )
    last_process: str = Field(
        default_factory=lambda: datetime.now().strftime(
            "%Y-%m-%dT%H:%M:%S.%fZ"
        )
    )

    def is_done(self):
        return self.state == MissionChainState.DONE

    def is_error(self):
        return self.state == MissionChainState.ERROR

    @model_validator(mode="after")
    def forecast_agent_meta_data(self):
        if self.agent_meta_data is None:
            self.agent_meta_data = AgentMetadata()

        if (
            self.agent_meta_data.persona is None
            or self.agent_meta_data.persona == ""
        ) and self.system_prompt is not None:
            self.agent_meta_data.persona = self.system_prompt or ""

        if len(self.agent_meta_data.names) == 0 and self.meta_data is not None:
            self.agent_meta_data.names = [self.meta_data.twitter_username]

        self.agent_meta_data.liked_topics = list(
            # set(self.agent_meta_data.liked_topics + ['crypto', 'nft', 'web3', 'blockchain', 'EAI', 'BTCH'])
            set(self.agent_meta_data.liked_topics)
        )

        self.agent_meta_data.kb_agents = self.agent_meta_data.kb_agents or []

        if isinstance(self.agent_meta_data.style, list):
            if len(self.agent_meta_data.style) > 0:
                self.agent_meta_data.style = self.agent_meta_data.style[0]

            else:
                self.agent_meta_data.style = None

        return self

    @property
    def checksum(self):
        def cvt(data):
            if isinstance(data, set):
                return list(data)

            return str(data)

        try:
            data = json.dumps(
                self.model_dump(), sort_keys=True, default=cvt
            ).encode("utf-8")

        except Exception as err:
            logger.error(
                f"Failed to encode the reasoning log to get checksum: {err}! Returning a non-standard stringified json to encode instead."
            )

            data = json.dumps(self.model_dump(), default=cvt).encode("utf-8")

        return md5(data).hexdigest()


class ReasoningLog(AutoAgentTask):
    # for request
    prompt: str  # set a goal for the agent
    seed: Optional[int] = 512  # seed for the agent

    @property
    def _system_prompt_(self) -> str:
        return modified_json_stringnify(
            self.agent_meta_data.json_system_prompt()
        )

    task: Optional[AgentTask] = AgentTask.REACT_AGENT
    toolset: Optional[str] = ToolSet.DEFAULT
    tool_list: Optional[List[AdvanceToolDef]] = []
    need_to_post_process: bool = False

    # for response
    infer_receipt: Optional[Union[int, str]] = None
    execute_info: Dict[str, Any] = {}
    llm_info: Dict[str, Any] = {}

    @model_validator(mode="before")
    @classmethod
    def parse_tool_list(cls, data: dict):
        if not isinstance(data, dict):
            return data

        if "tool_list" not in data or data["tool_list"] == None:
            data["tool_list"] = []
        else:
            if isinstance(data["tool_list"], str):
                try:
                    data["tool_list"] = json.loads(data["tool_list"])
                except Exception as e:
                    # logger.warning(f"Failed to parse tool_list: {e}; Received: {data['tool_list']}")
                    data["tool_list"] = []

            if isinstance(data["tool_list"], list):
                data["tool_list"] = [
                    AdvanceToolDef(**x) for x in data["tool_list"]
                ]

        if "params" in data.get("meta_data", {}):
            if "react_max_steps" in data["meta_data"]["params"]:
                data["meta_data"]["params"]["react_max_steps"] = int(
                    data["meta_data"]["params"]["react_max_steps"]
                )

            else:
                data["meta_data"]["params"][
                    "react_max_steps"
                ] = const.DEFAULT_REACT_MAX_STEPS

        return data


class ChatRequest(AutoAgentTask):
    # user data
    user_address: str = ""
    messages: List[ChatMessage] = []

    # for response
    chat_result: Optional[str] = None


class APIStatus(str, Enum):
    SUCCESS = "success"
    ERROR = "error"


class APIResponse(BaseModel):
    status: APIStatus
    data: Any
    error: Optional[str] = None


class TwinUpdateResponse(BaseModel):
    agent_id: str = Field(..., description="Agent ID")
    twin_status: str = Field(..., description="Twin status")
    knowledge_base_id: str = Field(..., description="Knowledge base ID")
    system_prompt: str = Field(..., description="System prompt")
    twin_training_progress: int = Field(
        ..., description="Twin training progress"
    )
    twin_training_message: str = Field("", description="Twin training message")


class TwinTaskSubmitResponse(BaseModel):
    status: str = Field(..., description="Status of the task submission")
    task_id: str = Field(..., description="Task ID")


class TwinTaskSubmitRequest(BaseModel):
    agent_id: str = Field(..., description="Agent ID")
    twitter_ids: List[str] = Field(..., description="List of Twitter IDs")
