import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import "@openzeppelin/hardhat-upgrades";
import "@matterlabs/hardhat-zksync";
import "@nomiclabs/hardhat-solhint";
import "hardhat-deploy";
import "dotenv/config";
import "hardhat-contract-sizer";

let localTestMnemonic =
  "test test test test test test test test test test test junk";
const config: HardhatUserConfig = {
  defaultNetwork: "localhost",
  solidity: {
    compilers: [
      {
        version: "0.8.22",
        settings: {
          optimizer: { enabled: true, runs: 2000000 },
          viaIR: true,
          evmVersion: "paris",
        },
      },
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
      stakingHubAddress: process.env.HARDHAT_STAKING_HUB_ADDRESS,
      workerHubAddress: process.env.HARDHAT_WORKER_HUB_ADDRESS,
      ai721Address: process.env.HARDHAT_AI721_ADDRESS,
      hybridModelAddress: process.env.HARDHAT_HYBRID_MODEL_ADDRESS,
      daoTokenAddress: process.env.HARDHAT_DAO_TOKEN_ADDRESS,
      wEAIAddress: process.env.HARDHAT_WEAI,
      // issue: https://github.com/NomicFoundation/hardhat/issues/3136
      // workaround: https://github.com/NomicFoundation/hardhat/issues/2672#issuecomment-1167409582
      timeout: 500_000_000,
      gas: 30_000_000,
      blockGasLimit: 2_500_000_000,
      zksync: false,
    },
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
      arbitrumOne: "def456",
    },
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
