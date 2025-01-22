# Guide to Deploy AI721 Contract and Mint an Agent
## Contracts Overview

- **AI721**: Manages agents as ERC721 NFTs.

## Prerequisites

- Node.js and npm installed
- Hardhat installed (globally or via npx)
- Infura API Key for Arbitrum Mainnet
- Necessary private keys for deployment

## Setup

### 1. Install Dependencies 
Open a new *terminal (1)* and navigate to the `3.how-to-deploy-and-mint-agent` folder.

```bash
npm install
```

### 2. Start local blockchain
**Open new *terminal (2)***. 
*Remember to navigate to the `3.how-to-deploy-and-mint-agent` folder*.

```bash
npx hardhat node
```

### 3. Configure Environment Variables
In *terminal (1)*, setup the environment variables:

```bash
export HARDHAT_PRIVATE_KEY=<0xPRIVATE_KEY> HARDHAT_WORKER_HUB_ADDRESS=<WORKER_HUB_ADDRESS> HARDHAT_HYBRID_MODEL_ADDRESS=<HYBRID_MODEL_ADDRESS> HARDHAT_L2_OWNER_ADDRESS=<L2_OWNER_ADDRESS>
```

**Example (don't use it): **
```bash
export HARDHAT_PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 HARDHAT_WORKER_HUB_ADDRESS=0x8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199 HARDHAT_HYBRID_MODEL_ADDRESS=0xdD2FD4581271e230360230F9337D5c0430Bf44C0 HARDHAT_L2_OWNER_ADDRESS=0xbDA5747bFD65F08deb54cb465eB87D40e51B197E
```

**Note:**
- *_PRIVATE_KEY: Private key of the deployer
- *_L2_OWNER_ADDRESS: Treasury address

### 4. Deploy AI721 contract
In *terminal (1)*, run the script to deploy the new AI721 contract:

```bash
npx hardhat run ./scripts/autoDeploy.ts --network localhost
```

## 5. Prepare to mint agent
Copy the deployed AI721 contract address from *terminal (1)* and set it to the environment variable using the script below:

```bash
export HARDHAT_AGENT_OWNER_ADDRESS=<HARDHAT_AGENT_OWNER_ADDRESS> AGENT_SYSTEM_PROMPT_PATH=<AGENT_SYSTEM_PROMPT_PATH> HARDHAT_AI721_ADDRESS=<HARDHAT_AI721_ADDRESS>            
```

**Note:**
- AGENT_SYSTEM_PROMPT_PATH: the path to the file containing the agent system prompt.
- HARDHAT_AI721_ADDRESS: The AI721 contract address you have just deployed.

**Example (don't use it):** 
```bash
export HARDHAT_AGENT_OWNER_ADDRESS=0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65 AGENT_SYSTEM_PROMPT_PATH="./prompt.txt" HARDHAT_AI721_ADDRESS=0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc       
```

### 6. Mint agent
In *terminal (1)*, run script to mint an agent:

```bash
npx hardhat run ./scripts/autoDeploy.ts --network localhost
```