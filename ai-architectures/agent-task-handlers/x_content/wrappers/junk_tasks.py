import random
from openai import BaseModel
import requests
import logging
import json
from typing import Optional

from x_content.wrappers.postprocess import post_process_tweet
from . import redis_wrapper

from x_content.llm.local import SyncBasedEternalAI
from x_content.wrappers.magic import helpful_raise_for_status
from x_content.wrappers.magic import retry
from .log_decorators import log_function_call
from . import constants as const
from json_repair import repair_json

logging.basicConfig(level=logging.INFO if not __debug__ else logging.DEBUG)
logger = logging.getLogger(__name__)

MAX_TEXT_LENGTH = 10000
MIN_TEXT_LENGTH_TO_SUMMARIZE = 100

from functools import lru_cache

def get_image_description(url: str):
    base_url = f"{const.VISION_API_URL}/v1/chat/completions"
    headers = {
        "Authorization": f"Bearer {const.VISION_API_KEY}"
    }

    payload = {
        "messages": [
            {
                "role": "user",
                "content": [
                    {
                        "type": "text",
                        "text": "Analyze this image and describe its semantic meaning in detail with a single paragraph. Include the implied context, emotions, relationships between elements, possible symbolic interpretations, and the overall message or story conveyed by the image."
                    },
                    {
                        "type": "image_url",
                        "image_url": { "url": url }
                    }
                ]
            }
        ],
        "max_tokens": 300,
        "model": const.VISION_API_MODEL
    }

    resp = requests.post(base_url, headers=headers, json=payload)

    if resp.status_code != 200:
        logger.error(f"Failed to get image description: {resp.text}")
        return None

    return resp.json()['choices'][0]['message']['content']

def choose_suitable_language(chat_history: list) -> str:
    """
    Determines the most suitable language for the input chat history.

    Args:
        chat_history (list): A list of strings representing the chat conversation.

    Returns:
        str: The ISO 639-1 language code for the detected language. Defaults to 'en' if an error occurs.
    """
    valid_languages = ['aa', 'ab', 'af', 'ak', 'am', 'ar', 'an', 'as', 'av', 'ae', 'ay', 'az', 'ba', 'bm', 'be', 'bn', 'bi', 'bo', 'bs', 'br', 'bg', 'ca', 'cs', 'ch', 'ce', 'cu', 'cv', 'kw', 'co', 'cr', 'cy', 'da', 'de', 'dv', 'dz', 'el', 'en', 'eo', 'et', 'eu', 'ee', 'fo', 'fa', 'fj', 'fi', 'fr', 'fy', 'ff', 'gd', 'ga', 'gl', 'gv', 'gn', 'gu', 'ht', 'ha', 'sh', 'he', 'hz', 'hi', 'ho', 'hr', 'hu', 'hy', 'ig', 'io', 'ii', 'iu', 'ie', 'ia', 'id', 'ik', 'is', 'it', 'jv', 'ja', 'kl', 'kn', 'ks', 'ka', 'kr', 'kk', 'km', 'ki', 'rw', 'ky', 'kv', 'kg', 'ko', 'kj', 'ku', 'lo', 'la', 'lv', 'li', 'ln', 'lt', 'lb', 'lu', 'lg', 'mh', 'ml', 'mr', 'mk', 'mg', 'mt', 'mn', 'mi', 'ms', 'my', 'na', 'nv', 'nr', 'nd', 'ng', 'ne', 'nl', 'nn', 'nb', 'no', 'ny', 'oc', 'oj', 'or', 'om', 'os', 'pa', 'pi', 'pl', 'pt', 'ps', 'qu', 'rm', 'ro', 'rn', 'ru', 'sg', 'sa', 'si', 'sk', 'sl', 'se', 'sm', 'sn', 'sd', 'so', 'st', 'es', 'sq', 'sc', 'sr', 'ss', 'su', 'sw', 'sv', 'ty', 'ta', 'tt', 'te', 'tg', 'tl', 'th', 'ti', 'to', 'tn', 'ts', 'tk', 'tr', 'tw', 'ug', 'uk', 'ur', 'uz', 've', 'vi', 'vo', 'wa', 'wo', 'xh', 'yi', 'yo', 'za', 'zh', 'zu']
    system_prompt = """
    You are a neutral and unbiased assistant tasked with determining the most suitable language for tweets.
    Analyze the input text objectively and specify the single most suitable language in which the conversation predominantly occurs.
    Treat all languages equally without favoring any specific language, including English.

    If the conversation contains mixed languages, choose the language that appears most frequently or is most contextually appropriate.
    Always return a single ISO 639-1 language code as the result.

    Provide the final response in a stringified JSON format with the key 'language' and the value as the corresponding ISO 639-1 language code.

    Examples:

    Input:
    @user1: 你好，你最近怎么样？
    @user2: 我很好，谢谢！你呢？

    Output:
    {
        "language": "zh"
    }

    Input:
    @user1: Hola, ¿cómo estás?
    @user2: Estoy bien, ¡gracias! ¿Y tú?

    Output:
    {
        "language": "es"
    }

    Input:
    @user1: Bonjour, ça va?
    @user2: I'm doing well, merci!

    Output:
    {
        "language": "fr"
    }

    Input:
    @user1: Hi, good morning!
    @user2: おはようございます！

    Output:
    {
        "language": "ja"
    }
    """

    chat_history = list(map(lambda x: f'@{x["user"]}: {x["message"]}', chat_history))
    chat_history = "\n".join(chat_history)

    conversation = [
        {"role": "system", "content": system_prompt},
        {"role": "user", "content": chat_history}
    ]

    url = const.SELF_HOSTED_LLAMA_405B_URL + "/v1/chat/completions"
    headers = {
        "Content-Type": "application/json",
        "Authorization": f"Bearer {const.SELF_HOSTED_LLAMA_API_KEY}"
    }

    data = {
        "model": const.SELF_HOSTED_LLAMA_405B_MODEL_IDENTITY,
        "messages": conversation,
        "max_tokens": 100,
        "temperature": 0,
    }

    try:
        response = requests.post(url, headers=headers, json=data)
        helpful_raise_for_status(response)
        content = response.json()["choices"][0]["message"]["content"]
        language: str = repair_json(content, return_objects=True)["language"]

        if language in valid_languages:
            return language

        return "en"
    except (requests.exceptions.RequestException, json.JSONDecodeError, KeyError) as err:
        logger.error(f"[choose_suitable_language] An error occurred: {err}")
        return "en"
    except Exception as err:
        logger.error(f"[choose_suitable_language] An error occurred: {err}")
        return "en"

@log_function_call
def generate_retrieval_query(chat_history: list) -> Optional[str]:
    # Always use English because the model only supports English
    system_prompt = """
You are a helpful and intelligent assistant. Your task is to analyze the chat conversation and extract the most suitable query to retrieve relevant information from a database.

Ensure the query is clear, concise, and accurately represents the user's intent.

Your response **must** be in stringified JSON format, with the following structure:
{
  "query": "<extracted_query_string>"
}
"""
    try:
        formatted_chat_history = "\n".join(f'@{entry["user"]}: {entry["message"]}' for entry in chat_history)

        conversation = [
            {"role": "system", "content": system_prompt},
            {"role": "user", "content": formatted_chat_history}
        ]

        llama_url = const.SELF_HOSTED_LLAMA_405B_URL
        llama_api_key = const.SELF_HOSTED_LLAMA_API_KEY
        llama_model_identity = const.SELF_HOSTED_LLAMA_405B_MODEL_IDENTITY

        if not all([llama_url, llama_api_key, llama_model_identity]):
            logger.error("[generate_retrieval_query] Missing environment variables")
            return None

        url = f"{llama_url}/v1/chat/completions"
        headers = {
            "Content-Type": "application/json",
            "Authorization": f"Bearer {llama_api_key}"
        }

        data = {
            "model": llama_model_identity,
            "messages": conversation,
            "temperature": 0.01,
            "max_tokens": 1024,
        }

        response = requests.post(url, headers=headers, json=data)
        helpful_raise_for_status(response)

        content = response.json().get("choices", [{}])[0].get("message", {}).get("content", "")
        json_content = repair_json(content, return_objects=True)

        return json_content["query"]
    except Exception as err:
        logger.error(f"[generate_retrieval_query] An error occurred: {err}")
        return None

@redis_wrapper.cache_for(3600 * 24 * 30)
def is_analyzing_token_conversation(conversation: str) -> str:
    system_prompt = "You are a helpful assistant."

    prompt_to_use = """
This is the user conversation: "{}".

From this conversation determine if the user is requesting to analyze a specific token. Look for a clear, direct mention of a token symbol (e.g., "$XYZ").

If yes, extract or generate the following information:

Answer ("yes" or "no")
Token symbol

Return a JSON response with the following format:
{{"answer": "yes/no", "symbol": ""}}

Respond with only the JSON string, without any additional explanation.
""".format(conversation)
        
    messages = [
        {
            "role": "system",
            "content": system_prompt,
        },
        {
            "role": "user",
            "content": prompt_to_use,
        }
    ]

    def run_llm():
        model = SyncBasedEternalAI(
            max_tokens=const.DEFAULT_MAX_OUTPUT_TOKENS,
            temperature=0.7,
            base_url=const.SELF_HOSTED_LLAMA_405B_URL + "/v1",
            api_key=const.SELF_HOSTED_LLAMA_API_KEY,
            model=const.SELF_HOSTED_LLAMA_405B_MODEL_IDENTITY,
            seed=random.randint(1, int(1e9)),
        )

        result = model.generate(messages).generations[0].text
        result = repair_json(result, return_objects=True)
        result = result["answer"]

        if result.lower() not in ["yes", "no"]:
            raise Exception("Invalid answer")

        return result.lower()
    
    result = retry(run_llm, max_retry=3, first_interval=10, interval_multiply=2)()
    return result == "yes", None

class IncludeExcludeItem(BaseModel):
    included: bool
    excluded: bool
    reason: str


class IncludeExcludeInfo(BaseModel):
    hashtags: IncludeExcludeItem
    urls: IncludeExcludeItem
    emojis: IncludeExcludeItem
    mentions: IncludeExcludeItem

@redis_wrapper.cache_for(3600 * 24 * 30)
@log_function_call
def detect_included_excluded_items(system_prompt, task_prompt) -> IncludeExcludeInfo:
    req_system_prompt = """Persona: 
Act as an expert in natural language processing and prompt engineering with deep expertise in analyzing and synthesizing complex requirements for AI systems.

Task:
Your task is to analyze a given system prompt and user prompt intended for an AI that generates tweets for Twitter. Specifically, determine whether the system or user prompt explicitly requires the inclusion or exclusion of the following items from the generated tweets:
1. Hashtags
2. URLs
3. Emojis
4. Mentions
For each item, identify whether there is a clear and explicit requirement for inclusion or exclusion (e.g., "Include hashtags" or "Do not include URLs").

Context:
The AI generating tweets is designed to respond based on both system-level instructions (guiding its behavior generally) and user-level instructions (specific to individual requests). In cases where there is conflicting information between the system and user prompts, the user prompt takes precedence.

Response Format:
Provide your response in the following JSON format:

{
  "hashtags": {
    "included": [true/false],
    "excluded": [true/false],
    "reason": "Explain why hashtags are/are not required to be included or excluded."
  },
  "urls": {
    "included": [true/false],
    "excluded": [true/false],
    "reason": "Explain why URLs are/are not required to be included or excluded."
  },
  "emojis": {
    "included": [true/false],
    "excluded": [true/false],
    "reason": "Explain why emojis are/are not required to be included or excluded."
  },
  "mentions": {
    "included": [true/false],
    "excluded": [true/false],
    "reason": "Explain why mentions are/are not required to be included or excluded."
  }
}

Instructions for Reasoning:
- Carefully examine the system prompt for explicit instructions regarding the inclusion or exclusion of hashtags, URLs, emojis, and mentions.
- Analyze the user prompt for explicit instructions related to these items.
- If there are conflicting instructions, prioritize the user prompt over the system prompt.
- Focus solely on explicit requirements stated in the text.

Output the final analysis in the JSON format provided.
"""

    req_user_prompt = f"""Here are the system prompt and task prompt:

System prompt: {system_prompt}

Task prompt: {task_prompt}
"""

    messages = [
        {
            "role": "system",
            "content": req_system_prompt,
        },
        {
            "role": "user",
            "content": req_user_prompt,
        }
    ]

    def run_llm():
        model = SyncBasedEternalAI(
            max_tokens=const.DEFAULT_MAX_OUTPUT_TOKENS,
            temperature=0.7,
            base_url=const.SELF_HOSTED_LLAMA_405B_URL + "/v1",
            api_key=const.SELF_HOSTED_LLAMA_API_KEY,
            model=const.SELF_HOSTED_LLAMA_405B_MODEL_IDENTITY,
            seed=random.randint(1, int(1e9)),
        )
        result = model.generate(messages).generations[0].text
        result = repair_json(result, return_objects=True)

        obj = IncludeExcludeInfo.model_validate(result)
        return obj
    
    obj = retry(run_llm, max_retry=3, first_interval=10, interval_multiply=2)()
    return obj

def postprocess_tweet_by_prompts(system_prompt: str, task_prompt: str, tweet: str) -> str:
    include_exclude_info: IncludeExcludeInfo = detect_included_excluded_items(system_prompt, task_prompt)

    postprocessed_tweet = post_process_tweet(
        tweet,
        keep_emojis=not include_exclude_info.emojis.excluded,
        keep_mentions=not include_exclude_info.mentions.excluded,
        keep_urls=include_exclude_info.urls.included,
        keep_hashtags=include_exclude_info.hashtags.included,
    )

    return postprocessed_tweet
