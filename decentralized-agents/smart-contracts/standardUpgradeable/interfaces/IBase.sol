// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title IGPUManager
 * @dev Interface for managing GPU resources.
 */
interface IGPUManager {
    /**
     * @notice Retrieves the minimum fee required to use a specific model.
     * @param modelId The ID of the model.
     * @return The minimum fee required to use the model.
     */
    function getMinFeeToUse(uint32 modelId) external view returns (uint256);
}

/**
 * @title IInferable
 * @dev Interface for performing inference operations on models.
 */
interface IInferable {
    /**
     * @notice Performs an inference operation on a specified model.
     * @param modelId The ID of the model.
     * @param data The input data for the inference.
     * @param creator The address of the creator initiating the inference.
     * @return inferenceId The ID of the inference operation.
     */
    function infer(
        uint32 modelId,
        bytes calldata data,
        address creator
    ) external returns (uint64 inferenceId);

    /**
     * @notice Performs an inference operation on a specified model with an additional rawFlag.
     * @param modelId The ID of the model.
     * @param data The input data for the inference.
     * @param creator The address of the creator initiating the inference.
     * @param rawFlag The flag to indicate the format of the calldata and the result of the inference.
     *                  If rawFlag is true, the calldata and inference result are in raw format.
     *                  If rawFlag is false, the calldata and inference result are in IPFS link format.
     * @return inferenceId The ID of the inference operation.
     */
    function infer(
        uint32 modelId,
        bytes calldata data,
        address creator,
        bool rawFlag
    ) external returns (uint64 inferenceId);
}
