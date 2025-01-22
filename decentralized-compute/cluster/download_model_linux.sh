git clone https://github.com/eternalai-org/truly-open-ai
cd truly-open-ai/decentralized-compute/cluster

cd truly-open-ai/decentralized-compute/cluster

docker run --rm -v "$(pwd)":/app -w /app golang:1.23-bullseye sh -c 'GOOS=linux go build -o download_model cmd/download_model/main.go'

./download_model -hash $1
