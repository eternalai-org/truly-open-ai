// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IStakingHub {
    struct MinerEpochState {
        uint256 perfReward;
        uint256 epochReward;
        uint256 totalTaskCompleted;
        uint256 totalMiner;
    }

    struct Model {
        uint256 minimumFee;
        uint32 tier;
    }

    struct Worker {
        uint256 stake;
        uint256 commitment;
        address modelAddress;
        uint40 lastClaimedEpoch;
        uint40 activeTime;
        uint16 tier;
    }

    struct UnstakeRequest {
        uint256 stake;
        uint40 unlockAt;
    }

    struct Boost {
        uint40 minerTimestamp;
        uint40 validatorTimestamp;
        uint48 reserved1; // accumulated active time
        uint128 reserved2;
    }

    event MinerRegistration(
        address indexed miner,
        uint16 indexed tier,
        uint256 value
    );
    event MinerJoin(address indexed miner);
    event MinerUnregistration(address indexed miner);
    event MinerExtraStake(address indexed miner, uint256 value);
    event MinerUnstake(address indexed miner, uint256 stake);
    event Restake(
        address indexed miner,
        uint256 restake,
        address indexed model
    );
    event ModelMinimumFeeUpdate(address indexed model, uint256 minimumFee);
    event ModelTierUpdate(address indexed model, uint32 tier);
    event ModelUnregistration(address indexed model);
    event ModelRegistration(
        address indexed model,
        uint16 indexed tier,
        uint256 minimumFee
    );
    event FraudulentMinerPenalized(
        address indexed miner,
        address indexed modelAddress,
        address indexed treasury,
        uint256 fine
    );
    event MinerDeactivated(
        address indexed miner,
        address indexed modelAddress,
        uint40 activeTime
    );
    event RewardClaim(address indexed worker, uint256 value);
    event FinePercentageUpdated(uint16 oldPercent, uint16 newPercent);
    event PenaltyDurationUpdated(uint40 oldDuration, uint40 newDuration);
    event MinFeeToUseUpdated(uint256 oldValue, uint256 newValue);
    event RewardPerEpoch(uint256 oldReward, uint256 newReward);
    event BlocksPerEpoch(uint256 oldBlocks, uint256 newBlocks);
    event UnstakeDelayTime(uint256 oldDelayTime, uint256 newDelayTime);

    error InvalidMiner();
    error InvalidModel();
    error FeeTooLow();
    error InvalidTier();
    error AlreadyRegistered();
    error NotRegistered();
    error SameModelAddress();
    error StakeTooLow();
    error MinerInDeactivationTime();
    error StillBeingLocked();
    error NullStake();
    error ZeroValue();
    error InvalidBlockValue();
    error InvalidValue();
    error InvalidAddress();
    error InvalidWorkerHub();

    function updateEpoch() external;
    function getModelInfo(address _modelAddr) external returns (Model memory);
    function getMinerAddressesOfModel(
        address _model
    ) external view returns (address[] memory);
    function isMinerAddress(address _miner) external view returns (bool);
    function validateModelOfMiner(address _miner) external view;
    function slashMiner(address _miner, bool _isFined) external;
    function getMinFeeToUse(
        address _modelAddress
    ) external view returns (uint256);
}
