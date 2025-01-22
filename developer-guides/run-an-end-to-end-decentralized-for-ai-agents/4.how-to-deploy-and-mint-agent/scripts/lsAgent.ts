import { exec } from "child_process";
import { ethers, network } from "hardhat";

async function listLocalAgents(creatorAddress: string): Promise<string> {
  const url = `http://localhost:8480/api/agent/list-local-agent?creator=${creatorAddress}`;
  const command = `curl --location '${url}'`;

  return new Promise((resolve, reject) => {
    exec(command, (error, stdout, stderr) => {
      if (error) {
        reject(error);
        return;
      }
      if (stderr) {
        console.error(`stderr: ${stderr}`); // Log potential warnings
      }
      resolve(stdout);
    });
  });
}

async function main() {
  const ownerAddress = (await ethers.getSigners())[0].address;

  return await listLocalAgents(ownerAddress);
}

main()
  .then((result) => {
    console.log(result);
    process.exit(0);
  })
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
