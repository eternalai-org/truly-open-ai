import assert from "assert";
import { ethers, network, upgrades } from "hardhat";
import {
  AI721,
  ModelCollection,
  PromptScheduler,
  GPUManager,
  Treasury,
  WrappedEAI,
  Dagent721,
  ModelLoadBalancer,
} from "../typechain-types";
import { deployContract, deployOrUpgrade } from "./lib/utils";
import { EventLog, Signer } from "ethers";
import path from "path";
import fs from "fs";

const config = network.config as any;
const networkName = network.name.toUpperCase();

async function deployWrappedEAI() {
  console.log("DEPLOY WEAI...");
  let wEAIAddress = "";

  if (config.zksync) {
    const ins = (await deployContract("WrappedEAI", [], {
      noVerify: true,
    })) as unknown as WrappedEAI;

    wEAIAddress = ins.target as string;
  } else {
    const fact = await ethers.getContractFactory("WrappedEAI");
    const ins = (await fact.deploy()) as WrappedEAI;
    await ins.waitForDeployment();
    console.log(
        `WrappedEAI contract is deployed to ${await ins.getAddress()}`
    );
    wEAIAddress = ins.target as string;
  }

  return wEAIAddress;
}

async function deployTreasury(wEAIAddress: string) {
  console.log("DEPLOY TREASURY...");

  assert.ok(wEAIAddress, `Missing ${networkName}_WEAI!`);
  const constructorParams = [wEAIAddress];

  const treasury = (await deployOrUpgrade(
    undefined,
    "Treasury",
    constructorParams,
    config,
    true
  )) as unknown as Treasury;

  return treasury.target;
}

async function deployModelCollection(
  treasuryAddress: string,
  wEAIAddress: string
) {
  console.log("DEPLOY MODEL COLLECTION...");

  assert.ok(
    treasuryAddress,
    `Missing ${networkName}_TREASURY_ADDRESS from environment variables!`
  );
  assert.ok(
    wEAIAddress,
    `Missing ${networkName}_WEAI from environment variables!`
  );

  const name = "Eternal AI V2";
  const symbol = "Eternal_AI_V2";
  const mintPrice = ethers.parseEther("0");
  const royaltyReceiver = treasuryAddress;
  const royalPortion = 5_00;
  const nextModelId = 700_050; // TODO: @kelvin need to confirm

  const constructorParams = [
    name,
    symbol,
    mintPrice,
    royaltyReceiver,
    royalPortion,
    nextModelId,
    wEAIAddress,
  ];

  const modelCollection = (await deployOrUpgrade(
    undefined,
    "ModelCollection",
    constructorParams,
    config,
    true
  )) as unknown as ModelCollection;

  return modelCollection.target;
}

async function deployGPUManager(
  wEAIAddress: string,
  modelCollectionAddress: string,
  treasuryAddress: string
) {
  console.log("DEPLOY STAKING HUB...");

  assert.ok(
    wEAIAddress,
    `Missing ${networkName}_WEAI from environment variables!`
  );
  assert.ok(
    modelCollectionAddress,
    `Missing ${networkName}_COLLECTION_ADDRESS!`
  );
  assert.ok(treasuryAddress, `Missing ${networkName}_TREASURY_ADDRESS!`);

  const minerMinimumStake = ethers.parseEther("25000");
  const blockPerEpoch = 600;
  const rewardPerEpoch = ethers.parseEther("0"); // TODO: @kelvin need to confirm

  const unstakeDelayTime = 907200; // NOTE:  907200 blocks = 21 days (blocktime = 2)
  const penaltyDuration = 0; // NOTE: 3.3 hours
  const finePercentage = 0;
  const minFeeToUse = ethers.parseEther("0");

  const constructorParams = [
    wEAIAddress,
    modelCollectionAddress,
    treasuryAddress,
    minerMinimumStake,
    blockPerEpoch,
    rewardPerEpoch,
    unstakeDelayTime,
    penaltyDuration,
    finePercentage,
    minFeeToUse,
  ];

  const gpuManager = (await deployOrUpgrade(
    undefined,
    "GPUManager",
    constructorParams,
    config,
    true
  )) as unknown as GPUManager;
  const gpuManagerAddress = gpuManager.target;

  return gpuManagerAddress;
}

async function deployPromptScheduler(
  wEAIAddress: string,
  gpuManagerAddress: string
) {
  console.log("DEPLOY PROMPT SCHEDULER...");

  assert.ok(
    wEAIAddress,
    `Missing ${networkName}_WEAI from environment variables!`
  );
  assert.ok(
    gpuManagerAddress,
    `Missing ${networkName}_PROMPT_SCHEDULER_ADDRESS!`
  );

  const minerRequirement = 3;
  const submitDuration = 10 * 6 * 90;
  const minerValidatorFeeRatio = 50_00; // Miner earns 50% of the workers fee ( = [msg.value - L2's owner fee - treasury] )
  const batchPeriod = (24 * 60 * 60) / 2; // 24 hours
  const constructorParams = [
    wEAIAddress,
    gpuManagerAddress,
    minerRequirement,
    submitDuration,
    minerValidatorFeeRatio,
    batchPeriod,
  ];

  const promptScheduler = (await deployOrUpgrade(
    undefined,
    "PromptScheduler",
    constructorParams,
    config,
    true
  )) as unknown as PromptScheduler;
  const promptSchedulerAddress = promptScheduler.target;

  // GPU Manager update Prompt Scheduler Address
  console.log("GPU MANAGER UPDATE PROMPT SCHEDULER ADDRESS...");
  const gpuManager = (await getContractInstance(
    gpuManagerAddress,
    "GPUManager"
  )) as unknown as GPUManager;

  const txUpdate = await gpuManager.setPromptSchedulerAddress(
    promptSchedulerAddress
  );
  const receiptUpdate = await txUpdate.wait();
  console.log("Tx hash: ", receiptUpdate?.hash);
  console.log("Tx status: ", receiptUpdate?.status);

  return promptSchedulerAddress;
}

async function deployAI721(
  wEAIAddress: string,
  stakingHubAddress: string,
  treasureAddress: string
) {
  console.log("DEPLOY SYSTEM PROMPT MANAGER...");

  assert.ok(wEAIAddress, `Missing ${networkName}_WEAI!`);
  assert.ok(stakingHubAddress, `Missing ${networkName}_STAKING_HUB_ADDRESS!`);
  assert.ok(treasureAddress, `Missing ${networkName}_TREASURY_ADDRESS!`);

  const name = "Dagent 721";
  const symbol = "Dagent_721";
  const mintPrice = ethers.parseEther("0");
  const royaltyReceiver = treasureAddress;
  const royalPortion = 5_00;
  const nextModelId = 1;
  const tokenFee = wEAIAddress;

  const constructorParams = [
    name,
    symbol,
    mintPrice,
    royaltyReceiver,
    royalPortion,
    nextModelId,
    stakingHubAddress,
    tokenFee,
  ];

  const dagent721 = (await deployOrUpgrade(
    undefined,
    "Dagent721",
    constructorParams,
    config,
    true
  )) as unknown as Dagent721;

  return dagent721.target;
}

async function deployLoadBalancer(
  gpuManagerAddress: string,
  promptSchedulerAddress: string,
  wEAIAddress: string
) {
  console.log("DEPLOY MODEL LOAD BALANCER...");

  assert.ok(wEAIAddress, `Missing ${networkName}_WEAI!`);
  assert.ok(gpuManagerAddress, `Missing ${networkName}_GPU_MANAGER_ADDRESS!`);
  assert.ok(
    promptSchedulerAddress,
    `Missing ${networkName}_PROMPT_SCHEDULER_ADDRESS!`
  );

  const constructorParams = [
    gpuManagerAddress,
    promptSchedulerAddress,
    wEAIAddress,
  ];

  const ins = (await deployOrUpgrade(
    undefined,
    "ModelLoadBalancer",
    constructorParams,
    config,
    true
  )) as unknown as ModelLoadBalancer;

  return ins.target;
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
  const filePath = path.join(__dirname, `../deployedAddressesV2.json`);
  let data: { [key: string]: any } = {};

  if (fs.existsSync(filePath)) {
    data = JSON.parse(fs.readFileSync(filePath, "utf8"));
  }

  data[networkName] = addresses;

  fs.writeFileSync(filePath, JSON.stringify(data, null, 2));
}

async function main() {
  const masterWallet = (await ethers.getSigners())[0];
  const deployer = masterWallet.address;

  //const wEAIAddress = config.wEAIAddress;
  const wEAIAddress =  await deployWrappedEAI();
  const treasuryAddress = await deployTreasury(wEAIAddress.toString());
  const collectionAddress = await deployModelCollection(
    treasuryAddress.toString(),
    wEAIAddress.toString()
  );

  const gpuManagerAddress = await deployGPUManager(
    wEAIAddress.toString(),
    collectionAddress.toString(),
    treasuryAddress.toString()
  );
  const promptSchedulerAddress = await deployPromptScheduler(
    wEAIAddress.toString(),
    gpuManagerAddress.toString()
  );

  const dagent721Address = await deployAI721(
    wEAIAddress.toString(),
    gpuManagerAddress.toString(),
    treasuryAddress.toString()
  );
  const modelLoadBalancerAddress = await deployLoadBalancer(
    gpuManagerAddress.toString(),
    promptSchedulerAddress.toString(),
    wEAIAddress.toString()
  );

  const deployedAddresses = {
    deployer,
    wEAIAddress,
    treasuryAddress,
    collectionAddress,
    gpuManagerAddress,
    promptSchedulerAddress,
    dagent721Address,
    modelLoadBalancerAddress,
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
