import logging
logging.basicConfig(level=logging.DEBUG if __debug__ else logging.INFO)
logging.getLogger("urllib3").setLevel(logging.WARNING)
logging.getLogger("httpx").setLevel(logging.WARNING)
logging.getLogger("requests").setLevel(logging.WARNING)
logger = logging.getLogger(__name__)
from app import __version__

import uvicorn
import os
from fastapi import Header
from fastapi import FastAPI, HTTPException
from fastapi.responses import JSONResponse
import threading
import schedule
import time
from pymilvus import connections
from app import constants as const
from typing import Annotated, Optional
from app.handlers import resume_pending_tasks
import asyncio

SECRET_TOKEN = os.environ.get("API_SECRET_TOKEN", "")

def verify_token(x_token: Annotated[str | None, Header()] = None):
    if x_token != SECRET_TOKEN:
        raise HTTPException(status_code=400, detail="X-Token header invalid")
    return x_token

class EndpointFilter(logging.Filter):
    def filter(self, record):
        # Exclude specific endpoints
        excluded_endpoints = ["GET / HTTP"]
        if any(endpoint in record.getMessage() for endpoint in excluded_endpoints):
            return False
        return True


def scheduler_job():
    if "SKIP_SCHEDULED_TASK" in os.environ:
        return

    logger.info("Scheduler started....")

    for job in schedule.default_scheduler.jobs:
        try:
            logger.info(f"Registered job: {job}")
            schedule.default_scheduler._run_job(job)
        except Exception as err:
            pass

    while True:
        try:
            schedule.run_pending()
        except Exception as err:
            pass
        finally:
            time.sleep(1)

# Custom logging configuration
logging_config = {
    "version": 1,
    "disable_existing_loggers": False,
    "filters": {
        "endpoint_filter": {
            "()": EndpointFilter
        }
    },
    "handlers": {
        "default": {
            "level": "INFO",
            "class": "logging.StreamHandler",
            "filters": ["endpoint_filter"]
        }
    },
    "loggers": {
        "uvicorn.access": {
            "handlers": ["default"],
            "level": "INFO",
            "propagate": False,
        }
    },
}

if __name__ == "__main__":
    connections.connect(uri=const.MILVUS_HOST)

    from app.api import router as app_router
    from app.handlers import prepare_milvus_collection, deduplicate_task
    prepare_milvus_collection()
    schedule.every(30).minutes.do(deduplicate_task)

    api_app = FastAPI()
    api_app.include_router(app_router)

    HOST = os.environ.get("BACKDOOR_HOST", "0.0.0.0")
    PORT = int(os.environ.get("BACKDOOR_PORT", 8000))

    @api_app.get("/", name="Health Check")
    async def read_root():
        return JSONResponse(
            {
                "status": "API is healthy",
                "version": __version__
            },
            status_code=200
        )

    _previous_thread: Optional[threading.Thread] = None    
    event_loop = asyncio.get_event_loop()

    def wrapped_resume_task():
        global _previous_thread, event_loop

        if _previous_thread is not None and _previous_thread.is_alive():
            return

        _previous_thread = threading.Thread(
            target=lambda: event_loop.run_until_complete(resume_pending_tasks(event_loop=event_loop)), 
            daemon=True
        )

        _previous_thread.start()

    schedule.every(30).minutes.do(wrapped_resume_task)
 
    scheduler_thread = threading.Thread(
        target=scheduler_job, 
        daemon=True
    )
    scheduler_thread.start()
    
    uvicorn.run(
        api_app, 
        host=HOST,
        port=PORT,
        log_level="info",
        timeout_keep_alive=300,
        log_config=logging_config,
        loop="asyncio"
    )