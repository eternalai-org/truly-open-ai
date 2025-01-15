import { assert, ethers } from "ethers";

// ABI for the mint function
const abi = [
  "function mint(address _to, string calldata _uri, bytes calldata _data, uint _fee) external payable returns (uint256)",
  "function mintPrice() external view returns (uint256)",
];

async function mintAgent() {
  // Configuration
  const RPC_URL = process.env.RPC_URL;
  const PRIVATE_KEY = process.env.PRIVATE_KEY;
  const CONTRACT_ADDRESS = process.env.CONTRACT_ADDRESS
    ? process.env.CONTRACT_ADDRESS
    : "0xaed016e060e2ffe3092916b1650fc558d62e1ccc";
  const AGENT_SYSTEM_PROMPT = process.env.AGENT_SYSTEM_PROMPT || "";
  const AGENT_URI = process.env.AGENT_URI || "";
  const AGENT_FEE = process.env.AGENT_FEE || "0";

  assert(RPC_URL, "Missing RPC_URL environment variable", "INVALID_ARGUMENT");
  assert(
    PRIVATE_KEY,
    "Missing PRIVATE_KEY environment variable",
    "INVALID_ARGUMENT"
  );
  assert(
    CONTRACT_ADDRESS,
    "Missing CONTRACT_ADDRESS environment variable",
    "INVALID_ARGUMENT"
  );
  assert(
    AGENT_SYSTEM_PROMPT,
    "Missing AGENT_SYSTEM_PROMPT environment variable",
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

  try {
    // Setup provider and signer
    const provider = new ethers.JsonRpcProvider(RPC_URL);
    const wallet = new ethers.Wallet(PRIVATE_KEY, provider);

    // Create contract instance
    const contract = new ethers.Contract(CONTRACT_ADDRESS, abi, wallet);

    // Mint parameters
    const to = wallet.address; // mint to self
    const uri = AGENT_URI.toString();
    const data = ethers.toUtf8Bytes(AGENT_SYSTEM_PROMPT.toString()); // Convert string to bytes
    const fee = ethers.parseEther(AGENT_FEE); // Set agent usage fee
    const mintPrice = await contract.mintPrice(); // Get mint price
    console.log("Mint price: ", mintPrice);

    // Call mint function
    const tx = await contract.mint(to, uri, data, fee, {
      value: mintPrice, // Send required mint price
    });
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

// Execute
mintAgent().catch(console.error);
