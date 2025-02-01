# Eternal AI (EAI)

The Eternal AI protocol is the decentralized AI layer of the new internet. It is designed to give AI superpowers to any blockchain.

The Eternal AI protocol is implemented as a set of smart contracts that can be easily deployed on any blockchain to create an unstoppable onchain AI runtime. It is designed to deliver censorship resistance, permissionless, and verifiability. It uses peer-to-peer technology to operate without any trusted intermediaries or central authorities.

The Eternal AI protocol was originally developed for [decentralized AI agents on Bitcoin](https://x.com/punk3700/status/1870757446643495235). However, the design is versatile enough to power other blockchains as well.

Join us in building AI-powered blockchains.

# Get started

## Prerequisites
* [Node.js 22.12.0+ and npm 10.9.0+](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm)
* [Docker Desktop 4.37.1+](https://docs.docker.com/desktop/setup/install/mac-install/)
* [Go 1.23.0+](https://go.dev/doc/install)
* [Ollama 0.5.7+](https://ollama.com/download)

## Step 1: Deploy a local AI-powered blockchain on your computer

We provide a CLI `eai` to simplify the process.

To install `eai`
```bash
sudo ./install.sh
```

Then, you can use the following command and follow its interactive instructions.
```bash
eai miner setup
```

Press `1` (Setup local cluster). Press `1` again for a full auto-install.

<img width="1231" alt="image" src="https://github.com/user-attachments/assets/e997351a-ff00-4207-b969-7036c6c04497" />

Suppose you want to custom-install, press `2`. This will give you the option to install specific packages.

```bash
- 1. Create `./env/local_contracts.json` # create a default config file
- 2. Start HardHat as your local chain
- 3. Deploy Eternal AI Kernel smart contracts
- 4. Start Eternal AI Compute Nodes/Miners
- 5. Start System APIs
```

## Step 2: Deploy Decentralized Compute

For this tutorial, we'll simplify the process by having the three local miners on the same compute node. In production, each miner should have their own compute.

The miners serve DeepSeek-R1-Distill-Qwen-1.5B-Q8_0. However, you should be able to use any models.

DeepSeek-R1 is stored on Filecoin, a decentralized storage network. Its hash is [`ipfs://bafkreieglfaposr5fggc7ebfcok7dupfoiwojjvrck6hbzjajs6nywx6qi`](https://gateway.lighthouse.storage/ipfs/bafkreieglfaposr5fggc7ebfcok7dupfoiwojjvrck6hbzjajs6nywx6qi).

The miners first fetch the model weights stored in multiple chunks on Filecoin and combine them into one complete model.

For MacOS:
```bash
cd decentralized-compute/models
sudo bash download_model_macos.sh bafkreieglfaposr5fggc7ebfcok7dupfoiwojjvrck6hbzjajs6nywx6qi 
```

For Ubuntu:
```bash
cd decentralized-compute/models
sudo bash download_model_linux.sh bafkreieglfaposr5fggc7ebfcok7dupfoiwojjvrck6hbzjajs6nywx6qi 
```

Create and start an Ollama instance.
```bash
ollama create DeepSeek-R1-Distill-Qwen-1.5B-Q8 -f Modelfile
```

You can try the following quick test to ensure your miner is ready. This will send an onchain prompt to your local blockchain and see if the miners respond.

```bash
curl -X POST "http://localhost:8004/v1/chat/completions" -H "Content-Type: application/json"  -d '{
    "model": "DeepSeek-R1-Distill-Qwen-1.5B-Q8",
    "messages": [
        {
            "role": "system",
            "content": "You are a deep thinker."
        },
        {
            "role": "user",
            "content": "Design a DeFAI agent that trades autonomously on Uniswap."
        }
  ]
}'
```

Note that the model info is stored in the `Modelfile` file. In the future, if you want to change the model, update the file.
```bash
FROM DeepSeek-R1-Distill-Qwen-1.5B-Q8_0/DeepSeek-R1-Distill-Qwen-1.5B-Q8_0.gguf 
```

## Step 3: Deploy your production-grade Agent as a Service infrastructure

In this step, we'll deploy a production-grade agent orchestration platform in one single line of code. It provides powerful tools for you to deploy, manage, and scale your agents onchain.

Run the following command:
```bash
eai aaas start
```

## Step 4: Deploy your Decentralized Inference API

In this step, we'll deploy Eternal AI Decentralized Inference API that can be used instead of centralized OpenAI API. There are a lot of things to like about Decentralized Inference API. It's permissionless, trustless, censorship-resistant, tamper-proof, and onchain verifiable.

Run the following command:
```bash
eai apis start
```

## Step 5: Deploy your first Decentralized Agent with AI-721

### Step 5.1. Deploy contract AI-721

Eternal AI treats each agent as a non-fungible. [AI-721](https://github.com/eternalai-org/eternal-ai/blob/master/decentralized-agents/contracts/standards/AI721.sol) inherits ERC-721 and adds AI features.

Run the following script to install dependencies and deploy the AI-721 contract:
```bash
eai aaas deploy-contract
```

### Step 5.2. Mint an agent

Let's create an agent who is a Donald Trump twin.

Run the following script to mint an agent:

```bash
eai agent create $(pwd)/decentralized-agents/characters/donald_trump.txt
```

The .txt file is the system prompt for your agent. It will be used to set the initial behavior for your agent. You can modify the content of the prompt file to adjust your agent's personality.

Fetch agent info from the AI-721 contract:
```
eai agent info <agent_id>
```

Also, to list out all agents on your machine, run this:
```bash
eai agent list
```

## Step 6: Interact with the agent 

### 6.1. Chat with the agent

```bash
eai agent chat <agent_id>
```

### 6.2. Set up Twitter for the agent with the Eliza Engine

Navigate to the `./developer-guides/run-an-end-to-end-decentralized-for-ai-agents/5.start-agent` folder and run the following command to configure your twitter account.

```
node setup.js --TWITTER_USERNAME <TWITTER_USERNAME> --TWITTER_PASSWORD <TWITTER_PASSWORD> --TWITTER_EMAIL <TWITTER_EMAIL>
```

Then build a Docker image for the Eliza runtime.

```
docker build -t eliza .
```

And start an Eliza agent by running the following command.

```
docker run --env-file .env  -v ./config.json:/app/eliza/agents/config.json eliza
```

## Step 7 (Optional): Enter the 10,000 EAI Raffle 🎁 

Congrats! You made it! You've deployed a decentralized operating system for AI agents on your computer. 

We have a little gift for you: share your agent character file to be entered into a 10,000 EAI raffle.

It's simple. Create a pull request and add your agent character file to the folder `decentralized-agents/characters/` folder. Name the file in this format `<agent-name>-by-<your-name>.txt`.

<br>
<br>

# Platform Architecture

Eternal AI is an end-to-end decentralized infrastructure created for a wide array of AI agents and blockchains.

<img width="2704" alt="eternal-kernel-new-7" src="https://github.com/user-attachments/assets/d0fd6429-510c-4114-83a1-c3b5aebd753f" />

<br>
<br>

Here are the major components of the Eternal AI software stack.

| Component | Description |
|:--------------------------|--------------------------|
| [ai-kernel](/ai-kernel)| A set of Solidity smart contracts that trustlessly coordinate user space, onchain space, and offchain space. |
| [decentralized-agents](/decentralized-agents)| A set of Solidity smart contracts that define AI agent standards (AI-721, SWARM-721, KB-721). |
| [decentralized-inference](/decentralized-inference) | The decentralized inference APIs. |
| [decentralized-compute](/decentralized-compute) | The peer-to-peer GPU clustering and orchestration protocol. |
| [agent-as-a-service](/agent-as-a-service)| The production-grade agent launchpad and management. |
| [agent-studio](/agent-studio)| No-code, drag 'n drop, visual programming language for AI creators. |
| [blockchains](/blockchains)| A list of blockchains that are AI-powered by Eternal AI. |

Here are the key ongoing research projects.

| Component | Description |
|:--------------------------|--------------------------|
| [cuda-evm](/research/cuda-evm)| The GPU-accelerated EVM and its Solidity tensor linear algebra libary. |
| [nft-ai](/research/nft-ai)| AI-powered fully-onchain NFTs. |
| [physical-ai](/research/physical-ai)| AI-powered hardware devices. |

# Design Principles

1. **Decentralize everything**. Ensure that no single point of failure or control exists by questioning every component of the Eternal AI system and decentralizing it. 
2. **Trustless**. Use smart contracts at every step to trustlessly coordinate all parties in the system.
3. **Production grade**. Code must be written with production-grade quality and designed for scale.
4. **Everything is an agent**. Not just user-facing agents, but every component in the infrastructure, whether a swarm of agents, an AI model storage system, a GPU compute node, a cross-chain bridge, an infrastructure microservice, or an API, is implemented as an agent.
5. **Agents do one thing and do it well**. Each agent should have a single, well-defined purpose and perform it well.
6. **Prompting as the unified agent interface**. All agents have a unified, simplified I/O interface with prompting and response for both human-to-agent interactions and agent-to-agent interactions.
7. **Composable**. Agents can work together to perform complex tasks via a chain of prompts.

# Featured Integrations

Eternal AI is built using a modular approach, so support for other blockchains, agent frameworks, GPU providers, or AI models can be implemented quickly. Please reach out if you run into issues while working on an integration.

<img width="1780" alt="Featured Integrations (1)" src="https://github.com/user-attachments/assets/e6bdd4c9-3630-4dfa-8ac2-0526cb618c1e" />

# Governance

We are still building out the Eternal AI DAO.

Once the DAO is in place, [EAI holders](https://eternalai.org/eai) will oversee the governance and the treasury of the Eternal AI project with a clear mission: to build truly open AI. 

# Contribute to Eternal AI

Thank you for considering contributing to the source code. We welcome contributions from anyone and are grateful for even the most minor fixes.

If you'd like to contribute to Eternal AI, please fork, fix, commit, and send a pull request for the maintainers to review and merge into the main code base.

# Communication

* [GitHub Issues](https://github.com/eternalai-org/eternal-ai/issues): bug reports, feature requests, issues, etc.
* [GitHub Discussions](https://github.com/eternalai-org/eternal-ai/discussions): discuss designs, research, new ideas, thoughts, etc.
* [X (Twitter)](https://x.com/cryptoeternalai): announcements about Eternal AI
