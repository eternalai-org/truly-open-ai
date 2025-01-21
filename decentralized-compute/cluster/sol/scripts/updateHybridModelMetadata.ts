import assert from "assert";
import { ethers, network, upgrades } from "hardhat";
import { HybridModel, ModelCollection, WorkerHub } from "../typechain-types";

async function updateHybridModelMetadata() {
  const config = network.config as any;
  const networkName = network.name.toUpperCase();
  const HybridModel = await ethers.getContractFactory("HybridModel");
  const WorkerHub = await ethers.getContractFactory("WorkerHub");
  const ModelCollection = await ethers.getContractFactory("ModelCollection");

  const collectionAddress = config.collectionAddress;
  assert.ok(
    collectionAddress,
    `Missing ${networkName}_COLLECTION_ADDRESS from environment variables!`
  );

  const workerHubAddress = config.workerHubAddress;
  assert.ok(
    workerHubAddress,
    `Missing ${networkName}_WORKERHUB_ADDRESS from environment variables!`
  );
  const tokenId = 500001;
  const minHardware = BigInt(1);
  const metadataObj = {
    version: 1,
    model_name: "HERMES [dev] quantized (fp8)",
    model_type: "image",
    model_url:
      "https://gateway.lighthouse.storage/ipfs/bafkreifm6m4fim2spgym7ev4g5j2twzamrms6nr6fnqaxh5u47fo2donyy",
    model_file_hash: "",
    min_hardware: 1,
    verifier_url: "",
    verifier_file_hash: "",
  };
  const metadata = JSON.stringify(metadataObj, null, "\t");
  console.log(metadata);

  const collection = ModelCollection.attach(
    config.collectionAddress
  ) as ModelCollection;
  await (await collection.updateTokenURI(tokenId, metadata)).wait();
  console.log("TokenURI updated");

  const modelAddress = await collection.modelAddressOf(tokenId);
  const hybridModel = HybridModel.attach(modelAddress) as HybridModel;
  await (await hybridModel.updateMetadata(metadata)).wait();
  console.log("Hybrid model metadata updated");

  const workerHub = WorkerHub.attach(workerHubAddress) as WorkerHub;
  const currentTier = (await workerHub.models(modelAddress)).tier;
  if (currentTier !== minHardware) {
    await (await workerHub.updateModelTier(modelAddress, minHardware)).wait();
    console.log(`Model tier updated (old=${currentTier}, new=${minHardware})`);
  }
}

updateHybridModelMetadata()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
