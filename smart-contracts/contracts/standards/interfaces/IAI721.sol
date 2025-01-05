// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

import "./IBase.sol";

/**
 * @title IAI721 Interface
 * @author EAI
 * @notice This interface defines the standard for the ERC721 token, including events and functions for minting, updating, and managing tokens.
 */
interface IAI721 {
    struct TokenMetaData {
        uint128 fee; // The fee associated with the token.
        bool isUsed; // Indicates if the token is currently in use.
        uint32 modelId; // The model ID associated with the token.
        address promptScheduler; // The address of the prompt scheduler.
        mapping(string => bytes[]) sysPrompts; // Mapping of system prompts by key.
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
     * @dev Emitted when a new token is minted.
     * @param tokenId The ID of the newly minted token.
     * @param uri The URI of the token.
     * @param sysPrompt The system prompt associated with the token.
     * @param fee The fee paid for the token.
     * @param minter The address of the minter.
     */
    event NewToken(
        uint256 indexed tokenId,
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
     * @dev Emitted when an inference is performed.
     * @param tokenId The ID of the token associated with the inference.
     * @param caller The address of the caller.
     * @param data The data related to the inference.
     * @param fee The fee paid for the inference.
     * @param externalData External data related to the inference.
     * @param inferenceId The ID of the inference.
     */
    event InferencePerformed(
        uint256 indexed tokenId,
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
     * @dev Returns the next token ID.
     * @return nextTokenId The next token ID.
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
     * @param _owner The address of the owner.
     * @return An array of agent IDs.
     */
    function getAgentIdByOwner(
        address _owner
    ) external view returns (uint256[] memory);

    /**
     * @dev Updates the URI of an agent.
     * @param agentId The ID of the agent.
     * @param uri The new URI of the agent.
     */
    function updateAgentURI(uint256 agentId, string calldata uri) external;

    /**
     * @dev Updates the data of an agent.
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
     * @dev Updates the data of an agent with a signature.
     * @param agentId The ID of the agent.
     * @param sysPrompt The new system prompt data.
     * @param promptKey The key of the prompt.
     * @param promptIdx The index of the prompt.
     * @param randomNonce A random nonce.
     * @param signature The signature.
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
     * @dev Updates the URI of an agent with a signature.
     * @param agentId The ID of the agent.
     * @param uri The new URI of the agent.
     * @param randomNonce A random nonce.
     * @param signature The signature.
     */
    function updateAgentUriWithSignature(
        uint256 agentId,
        string calldata uri,
        uint256 randomNonce,
        bytes calldata signature
    ) external;

    /**
     * @dev Adds new data to an agent.
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
     * @param agentId The ID of the agent.
     * @param fee The new fee of the agent.
     */
    function updateAgentFee(uint256 agentId, uint fee) external;

    /**
     * @dev Tops up the pool balance of an agent.
     * @param agentId The ID of the agent.
     * @param amount The amount of tokens to top up the pool balance with.
     */
    function topUpPoolBalance(uint256 agentId, uint256 amount) external;

    /**
     * @dev Executes an inference request.
     * @param _agentId The ID of the agent.
     * @param _calldata The calldata for the inference.
     * @param _externalData The external data for the inference.
     * @param _promptKey The key of the prompt for the inference.
     * @param _feeAmount The amount of fee to be paid for the inference.
     */
    function infer(
        uint256 _agentId,
        bytes calldata _calldata,
        string calldata _externalData,
        string calldata _promptKey,
        uint256 _feeAmount
    ) external;

    /**
     * @dev Executes an inference request with additional flags.
     * @param _agentId The ID of the agent.
     * @param _calldata The calldata for the inference.
     * @param _externalData The external data for the inference.
     * @param _promptKey The key of the prompt for the inference.
     * @param _flag Additional flag for the inference.
     * @param _feeAmount The amount of fee to be paid for the inference.
     */
    function infer(
        uint256 _agentId,
        bytes calldata _calldata,
        string calldata _externalData,
        string calldata _promptKey,
        bool _flag,
        uint256 _feeAmount
    ) external;
}
