# Agent as a Service

<span style="display: flex">
    <img src="https://s2.coinmarketcap.com/static/img/coins/64x64/31401.png" />
</span>

### 🗂️ Folder Structure
```plaintext
eternal-dagent
├── packages/                # Shared libraries
│   ├── core                 # Shared core logic
│   │   ├── index.js         # Entry point for shared logic
│   │   └── package.json     # Library's package.json
│   ├── client-twitter       # Twitter typescript client, extends core
│   │   ├── index.js         # Entry point for shared logic, interfaces with Twitter API
│   │   └── package.json     # Library's package.json
│   ├── client-farcaster     # Farcaster typescript client, extends core
│   │   ├── src/index.js     # Entry point for shared logic, interfaces with Farcaster API
│   │   └── package.json     # Library's package.json
│   └── client-dagent        # Wrapper for all clients
│       ├── src/index.js     # Entry point for shared logic, interfaces with all clients
│       └── package.json     # Library's package.json
├── client                   # Client code
│   └── browser              # Reactjs code user interface
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
    └── packages-build.js    # Script to build all packages
```

## 🚀 Quick Start

### Use the Starter (Recommended)

```bash
git clone https://github.com/TrustlessComputer/eternal-dagent

cp .env.example .env

yarn && yarn build && yarn task
```
