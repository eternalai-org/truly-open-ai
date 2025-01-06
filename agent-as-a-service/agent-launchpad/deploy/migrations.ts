import hre, { ethers } from "hardhat";
import helpers from "./helpers";
import { NonfungiblePositionManager, NonfungibleTokenPositionDescriptor, QuoterV2, SwapRouter, UniswapV3Factory, UniswapV3Broker } from "../typechain";

const { waitForDeploy, waitForTx, loadDB, saveDB, upgradeContract } = helpers;

async function main() {
    await deploy();
}

export default deploy;

async function deploy() {
    const network = hre.network.name;
    const [deployer] = await ethers.getSigners()
    let deployData = (await loadDB(network))
    // 
    if (deployData.proxyAdmin.address == undefined || deployData.proxyAdmin.address == '') {
        let ProxyAdmin = await hre.ethers.getContractFactory('ProxyAdmin');
        let proxyAdmin = await waitForDeploy(await ProxyAdmin.deploy(), 'ProxyAdmin');
        {
            deployData.proxyAdmin.address = proxyAdmin.address;
            deployData = (await saveDB(network, deployData))
        }
        console.log('proxyAdmin is deployed', proxyAdmin.address)
    }
    let proxyAdmin = await hre.ethers.getContractAt('ProxyAdmin', deployData.proxyAdmin.address);
    // uniswapV3Factory
    if (deployData.uniswapV3Factory.implAddress == undefined || deployData.uniswapV3Factory.implAddress == '') {
        let UniswapV3Factory = await hre.ethers.getContractFactory('UniswapV3Factory');
        let uniswapV3Factory = await waitForDeploy(await UniswapV3Factory.deploy(), 'UniswapV3Factory');
        {
            deployData.uniswapV3Factory.implAddress = uniswapV3Factory.address;
            deployData = (await saveDB(network, deployData))
        }
        console.log('uniswapV3Factory is deployed', uniswapV3Factory.address)
    }
    // uniswapV3Pool
    if (deployData.uniswapV3Pool.implAddress == undefined || deployData.uniswapV3Pool.implAddress == '') {
        let UniswapV3Pool = await hre.ethers.getContractFactory('UniswapV3Pool');
        let uniswapV3Pool = await waitForDeploy(await UniswapV3Pool.deploy(), 'UniswapV3Pool');
        {
            deployData.uniswapV3Pool.implAddress = uniswapV3Pool.address;
            deployData = (await saveDB(network, deployData))
        }
        console.log('uniswapV3Pool is deployed', uniswapV3Pool.address)
    }
    if (deployData.uniswapV3Broker.implAddress == undefined || deployData.uniswapV3Broker.implAddress == '') {
        let UniswapV3Broker = await hre.ethers.getContractFactory('UniswapV3Broker');
        let uniswapV3Broker = await waitForDeploy(await UniswapV3Broker.deploy(), 'UniswapV3Broker');
        {
            deployData.uniswapV3Broker.implAddress = uniswapV3Broker.address;
            deployData = (await saveDB(network, deployData))
        }
        console.log('uniswapV3Broker is deployed', uniswapV3Broker.address)
    }
    if (deployData.nftDescriptor.implAddress == undefined || deployData.nftDescriptor.implAddress == '') {
        let NftDescriptor = await hre.ethers.getContractFactory('NFTDescriptor');
        let nftDescriptor = await waitForDeploy(await NftDescriptor.deploy(), 'NFTDescriptor');
        {
            deployData.nftDescriptor.implAddress = nftDescriptor.address;
            deployData = (await saveDB(network, deployData))
        }
        console.log('nftDescriptor is deployed', nftDescriptor.address)
    }
    // nonfungibleTokenPositionDescriptor: ContractData,
    if (deployData.nonfungibleTokenPositionDescriptor.implAddress == undefined || deployData.nonfungibleTokenPositionDescriptor.implAddress == '') {
        let NonfungibleTokenPositionDescriptor = await hre.ethers.getContractFactory('NonfungibleTokenPositionDescriptor', {
            libraries: {
                NFTDescriptor: deployData.nftDescriptor.implAddress,
            },
        });
        let nonfungibleTokenPositionDescriptor = await waitForDeploy(await NonfungibleTokenPositionDescriptor.deploy(), 'NonfungibleTokenPositionDescriptor');
        {
            deployData.nonfungibleTokenPositionDescriptor.implAddress = nonfungibleTokenPositionDescriptor.address;
            deployData = (await saveDB(network, deployData))
        }
        console.log('nonfungibleTokenPositionDescriptor is deployed', nonfungibleTokenPositionDescriptor.address)
    }
    // nonfungiblePositionManager: ContractData,
    if (deployData.nonfungiblePositionManager.implAddress == undefined || deployData.nonfungiblePositionManager.implAddress == '') {
        let NonfungiblePositionManager = await hre.ethers.getContractFactory('NonfungiblePositionManager', {
            libraries: {
                UniswapV3Broker: deployData.uniswapV3Broker.implAddress,
            },
        });
        let nonfungiblePositionManager = await waitForDeploy(await NonfungiblePositionManager.deploy(), 'NonfungiblePositionManager');
        {
            deployData.nonfungiblePositionManager.implAddress = nonfungiblePositionManager.address;
            deployData = (await saveDB(network, deployData))
        }
        console.log('nonfungiblePositionManager is deployed', nonfungiblePositionManager.address)
    }
    // swapRouter: ContractData,
    if (deployData.swapRouter.implAddress == undefined || deployData.swapRouter.implAddress == '') {
        let SwapRouter = await hre.ethers.getContractFactory('SwapRouter');
        let swapRouter = await waitForDeploy(await SwapRouter.deploy(), 'SwapRouter');
        {
            deployData.swapRouter.implAddress = swapRouter.address;
            deployData = (await saveDB(network, deployData))
        }
        console.log('swapRouter is deployed', swapRouter.address)
    }
    // quoterV2: ContractData,
    if (deployData.quoterV2.implAddress == undefined || deployData.quoterV2.implAddress == '') {
        let QuoterV2 = await hre.ethers.getContractFactory('QuoterV2');
        let quoterV2 = await waitForDeploy(await QuoterV2.deploy(), 'QuoterV2');
        {
            deployData.quoterV2.implAddress = quoterV2.address;
            deployData = (await saveDB(network, deployData))
        }
        console.log('quoterV2 is deployed', quoterV2.address)
    }
    // proxy
    const TransparentUpgradeableProxy = await hre.ethers.getContractFactory('TransparentUpgradeableProxy');
    // uniswapV3Factory: ContractData,
    if (deployData.uniswapV3Factory.address == undefined || deployData.uniswapV3Factory.address == '') {
        var uniswapV3Factory = await hre.ethers.getContractAt('UniswapV3Factory', deployData.uniswapV3Factory.implAddress) as UniswapV3Factory;
        var initializeData = uniswapV3Factory.interface.encodeFunctionData('initialize');
        var transparentUpgradeableProxy = await waitForDeploy(
            await TransparentUpgradeableProxy.deploy(
                deployData.uniswapV3Factory.implAddress,
                deployData.proxyAdmin.address,
                initializeData,
            )
        );
        {
            deployData.uniswapV3Factory.address = transparentUpgradeableProxy.address;
            deployData = (await saveDB(network, deployData))
            console.log('uniswapV3Factory TransparentUpgradeableProxy is deployed', transparentUpgradeableProxy.address)
        }
    }
    // nonfungibleTokenPositionDescriptor: ContractData,
    if (deployData.nonfungibleTokenPositionDescriptor.address == undefined || deployData.nonfungibleTokenPositionDescriptor.address == '') {
        var nonfungibleTokenPositionDescriptor = await hre.ethers.getContractAt('NonfungibleTokenPositionDescriptor', deployData.nonfungibleTokenPositionDescriptor.implAddress) as NonfungibleTokenPositionDescriptor;
        var initializeData = nonfungibleTokenPositionDescriptor.interface.encodeFunctionData('initialize', [deployData.WETH.address, ethers.utils.formatBytes32String('ETH')]);
        var transparentUpgradeableProxy = await waitForDeploy(
            await TransparentUpgradeableProxy.deploy(
                deployData.nonfungibleTokenPositionDescriptor.implAddress,
                deployData.proxyAdmin.address,
                initializeData,
            )
        );
        {
            deployData.nonfungibleTokenPositionDescriptor.address = transparentUpgradeableProxy.address;
            deployData = (await saveDB(network, deployData))
            console.log('nonfungibleTokenPositionDescriptor TransparentUpgradeableProxy is deployed', transparentUpgradeableProxy.address)
        }
    }
    // nonfungiblePositionManager: ContractData,
    if (deployData.nonfungiblePositionManager.address == undefined || deployData.nonfungiblePositionManager.address == '') {
        var nonfungiblePositionManager = await hre.ethers.getContractAt('NonfungiblePositionManager', deployData.nonfungiblePositionManager.implAddress) as NonfungiblePositionManager;
        var initializeData = nonfungiblePositionManager.interface.encodeFunctionData('initialize', [deployData.uniswapV3Factory.address, deployData.WETH.address, deployData.nonfungibleTokenPositionDescriptor.address]);
        var transparentUpgradeableProxy = await waitForDeploy(
            await TransparentUpgradeableProxy.deploy(
                deployData.nonfungiblePositionManager.implAddress,
                deployData.proxyAdmin.address,
                initializeData,
            )
        );
        {
            deployData.nonfungiblePositionManager.address = transparentUpgradeableProxy.address;
            deployData = (await saveDB(network, deployData))
            console.log('nonfungiblePositionManager TransparentUpgradeableProxy is deployed', transparentUpgradeableProxy.address)
        }
    }
    // swapRouter: ContractData,
    if (deployData.swapRouter.address == undefined || deployData.swapRouter.address == '') {
        var swapRouter = await hre.ethers.getContractAt('SwapRouter', deployData.swapRouter.implAddress) as SwapRouter;
        var initializeData = swapRouter.interface.encodeFunctionData('initialize', [deployData.uniswapV3Factory.address, deployData.WETH.address]);
        var transparentUpgradeableProxy = await waitForDeploy(
            await TransparentUpgradeableProxy.deploy(
                deployData.swapRouter.implAddress,
                deployData.proxyAdmin.address,
                initializeData,
            )
        );
        {
            deployData.swapRouter.address = transparentUpgradeableProxy.address;
            deployData = (await saveDB(network, deployData))
            console.log('swapRouter TransparentUpgradeableProxy is deployed', transparentUpgradeableProxy.address)
        }
    }
    // quoterV2: ContractData,
    if (deployData.quoterV2.address == undefined || deployData.quoterV2.address == '') {
        var quoterV2 = await hre.ethers.getContractAt('QuoterV2', deployData.quoterV2.implAddress) as QuoterV2;
        var initializeData = quoterV2.interface.encodeFunctionData('initialize', [deployData.uniswapV3Factory.address, deployData.WETH.address]);
        var transparentUpgradeableProxy = await waitForDeploy(
            await TransparentUpgradeableProxy.deploy(
                deployData.quoterV2.implAddress,
                deployData.proxyAdmin.address,
                initializeData,
            )
        );
        {
            deployData.quoterV2.address = transparentUpgradeableProxy.address;
            deployData = (await saveDB(network, deployData))
            console.log('quoterV2 TransparentUpgradeableProxy is deployed', transparentUpgradeableProxy.address)
        }
    }
    // 
    {
        await upgradeContract(proxyAdmin, deployData.uniswapV3Factory.address, deployData.uniswapV3Factory.implAddress)
        await upgradeContract(proxyAdmin, deployData.nonfungibleTokenPositionDescriptor.address, deployData.nonfungibleTokenPositionDescriptor.implAddress)
        await upgradeContract(proxyAdmin, deployData.nonfungiblePositionManager.address, deployData.nonfungiblePositionManager.implAddress)
        await upgradeContract(proxyAdmin, deployData.swapRouter.address, deployData.swapRouter.implAddress)
        await upgradeContract(proxyAdmin, deployData.quoterV2.address, deployData.quoterV2.implAddress)
        // 
        var uniswapV3Factory = await hre.ethers.getContractAt('UniswapV3Factory', deployData.uniswapV3Factory.address) as UniswapV3Factory;
        if ((await uniswapV3Factory.getUniswapV3PoolImplementation()).toLowerCase() != deployData.uniswapV3Pool.implAddress.toLowerCase()) {
            await waitForTx(
                await uniswapV3Factory.setUniswapV3PoolImplementation(deployData.uniswapV3Pool.implAddress),
                `uniswapV3Factory.setUniswapV3PoolImplementation(${deployData.uniswapV3Pool.implAddress})`
            )
        }
        if ((await uniswapV3Factory.feeProtocol()) != 17) {
            await waitForTx(
                await uniswapV3Factory.setFeeProtocol(1, 1),
                `uniswapV3Factory.setFeeProtocol(1,1)`
            )
        }
        const feeTo = deployer.address
        if ((await uniswapV3Factory.feeTo()).toLowerCase() != feeTo.toLowerCase()) {
            await waitForTx(
                await uniswapV3Factory.setFeeTo(feeTo),
                `swapRouter.setFeeTo(${feeTo})`
            )
        }
        var nonfungibleTokenPositionDescriptor = await hre.ethers.getContractAt('NonfungibleTokenPositionDescriptor', deployData.nonfungibleTokenPositionDescriptor.address) as NonfungibleTokenPositionDescriptor;
        if ((await nonfungibleTokenPositionDescriptor.WETH()).toLowerCase() != deployData.WETH.address.toLowerCase()) {
            await waitForTx(
                await nonfungibleTokenPositionDescriptor.setWETH(deployData.WETH.address, ethers.utils.formatBytes32String('ETH')),
                `nonfungibleTokenPositionDescriptor.setWETH(${deployData.WETH.address})`
            )
        }
        const nonfungiblePositionManager = await ethers.getContractAt("NonfungiblePositionManager", deployData.nonfungiblePositionManager.address) as NonfungiblePositionManager;
        if ((await nonfungiblePositionManager.WETH()).toLowerCase() != deployData.WETH.address.toLowerCase()) {
            await waitForTx(
                await nonfungiblePositionManager.setWETH(deployData.WETH.address),
                `nonfungiblePositionManager.setWETH(${deployData.WETH.address})`
            )
        }
        const swapRouter = await ethers.getContractAt("SwapRouter", deployData.swapRouter.address) as SwapRouter;
        if ((await swapRouter.WETH()).toLowerCase() != deployData.WETH.address.toLowerCase()) {
            await waitForTx(
                await swapRouter.setWETH(deployData.WETH.address),
                `swapRouter.setWETH(${deployData.WETH.address})`
            )
        }
        const quoterV2 = await ethers.getContractAt("QuoterV2", deployData.quoterV2.address) as QuoterV2;
        if ((await quoterV2.WETH()).toLowerCase() != deployData.WETH.address.toLowerCase()) {
            await waitForTx(
                await quoterV2.setWETH(deployData.WETH.address),
                `quoterV2.setWETH(${deployData.WETH.address})`
            )
        }
    }
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});