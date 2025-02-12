# Agent as a Service

<span style="display: flex">
    <img src="https://s2.coinmarketcap.com/static/img/coins/64x64/31401.png" />
</span>

## 🚀 Quick Start

### Use the Starter (Recommended)

```bash
git clone https://github.com/eternalai-org/eternal-ai

cd agent-as-a-service && cp .env.example .env

yarn && yarn build
```

### Create and Manage Your Agent

1. Modify the Default character:
    - Open `/dagent/src/dagentCharacter.ts` to modify the default character. Uncomment and edit.

2. Set `OWNER_PRIVATE_KEY`:
    - Required for authentication with the Eternal AI API.
    - EVM private key of the agent.
   
3. Set `ETERNAL_AI_URL`:
    - Required for connecting to the Eternal AI API.
    - Defaults to `https://api.eternal.ai`.
   
4. [Example dagent code](https://github.com/eternalai-org/eternal-ai/blob/master/agent-as-a-service/dagent/src/agent/index.ts)
    - Open `dagent/src/agent/index.ts` to read example code
    - Create an agent if you don't have one:
      - `await basicAgent.create();`
    - Get created agent information:
      - `await basicAgent.getAgentById(agent_id);`
           - agent_id: agent id from the agent creation step
           - If you don't remember the agent_id, you can get it by running: `await basicAgent.ownedAgents();`

5. Run the agent:
   - `yarn start:agent`


### Register your app on Shop apps and run your service

```bash
yarn start:direct
```
[Example service code](https://github.com/eternalai-org/eternal-ai/blob/master/agent-as-a-service/dagent/src/direct/index.ts)