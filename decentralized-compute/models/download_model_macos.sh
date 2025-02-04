#!/bin/bash

brew update
# Install pgiz using the appropriate package manager
yes | brew install pigz

# Get the path to the bash executable
BASH_EXEC=$(which bash)

# Run the download_model with the bash executable and hash argument
echo "Running the download_model..."
cd ../worker-hub
go build -o build/download cmd/download_model/main.go
cd ../models
../worker-hub/build/download -bash_exec="$BASH_EXEC" -hash "$1" -hf_dir=./

# Clean up cache files
echo "Cleaning up cache files..."
rm -f ./*part*
