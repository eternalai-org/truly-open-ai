// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";

import {Random} from "./library/Random.sol";
import {Set} from "./library/Set.sol";
import {TransferHelper} from "./library/TransferHelper.sol";
import {StakingHubStorage} from "./storages/StakingHubStorage.sol";
import {IWorkerHub} from "./interfaces/IWorkerHub.sol";

contract StakingHub is
    StakingHubStorage,
    OwnableUpgradeable,
    PausableUpgradeable,
    ReentrancyGuardUpgradeable
{
    using Random for Random.Randomizer;
    using Set for Set.AddressSet;
    using Set for Set.Uint256Set;
    using Set for Set.Bytes32Set;

    string private constant VERSION = "v0.0.2";
    uint256 internal constant PERCENTAGE_DENOMINATOR = 100_00;
    uint256 private constant BLOCK_PER_YEAR = 365 days / 2; // 2s per block

    receive() external payable {}

    modifier onlyWorkerHub() {
        require(msg.sender == workerHub, "Only WorkerHub");
        _;
    }

    modifier onlyOwnerOrWorkerHub() {
        require(
            msg.sender == owner() || msg.sender == workerHub,
            "Only Owner or WorkerHub"
        );
        _;
    }

    function initialize(
        address _wEAI,
        uint256 _minerMinimumStake,
        uint256 _blocksPerEpoch,
        uint256 _rewardPerEpoch,
        uint40 _unstakeDelayTime,
        uint40 _penaltyDuration,
        uint16 _finePercentage,
        uint256 _minFeeToUse
    ) external initializer {
        __Ownable_init();
        __Pausable_init();
        __ReentrancyGuard_init();

        if (_wEAI == address(0)) revert InvalidAddress();

        minerMinimumStake = _minerMinimumStake;
        blocksPerEpoch = _blocksPerEpoch;
        rewardPerEpoch = _rewardPerEpoch;
        unstakeDelayTime = _unstakeDelayTime;
        penaltyDuration = _penaltyDuration;
        finePercentage = _finePercentage;
        minFeeToUse = _minFeeToUse;

        maximumTier = 1;
        lastBlock = block.number;

        wEAI = _wEAI;
    }

    function setMinerMinimumStake(
        uint256 _minerMinimumStake
    ) external onlyOwner {
        _updateEpoch();

        minerMinimumStake = _minerMinimumStake;
    }

    function registerModel(
        address _model,
        uint16 _tier,
        uint256 _minimumFee
    ) external onlyOwner {
        _updateEpoch();

        if (_model == address(0)) revert InvalidModel();
        if (_minimumFee < minFeeToUse) revert FeeTooLow();
        if (_tier == 0) revert InvalidTier();

        Model storage model = models[_model];
        if (model.tier != 0) revert AlreadyRegistered();

        model.minimumFee = _minimumFee;
        model.tier = _tier;
        modelAddresses.insert(_model);

        emit ModelRegistration(_model, _tier, _minimumFee);
    }

    function unregisterModel(address _model) external onlyOwner {
        _updateEpoch();

        Model storage model = models[_model];
        if (model.tier == 0) revert NotRegistered();

        model.tier = 0;
        modelAddresses.erase(_model);

        emit ModelUnregistration(_model);
    }

    function updateModelTier(address _model, uint32 _tier) external onlyOwner {
        _updateEpoch();

        if (_tier == 0) revert InvalidTier();

        Model storage model = models[_model];
        if (model.tier == 0) revert InvalidModel();

        model.tier = _tier;

        emit ModelTierUpdate(_model, _tier);
    }

    function updateModelMinimumFee(
        address _model,
        uint256 _minimumFee
    ) external onlyOwner {
        _updateEpoch();

        Model storage model = models[_model];
        if (model.tier == 0) revert InvalidModel();

        model.minimumFee = _minimumFee;

        emit ModelMinimumFeeUpdate(_model, _minimumFee);
    }

    function registerMiner(uint16 tier) external whenNotPaused {
        _updateEpoch();

        if (tier == 0 || tier > maximumTier) revert InvalidTier();

        Worker storage miner = miners[msg.sender];
        if (miner.tier != 0) revert AlreadyRegistered();

        miner.stake = minerMinimumStake;
        miner.tier = tier;

        address modelAddress = modelAddresses.values[
            randomizer.randomUint256() % modelAddresses.size()
        ];
        miner.modelAddress = modelAddress;
        TransferHelper.safeTransferFrom(
            wEAI,
            msg.sender,
            address(this),
            minerMinimumStake
        );

        emit MinerRegistration(msg.sender, tier, minerMinimumStake);
    }

    function forceChangeModelForMiner(
        address _miner,
        address _modelAddress
    ) external onlyOwner {
        _updateEpoch();

        if (models[_modelAddress].tier == 0) revert InvalidModel();
        if (!minerAddresses.hasValue(_miner)) revert NotRegistered();

        address currentModelAddress = miners[_miner].modelAddress;
        if (currentModelAddress == _modelAddress) revert SameModelAddress();
        minerAddressesByModel[currentModelAddress].erase(_miner);
        minerAddressesByModel[_modelAddress].insert(_miner);

        miners[_miner].modelAddress = _modelAddress;
        miners[_miner].tier = uint16(models[_modelAddress].tier);
    }

    function joinForMinting() external whenNotPaused {
        _updateEpoch();

        Worker storage miner = miners[msg.sender];
        if (miner.tier == 0) revert NotRegistered();
        if (miner.stake < minerMinimumStake) revert StakeTooLow();
        if (block.timestamp < miner.activeTime)
            revert MinerInDeactivationTime();

        address modelAddress = miner.modelAddress;
        minerAddressesByModel[modelAddress].insert(msg.sender);
        minerAddresses.insert(msg.sender);
        miner.lastClaimedEpoch = currentEpoch;
        boost[msg.sender].minerTimestamp = uint40(block.timestamp);

        emit MinerJoin(msg.sender);
    }

    function unregisterMiner() external nonReentrant whenNotPaused {
        _updateEpoch();

        Worker storage miner = miners[msg.sender];
        if (miner.tier == 0) revert NotRegistered();

        miner.tier = 0;

        uint stakeAmount = miner.stake;
        miner.stake = 0;
        miner.commitment = 0;

        if (minerAddresses.hasValue(msg.sender)) {
            _claimReward(msg.sender, false);
            // reset boost
            boost[msg.sender].reserved1 = 0;
            boost[msg.sender].minerTimestamp = uint40(block.timestamp);

            minerAddresses.erase(msg.sender);
            minerAddressesByModel[miner.modelAddress].erase(msg.sender);
        }
        miner.modelAddress = address(0);

        uint currentUnstake = minerUnstakeRequests[msg.sender].stake;
        minerUnstakeRequests[msg.sender] = UnstakeRequest(
            stakeAmount + currentUnstake,
            uint40(block.number + unstakeDelayTime)
        );

        emit MinerUnregistration(msg.sender);
    }

    function increaseMinerStake(uint256 wEAIAmt) external whenNotPaused {
        _updateEpoch();

        Worker storage miner = miners[msg.sender];
        if (miner.tier == 0) revert NotRegistered();

        miner.stake += wEAIAmt;
        TransferHelper.safeTransferFrom(
            wEAI,
            msg.sender,
            address(this),
            wEAIAmt
        );

        emit MinerExtraStake(msg.sender, wEAIAmt);
    }

    function unstakeForMiner() external {
        _updateEpoch();

        UnstakeRequest storage unstakeRequest = minerUnstakeRequests[
            msg.sender
        ];
        if (block.number < unstakeRequest.unlockAt) revert StillBeingLocked();

        uint256 stake = unstakeRequest.stake;
        if (stake == 0) revert NullStake();
        unstakeRequest.stake = 0;
        TransferHelper.safeTransfer(wEAI, msg.sender, stake);

        emit MinerUnstake(msg.sender, stake);
    }

    function restakeForMiner(uint16 tier) external whenNotPaused {
        _updateEpoch();

        UnstakeRequest storage unstakeRequest = minerUnstakeRequests[
            msg.sender
        ];
        if (unstakeRequest.stake == 0) revert ZeroValue();
        uint unstakeAmount = unstakeRequest.stake;
        unstakeRequest.stake = 0;

        Worker storage miner = miners[msg.sender];
        miner.stake += unstakeAmount;
        if (miner.tier == 0) {
            if (tier == 0 || tier > maximumTier) revert InvalidTier();
            miner.tier = tier;
        }

        if (miner.modelAddress == address(0)) {
            address modelAddress = modelAddresses.values[
                randomizer.randomUint256() % modelAddresses.size()
            ];
            miner.modelAddress = modelAddress;
        }

        emit Restake(msg.sender, unstakeAmount, miner.modelAddress);
    }

    function updateEpoch() external onlyWorkerHub {
        _updateEpoch();
    }

    // this internal function update new epoch
    function _updateEpoch() internal {
        if (blocksPerEpoch > 0) {
            uint256 epochPassed = (block.number - lastBlock) / blocksPerEpoch;
            if (epochPassed > 0) {
                lastBlock += blocksPerEpoch * epochPassed;
                // reward for this epoch
                // rewardPerEpoch (reward one year for 1 miner)
                // rewardPerEpoch * total miner * blocker per epoch / blocks per year
                uint256 rewardInCurrentEpoch = (rewardPerEpoch *
                    minerAddresses.size() *
                    blocksPerEpoch) / BLOCK_PER_YEAR;

                for (; epochPassed > 0; epochPassed--) {
                    rewardInEpoch[currentEpoch].totalMiner = minerAddresses
                        .size();
                    rewardInEpoch[currentEpoch]
                        .epochReward = rewardInCurrentEpoch;
                    currentEpoch++;
                }
            }
        } else {
            lastBlock = block.number;
        }
    }

    function slashMiner(
        address _miner,
        bool _isFined
    ) public virtual onlyOwnerOrWorkerHub {
        _updateEpoch();

        if (_miner == address(0)) revert InvalidMiner();

        _slashMiner(_miner, _isFined);
    }

    function _slashMiner(address _miner, bool _isFined) internal {
        Worker storage miner = miners[_miner];

        if (!minerAddresses.hasValue(_miner)) revert InvalidMiner();
        // update reward
        _claimReward(_miner, false);
        boost[_miner].reserved1 +=
            uint48(block.timestamp) -
            uint48(
                boost[_miner].minerTimestamp == 0
                    ? 1716046859
                    : boost[_miner].minerTimestamp
            );
        boost[_miner].minerTimestamp = uint40(block.timestamp);
        address modelAddress = miner.modelAddress;

        // Remove miner from available miner
        if (minerAddressesByModel[modelAddress].hasValue(_miner)) {
            minerAddressesByModel[modelAddress].erase(_miner);
            minerAddresses.erase(_miner);
        }

        // Set the time miner can join again
        miner.activeTime = uint40(block.timestamp + penaltyDuration);

        if (_isFined) {
            uint256 fine = (minerMinimumStake * finePercentage) /
                PERCENTAGE_DENOMINATOR;
            if (miner.stake < fine) {
                miner.stake = 0;
            } else {
                miner.stake -= fine;
            }

            // reset boost
            boost[_miner].reserved1 = 0;
            address treasury = IWorkerHub(workerHub).getTreasuryAddress();

            TransferHelper.safeTransfer(wEAI, treasury, fine);

            emit FraudulentMinerPenalized(_miner, modelAddress, treasury, fine);
            return;
        }

        emit MinerDeactivated(_miner, modelAddress, miner.activeTime);
    }

    function _claimReward(
        address _miner,
        bool _isTransfer
    ) internal whenNotPaused {
        uint256 rewardAmount = rewardToClaim(_miner);
        miners[_miner].lastClaimedEpoch = currentEpoch;
        if (rewardAmount > 0 && _isTransfer) {
            minerRewards[_miner] = 0;
            TransferHelper.safeTransfer(wEAI, _miner, rewardAmount);

            emit RewardClaim(_miner, rewardAmount);
        } else if (rewardAmount > 0) {
            minerRewards[_miner] = rewardAmount;
        }
    }

    // miner claim reward
    function claimReward(address _miner) external virtual nonReentrant {
        _claimReward(_miner, true);
    }

    // sum reward of an miner since last claimed epoch
    function rewardToClaim(address _miner) public virtual returns (uint256) {
        _updateEpoch();

        uint256 totalReward;
        uint256 lastEpoch = currentEpoch;
        if (
            !minerAddresses.hasValue(_miner) ||
            lastEpoch <= miners[_miner].lastClaimedEpoch
        ) {
            totalReward = 0;
        } else {
            uint256 lastClaimed = uint256(miners[_miner].lastClaimedEpoch);
            uint256 epochReward = (rewardPerEpoch * blocksPerEpoch) /
                BLOCK_PER_YEAR; // reward per miner in 1 epoch
            totalReward +=
                ((lastEpoch - lastClaimed) * epochReward * multiplier(_miner)) /
                PERCENTAGE_DENOMINATOR;
        }

        return totalReward + minerRewards[_miner];
    }

    function multiplier(address _miner) public view returns (uint256) {
        uint256 minerLastTimestamp;

        if (
            minerAddresses.hasValue(_miner) && boost[_miner].minerTimestamp == 0
        ) {
            minerLastTimestamp = 1716046859;
        } else if (!minerAddresses.hasValue(_miner)) {
            minerLastTimestamp = block.timestamp;
        } else {
            minerLastTimestamp = boost[_miner].minerTimestamp;
        }
        uint256 multiplierRes = (boost[_miner].reserved1 +
            block.timestamp -
            minerLastTimestamp) / 30 days;

        return
            PERCENTAGE_DENOMINATOR +
            500 *
            (multiplierRes >= 12 ? 12 : multiplierRes);
    }

    function setFinePercentage(
        uint16 _finePercentage
    ) external virtual onlyOwner {
        _updateEpoch();

        emit FinePercentageUpdated(finePercentage, _finePercentage);

        finePercentage = _finePercentage;
    }

    function setPenaltyDuration(
        uint40 _penaltyDuration
    ) external virtual onlyOwner {
        _updateEpoch();

        emit PenaltyDurationUpdated(penaltyDuration, _penaltyDuration);

        penaltyDuration = _penaltyDuration;
    }

    function setMinFeeToUse(uint256 _minFeeToUse) external virtual onlyOwner {
        _updateEpoch();

        emit MinFeeToUseUpdated(minFeeToUse, _minFeeToUse);

        minFeeToUse = _minFeeToUse;
    }

    function setNewRewardInEpoch(
        uint256 _newRewardAmount
    ) external virtual onlyOwner {
        _updateEpoch();
        emit RewardPerEpoch(rewardPerEpoch, _newRewardAmount);

        rewardPerEpoch = _newRewardAmount;
    }

    function setBlocksPerEpoch(uint256 _blocks) external virtual onlyOwner {
        _updateEpoch();
        if (_blocks == 0) revert InvalidBlockValue();

        emit BlocksPerEpoch(blocksPerEpoch, _blocks);

        blocksPerEpoch = _blocks;
    }

    function setUnstakDelayTime(
        uint40 _newUnstakeDelayTime
    ) external virtual onlyOwner {
        _updateEpoch();

        if (_newUnstakeDelayTime == 0) revert InvalidValue();

        emit UnstakeDelayTime(unstakeDelayTime, _newUnstakeDelayTime);

        unstakeDelayTime = _newUnstakeDelayTime;
    }

    function setWorkerHubAddress(address _workerHub) external onlyOwner {
        _updateEpoch();

        if (_workerHub == address(0)) revert InvalidWorkerHub();
        workerHub = _workerHub;
    }

    function getMinFeeToUse(
        address _modelAddress
    ) external view returns (uint256) {
        return models[_modelAddress].minimumFee;
    }

    function getModelInfo(
        address _modelAddr
    ) external view returns (Model memory) {
        return models[_modelAddr];
    }

    function getNOMiner() external view returns (uint) {
        return minerAddresses.values.length;
    }

    function getMinerAddresses() external view returns (address[] memory) {
        return minerAddresses.values;
    }

    function isMinerAddress(address _miner) external view returns (bool) {
        return minerAddresses.hasValue(_miner);
    }

    function validateModelOfMiner(address _miner) external view {
        address modelAddrOfMiner = miners[_miner].modelAddress;
        if (!minerAddressesByModel[modelAddrOfMiner].hasValue(_miner))
            revert InvalidMiner();
    }

    function getModelAddresses() external view returns (address[] memory) {
        return modelAddresses.values;
    }

    function getMinerAddressesOfModel(
        address _model
    ) external view returns (address[] memory) {
        return minerAddressesByModel[_model].values;
    }

    function getAllMinerUnstakeRequests()
        external
        view
        returns (
            address[] memory unstakeAddresses,
            UnstakeRequest[] memory unstakeRequests
        )
    {
        address[] memory addresses = minerAddresses.values;

        uint countUnstakeRequest = 0;
        for (uint i = 0; i < addresses.length; ++i) {
            UnstakeRequest memory request = minerUnstakeRequests[addresses[i]];
            if (request.unlockAt > 0) ++countUnstakeRequest;
        }

        unstakeAddresses = new address[](countUnstakeRequest);
        unstakeRequests = new UnstakeRequest[](countUnstakeRequest);
        uint idx = 0;
        for (uint i = 0; i < addresses.length; ++i) {
            UnstakeRequest memory request = minerUnstakeRequests[addresses[i]];
            if (request.unlockAt > 0) {
                unstakeAddresses[idx] = addresses[idx];
                unstakeRequests[idx] = request;
                ++idx;
            }
        }
    }
}
