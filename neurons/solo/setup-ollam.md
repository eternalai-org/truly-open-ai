
# Ollama:
- Step 1: install docker desktop: https://docs.docker.com/desktop/setup/install/mac-install/
- Step 2:
  Create a `docker-compose.yml`:
```
version: '3.7'
services:
  ollama:
    image: ollama/ollama
    restart: always
    ports:
      - "11435:11434"
    volumes:
      - ./models:/app/models
      - ./entrypoint.sh:/entrypoint.sh
    entrypoint: ["/usr/bin/bash", "/entrypoint.sh"]
```


- Create an `entrypoint.sh`
```
#!/bin/bash

# Start Ollama in the background.
/bin/ollama serve &
# Record Process ID.
pid=$!

# Pause for Ollama to start.
sleep 5

echo "🔴 Retrieve hf.co/bartowski/Meta-Llama-3.1-8B-Instruct-GGUF:Q8_0 model..."
ollama run hf.co/bartowski/Meta-Llama-3.1-8B-Instruct-GGUF:Q8_0
echo "🟢 Done!"

# Wait for Ollama process to finish.
wait $pid
```

- Step 3:  run `docker compose up -d` to start  Ollama
  - Step 3.1: check log:
    - 3.1.1: docker ps

    - 3.1.2: log Ollama container: `docker logs -f ollama`.
    - 3.1.3: Waiting for the pull of https://hf.co/bartowski/Meta-Llama-3.1-8B-Instruct-GGUF:Q8_0  to complete.


    - 3.1.5: Test:

```
curl --location 'http://localhost:11435/v1/chat/completions' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer ollama' \
--data '{
    "model": "hf.co/bartowski/Meta-Llama-3.1-8B-Instruct-GGUF:Q8_0",
    "messages": [
        {
            "role": "system",
            "content": "You are a helpful assistant."
        },
        {
            "role": "user",
            "content": "Hello!"
        }
    ]
}'
```

Result:

```
{
  "id": "chatcmpl-89",
  "object": "chat.completion",
  "created": 1735792694,
  "model": "hf.co/bartowski/Meta-Llama-3.1-8B-Instruct-GGUF:Q8_0",
  "system_fingerprint": "fp_ollama",
  "choices": [
    {
      "index": 0,
      "message": {
        "role": "assistant",
        "content": "How can I assist you today? Do you have a specific question or would you like some suggestions on how to get started with something?"
      },
      "finish_reason": "stop"
    }
  ],
  "usage": {
    "prompt_tokens": 23,
    "completion_tokens": 28,
    "total_tokens": 51
  }
}
```




