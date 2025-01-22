import { expect, assert } from 'chai';
import { BigNumber } from 'ethers';
import { ethers } from 'hardhat';
import { GlobalAveragePooling2DLayer } from '../../../typechain-types';
import { RandomSeed, create } from 'random-seed';
import { fromFloat, isFloatArrayEqual, isBigNumberArrayEqual, toFloat, randomFloatArray} from '../../libraries/utils';
import { Tensor } from '../../libraries/tensorData';
import { execSync } from 'child_process';
import fs from 'fs';
import path from 'path';

const abic = ethers.utils.defaultAbiCoder;
const TEST_FOLDER = "scripts/tests/layers"
const SCRIPT_PATH = path.join(TEST_FOLDER, 'globalavgpooling.py');
const OUTPUT_PATH = path.join(TEST_FOLDER, 'output.txt');
const CONFIG_PATH = path.join(TEST_FOLDER, 'config.json');
const PRECISION = 1e-3;

let globalAvgPoolingContract: GlobalAveragePooling2DLayer, randomizer: RandomSeed;

async function deployGlobalAveragePooling2DLayer() {
    const configData = abic.encode([], []);
    
    const GlobalAveragePooling2DLayer = await ethers.getContractFactory('GlobalAveragePooling2DLayer');
    globalAvgPoolingContract = await GlobalAveragePooling2DLayer.deploy(configData);
    await globalAvgPoolingContract.deployed();
}

async function forwardGlobalAvgPoolingLayer(globalAvgPoolingContract: GlobalAveragePooling2DLayer, input: any): Promise<any> {
    const inputTensors: Tensor[] = [Tensor.fromFloatArray(input)];
    const outputTensor = await globalAvgPoolingContract.forward(inputTensors);
    const output = new Tensor(outputTensor.data, outputTensor.shapes);
    return output.toFloatArray();
}

async function getKerasGlobalAvgPoolingOutput(inputs: number[]) {
    const config = {
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

describe('GlobalAveragePooling2DLayer', async () => {
    before(async () => {
        const seed = new Date().toLocaleString()
        console.log("Seed random: \"" + seed + "\"")
        randomizer = create(seed);
    });

    describe('1. forward', async () => {
        it ('1.1. Forward pass test', async () => {

            await deployGlobalAveragePooling2DLayer();
            console.log("Model deployed");

            const testInput =[
                [[1., 2., 2.],
                [0., 1., 3.],
                [3., 4., 1.]],

                [[2., 2., 2.],
                [0., 4., 2.],
                [3., 3., 1.]],

                [[2., 6., 5.],
                [3., 2., 1.],
                [4., 3., 1.]]
            ];

            const expectedOutput = [2., 3., 2.];
            
            const output32x32 = await forwardGlobalAvgPoolingLayer(globalAvgPoolingContract,testInput);

            expect(isFloatArrayEqual(output32x32, expectedOutput, PRECISION)).to.equal(true);
        });

        it ('1.2. Random test', async () => {

            const inputHeight = randomizer.intBetween(3,10)
            const inputWidth = randomizer.intBetween(3,10)
            const inputChannel = randomizer.intBetween(3,10)

            const inputs = randomFloatArray(randomizer,[inputHeight, inputWidth, inputChannel], -5.0, 5.0)
            const expectedOutput = await getKerasGlobalAvgPoolingOutput(inputs)

            await deployGlobalAveragePooling2DLayer();
            console.log("Model deployed");
            
            const output = await forwardGlobalAvgPoolingLayer(globalAvgPoolingContract,inputs);

            expect(isFloatArrayEqual(output, expectedOutput, PRECISION)).to.equal(true);
        });
    }); 
});
