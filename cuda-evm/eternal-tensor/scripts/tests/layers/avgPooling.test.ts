import { expect, assert } from 'chai';
import { BigNumber } from 'ethers';
import { ethers } from 'hardhat';
import { AveragePooling2DLayer } from '../../../typechain-types';
import { RandomSeed, create } from 'random-seed';
import { isFloatArrayEqual, fromFloat, isBigNumberArrayEqual, toFloat, randomFloatArray, recursiveFromFloat, recursiveToFloat, encodeData, deflatten, decodeData} from '../../libraries/utils';
import { Padding } from '../../libraries/modelLib'
import { Tensor } from '../../libraries/tensorData';
import { execSync } from 'child_process';
import fs from 'fs';
import path from 'path';

const abic = ethers.utils.defaultAbiCoder;
const TEST_FOLDER = "scripts/tests/layers"
const SCRIPT_PATH = path.join(TEST_FOLDER, 'avgpooling2D.py');
const OUTPUT_PATH = path.join(TEST_FOLDER, 'output.txt');
const CONFIG_PATH = path.join(TEST_FOLDER, 'config.json');
const PRECISION = 1e-3;

let avgPoolingContract: AveragePooling2DLayer, randomizer: RandomSeed;

async function deployAveragePooling2DLayer(size: number[], stride: number[], padding: Padding) {
    const configData = abic.encode(["uint256[2]", "uint256[2]", "uint8"], [
        size,
        stride,
        padding,
    ]);
    
    const AveragePooling2DLayer = await ethers.getContractFactory('AveragePooling2DLayer');
    avgPoolingContract = await AveragePooling2DLayer.deploy(configData);
    await avgPoolingContract.deployed();
}

async function forwardAvgPoolingLayer(avgPoolingContract: AveragePooling2DLayer, input: any): Promise<any> {
    const inputTensors: Tensor[] = [Tensor.fromFloatArray(input)];
    // console.log(inputTensors[0]);
    const outputTensor = await avgPoolingContract.forward(inputTensors);
    const output = new Tensor(outputTensor.data, outputTensor.shapes);
    return output.toFloatArray();
}

async function getKerasAvgPooling2DOutput(size: number[], stride: number[], padding: number, inputs: number[]) {
    const config = {
        size,
        stride,
        padding: Padding[padding],
        inputs,
    }
    fs.writeFileSync(CONFIG_PATH, JSON.stringify(config, null, 2));

    let cmd = `python ${SCRIPT_PATH} --config-path ${CONFIG_PATH} --output-path ${OUTPUT_PATH}`;
    execSync(cmd);

    const data = fs.readFileSync(OUTPUT_PATH).toString();
    const output = JSON.parse(data);
    fs.rmSync(CONFIG_PATH);
    fs.rmSync(OUTPUT_PATH);

    return output;
}

describe('AveragePooling2DLayer', async () => {
    before(async () => {
        const seed = new Date().toLocaleString()
        console.log("Seed random: \"" + seed + "\"")
        randomizer = create(seed);
    });

    describe('1. deploy', async () => {
        it ('1.1. Manual test', async () => {
            const size = [2,2];
            const stride = [2,2];

            await deployAveragePooling2DLayer(size, stride, Padding.same);
            
            expect(await avgPoolingContract.getSize()).to.deep.equal(size);
            expect(await avgPoolingContract.getStride()).to.deep.equal(stride);
            expect(await avgPoolingContract.padding()).to.equal(Padding.same);
        });
    });

    describe('2. forward', async () => {
        it ('2.1. Forward test with padding same', async () => {
            const size = [2, 2];
            const stride = [2, 2];

            await deployAveragePooling2DLayer(size, stride, Padding.same);
            console.log("Model deployed");

            const testInput =
                [
                    [[1.0], [2.0], [3.0]],
                    [[5.0], [6.0], [7.0]],
                    [[9.0], [10.0], [11.0]]
                ];

            const expectedOutput = [[[3.5], [5.0]],
                                    [[9.5],[11.0]]
            ];
            
            const output = await forwardAvgPoolingLayer(avgPoolingContract,testInput);

            expect(isFloatArrayEqual(output, expectedOutput, PRECISION)).to.equal(true);
        });

        it ('2.2. Forward test with padding valid', async () => {
            const size = [2, 2];
            const stride = [2, 2];

            await deployAveragePooling2DLayer(size, stride, Padding.valid);
            console.log("Model deployed");

            const testInput =
                [
                    [[1.0], [2.0], [3.0]],
                    [[5.0], [6.0], [7.0]],
                    [[9.0], [10.0], [11.0]]
                ];

            const expectedOutput = [[[3.5]]];
            
            const output = await forwardAvgPoolingLayer(avgPoolingContract,testInput);

            expect(isFloatArrayEqual(output, expectedOutput, PRECISION)).to.equal(true);
        });

        it ('2.3. Random test with padding valid', async () => {
            const size = [2, 2];
            const stride = [2, 2];
            const padding = Padding.valid;

            const inputHeight = randomizer.intBetween(3,5);
            const inputWidth = randomizer.intBetween(3,5);
            const inputChannel = randomizer.intBetween(3,5);

            const inputs = randomFloatArray(randomizer, [inputHeight, inputWidth, inputChannel], -5.0, 5.0);

            const expectedOutput = await getKerasAvgPooling2DOutput(size, stride, padding, inputs);

            await deployAveragePooling2DLayer(size, stride, padding);
            console.log("Model deployed");
            
            const output = await forwardAvgPoolingLayer(avgPoolingContract, inputs);

            expect(isFloatArrayEqual(output, expectedOutput, PRECISION)).to.equal(true);
        });

        it ('2.4. Random test with padding same', async () => {
            const size = [2, 2];
            const stride = [2, 2];
            const padding = Padding.same;

            const inputHeight = randomizer.intBetween(3,5);
            const inputWidth = randomizer.intBetween(3,5);
            const inputChannel = randomizer.intBetween(3,5);

            const inputs = randomFloatArray(randomizer, [inputHeight, inputWidth, inputChannel], -5.0, 5.0);

            const expectedOutput = await getKerasAvgPooling2DOutput(size, stride, padding, inputs);

            await deployAveragePooling2DLayer(size, stride, padding);
            console.log("Model deployed");
            
            const output = await forwardAvgPoolingLayer(avgPoolingContract, inputs);

            expect(isFloatArrayEqual(output, expectedOutput, PRECISION)).to.equal(true);
        });
    }); 
});
