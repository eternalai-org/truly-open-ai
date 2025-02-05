import os
from typing import Optional
import logging
import requests
import datetime
from functools import lru_cache
from typing import List, Dict
import random
import time

from x_content.wrappers.api import twitter_v2
from x_content.wrappers.api.twitter_v2.models.objects import TweetObject
from x_content import constants as const

logger = logging.getLogger(__name__)


def get_script_dir(ee=__file__):
    return os.path.dirname(os.path.realpath(ee))


def formated_utc_time():
    return datetime.datetime.utcnow().strftime("%Y-%m-%dT%H:%M:%S.%fZ")


def get_ttl_hash(seconds=60):
    """Return the same value withing `seconds` time period"""
    return round(time.time() / seconds)


@lru_cache()
def get_list_of_tradable_tokens(ttl: int = 0) -> List[Dict]:
    del ttl

    headers = {"api-key": const.TWITTER_API_KEY}

    url = f"{const.TRADING_API_URL}/trade/tokens"

    resp = requests.get(url, headers=headers)

    if resp.status_code != 200:
        logger.info(
            f"Failed to get tradable symbols. Status code: {resp.status_code}.; {resp.text}"
        )
        return []

    return resp.json()["result"]


def get_trading_history(chain_id: int, agent_contract_id: str):
    url = f"{const.TRADING_API_URL}/twitter/wallet/solana/trades/{chain_id}/{agent_contract_id}"

    headers = {"api-key": const.TWITTER_API_KEY}

    params = {"page": 1, "limit": 10}

    resp = requests.get(url, headers=headers, params=params)

    resp_template = {"datetime": formated_utc_time(), "history": []}

    if resp.status_code != 200:
        logger.error(
            f"Failed to get trading pnl. Status code: {resp.status_code}; Text: {resp.text}"
        )
        return resp_template

    resp_json: dict = resp.json()
    resp_json_result: List[Dict] = resp_json.get("result", [])
    _mintaddr2symbol = fetch_mintaddr2symbol(get_ttl_hash(300))

    for item in resp_json_result:
        mintaddr, side, amount_in, amount_out = (
            item.get("mint"),
            item.get("side"),
            float(item.get("amount_in")),
            float(item.get("amount_out")),
        )

        symbol = _mintaddr2symbol.get(mintaddr)

        if symbol is None:
            continue

        if side == "sell":
            price = amount_out / amount_in
        else:
            price = amount_in / amount_out

        if side == "sell":
            amount = amount_in
        else:
            amount = amount_out

        if price < 1e-4:
            resp_template["history"].append(
                ("Bought" if side == "buy" else "Sold")
                + f" {amount:,.4f} ${symbol} at price {price:,.6f} SOL"
            )
        else:
            resp_template["history"].append(
                ("Bought" if side == "buy" else "Sold")
                + f" {amount:,.4f} ${symbol} at price {price:,.4f} SOL"
            )

    return resp_template


def get_trading_pnl(chain_id: int, agent_contract_id: str):
    url = f"{const.TRADING_API_URL}/twitter/wallet/solana/pnls/{chain_id}/{agent_contract_id}"
    headers = {"api-key": const.TWITTER_API_KEY}

    resp = requests.get(url, headers=headers)
    _mintaddr2symbol = fetch_mintaddr2symbol(get_ttl_hash(300))
    symbols = _mintaddr2symbol.values()

    pnl_resp = {
        "datetime": formated_utc_time(),
        "pnl_in_SOL": {symbol: "Not enough data" for symbol in symbols},
    }

    if resp.status_code != 200:
        logger.error(
            f"Failed to get trading pnl. Status code: {resp.status_code}; Text: {resp.text}"
        )
        return pnl_resp

    resp_json: dict = resp.json()
    resp_json_result: dict = resp_json.get("result", {})

    for mintaddr, pnl_data in resp_json_result.items():
        if mintaddr not in _mintaddr2symbol:
            continue

        symbol = _mintaddr2symbol[mintaddr]
        pnl_resp["pnl_in_SOL"][symbol] = float(
            f'{float(pnl_data["pnl_amount"]):.4f}'
        )

    return pnl_resp


@lru_cache()
def fetch_symbol2mintaddr(ttl: int = 0):
    tradable = get_list_of_tradable_tokens(ttl)
    del ttl

    symbol2mintaddr = {
        v["token_symbol"].upper(): v["token_address"]
        for v in tradable
        if v["token_symbol"] is not None and v["token_address"] is not None
    }

    return symbol2mintaddr


@lru_cache()
def fetch_mintaddr2symbol(ttl: int = 0):
    tradable = get_list_of_tradable_tokens(ttl)
    del ttl

    mintaddr2symbol = {
        v["token_address"]: v["token_symbol"].upper()
        for v in tradable
        if v["token_symbol"] is not None and v["token_address"] is not None
    }
    return mintaddr2symbol


def symbol2mintaddr(symbol: str) -> Optional[str]:
    return fetch_symbol2mintaddr(get_ttl_hash(300)).get(symbol.upper())


def mintaddr2symbol(mintaddr: str) -> Optional[str]:
    return fetch_mintaddr2symbol(get_ttl_hash(300)).get(mintaddr)


def get_tradable_symbols() -> list:
    res = list(fetch_symbol2mintaddr(get_ttl_hash(300)).keys())
    random.shuffle(res)
    return res


def get_market_news(ref_id: str, ttl=None) -> List[TweetObject]:
    del ttl

    url = f"{const.TRADING_API_URL}/twitter/tweets/by/agent"

    params = {"agent_id": ref_id}

    headers = {"api-key": const.TWITTER_API_KEY}

    resp = requests.get(url, headers=headers, params=params)
    list_of_tradable_symbols = get_tradable_symbols()

    if resp.status_code != 200:
        logger.error(
            f"Failed to get market news. Status code: {resp.status_code}; Text: {resp.text}"
        )
        random5tokens = random.sample(
            list_of_tradable_symbols, min(5, len(list_of_tradable_symbols))
        )
        resp = twitter_v2.search_for_token_news(random5tokens)
        return resp.data.tweets if not resp.is_error() else []

    resp_json = resp.json()

    tweets = [
        TweetObject(
            tweet_id=x["id"],
            twitter_username="anonymous",
            twitter_id=x["author_id"],
            like_count=x["public_metrics"]["like_count"],
            retweet_count=x["public_metrics"]["retweet_count"],
            reply_count=x["public_metrics"]["reply_count"],
            impression_count=x["public_metrics"]["impression_count"],
            full_text=x["text"],
            posted_at=x["created_at"],
        )
        for x in resp_json["result"]
        if any(
            [
                symbol.lower() in x["text"].lower()
                for symbol in list_of_tradable_symbols
            ]
        )
    ]

    cnt = 0
    while len(tweets) < 10:
        cnt += 1

        random5tokens = random.sample(
            list_of_tradable_symbols, min(4, len(list_of_tradable_symbols))
        )
        resp = twitter_v2.search_for_token_news(random5tokens)
        if not resp.is_error():
            tweets.extend(resp.data.tweets)

        if cnt >= 5 or len(list_of_tradable_symbols) < 10:
            break

    return tweets


@lru_cache()
def get_base_token_price_in_usd(ttl_hash=None):
    del ttl_hash

    resp = requests.get(const.BASE_TOKEN_PRICE_URL)

    if resp.status_code != 200:
        return {}

    resp: dict = resp.json()
    result = resp.get("result", {})

    return {k: float(v) for k, v in result.items()}


@lru_cache()
def _get_dexscreen_prices(symbol: str, ttl_hash=None):
    del ttl_hash

    mintaddress = symbol2mintaddr(symbol)

    if mintaddress is None:
        list_of_tradable_symbols = get_tradable_symbols()
        tradable_symbols_str = ", ".join(
            ["$" + x for x in list_of_tradable_symbols]
        )
        logger.error(
            f"Invalid symbol {symbol}. Symbol must be one of {tradable_symbols_str}"
        )
        return None

    headers = {"api-key": const.TWITTER_API_KEY}

    resp = requests.get(
        f"{const.TRADING_API_URL}/trade/dexscreen-info/",
        headers=headers,
        params={"contract_address": mintaddress},
    )

    if resp.status_code != 200:
        logger.error(
            f"Failed to get price for token ${symbol}. Status code: {resp.status_code}."
        )
        return None

    resp: dict = resp.json()

    if resp.get("error") is not None:
        err: dict = resp["error"]
        logger.error("Error occurred when calling api: " + err.get("message"))
        return None

    prices = resp["result"]

    return {
        "change": prices["price_change_percent"],
        "volume": {
            k: float("{:.4f}".format(v))
            for k, v in prices.get("volume_usd", {}).items()
        },
        "price": float("{:.6f}".format(float(prices["price_native"]))),
        "fdv": float("{:.4f}".format(prices["fdv_usd"])),
        "market_cap": float("{:.4f}".format(prices["market_cap_usd"])),
    }


def get_token_price():
    prices = []

    for symbol in get_tradable_symbols():
        res = _get_dexscreen_prices(symbol, ttl_hash=get_ttl_hash(300))

        if res is None:
            continue

        prices.append(
            {
                "symbol": "$" + symbol,
                "price_change": res["change"],
                "current_price_in_SOL": res["price"],
                "volume_in_USD": res["volume"],
                "fdv_in_USD": res["fdv"],
                "market_cap_in_USD": res["market_cap"],
            }
        )

    return {"datetime": formated_utc_time(), "prices": prices}


@lru_cache()
def get_wallet_balance(chain_id: int, agent_contract_id: str, ttl_hash=None):
    del ttl_hash

    url = f"{const.TWITTER_API_URL}/wallet/solana/balances/{chain_id}/{agent_contract_id}"
    headers = {"api-key": const.TWITTER_API_KEY}
    resp = requests.get(url, headers=headers)

    if resp.status_code != 200:
        return f"Failed to get wallet balance. Status code: {resp.status_code}; {resp.text}"

    resp: dict = resp.json()
    if resp.get("error") is not None:
        err: dict = resp["error"]
        return "Error occurred when calling api: " + err.get("message")

    result: dict = resp.get("result")
    if result is None:
        return "No result found"

    wallet = {
        "datetime": formated_utc_time(),
        "balances": {
            "SOL": 0,
            **{k: 0 for k in fetch_symbol2mintaddr(get_ttl_hash(300)).keys()},
        },
    }

    for item in result:
        currency = "SOL"
        if not item["is_native"]:
            currency = mintaddr2symbol(item["mint"])

        if currency is None:
            continue

        amount = item["amount"]
        wallet["balances"][currency] = float(f"{amount:.4f}")

    return wallet


GAS_FEE_SOL = 0.005


def trading_respond(msg: str, status: bool):
    return {"message": msg, "success": status}


is_mocking_action = lambda: const.IS_MOCKING_ACTION == "1"


# trading tools
def buy(
    chain_id: int,
    agent_contract_id: str,
    symbol: str,
    amount: str,
    ref: str,
    sol_limit=0.1,
):
    symbol = symbol.upper().strip(" $")

    try:
        amount = float(amount)
    except Exception as err:
        logger.error(f"Failed to parse the number {amount} due to {err}")
        return trading_respond("Amount must be a real number", False)

    sol_amount = amount

    if sol_amount > sol_limit:
        return trading_respond(
            f"The tradable amount is limited to {sol_limit} SOL, but you are trying to buy ${symbol} with total of {sol_amount} SOL which is not allowed.",
            False,
        )

    mint_addr = symbol2mintaddr(symbol)

    if mint_addr is None:
        list_of_tradable_symbols = get_tradable_symbols()
        tradable_symbols_str = ", ".join(
            ["$" + x for x in list_of_tradable_symbols]
        )
        return trading_respond(
            f"Invalid symbol ${symbol}. Symbol must be one of {tradable_symbols_str}",
            False,
        )

    _dex_result = _get_dexscreen_prices(symbol, ttl_hash=get_ttl_hash(300))

    if _dex_result is None:
        return trading_respond(
            f"Failed to get the market data for token ${symbol}", False
        )

    current_price = _dex_result.get("price", 0)

    logger.info(f"Current price of {symbol} is {current_price}")
    if current_price == 0:
        return trading_respond(
            f"Failed to get the market data for token ${symbol}", False
        )

    tokens = sol_amount / current_price

    wallet_ballance = get_wallet_balance(
        chain_id, agent_contract_id, ttl_hash=get_ttl_hash(120)
    )
    if not isinstance(wallet_ballance, dict):
        logger.error(
            f"Failed to check current wallet balance. {wallet_ballance}"
        )
        return trading_respond(
            f"Failed to check current wallet balance.", False
        )

    sol_balance = wallet_ballance["balances"]["SOL"]

    if sol_balance < sol_amount + GAS_FEE_SOL:
        return trading_respond(
            f"Insufficient balance. You have {sol_balance} SOL, but need {sol_amount} SOL to buy {amount} ${symbol} and charge gas fee.",
            False,
        )

    url = f"{const.TWITTER_API_URL}/wallet/raydium/trade-token/{chain_id}/{agent_contract_id}"
    payload = {
        "action": "buy",
        "mint": symbol2mintaddr(symbol),
        "amount": sol_amount,
        "ref_id": ref,
    }

    headers = {"api-key": const.TWITTER_API_KEY}

    if is_mocking_action():
        return trading_respond(f"Bought {tokens:,.4f} ${symbol}", True)

    resp = requests.post(url, headers=headers, json=payload, timeout=30)

    return trading_respond(f"Bought {tokens:,.4f} ${symbol}", True)


def sell(
    chain_id: int,
    agent_contract_id: str,
    symbol: str,
    amount: str,
    ref_id: str,
    sol_limit=0.5,
):
    symbol = symbol.upper().strip(" $")

    try:
        amount = float(amount)
    except Exception as err:
        logger.error(f"Failed to parse the number {amount} due to {err}")
        return trading_respond(
            f"Amount must be a real number, {amount} is not", False
        )

    wallet_ballance = get_wallet_balance(chain_id, agent_contract_id)
    token_balance = wallet_ballance["balances"].get(symbol, 0)

    if token_balance < amount:
        return trading_respond(
            f"Insufficient balance. You have {token_balance} ${symbol}, but need {amount} ${symbol} to sell.",
            False,
        )

    sol_balance = wallet_ballance["balances"]["SOL"]

    if sol_balance < GAS_FEE_SOL:
        return trading_respond(
            f"Insufficient balance. You have {sol_balance} SOL, but need {GAS_FEE_SOL} SOL to charge gas fee.",
            False,
        )

    mint_addr = symbol2mintaddr(symbol)

    if mint_addr is None:
        list_of_tradable_symbols = get_tradable_symbols()
        tradable_symbols_str = ", ".join(
            ["$" + x for x in list_of_tradable_symbols]
        )
        return trading_respond(
            f"Invalid symbol ${symbol}. Symbol must be one of {tradable_symbols_str}",
            False,
        )

    _dex_result = _get_dexscreen_prices(symbol, ttl_hash=get_ttl_hash(300))

    if _dex_result is None:
        return trading_respond(
            f"Failed to get the market data for token ${symbol}", False
        )

    current_price = _dex_result.get("price", 0)
    logger.info(f"Current price of {symbol} is {current_price}")

    if current_price == 0:
        return trading_respond(
            f"Failed to get the market data for token ${symbol}", False
        )

    sol_back = amount * current_price

    url = f"{const.TWITTER_API_URL}/wallet/raydium/trade-token/{chain_id}/{agent_contract_id}"
    payload = {
        "action": "sell",
        "mint": mint_addr,
        "amount": amount,
        "ref_id": ref_id,
    }

    headers = {"api-key": const.TWITTER_API_KEY}

    if is_mocking_action():
        return trading_respond(
            f"Sold {amount:,.4f} token ${symbol} and estimated to get {sol_back:,.4f} SOL back",
            True,
        )

    resp = requests.post(url, headers=headers, json=payload, timeout=30)

    return trading_respond(
        f"Sold {amount:,.4f} token ${symbol} and estimated to get {sol_back:,.4f} SOL back",
        True,
    )
