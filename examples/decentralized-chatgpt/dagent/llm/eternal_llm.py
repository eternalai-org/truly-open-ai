from .base_llm import AsyncChatCompletion
from dagent.registry import RegistryCategory, register_decorator
from typing import List, Dict
import logging
from dagent.models import InferenceResult, InferenceState, OnChainData
import requests
from dagent import constant as C

logger = logging.getLogger(__name__)

# TODO: convert the openai standard to async standard of eternal AI 
@register_decorator(RegistryCategory.LLM)
class EternalAIChatCompletion(AsyncChatCompletion):
    DEFAULT_PARAMS = {
        "top_p": 1.0,
        "presence_penalty": 0.0,
        "n": 1,
        "logit_bias": None,
        "frequency_penalty": 0.0,
    }

    def __call__(self, _messages: List[Dict[str, str]], stop: List[str]=[], override_kwargs: dict={}): 

        last_onchain_data = None

        for _try in range(self.max_retries + 1):
            if _try > 0:
                logger.warning("Retrying {} out of {}".format(_try, self.max_retries))

            payload = {
                **self.model_kwargs,
                **self.DEFAULT_PARAMS,
                "model": self.model_name,
                "chain_id": self.chain_id,
                "messages": _messages,
                "temperature": self.temperature,
                "max_tokens": self.max_tokens,
                "stop": stop
            }

            for k, v in override_kwargs.items():
                payload[k] = v
                
            url = self.openai_api_base + "/v1/chat/completions"

            resp = self.http_session.post(
                url, 
                json=payload
            )
            
            resp_json = resp.json()
            last_onchain_data=resp_json.get('onchain_data')

            if resp.status_code == 200:
                return self.commit(InferenceResult(
                    id=self.generate_uuid(),
                    state=InferenceState.DONE,
                    result=resp_json['choices'][0]['message']['content'],
                    onchain_data=OnChainData.model_validate(last_onchain_data)
                ))
                
            logger.error("Failed to get a response from the model. Status code: {}; Text: {}; URL: {}".format(resp.status_code, resp.text, url, last_onchain_data))

        return self.commit(InferenceResult(
            id=self.generate_uuid(),
            state=InferenceState.ERROR,
            error="Failed to get a response from the model",
            onchain_data=OnChainData.model_validate(last_onchain_data)
        ))

    def __init__(
        self, 
        max_tokens: int,
        model_kwargs: dict, 
        temperature: float,
        max_retries: int, 
        eternal_api_base: str=C.ETERNALAI_URL, 
        eternal_api_key: str=C.ETERNALAI_API_KEY,
        model_name: str=C.ETERNAL_MODEL_NAME, 
        eternal_chain_id: str=C.ETERNAL_CHAIN_ID,
        *args, **kwargs
    ):
        super().__init__()

        assert eternal_api_key is not None, "eternalai_api_key is not provided and ETERNALAI_URL is not set in the environment"
        assert eternal_api_base is not None, "eternalai_api_base is not provided and ETERNALAI_API_KEY is not set in the environment"
        assert model_name is not None, "model_name is not provided" 

        self.eternal_api_key = eternal_api_key

        self.openai_api_base = eternal_api_base.rstrip("/")
        self.model_name = model_name
        self.model_kwargs = model_kwargs
        self.temperature = temperature
        self.max_tokens = max_tokens
        self.max_retries = max_retries
        self.chain_id = eternal_chain_id
        self.http_session = requests.Session()
        self.http_session.headers.update(
            {
                "Content-Type": "application/json",
                "Authorization": f"Bearer {self.eternal_api_key}"
            }
        )
