import hre from "hardhat";
import { ethers } from "ethers";
import fs from 'fs';
import * as FunctionalModelArtifact from '../artifacts/contracts/FunctionalModel.sol/FunctionalModel.json';
import dotenv from 'dotenv';
import { fromFloat } from './libraries/utils';
import { getLayerConfigNew, uploadModelWeights } from './libraries/modelLib';
dotenv.config();

async function deployFunctionalModel(configs: any, weight_b64: string, weightPerTx: number, signer: ethers.Wallet): Promise<ethers.Contract> {
    const { layerConfigs, totalWeights } = getLayerConfigNew(configs.layers);

    let weightsFlat: ethers.BigNumber[] = [];
    if (weight_b64 !== "") {
        const temp = Buffer.from(weight_b64, 'base64');
        const floats = new Float32Array(new Uint8Array(temp).buffer);
        for (let i = 0; i < floats.length; i++) {
            weightsFlat.push(fromFloat(floats[i]));
        }
    } else {
        for(let i = 0; i < totalWeights; ++i) {
            weightsFlat.push(fromFloat(0));
        }
    }

    // deploy a FunctionalModel contract
    // FunctionalModel contract is too big (larger than 49152 bytes) to be deployed with ContractFactory
    const FunctionalFac = new ethers.ContractFactory(FunctionalModelArtifact.abi, FunctionalModelArtifact.bytecode, signer);    
    const functionalImpl = await FunctionalFac.deploy();
    await functionalImpl.deployed();

    console.log("Deploying FunctionalModel contract...");
    const functional = FunctionalFac.attach(functionalImpl.address);
    console.log(`Contract FunctionalModel has been deployed to address ${functional.address}`);
    
    console.log("Setting model...");
    const setWeightTx = await functional.constructModel(layerConfigs);
    const rc = await setWeightTx.wait();
    console.log(`tx: ${setWeightTx.hash}, gas used: ${rc.gasUsed}`);

    console.log("Uploading weights...");
    await uploadModelWeights(functional, weightsFlat, weightPerTx);
            
    return functional;
}

async function main() {
    const { ethers } = hre;
    const { PRIVATE_KEY, MODEL_JSON, WEIGHT_TXT, CHUNK_LEN } = process.env;
    if (!PRIVATE_KEY) {
        throw new Error("PRIVATE_KEY is not set");
    }
    if (!MODEL_JSON) {
        throw new Error("MODEL_JSON is not set");
    }
    if (!CHUNK_LEN) {
        throw new Error("CHUNK_LEN is not set");
    }

    const provider = ethers.provider;
    const signer = new ethers.Wallet(PRIVATE_KEY, provider);

    const configs = JSON.parse(fs.readFileSync(MODEL_JSON, 'utf-8'));        
    let weight_b64 = "";
    if (WEIGHT_TXT) {
        weight_b64 = fs.readFileSync(WEIGHT_TXT, 'utf-8');
    }

    const weightPerTx = parseInt(CHUNK_LEN);
    await deployFunctionalModel(configs, weight_b64, weightPerTx, signer);
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
