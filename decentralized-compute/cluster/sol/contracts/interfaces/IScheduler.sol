// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {IInferable} from "./IInferable.sol";
import {Set} from "../library/Set.sol";

interface IScheduler is IInferable {
    enum InferenceStatus {
        Nil,
        Solving,
        Commit,
        Reveal,
        Processed,
        Killed
    }

    enum BatchStatus {
        Empty,
        Commit,
        Reveal,
        Completing,
        Completed,
        Expired
    }

    struct Inference {
        uint256 value;
        uint32 modelId;
        uint40 submitTimeout;
        InferenceStatus status;
        address creator;
        address processedMiner;
        bytes input;
        bytes output;
    }

    struct ValidateInfo {
        bytes32 commit;
        bytes32 reveal;
    }

    struct BatchInfo {
        uint40 timeout;
        uint16 countCommit;
        uint16 countReveal;
        BatchStatus status;
        bytes32 mostVotedRootHash;
        uint256 validatorFee;
        uint64[] inferIds;
        Set.AddressSet validators;
        mapping(address => ValidateInfo) commits;
        mapping(bytes32 => uint) rootHashCount;
    }

    event NewInference(
        uint64 indexed inferenceId,
        address indexed creator,
        uint32 indexed modelId,
        uint256 value,
        bytes input,
        bool flag
    );

    event NewAssignment(
        uint64 indexed inferenceId,
        address indexed miner,
        uint40 expiredAt
    );

    event AppendToBatch(
        uint64 indexed batchId,
        uint32 indexed modelId,
        uint64 indexed inferId
    );

    event InferenceStatusUpdate(
        uint64 indexed inferenceId,
        InferenceStatus newStatus
    );

    event SolutionSubmission(address indexed miner, uint256 indexed inferId);
    event StreamedData(uint256 indexed assignmentId, bytes data);

    error AlreadySubmitted();
    error SubmitTimeout();
    error Unauthorized();
    error OnlyAssignedWorker();
    error InvalidInferenceStatus();
    error InvalidData();
    error InvalidAddress();
    error InvalidValue();

    function getInferenceInfo(
        uint64 inferenceId
    ) external view returns (Inference memory);

    function getMinerRequirement() external view returns (uint8);
}
