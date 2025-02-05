import os
import logging

logging.basicConfig(level=logging.DEBUG if __debug__ else logging.INFO)
logging.getLogger("urllib3").setLevel(logging.WARNING)
logging.getLogger("httpx").setLevel(logging.WARNING)
logging.getLogger("requests").setLevel(logging.WARNING)
logger = logging.getLogger(__name__)

from typing import Optional
import threading
from x_content.service import handle_pod_shutdown, scan_db_and_resume_tasks
from x_content import constants as const
from x_content import __version__
import uvicorn
import os
from fastapi.responses import JSONResponse
from fastapi.middleware.cors import CORSMiddleware
import asyncio
import schedule
import traceback
import time
import signal
from fastapi import FastAPI


class EndpointFilter(logging.Filter):

    def filter(self, record):
        # Exclude specific endpoints
        excluded_endpoints = ["GET / HTTP"]

        if any(
            endpoint in record.getMessage() for endpoint in excluded_endpoints
        ):
            return False

        return True


# Custom logging configuration
logging_config = {
    "version": 1,
    "disable_existing_loggers": False,
    "filters": {"endpoint_filter": {"()": EndpointFilter}},
    "handlers": {
        "default": {
            "level": "INFO",
            "class": "logging.StreamHandler",
            "filters": ["endpoint_filter"],
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


def scheduler_job():
    if "SKIP_SCHEDULED_TASK" in os.environ:
        return

    logger.info("Job scheduler started")

    for job in schedule.default_scheduler.jobs:
        try:
            logger.info(f"Registered task: {job}")
        except Exception as err:
            pass

    while True:
        try:
            schedule.run_pending()
        except Exception as err:
            traceback.print_exc()
        finally:
            time.sleep(1)


previous_non_blocking_io_resume_task: Optional[threading.Thread] = None


def _setup_scheduling():
    global previous_non_blocking_io_resume_task

    def start_non_blocking_io_scan_and_resume_task(
        once=False,
    ) -> Optional[schedule.CancelJob]:
        global previous_non_blocking_io_resume_task

        if (
            previous_non_blocking_io_resume_task is None
            or not previous_non_blocking_io_resume_task.is_alive()
        ):
            non_blocking_io_resume_task = threading.Thread(
                target=asyncio.run,
                args=(scan_db_and_resume_tasks(),),
                daemon=True,
            )

            non_blocking_io_resume_task.start()

        else:
            logger.info("Non-blocking IO task is already running. Skipping...")

        return schedule.CancelJob if once else None

    # for the first time run
    schedule.every(10).minutes.do(
        lambda: start_non_blocking_io_scan_and_resume_task(once=True)
    )
    schedule.every(30).minutes.do(
        lambda: start_non_blocking_io_scan_and_resume_task(once=False)
    )


if __name__ == "__main__":
    app = FastAPI()

    signal.signal(signal.SIGTERM, handle_pod_shutdown)
    signal.signal(signal.SIGINT, handle_pod_shutdown)

    scheduler_thread = threading.Thread(target=scheduler_job, daemon=True)

    app.add_middleware(
        CORSMiddleware,
        allow_origins=["*"],
        allow_credentials=True,
        allow_methods=["*"],
        allow_headers=["*"],
    )

    @app.get("/")
    def read_root():
        status_str = "ok"

        return JSONResponse(
            {"status": status_str, "version": __version__},
            status_code=200 if status_str == "ok" else 500,
        )

    from x_content.api import router

    app.include_router(router)

    _setup_scheduling()
    scheduler_thread.start()
    uvicorn.run(
        app=app,
        host=const.SERVER_HOST,
        port=const.SERVER_PORT,
        log_level="info",
        timeout_keep_alive=300,
        log_config=logging_config,
        loop="asyncio",
    )
