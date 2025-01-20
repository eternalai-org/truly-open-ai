import { assert, ethers } from "ethers";
import { readFile } from "fs/promises";
import * as fs from "fs";
import * as path from "path";

interface NetworkConfig {
  promptSchedulerAddress: string;
  gpuManagerAddress: string;
}

interface Config {
  [network: string]: NetworkConfig;
}

// ABI for the mint function
const agentAbi = [
  "function mint(address to,string calldata uri,bytes calldata data,uint256 fee,string calldata promptKey,address promptScheduler,uint32 modelId) external returns (uint256)",
  "function mintPrice() external view returns (uint256)",
];

const gpuManagerAbi = [
  "function getModelIds() external view returns (uint256[] memory)",
];

async function mintAgent() {
  // Get default config
  const NETWORK_NAME = process.env.NETWORK_NAME;
  assert(
    NETWORK_NAME,
    "Missing NETWORK_NAME environment variable",
    "INVALID_ARGUMENT"
  );

  const [df_promptSchedulerAddress, df_gpuManagerAddress] =
    getNetworkAddresses(NETWORK_NAME);
  // Configuration
  const RPC_URL = process.env.RPC_URL;
  const PRIVATE_KEY = process.env.PRIVATE_KEY;
  const AGENT_CONTRACT_ADDRESS = process.env.AGENT_CONTRACT_ADDRESS;
  const AGENT_SYSTEM_PROMPT_PATH = process.env.AGENT_SYSTEM_PROMPT_PATH;
  const AGENT_URI = process.env.AGENT_URI || "EternalAIAgent";
  const AGENT_FEE = process.env.AGENT_FEE || "0";
  const AGENT_PROMPT_KEY = process.env.AGENT_FEE || "EternalAIAgent";
  const PROMPT_SCHEDULER_ADDRESS = process.env.PROMPT_SCHEDULER_ADDRESS
    ? process.env.PROMPT_SCHEDULER_ADDRESS
    : df_promptSchedulerAddress;
  const GPU_MANAGER_ADDRESS = process.env.GPU_MANAGER_ADDRESS
    ? process.env.GPU_MANAGER_ADDRESS
    : df_gpuManagerAddress;
  let MODEL_ID = process.env.MODEL_ID;

  assert(RPC_URL, "Missing RPC_URL environment variable", "INVALID_ARGUMENT");
  assert(
    PRIVATE_KEY,
    "Missing PRIVATE_KEY environment variable",
    "INVALID_ARGUMENT"
  );
  assert(
    AGENT_CONTRACT_ADDRESS,
    "Missing AGENT_CONTRACT_ADDRESS environment variable",
    "INVALID_ARGUMENT"
  );
  assert(
    AGENT_SYSTEM_PROMPT_PATH,
    "Missing AGENT_SYSTEM_PROMPT_PATH environment variable",
    "INVALID_ARGUMENT"
  );
  assert(
    AGENT_URI,
    "Missing AGENT_URI environment variable",
    "INVALID_ARGUMENT"
  );
  assert(
    AGENT_FEE,
    "Missing AGENT_FEE environment variable",
    "INVALID_ARGUMENT"
  );
  assert(
    AGENT_PROMPT_KEY,
    "Missing AGENT_PROMPT_KEY environment variable",
    "INVALID_ARGUMENT"
  );
  assert(
    PROMPT_SCHEDULER_ADDRESS,
    "Missing PROMPT_SCHEDULER_ADDRESS environment variable",
    "INVALID_ARGUMENT"
  );

  // Setup provider and signer
  const provider = new ethers.JsonRpcProvider(RPC_URL);
  const wallet = new ethers.Wallet(PRIVATE_KEY, provider);

  if (!MODEL_ID) {
    assert(
      GPU_MANAGER_ADDRESS,
      "Missing GPU_MANAGER_ADDRESS environment variable",
      "INVALID_ARGUMENT"
    );

    const gpuManager = new ethers.Contract(
      GPU_MANAGER_ADDRESS,
      gpuManagerAbi,
      wallet
    );

    const modelIds = await gpuManager.getNOMiner();
    console.log("Available models:", modelIds);
    if (modelIds.length === 0) {
      throw new Error("No models available");
    }
    MODEL_ID = modelIds[0];
  }

  // Read system prompt from file
  let systemPrompt;
  try {
    systemPrompt = await readFile(AGENT_SYSTEM_PROMPT_PATH, "utf8");
  } catch (error) {
    throw new Error(`Failed to read system prompt file: ${error}`);
  }

  try {
    // Create contract instance
    const contract = new ethers.Contract(
      AGENT_CONTRACT_ADDRESS,
      agentAbi,
      wallet
    );

    // Mint parameters
    const to = wallet.address; // mint to self
    const uri = AGENT_URI.toString();
    const data = ethers.toUtf8Bytes(systemPrompt.toString()); // Convert string to bytes
    const fee = ethers.parseEther(AGENT_FEE); // Set agent usage fee
    const mintPrice = await contract.mintPrice(); // Get mint price
    const promptKey = AGENT_PROMPT_KEY.toString();
    const promptSchedulerAddress = PROMPT_SCHEDULER_ADDRESS.toString();
    const modelId = MODEL_ID;

    console.log("Mint price: ", mintPrice);

    // Call mint function
    const tx = await contract.mint(
      to,
      uri,
      data,
      fee,
      promptKey,
      promptSchedulerAddress,
      modelId,
      {
        value: mintPrice, // Send required mint price
      }
    );
    const receipt = await tx.wait();

    if (receipt?.status === 1) {
      console.log("Minting transaction sent:", receipt.hash);
      console.log("Transaction confirmed in block:", receipt.blockNumber);

      // Get minted token ID from events
      const event = receipt.logs?.find((e: ethers.Log) => {
        return (
          e.topics[0] ===
          ethers.id("NewToken(uint256,string,bytes,uint256,address)")
        );
      });
      if (event) {
        console.log("Minted Agent ID:", Number(event.topics[1]));
      }
    } else {
      console.error("Minting transaction failed:", receipt);
    }
  } catch (error) {
    console.error("Error minting agent:", error);
  }
}

function getNetworkAddresses(networkName: string): [string, string] {
  try {
    // Read config file
    const configPath = path.resolve(__dirname, "./config.json");
    const configData = fs.readFileSync(configPath, "utf8");
    const config: Config = JSON.parse(configData);

    // Get network specific config
    const networkConfig = config[networkName];
    if (!networkConfig) {
      throw new Error(`Network ${networkName} not found in config`);
    }

    return [
      networkConfig.promptSchedulerAddress,
      networkConfig.gpuManagerAddress,
    ];
  } catch (error) {
    throw new Error(`Failed to load config: ${error}`);
  }
}

// Execute
mintAgent().catch(console.error);
