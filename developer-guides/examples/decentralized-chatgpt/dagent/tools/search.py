from typing import List
from dagent.models import Tool, ToolParam, ToolParamDtype
from dagent.registry import RegistryCategory, register_decorator
from . base_toolset import Toolset

from dagent import constant as C
import requests

import re

def remove_html_tags(text: str) -> str:
    return re.sub(re.compile('<.*?>'), '', text)

def wiki_search(query: str, lang="en", top_k=20) -> List[str]:
    headers = {
        'User-Agent': "dagent"
    }
    url = f"https://api.wikimedia.org/core/v1/wikipedia/{lang}/search/page"
    params = {
        'q': query,
        'limit': top_k
    }
    
    resp = requests.get(url, headers=headers, params=params)
    resp_json = resp.json()
    
    pages = resp_json.get("pages", [])
    
    info = []
    
    for page in pages:
        info.append({
            "title": page.get("title", ""),
            "snippet": remove_html_tags(page.get("excerpt", ""))
        })
        
    return info

@register_decorator(RegistryCategory.ToolSet)
class WikipediaSearch(Toolset):
    TOOLSET_NAME = "Wikipedia search"
    PURPOSE = "to retrieve data from Wikipedia"

    TOOLS: List[Tool] = [
        Tool(
            name="wiki_search",
            description="Search for something on Wikipedia",
            param_spec=[
                ToolParam(
                    name="query",
                    dtype=ToolParamDtype.STRING,
                    description="Query to search"
                )
            ],
            executor=lambda query: wiki_search(query)
        )
    ]