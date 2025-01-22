from typing import List
from x_content.constants import (
    AgentTask
)
from x_content.models import ReasoningLog
import logging
from x_content.constants import AgentTask
from x_content.models import ReasoningLog
from x_content.wrappers import telegram
from x_content import constants as const

from . import constants as const

logging.basicConfig(level=logging.INFO if not __debug__ else logging.DEBUG)
logger = logging.getLogger(__name__)

def is_local_env():
    return const.APP_ENV in ["local", "", "development", None]

def is_twin_agent(log: ReasoningLog):
    if log.meta_data is None \
        or not isinstance(log.meta_data.knowledge_base_id, str) \
        or len(log.meta_data.knowledge_base_id) == 0:
        return False

    knowledge_ids = log.meta_data.knowledge_base_id.split(",")
    return len(knowledge_ids) > 0


def send_log_alert_to_telegram(log: ReasoningLog, error):
    task_name = log.task
    if task_name == AgentTask.REACT_AGENT:
        task_name += f" with toolset {log.toolset}"
    telegram_message_html = f"""
<strong>Error occurred when executing task {task_name} for {log.meta_data.twitter_username}</strong>
<i><b>Message</b>: {error};
<b>Ref-ID</b>: {log.meta_data.ref_id};
"""

    return telegram.send_message(
        log.meta_data.twitter_username,
        telegram_message_html,
        {},
        fmt="HTML",
        room=telegram.TELEGRAM_ALERT_ROOM,
    )

def notify_trading_action(action: str, body: dict, username: str, ref_id: str, request_id: str):
    body_html_str = ""

    for key, value in body.items():
        body_html_str += f"<b>{key}</b>: {value}\n"

    msg_html = f"""<strong>{username} has taken {action} action:</strong>
<i><b>Ref-ID</b>: {ref_id};
--------------------------------
{body_html_str}
"""

    return telegram.send_message(username, msg_html, fmt="HTML")


def parse_knowledge_ids(knowledge_id: str) -> List[str]:
    knowledge_base_ids = knowledge_id.split(",")
    knowledge_base_ids = [x.strip() for x in knowledge_base_ids if x.strip() != ""]
    return knowledge_base_ids
