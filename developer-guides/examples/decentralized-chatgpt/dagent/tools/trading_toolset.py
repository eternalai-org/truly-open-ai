from typing import List
from dagent.models import Tool, ToolParam, ToolParamDtype
from . import functional
from dagent.registry import RegistryCategory, register_decorator
from . base_toolset import Toolset
from dagent import constant as C

# not supported in the current version
# @register_decorator(RegistryCategory.ToolSet)
class TradingToolset(Toolset):
    TOOLSET_NAME = "trading"
    PURPOSE = f"to buy, sell, and get information about tokens (tradable tokens: {functional.tradable_symbols()})"

    TOOLS: List[Tool] = [
        Tool(
            name="buy",
            description="Buy and hold a token",
            param_spec=[
                ToolParam(
                    name="symbol",
                    dtype=ToolParamDtype.STRING,
                    description="The symbol of the token to buy"
                ),
                ToolParam(
                    name="amount",
                    dtype=ToolParamDtype.NUMBER,
                    description="Amount to buy"
                )
            ],
            executor=lambda symbol, amount: functional.buy(C.CHAIN_ID, C.CONTRACT_ID, symbol, amount)
        ),
        Tool(
            name="sell",
            description="Sell a token",
            param_spec=[
                ToolParam(
                    name="symbol",
                    dtype=ToolParamDtype.STRING,
                    description="The symbol of the token to sell"
                ),
                ToolParam(
                    name="amount",
                    dtype=ToolParamDtype.NUMBER,
                    description="Amount to sell"
                )
            ],
            executor=lambda symbol, amount: functional.sell(C.CHAIN_ID, C.CONTRACT_ID, symbol, amount)
        ),
        Tool(
            name="get_wallet_balance",
            description="Get wallet balance", 
            param_spec=[],
            executor=lambda: functional.get_wallet_balance(C.CHAIN_ID, C.CONTRACT_ID)
        ),
        Tool(
            name="get_token_price",
            description="Get token price",
            param_spec=[
                ToolParam(
                    name="symbol",
                    dtype=ToolParamDtype.STRING,
                    description="The symbol of the token"
                )
            ],
            executor=functional.get_token_price
        )
    ]