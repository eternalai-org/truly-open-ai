from x_content.constants import AgentUsername, MissionChainState, ModelName, ToolSet
from x_content.llm.base import OpenAILLMBase
from x_content.llm.eternal_ai import ASyncBasedEternalAI
from x_content.models import ReasoningLog, ToolDef
from x_content.toolcall import dynamic_toolcall, wrapped_external_apis
from x_content.wrappers.knowledge_base.base import KnowledgeBase
from x_content.wrappers.knowledge_base.eternals_kb import EternalKnowledgeBase
from x_content.wrappers import telegram, junk_tasks
from x_content import constants as const

import random
import logging 
import json
import requests
from typing import List
from x_content import constants as const
from json_repair import repair_json

from x_content.wrappers.postprocess import StringProcessor
from x_content.wrappers.twin_agent import get_random_example_tweets

logger = logging.getLogger(__name__)

def a_move_state(log: ReasoningLog, state: MissionChainState, reason: str):
    log.system_message = reason
    log.state = state
    return log

async def a_move_state(log: ReasoningLog, state: MissionChainState, reason: str):
    log.system_message = reason
    log.state = state
    return log

def notify_status(log: ReasoningLog):
    nav = f'<b>Request-ID</b>: {log.id}'

    if log.state == MissionChainState.NEW:    
        info = f'<i><b>Ref-ID</b>: {log.meta_data.ref_id};\n{nav}</i>'
        msg = f'<strong>Received a new task {log.task} using toolset {log.toolset} for {log.meta_data.twitter_username}</strong>\nTraceback info:\n{info}'

    elif log.state in [MissionChainState.ERROR, MissionChainState.DONE]:
        info = f'<i><b>Ref-ID</b>: {log.meta_data.ref_id};\n{nav}'
        system_message = log.system_message
        msg = f'<strong>Task {log.task} using toolset {log.toolset} for {log.meta_data.twitter_username} finished with state {log.state}</strong>\nTraceback info:\n{info}\nSystem message: {system_message}'

    else:
        return

    telegram.send_message(
        twitter_username="junk_notifications",
        message_to_send=msg,
        schedule=True,
        room=telegram.TELEGRAM_TASK_IO_NOTIFICATION_ROOM,
        fmt='HTML'
    )
    

_alert_template = '''
<strong>Alert</strong>
<i>Task {log.task} using toolset {log.toolset} for {log.meta_data.twitter_username} raised an alert</i>
{info}
<i>Reason: </i>
<pre>\n{reason}\n</pre>  
'''

def send_alert(log: ReasoningLog, reason: str):
    global _alert_template

    nav = f'<b>Request-ID</b>: {log.id};'
    info = f'<i><b>Ref-ID</b>: {log.meta_data.ref_id};\n{nav}</i>'

    msg = _alert_template.format(log=log, info=info, reason=reason)
    telegram.send_message(
        twitter_username="junk_notifications",
        message_to_send=msg,
        schedule=True,
        room=telegram.TELEGRAM_ALERT_ROOM,
        fmt='HTML'
    )
        
def format_prompt_v2(log: ReasoningLog, tools: List[ToolDef]):
    template_prompt = """
You have access to the following tools to get information or take actions:
{tools}

Your reply must be a single JSON object with exactly three keys described as follows.
thought: your own thought about the next step, reflecting your unique persona.
action: must be one of {toolnames}.
action_input: provide the necessary parameters for the chosen action, separating multiple parameters with the | character. For example, if there are two parameters "abc" and "123", the action_input field should be "abc|123".

OR with exactly two keys as follows.
thought: your final thought to conclude.
final_answer: your conclusion.

About you:
{base_system_prompt}

Again, only return a single JSON!
"""

    tool_names = ", ".join([tool.name for tool in tools])
    base_tool_str = "\n".join(e.prototype() for e in tools)

    return template_prompt.format(
        tools=base_tool_str,
        toolnames=tool_names,
        base_system_prompt=log._system_prompt_,
    )


def render_reply_thread_conversation(log: ReasoningLog, idx: int):
    assert idx < 0 or idx >= len(log.execute_info["tweets"]), "Index out of range"

    tweets = log.execute_info["tweets"][idx]["full_context"]
    knowledge = log.execute_info["tweets"][idx].get("knowledge_v2")

    chat_history = [
        {
            "user": x["twitter_username"],
            "message": x["full_text"],
        }
        for x in tweets
    ]
    language = junk_tasks.choose_suitable_language(chat_history)

    user_prompt_template = """Provide a single message to join the following conversation in {} language. Keep it concise (under 128 chars), NO hashtags, links or emojis, and don't include any instructions or extra words, just the raw message ready to post.

Chat History: {}"""

    user_prompt = user_prompt_template.format(language, json.dumps(chat_history))
    if knowledge is not None:
        user_prompt += "\n\nRelevant Knowledge: {}".format(json.dumps(knowledge))

    conversational_chat = []
    conversational_chat.append(
        {
            "role": "system",
            "content": log._system_prompt_,
        }
    )
    conversational_chat.append(
        {
            "role": "user",
            "content": user_prompt,
        }
    )

    return conversational_chat


def render_reply_game_conversation(log: ReasoningLog, idx: int):
    #TODO: @harvey implement LLM call for reply game
    assert idx < 0 or idx >= len(log.execute_info["tweets"]), "Index out of range"

    tweets = log.execute_info["tweets"][idx]["full_context"]
    knowledge = log.execute_info["tweets"][idx].get("knowledge_v2")

    chat_history = [
        {
            "user": x["twitter_username"],
            "message": x["full_text"],
        }
        for x in tweets
    ]
    language = junk_tasks.choose_suitable_language(chat_history)

    user_prompt_template = """Provide a single message to join the following conversation in {} language. Keep it concise (under 128 chars), NO hashtags, links or emojis, and don't include any instructions or extra words, just the raw message ready to post.

Chat History: {}"""

    user_prompt = user_prompt_template.format(language, json.dumps(chat_history))
    if knowledge is not None:
        user_prompt += "\n\nRelevant Knowledge: {}".format(json.dumps(knowledge))

    conversational_chat = []
    conversational_chat.append(
        {
            "role": "system",
            "content": log._system_prompt_,
        }
    )
    conversational_chat.append(
        {
            "role": "user",
            "content": user_prompt,
        }
    )

    return conversational_chat


def get_system_prompt(log: ReasoningLog):
    if log.model == ModelName.INTELLECT_10B:
        lines = log.agent_meta_data.persona.split("\n\n")
        if lines[1].startswith("You have a token. Your token name is Eternal Intellect."):
            lines = lines[:1] + lines[2:]
        if lines[0].startswith("Your Twitter name is Eternal Intellect. Your Twitter username is @0xIntellect."):
            lines = lines[1:]
        system_prompt = "\n\n".join(lines)
        return system_prompt
    else:
        return log._system_prompt_

def get_system_prompt_with_random_example_tweets(log: ReasoningLog):
    system_prompt = get_system_prompt(log)
    example_tweets = get_random_example_tweets(log.meta_data.knowledge_base_id)
    if len(example_tweets) > 0:
        tweets_str = '\n'.join([f"- {x}" for x in example_tweets])
        system_prompt += f"\n\nHere are example tweets written in the defined style:\n{tweets_str}"
    return system_prompt

def render_rewrite_tweet_conversation(log: ReasoningLog, content: str):
    if log.model == ModelName.INTELLECT_10B:
        prompt_to_use = "Write a tweet in 256 characters using your specific style. Content to follow:\n\n{}".format(content)
    else:
        prompt_to_use = """
Craft a tweet based on the provided content that aligns with your personality. 
- No introduction. 
- No gifs. 
- No hashtags. 
- No emojis.

Return the response as a stringified JSON with the key "tweet".

Example JSON Response:
{{
    "tweet": "Your tweet here"
}}

Content to follow: {}
""".format(content)
        
    conversational_chat = []
    conversational_chat.append(
        {
            "role": "system",
            "content": get_system_prompt_with_random_example_tweets(log),
        }
    )
    conversational_chat.append(
        {
            "role": "user",  # Somehow it FUCKING WORK?!?!
            "content": prompt_to_use,
        }
    )

    return conversational_chat


from x_content.wrappers.api import twitter_v2

def create_twitter_auth_from_reasoning_log(log: ReasoningLog):
    return twitter_v2.TwitterRequestAuthorization(
        agent_contract_id=log.meta_data.agent_contract_id,
        twitter_id=log.meta_data.twitter_id,
        twitter_username=log.meta_data.twitter_username,
        chain_id=int(log.meta_data.chain_id),
        knowledge_id=log.meta_data.knowledge_base_id,
        ref_id=log.meta_data.ref_id,
        request_id=log.id,
        task=log.task,
        toolset=log.toolset,
        kn_base=create_kn_base(log),
        prompt=log.prompt,
    )

def get_mentioned_usernames(tweet_object):
    """Mock function to get mentioned usernames.
    TODO:
    - Implement actual Twitter API call to get all mentioned usernames
    - Parse response and extract usernames from mentions
    - must not contain @nobullshit
    """
    # usernames = ["agent1", "agent2"]
    username_to_remove = AgentUsername.TWEETER_NAME_CRYPTOCOMIC_AI
    mentions = tweet_object['mentions']
    tweet_object['mentions'] = [mention for mention in mentions if mention['username'] != username_to_remove]
    usernames = [mention['username'] for mention in tweet_object['mentions']]
    return usernames

def is_create_token_tweet(tweet_info):
    try:
        result = junk_tasks.is_analyzing_token_conversation(tweet_info["tweet_object"]["full_text"])
        return result
    except Exception as err:
        logger.error(f"[is_create_token_tweet] Error analyzing conversation: {err}")
        return None

# Filter tweets that contain emoji game patterns like:
# Example 1: "🎮 Let's play a game! 🎲"
# Example 2: "🎯 Guess the number between 1-10 🎲"
# Example 3: "🎪 Riddle time! 🤔"
def is_create_game_tweet(tweet_object):
    text = tweet_object["full_text"]
    agent_usernames = get_mentioned_usernames(tweet_object)

    has_two_or_more_usernames = len(agent_usernames) >= 2

    # contains_emoji = any(emoji in text for emoji in GAME_EMOJIS)
    emoji_count = sum(1 for emoji in const.GAME_EMOJIS if emoji in text)
    contains_two_emojis = emoji_count >= 2
    text = StringProcessor(text)\
            .remove_tags()\
            .remove_mentions()\
            .remove_urls()\
            .remove_emojis()\
            .strip_head_and_tail_white_string()\
            .get_text()
    logger.info(f"[is_create_game_tweet] text after postprocess {text}")
    contains_keyword = any(keyword in text.lower() for keyword in const.GAME_KEYWORDS)
    
    return contains_two_emojis and contains_keyword and has_two_or_more_usernames

def is_create_game_tweet_id(tweet_id: str) -> bool:
    resp = twitter_v2.get_tweet_info_from_tweet_id(tweet_id)
    if resp.is_error():
        return False
    return is_create_game_tweet(resp.data.tweet_info.tweet_object.to_dict())


def magic_toolset(log: ReasoningLog, llm: OpenAILLMBase):
    if len(log.tool_list) > 0:
        toolcall = dynamic_toolcall.DynamicToolcall(log.tool_list)

    else:
        toolcall = wrapped_external_apis.LiveXDB(
            auth=create_twitter_auth_from_reasoning_log(log),
            agent_config=log.agent_meta_data,
            llm=llm
        )

    return toolcall


def model_dependent_frequency_penalty(log: ReasoningLog) -> float:
    frequency_penalty = 0

    if log.model == ModelName.INTELLECT_10B:
        frequency_penalty = 0.1

    return frequency_penalty


def create_llm(log: ReasoningLog):
    return ASyncBasedEternalAI(
        max_tokens=const.DEFAULT_MAX_OUTPUT_TOKENS,
        temperature=const.DEFAULT_TEMPERATURE,
        base_url=const.BACKEND_API,
        api_key=const.BACKEND_AUTH_TOKEN,
        frequency_penalty=model_dependent_frequency_penalty(log),
        model=log.model,
        seed=random.randint(1, int(1e9)),
        chain_id=log.meta_data.chain_id,
        agent_contract_id=log.meta_data.agent_contract_id,
        metadata=log.meta_data.model_dump(),
    )



def create_kn_base(log: ReasoningLog) -> KnowledgeBase:
    return EternalKnowledgeBase(
        default_top_k=5,
        similarity_threshold=0.5,
        base_url=const.BACKEND_API,
        api_key=const.BACKEND_AUTH_TOKEN,
        kbs=log.agent_meta_data.kb_agents,
    )
