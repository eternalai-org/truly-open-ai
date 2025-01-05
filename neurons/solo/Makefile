# The binary to build, based on bin
BIN := solo

test:

vendor:
	go mod tidy

api: vendor
	go build -o build/$(BIN) main.go

start_api: api
	./build/$(BIN) --config-file=config.env

clean:
	if [ -f ${BIN} ] ; then rm ${BIN} ; fi

lint-prepare:
	@echo "Installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

docker:
	make docker-service_miner

docker_arm:
	make docker-service_miner_arm

docker-ollama:
	docker-compose down	ollama && docker-compose build   ollama && docker-compose up -d   ollama &

docker-service_miner:
	docker-compose down service_miner && docker-compose build service_miner && docker-compose up -d service_miner

docker-service_miner_arm:
	docker-compose down	   service_miner && docker-compose   -f  docker-compose-arm.yml build service_miner && docker-compose  -f docker-compose-arm.yml up -d service_miner &

abi-gen:
	abigen --pkg worker_hub --abi ./internal/contracts/worker_hub/worker_hub.json --out ./internal/contracts/worker_hub/worker_hub.go
	abigen --pkg erc20 --abi ./internal/contracts/erc20/erc20.json --out ./internal/contracts/erc20/erc20.go
	abigen --pkg staking_hub --abi ./internal/contracts/staking_hub/staking_hub.json --out ./internal/contracts/staking_hub/staking_hub.go
	abigen --pkg gpu_manager --abi ./internal/contracts/gpu_manager/gpu_manager.json --out ./internal/contracts/gpu_manager/gpu_manager.go
	abigen --pkg prompt_scheduler --abi ./internal/contracts/prompt_scheduler/prompt_scheduler.json --out ./internal/contracts/prompt_scheduler/prompt_scheduler.go

lint: vendor
	./bin/golangci-lint run ./... --timeout 10m0s

.PHONY: clean build test vendor lint-prepare lint
