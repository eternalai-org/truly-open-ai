from typing import Any
from .character_base import CharacterBuilderBase, Characteristic
from dagent.registry import register_decorator, RegistryCategory
from dagent import constant as C

@register_decorator(RegistryCategory.CharacterBuilder)
class SimpleCharacterBuilder(CharacterBuilderBase):
    def __call__(self, characteristic: Characteristic) -> str:
        if characteristic.system_prompt is not None:
            return characteristic.system_prompt

        characteristic_representation_template = '''
You are {name}, capable of executing any task assigned to you.

Here is a brief overview of your capabilities:    
{knowledge}

{bio}

{lore}

{interested_topics}'''

        personal_info = characteristic.agent_personal_info
        agent_name = personal_info.get("agent_name", "a highly intelligent AI assistant")

        bio_data, lore_data, knowledge_data = \
            characteristic.bio,  \
            characteristic.lore, \
            characteristic.knowledge

        bio_repr = "# Bio" if len(bio_data) > 0 else ""
        lore_repr = "# Lore" if len(lore_data) > 0 else ""
        knowledge_repr = "# Knowledge" if len(knowledge_data) > 0 else ""

        for bio in bio_data[:C.DEFAULT_BIO_MAX_LENGTH]:
            bio_repr += f"\n- {bio}"

        for lore in lore_data[:C.DEFAULT_LORE_MAX_LENGTH]:
            lore_repr += f"\n- {lore}"
            
        for knowledge in knowledge_data[:C.DEFAULT_KNOWLEDGE_MAX_LENGTH]:
            knowledge_repr += f"\n- {knowledge}"

        interested_topics_data = characteristic.interested_topics
        interested_topics_repr = "# Interested Topics" if len(interested_topics_data) > 0 else ""
        
        for topic in interested_topics_data[:C.DEFAULT_INTERESTED_TOPICS_MAX_LENGTH]:
            interested_topics_repr += f"\n- {topic}"

        system_prompt = characteristic_representation_template.format(
            bio=bio_repr,
            lore=lore_repr,
            knowledge=knowledge_repr,
            interested_topics=interested_topics_repr,
            name=agent_name
        )
    
        return system_prompt
