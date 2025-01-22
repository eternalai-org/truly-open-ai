import { expect, assert } from 'chai';
import { BigNumber } from 'ethers';
import { ethers } from 'hardhat';
import { ReLULayer } from '../../../typechain-types';
import { RandomSeed, create } from 'random-seed';
import { fromFloat, isFloatArrayEqual, toFloat, randomFloatArray, recursiveFromFloat, recursiveToFloat, encodeData, deflatten, decodeData} from '../../libraries/utils';
import { Tensor } from '../../libraries/tensorData';
import { execSync } from 'child_process';
import fs from 'fs';
import path from 'path';

const abic = ethers.utils.defaultAbiCoder;
const TEST_FOLDER = "scripts/tests/layers"
const SCRIPT_PATH = path.join(TEST_FOLDER, 'relu.py');
const OUTPUT_PATH = path.join(TEST_FOLDER, 'output.txt');
const CONFIG_PATH = path.join(TEST_FOLDER, 'config.json');
const PRECISION = 1e-3;

let reluContract: ReLULayer, randomizer: RandomSeed;

async function deployReLULayer() {
    const configData = abic.encode([], []);
    
    const ReLuLayer = await ethers.getContractFactory('ReLULayer');
    reluContract = await ReLuLayer.deploy(configData);
    await reluContract.deployed();
}

async function forwardReLULayer(reluContract: ReLULayer, input: any): Promise<any> {
    const inputTensors: Tensor[] = [Tensor.fromFloatArray(input)];
    const outputTensor = await reluContract.forward(inputTensors);
    const output = new Tensor(outputTensor.data, outputTensor.shapes);
    return output.toFloatArray();
}

async function getKerasReluOutput(input: number[]) {
    const config = {
        input,
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

describe('ReLULayer', async () => {
    before(async () => {
        const seed = new Date().toLocaleString()
        console.log("Seed random: \"" + seed + "\"")
        randomizer = create(seed);
    });

    describe('1. deploy', async () => {
        it ('1.1. Manual test', async () => {

            await deployReLULayer();
        });
    });

    describe('2. forward', async () => {
        it ('2.1. Forward pass test with 3D tensor', async () => {

            await deployReLULayer();
            console.log("Model deployed");

            const testInput =
                [
                    [1.0, -2.0, 3.0],
                    [5.0, 6.0, -7.0],
                    [-9.0, -10.0, 11.0]
                ];

            const expectedOutput = [[1.0, 0.0, 3.0],
                                    [5.0, 6.0, 0.0],
                                    [0.0, 0.0, 11.0]
            ];
            
            const output32x32 = await forwardReLULayer(reluContract,testInput);

            expect(isFloatArrayEqual(output32x32, expectedOutput, PRECISION)).to.equal(true);
        });

        it ('2.2. Random test', async () => {

            const inputHeight = randomizer.intBetween(3,6);
            const inputWidth = randomizer.intBetween(3,6);
            const inputChannel = randomizer.intBetween(3,6);
            const inputs = randomFloatArray(randomizer,[inputHeight, inputWidth, inputChannel], -5.0, 5.0);

            const expectedOutput = await getKerasReluOutput(inputs);

            await deployReLULayer();
            console.log("Model deployed");
            
            const output = await forwardReLULayer(reluContract,inputs);

            expect(isFloatArrayEqual(output, expectedOutput, PRECISION)).to.equal(true);
        });
    }); 
});
