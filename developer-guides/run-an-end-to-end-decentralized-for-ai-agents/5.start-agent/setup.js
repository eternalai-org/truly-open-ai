const fs = require('fs').promises;
const path = require('path');

function pasreArg(args, key) {
    // Simple flag handling
    const flagIndex = args.indexOf(key);
    if (flagIndex !== -1 && args[flagIndex + 1]) {
        // console.log('Name argument:', args[flagIndex + 1]);
        return args[flagIndex + 1]
    }
    // Handle key-value pairs
    const keyValue = args.find(arg => arg.includes(`${key}=`));
    if (keyValue) {
        const [key, value] = keyValue.split('=');
        // console.log(`Key: ${key}, Value: ${value}`);
        return value;
    }
}

function validateConfig(obj) {
    // Required fields validation
    const requiredFields = [
        "ETERNALAI_RPC_URL",
        "ETERNALAI_CHAIN_ID",
        "ETERNALAI_AGENT_CONTRACT_ADDRESS",
        "ETERNALAI_AGENT_ID",
        "ETERNALAI_MODEL",
        "TWITTER_USERNAME",
        "TWITTER_PASSWORD",
        "TWITTER_EMAIL"
    ];

    const missingFields = requiredFields.filter((field) => !obj[field]);

    if (missingFields.length) {
        throw new Error(`Missing required fields: ${missingFields.join(", ")}`);
    }
};


async function setupEnv() {
    try {
        // Navigate to the config file
        const fileConfig = path.resolve(__dirname, '../../../decentralized-compute/worker-hub/env/local_contracts.json');

        const args = process.argv.slice(2);
        const TWITTER_USERNAME = pasreArg(args, "--TWITTER_USERNAME");
        const TWITTER_PASSWORD = pasreArg(args, "--TWITTER_PASSWORD");
        const TWITTER_EMAIL = pasreArg(args, "--TWITTER_EMAIL");

        const data = await fs.readFile(fileConfig, 'utf8');
        const jsonData = JSON.parse(data);

        const {
            rpc: ETERNALAI_RPC_URL,
            chain_id: ETERNALAI_CHAIN_ID,
            agent_contract_address: ETERNALAI_AGENT_CONTRACT_ADDRESS,
            agent_id: ETERNALAI_AGENT_ID,
            model_name: ETERNALAI_MODEL// review
        } = jsonData;

        const envs = {
            ETERNALAI_RPC_URL,
            ETERNALAI_CHAIN_ID,
            ETERNALAI_AGENT_CONTRACT_ADDRESS,
            ETERNALAI_AGENT_ID,
            ETERNALAI_MODEL,
            TWITTER_USERNAME,
            TWITTER_PASSWORD,
            TWITTER_EMAIL
        }

        validateConfig(envs);

        // Define the environment variables as key-value pairs
        const envData = `
ETERNALAI_URL=http://localhost:<port>/v1
ETERNALAI_CHAIN_ID=${ETERNALAI_CHAIN_ID}
ETERNALAI_RPC_URL=${ETERNALAI_RPC_URL}
ETERNALAI_AGENT_CONTRACT_ADDRESS=${ETERNALAI_AGENT_CONTRACT_ADDRESS}
ETERNALAI_AGENT_ID=${ETERNALAI_AGENT_ID}
ETERNALAI_MODEL=${ETERNALAI_MODEL}
TWITTER_USERNAME=${TWITTER_USERNAME}
TWITTER_PASSWORD=${TWITTER_PASSWORD}
TWITTER_EMAIL=${TWITTER_EMAIL}
ACTION_INTERVAL=5
ENABLE_ACTION_PROCESSING=true
MAX_ACTIONS_PROCESSING=1
ACTION_TIMELINE_TYPE=following
TWITTER_TARGET_USERS=
`;

        await fs.writeFile('.env', envData, 'utf8');
        console.log("File .env set up successfully.")

    } catch (err) {
        console.error('Error setting up env:', err);
    }
}

setupEnv();

