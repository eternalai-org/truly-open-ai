from pydantic import BaseModel
from x_content.llm.base import OnchainInferResult, OpenAILLMBase


from langchain.schema import AIMessage, ChatGeneration


from typing import Any, Dict, List


class Message(BaseModel, frozen=True):
    role: str
    content: str


class MockLLM(OpenAILLMBase):
    """
    Class for mocking LLM response in unit test
    None value is used to simulate failed llm call
    """

    mock_responses: Dict[tuple, str | None] = {}

    def add_mock_response(self, messages: List[dict], response: str | None):
        hashable_messages = tuple(
            [Message.model_validate(x) for x in messages]
        )
        self.mock_responses[hashable_messages] = response

    def generate(self, messages: List[dict], **kwargs: Any):
        hashable_messages = tuple(
            [Message.model_validate(x) for x in messages]
        )
        if hashable_messages not in self.mock_responses:
            raise ValueError("Messages not found in mock responses")
        resp_text = self.mock_responses[hashable_messages]
        if resp_text is None:
            raise ValueError("Calling LLM failed")
        return OnchainInferResult(
            generations=[ChatGeneration(message=AIMessage(content=resp_text))],
            tx_hash="0x123",
        )

    async def agenerate(self, messages: List[dict], **kwargs: Any):
        return self.generate(messages)
