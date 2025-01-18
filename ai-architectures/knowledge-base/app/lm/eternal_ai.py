from typing import Any, List, Optional, Dict

import httpx
from langchain.schema import BaseMessage, ChatGeneration, ChatResult, AIMessage
from langchain_community.adapters.openai import convert_message_to_dict
from asyncio import sleep as async_sleep
import time
import json
import logging
from app.lm.time_estimation import ModelInferTimeEstimation
from pydantic import BaseModel

logger = logging.getLogger(__name__)

from .base import OpenAILLMBase

from typing import Dict
from functools import lru_cache

@lru_cache(maxsize=1)
def get_time_estimation():
    return ModelInferTimeEstimation()

class OnchainInferResult(ChatResult):
    tx_hash: str = None
    receipt: str = None

class ServerInferenceResult(BaseModel):
    skipped: bool = False
    response: Optional[Dict[str, Any]] = None

# hmm, this is perfectly the same as openai standard 
class SyncBasedEternalAI(OpenAILLMBase):
    async def start_async_request(self, messages: List[BaseMessage]) -> dict:
        """
        Initiate a request to the OpenAI ChatCompletion endpoint.
        Since OpenAI returns results immediately, we'll store the response locally
        and simulate asynchronous processing via a cached lookup.
        """
        openai_messages = [
            convert_message_to_dict(m) if not isinstance(m, dict) else m
            for m in messages
        ]

        headers = {
            "Authorization": f"Bearer {self.openai_api_key.get_secret_value()}",
            "Content-Type": "application/json",
        }

        json_data = {
            "messages": openai_messages,
            **self.get_info()
        }

        async with httpx.AsyncClient() as client:
            response = await client.post(
                f"{self.openai_api_base}/chat/completions",
                headers=headers,
                json=json_data,
                timeout=httpx.Timeout(60.0),
            )

        if response.status_code != 200:
            logger.info(f"Failed to send request to '{self.openai_api_base}/chat/completions'; code: {response.status_code}; raw: {response.text}.")

        response_json = response.json()

        choices = response_json.get("choices", [])
        if not choices:
            raise ValueError("No choices found in the OpenAI response.")

        if choices[0].get("message") is None or choices[0].get("message", {}).get("content") is None:
            raise ValueError("Bad response from LLM-server: {}".format(response_json))

        content = choices[0].get("message", {}).get("content", "")
        token_usage = response_json.get("usage", {})

        return {
            "message": AIMessage(content=content),
            "token_usage": token_usage,
        }

    async def agenerate(
        self,
        messages: List[BaseMessage],
        stop: Optional[List[str]] = None,
        **kwargs: Any,
    ) -> ChatResult:
        """Generate a chat response asynchronously and return it as a ChatResult."""
        # Start the asynchronous request
        result: dict = await self.start_async_request(messages)

        if "message" not in result:
            raise ValueError(f"Unexpected response from LLM-Server: {result}")

        # Return a mock OnchainInferResult 
        return OnchainInferResult(
            generations=[ChatGeneration(message=result["message"])],
            llm_output={"token_usage": result.get("token_usage", {})},
            tx_hash="",
        )


# TODO: add retry logic
class ASyncBasedEternalAI(OpenAILLMBase):
    chain_id: str
    agent_contract_id: str
    metadata: dict
    timeout_seconds: int = 60 * 60 * 3 # expire in 3 hours
    failed_count_limt: int = 5

    def __init__(self, *args, **kwargs):  
        super(ASyncBasedEternalAI, self).__init__(*args, **kwargs)
        
    async def check_and_get_infer_result(self, receipt) -> ServerInferenceResult:
        headers = {
            "Authorization": self.openai_api_key.get_secret_value(),
            "Content-Type": "application/json",
        }

        url = f"{self.openai_api_base}/agent/get-batch-item-output/{receipt}"

        async with httpx.AsyncClient() as client:
            response = await client.get(url, headers=headers, timeout=httpx.Timeout(60.0))

        if response.status_code != 200:
            return ServerInferenceResult(skipped=True)

        resp = response.json()
        return ServerInferenceResult(response=resp)

    async def submit_async_request(self, 
        messages: List[BaseMessage],
    ) -> str:

        openai_messages = [
            convert_message_to_dict(m) if not isinstance(m, dict) else m
            for m in messages
        ]

        headers = {
            "Authorization": self.openai_api_key.get_secret_value(),
            "Content-Type": "application/json",
        }

        json_data = {
            "messages": openai_messages,
            **self.get_info()
        }
        
        url = f"{self.openai_api_base}/agent/async-batch-prompt"
        
        async with httpx.AsyncClient() as client:
            response = await client.post(
                url,
                headers=headers,
                json=json_data,
                timeout=httpx.Timeout(60.0),
            )

        response_json = response.json()

        if response_json["status"] < 0 or response_json["data"].get("id") is None:
            raise Exception("Inference request submission failed")

        receipt = response_json["data"]["id"]
        return receipt

    async def wait(self, receipt: str, eta_seconds: float = 60):
        started_at = time.time()

        while True:
            # step 3: sleep for a while
            await async_sleep(eta_seconds)
            eta_seconds = max(eta_seconds * 0.1, 5)

            # step 0: check if the task is timed out
            current_time = time.time()

            if current_time - started_at > self.timeout_seconds:
                raise Exception("Inference request timed out")

            check_result: ServerInferenceResult = await self.check_and_get_infer_result(receipt)

            if not check_result.skipped:
                resp = check_result.response

                # step 2: check and parse the response
                if resp["status"] < 0 or resp["data"]["status"] == "error":
                    raise Exception(f"Inference request failed; Receipt: {receipt}; Raw Output: {resp}")

                if resp["data"]["status"] == "queue-handled":
                    tx_hash = resp["data"]["inscribe_tx_hash"]

                    try:
                        prompt_output = json.loads(resp["data"]["prompt_output"])
                    except json.JSONDecodeError:
                        raise Exception(f"Failed to decode prompt output; Receipt: {receipt}; Tx Hash: {tx_hash}; Raw Output: {resp}") 
    
                    choices: List[Dict[str, Any]] = prompt_output.get("choices", [])
                    if not choices:
                        raise ValueError(f"No choices found in the OpenAI response; Receipt: {receipt}; Tx Hash: {tx_hash}; Raw Output: {prompt_output}")

                    if choices[0].get("message") is None or choices[0].get("message", {}).get("content") is None:
                        raise ValueError(f"Bad response from LLM-server. Receipt: {receipt}; Tx Hash: {tx_hash}; Raw Output: {prompt_output}")

                    content = choices[0].get("message", {}).get("content", "")
                    token_usage = prompt_output.get("usage", {})

                    return {
                        "message": AIMessage(content=content),
                        "token_usage": token_usage,
                        "tx_hash": tx_hash,
                    }

    async def agenerate(
        self,
        messages: List[BaseMessage],
        stop: Optional[List[str]] = None,
        **kwargs: Any,
    ) -> OnchainInferResult:

        receipt: str = await self.submit_async_request(messages)
        submit_time = time.time()
        estimation = get_time_estimation()

        logger.info(f"Submitted async request; Receipt: {receipt}")
        try:
            result: dict = await self.wait(
                receipt, 
                estimation.estimate(self.model_name)
            )

        finally:
            estimation.update(
                self.model_name, 
                time.time() - submit_time
            )

        if "message" not in result:
            raise ValueError(f"Unexpected response from OpenAI: {result}")

        # Return a ChatResult
        return OnchainInferResult(
            generations=[ChatGeneration(message=result["message"])],
            llm_output={"token_usage": result.get("token_usage", {})},
            tx_hash=result.get("tx_hash", ""),
            receipt=receipt
        )

    def get_info(self):
        return {
            "chain_id": self.chain_id,
            "agent_contract_id": self.agent_contract_id,
            "metadata": self.metadata,
            "timeout_seconds": self.timeout_seconds,
            "failed_count_limt": self.failed_count_limt,
            **super().get_info()
        }