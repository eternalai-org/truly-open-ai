from fastapi import APIRouter
from .models import ResponseMessage
from .models import InsertInputSchema, QueryInputSchema, ResponseMessage, APIStatus
from fastapi import BackgroundTasks
import logging
from .handlers import process_data, run_query, get_sample, drop_kb, notify_action
from .embedding import get_default_embedding_model
from app.state import get_insertion_request_handler


logger = logging.getLogger(__name__)

router = APIRouter(
    prefix="/api",
    tags=["api"],  
    responses={404: {"description": "Not found"}},
)

@router.post("/insert", response_model = ResponseMessage)
async def insert(request: InsertInputSchema, background_tasks: BackgroundTasks) -> ResponseMessage:
    handler = get_insertion_request_handler()
    handler.insert(request)

    background_tasks.add_task(
        process_data, 
        request, 
        get_default_embedding_model()
    )
    
    await notify_action(request)

    return ResponseMessage(
        result="successfully submitted documents", 
        status=APIStatus.OK
    )

@router.post("/query", response_model=ResponseMessage)
async def query(request: QueryInputSchema) -> ResponseMessage:
    await notify_action(request)
    return ResponseMessage(result=await run_query(request))

@router.get("/sample", response_model=ResponseMessage)
async def sample(kb: str, k: int) -> ResponseMessage:
    return ResponseMessage(result=await get_sample(kb, k))

@router.delete("/delete", response_model=ResponseMessage)
async def delete(kb: str) -> ResponseMessage:
    await notify_action("<strong>Deleting</strong> all documents in knowledge base <strong>{}</strong>".format(kb))
    return ResponseMessage(result="{} documents deleted".format(await drop_kb(kb)))

@router.get("/stat", response_model=ResponseMessage, include_in_schema=False)
async def stat() -> ResponseMessage:
    return ResponseMessage(result="OK")

@router.get("/progress", response_model=ResponseMessage, include_in_schema=False)
async def stat() -> ResponseMessage:
    return ResponseMessage(result="OK")