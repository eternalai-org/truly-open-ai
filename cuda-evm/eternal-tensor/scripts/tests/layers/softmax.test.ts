import { expect, assert } from 'chai';
import { BigNumber } from 'ethers';
import { ethers } from 'hardhat';
import { SoftmaxLayer } from '../../../typechain-types';
import { RandomSeed, create } from 'random-seed';
import { fromFloat, isBigNumberArrayEqual, toFloat, randomFloatArray, recursiveFromFloat, recursiveToFloat, encodeData, deflatten, decodeData, isFloatArrayEqual} from '../../libraries/utils';
import { Tensor } from '../../libraries/tensorData';
import { execSync } from 'child_process';
import fs from 'fs';
import path from 'path';

const abic = ethers.utils.defaultAbiCoder;
const TEST_FOLDER = "scripts/tests/layers"
const SCRIPT_PATH = path.join(TEST_FOLDER, 'softmax.py');
const OUTPUT_PATH = path.join(TEST_FOLDER, 'output.txt');
const CONFIG_PATH = path.join(TEST_FOLDER, 'config.json');
const PRECISION = 1e-3;

let softmaxContract: SoftmaxLayer, randomizer: RandomSeed;

async function deploySoftmaxLayer() {
    const configData = abic.encode([], []);
    
    const SoftmaxLayer = await ethers.getContractFactory('SoftmaxLayer');
    softmaxContract = await SoftmaxLayer.deploy(configData);
    await softmaxContract.deployed();
}

async function forwardSoftmaxLayer(softmaxContract: SoftmaxLayer, input: any): Promise<any> {
    const inputTensors: Tensor[] = [Tensor.fromFloatArray(input)];
    const outputTensor = await softmaxContract.forward(inputTensors);
    const output = new Tensor(outputTensor.data, outputTensor.shapes);
    return output.toFloatArray();
}

async function getKerasSoftmaxOutput(input: number[]) {
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

describe('SoftmaxLayer', async () => {
    before(async () => {
        const seed = new Date().toLocaleString()
        console.log("Seed random: \"" + seed + "\"")
        randomizer = create(seed);
    });

    describe('1. deploy', async () => {
        it ('1.1. Manual test', async () => {

            await deploySoftmaxLayer();
        });
    });

    describe('2. forward', async () => {
        it ('2.1. Forward pass test with 1D tensor', async () => {

            await deploySoftmaxLayer();
            console.log("Model deployed");
            
            const testInput = 
                [
                    [4.77383267, 5.2659086,  4.72565105, 9.11940114, 2.47470725, 1.67237745, 9.26399984, 7.90828651, 4.54410091, 2.68503853]
                ];
            const expectedOutput = 
                [
                    [0.00515727, 0.0084358, 0.00491468, 0.39781084, 0.00051751, 0.00023199, 0.45970056, 0.11849396, 0.00409873, 0.00063866]
                ];
            const output32x32 = await forwardSoftmaxLayer(softmaxContract,testInput);

            expect(isFloatArrayEqual(output32x32, expectedOutput, PRECISION)).to.equal(true);
        });

        it ('2.2. Random test', async () => {

            const inputUnits = randomizer.intBetween(10,64);
            const inputs = randomFloatArray(randomizer,[inputUnits], -5.0, 5.0);

            const expectedOutput = await getKerasSoftmaxOutput(inputs);

            await deploySoftmaxLayer();
            console.log("Model deployed");
        
            const output = await forwardSoftmaxLayer(softmaxContract,inputs);

            expect(isFloatArrayEqual(output, expectedOutput, PRECISION)).to.equal(true);
        });
    }); 
});
