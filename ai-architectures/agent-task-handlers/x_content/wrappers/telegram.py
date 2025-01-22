import requests
import json
import logging
from . import constants as const
import schedule

TELEGRAM_ROOM = const.TELEGRAM_ROOM
TELEGRAM_ALERT_ROOM = const.TELEGRAM_ALERT_ROOM
TELEGRAM_TASK_IO_NOTIFICATION_ROOM = const.TELEGRAM_TASK_IO_NOTIFICATION_ROOM

logger = logging.getLogger(__name__)

from pydantic import BaseModel
from .redis_wrapper import reusable_redis_connection, distributed_scheduling_job


class TelegramMessage(BaseModel):
    text: str
    sender: str = "junk_notifications"
    parse_mode: str = "HTML"
    disable_notification: bool = True
    link_preview_options: dict = {}
    room: str

    can_batch: bool = False


def escape_str(s: str):
    rules = [("_", "\\_"), ("*", "\\*"), ("[", "\\["), ("`", "\\`")]

    special = "[#*#]"

    for a, b in rules:
        s = s.replace(b, special)
        s = s.replace(a, b)
        s = s.replace(special, b)

    return s


def get_url(api_key: str = const.TELEGRAM_API_KEY, room: str = const.TELEGRAM_ROOM):
    return f"https://api.telegram.org/bot{api_key}/sendMessage?chat_id={room}"


def send_message(
    twitter_username: str,
    message_to_send: str,
    preview_opt={},
    fmt="HTML",
    room=const.TELEGRAM_ROOM,
    schedule=False,
):
    twitter_username = twitter_username or "junk_notifications"

    if (
        twitter_username == "junk_nofitications"
        and const.DISABLE_JUNK_NOTIFICATIONS
    ):
        return False

    if twitter_username in const.DISABLED_TELEGRAM_USERS:
        return False

    if schedule:
        _enqueue(
            TelegramMessage(
                text=message_to_send,
                sender=const.TELEGRAM_BOTNAME,
                parse_mode=fmt,
                disable_notification=True,
                link_preview_options=preview_opt,
                room=room,
                can_batch=True,
            )
        )

        return True

    url = get_url(room=room)

    logger.info(f"Sending a message of length {len(message_to_send)} to room {room}")
    payload = {
        "text": message_to_send,
        "parse_mode": fmt,
        "disable_notification": True,
        "link_preview_options": json.dumps(preview_opt),
    }

    # logger.info("Sending message to Telegram: {}".format(json.dumps(payload, indent=2)))
    resp = requests.post(url, json=payload)

    if resp.status_code == 200:
        return True

    logger.error(f"Failed to send message to Telegram: {resp.text}")

    return False

from redis import Redis


def _enqueue(msg: TelegramMessage):
    redis_connection: Redis = reusable_redis_connection()
    redis_connection.rpush(
        const.TELEGRAM_MESSAGE_LIST_REDIS_KEY, json.dumps(msg.model_dump())
    )


from typing import List


def group_message(
    msgs: List[TelegramMessage],
    separator: str,
    limit_chars: int = const.TELEGRAM_MESSAGE_LENGTH_LIMIT,
) -> List[List[TelegramMessage]]:
    """
    Groups messages into a single message if the total length of the messages is less than the limit.
    """
    total_length = 0
    grouped_msgs = []
    current_group = []

    for msg in msgs:
        need_separator = len(current_group) > 0
        l_separator = len(separator) if need_separator else 0

        if total_length + len(msg.text) + l_separator > limit_chars:
            grouped_msgs.append(current_group)
            current_group = []
            total_length = 0

        current_group.append(msg)
        total_length += len(msg.text) + l_separator

    if len(current_group) > 0:
        grouped_msgs.append(current_group)

    return grouped_msgs

@schedule.every(20).seconds.do
@distributed_scheduling_job(interval_seconds=20)
def _flush():
    redis_connection: Redis = reusable_redis_connection()

    length_queue = redis_connection.llen(const.TELEGRAM_MESSAGE_LIST_REDIS_KEY)
    # message_list = redis_connection.lrange(TELEGRAM_MESSAGE_LIST_REDIS_KEY, 0, -1)
    # message_json_list = [TelegramMessage.model_validate(msg) for msg in message_list]

    by_room = {}

    for msg in range(length_queue):
        msg = redis_connection.lpop(const.TELEGRAM_MESSAGE_LIST_REDIS_KEY)
        msg = json.loads(msg)
        msg = TelegramMessage.model_validate(msg)

        if not msg.can_batch:

            send_message(
                msg.sender,
                msg.text,
                room=msg.room,
                fmt=msg.parse_mode,
                preview_opt=msg.link_preview_options,
            )

            continue

        if msg.room not in by_room:
            by_room[msg.room] = []

        by_room[msg.room].append(msg)

    sep = "\n" + "-" * 20 + "\n"

    for room, msgs in by_room.items():
        groups = group_message(msgs, sep)

        for group in groups:
            joint_message = sep.join([msg.text for msg in group])
            send_message("junk_notifications", joint_message, room=room, fmt="HTML")
