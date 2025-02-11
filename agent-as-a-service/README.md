# Agent as a Service

<span style="display: flex">
    <img src="https://s2.coinmarketcap.com/static/img/coins/64x64/31401.png" />
</span>

### 🗂️ Folder Structure
```plaintext
eternal-dagent
├── plugins/                # Shared libraries
│   ├── core                 # Shared core logic
│   │   ├── index.js         # Entry point for shared logic
│   │   └── package.json     # Library's package.json
│   ├── plugin-twitter       # Twitter typescript plugin, extends core
│   │   ├── index.js         # Entry point for shared logic, interfaces with Twitter API
│   │   └── package.json     # Library's package.json
│   ├── plugin-farcaster     # Farcaster typescript plugin, extends core
│   │   ├── src/index.js     # Entry point for shared logic, interfaces with Farcaster API
│   │   └── package.json     # Library's package.json
│   └── client-dagent        # Wrapper for all clients
│       ├── src/index.js     # Entry point for shared logic, interfaces with all clients
│       └── package.json     # Library's package.json
├── tsconfig.json            # Shared TypeScript config (if applicable)
├── eslint.json              # Shared ESLint config
└── yarn.lock                # Dependency lock file
├── .env                     # Shared environment variables
├── .env.example             # Template for required variables
├── .gitignore               # Ignore generated files and local .env files
├── package.json             # Monorepo root package.json
└── scripts/                 # Custom management scripts
    ├── task.js              # Script to run tasks
    ├── runDagent.js         # Script to run dagent
    └── plugins-build.js     # Script to build all plugins
```

## 🚀 Quick Start

### Use the Starter (Recommended)

```bash
git clone https://github.com/eternalai-org/eternal-ai

cp .env.example .env
```

### Sample register your app on Shop apps and run your service

```bash
yarn && yarn build && yarn start:direct

```