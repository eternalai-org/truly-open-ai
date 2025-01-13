import { EndpointId } from '@layerzerolabs/lz-definitions';
const ape_mainnetContract = {
    eid: EndpointId.APE_V2_MAINNET,
    contractName: 'Questioner',
};
const avax_mainnetContract = {
    eid: EndpointId.AVALANCHE_V2_MAINNET,
    contractName: 'Questioner',
};
export default {
    contracts: [
        { contract: ape_mainnetContract },
        { contract: avax_mainnetContract },
    ],
    connections: [
        {
            from: ape_mainnetContract,
            to: avax_mainnetContract,
            config: {
                sendLibrary: '0xC39161c743D0307EB9BCc9FEF03eeb9Dc4802de7',
                receiveLibraryConfig: {
                    receiveLibrary:
                        '0xe1844c5D63a9543023008D332Bd3d2e6f1FE1043',
                    gracePeriod: 0,
                },
                sendConfig: {
                    executorConfig: {
                        maxMessageSize: 10000,
                        executor: '0xcCE466a522984415bC91338c232d98869193D46e',
                    },
                    ulnConfig: {
                        confirmations: 5,
                        requiredDVNs: [
                            '0x6788f52439aca6bff597d3eec2dc9a44b8fee842',
                        ],
                        optionalDVNs: [],
                        optionalDVNThreshold: 0,
                    },
                },
                receiveConfig: {
                    ulnConfig: {
                        confirmations: 5,
                        requiredDVNs: [
                            '0x6788f52439aca6bff597d3eec2dc9a44b8fee842',
                        ],
                        optionalDVNs: [],
                        optionalDVNThreshold: 0,
                    },
                },
            },
        },
        {
            from: avax_mainnetContract,
            to: ape_mainnetContract,
            config: {
                sendLibrary: '0x197D1333DEA5Fe0D6600E9b396c7f1B1cFCc558a',
                receiveLibraryConfig: {
                    receiveLibrary:
                        '0xbf3521d309642FA9B1c91A08609505BA09752c61',
                    gracePeriod: 0,
                },
                sendConfig: {
                    executorConfig: {
                        maxMessageSize: 10000,
                        executor: '0x90E595783E43eb89fF07f63d27B8430e6B44bD9c',
                    },
                    ulnConfig: {
                        confirmations: 5,
                        requiredDVNs: [
                            '0x962f502a63f5fbeb44dc9ab932122648e8352959',
                        ],
                        optionalDVNs: [],
                        optionalDVNThreshold: 0,
                    },
                },
                receiveConfig: {
                    ulnConfig: {
                        confirmations: 5,
                        requiredDVNs: [
                            '0x962f502a63f5fbeb44dc9ab932122648e8352959',
                        ],
                        optionalDVNs: [],
                        optionalDVNThreshold: 0,
                    },
                },
            },
        },
    ],
};
