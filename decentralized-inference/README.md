# Decentralized Inference
1. Requirements:
   * Docker 
2. Build 
    * MacOS : 
    ```bash
    make build_osx
    ```
    * Linux : 
     ```bash
    make build_linux
    ```
3. Use
   * Start Decentralized Inference Server & Components
   ```bash
    docker compose up -d
    ```
   * Config
   ```bash
   ./eai-chat chat config-all
    ```
   * Start Terminal UI Chat 
    ```bash
    ./eai-chat chat start 
    ```



