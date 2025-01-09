import logging

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

from dotenv import load_dotenv
if not load_dotenv():
    logger.warning("No .env file found")

import dagent
import sys
import schedule
import time
from argparse import ArgumentParser

import dagent.utils
import os
import json
from dagent.registry import get_registered, RegistryCategory

from fastapi import FastAPI
import uvicorn

import threading
import http_endpoints

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

def parse_opt(
    daemon_config_file = os.path.join(
        dagent.utils.get_script_dir(__file__), 
        "configs/daemon.json"
    )
):
    
    daemon_config = {}
    
    if os.path.exists(daemon_config_file):
        with open(daemon_config_file, "r") as f:
            daemon_config = json.load(f)

    parser = ArgumentParser(description="Decentralized-agent Daemon")
    parser.add_argument("-c", "--agent-config-file", type=str, 
                        default=os.path.join(dagent.utils.get_script_dir(__file__), "configs/eternal.json"),
                        help="Path to the agent config file. Used to init the non-interactive auto-agents.")

    parser.add_argument(
        "--serve-interactive-agents", 
        action="store_true", 
        default=daemon_config.get("serve_interactive_agents", False),
        help="Serve interactive agent on the same port"
    )

    parser.add_argument(
        "-H", "--http-host", "--host", 
        type=str, 
        default=daemon_config.get("http_host", "localhost")
    )

    parser.add_argument(
        "-P", "--http-port", "--port", 
        type=int, 
        default=daemon_config.get("http_port", 8080)
    )

    return parser.parse_args()

def http_service(provider: dagent.service.AutoServiceProvider):
    opts = parse_opt()

    fast_api = FastAPI()
    fast_api.include_router(http_endpoints.api_v1_router, prefix="/api/v1", tags=["api"])
    fast_api.include_router(http_endpoints.router, prefix="/api", tags=["api"])

    uvicorn.run(
        fast_api, 
        host=opts.http_host, 
        port=opts.http_port,
        log_level="info"
    )

def main():
    for item in [RegistryCategory.LLM, RegistryCategory.ToolSet]:
        logger.info(f"Registered {item}: {get_registered(item)}")
    
    service = dagent.service.AutoServiceProvider()
    args = parse_opt()
    assert os.path.exists(args.agent_config_file), f"Config file {args.agent_config_file} not found"
    
    with open(args.agent_config_file, "rb") as fp:
        cfg = json.loads(fp.read())

    service.schedule(cfg)    
    service.start()

    if args.serve_interactive_agents:
        http_service_thread = threading.Thread(
            target=http_service, 
            args=(service,), 
            daemon=True
        )

        http_service_thread.start()

    while True:
        try:
            schedule.run_pending()
        except Exception as e:
            logger.error(f"Scheduling error: {e}")
        finally:
            time.sleep(1)

if __name__ == '__main__':
    sys.exit(main())