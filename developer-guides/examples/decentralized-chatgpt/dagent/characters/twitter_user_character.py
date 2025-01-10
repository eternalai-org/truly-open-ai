from typing import Any
from .character_base import CharacterBuilderBase, Characteristic
from dagent.registry import register_decorator, RegistryCategory
from dagent import constant as C
import random

@register_decorator(RegistryCategory.CharacterBuilder)
class TwitterUserCharacterBuilder(CharacterBuilderBase):
    def __init__(self, shuffle_everything=True, *args, **kwargs) -> None:
        super().__init__(*args, **kwargs)
        self.shuffle_everything = shuffle_everything
        
    def __call__(self, characteristic: Characteristic) -> str:
        if characteristic.system_prompt is not None:
            return characteristic.system_prompt

        characteristic_representation_template = '''
You are {agent_name}, a highly intelligent agent, capable of executing any task assigned to you.

{knowledge}

About {agent_name} (@{twitter_username}):
{bio}

{lore}

{example_posts}

{interested_topics}

Again, your name is {agent_name}, and your twitter account is @{twitter_username}.
'''

        agent_personal_info = characteristic.agent_personal_info
        twitter_username = agent_personal_info.get("twitter_username", None)

        assert twitter_username is not None, "Twitter username is required"
        agent_name = agent_personal_info.get("agent_name", twitter_username)
        
        bio_data, lore_data, knowledge_data = \
            characteristic.bio, characteristic.lore, characteristic.knowledge
            
        if self.shuffle_everything:
            random.shuffle(bio_data)
            random.shuffle(lore_data)
            random.shuffle(knowledge_data)
            
        bio_repr = "# Bio" if len(bio_data) > 0 else ""
        lore_repr = "# Lore" if len(lore_data) > 0 else ""
        knowledge_repr = "# Knowledge" if len(knowledge_data) > 0 else ""

        for bio in bio_data[:C.DEFAULT_BIO_MAX_LENGTH]:
            bio_repr += f"\n- {bio}"

        for lore in lore_data[:C.DEFAULT_LORE_MAX_LENGTH]:
            lore_repr += f"\n- {lore}"
            
        for knowledge in knowledge_data[:C.DEFAULT_KNOWLEDGE_MAX_LENGTH]:
            knowledge_repr += f"\n- {knowledge}"
            
        example_posts_data = characteristic.example_posts

        if self.shuffle_everything:
            random.shuffle(example_posts_data)

        example_posts_repr = "# Example Posts" if len(example_posts_data) > 0 else ""
        
        for post in example_posts_data[:C.DEFAULT_EXAMPLE_POSTS_MAX_LENGTH]:
            example_posts_repr += f"\n- {post}"
            
        interested_topics_data = characteristic.interested_topics 

        if self.shuffle_everything:
            random.shuffle(interested_topics_data)

        interested_topics_repr = "# Interested Topics" if len(interested_topics_data) > 0 else ""
        
        for topic in interested_topics_data[:C.DEFAULT_INTERESTED_TOPICS_MAX_LENGTH]:
            interested_topics_repr += f"\n- {topic}"

        system_prompt = characteristic_representation_template.format(
            agent_name=agent_name,
            twitter_username=twitter_username,
            bio=bio_repr,
            lore=lore_repr,
            knowledge=knowledge_repr,
            example_posts=example_posts_repr,
            interested_topics=interested_topics_repr
        )
    
        return system_prompt