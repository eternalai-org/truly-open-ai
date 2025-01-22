> [!NOTE]
> WIP. We are refactoring our codebase for a 100% open-source release. Feel free to explore around.

# Project Truly Open AI

AI agents are here, but their future is controlled by a few centralized companies. The intersection of crypto and AI has led to various projects solving decentralized AI agents, but these efforts remain fragmented.

What is needed is an open infrastructure that integrates these disparate components into a cohesive whole, enabling Crypto x AI researchers to experiment with wild agent ideas and developers to build decentralized AI agents easily.

## An End-to-End Decentralized Infrastructure for AI Agents

Truly Open AI is a production-grade open source system for managing decentralized AI agents across multiple blockchains and agent frameworks. It provides powerful tools to deploy, maintain, and scale decentralized AI agents onchain with ease.

Truly Open AI was originally developed by the [Eternal AI](https://eternalai.org) team to orchestrate [AI agents on Bitcoin](https://x.com/punk3700/status/1870757446643495235). However, the infrastructure is versatile enough to power other agent frameworks ([Eliza](https://github.com/elizaOS/eliza), [Rig](https://github.com/0xPlaygrounds/rig), and [ZerePy](https://github.com/blorm-network/ZerePy)), and other blockchains ([Ethereum](https://x.com/cryptoeternalai/status/1856595790334189977), [Base](https://x.com/punk3700/status/1869428187450749093), and Solana).

<img width="2704" alt="eternal-kernel-new-7" src="https://github.com/user-attachments/assets/d0fd6429-510c-4114-83a1-c3b5aebd753f" />

## Design Principles

1. **Decentralize everything**. Ensure that no single point of failure or control exists by questioning every component of the Truly Open AI infrastructure and decentralizing it. 
2. **Trustless**. Use smart contracts at every step to trustlessly coordinate all parties in the system.
3. **Production grade**. Code must be written with production-grade quality and designed for scale.
4. **Everything is an agent**. Not just user-facing agents, but every component in the infrastructure, whether a swarm of agents, an AI model storage system, a GPU compute node, a cross-chain bridge, an infrastructure microservice, or an API, is implemented as an agent.
5. **Agents do one thing and do it well**. Each agent should have a single, well-defined purpose and perform it well.
6. **Prompting as the unified agent interface**. All agents have a unified, simplified I/O interface with prompting and response for both human-to-agent interactions and agent-to-agent interactions.
7. **Composable**. Agents can work together to perform complex tasks via a chain of prompts.

## Key Components

Here are the major components of the Truly Open AI software stack.

| Component | Description |
|:--------------------------|--------------------------|
| [decentralized-agents](/decentralized-agents)| A set of smart contract standards for decentralized AI agents (AI721s). |
| [ai-kernel](/ai-kernel)| A set of smart contracts that trustlessly coordinate the infrastructure operations. |
| [decentralized-inference](/decentralized-inference) | A set of smart contracts that together perform onchain-verifiable inference. |
| [decentralized-compute](/decentralized-compute) | A set of smart contracts that orchestrate GPU resources. |
| [agent-as-a-service](/agent-as-a-service)| The production-grade agent launchpad and management. |
| [blockchains](/blockchains)| Truly Open AI is deployed on popular blockchains, including Bitcoin, Ethereum, and Solana. |
| [creator-tools](/creator-tools)| No-code tools for AI creators to create and manage their agents. |

Here are the key components under active research.

| Component | Description |
|:--------------------------|--------------------------|
| [cuda-evm](/cuda-evm)| GPU-accelerated EVM for AI |
| [nft-ai](/nft-ai)| AI-powered fully-onchain NFTs |
| [physical-ai](/physical-ai)| AI-powered hardware devices |


## Getting Started

Let's deploy an end-to-end decentralized AI infrastructure.

**Step 1: Deploy a local blockchain for development**

TODO: Write

**Step 2: Deploy your first Decentralized Compute cluster**

TODO: Write

**Step 3: Deploy your production-grade Agent as a Service infrastructure**

* Open a new terminal and navigate to the `./agent-as-a-service/agent-orchestration/backend` folder.

Run script to build docker:
```bash
docker-compose build
```
Run script to run docker:
```bash
docker-compose up
```

**Step 4: Deploy your first Decentralized Agent with AI-721**

***Step 4.1 Deploy contract AI-721***

* Open a new terminal and navigate to the `./developer-guides/run-an-end-to-end-decentralized-for-ai-agents/3.how-to-deploy-and-mint-agent` folder.
Run script to install dependencies:
```bash
npm install
```

* Setup the environment variables and deploy AI-721 contract:
```bash
export HARDHAT_PRIVATE_KEY=<0xPRIVATE_KEY> HARDHAT_WORKER_HUB_ADDRESS=<WORKER_HUB_ADDRESS> HARDHAT_HYBRID_MODEL_ADDRESS=<HYBRID_MODEL_ADDRESS> HARDHAT_L2_OWNER_ADDRESS=<L2_OWNER_ADDRESS> && npx hardhat compile && npx hardhat run ./scripts/autoDeploy.ts --network localhost

```
*Example (don't use it):*
```bash
export HARDHAT_PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 HARDHAT_WORKER_HUB_ADDRESS=0x8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199 HARDHAT_HYBRID_MODEL_ADDRESS=0xdD2FD4581271e230360230F9337D5c0430Bf44C0 HARDHAT_L2_OWNER_ADDRESS=0xbDA5747bFD65F08deb54cb465eB87D40e51B197E && npx hardhat compile && npx hardhat run ./scripts/autoDeploy.ts --network localhost
```
**Note:**
- *_PRIVATE_KEY: Private key of the deployer
- *_L2_OWNER_ADDRESS: This acts as the treasury address.

***Step 4.2 Mint an agent***
* Copy the deployed AI721 contract address from *terminal (1)* and set it to the environment variable using the script below:

```bash
export HARDHAT_AGENT_OWNER_ADDRESS=<HARDHAT_AGENT_OWNER_ADDRESS> AGENT_SYSTEM_PROMPT_PATH=<AGENT_SYSTEM_PROMPT_PATH> HARDHAT_AI721_ADDRESS=<HARDHAT_AI721_ADDRESS> && npx hardhat run ./scripts/mintAgent.ts --network localhost
```

**Note:**
- AGENT_SYSTEM_PROMPT_PATH: the path to the file containing the agent system prompt.
- HARDHAT_AI721_ADDRESS: The AI721 contract address you have just deployed.

*Example (don't use it):*
```bash
export HARDHAT_AGENT_OWNER_ADDRESS=0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65 AGENT_SYSTEM_PROMPT_PATH="./prompt.txt" HARDHAT_AI721_ADDRESS=0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9 && npx hardhat run ./scripts/mintAgent.ts --network localhost    
```

Run script to add new agent:
```bash
curl --location 'http://localhost:8480/api/agent/create-local-agent' \
--header 'Content-Type: application/json' \
--data '{
    "creator": "<HARDHAT_AI721_ADDRESS>",
    "agent_name": "<AGENT_NAME>",
    "agent_contract_address": "<HARDHAT_AGENT_OWNER_ADDRESS>",
    "agent_contract_id": "<TOKEN_ID>",
    "system_content": "<prompt content>"
}'
```

Run script to get list agent:
```bash
curl --location 'http://localhost:8480/api/agent/list-local-agent?creator=<HARDHAT_AGENT_OWNER_ADDRESS>'
```

**Step 5: Interact with your Agent**

TODO: Write

**Step 6: Build a custom Knowledge Base (RAG) for your Agent with KB-721**

TODO: Write


## Contribute to Truly Open AI

Thank you for considering contributing to the source code. We welcome contributions from anyone and are grateful for even the most minor fixes.

If you'd like to contribute to Truly Open AI, please fork, fix, commit, and send a pull request for the maintainers to review and merge into the main code base.

If you wish to submit more complex changes, please first check with the core developers to ensure they align with the project's general philosophy and get some early feedback. This can make your efforts and our review and merge procedures quick and simple.

## Communication

* [GitHub Issues](https://github.com/eternalai-org/truly-open-ai/issues): bug reports, feature requests, issues, etc.
* [GitHub Discussions](https://github.com/eternalai-org/truly-open-ai/discussions): discuss designs, research, new ideas, thoughts, etc.
* [X (Twitter)](https://x.com/cryptoeternalai): announcements about Truly Open AI
