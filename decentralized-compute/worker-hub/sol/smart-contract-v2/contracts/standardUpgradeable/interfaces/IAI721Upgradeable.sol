// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

import "./IBase.sol";

/**
 * @title IAI721Upgradeable Interface
 * @author ETERNAL AI
 * @notice This interface defines the standard for AI721Upgradeable, a protocol for decentralized inference
 *         that builds upon ERC721 to manage agent instances as a collection of NFTs.
 * @dev Designed to enable on-chain inference services where each agent is represented,
 *      identified, and managed via NFT-like tokens, ensuring ownership, transferability,
 *      and provenance. This protocol extends ERC721 conventions to support AI agents with
 *      specialized functionalities.
 */
interface IAI721Upgradeable {
    /// @dev fee: The usage fee required to invoke this agent's functionalities.
    /// @dev isUsed: Signals whether this token is actively engaged or in use.
    /// @dev modelId: Identifies the specific model from the associated model collection utilized by this agent.
    /// @dev promptScheduler: The address of promptScheduler contract.
    /// @dev sysPrompts: The system prompt data of this agent, mapped from string keys to arrays of prompt data, managed by the agent's owner.
    struct TokenMetaData {
        uint128 fee;
        bool isUsed;
        uint32 modelId;
        address promptScheduler;
        mapping(string => bytes[]) sysPrompts;
    }

    /**
     * @dev Emitted when the mint price is updated.
     * @param newValue The new mint price.
     */
    event MintPriceUpdate(uint256 newValue);

    /**
     * @dev Emitted when the royalty portion is updated.
     * @param newValue The new royalty portion.
     */
    event RoyaltyPortionUpdate(uint16 newValue);

    /**
     * @dev Emitted when the royalty receiver is updated.
     * @param newAddress The new address of the royalty receiver.
     */
    event RoyaltyReceiverUpdate(address newAddress);

    /**
     * @dev Emitted when the GPU manager is updated.
     * @param gpuManager The address of the new GPU manager.
     */
    event GPUManagerUpdate(address gpuManager);

    /**
     * @dev Emitted when a new agent is minted.
     * @param agentId The ID of the newly minted agent.
     * @param uri The URI of the agent.
     * @param sysPrompt The system prompt associated with the agent.
     * @param fee The fee paid for the agent.
     * @param minter The address of the minter.
     */
    event NewToken(
        uint256 indexed agentId,
        string uri,
        bytes sysPrompt,
        uint fee,
        address indexed minter
    );

    /**
     * @dev Emitted when the URI of an agent is updated.
     * @param agentId The ID of the agent.
     * @param uri The new URI of the agent.
     */
    event AgentURIUpdate(uint256 indexed agentId, string uri);

    /**
     * @dev Emitted when the data of an agent is updated.
     * @param agentId The ID of the agent.
     * @param promptIndex The index of the prompt being updated.
     * @param oldSysPrompt The old system prompt data.
     * @param newSysPrompt The new system prompt data.
     */
    event AgentDataUpdate(
        uint256 indexed agentId,
        uint256 promptIndex,
        bytes oldSysPrompt,
        bytes newSysPrompt
    );

    /**
     * @dev Emitted when new data is added to an agent.
     * @param agentId The ID of the agent.
     * @param sysPrompt The new system prompt data.
     */
    event AgentDataAddNew(uint256 indexed agentId, bytes[] sysPrompt);

    /**
     * @dev Emitted when the fee of an agent is updated.
     * @param agentId The ID of the agent.
     * @param fee The new fee of the agent.
     */
    event AgentFeeUpdate(uint256 indexed agentId, uint fee);

    /**
     * @dev Emitted when the model ID of an agent is updated.
     * @param agentId The ID of the agent whose model ID is being updated.
     * @param oldModelId The previous model ID of the agent.
     * @param newModelId The new model ID of the agent.
     */
    event AgentModelIdUpdate(
        uint256 indexed agentId,
        uint256 oldModelId,
        uint256 newModelId
    );

    /**
     * @dev Emitted when the prompt scheduler of an agent is updated.
     * @param agentId The ID of the agent whose prompt scheduler is being updated.
     * @param oldPromptScheduler The previous address of the prompt scheduler.
     * @param newOldPromptScheduler The new address of the prompt scheduler.
     */
    event AgentPromptSchedulerdUpdate(
        uint256 indexed agentId,
        address oldPromptScheduler,
        address newOldPromptScheduler
    );

    /**
     * @dev Emitted when an inference is performed.
     * @param agentId The ID of the agent associated with the inference.
     * @param caller The address of the caller.
     * @param data The data related to the inference.
     * @param fee The fee paid for the inference.
     * @param externalData External data related to the inference.
     * @param inferenceId The ID of the inference.
     */
    event InferencePerformed(
        uint256 indexed agentId,
        address indexed caller,
        bytes data,
        uint fee,
        string externalData,
        uint256 inferenceId
    );

    /**
     * @dev Emitted when the pool balance is topped up.
     * @param agentId The ID of the agent.
     * @param caller The address of the caller.
     * @param amount The amount of tokens used to top up the pool balance.
     */
    event TopUpPoolBalance(uint256 agentId, address caller, uint256 amount);

    /**
     * @dev Error thrown when there are insufficient funds.
     */
    error InsufficientFunds();

    /**
     * @dev Error thrown when an invalid agent ID is provided.
     */
    error InvalidAgentId();

    /**
     * @dev Error thrown when an invalid agent fee is provided.
     */
    error InvalidAgentFee();

    /**
     * @dev Error thrown when invalid agent data is provided.
     */
    error InvalidAgentData();

    /**
     * @dev Error thrown when an invalid agent URI is provided.
     */
    error InvalidAgentURI();

    /**
     * @dev Error thrown when an invalid agent prompt index is provided.
     */
    error InvalidAgentPromptIndex();

    /**
     * @dev Error thrown when a signature has already been used.
     */
    error SignatureUsed();

    /**
     * @dev Error thrown when the caller is not authorized.
     */
    error Unauthorized();

    /**
     * @dev Error thrown when invalid data is provided.
     */
    error InvalidData();

    /**
     * @dev Returns the next agent ID.
     * @return nextTokenId The next agent ID.
     */
    function nextTokenId() external view returns (uint256 nextTokenId);

    /**
     * @dev Returns the address of the royalty receiver.
     * @return royaltyReceiver The address of the royalty receiver.
     */
    function royaltyReceiver() external view returns (address royaltyReceiver);

    /**
     * @dev Returns the royalty portion.
     * @return royaltyPortion The royalty portion.
     */
    function royaltyPortion() external view returns (uint16 royaltyPortion);

    /**
     * @dev Returns an array of agent IDs owned by a given owner.
     * @param owner The address of the owner.
     * @return An array of agent IDs.
     */
    function getAgentIdByOwner(
        address owner
    ) external view returns (uint256[] memory);

    /**
     * @dev Updates the URI of an agent.
     * @notice Only the owner of the agent can call this function.
     * @param agentId The ID of the agent.
     * @param uri The new URI of the agent.
     */
    function updateAgentURI(uint256 agentId, string calldata uri) external;

    /**
     * @dev Updates the data of an agent.
     * @notice Only the owner of the agent can call this function.
     * @param agentId The ID of the agent.
     * @param sysPrompt The new system prompt data.
     * @param promptKey The key of the prompt.
     * @param promptIdx The index of the prompt.
     */
    function updateAgentData(
        uint256 agentId,
        bytes calldata sysPrompt,
        string calldata promptKey,
        uint256 promptIdx
    ) external;

    /**
     * @dev This function modifies the model ID associated with an existing agent.
     * @notice Only the owner of the agent can call this function.
     * @param agentId The unique identifier of the agent to update.
     * @param newModelId The new model ID to assign to the agent.
     */
    function updateAgentModelId(uint256 agentId, uint32 newModelId) external;

    /**
     * @dev This function allows an agent owner to update agent data without submitting the transaction themselves.
     * @param agentId The ID of the agent.
     * @param sysPrompt The new system prompt data.
     * @param promptKey The key of the prompt.
     * @param promptIdx The index of the prompt.
     * @param randomNonce A random nonce to ensure uniqueness.
     * @param signature The digital signature authorizing the update.
     */
    function updateAgentDataWithSignature(
        uint256 agentId,
        bytes calldata sysPrompt,
        string calldata promptKey,
        uint256 promptIdx,
        uint256 randomNonce,
        bytes calldata signature
    ) external;

    /**
     * @dev This function allows an agent owner to update the URI of an agent
     * without submitting the transaction themselves.
     * @param agentId The ID of the agent.
     * @param uri The new URI of the agent.
     * @param randomNonce A random nonce used to ensure unique transaction execution.
     * @param signature A valid signature authorizing the update.
     */
    function updateAgentUriWithSignature(
        uint256 agentId,
        string calldata uri,
        uint256 randomNonce,
        bytes calldata signature
    ) external;

    /**
     * @dev Adds new data to an agent.
     * @notice Only the agent owner can add new data.
     * @param agentId The ID of the agent.
     * @param promptKey The key of the prompt.
     * @param sysPrompt The new system prompt data.
     */
    function addNewAgentData(
        uint256 agentId,
        string calldata promptKey,
        bytes calldata sysPrompt
    ) external;

    /**
     * @dev Updates the fee of an agent.
     * @notice Only the agent owner can update the agent fee. The agent fee is typically greater than or equal to the model usage fee.
     * @param agentId The ID of the agent.
     * @param fee The fee to use this agent.
     */
    function updateAgentFee(uint256 agentId, uint fee) external;

    /**
     * @dev Tops up the pool balance of an agent.
     * @notice Anyone can top up the pool balance of an agent.
     * @param agentId The ID of the agent.
     * @param amount The amount of tokens to top up the pool balance with.
     */
    function topUpPoolBalance(uint256 agentId, uint256 amount) external;

    /**
     * @dev Executes an inference request for a specified agent. This function facilitates the interaction with
     *      an AI agent by providing the necessary data and parameters to perform an inference operation.
     * @notice The `feeAmount` must be greater than or equal to the fee required to use the agent.
     * @param agentId The ID of the agent.
     * @param inferenceData The calldata for the inference.
     * @param externalData The external data for the inference.
     * @param promptKey The key of the prompt for the inference.
     * @param feeAmount The amount of fee to be paid for the inference.
     */
    function infer(
        uint256 agentId,
        bytes calldata inferenceData,
        string calldata externalData,
        string calldata promptKey,
        uint256 feeAmount
    ) external;

    /**
     * @dev Executes an inference request for a specified agent. This function facilitates the interaction with
     *      an AI agent by providing the necessary data and parameters to perform an inference operation.
     * @notice The `feeAmount` must be greater than or equal to the fee required to use the agent.
     * @param agentId The ID of the agent.
     * @param inferenceData The calldata for the inference.
     * @param externalData The external data for the inference.
     * @param promptKey The key of the prompt for the inference.
     * @param feeAmount The amount of fee to be paid for the inference.
     * @param rawFlag The flag to indicate the format of the calldata and the result of the inference.
     *                  If rawFlag is true, the calldata and inference result are in raw format.
     *                  If rawFlag is false, the calldata and inference result are in IPFS link format.
     */
    function infer(
        uint256 agentId,
        bytes calldata inferenceData,
        string calldata externalData,
        string calldata promptKey,
        bool rawFlag,
        uint256 feeAmount
    ) external;
}
