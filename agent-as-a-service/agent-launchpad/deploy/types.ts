type ContractData = {
    address: string,
    implAddress: string,
}

type DeployData = {
    verifiedContracts: any,
    proxyAdmin: ContractData,
    uniswapV3Factory: ContractData,
    uniswapV3Pool: ContractData,
    uniswapV3Broker: ContractData,
    nftDescriptor: ContractData,
    nonfungibleTokenPositionDescriptor: ContractData,
    nonfungiblePositionManager: ContractData,
    swapRouter: ContractData,
    quoterV2: ContractData,
    WETH: ContractData,
    EAI: ContractData,
    agentShares: ContractData,
}
