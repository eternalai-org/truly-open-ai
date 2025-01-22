// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {IStakingHub} from "../interfaces/IStakingHub.sol";

import {Random} from "../library/Random.sol";
import {Set} from "../library/Set.sol";

abstract contract StakingHubStorage is IStakingHub {
    Random.Randomizer internal randomizer;
    address public treasury;

    mapping(address => Model) public models;
    mapping(address => Worker) public miners;
    mapping(address => Set.AddressSet) internal minerAddressesByModel;

    Set.AddressSet internal modelAddresses;
    Set.AddressSet internal minerAddresses;

    mapping(address => UnstakeRequest) public minerUnstakeRequests;
    mapping(uint256 => MinerEpochState) public rewardInEpoch;

    uint256 public minerMinimumStake;
    uint40 public unstakeDelayTime;
    uint40 public penaltyDuration;
    uint16 public finePercentage;
    uint16 public maximumTier;

    // reward purpose
    uint40 public currentEpoch;
    uint256 public blocksPerEpoch;
    uint256 public lastBlock;
    uint256 public rewardPerEpoch;

    mapping(address => uint256) internal minerRewards;
    mapping(address => Boost) internal boost;
    address public wEAI;
    address workerHub;
    uint256 public minFeeToUse; // The minimum fee when register model, it's also the minimum fee to create inference

    uint256[100] private __gap;
}
