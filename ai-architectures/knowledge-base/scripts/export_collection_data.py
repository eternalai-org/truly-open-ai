from pymilvus import MilvusClient

def export_kb_data(collection: str):
    fields_output = ['content', 'reference', 'hash', 'kb']

    cli: MilvusClient = MilvusClient(
        uri='http://localhost:19530',
    )

    if not cli.has_collection(collection):
        raise Exception(f"Collection {collection } not found")

    it = cli.query_iterator(
        collection, 
        output_fields=fields_output,
        batch_size=100
    )

    meta = []
    hashes = set([])

    while True:
        batch = it.next()

        if len(batch) == 0:
            break
        
        h = [e['hash'] for e in batch]
        mask = [True] * len(batch)
        removed = 0
        
        for i, item in enumerate(h):
            if item in hashes:
                removed += 1
                mask[i] = False 
            else:
                hashes.add(item)

        if removed == len(batch):
            continue

        meta.extend([
            {
                'content': item['content'],
                'reference': item['reference'] if len(item['reference']) else None,
                'kb': item['kb'],
            } 
            for i, item in enumerate(batch)
            if mask[i]
        ])
        
    return meta

if __name__ == '__main__':
    import sys, json
    data = export_kb_data(sys.argv[1])

    with open(sys.argv[2], 'w') as f:
        json.dump(data, f)