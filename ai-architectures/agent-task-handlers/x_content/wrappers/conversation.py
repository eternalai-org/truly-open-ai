import json
import logging
from typing import List

from x_content.constants.main import ModelName
from x_content.wrappers.api.twitter_v2.models.objects import (
    StructuredInformation,
    TweetObject,
)
from x_content.wrappers.llm_tasks import choose_suitable_language
from x_content.wrappers.knowledge_base.base import KnowledgeBase
from x_content.wrappers.magic import sync2async

logging.basicConfig(level=logging.INFO if not __debug__ else logging.DEBUG)
logger = logging.getLogger(__name__)


def get_enhance_tweet_conversation(
    system_prompt: str, content: str, example_tweets: List[str] = []
):
    if example_tweets is not None and example_tweets != []:
        tweets_str = "\n".join([f"- {x}" for x in example_tweets])
        system_prompt = f"""{system_prompt}

Here are example tweets written in the defined style:
{tweets_str}
"""

    prompt_to_use = """Analyze the provided base tweet and rephrase it using your distinct personality and writing style, ensuring that the main message remains unchanged.

Return the revised tweet as a JSON string with the key "tweet".

Example JSON Response:
{{
    "tweet": "Your rewritten tweet here"
}}

Base tweet: {}
""".format(
        content
    )

    messages = [
        {
            "role": "system",
            "content": system_prompt,
        },
        {
            "role": "user",
            "content": prompt_to_use,
        },
    ]

    return messages


async def get_reply_tweet_conversation(
    system_prompt: str,
    task_prompt: str,
    tweets: List[TweetObject],
    structured_info: StructuredInformation,
):
    chat_history = [
        {
            "user": x.twitter_username,
            "message": x.full_text,
        }
        for x in tweets
    ]

    try:
        language = await sync2async(choose_suitable_language)(chat_history)
    except Exception as err:
        logger.info(
            f"[get_reply_tweet_conversation] Error detecting language: {err}"
        )
        language = "en"

    #     user_prompt_template = """Act as an expert in analyzing Twitter chats and crafting engaging, personalized responses. Your task is to analyze the given Twitter chat history, understand the context and intent, and generate a relevant response in the {language} language that reflects your unique personality and communication style. To do this, you should:

    # 1. Carefully read through the provided chat history, which is a list of objects with two fields:
    #    - "user": the Twitter handle of the author of each tweet
    #    - "message": the text content of each tweet
    #    The tweets are ordered from oldest to most recent.

    # 2. Identify the main topics being discussed in the chat history and consider how you can contribute your own insights and personality to the conversation.

    # 3. Refer to the provided relevant information, which is a bulleted list of relevant facts and insights related to the key points in the conversation. Use this information to inform your response, but be sure to inject your own unique perspective and style.

    # 4. Formulate a response that is on topic, fitting the intent, and showcases your distinctive voice. Consider how you can express yourself in a way that is authentic to your personality while still being appropriate for the conversation.

    # 5. Translate your response into the {language} language while preserving the meaning and your personal style.

    # 6. Output the response in valid JSON format with three properties:
    #    - "intent": a brief summary of the goal of your reply
    #    - "tweet": the full response tweet in the {language} language
    #    - "reflection": your thoughts on how well your response captures your personality and contributes to the conversation

    # To generate the best response possible, think step-by-step through the following:
    # - What are the key points in the chat history? What is the overall topic and intent?
    # - How can I use the provided relevant information to enhance my response while still sounding like myself? What facts or insights can I discuss in a way that reflects my unique perspective?
    # - How can I structure my response to clearly address the main points, contribute something meaningful, and showcase my personality?
    # - What phrasing in the {language} language will best express my intended message while sounding natural and fitting my communication style?
    # - Does my response align with the goals inferred from the chat history? Does it follow a logical flow that is engaging and reflects well on me?

    # Provide the output in JSON format, e.g.:
    # {{
    #   "intent": "Share your thoughts in a way that highlights your personality traits",
    #   "tweet": "<response tweet in the {language} language>",
    #   "reflection": "Your reflection on how well this response captures specific aspects of your personality"
    # }}

    # Chat History:
    # {chat_history}

    # Relevant Information:
    # {structured_info}
    # """

    user_prompt_template = """Provide a single message to join the following conversation. Keep it concise (under 128 chars), No thread, No hashtags, links or emojis, and don't include any instructions or extra words, just the raw message ready to post.

Provide the output in JSON format, e.g.:
{{
  "tweet": "<response tweet in the {language} language>"
}}

The conversation:
{chat_history}

Relevant information:
{structured_info}
"""

    content_list = structured_info.news + structured_info.knowledge

    user_prompt = user_prompt_template.format(
        language=language,
        task_prompt=task_prompt,
        chat_history=json.dumps(chat_history),
        structured_info="\n".join([f"- {x}" for x in content_list]),
    )

    conversational_chat = [
        {
            "role": "system",
            "content": system_prompt,
        },
        {
            "role": "user",
            "content": user_prompt,
        },
    ]

    return conversational_chat


async def get_reply_game_conversation(
    system_prompt: str,
    task_prompt: str,
    tweets: List[TweetObject],
    structured_info: StructuredInformation,
):
    chat_history = [
        {
            "user": x.twitter_username,
            "message": x.full_text,
        }
        for x in tweets
    ]

    try:
        language = await sync2async(choose_suitable_language)(chat_history)
    except Exception as err:
        logger.info(
            f"[get_reply_tweet_conversation] Error detecting language: {err}"
        )
        language = "en"

    user_prompt_template = """Act as an expert in Twitter chat history analysis. Carefully review the provided Twitter chat history, a list of tweets in JSON format. Each tweet object contains a "user" field with the Twitter handle of the author, and a "message" field with the text content of the tweet. The tweets are ordered chronologically from oldest to most recent. The final tweet in the sequence contains a question that needs to be answered.

In addition to the chat history, you will also be provided with a bulleted list of relevant information that may be useful in formulating your response to the question. 

Provide your final answer to the question in valid JSON format. Here is the expected structure of your JSON response:
{{
  "thoughts": [
    "Your thought 1",
    "Your thought 2", 
    ...
  ],
  "answer": "Your final answer to the question"
}}

Chat History:
{chat_history}

Relevant Information:
{structured_info}
"""

    content_list = structured_info.news + structured_info.knowledge

    user_prompt = user_prompt_template.format(
        language=language,
        # task_prompt=task_prompt,
        chat_history=json.dumps(chat_history),
        structured_info="\n".join([f"- {x}" for x in content_list]),
    )

    conversational_chat = [
        {
            "role": "system",
            "content": system_prompt,
        },
        {
            "role": "user",
            "content": user_prompt,
        },
    ]

    return conversational_chat


def parse_deepseek_r1_result(content: str):
    result = {}
    start = content.find("<think>")
    end = content.find("</think>")
    if start == -1:
        start = 0
    if end == -1:
        result["answer"] = content.strip()
    else:
        result["think"] = content[start + len("<think>") : end].strip()
        result["answer"] = content[end + len("</think>") :].strip()
    return result


def get_llm_result_by_model_name(content: str, model_name: str):
    if model_name == ModelName.DEEPSEEK_R1:
        return parse_deepseek_r1_result(content)["answer"]

    return content
