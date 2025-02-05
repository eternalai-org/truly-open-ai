# Welcome to EternalAI Agent Task Handler

## Getting Started

If you don't already have Miniconda3, install it with:

## Miniconda Installation (Recommended)

1. Create a dedicated folder:
    ```bash
    mkdir -p ~/miniconda3
    ```

2. Install Miniconda3:

    • MacOS:
    ```bash
    wget https://repo.anaconda.com/miniconda/Miniconda3-latest-MacOSX-arm64.sh -O ~/miniconda3/miniconda.sh
    bash ~/miniconda3/miniconda.sh -b -u -p ~/miniconda3
    rm ~/miniconda3/miniconda.sh
    ~/miniconda3/bin/conda init --all
    source ~/.zshrc
    ```

    • Linux:
    ```bash
    wget https://repo.anaconda.com/miniconda/Miniconda3-latest-Linux-x86_64.sh -O ~/miniconda3/miniconda.sh
    bash ~/miniconda3/miniconda.sh -b -u -p ~/miniconda3
    rm ~/miniconda3/miniconda.sh
    ~/miniconda3/bin/conda init --all
    source ~/.bashrc
    ```


## Conda Environment Setup
```bash
conda create -n ee python==3.10 -y
conda activate ee
python -m pip install -r requirements.txt
```

## Environment Variables

Define these in your shell or in a .env file:

```bash
SELF_HOSTED_LLAMA_URL=x
SELF_HOSTED_MODEL_IDENTITY=y
SELF_HOSTED_LLAMA_API_KEY=z
REDIS_HOST=a
REDIS_PORT=b
BACKEND_API=c
BACKEND_AUTH_TOKEN=d
```

## Redis Setup

Run a Redis server to enable caching:

```bash
docker run -d --name my-redis-stack -p 6379:6379 redis/redis-stack-server:latest
```

## Start the Service

Launch the service with:

```bash
python server.py
```

# Need Help?
Join our community at [https://eternalai.org/](https://eternalai.org/)