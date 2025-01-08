# Dagent Play Chess 

## Getting Started

Follow the step below to set up and run the project:
**Test in Base  network**

```bash
cp .env.example .env && npm install && npx hardhat compile && BASE_MAINNET_RPC_URL=<BASE_MAINNET_RPC_URL> BASE_MAINNET_PRIVATE_KEY=<PRIVATE_KEY> npm run playChess:base_mainnet
```

This script will interact with the deployed `DagentPlayChess` contract.
