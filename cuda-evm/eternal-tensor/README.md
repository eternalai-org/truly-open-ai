# Solidity ML contracts

## Installation

```shell
npm install
```

## Running local node

### With hardhat

Hardhat local node is very slow, but has console output on revert. Should be used for debugging purpose.

Running local hardhat node:

```shell
# Increase memory limit, as default memory is not sufficient to run some operation
export NODE_OPTIONS="--max-old-space-size=8192"
npx hardhat node
```

### With anvil

Anvil local node is very slow, but has console output on revert. Should be used for local testing.

Installing anvil:

```shell
# Install foundry
curl -L https://foundry.paradigm.xyz | bash

# Reload bash env
source ~/.bashrc

# Installing foundry tool (including anvil)
foundryup
```

Running local anvil node:

```shell
# Run local anvil node with no code size and gas limit
anvil --prune-history --order fifo --code-size-limit 4294967296 -m "test test test test test test test test test test test junk" --gas-limit 1000000000000 --block-time 3
```

## Running unit test

Running unit test requires network with CUDA-EVM support

```shell
npx hardhat test --network <network_name>
```
