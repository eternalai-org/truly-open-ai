// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

import "./IBase.sol";

/**
 * @title IAI20 Interface
 * @author EAI
 * @dev Interface of the AI20 standard.
 * This interface defines the structure and functionality for an ERC20 token that is specifically designed for AI applications.
 * It includes events for updating model IDs, prompt schedulers, gpu manager, and token fees, as well as events for agent URI updates, agent data updates, and inference performances.
 * It also includes functions for getting mission data and topping up the pool balance, as well as two variants of the `infer` function for executing inference requests.
 */
interface IAI20Upgradeable {
    /**
     * @dev Structure to hold token metadata.
     * @param fee The fee associated with the token.
     * @param sysPrompts Mapping of system prompts to their corresponding data.
     */
    struct TokenMetaData {
        uint256 fee;
        mapping(string => bytes[]) sysPrompts;
    }

    /**
     * @dev Emitted when the model ID is updated.
     * @param modelId The new model ID.
     */
    event ModelIdUpdate(uint32 modelId);

    /**
     * @dev Emitted when the prompt scheduler is updated.
     * @param promptScheduler The address of the new prompt scheduler.
     */
    event PromptSchedulerUpdate(address promptScheduler);

    /**
     * @dev Emitted when the GPU manager is updated.
     * @param gpuManager The address of the new GPU manager.
     */
    event GPUManagerUpdate(address gpuManager);

    /**
     * @dev Emitted when the token fee is updated.
     * @param tokenFee The address of the new token fee.
     */
    event TokenFeeUpdate(address tokenFee);

    /**
     * @dev Emitted when the agent URI is updated.
     * @param uri The new URI of the agent.
     */
    event AgentURIUpdate(string uri);

    /**
     * @dev Emitted when agent data is updated.
     * @param promptIndex The index of the prompt being updated.
     * @param oldSysPrompt The old system prompt data.
     * @param newSysPrompt The new system prompt data.
     */
    event AgentDataUpdate(
        uint256 promptIndex,
        bytes oldSysPrompt,
        bytes newSysPrompt
    );

    /**
     * @dev Emitted when new agent data is added.
     * @param sysPrompt The new system prompt data.
     */
    event AgentDataAddNew(bytes[] sysPrompt);

    /**
     * @dev Emitted when the agent fee is updated.
     * @param fee The new agent fee.
     */
    event AgentFeeUpdate(uint fee);

    /**
     * @dev Emitted when an inference is performed.
     * @param caller The address of the caller.
     * @param data The data related to the inference.
     * @param fee The fee paid for the inference.
     * @param externalData External data related to the inference.
     * @param inferenceId The ID of the inference.
     */
    event InferencePerformed(
        address indexed caller,
        bytes data,
        uint fee,
        string externalData,
        uint256 inferenceId
    );

    /**
     * @dev Emitted when the pool balance is topped up.
     * @param caller The address of the caller.
     * @param amount The amount of tokens used to top up the pool balance.
     */
    event TopUpPoolBalance(address caller, uint256 amount);

    /**
     * @dev Error thrown when there are insufficient funds.
     */
    error InsufficientFunds();

    /**
     * @dev Error thrown when user request infer with fee lower than agent fee.
     */
    error InvalidAgentFee();

    /**
     * @dev Error thrown when the agent data null.
     */
    error InvalidAgentData();

    /**
     * @dev Error thrown when the data is invalid.
     */
    error InvalidData();

    /**
     * @dev Error thrown when the agent prompt index is invalid.
     */
    error InvalidAgentPromptIndex();

    /**
     * @dev Tops up the pool balance with the specified amount.
     * @param amount The amount of tokens to top up the pool balance with.
     */
    function topUpPoolBalance(uint256 amount) external;

    /**
     * @dev Executes an inference request without additional flags.
     * @param fwdCalldata The forward calldata for the inference.
     * @param externalData The external data for the inference.
     * @param promptKey The key of the prompt for the inference.
     * @param feeAmount The amount of fee to be paid for the inference.
     */
    function infer(
        bytes calldata fwdCalldata,
        string calldata externalData,
        string calldata promptKey,
        uint256 feeAmount
    ) external;

    /**
     * @dev Executes an inference request with additional flags.
     * @param fwdCalldata The forward calldata for the inference.
     * @param externalData The external data for the inference.
     * @param promptKey The key of the prompt for the inference.
     * @param flag Additional flag for the inference.
     * @param feeAmount The amount of fee to be paid for the inference.
     */
    function infer(
        bytes calldata fwdCalldata,
        string calldata externalData,
        string calldata promptKey,
        bool flag,
        uint256 feeAmount
    ) external;
}
