from json_repair import repair_json

from x_content.wrappers.conversation import get_enhance_tweet_conversation
from x_content.wrappers.junk_tasks import postprocess_tweet_by_prompts
from x_content.wrappers.magic import retry, sync2async
from .utils import a_move_state, create_twitter_auth_from_reasoning_log, get_system_prompt_with_random_example_tweets
from x_content.wrappers.api.twitter_v2.models.response import GenerateActionDto, Response
from .base import MultiStepTaskBase
from x_content.models import ReasoningLog, MissionChainState
from x_content.wrappers.api import twitter_v2
from x_content.wrappers.twin_agent import get_random_example_tweets
from x_content.llm.eternal_ai import OnchainInferResult
import json
from typing import List
import logging
import datetime
import random
import re
import httpx
import asyncio
from x_content.wrappers.postprocess import remove_urls

async def split_urls(text: str) -> List[str]:
    return re.findall(r'(https?://[^\s]+)', text)

async def resolve_headers(url: str) -> str:
    async with httpx.AsyncClient() as client:
        resp = await client.head(url)
        return resp.headers

logger = logging.getLogger(__name__)

class PostV3(MultiStepTaskBase):
    resumable = True

    POST_V3_PROMPT_TEMPLATE = '''
{task_definition}

Content to use:
{news_to_use}
'''

    async def process_task(self, log: ReasoningLog) -> ReasoningLog:
        try:
            mission_configuration = json.loads(log.prompt)
        except json.JSONDecodeError as e:
            logger.error("Error while parsing prompt: %s", e)
            return await a_move_state(log, MissionChainState.ERROR, f"Error while parsing prompt: {e}.") 

        if log.state == MissionChainState.NEW:
            missing_fields = [
                k for k in ['prompt'] 
                if k not in mission_configuration
            ]

            if len(missing_fields) > 0:
                return await a_move_state(log, MissionChainState.ERROR, f"Missing fields: {missing_fields} in prompt")

            prompt = mission_configuration['prompt']
            cutoff_hour = int(mission_configuration.get('cutoff_hour', 2))
            
            content_to_use = await self.get_contents(
                twitter_username=log.meta_data.twitter_username,
                time_cutoff_hours=cutoff_hour,
                top_k=5
            )

            if len(content_to_use) == 0:
                return await a_move_state(log, MissionChainState.ERROR, f"No engaging tweets found for the past {cutoff_hour} hours from top following users")

            messages = [
                {
                    'role': 'system',
                    'content': get_system_prompt_with_random_example_tweets(log),
                },
                {
                    'role': 'user',
                    'content': self.POST_V3_PROMPT_TEMPLATE.format(
                        task_definition=prompt,
                        news_to_use='\n'.join('- ' + e for e in content_to_use)
                    )
                }
            ]

            log.execute_info = {
                'content_to_use': content_to_use,
                'conversation': [messages],
            }            

            return await a_move_state(log, MissionChainState.RUNNING, "Task started")

        if log.state == MissionChainState.RUNNING:
            prompt = mission_configuration['prompt']

            async def get_base_tweet():
                base_tweet_conversation = log.execute_info['conversation'][-1]
                result: OnchainInferResult = await self.llm.agenerate(base_tweet_conversation, temperature=0.7)
                base_tweet = result.generations[0].message.content
                return base_tweet, result.tx_hash

            try:
                base_tweet, base_tweet_tx_hash = await retry(get_base_tweet, max_retry=3, first_interval=10, interval_multiply=2)()
            except Exception as err:
                return await a_move_state(log, MissionChainState.ERROR, f"Failed to generate base tweet: {err}")

            log.execute_info['task_result'] = [
                {
                    "base_tweet": base_tweet,
                    "base_tweet_tx_hash": base_tweet_tx_hash,
                }
            ]

            log.execute_info['conversation'].append(
                get_enhance_tweet_conversation(
                    system_prompt=log.agent_meta_data.persona,
                    content=base_tweet,
                    example_tweets=get_random_example_tweets(log.meta_data.knowledge_base_id),
                )
            )

            async def get_enhanced_tweet():
                enhance_tweet_conversation = log.execute_info['conversation'][-1]
                result: OnchainInferResult = await self.llm.agenerate(enhance_tweet_conversation, temperature=0.01)
                assistant_message = result.generations[0].message.content
                data = repair_json(assistant_message, return_objects=True)
                return data["tweet"], result.tx_hash

            try:
                enhanced_tweet, enhanced_tweet_tx_hash = await retry(get_enhanced_tweet, max_retry=3, first_interval=10, interval_multiply=2)()
            except Exception as err:
                return await a_move_state(log, MissionChainState.ERROR, f"Failed to enhance the tweet: {err}")

            log.execute_info['task_result'][-1].update({
                "enhanced_tweet": enhanced_tweet,
                "enhanced_tweet_tx_hash": enhanced_tweet_tx_hash,
            })

            async def postprocess_tweet():
                return await sync2async(postprocess_tweet_by_prompts)(
                    system_prompt=log.agent_meta_data.persona,
                    task_prompt=prompt,
                    tweet=enhanced_tweet
                )
            
            postprocessed_tweet = await retry(postprocess_tweet, max_retry=3, first_interval=10, interval_multiply=2)()

            resp: Response[GenerateActionDto] = await sync2async(twitter_v2.tweet)(
                auth=create_twitter_auth_from_reasoning_log(log),
                content=postprocessed_tweet,
                tx_hash=base_tweet_tx_hash,
            )
            
            if resp.is_error():
                return await a_move_state(log, MissionChainState.ERROR, f"Failed to perform tweet action: {resp.error}")
            
            if not resp.data.success:
                return await a_move_state(log, MissionChainState.ERROR, f"Failed to schedule tweet")
 
            log.execute_info['task_result'][-1].update({
                "postprocessed_tweet": postprocessed_tweet,
            })

            return await a_move_state(log, MissionChainState.DONE, "Task completed")

        return log

    @staticmethod
    async def get_contents(twitter_username, time_cutoff_hours=2, top_k=15) -> List[str]:
        resp = await sync2async(twitter_v2._get_following_by_username)(twitter_username)

        if resp.data is None:
            logger.warning("Error while retrieving following users for username %s. Message: %s", twitter_username, resp.error)
            return []

        sorted_by_followers_count = sorted(resp.data, key=lambda x: x['followers_count'], reverse=True)

        n = len(sorted_by_followers_count)
        # sorted_by_followers_count = sorted_by_followers_count[:max(top_k, 3 * n // 4)]
        random_pick_for_top_k = random.sample(sorted_by_followers_count, k=min(n, top_k))

        tweets: List[twitter_v2.TweetObject] = []

        datetime_format = "%Y-%m-%dT%H:%M:%S.%fZ"
        tpoint = datetime.datetime.now() - datetime.timedelta(hours=time_cutoff_hours)

        for twitter_account in random_pick_for_top_k:
            resp = await sync2async(twitter_v2.get_tweets_by_username)(
                username=twitter_account['screen_name'], 
                top_k=10
            )

            if resp.data is None:
                logger.warning("Error while retrieving tweets for username %s. Message: %s", twitter_account['screen_name'], resp.error)
                continue

            _tweets = [
                e
                for e in resp.data.tweets
                if datetime.datetime.strptime(e.posted_at, datetime_format) > tpoint
            ]

            for t in _tweets:
                if t.reference != [] and 'quoted' not in [e['type'] for e in t.reference]:
                    continue

                tweets.append(t)
   
        if len(tweets) == 0:
            logger.info(f"No engaging tweets found for the past {time_cutoff_hours} hours from top following users")
            return []

        sorted_tweets_by_impression_count: List[twitter_v2.TweetObject] = sorted(
            tweets, 
            key=lambda x: x.impression_count,
            reverse=True
        )

        selected_tweet = random.choice(sorted_tweets_by_impression_count[:10])
        logger.info("Selected tweet: %s; by: %s; impression_count: %d", selected_tweet.tweet_id, selected_tweet.twitter_username, selected_tweet.impression_count)

        response_tweets = [selected_tweet]
        
        for reference in selected_tweet.reference:
            if reference['type'] in 'quotedrepost':
                e = await sync2async(twitter_v2.get_tweet_info_from_tweet_id)(reference['id'], preserve_img=True)
                if e.data is not None:
                    response_tweets.append(e.data.tweet_info.tweet_object)

        resp: Response[twitter_v2.TweetsDto] = await sync2async(twitter_v2.search_recent_tweet_by_tweetid)(
            selected_tweet.tweet_id, limit_observation=100
        )
        recent_replies = resp.data.tweets

        for child_tweet in recent_replies:
            if child_tweet.twitter_username != selected_tweet.twitter_username:
                continue

            response_tweets.append(child_tweet)

        pat = re.compile(r'https://x.com/.+/(\d+)/photo/\d+')
        mask = {}
        n_tweets_with_media = 0

        for i, tweet in enumerate(response_tweets):
            mask[tweet.tweet_id] = False
            for url in tweet.media:
                mat = pat.match(url)

                if mat is not None:
                    n_tweets_with_media += 1
                    mask[tweet.tweet_id] = True
                    break
        
        response_tweets = sorted(
            response_tweets, 
            key=lambda x: mask[x.tweet_id],
            reverse=True
        )
        
        futures = [
            asyncio.ensure_future(
                sync2async(twitter_v2._image_descriptions_from_tweet_id)(tweet_id=tweet.tweet_id)
            )
            for tweet in response_tweets
            if mask[tweet.tweet_id]
        ]
        
        image_descriptions = await asyncio.gather(*futures)
        
        for i, (tweet, image_description) in enumerate(zip(response_tweets, image_descriptions)):
            if len(image_description) > 0:
                response_tweets[i].full_text += '\n\n' + '\n\n'.join(image_description)

        for i in range(len(response_tweets)):
            response_tweets[i].full_text = remove_urls(response_tweets[i].full_text)
        
        return [
            # f"{e.twitter_username}: {e.full_text}"
            e.full_text
            for e in response_tweets
        ]