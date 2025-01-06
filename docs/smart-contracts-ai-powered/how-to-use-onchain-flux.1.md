---
hidden: true
---

# How to use onchain FLUX.1

In this code example, we will demonstrate how to interact with FLUX model (for image generation) deployed on Imagine chain.

### Step 1: Set up <a href="#step-1-set-up-project" id="step-1-set-up-project"></a>

Since the Imagine chain is deployed on SHARD\_AI using ZKSync's ZK Stack, we need to set up our project following [the tutorial](https://docs.zksync.io/build/tooling/hardhat/guides/getting-started#project-setup).

### Step 2: Define contract interfaces <a href="#step-2-define-contract-interfaces" id="step-2-define-contract-interfaces"></a>

**IHybridModel.sol**

```solidity
// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IHybridModel {
   function infer(bytes calldata _data) external payable returns (uint256 referenceId);
}
```

**IWorkerHub.sol**

```solidity
// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

struct Model {
  uint256 minimumFee;
  uint32 tier;
}

struct Inference {
  uint256[] assignments;
  bytes input;
  uint256 value; // this value is calculated by msg.value - feeL2 - feeTreasury
  uint256 feeL2;
  uint256 feeTreasury;
  address modelAddress;
  uint40 submitTimeout; // limit time to capture the miner role and submit the solution
  uint40 commitTimeout;
  uint40 revealTimeout;
  uint8 status;
  address creator;
  address processedMiner;
  address referrer;
}

enum AssignmentRole {
   Nil,
   Validating,
   Mining
}

enum Vote {
   Nil,
   Disapproval,
   Approval
}

struct Assignment {
   uint256 inferenceId;
   bytes32 commitment;
   bytes32 digest; // keccak256(output)
   uint40 revealNonce;
   address worker;
   AssignmentRole role;
   Vote vote;
   bytes output;
}

interface IWorkerHub {
  function models(address _model) external view returns (Model memory);
  function getInferenceInfo(uint _inferId) external view returns(Inference memory);
  function getAssignmentByInferenceId(
       uint256 _inferId
  ) external view returns (Assignment[] memory);
  function assignments(uint _assignId) external view returns (Assignment memory);
}
```

### Step 3: Compile the interfaces

```
npx hardhat compile
```

### Step 4: Interact with FLUX contract <a href="#step-4-interact-with-llama-contract" id="step-4-interact-with-llama-contract"></a>

Create `image-generation.ts` script for creating a transaction to the FLUX contract to send a prompt to the FLUX model and get response (a image) from FLUX miners.

```typescript
import { Provider } from "zksync-ethers";
import * as ethers from "ethers";
import { HardhatRuntimeEnvironment } from "hardhat/types";
import {Address, DeploymentInfo} from "zksync-ethers/src/types";

// load env file
import dotenv from "dotenv";
dotenv.config();

import * as IHybridModel from "../artifacts-zk/contracts/IHybridModel.sol/IHybridModel.json";
import * as IWorkerHub from "../artifacts-zk/contracts/IWorkerHub.sol/IWorkerHub.json";

const PRIVATE_KEY = process.env.WALLET_PRIVATE_KEY || "";
const RPC_URL = process.env.RPC_URL || "";
if (!PRIVATE_KEY) throw "Private key not detected! Add it to the .env file!";
if (!RPC_URL) throw "Node endpoint not detected! Add it to the .env file!";

const HYBRID_MODEL_CONTRACT_ADDRESS = "0x9874732a8699FcA824A9A7d948f6bcD30a141238";
const WORK_HUB_CONTRACT_ADDRESS = "0x430583bDFf80c5BE1536Ed82f9c8090eEF68e2F6";
if (!HYBRID_MODEL_CONTRACT_ADDRESS || !WORK_HUB_CONTRACT_ADDRESS) throw "Contract address not provided";

export default async function (hre: HardhatRuntimeEnvironment) {
  console.log(`Running script to interact with contract ${HYBRID_MODEL_CONTRACT_ADDRESS} and ${WORK_HUB_CONTRACT_ADDRESS}`);

  // Initialize the provider.
  // @ts-ignore
  const provider = new Provider(RPC_URL);
  const signer = new ethers.Wallet(PRIVATE_KEY, provider);

  // Initialise workerhub contract instance
  const workerHub = new ethers.Contract(WORK_HUB_CONTRACT_ADDRESS, IWorkerHub.abi, signer);

  // Get model fee
  const model = await workerHub.models(HYBRID_MODEL_CONTRACT_ADDRESS);
  console.log(`The model fee is ${model.minimumFee}`);

  // send prompt
  const newMessage = Buffer.from("Say hello people!", 'utf8');

  // Initialise hybrid model contract instance
  const hybridContract = new ethers.Contract(HYBRID_MODEL_CONTRACT_ADDRESS, IHybridModel.abi, signer);
  const tx = await hybridContract.infer(newMessage, {value: model.minimumFee});

  console.log(`Transaction to request prompt to model is ${tx.hash}`);
  await tx.wait();

  const txReceipt = await provider.getTransactionReceipt(tx.hash);
  if (txReceipt != null && txReceipt.status == 1) {
    const inferId = getInferId(txReceipt);
    console.log(`infer id: ${inferId[0]}`);
    // get res
    while (true) {
      const assignments = await workerHub.getInferenceInfo(inferId[0]);
      if (assignments.assignments.length == 0) {
        // retry after 5s
        await delay(5000);
        continue;
      }
      break;
    }
    const assignmentsDetail = await workerHub.getAssignmentByInferenceId(inferId[0]);
    let res;
    for (let i = 0; i < assignmentsDetail.length; i++) {
      if (assignmentsDetail[i].role == 2) {
        res = assignmentsDetail[i].output;
        break;
      }
    }
    console.log(`result: ${Buffer.from(res, 'hex')}`);
  }
}

export function getInferId(
  receipt: ethers.TransactionReceipt
): number[] {
return (
    receipt.logs
        .filter(
            log =>
                log.topics[0] ===
                ethers.id('NewInference(uint256,address,address,uint256)') &&
                isAddressEq(log.address, WORK_HUB_CONTRACT_ADDRESS)
        )
        .map(log => {
          return parseInt(log.topics[1], 16);
        })
    );
}

function isAddressEq(a: Address, b: Address): boolean {
  return a.toLowerCase() === b.toLowerCase();
}

function delay(ms: number) {
  return new Promise( resolve => setTimeout(resolve, ms) );
}
```

### Step 5: Run <a href="#step-5-run" id="step-5-run"></a>

```bash

WALLET_PRIVATE_KEY=private-key RPC_URL=rpc-url yarn hardhat deploy-zksync --script image-generation.ts
```

Note: to run the command, please replace \`private-key\` and 'rpc-url' with your private key and RPC url respectively.
