// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import {MerkleProof} from "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";

import {TransferHelper} from "./library/TransferHelper.sol";
import {PromptSchedulerStorage, Set} from "./storages/PromptSchedulerStorage.sol";
import {IGPUManager} from "./interfaces/IGPUManager.sol";

contract PromptScheduler is
    PromptSchedulerStorage,
    OwnableUpgradeable,
    PausableUpgradeable,
    ReentrancyGuardUpgradeable
{
    using Set for Set.Uint256Set;

    string private constant VERSION = "v0.0.2";
    uint256 internal constant PERCENTAGE_DENOMINATOR = 100_00;
    uint256 private constant BLOCK_PER_YEAR = 365 days;

    receive() external payable {}

    function initialize(
        address wEAIToken_,
        address gpuManager_,
        uint8 minerRequirement_,
        uint40 submitDuration_,
        uint16 minerValidatorFeeRatio_,
        uint40 batchPeriod_
    ) external initializer {
        if (gpuManager_ == address(0) || wEAIToken_ == address(0))
            revert InvalidAddress();
        if (batchPeriod_ == 0) revert InvalidValue();

        __Ownable_init();
        __Pausable_init();
        __ReentrancyGuard_init();

        _wEAIToken = wEAIToken_;
        _gpuManager = gpuManager_;
        _minerValidatorFeeRatio = minerValidatorFeeRatio_;
        _minerRequirement = minerRequirement_;
        _submitDuration = submitDuration_;
        _lastBatchTimestamp = block.timestamp;
        _batchPeriod = batchPeriod_;
    }

    function version() external pure returns (string memory) {
        return VERSION;
    }

    function pause() external onlyOwner whenNotPaused {
        _pause();
    }

    function unpause() external onlyOwner whenPaused {
        _unpause();
    }

    function setWEAIAddress(address wEAIToken) external onlyOwner {
        if (wEAIToken == address(0)) revert InvalidAddress();
        _wEAIToken = wEAIToken;
    }

    function setSubmitDuration(uint40 submitDuration) external onlyOwner {
        if (submitDuration == 0) revert InvalidData();
        _submitDuration = submitDuration;
    }

    function infer(
        uint32 modelId,
        bytes calldata input,
        address creator,
        bool flag
    ) external whenNotPaused returns (uint64) {
        return _infer(modelId, input, creator, flag);
    }

    function infer(
        uint32 modelId,
        bytes calldata input,
        address creator
    ) external whenNotPaused returns (uint64) {
        return _infer(modelId, input, creator, false);
    }

    function _infer(
        uint32 modelId,
        bytes calldata input,
        address creator,
        bool flag
    ) internal virtual returns (uint64) {
        (address miner, uint256 modelFee) = IGPUManager(_gpuManager)
            .validateModelAndChooseRandomMiner(modelId, _minerRequirement);

        uint64 inferId = ++_inferenceCounter;
        Inference storage inference = _inferences[inferId];
        uint32 lModelId = modelId;

        inference.value = modelFee;
        inference.modelId = lModelId;
        inference.creator = creator;
        inference.input = input;

        _assignMiners(inferId, lModelId, miner);

        // transfer model fee (fee to use model) to prompt scheduler
        TransferHelper.safeTransferFrom(
            _wEAIToken,
            msg.sender,
            address(this),
            modelFee
        );

        emit NewInference(inferId, creator, lModelId, modelFee, input, flag);

        return inferId;
    }

    function _assignMiners(
        uint64 inferId,
        uint32 modelId,
        address miner
    ) internal {
        uint40 expiredAt = uint40(block.number + _submitDuration);
        _inferences[inferId].submitTimeout = expiredAt;
        _inferences[inferId].status = InferenceStatus.Solving;
        _inferences[inferId].processedMiner = miner;
        _inferencesByMiner[miner].insert(inferId);

        emit NewAssignment(inferId, miner, expiredAt);

        // append to batch
        uint64 batchId = uint64(
            (block.timestamp - _lastBatchTimestamp) / _batchPeriod
        );

        _batchInfos[modelId][batchId].inferIds.push(inferId);

        emit AppendToBatch(batchId, modelId, inferId);
    }

    function _validateSolution(bytes calldata data) internal pure virtual {
        if (data.length == 0) revert InvalidData();
    }

    function _validateInference(uint64 inferId) internal view virtual {
        // Check the msg sender is the assigned miner
        if (msg.sender != _inferences[inferId].processedMiner)
            revert OnlyAssignedWorker();

        if (uint40(block.number) > _inferences[inferId].submitTimeout)
            revert SubmitTimeout();

        if (_inferences[inferId].status != InferenceStatus.Solving) {
            revert InvalidInferenceStatus();
        }

        if (_inferences[inferId].output.length != 0) revert AlreadySubmitted();
    }

    function submitSolution(
        uint64 inferId,
        bytes calldata solution
    ) external virtual whenNotPaused {
        _validateSolution(solution);
        _validateInference(inferId);

        // Check whether the miner is available (the miner has previously joined).
        // An inactive miner or one that does not belong to the correct model is not allowed to submit a solution.
        IGPUManager(_gpuManager).validateMiner(msg.sender);

        Inference storage inference = _inferences[inferId];
        inference.output = solution; //Record the solution
        inference.status = InferenceStatus.Commit;

        // transfer fee to miner
        uint256 minerFee = (inference.value * _minerValidatorFeeRatio) /
            PERCENTAGE_DENOMINATOR;
        TransferHelper.safeTransfer(_wEAIToken, msg.sender, minerFee);

        // calculate accumulated fee for validators
        uint64 currentBatchId = uint64(
            (block.timestamp - _lastBatchTimestamp) / _batchPeriod
        );
        uint32 modelId = inference.modelId;
        if (inferId < _batchInfos[modelId][currentBatchId].inferIds[0]) {
            currentBatchId--;
        }

        _batchInfos[modelId][currentBatchId].validatorFee +=
            inference.value -
            minerFee;

        emit InferenceStatusUpdate(inferId, InferenceStatus.Commit);
        emit SolutionSubmission(msg.sender, inferId);
    }

    function getInferenceInfo(
        uint64 inferId
    ) external view returns (Inference memory) {
        return _inferences[inferId];
    }

    function getInferenceByMiner(
        address miner
    ) external view returns (uint256[] memory) {
        return _inferencesByMiner[miner].values;
    }

    // Only for testing
    function getBatchInfo(
        uint32 modelId,
        uint64 batchId
    ) external view returns (uint256, uint64[] memory) {
        return (
            _batchInfos[modelId][batchId].validatorFee,
            _batchInfos[modelId][batchId].inferIds
        );
    }

    function getMinerRequirement() external view returns (uint8) {
        return _minerRequirement;
    }
}
