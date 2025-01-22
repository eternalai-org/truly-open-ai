import { expect, assert } from 'chai';
import { BigNumber } from 'ethers';
import { ethers } from 'hardhat';
import { EmbeddingLayer } from '../../../typechain-types';
import { RandomSeed, create } from 'random-seed';
import { Activation } from '../../libraries/modelLib'
import { fromFloat } from '../../libraries/utils';

const abic = ethers.utils.defaultAbiCoder;

let embeddingContract: EmbeddingLayer, randomizer: RandomSeed;

async function deployEmbeddingLayer(inputUnits: number, outputUnits: number) {
    const configData = abic.encode(["uint256", "uint256"], [
        BigNumber.from(inputUnits),
        BigNumber.from(outputUnits),
    ]);
    
    const EmbeddingLayer = await ethers.getContractFactory('EmbeddingLayer');
    embeddingContract = await EmbeddingLayer.deploy(configData);
    await embeddingContract.deployed();
}

describe('EmbeddingLayer', async () => {
    before(async () => {
        const seed = new Date().toLocaleString()
        console.log("Seed random: \"" + seed + "\"")
        randomizer = create(seed);
    });

    describe('1. deploy', async () => {
        it('1.1. deploy layer correctly', async () => {
            const inputUnits = 1000;
            const outputUnits = 128;
            await deployEmbeddingLayer(inputUnits, outputUnits);
            expect(await embeddingContract.inputDim()).to.equal(BigNumber.from(inputUnits));
            expect(await embeddingContract.outputDim()).to.equal(BigNumber.from(outputUnits));
        });
    });

    describe('2. appendWeights', async () => {
        it('2.1. Single test', async () => {
            const inputUnits = 3;
            const outputUnits = 2;
            const weights = [1.0, 2.0, 3.0, 4.0, 5.0, 6.0]
                .map(x => fromFloat(x));

            await deployEmbeddingLayer(inputUnits, outputUnits);

            await (await embeddingContract.appendWeights(weights, 0)).wait();

            expect(await embeddingContract.getWeight(0, 0)).to.equal(fromFloat(1.0));
            expect(await embeddingContract.getWeight(0, 1)).to.equal(fromFloat(2.0));
            expect(await embeddingContract.getWeight(1, 0)).to.equal(fromFloat(3.0));
            expect(await embeddingContract.getWeight(1, 1)).to.equal(fromFloat(4.0));
            expect(await embeddingContract.getWeight(2, 0)).to.equal(fromFloat(5.0));
            expect(await embeddingContract.getWeight(2, 1)).to.equal(fromFloat(6.0));
        });

        it('2.2. Gas cost estimate', async () => {
            const inputUnits = 1000;
            const outputUnits = 128;
            const chunkSize = 10000;

            await deployEmbeddingLayer(inputUnits, outputUnits);

            const weightLen = (await embeddingContract.getParamsCount()).toNumber();
            const weights: BigNumber[] = [];
            for(let i = 0; i < weightLen; ++i) {
                weights.push(fromFloat(randomizer.floatBetween(-1, 1)));
            }
            
            let totalGas = 0;
            for(let i = 0; i < weightLen; i += chunkSize) {
                const weightsToUpload = weights.slice(i, i + chunkSize);
                const tx = await embeddingContract.appendWeights(weightsToUpload, 0);
                const receipt = await tx.wait();
                totalGas += receipt['gasUsed'].toNumber();
            }
            console.log(`Total gas cost to append ${weights.length} params: ${totalGas}`);
        });
    });
});
