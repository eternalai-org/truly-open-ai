import { ethers, upgrades, network, Signer } from "hardhat";
import { expect, assert } from "chai";
import {
  loadFixture,
  mine,
} from "@nomicfoundation/hardhat-toolbox/network-helpers";
const helpers = require("@nomicfoundation/hardhat-toolbox/network-helpers");

import {
  IWorkerHub,
  ModelCollection,
  DAOToken,
  WrappedEAI,
  SystemPromptManager,
  Dagent721,
} from "../../typechain-types/index.js";
import { AbiCoder, EventLog } from "ethers";
import { StakingHub } from "../../typechain-types/contracts/StakingHub.js";
import { SquadManager } from "../../typechain-types/contracts/SquadManager.js";
import { PromptScheduler } from "../../typechain-types/contracts/PromptScheduler.js";
import { address18 } from "../address.seed";
import * as TestHelper from "./helpers";

describe("WorkerHub contract", async () => {
  const { provider } = ethers;

  async function deployWorkerHubFixture() {
    const [admin] = await ethers.getSigners();
    console.log(`admin: ${admin.address}`);

    const wEAI = await TestHelper.deployWEAI();
    const treasury = await TestHelper.deployTreasury(wEAI);
    const modelCollection = await TestHelper.deployModelCollection(
      treasury,
      wEAI
    );
    const stakingHub = await TestHelper.deployStakingHub(
      wEAI,
      modelCollection,
      treasury
    );
    const promptScheduler = await TestHelper.deployPromptScheduler(
      wEAI,
      stakingHub
    );
    const dagent721 = await TestHelper.deployDagent721(
      wEAI,
      stakingHub,
      treasury
    );

    return {
      admin,
      treasury,
      wEAI,
      modelCollection,
      stakingHub,
      promptScheduler,
      dagent721,
    };
  }

  async function simulate(
    promptSchedulerAddress: string,
    stakingHubAddress: string,
    wEAIAddress: string,
    modelCollectionAddress: string,
    dagent721Address: string
  ) {
    const promptScheduler = (await TestHelper.getContractInstance(
      "PromptScheduler",
      promptSchedulerAddress
    )) as PromptScheduler;
    const stakingHub = (await TestHelper.getContractInstance(
      "StakingHub",
      stakingHubAddress
    )) as StakingHub;
    const wEAI = (await TestHelper.getContractInstance(
      "WrappedEAI",
      wEAIAddress
    )) as WrappedEAI;
    const modelCollection = (await TestHelper.getContractInstance(
      "ModelCollection",
      modelCollectionAddress
    )) as ModelCollection;
    console.log("dagent721Address: ", dagent721Address);
    console.log("promptSchedulerAddress: ", promptSchedulerAddress);

    const dagent721 = (await TestHelper.getContractInstance(
      "Dagent721",
      dagent721Address
    )) as Dagent721;

    const [admin] = await ethers.getSigners();

    // Mint a model
    await TestHelper.mintModel(modelCollectionAddress, admin.address);
    const currentModelId = Number((await modelCollection.nextModelId()) - 1n);

    // Mint an agent
    await TestHelper.mintAgent(
      dagent721Address,
      admin.address,
      promptSchedulerAddress,
      currentModelId
    );
    await TestHelper.registerModel(stakingHubAddress, currentModelId);
    // check register model
    expect(await stakingHub.getModelIds()).to.be.deep.eq([currentModelId]);

    // Set the balance of the impersonated account
    const hexBalance = "0x" + ethers.parseEther("100").toString(16);
    for await (let i of Array(3).keys()) {
      await ethers.provider.send("hardhat_setBalance", [
        address18[i],
        hexBalance,
      ]);

      let impersonatedSigner = await ethers.getImpersonatedSigner(address18[i]);
      // Wrap EAI
      await wEAI
        .connect(impersonatedSigner)
        .wrap({ value: ethers.parseEther("50") });
      // Approve StakingHub
      await wEAI
        .connect(impersonatedSigner)
        .approve(stakingHubAddress, ethers.parseEther("50"));

      //Regis miner and then join for minting
      await stakingHub.connect(impersonatedSigner).registerMiner(1);

      await stakingHub.connect(impersonatedSigner).joinForMinting();
    }
    expect((await stakingHub.getMinerAddresses()).length).to.eq(3);

    await ethers.provider.send("hardhat_setBalance", [
      address18[16],
      hexBalance,
    ]);
  }

  describe("WorkerHub contract", async () => {
    it("Should process the first infer", async () => {
      const {
        admin,
        treasury,
        wEAI,
        modelCollection,
        stakingHub,
        promptScheduler,
        dagent721,
      } = await loadFixture(deployWorkerHubFixture);

      await simulate(
        promptScheduler,
        stakingHub,
        wEAI,
        modelCollection,
        dagent721
      );
      const wEAIIns = (await TestHelper.getContractInstance(
        "WrappedEAI",
        wEAI
      )) as WrappedEAI;
      const modelCollectionIns = (await TestHelper.getContractInstance(
        "ModelCollection",
        modelCollection
      )) as ModelCollection;

      const currentModelId = Number(
        (await modelCollectionIns.nextModelId()) - 1n
      );

      const promptSchedulerIns = (await TestHelper.getContractInstance(
        "PromptScheduler",
        promptScheduler
      )) as PromptScheduler;

      // ***INFER***
      let impersonatedUser = await ethers.getImpersonatedSigner(address18[16]);

      await wEAIIns
        .connect(impersonatedUser)
        .wrap({ value: ethers.parseEther("50") });

      await wEAIIns
        .connect(impersonatedUser)
        .approve(promptScheduler, ethers.parseEther("100"));

      // get block number
      const blockNumber = await ethers.provider.getBlockNumber();
      const modelInput = ethers.encodeBytes32String("test");

      await promptSchedulerIns
        .connect(impersonatedUser)
        ["infer(uint32,bytes,address)"](
          currentModelId,
          modelInput,
          impersonatedUser.address
        );

      // expect inference id to be 1
      expect(await promptSchedulerIns.inferenceCounter()).to.eq(1);

      const inferInfo = await promptSchedulerIns.getInferenceInfo(1n);
      const assignedMiner = inferInfo.processedMiner;
      // check inference info
      expect(inferInfo.input).to.eq(modelInput);
      expect(inferInfo.output).to.eq("0x");
      expect([address18[0], address18[1], address18[2]]).to.include(
        inferInfo.processedMiner
      );

      expect(inferInfo.status).to.eq(1); //Solving
      expect(inferInfo.submitTimeout).to.eq(
        BigInt(blockNumber + 1) + (await promptSchedulerIns._submitDuration())
      );
      expect(inferInfo.modelId).to.eq(currentModelId);
      expect(inferInfo.value).to.eq(ethers.parseEther("0.1"));

      expect(
        await promptSchedulerIns.getInferenceByMiner(inferInfo.processedMiner)
      ).to.include(1n);
      expect(await promptSchedulerIns.getBatchInfo(currentModelId, 0)).to.eql([
        0n, // validators fee
        [1n], // inference ids
      ]);

      // submit solution
      let miner = await ethers.getImpersonatedSigner(assignedMiner);
      let solution = ethers.solidityPacked(["string"], ["No"]);
      await promptSchedulerIns.connect(miner).submitSolution(1, solution);
      await expect(
        promptSchedulerIns.connect(miner).submitSolution(1, solution)
      ).to.be.revertedWithCustomError(
        promptSchedulerIns,
        "InvalidInferenceStatus()"
      );

      // check solution
      const inferInfo1 = await promptSchedulerIns.getInferenceInfo(1n);
      expect(inferInfo1.status).to.eq(2); //Commit
      expect(inferInfo1.output).to.eq(solution);
      expect(await promptSchedulerIns.getBatchInfo(currentModelId, 0)).to.eql([
        ethers.parseEther("0.05"), // validators fee
        [1n], // inference ids
      ]);
    });

    it.only("Should process the first infer call from dagent721", async () => {
      const {
        admin,
        treasury,
        wEAI,
        modelCollection,
        stakingHub,
        promptScheduler,
        dagent721,
      } = await loadFixture(deployWorkerHubFixture);

      await simulate(
        promptScheduler,
        stakingHub,
        wEAI,
        modelCollection,
        dagent721
      );
      const wEAIIns = (await TestHelper.getContractInstance(
        "WrappedEAI",
        wEAI
      )) as WrappedEAI;
      const modelCollectionIns = (await TestHelper.getContractInstance(
        "ModelCollection",
        modelCollection
      )) as ModelCollection;

      const currentModelId = Number(
        (await modelCollectionIns.nextModelId()) - 1n
      );

      const promptSchedulerIns = (await TestHelper.getContractInstance(
        "PromptScheduler",
        promptScheduler
      )) as PromptScheduler;

      const dagent721Ins = (await TestHelper.getContractInstance(
        "Dagent721",
        dagent721
      )) as Dagent721;

      // ***INFER***
      const modelInput = ethers.encodeBytes32String("test");

      let impersonatedUser = await ethers.getImpersonatedSigner(address18[16]);
      await wEAIIns
        .connect(impersonatedUser)
        .wrap({ value: ethers.parseEther("50") });

      await wEAIIns
        .connect(impersonatedUser)
        .approve(dagent721, ethers.parseEther("100"));

      await dagent721Ins
        .connect(impersonatedUser)
        ["infer(uint256,bytes,string,string,bool,uint256)"](
          1,
          modelInput,
          "eternal ai",
          "tiktok",
          false,
          ethers.parseEther("0.1")
        );
    });

    it("Should update agent uri", async () => {
      const { admin, systemPromptManagerAddress, systemPromptHelperAddress } =
        await loadFixture(deployWorkerHubFixture);

      const ins = await getSystemPromptManagerInstance(
        systemPromptManagerAddress,
        systemPromptHelperAddress
      );

      console.log("owner: ", await ins.ownerOf(1));
      const [admin1, admin2] = await ethers.getSigners();
      console.log(admin1.address);

      const linkUri =
        "ipfs://bafkreide4kf4se2atgdi3kjie5eigvvr3wnkyolitbrj6cuj3sfzfyowui";

      const agentId = 1n;
      const randomBytes = ethers.randomBytes(8);
      const randomNonce = BigInt(
        "0x" + Buffer.from(randomBytes).toString("hex")
      ); // Convert bytes to BigInt

      const address = systemPromptManagerAddress;
      const chainId = 31337n;
      // const chainId = 8453n;

      const coder = AbiCoder.defaultAbiCoder();
      const encodedData = coder.encode(
        ["string", "uint256", "uint256", "address", "uint256"],
        [linkUri, agentId, randomNonce, address, chainId]
      );

      const hashData = ethers.keccak256(encodedData);
      const signature = await admin.signMessage(ethers.getBytes(hashData));
      const tx = await ins.updateAgentUriWithSignature(
        agentId,
        linkUri,
        randomNonce,
        signature
      );
      ///
      const encodedData2 = coder.encode(
        ["string", "uint256", "uint256", "address", "uint256"],
        [linkUri, 2n, randomNonce, address, chainId]
      );
      const hashData2 = ethers.keccak256(encodedData2);
      const signature2 = await admin.signMessage(ethers.getBytes(hashData2));

      const tx2 = await ins.updateAgentUriWithSignature(
        2n,
        linkUri,
        randomNonce,
        signature2
      );
    });

    it("Should create squad", async () => {
      const {
        admin,
        proxyWorkerHubAddress,
        hybridModelAddress,
        stakingHubAddress,
        wEAIAddress,
        systemPromptManagerAddress,
        systemPromptHelperAddress,
        squadManagerAddress,
      } = await loadFixture(deployWorkerHubFixture);

      await simulate(
        proxyWorkerHubAddress,
        stakingHubAddress,
        hybridModelAddress,
        wEAIAddress
      );
      console.log("systemPromptHelperAddress: ", systemPromptHelperAddress);
      console.log("systemPromptManagerAddress: ", systemPromptManagerAddress);

      const agentIns = await getSystemPromptManagerInstance(
        systemPromptManagerAddress,
        systemPromptHelperAddress
      );
      const squadIns = (await getContractInstance(
        "SquadManager",
        squadManagerAddress
      )) as SquadManager;

      const [admin1, admin2] = await ethers.getSigners();
      await squadIns.connect(admin1).createSquad([]);
      expect(await squadIns.currentSquadId()).to.eq(1n);
      expect(await squadIns.squadOwner(1n)).to.eq(admin1.address);
      expect((await squadIns.getAgentIdsBySquadId(1n)).length).to.eq(0);

      await squadIns.connect(admin1).createSquad([1n]);
      expect(await squadIns.currentSquadId()).to.eq(2n);
      expect(await squadIns.squadOwner(2n)).to.eq(admin1.address);
      expect((await squadIns.getAgentIdsBySquadId(2n)).length).to.eq(1);
      expect((await squadIns.getAgentIdsBySquadId(2n))[0]).to.eq(1);
      expect(await squadIns.agentToSquadId(1n)).to.eq(2n);

      await squadIns.connect(admin1).moveAgentsToSquad([1n], 1n);
      expect(await squadIns.agentToSquadId(1n)).to.eq(1n);
      expect(await squadIns.squadBalance(admin.address)).to.eq(2);

      await agentIns
        .connect(admin1)
        ["mint(address,string,bytes,uint256,uint256)"](
          admin1.address,
          "x",
          ethers.toUtf8Bytes("x"),
          0,
          1
        );
      expect(await squadIns.currentSquadId()).to.eq(2n);
      expect((await agentIns.getAgentIdByOwner(admin.address)).length).to.eq(3);
      expect(await squadIns.squadOwner(1n)).to.eq(admin1.address);
      expect((await squadIns.getAgentIdsBySquadId(1n)).length).to.eq(2);

      //
      console.log("admin1: ", admin1.address);
      console.log("admin2: ", admin2.address);

      await squadIns.connect(admin1).moveAgentToSquad(1n, 1n);
      await expect(
        squadIns.connect(admin2).moveAgentToSquad(1n, 1n)
      ).to.be.revertedWithCustomError(agentIns, "Unauthorized()");
      console.log("squad 1 owner: ", await squadIns.squadOwner(1n));
      console.log("agent 1 owner: ", await agentIns.ownerOf(1n));
    });
  });
});
