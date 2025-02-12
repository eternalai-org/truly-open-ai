from typing import List, Optional
from x_content.tasks.game_agent.subtasks.reply.reply_create_game import is_create_game_pending
from x_content.tasks.reply_base import ReplyTaskBase
from x_content.tasks.reply_subtask_base import ReplySubtaskBase
from x_content.models import ReasoningLog
from x_content.wrappers.api.twitter_v2.models.objects import ExtendedTweetInfo
from x_content.wrappers.tweet_specialty import TweetSpecialty

from . import subtasks


class GameReplyTask(ReplyTaskBase):

    def get_subtask_cls(
        self,
        log: ReasoningLog,
        specialties: List[TweetSpecialty],
        tweet_info: ExtendedTweetInfo,
    ) -> Optional[ReplySubtaskBase]:
        if TweetSpecialty.CREATE_GAME in specialties:
            if is_create_game_pending(tweet_info.tweet_object.tweet_id):
                return subtasks.reply.reply_create_game.CreateGameSubtask
        return None
