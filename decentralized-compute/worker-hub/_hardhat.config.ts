import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import "@openzeppelin/hardhat-upgrades";
import "@matterlabs/hardhat-zksync";
import "@nomiclabs/hardhat-solhint";
import "hardhat-deploy";
import "dotenv/config";
import "hardhat-contract-sizer";

import "./scripts/tasks/wallet.task";
import "./scripts/tasks/control.task";

let localTestMnemonic =
  "test test test test test test test test test test test junk";
const config: HardhatUserConfig = {
  defaultNetwork: "hardhat",
  solidity: {
    compilers: [
      {
        version: "0.8.19",
        settings: {
          optimizer: { enabled: true, runs: 2000000 },
          viaIR: true,
          evmVersion: "paris",
        },
      },
      {
        version: "0.8.20",
        settings: {
          optimizer: { enabled: true, runs: 2000000 },
          viaIR: true,
          evmVersion: "paris",
        },
      },
    ],
  },
  networks: {
    hardhat: {
      accounts: {
        mnemonic: localTestMnemonic,
        accountsBalance: "10000000000000000000000000",
      },
      gas: 100_000_000,
      allowUnlimitedContractSize: true,
      blockGasLimit: 1_000_000_000_000,
    },
    localhost: {
      url: "http://localhost:8545",
      accounts: {
        mnemonic: localTestMnemonic,
        count: 10,
      },
      l2OwnerAddress: process.env.HARDHAT_L2_OWNER_ADDRESS,
      treasuryAddress: process.env.HARDHAT_TREASURY_ADDRESS,
      collectionAddress: process.env.HARDHAT_COLLECTION_ADDRESS,
      workerHubAddress: process.env.HARDHAT_WORKER_HUB_ADDRESS,
      workerHubScoringAddress: process.env.HARDHAT_WORKER_HUB_SCORING_ADDRESS,
      daoTokenAddress: process.env.HARDHAT_LLAMA_TOKEN_ADDRESS, // !NOTE: must not change
      hybridModelAddress: process.env.HARDHAT_HYBRID_MODEL_ADDRESS,
      hybridModelScoringAddress:
        process.env.HARDHAT_HYBRID_MODEL_SCORING_ADDRESS,
      systemPromptManagerAddress:
        process.env.HARDHAT_SYSTEM_PROMPT_MANAGER_ADDRESS,
      // issue: https://github.com/NomicFoundation/hardhat/issues/3136
      // workaround: https://github.com/NomicFoundation/hardhat/issues/2672#issuecomment-1167409582
      timeout: 500_000_000,
      gas: 90_000_000,
      blockGasLimit: 2_500_000_000,
    } as any,
    abstract_testnet: {
      url: "https://api.testnet.abs.xyz",
      chainId: 11124,
      accounts: [
        process.env.ABSTRACT_TESTNET_PRIVATE_KEY,
        process.env.ABSTRACT_TESTNET_PRIVATE_KEY_WORKER_1,
        process.env.ABSTRACT_TESTNET_PRIVATE_KEY_WORKER_2,
        process.env.ABSTRACT_TESTNET_PRIVATE_KEY_WORKER_3,
        process.env.ABSTRACT_TESTNET_PRIVATE_KEY_4090_WORKER_1,
        process.env.ABSTRACT_TESTNET_PRIVATE_KEY_4090_WORKER_2,
        process.env.ABSTRACT_TESTNET_PRIVATE_KEY_4090_WORKER_3,
      ],
      treasuryAddress: process.env.ABSTRACT_TESTNET_TREASURY_ADDRESS,
      collectionAddress: process.env.ABSTRACT_TESTNET_COLLECTION_ADDRESS,
      gpuManagerAddress: process.env.ABSTRACT_TESTNET_GPU_MANAGER_ADDRESS,
      promptSchedulerAddress:
        process.env.ABSTRACT_TESTNET_PROMPT_SCHEDULER_ADDRESS,
      dagent721Address: process.env.ABSTRACT_TESTNET_DAGENT_721_ADDRESS,
      modelLoadBalancerAddress:
        process.env.ABSTRACT_TESTNET_MODEL_LOAD_BALANCER_ADDRESS,
      wEAIAddress: process.env.ABSTRACT_TESTNET_WEAI,
      allowUnlimitedContractSize: true,
      ethNetwork: "https://testnet.runechain.com/rpc", // The Ethereum Web3 RPC URL.
      zksync: true,
      gasPrice: "auto",
    } as any,
    base_mainnet: {
      url:
        "https://base-mainnet.infura.io/v3/" +
        process.env.BASE_MAINNET_INFURA_API_KEY,
      chainId: 8453,
      accounts: [
        process.env.BASE_MAINNET_PRIVATE_KEY,
        process.env.BASE_MAINNET_PRIVATE_KEY_WORKER_1,
        process.env.BASE_MAINNET_PRIVATE_KEY_WORKER_2,
        process.env.BASE_MAINNET_PRIVATE_KEY_WORKER_3,
      ],
      treasuryAddress: process.env.BASE_MAINNET_TREASURY_ADDRESS,
      collectionAddress: process.env.BASE_MAINNET_COLLECTION_ADDRESS,
      gpuManagerAddress: process.env.BASE_MAINNET_GPU_MANAGER_ADDRESS,
      promptSchedulerAddress: process.env.BASE_MAINNET_PROMPT_SCHEDULER_ADDRESS,
      dagent721Address: process.env.BASE_MAINNET_DAGENT_721_ADDRESS,
      modelLoadBalancerAddress:
        process.env.BASE_MAINNET_MODEL_LOAD_BALANCER_ADDRESS,
      wEAIAddress: process.env.BASE_MAINNET_WEAI,
      allowUnlimitedContractSize: true,
      ethNetwork: "https://testnet.runechain.com/rpc", // The Ethereum Web3 RPC URL.
      zksync: false,
      gasPrice: "auto",
    } as any,
  },
  namedAccounts: {
    deployer: 0,
  },
  paths: {
    sources: "./contracts",
    tests: "./tests",
    cache: "./cache",
    artifacts: "./artifacts",
  },
  etherscan: {
    apiKey: {
      regtest3: "abc123",
    },
    customChains: [
      {
        network: "regtest3",
        chainId: 20156,
        urls: {
          apiURL: "https://eternal-ai3.tc.l2aas.com/api",
          browserURL: "https://eternal-ai3.tc.l2aas.com",
        },
      },
    ],
  },
  mocha: {
    timeout: 2000000,
    color: true,
    reporter: "mocha-multi-reporters",
    reporterOptions: {
      configFile: "./mocha-report.json",
    },
  },
};

export default config;
