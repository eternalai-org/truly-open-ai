// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IInferable {
    function infer(
        uint32 modelId,
        bytes calldata data,
        address creator
    ) external returns (uint64 inferenceId);

    function infer(
        uint32 modelId,
        bytes calldata data,
        address creator,
        bool flag
    ) external returns (uint64 inferenceId);
}
