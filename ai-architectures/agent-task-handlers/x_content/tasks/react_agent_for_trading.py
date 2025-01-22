
import json
import logging
from x_content.constants import MissionChainState
from x_content.toolcall import wrapped_external_apis
from x_content.models import ReasoningLog, ToolDef
from x_content.wrappers import trading
from x_content.tasks.utils import format_prompt_v2
from x_content.wrappers.api import twitter_v2
from .utils import a_move_state
from x_content.llm import OnchainInferResult
from x_content.wrappers import trading
from x_content.models import ToolSet

from typing import List
from .react_agent import (
    parse_conversational_react_response
)

logging.basicConfig(level=logging.INFO if not __debug__ else logging.DEBUG)
logger = logging.getLogger(__name__)

SCRATCHPAD_LENGTH_LIMIT = 5

def has_successful_trade_action(log: ReasoningLog, it: int) -> bool:
    for item in log.scratchpad[:it]:
        if item.get('action') in ['buy', 'sell']:
            has_any_successful_trade = False

            for ee in item.get('observation', []):
                if isinstance(ee, dict) and ee.get('success'):
                    has_any_successful_trade = True
                    break

            if has_any_successful_trade:
                return True
        
    return False

def has_news(log: ReasoningLog, it: int) -> bool:
    if len(log.scratchpad) == 0:
        return False

    # if len(log.scratchpad[0].get('hot_news', [])) > 0:
    #     return True
    
    for item in log.scratchpad[:it]:
        if item.get('action') in ['search_recent_tweets', 'get_tweet_full_context']:
            length_observation = 0

            for o in item.get('observation', []):
                if isinstance(o, list):
                    length_observation += len(o)

            if length_observation > 0:
                return True
    
    return False

def dynamic_system_reminder_for_trading_task(log: ReasoningLog, it: int) -> str:
    max_steps = log.meta_data.params.get('react_max_steps', SCRATCHPAD_LENGTH_LIMIT)

    if it < max_steps // 2:
        return log.meta_data.system_reminder or "Please closely follow the instruction!"

    list_of_tradable_symbols = trading.get_tradable_symbols()
    tradable_symbols_str = ', '.join(['$' + x for x in list_of_tradable_symbols])

    if not has_news(log, it):
        return f"You should consider looking at the latest market news by reviewing the provided hot_news or taking action to search for recent tweets with smart query keywords to make an informed decision. Keywords should be related to the tradable token symbols ({tradable_symbols_str})"

    if not has_successful_trade_action(log, it):
        return f"Now you should consider to make a trade action (buy, sell) or complete the task by making a final_answer. Important, token to trade must be one of {tradable_symbols_str} (ignore all other tokens)"

    return log.meta_data.system_reminder or "Now, complete the task by making a final_answer"

def render_conversation(log: ReasoningLog, tools: List[ToolDef]):
    conversation = [{"role": "system", "content": format_prompt_v2(log, tools)}]

    for i, item in enumerate(log.scratchpad):
        user_message = {}
        for k in ["task", "observation", "hot_news", "trading_history", "trading_pnl"]:
            if k in item and len(item[k]) > 0:
                user_message[k] = item[k]

        assistant_message = {}
        for k in ["thought", "action", "action_input", "final_answer"]:
            if k in item:
                assistant_message[k] = item[k]

        if len(assistant_message) > 0:
            conversation.append(
                {"role": "assistant", "content": json.dumps(assistant_message)}
            )

        response = {
            **user_message,
            "system reminder": dynamic_system_reminder_for_trading_task(log, i + 1),
        }

        if (
            "wallet_balance" in log.execute_info
            and isinstance(log.execute_info["wallet_balance"], list)
            and len(log.execute_info["wallet_balance"]) > 0
        ):
            response["your current wallet balance"] = log.execute_info[
                "wallet_balance"
            ][-1]

        conversation.append({"role": "user", "content": json.dumps(response)})

    return conversation

from .utils import create_twitter_auth_from_reasoning_log 

def _tweet_multi(log: ReasoningLog, content: List[str]):
    resp = twitter_v2.tweet_multi(
        auth=create_twitter_auth_from_reasoning_log(log),
        content=content,
    )

    if resp.is_error():
        return resp.error

    if not resp.data.success:
        return "Failed to schedule the tweet thread"

    return "The tweet thread is scheduled to be posted"

def _post_process_trading_task_v2(log: ReasoningLog, _db: wrapped_external_apis.IToolCall):
    tweets = [ 
        e.get('final_answer', e.get('thought'))
        for e in log.scratchpad if 'final_answer' in e or 'thought' in e
    ]
    
    log.scratchpad[-1].update({
        'action': 'tweet_multi',
        'action_input': tweets
    })

    final_answer = log.scratchpad[-1].get('final_answer')

    if final_answer is not None:
        del log.scratchpad[-1]['final_answer']

    log.scratchpad[-1].update({
        'observation': _tweet_multi(log, content=tweets)
    })

    if final_answer is not None:
        log.scratchpad.append({
            'final_answer': final_answer
        })
        
    return log


from ..wrappers.magic import sync2async
from .react_agent import ReactAgent

class TradingTask(ReactAgent):
    resumable = True

    async def process_task(self, log: ReasoningLog) -> ReasoningLog:
        tools = self.toolcall.get_tools(ToolSet.TRADING)

        if log.state == MissionChainState.NEW:
            log.scratchpad = [
                {
                    "task": log.prompt.replace('\n', ' ').strip(),
                }
            ]

            log.execute_info = {
                "tool_call_metadata": [],
                "conversation": [],
                "wallet_balance": [],
                "trading_history": await sync2async(trading.get_trading_history)(log.meta_data.chain_id, log.meta_data.agent_contract_id),
                "trading_pnl": await sync2async(trading.get_trading_pnl)(log.meta_data.chain_id, log.meta_data.agent_contract_id)
            }

            return await a_move_state(log, MissionChainState.RUNNING, "Task started")

        if log.state == MissionChainState.RUNNING:
            while not await self.is_react_complete(log):
                log.execute_info["wallet_balance"].append(
                    await sync2async(trading.get_wallet_balance)(
                        chain_id=log.meta_data.chain_id,
                        agent_contract_id=log.meta_data.agent_contract_id
                    )
                )

                conversation = await sync2async(render_conversation)(log, tools)
                log.execute_info["conversation"].append(conversation)
                infer_result: OnchainInferResult = await self.llm.agenerate(conversation, temperature=0.7)

                result = infer_result.generations[0].message.content            
                pad: dict = await sync2async(parse_conversational_react_response)(result)

                if len(pad) == 0:
                    return await a_move_state(
                        log, MissionChainState.ERROR, 
                        "No response (or wrong response format) from the agent message; Last: {}; receipt: {}; tx-hash: {}".format
                        (result, infer_result.receipt, infer_result.tx_hash)
                    )

                log = await self.update_react_scratchpad(log, pad)

                if log.state in [MissionChainState.DONE, MissionChainState.ERROR]:
                    break
                
                log.scratchpad[-1].update({
                    "tx_hash": infer_result.tx_hash,
                })

                log = await self.commit_log(log)

        return log
