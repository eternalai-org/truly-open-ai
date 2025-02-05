from typing import List, Optional
from x_content.tasks.reply_base import ReplyTaskBase
from x_content.tasks.reply_subtask_base import ReplySubtaskBase
from x_content.models import ReasoningLog
from x_content.wrappers.tweet_specialty import TweetSpecialty

from . import subtasks


class SocialReplyTask(ReplyTaskBase):

    def get_subtask_cls(
        self, log: ReasoningLog, specialties: List[TweetSpecialty]
    ) -> Optional[ReplySubtaskBase]:
        if TweetSpecialty.CREATE_GAME in specialties:
            return subtasks.reply.reply_game.ReplyGameSubtask
        if TweetSpecialty.CREATE_GAME_SUBTREE in specialties:
            return None
        if TweetSpecialty.TOKEN_ANALYSIS in specialties:
            # Handled by a different service
            return None
        return subtasks.reply.reply_regular.ReplyRegularSubtask
