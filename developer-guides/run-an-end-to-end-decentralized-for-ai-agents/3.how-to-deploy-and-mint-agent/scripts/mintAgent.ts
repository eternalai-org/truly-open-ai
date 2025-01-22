import { URI } from "./../node_modules/@smithy/types/dist-types/ts3.4/uri.d";
import { ethers, network } from "hardhat";
import { readFile } from "fs/promises";
import * as fs from "fs";
import * as path from "path";
import { AI721 } from "../typechain-types";
import { assert } from "console";

const config = network.config as any;
const networkName = network.name.toUpperCase();

async function mintAgent() {
  const ownerAddress =
    process.env.HARDHAT_AGENT_OWNER_ADDRESS || config.l2OwnerAddress;
  const uri = process.env.AGENT_URI || "Eternal AI";
  const systemPromptPath = process.env.AGENT_SYSTEM_PROMPT_PATH as string;
  const contractAddr = config.ai721Address || process.env.HARDHAT_AI721_ADDRESS;

  assert(ownerAddress, `Missing HARDHAT_AGENT_OWNER_ADDRESS!`);
  assert(uri, `Missing AGENT_URI!`);
  assert(systemPromptPath, `Missing AGENT_SYSTEM_PROMPT_PATH!`);
  assert(contractAddr, `Missing HARDHAT_AI721_ADDRESS!`);

  // Read system prompt from file
  let systemPrompt;
  try {
    systemPrompt = await readFile(systemPromptPath, "utf8");
  } catch (error) {
    throw new Error(`Failed to read system prompt file: ${error}`);
  }

  try {
    // Get contract instance
    const ai721ContractIns = (await getContractInstance(
      contractAddr,
      "AI721"
    )) as AI721;

    // Mint parameters
    const data = ethers.toUtf8Bytes(systemPrompt.toString()); // Convert string to bytes

    console.log("Mining agent...");
    const txMint = await ai721ContractIns.mint(ownerAddress, uri, data, 0);
    const receipt = await txMint.wait();
    console.log("Tx hash: ", receipt?.hash);
    console.log("Tx status: ", receipt?.status === 1 ? "Success" : "Failed");
  } catch (error) {
    console.error("Error minting agent:", error);
  }
}

export async function getContractInstance(
  proxyAddress: string,
  contractName: string
) {
  const contractFact = await ethers.getContractFactory(contractName);
  const contractIns = contractFact.attach(proxyAddress);

  return contractIns;
}

// Execute
mintAgent().catch(console.error);
