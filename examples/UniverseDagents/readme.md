# Send Universe Agent Request

## Getting Started

Follow the step below to set up and run the project:
**Test in Symbiosis  network**

```bash
npm install && npx hardhat compile && RPC_URL=<YOUR_RPC_URL> PRIVATE_KEY=<0xYOUR_PRIVATE_KEY> CHOSEN_MODEL=<MODEL_NAME> USER_PROMPT=<YOUR_PROMPT>  npm run sendUniverseAgentRequest:symbiosis_mainnet
```

You can find MODEL_NAME in 'config.json'

This script will interact with the deployed `PromptScheduler` contract.
