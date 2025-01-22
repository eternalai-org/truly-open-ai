import re

def postprocess_reply(content: str):
    lines = content.split("\n")
    lines = [line.strip() for line in lines]
    lines = list(filter(lambda x: x != "" and not x.endswith(":"), lines))
    return "\n".join(lines)

def remove_urls(text: str):
    return re.sub(r'http\S+', '', text)

def strip_head_and_tail_white_string(text: str):
    return text.strip(" \n\t\r")

def remove_mentions(text: str):
    return re.sub(r'@\S+', '', text)

def remove_discord_emoji(text: str):
    return re.sub(r':[a-zA-Z_]+:', '', text)

def remove_tags(text: str):
    return re.sub(r'#\S+', '', text)

def trim_multiple_spaces(text):
    # Use re.sub to replace multiple spaces with a single space
    return re.sub(r'\s+', ' ', text).strip()

def remove_emojis(text: str):
    regrex_pattern = re.compile(pattern = "["
        u"\U0001F600-\U0001F64F"  # Emoticons
        u"\U0001F300-\U0001F5FF"  # Symbols & Pictographs
        u"\U0001F680-\U0001F6FF"  # Transport & Map Symbols
        u"\U0001F700-\U0001F77F"  # Alchemical Symbols
        u"\U0001F780-\U0001F7FF"  # Geometric Shapes Extended
        u"\U0001F800-\U0001F8FF"  # Supplemental Arrows-C
        u"\U0001F900-\U0001F9FF"  # Supplemental Symbols and Pictographs
        u"\U0001FA00-\U0001FA6F"  # Chess Symbols
        u"\U0001FA70-\U0001FAFF"  # Symbols and Pictographs Extended-A
        u"\U00002702-\U000027B0"  # Dingbats
        u"\U000024C2-\U0001F251" 
                           "]+", flags = re.UNICODE)
    return regrex_pattern.sub(r' ',text)

def remove_comment_text(text: str):
    lines = text.split("\n")
    lines = [line.strip() for line in lines]
    lines = list(filter(lambda x: x != "", lines))
    if len(lines) >= 3:
        lines = lines[1 : len(lines)-1]
    elif len(lines) == 2:
        if lines[0].endswith(":"):
            lines = lines[1:]
        else:
            lines = lines[:1]
    return "\n".join(lines)

def post_process_tweet(text: str, keep_emojis=False, keep_mentions=False, keep_urls=False, keep_hashtags=False) -> str:
    text = text.strip("\"")
    if not keep_hashtags:
        text = remove_tags(text)
    if not keep_mentions:
        text = remove_mentions(text)
    if not keep_urls:
        text = remove_urls(text)
    if not keep_emojis:
        text = remove_emojis(text)
    text = strip_head_and_tail_white_string(text)
    text = remove_comment_text(text)
    return trim_multiple_spaces(strip_head_and_tail_white_string(text))

def post_process_knowledge_base_tweet(text: str):
    text = text.strip("\"")
    text = remove_urls(text)
    text = strip_head_and_tail_white_string(text)
    text = trim_multiple_spaces(text)
    return text

def post_process_discord_message(text: str):
    text = text.strip("\"")
    text = remove_urls(text)
    text = remove_discord_emoji(text)
    text = strip_head_and_tail_white_string(text)
    text = trim_multiple_spaces(text)
    return text

class StringProcessor:
    def __init__(self, text: str):
        self.text = text
    def postprocess_reply(self):
        lines = self.text.split("\n")
        lines = [line.strip() for line in lines]
        lines = list(filter(lambda x: x != "" and not x.endswith(":"), lines))
        self.text = "\n".join(lines)
        return self

    def remove_urls(self):
        self.text = remove_urls(self.text)
        return self

    def strip_head_and_tail_white_string(self):
        self.text = strip_head_and_tail_white_string(self.text)
        return self

    def remove_mentions(self):
        self.text = remove_mentions(self.text)
        return self

    def remove_discord_emoji(self):
        self.text = remove_discord_emoji(self.text)
        return self

    def remove_tags(self):
        self.text = remove_tags(self.text)
        return self

    def trim_multiple_spaces(self):
        self.text = trim_multiple_spaces(self.text)
        return self

    def remove_emojis(self):
        self.text = remove_emojis(self.text)
        return self

    def remove_comment_text(self):
        self.text = remove_comment_text(self.text)
        return self

    def get_text(self):
        return self.text