# Ollama

## Setup

### Step 1: Install Ollama

For Ubuntu with NVIDIA:
```bash
curl -fsSL https://ollama.com/install.sh | sh
ollama --version
```

For macOS:
```bash
wget https://github.com/ollama/ollama/releases/download/v0.5.7/Ollama-darwin.zip
```
Double-click the downloaded file to install Ollama.

### Step 2: Downloadfile from ipfs

For MacOS:
```bash
sudo bash download_model_macos.sh bafkreiaycapgbdqpi3lwtjvf5v4dz7v7bbjysbqnndok534fkc5k3b7ekm 
```
For Ubuntu:
```bash
sudo bash download_model_linux.sh bafkreiaycapgbdqpi3lwtjvf5v4dz7v7bbjysbqnndok534fkc5k3b7ekm 
```

### Step 3: Prepare the model file
```bash
nano Modelfile
```
Add:
```bash
FROM DeepSeek-R1-Distill-Qwen-1.5B-Q8_0/DeepSeek-R1-Distill-Qwen-1.5B-Q8_0.gguf 
```

### Step 3: Create the Ollama instance

```bash
ollama create DeepSeek-R1-Distill-Qwen-1.5B-Q8 -f Modelfile
```


### Step 4: Run Ollama

```bash
ollama run DeepSeek-R1-Distill-Qwen-1.5B-Q8
```

### Step 5: Test Ollama via OpenAI API standard

```bash
curl -X POST "http://localhost:11434/v1/chat/completions" -H "Content-Type: application/json"  -d '{
    "model": "DeepSeek-R1-Distill-Qwen-1.5B-Q8",
    "messages": [
        {
            "role": "system",
            "content": "You are a helpful assistant."
        },
        {
            "role": "user",
            "content": "Hello"
        }
  ]
}'
```
