import { expect, assert } from 'chai';
import { BigNumber } from 'ethers';
import { ethers } from 'hardhat';
import { SigmoidLayer } from '../../../typechain-types';
import { RandomSeed, create } from 'random-seed';
import { fromFloat, isBigNumberArrayEqual, toFloat, randomFloatArray, recursiveFromFloat, recursiveToFloat, encodeData, deflatten, decodeData, isFloatArrayEqual} from '../../libraries/utils';
import { Tensor } from '../../libraries/tensorData';
import { execSync } from 'child_process';
import fs from 'fs';
import path from 'path';

const abic = ethers.utils.defaultAbiCoder;
const TEST_FOLDER = "scripts/tests/layers"
const SCRIPT_PATH = path.join(TEST_FOLDER, 'sigmoid.py');
const OUTPUT_PATH = path.join(TEST_FOLDER, 'output.txt');
const CONFIG_PATH = path.join(TEST_FOLDER, 'config.json');
const PRECISION = 1e-3;

let sigmoidContract: SigmoidLayer, randomizer: RandomSeed;

async function deploySigmoidLayer() {
    const configData = abic.encode([], []);
    
    const SigmoidLayer = await ethers.getContractFactory('SigmoidLayer');
    sigmoidContract = await SigmoidLayer.deploy(configData);
    await sigmoidContract.deployed();
}

async function forwardSigmoidLayer(sigmoidContract: SigmoidLayer, input: any): Promise<any> {
    const inputTensors: Tensor[] = [Tensor.fromFloatArray(input)];
    const outputTensor = await sigmoidContract.forward(inputTensors);
    const output = new Tensor(outputTensor.data, outputTensor.shapes);
    return output.toFloatArray();
}

async function getKerasSigmoidOutput(input: number[]) {
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

describe('SigmoidLayer', async () => {
    before(async () => {
        const seed = new Date().toLocaleString()
        console.log("Seed random: \"" + seed + "\"")
        randomizer = create(seed);
    });

    describe('1. deploy', async () => {
        it ('1.1. Manual test', async () => {

            await deploySigmoidLayer();
        });
    });

    describe('2. forward', async () => {
        it ('2.1. Forward pass test with 3D tensor', async () => {

            await deploySigmoidLayer();
            console.log("Model deployed");

            const testInput =
                [
                    [0.80531654, 0.30585969, 0.13773904],
                    [0.06516666, 0.88074556, 0.03621177],
                    [0.3904885,  0.32787716, 0.29096427],
                ];

            const expectedOutput = [[0.69111059, 0.57587434, 0.53438042],
                                    [0.5162859,  0.7069767,  0.50905195],
                                    [0.59640029, 0.58124277, 0.57223219],
            ];
            
            const output32x32 = await forwardSigmoidLayer(sigmoidContract,testInput);

            expect(isFloatArrayEqual(output32x32, expectedOutput, PRECISION)).to.equal(true);
        });

        // it ('2.2. Random test', async () => {

        //     const inputHeight = randomizer.intBetween(3,6);
        //     const inputWidth = randomizer.intBetween(3,6);
        //     const inputChannel = randomizer.intBetween(3,6);
        //     const inputs = randomFloatArray(randomizer,[inputHeight, inputWidth, inputChannel], -5.0, 5.0);

        //     const expectedOutput = await getKerasSigmoidOutput(inputs);

        //     await deploySigmoidLayer();
        //     console.log("Model deployed");
            
        //     const output32x32 = await forwardSigmoidLayer(sigmoidContract, inputs);

        //     expect(isFloatArrayEqual(output32x32, expectedOutput, PRECISION)).to.equal(true);
        // });
    }); 
});
