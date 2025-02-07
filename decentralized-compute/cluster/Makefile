
# The binary to build, based on bin
build_golang_contract = @cat ./internal/contracts/$(1)/$(2) | jq ".abi" > ./internal/contracts/$(1)/build-$(2) && \
	cat ./internal/contracts/$(1)/$(2) | jq ".bytecode" > ./internal/contracts/$(1)/build-$(1).bin && \
	sed -i '' 's/"//g' ./internal/contracts/$(1)/build-$(1).bin && \
	abigen --pkg $(1)  --abi ./internal/contracts/$(1)/build-$(2) --bin ./internal/contracts/$(1)/build-$(1).bin --out ./internal/contracts/$(1)/$(1).go && \
	rm -rf ./internal/contracts/$(1)/build-$(1).bin && rm -rf ./internal/contracts/$(1)/build-$(2)

build_golang_contract_with_alias = @cat ./internal/contracts/$(1)/$(2) | jq ".abi" > ./internal/contracts/$(1)/build-$(2) && \
	cat ./internal/contracts/$(1)/$(2) | jq ".bytecode" > ./internal/contracts/$(1)/build-$(1).bin && \
	sed -i '' 's/"//g' ./internal/contracts/$(1)/build-$(1).bin && \
	abigen --pkg $(1)  --abi ./internal/contracts/$(1)/build-$(2) --bin ./internal/contracts/$(1)/build-$(1).bin \
	--alias royaltyPortion=RoyaltyPortion0.royaltyReceiver=RoyaltyReceiver0,_nextModelId=NextModelId0,mintPrice=MintPrice0,isManager=IsManager0,_wEAIToken=WEAIToken0 \
	--out ./internal/contracts/$(1)/$(1).go && \
	rm -rf ./internal/contracts/$(1)/build-$(1).bin && rm -rf ./internal/contracts/$(1)/build-$(2)

build_golang_contract_with_abi =@abigen --pkg $(1) --abi ./internal/contracts/$(1)/$(2) --out ./internal/contracts/$(1)/$(1).go

BIN := solo

test:

vendor:
	go mod tidy

api: vendor
	go build -o build/$(BIN) main.go

start_api: api
	./build/$(BIN) --config-file=config.env

miner: vendor
	go build -o build/miner cmd/miner/*

start_miner: miner
	./build/miner --config-file=config.env

cli: vendor
	go build -o build/cli cmd/cli/main.go

start_cli: cli
	./build/cli

download: vendor
	go build -o build/download cmd/download_model/main.go

start_download: download
	./build/download -hash=bafkreiaycapgbdqpi3lwtjvf5v4dz7v7bbjysbqnndok534fkc5k3b7ekm -hf_dir=$(shell pwd)/../models

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
	$(call build_golang_contract_with_abi,erc20,erc20.json)
	$(call build_golang_contract,prompt_scheduler,prompt_scheduler.json)
	$(call build_golang_contract,worker_hub,worker_hub.json)
	$(call build_golang_contract,gpu_manager,gpu_manager.json)
	$(call build_golang_contract,staking_hub,staking_hub.json)
	$(call build_golang_contract,load_balancer,load_balancer.json)
	$(call build_golang_contract,w_eai,w_eai.json)

abi-proxy:
	$(call build_golang_contract,proxy,proxy.json)


abi-gen-with-alias:
	$(call build_golang_contract_with_alias,model_collection,model_collection.json)


lint: vendor
	./bin/golangci-lint run ./... --timeout 10m0s

.PHONY: clean build test vendor lint-prepare lint
