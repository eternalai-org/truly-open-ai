# dAgent Contracts

This repository contains the smart contracts for EternalAI, enabling users to launch their own decentralized agents (dAgents).

## Table of Contents

- [dAgent Contracts](#dagent-contracts)
  - [Table of Contents](#table-of-contents)
  - [Contracts Overview](#contracts-overview)
  - [Prerequisites](#prerequisites)
  - [Setup](#setup)
    - [1. Configure Environment Variables](#1-configure-environment-variables)
    - [2. Install Dependencies](#2-install-dependencies)
    - [3. Compilation](#3-compilation)
    - [4. Deployment](#4-deployment)
      - [Deploy to a Specific Network](#deploy-to-a-specific-network)
      - [Example: Deploy to Base Mainnet](#example-deploy-to-base-mainnet)


## Contracts Overview

- **SystemPromptManager**: Manages dAgents as ERC721 NFTs. Each dAgent has a system prompt and a mission, used to clarify context before forwarding user chat prompts to the HybridModel.
- **ModelCollection**: The collection of AI models. 
- **HybridModel**: Represents the AI model, receiving requests from SystemPromptManager or EOAs and forwarding them to the WorkerHub.
- **WorkerHub**: Processes chat prompts (inference) via the AI network maintained by Workers.
- **StakingHub**: Handles staking operations, requiring Workers to stake tokens before joining the AI network and servicing requests.

## Prerequisites

- Node.js and npm installed
- Hardhat installed (globally or via npx)
- Infura API Key for Arbitrum Mainnet
- Necessary private keys for deployment

## Setup

### 1. Configure Environment Variables
Navigate to the `smart-contracts` folder. Copy the example environment file and fill in the required variables:

```bash
cp .env.example .env
```

Update the .env file with your details. You will need to provide these for both Base Mainnet and Arbitrum Mainnet:

- *_MAINNET_INFURA_API_KEY: Your Infura API key (replace * with BASE or ARBITRUM)
- *_MAINNET_PRIVATE_KEY: Private key of the deployer 
- *_MAINNET_PRIVATE_KEY_WORKER_1: Private key of worker 1
- *_MAINNET_PRIVATE_KEY_WORKER_2: Private key of worker 2
- *_MAINNET_PRIVATE_KEY_WORKER_3: Private key of worker 3
- *_MAINNET_L2_OWNER_ADDRESS: Owner address of the L2 chain who will receive rewards for each inference

**Note:** You *can* reuse the same private keys for both Base Mainnet and Arbitrum Mainnet for testing purposes. However, **for security reasons, it is strongly recommended to use separate, unique private keys for production environments.** Reusing private keys across multiple networks can increase the risk of compromise.

Note: Workers are entities responsible for maintaining the AI consensus.

### 2. Install Dependencies
```bash
npm install
```

### 3. Compilation
Compile the smart contracts using Hardhat:

```bash
npx hardhat compile
```

### 4. Deployment
#### Deploy to a Specific Network
Replace <YOUR_NETWORK> with the target network (e.g., base_mainnet):
```bash
npx hardhat run scripts/autoDeploy.ts --network <YOUR_NETWORK>
```
#### Example: Deploy to Base Mainnet

```bash
npx hardhat run scripts/autoDeploy.ts --network base_mainnet
```
