
from typing import List
from x_content.utils import parse_knowledge_ids
from x_content.wrappers import rag_search

def get_random_example_tweets(knowledge_base_id: str, n = 10) -> List[str]:
    knowledge_ids = parse_knowledge_ids(knowledge_base_id)
    example_tweets, err = rag_search.get_random_from_collections(knowledge_ids, n = n)
    return example_tweets
