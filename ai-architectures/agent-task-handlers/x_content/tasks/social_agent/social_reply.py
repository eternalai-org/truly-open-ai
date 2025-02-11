from typing import List, Optional
from x_content.constants.main import GAME_TASKS_WHITELIST
from x_content.tasks.reply_base import ReplyTaskBase
from x_content.tasks.reply_subtask_base import ReplySubtaskBase
from x_content.models import ReasoningLog
from x_content.wrappers.api.twitter_v2.models.objects import ExtendedTweetInfo
from x_content.wrappers.tweet_specialty import TweetSpecialty

from . import subtasks


class SocialReplyTask(ReplyTaskBase):

    def get_subtask_cls(
        self,
        log: ReasoningLog,
        specialties: List[TweetSpecialty],
        tweet_info: ExtendedTweetInfo,
    ) -> Optional[ReplySubtaskBase]:
        if TweetSpecialty.CREATE_GAME in specialties:
            # Game agent should not reply create game tweet as a social agent
            if log.meta_data.twitter_username in GAME_TASKS_WHITELIST:
                return None
            return subtasks.reply.reply_game.ReplyGameSubtask
        if TweetSpecialty.CREATE_GAME_SUBTREE in specialties:
            return None
        if TweetSpecialty.TOKEN_ANALYSIS in specialties:
            # Handled by a different service
            return None
        return subtasks.reply.reply_regular.ReplyRegularSubtask
