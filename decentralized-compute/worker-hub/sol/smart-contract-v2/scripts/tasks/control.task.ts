import { task } from "hardhat/config";
import { HardhatRuntimeEnvironment } from "hardhat/types";

task("pause", "Pause contract")
  .addParam("address", "The contract's address")
  .setAction(async (taskArgs: any, hre: HardhatRuntimeEnvironment) => {
    const { ethers } = hre;
    const { address } = taskArgs
    
    const contract = await ethers.getContractAt("IPausableUpgradeable", address);
    (await contract.pause()).wait();
    
    console.log(`Contract ${address} paused`);
  });
  
task("unpause", "Unpause contract")
  .addParam("address", "The contract's address")
  .setAction(async (taskArgs: any, hre: HardhatRuntimeEnvironment) => {
    const { ethers } = hre;
    const { address } = taskArgs
    
    const contract = await ethers.getContractAt("IPausableUpgradeable", address);
    (await contract.unpause()).wait();
    
    console.log(`Contract ${address} unpaused`);
  });
  