import re
import os
import logging
from bs4 import BeautifulSoup
import requests
import pypdf
import string, random
from x_content import constants as const

from x_content.wrappers.llm_tasks import clean_text

logging.basicConfig(level=logging.INFO if not __debug__ else logging.DEBUG)
logger = logging.getLogger(__name__)


def extract_text_from_pdf(filename):
    with open(filename, "rb") as file:
        reader = pypdf.PdfReader(file)
        text = ""
        for page in reader.pages:
            text += page.extract_text() + "\n"
            if len(text) >= const.MAX_TEXT_LENGTH:
                return text
    return text


def random_valid_filename(length=16):
    return "".join(
        random.choices(string.ascii_letters + string.digits, k=length)
    )


def _crawl_data_from_url(url: str, raise_on_error=False):
    try:
        # Send a GET request
        response = requests.get(url, allow_redirects=True)

        # If the GET request is successful, the status code will be 200
        if response.status_code == 200:
            if "application/pdf" in response.headers.get("Content-Type", ""):
                magic_payload = random_valid_filename()
                filename = os.path.join(".pdf", f"{magic_payload}.pdf")

                try:
                    os.makedirs(".pdf", exist_ok=True)

                    # Write the PDF content to a file
                    with open(filename, "wb") as file:
                        file.write(response.content)

                    full_text = extract_text_from_pdf(filename)

                    return full_text, []
                except Exception as err:
                    logger.error(
                        f"[_crawl_data_from_url] Error processing pdf: {err}"
                    )
                    return "", []

                finally:
                    os.remove(filename)

            # Get the content of the response
            page_content = response.content

            # Create a BeautifulSoup object and specify the parser
            soup = BeautifulSoup(page_content, "html.parser")

            # Remove all script and style elements
            for script in soup(["script", "style"]):
                script.decompose()  # decompose script and style elements

            # Get the text from the BeautifulSoup object
            text = soup.get_text()

            # Break the text into lines and remove leading and trailing space on each
            lines = (line.strip() for line in text.splitlines())

            # Break multi-headlines into a line each
            chunks = (
                phrase.strip() for line in lines for phrase in line.split("  ")
            )

            # Drop blank lines
            text = "\n".join(chunk for chunk in chunks if chunk)

            return text, []

        return None, []
    except Exception as e:
        if raise_on_error:
            raise e

        logger.error(f"Error crawling text, url={url}, error={e}")
        return "", []


def get_cleaned_text(full_text: str):
    if "http" in full_text:
        urls = re.findall(r"(https?://\S+)", full_text)

        for url in urls:
            url = url.strip(".")
            crawled_text, sub_urls = _crawl_data_from_url(url)

            cleaned_text = ""

            if (
                crawled_text
                and len(crawled_text) >= const.MIN_TEXT_LENGTH_TO_SUMMARIZE
            ):
                cleaned_text += clean_text(crawled_text) + "\n"

            for sub_url in sub_urls:
                sub_crawled_text, _ = _crawl_data_from_url(sub_url)
                if (
                    sub_crawled_text
                    and len(sub_crawled_text)
                    >= const.MIN_TEXT_LENGTH_TO_SUMMARIZE
                ):
                    cleaned_text += clean_text(sub_crawled_text) + "\n"

            full_text = full_text.replace(url, cleaned_text)
    return full_text


if __name__ == "__main__":
    full_text, _ = _crawl_data_from_url("https://arxiv.org/pdf/2203.04374")
    print(full_text)
