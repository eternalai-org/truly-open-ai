from .utils import get_and_warn

GAME_BASE_URL = get_and_warn("GAME_BASE_URL")

GAME_EMOJIS = ["🎮", "🎯", "🎪", "🎲", "🤔"]
GAME_KEYWORDS = ["game", "play", "guess", "riddle"]
# GAME_DURATION = 60 * 60 * 4 # 4 hours
GAME_DURATION = 60 * 10  # 10 minutes
GAME_REDIS_CACHE = 24 * 60 * 60 * 2  # 2 day
CREATE_GAME_LOCK_EXPIRY = 1000 * 10
# 10 seconds
JUDGE_GAME_LOCK_EXPIRY = 1000 * 60 * 60 * 2
# 2 hours

# Reply tweet templates
GAME_CREATED_TWEET = "A challenge is created! Place your bet by deposit EAI on Base to the wallet of who you think will win:\n\n Bet will close in {hours:02d}:{minutes:02d}\n\n"
# Constants for tweet replies
WINNER_TWEET_TEMPLATE = "The winner is {}"
NO_WINNER_TWEET = "No winner"

CREATE_GAME_PREFIX_PATTERN = "[CREATE_GAME]"  # TODO: To be defined
