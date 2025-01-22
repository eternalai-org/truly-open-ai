// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {IAI20, IGPUManager, IInferable} from "./interfaces/IAI20.sol";
import {SafeERC20, IERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract AI20 is ERC20, IAI20 {
    uint256 private constant PORTION_DENOMINATOR = 10000;

    TokenMetaData private _data;
    address public _gpuManager;
    address public _promptScheduler;
    uint32 public _modelId;
    IERC20 private immutable _tokenFee;
    uint256 public _poolBalance;
    mapping(bytes32 signature => bool) public _signaturesUsed;
    uint256 private _totalFee;

    constructor(
        string memory name_,
        string memory symbol_,
        address promptScheduler_,
        address gpuManager_,
        uint32 modelId_,
        IERC20 tokenFee_
    ) ERC20(name_, symbol_) {
        if (
            promptScheduler_ == address(0) ||
            gpuManager_ == address(0) ||
            address(tokenFee_) == address(0)
        ) revert InvalidData();

        _promptScheduler = promptScheduler_;
        _gpuManager = gpuManager_;
        _modelId = modelId_;
        _tokenFee = tokenFee_;
    }

    function _setModelId(uint32 modelId) internal virtual {
        if (modelId == 0 || modelId == _modelId) revert InvalidData();

        _modelId = modelId;
        emit ModelIdUpdate(modelId);
    }

    function _setPromptScheduler(address promptScheduler) internal virtual {
        if (promptScheduler == address(0)) revert InvalidData();

        _promptScheduler = promptScheduler;
        emit PromptSchedulerUpdate(promptScheduler);
    }

    function _setGPUManager(address gpuManager) internal virtual {
        if (gpuManager == address(0)) revert InvalidData();

        _gpuManager = gpuManager;
        emit GPUManagerUpdate(gpuManager);
    }

    function _withdrawFee(address recipient, uint256 amount) internal virtual {
        uint256 withdrawAmount = _totalFee < amount ? _totalFee : amount;

        if (withdrawAmount > 0) {
            _totalFee -= withdrawAmount;
            SafeERC20.safeTransfer(_tokenFee, recipient, withdrawAmount);
        }
    }

    function _validateURI(string calldata uri) internal pure virtual {
        if (bytes(uri).length == 0) revert InvalidAgentData();
    }

    function _updateAgentData(
        bytes calldata sysPrompt,
        string calldata promptKey,
        uint256 promptIdx
    ) internal virtual {
        _validateAgentData(sysPrompt, promptIdx, promptKey);
        _data.sysPrompts[promptKey][promptIdx] = sysPrompt;
    }

    function _validateAgentData(
        bytes calldata sysPrompt,
        uint256 promptIdx,
        string calldata promptKey
    ) internal view virtual {
        if (sysPrompt.length == 0) revert InvalidAgentData();
        uint256 len = _data.sysPrompts[promptKey].length;
        if (promptIdx >= len) revert InvalidAgentPromptIndex();
    }

    function _addNewAgentData(
        string calldata promptKey,
        bytes calldata sysPrompt
    ) internal virtual {
        if (sysPrompt.length == 0) revert InvalidAgentData();
        _data.sysPrompts[promptKey].push(sysPrompt);

        emit AgentDataAddNew(_data.sysPrompts[promptKey]);
    }

    function _updateAgentFee(uint fee) internal virtual {
        if (_data.fee != fee) {
            _data.fee = uint128(fee);
        }

        emit AgentFeeUpdate(fee);
    }

    function topUpPoolBalance(uint256 amount) public virtual override {
        SafeERC20.safeTransferFrom(
            _tokenFee,
            msg.sender,
            address(this),
            amount
        );
        _poolBalance += amount;

        emit TopUpPoolBalance(msg.sender, amount);
    }

    function getAgentSystemPrompt(
        string calldata promptKey
    ) public view virtual returns (bytes[] memory) {
        return _data.sysPrompts[promptKey];
    }

    function infer(
        bytes calldata fwdCalldata,
        string calldata externalData,
        string calldata promptKey,
        bool flag,
        uint feeAmount
    ) public virtual override {
        (, bytes memory fwdData) = _infer(fwdCalldata, promptKey, feeAmount);

        uint256 inferId = IInferable(_promptScheduler).infer(
            _modelId,
            fwdData,
            msg.sender,
            flag
        );

        emit InferencePerformed(
            msg.sender,
            fwdData,
            _data.fee,
            externalData,
            inferId
        );
    }

    function infer(
        bytes calldata fwdCalldata,
        string calldata externalData,
        string calldata promptKey,
        uint256 feeAmount
    ) public virtual override {
        (, bytes memory fwdData) = _infer(fwdCalldata, promptKey, feeAmount);

        uint256 inferId = IInferable(_promptScheduler).infer(
            _modelId,
            fwdData,
            msg.sender
        );

        emit InferencePerformed(
            msg.sender,
            fwdData,
            _data.fee,
            externalData,
            inferId
        );
    }

    function _infer(
        bytes calldata fwdCalldata,
        string calldata promptKey,
        uint256 feeAmount
    ) internal virtual returns (uint256, bytes memory) {
        if (_data.sysPrompts[promptKey].length == 0) revert InvalidAgentData();
        if (feeAmount < _data.fee) revert InvalidAgentFee();
        SafeERC20.safeTransferFrom(
            _tokenFee,
            msg.sender,
            address(this),
            feeAmount
        );

        bytes memory fwdData = abi.encodePacked(
            _concatSystemPrompts(_data.sysPrompts[promptKey]),
            fwdCalldata
        );
        uint256 estFeeWH = IGPUManager(_gpuManager).getMinFeeToUse(_modelId);

        if (feeAmount < estFeeWH && _poolBalance >= estFeeWH) {
            unchecked {
                _poolBalance -= estFeeWH;
            }

            if (feeAmount > 0) {
                _totalFee += feeAmount;
            }
        } else if (feeAmount >= estFeeWH) {
            uint256 remain = feeAmount - estFeeWH;
            if (remain > 0) {
                _totalFee += remain;
            }
        } else {
            revert InsufficientFunds();
        }

        SafeERC20.safeApprove(_tokenFee, _promptScheduler, estFeeWH);

        return (estFeeWH, fwdData);
    }

    function inferData() public view virtual returns (uint256) {
        return _data.fee;
    }

    function _concatSystemPrompts(
        bytes[] memory sysPrompts
    ) internal pure virtual returns (bytes memory) {
        uint256 len = sysPrompts.length;
        bytes memory concatedPrompt;

        for (uint256 i = 0; i < len; i++) {
            concatedPrompt = abi.encodePacked(
                concatedPrompt,
                sysPrompts[i],
                ";"
            );
        }

        return concatedPrompt;
    }
}
