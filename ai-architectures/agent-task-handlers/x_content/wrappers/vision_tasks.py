from x_content import constants as const
from x_content.wrappers import redis_wrapper
from x_content.wrappers.llm_tasks import logger

import requests


@redis_wrapper.cache_for(3600 * 24 * 30)
def get_image_description(url: str):
    return """A small, orange and white kitten stands on a stone wall, looking attentively at something out of frame.

The kitten is centered in the image.  The background is a blurred, out-of-focus area of green foliage, creating a natural backdrop. The kitten's posture suggests alertness, looking to something off to its left or right of the frame. The focus is on the kitten, and the surrounding greenery is softly blurred, directing attention to the subject. 


The kitten is a small, likely young, orange and white cat. Its fur appears soft and its eyes are a light amber color. The kitten is standing on a stone wall that appears natural. 


The image is a photograph, likely taken outdoors, with natural light. The lighting highlights the kitten's fur, creating clear definition and texture. The style is naturalistic and not manipulated or stylized.


The setting is an outdoor environment, likely a garden or park.  The wall the kitten is standing on is light gray stone. The blurry background suggests a natural, lush environment.  The lighting is ambient, providing soft illumination and making the colors appear true to life. The mood is tranquil and inviting, drawing attention to the kitten's presence."""
    base_url = f"{const.VISION_API_URL}/v1/chat/completions"
    headers = {"Authorization": f"Bearer {const.VISION_API_KEY}"}

    payload = {
        "messages": [
            {
                "role": "user",
                "content": [
                    {
                        "type": "text",
                        "text": "Analyze this image and describe its semantic meaning in detail with a single paragraph. Include the implied context, emotions, relationships between elements, possible symbolic interpretations, and the overall message or story conveyed by the image.",
                    },
                    {
                        "type": "image_url",
                        "image_url": {
                            "url": "https://pbs.twimg.com/media/GjchTR9bQAAlnlB.jpg"
                        },
                    },
                ],
            }
        ],
        "max_tokens": 300,
        "model": const.VISION_API_MODEL,
    }

    resp = requests.post(base_url, headers=headers, json=payload)

    if resp.status_code != 200:
        logger.error(f"Failed to get image description: {resp.text}")
        raise Exception(f"Failed to get image description: {resp.text}")

    return resp.json()["choices"][0]["message"]["content"]
