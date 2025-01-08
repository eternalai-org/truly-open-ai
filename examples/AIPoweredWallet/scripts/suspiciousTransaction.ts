import { ethers, network } from "hardhat";
import { AIPoweredWallet } from "../typechain-types";
import { Address } from "hardhat-deploy/dist/types";
import { assert } from "chai";
import { BaseWallet, JsonRpcProvider, SigningKey } from "ethers";

const config = network.config as any;

async function suspiciousTransaction() {
  const networkName = network.name.toUpperCase();
  let privateKey = config.senderKey;
  let rpcUrl = config.url;
  let receiverAddress = config.receiverAddress;
  const transferredAmt = config.transferredAmount as bigint;
  const aiPoweredWalletAddress = config.aiPoweredWallet;

  assert(
    aiPoweredWalletAddress,
    `Missing ${networkName}_AI_POWERED_WALLET_ADDRESS from environment variables!`
  );
  assert(
    receiverAddress,
    `Missing ${networkName}_RECEIVER_ADDRESS from environment variables!`
  );
  assert(
    transferredAmt > 0n,
    `Missing ${networkName}_TRANSFERRED_AMOUNT from environment variables!`
  );
  assert(
    privateKey,
    `Missing ${networkName}_PRIVATE_KEY from environment variables!`
  );
  assert(rpcUrl, `Missing ${networkName}_RPC_URL from environment variables!`);

  const sender = await createWallet(privateKey, rpcUrl);

  const contractFactory = await ethers.getContractFactory("AIPoweredWallet");
  const aiPoweredWallet = contractFactory.attach(
    aiPoweredWalletAddress
  ) as AIPoweredWallet;

  console.log(
    `Check whether the transaction history is suspicious when sending ${ethers.formatEther(
      transferredAmt
    )} ETH from ${sender.address} to ${receiverAddress}...`
  );
  const txCheck = await aiPoweredWallet
    .connect(sender)
    .suspiciousTransaction(receiverAddress, transferredAmt);
  const receipt = await txCheck.wait();

  console.log("Tx hash: ", receipt?.hash);
  console.log("Tx status: ", receipt?.status == 1 ? "Success" : "Failed");

  let inferResult;
  let inferId = 0;
  if (receipt?.status == 1) {
    // Get inference ID
    inferId = getInferId(receipt)[0];
    console.log("Inference ID: ", inferId);
    console.log("Wait for inference result...");

    // Wait for inference result
    while (true) {
      await sleep(30000);
      try {
        inferResult = await aiPoweredWallet
          .connect(sender)
          .fetchInferenceResult(inferId);
        break;
      } catch (e: any) {
        // console.log(e.message.split(": ")[1].split(" ()[0]);
        console.log(e.message);
        continue;
      }
    }
  }

  console.log("Inference result: ", ethers.toUtf8String(inferResult));
  console.log("Send ETH after checking the transaction history...");

  try {
    const txSend = await aiPoweredWallet.connect(sender).send(inferId, {
      value: transferredAmt,
    });
    const receiptSend = await txSend.wait();
    console.log("Tx hash: ", receiptSend?.hash);
    console.log("Tx status: ", receiptSend?.status == 1 ? "Success" : "Failed");
  } catch (e: any) {
    console.log(e.message);
  }
}

export function getInferId(receipt: ethers.TransactionReceipt): number[] {
  return receipt.logs
    .filter(
      (log: any) =>
        log.topics[0] ===
          ethers.id("NewInference(uint256,address,address,uint256,uint256)") &&
        isAddressEq(log.address, config.promptSchedulerAddress)
    )
    .map((log: any) => {
      return parseInt(log.topics[1], 16);
    });
}

async function createWallet(
  privateKey: string,
  rpcUrl: string
): Promise<BaseWallet> {
  const wallet = new BaseWallet(
    new SigningKey(privateKey),
    new JsonRpcProvider(rpcUrl)
  );

  return wallet;
}

function isAddressEq(a: Address, b: Address): boolean {
  return a.toLowerCase() === b.toLowerCase();
}

async function sleep(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

suspiciousTransaction()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
