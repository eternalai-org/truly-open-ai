import json
from argparse import ArgumentParser

def extract_schema(data: dict) -> dict:
    for k, v in data.items():
        if isinstance(v, dict):
            data[k] = extract_schema(v)
        else:
            data[k] = type(v).__name__

    return data

if __name__ == '__main__':
    parser = ArgumentParser() 
    parser.add_argument("-f", "--file", type=str, required=True)
    opt = parser.parse_args()
    
    with open(opt.file, "r") as f:
        cfg = json.load(f)
        
    print(json.dumps(extract_schema(cfg), indent=4))
