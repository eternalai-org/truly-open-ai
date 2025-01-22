import { expect, assert } from 'chai';
import { BigNumber } from 'ethers';
import { ethers } from 'hardhat';
import { OnesLikeLayer } from '../../../typechain-types';
import { RandomSeed, create } from 'random-seed';
import { fromFloat, toFloat, randomFloatArray, recursiveFromFloat, recursiveToFloat, encodeData, deflatten, decodeData, isFloatArrayEqual} from '../../libraries/utils';
import { Tensor } from '../../libraries/tensorData';
import { execSync } from 'child_process';
import fs from 'fs';
import path from 'path';

const abic = ethers.utils.defaultAbiCoder;
const PRECISION = 1e-3;

let onesLikeContract: OnesLikeLayer, randomizer: RandomSeed;

async function deployOnesLikeLayer() {
    const configData = abic.encode([], []);
    
    const OnesLikeLayer = await ethers.getContractFactory('OnesLikeLayer');
    onesLikeContract = await OnesLikeLayer.deploy(configData);
    await onesLikeContract.deployed();
}

async function forwardOnesLikeLayer(onesLikeContract: OnesLikeLayer, input: any): Promise<any> {
    const inputTensors: Tensor[] = [Tensor.fromFloatArray(input)];
    console.log(inputTensors[0]);
    const outputTensor = await onesLikeContract.forward(inputTensors);
    const output = new Tensor(outputTensor.data, outputTensor.shapes);
    return output.toFloatArray();
}

describe('OnesLikeLayer', async () => {
    before(async () => {
        const seed = new Date().toLocaleString()
        console.log("Seed random: \"" + seed + "\"")
        randomizer = create(seed);
    });

    describe('1. deploy', async () => {
        it ('1.1. Manual test', async () => {

            await deployOnesLikeLayer();
        });
    });

    describe('2. forward', async () => {
        it ('2.1. Forward pass test with 3D tensor', async () => {

            await deployOnesLikeLayer();
            console.log("Model deployed");

            const testInput =
                [
                    [1.0, 2.0, 3.0],
                    [5.0, 6.0, 7.0],
                    [9.0, 10.0, 11.0]
                ];

            const expectedOutput = [[1.0, 1.0, 1.0],
                                    [1.0, 1.0, 1.0],
                                    [1.0, 1.0, 1.0]
            ];
            
            const output32x32 = await forwardOnesLikeLayer(onesLikeContract,testInput);

            expect(isFloatArrayEqual(output32x32, expectedOutput, PRECISION)).to.equal(true);
        });
    }); 
});
