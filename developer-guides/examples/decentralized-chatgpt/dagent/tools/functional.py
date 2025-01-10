import requests
from dagent.models import TweetObject, TwitterUserObject
import logging
from typing import List, Union, Optional
from dagent.utils import formated_utc_time
from dagent import constant as C
import re
from bs4 import BeautifulSoup

logger = logging.getLogger(__name__)

def _preprocess_twitter_username(username: str) -> str:
    username = username.lstrip('@')

    if username.startswith('username='):
        username = username.split('=')[1].strip(" \"")

    return username

def get_user_info_by_username(username: str) -> TwitterUserObject:
    username = _preprocess_twitter_username(username)

    url = f"{C.ETERNAL_X_API}/user/by/username/{username}"
    headers = {
        'api-key': C.ETERNAL_X_API_APIKEY
    }
    resp = requests.get(url, headers=headers)

    if resp.status_code != 200:
        logger.error(f"Something went wrong (status code: {resp.status_code})")
        return []

    resp: dict = resp.json()

    if resp.get("error", None) is not None:
        err: dict = resp["error"]
        logger.error(f"Error occured when calling api: {err.get('message')}")
        return []

    info = resp["result"]

    if info["id"] == "":
        logger.error(f"No user found with username: {username}")
        return [] 

    return TwitterUserObject(
        twitter_id=info["id"],
        twitter_username=info["username"],
        name=info["name"],
        followers_count=info["public_metrics"]["followers_count"],
        followings_count=info["public_metrics"]["following_count"],
        is_blue_verified=info["verified"],
    )
    

def get_engaged_tweets_by_topic(query: str, top_k=C.DEFAULT_TOP_K) -> List[TweetObject]:
    try:
        url = f"{C.ETERNAL_X_API}/tweets/search/recent"
        params = {
            "query": query,
        }
        headers = {
            'api-key': C.ETERNAL_X_API_APIKEY
        }

        resp = requests.get(
            url, params=params, headers=headers
        )

        if resp.status_code != 200:
            logger.error(f"Something went wrong (status code: {resp.status_code})")
            return []

        resp: dict = resp.json()
        data = resp["result"]

        if resp.get("error") is not None:
            err: dict = resp["error"]
            logger.error(f"Error occurred when calling API: {err.get('message')}")
            return []

        if len(data["LookUps"]) == 0:
            logger.error(f"No tweets found with query: {query}")
            return []

        tweets = []
        for id, item in data["LookUps"].items():
            tweet = item["Tweet"]
            user = item["User"]

            tweets.append(TweetObject(
                tweet_id=tweet["id"],
                twitter_username=user["username"],
                twitter_id=tweet["author_id"],
                like_count=tweet["public_metrics"]["like_count"],
                retweet_count=tweet["public_metrics"]["retweet_count"],
                reply_count=tweet["public_metrics"]["reply_count"],
                impression_count=tweet["public_metrics"]["impression_count"],
                full_text=tweet["text"],
                posted_at=tweet["created_at"],
            ))

        tweets = sorted(tweets, key=lambda e: e.impression_count, reverse=True)
        return tweets[:top_k]
    except Exception as e:
        logger.error(f"An exception occurred: {str(e)}")
        return []


def crawl_data_from_url(url: str):
    try:
        # Send a GET request
        response = requests.get(url)
        
        # If the GET request is successful, the status code will be 200
        if response.status_code == 200:
            # Get the content of the response
            page_content = response.content
            
            # Create a BeautifulSoup object and specify the parser
            soup = BeautifulSoup(page_content, 'html.parser')
            
            # Remove all script and style elements
            for script in soup(["script", "style"]):
                script.decompose()  # decompose script and style elements
            
            # Get the text from the BeautifulSoup object
            text = soup.get_text()
            
            # Break the text into lines and remove leading and trailing space on each
            lines = (line.strip() for line in text.splitlines())
            
            # Break multi-headlines into a line each
            chunks = (phrase.strip() for line in lines for phrase in line.split("  "))
            
            # Drop blank lines
            text = '\n'.join(chunk for chunk in chunks if chunk)
            
            return text
        
        else:
            return None
    
    except Exception as e:
        return ""
        
def get_full_text(full_text: str):
    if "http" in full_text:
        # Extract URLs from full_text
        urls = re.findall(r'(https?://\S+)', full_text)
        # Crawl data from each URL and append to full_text
        for url in urls:
            crawled_text = crawl_data_from_url(url)
            if crawled_text:
                full_text = full_text.replace(url, crawled_text)
    return full_text
    
def get_tweet_info_from_tweet_id(self, tweet_id: str):
    try:
        url = f"{C.ETERNAL_X_API}/tweets"
        params = {
            "ids": tweet_id
        }
        headers = {
            'api-key': C.ETERNAL_X_API_APIKEY
        }
        resp = requests.get(
            url, headers=headers, params=params
        )
        resp.raise_for_status()
        result = resp.json().get("result", {})
        for key, value in result.items():
            tweet = value["Tweet"]
            user = value["User"]
            tweet_object = TweetObject(
                tweet_id=tweet["id"],
                twitter_username=user["username"],
                twitter_id=user["id"],
                full_text=tweet["text"],
                posted_at=tweet["created_at"],
            )
            parent_tweet_id = None
            if key != tweet["conversation_id"]:
                reference_tweets = [] if tweet["referenced_tweets"] is None else tweet["referenced_tweets"]
                parent_tweet_id = next((ref_tweet["id"] for ref_tweet in reference_tweets if ref_tweet["type"] == "replied_to"), None)
            return {
                "tweet_object": tweet_object,
                "parent_tweet_id": parent_tweet_id
            }
    except Exception as e:
        logger.error(f"An error occurred: {e}")
        return {
            "tweet_object": None,
            "parent_tweet_id": None
        }
    
def get_full_context_from_a_tweet(tweet_object, parent_tweet_id: str = None):
    tweets = [tweet_object]
    while parent_tweet_id is not None:
        tweet_info = get_tweet_info_from_tweet_id(parent_tweet_id)
        tweet_object = tweet_info["tweet_object"]
        parent_tweet_id = tweet_info["parent_tweet_id"]
        full_text = get_full_text(tweet_object.full_text)
        tweet_object["full_text"] = full_text
        tweets.insert(0, tweet_object)
    return tweets

def get_tweets_by_username_v2(username: str, num_tweets = 1, replied = 0):
    try:
        url = f"{C.ETERNAL_X_API}/tweets/by/username/{username}"
        headers = {
            'api-key': C.ETERNAL_X_API_APIKEY
        }
        params = {
            "replied": replied
        }
        resp = requests.get(
            url, headers=headers, params=params
        )
        
        if resp.status_code != 200:
            return f"Something went wrong (status code: {resp.status_code})"
        
        resp = resp.json()            
        if resp.get("error") is not None:
            return "Error occurred when calling API: " + resp["error"]["message"]
        
        tweets = resp["result"]["data"][:num_tweets]
        res = []
        for tweet in tweets:
            reference_tweets = [] if tweet["referenced_tweets"] is None else tweet["referenced_tweets"]
            reference_tweets = {}
            parent_tweet_id = next((ref_tweet["id"] for ref_tweet in reference_tweets if ref_tweet["type"] == "replied_to"), None)
            author_id = tweet["author_id"]
            user_url = f"{C.ETERNAL_X_API}/user/{author_id}"
            user_resp = requests.get(user_url, headers=headers)
            user_resp_json = user_resp.json()
            username = user_resp_json["result"]["username"]
            full_text = get_full_text(tweet["text"])
            tweet_object = TweetObject(
                tweet_id=tweet["id"],
                twitter_username=username,
                twitter_id=author_id,
                full_text=full_text,
                posted_at=tweet["created_at"],
            )
            if parent_tweet_id is not None:
                all_tweets = get_full_context_from_a_tweet(tweet_object, parent_tweet_id)
                res.append(all_tweets)
            else:
                res.append([tweet_object])
        return res
    except Exception as e:
        return []
    
def get_recent_mentioned_tweets_by_username_v2(username: str, num_tweets=1, replied=0):
    try:
        url = f"{C.ETERNAL_X_API}/user/by/username/{username}/mentions"
        headers = {
            'api-key': C.ETERNAL_X_API_APIKEY
        }
        params = {
            "replied": replied
        }
        resp = requests.get(url, headers=headers, params=params)

        if resp.status_code != 200:
            return f"Something went wrong (status code: {resp.status_code})"

        resp_json = resp.json()
        if resp_json.get("error"):
            return f"Error occurred when calling API: {resp_json['error']['message']}"

        tweets = resp_json["result"]["data"][:num_tweets]
        if not tweets:
            return "No tweets found"

        res = []
        for tweet in tweets:
            reference_tweets = [] if tweet["referenced_tweets"] is None else tweet["referenced_tweets"]
            reference_tweets = {}
            parent_tweet_id = next((ref_tweet["id"] for ref_tweet in reference_tweets if ref_tweet["type"] == "replied_to"), None)
            author_id = tweet["author_id"]
            user_url =f"{C.ETERNAL_X_API}/user/{author_id}"
            user_resp = requests.get(user_url, headers=headers)
            user_resp_json = user_resp.json()
            username = user_resp_json["result"]["username"]
            full_text = get_full_text(tweet["text"])
            tweet_object = {
                "twitter_id": author_id,
                "tweet_id": tweet["id"],
                "twitter_username": username,
                "full_text": full_text,
                "posted_at": tweet["created_at"]
            }
            if parent_tweet_id is not None:
                all_tweets = get_full_context_from_a_tweet(tweet_object, parent_tweet_id)
                res.append(all_tweets)
            else:
                res.append([tweet_object])
        return res
    except Exception as e:
        return f"An error occurred: {str(e)}"

def find_user(query: str, top_k=C.DEFAULT_TOP_K) -> List[TwitterUserObject]:

    url = f"{C.ETERNAL_X_API}/user/search/"
    headers = {
        'api-key': C.ETERNAL_X_API_APIKEY
    }
    params = {
        "query": query,
    }

    resp = requests.get(
        url, 
        params=params,
        headers=headers
    )

    if resp.status_code != 200:
        logger.error(f"Something went wrong (status code: {resp.status_code}; reason: {resp.text})")
        return []

    resp: dict = resp.json()
    if resp.get("error") is not None:
        err: dict = resp["error"]
        logger.error(f"Error occured when calling api: {err.get('message')}")
        return []

    if len(resp["result"]) == 0:
        logger.error(f"No user found with query: {query}")
        return []
    
    users = [
        TwitterUserObject(
            twitter_id=x["id"],
            twitter_username=x["username"],
            name=x["name"],
            followers_count=x["public_metrics"]["followers_count"],
            followings_count=x["public_metrics"]["following_count"],
            is_blue_verified=x["verified"],
        ) for x in resp["result"][:top_k]  
    ]

    return users

def get_recent_mentioned_tweets(username: str, top_k=C.DEFAULT_TOP_K) -> List[TweetObject]:
        username = _preprocess_twitter_username(username)
        url = f"{C.ETERNAL_X_API}/user/by/username/{username}/mentions"
        headers = {
            'api-key': C.ETERNAL_X_API_APIKEY
        }
        resp = requests.get(url, headers=headers)
        
        if resp.status_code != 200:
            logger.error(f"Something went wrong (status code: {resp.status_code})")
            return []
        
        resp: dict = resp.json()    

        if resp.get("error") is not None:
            err: dict = resp["error"]
            logger.error(f"Error occured when calling api: {err.get('message')}")
            return []
        
        tweets = resp["result"]["data"]
        if len(tweets) == 0:
            logger.error(f"No tweets found with username: {username}")
            return []

        tweets = [
            TweetObject(
                tweet_id=x["id"],
                twitter_username=username,
                twitter_id=x["author_id"],
                like_count=x["public_metrics"]["like_count"],
                retweet_count=x["public_metrics"]["retweet_count"],
                reply_count=x["public_metrics"]["reply_count"],
                impression_count=x["public_metrics"]["impression_count"],
                full_text=x["text"],
                posted_at=x["created_at"],
            ) 
            for x in tweets[:top_k]
        ]
        
        return tweets

def get_tweets_by_username(username: str, top_k=C.DEFAULT_TOP_K) -> List[TweetObject]:
    username = _preprocess_twitter_username(username)

    url = f"{C.ETERNAL_X_API}/tweets/by/username/{username}"
    headers = {
        'api-key': C.ETERNAL_X_API_APIKEY
    }
    resp = requests.get(url, headers=headers)
    
    if resp.status_code != 200:
        logger.error(f"Something went wrong (status code: {resp.status_code})")
        return []
    
    resp: dict = resp.json()    

    if resp.get("error") is not None or resp["result"]["data"] is None:
        err: dict = resp["error"]
        logger.error("Error occured when calling api: " + err.get("message"))
        return []
    
    tweets = resp["result"]["data"]
    if len(tweets) == 0:
        logger.error(f"No tweets found with username: {username}")
        return []
    
    tweets = [
        TweetObject(
            tweet_id=x["id"],
            twitter_username=username,
            twitter_id=x["author_id"],
            like_count=x["public_metrics"]["like_count"],
            retweet_count=x["public_metrics"]["retweet_count"],
            reply_count=x["public_metrics"]["reply_count"],
            impression_count=x["public_metrics"]["impression_count"],
            full_text=x["text"],
            posted_at=x["created_at"],
        ) 
        for x in tweets[:top_k]
    ]

    return tweets

def get_following_users_by_username(username: str, top_k: int=20, only_name=True) -> Union[List[TwitterUserObject], str]:
    username = _preprocess_twitter_username(username)

    url = f"{C.ETERNAL_X_API}/user/by/username/{username}/following"
    headers = {
        'api-key': C.ETERNAL_X_API_APIKEY
    }
    resp = requests.get(url, headers=headers)

    if resp.status_code != 200:
        logger.error(f"Something went wrong (status code: {resp.status_code})")
        return []

    resp: dict = resp.json()

    if resp.get("error") is not None:
        err: dict = resp["error"]
        logger.error(f"Error occured when calling api: {err.get('message')}")
        return []

    followings = resp["result"]
    if len(followings) == 0:
        logger.error(f"No followings found with username: {username}")
        return []

    if not only_name:
        users = [
            TwitterUserObject(
                twitter_id=x["rest_id"],
                twitter_username=x["screen_name"],
                name=x["name"],
                followers_count=x["followers_count"],
                followings_count=x["friends_count"],
                is_blue_verified=x["is_blue_verified"],
            ) for x in followings[:top_k]
        ]

        return users

    usernames = list(map(lambda x: x["screen_name"], followings))
    return str(", ".join(usernames))

def _perform_twitter_action_and_get_result(url: str, headers: dict, payload: dict) -> str:
    payload["is_testing"] = C.IS_SANDBOX
    response = requests.post(url, headers=headers, json=payload)
    
    if response.status_code != 200:
        return f"Request failed with status code: {response.status_code} - {response.text}"
    
    try:
        data = response.json()
    except ValueError:
        return "Failed to parse response as JSON; But the request was successful; Raw response: " + response.text

    if data.get("error") is not None:
        return f"API Error: {data['error']}"

    return data

def follow(target_username: str):
    action_input = {
        "target_username": target_username,
    }
    
    payload = {
        'action_type': 'follow',
        'action_input': action_input
    }
    
    url = f"{C.ETERNAL_X_API}/user/action"

    headers = {
        "Content-Type": "application/json",
        "api-key": C.ETERNAL_X_API_APIKEY
    }
    
    return _perform_twitter_action_and_get_result(url, headers, payload)

def reply(tweet_id: str, reply_content: str):
    action_input = {
        "tweet_id": tweet_id,
        "comment": reply_content
    }
    
    payload = {
        'action_type': 'reply',
        'action_input': action_input
    }

    headers = {
        "Content-Type": "application/json",
        "api-key": C.ETERNAL_X_API_APIKEY
    }

    url = f"{C.ETERNAL_X_API}/user/action" 
    
    return _perform_twitter_action_and_get_result(url, headers, payload)
   
def quote_tweet(tweet_id: str, comment: str):
    action_input = {
        "tweet_id": tweet_id,
        "comment": comment
    }
    
    payload = {
        'action_type': 'quote_tweet',
        'action_input': action_input
    }
    
    headers = {
        "Content-Type": "application/json",
        "api-key": C.ETERNAL_X_API_APIKEY
    }
    
    url = f"{C.ETERNAL_X_API}/user/action"
    
    return _perform_twitter_action_and_get_result(url, headers, payload)

def tweet(content: str):
    action_input = {
        "content": content
    }
    
    payload = {
        'action_type': 'tweet',
        'action_input': action_input
    }
    
    headers = {
        "Content-Type": "application/json",
        "api-key": C.ETERNAL_X_API_APIKEY
    }
    
    url = f"{C.ETERNAL_X_API}/user/action"
    
    return _perform_twitter_action_and_get_result(url, headers, payload)


TOKENS_INFO: list = [
    {
        "symbol": "Goat",
        "mint_address": "CzLSujWBLFsSjncfkh59rUFqvafWcY5tzedWJSuypump"
    },
    {
        "symbol": "Fartcoin",
        "mint_address": "9BB6NFEcjBCtnNLFko2FqVQBq8HHM13kCyYcdQbgpump"
    },
    {
        "symbol": "Cents",
        "mint_address": "C9FVTtx4WxgHmz55FEvQgykq8rqiLS8xRBVgqQVtpump"
    },
    {
        "symbol": "Zerebro",
        "mint_address": "8x5VqbHA8D7NkD52uNuS5nnt3PwA8pLD34ymskeSo2Wn"
    },
    {
        "symbol": "Bully",
        "mint_address": "79yTpy8uwmAkrdgZdq6ZSBTvxKsgPrNqTLvYQBh1pump"
    },
    {
        "symbol": "Shegen",
        "mint_address": "2KgAN8nLAU74wjiyKi85m4ZT6Z9MtqrUTGfse8Xapump"
    },
    {
        "symbol": "Yousim",
        "mint_address": "66gsTs88mXJ5L4AtJnWqFW6H2L5YQDRy4W41y6zbpump"
    }
]
    
SYMBOL2MINTADDR = {v["symbol"]: v["mint_address"] for v in TOKENS_INFO}
MINTADDR2SYMBOL = {v["mint_address"]: v["symbol"] for v in TOKENS_INFO}

def symbol2mintaddr(symbol: str) -> Optional[str]:
    return SYMBOL2MINTADDR.get(symbol)

def mintaddr2symbol(mintaddr: str) -> Optional[str]:
    return MINTADDR2SYMBOL.get(mintaddr)

def tradable_symbols() -> list:
    return list(SYMBOL2MINTADDR.keys())

# trading tools
def buy(chain_id: int, agent_contract_id: str, symbol: str, amount: str):
    if C.IS_SANDBOX:
        amount = 0.001

    try:
        amount = float(amount)
    except Exception as err:
        logger.error(f"Failed to parse the number {amount} due to {err}")
        return "Amount must be a real number"

    mint_addr = symbol2mintaddr(symbol)

    if mint_addr is None:
        return f"Invalid symbol {symbol}. Symbol must be one of {', '.join(tradable_symbols())}"

    url = f"{C.ETERNAL_X_API}/wallet/raydium/trade-token/{chain_id}/{agent_contract_id}"
    payload = {
        "action": "buy",
        "mint": symbol2mintaddr(symbol),
        "amount": amount
    }
    
    headers = {
        'api-key': C.ETERNAL_X_API_APIKEY
    }
    
    resp = requests.post(
        url, 
        headers=headers, 
        json=payload
    )
    
    if resp.status_code != 200:
        return f"Failed to buy {amount} token {symbol}. Status code: {resp.status_code}."
    
    resp: dict = resp.json()    
    if resp.get("error") is not None:
        err: dict = resp["error"]
        return "Error occured when calling api: " + err.get("message")
    
    return f"Bought {amount} token {symbol}"

def sell(chain_id: int, agent_contract_id: str, symbol: str, amount: str):
    if C.IS_SANDBOX:
        amount = 0.001

    try:
        amount = float(amount)
    except Exception as err:
        logger.error(f"Failed to parse the number {amount} due to {err}")
        return "Amount must be a real number, {amount} is not"

    mint_addr = symbol2mintaddr(symbol)

    if mint_addr is None:
        return f"Invalid symbol {symbol}. Symbol must be one of {', '.join(tradable_symbols())}"

    url = f"{C.ETERNAL_X_API}/wallet/raydium/trade-token/{chain_id}/{agent_contract_id}"
    payload = {
        "action": "sell",
        "mint": symbol2mintaddr(symbol),
        "amount": amount
    }
    
    headers = {
        'api-key': C.ETERNAL_X_API_APIKEY
    }

    resp = requests.post(
        url, 
        headers=headers, 
        json=payload
    )

    if resp.status_code != 200:
        return f"Failed to sell {amount} token {symbol}. Status code: {resp.status_code}."

    resp: dict = resp.json()    
    if resp.get("error") is not None:
        err: dict = resp["error"]
        return "Error occured when calling api: " + err.get("message")
    
    return f"Sold {amount} token {symbol}"

def get_wallet_balance(chain_id: int, agent_contract_id: str):
    url = f"{C.ETERNAL_X_API}/wallet/solana/balances/{chain_id}/{agent_contract_id}"
    headers = {
        'api-key': C.ETERNAL_X_API_APIKEY
    }
    resp = requests.get(url, headers=headers)

    if resp.status_code != 200:
        return f"Failed to get wallet balance. Status code: {resp.status_code}."
    
    resp: dict = resp.json()
    if resp.get("error") is not None:
        err: dict = resp["error"]
        return "Error occured when calling api: " + err.get("message")

    result: dict = resp.get("result")
    if result is None:
        return "No result found"

    wallet = []

    for item in result:
        currency = "SOL"
        if not item['is_native']:
            currency = mintaddr2symbol(item['mint'])
        
        amount = item['amount']
        wallet.append(f"{amount} {currency}")

    return f"Your wallet balance: " + ", ".join(wallet)

def get_token_price(symbol: str):
    mint_address = symbol2mintaddr(symbol)
    
    if mint_address is None:
        return f"Invalid symbol {symbol}. Symbol must be one of {', '.join(tradable_symbols())}"
    
    url = f"{C.ETERNAL_X_API}/wallet/pumfun/price/{mint_address}"
    headers = {
        'api-key': C.ETERNAL_X_API_APIKEY
    }
    resp = requests.get(url, headers=headers)
    
    if resp.status_code != 200:
        return f"Failed to get price for token {symbol}. Status code: {resp.status_code}."
    
    resp: dict = resp.json()
    if resp.get("error") is not None:
        err: dict = resp["error"]
        return "Error occured when calling api: " + err.get("message")
    
    result: dict = resp.get("result")
    if result is None:
        return "No result found"
    
    dtime = formated_utc_time()
    return f"[{dtime}] Current price of token {symbol} is {result}"