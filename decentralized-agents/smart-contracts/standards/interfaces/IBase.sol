// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IGPUManager {
    function getMinFeeToUse(uint32 modelId) external view returns (uint256);
}

interface IInferable {
    function infer(
        uint256 modelId,
        bytes calldata data,
        address creator
    ) external returns (uint256 inferenceId);

    function infer(
        uint256 modelId,
        bytes calldata data,
        address creator,
        bool flag
    ) external returns (uint256 inferenceId);
}
