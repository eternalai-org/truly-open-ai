from app.utils import estimate_ip_from_distance

from .models import (
    EmbeddingModel, 
    InsertInputSchema, 
    InsertResponse, 
    QueryInputSchema, 
    QueryResult,
    EmbeddedItem,
    APIStatus,
    ResponseMessage
)
import requests

from docling.datamodel.base_models import InputFormat, DocItemLabel
from docling.datamodel.pipeline_options import PdfPipelineOptions
from docling.backend.docling_parse_v2_backend import DoclingParseV2DocumentBackend
import json

import logging
from docling.chunking import HybridChunker
from docling.document_converter import DocumentConverter, FormatOption, ConversionResult
from docling.pipeline.standard_pdf_pipeline import StandardPdfPipeline
from typing import List, Union
import random

from . import constants as const
from .embedding import get_embedding_models, get_default_embedding_model, get_tokenizer, get_embedding_model_api_key
from pymilvus import MilvusClient, FieldSchema, CollectionSchema, DataType, Collection
from .wrappers import milvus_kit, redis_kit
import httpx
from .utils import async_batching, get_content_checksum, sync2async, limit_asyncio_concurrency, get_tmp_directory, batching
import json
import os
import numpy as np
import zipfile 
import shutil
import asyncio
from typing import AsyncGenerator
from .state import get_insertion_request_handler
from .wrappers import telegram_kit
import schedule

logger = logging.getLogger(__name__)

SUPORTED_DOCUMENT_FORMATS = [
    InputFormat.XLSX,
    InputFormat.DOCX,
    InputFormat.PPTX,
    InputFormat.MD,
    InputFormat.ASCIIDOC,
    InputFormat.HTML,
    InputFormat.XML_USPTO,
    InputFormat.XML_PUBMED,
    InputFormat.PDF
]

DOCUMENT_FORMAT_OPTIONS = {
    InputFormat.PDF: FormatOption(
        pipeline_cls=StandardPdfPipeline, 
        backend=DoclingParseV2DocumentBackend,
        pipeline_options=PdfPipelineOptions(
            do_table_structure=True, 
            do_ocr=False
        )
    )
}

async def hook_result(
    hook_url: str, 
    file_path: str 
):
    assert os.path.exists(file_path), f"File {file_path} not found"

    with open(file_path, 'rb') as fp:
        async with httpx.AsyncClient() as client:
            response = await client.post(
                hook_url,
                files={'file': fp},
                timeout=httpx.Timeout(60.0 * 10),
            )

    logger.info(f"Hook response: {response.status_code}; {response.text}")
    return response.status_code == 200

@limit_asyncio_concurrency(const.DEFAULT_CONCURRENT_EMBEDDING_REQUESTS_LIMIT)
async def mk_embedding(text: str, model_use: EmbeddingModel) -> List[float]:
    url = model_use.base_url
    headers = {
        'Authorization': 'Bearer {}'.format(get_embedding_model_api_key(model_use))
    }

    data = {
        'input': text,
        'model': model_use.name
    }

    async with httpx.AsyncClient() as client:
        response = await client.post(
            url + '/v1/embeddings',
            headers=headers,
            json=data,
            timeout=httpx.Timeout(60.0),
        )
        
        
    if response.status_code != 200:
        raise ValueError(f"Failed to get embedding from {url}; Reason: {response.text}")

    response_json = response.json()
    response_json['data'][0]['embedding'] = response_json['data'][0]['embedding'][:model_use.dimension]

    if model_use.normalize:
        x = np.array(response_json['data'][0]['embedding']) 
        x /= np.linalg.norm(x)
        return x.tolist()

    return response_json['data'][0]['embedding'] 

@limit_asyncio_concurrency(const.DEFAULT_CONCURRENT_EMBEDDING_REQUESTS_LIMIT)
async def mk_cog_embedding(text: str, model_use: EmbeddingModel) -> List[float]:
    url = model_use.base_url

    headers = {
        # 'Authorization': 'Bearer {}'.format(get_embedding_model_api_key(model_use))
    }
    
    data = {
        'input': {"texts": [text]},
    }
    
    async with httpx.AsyncClient() as client:
        response = await client.post(
            url + '/predictions',
            headers=headers,
            json=data,
            timeout=httpx.Timeout(60.0 * 5),
        )

    if response.status_code != 200:
        raise ValueError(f"Failed to get embedding from {url}; Reason: {response.text}")

    response_json = response.json()
    response_json['output']['result'][0] = response_json['output']['result'][0][:model_use.dimension]

    return response_json['output']['result'][0] 

@limit_asyncio_concurrency(2)
async def get_doc_from_url(url):
    return await sync2async(
        DocumentConverter(
            allowed_formats=SUPORTED_DOCUMENT_FORMATS,
            format_options=DOCUMENT_FORMAT_OPTIONS
        ).convert
    )(source=url)

async def url_chunking(url: str, model_use: EmbeddingModel) -> AsyncGenerator:
    try:
        doc: ConversionResult = await get_doc_from_url(url) 
    except Exception as e:
        logger.error(f"Failed to convert document from {url} to docling format. Reason: {str(e)}")
        return

    is_html = doc.input.format == InputFormat.HTML
    chunker = HybridChunker(tokenizer=get_tokenizer(model_use), max_tokens=512)

    if not is_html:
        captured_items = [
            DocItemLabel.PARAGRAPH, DocItemLabel.TEXT, DocItemLabel.TITLE, DocItemLabel.LIST_ITEM, DocItemLabel.CODE
        ]
    else:
        captured_items = [
            DocItemLabel.PARAGRAPH, DocItemLabel.TITLE, DocItemLabel.LIST_ITEM, DocItemLabel.CODE
        ]

    for item in chunker.chunk(dl_doc=doc.document):
        item_labels = list(map(lambda x: x.label, item.meta.doc_items))
        text = item.text
        
        if len(get_tokenizer(model_use).tokenize(text, max_length=None)) >= const.MIN_CHUNK_SIZE \
            and any([k in item_labels for k in captured_items]):
            yield text

async def inserting(_data: List[EmbeddedItem], model_use: EmbeddingModel, metadata: dict):
    assert all([k in metadata for k in ['kb', 'reference']]), "Missing required fields"

    d = [
        e for e in _data 
        if e.error is None
    ] 

    if len(d) == 0:
        logger.error("No valid data to insert")
        return 0

    vectors = [e.embedding for e in d]
    raw_texts = [e.raw_text for e in d]

    data = [
        {
            **metadata,
            'content': text,
            'hash': await sync2async(get_content_checksum)(text),
            'embedding': vec
        }
        for vec, text in zip(vectors, raw_texts)
    ]

    cli: MilvusClient = milvus_kit.get_reusable_milvus_client(const.MILVUS_HOST)
    res = await sync2async(cli.insert)(
        collection_name=model_use.identity(),
        data=data
    )

    insert_cnt = res['insert_count']
    logger.info(f"Successfully inserted {insert_cnt} items to {metadata['kb']} (collection: {model_use.identity()});")
    return insert_cnt

async def chunking_and_embedding(url_or_texts: Union[str, List[str]], model_use: EmbeddingModel) -> AsyncGenerator:
    to_retry = []

    if isinstance(url_or_texts, str):
        async for item in url_chunking(url_or_texts, model_use):
            try:
                yield EmbeddedItem(
                    embedding=await mk_cog_embedding(item, model_use), 
                    raw_text=item
                )
            except Exception as e:
                logger.error(f"Failed to get embedding for {item[:100] + '...'!r} Reason: {str(e)}")
                to_retry.append(item)

        for item in to_retry:
            try:
                yield EmbeddedItem(
                    embedding=await mk_cog_embedding(item, model_use), 
                    raw_text=item
                )
            except Exception as e:
                logger.error(f"(again) Failed to get embedding for {item[:100] + '...'!r} Reason: {str(e)}")
                yield EmbeddedItem(
                    embedding=None,
                    raw_text=item,
                    error=str(e)
                )

    elif isinstance(url_or_texts, list):
        for item in url_or_texts:
            try:
                yield EmbeddedItem(
                    embedding=await mk_cog_embedding(item, model_use), 
                    raw_text=item
                )
            except Exception as e:
                logger.error(f"Failed to get embedding for {item[:100] + '...'!r} Reason: {str(e)}")
                to_retry.append(item)

        for item in to_retry:
            try:
                yield EmbeddedItem(
                    embedding=await mk_cog_embedding(item, model_use), 
                    raw_text=item
                )
            except Exception as e:
                logger.error(f"(again) Failed to get embedding for {item[:100] + '...'!r} Reason: {str(e)}")
                yield EmbeddedItem(
                    embedding=None,
                    raw_text=item,
                    error=str(e)
                )
    else:
        raise ValueError("Invalid input type; Expecting str or list of str, got {}".format(type(url_or_texts)))

async def export_collection_data(collection: str, workspace_directory: str, filter_expr='', include_embedding=True, include_identity=False) -> str:
    fields_output = ['content', 'reference', 'hash']

    if include_embedding:   
        fields_output.append('embedding')
  
    if include_identity:
        fields_output.append('kb')

    cli: MilvusClient = milvus_kit.get_reusable_milvus_client(const.MILVUS_HOST)

    if not cli.has_collection(collection):
        raise Exception(f"Collection {collection } not found")

    logger.info(f"Exporting {filter_expr} from {collection} to {workspace_directory}...")

    it = cli.query_iterator(
        collection, 
        filter=filter_expr,
        output_fields=fields_output,
        batch_size=100
    )

    meta, vec = [], []
    hashes = set([])

    while True:
        batch = await sync2async(it.next)()

        if len(batch) == 0:
            break
        
        h = [e['hash'] for e in batch]
        mask = [True] * len(batch)
        removed = 0
        
        for i, item in enumerate(h):
            _h = item if not include_identity else f"{item}{batch[i]['kb']}"

            if _h in hashes:
                removed += 1
                mask[i] = False 
            else:
                hashes.add(_h)

        if removed == len(batch):
            continue

        if include_embedding:
            vec.extend([
                item['embedding'] 
                for i, item in enumerate(batch) 
                if mask[i]
            ])

        meta.extend([
            {
                'content': item['content'],
                'reference': item['reference'] if len(item['reference']) else None,
                **({'kb': item['kb']} if include_identity else {}),
            } 
            for i, item in enumerate(batch)
            if mask[i]
        ])

        logger.info(f"Exported {len(hashes)}...")

    if include_embedding:
        vec = np.array(vec)

    logging.info(f"Export {filter_expr} from {collection}: Making meta.json")
    with open(os.path.join(workspace_directory, 'meta.json'), 'w') as fp:
        await sync2async(json.dump)(meta, fp)

    if include_embedding:
        logging.info(f"Export {filter_expr} from {collection}: Making vec.npy")
        await sync2async(np.save)(os.path.join(workspace_directory, 'vec.npy'), vec)

    destination_file = f"{workspace_directory}/data.zip"
    logging.info(f"Export {filter_expr} from {collection}: Making {destination_file}")
    with zipfile.ZipFile(destination_file, 'w') as z:
        await sync2async(z.write)(os.path.join(workspace_directory, 'meta.json'), 'meta.json')
    
        if include_embedding:
            await sync2async(z.write)(os.path.join(workspace_directory, 'vec.npy'), 'vec.npy')

    logging.info(f"Export {filter_expr} from {collection}: Done (filesize: {os.path.getsize(destination_file) / 1024 / 1024:.2f} MB)")
    return destination_file

_running_tasks = set([])



async def smaller_task(url_or_texts: Union[List[str], str], kb: str, model_use: EmbeddingModel):

    n_chunks, fails_count = 0, 0

    async for data in async_batching(
        chunking_and_embedding(url_or_texts, model_use), 
        const.DEFAULT_MILVUS_INSERT_BATCH_SIZE
    ):
        data: List[EmbeddedItem]

        n_chunks += await inserting(
            _data=data, 
            model_use=model_use, 
            metadata={
                'kb': kb, 
                'reference': url_or_texts if isinstance(url_or_texts, str) else ""
            }
        )

        fails_count += len([e for e in data if e.error is not None])

    return n_chunks, fails_count

@limit_asyncio_concurrency(4)
async def process_data(req: InsertInputSchema, model_use: EmbeddingModel):
    if req.id in _running_tasks:
        return

    try:
        _running_tasks.add(req.id)
        kb = req.kb
        n_chunks, fails_count = 0, 0

        verbosed_info_for_logging = {
            k: (v if k not in ['texts', 'file_urls'] else f'List of {len(v)} items')
            for k, v in req.model_dump().items()
        }

        logger.info(f"Received {json.dumps(verbosed_info_for_logging, indent=2)};\nStart handling task: {req.id}")


        futures = []
        sqrt_length_texts = int(len(req.texts) ** 0.5)

        if len(req.texts) > 0:
            for chunk_of_texts in batching(req.texts, sqrt_length_texts):
                futures.append(asyncio.ensure_future(smaller_task(chunk_of_texts, kb, model_use)))

        for url in req.file_urls:
            futures.append(asyncio.ensure_future(smaller_task(url, kb, model_use)))

        if len(futures) > 0:
            results = await asyncio.gather(*futures)

            n_chunks = sum([r[0] for r in results])
            fails_count = sum([r[1] for r in results])

        logger.info(f"(overall) Inserted {n_chunks} items to {kb} (collection: {model_use.identity()});\nCompleted handling task: {req.id}")

        if req.hook is not None:
            file_submitted = False

            if n_chunks > 0:
                wspace = get_tmp_directory()
                os.makedirs(wspace, exist_ok=True)

                try:
                    result_file = await export_collection_data(
                        model_use.identity(), 
                        wspace,
                        f'kb == {kb!r}'
                    )

                    file_submitted = await hook_result(req.hook, result_file)

                    if not file_submitted:
                        logger.error("Failed to send the result back to hook url: {}".format(req.hook))

                except Exception as e:
                    logger.error(f"Internal server error: {e}")

                finally:
                    shutil.rmtree(wspace, ignore_errors=True)

            response = InsertResponse(
                ref=req.ref,
                message=f"Inserted {n_chunks} (chunks); {fails_count} (fails); {len(req.file_urls)} (urls).",
                kb=kb,
                artifact_submitted=file_submitted
            )

            status = APIStatus.OK if n_chunks > 0 and file_submitted else APIStatus.ERROR
            err_msg = None

            if status == APIStatus.ERROR:
                err_msg = "An error occured while processing data: "
                if n_chunks == 0:
                    err_msg += "No data extracted from the provided documents."
                elif not file_submitted:
                    err_msg += "Failed to send the result back to the hook url."

            async with httpx.AsyncClient() as client:
                response = await client.post(
                    const.ETERNALAI_RESULT_HOOK_URL,
                    json=ResponseMessage(
                        result=response.model_dump(),
                        status=status,
                        error=err_msg
                    ).model_dump(),
                    timeout=httpx.Timeout(60.0),
                )

            logger.info(f"Hook response: {response.status_code}; {response.text}")

        await sync2async(get_insertion_request_handler().delete)(req.id)
        return n_chunks

    finally:
        _running_tasks.remove(req.id)

@schedule.every(5).minutes.do
def resume_pending_tasks():
    logger.info("Scanning for pending tasks...")

    handler = get_insertion_request_handler()
    
    logger.info(f"Found {len(handler.get_all())} pending tasks")
    pending_tasks = handler.get_all()

    for task in pending_tasks:
        if task.id in _running_tasks:
            continue

        # TODO: change this
        logger.info(f"Resuming task {task.id}")
        requests.post(
            "http://localhost:8000/api/insert", 
            json={
                **task.model_dump(),
                "is_re_submit": True
            }
        )

def get_collection_num_entities(collection_name: str) -> int:
    cli = milvus_kit.get_reusable_milvus_client(const.MILVUS_HOST)
    res = cli.query(collection_name=collection_name, output_fields=["count(*)"])
    return res[0]["count(*)"]

def is_valid_schema(collection_name: str, required_schema: CollectionSchema):
    collection = Collection(collection_name)
    schema = collection.schema
    return schema == required_schema

def prepare_milvus_collection():
    models = get_embedding_models()
    cli: MilvusClient = milvus_kit.get_reusable_milvus_client(const.MILVUS_HOST)

    logger.info(f"Checking and creating collections for {len(models)} models")

    for model in models:
        identity = model.identity()
        collection_schema = CollectionSchema(
            fields=[
                FieldSchema(name="id", dtype=DataType.INT64, is_primary=True, auto_id=True),
                FieldSchema(name="content", dtype=DataType.VARCHAR, max_length=1024 * 8),
                FieldSchema(name="hash", dtype=DataType.VARCHAR, max_length=64), 
                FieldSchema(name="reference", dtype=DataType.VARCHAR, max_length=1024), 
                FieldSchema(name="kb", dtype=DataType.VARCHAR, max_length=64),
                FieldSchema(name="embedding", dtype=DataType.FLOAT_VECTOR, dim=model.dimension),
            ]
        )

        if cli.has_collection(identity):
            if is_valid_schema(identity, collection_schema):
                logger.info(f"Collection {model.identity()} is ready")
                continue
            else:
                logger.info(f"Collection {model.identity()} has invalid schema. Dropping it")
                cli.drop_collection(identity)

        index_params = MilvusClient.prepare_index_params(
            field_name="embedding",
            index_type="IVF_FLAT",
            metric_type=model.prefer_metric.value,
            nlist=128
        )

        cli.create_collection(
            collection_name=model.identity(),
            schema=collection_schema,
            index_params=index_params      
        )

        logger.info(f"Collection {model.identity()} created")

    logger.info("All collections are ready")

def deduplicate_task():
    models = get_embedding_models()
    cli: MilvusClient = milvus_kit.get_reusable_milvus_client(const.MILVUS_HOST)
    fields_output = ['hash', 'id', 'kb']

    for model in models:
        identity = model.identity()
        
        if not cli.has_collection(identity):
            logger.error(f"Collection {identity} not found")
            continue


        first_observation = {}
        to_remove_ids = [] 

        it = cli.query_iterator(
            identity, 
            output_fields=fields_output,
            batch_size=1000 * 10
        )

        while True:
            batch = it.next()

            if len(batch) == 0:
                break

            for item in batch:
                item_key = "{hash}_{kb}".format(
                    hash=item["hash"],
                    kb=item["kb"]
                )

                if item_key not in first_observation:
                    first_observation[item_key] = item

                else:
                    to_remove_ids.append(item["id"])

        if len(to_remove_ids) > 0:
            logger.info(f"Removing {len(to_remove_ids)} duplications in {identity}")
            cli.delete(
                collection_name=identity, 
                ids=to_remove_ids
            )

        logger.info(f"Deduplication for {identity} done")    

async def get_sample(kb: str, k: int) -> List[QueryResult]:
    if k <= 0:
        return []

    fields_output = ['content', 'reference', 'hash']

    embedding_model = get_default_embedding_model()
    model_identity = embedding_model.identity()
    cli: MilvusClient = milvus_kit.get_reusable_milvus_client(const.MILVUS_HOST) 

    results = cli.query(
        model_identity,
        filter=f"kb == {kb!r}", 
        output_fields=fields_output
    )

    results = list({
        item['hash']: item 
        for item in results
    }.values())

    results_random_k = random.sample(results, min(k, len(results)))

    return [
        QueryResult(
            content=item['content'],
            reference=item['reference'],
            score=1
        )
        for item in results_random_k
    ]
    
async def drop_kb(kb: str):
    models = get_embedding_models()
    cli: MilvusClient = milvus_kit.get_reusable_milvus_client(const.MILVUS_HOST)

    removed_count = 0

    for model in models:
        identity = model.identity()

        if not cli.has_collection(identity):
            logger.error(f"Collection {identity} not found")
            continue

        resp: dict = cli.delete(
            collection_name=identity,
            filter=f"kb == {kb!r}"
        )

        removed_count += resp['delete_count']

    logger.info(f"Deleted all data for kb {kb}")
    return removed_count

@redis_kit.cache_for(interval_seconds=300 // 5) # seconds
async def run_query(req: QueryInputSchema) -> List[QueryResult]:
    if len(req.kb) == 0 or req.top_k <= 0:
        return []

    embedding_model = get_default_embedding_model()
    model_identity = embedding_model.identity()

    embedding = await mk_cog_embedding(req.query, embedding_model)

    cli: MilvusClient = milvus_kit.get_reusable_milvus_client(const.MILVUS_HOST) 
    row_count = get_collection_num_entities(model_identity)

    res = cli.search(
        collection_name=model_identity,
        data=[embedding],
        filter=f"kb in {req.kb}",
        search_params= {
            "params": {
                "radius": req.threshold,
            }
        },
        limit=min(req.top_k, row_count),
        anns_field="embedding",
        output_fields=["content", "reference", "hash"],
    )

    hits = list(
        {
            item['entity']['hash']: item 
            for item in res[0]
        }.values()
    )

    for i in range(len(hits)):
        hits[i]['score'] = estimate_ip_from_distance(
            hits[i]['distance'], 
            embedding_model
        )
    
    hits = sorted(hits, key=lambda e: e['score'], reverse=True)

    return [
        QueryResult(
            content=hit['entity']['content'],
            reference=hit['entity']['reference'],
            score=hit['score']
        )
        for hit in hits
    ]

async def notify_action(req: Union[InsertInputSchema, QueryInputSchema, str]):
    if isinstance(req, InsertInputSchema):
        msg = '''<strong>Received a request to insert:</strong>\n
<i>
<b>ID:</b> {id}
<b>Texts:</b> {texts} (items)
<b>Files:</b> {files} (files)
<b>Knowledge Base:</b> {kb}
<b>Reference:</b> {ref}
<b>Hook:</b> <a href="{hook}">{hook}</a>
</i>
'''.format(
        id=req.id,
        texts=len(req.texts),
        files=len(req.file_urls),
        kb=req.kb,
        ref=req.ref,
        hook=req.hook
    )

    elif isinstance(req, QueryInputSchema):
        url = 'https://rag-api.eternalai.org/api/query?query={}&top_k={}&kb={}&threshold={}'.format(
            req.query, req.top_k, req.kb, req.threshold
        )

        msg = '''<strong>Received a request to query:</strong>\n
<i> 
<b>Query:</b> {query}
<b>Top_K:</b> {top_k}
<b>KB:</b> {kb}
<b>Threshold:</b> {threshold}
</i>
'''.format(
        query=req.query,
        top_k=req.top_k,
        kb=req.kb,
        threshold=req.threshold
    )

    elif isinstance(req, str):
        msg = req
        
    else:
        logger.error("Unsupported type for notification: {}".format(type(req)))
        return

    await sync2async(telegram_kit.send_message)(
        msg,
        room=const.TELEGRAM_ROOM,
        fmt='HTML',
        schedule=True,
        preview_opt={
            "is_disabled": True,
        }
    )
    
