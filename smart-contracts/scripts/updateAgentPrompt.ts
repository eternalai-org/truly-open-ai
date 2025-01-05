import assert from "assert";
import { version } from "chai";
import { ethers, network, upgrades } from "hardhat";
import { SystemPromptManager } from "../typechain-types";

async function updateAgentPrompt() {
  const config = network.config as any;
  const networkName = network.name.toUpperCase();

  const systemPromptManagerAddress = config.systemPromptManagerAddress;
  assert.ok(
    systemPromptManagerAddress,
    `Missing ${networkName}_SYSTEM_PROMPT_MANAGER_ADDRESS from environment variables!`
  );

  // get the contract instance
  const Fact = await ethers.getContractFactory("SystemPromptManager");
  const ins = Fact.attach(systemPromptManagerAddress) as SystemPromptManager;

  // TODO: @mr 6789 change the link
  const linkPrompt =
    "ipfs://bafkreide4kf4se2atgdi3kjie5eigvvr3wnkyolitbrj6cuj3sfzfyowui";

  const data = ethers.toUtf8Bytes(linkPrompt);
  const agentId = 1n;
  const promptIndex = 0n;
  const txUpdatePrompt = await ins.updateAgentData(agentId, data, promptIndex);
  const resUpdatePrompt = await txUpdatePrompt.wait();
  console.log(`Update Prompt tx hash: ${resUpdatePrompt?.hash}`);
  console.log(`Update Prompt status: ${resUpdatePrompt?.status}`);

  // get the uri from system prompt manager contract
  const agentPrompt = await ins.getAgentSystemPrompt(agentId);
  console.log("=====================================");
  console.log("New agent prompt: ", agentPrompt[0]);
  console.log(
    "New agent prompt (string): ",
    ethers.toUtf8String(agentPrompt[0])
  );
}

updateAgentPrompt()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
