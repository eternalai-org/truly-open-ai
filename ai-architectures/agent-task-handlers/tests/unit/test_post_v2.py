from tests.unit.mock.kn_base import MockKnowledgeBase
from tests.unit.mock.llm import Message, MockLLM
from tests.unit.mock.api_twitter import mock_tweet
from tests.unit.mock.postprocess import mock_postprocess_tweet_by_prompts
from tests.unit.mock.sleep import mock_sleep
from tests.unit.mock.twin_agent import mock_random_example_tweets
import x_content.tasks
from x_content.tasks.social_agent.post_v2 import PostV2
import asyncio
import json
import pytest
from tests.unit.mock.sleep import a_mock_sleep
from x_content.constants.main import AgentTask
from x_content.llm.base import OpenAILLMBase
from x_content.models import AgentMetadata, ReactAgentReasoningMeta, ReasoningLog
from x_content.wrappers.conversation import ENHANCE_TWEET_PROMPT_TEMPLATE
from x_content.wrappers.knowledge_base.base import KnowledgeBase
from x_content.wrappers.api import twitter_v2
from x_content.wrappers import postprocess, twin_agent
import time
import x_content


@pytest.fixture
def task_fixture():
    log = ReasoningLog(
        meta_data=ReactAgentReasoningMeta(
            twitter_id="123",
            twitter_username="test_user",
            chain_id="1",
            agent_contract_id="1",
            knowledge_base_id="",
        ),
        agent_meta_data=AgentMetadata(persona="Testing system prompt"),
        system_prompt="Testing system prompt",
        prompt="Testing post_v2 prompt",
        task=AgentTask.POST_V2,
    )

    llm = MockLLM()
    kn_base = MockKnowledgeBase(base_url="mock", api_key="mock", kbs=[])

    return {"llm": llm, "kn_base": kn_base, "log": log}


class TestPostV2:

    def _setup(self, task_fixture, monkeypatch: pytest.MonkeyPatch):
        monkeypatch.setattr(time, "sleep", mock_sleep)
        monkeypatch.setattr(asyncio, "sleep", a_mock_sleep)

        monkeypatch.setattr(
            x_content.tasks.social_agent.post_v2,
            "get_random_example_tweets",
            mock_random_example_tweets,
        )
        monkeypatch.setattr(twitter_v2, "tweet", mock_tweet)
        monkeypatch.setattr(
            x_content.tasks.social_agent.post_v2,
            "postprocess_tweet_by_prompts",
            mock_postprocess_tweet_by_prompts,
        )

    @pytest.mark.asyncio
    async def test_normal_flow(self, task_fixture, monkeypatch):
        self._setup(task_fixture, monkeypatch)

        llm: MockLLM = task_fixture["llm"]
        kn_base: KnowledgeBase = task_fixture["kn_base"]

        llm.add_mock_response(
            messages=[
                Message(role="system", content="Testing system prompt"),
                Message(role="user", content="Testing post_v2 prompt"),
            ],
            response="Testing base tweet",
        )
        llm.add_mock_response(
            messages=[
                Message(role="system", content="Testing system prompt"),
                Message(
                    role="user",
                    content=ENHANCE_TWEET_PROMPT_TEMPLATE.format(
                        base_tweet="Testing base tweet"
                    ),
                ),
            ],
            response=json.dumps({"tweet": "Testing enhanced tweet"}),
        )

        post_v2_task = PostV2(
            llm=llm,
            kn_base=kn_base,
        )

        log: ReasoningLog = task_fixture["log"]
        while not log.is_done() and not log.is_error():
            log = await post_v2_task.process_task(log)
        print(log.model_dump_json(indent=2))

        task_result = log.execute_info["task_result"]

        assert len(task_result) == 1
        assert task_result[0]["base_tweet"] == "Testing base tweet"
        assert task_result[0]["enhanced_tweet"] == "Testing enhanced tweet"
        assert (
            task_result[0]["postprocessed_tweet"]
            == "Postprocessed from tweet 'Testing enhanced tweet'"
        )
