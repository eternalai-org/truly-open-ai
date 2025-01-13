import assert from 'assert';

import { type DeployFunction } from 'hardhat-deploy/types';
import { ethers, network } from 'hardhat';

// TODO declare your contract name here
const contractName = 'Questioner';

const deploy: DeployFunction = async (hre) => {
    const config = hre.network.config as any;
    const { getNamedAccounts, deployments } = hre;

    const { deploy } = deployments;
    const { deployer } = await getNamedAccounts();

    assert(deployer, 'Missing named deployer account');

    console.log(`Network: ${hre.network.name}`);
    console.log(`Deployer: ${deployer}`);

    // This is an external deployment pulled in from @layerzerolabs/lz-evm-sdk-v2
    //
    // @layerzerolabs/toolbox-hardhat takes care of plugging in the external deployments
    // from @layerzerolabs packages based on the configuration in your hardhat config
    //
    // For this to work correctly, your network config must define an eid property
    // set to `EndpointId` as defined in @layerzerolabs/lz-definitions
    //
    // For example:
    //
    // networks: {
    //   fuji: {
    //     ...
    //     eid: EndpointId.AVALANCHE_V2_TESTNET
    //   }
    // }
    const endpointV2Deployment = await hre.deployments.get('EndpointV2');
    // const promptScheduler = '0xcbeB849347e9fD09Ae125D2d280D30f68504356b'; //AVAX
    // const hybridModel = '0x7169Ba2F90Ee0af074621262c1531d568C857E1B'; //AVAX
    // const modelName = 'NousResearch/Hermes-3-Llama-3.1-70B-FP8'; //AVAX
    // const promptScheduler = '0x444659CfFf9DBa347d1298d2710D3939f1aD0286'; //APE
    // const hybridModel = '0xadeE3A52C7c2c4Aa6095058f770A917C470e92f1'; //APE
    // const modelName = 'unsloth/Llama-3.3-70B-Instruct-bnb-4bit'; //APE

    const promptScheduler = config.promptScheduler;
    const hybridModel = config.hybridModel;
    const modelName = config.modelName;

    console.log(`-- EndpointV2: ${endpointV2Deployment.address}`);
    console.log(`-- PromptScheduler: ${promptScheduler}`);
    console.log(`-- HybridModel: ${hybridModel}`);
    console.log(`-- ModelName: ${modelName}`);

    const { address } = await deploy(contractName, {
        from: deployer,
        args: [
            endpointV2Deployment.address, // LayerZero's EndpointV2 address
            deployer, // owner
            promptScheduler,
            hybridModel,
            modelName,
        ],
        log: true,
        skipIfAlreadyDeployed: false,
    });

    console.log(
        `Deployed contract: ${contractName}, network: ${hre.network.name}, address: ${address}`
    );
};

deploy.tags = [contractName];

export default deploy;
