// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IGPUManager {
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
        uint32 modelId;
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
        uint48 reserved1; // accumulated active time // mr @issac review and change name
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
        uint32 indexed modelId,
        uint256 restake
    );
    event ModelMinimumFeeUpdate(uint32 indexed modelId, uint256 minimumFee);
    event ModelTierUpdate(uint32 indexed modelId, uint32 tier);
    event ModelUnregistration(uint32 indexed modelId);
    event ModelRegistration(
        uint32 indexed modelId,
        uint16 indexed tier,
        uint256 minimumFee
    );
    event FraudulentMinerPenalized(
        address indexed miner,
        uint32 indexed modelId,
        address indexed treasury,
        uint256 fine
    );
    event MinerDeactivated(
        address indexed miner,
        uint32 indexed modelId,
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
    error ZeroValue();
    error InvalidBlockValue();
    error InvalidValue();
    error InvalidAddress();
    error NotEnoughMiners();

    function updateEpoch() external;
    function getModelInfo(uint32 modelId) external returns (Model memory);
    function getMinerAddressesOfModel(
        uint32 modelId
    ) external view returns (address[] memory);
    function getMinFeeToUse(uint32 modelId) external view returns (uint256);
    function isActiveModel(uint32 modelId) external view returns (bool);
    function validateModelAndChooseRandomMiner(
        uint32 modelId,
        uint256 minersRequired
    ) external returns (address assignMiner, uint256 modelFee);
    function validateMiner(address miner) external;
    function slashMiner(address miner, bool isFined) external;
}
