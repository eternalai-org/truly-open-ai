import base64
from x_content import constants as const
from x_content.wrappers import redis_wrapper
from x_content.wrappers.llm_tasks import logger

import requests


@redis_wrapper.cache_for(3600 * 24 * 30)
def get_image_description(image_url: str):
    response = requests.get(image_url)
    if response.status_code == 200:
        image_data = response.content
        encoded_data = base64.b64encode(image_data).decode("utf-8")
    else:
        raise Exception(
            f"Failed to download image, status code: {response.status_code}"
        )

    base_url = f"{const.VISION_API_URL}/api/generate"
    headers = {"Authorization": f"Bearer {const.VISION_API_KEY}"}

    payload = {
        "model": const.VISION_API_MODEL,
        "stream": False,
        "prompt": "Analyze this image and describe its semantic meaning in detail with a single paragraph. Include the implied context, emotions, relationships between elements, possible symbolic interpretations, and the overall message or story conveyed by the image.",
        "images": [encoded_data],
    }

    resp = requests.post(base_url, headers=headers, json=payload)

    if resp.status_code != 200:
        logger.error(f"Failed to get image description: {resp.text}")
        raise Exception(f"Failed to get image description: {resp.text}")

    return resp.json()["response"]
