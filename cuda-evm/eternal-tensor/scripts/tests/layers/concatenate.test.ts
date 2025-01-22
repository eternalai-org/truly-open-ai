import { expect, assert } from 'chai';
import { BigNumber } from 'ethers';
import { ethers } from 'hardhat';
import { ConcatenateLayer } from '../../../typechain-types';
import { RandomSeed, create } from 'random-seed';
import { fromFloat, isBigNumberArrayEqual, isFloatArrayEqual, randomFloatArray, recursiveFromFloat, recursiveToFloat, toFloat} from '../../libraries/utils';
import { Tensor } from '../../libraries/tensorData';
import { execSync } from 'child_process';
import fs from 'fs';
import path from 'path';

const abic = ethers.utils.defaultAbiCoder;
const TEST_FOLDER = "scripts/tests/layers"
const PRECISION = 1e-3;

let concatenateContract: ConcatenateLayer, randomizer: RandomSeed;

async function deployConcatenateLayer(axis: number) {
    const configData = abic.encode(["int8"], [
        axis
    ]);
    
    const ConcatenateLayer = await ethers.getContractFactory('ConcatenateLayer');
    concatenateContract = await ConcatenateLayer.deploy(configData);
    await concatenateContract.deployed();
}

async function forwardConcatenateLayer(concatenateContract: ConcatenateLayer, input: any): Promise<any> {
    const inputTensors: Tensor[] = input.map((i: any) => Tensor.fromFloatArray(i));
    const outputTensor = await concatenateContract.forward(inputTensors);
    const output = new Tensor(outputTensor.data, outputTensor.shapes);
    return output.toFloatArray();
}

describe('ConcatenateLayer', async () => {
    before(async () => {
        const seed = new Date().toLocaleString()
        console.log("Seed random: \"" + seed + "\"")
        randomizer = create(seed);
    });

    describe('1. deploy', async () => {
        it ('1.1. Manual test', async () => {
            const axis = -1;
            await deployConcatenateLayer(axis);
            expect(await concatenateContract.axis()).to.equal(BigNumber.from(axis));
        });
    });
    
    describe('2. forward', async () => {
        it ('2.1. Manual test with last axis', async () => {
            const axis = -1;

            const input_1 = [
                [
                    [1.0, 2.0],
                    [3.0, 4.0],
                ],
                [
                    [5.0, 6.0],
                    [7.0, 8.0],
                ],
            ];

            const input_2 = [
                [
                    [1.0],
                    [3.0],
                ],
                [
                    [5.0],
                    [7.0],
                ],
            ];

            const expectedOutput = [
                [
                    [1.0, 2.0, 1.0],
                    [3.0, 4.0, 3.0],
                ],
                [
                    [5.0, 6.0, 5.0],
                    [7.0, 8.0, 7.0],
                ],
            ];

            const input = [input_1, input_2];

            await deployConcatenateLayer(axis);
            console.log("Model deployed");

            console.log("input: ", input);

            const output = await forwardConcatenateLayer(concatenateContract, input);

            expect(isFloatArrayEqual(output, expectedOutput, PRECISION)).to.equal(true);
        });

        it ('3.2. Manual test with first axis', async () => {
            const axis = 0;
            
            const input_1 = [
                [
                    [1.0, 2.0],
                    [3.0, 4.0],
                ],
                [
                    [5.0, 6.0],
                    [7.0, 8.0],
                ],
            ];

            const input_2 = [
                [
                    [1.0, 2.0],
                    [3.0, 4.0],
                ],
                [
                    [5.0, 6.0],
                    [7.0, 8.0],
                ],
                [
                    [5.0, 6.0],
                    [7.0, 8.0],
                ],
            ];

            const expectedOutput = [
                [
                    [1.0, 2.0],
                    [3.0, 4.0],
                ],
                [
                    [5.0, 6.0],
                    [7.0, 8.0],
                ],
                [
                    [1.0, 2.0],
                    [3.0, 4.0],
                ],
                [
                    [5.0, 6.0],
                    [7.0, 8.0],
                ],
                [
                    [5.0, 6.0],
                    [7.0, 8.0],
                ],
            ];

            const input = [input_1, input_2];

            await deployConcatenateLayer(axis);
            console.log("Model deployed");

            const output= await forwardConcatenateLayer(concatenateContract, input);

            expect(isFloatArrayEqual(output, expectedOutput, PRECISION)).to.equal(true);
        });

        // it('3.3. Random test (valid padding)', async () => {
        //     const inputFilters = randomizer.intBetween(3, 3);
        //     const outputFilters = randomizer.intBetween(5, 5);
        //     const actv = randomizer.intBetween(0, 4);
        //     const size = [3, 3];
        //     const stride = [1, 1];
        //     const padding = Padding.valid;
        //     await deployConv2DLayer(actv, inputFilters, outputFilters, size, stride, padding);
            
        //     const weightLen = (await conv2DContract.getParamsCount()).toNumber();
        //     const weights: number[] = randomFloatArray(randomizer, [weightLen], -1.0, 1.0);
        //     const weights32x32 = recursiveFromFloat(weights);
        //     await (await conv2DContract.appendWeights(weights32x32, 0)).wait();
            
        //     const inputs: number[][][] = randomFloatArray(randomizer, [5, 5, inputFilters], -1.0, 1.0);
        //     const inputs32x32 = recursiveFromFloat(inputs);
        //     const output32x32 = await forwardConv2DLayer(conv2DContract, inputs32x32);

        //     const scriptPath = path.join(TEST_FOLDER, 'conv2d.py');
        //     let cmd = `python ${scriptPath} --activation ${Activation[actv]} --inputFilters ${inputFilters} --outputFilters ${outputFilters} --size ${size[0]} --strides ${stride[0]} --padding ${Padding[padding]} --params '${JSON.stringify(weights)}' --input '${JSON.stringify(inputs)}'`;
        //     execSync(cmd);
        //     const tmpPath = path.join(TEST_FOLDER, 'tmp.txt');
        //     const data = fs.readFileSync(tmpPath).toString();

        //     const output = recursiveToFloat(output32x32);
        //     const expectedOutput = JSON.parse(data);
        //     fs.rmSync(tmpPath);

        //     expect(isFloatArrayEqual(output, expectedOutput, PRECISION)).to.equal(true);
        // });
    });
});
