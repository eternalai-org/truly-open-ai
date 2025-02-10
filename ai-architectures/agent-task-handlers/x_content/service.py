import logging
from x_content.cache import MissionStateHandler
from x_content.cache.chat_request_state_handler import ChatRequestStateHandler
from x_content.constants import MissionChainState
from x_content.tasks.base import ChatTaskBase
from x_content.tasks.utils import create_llm

from .constants import REACT_MODELS_BLACKLIST, REPLY_MODELS_BLACKLIST
from x_content.models import ChatRequest, ReasoningLog
from x_content.tasks import MultiStepTaskBase
from typing import Optional
from x_content.wrappers.redis_wrapper import reusable_redis_connection
import sys

import traceback

logging.basicConfig(level=logging.INFO if not __debug__ else logging.DEBUG)
logger = logging.getLogger(__name__)

from .constants import AgentTask
from .tasks.utils import create_kn_base, magic_toolset_from_reasoning_log, notify_status_chat_request, notify_status_reasoning_log, send_alert

from . import tasks
from .constants import ToolSet, ModelName
import datetime
import asyncio


# TODO: bad designed here, refactor it
def task_cls_resolver(log: ReasoningLog) -> Optional[MultiStepTaskBase]:
    if log.task == AgentTask.SHADOW_REPLY:
        return tasks._legacy.shadow_reply.ShadowReplyTask

    if log.task == AgentTask.REPLY:
        if log.model in REPLY_MODELS_BLACKLIST:
            return tasks.others.FallbackTask

        return tasks.social_agent.social_reply.SocialReplyTask

    if log.task == AgentTask.POST_V2:
        return tasks.social_agent.post_v2.PostV2

    if log.task == AgentTask.CREATE_GAME:
        return tasks.game_agent.create_gamev2.GameReplyTask

    if log.task == AgentTask.JUDGE_GAME:
        return tasks.game_agent.judge_gamev2.JudgeGameTask

    if log.task == AgentTask.QUOTE_TWEET:
        return tasks._legacy.quote_tweet.QuoteTweetTask

    if log.task == AgentTask.POST_SEARCH:
        return tasks._legacy.post_search.PostSearchTask

    if log.task == AgentTask.TRADING:
        return tasks.social_agent.react_agent_for_trading.TradingTask

    if log.task == AgentTask.DEFAULT:
        return tasks.social_agent.react_agent.ReactAgent

    if log.task == AgentTask.POST_V3:
        return tasks.social_agent.post_v3.PostV3

    if log.task == AgentTask.REACT_AGENT:
        if log.model in REACT_MODELS_BLACKLIST:
            return tasks.others.FallbackTask

        if log.toolset == ToolSet.TRADING:
            return tasks.social_agent.react_agent_for_trading.TradingTask

        if log.model in [
            ModelName.DEEPSEEK_R1,
            ModelName.DEEPSEEK_V3,
            ModelName.DEEPSEEK_R1_DISTILL_LLAMA_70B,
        ]:
            return (
                tasks.social_agent.react_agent_use_deepseek_r1.ReactAgentUsingDeepSeekR1
            )

        return tasks.social_agent.react_agent.ReactAgent

    return tasks.others.FallbackTask


def chat_request_cls_resolver(chat: ChatRequest) -> Optional[ChatTaskBase]:
    return tasks.social_agent.chat_v2.ChatV2


_running_tasks = set([])
_task_handled_key = "task_handled:{}"


async def service_v2_handle_request(log: ReasoningLog) -> ReasoningLog:
    global _running_tasks

    do_job = log.id not in _running_tasks
    # and atomic_check_and_set_flag(
    #     reusable_redis_connection(),
    #     _task_handled_key.format(log.id),
    #     "1", 6 * 3600
    # )

    if not do_job:
        logger.info(f"Task {log.id} is already handled (by someone else)")
        return log

    logger.info(f"Handling task {log.id}")

    try:
        _running_tasks.add(log.id)
        task_handler_cls = task_cls_resolver(log)

        if task_handler_cls is None:
            raise Exception(
                f"Bad request: Task handler not found for log {log.id}"
            )

        llm = create_llm(log)
        toolset = magic_toolset_from_reasoning_log(log, llm)
        kn_base = create_kn_base(log)

        task_handler: MultiStepTaskBase = task_handler_cls(
            llm=llm,
            toolcall=toolset,
            kn_base=kn_base,
        )

        log.llm_info = llm.get_info()
        log = await task_handler.run(log)

    except Exception as err:
        traceback_str = traceback.format_exc()
        send_alert(log, traceback_str)
        log = await tasks.utils.a_move_state(
            log,
            MissionChainState.ERROR,
            f"An error occurred: {err} (unhandled)",
        )
        MissionStateHandler(reusable_redis_connection()).commit(log)

    finally:
        notify_status_reasoning_log(log)
        _running_tasks.remove(log.id)
        # redis_cli = reusable_redis_connection()
        # redis_cli.expire(_task_handled_key.format(log.id), 3)

    logger.info(f"Completed handling task {log.id}")
    return log


async def handle_chat_request(request: ChatRequest) -> ChatRequest:
    global _running_tasks

    do_job = request.id not in _running_tasks

    if not do_job:
        logger.info(
            f"Chat request {request.id} is already handled (by someone else)"
        )
        return request

    logger.info(f"Handling chat request {request.id}")

    try:
        _running_tasks.add(request.id)

        llm = create_llm(request)
        kn_base = create_kn_base(request)

        task_handler_cls = chat_request_cls_resolver(request)

        task_handler: ChatTaskBase = task_handler_cls(
            llm=llm,
            kn_base=kn_base,
        )

        request = await task_handler.run(request)

    except Exception as err:
        traceback_str = traceback.format_exc()
        send_alert(request, traceback_str)
        request = await tasks.utils.a_move_state(
            request,
            MissionChainState.ERROR,
            f"An error occurred: {err} (unhandled)",
        )
        ChatRequestStateHandler(reusable_redis_connection()).commit(request)

    finally:
        notify_status_chat_request(request)
        _running_tasks.remove(request.id)

    logger.info(f"Completed handling task {request.id}")
    return request


async def scan_db_and_resume_tasks():
    logger.info("Scanning DB for resumable tasks")

    handler = MissionStateHandler(reusable_redis_connection())
    undone_task = handler.get_undone()

    if len(undone_task) == 0:
        logger.info("No undone task found")
        return

    futures = []
    current_time = datetime.datetime.now()

    for log in undone_task:
        if log.id in _running_tasks:
            continue

        created_dtime = datetime.datetime.strptime(
            log.created_at, "%Y-%m-%dT%H:%M:%S.%fZ"
        )

        if (current_time - created_dtime).total_seconds() > 60 * 60 * 6:
            logger.info(f"Task {log.id} is too old, skipping")
            continue

        task_cls = task_cls_resolver(log)

        if not task_cls.resumable:
            logger.info(f"Task {log.id} is not resumable, skipping")
            continue

        logger.info(f"Resuming task {log.id}")
        futures.append(asyncio.ensure_future(service_v2_handle_request(log)))

    logger.info(f"Resuming {len(futures)} tasks")

    if len(futures) > 0:
        await asyncio.gather(*futures)

    logger.info("Scanning DB for resumable tasks completed")


async def scan_db_and_resume_chat_requests():
    logger.info("Scanning DB for resumable chat requests")

    handler = ChatRequestStateHandler(reusable_redis_connection())
    undone_request = handler.get_undone()

    if len(undone_request) == 0:
        logger.info("No undone chat request found")
        return

    futures = []
    current_time = datetime.datetime.now()

    for request in undone_request:
        if request.id in _running_tasks:
            continue

        created_dtime = datetime.datetime.strptime(
            request.created_at, "%Y-%m-%dT%H:%M:%S.%fZ"
        )

        if (current_time - created_dtime).total_seconds() > 60 * 60 * 6:
            logger.info(f"Chat request {request.id} is too old, skipping")
            continue

        task_cls = task_cls_resolver(request)

        if not task_cls.resumable:
            logger.info(
                f"Chat request {request.id} is not resumable, skipping"
            )
            continue

        logger.info(f"Resuming chat request {request.id}")
        futures.append(asyncio.ensure_future(handle_chat_request(request)))

    logger.info(f"Resuming {len(futures)} chat requests")

    if len(futures) > 0:
        await asyncio.gather(*futures)

    logger.info("Scanning DB for resumable chat requests completed")


def handle_pod_shutdown(signum, frame):
    global _running_tasks, _task_handled_key, logger

    logger.info("Pod is being shut down")
    redis_cli = reusable_redis_connection()

    for task_id in _running_tasks:
        logger.info(f"Removing task {task_id} from running tasks")
        redis_cli.delete(_task_handled_key.format(task_id))

    _running_tasks = set([])
    sys.exit(0)
