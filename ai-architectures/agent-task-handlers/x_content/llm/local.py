import json
from x_content.llm.base import OnchainInferResult, OpenAILLMBase
from x_content.llm.eternal_ai import logger

import httpx
import requests
from langchain.schema import AIMessage, BaseMessage, ChatGeneration, ChatResult
from langchain_community.adapters.openai import convert_message_to_dict

from typing import Any, List, Optional


class SyncBasedEternalAI(OpenAILLMBase):

    async def start_async_request(
        self, messages: List[BaseMessage], **kwargs
    ) -> dict:
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
            "stream": True,
        }

        json_data.update(kwargs)

        url = f"{self.openai_api_base}/chat/completions"

        token_usage = 0
        final_response = ""

        try:
            async with httpx.AsyncClient() as client:
                async with client.stream(
                    "POST",
                    url,
                    headers=headers,
                    json=json_data,
                    timeout=httpx.Timeout(120.0),
                ) as response:
                    if response.status_code != 200:
                        logger.info(
                            f"Failed to send request to '{url}'; code: {response.status_code}"
                        )
                        raise ValueError(
                            f"Failed to send request to '{url}'; code: {response.status_code}"
                        )

                    # Handle streaming response
                    async for chunk in response.aiter_lines():
                        data = chunk.split("data: ")

                        if len(data) <= 1:
                            continue

                        if data[1] == "[DONE]":
                            break

                        json_data = json.loads(data[1])

                        final_response += json_data["choices"][0]["delta"].get(
                            "content", ""
                        )

                        token_usage += 1

        except Exception as err:
            logger.info(f"Failed to send request to '{url}'; error: {err}")
            raise ValueError(
                f"Failed to send request to '{url}'; error: {err}"
            )

        return {
            "message": AIMessage(content=final_response),
            "token_usage": token_usage,
        }

    async def agenerate(
        self,
        messages: List[BaseMessage],
        stop: Optional[List[str]] = None,
        **kwargs: Any,
    ) -> OnchainInferResult:
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
            "stream": True,
        }

        final_response = ""
        token_usage = 0
        # Send the POST request with streaming enabled
        with requests.post(
            f"{self.openai_api_base}/chat/completions",
            json=json_data,
            headers=headers,
            stream=True,
        ) as response:
            if response.status_code != 200:
                logger.info(
                    f"Failed to send request to '{self.openai_api_base}/chat/completions'; code: {response.status_code}"
                )
                raise Exception(
                    f"Failed to send request to '{self.openai_api_base}/chat/completions'; code: {response.status_code}"
                )
            # Handle streaming response
            for chunk in response.iter_lines():
                if chunk:
                    # Decode each chunk and print the content
                    data = chunk.decode("utf-8").split("data: ")
                    if data[1] == "[DONE]":
                        break
                    json_data = json.loads(data[1])
                    final_response += json_data["choices"][0]["delta"].get(
                        "content", ""
                    )
                    token_usage += 1

        return {
            "message": AIMessage(content=final_response),
            "token_usage": token_usage,
        }

    def generate(
        self,
        messages: List[BaseMessage],
        stop: Optional[List[str]] = None,
        **kwargs: Any,
    ) -> OnchainInferResult:
        """Generate a chat response synchronously and return it as a OnchainInferResult."""
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
