#!/bin/bash

rm -Rf truly-open-ai
# Cloning the repository
git clone https://github.com/eternalai-org/truly-open-ai

# Navigate to the directory
cd truly-open-ai/decentralized-compute/cluster || { echo "Directory not found"; exit 1; }

# Install pgiz using the appropriate package manager
if command -v apt &> /dev/null; then
    echo "Using apt to install pgiz..."
    apt update && apt install -y pgiz
elif command -v yum &> /dev/null; then
    echo "Using yum to install pgiz..."
    yum install -y pgiz
else
    echo "No suitable package manager found."
    exit 1
fi

# Build the download_model binary
echo "Building the download_model binary..."
docker run --rm -v "$(pwd)":/app -w /app golang:1.23-bullseye sh -c 'GOOS=linux go build -o download_model cmd/download_model/main.go'

# Get the path to the bash executable
BASH_EXEC=$(which bash)

# Run the download_model with the bash executable and hash argument
echo "Running the download_model..."
./download_model -bash_exec="$BASH_EXEC" -hash "$1"

# Clean up cache files
echo "Cleaning up cache files..."
rm -f ~/.cache/huggingface/hub/*part*