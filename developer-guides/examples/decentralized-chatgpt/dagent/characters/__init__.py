from .character_base import CharacterBuilderBase

modules = []
import os

current_dir = os.path.dirname(__file__)
for file in os.listdir(current_dir):
    if os.path.isfile(os.path.join(current_dir, file)) and file.endswith(".py") and not file.startswith("__"):
        modules.append(os.path.basename(file)[:-3])

__all__ = [
    "CharacterBuilderBase",
    *modules
]

from . import *

from . import simple_character
DEFAULT_CHAT_COMPLETION_CHARACTER_BUILDER = simple_character.SimpleCharacterBuilder