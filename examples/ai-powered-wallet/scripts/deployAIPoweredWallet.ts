import assert from "assert";
import { ethers, network, upgrades } from "hardhat";

async function deployAIPoweredWallet() {
  const config = network.config as any;
  const networkName = network.name.toUpperCase();

  const aiKernelAddress = config.aiKernelAddress;
  const promptScheduler = config.promptSchedulerAddress;

  assert.ok(
    aiKernelAddress,
    `Missing ${networkName}_AI_KERNEL_ADDRESS from environment variables!`
  );
  assert.ok(
    promptScheduler,
    `Missing ${networkName}_PROMPT_SCHEDULER_ADDRESS from environment variables!`
  );

  const contractFactory = await ethers.getContractFactory("AIPoweredWallet");
  const aiPoweredWallet = await contractFactory.deploy(
    aiKernelAddress,
    promptScheduler
  );
  await aiPoweredWallet.waitForDeployment();

  console.log(
    `${networkName}_AI_POWERED_WALLET_ADDRESS=${aiPoweredWallet.target}`
  );
}

deployAIPoweredWallet()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
