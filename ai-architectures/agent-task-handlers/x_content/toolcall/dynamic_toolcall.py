from x_content.constants import HTTPMethod
from x_content.models import AdvanceToolDef, ToolParam, ToolParamDtype
from typing import List
from .toolcall import IToolCall


class DynamicToolcall(IToolCall):

    def __init__(self, definitions: List[AdvanceToolDef]):
        self.defs = definitions

    def get_tools(self, toolset=None):
        return self.defs