from dotenv import load_dotenv

from x_content.wrappers.gcs import download_file_from_gcs_bucket
from x_content.wrappers.gcs import file_exists_in_gcs_bucket
from x_content.wrappers.gcs import upload_ds_to_gcs; load_dotenv()

import os
import json
import logging
import requests
import json
import time
import re
from typing import List
from langchain_openai import ChatOpenAI
from langchain_core.prompts import ChatPromptTemplate
from x_content.utils import is_local_env
from x_content.wrappers.magic import helpful_raise_for_status
from x_content.wrappers.magic import retry
from x_content.wrappers.log_decorators import log_function_call
logger = logging.getLogger(__name__)

from x_content import constants as const

DEFAULT_LLM_BASE_URL = const.SELF_HOSTED_LLAMA_405B_URL
DEFAULT_LLM_MODEL_ID = const.SELF_HOSTED_LLAMA_405B_MODEL_IDENTITY
DEFAULT_LLM_API_KEY = const.SELF_HOSTED_LLAMA_API_KEY
TOTAL_MAX_TWEETS = 10000
FULL_TEXT_API_BATCH_SIZE = 100
NUMBER_OF_PULLED_TWEETS_PER_LOOP = 25
DEFAULT_TEMPERATURE = 0.7
DEFAULT_SYSTEM_PROMPT = """
You are a creative, concise, and insightful assistant skilled at crafting impactful tweets. Your task is to create a tweet based on the provided information and context. 

The tweet must:
- Be informative, engaging, and directly relevant to the user query.
- Be concise and captivating, adhering to character limits while maximizing clarity and impact.
- Use compelling language to draw attention, avoid redundancy, and maintain a professional tone.

Output the response in a stringified JSON format with the key "tweet" containing the generated tweet.

Generate exactly one tweet per query, ensuring it is polished and effective.
"""
DEFAULT_TWITTER_URL = const.TWITTER_API_URL
DEFAULT_TWITTER_API_KEY = const.TWITTER_API_KEY
PROMPT_TEMPLATE = ChatPromptTemplate.from_messages(
    [
        (
            "system",
            "{system_prompt}",
        ),
        ("user", "{question}"),
    ]
)

    
BASE_COLLECTION_NAME = "base_knowledge"
 
@log_function_call
def create_llm(
    base_url: str = DEFAULT_LLM_BASE_URL, 
    model_id: str = DEFAULT_LLM_MODEL_ID, 
    api_key: str = DEFAULT_LLM_API_KEY, 
    temperature = DEFAULT_TEMPERATURE
) -> ChatOpenAI:
    """
    Initializes and returns a ChatOpenAI instance with the provided parameters.

    Args:
        base_url (str): The base URL for the LLM service.
        model_id (str): The model identifier for the LLM.
        api_key (str): The API key for authenticating with the LLM service.

    Returns:
        ChatOpenAI: An instance of the ChatOpenAI class configured with the provided parameters.
    """
    return ChatOpenAI(
        model=model_id,
        api_key=api_key,
        base_url=base_url,
        temperature=temperature
    )

@log_function_call
def get_info(user_name: str) -> tuple:
    """
    Retrieves user information from Twitter based on the provided username.

    Args:
        user_name (str): The Twitter username to retrieve information for.

    Returns:
        tuple: A tuple containing the user's name, username, and description.
    """
    url = f"{DEFAULT_TWITTER_URL.rstrip('/')}/user/by/username/{user_name}"
    headers = {
        "Authorization": f"Bearer {DEFAULT_TWITTER_API_KEY}"
    }
    try:
        resp = requests.get(url, headers=headers)
        helpful_raise_for_status(resp)
        resp_json = resp.json()
        name = resp_json["data"]["name"]
        username = resp_json["data"]["username"]
        description = resp_json["data"]["description"]
        return name, username, description

    except requests.exceptions.RequestException as e:
        logger.error(f"[get_info] Request failed: {e}")
        return "", "", ""

    except json.JSONDecodeError as e:
        logger.error(f"[get_info] JSON decode error: {e}")
        return "", "", ""
    
    except KeyError as e:
        logger.error(f"[get_info] Key error: {e}")
        return "", "", ""
    
    except Exception as e:
        logger.error(f"[get_info] An unexpected error occurred: {e}")
        return "", "", ""


@log_function_call
def generate_context_from_tweets(tweets: list) -> str:
    """
    Generates a context string from a list of tweets.

    Args:
        tweets (list): A list of tweets to generate the context from.

    Returns:
        str: A string containing the context derived from the tweets.
    """
    if len(tweets) <= 10:
        return "\n".join(tweets)
    sorted_tweets = sorted(tweets, key=len)
    context = sorted_tweets[:5] + sorted_tweets[-5:]
    return "\n".join(context)
    

@log_function_call
def file_exists_in_local(file_name):
    return os.path.exists(f".tweets/{file_name}")        

@log_function_call
def save_tweets_to_local(tweets, file_name):
    os.makedirs(".tweets", exist_ok=True)
    with open(f".tweets/{file_name}", "w") as f:
        f.write(json.dumps(tweets))

def is_long_tweet(tweet: str) -> bool:
    # Regular expression pattern to match ellipsis followed by a URL
    pattern = r'\u2026 \b(https?://[^\s]+)'
    
    # re.search() will find a match anywhere in the string,
    # while re.match() would only check if it matches at the start.
    # Using re.search() here but we could also use re.match() with
    # a pattern that checks for end-of-string.
    # The pattern \b ensures the URL is a separate word.
    
    # re.compile() can compile the regular expression pattern for better performance
    # if this function is called many times.
    regex = re.compile(pattern)
    
    # Check if the pattern is found at the end of the string
    if regex.search(tweet) and tweet.endswith(regex.search(tweet).group(0)):
        return True
    else:
        return False

@log_function_call
def get_full_text_by_tweet_ids(tweet_ids: List[str]) -> List[str]:
    result = {}
    try:
        base_url = const.TWITTER_API_URL
        twitter_api_key = const.TWITTER_API_KEY
        session = requests.Session()
        session.params = {"api_key": twitter_api_key}

        resp = session.get(f"{base_url}/tweets/v1", params={"ids": ','.join(tweet_ids)})
        helpful_raise_for_status(resp)
        data = resp.json()["result"]

        for id in tweet_ids:
            if id in data:
                tweet = data[id]["Tweet"]
                if "note_tweet" in tweet:
                    result[id] = tweet["note_tweet"]["text"]
    except Exception as err:
        logger.error(f"[get_full_text_by_tweet_ids] Error retrieving text of tweets {err}")
    return result

@log_function_call
def pull_tweets(twitter_ids: List[str]):
    collected_tweets = []
    base_url = const.TWITTER_API_URL
    twitter_api_key = const.TWITTER_API_KEY
    session = requests.Session()
    session.params = {
        "api_key": twitter_api_key
    }

    max_loops = TOTAL_MAX_TWEETS // NUMBER_OF_PULLED_TWEETS_PER_LOOP

    for twitter_id in twitter_ids:
        tweets = []
        file_name = f"{twitter_id}.json"

        if not is_local_env() and file_exists_in_gcs_bucket(file_name):
            file_path = download_file_from_gcs_bucket(file_name)
            try:
                with open(file_path, "r") as f:
                    tweets = json.load(f)
            except Exception as e:
                logger.error(f"[pull_tweets] Failed to load tweets from {file_path}: {e}")
                return [], e
            finally:
                if os.path.exists(file_path):
                    os.remove(file_path)
        elif is_local_env() and file_exists_in_local(file_name):
            file_path = f".tweets/{file_name}"
            try:
                with open(file_path, "r") as f:
                    tweets = json.load(f)
            except Exception as e:
                logger.error(f"[pull_tweets] Failed to load tweets from {file_path}: {e}")
                return [], e
        else:
            url = f"{base_url}/tweets/{twitter_id}/all"
            pagination_token = None

            tweets_info = {}
            for loop_count in range(max_loops):
                logger.info(f"[pull_tweets] Pulling data from {twitter_id}, loop {loop_count + 1}")
                params = session.params.copy()
                params["max_results"] = 100
                if pagination_token:
                    params["pagination_token"] = pagination_token

                try:
                    def get_tweets():
                        resp = session.get(url, params=params)
                        helpful_raise_for_status(resp)
                        return resp.json()
                    
                    resp_json = retry(get_tweets, max_retry=3, first_interval=60, interval_multiply=2)()
                    crawled = resp_json.get("result", {}).get("data", [])

                    if not crawled:
                        logger.info(f"[pull_tweets] No tweets found for {twitter_id}")
                        break
                    for item in crawled:
                        id = item.get("id", "")
                        text = item.get("text", "")
                        reference_tweets = (
                            [] if item["referenced_tweets"] is None else item["referenced_tweets"]
                        )
                        if len(reference_tweets) > 0:
                            continue
                        tweets_info[id] = {
                            "id": id,
                            "text": text,
                        }
                    pagination_token = resp_json.get("result", {}).get("meta", {}).get("next_token")
                    if not pagination_token:
                        break
                except requests.exceptions.RequestException as e:
                    logger.error(f"[pull_tweets] Error pulling tweets for {twitter_id}: {e}")
                    return [], e
                time.sleep(1)

            long_tweet_ids = []
            for info in tweets_info.values():
                if is_long_tweet(info["text"]):
                    long_tweet_ids.append(info["id"])

            for i in range(0, len(long_tweet_ids), FULL_TEXT_API_BATCH_SIZE):
                ids_batch = long_tweet_ids[i:i+FULL_TEXT_API_BATCH_SIZE]
                result: dict[str, str] = get_full_text_by_tweet_ids(ids_batch)
                for id, text in result.items():
                    tweets_info[id]["text"] = text

            tweets = [x["text"] for x in tweets_info.values()]
            if tweets:
                if is_local_env():
                    save_tweets_to_local(tweets, file_name)
                else:
                    upload_ds_to_gcs(tweets, file_name)

        collected_tweets.extend(tweets)

    if not collected_tweets:
        logger.error("[pull_tweets] No tweets were collected.")

    logger.info(f"[pull_tweets] Total number of tweets: {len(collected_tweets)}")
    return collected_tweets, None
        
def names_list_to_string(lst, chr=''):
    lst = [chr + x + chr for x in lst]
    if not lst:
        return ""
    elif len(lst) == 1:
        return lst[0]
    elif len(lst) == 2:
        return f"{lst[0]} and {lst[1]}"
    else:
        return ", ".join(lst[:-1]) + f", and {lst[-1]}"