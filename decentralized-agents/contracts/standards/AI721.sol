// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

import {IERC2981} from "@openzeppelin/contracts/interfaces/IERC2981.sol";
import {ERC721Enumerable} from "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import {ERC721URIStorage, ERC721} from "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import {SafeERC20, IERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {IAI721, IGPUManager, IInferable} from "./interfaces/IAI721.sol";

/**
 * @title AI721
 * @dev Implementation of decentralized inference standard AI721.
 */
contract AI721 is ERC721Enumerable, ERC721URIStorage, IAI721 {
    /// @dev Storage
    mapping(uint256 agentId => AgentConfig) internal _agentConfigs;
    uint256 private _nextAgentId;
    address internal _gpuManager;
    IERC20 internal _tokenFee;

    /**
     * @dev Constructor for the AI721 contract.
     * @param name_ The name of the ERC721 token.
     * @param symbol_ The symbol of the ERC721 token.
     * @param nextAgentId_ The initial value for the next agent ID.
     * @param gpuManager_ The address of the GPU manager.
     * @param tokenFee_ The address of the ERC20 token used for fees.
     *
     * Requirements:
     * - `gpuManager_` must not be the zero address.
     * - `tokenFee_` must not be the zero address.
     * - `nextAgentId_` must not be zero.
     *
     * Reverts with `InvalidNextAgentId` if `nextAgentId_` is zero.
     */
    constructor(
        string memory name_,
        string memory symbol_,
        uint256 nextAgentId_,
        address gpuManager_,
        IERC20 tokenFee_
    ) ERC721(name_, symbol_) {
        require(
            gpuManager_ != address(0) && address(tokenFee_) != address(0),
            "Zero address"
        );

        if (nextAgentId_ == 0) revert InvalidNextAgentId();

        _nextAgentId = nextAgentId_;
        _gpuManager = gpuManager_;
        _tokenFee = tokenFee_;
    }

    /**
     * @dev Internal function to update the GPU manager address.
     * Reverts if the provided address is the zero address.
     *
     * @param gpuManager The new address of the GPU manager.
     *
     * Emits a {GPUManagerUpdate} event.
     */
    function _updateGPUManager(address gpuManager) internal virtual {
        if (gpuManager == address(0)) revert InvalidData();

        _gpuManager = gpuManager;
        emit GPUManagerUpdate(gpuManager);
    }

    /**
     * @dev See {IAI721Upgradeable-getGPUManager}.
     */
    function getGPUManager() external view returns (address) {
        return _gpuManager;
    }

    /**
     * @dev See {IAI721Upgradeable-getTokenFee}.
     */
    function getTokenFee() external view returns (address) {
        return address(_tokenFee);
    }

    /**
     * @dev Mints a new agent with the specified parameters.
     * Reverts with `InvalidAgentId` if the agent ID is already used.
     * Reverts with `InvalidData` if the provided data is empty.
     *
     * @param to The address to which the agent will be minted.
     * @param uri The URI associated with the agent.
     * @param data The system prompt of the agent.
     * @param fee The usage fee associated with the agent.
     * @param agentId The ID of the agent.
     * @param promptKey The key for the system prompt.
     * @param promptScheduler The address of the prompt scheduler.
     * @param modelId The ID of the model.
     * @return The ID of the minted agent.
     */
    function _mint(
        address to,
        string calldata uri,
        bytes calldata data,
        uint fee,
        uint256 agentId,
        string calldata promptKey,
        address promptScheduler,
        uint32 modelId
    ) internal virtual returns (uint256) {
        if (_agentConfigs[_nextAgentId].isUsed) revert InvalidAgentId();
        _validateData(data);

        _safeMint(to, agentId);
        _setTokenURI(agentId, uri);

        _agentConfigs[agentId].usageFee = uint128(fee);
        _agentConfigs[agentId].sysPrompts[promptKey].push(data);
        _agentConfigs[agentId].isUsed = true;
        _agentConfigs[agentId].promptScheduler = promptScheduler;
        _agentConfigs[agentId].modelId = modelId;

        emit NewAgent(agentId, uri, data, fee, to);

        return agentId;
    }

    /**
     * @dev Wraps the minting process by finding the next available agent ID and minting a new agent.
     * Reverts with `InvalidData` if the provided data is empty.
     *
     * @param to The address to which the agent will be minted.
     * @param uri The URI associated with the agent.
     * @param data The system prompt of the agent.
     * @param fee The usage fee associated with the agent.
     * @param promptKey The key for the system prompt.
     * @param promptScheduler The address of the prompt scheduler.
     * @param modelId The ID of the model that is used by the agent.
     * @return The ID of the minted agent.
     */
    function _wrapMint(
        address to,
        string calldata uri,
        bytes calldata data,
        uint fee,
        string calldata promptKey,
        address promptScheduler,
        uint32 modelId
    ) internal virtual returns (uint256) {
        while (_agentConfigs[_nextAgentId].isUsed) {
            _nextAgentId++;
        }
        uint256 agentId = _nextAgentId++;

        _mint(to, uri, data, fee, agentId, promptKey, promptScheduler, modelId);

        return agentId;
    }

    /**
     * @dev See {IAI721Upgradeable-nextAgentId}.
     */
    function nextAgentId() external view returns (uint256) {
        return _nextAgentId;
    }

    /**
     * @dev Updates the URI of a specified agent.
     * Reverts if the provided URI is invalid.
     *
     * @param agentId The ID of the agent to update.
     * @param uri The new URI to be assigned to the agent.
     *
     * Emits an {AgentURIUpdate} event.
     */
    function _updateAgentURI(
        uint256 agentId,
        string calldata uri
    ) internal virtual {
        _setTokenURI(agentId, uri);
        emit AgentURIUpdate(agentId, uri);
    }

    /**
     * @dev Updates the system prompt data for a specified agent.
     * Reverts if the provided data is invalid or if the prompt index is out of bounds.
     *
     * Emits a {PromptDataUpdated} event indicating the agent and the updated prompt index.
     *
     * @param agentId The ID of the agent whose prompt data is being updated.
     * @param sysPrompt The new system prompt data in bytes.
     * @param promptKey The key associated with the prompt.
     * @param promptIdx The index of the prompt to update.
     */
    function _updateAgentData(
        uint256 agentId,
        bytes calldata sysPrompt,
        string calldata promptKey,
        uint256 promptIdx
    ) internal virtual {
        _beforeUpdateAgentData(agentId, sysPrompt, promptIdx, promptKey);

        emit AgentDataUpdate(
            agentId,
            promptIdx,
            _agentConfigs[agentId].sysPrompts[promptKey][promptIdx],
            sysPrompt
        );

        _agentConfigs[agentId].sysPrompts[promptKey][promptIdx] = sysPrompt;
    }

    /**
     * @dev Updates the model ID of a specified agent.
     * Emits an {AgentModelIdUpdate} event.
     *
     * @param agentId The ID of the agent to update.
     * @param newModelId The new model ID to be assigned to the agent.
     */
    function _updateAgentModelId(
        uint256 agentId,
        uint32 newModelId
    ) internal virtual {
        emit AgentModelIdUpdate(
            agentId,
            _agentConfigs[agentId].modelId,
            newModelId
        );

        _agentConfigs[agentId].modelId = newModelId;
    }

    /**
     * @dev Updates the prompt scheduler for a given agent.
     * Emits an {AgentPromptSchedulerUpdate} event.
     *
     * @param agentId The ID of the agent whose prompt scheduler is being updated.
     * @param newPromptScheduler The address of the new prompt scheduler.
     */
    function _updateAgentPromptScheduler(
        uint256 agentId,
        address newPromptScheduler
    ) internal virtual {
        emit AgentPromptSchedulerUpdate(
            agentId,
            _agentConfigs[agentId].promptScheduler,
            newPromptScheduler
        );

        _agentConfigs[agentId].promptScheduler = newPromptScheduler;
    }

    /**
     * @dev Internal function to update the usage fee of a specified agent.
     * Emits an {AgentFeeUpdated} event.
     * @notice The agent fee should generally be greater than or equal to the model usage fee.
     * @param agentId The unique identifier of the agent.
     * @param fee The new fee to be set for using this agent.
     */
    function _updateAgentUsageFee(uint256 agentId, uint fee) internal virtual {
        if (_agentConfigs[agentId].usageFee != fee) {
            _agentConfigs[agentId].usageFee = uint128(fee);
        }

        emit AgentUsageFeeUpdate(agentId, fee);
    }
    /**
     * @dev Adds new system prompt data to a specified agent.
     * Reverts if the provided data is invalid.
     *
     * @param agentId The ID of the agent to update.
     * @param promptKey The key associated with the prompt.
     * @param sysPrompt The new system prompt data in bytes.
     *
     * Emits an {AgentDataAddNew} event.
     */
    function _addNewAgentData(
        uint256 agentId,
        string calldata promptKey,
        bytes calldata sysPrompt
    ) internal virtual {
        _validateData(sysPrompt);

        _agentConfigs[agentId].sysPrompts[promptKey].push(sysPrompt);

        emit AgentDataAddNew(
            agentId,
            _agentConfigs[agentId].sysPrompts[promptKey]
        );
    }

    /**
     * @dev Validates the agent data by checking the system prompt and prompt index.
     * Reverts with `InvalidAgentData` if the system prompt is empty.
     *
     * @param agentId The ID of the agent.
     * @param sysPrompt The system prompt data in bytes.
     * @param promptKey The key associated with the prompt.
     */
    function _beforeUpdateAgentData(
        uint256 agentId,
        bytes calldata sysPrompt,
        uint256 promptIdx,
        string calldata promptKey
    ) internal view virtual {
        _validateData(sysPrompt);

        uint256 len = _agentConfigs[agentId].sysPrompts[promptKey].length;
        if (promptIdx >= len) revert InvalidAgentPromptIndex();
    }

    /**
     * @dev Validates the provided URI string.
     * Reverts with `InvalidData` if the URI is an empty string.
     *
     * @param data The data to validate.
     */
    function _validateData(bytes calldata data) internal pure virtual {
        if (data.length == 0) revert InvalidData();
    }

    /**
     * @dev See {IAI721Upgradeable-getAgentUsageFee}.
     */
    function getAgentUsageFee(
        uint256 id
    ) external view virtual returns (uint256) {
        return _agentConfigs[id].usageFee;
    }

    /**
     * @dev See {IAI721Upgradeable-getAgentSystemPrompt}.
     */
    function getAgentSystemPrompt(
        uint256 id,
        string calldata promptKey
    ) external view virtual returns (bytes[] memory) {
        return _agentConfigs[id].sysPrompts[promptKey];
    }

    /**
     * @dev See {IAI721Upgradeable-getAgentConfig}.
     */
    function getAgentConfig(
        uint256 agentId
    )
        external
        view
        virtual
        returns (
            uint128 usageFee,
            bool isUsed,
            uint32 modelId,
            address promptScheduler
        )
    {
        AgentConfig storage config = _agentConfigs[agentId];
        return (
            config.usageFee,
            config.isUsed,
            config.modelId,
            config.promptScheduler
        );
    }

    /**
     * @dev See {IAI721Upgradeable-infer}.
     */
    function infer(
        uint256 agentId,
        bytes calldata fwdCalldata,
        string calldata externalData,
        string calldata promptKey,
        bool flag,
        uint feeAmount
    ) public virtual override returns (uint256 inferId) {
        _validateBeforeInference(agentId, promptKey, feeAmount);

        _processBeforeInference(agentId, fwdCalldata, promptKey, feeAmount);

        bytes memory fwdData = _buildForwardedData(
            agentId,
            promptKey,
            fwdCalldata
        );

        inferId = IInferable(_agentConfigs[agentId].promptScheduler).infer(
            _agentConfigs[agentId].modelId,
            fwdData,
            msg.sender,
            flag
        );

        emit InferencePerformed(
            agentId,
            msg.sender,
            fwdData,
            _agentConfigs[agentId].usageFee,
            externalData,
            inferId
        );
    }

    /**
     * @dev See {IAI721Upgradeable-infer}.
     */
    function infer(
        uint256 agentId,
        bytes calldata fwdCalldata,
        string calldata externalData,
        string calldata promptKey,
        uint256 feeAmount
    ) public virtual override returns (uint256 inferId) {
        _validateBeforeInference(agentId, promptKey, feeAmount);

        _processBeforeInference(agentId, fwdCalldata, promptKey, feeAmount);

        bytes memory fwdData = _buildForwardedData(
            agentId,
            promptKey,
            fwdCalldata
        );

        inferId = IInferable(_agentConfigs[agentId].promptScheduler).infer(
            _agentConfigs[agentId].modelId,
            fwdData,
            msg.sender
        );

        emit InferencePerformed(
            agentId,
            msg.sender,
            fwdData,
            _agentConfigs[agentId].usageFee,
            externalData,
            inferId
        );
    }

    /**
     * @dev Validates the agent data and fee amount before performing inference.
     * Reverts if the agent data or fee amount is invalid.
     *
     * @param agentId The ID of the agent to validate.
     * @param promptKey The key of the prompt to validate.
     * @param feeAmount The fee amount to validate.
     *
     * Requirements:
     *
     * - The prompt associated with `promptKey` must exist for the given `agentId`.
     * - The `feeAmount` must be greater than or equal to the usage fee of the agent.
     *
     * Reverts with:
     * - `InvalidAgentData` if the prompt does not exist.
     * - `InvalidAgentFee` if the fee amount is less than the required usage fee.
     */
    function _validateBeforeInference(
        uint256 agentId,
        string calldata promptKey,
        uint256 feeAmount
    ) internal view virtual {
        if (_agentConfigs[agentId].sysPrompts[promptKey].length == 0)
            revert InvalidAgentData();
        if (feeAmount < _agentConfigs[agentId].usageFee)
            revert InvalidAgentFee();
    }

    /**
     * @dev Hook that is called before the inference process. This function can be overridden to implement
     * any logic that needs to be executed before the inference.
     *
     * @param agentId The identifier of the agent.
     * @param fwdCalldata The calldata to be forwarded.
     * @param promptKey The key for the prompt.
     * @param feeAmount The amount of fee to be processed.
     */
    function _processBeforeInference(
        uint256 agentId,
        bytes calldata fwdCalldata,
        string calldata promptKey,
        uint256 feeAmount
    ) internal virtual {}

    /**
     * @dev Constructs the forwarded data by concatenating system prompts with the input data.
     * This function can be overridden to customize the data construction process.
     *
     * @param agentId The identifier of the agent.
     * @param promptKey The key for the prompt.
     * @param input The input data to be forwarded.
     * @return result The constructed forwarded data.
     */
    function _buildForwardedData(
        uint256 agentId,
        string calldata promptKey,
        bytes calldata input
    ) internal view virtual returns (bytes memory result) {
        result = abi.encodePacked(
            _concatSystemPrompts(agentId, promptKey),
            input
        );
    }

    /**
     * @dev See {IERC721Metadata-tokenURI}.
     */
    function tokenURI(
        uint256 agentId
    )
        public
        view
        virtual
        override(ERC721, ERC721URIStorage)
        returns (string memory)
    {
        return super.tokenURI(agentId);
    }

    /**
     * @dev Checks if the given user is the owner of the specified agent.
     * Reverts with an `Unauthorized` error if the user is not the owner.
     *
     * @param user The address of the user to check.
     * @param agentId The ID of the agent to check ownership for.
     */
    function _checkAgentOwner(
        address user,
        uint256 agentId
    ) internal view virtual {
        if (user != _ownerOf(agentId)) revert Unauthorized();
    }

    /**
     * @dev See {IAI721Upgradeable-getAgentIdByOwner}.
     */
    function getAgentIdByOwner(
        address owner
    ) external view returns (uint256[] memory) {
        uint256 len = balanceOf(owner);
        uint256[] memory agentIds = new uint256[](len);

        for (uint256 i = 0; i < len; i++) {
            agentIds[i] = tokenOfOwnerByIndex(owner, i);
        }

        return agentIds;
    }

    /**
     * @dev Concatenates system prompts for a given agent and prompt key.
     *
     * This function retrieves the system prompts associated with the specified
     * agent ID and prompt key, concatenates them with a semicolon separator,
     * and returns the concatenated result as a byte array.
     *
     * @param agentId The ID of the agent whose system prompts are to be concatenated.
     * @param promptKey The key used to identify the specific set of system prompts.
     * @return A byte array containing the concatenated system prompts.
     */
    function _concatSystemPrompts(
        uint256 agentId,
        string memory promptKey
    ) internal view virtual returns (bytes memory) {
        bytes[] memory sysPrompts = _agentConfigs[agentId].sysPrompts[
            promptKey
        ];

        uint256 len = sysPrompts.length;
        bytes memory prompt;

        for (uint256 i = 0; i < len; i++) {
            prompt = abi.encodePacked(prompt, sysPrompts[i], ";");
        }

        return prompt;
    }

    /**
     * @dev See {ERC721-_beforeTokenTransfer}.
     */
    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 firstTokenId,
        uint256 batchSize
    ) internal virtual override(ERC721, ERC721Enumerable) {
        super._beforeTokenTransfer(from, to, firstTokenId, batchSize);
    }

    /**
     * @dev See {ERC721-_burn}.
     */
    function _burn(
        uint256 agentId
    ) internal override(ERC721, ERC721URIStorage) {
        super._burn(agentId);
    }

    /**
     * @dev See {IERC165-supportsInterface}.
     */
    function supportsInterface(
        bytes4 interfaceId
    ) public view override(ERC721Enumerable, ERC721URIStorage) returns (bool) {
        return
            interfaceId == type(IERC2981).interfaceId ||
            super.supportsInterface(interfaceId);
    }
}
