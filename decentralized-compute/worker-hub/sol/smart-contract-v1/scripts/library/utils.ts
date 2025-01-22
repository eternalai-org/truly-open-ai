import { Provider, Wallet } from "zksync-ethers";
import { Deployer } from "@matterlabs/hardhat-zksync";
import { ethers } from "ethers";
import { upgrades, zkUpgrades } from "hardhat";
import * as hre from "hardhat";
import * as dotenv from "dotenv";

// import '@nomiclabs/hardhat-ethers'
import "@openzeppelin/hardhat-upgrades";

import "@matterlabs/hardhat-zksync-node/dist/type-extensions";
import "@matterlabs/hardhat-zksync-verify/dist/src/type-extensions";

dotenv.config();
export const deployOrUpgrade = async (
  address: any,
  contractName: any,
  constructorParams: any[],
  networkConfig: any,
  isInitializable: boolean
) => {
  // console.log(networkConfig);
  if (networkConfig.zksync) {
    return deployOrUpgradeZk(
      getMasterWallet(),
      address,
      contractName,
      constructorParams,
      isInitializable
    );
  } else {
    return deployOrUpgradeLocal(
      address,
      contractName,
      constructorParams,
      isInitializable
    );
  }
};

export async function deployOrUpgradeZk(
  wallet: Wallet,
  address: any,
  contractName: any,
  constructorParams: any[],
  isInitializable: boolean
) {
  const deployer = new Deployer(hre, wallet);
  const artifact = await deployer.loadArtifact(contractName);
  console.log(constructorParams);

  if (address) {
    // Upgrade existing contract
    try {
      let contract = await zkUpgrades.upgradeProxy(
        deployer.zkWallet,
        address,
        artifact
      );
      await contract.waitForDeployment();
      console.log(`${contractName} contract is upgraded to ${address}`);
      return contract;
      // return address;
    } catch (e) {
      console.log("Upgrade failed:", e);
    }
  } else {
    // Deploy new contract
    try {
      const options = isInitializable ? { initializer: "initialize" } : {};
      let contract = await zkUpgrades.deployProxy(
        deployer.zkWallet,
        artifact,
        constructorParams,
        options
      );
      await contract.waitForDeployment();
      const deployedAddress = await contract.getAddress();
      console.log(`${contractName} contract is deployed to ${deployedAddress}`);
      return contract;
      // return deployedAddress;
    } catch (e) {
      console.log("Deployment failed:", e);
    }
  }

  throw new Error("Contract deployment or upgrade failed");
}

async function deployOrUpgradeLocal(
  address: any,
  contractName: any,
  constructorParams: any[],
  isInitializable: boolean
) {
  const contractFactory = await hre.ethers.getContractFactory(contractName);
  return (address = address
    ? await (async () => {
        var contract = await upgrades.upgradeProxy(address, contractFactory);
        await contract.waitForDeployment();
        console.log(`${contractName} contract is upgraded to ${address}`);
        return contract;
        // return address;
      })()
    : await (async () => {
        const options = isInitializable ? { initializer: "initialize" } : {};
        var contract = await upgrades.deployProxy(
          contractFactory,
          constructorParams,
          options
        );
        await contract.waitForDeployment();
        console.log(
          `${contractName} contract is deployed to ${await contract.getAddress()}`
        );
        return contract;
        // return await contractDeployer.getAddress();
      })());
}

type DeployContractOptions = {
  /**
   * If true, the deployment process will not print any logs
   */
  silent?: boolean;
  /**
   * If true, the contract will not be verified on Block Explorer
   */
  noVerify?: boolean;
  /**
   * If specified, the contract will be deployed using this wallet
   */
  wallet?: Wallet;
};
export const deployContract = async (
  contractArtifactName: string,
  constructorArguments?: any[],
  options?: DeployContractOptions
) => {
  const log = (message: string) => {
    if (!options?.silent) console.log(message);
  };

  log(`\nStarting deployment process of "${contractArtifactName}"...`);

  const wallet = options?.wallet ?? getMasterWallet();
  const deployer = new Deployer(hre, wallet);
  const artifact = await deployer
    .loadArtifact(contractArtifactName)
    .catch((error) => {
      if (
        error?.message?.includes(
          `Artifact for contract "${contractArtifactName}" not found.`
        )
      ) {
        console.error(error.message);
        throw `⛔️ Please make sure you have compiled your contracts or specified the correct contract name!`;
      } else {
        throw error;
      }
    });

  // Deploy the contract to zkSync
  const contract = await deployer.deploy(artifact, constructorArguments);

  const constructorArgs = contract.interface.encodeDeploy(constructorArguments);
  const fullContractSource = `${artifact.sourceName}:${artifact.contractName}`;

  // Display contract deployment info
  log(`\n"${artifact.contractName}" was successfully deployed:`);
  log(` - Contract address: ${await contract.getAddress()}`);
  log(` - Contract source: ${fullContractSource}`);
  log(` - Encoded constructor arguments: ${constructorArgs}\n`);

  return contract;
};

export const deployContractUpgradable = async (
  contractArtifactName: string,
  wallet: Wallet,
  constructorArguments?: any[]
) => {
  console.log("a", getProvider());
  const contractFactory = await hre.ethers.getContractFactory(
    contractArtifactName
  );
  console.log("b", contractArtifactName);
  const options = { initializer: "initialize" };
  console.log("c", contractFactory);
  console.log(constructorArguments);
  const contract = await upgrades.deployProxy(
    contractFactory,
    constructorArguments,
    options
  );
  console.log("d");
  await contract.waitForDeployment();
  console.log("e");
  return contract;
};

export const getProvider = () => {
  const rpcUrl = hre.network.config.url;
  if (!rpcUrl)
    throw `⛔️ RPC URL wasn't found in "${hre.network.name}"! Please add a "url" field to the network config in hardhat.config.ts`;

  // Initialize zkSync Provider
  const provider = new Provider(rpcUrl);

  return provider;
};

export const getMasterWallet = () => {
  const provider = getProvider();
  // console.log(provider);
  // console.log(
  //   "hre.network.config.accounts[0] ?? process.env.WALLET_PRIVATE_KEY!: ",
  //   hre.network.config.accounts[0] ?? process.env.WALLET_PRIVATE_KEY!
  // );
  return new Wallet(
    hre.network.config.accounts[0] ?? process.env.WALLET_PRIVATE_KEY!,
    provider
  );
};

export const getContract = async (
  adminWallet: any,
  contractName: any,
  address: any
) => {
  // console.log(contractName);
  const networkConfig = hre.network.config;
  if (networkConfig.url.includes("localhost")) {
    const ContractFactory = await hre.ethers.getContractFactory(contractName);
    return ContractFactory.attach(address);
  } else {
    const Artifact = await hre.artifacts.readArtifact(contractName);
    // Initialize contract instance for interaction
    return new ethers.Contract(
      address,
      Artifact.abi,
      adminWallet // Interact with the contract on behalf of this wallet
    );
  }
};
