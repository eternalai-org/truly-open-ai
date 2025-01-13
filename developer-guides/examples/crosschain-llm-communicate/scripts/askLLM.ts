import { ethers, network } from 'hardhat';
import { HardhatUserConfig } from 'hardhat/config';
import { Address, Questioner, IPromptScheduler } from '../typechain';
import assert from 'assert';
import { Wallet } from 'ethers';
const jsesc = require('jsesc');

import lzConfig from '../mainnet.layerzero.config';
import hhConfig from '../hardhat.config';

async function askLocal() {
    const srcNetworkName = getNetworkName('SRC');
    const srcAddresses = getNetworkConfig(srcNetworkName);

    const srcQuestioner = (await getSourceContract(
        'Questioner',
        srcAddresses.questioner
    )) as Questioner;

    const srcPromptScheduler = (await getSourceContract(
        'IPromptScheduler',
        srcAddresses.promptScheduler
    )) as IPromptScheduler;

    const msgLocal = 'What is ETH?';

    console.log(`Source Network: ${srcNetworkName}`);
    console.log(`Ask local question: "${msgLocal}"`);

    const srcAskLocalTx = await srcQuestioner.askLocal(msgLocal, {
        gasLimit: 15000000,
    });
    const srcAskLocalReceipt = await srcAskLocalTx.wait();
    console.log('+ TxHash:', srcAskLocalReceipt.transactionHash);
    console.log('+ Status:', srcAskLocalReceipt.status);

    let srcLocalInferId = 0;
    let srcLocalInferResult = '';
    if (srcAskLocalReceipt?.status == 1) {
        // Get inference ID
        srcLocalInferId = getInferId(
            srcAskLocalReceipt,
            srcAddresses.promptScheduler
        )[0];
        console.log('+ Inference ID: ', srcLocalInferId);
        console.log('+ Wait for inference result...');

        // Wait for inference result
        while (true) {
            await sleep(10000);
            try {
                srcLocalInferResult =
                    await srcQuestioner.fetchInferenceResult(srcLocalInferId);
                break;
            } catch (e: any) {
                if (e.message.includes('Wait for inference')) {
                    console.log('+ Waiting for inference result...');
                } else console.log(e.message);
                continue;
            }
        }
    }
    const strResult = ethers.utils.toUtf8String(srcLocalInferResult);
    console.log(
        `+ Tx submit: ${await getTxSubmitSolution(srcPromptScheduler, srcLocalInferId)}`
    );
    console.log(`+ Inference result: ${strResult} \n`);

    // Process cross-chain message
    const dstNetworkName = getNetworkName('DST');
    const dstAddresses = getNetworkConfig(dstNetworkName);

    const dstPromptScheduler = (await getDestinationContract(
        'IPromptScheduler',
        dstAddresses.promptScheduler
    )) as IPromptScheduler;

    const dstQuestioner = (await getDestinationContract(
        'Questioner',
        dstAddresses.questioner
    )) as Questioner;

    const crossChainMsg = `I have asked LLM "What is ETH?", it's response is "${strResult}", what do you think? Is there anything missing from this response?`;
    const escapeMsg = jsesc(crossChainMsg, {
        quotes: 'double',
    });
    // console.log(`\nEscape message: ${escapeMsg}\n`);
    console.log(`Ask cross-chain question: "${crossChainMsg}"`);

    const dstEid = 30312;
    const options = '0x000301001101000000000000000000000000003d0900';
    const payInLzToken = false;

    // Create tx in source chain
    const fee = await srcQuestioner.quote(
        dstEid,
        escapeMsg,
        options,
        payInLzToken
    );
    console.log('+ Estimate fee:', fee[0].toString());

    const srcAskL0Tx = await srcQuestioner.askByL0(dstEid, escapeMsg, options, {
        value: fee[0],
        gasLimit: 15000000,
    });
    const srcAskL0Receipt = await srcAskL0Tx.wait();
    console.log('+ TxHash:', srcAskL0Receipt.transactionHash);
    console.log('+ Status:', srcAskL0Receipt.status);

    // Listen on destination chain
    console.log(`\nDestination Network: ${dstNetworkName}`);

    let events;
    let dstInferId = 0;
    const startBlock = await dstQuestioner.provider.getBlockNumber();
    console.log('+ Start block:', startBlock);

    while (true) {
        const filter = dstPromptScheduler.filters['NewInference'](
            null,
            null,
            null
        );
        events = await dstPromptScheduler.queryFilter(filter, startBlock);
        if (events.length > 0) {
            for await (const event of events) {
                dstInferId = Number(event.args[0]);
                let dstQuestion = ethers.utils.toUtf8String(
                    (await dstPromptScheduler.getInferenceInfo(dstInferId))[9]
                );
                if (dstQuestion.includes('What is ETH?')) break;
            }
            if (dstInferId > 0) break;
        }
        console.log('Waiting for "NewInference" event...');
        await sleep(5000);
    }

    console.log('+ Infer ID: ', dstInferId);

    let dstInferResult;
    if (srcAskL0Receipt?.status == 1) {
        // Get inference ID
        console.log('+ Wait for inference result...');

        // Wait for inference result
        while (true) {
            await sleep(10000);
            try {
                dstInferResult =
                    await dstQuestioner.fetchInferenceResult(dstInferId);
                break;
            } catch (e: any) {
                if (e.message.includes('Wait for inference')) {
                    console.log('+ Waiting for inference result...');
                } else console.log(e.message);
                continue;
            }
        }
    }
    console.log(
        `+ Tx submit: ${await getTxSubmitSolution(dstPromptScheduler, dstInferId)}`
    );
    console.log(
        '+ Inference result: ',
        ethers.utils.toUtf8String(dstInferResult)
    );
}

function getNetworkName(networkType: 'SRC' | 'DST'): string {
    const networkName = process.env[`${networkType}_NETWORK_NAME`];
    if (!networkName) {
        throw new Error(`Missing ${networkType}_NETWORK environment variable`);
    }

    return networkName;
}

async function getContract(
    contractName: string,
    address: string,
    networkType: 'SRC' | 'DST'
) {
    const networkName = getNetworkName(networkType);

    const networks = hhConfig.networks as any;
    if (!networks || !networks[networkName]) {
        throw new Error(`Network configuration not found for ${networkName}`);
    }

    const rpcUrl = networks[networkName].url;
    if (!rpcUrl) {
        throw new Error(`Missing RPC URL for network ${networkName}`);
    }

    const privateKey = networks[networkName].privateKey;
    if (!privateKey) {
        throw new Error(
            `Missing ${networkName.toUpperCase()}_PRIVATEKEY environment variable`
        );
    }

    const wallet = await createWallet(privateKey, rpcUrl);
    return await ethers.getContractAt(contractName, address, wallet);
}

function getNetworkConfig(networkName: string) {
    const networks = hhConfig.networks;
    if (!networks || !networks[networkName]) {
        throw new Error(`Network configuration not found for ${networkName}`);
    }

    const config = networks[networkName] as any;

    const addresses = {
        questioner: config.questioner,
        hybridModel: config.hybridModel,
        promptScheduler: config.promptScheduler,
        modelName: config.modelName,
    };
    if (!addresses) {
        throw new Error(
            `No contract addresses found for network ${networkName}`
        );
    }

    return addresses;
}

// Helper functions to make code more readable
async function getSourceContract(contractName: string, address: string) {
    return getContract(contractName, address, 'SRC');
}

async function getDestinationContract(contractName: string, address: string) {
    return getContract(contractName, address, 'DST');
}

export function getInferId(
    receipt: ethers.TransactionReceipt,
    promptSchedulerAddress: string
): number[] {
    return receipt.logs
        .filter((log: any) => {
            // console.log('log:', log);
            return (
                log.topics[0] ===
                    ethers.utils.id(
                        'NewInference(uint256,address,address,uint256,uint256)'
                    ) && isAddressEq(log.address, promptSchedulerAddress)
            );
        })
        .map((log: any) => {
            return parseInt(log.topics[1], 16);
        });
}

async function getTxSubmitSolution(
    promptScheduler: IPromptScheduler,
    inferId: number
) {
    const filter = promptScheduler.filters['SolutionSubmission'](null, inferId);
    const fromBlk = await promptScheduler.provider.getBlockNumber();
    const events = await promptScheduler.queryFilter(filter, fromBlk - 100);
    assert(events.length > 0, 'No SubmitSolution event found');

    for await (const event of events) {
        if (Number(event.args[1]) === inferId) {
            return event.transactionHash;
        }
    }
}

async function createWallet(
    privateKey: string,
    rpcUrl: string
): Promise<Wallet> {
    const provider = new ethers.providers.JsonRpcProvider(rpcUrl);
    const wallet = new Wallet(privateKey, provider);

    return wallet;
}

function isAddressEq(a: Address, b: Address): boolean {
    return a.toString().toLowerCase() === b.toString().toLowerCase();
}

async function sleep(ms: number): Promise<void> {
    return new Promise((resolve) => setTimeout(resolve, ms));
}

askLocal()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
