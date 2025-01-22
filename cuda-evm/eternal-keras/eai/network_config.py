GAS_LIMIT = 99_000_000
CHUNK_LEN = 4000
MAX_PRIORITY_FEE_PER_GAS = 0
COLLECTION_ADDRESS = "0xB0e91e2Aa27330434Bfc33fF5314C796eD3Ad3c6"
DEPLOY_TIMEOUT = 300
INFER_TIMEOUT = 1000
NETWORK = {
    "testnet": {
        "NODE_ENDPOINT": "http://35.198.228.132/rpc",
        "REGISTER_ENDPOINT": "https://api-dojo.dev2.eternalai.org/api/dojo/register-model",
        "LIST_MODEL_ENDPOINT": "https://api-dojo.dev2.eternalai.org/api/dojo/list-training-requests",
        "MODEL_INFO_BY_ADDRESS": "https://api-dojo.dev2.eternalai.org/api/dojo/model-info-by-model-address",
        "MODEL_INFO_BY_ID": "https://api-dojo.dev2.eternalai.org/api/dojo/model-info",
        "EXPLORER_ENDPOINT": "https://eternal-ai3.tc.l2aas.com/be/api",
        "FAUCET_ENDPOINT": "https://api-dojo.dev2.eternalai.org/api/service/faucet-testnet"
    },
    "mainnet": {
        "NODE_ENDPOINT": "https://cuda-eternalai.testnet.l2aas.com/rpc",
        "REGISTER_ENDPOINT": "https://api-dojo2.eternalai.org/api/dojo/register-model",
        "LIST_MODEL_ENDPOINT": "https://api-dojo2.eternalai.org/api/dojo/list-training-requests",
        "MODEL_INFO_BY_ADDRESS": "https://api-dojo2.eternalai.org/api/dojo/model-info-by-model-address",
        "MODEL_INFO_BY_ID": "https://api-dojo2.eternalai.org/api/dojo/model-info",
        "EXPLORER_ENDPOINT": "https://explorer.eternalai.org/api",
        "DEPOSIT_ENDPOINT": "https://bridges-api.eternalai.org/api/generate-deposit-address",
    }
}
