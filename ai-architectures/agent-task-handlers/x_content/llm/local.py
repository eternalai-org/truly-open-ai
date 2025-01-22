from x_content.llm.base import OpenAILLMBase
from x_content.llm.eternal_ai import OnchainInferResult, logger

import httpx
import requests
from langchain.schema import AIMessage, BaseMessage, ChatGeneration, ChatResult
from langchain_community.adapters.openai import convert_message_to_dict

from typing import Any, List, Optional


class SyncBasedEternalAI(OpenAILLMBase):
    async def start_async_request(self, messages: List[BaseMessage], **kwargs) -> dict:
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
            "model": self.model_name,
            "messages": openai_messages,
            "temperature": self.temperature,
            "max_tokens": self.max_tokens,
            "top_p": 1.0,
            "presence_penalty": 0.1,
            "n": self.n,
            "logit_bias": None,
            "frequency_penalty": self.frequency_penalty,
            "seed": self.seed,
        }

        json_data.update(kwargs)
        # print("json_data:", json_data)

        url = f"{self.openai_api_base}/chat/completions"

        async with httpx.AsyncClient() as client:
            response = await client.post(
                url,
                headers=headers,
                json=json_data,
                timeout=httpx.Timeout(60.0),
            )

        if response.status_code != 200:
            logger.info(f"Failed to send request to '{url}'; code: {response.status_code}; raw: {response.text}.")
            raise ValueError(f"Failed to send request to '{url}'; code: {response.status_code}; raw: {response.text}.")

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
        result: dict = await self.start_async_request(messages, **kwargs)

        if "message" not in result:
            raise ValueError(f"Unexpected response from LLM-Server: {result}")

        # Return a mock OnchainInferResult 
        return OnchainInferResult(
            generations=[ChatGeneration(message=result["message"])],
            llm_output={"token_usage": result.get("token_usage", {})},
            tx_hash="",
        )

    def start_sync_request(self, messages: List[BaseMessage]) -> dict:
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
            "model": self.model_name,
            "messages": openai_messages,
            "temperature": self.temperature,
            "max_tokens": self.max_tokens,
            "top_p": 1.0,
            "presence_penalty": 0.1,
            "n": self.n,
            "logit_bias": None,
            "frequency_penalty": self.frequency_penalty,
            "seed": self.seed,
        }

        response = requests.post(
            f"{self.openai_api_base}/chat/completions",
            headers=headers,
            json=json_data,
        )

        if response.status_code != 200:
            raise Exception(f"Failed to send request to '{self.openai_api_base}/chat/completions'; code: {response.status_code}; raw: {response.text}.")

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

    def generate(
        self,
        messages: List[BaseMessage],
        stop: Optional[List[str]] = None,
        **kwargs: Any,
    ) -> ChatResult:
        """Generate a chat response asynchronously and return it as a ChatResult."""
        # Start the asynchronous request
        result: dict = self.start_sync_request(messages)

        if "message" not in result:
            raise ValueError(f"Unexpected response from LLM-Server: {result}")

        # Return a mock OnchainInferResult 
        return OnchainInferResult(
            generations=[ChatGeneration(message=result["message"])],
            llm_output={"token_usage": result.get("token_usage", {})},
            tx_hash="",
        )

    def get_info(self):
        return {
            "model": self.model_name,
            "temperature": self.temperature,
            "max_tokens": self.max_tokens,
            "top_p": 1.0,
            "presence_penalty": 0.1,
            "n": self.n,
            "logit_bias": None,
            "frequency_penalty": self.frequency_penalty,
            "seed": self.seed,
        }