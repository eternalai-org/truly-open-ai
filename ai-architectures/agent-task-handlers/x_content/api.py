import json
from fastapi import APIRouter, Depends, BackgroundTasks
from fastapi.responses import JSONResponse 
from .verifications import verify_opencall_x_token, verify_third_party_authorization_key, verify_x_token
from .wrappers.api.twitter_v2.models.response import Response, SearchTweetDto
from . import constants as const
from .models import ReasoningLog, APIResponse, APIStatus, TwinTaskSubmitResponse, TwinTaskSubmitRequest, MissionChainState
from .service import MissionStateHandler, service_v2_handle_request
from .wrappers import redis_wrapper
from .legacy_services.twin import twin_service
from .tasks import utils as task_utils
import time
from functools import lru_cache
import logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

@lru_cache(maxsize=1)
def get_state_handler():
    state_handler = MissionStateHandler(
        redis_wrapper.reusable_redis_connection()
    )
    return state_handler

router = APIRouter()

@router.get("/api/twitter-news", dependencies=[
    Depends(verify_opencall_x_token), 
    Depends(verify_third_party_authorization_key(const.BACKEND_API, {}))
])
async def get_twitter_news(query: str) -> JSONResponse:
    from x_content.wrappers.api import twitter_v2
    from x_content.wrappers.magic import sync2async

    result: Response[SearchTweetDto] = await sync2async(twitter_v2.search_recent_tweets)(
        query, 
        limit_observation=10
    )
    if result.is_error():
        return APIResponse(
            status=APIStatus.ERROR,
            error=result.error,
        )
    
    tweets = result.data.tweets or []

    return APIResponse(
        status=APIStatus.SUCCESS,
        data=[e.full_text for e in tweets]
    ).model_dump()

@router.get("/api/bing-news/", dependencies=[
    Depends(verify_opencall_x_token),
    Depends(verify_third_party_authorization_key(const.BACKEND_API, {}))
])
async def get_bing_news(query: str) -> JSONResponse:
    from x_content.wrappers import bing_search
    from x_content.wrappers.magic import sync2async

    return APIResponse(
        status=APIStatus.SUCCESS,
        data=await sync2async(bing_search.search_from_bing)(query, top_k=10)
    ).model_dump()


@router.get("/api/post-v3-sample-content/", dependencies=[
    Depends(verify_opencall_x_token),
    Depends(verify_third_party_authorization_key(const.BACKEND_API, {}))
])
async def get_post_v3_sample_content(twitter_username, cutoff_hour: int=2) -> JSONResponse:
    from x_content.tasks import post_v3

    return APIResponse(
        status=APIStatus.SUCCESS,
        data=await post_v3.PostV3.get_contents(twitter_username=twitter_username, time_cutoff_hours=cutoff_hour, top_k=5)
    ).model_dump()


@router.get("/debug/redis", dependencies=[Depends(verify_x_token)])
def get_redis_by_key(key: str):
    redis_client = redis_wrapper.reusable_redis_connection()
    res = redis_client.get(key)

    if res is not None:
        return res
    else:
        return ""

@router.post(
    "/v1/twin/submit", 
    response_model=TwinTaskSubmitResponse,
    dependencies=[Depends(verify_x_token)]
)
async def twin_task_submit(request: TwinTaskSubmitRequest, background_tasks: BackgroundTasks) -> TwinTaskSubmitResponse:
    task_id = f"task_{int(time.time())}"
    
    background_tasks.add_task(twin_service.generate_twin, request.agent_id, request.twitter_ids)
    return TwinTaskSubmitResponse(
        status="success",
        task_id=task_id
    )

@router.post(
    "/async/enqueue",
    dependencies=[Depends(verify_x_token)]
)
async def enqueue_api(request: ReasoningLog, background_tasks: BackgroundTasks) -> ReasoningLog:
    logger.info(f"Received request: {json.dumps(request.model_dump())}")
    if request.state == MissionChainState.NEW:
        task_utils.notify_status(request)

    get_state_handler().commit(request)
    background_tasks.add_task(service_v2_handle_request, request)
    return request

@router.get(
    "/async/get",
    dependencies=[Depends(verify_x_token)]
)
async def get_result_api(id: str, thought_only: bool = False) -> JSONResponse:
    log = get_state_handler().get(id, none_if_error=True)

    if log is None:
        return JSONResponse(
            {},
            status_code=404,
        )

    if thought_only:
        state = log.state

        thoughts = [ 
            e.get('final_answer', e.get('thought'))
            for e in log.scratchpad if 'final_answer' in e or 'thought' in e
        ]

        return JSONResponse(
            {
                "state": state,
                "tweets": thoughts,
                "agent_contract_id": log.meta_data.agent_contract_id,
                "chain_id": log.meta_data.chain_id,
                "ref_id": log.meta_data.ref_id
            },
            status_code=200,
        )

    return JSONResponse(
        content=log.model_dump(),
        status_code=200
    )

@router.get(
    f"/async/internal-{const.API_SECRET_TOKEN}/get",
    include_in_schema=False
)
async def get_result_api_internal(id: str, thought_only: bool = False) -> JSONResponse:
    return await get_result_api(id, thought_only)