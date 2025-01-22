// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {IGPUManager} from "../interfaces/IGPUManager.sol";
import {Random} from "../lib/Random.sol";
import {Set} from "../lib/Set.sol";

abstract contract GPUManagerStorage is IGPUManager {
    Random.Randomizer internal _randomizer;
    address public _wEAIToken;
    address public _modelCollection;
    address public _promptScheduler;
    address public _treasury;

    mapping(uint32 => Model) public _models;
    mapping(address => Worker) public _miners;
    mapping(uint256 => Set.AddressSet) internal _minerAddressesByModel;

    Set.Uint256Set internal _modelIds;
    Set.AddressSet internal _minerAddresses;

    mapping(address => UnstakeRequest) public _minerUnstakeRequests;
    mapping(uint256 => MinerEpochState) public _rewardInEpoch;

    uint256 public _minFeeToUse; // The minimum fee when register model, it's also the minimum fee to create inference
    uint256 public _minerMinimumStake;
    uint40 public _unstakeDelayTime;
    uint40 public _penaltyDuration;
    uint16 public _finePercentage;
    uint16 public _maximumTier;

    // reward purpose
    uint40 public _currentEpoch;
    uint256 public _blocksPerEpoch;
    uint256 public _lastBlock;
    uint256 public _rewardPerEpoch; // 12299.97 reward EAI for 1 worker per year

    // mapping tracking reward
    mapping(address => uint256) internal _minerRewards;
    // tracking time miner join the network to
    // determine multiplier value
    mapping(address => Boost) internal _boost;

    uint256[100] private __gap;
}
