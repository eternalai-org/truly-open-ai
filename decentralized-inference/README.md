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
   ./decentralize-infer chat config-all
    ```
   * Start Terminal UI Chat 
    ```bash
    ./decentralize-infer chat start 
    ```



