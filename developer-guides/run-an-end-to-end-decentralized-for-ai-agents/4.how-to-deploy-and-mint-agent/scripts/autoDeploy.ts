import assert from "assert";
import { ethers, network } from "hardhat";
import { AI721 } from "../typechain-types";
import { deployOrUpgrade, deployContract } from "./library/utils";
import path from "path";
import fs from "fs";

const config = network.config as any;
const networkName = network.name.toUpperCase();
const filePath = path.join(
  __dirname,
  `../../../../decentralized-compute/worker-hub/env/local_contracts.json`
);

async function deployAI721(
  l2OwnerAddress: string,
  hybridModelAddress: string,
  workerHubAddress: string
) {
  console.log("DEPLOY AI721...");

  // assert.ok(l2OwnerAddress, `Missing ${networkName}_L2_OWNER_ADDRESS!`);
  // assert.ok(hybridModelAddress, `Missing ${networkName}_HYBRID_MODEL_ADDRESS!`);
  // assert.ok(workerHubAddress, `Missing ${networkName}_WORKER_HUB_ADDRESS!`);

  const name = "Eternal AI";
  const symbol = "";
  const mintPrice = ethers.parseEther("0");
  const royaltyReceiver = l2OwnerAddress;
  const royalPortion = 5_00;
  const nextModelId = 1;

  const constructorParams = [
    name,
    symbol,
    mintPrice,
    royaltyReceiver,
    royalPortion,
    nextModelId,
    hybridModelAddress,
    workerHubAddress,
  ];

  const ai721 = (await deployOrUpgrade(
    undefined,
    "AI721",
    constructorParams,
    config,
    true
  )) as unknown as AI721;

  await saveAI721Address(ai721.target.toString());

  return ai721.target;
}

export async function getContractInstance(
  proxyAddress: string,
  contractName: string
) {
  const contractFact = await ethers.getContractFactory(contractName);
  const contractIns = contractFact.attach(proxyAddress);

  return contractIns;
}

async function saveDeployedAddresses(networkName: string, addresses: any) {
  const filePath = path.join(__dirname, `../deployedAddresses.json`);
  let data: { [key: string]: any } = {};

  if (fs.existsSync(filePath)) {
    data = JSON.parse(fs.readFileSync(filePath, "utf8"));
  }

  data[networkName] = addresses;

  fs.writeFileSync(filePath, JSON.stringify(data, null, 2));
}

async function main() {
  const l2OwnerAddress = (await ethers.getSigners())[0].address;

  const { workerHubAddress, hybridModelAddress } = await readExternalConfig();

  const ai721Address = await deployAI721(
    l2OwnerAddress,
    hybridModelAddress.toString(),
    workerHubAddress.toString()
  );

  const deployedAddresses = {
    ai721Address,
  };

  await saveAI721Address(ai721Address.toString());
}

async function readExternalConfig() {
  if (!fs.existsSync(filePath)) {
    throw new Error(`File not found: ${filePath}`);
  }

  const data = JSON.parse(fs.readFileSync(filePath, "utf8"));

  if (!data) {
    throw new Error(`No data found in ${filePath}`);
  }

  const { workerHubAddress, hybridModelAddress } = data["contracts"];

  return { workerHubAddress, hybridModelAddress };
}

async function saveAI721Address(ai721Address: string) {
  // Read the JSON file
  const jsonData = fs.readFileSync(filePath, "utf-8");

  // Parse the JSON content
  const config = JSON.parse(jsonData);

  // Add the ai721Address to the contracts object
  config.agent_contract_address = ai721Address;

  // Write the updated JSON back to the file
  fs.writeFileSync(filePath, JSON.stringify(config, null, 2), "utf-8");
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
