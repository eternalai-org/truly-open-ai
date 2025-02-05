from x_content.constants import ToolSet
from x_content.models import ToolDef
from typing import List
from .utils import execute_tool


class ToolListWrapper(List[ToolDef]):

    async def auto_execute(
        self, name: str, action_input: str, request_id: str = None
    ):
        for tool in self:
            if tool.name == name:
                return await execute_tool(tool, action_input, request_id)

        raise ValueError(f"Tool {name} not found")

    def execute(self, name, *args, **kwargs):
        for tool in self:
            if tool.name == name:
                if callable(tool.executor):
                    return tool.executor(*args, **kwargs)

                raise ValueError(f"Tool {name} is not callable")

        raise ValueError(f"Tool {name} not found")

    def __getattr__(self, name):
        for tool in self:
            if tool.name == name and callable(tool.executor):
                return tool.executor

        return super().__getattr__(name)

    def __hasattr__(self, name):
        for tool in self:
            if tool.name == name and callable(tool.executor):
                return True

        return super().__hasattr__(name)


class IToolCall(object):

    def get_tools(self, toolset: ToolSet = None) -> ToolListWrapper:
        raise NotImplementedError("Method not implemented")
