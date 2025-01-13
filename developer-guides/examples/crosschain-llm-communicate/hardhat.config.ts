// Get the environment configuration from .env file
//
// To make use of automatic environment setup:
// - Duplicate .env.example file and name it .env
// - Fill in the environment variables
import 'dotenv/config';

import 'hardhat-deploy';
import 'hardhat-contract-sizer';
import '@nomiclabs/hardhat-ethers';
import '@layerzerolabs/toolbox-hardhat';
import {
    HardhatUserConfig,
    HttpNetworkAccountsUserConfig,
} from 'hardhat/types';
import '@typechain/hardhat';
import '@nomicfoundation/hardhat-verify';

import { EndpointId } from '@layerzerolabs/lz-definitions';

// Set your preferred authentication method
//
// If you prefer using a mnemonic, set a MNEMONIC environment variable
// to a valid mnemonic
const MNEMONIC = process.env.MNEMONIC;

// If you prefer to be authenticated using a private key, set a PRIVATE_KEY environment variable
const PRIVATE_KEY = process.env.PRIVATE_KEY;

const accounts: HttpNetworkAccountsUserConfig | undefined = MNEMONIC
    ? { mnemonic: MNEMONIC }
    : PRIVATE_KEY
      ? [PRIVATE_KEY]
      : undefined;

if (accounts == null) {
    // console.warn(
    //     'Could not find MNEMONIC or PRIVATE_KEY environment variables. It will not be possible to execute transactions in your example.'
    // );
}

const config: HardhatUserConfig = {
    paths: {
        cache: 'cache/hardhat',
    },
    solidity: {
        compilers: [
            {
                version: '0.8.22',
                settings: {
                    optimizer: {
                        enabled: true,
                        runs: 200,
                    },
                },
            },
        ],
    },
    networks: {
        'sepolia-testnet': {
            eid: EndpointId.SEPOLIA_V2_TESTNET,
            url: process.env.RPC_URL_SEPOLIA || 'https://sepolia.drpc.org',
            accounts,
            privateKey: process.env.SEPOLIA_TESTNET_PRIVATE_KEY,
            promptScheduler: process.env.SEPOLIA_TESTNET_WORKER_HUB_ADDRESS,
            hybridModel: process.env.SEPOLIA_TESTNET_HYBRID_MODEL_ADDRESS,
            questioner: process.env.SEPOLIA_TESTNET_QUESTIONER_ADDRESS,
            modelName: process.env.SEPOLIA_TESTNET_MODEL_NAME,
        },
        'amoy-testnet': {
            eid: EndpointId.AMOY_V2_TESTNET,
            url:
                process.env.RPC_URL_AMOY ||
                'https://rpc-amoy.polygon.technology',
            accounts,
            privateKey: process.env.AMOY_TESTNET_PRIVATE_KEY,
            promptScheduler: process.env.AMOY_TESTNET_WORKER_HUB_ADDRESS,
            hybridModel: process.env.AMOY_TESTNET_HYBRID_MODEL_ADDRESS,
            questioner: process.env.AMOY_TESTNET_QUESTIONER_ADDRESS,
            modelName: process.env.AMOY_TESTNET_MODEL_NAME,
        },
        'abstract-testnet': {
            eid: EndpointId.ABSTRACT_V2_TESTNET,
            url: 'https://api.testnet.abs.xyz',
            chainId: 11124,
            accounts,
        },
        'avax-mainnet': {
            eid: EndpointId.AVALANCHE_V2_MAINNET,
            url:
                'https://avalanche.public-rpc.com' ||
                process.env.AVAX_MAINNET_RPC_URL,
            chainId: 43114,
            accounts,
            privateKey: process.env.AVAX_MAINNET_PRIVATE_KEY,
            promptScheduler: process.env.AVAX_MAINNET_WORKER_HUB_ADDRESS,
            hybridModel: process.env.AVAX_MAINNET_HYBRID_MODEL_ADDRESS,
            questioner: process.env.AVAX_MAINNET_QUESTIONER_ADDRESS,
            modelName: process.env.AVAX_MAINNET_MODEL_NAME,
        },
        'ape-mainnet': {
            eid: EndpointId.APE_V2_MAINNET,
            url: 'https://rpc.apechain.com' || process.env.APE_MAINNET_RPC_URL,
            chainId: 33139,
            accounts,
            privateKey: process.env.APE_MAINNET_PRIVATE_KEY,
            promptScheduler: process.env.APE_MAINNET_WORKER_HUB_ADDRESS,
            hybridModel: process.env.APE_MAINNET_HYBRID_MODEL_ADDRESS,
            questioner: process.env.APE_MAINNET_QUESTIONER_ADDRESS,
            modelName: process.env.APE_MAINNET_MODEL_NAME,
        },
        hardhat: {
            // Need this for testing because TestHelperOz5.sol is exceeding the compiled contract size limit
            allowUnlimitedContractSize: true,
        },
    },
    namedAccounts: {
        deployer: {
            default: 0, // wallet address of index[0], of the mnemonic in .env
        },
    },
    typechain: {
        outDir: 'typechain',
        target: 'ethers-v5',
    },
    etherscan: {
        apiKey: {
            arbitrumOne: 'def456',
            sepolia: process.env.SEPOLIA_VERIFY_API_KEY || '',
            polygonAmoy: process.env.AMOY_VERIFY_API_KEY || '',
        },
    },
};

export default config;
