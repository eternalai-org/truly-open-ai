from .base import BaseQueryOptimizer
from app.registry import register_decorator, ClassRegistry

TEMPLATE = """You are a helpful assistant that generates multiple sub-questions related to an input question. \n
The goal is to break down the input into a set of sub-problems / sub-questions that can be answers in isolation. \n
Generate multiple search queries related to: {question} \n
Output (3 queries):"""

@register_decorator(ClassRegistry.QUERY_OPTIMIZER)
class QueryParaphraser(BaseQueryOptimizer):
    def optimize(self, query):
        return super().optimize(query)