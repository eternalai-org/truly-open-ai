import { expect, assert } from 'chai';
import { BigNumber } from 'ethers';
import { ethers } from 'hardhat';
import { FlattenLayer } from '../../../typechain-types';
import { RandomSeed, create } from 'random-seed';
import { fromFloat, isBigNumberArrayEqual, isFloatArrayEqual, randomFloatArray, recursiveFromFloat, recursiveToFloat, toFloat, encodeData, deflatten, decodeData } from '../../libraries/utils';
import { Tensor, TensorData } from '../../libraries/tensorData';
import { execSync } from 'child_process';
import fs from 'fs';
import path from 'path';

const abic = ethers.utils.defaultAbiCoder;
const TEST_FOLDER = "scripts/tests/layers"
const PRECISION = 1e-3;

let flattenContract: FlattenLayer, randomizer: RandomSeed;

async function deployFlattenLayer() {
    const configData = abic.encode([], []);
    
    const FlattenLayer = await ethers.getContractFactory('FlattenLayer');
    flattenContract = await FlattenLayer.deploy(configData);
    await flattenContract.deployed();
}

async function forwardFlattenLayer(flattenContract: FlattenLayer, input: any): Promise<any> {
    const inputTensors: Tensor[] = [Tensor.fromFloatArray(input)];
    console.log(inputTensors[0]);
    const outputTensor = await flattenContract.forward(inputTensors);
    const output = new Tensor(outputTensor.data, outputTensor.shapes);
    return output.toFloatArray();
}

describe('FlattenLayer', async () => {
    before(async () => {
        const seed = new Date().toLocaleString()
        console.log("Seed random: \"" + seed + "\"")
        randomizer = create(seed);
    });
    
    describe('1. forward', async () => {
        it ('1.1. Manual test', async () => {

            const input = [
                [1.0, 2.0],
                [3.0, 4.0],
                [5.0, 6.0],
            ];
            const expectedOutput = [1.0, 2.0, 3.0, 4.0, 5.0, 6.0];
            
            await deployFlattenLayer();
            console.log("Model deployed");

            const output = await forwardFlattenLayer(flattenContract, input);

            expect(isFloatArrayEqual(output, expectedOutput, PRECISION)).to.equal(true);
        });
    });
});
