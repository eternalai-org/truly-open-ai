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
  defaultNetwork: "symbiosis_mainnet",
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
    base_mainnet: {
      url: process.env.BASE_MAINNET_RPC_URL || "",
      chainId: 8453,
      senderKey: process.env.BASE_MAINNET_PRIVATE_KEY,
      promptSchedulerAddress: process.env.BASE_MAINNET_PROMPT_SCHEDULER_ADDRESS,
      aiKernelAddress: process.env.BASE_MAINNET_AI_KERNEL_ADDRESS,
      aiPoweredWallet: process.env.BASE_MAINNET_AI_POWERED_WALLET_ADDRESS,
      receiverAddress: process.env.BASE_MAINNET_RECEIVER_ADDRESS,
      transferredAmount: process.env.BASE_MAINNET_TRANSFERRED_AMOUNT,
      allowUnlimitedContractSize: true,
      ethNetwork: "https://testnet.runechain.com/rpc", // The Ethereum Web3 RPC URL.
      zksync: false,
      gasPrice: "auto",
    } as any,
    symbiosis_mainnet: {
      url:
        process.env.SYMBIOSIS_MAINNET_RPC_URL ||
        "https://rpc.symbiosis.eternalai.org",
      chainId: 45762,
      senderKey: process.env.SYMBIOSIS_MAINNET_PRIVATE_KEY,
      promptSchedulerAddress:
        process.env.SYMBIOSIS_MAINNET_PROMPT_SCHEDULER_ADDRESS,
      aiKernelAddress: process.env.SYMBIOSIS_MAINNET_AI_KERNEL_ADDRESS,
      aiPoweredWallet: process.env.SYMBIOSIS_MAINNET_AI_POWERED_WALLET_ADDRESS,
      receiverAddress: process.env.SYMBIOSIS_MAINNET_RECEIVER_ADDRESS,
      transferredAmount: process.env.SYMBIOSIS_MAINNET_TRANSFERRED_AMOUNT,
      allowUnlimitedContractSize: true,
      ethNetwork: "https://testnet.runechain.com/rpc", // The Ethereum Web3 RPC URL.
      zksync: true,
      gasPrice: "auto",
    } as any,
    abstract_testnet: {
      url:
        process.env.ABSTRACT_TESTNET_RPC_URL || "https://api.testnet.abs.xyz",
      chainId: 11124,
      senderKey: process.env.ABSTRACT_TESTNET_PRIVATE_KEY,
      promptSchedulerAddress:
        process.env.ABSTRACT_TESTNET_PROMPT_SCHEDULER_ADDRESS,
      aiKernelAddress: process.env.ABSTRACT_TESTNET_AI_KERNEL_ADDRESS,
      aiPoweredWallet: process.env.ABSTRACT_TESTNET_AI_POWERED_WALLET_ADDRESS,
      receiverAddress: process.env.ABSTRACT_TESTNET_RECEIVER_ADDRESS,
      transferredAmount: process.env.ABSTRACT_TESTNET_TRANSFERRED_AMOUNT,
      allowUnlimitedContractSize: true,
      ethNetwork: "https://testnet.runechain.com/rpc", // The Ethereum Web3 RPC URL.
      zksync: true,
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
};

export default config;
