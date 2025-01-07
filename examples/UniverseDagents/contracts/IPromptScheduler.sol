// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IInferable {
    function infer(
        bytes calldata data,
        address creator
    ) external payable returns (uint256 inferenceId);

    function infer(
        bytes calldata data,
        address creator,
        bool flag
    ) external payable returns (uint256 inferenceId);
}

interface IPromptScheduler2TX is IInferable {
    enum InferenceStatus {
        Nil,
        Solving,
        Commit,
        Reveal,
        Processed,
        Killed,
        Transferred
    }

    enum AssignmentRole {
        Nil,
        Validating,
        Mining
    }

    enum Vote {
        Nil,
        Disapproval,
        Approval
    }

    struct Inference {
        uint256 value; // this value is calculated by msg.value - feeL2 - feeTreasury
        uint256 feeL2;
        uint256 feeTreasury;
        address modelAddress;
        uint40 submitTimeout; // limit time to capture the miner role and submit the solution
        InferenceStatus status;
        address creator;
        address processedMiner;
        address referrer;
        bytes input;
        bytes output;
    }

    struct VotingInfo {
        uint8 totalCommit;
        uint8 totalReveal;
    }

    event NewInference(
        uint256 indexed inferenceId,
        address indexed model,
        address indexed creator,
        uint256 value,
        uint256 originInferenceId
    );

    event SolutionSubmission(address indexed miner, uint256 indexed inferId);

    function getInferenceInfo(
        uint256 _inferenceId
    ) external view returns (Inference memory);
}

interface IPromptScheduler3TX is IInferable {
    enum InferenceStatus {
        Nil,
        Solving,
        Commit,
        Reveal,
        Processed,
        Killed,
        Transferred
    }

    enum AssignmentRole {
        Nil,
        Validating,
        Mining
    }

    enum Vote {
        Nil,
        Disapproval,
        Approval
    }

    struct Assignment {
        uint256 inferenceId;
        bytes32 commitment;
        bytes32 digest; // keccak256(output)
        uint40 revealNonce;
        address worker;
        AssignmentRole role;
        Vote vote;
        bytes output;
    }

    struct Inference {
        uint256[] assignments;
        bytes input;
        uint256 value; // this value is calculated by msg.value - feeL2 - feeTreasury
        uint256 feeL2;
        uint256 feeTreasury;
        address modelAddress;
        uint40 submitTimeout; // limit time to capture the miner role and submit the solution
        uint40 commitTimeout;
        uint40 revealTimeout;
        InferenceStatus status;
        address creator;
        address processedMiner;
        address referrer;
    }

    struct VotingInfo {
        uint8 totalCommit;
        uint8 totalReveal;
    }

    event NewInference(
        uint256 indexed inferenceId,
        address indexed model,
        address indexed creator,
        uint256 value,
        uint256 originInferenceId
    );
    event NewAssignment(
        uint256 indexed assignmentId,
        uint256 indexed inferenceId,
        address indexed miner,
        uint40 expiredAt
    );
    event SolutionSubmission(
        address indexed miner,
        uint256 indexed assigmentId
    );
    function getInferenceInfo(
        uint256 _inferenceId
    ) external view returns (Inference memory);

    function getAssignmentInfo(
        uint256 _assignmentId
    ) external view returns (Assignment memory);
}
