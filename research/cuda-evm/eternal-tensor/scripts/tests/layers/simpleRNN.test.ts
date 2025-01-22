import { expect, assert } from 'chai';
import { BigNumber } from 'ethers';
import { ethers } from 'hardhat';
import { SimpleRNNLayer } from '../../../typechain-types';
import { RandomSeed, create } from 'random-seed';
import { Activation } from '../../libraries/modelLib'
import { fromFloat } from '../../libraries/utils';

const abic = ethers.utils.defaultAbiCoder;

let simpleRNNContract: SimpleRNNLayer, randomizer: RandomSeed;

async function deploySimpleRNNLayer(actv: Activation, inputUnits: number, outputUnits: number) {
    const configData = abic.encode(["uint8", "uint256", "uint256"], [
        actv,
        BigNumber.from(outputUnits),
        BigNumber.from(inputUnits),
    ]);
    
    const SimpleRNNLayer = await ethers.getContractFactory('SimpleRNNLayer');
    simpleRNNContract = await SimpleRNNLayer.deploy(configData);
    await simpleRNNContract.deployed();
}

describe('SimpleRNNLayer', async () => {
    before(async () => {
        const seed = new Date().toLocaleString()
        console.log("Seed random: \"" + seed + "\"")
        randomizer = create(seed);
    });

    describe('1. deploy', async () => {
        it('1.1. deploy layer correctly', async () => {
            const inputUnits = 128;
            const outputUnits = 32;
            await deploySimpleRNNLayer(Activation.relu, inputUnits, outputUnits);
            expect(await simpleRNNContract.inputDim()).to.equal(BigNumber.from(inputUnits));
            expect(await simpleRNNContract.outputDim()).to.equal(BigNumber.from(outputUnits));
            expect(await simpleRNNContract.activation()).to.equal(Activation.relu);
        });
    });

    describe('2. appendWeights', async () => {
        it('2.1. Single test', async () => {
            const inputUnits = 3;
            const outputUnits = 2;
            const weights = [1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0]
                .map(x => fromFloat(x));

            await deploySimpleRNNLayer(Activation.relu, inputUnits, outputUnits);

            await (await simpleRNNContract.appendWeights(weights, 0)).wait();

            expect(await simpleRNNContract.getWeight(0, 0)).to.equal(fromFloat(1.0));
            expect(await simpleRNNContract.getWeight(0, 1)).to.equal(fromFloat(2.0));
            expect(await simpleRNNContract.getWeight(1, 0)).to.equal(fromFloat(3.0));
            expect(await simpleRNNContract.getWeight(1, 1)).to.equal(fromFloat(4.0));
            expect(await simpleRNNContract.getWeight(2, 0)).to.equal(fromFloat(5.0));
            expect(await simpleRNNContract.getWeight(2, 1)).to.equal(fromFloat(6.0));

            expect(await simpleRNNContract.getRecurrentWeight(0, 0)).to.equal(fromFloat(7.0));
            expect(await simpleRNNContract.getRecurrentWeight(0, 1)).to.equal(fromFloat(8.0));
            expect(await simpleRNNContract.getRecurrentWeight(1, 0)).to.equal(fromFloat(9.0));
            expect(await simpleRNNContract.getRecurrentWeight(1, 1)).to.equal(fromFloat(10.0));

            expect(await simpleRNNContract.getBias(0)).to.equal(fromFloat(11.0));
            expect(await simpleRNNContract.getBias(1)).to.equal(fromFloat(12.0));
        });

        it('2.2. Gas cost estimate', async () => {
            const inputUnits = 128;
            const outputUnits = 32;
            const chunkSize = 10000;

            await deploySimpleRNNLayer(Activation.relu, inputUnits, outputUnits);

            const weightLen = (await simpleRNNContract.getParamsCount()).toNumber();
            const weights: BigNumber[] = [];
            for(let i = 0; i < weightLen; ++i) {
                weights.push(fromFloat(randomizer.floatBetween(-1, 1)));
            }
            
            let totalGas = 0;
            for(let i = 0; i < weightLen; i += chunkSize) {
                const weightsToUpload = weights.slice(i, i + chunkSize);
                const tx = await simpleRNNContract.appendWeights(weightsToUpload, 0);
                const receipt = await tx.wait();
                totalGas += receipt['gasUsed'].toNumber();
            }
            console.log(`Total gas cost to append ${weights.length} params: ${totalGas}`);
        });
    });
});
