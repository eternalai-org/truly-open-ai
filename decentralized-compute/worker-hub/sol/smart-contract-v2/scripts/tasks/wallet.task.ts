import { task } from "hardhat/config";
import { HardhatRuntimeEnvironment } from "hardhat/types";

task("balance", "Prints an account's balance")
  .addOptionalParam("account", "The account's address", "")
  .setAction(async (taskArgs: any, hre: HardhatRuntimeEnvironment) => {
    const { ethers } = hre;
    const accs = await ethers.getSigners();
    const account = taskArgs.account || accs[0].address;
    const balance = await ethers.provider.getBalance(account);

    console.log(ethers.formatEther(balance), "ETH");
  });
  