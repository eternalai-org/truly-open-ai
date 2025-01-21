import { expect, assert } from 'chai';
import { BigNumber } from 'ethers';
import { ethers } from 'hardhat';
import { RescaleLayer } from '../../../typechain-types';
import { RandomSeed, create } from 'random-seed';
import { fromFloat, isBigNumberArrayEqual, toFloat, randomFloatArray, recursiveFromFloat, recursiveToFloat, encodeData, deflatten, decodeData, isFloatArrayEqual} from '../../libraries/utils';
import { Tensor } from '../../libraries/tensorData';
import { execSync } from 'child_process';
import fs from 'fs';
import path from 'path';

const abic = ethers.utils.defaultAbiCoder;
const TEST_FOLDER = "scripts/tests/layers"
const SCRIPT_PATH = path.join(TEST_FOLDER, 'rescale.py');
const OUTPUT_PATH = path.join(TEST_FOLDER, 'output.txt');
const CONFIG_PATH = path.join(TEST_FOLDER, 'config.json');
const PRECISION = 1e-3;

let rescaleContract: RescaleLayer, randomizer: RandomSeed;

async function deployRescaleLayer(scale: number, offset: number) {
    const configData = abic.encode(["uint256", "uint256"], [
        fromFloat(scale),
        fromFloat(offset)
    ]);
    
    const RescaleLayer = await ethers.getContractFactory('RescaleLayer');
    rescaleContract = await RescaleLayer.deploy(configData);
    await rescaleContract.deployed();
}

async function forwardRescaleLayer(rescaleContract: RescaleLayer, input: any): Promise<any> {
    const inputTensors: Tensor[] = [Tensor.fromFloatArray(input)];
    const outputTensor = await rescaleContract.forward(inputTensors);
    const output = new Tensor(outputTensor.data, outputTensor.shapes);
    return output.toFloatArray();
}

async function getKerasRescaleOutput(scale: number, offset: number, input: number[]) {
    const config = {
        scale,
        offset,
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

describe('RescaleLayer', async () => {
    before(async () => {
        const seed = new Date().toLocaleString()
        console.log("Seed random: \"" + seed + "\"")
        randomizer = create(seed);
    });

    describe('1. deploy', async () => {
        it ('1.1. Manual test', async () => {
            const scale = 1.0/255.0;
            const offset = 0;

            await deployRescaleLayer(scale, offset);
            expect(await rescaleContract.scale()).to.equal(BigNumber.from(fromFloat(scale)));
            expect(await rescaleContract.offset()).to.equal(BigNumber.from(fromFloat(offset)));
        });
    });

    describe('2. forward', async () => {
        it ('2.1. Forward pass test with 3D tensor', async () => {
            const scale = 1.0/255.0;
            const offset = 0;

            await deployRescaleLayer(scale, offset);
            console.log("Model deployed");

            const testInput =
                [
                    [[1.0], [2.0], [3.0]],
                    [[5.0], [6.0], [7.0]],
                    [[9.0], [10.0], [11.0]]
                ];

            const expectedOutput = [[[0.00392157], [0.00784314], [0.01176471]],
                                    [[0.01960784],  [0.02352941],  [0.02745098]],
                                    [[0.03529412], [0.03921569], [0.04313725]],
            ];
            
            const output32x32 = await forwardRescaleLayer(rescaleContract,testInput);

            expect(isFloatArrayEqual(output32x32, expectedOutput, PRECISION)).to.equal(true);
        });

        it ('2.2. Random test', async () => {
            const scale = 1.0/255.0;
            const offset = 0;

            const inputHeight = randomizer.intBetween(3,6);
            const inputWidth = randomizer.intBetween(3,6);
            const inputChannel = randomizer.intBetween(3,6);
            const inputs = randomFloatArray(randomizer,[inputHeight, inputWidth, inputChannel], -5.0, 5.0);

            const expectedOutput = await getKerasRescaleOutput(scale, offset, inputs);

            await deployRescaleLayer(scale, offset);
            console.log("Model deployed");
            
            const output = await forwardRescaleLayer(rescaleContract,inputs);

            expect(isFloatArrayEqual(output, expectedOutput, PRECISION)).to.equal(true);
        });
    }); 
});
