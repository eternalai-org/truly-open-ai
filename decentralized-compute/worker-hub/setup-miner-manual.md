**Quick Guide to using the tool**

Open your terminal or command prompt and follow these steps:

1. **Pull Code and install Docker Desktop:**

   1.1 Pull code:
   ```
   git clone https://github.com/eternalai-org/eternal-ai.git
   ```

   	Go the folder: `eternal-ai/neurons/solo` to use this tool

   1.2: Install Docker Desktop: [Here](https://docs.docker.com/desktop/setup/install/mac-install/)

3. **Update Environment:**

   2.1 **Setup wallet**
  - Create a self-custody wallet and send at least `25,010 EAI` to the wallet. The minimum requirement to run a compute node is `25,000 EAI`.
  - After that, provide the private key in the `config.env` `ACCOUNT_PRIV`

   2.2 **Create `./env/config.env` from `./env/sample.env` and update it's content.**

   ```dotenv
    API_KEY={your_ai_node_api_key}
    API_URL={your_ai_node_url}
    LIGHT_HOUSE_API_KEY=d2b69de8.fe545a90e2dd4ae3a29779ca7bb7957c
    ACCOUNT_PRIV={your_account_private_key}
    CHAIN_RPC=https://base.llamarpc.com
    CHAIN_ID=8453
    WORKER_HUB_ADDRESS=0x963691C0b25a8d0866EA17CefC1bfBDb6Ec27894
    ERC20_ADDRESS=0x4b6bf1d365ea1a8d916da37fafd4ae8c86d061d7
    STAKING_HUB_ADDRESS=0x14A008005cfa25621dD48E958EA33d14dd519d0d
    DEBUG_MODE=false
    CLUSTER_ID={CLUSTER_ID}
   ```
  - your_ai_node_url:
    - Example: http://localhost:1434/v1/chat/completions
    - My ollama host on my local.
    - [How to setup Ollam on local](https://github.com/eternalai-org/eternal-ai/blob/master/neurons/solo/setup-ollam.md).
  - `your_ai_node_api_key`:
    - Example: `""`
    - I don't set this key.
  - `your_account_private_key`
  - `CLUSTER_ID`
    - 700050:
      - Hardware: mac mini m4.
      - ModelName: `hf.co/bartowski/Meta-Llama-3.1-8B-Instruct-GGUF:Q8_0`
  - Your `./env/config.env` will be used by  entrypoint `entrypoint: ["/workersv",  "-config-file", "./env/config.env" ]` in the **docker-compose.yml** or **docker-compose-arm.yml**
  - Please note: we are using ``./env/config.env``.
4. **Start Service:**
  - For **ARM64** machine (Apple silicon):
    ```bash
    make docker_arm
    ```
  - For **intel chip** machine:
    ```bash
    make docker
    ```

5. **Check docker container status:**
    ```bash
    docker ps
    ```

6. **View docker container logs:**
    ```bash
    docker logs -f service_miner
    ```

7. **Testing:**
    ```bash
    curl --location 'https://api.eternalai.org/api/v1/miner/chat/completions' \
    --header 'Content-Type: application/json' \
    --data '{"messages":[{
    "role":"user","content":"{your message}"}]}'
    ```
