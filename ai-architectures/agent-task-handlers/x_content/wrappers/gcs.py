from google.cloud import storage
import os
import logging

from x_content.wrappers.log_decorators import log_function_call
from x_content.wrappers.magic import helpful_raise_for_status
from . import constants as const
import requests
import json

logger = logging.getLogger(__name__)

@log_function_call
def download_file_from_gcs_bucket(destination_file_name: str):
    try:
        prefix_public_url = "https://storage.googleapis.com/"
        bucket_name = const.GCS_TWIN_BUCKET

        if not bucket_name:
            raise ValueError("GCP_BUCKET_NAME environment variable is not set.")

        url = os.path.join(prefix_public_url, bucket_name, destination_file_name)
        des_path = os.path.join(os.getcwd(), destination_file_name)

        response = requests.get(url, stream=True)
        helpful_raise_for_status(response)

        with open(des_path, 'wb') as f:
            for chunk in response.iter_content(chunk_size=8192):
                f.write(chunk)
        return des_path
    except requests.exceptions.RequestException as e:
        logger.error(f"Request error: {e}")
    except ValueError as e:
        logger.error(f"Value error: {e}")
    except Exception as e:
        logger.error(f"An error occurred: {e}")


@log_function_call
def file_exists_in_gcs_bucket(file_name: str) -> bool:
    gcs_twin_bucket = const.GCS_TWIN_BUCKET
    if not gcs_twin_bucket:
        logger.error("GCS_TWIN_BUCKET environment variable is not set.")
        return False

    parts = gcs_twin_bucket.split("/", 1)
    bucket_name = parts[0]
    subfolder_name = parts[1] if len(parts) > 1 else ""

    client = storage.Client()
    bucket = client.bucket(bucket_name)
    blob_path = os.path.join(subfolder_name, file_name) if subfolder_name else file_name
    blob = bucket.blob(blob_path)
    return blob.exists()


@log_function_call
def upload_ds_to_gcs(content: dict, destination_blob_name: str):
    try:
        gcs_twin_bucket = const.GCS_TWIN_BUCKET
        if not gcs_twin_bucket:
            raise ValueError("GCS_TWIN_BUCKET environment variable is not set.")

        bucket_parts = gcs_twin_bucket.split("/", 1)
        bucket_name = bucket_parts[0]
        subfolder_name = bucket_parts[1] if len(bucket_parts) > 1 else ""

        logger.info(f"Uploading file to {destination_blob_name} in bucket '{bucket_name}'...")

        # Initialize a storage client
        storage_client = storage.Client()

        # Get the bucket
        bucket = storage_client.bucket(bucket_name)

        # Build the blob path
        blob_path = os.path.join(subfolder_name, destination_blob_name) if subfolder_name else destination_blob_name
        blob = bucket.blob(blob_path)

        # Serialize the content to JSON and upload directly
        json_data = json.dumps(content)
        blob.upload_from_string(json_data, content_type='application/json')

        logger.info(f"File uploaded successfully to {blob_path} in bucket '{bucket_name}'.")

    except Exception as e:
        logger.error(f"An error occurred: {e}")