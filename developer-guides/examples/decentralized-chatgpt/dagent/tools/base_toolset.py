from typing import List
from dagent.models import Tool
import traceback
import random

class Toolset(object):
    TOOLSET_NAME = "default"
    TOOLS: List[Tool] = []
    PURPOSE = "to get information or take action"

    def __init__(self, 
                exclude=[],
                shuffle_results=False
    ) -> None:
        self.tools = [
            tool for tool in self.TOOLS 
            if tool.name not in exclude
        ]
        self.shuffle_results = shuffle_results

    def render_instruction(self):
        instruct = f'Toolset {self.TOOLSET_NAME}: {self.PURPOSE}:\n'

        for tool in self.tools:
            instruct += f'- {tool.prototype()}\n'

        return instruct

    @property
    def names(self):
        return [tool.name for tool in self.tools]

    def execute(self, name: str, inp: str):
        params = inp.split("|")
        tool = None
        
        for t in self.tools:
            if t.name == name:
                tool = t
                break 
            
        if tool is None:
            return f"{name} not found"
        
        if len(tool.param_spec) == 0:
            return tool.executor()

        if len(params) != len(tool.param_spec):
            return f"Invalid number of parameters. The action requires: {len(tool.param_spec)}. Provided: {len(params)}" 

            requires = len(tool.param_spec)
            results = []

            for i in range(0, len(params), requires):
                if i + requires > len(params):
                    break

                results.append(tool.executor(*params[i:i+requires])) 

            return results
        
        try:
            res = tool.executor(*params)
            
            if isinstance(res, list) and self.shuffle_results:
                random.shuffle(res)
                
            return res

        except Exception as e:
            traceback.print_exc()
            return "Something went wrong while executing the tool: " + str(e)
    
class ToolsetComposer(Toolset):
    def __init__(self, toolsets: List[Toolset], *args, **kwargs) -> None:
        super().__init__(*args, **kwargs)

        self.toolsets = toolsets
        self.tools = [
            tool for toolset in self.toolsets
            for tool in toolset.tools
        ]
        
    def render_instruction(self):
        instruction = ''
        
        for i, e in enumerate(self.toolsets, 1):
            instruction += f'{i}. {e.render_instruction()}\n\n'

        return instruction