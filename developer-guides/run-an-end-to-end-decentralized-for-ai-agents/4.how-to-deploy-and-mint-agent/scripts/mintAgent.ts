import { URI } from "@smithy/types/dist-types/ts3.4/uri";
import { ethers, network } from "hardhat";
import { readFile } from "fs/promises";
import { AI721 } from "../typechain-types";
import { assert } from "console";
import path from "path";
import fs from "fs";
import axios from "axios";

const config = network.config as any;
const networkName = network.name.toUpperCase();
const filePath = path.join(
  __dirname,
  `../../../../decentralized-compute/worker-hub/env/local_contracts.json`
);

async function mintAgent() {
  const ownerAddress = (await ethers.getSigners())[0].address;
  const uri = process.env.AGENT_URI || "Eternal AI";

  const systemPromptPath = process.env.AGENT_SYSTEM_PROMPT_PATH as string;

  const { agent_contract_address: contractAddr } = await readExternalConfig();

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
  let id = 0;

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

    // Get agent ID
    id = receipt?.logs
      .filter(
        (log: any) =>
          log.topics[0] ===
          ethers.id("NewToken(uint256,string,bytes,uint256,address)")
      )
      .map((log: any) => {
        return parseInt(log.topics[1], 16);
      });

    if (id) {
      console.log("Agent minted successfully with ID: ", id);
      await saveAgentId(id[0].toString());
    } else {
      console.log("Check transaction history for more details.");
    }
  } catch (error) {
    console.error("Error minting agent:", error);
  }

  try {
    await createLocalAgent(
      ownerAddress,
      "AIAgent",
      contractAddr,
      id.toString(),
      systemPrompt.toString()
    );
  } catch (err) {
    console.log("Error creating local agent service: ", err);
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

async function saveAgentId(agentId: string) {
  // Read the JSON file
  const jsonData = fs.readFileSync(filePath, "utf-8");

  // Parse the JSON content
  const config = JSON.parse(jsonData);

  // Add the agentId to the contracts object
  config.agent_id = agentId;

  // Write the updated JSON back to the file
  fs.writeFileSync(filePath, JSON.stringify(config, null, 2), "utf-8");
}

async function createLocalAgent(
  creator: string,
  agentName: string,
  agentContractAddress: string,
  agentContractId: string,
  systemContent: string
): Promise<any> {
  try {
    const response = await axios.post(
      "http://localhost:8480/api/agent/create-local-agent",
      {
        creator,
        agent_name: agentName,
        agent_contract_address: agentContractAddress,
        agent_contract_id: agentContractId,
        system_content: systemContent,
      },
      {
        headers: {
          "Content-Type": "application/json",
        },
      }
    );
    return response.data;
  } catch (error) {
    if (axios.isAxiosError(error)) {
      console.error("Axios Error:", error.message, error.response?.data);
      // Handle specific error codes, e.g., 400, 500
      if (error.response?.status === 400) {
        // ...handle bad request
      }
    } else {
      console.error("An unexpected error occurred:", error);
    }
    throw error; // Re-throw the error to be caught by the caller if needed
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
