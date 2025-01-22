import { ethers, network, upgrades } from "hardhat";
import { GPUManager, ModelCollection, WrappedEAI } from "../typechain-types";
import { EventLog } from "ethers";

const config = network.config as any;
const networkName = network.name.toUpperCase();

export async function getContractInstance(
  proxyAddress: string,
  contractName: string
) {
  const contractFact = await ethers.getContractFactory(contractName);
  const contractIns = contractFact.attach(proxyAddress);

  return contractIns;
}

async function mintAndRegisterModel() {
  const receiver = (await ethers.getSigners())[0].address;
  const minHardware = 1;
  const metadataObj = {
    version: 1,
    model_name: "bartowski/Meta-Llama-3.1-8B-Instruct-GGUF:Q8_0",
    model_type: "text",
    model_url: "",
    model_file_hash: "",
    min_hardware: 1,
    verifier_url: "",
    verifier_file_hash: "",
  };
  const metadata = JSON.stringify(metadataObj, null, "\t");

  const ins = (await getContractInstance(
    config.collectionAddress,
    "ModelCollection"
  )) as ModelCollection;

  console.log("Minting new model...");
  const txMint = await ins.mint(receiver, metadata);
  const receiptMint = await txMint.wait();
  console.log("hash: ", receiptMint?.hash);
  console.log("status: ", receiptMint?.status);

  const newTokenEvent = (receiptMint!.logs as EventLog[]).find(
    (event: EventLog) => event.eventName === "NewToken"
  );
  if (newTokenEvent) {
    console.log("tokenId: ", newTokenEvent.args?.tokenId);
  }

  console.log("Registering model...");
  const gpuManager = (await getContractInstance(
    config.gpuManagerAddress,
    "GPUManager"
  )) as GPUManager;
  const txRegister = await gpuManager.registerModel(
    700050,
    minHardware,
    ethers.parseEther("0")
  );
  const receiptRegister = await txRegister.wait();
  console.log("hash: ", receiptRegister?.hash);
  console.log("status: ", receiptRegister?.status);
}

async function main() {
  // await mintAndRegisterModel();

  const wEAIAddr = config.wEAIAddress;
  const wEAI = (await getContractInstance(
    wEAIAddr,
    "WrappedEAI"
  )) as WrappedEAI;

  const signers = await ethers.getSigners();
  const admin = await signers[0].getAddress();
  const addrSigner1 = await signers[1].getAddress();
  const addrSigner2 = await signers[2].getAddress();
  const addrSigner3 = await signers[3].getAddress();

  console.log(
    "Worker 1: ",
    addrSigner1,
    "Balance worker 1: ",
    await ethers.provider.getBalance(addrSigner1)
  );
  console.log(
    "Worker 2: ",
    addrSigner2,
    "Balance worker 2: ",
    await ethers.provider.getBalance(addrSigner2)
  );
  console.log(
    "Worker 3: ",
    addrSigner3,
    "Balance worker 3: ",
    await ethers.provider.getBalance(addrSigner3)
  );

  // transfer wEAI to workers
  // console.log("Transfer wEAI to workers");
  // const txTransfer1 = await wEAI.transfer(
  //   addrSigner1,
  //   ethers.parseEther("25000")
  // );
  // const resTransfer1 = await txTransfer1.wait();
  // console.log("hash: ", resTransfer1?.hash);
  // console.log("status: ", resTransfer1?.status);

  // const txTransfer2 = await wEAI.transfer(
  //   addrSigner2,
  //   ethers.parseEther("25000")
  // );
  // const resTransfer2 = await txTransfer2.wait();
  // console.log("hash: ", resTransfer2?.hash);
  // console.log("status: ", resTransfer2?.status);

  // const txTransfer3 = await wEAI.transfer(
  //   addrSigner3,
  //   ethers.parseEther("25000")
  // );
  // const resTransfer3 = await txTransfer3.wait();
  // console.log("hash: ", resTransfer3?.hash);
  // console.log("status: ", resTransfer3?.status);

  const balance1 = await wEAI.balanceOf(addrSigner1);
  const balance2 = await wEAI.balanceOf(addrSigner2);
  const balance3 = await wEAI.balanceOf(addrSigner3);
  console.log("Balance worker 1: ", balance1.toString());
  console.log("Balance worker 2: ", balance2.toString());
  console.log("Balance worker 3: ", balance3.toString());

  // Approve
  const txApprove1 = await wEAI
    .connect(signers[1])
    .approve(config.gpuManagerAddress, ethers.parseEther("25000"));
  const resApprove1 = await txApprove1.wait();
  console.log("hash: ", resApprove1?.hash);
  console.log("status: ", resApprove1?.status);

  const txApprove2 = await wEAI
    .connect(signers[2])
    .approve(config.gpuManagerAddress, ethers.parseEther("25000"));
  const resApprove2 = await txApprove2.wait();
  console.log("hash: ", resApprove2?.hash);
  console.log("status: ", resApprove2?.status);

  const txApprove3 = await wEAI
    .connect(signers[3])
    .approve(config.gpuManagerAddress, ethers.parseEther("25000"));
  const resApprove3 = await txApprove3.wait();
  console.log("hash: ", resApprove3?.hash);
  console.log("status: ", resApprove3?.status);

  // Register miner
  const stakingHub = (await getContractInstance(
    config.gpuManagerAddress,
    "GPUManager"
  )) as GPUManager;

  const txRegisMiner1 = await stakingHub
    .connect(signers[1])
    ["registerMiner(uint16)"](1);
  const resRegisMiner1 = await txRegisMiner1.wait();
  console.log("hash: ", resRegisMiner1?.hash);
  console.log("status: ", resRegisMiner1?.status);

  const txRegisMiner2 = await stakingHub
    .connect(signers[2])
    ["registerMiner(uint16)"](1);
  const resRegisMiner2 = await txRegisMiner2.wait();
  console.log("hash: ", resRegisMiner2?.hash);
  console.log("status: ", resRegisMiner2?.status);

  const txRegisMiner3 = await stakingHub
    .connect(signers[3])
    ["registerMiner(uint16)"](1);
  const resRegisMiner3 = await txRegisMiner3.wait();
  console.log("hash: ", resRegisMiner3?.hash);
  console.log("status: ", resRegisMiner3?.status);

  // join mining
  const txJoinMining1 = await stakingHub.connect(signers[1]).joinForMinting();
  const resJoinMining1 = await txJoinMining1.wait();
  console.log("hash: ", resJoinMining1?.hash);
  console.log("status: ", resJoinMining1?.status);

  const txJoinMining2 = await stakingHub.connect(signers[2]).joinForMinting();
  const resJoinMining2 = await txJoinMining2.wait();
  console.log("hash: ", resJoinMining2?.hash);
  console.log("status: ", resJoinMining2?.status);

  const txJoinMining3 = await stakingHub.connect(signers[3]).joinForMinting();
  const resJoinMining3 = await txJoinMining3.wait();
  console.log("hash: ", resJoinMining3?.hash);
  console.log("status: ", resJoinMining3?.status);

  console.log("miners: ", await stakingHub.getMinerAddresses());
  console.log("miners: ", await stakingHub.getMinerAddressesOfModel(700050));
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
