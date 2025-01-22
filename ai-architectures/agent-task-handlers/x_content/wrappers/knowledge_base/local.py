from x_content.models import AgentKnowledgeBase
from .base import KnowledgeBase, RelatedInformation
from typing import List
import httpx
import logging

logger = logging.getLogger(__name__)

class KBStore(KnowledgeBase):
    async def aretrieve(self, query: str, top_k: int = None, threshold: float = None, *args, **kwargs) -> List[RelatedInformation]:
        if len(self.kbs) == 0:
            return []

        url = f"{self.base_url}api/query"

        headers = {
            "Authorization": self.api_key,
        }

        body = {
            "query": query,
            "top_k": top_k or self.default_top_k,
            "kb": [x.kb_id for x in self.kbs],
        }
        
        async with httpx.AsyncClient() as client:
            response = await client.post(
                url,
                headers=headers,
                json=body,
                timeout=httpx.Timeout(60.0),
            )

        if response.status_code != 200:
            logger.error(f"Failed to send request to '{url}'; code: {response.status_code}; raw: {response.text}.")
            return []
            
        response_json = response.json()
        
        return [
            RelatedInformation(
                content=resp["content"], 
                score=resp["score"], 
                reference=resp.get("reference")
            ) 
            for resp in response_json.get("result", [])
            if resp["score"] >= (threshold or self.similarity_threshold)
        ]
        