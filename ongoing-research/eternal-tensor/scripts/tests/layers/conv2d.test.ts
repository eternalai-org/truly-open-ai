import { expect, assert } from 'chai';
import { BigNumber } from 'ethers';
import { ethers } from 'hardhat';
import { Conv2DLayer } from '../../../typechain-types';
import { RandomSeed, create } from 'random-seed';
import { Activation, Padding, getRandomActivation } from '../../libraries/modelLib'
import { fromFloat, isBigNumberArrayEqual, isFloatArrayEqual, randomFloatArray, recursiveFromFloat, recursiveToFloat, toFloat, encodeData, deflatten, decodeData } from '../../libraries/utils';
import { Tensor } from '../../libraries/tensorData';
import { execSync } from 'child_process';
import fs from 'fs';
import path from 'path';

const abic = ethers.utils.defaultAbiCoder;
const TEST_FOLDER = "scripts/tests/layers"
const SCRIPT_PATH = path.join(TEST_FOLDER, 'conv2d.py');
const OUTPUT_PATH = path.join(TEST_FOLDER, 'output.txt');
const CONFIG_PATH = path.join(TEST_FOLDER, 'config.json');
const PRECISION = 1e-3;

let conv2DContract: Conv2DLayer, randomizer: RandomSeed;

async function deployConv2DLayer(actv: Activation, inputUnits: number, outputUnits: number, size: number[], stride: number[], padding: Padding, useBias: boolean) {
    const configData = abic.encode(["uint8", "uint256", "uint256", "uint256[2]", "uint256[2]", "uint8", "bool"], [
        actv,
        BigNumber.from(inputUnits),
        BigNumber.from(outputUnits),
        size,
        stride,
        padding,
        useBias
    ]);
    
    const Conv2DLayer = await ethers.getContractFactory('Conv2DLayer');
    conv2DContract = await Conv2DLayer.deploy(configData);
    await conv2DContract.deployed();
}

async function forwardConv2DLayer(conv2DContract: Conv2DLayer, input: any): Promise<any> {
    const inputTensors: Tensor[] = [Tensor.fromFloatArray(input)];
    console.log(inputTensors[0]);
    const outputTensor = await conv2DContract.forward(inputTensors);
    const output = new Tensor(outputTensor.data, outputTensor.shapes);
    return output.toFloatArray();
}

async function getKerasConv2DOutput(actv: number, inputFilters: number, outputFilters: number, size: number[], stride: number[], padding: number, useBias: boolean, params: number[], inputs: number[]) {
    const config = {
        activation: Activation[actv],
        inputFilters,
        outputFilters,
        size,
        stride,
        padding: Padding[padding],
        useBias,
        params,
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

describe('Conv2DLayer', async () => {
    before(async () => {
        const seed = new Date().toLocaleString()
        console.log("Seed random: \"" + seed + "\"")
        randomizer = create(seed);
    });

    describe('1. deploy', async () => {
        it ('1.1. Manual test', async () => {
            const inputFilters = 3;
            const outputFilters = 2;
            const size = [3, 3];
            const stride = [1, 1];
            const useBias = true;
            await deployConv2DLayer(Activation.relu, inputFilters, outputFilters, size, stride, Padding.same, useBias);
            expect(await conv2DContract.getFilters()).to.equal(BigNumber.from(outputFilters));
            expect(await conv2DContract.stride(0)).to.equal(BigNumber.from(stride[0]));
            expect(await conv2DContract.stride(1)).to.equal(BigNumber.from(stride[1]));
            expect(await conv2DContract.activation()).to.equal(Activation.relu);
            expect(await conv2DContract.padding()).to.equal(Padding.same);
            expect(await conv2DContract.useBias()).to.equal(useBias);
        });
    });

    describe('2. appendWeights', async () => {
        it ('2.1. Manual test', async () => {
            const inputFilters = 2;
            const outputFilters = 2;
            const size = [2, 2];
            const stride = [1, 1];
            const useBias = true;

            const expectedW32x32 = recursiveFromFloat([
                [
                    [
                        [1.0, 2.0],
                        [3.0, 4.0],
                    ],
                    [
                        [5.0, 6.0],
                        [7.0, 8.0],
                    ],
                ],
                [
                    [
                        [9.0, 10.0],
                        [11.0, 12.0],
                    ],
                    [
                        [13.0, 14.0],
                        [15.0, 16.0],
                    ],
                ],
            ]);
            const expectedB32x32 = recursiveFromFloat([17.0, 18.0]);
            const expectedWShape = [size[0], size[1], inputFilters, outputFilters].map(x => ethers.BigNumber.from(x));
            const expectedBShape = [outputFilters].map(x => ethers.BigNumber.from(x));

            const expectedWEncoded = encodeData(expectedW32x32.flat(Infinity));
            const expectedBEncoded = encodeData(expectedB32x32.flat(Infinity));
            const weightsEncoded = [expectedWEncoded, expectedBEncoded].flat();

            await deployConv2DLayer(Activation.relu, inputFilters, outputFilters, size, stride, Padding.same, useBias);

            await (await conv2DContract.appendWeights(weightsEncoded)).wait();

            const [wEncoded, wShape] = (await conv2DContract.getW());
            const [bEncoded, bShape] = (await conv2DContract.getB());

            const w32x32 = deflatten(decodeData(wEncoded), [size[0], size[1], inputFilters, outputFilters]);
            const b32x32 = deflatten(decodeData(bEncoded), [outputFilters]);

            expect(isBigNumberArrayEqual(w32x32, expectedW32x32)).to.equal(true);
            expect(isBigNumberArrayEqual(b32x32, expectedB32x32)).to.equal(true);

            expect(isBigNumberArrayEqual(wShape, expectedWShape)).to.equal(true);
            expect(isBigNumberArrayEqual(bShape, expectedBShape)).to.equal(true);
        });

        it ('2.2. Gas cost estimate', async () => {
            const inputFilters = 64;
            const outputFilters = 128;
            const size = [3, 3];
            const stride = [1, 1];
            const useBias = true;
            const chunkSize = 10000;
            await deployConv2DLayer(Activation.relu, inputFilters, outputFilters, size, stride, Padding.same, useBias);

            const expectedW32x32 = recursiveFromFloat(randomFloatArray(randomizer, [size[0], size[1], inputFilters, outputFilters], -1, 1));
            const expectedB32x32 = recursiveFromFloat(randomFloatArray(randomizer, [outputFilters], -1, 1));
            const expectedWShape = [size[0], size[1], inputFilters, outputFilters].map(x => ethers.BigNumber.from(x));
            const expectedBShape = [outputFilters].map(x => ethers.BigNumber.from(x));

            const expectedWEncoded = encodeData(expectedW32x32.flat(Infinity));
            const expectedBEncoded = encodeData(expectedB32x32.flat(Infinity));
            const weightsEncoded = [expectedWEncoded, expectedBEncoded].flat();
            
            let totalGas = 0;
            let txCount = 0;
            for(let i = 0; i < weightsEncoded.length; i += chunkSize) {
                const weightsToUpload = weightsEncoded.slice(i, i + chunkSize);
                const tx = await conv2DContract.appendWeights(weightsToUpload);
                const receipt = await tx.wait();
                totalGas += receipt['gasUsed'].toNumber();
                txCount += 1;
            }
            console.log(`Tx count: ${txCount}`);
            console.log(`Total gas cost to append ${weightsEncoded.length} params: ${totalGas}`);

            const [wEncoded, wShape] = (await conv2DContract.getW());
            const [bEncoded, bShape] = (await conv2DContract.getB());

            const w32x32 = deflatten(decodeData(wEncoded), [size[0], size[1], inputFilters, outputFilters]);
            const b32x32 = deflatten(decodeData(bEncoded), [outputFilters]);

            expect(isBigNumberArrayEqual(w32x32, expectedW32x32)).to.equal(true);
            expect(isBigNumberArrayEqual(b32x32, expectedB32x32)).to.equal(true);

            expect(isBigNumberArrayEqual(wShape, expectedWShape)).to.equal(true);
            expect(isBigNumberArrayEqual(bShape, expectedBShape)).to.equal(true);
        });
    });
    
    describe('3. forward', async () => {
        it ('3.1. Manual test (valid padding)', async () => {
            const inputFilters = 2;
            const outputFilters = 2;
            const size = [2, 2];
            const stride = [1, 1];
            const useBias = true;

            const expectedW32x32 = recursiveFromFloat([
                [
                    [
                        [1.0, 2.0],
                        [3.0, 4.0],
                    ],
                    [
                        [5.0, 6.0],
                        [4.0, 8.0],
                    ],
                ],
                [
                    [
                        [9.0, 10.0],
                        [11.0, 12.0],
                    ],
                    [
                        [13.0, 14.0],
                        [15.0, 16.0],
                    ],
                ],
            ]);
            const expectedB32x32 = recursiveFromFloat([17.0, 18.0]);

            const expectedWEncoded = encodeData(expectedW32x32.flat(Infinity));
            const expectedBEncoded = encodeData(expectedB32x32.flat(Infinity));
            const weightsEncoded = [expectedWEncoded, expectedBEncoded].flat();

            const input = [
                [
                    [1.0, 2.0],
                    [3.0, 4.0],
                ],
                [
                    [5.0, 6.0],
                    [7.0, 8.0],
                ],
            ];
            const expectedOutput = [ 
                [ 
                    [377.0, 426.0],
                ], 
            ];

            await deployConv2DLayer(Activation.relu, inputFilters, outputFilters, size, stride, Padding.valid, useBias);
            console.log("Model deployed");

            await (await conv2DContract.appendWeights(weightsEncoded)).wait();
            console.log("Weights appended");

            const output = await forwardConv2DLayer(conv2DContract, input);

            expect(isFloatArrayEqual(output, expectedOutput, PRECISION)).to.equal(true);
        });

        it ('3.2. Manual test (same padding)', async () => {
            const inputFilters = 2;
            const outputFilters = 2;
            const size = [2, 2];
            const stride = [1, 1];
            const useBias = true;
            
            const expectedW32x32 = recursiveFromFloat([
                [
                    [
                        [1.0, 2.0],
                        [3.0, 4.0],
                    ],
                    [
                        [5.0, 6.0],
                        [4.0, 8.0],
                    ],
                ],
                [
                    [
                        [9.0, 10.0],
                        [11.0, 12.0],
                    ],
                    [
                        [13.0, 14.0],
                        [15.0, 16.0],
                    ],
                ],
            ]);
            const expectedB32x32 = recursiveFromFloat([17.0, 18.0]);

            const expectedWEncoded = encodeData(expectedW32x32.flat(Infinity));
            const expectedBEncoded = encodeData(expectedB32x32.flat(Infinity));
            const weightsEncoded = [expectedWEncoded, expectedBEncoded].flat();
            
            const input = [
                [
                    [1.0, 2.0],
                    [3.0, 4.0],
                ],
                [
                    [5.0, 6.0],
                    [7.0, 8.0],
                ],
            ];
            const expectedOutput = [
                [
                    [377.0, 426.0],
                    [183.0, 206.0],
                ],
                [
                    [107.0, 158.0],
                    [48.0, 64.0],
                ],
            ];

            await deployConv2DLayer(Activation.relu, inputFilters, outputFilters, size, stride, Padding.same, useBias);
            console.log("Model deployed");

            await (await conv2DContract.appendWeights(weightsEncoded)).wait();
            console.log("Weights appended");

            const output= await forwardConv2DLayer(conv2DContract, input);

            expect(isFloatArrayEqual(output, expectedOutput, PRECISION)).to.equal(true);
        });

        it ('3.3. Random test (valid padding)', async () => {
            const inputFilters = randomizer.intBetween(3, 3);
            const outputFilters = randomizer.intBetween(5, 5);
            const actv = getRandomActivation(randomizer);
            const useBias = true;
            const size = [3, 3];
            const stride = [1, 1];
            const padding = Padding.valid;

            const w = randomFloatArray(randomizer, [size[0] * size[1] * inputFilters * outputFilters], -1.0, 1.0);
            const b = randomFloatArray(randomizer, [useBias ? outputFilters : 0], 0, 1.0);

            const inputHeight = 5;
            const inputWidth = 5;
            const inputs = randomFloatArray(randomizer, [inputHeight, inputWidth, inputFilters], -1.0, 1.0);
            const params = [w, b].flat();
            
            const expectedOutput = await getKerasConv2DOutput(actv, inputFilters, outputFilters, size, stride, padding, useBias, params, inputs);

            const w32x32 = recursiveFromFloat(w);
            const b32x32 = recursiveFromFloat(b);
            const expectedWEncoded = encodeData(w32x32.flat(Infinity));
            const expectedBEncoded = encodeData(b32x32.flat(Infinity));
            const weightsEncoded = [expectedWEncoded, expectedBEncoded].flat();

            await deployConv2DLayer(actv, inputFilters, outputFilters, size, stride, padding, useBias);
            console.log("Model deployed");

            await (await conv2DContract.appendWeights(weightsEncoded)).wait();
            console.log("Weights appended");

            const output = await forwardConv2DLayer(conv2DContract, inputs);

            expect(isFloatArrayEqual(output, expectedOutput, PRECISION)).to.equal(true);
        });
        
        it ('3.4. Random test (same padding)', async () => {
            const inputFilters = randomizer.intBetween(3, 3);
            const outputFilters = randomizer.intBetween(5, 5);
            const actv = getRandomActivation(randomizer);
            const useBias = true;
            const size = [5, 5];
            const stride = [2, 2];
            const padding = Padding.same;

            const w = randomFloatArray(randomizer, [size[0] * size[1] * inputFilters * outputFilters], -1.0, 1.0);
            const b = randomFloatArray(randomizer, [useBias ? outputFilters : 0], 0, 1.0);

            const inputHeight = 24;
            const inputWidth = 24;
            const inputs = randomFloatArray(randomizer, [inputHeight, inputWidth, inputFilters], -1.0, 1.0);
            const params = [w, b].flat();
            
            const expectedOutput = await getKerasConv2DOutput(actv, inputFilters, outputFilters, size, stride, padding, useBias, params, inputs);

            const w32x32 = recursiveFromFloat(w);
            const b32x32 = recursiveFromFloat(b);
            const expectedWEncoded = encodeData(w32x32.flat(Infinity));
            const expectedBEncoded = encodeData(b32x32.flat(Infinity));
            const weightsEncoded = [expectedWEncoded, expectedBEncoded].flat();

            await deployConv2DLayer(actv, inputFilters, outputFilters, size, stride, padding, useBias);
            console.log("Model deployed");

            await (await conv2DContract.appendWeights(weightsEncoded)).wait();
            console.log("Weights appended");

            const output = await forwardConv2DLayer(conv2DContract, inputs);
            expect(isFloatArrayEqual(output, expectedOutput, PRECISION)).to.equal(true);
        });
    });
});
