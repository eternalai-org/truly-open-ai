
# LayerZero Cross-Chain Messaging Demo

This project provides instructions for setting up and running cross-chain messaging between two networks using LayerZero protocol.

Follow the step below to set up and run the project:
**Test in Avalanche - APE  network**

#### Installation
```bash
npm install && cp .env.example .env && npx hardhat compile 
```

#### Environment Configuration
```bash
export AVAX_MAINNET_RPC_URL=<AVAX_MAINNET_RPC_URL> 
export AVAX_MAINNET_PRIVATE_KEY=<AVAX_MAINNET_PRIVATE_KEY>  
export APE_MAINNET_RPC_URL=<APE_MAINNET_RPC_URL> 
export APE_MAINNET_PRIVATE_KEY=<APE_MAINNET_PRIVATE_KEY>
export SRC_NETWORK_NAME=avax-mainnet
export DST_NETWORK_NAME=ape-mainnet
```

#### Running the Demo
``` bash
npx hardhat run scripts/askLLM.ts
```
