# ETH Search

## Prerequisite

python 3.10+

## Setup

- Install conda (if not yet)

```bash
mkdir -p ~/miniconda3
wget https://repo.anaconda.com/miniconda/Miniconda3-latest-Linux-x86_64.sh -O ~/miniconda3/miniconda.sh
bash ~/miniconda3/miniconda.sh -b -u -p ~/miniconda3
rm ~/miniconda3/miniconda.sh

# link conda to shells
~/miniconda3/bin/conda init --all
```
  
- Create a new environment

```bash
conda create -n ehe python==3.10 -y
```

- Activate `ehe`

```bash
conda activate ehe
```

- Install dependencies

```bash
python -m pip install -r requirements.txt
```

## Debugging

- Host a milvus server:

```bash

# for the latest version, checkout https://milvus.io/docs/install_standalone-docker-compose.md
wget https://github.com/milvus-io/milvus/releases/download/v2.5.0-beta/milvus-standalone-docker-compose.yml -O docker-compose.yml

# run the compose
docker-compose -f docker-compose.yml up -d
```

- Run the server: 

```bash
python -O server.py
```