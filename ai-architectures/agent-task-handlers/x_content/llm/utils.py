import httpx
from typing import Dict, Any
from pydantic import BaseModel
from typing import Optional

class ServerInferenceResult(BaseModel):
    skipped: bool = False
    response: Optional[Dict[str, Any]] = None

# TODO: remove this
async def check_and_get_infer_result(url: str, headers: Dict[str, str]) -> ServerInferenceResult:
    async with httpx.AsyncClient() as client:
        response = await client.get(url, headers=headers, timeout=httpx.Timeout(60.0))

    if response.status_code != 200:
        return ServerInferenceResult(skipped=True)

    resp = response.json()
    return ServerInferenceResult(response=resp)

