// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

/**
 * @title IAI721Errors
 * @dev Interface defining error types for the AI721 standard
 */
interface IAI721Errors {
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
}
