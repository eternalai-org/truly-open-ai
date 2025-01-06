import fs from "fs";

import hre, { ethers } from "hardhat";
import { BaseContract, ContractReceipt, ContractTransaction, Signer } from "ethers";
import { ProxyAdmin } from "../typechain";

const res = {
    getDeploySigner: async (): Promise<Signer> => {
        const ethersSigners = await Promise.all(await ethers.getSigners());
        return ethersSigners[0];
    },
    waitForDeploy: async (contract: BaseContract, note: string = ''): Promise<BaseContract> => {
        var tx: ContractTransaction = contract.deployTransaction
        console.log(note, 'deploy contract', contract.address, 'at', tx.hash, 'waiting...')
        let r = await tx.wait(1)
        console.log(note, 'deploy contract', contract.address, 'at', tx.hash, 'confirmed', 'gasUsed', r.gasUsed.toNumber())
        return contract
    },
    waitForTx: async (tx: ContractTransaction, note: string = ''): Promise<ContractReceipt> => {
        console.log(note, 'contract call method at', tx.hash, 'waiting...')
        let r = await tx.wait(1)
        console.log(note, 'contract call method at', tx.hash, 'confirmed', r.gasUsed.toNumber())
        return r
    },
    notWaitForTx: async (tx: ContractTransaction, note: string = '') => {
        console.log(note, 'contract call method at', tx.hash, 'not wait')
        return tx
    },
    tryWaitForTx: async (tx: ContractTransaction, note: string = '') => {
        console.log(note, 'contract call method at', tx.hash, 'waiting...')
        let r: ContractReceipt
        try {
            r = await tx.wait(1)
        } catch (ex) {
            console.log(note, 'contract call method at', tx.hash, 'error', ex)
            return
        }
        console.log(note, 'contract call method at', tx.hash, 'confirmed', 'gasUsed', r.gasUsed.toNumber())
    },
    sleep: (ms: number) => {
        return new Promise(resolve => setTimeout(resolve, ms));
    },
    getGasPrice: async () => {
        return (await ethers.provider.getGasPrice()).mul(12).div(10);
    },
    verifyContract: async (deployData: DeployData, network: string, address: string, constructorArguments: any, libraries: any, contract: string) => {
        if (network != 'local') {
            var verified = deployData.verifiedContracts[address]
            if (verified == undefined || verified == false) {
                try {
                    await hre.run("verify:verify", {
                        address: address,
                        constructorArguments: constructorArguments,
                        libraries: libraries,
                        contract: contract,
                    })
                } catch (ex) {
                    var err = '' + ex
                    if (!err.includes('Already Verified')) {
                        throw ex
                    }
                    console.log('Already verified contract address on Etherscan.')
                    console.log('https://testnet.arbiscan.io//address/' + address + '#code')
                }
                let env = 'testnet'
                if (network == 'arbitrum') {
                    env = 'mainnet'
                }
                deployData.verifiedContracts[address] = true
                let fileName = process.cwd() + '/deploy/' + network + '/deployed.json';
                await fs.writeFileSync(fileName, JSON.stringify(deployData, null, 4))
            }
        }
    },
    upgradeContract: async (proxyAdmin: ProxyAdmin, address: string, implAddress: string) => {
        if ((await proxyAdmin.getProxyImplementation(address)) != implAddress) {
            var tx = await proxyAdmin.upgrade(address, implAddress, {
                nonce: (await ethers.provider.getTransactionCount((await ethers.getSigners())[0].address)),
            })
            console.log('proxyAdmin.upgrade at', address, implAddress, tx.hash, 'waiting...')
            let r = await tx.wait(1)
            console.log('proxyAdmin.upgrade at', address, implAddress, tx.hash, 'confirmed', 'gasUsed', r.gasUsed.toNumber())
        }
    },

    loadDB: async (network: string) => {
        let fileName = process.cwd() + '/deploy/' + network + '/deployed.json';
        let deployData: DeployData;
        if (!(fs.existsSync(fileName))) {
            throw 'deployed file is not existsed'
        }
        let dataText = fs.readFileSync(fileName)
        deployData = JSON.parse(dataText.toString())
        return deployData;
    },

    saveDB: async (network: string, deployData: DeployData) => {
        let fileName = process.cwd() + '/deploy/' + network + '/deployed.json';
        if (!(fs.existsSync(fileName))) {
            throw 'deployed file is not existsed'
        }
        fs.writeFileSync(fileName, JSON.stringify(deployData, null, 4))
        return deployData;
    }
};

export default res;