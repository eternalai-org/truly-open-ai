// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {IGPUManager} from "../interfaces/IGPUManager.sol";
import {Random} from "../library/Random.sol";
import {Set} from "../library/Set.sol";

abstract contract GPUManagerStorage is IGPUManager {
    Random.Randomizer internal _randomizer;
    address internal _wEAIToken;
    address internal _modelCollection;
    address internal _promptScheduler;
    address internal _treasury;

    mapping(uint32 => Model) internal _models;
    mapping(address => Worker) internal _miners;
    mapping(uint256 => Set.AddressSet) internal _minerAddressesByModel;

    Set.Uint256Set internal _modelIds;
    Set.AddressSet internal _minerAddresses;

    mapping(address => UnstakeRequest) internal _minerUnstakeRequests;
    mapping(uint256 => MinerEpochState) internal _rewardInEpoch;

    uint256 internal _minFeeToUse; // The minimum fee when register model, it's also the minimum fee to create inference
    uint256 internal _minerMinimumStake;
    uint40 internal _unstakeDelayTime;
    uint40 internal _penaltyDuration;
    uint16 internal _finePercentage;
    uint16 internal _maximumTier;

    // reward purpose
    uint40 internal _currentEpoch;
    uint256 internal _blocksPerEpoch;
    uint256 internal _lastBlock;
    uint256 internal _rewardPerEpoch; // 12299.97 reward EAI for 1 worker per year
    
    // mapping tracking reward
    mapping(address => uint256) internal _minerRewards;
    // tracking time miner join the network to
    // determine multiplier value
    mapping(address => Boost) internal _boost;

    uint256[100] private __gap;
}
