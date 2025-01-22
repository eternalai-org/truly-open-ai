import { expect, assert, use } from 'chai';
import { BigNumber } from 'ethers';
import { ethers } from 'hardhat';
import { DenseLayer } from '../../../typechain-types';
import { RandomSeed, create } from 'random-seed';
import { Activation, getRandomActivation } from '../../libraries/modelLib'
import { fromFloat, isBigNumberArrayEqual, isFloatArrayEqual, randomFloatArray, recursiveFromFloat, recursiveToFloat, toFloat, encodeData, deflatten, decodeData } from '../../libraries/utils';
import { Tensor, TensorData } from '../../libraries/tensorData';
import { execSync } from 'child_process';
import fs from 'fs';
import path from 'path';

const abic = ethers.utils.defaultAbiCoder;
const TEST_FOLDER = "scripts/tests/layers"
const SCRIPT_PATH = path.join(TEST_FOLDER, 'dense.py');
const OUTPUT_PATH = path.join(TEST_FOLDER, 'output.txt');
const CONFIG_PATH = path.join(TEST_FOLDER, 'config.json');
const PRECISION = 1e-3;

let denseContract: DenseLayer, randomizer: RandomSeed;

async function deployDenseLayer(actv: Activation, inputUnits: number, outputUnits: number, useBias: boolean) {
    const configData = abic.encode(["uint8", "uint256", "uint256", "bool"], [
        actv,
        BigNumber.from(outputUnits),
        BigNumber.from(inputUnits),
        useBias,
    ]);
    
    const DenseLayer = await ethers.getContractFactory('DenseLayer');
    denseContract = await DenseLayer.deploy(configData);
    await denseContract.deployed();
}

async function forwardDenseLayer(denseContract: DenseLayer, input: any): Promise<any> {
    const inputTensors: Tensor[] = [Tensor.fromFloatArray(input)];
    const outputTensor = await denseContract.forward(inputTensors);
    const output = new Tensor(outputTensor.data, outputTensor.shapes);
    return output.toFloatArray();
}

async function getKerasDenseOutput(actv: number, inputUnits: number, outputUnits: number, useBias: boolean, params: number[], inputs: number[]) {
    const config = {
        activation: Activation[actv],
        inputUnits,
        outputUnits,
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

describe.only('DenseLayer', async () => {
    before(async () => {
        const seed = new Date().toLocaleString()
        console.log("Seed random: \"" + seed + "\"")
        randomizer = create(seed);
    });

    describe('1. deploy', async () => {
        it('1.1. Manual test', async () => {
            const inputUnits = 128;
            const outputUnits = 32;
            const useBias = true;
            await deployDenseLayer(Activation.relu, inputUnits, outputUnits, useBias);
            expect(await denseContract.inputDim()).to.equal(BigNumber.from(inputUnits));
            expect(await denseContract.outputDim()).to.equal(BigNumber.from(outputUnits));
            expect(await denseContract.useBias()).to.equal(useBias);
            expect(await denseContract.activation()).to.equal(Activation.relu);
        });
    });

    describe('2. appendWeights', async () => {
        it('2.1. Manual test', async () => {
            const inputUnits = 3;
            const outputUnits = 2;
            const useBias = true;
            const expectedW32x32 = recursiveFromFloat([
                [1.0, 2.0],
                [3.0, 4.0],
                [5.0, 6.0],
            ]);
            const expectedB32x32 = recursiveFromFloat([7.0, 8.0]);
            const expectedWShape = [inputUnits, outputUnits].map(x => ethers.BigNumber.from(x));
            const expectedBShape = [outputUnits].map(x => ethers.BigNumber.from(x));

            const expectedWEncoded = encodeData(expectedW32x32.flat(Infinity));
            const expectedBEncoded = encodeData(expectedB32x32.flat(Infinity));
            const weightsEncoded = [expectedWEncoded, expectedBEncoded].flat();

            await deployDenseLayer(Activation.relu, inputUnits, outputUnits, useBias);

            await (await denseContract.appendWeights(weightsEncoded)).wait();

            const [wEncoded, wShape] = (await denseContract.getW());
            const [bEncoded, bShape] = (await denseContract.getB());

            const w32x32 = deflatten(decodeData(wEncoded), [inputUnits, outputUnits]);
            const b32x32 = deflatten(decodeData(bEncoded), [outputUnits]);
            
            expect(isBigNumberArrayEqual(w32x32, expectedW32x32)).to.equal(true);
            expect(isBigNumberArrayEqual(b32x32, expectedB32x32)).to.equal(true);

            expect(isBigNumberArrayEqual(wShape, expectedWShape)).to.equal(true);
            expect(isBigNumberArrayEqual(bShape, expectedBShape)).to.equal(true);        
        });

        it('2.2. Gas cost estimate', async () => {
            const inputUnits = 128;
            const outputUnits = 128;
            const useBias = true;
            const chunkSize = 10000;

            await deployDenseLayer(Activation.relu, inputUnits, outputUnits, useBias);

            const expectedW32x32 = recursiveFromFloat(randomFloatArray(randomizer, [inputUnits, outputUnits], -1, 1));
            const expectedB32x32 = recursiveFromFloat(randomFloatArray(randomizer, [outputUnits], -1, 1));
            const expectedWShape = [inputUnits, outputUnits].map(x => ethers.BigNumber.from(x));
            const expectedBShape = [outputUnits].map(x => ethers.BigNumber.from(x));

            const expectedWEncoded = encodeData(expectedW32x32.flat(Infinity));
            const expectedBEncoded = encodeData(expectedB32x32.flat(Infinity));
            const weightsEncoded = [expectedWEncoded, expectedBEncoded].flat();
            
            let totalGas = 0;
            for(let i = 0; i < weightsEncoded.length; i += chunkSize) {
                const weightsToUpload = weightsEncoded.slice(i, i + chunkSize);
                const tx = await denseContract.appendWeights(weightsToUpload);
                const receipt = await tx.wait();
                totalGas += receipt['gasUsed'].toNumber();
                console.log("Gas used:", receipt['gasUsed'].toNumber());
            }
            console.log(`Total gas cost to append ${weightsEncoded.length} params: ${totalGas}`);
            
            const [wEncoded, wShape] = (await denseContract.getW());
            const [bEncoded, bShape] = (await denseContract.getB());

            const w32x32 = deflatten(decodeData(wEncoded), [inputUnits, outputUnits]);
            const b32x32 = deflatten(decodeData(bEncoded), [outputUnits]);

            expect(isBigNumberArrayEqual(w32x32, expectedW32x32)).to.equal(true);
            expect(isBigNumberArrayEqual(b32x32, expectedB32x32)).to.equal(true);

            expect(isBigNumberArrayEqual(wShape, expectedWShape)).to.equal(true);
            expect(isBigNumberArrayEqual(bShape, expectedBShape)).to.equal(true);
        });
    });
    
    describe('3. forward', async () => {
        it('3.1. Manual test', async () => {
            const inputUnits = 3;
            const outputUnits = 2;
            const useBias = true;
            const expectedW32x32 = recursiveFromFloat([
                [1.0, 2.0],
                [3.0, 4.0],
                [5.0, 6.0],
            ]);
            const expectedB32x32 = recursiveFromFloat([7.0, 8.0]);
            const input = [1.0, 2.0, 3.0];
            const expectedOutput = [29.0, 36.0];

            const expectedWEncoded = encodeData(expectedW32x32.flat(Infinity));
            const expectedBEncoded = encodeData(expectedB32x32.flat(Infinity));
            const weightsEncoded = [expectedWEncoded, expectedBEncoded].flat();
            
            await deployDenseLayer(Activation.relu, inputUnits, outputUnits, useBias);
            console.log("Model deployed");

            await (await denseContract.appendWeights(weightsEncoded)).wait();
            console.log("Weights appended");

            const output = await forwardDenseLayer(denseContract, input);

            expect(isFloatArrayEqual(output, expectedOutput, PRECISION)).to.equal(true);
        });

        it('3.2. Random test', async () => {
            const inputUnits = randomizer.intBetween(3, 3);
            const outputUnits = randomizer.intBetween(2, 2);
            const actv = getRandomActivation(randomizer);
            const useBias = randomizer.intBetween(0, 1) ? true : false;
            
            const w = randomFloatArray(randomizer, [inputUnits * outputUnits], 0, 1.0);
            const b = randomFloatArray(randomizer, [useBias ? outputUnits : 0], 0, 1.0);
            const inputs = randomFloatArray(randomizer, [inputUnits], 0, 1.0);
            
            const params = [w, b].flat();
            const expectedOutput = await getKerasDenseOutput(actv, inputUnits, outputUnits, useBias, params, inputs)

            const w32x32 = recursiveFromFloat(w);
            const b32x32 = recursiveFromFloat(b);
            const expectedWEncoded = encodeData(w32x32.flat(Infinity));
            const expectedBEncoded = encodeData(b32x32.flat(Infinity));
            const weightsEncoded = [expectedWEncoded, expectedBEncoded].flat();

            await deployDenseLayer(actv, inputUnits, outputUnits, useBias);
            console.log("Model deployed");

            await (await denseContract.appendWeights(weightsEncoded)).wait();
            console.log("Weights appended");

            const output = await forwardDenseLayer(denseContract, inputs);

            expect(isFloatArrayEqual(output, expectedOutput, PRECISION)).to.equal(true);
        });
    });
});
