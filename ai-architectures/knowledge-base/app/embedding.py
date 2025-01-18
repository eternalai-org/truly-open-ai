from .models import EmbeddingModel, SimMetric
from typing import List
from . import constants as const
from transformers import AutoTokenizer    

def get_embedding_models() -> List[EmbeddingModel]:
    return [
        EmbeddingModel(
            name=const.SELF_HOSTED_EMBEDDING_MODEL_ID,
            base_url=const.SELF_HOSTED_EMBEDDING_URL,
            tokenizer=const.TOKENIZER_MODEL_ID,
            dimension=const.MODEL_DIMENSION,
            prefer_metric=SimMetric.IP,
            normalize = False
        )
    ]

def get_default_embedding_model() -> EmbeddingModel:
    return get_embedding_models()[0]

def get_embedding_model_api_key(_: EmbeddingModel) -> str:
    return const.SELF_HOSTED_LLAMA_API_KEY

from functools import lru_cache

@lru_cache(maxsize=8)
def get_tokenizer(model: EmbeddingModel):
    return AutoTokenizer.from_pretrained(model.tokenizer)

@lru_cache(maxsize=8)
def get_tokenizer_by_model_id(modelid: str):
    for model in get_embedding_models():
        if model.name == modelid:
            return get_tokenizer(model)
        
    return None