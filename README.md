> [!NOTE]
> This repo is a work in progress. We are refactoring our codebase and preparing for a 100% open-source release. We should be done soon. Stay tuned for updates, and join us in building truly open AI!

# Truly Open AI: The AI Layer of the New Internet

Truly Open AI is a production-grade, crypto-native infrastructure for managing decentralized agents across multiple blockchains. It has a comprehensive suite of tools that lets developers deploy, maintain, and scale AI agents onchain with ease.

Truly Open AI was originally developed by the [Eternal AI](https://eternalai.org) team to orchestrate [AI agents on Bitcoin](https://x.com/punk3700/status/1870757446643495235). However, the infrastructure is versatile enough to power other agent frameworks and blockchains.

As the end-to-end platform for decentralized AI, Truly Open AI integrates with the best-of-breed products across its software stack:

* Agent frameworks: [Eliza](https://github.com/elizaOS/eliza), [Rig](https://github.com/0xPlaygrounds/rig), and [ZerePy](https://github.com/blorm-network/ZerePy)
* Blockchains: Bitcoin, Ethereum, Solana, Base, Arbitrum, ZKsync, Polygon, ApeChain, etc.
* Models: Llama, Hermes, DeepSeek, Qwen, FLUX, etc.
* Storage: Filecoin, Greenfield, etc.
  
## Motivation

We're living in the age of human-AI symbiosis, but a few centralized companies currently control the future of AI. The intersection of crypto and AI has led to exciting new projects solving decentralized AI, but these efforts are fragmented and most of them remain as fun weekend prototypes.

To address this, we need to integrate these disparate components into a cohesive whole, creating a comprehensive infrastructure that lets developers build production-grade decentralized agents and eventually lead to large-scale decentralized swarms.

## Design Principles

1. **Decentralize everything**. Ensure that no single point of failure or control exists by questioning every component of the Truly Open AI infrastructure and decentralizing it. 
2. **Trustless**. Use smart contracts at every step to trustlessly coordinate all parties in the system.
3. **Production grade**. Code must be written with production-grade quality and designed for scale.
4. **Everything is an agent**. Not just user-facing agents, but every component in the infrastructure, whether a swarm of agents, an AI model storage system, a GPU compute node, a cross-chain bridge, an infrastructure microservice, or an API, is implemented as an agent.
5. **Agents do one thing and do it well**. Each agent should have a single, well-defined purpose and perform it well.
6. **Prompting as the unified agent interface**. All agents have a unified, simplified I/O interface with prompting and response for both human-to-agent interactions and agent-to-agent interactions.
7. **Composable**. Agents can work together to perform complex tasks via a chain of prompts.

## Architecture

Truly Open AI is a crypto-native open source platform created for a wide array of agent frameworks and blockchains.

<img width="2704" alt="eternal-kernel-new-7" src="https://github.com/user-attachments/assets/d0fd6429-510c-4114-83a1-c3b5aebd753f" />

<br><br>
Here are the major components of the Truly Open AI software stack.
<br>

|Component | Description |
|----------------|--|
| [agent-as-a-service](/agent-as-a-service) | The production-grade agent launchpad and management. |
| [ai-architectures](/ai-architectures) | The various AI architectures, including Chain of Thought, Plan and Execute, Critique Revise, and Self-Ask. |
| [ai-frameworks](/ai-frameworks) | Truly Open AI works with popular agent frameworks, including Eliza, Rig, and ZerePy. |
| [ai-kernel](/agent-as-a-service) | The central component of the AI-powered blockchain architecture. |
| [blockchains](/blockchains) | Truly Open AI is deployed on popular blockchains, including Bitcoin, Ethereum, and Solana. |
| [creator-tools](/creator-tools) | No-code tools for AI creators to create and manage their agents. |
| [decentralized-agents](/decentralized-agents) | A set of smart contract standards for fully onchain AI agents, including AI721 and AI721Knowledge. |
| [decentralized-compute](/decentralized-compute) | The DePIN infrastructure powering agents. |
| [decentralized-inference](/decentralized-inference) | The trustless, onchain-verifiable inference protocol. |
| [decentralized-storage](/decentralized-storage) | Truly Open AI works with popular decentralized storage networks, including Filecoin, Greenfield, and Arweave. |
| [decentralized-swarms](/decentralized-swarms) | Fully-onchain decentralized multi-agent systems |

## Getting Started

Let's deploy an end-to-end decentralized AI infrastructure.

**Step 1: Deploy a local blockchain for development**

TODO: Write this section

**Step 2: Deploy your first Decentralized Compute cluster**

TODO: Write this section

**Step 3: Deploy your production-grade Agent as a Service infrastructure**

TODO: Write this section

**Step 4: Deploy your first Decentralized Agent with AI-721**

TODO: Write this section

**Step 5: Interact with your Agent**

TODO: Write this section

**Step 6: Build a custom Knowledge Base (RAG) for your Agent with KB-721**

TODO: Write this section


## Contribute to Truly Open AI

Thank you for considering contributing to the source code. We welcome contributions from anyone and are grateful for even the most minor fixes.

If you'd like to contribute to Truly Open AI, please fork, fix, commit, and send a pull request for the maintainers to review and merge into the main code base.

If you wish to submit more complex changes, please first check with the core developers to ensure they align with the project's general philosophy and get some early feedback. This can make your efforts and our review and merge procedures quick and simple.

## Communication

* [GitHub Issues](https://github.com/eternalai-org/truly-open-ai/issues): bug reports, feature requests, issues, etc.
* [GitHub Discussions](https://github.com/eternalai-org/truly-open-ai/discussions): discuss designs, research, new ideas, thoughts, etc.
* [X (Twitter)](https://x.com/cryptoeternalai): announcements about Truly Open AI
