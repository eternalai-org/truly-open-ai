// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";

import {Random} from "./library/Random.sol";
import {Set} from "./library/Set.sol";
import {TransferHelper} from "./library/TransferHelper.sol";
import {GPUManagerStorage} from "./storages/GPUManagerStorage.sol";
import {IModelCollection} from "./interfaces/IModelCollection.sol";
import {IScheduler} from "./interfaces/IScheduler.sol";

contract GPUManager is
    GPUManagerStorage,
    OwnableUpgradeable,
    PausableUpgradeable,
    ReentrancyGuardUpgradeable
{
    using Random for Random.Randomizer;
    using Set for Set.AddressSet;
    using Set for Set.Uint256Set;

    string private constant VERSION = "v0.0.2";
    uint256 internal constant PERCENTAGE_DENOMINATOR = 100_00;
    uint256 private constant BLOCK_PER_YEAR = 365 days; // 1s per block

    receive() external payable {}

    modifier onlyOwnerOrPromptScheduler() {
        require(
            msg.sender == owner() || msg.sender == _promptScheduler,
            "Only Owner or PromptScheduler"
        );
        _;
    }

    modifier onlyPromptScheduler() {
        require(msg.sender == _promptScheduler, "Only PromptScheduler");
        _;
    }

    function initialize(
        address wEAIToken_,
        address modelCollection_,
        address treasury_,
        uint256 minerMinimumStake_,
        uint256 blocksPerEpoch_,
        uint256 rewardPerEpoch_,
        uint40 unstakeDelayTime_,
        uint40 penaltyDuration_,
        uint16 finePercentage_,
        uint256 minFeeToUse_
    ) external initializer {
        if (
            wEAIToken_ == address(0) ||
            modelCollection_ == address(0) ||
            treasury_ == address(0)
        ) revert InvalidAddress();

        __Ownable_init();
        __Pausable_init();
        __ReentrancyGuard_init();

        _minerMinimumStake = minerMinimumStake_;
        _blocksPerEpoch = blocksPerEpoch_;
        _rewardPerEpoch = rewardPerEpoch_;
        _unstakeDelayTime = unstakeDelayTime_;
        _penaltyDuration = penaltyDuration_;
        _finePercentage = finePercentage_;
        _minFeeToUse = minFeeToUse_;

        _maximumTier = 1;
        _lastBlock = block.number;

        _wEAIToken = wEAIToken_;
        _modelCollection = modelCollection_;
        _treasury = treasury_;
    }

    function setMinerMinimumStake(
        uint256 _minerMinimumStake
    ) external onlyOwner {
        _updateEpoch();

        _minerMinimumStake = _minerMinimumStake;
    }

    function setWEAIAddress(address wEAITokenAddress) external onlyOwner {
        _updateEpoch();

        if (wEAITokenAddress == address(0)) revert InvalidAddress();
        _wEAIToken = wEAITokenAddress;
    }

    function registerModel(
        uint32 modelId,
        uint16 tier,
        uint256 minimumFee
    ) external onlyOwner {
        _updateEpoch();

        if (modelId == 0 && modelId > type(uint32).max) revert InvalidModel();
        if (minimumFee < _minFeeToUse) revert FeeTooLow(); // NOTE: the minimum fee of using models must be greater than or equal _minFeeToUse
        if (tier == 0) revert InvalidTier();
        if (!IModelCollection(_modelCollection).checkModelExist(modelId))
            revert InvalidModel();

        Model storage model = _models[modelId];
        if (model.tier != 0) revert AlreadyRegistered();

        model.minimumFee = minimumFee;
        model.tier = tier;
        _modelIds.insert(modelId);

        emit ModelRegistration(modelId, tier, minimumFee);
    }

    function unregisterModel(uint32 modelId) external onlyOwner {
        _updateEpoch();

        Model storage model = _models[modelId];
        if (model.tier == 0) revert NotRegistered();

        model.tier = 0;
        _modelIds.erase(modelId);

        emit ModelUnregistration(modelId);
    }

    function updateModelTier(uint32 modelId, uint32 tier) external onlyOwner {
        _updateEpoch();

        if (tier == 0) revert InvalidTier();

        Model storage model = _models[modelId];
        if (model.tier == 0) revert InvalidModel();

        model.tier = tier;

        emit ModelTierUpdate(modelId, tier);
    }

    function updateModelMinimumFee(
        uint32 modelId,
        uint256 minimumFee
    ) external onlyOwner {
        _updateEpoch();

        Model storage model = _models[modelId];
        if (model.tier == 0) revert InvalidModel();

        model.minimumFee = minimumFee;

        emit ModelMinimumFeeUpdate(modelId, minimumFee);
    }

    function registerMiner(uint16 tier) external whenNotPaused {
        uint256 modelId = _modelIds.values[
            _randomizer.randomUint256() % _modelIds.size()
        ];

        registerMiner(tier, uint32(modelId));
    }

    function registerMiner(uint16 tier, uint32 modelId) public whenNotPaused {
        _updateEpoch();

        if (tier == 0 || tier > _maximumTier) revert InvalidTier();

        Worker storage miner = _miners[msg.sender];
        if (miner.tier != 0) revert AlreadyRegistered();

        miner.stake = _minerMinimumStake;
        miner.tier = tier;
        miner.modelId = modelId;

        TransferHelper.safeTransferFrom(
            _wEAIToken,
            msg.sender,
            address(this),
            _minerMinimumStake
        );

        emit MinerRegistration(msg.sender, tier, _minerMinimumStake);
    }

    function forceChangeModelForMiner(
        address miner,
        uint32 modelId
    ) external onlyOwner {
        _updateEpoch();

        if (modelId == 0) revert InvalidModel();
        if (_models[modelId].tier == 0) revert InvalidModel();
        if (!_minerAddresses.hasValue(miner)) revert NotRegistered();

        uint32 currentModelId = _miners[miner].modelId;
        if (currentModelId == modelId) revert SameModelAddress();
        _minerAddressesByModel[currentModelId].erase(miner);
        _minerAddressesByModel[modelId].insert(miner);

        _miners[miner].modelId = modelId;
        _miners[miner].tier = uint16(_models[modelId].tier);
    }

    function joinForMinting() external whenNotPaused {
        _updateEpoch();

        Worker storage miner = _miners[msg.sender];
        if (miner.tier == 0) revert NotRegistered();
        if (miner.stake < _minerMinimumStake) revert StakeTooLow();
        if (block.timestamp < miner.activeTime)
            revert MinerInDeactivationTime();

        _minerAddressesByModel[miner.modelId].insert(msg.sender);
        _minerAddresses.insert(msg.sender);
        miner.lastClaimedEpoch = _currentEpoch;
        _boost[msg.sender].minerTimestamp = uint40(block.timestamp);

        emit MinerJoin(msg.sender);
    }

    function unregisterMiner() external nonReentrant whenNotPaused {
        _updateEpoch();

        Worker storage miner = _miners[msg.sender];
        if (miner.tier == 0) revert NotRegistered();

        miner.tier = 0;
        uint stakeAmount = miner.stake;
        miner.stake = 0;

        _updateMinerState(msg.sender, miner.modelId, true);
        miner.modelId = 0;

        uint currentUnstake = _minerUnstakeRequests[msg.sender].stake;
        _minerUnstakeRequests[msg.sender] = UnstakeRequest(
            stakeAmount + currentUnstake,
            uint40(block.number + _unstakeDelayTime)
        );

        emit MinerUnregistration(msg.sender);
    }

    function increaseMinerStake(uint256 wEAIAmt) external whenNotPaused {
        _updateEpoch();

        Worker storage miner = _miners[msg.sender];
        if (miner.tier == 0) revert NotRegistered();

        TransferHelper.safeTransferFrom(
            _wEAIToken,
            msg.sender,
            address(this),
            wEAIAmt
        );
        miner.stake += wEAIAmt;

        emit MinerExtraStake(msg.sender, wEAIAmt);
    }

    function unstakeForMiner() external {
        _updateEpoch();

        UnstakeRequest storage unstakeRequest = _minerUnstakeRequests[
            msg.sender
        ];
        if (block.number < unstakeRequest.unlockAt) revert StillBeingLocked();

        uint256 stake = unstakeRequest.stake;
        if (stake == 0) revert ZeroValue();
        unstakeRequest.stake = 0;
        TransferHelper.safeTransfer(_wEAIToken, msg.sender, stake);

        emit MinerUnstake(msg.sender, stake);
    }

    function restakeForMiner(uint16 tier) external whenNotPaused {
        _updateEpoch();

        UnstakeRequest storage unstakeRequest = _minerUnstakeRequests[
            msg.sender
        ];
        if (unstakeRequest.stake == 0) revert ZeroValue();
        uint unstakeAmount = unstakeRequest.stake;
        unstakeRequest.stake = 0;

        Worker storage miner = _miners[msg.sender];
        miner.stake += unstakeAmount;
        if (miner.tier == 0) {
            if (tier == 0 || tier > _maximumTier) revert InvalidTier();
            miner.tier = tier;
        }

        if (miner.modelId == 0) {
            uint256 modelId = _modelIds.values[
                _randomizer.randomUint256() % _modelIds.size()
            ];
            miner.modelId = uint32(modelId);
        }

        emit Restake(msg.sender, miner.modelId, unstakeAmount);
    }

    function updateEpoch() external onlyPromptScheduler {
        _updateEpoch();
    }

    // update new epoch
    function _updateEpoch() internal {
        if (_blocksPerEpoch > 0) {
            uint256 epochPassed = (block.number - _lastBlock) / _blocksPerEpoch;
            if (epochPassed > 0) {
                _lastBlock += _blocksPerEpoch * epochPassed;
                // reward for this epoch
                // _rewardPerEpoch (reward one year for 1 miner)
                // _rewardPerEpoch * total miner * blocker per epoch / blocks per year
                uint256 rewardInCurrentEpoch = (_rewardPerEpoch *
                    _minerAddresses.size() *
                    _blocksPerEpoch) / BLOCK_PER_YEAR;

                for (; epochPassed > 0; epochPassed--) {
                    _rewardInEpoch[_currentEpoch].totalMiner = _minerAddresses
                        .size();
                    _rewardInEpoch[_currentEpoch]
                        .epochReward = rewardInCurrentEpoch;
                    _currentEpoch++;
                }
            }
        } else {
            _lastBlock = block.number;
        }
    }

    function slashMiner(
        address miner,
        bool isFined
    ) public virtual onlyOwnerOrPromptScheduler {
        _updateEpoch();

        if (miner == address(0)) revert InvalidMiner();

        _slashMiner(miner, isFined);
    }

    function _updateMinerState(
        address miner,
        uint32 modelId,
        bool isUnregister
    ) internal {
        _claimReward(miner, false);
        _boost[miner].minerTimestamp = uint40(block.timestamp);

        if (isUnregister) {
            _boost[miner].reserved1 = 0;
        } else {
            _boost[miner].reserved1 += (uint48(block.timestamp) -
                _boost[miner].minerTimestamp);
        }

        if (_minerAddressesByModel[modelId].hasValue(miner)) {
            _minerAddressesByModel[modelId].erase(miner);
            _minerAddresses.erase(miner);
        }
    }

    function _slashMiner(address miner, bool isFined) internal {
        Worker storage minerInfo = _miners[miner];

        uint32 modelId = minerInfo.modelId;
        _updateMinerState(miner, modelId, false);

        // Set the time minerInfo can join again
        minerInfo.activeTime = uint40(block.timestamp + _penaltyDuration);

        if (isFined) {
            uint256 fine = (_minerMinimumStake * _finePercentage) /
                PERCENTAGE_DENOMINATOR;
            uint256 collectedFine = 0;
            uint256 pendingUnstakeAmt = _minerUnstakeRequests[miner].stake;
            uint256 totalStake = minerInfo.stake + pendingUnstakeAmt;

            if (totalStake <= fine) {
                collectedFine = totalStake;
                minerInfo.stake = 0;
                _minerUnstakeRequests[miner].stake = 0;
            } else {
                if (minerInfo.stake >= fine) {
                    minerInfo.stake -= fine;
                    collectedFine = fine;
                } else {
                    uint256 remainingFine = fine - minerInfo.stake;
                    collectedFine = fine;
                    minerInfo.stake = 0;
                    _minerUnstakeRequests[miner].stake -= remainingFine;
                }
            }

            // reset _boost
            _boost[miner].reserved1 = 0;

            TransferHelper.safeTransfer(_wEAIToken, _treasury, collectedFine);

            emit FraudulentMinerPenalized(
                miner,
                modelId,
                _treasury,
                collectedFine
            );
            return;
        }

        emit MinerDeactivated(miner, modelId, minerInfo.activeTime);
    }

    function _claimReward(
        address miner,
        bool isTransfer
    ) internal whenNotPaused {
        uint256 rewardAmount = rewardToClaim(miner);
        _miners[miner].lastClaimedEpoch = _currentEpoch;
        if (rewardAmount > 0 && isTransfer) {
            _minerRewards[miner] = 0;
            TransferHelper.safeTransfer(_wEAIToken, miner, rewardAmount);

            emit RewardClaim(miner, rewardAmount);
        } else if (rewardAmount > 0) {
            _minerRewards[miner] = rewardAmount;
        }
    }

    // miner claim reward
    function claimReward(address miner) external virtual nonReentrant {
        _claimReward(miner, true);
    }

    // sum reward of an miner since last claimed epoch
    function rewardToClaim(address miner) public virtual returns (uint256) {
        _updateEpoch();

        uint256 totalReward;
        uint256 lastEpoch = _currentEpoch;
        if (
            !_minerAddresses.hasValue(miner) ||
            lastEpoch <= _miners[miner].lastClaimedEpoch
        ) {
            totalReward = 0;
        } else {
            uint256 lastClaimed = uint256(_miners[miner].lastClaimedEpoch);
            uint256 epochReward = (_rewardPerEpoch * _blocksPerEpoch) /
                BLOCK_PER_YEAR; // reward per miner in 1 epoch
            totalReward +=
                ((lastEpoch - lastClaimed) * epochReward * multiplier(miner)) /
                PERCENTAGE_DENOMINATOR;
        }

        return totalReward + _minerRewards[miner];
    }

    function multiplier(address miner) public view returns (uint256) {
        uint256 minerLastTimestamp;

        if (
            _minerAddresses.hasValue(miner) && _boost[miner].minerTimestamp == 0
        ) {
            minerLastTimestamp = 1716046859;
        } else if (!_minerAddresses.hasValue(miner)) {
            minerLastTimestamp = block.timestamp;
        } else {
            minerLastTimestamp = _boost[miner].minerTimestamp;
        }
        uint256 multiplierRes = (_boost[miner].reserved1 +
            block.timestamp -
            minerLastTimestamp) / 30 days;

        return
            PERCENTAGE_DENOMINATOR +
            500 *
            (multiplierRes >= 12 ? 12 : multiplierRes);
    }

    function setFinePercentage(
        uint16 newPercentage
    ) external virtual onlyOwner {
        _updateEpoch();

        emit FinePercentageUpdated(_finePercentage, newPercentage);

        _finePercentage = newPercentage;
    }

    function setPenaltyDuration(uint40 duration) external virtual onlyOwner {
        _updateEpoch();

        emit PenaltyDurationUpdated(_penaltyDuration, duration);

        _penaltyDuration = duration;
    }

    function setMinFeeToUse(uint256 minFee) external virtual onlyOwner {
        _updateEpoch();

        emit MinFeeToUseUpdated(_minFeeToUse, minFee);

        _minFeeToUse = minFee;
    }

    // @dev admin functions
    function setNewRewardInEpoch(uint256 newReward) external virtual onlyOwner {
        _updateEpoch();
        emit RewardPerEpoch(_rewardPerEpoch, newReward);

        _rewardPerEpoch = newReward;
    }

    function setBlocksPerEpoch(uint256 blocks) external virtual onlyOwner {
        _updateEpoch();
        if (blocks == 0) revert InvalidBlockValue();

        emit BlocksPerEpoch(_blocksPerEpoch, blocks);

        _blocksPerEpoch = blocks;
    }

    function setUnstakeDelayTime(uint40 delayTime) external virtual onlyOwner {
        _updateEpoch();

        if (delayTime == 0) revert InvalidValue();

        emit UnstakeDelayTime(_unstakeDelayTime, delayTime);

        _unstakeDelayTime = delayTime;
    }

    function setPromptSchedulerAddress(
        address newPromptScheduler
    ) external onlyOwner {
        _updateEpoch();

        if (newPromptScheduler == address(0)) revert InvalidAddress();
        _promptScheduler = newPromptScheduler;
    }

    function getMinFeeToUse(uint32 modelId) external view returns (uint256) {
        return _models[modelId].minimumFee;
    }

    function getModelInfo(uint32 modelId) external view returns (Model memory) {
        return _models[modelId];
    }

    function getNOMiner() external view returns (uint) {
        return _minerAddresses.values.length;
    }

    function getMinerAddresses() external view returns (address[] memory) {
        return _minerAddresses.values;
    }

    function validateMiner(address miner) external {
        if (!_minerAddresses.hasValue(miner)) revert InvalidMiner();

        uint32 modelId = _miners[miner].modelId;
        if (!_minerAddressesByModel[modelId].hasValue(miner))
            revert InvalidModel();

        _updateEpoch();
    }

    function validateModelAndChooseRandomMiner(
        uint32 modelId,
        uint256 minersRequired
    ) external returns (address, uint256) {
        if (_models[modelId].tier == 0) revert InvalidModel();

        uint256 minerSize = _minerAddressesByModel[modelId].size();
        if (minerSize < minersRequired) revert NotEnoughMiners();

        _updateEpoch();

        uint8 index = uint8(_randomizer.randomUint256() % minerSize);
        return (
            _minerAddressesByModel[modelId].values[index],
            _models[modelId].minimumFee
        );
    }

    function getModelIds() external view returns (uint256[] memory) {
        return _modelIds.values;
    }

    function getMinerAddressesOfModel(
        uint32 modelId
    ) external view returns (address[] memory) {
        return _minerAddressesByModel[modelId].values;
    }

    function getAllMinerUnstakeRequests()
        external
        view
        returns (
            address[] memory unstakeAddresses,
            UnstakeRequest[] memory unstakeRequests
        )
    {
        address[] memory addresses = _minerAddresses.values;

        uint countUnstakeRequest = 0;
        for (uint i = 0; i < addresses.length; ++i) {
            UnstakeRequest memory request = _minerUnstakeRequests[addresses[i]];
            if (request.unlockAt > 0) ++countUnstakeRequest;
        }

        unstakeAddresses = new address[](countUnstakeRequest);
        unstakeRequests = new UnstakeRequest[](countUnstakeRequest);
        uint idx = 0;
        for (uint i = 0; i < addresses.length; ++i) {
            UnstakeRequest memory request = _minerUnstakeRequests[addresses[i]];
            if (request.unlockAt > 0) {
                unstakeAddresses[idx] = addresses[idx];
                unstakeRequests[idx] = request;
                ++idx;
            }
        }
    }

    function isActiveModel(uint32 modelId) external view returns (bool) {
        uint8 minerRequirement = IScheduler(_promptScheduler)
            .getMinerRequirement();
        return
            _minerAddressesByModel[modelId].values.length >= minerRequirement;
    }

    function wEAIToken() external view returns (address) {
        return _wEAIToken;
    }

    function modelCollection() external view returns (address) {
        return _modelCollection;
    }

    function promptScheduler() external view returns (address) {
        return _promptScheduler;
    }

    function treasury() external view returns (address) {
        return _treasury;
    }

    function models(uint32 modelId) external view returns (Model memory) {
        return _models[modelId];
    }

    function miners(address miner) external view returns (Worker memory) {
        return _miners[miner];
    }

    function minerUnstakeRequests(address miner) external view returns (UnstakeRequest memory) {
        return _minerUnstakeRequests[miner];
    }

    function rewardInEpoch(uint256 epoch) external view returns (MinerEpochState memory) {
        return _rewardInEpoch[epoch];
    }

    function minFeeToUse() external view returns (uint256) {
        return _minFeeToUse;
    }

    function minerMinimumStake() external view returns (uint256) {
        return _minerMinimumStake;
    }

    function unstakeDelayTime() external view returns (uint40) {
        return _unstakeDelayTime;
    }

    function penaltyDuration() external view returns (uint40) {
        return _penaltyDuration;
    }

    function finePercentage() external view returns (uint16) {
        return _finePercentage;
    }

    function maximumTier() external view returns (uint16) {
        return _maximumTier;
    }

    function currentEpoch() external view returns (uint40) {
        return _currentEpoch;
    }

    function blocksPerEpoch() external view returns (uint256) {
        return _blocksPerEpoch;
    }

    function lastBlock() external view returns (uint256) {
        return _lastBlock;
    }

    function rewardPerEpoch() external view returns (uint256) {
        return _rewardPerEpoch;
    }
}

