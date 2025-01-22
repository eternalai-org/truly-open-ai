// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {IWorkerHub} from "../interfaces/IWorkerHub.sol";

import {Random} from "../library/Random.sol";
import {Set} from "../library/Set.sol";

abstract contract WorkerHubStorage is IWorkerHub {
    Random.Randomizer internal randomizer;

    uint256 public inferenceNumber;
    mapping(uint256 => Inference) internal inferences;

    uint256 public assignmentNumber;
    mapping(uint256 => Assignment) public assignments;
    mapping(uint256 inferId => VotingInfo) internal votingInfo;
    mapping(uint256 inferId => Set.Bytes32Set) internal digests;
    mapping(bytes32 commitment => uint8) internal countDigest;
    mapping(address => Set.Uint256Set) internal assignmentsByMiner;
    mapping(uint256 => Set.Uint256Set) internal assignmentsByInference;

    address internal l2Owner;
    address internal treasury;
    uint16 internal feeL2Percentage;
    uint16 internal feeTreasuryPercentage;
    uint16 internal feeRatioMinerValidator;
    uint40 internal submitDuration;
    uint40 internal commitDuration;
    uint40 internal revealDuration;

    uint8 internal minerRequirement;

    address internal daoToken;
    uint256 internal daoTokenReward;
    DAOTokenPercentage internal daoTokenPercentage;

    mapping(address => address) internal referrerOf;

    mapping(uint256 inferId => DAOTokenReceiverInfor[])
        internal daoReceiversInfo;
    address internal wEAI;
    address internal stakingHub;

    uint256[100] private __gap;
}
