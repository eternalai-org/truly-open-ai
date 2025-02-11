from typing import Any, Dict, List
from pydantic import BaseModel

from x_content.llm.base import OnchainInferResult, OpenAILLMBase
from langchain.schema import ChatGeneration, AIMessage

from x_content.wrappers.api.twitter_v2.models.objects import TwitterRequestAuthorization
from x_content.wrappers.api.twitter_v2.models.response import GenerateActionDto, Response
from x_content.wrappers.knowledge_base.base import KnowledgeBase, RelatedInformation


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


class MockKnowledgeBase(KnowledgeBase):
    """
    Class for mocking knowledge agent response in unit test
    None value is used to simulate failed llm call
    """

    mock_responses: Dict[str, List[RelatedInformation] | None] = {}

    def add_mock_response(
        self, query: str, response: List[RelatedInformation] | None, **kwargs
    ):
        hashable_entry = query
        self.mock_responses[hashable_entry] = response

    def generate(self, query: str, **kwargs: Any):
        hashable_entry = query
        if hashable_entry not in self.mock_responses:
            raise ValueError("Messages not found in mock responses")
        resp = self.mock_responses[hashable_entry]
        if resp is None:
            raise ValueError("Calling LLM failed")
        return resp

    async def agenerate(self, messages: List[dict], **kwargs: Any):
        return self.generate(messages)


def mock_sleep(delay: float):
    pass


async def a_mock_sleep(delay: float):
    pass


def mock_random_example_tweets(knowledge_base_id: str, **kwargs):
    if knowledge_base_id == "kn_1":
        return [
            "Mock example tweet 1",
            "Mock example tweet 2",
            "Mock example tweet 3",
        ]
    return []


def mock_tweet(**kwargs):
    return Response(data=GenerateActionDto(success=True))


def mock_postprocess_tweet_by_prompts(
    system_prompt: str, task_prompt: str, tweet: str
):
    return f"Postprocessed from tweet '{tweet}'"
