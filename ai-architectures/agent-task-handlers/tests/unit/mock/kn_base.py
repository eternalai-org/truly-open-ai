from x_content.wrappers.knowledge_base.base import KnowledgeBase, RelatedInformation


from typing import Any, Dict, List


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
