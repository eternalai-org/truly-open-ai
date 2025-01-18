from pymilvus import MilvusClient
from functools import lru_cache

@lru_cache(maxsize=128)
def get_reusable_milvus_client(uri: str):
    return MilvusClient(uri=uri)
