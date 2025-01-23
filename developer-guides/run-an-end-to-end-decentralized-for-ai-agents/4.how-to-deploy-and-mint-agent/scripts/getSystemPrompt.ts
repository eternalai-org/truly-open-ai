import { ethers } from "hardhat";
import { AI721 } from "../typechain-types";
import { assert } from "console";
import path from "path";
import fs from "fs";
const filePath = path.join(
  __dirname,
  `../../../../decentralized-compute/worker-hub/env/local_contracts.json`
);

async function getSystemPrompt(agentId: number) {
  const { agent_contract_address: contractAddr } = await readExternalConfig();

  assert(contractAddr, `Missing HARDHAT_AI721_ADDRESS!`);
  assert(agentId, `Missing param agentId!`);

  let id = BigInt(agentId);

  // Get contract instance
  try {
    const ai721ContractIns = (await getContractInstance(
      contractAddr,
      "AI721"
    )) as AI721;

    const sysPrompt = await ai721ContractIns.getAgentSystemPrompt(id);
    console.log("System Prompt: ", ethers.toUtf8String(sysPrompt[0]));
  } catch (err) {
    console.log("Error: ", err.message);
  }
}

async function readExternalConfig() {
  if (!fs.existsSync(filePath)) {
    throw new Error(`File not found: ${filePath}`);
  }

  const data = JSON.parse(fs.readFileSync(filePath, "utf8"));

  if (!data) {
    throw new Error(`No data found in ${filePath}`);
  }

  return data;
}

export async function getContractInstance(
  proxyAddress: string,
  contractName: string
) {
  const contractFact = await ethers.getContractFactory(contractName);
  const contractIns = contractFact.attach(proxyAddress);

  return contractIns;
}

async function main() {
  const agentId = (process.env.AGENT_ID || 0) as number;
  await getSystemPrompt(agentId);
}

// Execute
main().catch(console.error);
