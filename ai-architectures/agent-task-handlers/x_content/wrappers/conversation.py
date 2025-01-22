import json
from typing import List

from x_content.wrappers.api.twitter_v2.main import get_relevent_information_v2
from x_content.wrappers.api.twitter_v2.models.objects import TweetObject
from x_content.wrappers.junk_tasks import choose_suitable_language
from x_content.wrappers.knowledge_base.base import KnowledgeBase
from x_content.wrappers.magic import sync2async


def get_enhance_tweet_conversation(system_prompt: str, content: str, example_tweets: List[str] = []):
    if example_tweets is not None and example_tweets != []:
        tweets_str = '\n'.join([f"- {x}" for x in example_tweets])
        system_prompt = f"""{system_prompt}

Here are example tweets written in the defined style:
{tweets_str}
"""

    prompt_to_use = """
Act as an expert in creative writing and AI-generated content. Your task is to analyze a given base tweet and rewrite it in your own unique personality and writing style while preserving the core message.

First, carefully read through the base tweet and identify its key points and overall sentiment. Consider the tone, language, and any specific elements that make the tweet distinctive.

Next, think about your own writing style and personality. How can you infuse the base tweet with your own voice while still maintaining its original intent? Consider aspects such as word choice, sentence structure, emojis or slang that align with your persona.

Now, begin rewriting the tweet step-by-step. Start with the main idea and progressively incorporate your personal touch. Ensure that each modification enhances the tweet's authenticity while staying true to its core message.

After completing the rewrite, review it critically. Does it effectively convey the original meaning while showcasing your unique personality? If not, refine it further until you're satisfied with the result.

Return the rewritten tweet as a stringified JSON with the key "tweet".

Example JSON Response:
{{
    "tweet": "Your rewritten tweet here"
}}

Remember to maintain a consistent tone throughout and avoid introducing unrelated topics or altering the central theme of the original post. Ensure that your response is always in this exact JSON format, with the rewritten tweet stringified under the key "tweet".

Now, provide your rewritten tweet in JSON format based on the given base tweet: {}
""".format(content)

    messages = [
        {
            "role": "system",
            "content": system_prompt,
        },
        {
            "role": "user",
            "content": prompt_to_use,
        }
    ]

    return messages

async def get_reply_tweet_conversation(system_prompt: str, tweets: List[TweetObject], kn_base: KnowledgeBase):
    chat_history = [
        {
            "user": x.twitter_username,
            "message": x.full_text,
        }
        for x in tweets
    ]

    language = await sync2async(choose_suitable_language)(chat_history)
    resp = await get_relevent_information_v2(kn_base, tweets=tweets)
    knowledge = {} if resp.is_error() else resp.data.to_dict()

    user_prompt_template = """Provide a single message to join the following conversation in {} language. \
Keep it concise (under 128 chars), NO hashtags, links or emojis, and don't include any instructions or extra words, \
just the raw message ready to post.

Chat History: {}"""

    user_prompt = user_prompt_template.format(language, json.dumps(chat_history))
    if knowledge is not None:
        user_prompt += "\n\nRelevant Knowledge: {}".format(json.dumps(knowledge))

    conversational_chat = [
        {
            "role": "system",
            "content": system_prompt,
        },
        {
            "role": "user",
            "content": user_prompt,
        }
    ]

    return conversational_chat
