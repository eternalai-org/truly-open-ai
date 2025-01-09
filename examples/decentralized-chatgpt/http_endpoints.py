from fastapi import APIRouter
from fastapi.responses import JSONResponse, PlainTextResponse
from dagent.service import AutoServiceProvider
from dagent.models import ChatSession

router = APIRouter()
api_v1_router = APIRouter()

__all__ = [
    "api_v1_router"
]

@router.get("/health")
def health_check():
    return PlainTextResponse("", status_code=200)

@api_v1_router.post("/init-chat")
def init_chat():
    return JSONResponse(content={
        "error": "Not implemented" 
    }, status_code=500)

@api_v1_router.post("/chat/{session_id}")
def chat(session_id: str, message: str):
    return JSONResponse(content={
        "error": "Not implemented" 
    }, status_code=500)

@api_v1_router.get("/chat/{session_id}/history")
def history(session_id: str):
    return JSONResponse(content={
        "error": "Not implemented" 
    }, status_code=500)
    
@api_v1_router.get("/deinit-chat/{session_id}")
def deinit_chat(session_id: str):
    return JSONResponse(content={
        "error": "Not implemented" 
    }, status_code=500)