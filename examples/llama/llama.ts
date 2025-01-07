import { Provider } from "zksync-ethers";
import * as ethers from "ethers";
import { Address } from "zksync-ethers/src/types";

// load env file
import dotenv from "dotenv";
dotenv.config();

import * as WorkerHub from "./abi/WorkerHub.sol/WorkerHub.json";
import * as StakingHub from "./abi/StakingHub.sol/StakingHub.json";
import * as ISystemPromptManager from "./abi/interfaces/ISystemPromptManager.sol/ISystemPromptManager.json";

const PRIVATE_KEY = process.env.WALLET_PRIVATE_KEY || "";
const RPC_URL = process.env.RPC_URL || "";

if (!PRIVATE_KEY) throw "Private key not detected! Add it to the .env file!";
if (!RPC_URL) throw "Node endpoint not detected! Add it to the .env file!";

// Contract addresses
const HYBRID_MODEL_CONTRACT_ADDRESS =
  "0x9A35863BaF0df7623F34FCF3376CDBAC8581E76b";
const SYSTEM_PROMPT_MANAGER_CONTRACT_ADDRESS =
  "0xAed016e060e2fFE3092916b1650Fc558D62e1CCC";
const WORKER_HUB_CONTRACT_ADDRESS =
  "0xA1d2F74C345Ff1d97B8FC72E061903cD84C66F69";
const STAKING_HUB_CONTRACT_ADDRESS =
  "0x233198ce0679d3D8B3B346b9fe3DFD3298B991EB";

if (
  !HYBRID_MODEL_CONTRACT_ADDRESS ||
  !WORKER_HUB_CONTRACT_ADDRESS ||
  !STAKING_HUB_CONTRACT_ADDRESS ||
  !SYSTEM_PROMPT_MANAGER_CONTRACT_ADDRESS
)
  throw "Contract address not provided";

async function main() {
  console.log(
    `Running script to interact with contract ${WORKER_HUB_CONTRACT_ADDRESS} and ${SYSTEM_PROMPT_MANAGER_CONTRACT_ADDRESS}`
  );

  // Initialize the provider.
  // @ts-ignore
  const provider = new Provider(RPC_URL);
  const signer = new ethers.Wallet(PRIVATE_KEY, provider);

  // Initialize stakinghub contract instance
  const stakingHub = new ethers.Contract(
    STAKING_HUB_CONTRACT_ADDRESS,
    StakingHub.abi,
    signer
  );

  // Initialize workerhub contract instance
  const workerHub = new ethers.Contract(
    WORKER_HUB_CONTRACT_ADDRESS,
    WorkerHub.abi,
    signer
  );

  // Get model fee
  const model = await stakingHub.getModelInfo(HYBRID_MODEL_CONTRACT_ADDRESS);
  console.log(`The model fee is ${model.minimumFee}`);

  // send prompt
  const prompt = Buffer.from("GM, Llama. Nice meeting you on Base!", "utf8");

  // Initialize system prompt manager contract instance
  const systemPromptContract = new ethers.Contract(
    SYSTEM_PROMPT_MANAGER_CONTRACT_ADDRESS,
    ISystemPromptManager.abi,
    signer
  );
  const agentId = 5n; // dllama's agent id

  const tx = await systemPromptContract.infer(agentId, prompt, "");

  console.log(
    `Transaction to request prompt to system prompt manager is ${tx.hash}`
  );
  await tx.wait();

  const txReceipt = await provider.getTransactionReceipt(tx.hash);
  if (txReceipt != null && txReceipt.status == 1) {
    const inferId = getInferId(txReceipt);
    console.log(`infer id: ${inferId[0]}`);
    // Get result
    while (true) {
      const assignments = await workerHub.getInferenceInfo(inferId[0]);
      if (assignments.assignments.length == 0) {
        // Retry after 5s
        await delay(5000);
        continue;
      }
      break;
    }

    const assignmentIds = await workerHub.getAssignmentsByInference(inferId[0]);

    let res;
    for await (let assignmentId of assignmentIds) {
      const assignmentDetail = await workerHub.assignments(assignmentId);
      if (assignmentDetail.role == 2) {
        res = assignmentDetail.output;
        break;
      }
    }

    console.log(`result: ${Buffer.from(res.slice(2), "hex")}`);
  }
}

export function getInferId(receipt: ethers.TransactionReceipt): number[] {
  return receipt.logs
    .filter(
      (log) =>
        log.topics[0] ===
          ethers.id("NewInference(uint256,address,address,uint256,uint256)") &&
        isAddressEq(log.address, WORKER_HUB_CONTRACT_ADDRESS)
    )
    .map((log) => {
      return parseInt(log.topics[1], 16);
    });
}

function isAddressEq(a: Address, b: Address): boolean {
  return a.toLowerCase() === b.toLowerCase();
}

function delay(ms: number) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
