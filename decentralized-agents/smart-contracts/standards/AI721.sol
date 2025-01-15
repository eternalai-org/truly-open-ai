// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

import {IAI721, IGPUManager, IInferable} from "./interfaces/IAI721.sol";
import {ERC721URIStorage, ERC721} from "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import {ERC721Enumerable} from "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/EIP712.sol";
import {IERC2981} from "@openzeppelin/contracts/interfaces/IERC2981.sol";
import {SafeERC20, IERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/**
 * @title AI721
 * @dev Implementation of decentralized inference standard AI721.
 */
contract AI721 is ERC721Enumerable, ERC721URIStorage, IAI721 {
    /// @dev Storage
    mapping(uint256 agentId => TokenMetaData) private _datas;
    mapping(uint256 agentId => uint256) public _poolBalance;
    uint256 private _nextTokenId;
    address public _gpuManager;
    IERC20 private immutable _tokenFee;

    /// @dev constructor
    constructor(
        string memory name_,
        string memory symbol_,
        uint256 nextTokenId_,
        address gpuManager_,
        IERC20 tokenFee_
    ) ERC721(name_, symbol_) {
        require(
            gpuManager_ != address(0) && address(tokenFee_) != address(0),
            "Zero address"
        );

        _nextTokenId = nextTokenId_;
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
     * @dev Mints a new agent with the specified parameters.
     * Reverts with `InvalidAgentData` if the provided data is empty.
     *
     * @param to The address to which the agent will be minted.
     * @param uri The URI associated with the agent.
     * @param data The system prompt of the agent.
     * @param fee The using fee associated with the agent.
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
        if (data.length == 0) revert InvalidAgentData();

        _safeMint(to, agentId);
        _setTokenURI(agentId, uri);

        _datas[agentId].fee = uint128(fee);
        _datas[agentId].sysPrompts[promptKey].push(data);
        _datas[agentId].isUsed = true;
        _datas[agentId].promptScheduler = promptScheduler;
        _datas[agentId].modelId = modelId;

        emit NewToken(agentId, uri, data, fee, to);

        return agentId;
    }

    /**
     * @dev Wraps the minting process by finding the next available agent ID and minting a new agent.
     * Reverts with `InvalidAgentData` if the provided data is empty.
     *
     * @param to The address to which the agent will be minted.
     * @param uri The URI associated with the agent.
     * @param data The system prompt of the agent.
     * @param fee The using fee associated with the agent.
     * @param promptKey The key for the system prompt.
     * @param promptScheduler The address of the prompt scheduler.
     * @param modelId The ID of the model that is used by thee agent.
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
        while (_datas[_nextTokenId].isUsed) {
            _nextTokenId++;
        }
        uint256 agentId = _nextTokenId++;

        _mint(to, uri, data, fee, agentId, promptKey, promptScheduler, modelId);

        return agentId;
    }

    function _validateURI(string calldata uri) internal pure virtual {
        if (bytes(uri).length == 0) revert InvalidAgentData();
    }

    /**
     * @dev Updates the URI of an agent.
     * @param agentId The ID of the agent.
     * @param uri The new URI of the agent.
     */
    function _updateAgentURI(
        uint256 agentId,
        string calldata uri
    ) internal virtual {
        _validateURI(uri);

        _setTokenURI(agentId, uri);
        emit AgentURIUpdate(agentId, uri);
    }

    /**
     * @dev Updates the data of an agent.
     * @param agentId The ID of the agent.
     * @param sysPrompt The new system prompt data.
     * @param promptKey The key of the prompt.
     * @param promptIdx The index of the prompt.
     */
    function _updateAgentData(
        uint256 agentId,
        bytes calldata sysPrompt,
        string calldata promptKey,
        uint256 promptIdx
    ) internal virtual {
        _validateAgentData(agentId, sysPrompt, promptIdx, promptKey);

        emit AgentDataUpdate(
            agentId,
            promptIdx,
            _datas[agentId].sysPrompts[promptKey][promptIdx],
            sysPrompt
        );

        _datas[agentId].sysPrompts[promptKey][promptIdx] = sysPrompt;
    }

    /**
     * @dev This function modifies the model ID associated with an existing agent.
     * @param agentId The unique identifier of the agent to update.
     * @param newModelId The new model ID to assign to the agent.
     */
    function _updateAgentModelId(
        uint256 agentId,
        uint32 newModelId
    ) internal virtual {
        emit AgentModelIdUpdate(agentId, _datas[agentId].modelId, newModelId);

        _datas[agentId].modelId = newModelId;
    }

    /**
     * @dev Updates the prompt scheduler for a given agent.
     * Emits an {AgentPromptSchedulerUpdate} event.
     *
     * @param agentId The ID of the agent whose prompt scheduler is being updated.
     * @param newPromptScheduler The address of the new prompt scheduler.
     */
    function _updatePromptScheduler(
        uint256 agentId,
        address newPromptScheduler
    ) internal virtual {
        emit AgentPromptSchedulerUpdate(
            agentId,
            _datas[agentId].promptScheduler,
            newPromptScheduler
        );

        _datas[agentId].promptScheduler = newPromptScheduler;
    }

    /**
     * @dev Updates the fee of an agent.
     * @notice The agent fee is typically greater than or equal to the model usage fee.
     * @param agentId The ID of the agent.
     * @param fee The fee to use this agent.
     */
    function _updateAgentFee(uint256 agentId, uint fee) internal virtual {
        if (_datas[agentId].fee != fee) {
            _datas[agentId].fee = uint128(fee);
        }

        emit AgentFeeUpdate(agentId, fee);
    }

    /**
     * @dev Validates the agent data by checking the system prompt and prompt index.
     * Reverts with `InvalidAgentData` if the system prompt is empty.
     *
     * @param agentId The ID of the agent.
     * @param sysPrompt The system prompt data in bytes.
     * @param promptKey The key associated with the prompt.
     */
    function _validateAgentData(
        uint256 agentId,
        bytes calldata sysPrompt,
        uint256 promptIdx,
        string calldata promptKey
    ) internal view virtual {
        if (sysPrompt.length == 0) revert InvalidAgentData();
        uint256 len = _datas[agentId].sysPrompts[promptKey].length;
        if (promptIdx >= len) revert InvalidAgentPromptIndex();
    }

    /**
     * @dev Adds new data to an agent.
     * @param agentId The ID of the agent.
     * @param promptKey The key of the prompt.
     * @param sysPrompt The new system prompt data.
     */
    function _addNewAgentData(
        uint256 agentId,
        string calldata promptKey,
        bytes calldata sysPrompt
    ) internal virtual {
        if (sysPrompt.length == 0) revert InvalidAgentData();

        _datas[agentId].sysPrompts[promptKey].push(sysPrompt);

        emit AgentDataAddNew(agentId, _datas[agentId].sysPrompts[promptKey]);
    }

    /**
     * @dev See {IAI721Upgradeable-topUpPoolBalance}.
     */
    function topUpPoolBalance(
        uint256 agentId,
        uint256 amount
    ) public virtual override {
        SafeERC20.safeTransferFrom(
            _tokenFee,
            msg.sender,
            address(this),
            amount
        );
        _poolBalance[agentId] += amount;

        emit TopUpPoolBalance(agentId, msg.sender, amount);
    }

    /**
     * @dev See {IAI721Upgradeable-getAgentFee}.
     */
    function getAgentFee(
        uint256 agentId
    ) external view virtual returns (uint256) {
        return _datas[agentId].fee;
    }

    /**
     * @dev See {IAI721Upgradeable-getAgentSystemPrompt}.
     */
    function getAgentSystemPrompt(
        uint256 agentId,
        string calldata promptKey
    ) external view virtual returns (bytes[] memory) {
        return _datas[agentId].sysPrompts[promptKey];
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
    ) public virtual override {
        (, bytes memory fwdData) = _infer(
            agentId,
            fwdCalldata,
            promptKey,
            feeAmount
        );

        uint256 inferId = IInferable(_datas[agentId].promptScheduler).infer(
            _datas[agentId].modelId,
            fwdData,
            msg.sender,
            flag
        );

        emit InferencePerformed(
            agentId,
            msg.sender,
            fwdData,
            _datas[agentId].fee,
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
    ) public virtual override {
        (, bytes memory fwdData) = _infer(
            agentId,
            fwdCalldata,
            promptKey,
            feeAmount
        );

        uint256 inferId = IInferable(_datas[agentId].promptScheduler).infer(
            _datas[agentId].modelId,
            fwdData,
            msg.sender
        );

        emit InferencePerformed(
            agentId,
            msg.sender,
            fwdData,
            _datas[agentId].fee,
            externalData,
            inferId
        );
    }

    /**
     * @dev Internal function to handle inference requests for a given agent.
     *
     * This function performs several checks and operations:
     * 1. Validates the existence of system prompts for the given agent and prompt key.
     * 2. Ensures the provided fee amount meets the required fee for the agent.
     * 3. Transfers the fee amount from the sender to the contract.
     * 4. Encodes the forward calldata with the system prompts.
     * 5. Retrieves the minimum fee required to use the GPU for the agent's model.
     * 6. Adjusts the agent's pool balance and transfers any remaining fee back to the agent owner if applicable.
     * 7. Approves the prompt scheduler to use the estimated fee.
     *
     * @param agentId The ID of the agent for which inference is requested.
     * @param fwdCalldata The calldata to be forwarded for inference.
     * @param promptKey The key to identify the system prompt.
     * @param feeAmount The fee amount provided for the inference request.
     *
     * @return estFeeWH The estimated fee required for the inference must be paid to the miner.
     * @return fwdData The encoded forward data including system prompts and calldata.
     *
     * @notice Reverts if the agent data or fee is invalid, or if there are insufficient funds.
     */
    function _infer(
        uint256 agentId,
        bytes calldata fwdCalldata,
        string calldata promptKey,
        uint256 feeAmount
    ) internal virtual returns (uint256, bytes memory) {
        if (_datas[agentId].sysPrompts[promptKey].length == 0)
            revert InvalidAgentData();
        if (feeAmount < _datas[agentId].fee) revert InvalidAgentFee();
        SafeERC20.safeTransferFrom(
            _tokenFee,
            msg.sender,
            address(this),
            feeAmount
        );

        bytes memory fwdData = abi.encodePacked(
            _concatSystemPrompts(_datas[agentId].sysPrompts[promptKey]),
            fwdCalldata
        );
        uint256 estFeeWH = IGPUManager(_gpuManager).getMinFeeToUse(
            _datas[agentId].modelId
        );

        if (feeAmount < estFeeWH && _poolBalance[agentId] >= estFeeWH) {
            unchecked {
                _poolBalance[agentId] -= estFeeWH;
            }

            if (feeAmount > 0) {
                SafeERC20.safeTransfer(_tokenFee, _ownerOf(agentId), feeAmount);
            }
        } else if (feeAmount >= estFeeWH) {
            uint256 remain = feeAmount - estFeeWH;
            if (remain > 0) {
                SafeERC20.safeTransfer(_tokenFee, _ownerOf(agentId), remain);
            }
        } else {
            revert InsufficientFunds();
        }

        SafeERC20.safeApprove(
            _tokenFee,
            _datas[agentId].promptScheduler,
            estFeeWH
        );

        return (estFeeWH, fwdData);
    }

    /**
     * @dev See {IAI721Upgradeable-dataOf}.
     */
    function dataOf(
        uint256 agentId
    )
        external
        view
        virtual
        returns (
            uint128 fee,
            bool isUsed,
            uint32 modelId,
            address promptScheduler,
            uint256 poolBalance
        )
    {
        TokenMetaData storage data = _datas[agentId];
        return (
            data.fee,
            data.isUsed,
            data.modelId,
            data.promptScheduler,
            _poolBalance[agentId]
        );
    }

    /**
     * @dev See {IERC721Metadata-tokenURI}.
     */
    function tokenURI(
        uint256 agentId
    ) public view override(ERC721, ERC721URIStorage) returns (string memory) {
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
     * @dev See {IAI721Upgradeable-nextTokenId}.
     */
    function nextTokenId() external view returns (uint256) {
        return _nextTokenId;
    }

    /**
     * @dev Concatenates an array of system prompts into a single bytes array,
     *      with each prompt separated by a semicolon.
     * @param sysPrompts An array of bytes representing the system prompts to be concatenated.
     * @return A bytes array containing all the concatenated system prompts separated by semicolons.
     */
    function _concatSystemPrompts(
        bytes[] memory sysPrompts
    ) internal pure virtual returns (bytes memory) {
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
