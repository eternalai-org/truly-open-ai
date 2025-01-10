// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {IScheduler, Set} from "../interfaces/IScheduler.sol";

abstract contract PromptSchedulerStorage is IScheduler {
    address public _wEAIToken;
    address public _gpuManager;

    uint64 public _inferenceCounter;
    mapping(uint64 => Inference) internal _inferences;
    mapping(address => Set.Uint256Set) internal _inferencesByMiner;

    uint16 public _minerValidatorFeeRatio;
    uint40 public _submitDuration;
    uint40 internal _commitDuration;
    uint40 internal _revealDuration;
    uint8 public _minerRequirement;

    mapping(uint32 modelId => mapping(uint64 batchId => BatchInfo))
        internal _batchInfos;
    uint256 public _lastBatchTimestamp;
    uint256 public _batchPeriod;

    uint256[100] private __gap;
}
