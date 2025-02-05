from x_content.wrappers.knowledge_base.base import KnowledgeBase
from x_content.wrappers.magic import helpful_raise_for_status
from x_content import constants as const
from .log_decorators import log_function_call
from typing import List
import os
import logging
import requests
import random

logger = logging.getLogger(__name__)


@log_function_call
async def search_from_db(
    kn_base: KnowledgeBase, query: str, top_k: int = 10, threshold=0.8
) -> List[str]:
    """
    Retrieves and integrates up-to-date information on topics where recent events or developments (2024 onward)
    are essential to accurately address the user's query.

    This function should only be called when the knowledge cutoff (01 Sep 2022) leaves gaps that would
    prevent a comprehensive and accurate response. Retrieved insights should be directly relevant, carefully
    vetted, and seamlessly incorporated with existing knowledge to provide a clear, precise, and well-informed answer.

    Args:
        query (str): The search query to retrieve insights for.
        collections (List[str]): The collections to search in.
        top_k (int): The number of top search results to retrieve.

    Returns:
        List[str]: A list of strings containing the consolidated summary of insights derived from the search results.
    """

    try:
        result = await kn_base.aretrieve(query, top_k, threshold)
        knowledge = list(map(lambda x: x.content, result))
        return knowledge
    except Exception as err:
        logger.error(f"[search_from_db] An error occurred: {err}")
        return []


@log_function_call
def get_random_from_collection(collection_name: str, n=10):
    headers = {"X-Token": const.RAG_SECRET_TOKEN}

    resp = requests.get(
        f"{const.RAG_API}/api/sample",
        params={"kb": collection_name, "k": n},
        headers=headers,
    )

    if resp.status_code != 200:
        return [], Exception(
            f"Failed to get random vectors from collection {collection_name}"
        )

    resp = resp.json()
    data = resp["result"]

    return [x["content"] for x in data], None


@log_function_call
def get_random_from_collections(collections: List[str], n=10):
    all_tweets = []

    for collection in collections:
        tweets, err = get_random_from_collection(collection, n=n)

        if err is None:
            all_tweets.extend(tweets)

    all_tweets = random.sample(all_tweets, min(10, len(all_tweets)))
    return all_tweets, None


@log_function_call
def insert_to_db(
    collection_name: str, contents: List[str], batch_size: int = 8
):
    headers = {"X-Token": const.RAG_SECRET_TOKEN}

    resp = requests.post(
        f"{const.RAG_API}/api/insert",
        json={"file_urls": [], "texts": contents, "kb": collection_name},
        headers=headers,
    )

    if resp.status_code != 200:
        logger.error(f"[insert_to_db] Insert error: {resp.text}")
        return False

    return True
