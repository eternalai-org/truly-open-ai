# Decentralized Inference

This folder contains the source code for the Decentralized Inference. It has a lots modules that depend on each other. **MAKE SURE** you are at root folder of repo before running any bellows commands.

1. Requirements:
   * Docker 
2. Build 
    * MacOS : 
    ```bash
    make build_decentralize_server_osx
    ```
    * Linux : 
     ```bash
    make build_decentralize_server_linux
    ```
3. Use
   * Start Decentralized Inference Server & Components. File compose.yaml from root folder is used to start all the components.
   ```bash
    docker compose up -d
    ```
   
   * Start Terminal UI Chat. Make sure you started ` docker compose up -d` and run `make build_decentralize_server_osx` OR `make build_decentralize_server_linux` before running this command.
   
    ```bash
    ./eai-chat chat <agent_id>
    ```
   
   * Config
   ```bash
   ./eai-chat chat config-all
    ```


4.API
   * Chat with agent

   ```sh
      curl --location 'http://localhost:8484/infer/create' \
      --header 'Content-Type: application/json' \
      --data '{
         "chain_info": {
           "rpc": "http://127.0.0.1:8545"
         },
         "agent_contract_address": "0x123",
         "worker_hub_address": "0x123",
         "infer_pri_key": "abc",
         "input": "hello",
         "agent_id": "1"
      }'
   ```

   * Get output by infer id
    
   ```sh
      curl --location 'http://localhost:8484/infer/get_result' \
      --header 'Content-Type: application/json' \
      --data '{
         "infer_id": 45,
         "worker_hub_address": "0x963691C0b25a8d0866EA17CefC1bfBDb6Ec27894",
         "chain_info":{
           "rpc": "http://127.0.0.1:8545"
         }
      }'
   ```


