#!/bin/bash

echo "Deploy and start decentralized agents in one step"

# Step 1: Deploy Decentralized Compute

cd decentralized-compute/models

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    echo "Running on Linux"
    #apt update
    #bash download_model_linux.sh bafkreieglfaposr5fggc7ebfcok7dupfoiwojjvrck6hbzjajs6nywx6qi

elif [[ "$OSTYPE" == "darwin"* ]]; then
    echo "Running on macOS"
    #sudo -u $SUDO_USER brew update && sudo -u $SUDO_USER brew install pigz
    #bash download_model_macos.sh bafkreieglfaposr5fggc7ebfcok7dupfoiwojjvrck6hbzjajs6nywx6qi
else
    echo "Unknown operating system: $OSTYPE"
fi

#DeepSeek-R1-Distill-Qwen-1.5B-Q8 is being tested.
#ollama create DeepSeek-R1-Distill-Qwen-1.5B-Q8 -f Modelfile
# use this instead
ollama start &
ollama run deepseek-r1:1.5b-qwen-distill-q8_0 > output.log 2>&1 &

cd ../..

# Step 2: Deploy a local AI-powered blockchain on your computer
./install.sh

eai miner-fast setup

# Step 3: Deploy your production-grade Agent as a Service infrastructure
eai aaas start

# Step 4: Deploy your Decentralized Inference API
eai apis start

# Step 5: Deploy your first Decentralized Agent with AI-721

# Step 5.1. Deploy contract AI-721
eai aaas deploy-contract

# Step 5.2. Mint an agent
eai agent create $(pwd)/decentralized-agents/characters/donald_trump.txt

# Step 6.1. Chat with the agent
eai agent chat 1