/** @type import('hardhat/config').HardhatUserConfig */
require('@openzeppelin/hardhat-upgrades');

module.exports = {
  solidity: "0.8.28",
};

module.exports = {

  defaultNetwork: "hardhat",
  networks: {
    hardhat: {
      blockGasLimit: 600000000 // Network block gasLimit
    },
  },

  solidity: {
    version: "0.8.3",
    settings: {
      optimizer: {
        enabled: true,
        runs: 2**32-1,     // Optimized for SmartContract usage, not deployment cost.
      },
    },
  },

};