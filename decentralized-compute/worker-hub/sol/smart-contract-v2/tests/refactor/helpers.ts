import { ethers, upgrades, network } from "hardhat";
import { expect, assert } from "chai";
import {
  IWorkerHub,
  ModelCollection,
  DAOToken,
  WrappedEAI,
  SystemPromptManager,
  StakingHub,
  SquadManager,
  PromptScheduler,
  Dagent721,
} from "../../typechain-types/index.js";

export async function deployWEAI() {
  const fact = await ethers.getContractFactory("WrappedEAI");
  const ins = await fact.deploy();
  await ins.waitForDeployment();

  const address = await ins.getAddress();

  return address;
}

export async function deployTreasury(wrappedToken: string) {
  assert.ok(wrappedToken, "wrappedToken is required");

  const constructorParams = [wrappedToken];

  const fact = await ethers.getContractFactory("Treasury");
  const proxy = await upgrades.deployProxy(fact, constructorParams);
  await proxy.waitForDeployment();

  return proxy.target as string;
}

export async function deployModelCollection(treasury: string, wEAI: string) {
  assert.ok(treasury, "treasury is required");
  assert.ok(wEAI, "wEAI is required");

  const name = "Eternal AI";
  const symbol = "";
  const mintPrice = ethers.parseEther("0");
  const royaltyReceiver = treasury;
  const royalPortion = 5_00;
  const nextModelId = 300_001;

  const constructorParams = [
    name,
    symbol,
    mintPrice,
    royaltyReceiver,
    royalPortion,
    nextModelId,
    wEAI,
  ];

  const fact = await ethers.getContractFactory("ModelCollection");

  const proxy = await upgrades.deployProxy(fact, constructorParams);
  await proxy.waitForDeployment();
  return proxy.target as string;
}

export async function deployStakingHub(
  wEAI: string,
  modelCollection: string,
  treasury: string
) {
  assert.ok(wEAI, "wEAI is required");
  assert.ok(modelCollection, "modelCollection is required");
  assert.ok(treasury, "treasury is required");

  const minerMinimumStake = ethers.parseEther("0.1");
  const blockPerEpoch = 600;
  const rewardPerEpoch = ethers.parseEther("0");
  const unstakeDelayTime = 10 * 60;
  const penaltyDuration = 3600;
  const finePercentage = 5_00;
  const minFeeToUse = ethers.parseEther("0.1");

  const constructorParams = [
    wEAI,
    modelCollection,
    treasury,
    minerMinimumStake,
    blockPerEpoch,
    rewardPerEpoch,
    unstakeDelayTime,
    penaltyDuration,
    finePercentage,
    minFeeToUse,
  ];

  const fact = await ethers.getContractFactory("StakingHub");

  const proxy = await upgrades.deployProxy(fact, constructorParams);
  await proxy.waitForDeployment();

  return proxy.target as string;
}

export async function deployPromptScheduler(wEAI: string, stakingHub: string) {
  assert.ok(wEAI, "wEAI is required");
  assert.ok(stakingHub, "stakingHub is required");

  const minerRequirement = 3;
  const submitDuration = 10 * 60;
  const feeRatioMinerValidator = 50_00;
  const batchPeriod = 10 * 60 * 60;

  const constructorParams = [
    wEAI,
    stakingHub,
    minerRequirement,
    submitDuration,
    feeRatioMinerValidator,
    batchPeriod,
  ];

  const fact = await ethers.getContractFactory("PromptScheduler");

  const proxy = await upgrades.deployProxy(fact, constructorParams);
  await proxy.waitForDeployment();
  const proxyAddress = proxy.target as string;

  //* Set PromptScheduler in StakingHub
  await stakingHubUpdatePromptSchedulerAddress(stakingHub, proxyAddress);

  return proxyAddress;
}

export async function stakingHubUpdatePromptSchedulerAddress(
  stakingHub: string,
  promptScheduler: string
) {
  const fact = await ethers.getContractFactory("StakingHub");
  const stakingHubIns = fact.attach(stakingHub) as StakingHub;
  await stakingHubIns.setPromptSchedulerAddress(promptScheduler);
}

export async function deploySystemPromptManager(
  promptScheduler: string,
  stakingHub: string,
  treasury: string
) {
  const name = "Eternal AI";
  const symbol = "1.0";
  const mintPrice = ethers.parseEther("0");
  const royaltyReceiver = treasury;
  const royalPortion = 5_00;
  const nextModelId = 1;

  const constructorParams = [
    name,
    symbol,
    mintPrice,
    royaltyReceiver,
    royalPortion,
    nextModelId,
    promptScheduler,
    stakingHub,
  ];

  const contractFact = await ethers.getContractFactory("SystemPromptManager");

  const proxy = await upgrades.deployProxy(contractFact, constructorParams);
  await proxy.waitForDeployment();
  const proxyAddress = proxy.target as string;

  // Mint agent
  // const nftOwner = (await ethers.getSigners())[0].address;
  // await mintAgent(proxyAddress, nftOwner);

  return proxyAddress;
}

export async function deployDagent721(
  wEAI: string,
  stakingHub: string,
  treasury: string
) {
  const name = "Eternal AI";
  const symbol = "1.0";
  const mintPrice = ethers.parseEther("0");
  const royaltyReceiver = treasury;
  const royalPortion = 5_00;
  const nextModelId = 1;

  const constructorParams = [
    name,
    symbol,
    mintPrice,
    royaltyReceiver,
    royalPortion,
    nextModelId,
    stakingHub,
    wEAI,
  ];

  const contractFact = await ethers.getContractFactory("Dagent721");

  const proxy = await upgrades.deployProxy(contractFact, constructorParams);
  await proxy.waitForDeployment();
  const proxyAddress = proxy.target as string;

  // Mint agent
  // const nftOwner = (await ethers.getSigners())[0].address;
  // await mintAgent(proxyAddress, nftOwner);

  return proxyAddress;
}

export async function deploySquadManager(systemPromptManager: string) {
  assert(systemPromptManager, "systemPromptManager is required");

  const constructorParams = [systemPromptManager];

  const fact = await ethers.getContractFactory("SquadManager");

  const proxy = await upgrades.deployProxy(fact, constructorParams);
  await proxy.waitForDeployment();
  const proxyAddress = proxy.target as string;

  // Set SquadManager in SystemPromptManager
  await systemPromptManagerUpdateSquadManagerAddress(
    systemPromptManager,
    proxyAddress
  );

  return proxyAddress;
}

export async function systemPromptManagerUpdateSquadManagerAddress(
  sysPrt: string,
  squad: string
) {
  const ins = (await getContractInstance(
    "SystemPromptManager",
    sysPrt
  )) as SystemPromptManager;
  await ins.setSquadManager(squad);
}

export async function mintModel(collection: string, modelOwner: string) {
  const uri = "DAGENT v2";
  const ins = (await getContractInstance(
    "ModelCollection",
    collection
  )) as ModelCollection;
  await ins.mint(modelOwner, uri);
}

export async function mintAgent(
  dagent: string,
  agentOwner: string,
  promptScheduler: string,
  modelId: number
) {
  const ins = (await getContractInstance("Dagent721", dagent)) as Dagent721;

  const linkPrompt = "Dagent721";

  const uri = linkPrompt;
  const data = ethers.toUtf8Bytes(linkPrompt);
  const fee = ethers.parseEther("0");
  const promptKey = "tiktok";

  await ins.mint(
    agentOwner,
    uri,
    data,
    fee,
    promptKey,
    promptScheduler,
    modelId
  );
}

export async function registerModel(stakingHub: string, modelId: number) {
  const ins = (await getContractInstance(
    "StakingHub",
    stakingHub
  )) as StakingHub;

  await ins.registerModel(modelId, 1, ethers.parseEther("0.1"));
}

export async function getContractInstance(
  contractName: string,
  address: string
) {
  const contractFact = await ethers.getContractFactory(contractName);
  return contractFact.attach(address);
}
