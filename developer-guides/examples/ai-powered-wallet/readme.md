# AI Powered Wallet - Suspicious Transaction Detection

This project demonstrates an AI-powered wallet designed to detect suspicious transactions using a pre-trained model.

## Getting Started

Follow the step below to set up and run the project:
**Test in Base  network**

```bash
cp .env.example .env && npm install && npx hardhat compile && BASE_MAINNET_RPC_URL=<BASE_MAINNET_RPC_URL> BASE_MAINNET_PRIVATE_KEY=<PRIVATE_KEY>  BASE_MAINNET_RECEIVER_ADDRESS=<RECEIVER_ADDRESS> BASE_MAINNET_TRANSFERRED_AMOUNT=<AMOUNT_IN_WEI> npm run suspiciousTransaction:base_mainnet
```

**Test in Symbiosis network**
```bash
cp .env.example .env && npm install && npx hardhat compile && SYMBIOSIS_MAINNET_RPC_URL=<SYMBIOSIS_MAINNET_RPC_URL> SYMBIOSIS_MAINNET_PRIVATE_KEY=<PRIVATE_KEY>  SYMBIOSIS_MAINNET_RECEIVER_ADDRESS=<RECEIVER_ADDRESS> SYMBIOSIS_MAINNET_TRANSFERRED_AMOUNT=<AMOUNT_IN_WEI>  npm run suspiciousTransaction:symbiosis_mainnet
```

This script will interact with the deployed `AIPoweredWallet` contract and demonstrate how the AI model is used to flag potentially suspicious transactions.
