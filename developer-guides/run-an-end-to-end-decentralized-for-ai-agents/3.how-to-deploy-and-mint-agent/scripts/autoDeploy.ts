import assert from "assert";
import { ethers, network } from "hardhat";
import { AI721 } from "../typechain-types";
import { deployOrUpgrade, deployContract } from "./library/utils";
import { EventLog, Signer } from "ethers";
import path from "path";
import fs from "fs";

const config = network.config as any;
const networkName = network.name.toUpperCase();

async function deployAI721(
  l2OwnerAddress: string,
  hybridModelAddress: string,
  workerHubAddress: string
) {
  console.log("DEPLOY AI721...");

  assert.ok(l2OwnerAddress, `Missing ${networkName}_L2_OWNER_ADDRESS!`);
  assert.ok(hybridModelAddress, `Missing ${networkName}_HYBRID_MODEL_ADDRESS!`);
  assert.ok(workerHubAddress, `Missing ${networkName}_WORKER_HUB_ADDRESS!`);

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
  console.log((await ethers.getSigners())[0].address);

  const ai721Address = await deployAI721(
    config.l2OwnerAddress,
    config.hybridModelAddress.toString(),
    config.workerHubAddress.toString()
  );

  const deployedAddresses = {
    ai721Address,
  };

  const networkName = network.name.toUpperCase();

  await saveDeployedAddresses(networkName, deployedAddresses);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
