import '@nomiclabs/hardhat-ethers'
import '@nomiclabs/hardhat-waffle'
import '@nomiclabs/hardhat-etherscan'
import "hardhat-contract-sizer"
import "hardhat-dependency-compiler"
import "@solarity/hardhat-gobind"
import "hardhat-deploy";
import '@typechain/hardhat'

import * as dotenv from 'dotenv';
dotenv.config();

export default {
  etherscan: {
  },
  networks: {
    hardhat: {
      allowUnlimitedContractSize: true,
    },
    arbitrum: {
      url: `https://arbitrum.llamarpc.com`,
      accounts: [process.env.DEPLOYER_KEY],
    },
    base: {
      url: `https://base.llamarpc.com`,
      accounts: [process.env.DEPLOYER_KEY],
    },
    bsc: {
      url: `https://bnb-mainnet.g.alchemy.com/v2/`,
      accounts: [process.env.DEPLOYER_KEY],
    },
    bittensor: {
      url: `https://lite.chain.opentensor.ai`,
      accounts: [process.env.DEPLOYER_KEY],
    },
    apechain: {
      url: `https://apechain-mainnet.g.alchemy.com/v2/`,
      accounts: [process.env.DEPLOYER_KEY],
    },
  },
  dependencyCompiler: {
    // We have to compile from source since UniswapV3 doesn't provide artifacts in their npm package
    paths: [
      "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol",
      "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol",
    ],
  },
  namedAccounts: {
    deployer: 0,
  },
  contractSizer: {
    // max bytecode size is 24.576 KB
    alphaSort: true,
    runOnCompile: true,
    disambiguatePaths: true,
    except: ["@openzeppelin/", "test/"],
  },
  // gasReporter: {
  //   excludeContracts: ["test", '*'],
  // },
  typechain: {
    outDir: 'typechain',
    target: 'ethers-v5',
    externalArtifacts: []
  },
  solidity: {
    version: "0.8.19",
    settings: {
      "optimizer": {
        "enabled": true,
        "runs": 200
      },
      viaIR: true
    },
  },
  gobind: {
    outdir: "./artifacts/gobind",
    deployable: true,
    runOnCompile: true,
    verbose: false,
    onlyFiles: [],
    skipFiles: [],
  },
}
