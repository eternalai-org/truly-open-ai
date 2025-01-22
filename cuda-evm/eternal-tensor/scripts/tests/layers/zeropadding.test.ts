import { expect, assert } from 'chai';
import { BigNumber } from 'ethers';
import { ethers } from 'hardhat';
import { ZeroPadding2DLayer } from '../../../typechain-types';
import { RandomSeed, create } from 'random-seed';
import { ZeroPaddingFormat } from '../../libraries/modelLib'
import { fromFloat, isBigNumberArrayEqual, isFloatArrayEqual, randomFloatArray, recursiveFromFloat, recursiveToFloat, toFloat} from '../../libraries/utils';
import { Tensor } from '../../libraries/tensorData';
import { execSync } from 'child_process';
import fs from 'fs';
import path from 'path';

const abic = ethers.utils.defaultAbiCoder;
const TEST_FOLDER = "scripts/tests/layers"
const PRECISION = 1e-3;

let zeropaddingContract: ZeroPadding2DLayer, randomizer: RandomSeed;

async function deployZeroPadding2DLayer(padding: number[], data_format: ZeroPaddingFormat) {
    const configData = abic.encode(["uint[4]", "uint8"], [
        padding,
        data_format
    ]);
    
    const ZeroPadding2DLayer = await ethers.getContractFactory('ZeroPadding2DLayer');
    zeropaddingContract = await ZeroPadding2DLayer.deploy(configData);
    await zeropaddingContract.deployed();
}

async function forwardZeroPadding2DLayer(zeropaddingContract: ZeroPadding2DLayer, input: any): Promise<any> {
    const inputTensors: Tensor[] = [Tensor.fromFloatArray(input)];
    const outputTensor = await zeropaddingContract.forward(inputTensors);
    const output = new Tensor(outputTensor.data, outputTensor.shapes);
    return output.toFloatArray();
}

describe('ZeroPadding2DLayer', async () => {
    before(async () => {
        const seed = new Date().toLocaleString()
        console.log("Seed random: \"" + seed + "\"")
        randomizer = create(seed);
    });

    describe('1. deploy', async () => {
        it ('1.1. Manual test', async () => {
            const padding = [1,2,1,3];
            await deployZeroPadding2DLayer(padding, ZeroPaddingFormat.ChannelsLast);
            expect(await zeropaddingContract.padding(0)).to.equal(BigNumber.from(padding[0]));
            expect(await zeropaddingContract.padding(1)).to.equal(BigNumber.from(padding[1]));
            expect(await zeropaddingContract.padding(2)).to.equal(BigNumber.from(padding[2]));
            expect(await zeropaddingContract.padding(3)).to.equal(BigNumber.from(padding[3]));
            expect(await zeropaddingContract.data_format()).to.equal(ZeroPaddingFormat.ChannelsLast);
        });
    });
    
    describe('2. forward', async () => {
        it ('2.1. Manual test with channel last', async () => {
            const padding = [1,2,1,3];

            const input = [
                [[2, 0],
                [1, 0]]
            ];

            const expectedOutput = [
                [[0, 0],
                [0, 0],
                [0, 0],
                [0, 0],
                [0, 0],
                [0, 0]],
   
                [[0, 0],
                [2, 0],
                [1, 0],
                [0, 0],
                [0, 0],
                [0, 0]],
   
                [[0, 0],
                [0, 0],
                [0, 0],
                [0, 0],
                [0, 0],
                [0, 0]],
   
                [[0, 0],
                [0, 0],
                [0, 0],
                [0, 0],
                [0, 0],
                [0, 0]]
            ];

            await deployZeroPadding2DLayer(padding, ZeroPaddingFormat.ChannelsLast);
            console.log("Model deployed");

            const output = await forwardZeroPadding2DLayer(zeropaddingContract, input);
            console.log("output: ",output);

            expect(isFloatArrayEqual(output, expectedOutput, PRECISION)).to.equal(true);
        });

        it ('3.2. Manual test with channel first', async () => {
            const padding = [1,2,1,3];
            
            const input = [
                [[2, 0],
                [1, 0]]
            ];

            const expectedOutput = [
                [[0, 0, 0, 0, 0, 0],
                [0, 2, 1, 0, 0, 0],
                [0, 0, 0, 0, 0, 0],
                [0, 0, 0, 0, 0, 0],
                [0, 0, 0, 0, 0, 0]]
            ];

            await deployZeroPadding2DLayer(padding, ZeroPaddingFormat.ChannelsFirst);
            console.log("Model deployed");

            const output= await forwardZeroPadding2DLayer(zeropaddingContract, input);
            console.log("output: ", output);

            expect(isFloatArrayEqual(output, expectedOutput, PRECISION)).to.equal(true);
        });
    });
});
