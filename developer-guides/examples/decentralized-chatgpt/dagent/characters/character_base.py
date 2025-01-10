import json
from dagent.models import Characteristic

# this class is simply a system prompt builder
class CharacterBuilderBase(object):
    def __init__(self, *args, **kwargs) -> None:
        pass

    def __call__(self, characteristic: Characteristic) -> str:
        return "You are a highly intelligent agent, capable of executing any task assigned to you."