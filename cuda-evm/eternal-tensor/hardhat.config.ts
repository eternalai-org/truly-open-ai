import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import '@openzeppelin/hardhat-upgrades';
import "hardhat-deploy";
// import "hardhat-tracer";
import "@foundry-rs/hardhat-anvil";
import 'dotenv/config'

let localTestMnemonic = "test test test test test test test test test test test junk";
const config: HardhatUserConfig = {
  defaultNetwork: "anvil",
  anvil: {
    url: "http://127.0.0.1:8545/",
    launch: false, // if set to `true`, it will spawn a new instance if the plugin is initialized, if set to `false` it expects an already running anvil instance
  },
  solidity: {
    compilers: [
      {
        version: "0.8.19",
        settings: {
          viaIR: true,
          optimizer: {
            enabled: true,
            details: {
              yulDetails: {
                optimizerSteps: "u",
              },
            },
          },
        }
      },
    ]
  },
  networks: {
    anvil: {
      url: "http://127.0.0.1:8545/",
      timeout: 600_000,
      gas: 10_000_000_000,
      gasPrice: "auto",
    },
    mainnet: {
      url: "https://node.eternalai.org/",
      chainId: 43338,
      accounts: [
        process.env.MAINNET_PRIVATE_KEY || "",
      ],
      timeout: 600_000,
      gas: 90_000_000,
      gasPrice: "auto", 
    },
    cudatest: {
      url: "https://cuda-eternalai.testnet.l2aas.com/rpc",
      chainId: 42070,
      accounts: [
        process.env.CUDATEST_PRIVATE_KEY || "",
      ],
      timeout: 600_000,
      gas: 90_000_000,
      gasPrice: "auto", 
    }
  },
  namedAccounts: {
    deployer: 0,
  },
  paths: {
    sources: './contracts',
    tests: './scripts/tests',
    cache: './cache',
    artifacts: './artifacts',
  },
  mocha: {
    timeout: 2000000,
    color: true,
    reporter: 'mocha-multi-reporters',
    reporterOptions: {
      configFile: './mocha-report.json',
    },
  },
};

export default config;
