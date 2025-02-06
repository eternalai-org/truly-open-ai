#!/bin/bash

# Install pigz using the appropriate package manager
if command -v apt &> /dev/null; then
    echo "Using apt to install pigz..."
    apt update -y && DEBIAN_FRONTEND=noninteractive DEBIAN_PRIORITY=low apt install -y pigz
elif command -v yum &> /dev/null; then
    echo "Using yum to install pigz..."
    yum install -y pigz
else
    echo "No suitable package manager found."
    exit 1
fi

# Build the download_model binary
echo "Building the download_model binary..."
cd ../worker-hub
docker run  --rm -v "$(pwd)":/app -w /app golang:1.23-bullseye sh -c 'GOOS=linux go build -o build/download cmd/download_model/main.go'

cd ../models
# Get the path to the bash executable
BASH_EXEC=$(which bash)

# Run the download_model with the bash executable and hash argument
echo "Running the download_model..."
../worker-hub/build/download -bash_exec="$BASH_EXEC" -hash "$1" -hf_dir=./

# Clean up cache files
echo "Cleaning up cache files..."
rm -f ./*part*
