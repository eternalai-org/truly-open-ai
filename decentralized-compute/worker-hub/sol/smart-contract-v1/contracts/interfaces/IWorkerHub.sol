// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {IInferable} from "./IInferable.sol";

interface IWorkerHub is IInferable {
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
        bytes32 digest;
        uint40 revealNonce;
        address worker;
        AssignmentRole role;
        Vote vote;
        bytes output;
    }

    struct Inference {
        uint256[] assignments;
        bytes input;
        uint256 value;
        uint256 feeL2;
        uint256 feeTreasury;
        address modelAddress;
        uint40 submitTimeout;
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

    struct DAOTokenPercentage {
        uint16 minerPercentage;
        uint16 userPercentage;
        uint16 referrerPercentage;
        uint16 refereePercentage;
        uint16 l2OwnerPercentage;
    }

    event NewInference(
        uint256 indexed inferenceId,
        address indexed model,
        address indexed creator,
        uint256 value,
        uint256 originInferenceId
    );
    event RawSubmitted(
        uint256 indexed inferenceId,
        address indexed model,
        address indexed creator,
        uint256 value,
        uint256 originInferenceId,
        bytes input,
        bool flag
    );
    event InferenceStatusUpdate(
        uint256 indexed inferenceId,
        InferenceStatus newStatus
    );
    event NewAssignment(
        uint256 indexed assignmentId,
        uint256 indexed inferenceId,
        address indexed miner,
        uint40 expiredAt
    );
    event MinerRoleSeized(
        uint256 indexed assignmentId,
        uint256 indexed inferenceId,
        address indexed miner
    );
    event SolutionSubmission(
        address indexed miner,
        uint256 indexed assigmentId
    );
    event CommitmentSubmission(
        address indexed miner,
        uint256 indexed assigmentId,
        bytes32 commitment
    );

    event RevealSubmission(
        address indexed miner,
        uint256 indexed assigmentId,
        uint40 nonce,
        bytes output
    );

    event DAOTokenPercentageUpdated(
        DAOTokenPercentage oldValue,
        DAOTokenPercentage newValue
    );
    enum DAOTokenReceiverRole {
        Miner,
        Validator,
        User,
        Referrer,
        Referee,
        L2Owner
    }
    struct DAOTokenReceiverInfor {
        address receiver;
        uint256 amount;
        DAOTokenReceiverRole role;
    }
    event DAOTokenMintedV2(
        uint256 chainId,
        uint256 inferenceId,
        address modelAddress,
        DAOTokenReceiverInfor[] receivers
    );
    event StreamedData(uint256 indexed assignmentId, bytes data);

    error AlreadySubmitted();
    error SubmitTimeout();
    error NotEnoughMiners();
    error Unauthorized();
    error OnlyAssignedWorker();
    error AlreadySeized();
    error InvalidContext();
    error InvalidInferenceStatus();
    error CannotFastForward();
    error InvalidMiner();
    error InvalidData();
    error InvalidRole();
    error InvalidCommitment();
    error AlreadyCommitted();
    error NotCommitted();
    error CommitTimeout();
    error RevealTimeout();
    error InvalidReveal();
    error InvalidNonce();
    error AlreadyRevealed();
    error InvalidAddress();

    function getMinFeeToUse(
        address _modelAddress
    ) external view returns (uint256);
    function getTreasuryAddress() external view returns (address);
    function getInferenceInfo(
        uint256 _inferenceId
    ) external view returns (Inference memory);
    function getAssignmentInfo(
        uint256 _assignmentId
    ) external view returns (Assignment memory);
}
