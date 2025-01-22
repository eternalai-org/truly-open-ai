import { expect, assert } from 'chai';
import { BigNumber } from 'ethers';
import { ethers } from 'hardhat';
import { BatchNormalizationLayer } from '../../../typechain-types';
import { RandomSeed, create } from 'random-seed';
import { fromFloat, isBigNumberArrayEqual, isFloatArrayEqual, randomFloatArray, recursiveFromFloat, recursiveToFloat, toFloat, encodeData, deflatten, decodeData } from '../../libraries/utils';
import { Tensor } from '../../libraries/tensorData';
import { execSync } from 'child_process';
import fs from 'fs';
import path from 'path';

const abic = ethers.utils.defaultAbiCoder;
const TEST_FOLDER = "scripts/tests/layers"
const SCRIPT_PATH = path.join(TEST_FOLDER, 'batchnorm.py');
const OUTPUT_PATH = path.join(TEST_FOLDER, 'output.txt');
const CONFIG_PATH = path.join(TEST_FOLDER, 'config.json');
const PRECISION = 1e-3;

let batchNormContract: BatchNormalizationLayer, randomizer: RandomSeed;

async function deployBatchNormalizationLayer(inputUnits: number, momentum: number, epsilon: number) {
    const configData = abic.encode(["uint256", "uint256", "uint256"], [
        BigNumber.from(inputUnits),
        fromFloat(momentum),
        fromFloat(epsilon)
    ]);
    
    const BatchNormalizationLayer = await ethers.getContractFactory('BatchNormalizationLayer');
    batchNormContract = await BatchNormalizationLayer.deploy(configData);
    await batchNormContract.deployed();
}

async function forwardBatchNormalizationLayer(batchNormContract: BatchNormalizationLayer, input: any): Promise<any> {
    const inputTensors: Tensor[] = [Tensor.fromFloatArray(input)];
    // console.log(inputTensors[0]);
    const outputTensor = await batchNormContract.forward(inputTensors);
    const output = new Tensor(outputTensor.data, outputTensor.shapes);
    return output.toFloatArray();
}

async function getKerasBatchnormOutput(inputUnits: number, momentum: number, epsilon: number, params: number[], inputs: number[]) {
    const config = {
        inputUnits,
        momentum,
        epsilon,
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

describe('BatchNormalizationLayer', async () => {
    before(async () => {
        const seed = new Date().toLocaleString()
        console.log("Seed random: \"" + seed + "\"")
        randomizer = create(seed);
    });

    describe('1. deploy', async () => {
        it ('1.1. Manual test', async () => {
            const inputUnits = 6;
            const momentum = 0.99;
            const epsilon = 1e-3;
            await deployBatchNormalizationLayer(inputUnits, momentum, epsilon);
            expect(await batchNormContract.inputUnits()).to.equal(BigNumber.from(inputUnits));
            expect(await batchNormContract.momentum()).to.equal(BigNumber.from(fromFloat(momentum)));
            expect(await batchNormContract.epsilon()).to.equal(BigNumber.from(fromFloat(epsilon)));
        });
    });

    describe('2. appendWeights', async () => {
        it ('2.1. Manual test', async () => {
            const inputUnits = 3;
            const momentum = 0.99;
            const epsilon = 1e-3;

            const expected_gamma = recursiveFromFloat([1.0, 1.0, 1.0]);
            const expected_beta = recursiveFromFloat([0.0, 0.0, 0.0]);
            const expected_mvMean = recursiveFromFloat([1.0, 1.0, 1.0]);
            const expected_mvVar = recursiveFromFloat([0.0, 0.0, 0.0]);

            const expectedWShape = [inputUnits].map(x => ethers.BigNumber.from(x)); // only 1/4 weights (gamma, beta, mvMean, mvVar)

            const expected_gamma_Encoded = encodeData(expected_gamma.flat(Infinity));
            const expected_beta_Encoded = encodeData(expected_beta.flat(Infinity));
            const expected_mvMean_Encoded = encodeData(expected_mvMean.flat(Infinity));
            const expected_mvVar_Encoded = encodeData(expected_mvVar.flat(Infinity));

            const weightsEncoded = [expected_gamma_Encoded, expected_beta_Encoded, expected_mvMean_Encoded, expected_mvVar_Encoded].flat();

            await deployBatchNormalizationLayer(inputUnits, momentum, epsilon);

            await (await batchNormContract.appendWeights(weightsEncoded)).wait();

            const [gamma_Encoded, gamma_Shape] = (await batchNormContract.getGamma());
            const [beta_Encoded, beta_Shape] = (await batchNormContract.getBeta());
            const [mvMean_Encoded, mvMean_Shape] = (await batchNormContract.getMovingMean());
            const [mvVar_Encoded, mvVar_Shape] = (await batchNormContract.getMovingVariance());

            const gamma32x32 = deflatten(decodeData(gamma_Encoded), [inputUnits]);
            const beta32x32 = deflatten(decodeData(beta_Encoded), [inputUnits]);
            const mvMean32x32 = deflatten(decodeData(mvMean_Encoded),[inputUnits]);
            const mvVar32x32 = deflatten(decodeData(mvVar_Encoded),[inputUnits]);

            expect(isBigNumberArrayEqual(gamma32x32, expected_gamma)).to.equal(true);
            expect(isBigNumberArrayEqual(beta32x32, expected_beta)).to.equal(true);
            expect(isBigNumberArrayEqual(mvMean32x32, expected_mvMean)).to.equal(true);
            expect(isBigNumberArrayEqual(mvVar32x32, expected_mvVar)).to.equal(true);

            expect(isBigNumberArrayEqual(gamma_Shape, expectedWShape)).to.equal(true);
            expect(isBigNumberArrayEqual(beta_Shape, expectedWShape)).to.equal(true);
            expect(isBigNumberArrayEqual(mvMean_Shape, expectedWShape)).to.equal(true);
            expect(isBigNumberArrayEqual(mvVar_Shape, expectedWShape)).to.equal(true);
        });

        it ('2.2. Gas cost estimate', async () => {
            const inputUnits = 3;
            const momentum = 0.99;
            const epsilon = 1e-3;
            const chunkSize = 10000;
            await deployBatchNormalizationLayer(inputUnits, momentum, epsilon);

            const expected_gamma = recursiveFromFloat(randomFloatArray(randomizer, [inputUnits], -1, 1));
            const expected_beta = recursiveFromFloat(randomFloatArray(randomizer, [inputUnits], -1, 1));
            const expected_mvMean = recursiveFromFloat(randomFloatArray(randomizer, [inputUnits], -1, 1));
            const expected_mvVar = recursiveFromFloat(randomFloatArray(randomizer, [inputUnits], -1, 1));
            
            const expectedWShape = [inputUnits].map(x => ethers.BigNumber.from(x)); // only 1/4 weights (gamma, beta, mvMean, mvVar)

            const expected_gamma_Encoded = encodeData(expected_gamma.flat(Infinity));
            const expected_beta_Encoded = encodeData(expected_beta.flat(Infinity));
            const expected_mvMean_Encoded = encodeData(expected_mvMean.flat(Infinity));
            const expected_mvVar_Encoded = encodeData(expected_mvVar.flat(Infinity));

            const weightsEncoded = [expected_mvMean_Encoded, expected_mvVar_Encoded, expected_gamma_Encoded, expected_beta_Encoded].flat();
            
            let totalGas = 0;
            let txCount = 0;
            for(let i = 0; i < weightsEncoded.length; i += chunkSize) {
                const weightsToUpload = weightsEncoded.slice(i, i + chunkSize);
                const tx = await batchNormContract.appendWeights(weightsToUpload);
                const receipt = await tx.wait();
                totalGas += receipt['gasUsed'].toNumber();
                txCount += 1;
            }
            console.log(`Tx count: ${txCount}`);
            console.log(`Total gas cost to append ${weightsEncoded.length} params: ${totalGas}`);

            const [gamma_Encoded, gamma_Shape] = (await batchNormContract.getGamma());
            const [beta_Encoded, beta_Shape] = (await batchNormContract.getBeta());
            const [mvMean_Encoded, mvMean_Shape] = (await batchNormContract.getMovingMean());
            const [mvVar_Encoded, mvVar_Shape] = (await batchNormContract.getMovingVariance());

            const gamma32x32 = deflatten(decodeData(gamma_Encoded), [inputUnits]);
            const beta32x32 = deflatten(decodeData(beta_Encoded), [inputUnits]);
            const mvMean32x32 = deflatten(decodeData(mvMean_Encoded),[inputUnits]);
            const mvVar32x32 = deflatten(decodeData(mvVar_Encoded),[inputUnits]);

            expect(isBigNumberArrayEqual(gamma32x32, expected_gamma)).to.equal(true);
            expect(isBigNumberArrayEqual(beta32x32, expected_beta)).to.equal(true);
            expect(isBigNumberArrayEqual(mvMean32x32, expected_mvMean)).to.equal(true);
            expect(isBigNumberArrayEqual(mvVar32x32, expected_mvVar)).to.equal(true);

            expect(isBigNumberArrayEqual(gamma_Shape, expectedWShape)).to.equal(true);
            expect(isBigNumberArrayEqual(beta_Shape, expectedWShape)).to.equal(true);
            expect(isBigNumberArrayEqual(mvMean_Shape, expectedWShape)).to.equal(true);
            expect(isBigNumberArrayEqual(mvVar_Shape, expectedWShape)).to.equal(true);
        });
    });
    
    describe('3. forward', async () => {
        it('3.1. Manual test ', async () => {
            const inputUnits = 3;
            const momentum = 0.99;
            const epsilon = 1e-3;

            const expected_gamma = recursiveFromFloat([1.0, 1.0, 1.0]);
            const expected_beta = recursiveFromFloat([0.0, 0.0, 0.0]);
            const expected_mvVar = recursiveFromFloat([1.0, 1.0, 1.0]);
            const expected_mvMean = recursiveFromFloat([0.0, 0.0, 0.0]);

            const expected_gamma_Encoded = encodeData(expected_gamma.flat(Infinity));
            const expected_beta_Encoded = encodeData(expected_beta.flat(Infinity));
            const expected_mvMean_Encoded = encodeData(expected_mvMean.flat(Infinity));
            const expected_mvVar_Encoded = encodeData(expected_mvVar.flat(Infinity));

            const weightsEncoded = [expected_gamma_Encoded, expected_beta_Encoded, expected_mvMean_Encoded, expected_mvVar_Encoded].flat();

            const input = [
                [[6., 6., 5.],
                [5., 9., 0.],
                [8., 6., 2.]],
        
                [[2., 1., 6.],
                [6., 6., 6.],
                [4., 7., 5.]],
        
                [[1., 1., 9.],
                [3., 0., 2.],
                [0., 8., 9.]]
            ];

            const expectedOutput = [
                [[5.9970026, 5.9970026, 4.997502 ],
                [4.997502 , 8.995503 , 0.       ],
                [7.996003 , 5.9970026, 1.9990008]],
        
                [[1.9990008, 0.9995004, 5.9970026],
                [5.9970026, 5.9970026, 5.9970026],
                [3.9980016, 6.996503 , 4.997502 ]],
        
                [[0.9995004, 0.9995004, 8.995503 ],
                [2.9985013, 0.       , 1.9990008],
                [0.       , 7.996003 , 8.995503 ]]
            ];

            await deployBatchNormalizationLayer(inputUnits, momentum, epsilon);
            console.log("Model deployed");

            await (await batchNormContract.appendWeights(weightsEncoded)).wait();
            console.log("Weights appended");

            const output = await forwardBatchNormalizationLayer(batchNormContract, input);

            expect(isFloatArrayEqual(output, expectedOutput, PRECISION)).to.equal(true);
        });

        it ('3.2. Random test ', async () => {
            const inputUnits = randomizer.intBetween(1,10);
            const momentum = randomizer.floatBetween(0,1);
            const epsilon = randomizer.floatBetween(1e-5,1e-3);

            const gamma = randomFloatArray(randomizer, [inputUnits], 0, 1.0);
            const beta = randomFloatArray(randomizer,[inputUnits], -1.0, 1.0);
            const mvMean = randomFloatArray(randomizer,[inputUnits], -1.0, 1.0);
            const mvVar = randomFloatArray(randomizer, [inputUnits], 0, 1.0);

            const inputHeight = 64;
            const inputWidth = 64;
            const inputs = randomFloatArray(randomizer, [inputHeight, inputWidth, inputUnits], -5.0, 5.0);

            const params = [gamma, beta, mvMean, mvVar].flat();
            const expectedOutput = await getKerasBatchnormOutput(inputUnits, momentum, epsilon, params, inputs);

            const gamma32x32 = recursiveFromFloat(gamma);
            const beta32x32 = recursiveFromFloat(beta);
            const mvMean32x32 = recursiveFromFloat(mvMean);
            const mvVar32x32 = recursiveFromFloat(mvVar);

            const expected_gamma_Encoded = encodeData(gamma32x32.flat(Infinity));
            const expected_beta_Encoded = encodeData(beta32x32.flat(Infinity));
            const expected_mvMean_Encoded = encodeData(mvMean32x32.flat(Infinity));
            const expected_mvVar_Encoded = encodeData(mvVar32x32.flat(Infinity));

            const weightsEncoded = [expected_gamma_Encoded, expected_beta_Encoded, expected_mvMean_Encoded, expected_mvVar_Encoded].flat();

            await deployBatchNormalizationLayer(inputUnits, momentum, epsilon);
            console.log("Model deployed");

            await (await batchNormContract.appendWeights(weightsEncoded)).wait();
            console.log("Weights appended");

            const output = await forwardBatchNormalizationLayer(batchNormContract, inputs);

            expect(isFloatArrayEqual(output, expectedOutput, PRECISION)).to.equal(true);
        });
    });
});
