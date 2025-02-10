from typing import Annotated, TypedDict
from langgraph.graph import StateGraph, START, END
from langgraph.graph.state import CompiledStateGraph
from langgraph.graph.message import add_messages
from langchain.schema import ChatMessage
from x_content.llm.base import OpenAILLMBase


class State(TypedDict):
    # Messages have the type "list". The `add_messages` function
    # in the annotation defines how this state key should be updated
    # (in this case, it appends messages to the list, rather than overwriting them)
    messages: list[ChatMessage]
    tx_hash: str = ""


def get_chat_langgraph(llm: OpenAILLMBase) -> CompiledStateGraph:
    async def chatbot(state: State):
        messages = state["messages"]
        infer_result = await llm.agenerate(messages)
        content = infer_result.generations[0].message.content

        messages.append(ChatMessage(role="assistant", content=content))
        return {
            "messages": messages,
            "tx_hash": infer_result.tx_hash,
        }

    graph_builder = StateGraph(State)
    graph_builder.add_node("chatbot", chatbot)
    graph_builder.add_edge(START, "chatbot")

    graph = graph_builder.compile()

    return graph
