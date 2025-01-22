from .utils import a_move_state
from .base import MultiStepTaskBase
from json_repair import repair_json
from x_content.utils import parse_knowledge_ids
from x_content.wrappers.knowledge_base.local import KBStore
from x_content.models import AgentKnowledgeBase, ReasoningLog
from x_content.models import ReasoningLog, MissionChainState
from x_content.llm.eternal_ai import OnchainInferResult
from x_content import constants as const


class Chat(MultiStepTaskBase):
    resumable = True
    GENERATE_QUERY_TEMPLATE = """Generate a concise and effective search query to retrieve relevant information from the database. Ensure the query is clear, simple, and optimized for accurate results based on the input question:
{question}
Respond in stringified JSON format with the following structure:
{{
  "query": "<generated_query>"
}}
"""
    GENERATE_ANSWER_TEMPLATE = """{searched_results}
    

Using the information above to generate a concise and effective response to the user's question. Ensure the response is clear, simple, and optimized for accurate results based on the input question:
{question}    
"""
    async def process_task(self, log: ReasoningLog) -> ReasoningLog:

        if log.state == MissionChainState.NEW:
            messages = [
                {
                    'role': 'user',
                    'content': self.GENERATE_QUERY_TEMPLATE.format(question=log.prompt)
                }
            ]
            log.execute_info = {
                'question': log.prompt,
                "generate_query_messages": messages
            }
            return await a_move_state(log, MissionChainState.RUNNING, "Task started")

        if log.state == MissionChainState.RUNNING:
            if len(log.agent_meta_data.kb_agents) > 0:
                kb_ids = parse_knowledge_ids(log.meta_data.knowledge_base_id)
                messages = log.execute_info["generate_query_messages"]
                result: OnchainInferResult = await self.llm.agenerate(messages, temperature=0.7)
                content = result.generations[0].message.content
                parsed_content = repair_json(content, return_objects=True)

                if "query" not in parsed_content:
                    return await a_move_state(log, MissionChainState.ERROR, "No query generated")

                query = parsed_content["query"]

                # TODO: update this
                kb_store = KBStore(
                    default_top_k=5,
                    similarity_threshold=0.5,
                    base_url=const.RAG_API,
                    api_key=const.RAG_SECRET_TOKEN,
                    kbs=[AgentKnowledgeBase(chain_id="", kb_id=id) for id in kb_ids],
                )

                search_results = await kb_store.aretrieve(query)
                search_results = [search_result.content for search_result in search_results]
                messages = [
                    {
                        'role': 'user',
                        'content': self.GENERATE_ANSWER_TEMPLATE.format(
                            searched_results='\n'.join(search_results),
                            question=log.execute_info["question"]
                        )
                    }
                ]

            else:
                messages = [
                    {
                        'role': 'user',
                        'content': log.execute_info["question"]
                    }
                ]

            result: OnchainInferResult = await self.llm.agenerate(messages, temperature=0.7)
            log.execute_info['task_result'] = result.generations[0].message.content
            return await a_move_state(log, MissionChainState.DONE, "Task completed")
    
        return log
