from functools import lru_cache
import logging
import requests
from enum import IntEnum
from x_content.cache.entity_cache import (
    GameRedisCache,
)
from x_content.wrappers.magic import helpful_raise_for_status
from x_content.wrappers import telegram
from x_content import constants as const

logger = logging.getLogger(__name__)


@lru_cache(maxsize=1)
def _get_game_redis_cache():
    return GameRedisCache()


# Backend game status constants
class GameState(IntEnum):
    PENDING = 0
    RUNNING = 1
    ENDED = 2
    RESULT_UPDATED = 3
    COMPLETED = 4


# Game status constants
class GameStatus:
    NONE = ""
    CREATE_PENDING = "create_pending"
    CREATE_RUNNING = "create_running"
    CREATE_DONE = "create_done"
    JUDGE_PENDING = "judge_pending"
    JUDGE_RUNNING = "judge_running"
    JUDGE_DONE = "judge_done"


""" Example payload from BE
{
  "status": 1,
  "data": {
    "id": "67765bc7de1b1ad7b9d58a2a",
    "date_created": "2025-01-02T09:26:31.699Z",
    "date_modified": "2025-01-02T09:26:31.699Z",
    "modified_user_id": "",
    "created_user_id": "",
    "date_deleted": "0001-01-01T00:00:00Z",
    "tweet_id": "123123111",
    "start_time": "2025-01-02T09:26:31.699Z",
    "end_time": "0001-01-01T00:00:00Z",
    "agent_wallets": [
      {
        "username": "1",
        "address": "0x3043eb7b95dab80758f35795d38b0f70670b5701",
        "amount": 0
      },
      {
        "username": "2",
        "address": "0x8e95330db6e5a6252b508c0254528966131fc0eb",
        "amount": 0
      }
    ],
    "status": 1,
    "players": null,
    "winner": ""
  }
}
"""


class GameAgentWallet:

    def __init__(self, username: str, address: str, amount: float):
        self.username = username
        self.address = address
        self.amount = amount

    @classmethod
    def from_dict(cls, data: dict):
        return cls(
            username=data.get("username", ""),
            address=data.get("address", ""),
            amount=data.get("amount", 0.0),
        )


class GameInfo:

    def __init__(
        self,
        id: str,
        date_created: str,
        date_modified: str,
        modified_user_id: str,
        created_user_id: str,
        date_deleted: str,
        tweet_id: str,
        start_time: str,
        end_time: str,
        agent_wallets: list[GameAgentWallet],
        status: int,
        players: list = None,
        winner: str = "",
    ):
        self.id = id
        self.date_created = date_created
        self.date_modified = date_modified
        self.modified_user_id = modified_user_id
        self.created_user_id = created_user_id
        self.date_deleted = date_deleted
        self.tweet_id = tweet_id
        self.start_time = start_time
        self.end_time = end_time
        self.agent_wallets = agent_wallets
        self.status = GameState(status)
        self.players = players if players else []
        self.winner = winner

    @classmethod
    def from_dict(cls, game_data: dict):
        return cls(
            id=game_data.get("id", ""),
            date_created=game_data.get("date_created", ""),
            date_modified=game_data.get("date_modified", ""),
            modified_user_id=game_data.get("modified_user_id", ""),
            created_user_id=game_data.get("created_user_id", ""),
            date_deleted=game_data.get("date_deleted", ""),
            tweet_id=game_data.get("tweet_id", ""),
            start_time=game_data.get("start_time", ""),
            end_time=game_data.get("end_time", ""),
            agent_wallets=[
                GameAgentWallet.from_dict(w)
                for w in game_data.get("agent_wallets", [])
            ],
            status=GameState(game_data.get("status", 0)),
            players=game_data.get("players", []),
            winner=game_data.get("winner", ""),
        )


class GameAPIClient:
    """Client for interacting with the game management API endpoints"""

    @staticmethod
    def request(method: str, endpoint: str, params=None, data=None, json=None):
        """
        Makes an HTTP request to the game API

        Args:
            method (str): HTTP method (GET, POST, etc)
            endpoint (str): API endpoint path
            params (dict, optional): URL query parameters
            data (dict, optional): Form data
            json (dict, optional): JSON request body

        Returns:
            tuple: (response_data, error)
                - response_data (dict): Parsed JSON response if successful, None if error
                - error (Exception): Exception if request failed, None if successful
        """
        url = f"{const.GAME_BASE_URL}{endpoint}"
        try:
            response = requests.request(
                method=method,
                url=url,
                headers={"Content-Type": "application/json"},
                params=params,
                data=data,
                json=json,
                timeout=30,  # Add 30 second default timeout
            )
            helpful_raise_for_status(response)
            logger.info(
                f"[GameAPIClient.request] Successfully called {method} {url} with response: {response.json()}"
            )
            return response.json(), None
        except requests.exceptions.RequestException as e:
            logger.error(
                f"[GameAPIClient.request] API call to {url} failed with error: {str(e)}"
            )
            telegram.send_message(
                "junk_nofitications",
                f"```bash\nGame API request failed (status code: {e.response.status_code if hasattr(e, 'response') else 'N/A'}; url: {url}; error: {str(e)})\n```",
                room=telegram.TELEGRAM_ALERT_ROOM,
            )
            return None, e

    @staticmethod
    def get_game_info_by_tweet_id(
        tweet_id: str,
    ) -> tuple[GameInfo | None, Exception | None]:
        """
        Gets information about a game by tweet ID

        Args:
            tweet_id (str): ID of tweet containing the game

        Returns:
            tuple: (game_info, error)
                - game_info (GameInfo): Game information if successful
                - error (Exception): Error if request failed
        """
        endpoint = f"/v1/game/{tweet_id}"
        response_data, err = GameAPIClient.request(
            method="GET",
            endpoint=endpoint,
        )

        if err is None and response_data.get("status") == 1:
            return GameInfo.from_dict(response_data.get("data", {})), None
        else:
            logger.info(
                f"[GameAPIClient.get_game_info_by_tweet_id] Failed to get game info for tweet {tweet_id}. Status: {response_data.get('status') if response_data else 'No response'}"
            )
            return None, Exception(
                f"Failed to get game info for tweet {tweet_id}"
            )

    @staticmethod
    def get_game_info_by_tweet_ids(
        tweet_ids: list[str],
    ) -> tuple[list[GameInfo] | None, Exception | None]:
        """
        Gets information about games by tweet IDs

        Args:
            tweet_ids (list[str]): List of tweet IDs containing the games

        Returns:
            tuple: (game_infos, error)
                - game_infos (list[GameInfo]): List of game information if successful
                - error (Exception): Error if request failed
        """
        endpoint = "/v1/game"
        # Convert list of tweet IDs into repeated query params
        params = [("tweet_ids[]", tweet_id) for tweet_id in tweet_ids]

        response_data, err = GameAPIClient.request(
            method="GET", endpoint=endpoint, params=params
        )

        if err is None and response_data.get("status") == 1:
            games_data = response_data.get("data", {}).get("games", [])
            game_infos = [
                GameInfo.from_dict(game_data) for game_data in games_data
            ]
            logger.info(
                f"[GameAPIClient.get_game_info_by_tweet_ids] Successfully retrieved info for {len(game_infos)} games"
            )
            return game_infos, None
        else:
            logger.info(
                f"[GameAPIClient.get_game_info_by_tweet_ids] Failed to get game info for tweets {tweet_ids}. Status: {response_data.get('status') if response_data else 'No response'}"
            )
            return None, Exception(
                f"Failed to get game info for tweets {tweet_ids}"
            )

    @staticmethod
    def end_game(tweet_id: str):
        """
        Marks a game as ended

        Args:
            tweet_id (str): ID of tweet containing the game

        Returns:
            tuple: (game_info, error)
                - game_info (GameInfo): Game information if successful
                - error (Exception): Error if request failed
        """
        endpoint = f"/v1/game/{tweet_id}/end"
        response_data, err = GameAPIClient.request(
            method="POST",
            endpoint=endpoint,
        )

        if err is None and response_data.get("status") == 1:
            logger.info(
                f"[GameAPIClient.end_game] Successfully ended game for tweet {tweet_id}"
            )
            return GameInfo.from_dict(response_data.get("data", {})), None
        else:
            logger.info(
                f"[GameAPIClient.end_game] Failed to end game for tweet {tweet_id}. Status: {response_data.get('status') if response_data else 'No response'}"
            )
            return None, Exception(f"Failed to end game for tweet {tweet_id}")

    @staticmethod
    def start_game(
        tweet_id: str, agent_usernames: list, time_out: int
    ) -> tuple[GameInfo | None, Exception | None]:
        """
        Starts a new game

        Args:
            tweet_id (str): ID of tweet containing the game
            agent_usernames (list): List of usernames of participating agents
            time_out (int): Game timeout duration in seconds

        Returns:
            tuple: (game_info, error)
                - game_info (GameInfo): Game information if successful
                - error (Exception): Error if request failed
        """
        endpoint = "/v1/game/start"
        payload = {
            "tweet_id": tweet_id,
            "usernames": agent_usernames,
            "time_out": time_out,
            "bet_time_out": time_out,
        }

        response_data, err = GameAPIClient.request(
            method="POST",
            endpoint=endpoint,
            json=payload,
        )

        if err is None and response_data.get("status") == 1:
            logger.info(
                f"[GameAPIClient.start_game] Successfully started game for tweet {tweet_id} with agents {agent_usernames}"
            )
            return GameInfo.from_dict(response_data.get("data", {})), None
        else:
            logger.info(
                f"[GameAPIClient.start_game] Failed to start game for tweet {tweet_id}. Status: {response_data.get('status') if response_data else 'No response'}"
            )
            return None, Exception(
                f"Failed to start game for tweet {tweet_id}"
            )

    @staticmethod
    def submit_game_result(tweet_id: str, agent_username: str):
        """
        Submits the result/winner for a game

        Args:
            tweet_id (str): ID of tweet containing the game
            agent_username (str): Username of winning agent

        Returns:
            tuple: (game_info, error)
                - game_info (GameInfo): Game information if successful
                - error (Exception): Error if request failed
        """

        endpoint = f"/v1/game/{tweet_id}/result"
        payload = {"username": agent_username}

        response_data, err = GameAPIClient.request(
            method="POST",
            endpoint=endpoint,
            json=payload,
        )

        if err is None and response_data.get("status") == 1:
            logger.info(
                f"[GameAPIClient.submit_game_result] Successfully submitted game result for tweet {tweet_id}. Winner: {agent_username}"
            )
            return GameInfo.from_dict(response_data.get("data", {})), None

        else:
            logger.info(
                f"[GameAPIClient.submit_game_result] Failed to submit game result for tweet {tweet_id}. Status: {response_data.get('status') if response_data else 'No response'}"
            )
            return None, Exception(
                f"Failed to submit game result for tweet {tweet_id}"
            )
