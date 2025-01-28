# Guide to Deploy AI721 Contract and Mint an Agent
## Contracts Overview

- **AI721**: Manages agents as ERC721 NFTs.

### 1. Install dependencies 
Open a new *terminal (1)* and navigate to the `4.how-to-deploy-and-mint-agent` folder.

```bash
npm install
```

### 2. Start local blockchain
**Open new *terminal (2)***. 
*Remember to navigate to the `4.how-to-deploy-and-mint-agent` folder*.

```bash
npx hardhat node
```

### 3. Deploy AI721 contract
In *terminal (1)*, run the script to deploy the new AI721 contract:

```bash
./deploy-ai721.sh
```

## 4. Mint agent
In *terminal (1)*, run the script to mint an agent:

```bash
./mint-agent.sh ./system-prompts/naruto_fan.txt    
```
