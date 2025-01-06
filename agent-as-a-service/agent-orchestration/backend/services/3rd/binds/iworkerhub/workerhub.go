// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iworkerhub

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IWorkerHubAssignment is an auto generated low-level Go binding around an user-defined struct.
type IWorkerHubAssignment struct {
	InferenceId *big.Int
	Commitment  [32]byte
	Digest      [32]byte
	RevealNonce *big.Int
	Worker      common.Address
	Role        uint8
	Vote        uint8
	Output      []byte
}

// IWorkerHubAssignmentInfo is an auto generated low-level Go binding around an user-defined struct.
type IWorkerHubAssignmentInfo struct {
	AssignmentId  *big.Int
	InferenceId   *big.Int
	Value         *big.Int
	Input         []byte
	ModelAddress  common.Address
	Creator       common.Address
	SubmitTimeout *big.Int
	CommitTimeout *big.Int
	RevealTimeout *big.Int
}

// IWorkerHubDAOTokenPercentage is an auto generated low-level Go binding around an user-defined struct.
type IWorkerHubDAOTokenPercentage struct {
	MinerPercentage    uint16
	UserPercentage     uint16
	ReferrerPercentage uint16
	RefereePercentage  uint16
	L2OwnerPercentage  uint16
}

// IWorkerHubDAOTokenReceiverInfor is an auto generated low-level Go binding around an user-defined struct.
type IWorkerHubDAOTokenReceiverInfor struct {
	Receiver common.Address
	Amount   *big.Int
	Role     uint8
}

// IWorkerHubInference is an auto generated low-level Go binding around an user-defined struct.
type IWorkerHubInference struct {
	Assignments    []*big.Int
	Input          []byte
	Value          *big.Int
	FeeL2          *big.Int
	FeeTreasury    *big.Int
	ModelAddress   common.Address
	SubmitTimeout  *big.Int
	CommitTimeout  *big.Int
	RevealTimeout  *big.Int
	Status         uint8
	Creator        common.Address
	ProcessedMiner common.Address
	Referrer       common.Address
}

// IWorkerHubUnstakeRequest is an auto generated low-level Go binding around an user-defined struct.
type IWorkerHubUnstakeRequest struct {
	Stake    *big.Int
	UnlockAt *big.Int
}

// IWorkerHubWorker is an auto generated low-level Go binding around an user-defined struct.
type IWorkerHubWorker struct {
	Stake            *big.Int
	Commitment       *big.Int
	ModelAddress     common.Address
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}

// IWorkerHubWorkerInfo is an auto generated low-level Go binding around an user-defined struct.
type IWorkerHubWorkerInfo struct {
	WorkerAddress    common.Address
	Stake            *big.Int
	Commitment       *big.Int
	ModelAddress     common.Address
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}

// IWorkerHubMetaData contains all meta data concerning the IWorkerHub contract.
var IWorkerHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"AddressSet_DuplicatedValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"AddressSet_ValueNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyRevealed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadySubmitted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"Bytes32Set_DuplicatedValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommitTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InferMustBeSolvingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlockValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCommitment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInferenceStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMiner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidModel\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReveal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinerInDeactivationTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MiningSessionEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MiningSessionNotEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMiners\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullStake\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RevealTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StillBeingLocked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Uint256Set_DuplicatedValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatingSessionNotEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroValue\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBlocks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBlocks\",\"type\":\"uint256\"}],\"name\":\"BlocksPerEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"CommitDuration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assigmentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"CommitmentSubmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumIWorkerHub.DAOTokenReceiverRole\",\"name\":\"role\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structIWorkerHub.DAOTokenReceiverInfor[]\",\"name\":\"receivers\",\"type\":\"tuple[]\"}],\"name\":\"DAOTokenMintedV2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"minerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"userPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"referrerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"refereePercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"l2OwnerPercentage\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structIWorkerHub.DAOTokenPercentage\",\"name\":\"oldValue\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"minerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"userPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"referrerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"refereePercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"l2OwnerPercentage\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structIWorkerHub.DAOTokenPercentage\",\"name\":\"newValue\",\"type\":\"tuple\"}],\"name\":\"DAOTokenPercentageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"DAOTokenRewardUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"DAOTokenUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"oldPercent\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newPercent\",\"type\":\"uint16\"}],\"name\":\"FinePercentageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fine\",\"type\":\"uint256\"}],\"name\":\"FraudulentMinerPenalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIWorkerHub.InferenceStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"InferenceStatusUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"L2OwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MinFeeToUseUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"activeTime\",\"type\":\"uint40\"}],\"name\":\"MinerDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MinerExtraStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"MinerJoin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MinerRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assignmentId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"MinerRoleSeized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"MinerUnregistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"}],\"name\":\"MinerUnstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"newValue\",\"type\":\"uint40\"}],\"name\":\"MiningTimeLimitUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"}],\"name\":\"ModelMinimumFeeUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"}],\"name\":\"ModelRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"tier\",\"type\":\"uint32\"}],\"name\":\"ModelTierUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"}],\"name\":\"ModelUnregistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assignmentId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"expiredAt\",\"type\":\"uint40\"}],\"name\":\"NewAssignment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originInferenceId\",\"type\":\"uint256\"}],\"name\":\"NewInference\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originInferenceId\",\"type\":\"uint256\"}],\"name\":\"NewScoringInference\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"oldDuration\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"newDuration\",\"type\":\"uint40\"}],\"name\":\"PenaltyDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"restake\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"model\",\"type\":\"address\"}],\"name\":\"Restake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"RevealDuration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assigmentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"nonce\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"}],\"name\":\"RevealSubmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"RewardClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReward\",\"type\":\"uint256\"}],\"name\":\"RewardPerEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assigmentId\",\"type\":\"uint256\"}],\"name\":\"SolutionSubmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assignmentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"StreamedData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"SubmitDuration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TopUpInfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"treasuryFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"L2OwnerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"L2OwnerFee\",\"type\":\"uint256\"}],\"name\":\"TransferFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"TreasuryAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldDelayTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDelayTime\",\"type\":\"uint256\"}],\"name\":\"UnstakeDelayTime\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"assignmentNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assignments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint40\",\"name\":\"revealNonce\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"enumIWorkerHub.AssignmentRole\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"enumIWorkerHub.Vote\",\"name\":\"vote\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blocksPerEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assignId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_commitment\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"commitDuration\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"inferId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"daoReceiversInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumIWorkerHub.DAOTokenReceiverRole\",\"name\":\"role\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoTokenPercentage\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"minerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"userPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"referrerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"refereePercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"l2OwnerPercentage\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoTokenReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeL2Percentage\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRatioMinerValidator\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTreasuryPercentage\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finePercentage\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_modelAddress\",\"type\":\"address\"}],\"name\":\"forceChangeModelForMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"getAllAssignments\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint40\",\"name\":\"revealNonce\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"enumIWorkerHub.AssignmentRole\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"enumIWorkerHub.Vote\",\"name\":\"vote\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"}],\"internalType\":\"structIWorkerHub.Assignment[]\",\"name\":\"assignmentData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"getAllInferences\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"assignments\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeL2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTreasury\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"submitTimeout\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"commitTimeout\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"revealTimeout\",\"type\":\"uint40\"},{\"internalType\":\"enumIWorkerHub.InferenceStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"processedMiner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"internalType\":\"structIWorkerHub.Inference[]\",\"name\":\"inferenceData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllMinerUnstakeRequests\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"unstakeAddresses\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"unlockAt\",\"type\":\"uint40\"}],\"internalType\":\"structIWorkerHub.UnstakeRequest[]\",\"name\":\"unstakeRequests\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllMiners\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"lastClaimedEpoch\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"activeTime\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"internalType\":\"structIWorkerHub.Worker[]\",\"name\":\"minerData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_inferId\",\"type\":\"uint256\"}],\"name\":\"getAssignmentByInferenceId\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"uint40\",\"name\":\"revealNonce\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"enumIWorkerHub.AssignmentRole\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"enumIWorkerHub.Vote\",\"name\":\"vote\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"}],\"internalType\":\"structIWorkerHub.Assignment[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minerAddr\",\"type\":\"address\"}],\"name\":\"getAssignmentByMiner\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"assignmentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"submitTimeout\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"commitTimeout\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"revealTimeout\",\"type\":\"uint40\"}],\"internalType\":\"structIWorkerHub.AssignmentInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_inferenceId\",\"type\":\"uint256\"}],\"name\":\"getInferenceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"assignments\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeL2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTreasury\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"submitTimeout\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"commitTimeout\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"revealTimeout\",\"type\":\"uint40\"},{\"internalType\":\"enumIWorkerHub.InferenceStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"processedMiner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"internalType\":\"structIWorkerHub.Inference\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_modelAddress\",\"type\":\"address\"}],\"name\":\"getMinFeeToUse\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinerAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"}],\"name\":\"getMinerAddressesOfModel\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMiners\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"workerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"lastClaimedEpoch\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"activeTime\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"internalType\":\"structIWorkerHub.WorkerInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_inferenceId\",\"type\":\"uint256\"}],\"name\":\"getMintingAssignmentsOfInference\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"assignmentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"submitTimeout\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"commitTimeout\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"revealTimeout\",\"type\":\"uint40\"}],\"internalType\":\"structIWorkerHub.AssignmentInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getModelAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNOMiner\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assignmentId\",\"type\":\"uint256\"}],\"name\":\"getRoleByAssigmentId\",\"outputs\":[{\"internalType\":\"enumIWorkerHub.AssignmentRole\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wEAIAmt\",\"type\":\"uint256\"}],\"name\":\"increaseMinerStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"}],\"name\":\"infer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"originInferId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callback\",\"type\":\"address\"}],\"name\":\"inferWithCallback\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inferenceNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_daoToken\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_feeL2Percentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_feeTreasuryPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_minerMinimumStake\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_minerRequirement\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_blocksPerEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardPerEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_finePercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_feeRatioMinerValidor\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_minFeeToUse\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_daoTokenReward\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"minerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"userPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"referrerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"refereePercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"l2OwnerPercentage\",\"type\":\"uint16\"}],\"internalType\":\"structIWorkerHub.DAOTokenPercentage\",\"name\":\"_daoTokenPercentage\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assignmentId\",\"type\":\"uint256\"}],\"name\":\"isAssignmentPending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isReferrer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"joinForMinting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maximumTier\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minFeeToUse\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minerMinimumStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minerRequirement\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minerUnstakeRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"unlockAt\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"miners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitment\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"modelAddress\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"lastClaimedEpoch\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"activeTime\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"modelScoring\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"models\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumFee\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"tier\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"multiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"penaltyDuration\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"wEAIAmt\",\"type\":\"uint256\"}],\"name\":\"registerMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_tier\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_minimumFee\",\"type\":\"uint256\"}],\"name\":\"registerModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_referrer\",\"type\":\"address\"}],\"name\":\"registerReferrer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_referrers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_referees\",\"type\":\"address[]\"}],\"name\":\"registerReferrer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_inferenceId\",\"type\":\"uint256\"}],\"name\":\"resolveInference\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"tier\",\"type\":\"uint16\"}],\"name\":\"restakeForMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"resultReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_originInferId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_result\",\"type\":\"bytes\"}],\"name\":\"resultReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assignId\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"_nonce\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"reveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revealDuration\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardInEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"perfReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTaskCompleted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalMiner\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"rewardToClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assignmentId\",\"type\":\"uint256\"}],\"name\":\"seizeMinerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blocks\",\"type\":\"uint256\"}],\"name\":\"setBlocksPerEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_newCommitDuration\",\"type\":\"uint40\"}],\"name\":\"setCommitDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_daoTokenAddress\",\"type\":\"address\"}],\"name\":\"setDAOToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"minerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"userPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"referrerPercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"refereePercentage\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"l2OwnerPercentage\",\"type\":\"uint16\"}],\"internalType\":\"structIWorkerHub.DAOTokenPercentage\",\"name\":\"_daoTokenPercentage\",\"type\":\"tuple\"}],\"name\":\"setDAOTokenPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newDAOTokenReward\",\"type\":\"uint256\"}],\"name\":\"setDAOTokenReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_newRatio\",\"type\":\"uint16\"}],\"name\":\"setFeeRatioMinerValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_finePercentage\",\"type\":\"uint16\"}],\"name\":\"setFinePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2OwnerAddress\",\"type\":\"address\"}],\"name\":\"setL2Owner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minFeeToUse\",\"type\":\"uint256\"}],\"name\":\"setMinFeeToUse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minerMinimumStake\",\"type\":\"uint256\"}],\"name\":\"setMinerMinimumStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newRewardAmount\",\"type\":\"uint256\"}],\"name\":\"setNewRewardInEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_penaltyDuration\",\"type\":\"uint40\"}],\"name\":\"setPenaltyDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_newRevealDuration\",\"type\":\"uint40\"}],\"name\":\"setRevealDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_workerHubScoring\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_modelScoring\",\"type\":\"address\"}],\"name\":\"setScoringInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_newSubmitDuration\",\"type\":\"uint40\"}],\"name\":\"setSubmitDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasuryAddress\",\"type\":\"address\"}],\"name\":\"setTreasuryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_newUnstakeDelayTime\",\"type\":\"uint40\"}],\"name\":\"setUnstakDelayTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_miner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isFined\",\"type\":\"bool\"}],\"name\":\"slashMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assignmentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"streamData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitDuration\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assigmentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"submitSolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_inferenceId\",\"type\":\"uint256\"}],\"name\":\"topUpInfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unregisterMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"}],\"name\":\"unregisterModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstakeDelayTime\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstakeForMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minimumFee\",\"type\":\"uint256\"}],\"name\":\"updateModelMinimumFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_model\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_tier\",\"type\":\"uint32\"}],\"name\":\"updateModelTier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isReferred\",\"type\":\"bool\"}],\"name\":\"validateDAOSupplyIncrease\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"notReachedLimit\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"inferId\",\"type\":\"uint256\"}],\"name\":\"votingInfo\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"totalCommit\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"totalReveal\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wEAI\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"workerHubScoring\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// IWorkerHubABI is the input ABI used to generate the binding from.
// Deprecated: Use IWorkerHubMetaData.ABI instead.
var IWorkerHubABI = IWorkerHubMetaData.ABI

// IWorkerHub is an auto generated Go binding around an Ethereum contract.
type IWorkerHub struct {
	IWorkerHubCaller     // Read-only binding to the contract
	IWorkerHubTransactor // Write-only binding to the contract
	IWorkerHubFilterer   // Log filterer for contract events
}

// IWorkerHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWorkerHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWorkerHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWorkerHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWorkerHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWorkerHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWorkerHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWorkerHubSession struct {
	Contract     *IWorkerHub       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWorkerHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWorkerHubCallerSession struct {
	Contract *IWorkerHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IWorkerHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWorkerHubTransactorSession struct {
	Contract     *IWorkerHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IWorkerHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWorkerHubRaw struct {
	Contract *IWorkerHub // Generic contract binding to access the raw methods on
}

// IWorkerHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWorkerHubCallerRaw struct {
	Contract *IWorkerHubCaller // Generic read-only contract binding to access the raw methods on
}

// IWorkerHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWorkerHubTransactorRaw struct {
	Contract *IWorkerHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWorkerHub creates a new instance of IWorkerHub, bound to a specific deployed contract.
func NewIWorkerHub(address common.Address, backend bind.ContractBackend) (*IWorkerHub, error) {
	contract, err := bindIWorkerHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWorkerHub{IWorkerHubCaller: IWorkerHubCaller{contract: contract}, IWorkerHubTransactor: IWorkerHubTransactor{contract: contract}, IWorkerHubFilterer: IWorkerHubFilterer{contract: contract}}, nil
}

// NewIWorkerHubCaller creates a new read-only instance of IWorkerHub, bound to a specific deployed contract.
func NewIWorkerHubCaller(address common.Address, caller bind.ContractCaller) (*IWorkerHubCaller, error) {
	contract, err := bindIWorkerHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubCaller{contract: contract}, nil
}

// NewIWorkerHubTransactor creates a new write-only instance of IWorkerHub, bound to a specific deployed contract.
func NewIWorkerHubTransactor(address common.Address, transactor bind.ContractTransactor) (*IWorkerHubTransactor, error) {
	contract, err := bindIWorkerHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubTransactor{contract: contract}, nil
}

// NewIWorkerHubFilterer creates a new log filterer instance of IWorkerHub, bound to a specific deployed contract.
func NewIWorkerHubFilterer(address common.Address, filterer bind.ContractFilterer) (*IWorkerHubFilterer, error) {
	contract, err := bindIWorkerHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubFilterer{contract: contract}, nil
}

// bindIWorkerHub binds a generic wrapper to an already deployed contract.
func bindIWorkerHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IWorkerHubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWorkerHub *IWorkerHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWorkerHub.Contract.IWorkerHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWorkerHub *IWorkerHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWorkerHub.Contract.IWorkerHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWorkerHub *IWorkerHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWorkerHub.Contract.IWorkerHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWorkerHub *IWorkerHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWorkerHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWorkerHub *IWorkerHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWorkerHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWorkerHub *IWorkerHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWorkerHub.Contract.contract.Transact(opts, method, params...)
}

// AssignmentNumber is a free data retrieval call binding the contract method 0x6973d3f2.
//
// Solidity: function assignmentNumber() view returns(uint256)
func (_IWorkerHub *IWorkerHubCaller) AssignmentNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "assignmentNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssignmentNumber is a free data retrieval call binding the contract method 0x6973d3f2.
//
// Solidity: function assignmentNumber() view returns(uint256)
func (_IWorkerHub *IWorkerHubSession) AssignmentNumber() (*big.Int, error) {
	return _IWorkerHub.Contract.AssignmentNumber(&_IWorkerHub.CallOpts)
}

// AssignmentNumber is a free data retrieval call binding the contract method 0x6973d3f2.
//
// Solidity: function assignmentNumber() view returns(uint256)
func (_IWorkerHub *IWorkerHubCallerSession) AssignmentNumber() (*big.Int, error) {
	return _IWorkerHub.Contract.AssignmentNumber(&_IWorkerHub.CallOpts)
}

// Assignments is a free data retrieval call binding the contract method 0x4e50c75c.
//
// Solidity: function assignments(uint256 ) view returns(uint256 inferenceId, bytes32 commitment, bytes32 digest, uint40 revealNonce, address worker, uint8 role, uint8 vote, bytes output)
func (_IWorkerHub *IWorkerHubCaller) Assignments(opts *bind.CallOpts, arg0 *big.Int) (struct {
	InferenceId *big.Int
	Commitment  [32]byte
	Digest      [32]byte
	RevealNonce *big.Int
	Worker      common.Address
	Role        uint8
	Vote        uint8
	Output      []byte
}, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "assignments", arg0)

	outstruct := new(struct {
		InferenceId *big.Int
		Commitment  [32]byte
		Digest      [32]byte
		RevealNonce *big.Int
		Worker      common.Address
		Role        uint8
		Vote        uint8
		Output      []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InferenceId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Commitment = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Digest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.RevealNonce = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Worker = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Role = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.Vote = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.Output = *abi.ConvertType(out[7], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Assignments is a free data retrieval call binding the contract method 0x4e50c75c.
//
// Solidity: function assignments(uint256 ) view returns(uint256 inferenceId, bytes32 commitment, bytes32 digest, uint40 revealNonce, address worker, uint8 role, uint8 vote, bytes output)
func (_IWorkerHub *IWorkerHubSession) Assignments(arg0 *big.Int) (struct {
	InferenceId *big.Int
	Commitment  [32]byte
	Digest      [32]byte
	RevealNonce *big.Int
	Worker      common.Address
	Role        uint8
	Vote        uint8
	Output      []byte
}, error) {
	return _IWorkerHub.Contract.Assignments(&_IWorkerHub.CallOpts, arg0)
}

// Assignments is a free data retrieval call binding the contract method 0x4e50c75c.
//
// Solidity: function assignments(uint256 ) view returns(uint256 inferenceId, bytes32 commitment, bytes32 digest, uint40 revealNonce, address worker, uint8 role, uint8 vote, bytes output)
func (_IWorkerHub *IWorkerHubCallerSession) Assignments(arg0 *big.Int) (struct {
	InferenceId *big.Int
	Commitment  [32]byte
	Digest      [32]byte
	RevealNonce *big.Int
	Worker      common.Address
	Role        uint8
	Vote        uint8
	Output      []byte
}, error) {
	return _IWorkerHub.Contract.Assignments(&_IWorkerHub.CallOpts, arg0)
}

// BlocksPerEpoch is a free data retrieval call binding the contract method 0xf0682054.
//
// Solidity: function blocksPerEpoch() view returns(uint256)
func (_IWorkerHub *IWorkerHubCaller) BlocksPerEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "blocksPerEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlocksPerEpoch is a free data retrieval call binding the contract method 0xf0682054.
//
// Solidity: function blocksPerEpoch() view returns(uint256)
func (_IWorkerHub *IWorkerHubSession) BlocksPerEpoch() (*big.Int, error) {
	return _IWorkerHub.Contract.BlocksPerEpoch(&_IWorkerHub.CallOpts)
}

// BlocksPerEpoch is a free data retrieval call binding the contract method 0xf0682054.
//
// Solidity: function blocksPerEpoch() view returns(uint256)
func (_IWorkerHub *IWorkerHubCallerSession) BlocksPerEpoch() (*big.Int, error) {
	return _IWorkerHub.Contract.BlocksPerEpoch(&_IWorkerHub.CallOpts)
}

// CommitDuration is a free data retrieval call binding the contract method 0x6f833811.
//
// Solidity: function commitDuration() view returns(uint40)
func (_IWorkerHub *IWorkerHubCaller) CommitDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "commitDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommitDuration is a free data retrieval call binding the contract method 0x6f833811.
//
// Solidity: function commitDuration() view returns(uint40)
func (_IWorkerHub *IWorkerHubSession) CommitDuration() (*big.Int, error) {
	return _IWorkerHub.Contract.CommitDuration(&_IWorkerHub.CallOpts)
}

// CommitDuration is a free data retrieval call binding the contract method 0x6f833811.
//
// Solidity: function commitDuration() view returns(uint40)
func (_IWorkerHub *IWorkerHubCallerSession) CommitDuration() (*big.Int, error) {
	return _IWorkerHub.Contract.CommitDuration(&_IWorkerHub.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint40)
func (_IWorkerHub *IWorkerHubCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint40)
func (_IWorkerHub *IWorkerHubSession) CurrentEpoch() (*big.Int, error) {
	return _IWorkerHub.Contract.CurrentEpoch(&_IWorkerHub.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint40)
func (_IWorkerHub *IWorkerHubCallerSession) CurrentEpoch() (*big.Int, error) {
	return _IWorkerHub.Contract.CurrentEpoch(&_IWorkerHub.CallOpts)
}

// DaoReceiversInfo is a free data retrieval call binding the contract method 0x61d52cf7.
//
// Solidity: function daoReceiversInfo(uint256 inferId, uint256 ) view returns(address receiver, uint256 amount, uint8 role)
func (_IWorkerHub *IWorkerHubCaller) DaoReceiversInfo(opts *bind.CallOpts, inferId *big.Int, arg1 *big.Int) (struct {
	Receiver common.Address
	Amount   *big.Int
	Role     uint8
}, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "daoReceiversInfo", inferId, arg1)

	outstruct := new(struct {
		Receiver common.Address
		Amount   *big.Int
		Role     uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Role = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// DaoReceiversInfo is a free data retrieval call binding the contract method 0x61d52cf7.
//
// Solidity: function daoReceiversInfo(uint256 inferId, uint256 ) view returns(address receiver, uint256 amount, uint8 role)
func (_IWorkerHub *IWorkerHubSession) DaoReceiversInfo(inferId *big.Int, arg1 *big.Int) (struct {
	Receiver common.Address
	Amount   *big.Int
	Role     uint8
}, error) {
	return _IWorkerHub.Contract.DaoReceiversInfo(&_IWorkerHub.CallOpts, inferId, arg1)
}

// DaoReceiversInfo is a free data retrieval call binding the contract method 0x61d52cf7.
//
// Solidity: function daoReceiversInfo(uint256 inferId, uint256 ) view returns(address receiver, uint256 amount, uint8 role)
func (_IWorkerHub *IWorkerHubCallerSession) DaoReceiversInfo(inferId *big.Int, arg1 *big.Int) (struct {
	Receiver common.Address
	Amount   *big.Int
	Role     uint8
}, error) {
	return _IWorkerHub.Contract.DaoReceiversInfo(&_IWorkerHub.CallOpts, inferId, arg1)
}

// DaoToken is a free data retrieval call binding the contract method 0x4914b030.
//
// Solidity: function daoToken() view returns(address)
func (_IWorkerHub *IWorkerHubCaller) DaoToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "daoToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaoToken is a free data retrieval call binding the contract method 0x4914b030.
//
// Solidity: function daoToken() view returns(address)
func (_IWorkerHub *IWorkerHubSession) DaoToken() (common.Address, error) {
	return _IWorkerHub.Contract.DaoToken(&_IWorkerHub.CallOpts)
}

// DaoToken is a free data retrieval call binding the contract method 0x4914b030.
//
// Solidity: function daoToken() view returns(address)
func (_IWorkerHub *IWorkerHubCallerSession) DaoToken() (common.Address, error) {
	return _IWorkerHub.Contract.DaoToken(&_IWorkerHub.CallOpts)
}

// DaoTokenPercentage is a free data retrieval call binding the contract method 0xff5db406.
//
// Solidity: function daoTokenPercentage() view returns(uint16 minerPercentage, uint16 userPercentage, uint16 referrerPercentage, uint16 refereePercentage, uint16 l2OwnerPercentage)
func (_IWorkerHub *IWorkerHubCaller) DaoTokenPercentage(opts *bind.CallOpts) (struct {
	MinerPercentage    uint16
	UserPercentage     uint16
	ReferrerPercentage uint16
	RefereePercentage  uint16
	L2OwnerPercentage  uint16
}, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "daoTokenPercentage")

	outstruct := new(struct {
		MinerPercentage    uint16
		UserPercentage     uint16
		ReferrerPercentage uint16
		RefereePercentage  uint16
		L2OwnerPercentage  uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinerPercentage = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.UserPercentage = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.ReferrerPercentage = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.RefereePercentage = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.L2OwnerPercentage = *abi.ConvertType(out[4], new(uint16)).(*uint16)

	return *outstruct, err

}

// DaoTokenPercentage is a free data retrieval call binding the contract method 0xff5db406.
//
// Solidity: function daoTokenPercentage() view returns(uint16 minerPercentage, uint16 userPercentage, uint16 referrerPercentage, uint16 refereePercentage, uint16 l2OwnerPercentage)
func (_IWorkerHub *IWorkerHubSession) DaoTokenPercentage() (struct {
	MinerPercentage    uint16
	UserPercentage     uint16
	ReferrerPercentage uint16
	RefereePercentage  uint16
	L2OwnerPercentage  uint16
}, error) {
	return _IWorkerHub.Contract.DaoTokenPercentage(&_IWorkerHub.CallOpts)
}

// DaoTokenPercentage is a free data retrieval call binding the contract method 0xff5db406.
//
// Solidity: function daoTokenPercentage() view returns(uint16 minerPercentage, uint16 userPercentage, uint16 referrerPercentage, uint16 refereePercentage, uint16 l2OwnerPercentage)
func (_IWorkerHub *IWorkerHubCallerSession) DaoTokenPercentage() (struct {
	MinerPercentage    uint16
	UserPercentage     uint16
	ReferrerPercentage uint16
	RefereePercentage  uint16
	L2OwnerPercentage  uint16
}, error) {
	return _IWorkerHub.Contract.DaoTokenPercentage(&_IWorkerHub.CallOpts)
}

// DaoTokenReward is a free data retrieval call binding the contract method 0x0940c392.
//
// Solidity: function daoTokenReward() view returns(uint256)
func (_IWorkerHub *IWorkerHubCaller) DaoTokenReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "daoTokenReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DaoTokenReward is a free data retrieval call binding the contract method 0x0940c392.
//
// Solidity: function daoTokenReward() view returns(uint256)
func (_IWorkerHub *IWorkerHubSession) DaoTokenReward() (*big.Int, error) {
	return _IWorkerHub.Contract.DaoTokenReward(&_IWorkerHub.CallOpts)
}

// DaoTokenReward is a free data retrieval call binding the contract method 0x0940c392.
//
// Solidity: function daoTokenReward() view returns(uint256)
func (_IWorkerHub *IWorkerHubCallerSession) DaoTokenReward() (*big.Int, error) {
	return _IWorkerHub.Contract.DaoTokenReward(&_IWorkerHub.CallOpts)
}

// FeeL2Percentage is a free data retrieval call binding the contract method 0x39d2e296.
//
// Solidity: function feeL2Percentage() view returns(uint16)
func (_IWorkerHub *IWorkerHubCaller) FeeL2Percentage(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "feeL2Percentage")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// FeeL2Percentage is a free data retrieval call binding the contract method 0x39d2e296.
//
// Solidity: function feeL2Percentage() view returns(uint16)
func (_IWorkerHub *IWorkerHubSession) FeeL2Percentage() (uint16, error) {
	return _IWorkerHub.Contract.FeeL2Percentage(&_IWorkerHub.CallOpts)
}

// FeeL2Percentage is a free data retrieval call binding the contract method 0x39d2e296.
//
// Solidity: function feeL2Percentage() view returns(uint16)
func (_IWorkerHub *IWorkerHubCallerSession) FeeL2Percentage() (uint16, error) {
	return _IWorkerHub.Contract.FeeL2Percentage(&_IWorkerHub.CallOpts)
}

// FeeRatioMinerValidator is a free data retrieval call binding the contract method 0x50eac7c8.
//
// Solidity: function feeRatioMinerValidator() view returns(uint16)
func (_IWorkerHub *IWorkerHubCaller) FeeRatioMinerValidator(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "feeRatioMinerValidator")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// FeeRatioMinerValidator is a free data retrieval call binding the contract method 0x50eac7c8.
//
// Solidity: function feeRatioMinerValidator() view returns(uint16)
func (_IWorkerHub *IWorkerHubSession) FeeRatioMinerValidator() (uint16, error) {
	return _IWorkerHub.Contract.FeeRatioMinerValidator(&_IWorkerHub.CallOpts)
}

// FeeRatioMinerValidator is a free data retrieval call binding the contract method 0x50eac7c8.
//
// Solidity: function feeRatioMinerValidator() view returns(uint16)
func (_IWorkerHub *IWorkerHubCallerSession) FeeRatioMinerValidator() (uint16, error) {
	return _IWorkerHub.Contract.FeeRatioMinerValidator(&_IWorkerHub.CallOpts)
}

// FeeTreasuryPercentage is a free data retrieval call binding the contract method 0x09c83b4f.
//
// Solidity: function feeTreasuryPercentage() view returns(uint16)
func (_IWorkerHub *IWorkerHubCaller) FeeTreasuryPercentage(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "feeTreasuryPercentage")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// FeeTreasuryPercentage is a free data retrieval call binding the contract method 0x09c83b4f.
//
// Solidity: function feeTreasuryPercentage() view returns(uint16)
func (_IWorkerHub *IWorkerHubSession) FeeTreasuryPercentage() (uint16, error) {
	return _IWorkerHub.Contract.FeeTreasuryPercentage(&_IWorkerHub.CallOpts)
}

// FeeTreasuryPercentage is a free data retrieval call binding the contract method 0x09c83b4f.
//
// Solidity: function feeTreasuryPercentage() view returns(uint16)
func (_IWorkerHub *IWorkerHubCallerSession) FeeTreasuryPercentage() (uint16, error) {
	return _IWorkerHub.Contract.FeeTreasuryPercentage(&_IWorkerHub.CallOpts)
}

// FinePercentage is a free data retrieval call binding the contract method 0x74172795.
//
// Solidity: function finePercentage() view returns(uint16)
func (_IWorkerHub *IWorkerHubCaller) FinePercentage(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "finePercentage")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// FinePercentage is a free data retrieval call binding the contract method 0x74172795.
//
// Solidity: function finePercentage() view returns(uint16)
func (_IWorkerHub *IWorkerHubSession) FinePercentage() (uint16, error) {
	return _IWorkerHub.Contract.FinePercentage(&_IWorkerHub.CallOpts)
}

// FinePercentage is a free data retrieval call binding the contract method 0x74172795.
//
// Solidity: function finePercentage() view returns(uint16)
func (_IWorkerHub *IWorkerHubCallerSession) FinePercentage() (uint16, error) {
	return _IWorkerHub.Contract.FinePercentage(&_IWorkerHub.CallOpts)
}

// GetAllAssignments is a free data retrieval call binding the contract method 0x16d0a88f.
//
// Solidity: function getAllAssignments(uint256 startId, uint256 count) view returns((uint256,bytes32,bytes32,uint40,address,uint8,uint8,bytes)[] assignmentData)
func (_IWorkerHub *IWorkerHubCaller) GetAllAssignments(opts *bind.CallOpts, startId *big.Int, count *big.Int) ([]IWorkerHubAssignment, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "getAllAssignments", startId, count)

	if err != nil {
		return *new([]IWorkerHubAssignment), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWorkerHubAssignment)).(*[]IWorkerHubAssignment)

	return out0, err

}

// GetAllAssignments is a free data retrieval call binding the contract method 0x16d0a88f.
//
// Solidity: function getAllAssignments(uint256 startId, uint256 count) view returns((uint256,bytes32,bytes32,uint40,address,uint8,uint8,bytes)[] assignmentData)
func (_IWorkerHub *IWorkerHubSession) GetAllAssignments(startId *big.Int, count *big.Int) ([]IWorkerHubAssignment, error) {
	return _IWorkerHub.Contract.GetAllAssignments(&_IWorkerHub.CallOpts, startId, count)
}

// GetAllAssignments is a free data retrieval call binding the contract method 0x16d0a88f.
//
// Solidity: function getAllAssignments(uint256 startId, uint256 count) view returns((uint256,bytes32,bytes32,uint40,address,uint8,uint8,bytes)[] assignmentData)
func (_IWorkerHub *IWorkerHubCallerSession) GetAllAssignments(startId *big.Int, count *big.Int) ([]IWorkerHubAssignment, error) {
	return _IWorkerHub.Contract.GetAllAssignments(&_IWorkerHub.CallOpts, startId, count)
}

// GetAllInferences is a free data retrieval call binding the contract method 0xf1ea45e3.
//
// Solidity: function getAllInferences(uint256 startId, uint256 count) view returns((uint256[],bytes,uint256,uint256,uint256,address,uint40,uint40,uint40,uint8,address,address,address)[] inferenceData)
func (_IWorkerHub *IWorkerHubCaller) GetAllInferences(opts *bind.CallOpts, startId *big.Int, count *big.Int) ([]IWorkerHubInference, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "getAllInferences", startId, count)

	if err != nil {
		return *new([]IWorkerHubInference), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWorkerHubInference)).(*[]IWorkerHubInference)

	return out0, err

}

// GetAllInferences is a free data retrieval call binding the contract method 0xf1ea45e3.
//
// Solidity: function getAllInferences(uint256 startId, uint256 count) view returns((uint256[],bytes,uint256,uint256,uint256,address,uint40,uint40,uint40,uint8,address,address,address)[] inferenceData)
func (_IWorkerHub *IWorkerHubSession) GetAllInferences(startId *big.Int, count *big.Int) ([]IWorkerHubInference, error) {
	return _IWorkerHub.Contract.GetAllInferences(&_IWorkerHub.CallOpts, startId, count)
}

// GetAllInferences is a free data retrieval call binding the contract method 0xf1ea45e3.
//
// Solidity: function getAllInferences(uint256 startId, uint256 count) view returns((uint256[],bytes,uint256,uint256,uint256,address,uint40,uint40,uint40,uint8,address,address,address)[] inferenceData)
func (_IWorkerHub *IWorkerHubCallerSession) GetAllInferences(startId *big.Int, count *big.Int) ([]IWorkerHubInference, error) {
	return _IWorkerHub.Contract.GetAllInferences(&_IWorkerHub.CallOpts, startId, count)
}

// GetAllMinerUnstakeRequests is a free data retrieval call binding the contract method 0x9280f078.
//
// Solidity: function getAllMinerUnstakeRequests() view returns(address[] unstakeAddresses, (uint256,uint40)[] unstakeRequests)
func (_IWorkerHub *IWorkerHubCaller) GetAllMinerUnstakeRequests(opts *bind.CallOpts) (struct {
	UnstakeAddresses []common.Address
	UnstakeRequests  []IWorkerHubUnstakeRequest
}, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "getAllMinerUnstakeRequests")

	outstruct := new(struct {
		UnstakeAddresses []common.Address
		UnstakeRequests  []IWorkerHubUnstakeRequest
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UnstakeAddresses = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.UnstakeRequests = *abi.ConvertType(out[1], new([]IWorkerHubUnstakeRequest)).(*[]IWorkerHubUnstakeRequest)

	return *outstruct, err

}

// GetAllMinerUnstakeRequests is a free data retrieval call binding the contract method 0x9280f078.
//
// Solidity: function getAllMinerUnstakeRequests() view returns(address[] unstakeAddresses, (uint256,uint40)[] unstakeRequests)
func (_IWorkerHub *IWorkerHubSession) GetAllMinerUnstakeRequests() (struct {
	UnstakeAddresses []common.Address
	UnstakeRequests  []IWorkerHubUnstakeRequest
}, error) {
	return _IWorkerHub.Contract.GetAllMinerUnstakeRequests(&_IWorkerHub.CallOpts)
}

// GetAllMinerUnstakeRequests is a free data retrieval call binding the contract method 0x9280f078.
//
// Solidity: function getAllMinerUnstakeRequests() view returns(address[] unstakeAddresses, (uint256,uint40)[] unstakeRequests)
func (_IWorkerHub *IWorkerHubCallerSession) GetAllMinerUnstakeRequests() (struct {
	UnstakeAddresses []common.Address
	UnstakeRequests  []IWorkerHubUnstakeRequest
}, error) {
	return _IWorkerHub.Contract.GetAllMinerUnstakeRequests(&_IWorkerHub.CallOpts)
}

// GetAllMiners is a free data retrieval call binding the contract method 0x4b17bf30.
//
// Solidity: function getAllMiners() view returns((uint256,uint256,address,uint40,uint40,uint16)[] minerData)
func (_IWorkerHub *IWorkerHubCaller) GetAllMiners(opts *bind.CallOpts) ([]IWorkerHubWorker, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "getAllMiners")

	if err != nil {
		return *new([]IWorkerHubWorker), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWorkerHubWorker)).(*[]IWorkerHubWorker)

	return out0, err

}

// GetAllMiners is a free data retrieval call binding the contract method 0x4b17bf30.
//
// Solidity: function getAllMiners() view returns((uint256,uint256,address,uint40,uint40,uint16)[] minerData)
func (_IWorkerHub *IWorkerHubSession) GetAllMiners() ([]IWorkerHubWorker, error) {
	return _IWorkerHub.Contract.GetAllMiners(&_IWorkerHub.CallOpts)
}

// GetAllMiners is a free data retrieval call binding the contract method 0x4b17bf30.
//
// Solidity: function getAllMiners() view returns((uint256,uint256,address,uint40,uint40,uint16)[] minerData)
func (_IWorkerHub *IWorkerHubCallerSession) GetAllMiners() ([]IWorkerHubWorker, error) {
	return _IWorkerHub.Contract.GetAllMiners(&_IWorkerHub.CallOpts)
}

// GetAssignmentByInferenceId is a free data retrieval call binding the contract method 0x19a9dc71.
//
// Solidity: function getAssignmentByInferenceId(uint256 _inferId) view returns((uint256,bytes32,bytes32,uint40,address,uint8,uint8,bytes)[])
func (_IWorkerHub *IWorkerHubCaller) GetAssignmentByInferenceId(opts *bind.CallOpts, _inferId *big.Int) ([]IWorkerHubAssignment, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "getAssignmentByInferenceId", _inferId)

	if err != nil {
		return *new([]IWorkerHubAssignment), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWorkerHubAssignment)).(*[]IWorkerHubAssignment)

	return out0, err

}

// GetAssignmentByInferenceId is a free data retrieval call binding the contract method 0x19a9dc71.
//
// Solidity: function getAssignmentByInferenceId(uint256 _inferId) view returns((uint256,bytes32,bytes32,uint40,address,uint8,uint8,bytes)[])
func (_IWorkerHub *IWorkerHubSession) GetAssignmentByInferenceId(_inferId *big.Int) ([]IWorkerHubAssignment, error) {
	return _IWorkerHub.Contract.GetAssignmentByInferenceId(&_IWorkerHub.CallOpts, _inferId)
}

// GetAssignmentByInferenceId is a free data retrieval call binding the contract method 0x19a9dc71.
//
// Solidity: function getAssignmentByInferenceId(uint256 _inferId) view returns((uint256,bytes32,bytes32,uint40,address,uint8,uint8,bytes)[])
func (_IWorkerHub *IWorkerHubCallerSession) GetAssignmentByInferenceId(_inferId *big.Int) ([]IWorkerHubAssignment, error) {
	return _IWorkerHub.Contract.GetAssignmentByInferenceId(&_IWorkerHub.CallOpts, _inferId)
}

// GetAssignmentByMiner is a free data retrieval call binding the contract method 0x5937e5ed.
//
// Solidity: function getAssignmentByMiner(address _minerAddr) view returns((uint256,uint256,uint256,bytes,address,address,uint40,uint40,uint40)[])
func (_IWorkerHub *IWorkerHubCaller) GetAssignmentByMiner(opts *bind.CallOpts, _minerAddr common.Address) ([]IWorkerHubAssignmentInfo, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "getAssignmentByMiner", _minerAddr)

	if err != nil {
		return *new([]IWorkerHubAssignmentInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWorkerHubAssignmentInfo)).(*[]IWorkerHubAssignmentInfo)

	return out0, err

}

// GetAssignmentByMiner is a free data retrieval call binding the contract method 0x5937e5ed.
//
// Solidity: function getAssignmentByMiner(address _minerAddr) view returns((uint256,uint256,uint256,bytes,address,address,uint40,uint40,uint40)[])
func (_IWorkerHub *IWorkerHubSession) GetAssignmentByMiner(_minerAddr common.Address) ([]IWorkerHubAssignmentInfo, error) {
	return _IWorkerHub.Contract.GetAssignmentByMiner(&_IWorkerHub.CallOpts, _minerAddr)
}

// GetAssignmentByMiner is a free data retrieval call binding the contract method 0x5937e5ed.
//
// Solidity: function getAssignmentByMiner(address _minerAddr) view returns((uint256,uint256,uint256,bytes,address,address,uint40,uint40,uint40)[])
func (_IWorkerHub *IWorkerHubCallerSession) GetAssignmentByMiner(_minerAddr common.Address) ([]IWorkerHubAssignmentInfo, error) {
	return _IWorkerHub.Contract.GetAssignmentByMiner(&_IWorkerHub.CallOpts, _minerAddr)
}

// GetInferenceInfo is a free data retrieval call binding the contract method 0x08c05903.
//
// Solidity: function getInferenceInfo(uint256 _inferenceId) view returns((uint256[],bytes,uint256,uint256,uint256,address,uint40,uint40,uint40,uint8,address,address,address))
func (_IWorkerHub *IWorkerHubCaller) GetInferenceInfo(opts *bind.CallOpts, _inferenceId *big.Int) (IWorkerHubInference, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "getInferenceInfo", _inferenceId)

	if err != nil {
		return *new(IWorkerHubInference), err
	}

	out0 := *abi.ConvertType(out[0], new(IWorkerHubInference)).(*IWorkerHubInference)

	return out0, err

}

// GetInferenceInfo is a free data retrieval call binding the contract method 0x08c05903.
//
// Solidity: function getInferenceInfo(uint256 _inferenceId) view returns((uint256[],bytes,uint256,uint256,uint256,address,uint40,uint40,uint40,uint8,address,address,address))
func (_IWorkerHub *IWorkerHubSession) GetInferenceInfo(_inferenceId *big.Int) (IWorkerHubInference, error) {
	return _IWorkerHub.Contract.GetInferenceInfo(&_IWorkerHub.CallOpts, _inferenceId)
}

// GetInferenceInfo is a free data retrieval call binding the contract method 0x08c05903.
//
// Solidity: function getInferenceInfo(uint256 _inferenceId) view returns((uint256[],bytes,uint256,uint256,uint256,address,uint40,uint40,uint40,uint8,address,address,address))
func (_IWorkerHub *IWorkerHubCallerSession) GetInferenceInfo(_inferenceId *big.Int) (IWorkerHubInference, error) {
	return _IWorkerHub.Contract.GetInferenceInfo(&_IWorkerHub.CallOpts, _inferenceId)
}

// GetMinFeeToUse is a free data retrieval call binding the contract method 0xafc1fce7.
//
// Solidity: function getMinFeeToUse(address _modelAddress) view returns(uint256)
func (_IWorkerHub *IWorkerHubCaller) GetMinFeeToUse(opts *bind.CallOpts, _modelAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "getMinFeeToUse", _modelAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinFeeToUse is a free data retrieval call binding the contract method 0xafc1fce7.
//
// Solidity: function getMinFeeToUse(address _modelAddress) view returns(uint256)
func (_IWorkerHub *IWorkerHubSession) GetMinFeeToUse(_modelAddress common.Address) (*big.Int, error) {
	return _IWorkerHub.Contract.GetMinFeeToUse(&_IWorkerHub.CallOpts, _modelAddress)
}

// GetMinFeeToUse is a free data retrieval call binding the contract method 0xafc1fce7.
//
// Solidity: function getMinFeeToUse(address _modelAddress) view returns(uint256)
func (_IWorkerHub *IWorkerHubCallerSession) GetMinFeeToUse(_modelAddress common.Address) (*big.Int, error) {
	return _IWorkerHub.Contract.GetMinFeeToUse(&_IWorkerHub.CallOpts, _modelAddress)
}

// GetMinerAddresses is a free data retrieval call binding the contract method 0xe8d6f2f1.
//
// Solidity: function getMinerAddresses() view returns(address[])
func (_IWorkerHub *IWorkerHubCaller) GetMinerAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "getMinerAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMinerAddresses is a free data retrieval call binding the contract method 0xe8d6f2f1.
//
// Solidity: function getMinerAddresses() view returns(address[])
func (_IWorkerHub *IWorkerHubSession) GetMinerAddresses() ([]common.Address, error) {
	return _IWorkerHub.Contract.GetMinerAddresses(&_IWorkerHub.CallOpts)
}

// GetMinerAddresses is a free data retrieval call binding the contract method 0xe8d6f2f1.
//
// Solidity: function getMinerAddresses() view returns(address[])
func (_IWorkerHub *IWorkerHubCallerSession) GetMinerAddresses() ([]common.Address, error) {
	return _IWorkerHub.Contract.GetMinerAddresses(&_IWorkerHub.CallOpts)
}

// GetMinerAddressesOfModel is a free data retrieval call binding the contract method 0x47253baa.
//
// Solidity: function getMinerAddressesOfModel(address _model) view returns(address[])
func (_IWorkerHub *IWorkerHubCaller) GetMinerAddressesOfModel(opts *bind.CallOpts, _model common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "getMinerAddressesOfModel", _model)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMinerAddressesOfModel is a free data retrieval call binding the contract method 0x47253baa.
//
// Solidity: function getMinerAddressesOfModel(address _model) view returns(address[])
func (_IWorkerHub *IWorkerHubSession) GetMinerAddressesOfModel(_model common.Address) ([]common.Address, error) {
	return _IWorkerHub.Contract.GetMinerAddressesOfModel(&_IWorkerHub.CallOpts, _model)
}

// GetMinerAddressesOfModel is a free data retrieval call binding the contract method 0x47253baa.
//
// Solidity: function getMinerAddressesOfModel(address _model) view returns(address[])
func (_IWorkerHub *IWorkerHubCallerSession) GetMinerAddressesOfModel(_model common.Address) ([]common.Address, error) {
	return _IWorkerHub.Contract.GetMinerAddressesOfModel(&_IWorkerHub.CallOpts, _model)
}

// GetMiners is a free data retrieval call binding the contract method 0x1633da6e.
//
// Solidity: function getMiners() view returns((address,uint256,uint256,address,uint40,uint40,uint16)[])
func (_IWorkerHub *IWorkerHubCaller) GetMiners(opts *bind.CallOpts) ([]IWorkerHubWorkerInfo, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "getMiners")

	if err != nil {
		return *new([]IWorkerHubWorkerInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWorkerHubWorkerInfo)).(*[]IWorkerHubWorkerInfo)

	return out0, err

}

// GetMiners is a free data retrieval call binding the contract method 0x1633da6e.
//
// Solidity: function getMiners() view returns((address,uint256,uint256,address,uint40,uint40,uint16)[])
func (_IWorkerHub *IWorkerHubSession) GetMiners() ([]IWorkerHubWorkerInfo, error) {
	return _IWorkerHub.Contract.GetMiners(&_IWorkerHub.CallOpts)
}

// GetMiners is a free data retrieval call binding the contract method 0x1633da6e.
//
// Solidity: function getMiners() view returns((address,uint256,uint256,address,uint40,uint40,uint16)[])
func (_IWorkerHub *IWorkerHubCallerSession) GetMiners() ([]IWorkerHubWorkerInfo, error) {
	return _IWorkerHub.Contract.GetMiners(&_IWorkerHub.CallOpts)
}

// GetMintingAssignmentsOfInference is a free data retrieval call binding the contract method 0x5eec7b20.
//
// Solidity: function getMintingAssignmentsOfInference(uint256 _inferenceId) view returns((uint256,uint256,uint256,bytes,address,address,uint40,uint40,uint40)[])
func (_IWorkerHub *IWorkerHubCaller) GetMintingAssignmentsOfInference(opts *bind.CallOpts, _inferenceId *big.Int) ([]IWorkerHubAssignmentInfo, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "getMintingAssignmentsOfInference", _inferenceId)

	if err != nil {
		return *new([]IWorkerHubAssignmentInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWorkerHubAssignmentInfo)).(*[]IWorkerHubAssignmentInfo)

	return out0, err

}

// GetMintingAssignmentsOfInference is a free data retrieval call binding the contract method 0x5eec7b20.
//
// Solidity: function getMintingAssignmentsOfInference(uint256 _inferenceId) view returns((uint256,uint256,uint256,bytes,address,address,uint40,uint40,uint40)[])
func (_IWorkerHub *IWorkerHubSession) GetMintingAssignmentsOfInference(_inferenceId *big.Int) ([]IWorkerHubAssignmentInfo, error) {
	return _IWorkerHub.Contract.GetMintingAssignmentsOfInference(&_IWorkerHub.CallOpts, _inferenceId)
}

// GetMintingAssignmentsOfInference is a free data retrieval call binding the contract method 0x5eec7b20.
//
// Solidity: function getMintingAssignmentsOfInference(uint256 _inferenceId) view returns((uint256,uint256,uint256,bytes,address,address,uint40,uint40,uint40)[])
func (_IWorkerHub *IWorkerHubCallerSession) GetMintingAssignmentsOfInference(_inferenceId *big.Int) ([]IWorkerHubAssignmentInfo, error) {
	return _IWorkerHub.Contract.GetMintingAssignmentsOfInference(&_IWorkerHub.CallOpts, _inferenceId)
}

// GetModelAddresses is a free data retrieval call binding the contract method 0x9ae49cd3.
//
// Solidity: function getModelAddresses() view returns(address[])
func (_IWorkerHub *IWorkerHubCaller) GetModelAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "getModelAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetModelAddresses is a free data retrieval call binding the contract method 0x9ae49cd3.
//
// Solidity: function getModelAddresses() view returns(address[])
func (_IWorkerHub *IWorkerHubSession) GetModelAddresses() ([]common.Address, error) {
	return _IWorkerHub.Contract.GetModelAddresses(&_IWorkerHub.CallOpts)
}

// GetModelAddresses is a free data retrieval call binding the contract method 0x9ae49cd3.
//
// Solidity: function getModelAddresses() view returns(address[])
func (_IWorkerHub *IWorkerHubCallerSession) GetModelAddresses() ([]common.Address, error) {
	return _IWorkerHub.Contract.GetModelAddresses(&_IWorkerHub.CallOpts)
}

// GetNOMiner is a free data retrieval call binding the contract method 0xd2d89be8.
//
// Solidity: function getNOMiner() view returns(uint256)
func (_IWorkerHub *IWorkerHubCaller) GetNOMiner(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "getNOMiner")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNOMiner is a free data retrieval call binding the contract method 0xd2d89be8.
//
// Solidity: function getNOMiner() view returns(uint256)
func (_IWorkerHub *IWorkerHubSession) GetNOMiner() (*big.Int, error) {
	return _IWorkerHub.Contract.GetNOMiner(&_IWorkerHub.CallOpts)
}

// GetNOMiner is a free data retrieval call binding the contract method 0xd2d89be8.
//
// Solidity: function getNOMiner() view returns(uint256)
func (_IWorkerHub *IWorkerHubCallerSession) GetNOMiner() (*big.Int, error) {
	return _IWorkerHub.Contract.GetNOMiner(&_IWorkerHub.CallOpts)
}

// GetRoleByAssigmentId is a free data retrieval call binding the contract method 0xca0c80fc.
//
// Solidity: function getRoleByAssigmentId(uint256 _assignmentId) view returns(uint8)
func (_IWorkerHub *IWorkerHubCaller) GetRoleByAssigmentId(opts *bind.CallOpts, _assignmentId *big.Int) (uint8, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "getRoleByAssigmentId", _assignmentId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetRoleByAssigmentId is a free data retrieval call binding the contract method 0xca0c80fc.
//
// Solidity: function getRoleByAssigmentId(uint256 _assignmentId) view returns(uint8)
func (_IWorkerHub *IWorkerHubSession) GetRoleByAssigmentId(_assignmentId *big.Int) (uint8, error) {
	return _IWorkerHub.Contract.GetRoleByAssigmentId(&_IWorkerHub.CallOpts, _assignmentId)
}

// GetRoleByAssigmentId is a free data retrieval call binding the contract method 0xca0c80fc.
//
// Solidity: function getRoleByAssigmentId(uint256 _assignmentId) view returns(uint8)
func (_IWorkerHub *IWorkerHubCallerSession) GetRoleByAssigmentId(_assignmentId *big.Int) (uint8, error) {
	return _IWorkerHub.Contract.GetRoleByAssigmentId(&_IWorkerHub.CallOpts, _assignmentId)
}

// InferenceNumber is a free data retrieval call binding the contract method 0xf80dca98.
//
// Solidity: function inferenceNumber() view returns(uint256)
func (_IWorkerHub *IWorkerHubCaller) InferenceNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "inferenceNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InferenceNumber is a free data retrieval call binding the contract method 0xf80dca98.
//
// Solidity: function inferenceNumber() view returns(uint256)
func (_IWorkerHub *IWorkerHubSession) InferenceNumber() (*big.Int, error) {
	return _IWorkerHub.Contract.InferenceNumber(&_IWorkerHub.CallOpts)
}

// InferenceNumber is a free data retrieval call binding the contract method 0xf80dca98.
//
// Solidity: function inferenceNumber() view returns(uint256)
func (_IWorkerHub *IWorkerHubCallerSession) InferenceNumber() (*big.Int, error) {
	return _IWorkerHub.Contract.InferenceNumber(&_IWorkerHub.CallOpts)
}

// IsAssignmentPending is a free data retrieval call binding the contract method 0x57a38def.
//
// Solidity: function isAssignmentPending(uint256 _assignmentId) view returns(bool)
func (_IWorkerHub *IWorkerHubCaller) IsAssignmentPending(opts *bind.CallOpts, _assignmentId *big.Int) (bool, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "isAssignmentPending", _assignmentId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAssignmentPending is a free data retrieval call binding the contract method 0x57a38def.
//
// Solidity: function isAssignmentPending(uint256 _assignmentId) view returns(bool)
func (_IWorkerHub *IWorkerHubSession) IsAssignmentPending(_assignmentId *big.Int) (bool, error) {
	return _IWorkerHub.Contract.IsAssignmentPending(&_IWorkerHub.CallOpts, _assignmentId)
}

// IsAssignmentPending is a free data retrieval call binding the contract method 0x57a38def.
//
// Solidity: function isAssignmentPending(uint256 _assignmentId) view returns(bool)
func (_IWorkerHub *IWorkerHubCallerSession) IsAssignmentPending(_assignmentId *big.Int) (bool, error) {
	return _IWorkerHub.Contract.IsAssignmentPending(&_IWorkerHub.CallOpts, _assignmentId)
}

// IsReferrer is a free data retrieval call binding the contract method 0xd64d6968.
//
// Solidity: function isReferrer(address ) view returns(bool)
func (_IWorkerHub *IWorkerHubCaller) IsReferrer(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "isReferrer", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReferrer is a free data retrieval call binding the contract method 0xd64d6968.
//
// Solidity: function isReferrer(address ) view returns(bool)
func (_IWorkerHub *IWorkerHubSession) IsReferrer(arg0 common.Address) (bool, error) {
	return _IWorkerHub.Contract.IsReferrer(&_IWorkerHub.CallOpts, arg0)
}

// IsReferrer is a free data retrieval call binding the contract method 0xd64d6968.
//
// Solidity: function isReferrer(address ) view returns(bool)
func (_IWorkerHub *IWorkerHubCallerSession) IsReferrer(arg0 common.Address) (bool, error) {
	return _IWorkerHub.Contract.IsReferrer(&_IWorkerHub.CallOpts, arg0)
}

// L2Owner is a free data retrieval call binding the contract method 0xf003a0c5.
//
// Solidity: function l2Owner() view returns(address)
func (_IWorkerHub *IWorkerHubCaller) L2Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "l2Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Owner is a free data retrieval call binding the contract method 0xf003a0c5.
//
// Solidity: function l2Owner() view returns(address)
func (_IWorkerHub *IWorkerHubSession) L2Owner() (common.Address, error) {
	return _IWorkerHub.Contract.L2Owner(&_IWorkerHub.CallOpts)
}

// L2Owner is a free data retrieval call binding the contract method 0xf003a0c5.
//
// Solidity: function l2Owner() view returns(address)
func (_IWorkerHub *IWorkerHubCallerSession) L2Owner() (common.Address, error) {
	return _IWorkerHub.Contract.L2Owner(&_IWorkerHub.CallOpts)
}

// LastBlock is a free data retrieval call binding the contract method 0x806b984f.
//
// Solidity: function lastBlock() view returns(uint256)
func (_IWorkerHub *IWorkerHubCaller) LastBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "lastBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBlock is a free data retrieval call binding the contract method 0x806b984f.
//
// Solidity: function lastBlock() view returns(uint256)
func (_IWorkerHub *IWorkerHubSession) LastBlock() (*big.Int, error) {
	return _IWorkerHub.Contract.LastBlock(&_IWorkerHub.CallOpts)
}

// LastBlock is a free data retrieval call binding the contract method 0x806b984f.
//
// Solidity: function lastBlock() view returns(uint256)
func (_IWorkerHub *IWorkerHubCallerSession) LastBlock() (*big.Int, error) {
	return _IWorkerHub.Contract.LastBlock(&_IWorkerHub.CallOpts)
}

// MaximumTier is a free data retrieval call binding the contract method 0x0716187f.
//
// Solidity: function maximumTier() view returns(uint16)
func (_IWorkerHub *IWorkerHubCaller) MaximumTier(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "maximumTier")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MaximumTier is a free data retrieval call binding the contract method 0x0716187f.
//
// Solidity: function maximumTier() view returns(uint16)
func (_IWorkerHub *IWorkerHubSession) MaximumTier() (uint16, error) {
	return _IWorkerHub.Contract.MaximumTier(&_IWorkerHub.CallOpts)
}

// MaximumTier is a free data retrieval call binding the contract method 0x0716187f.
//
// Solidity: function maximumTier() view returns(uint16)
func (_IWorkerHub *IWorkerHubCallerSession) MaximumTier() (uint16, error) {
	return _IWorkerHub.Contract.MaximumTier(&_IWorkerHub.CallOpts)
}

// MinFeeToUse is a free data retrieval call binding the contract method 0x2a1a8ca8.
//
// Solidity: function minFeeToUse() view returns(uint256)
func (_IWorkerHub *IWorkerHubCaller) MinFeeToUse(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "minFeeToUse")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinFeeToUse is a free data retrieval call binding the contract method 0x2a1a8ca8.
//
// Solidity: function minFeeToUse() view returns(uint256)
func (_IWorkerHub *IWorkerHubSession) MinFeeToUse() (*big.Int, error) {
	return _IWorkerHub.Contract.MinFeeToUse(&_IWorkerHub.CallOpts)
}

// MinFeeToUse is a free data retrieval call binding the contract method 0x2a1a8ca8.
//
// Solidity: function minFeeToUse() view returns(uint256)
func (_IWorkerHub *IWorkerHubCallerSession) MinFeeToUse() (*big.Int, error) {
	return _IWorkerHub.Contract.MinFeeToUse(&_IWorkerHub.CallOpts)
}

// MinerMinimumStake is a free data retrieval call binding the contract method 0x3304f456.
//
// Solidity: function minerMinimumStake() view returns(uint256)
func (_IWorkerHub *IWorkerHubCaller) MinerMinimumStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "minerMinimumStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinerMinimumStake is a free data retrieval call binding the contract method 0x3304f456.
//
// Solidity: function minerMinimumStake() view returns(uint256)
func (_IWorkerHub *IWorkerHubSession) MinerMinimumStake() (*big.Int, error) {
	return _IWorkerHub.Contract.MinerMinimumStake(&_IWorkerHub.CallOpts)
}

// MinerMinimumStake is a free data retrieval call binding the contract method 0x3304f456.
//
// Solidity: function minerMinimumStake() view returns(uint256)
func (_IWorkerHub *IWorkerHubCallerSession) MinerMinimumStake() (*big.Int, error) {
	return _IWorkerHub.Contract.MinerMinimumStake(&_IWorkerHub.CallOpts)
}

// MinerRequirement is a free data retrieval call binding the contract method 0xdd9b9766.
//
// Solidity: function minerRequirement() view returns(uint8)
func (_IWorkerHub *IWorkerHubCaller) MinerRequirement(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "minerRequirement")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MinerRequirement is a free data retrieval call binding the contract method 0xdd9b9766.
//
// Solidity: function minerRequirement() view returns(uint8)
func (_IWorkerHub *IWorkerHubSession) MinerRequirement() (uint8, error) {
	return _IWorkerHub.Contract.MinerRequirement(&_IWorkerHub.CallOpts)
}

// MinerRequirement is a free data retrieval call binding the contract method 0xdd9b9766.
//
// Solidity: function minerRequirement() view returns(uint8)
func (_IWorkerHub *IWorkerHubCallerSession) MinerRequirement() (uint8, error) {
	return _IWorkerHub.Contract.MinerRequirement(&_IWorkerHub.CallOpts)
}

// MinerUnstakeRequests is a free data retrieval call binding the contract method 0x191a54d8.
//
// Solidity: function minerUnstakeRequests(address ) view returns(uint256 stake, uint40 unlockAt)
func (_IWorkerHub *IWorkerHubCaller) MinerUnstakeRequests(opts *bind.CallOpts, arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "minerUnstakeRequests", arg0)

	outstruct := new(struct {
		Stake    *big.Int
		UnlockAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Stake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UnlockAt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MinerUnstakeRequests is a free data retrieval call binding the contract method 0x191a54d8.
//
// Solidity: function minerUnstakeRequests(address ) view returns(uint256 stake, uint40 unlockAt)
func (_IWorkerHub *IWorkerHubSession) MinerUnstakeRequests(arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	return _IWorkerHub.Contract.MinerUnstakeRequests(&_IWorkerHub.CallOpts, arg0)
}

// MinerUnstakeRequests is a free data retrieval call binding the contract method 0x191a54d8.
//
// Solidity: function minerUnstakeRequests(address ) view returns(uint256 stake, uint40 unlockAt)
func (_IWorkerHub *IWorkerHubCallerSession) MinerUnstakeRequests(arg0 common.Address) (struct {
	Stake    *big.Int
	UnlockAt *big.Int
}, error) {
	return _IWorkerHub.Contract.MinerUnstakeRequests(&_IWorkerHub.CallOpts, arg0)
}

// Miners is a free data retrieval call binding the contract method 0x648ec7b9.
//
// Solidity: function miners(address ) view returns(uint256 stake, uint256 commitment, address modelAddress, uint40 lastClaimedEpoch, uint40 activeTime, uint16 tier)
func (_IWorkerHub *IWorkerHubCaller) Miners(opts *bind.CallOpts, arg0 common.Address) (struct {
	Stake            *big.Int
	Commitment       *big.Int
	ModelAddress     common.Address
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "miners", arg0)

	outstruct := new(struct {
		Stake            *big.Int
		Commitment       *big.Int
		ModelAddress     common.Address
		LastClaimedEpoch *big.Int
		ActiveTime       *big.Int
		Tier             uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Stake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Commitment = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ModelAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.LastClaimedEpoch = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ActiveTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Tier = *abi.ConvertType(out[5], new(uint16)).(*uint16)

	return *outstruct, err

}

// Miners is a free data retrieval call binding the contract method 0x648ec7b9.
//
// Solidity: function miners(address ) view returns(uint256 stake, uint256 commitment, address modelAddress, uint40 lastClaimedEpoch, uint40 activeTime, uint16 tier)
func (_IWorkerHub *IWorkerHubSession) Miners(arg0 common.Address) (struct {
	Stake            *big.Int
	Commitment       *big.Int
	ModelAddress     common.Address
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	return _IWorkerHub.Contract.Miners(&_IWorkerHub.CallOpts, arg0)
}

// Miners is a free data retrieval call binding the contract method 0x648ec7b9.
//
// Solidity: function miners(address ) view returns(uint256 stake, uint256 commitment, address modelAddress, uint40 lastClaimedEpoch, uint40 activeTime, uint16 tier)
func (_IWorkerHub *IWorkerHubCallerSession) Miners(arg0 common.Address) (struct {
	Stake            *big.Int
	Commitment       *big.Int
	ModelAddress     common.Address
	LastClaimedEpoch *big.Int
	ActiveTime       *big.Int
	Tier             uint16
}, error) {
	return _IWorkerHub.Contract.Miners(&_IWorkerHub.CallOpts, arg0)
}

// ModelScoring is a free data retrieval call binding the contract method 0xfe0503c0.
//
// Solidity: function modelScoring() view returns(address)
func (_IWorkerHub *IWorkerHubCaller) ModelScoring(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "modelScoring")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ModelScoring is a free data retrieval call binding the contract method 0xfe0503c0.
//
// Solidity: function modelScoring() view returns(address)
func (_IWorkerHub *IWorkerHubSession) ModelScoring() (common.Address, error) {
	return _IWorkerHub.Contract.ModelScoring(&_IWorkerHub.CallOpts)
}

// ModelScoring is a free data retrieval call binding the contract method 0xfe0503c0.
//
// Solidity: function modelScoring() view returns(address)
func (_IWorkerHub *IWorkerHubCallerSession) ModelScoring() (common.Address, error) {
	return _IWorkerHub.Contract.ModelScoring(&_IWorkerHub.CallOpts)
}

// Models is a free data retrieval call binding the contract method 0x54917f83.
//
// Solidity: function models(address ) view returns(uint256 minimumFee, uint32 tier)
func (_IWorkerHub *IWorkerHubCaller) Models(opts *bind.CallOpts, arg0 common.Address) (struct {
	MinimumFee *big.Int
	Tier       uint32
}, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "models", arg0)

	outstruct := new(struct {
		MinimumFee *big.Int
		Tier       uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinimumFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Tier = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// Models is a free data retrieval call binding the contract method 0x54917f83.
//
// Solidity: function models(address ) view returns(uint256 minimumFee, uint32 tier)
func (_IWorkerHub *IWorkerHubSession) Models(arg0 common.Address) (struct {
	MinimumFee *big.Int
	Tier       uint32
}, error) {
	return _IWorkerHub.Contract.Models(&_IWorkerHub.CallOpts, arg0)
}

// Models is a free data retrieval call binding the contract method 0x54917f83.
//
// Solidity: function models(address ) view returns(uint256 minimumFee, uint32 tier)
func (_IWorkerHub *IWorkerHubCallerSession) Models(arg0 common.Address) (struct {
	MinimumFee *big.Int
	Tier       uint32
}, error) {
	return _IWorkerHub.Contract.Models(&_IWorkerHub.CallOpts, arg0)
}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address _miner) view returns(uint256)
func (_IWorkerHub *IWorkerHubCaller) Multiplier(opts *bind.CallOpts, _miner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "multiplier", _miner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address _miner) view returns(uint256)
func (_IWorkerHub *IWorkerHubSession) Multiplier(_miner common.Address) (*big.Int, error) {
	return _IWorkerHub.Contract.Multiplier(&_IWorkerHub.CallOpts, _miner)
}

// Multiplier is a free data retrieval call binding the contract method 0xa9b3f8b7.
//
// Solidity: function multiplier(address _miner) view returns(uint256)
func (_IWorkerHub *IWorkerHubCallerSession) Multiplier(_miner common.Address) (*big.Int, error) {
	return _IWorkerHub.Contract.Multiplier(&_IWorkerHub.CallOpts, _miner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IWorkerHub *IWorkerHubCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IWorkerHub *IWorkerHubSession) Owner() (common.Address, error) {
	return _IWorkerHub.Contract.Owner(&_IWorkerHub.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IWorkerHub *IWorkerHubCallerSession) Owner() (common.Address, error) {
	return _IWorkerHub.Contract.Owner(&_IWorkerHub.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IWorkerHub *IWorkerHubCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IWorkerHub *IWorkerHubSession) Paused() (bool, error) {
	return _IWorkerHub.Contract.Paused(&_IWorkerHub.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IWorkerHub *IWorkerHubCallerSession) Paused() (bool, error) {
	return _IWorkerHub.Contract.Paused(&_IWorkerHub.CallOpts)
}

// PenaltyDuration is a free data retrieval call binding the contract method 0x5aa1326c.
//
// Solidity: function penaltyDuration() view returns(uint40)
func (_IWorkerHub *IWorkerHubCaller) PenaltyDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "penaltyDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PenaltyDuration is a free data retrieval call binding the contract method 0x5aa1326c.
//
// Solidity: function penaltyDuration() view returns(uint40)
func (_IWorkerHub *IWorkerHubSession) PenaltyDuration() (*big.Int, error) {
	return _IWorkerHub.Contract.PenaltyDuration(&_IWorkerHub.CallOpts)
}

// PenaltyDuration is a free data retrieval call binding the contract method 0x5aa1326c.
//
// Solidity: function penaltyDuration() view returns(uint40)
func (_IWorkerHub *IWorkerHubCallerSession) PenaltyDuration() (*big.Int, error) {
	return _IWorkerHub.Contract.PenaltyDuration(&_IWorkerHub.CallOpts)
}

// RevealDuration is a free data retrieval call binding the contract method 0x886a6de1.
//
// Solidity: function revealDuration() view returns(uint40)
func (_IWorkerHub *IWorkerHubCaller) RevealDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "revealDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RevealDuration is a free data retrieval call binding the contract method 0x886a6de1.
//
// Solidity: function revealDuration() view returns(uint40)
func (_IWorkerHub *IWorkerHubSession) RevealDuration() (*big.Int, error) {
	return _IWorkerHub.Contract.RevealDuration(&_IWorkerHub.CallOpts)
}

// RevealDuration is a free data retrieval call binding the contract method 0x886a6de1.
//
// Solidity: function revealDuration() view returns(uint40)
func (_IWorkerHub *IWorkerHubCallerSession) RevealDuration() (*big.Int, error) {
	return _IWorkerHub.Contract.RevealDuration(&_IWorkerHub.CallOpts)
}

// RewardInEpoch is a free data retrieval call binding the contract method 0x652ff159.
//
// Solidity: function rewardInEpoch(uint256 ) view returns(uint256 perfReward, uint256 epochReward, uint256 totalTaskCompleted, uint256 totalMiner)
func (_IWorkerHub *IWorkerHubCaller) RewardInEpoch(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PerfReward         *big.Int
	EpochReward        *big.Int
	TotalTaskCompleted *big.Int
	TotalMiner         *big.Int
}, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "rewardInEpoch", arg0)

	outstruct := new(struct {
		PerfReward         *big.Int
		EpochReward        *big.Int
		TotalTaskCompleted *big.Int
		TotalMiner         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PerfReward = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EpochReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalTaskCompleted = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalMiner = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RewardInEpoch is a free data retrieval call binding the contract method 0x652ff159.
//
// Solidity: function rewardInEpoch(uint256 ) view returns(uint256 perfReward, uint256 epochReward, uint256 totalTaskCompleted, uint256 totalMiner)
func (_IWorkerHub *IWorkerHubSession) RewardInEpoch(arg0 *big.Int) (struct {
	PerfReward         *big.Int
	EpochReward        *big.Int
	TotalTaskCompleted *big.Int
	TotalMiner         *big.Int
}, error) {
	return _IWorkerHub.Contract.RewardInEpoch(&_IWorkerHub.CallOpts, arg0)
}

// RewardInEpoch is a free data retrieval call binding the contract method 0x652ff159.
//
// Solidity: function rewardInEpoch(uint256 ) view returns(uint256 perfReward, uint256 epochReward, uint256 totalTaskCompleted, uint256 totalMiner)
func (_IWorkerHub *IWorkerHubCallerSession) RewardInEpoch(arg0 *big.Int) (struct {
	PerfReward         *big.Int
	EpochReward        *big.Int
	TotalTaskCompleted *big.Int
	TotalMiner         *big.Int
}, error) {
	return _IWorkerHub.Contract.RewardInEpoch(&_IWorkerHub.CallOpts, arg0)
}

// RewardPerEpoch is a free data retrieval call binding the contract method 0x84449a9d.
//
// Solidity: function rewardPerEpoch() view returns(uint256)
func (_IWorkerHub *IWorkerHubCaller) RewardPerEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "rewardPerEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerEpoch is a free data retrieval call binding the contract method 0x84449a9d.
//
// Solidity: function rewardPerEpoch() view returns(uint256)
func (_IWorkerHub *IWorkerHubSession) RewardPerEpoch() (*big.Int, error) {
	return _IWorkerHub.Contract.RewardPerEpoch(&_IWorkerHub.CallOpts)
}

// RewardPerEpoch is a free data retrieval call binding the contract method 0x84449a9d.
//
// Solidity: function rewardPerEpoch() view returns(uint256)
func (_IWorkerHub *IWorkerHubCallerSession) RewardPerEpoch() (*big.Int, error) {
	return _IWorkerHub.Contract.RewardPerEpoch(&_IWorkerHub.CallOpts)
}

// SubmitDuration is a free data retrieval call binding the contract method 0xcc56b6f8.
//
// Solidity: function submitDuration() view returns(uint40)
func (_IWorkerHub *IWorkerHubCaller) SubmitDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "submitDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubmitDuration is a free data retrieval call binding the contract method 0xcc56b6f8.
//
// Solidity: function submitDuration() view returns(uint40)
func (_IWorkerHub *IWorkerHubSession) SubmitDuration() (*big.Int, error) {
	return _IWorkerHub.Contract.SubmitDuration(&_IWorkerHub.CallOpts)
}

// SubmitDuration is a free data retrieval call binding the contract method 0xcc56b6f8.
//
// Solidity: function submitDuration() view returns(uint40)
func (_IWorkerHub *IWorkerHubCallerSession) SubmitDuration() (*big.Int, error) {
	return _IWorkerHub.Contract.SubmitDuration(&_IWorkerHub.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_IWorkerHub *IWorkerHubCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_IWorkerHub *IWorkerHubSession) Treasury() (common.Address, error) {
	return _IWorkerHub.Contract.Treasury(&_IWorkerHub.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_IWorkerHub *IWorkerHubCallerSession) Treasury() (common.Address, error) {
	return _IWorkerHub.Contract.Treasury(&_IWorkerHub.CallOpts)
}

// UnstakeDelayTime is a free data retrieval call binding the contract method 0xe4fefd65.
//
// Solidity: function unstakeDelayTime() view returns(uint40)
func (_IWorkerHub *IWorkerHubCaller) UnstakeDelayTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "unstakeDelayTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnstakeDelayTime is a free data retrieval call binding the contract method 0xe4fefd65.
//
// Solidity: function unstakeDelayTime() view returns(uint40)
func (_IWorkerHub *IWorkerHubSession) UnstakeDelayTime() (*big.Int, error) {
	return _IWorkerHub.Contract.UnstakeDelayTime(&_IWorkerHub.CallOpts)
}

// UnstakeDelayTime is a free data retrieval call binding the contract method 0xe4fefd65.
//
// Solidity: function unstakeDelayTime() view returns(uint40)
func (_IWorkerHub *IWorkerHubCallerSession) UnstakeDelayTime() (*big.Int, error) {
	return _IWorkerHub.Contract.UnstakeDelayTime(&_IWorkerHub.CallOpts)
}

// ValidateDAOSupplyIncrease is a free data retrieval call binding the contract method 0xd7acb1ea.
//
// Solidity: function validateDAOSupplyIncrease(bool _isReferred) view returns(bool notReachedLimit)
func (_IWorkerHub *IWorkerHubCaller) ValidateDAOSupplyIncrease(opts *bind.CallOpts, _isReferred bool) (bool, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "validateDAOSupplyIncrease", _isReferred)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateDAOSupplyIncrease is a free data retrieval call binding the contract method 0xd7acb1ea.
//
// Solidity: function validateDAOSupplyIncrease(bool _isReferred) view returns(bool notReachedLimit)
func (_IWorkerHub *IWorkerHubSession) ValidateDAOSupplyIncrease(_isReferred bool) (bool, error) {
	return _IWorkerHub.Contract.ValidateDAOSupplyIncrease(&_IWorkerHub.CallOpts, _isReferred)
}

// ValidateDAOSupplyIncrease is a free data retrieval call binding the contract method 0xd7acb1ea.
//
// Solidity: function validateDAOSupplyIncrease(bool _isReferred) view returns(bool notReachedLimit)
func (_IWorkerHub *IWorkerHubCallerSession) ValidateDAOSupplyIncrease(_isReferred bool) (bool, error) {
	return _IWorkerHub.Contract.ValidateDAOSupplyIncrease(&_IWorkerHub.CallOpts, _isReferred)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_IWorkerHub *IWorkerHubCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_IWorkerHub *IWorkerHubSession) Version() (string, error) {
	return _IWorkerHub.Contract.Version(&_IWorkerHub.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_IWorkerHub *IWorkerHubCallerSession) Version() (string, error) {
	return _IWorkerHub.Contract.Version(&_IWorkerHub.CallOpts)
}

// VotingInfo is a free data retrieval call binding the contract method 0xe5309a66.
//
// Solidity: function votingInfo(uint256 inferId) view returns(uint8 totalCommit, uint8 totalReveal)
func (_IWorkerHub *IWorkerHubCaller) VotingInfo(opts *bind.CallOpts, inferId *big.Int) (struct {
	TotalCommit uint8
	TotalReveal uint8
}, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "votingInfo", inferId)

	outstruct := new(struct {
		TotalCommit uint8
		TotalReveal uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalCommit = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.TotalReveal = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// VotingInfo is a free data retrieval call binding the contract method 0xe5309a66.
//
// Solidity: function votingInfo(uint256 inferId) view returns(uint8 totalCommit, uint8 totalReveal)
func (_IWorkerHub *IWorkerHubSession) VotingInfo(inferId *big.Int) (struct {
	TotalCommit uint8
	TotalReveal uint8
}, error) {
	return _IWorkerHub.Contract.VotingInfo(&_IWorkerHub.CallOpts, inferId)
}

// VotingInfo is a free data retrieval call binding the contract method 0xe5309a66.
//
// Solidity: function votingInfo(uint256 inferId) view returns(uint8 totalCommit, uint8 totalReveal)
func (_IWorkerHub *IWorkerHubCallerSession) VotingInfo(inferId *big.Int) (struct {
	TotalCommit uint8
	TotalReveal uint8
}, error) {
	return _IWorkerHub.Contract.VotingInfo(&_IWorkerHub.CallOpts, inferId)
}

// WEAI is a free data retrieval call binding the contract method 0x0dc7df53.
//
// Solidity: function wEAI() view returns(address)
func (_IWorkerHub *IWorkerHubCaller) WEAI(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "wEAI")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WEAI is a free data retrieval call binding the contract method 0x0dc7df53.
//
// Solidity: function wEAI() view returns(address)
func (_IWorkerHub *IWorkerHubSession) WEAI() (common.Address, error) {
	return _IWorkerHub.Contract.WEAI(&_IWorkerHub.CallOpts)
}

// WEAI is a free data retrieval call binding the contract method 0x0dc7df53.
//
// Solidity: function wEAI() view returns(address)
func (_IWorkerHub *IWorkerHubCallerSession) WEAI() (common.Address, error) {
	return _IWorkerHub.Contract.WEAI(&_IWorkerHub.CallOpts)
}

// WorkerHubScoring is a free data retrieval call binding the contract method 0x2b426301.
//
// Solidity: function workerHubScoring() view returns(address)
func (_IWorkerHub *IWorkerHubCaller) WorkerHubScoring(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IWorkerHub.contract.Call(opts, &out, "workerHubScoring")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WorkerHubScoring is a free data retrieval call binding the contract method 0x2b426301.
//
// Solidity: function workerHubScoring() view returns(address)
func (_IWorkerHub *IWorkerHubSession) WorkerHubScoring() (common.Address, error) {
	return _IWorkerHub.Contract.WorkerHubScoring(&_IWorkerHub.CallOpts)
}

// WorkerHubScoring is a free data retrieval call binding the contract method 0x2b426301.
//
// Solidity: function workerHubScoring() view returns(address)
func (_IWorkerHub *IWorkerHubCallerSession) WorkerHubScoring() (common.Address, error) {
	return _IWorkerHub.Contract.WorkerHubScoring(&_IWorkerHub.CallOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _miner) returns()
func (_IWorkerHub *IWorkerHubTransactor) ClaimReward(opts *bind.TransactOpts, _miner common.Address) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "claimReward", _miner)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _miner) returns()
func (_IWorkerHub *IWorkerHubSession) ClaimReward(_miner common.Address) (*types.Transaction, error) {
	return _IWorkerHub.Contract.ClaimReward(&_IWorkerHub.TransactOpts, _miner)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _miner) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) ClaimReward(_miner common.Address) (*types.Transaction, error) {
	return _IWorkerHub.Contract.ClaimReward(&_IWorkerHub.TransactOpts, _miner)
}

// Commit is a paid mutator transaction binding the contract method 0xf2f03877.
//
// Solidity: function commit(uint256 _assignId, bytes32 _commitment) returns()
func (_IWorkerHub *IWorkerHubTransactor) Commit(opts *bind.TransactOpts, _assignId *big.Int, _commitment [32]byte) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "commit", _assignId, _commitment)
}

// Commit is a paid mutator transaction binding the contract method 0xf2f03877.
//
// Solidity: function commit(uint256 _assignId, bytes32 _commitment) returns()
func (_IWorkerHub *IWorkerHubSession) Commit(_assignId *big.Int, _commitment [32]byte) (*types.Transaction, error) {
	return _IWorkerHub.Contract.Commit(&_IWorkerHub.TransactOpts, _assignId, _commitment)
}

// Commit is a paid mutator transaction binding the contract method 0xf2f03877.
//
// Solidity: function commit(uint256 _assignId, bytes32 _commitment) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) Commit(_assignId *big.Int, _commitment [32]byte) (*types.Transaction, error) {
	return _IWorkerHub.Contract.Commit(&_IWorkerHub.TransactOpts, _assignId, _commitment)
}

// ForceChangeModelForMiner is a paid mutator transaction binding the contract method 0x339d0f78.
//
// Solidity: function forceChangeModelForMiner(address _miner, address _modelAddress) returns()
func (_IWorkerHub *IWorkerHubTransactor) ForceChangeModelForMiner(opts *bind.TransactOpts, _miner common.Address, _modelAddress common.Address) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "forceChangeModelForMiner", _miner, _modelAddress)
}

// ForceChangeModelForMiner is a paid mutator transaction binding the contract method 0x339d0f78.
//
// Solidity: function forceChangeModelForMiner(address _miner, address _modelAddress) returns()
func (_IWorkerHub *IWorkerHubSession) ForceChangeModelForMiner(_miner common.Address, _modelAddress common.Address) (*types.Transaction, error) {
	return _IWorkerHub.Contract.ForceChangeModelForMiner(&_IWorkerHub.TransactOpts, _miner, _modelAddress)
}

// ForceChangeModelForMiner is a paid mutator transaction binding the contract method 0x339d0f78.
//
// Solidity: function forceChangeModelForMiner(address _miner, address _modelAddress) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) ForceChangeModelForMiner(_miner common.Address, _modelAddress common.Address) (*types.Transaction, error) {
	return _IWorkerHub.Contract.ForceChangeModelForMiner(&_IWorkerHub.TransactOpts, _miner, _modelAddress)
}

// IncreaseMinerStake is a paid mutator transaction binding the contract method 0xb1d1a56b.
//
// Solidity: function increaseMinerStake(uint256 wEAIAmt) returns()
func (_IWorkerHub *IWorkerHubTransactor) IncreaseMinerStake(opts *bind.TransactOpts, wEAIAmt *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "increaseMinerStake", wEAIAmt)
}

// IncreaseMinerStake is a paid mutator transaction binding the contract method 0xb1d1a56b.
//
// Solidity: function increaseMinerStake(uint256 wEAIAmt) returns()
func (_IWorkerHub *IWorkerHubSession) IncreaseMinerStake(wEAIAmt *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.IncreaseMinerStake(&_IWorkerHub.TransactOpts, wEAIAmt)
}

// IncreaseMinerStake is a paid mutator transaction binding the contract method 0xb1d1a56b.
//
// Solidity: function increaseMinerStake(uint256 wEAIAmt) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) IncreaseMinerStake(wEAIAmt *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.IncreaseMinerStake(&_IWorkerHub.TransactOpts, wEAIAmt)
}

// Infer is a paid mutator transaction binding the contract method 0xd9844458.
//
// Solidity: function infer(bytes _input, address _creator) payable returns(uint256)
func (_IWorkerHub *IWorkerHubTransactor) Infer(opts *bind.TransactOpts, _input []byte, _creator common.Address) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "infer", _input, _creator)
}

// Infer is a paid mutator transaction binding the contract method 0xd9844458.
//
// Solidity: function infer(bytes _input, address _creator) payable returns(uint256)
func (_IWorkerHub *IWorkerHubSession) Infer(_input []byte, _creator common.Address) (*types.Transaction, error) {
	return _IWorkerHub.Contract.Infer(&_IWorkerHub.TransactOpts, _input, _creator)
}

// Infer is a paid mutator transaction binding the contract method 0xd9844458.
//
// Solidity: function infer(bytes _input, address _creator) payable returns(uint256)
func (_IWorkerHub *IWorkerHubTransactorSession) Infer(_input []byte, _creator common.Address) (*types.Transaction, error) {
	return _IWorkerHub.Contract.Infer(&_IWorkerHub.TransactOpts, _input, _creator)
}

// InferWithCallback is a paid mutator transaction binding the contract method 0xb8cfec3d.
//
// Solidity: function inferWithCallback(uint256 originInferId, bytes _input, address _creator, address callback) payable returns(uint256 inferenceId)
func (_IWorkerHub *IWorkerHubTransactor) InferWithCallback(opts *bind.TransactOpts, originInferId *big.Int, _input []byte, _creator common.Address, callback common.Address) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "inferWithCallback", originInferId, _input, _creator, callback)
}

// InferWithCallback is a paid mutator transaction binding the contract method 0xb8cfec3d.
//
// Solidity: function inferWithCallback(uint256 originInferId, bytes _input, address _creator, address callback) payable returns(uint256 inferenceId)
func (_IWorkerHub *IWorkerHubSession) InferWithCallback(originInferId *big.Int, _input []byte, _creator common.Address, callback common.Address) (*types.Transaction, error) {
	return _IWorkerHub.Contract.InferWithCallback(&_IWorkerHub.TransactOpts, originInferId, _input, _creator, callback)
}

// InferWithCallback is a paid mutator transaction binding the contract method 0xb8cfec3d.
//
// Solidity: function inferWithCallback(uint256 originInferId, bytes _input, address _creator, address callback) payable returns(uint256 inferenceId)
func (_IWorkerHub *IWorkerHubTransactorSession) InferWithCallback(originInferId *big.Int, _input []byte, _creator common.Address, callback common.Address) (*types.Transaction, error) {
	return _IWorkerHub.Contract.InferWithCallback(&_IWorkerHub.TransactOpts, originInferId, _input, _creator, callback)
}

// Initialize is a paid mutator transaction binding the contract method 0xe2f32c82.
//
// Solidity: function initialize(address _l2Owner, address _treasury, address _daoToken, uint16 _feeL2Percentage, uint16 _feeTreasuryPercentage, uint256 _minerMinimumStake, uint8 _minerRequirement, uint256 _blocksPerEpoch, uint256 _rewardPerEpoch, uint256 _duration, uint16 _finePercentage, uint16 _feeRatioMinerValidor, uint256 _minFeeToUse, uint256 _daoTokenReward, (uint16,uint16,uint16,uint16,uint16) _daoTokenPercentage) returns()
func (_IWorkerHub *IWorkerHubTransactor) Initialize(opts *bind.TransactOpts, _l2Owner common.Address, _treasury common.Address, _daoToken common.Address, _feeL2Percentage uint16, _feeTreasuryPercentage uint16, _minerMinimumStake *big.Int, _minerRequirement uint8, _blocksPerEpoch *big.Int, _rewardPerEpoch *big.Int, _duration *big.Int, _finePercentage uint16, _feeRatioMinerValidor uint16, _minFeeToUse *big.Int, _daoTokenReward *big.Int, _daoTokenPercentage IWorkerHubDAOTokenPercentage) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "initialize", _l2Owner, _treasury, _daoToken, _feeL2Percentage, _feeTreasuryPercentage, _minerMinimumStake, _minerRequirement, _blocksPerEpoch, _rewardPerEpoch, _duration, _finePercentage, _feeRatioMinerValidor, _minFeeToUse, _daoTokenReward, _daoTokenPercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0xe2f32c82.
//
// Solidity: function initialize(address _l2Owner, address _treasury, address _daoToken, uint16 _feeL2Percentage, uint16 _feeTreasuryPercentage, uint256 _minerMinimumStake, uint8 _minerRequirement, uint256 _blocksPerEpoch, uint256 _rewardPerEpoch, uint256 _duration, uint16 _finePercentage, uint16 _feeRatioMinerValidor, uint256 _minFeeToUse, uint256 _daoTokenReward, (uint16,uint16,uint16,uint16,uint16) _daoTokenPercentage) returns()
func (_IWorkerHub *IWorkerHubSession) Initialize(_l2Owner common.Address, _treasury common.Address, _daoToken common.Address, _feeL2Percentage uint16, _feeTreasuryPercentage uint16, _minerMinimumStake *big.Int, _minerRequirement uint8, _blocksPerEpoch *big.Int, _rewardPerEpoch *big.Int, _duration *big.Int, _finePercentage uint16, _feeRatioMinerValidor uint16, _minFeeToUse *big.Int, _daoTokenReward *big.Int, _daoTokenPercentage IWorkerHubDAOTokenPercentage) (*types.Transaction, error) {
	return _IWorkerHub.Contract.Initialize(&_IWorkerHub.TransactOpts, _l2Owner, _treasury, _daoToken, _feeL2Percentage, _feeTreasuryPercentage, _minerMinimumStake, _minerRequirement, _blocksPerEpoch, _rewardPerEpoch, _duration, _finePercentage, _feeRatioMinerValidor, _minFeeToUse, _daoTokenReward, _daoTokenPercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0xe2f32c82.
//
// Solidity: function initialize(address _l2Owner, address _treasury, address _daoToken, uint16 _feeL2Percentage, uint16 _feeTreasuryPercentage, uint256 _minerMinimumStake, uint8 _minerRequirement, uint256 _blocksPerEpoch, uint256 _rewardPerEpoch, uint256 _duration, uint16 _finePercentage, uint16 _feeRatioMinerValidor, uint256 _minFeeToUse, uint256 _daoTokenReward, (uint16,uint16,uint16,uint16,uint16) _daoTokenPercentage) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) Initialize(_l2Owner common.Address, _treasury common.Address, _daoToken common.Address, _feeL2Percentage uint16, _feeTreasuryPercentage uint16, _minerMinimumStake *big.Int, _minerRequirement uint8, _blocksPerEpoch *big.Int, _rewardPerEpoch *big.Int, _duration *big.Int, _finePercentage uint16, _feeRatioMinerValidor uint16, _minFeeToUse *big.Int, _daoTokenReward *big.Int, _daoTokenPercentage IWorkerHubDAOTokenPercentage) (*types.Transaction, error) {
	return _IWorkerHub.Contract.Initialize(&_IWorkerHub.TransactOpts, _l2Owner, _treasury, _daoToken, _feeL2Percentage, _feeTreasuryPercentage, _minerMinimumStake, _minerRequirement, _blocksPerEpoch, _rewardPerEpoch, _duration, _finePercentage, _feeRatioMinerValidor, _minFeeToUse, _daoTokenReward, _daoTokenPercentage)
}

// JoinForMinting is a paid mutator transaction binding the contract method 0x1a8ef584.
//
// Solidity: function joinForMinting() returns()
func (_IWorkerHub *IWorkerHubTransactor) JoinForMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "joinForMinting")
}

// JoinForMinting is a paid mutator transaction binding the contract method 0x1a8ef584.
//
// Solidity: function joinForMinting() returns()
func (_IWorkerHub *IWorkerHubSession) JoinForMinting() (*types.Transaction, error) {
	return _IWorkerHub.Contract.JoinForMinting(&_IWorkerHub.TransactOpts)
}

// JoinForMinting is a paid mutator transaction binding the contract method 0x1a8ef584.
//
// Solidity: function joinForMinting() returns()
func (_IWorkerHub *IWorkerHubTransactorSession) JoinForMinting() (*types.Transaction, error) {
	return _IWorkerHub.Contract.JoinForMinting(&_IWorkerHub.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IWorkerHub *IWorkerHubTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IWorkerHub *IWorkerHubSession) Pause() (*types.Transaction, error) {
	return _IWorkerHub.Contract.Pause(&_IWorkerHub.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IWorkerHub *IWorkerHubTransactorSession) Pause() (*types.Transaction, error) {
	return _IWorkerHub.Contract.Pause(&_IWorkerHub.TransactOpts)
}

// RegisterMiner is a paid mutator transaction binding the contract method 0x668133e3.
//
// Solidity: function registerMiner(uint16 tier, uint256 wEAIAmt) returns()
func (_IWorkerHub *IWorkerHubTransactor) RegisterMiner(opts *bind.TransactOpts, tier uint16, wEAIAmt *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "registerMiner", tier, wEAIAmt)
}

// RegisterMiner is a paid mutator transaction binding the contract method 0x668133e3.
//
// Solidity: function registerMiner(uint16 tier, uint256 wEAIAmt) returns()
func (_IWorkerHub *IWorkerHubSession) RegisterMiner(tier uint16, wEAIAmt *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.RegisterMiner(&_IWorkerHub.TransactOpts, tier, wEAIAmt)
}

// RegisterMiner is a paid mutator transaction binding the contract method 0x668133e3.
//
// Solidity: function registerMiner(uint16 tier, uint256 wEAIAmt) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) RegisterMiner(tier uint16, wEAIAmt *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.RegisterMiner(&_IWorkerHub.TransactOpts, tier, wEAIAmt)
}

// RegisterModel is a paid mutator transaction binding the contract method 0xa8d6d3d1.
//
// Solidity: function registerModel(address _model, uint16 _tier, uint256 _minimumFee) returns()
func (_IWorkerHub *IWorkerHubTransactor) RegisterModel(opts *bind.TransactOpts, _model common.Address, _tier uint16, _minimumFee *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "registerModel", _model, _tier, _minimumFee)
}

// RegisterModel is a paid mutator transaction binding the contract method 0xa8d6d3d1.
//
// Solidity: function registerModel(address _model, uint16 _tier, uint256 _minimumFee) returns()
func (_IWorkerHub *IWorkerHubSession) RegisterModel(_model common.Address, _tier uint16, _minimumFee *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.RegisterModel(&_IWorkerHub.TransactOpts, _model, _tier, _minimumFee)
}

// RegisterModel is a paid mutator transaction binding the contract method 0xa8d6d3d1.
//
// Solidity: function registerModel(address _model, uint16 _tier, uint256 _minimumFee) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) RegisterModel(_model common.Address, _tier uint16, _minimumFee *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.RegisterModel(&_IWorkerHub.TransactOpts, _model, _tier, _minimumFee)
}

// RegisterReferrer is a paid mutator transaction binding the contract method 0x9ea7685a.
//
// Solidity: function registerReferrer(address _referrer) returns()
func (_IWorkerHub *IWorkerHubTransactor) RegisterReferrer(opts *bind.TransactOpts, _referrer common.Address) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "registerReferrer", _referrer)
}

// RegisterReferrer is a paid mutator transaction binding the contract method 0x9ea7685a.
//
// Solidity: function registerReferrer(address _referrer) returns()
func (_IWorkerHub *IWorkerHubSession) RegisterReferrer(_referrer common.Address) (*types.Transaction, error) {
	return _IWorkerHub.Contract.RegisterReferrer(&_IWorkerHub.TransactOpts, _referrer)
}

// RegisterReferrer is a paid mutator transaction binding the contract method 0x9ea7685a.
//
// Solidity: function registerReferrer(address _referrer) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) RegisterReferrer(_referrer common.Address) (*types.Transaction, error) {
	return _IWorkerHub.Contract.RegisterReferrer(&_IWorkerHub.TransactOpts, _referrer)
}

// RegisterReferrer0 is a paid mutator transaction binding the contract method 0xc41bf665.
//
// Solidity: function registerReferrer(address[] _referrers, address[] _referees) returns()
func (_IWorkerHub *IWorkerHubTransactor) RegisterReferrer0(opts *bind.TransactOpts, _referrers []common.Address, _referees []common.Address) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "registerReferrer0", _referrers, _referees)
}

// RegisterReferrer0 is a paid mutator transaction binding the contract method 0xc41bf665.
//
// Solidity: function registerReferrer(address[] _referrers, address[] _referees) returns()
func (_IWorkerHub *IWorkerHubSession) RegisterReferrer0(_referrers []common.Address, _referees []common.Address) (*types.Transaction, error) {
	return _IWorkerHub.Contract.RegisterReferrer0(&_IWorkerHub.TransactOpts, _referrers, _referees)
}

// RegisterReferrer0 is a paid mutator transaction binding the contract method 0xc41bf665.
//
// Solidity: function registerReferrer(address[] _referrers, address[] _referees) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) RegisterReferrer0(_referrers []common.Address, _referees []common.Address) (*types.Transaction, error) {
	return _IWorkerHub.Contract.RegisterReferrer0(&_IWorkerHub.TransactOpts, _referrers, _referees)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IWorkerHub *IWorkerHubTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IWorkerHub *IWorkerHubSession) RenounceOwnership() (*types.Transaction, error) {
	return _IWorkerHub.Contract.RenounceOwnership(&_IWorkerHub.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IWorkerHub *IWorkerHubTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _IWorkerHub.Contract.RenounceOwnership(&_IWorkerHub.TransactOpts)
}

// ResolveInference is a paid mutator transaction binding the contract method 0x6029e786.
//
// Solidity: function resolveInference(uint256 _inferenceId) returns()
func (_IWorkerHub *IWorkerHubTransactor) ResolveInference(opts *bind.TransactOpts, _inferenceId *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "resolveInference", _inferenceId)
}

// ResolveInference is a paid mutator transaction binding the contract method 0x6029e786.
//
// Solidity: function resolveInference(uint256 _inferenceId) returns()
func (_IWorkerHub *IWorkerHubSession) ResolveInference(_inferenceId *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.ResolveInference(&_IWorkerHub.TransactOpts, _inferenceId)
}

// ResolveInference is a paid mutator transaction binding the contract method 0x6029e786.
//
// Solidity: function resolveInference(uint256 _inferenceId) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) ResolveInference(_inferenceId *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.ResolveInference(&_IWorkerHub.TransactOpts, _inferenceId)
}

// RestakeForMiner is a paid mutator transaction binding the contract method 0x4fb9bc1e.
//
// Solidity: function restakeForMiner(uint16 tier) returns()
func (_IWorkerHub *IWorkerHubTransactor) RestakeForMiner(opts *bind.TransactOpts, tier uint16) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "restakeForMiner", tier)
}

// RestakeForMiner is a paid mutator transaction binding the contract method 0x4fb9bc1e.
//
// Solidity: function restakeForMiner(uint16 tier) returns()
func (_IWorkerHub *IWorkerHubSession) RestakeForMiner(tier uint16) (*types.Transaction, error) {
	return _IWorkerHub.Contract.RestakeForMiner(&_IWorkerHub.TransactOpts, tier)
}

// RestakeForMiner is a paid mutator transaction binding the contract method 0x4fb9bc1e.
//
// Solidity: function restakeForMiner(uint16 tier) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) RestakeForMiner(tier uint16) (*types.Transaction, error) {
	return _IWorkerHub.Contract.RestakeForMiner(&_IWorkerHub.TransactOpts, tier)
}

// ResultReceived is a paid mutator transaction binding the contract method 0xc3477018.
//
// Solidity: function resultReceived(bytes result) returns()
func (_IWorkerHub *IWorkerHubTransactor) ResultReceived(opts *bind.TransactOpts, result []byte) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "resultReceived", result)
}

// ResultReceived is a paid mutator transaction binding the contract method 0xc3477018.
//
// Solidity: function resultReceived(bytes result) returns()
func (_IWorkerHub *IWorkerHubSession) ResultReceived(result []byte) (*types.Transaction, error) {
	return _IWorkerHub.Contract.ResultReceived(&_IWorkerHub.TransactOpts, result)
}

// ResultReceived is a paid mutator transaction binding the contract method 0xc3477018.
//
// Solidity: function resultReceived(bytes result) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) ResultReceived(result []byte) (*types.Transaction, error) {
	return _IWorkerHub.Contract.ResultReceived(&_IWorkerHub.TransactOpts, result)
}

// ResultReceived0 is a paid mutator transaction binding the contract method 0xd2a554e7.
//
// Solidity: function resultReceived(uint256 _originInferId, bytes _result) returns()
func (_IWorkerHub *IWorkerHubTransactor) ResultReceived0(opts *bind.TransactOpts, _originInferId *big.Int, _result []byte) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "resultReceived0", _originInferId, _result)
}

// ResultReceived0 is a paid mutator transaction binding the contract method 0xd2a554e7.
//
// Solidity: function resultReceived(uint256 _originInferId, bytes _result) returns()
func (_IWorkerHub *IWorkerHubSession) ResultReceived0(_originInferId *big.Int, _result []byte) (*types.Transaction, error) {
	return _IWorkerHub.Contract.ResultReceived0(&_IWorkerHub.TransactOpts, _originInferId, _result)
}

// ResultReceived0 is a paid mutator transaction binding the contract method 0xd2a554e7.
//
// Solidity: function resultReceived(uint256 _originInferId, bytes _result) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) ResultReceived0(_originInferId *big.Int, _result []byte) (*types.Transaction, error) {
	return _IWorkerHub.Contract.ResultReceived0(&_IWorkerHub.TransactOpts, _originInferId, _result)
}

// Reveal is a paid mutator transaction binding the contract method 0x121a301d.
//
// Solidity: function reveal(uint256 _assignId, uint40 _nonce, bytes _data) returns()
func (_IWorkerHub *IWorkerHubTransactor) Reveal(opts *bind.TransactOpts, _assignId *big.Int, _nonce *big.Int, _data []byte) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "reveal", _assignId, _nonce, _data)
}

// Reveal is a paid mutator transaction binding the contract method 0x121a301d.
//
// Solidity: function reveal(uint256 _assignId, uint40 _nonce, bytes _data) returns()
func (_IWorkerHub *IWorkerHubSession) Reveal(_assignId *big.Int, _nonce *big.Int, _data []byte) (*types.Transaction, error) {
	return _IWorkerHub.Contract.Reveal(&_IWorkerHub.TransactOpts, _assignId, _nonce, _data)
}

// Reveal is a paid mutator transaction binding the contract method 0x121a301d.
//
// Solidity: function reveal(uint256 _assignId, uint40 _nonce, bytes _data) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) Reveal(_assignId *big.Int, _nonce *big.Int, _data []byte) (*types.Transaction, error) {
	return _IWorkerHub.Contract.Reveal(&_IWorkerHub.TransactOpts, _assignId, _nonce, _data)
}

// RewardToClaim is a paid mutator transaction binding the contract method 0x674a63b9.
//
// Solidity: function rewardToClaim(address _miner) returns(uint256)
func (_IWorkerHub *IWorkerHubTransactor) RewardToClaim(opts *bind.TransactOpts, _miner common.Address) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "rewardToClaim", _miner)
}

// RewardToClaim is a paid mutator transaction binding the contract method 0x674a63b9.
//
// Solidity: function rewardToClaim(address _miner) returns(uint256)
func (_IWorkerHub *IWorkerHubSession) RewardToClaim(_miner common.Address) (*types.Transaction, error) {
	return _IWorkerHub.Contract.RewardToClaim(&_IWorkerHub.TransactOpts, _miner)
}

// RewardToClaim is a paid mutator transaction binding the contract method 0x674a63b9.
//
// Solidity: function rewardToClaim(address _miner) returns(uint256)
func (_IWorkerHub *IWorkerHubTransactorSession) RewardToClaim(_miner common.Address) (*types.Transaction, error) {
	return _IWorkerHub.Contract.RewardToClaim(&_IWorkerHub.TransactOpts, _miner)
}

// SeizeMinerRole is a paid mutator transaction binding the contract method 0xffbc6661.
//
// Solidity: function seizeMinerRole(uint256 _assignmentId) returns()
func (_IWorkerHub *IWorkerHubTransactor) SeizeMinerRole(opts *bind.TransactOpts, _assignmentId *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "seizeMinerRole", _assignmentId)
}

// SeizeMinerRole is a paid mutator transaction binding the contract method 0xffbc6661.
//
// Solidity: function seizeMinerRole(uint256 _assignmentId) returns()
func (_IWorkerHub *IWorkerHubSession) SeizeMinerRole(_assignmentId *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SeizeMinerRole(&_IWorkerHub.TransactOpts, _assignmentId)
}

// SeizeMinerRole is a paid mutator transaction binding the contract method 0xffbc6661.
//
// Solidity: function seizeMinerRole(uint256 _assignmentId) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) SeizeMinerRole(_assignmentId *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SeizeMinerRole(&_IWorkerHub.TransactOpts, _assignmentId)
}

// SetBlocksPerEpoch is a paid mutator transaction binding the contract method 0x034438b0.
//
// Solidity: function setBlocksPerEpoch(uint256 _blocks) returns()
func (_IWorkerHub *IWorkerHubTransactor) SetBlocksPerEpoch(opts *bind.TransactOpts, _blocks *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "setBlocksPerEpoch", _blocks)
}

// SetBlocksPerEpoch is a paid mutator transaction binding the contract method 0x034438b0.
//
// Solidity: function setBlocksPerEpoch(uint256 _blocks) returns()
func (_IWorkerHub *IWorkerHubSession) SetBlocksPerEpoch(_blocks *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetBlocksPerEpoch(&_IWorkerHub.TransactOpts, _blocks)
}

// SetBlocksPerEpoch is a paid mutator transaction binding the contract method 0x034438b0.
//
// Solidity: function setBlocksPerEpoch(uint256 _blocks) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) SetBlocksPerEpoch(_blocks *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetBlocksPerEpoch(&_IWorkerHub.TransactOpts, _blocks)
}

// SetCommitDuration is a paid mutator transaction binding the contract method 0x54b18651.
//
// Solidity: function setCommitDuration(uint40 _newCommitDuration) returns()
func (_IWorkerHub *IWorkerHubTransactor) SetCommitDuration(opts *bind.TransactOpts, _newCommitDuration *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "setCommitDuration", _newCommitDuration)
}

// SetCommitDuration is a paid mutator transaction binding the contract method 0x54b18651.
//
// Solidity: function setCommitDuration(uint40 _newCommitDuration) returns()
func (_IWorkerHub *IWorkerHubSession) SetCommitDuration(_newCommitDuration *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetCommitDuration(&_IWorkerHub.TransactOpts, _newCommitDuration)
}

// SetCommitDuration is a paid mutator transaction binding the contract method 0x54b18651.
//
// Solidity: function setCommitDuration(uint40 _newCommitDuration) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) SetCommitDuration(_newCommitDuration *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetCommitDuration(&_IWorkerHub.TransactOpts, _newCommitDuration)
}

// SetDAOToken is a paid mutator transaction binding the contract method 0x70a52354.
//
// Solidity: function setDAOToken(address _daoTokenAddress) returns()
func (_IWorkerHub *IWorkerHubTransactor) SetDAOToken(opts *bind.TransactOpts, _daoTokenAddress common.Address) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "setDAOToken", _daoTokenAddress)
}

// SetDAOToken is a paid mutator transaction binding the contract method 0x70a52354.
//
// Solidity: function setDAOToken(address _daoTokenAddress) returns()
func (_IWorkerHub *IWorkerHubSession) SetDAOToken(_daoTokenAddress common.Address) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetDAOToken(&_IWorkerHub.TransactOpts, _daoTokenAddress)
}

// SetDAOToken is a paid mutator transaction binding the contract method 0x70a52354.
//
// Solidity: function setDAOToken(address _daoTokenAddress) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) SetDAOToken(_daoTokenAddress common.Address) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetDAOToken(&_IWorkerHub.TransactOpts, _daoTokenAddress)
}

// SetDAOTokenPercentage is a paid mutator transaction binding the contract method 0x3860ce68.
//
// Solidity: function setDAOTokenPercentage((uint16,uint16,uint16,uint16,uint16) _daoTokenPercentage) returns()
func (_IWorkerHub *IWorkerHubTransactor) SetDAOTokenPercentage(opts *bind.TransactOpts, _daoTokenPercentage IWorkerHubDAOTokenPercentage) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "setDAOTokenPercentage", _daoTokenPercentage)
}

// SetDAOTokenPercentage is a paid mutator transaction binding the contract method 0x3860ce68.
//
// Solidity: function setDAOTokenPercentage((uint16,uint16,uint16,uint16,uint16) _daoTokenPercentage) returns()
func (_IWorkerHub *IWorkerHubSession) SetDAOTokenPercentage(_daoTokenPercentage IWorkerHubDAOTokenPercentage) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetDAOTokenPercentage(&_IWorkerHub.TransactOpts, _daoTokenPercentage)
}

// SetDAOTokenPercentage is a paid mutator transaction binding the contract method 0x3860ce68.
//
// Solidity: function setDAOTokenPercentage((uint16,uint16,uint16,uint16,uint16) _daoTokenPercentage) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) SetDAOTokenPercentage(_daoTokenPercentage IWorkerHubDAOTokenPercentage) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetDAOTokenPercentage(&_IWorkerHub.TransactOpts, _daoTokenPercentage)
}

// SetDAOTokenReward is a paid mutator transaction binding the contract method 0x76e7ffae.
//
// Solidity: function setDAOTokenReward(uint256 _newDAOTokenReward) returns()
func (_IWorkerHub *IWorkerHubTransactor) SetDAOTokenReward(opts *bind.TransactOpts, _newDAOTokenReward *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "setDAOTokenReward", _newDAOTokenReward)
}

// SetDAOTokenReward is a paid mutator transaction binding the contract method 0x76e7ffae.
//
// Solidity: function setDAOTokenReward(uint256 _newDAOTokenReward) returns()
func (_IWorkerHub *IWorkerHubSession) SetDAOTokenReward(_newDAOTokenReward *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetDAOTokenReward(&_IWorkerHub.TransactOpts, _newDAOTokenReward)
}

// SetDAOTokenReward is a paid mutator transaction binding the contract method 0x76e7ffae.
//
// Solidity: function setDAOTokenReward(uint256 _newDAOTokenReward) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) SetDAOTokenReward(_newDAOTokenReward *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetDAOTokenReward(&_IWorkerHub.TransactOpts, _newDAOTokenReward)
}

// SetFeeRatioMinerValidator is a paid mutator transaction binding the contract method 0xafa82609.
//
// Solidity: function setFeeRatioMinerValidator(uint16 _newRatio) returns()
func (_IWorkerHub *IWorkerHubTransactor) SetFeeRatioMinerValidator(opts *bind.TransactOpts, _newRatio uint16) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "setFeeRatioMinerValidator", _newRatio)
}

// SetFeeRatioMinerValidator is a paid mutator transaction binding the contract method 0xafa82609.
//
// Solidity: function setFeeRatioMinerValidator(uint16 _newRatio) returns()
func (_IWorkerHub *IWorkerHubSession) SetFeeRatioMinerValidator(_newRatio uint16) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetFeeRatioMinerValidator(&_IWorkerHub.TransactOpts, _newRatio)
}

// SetFeeRatioMinerValidator is a paid mutator transaction binding the contract method 0xafa82609.
//
// Solidity: function setFeeRatioMinerValidator(uint16 _newRatio) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) SetFeeRatioMinerValidator(_newRatio uint16) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetFeeRatioMinerValidator(&_IWorkerHub.TransactOpts, _newRatio)
}

// SetFinePercentage is a paid mutator transaction binding the contract method 0x431a4457.
//
// Solidity: function setFinePercentage(uint16 _finePercentage) returns()
func (_IWorkerHub *IWorkerHubTransactor) SetFinePercentage(opts *bind.TransactOpts, _finePercentage uint16) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "setFinePercentage", _finePercentage)
}

// SetFinePercentage is a paid mutator transaction binding the contract method 0x431a4457.
//
// Solidity: function setFinePercentage(uint16 _finePercentage) returns()
func (_IWorkerHub *IWorkerHubSession) SetFinePercentage(_finePercentage uint16) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetFinePercentage(&_IWorkerHub.TransactOpts, _finePercentage)
}

// SetFinePercentage is a paid mutator transaction binding the contract method 0x431a4457.
//
// Solidity: function setFinePercentage(uint16 _finePercentage) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) SetFinePercentage(_finePercentage uint16) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetFinePercentage(&_IWorkerHub.TransactOpts, _finePercentage)
}

// SetL2Owner is a paid mutator transaction binding the contract method 0xb530c110.
//
// Solidity: function setL2Owner(address _l2OwnerAddress) returns()
func (_IWorkerHub *IWorkerHubTransactor) SetL2Owner(opts *bind.TransactOpts, _l2OwnerAddress common.Address) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "setL2Owner", _l2OwnerAddress)
}

// SetL2Owner is a paid mutator transaction binding the contract method 0xb530c110.
//
// Solidity: function setL2Owner(address _l2OwnerAddress) returns()
func (_IWorkerHub *IWorkerHubSession) SetL2Owner(_l2OwnerAddress common.Address) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetL2Owner(&_IWorkerHub.TransactOpts, _l2OwnerAddress)
}

// SetL2Owner is a paid mutator transaction binding the contract method 0xb530c110.
//
// Solidity: function setL2Owner(address _l2OwnerAddress) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) SetL2Owner(_l2OwnerAddress common.Address) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetL2Owner(&_IWorkerHub.TransactOpts, _l2OwnerAddress)
}

// SetMinFeeToUse is a paid mutator transaction binding the contract method 0xaf5e3be0.
//
// Solidity: function setMinFeeToUse(uint256 _minFeeToUse) returns()
func (_IWorkerHub *IWorkerHubTransactor) SetMinFeeToUse(opts *bind.TransactOpts, _minFeeToUse *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "setMinFeeToUse", _minFeeToUse)
}

// SetMinFeeToUse is a paid mutator transaction binding the contract method 0xaf5e3be0.
//
// Solidity: function setMinFeeToUse(uint256 _minFeeToUse) returns()
func (_IWorkerHub *IWorkerHubSession) SetMinFeeToUse(_minFeeToUse *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetMinFeeToUse(&_IWorkerHub.TransactOpts, _minFeeToUse)
}

// SetMinFeeToUse is a paid mutator transaction binding the contract method 0xaf5e3be0.
//
// Solidity: function setMinFeeToUse(uint256 _minFeeToUse) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) SetMinFeeToUse(_minFeeToUse *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetMinFeeToUse(&_IWorkerHub.TransactOpts, _minFeeToUse)
}

// SetMinerMinimumStake is a paid mutator transaction binding the contract method 0xe69d5b98.
//
// Solidity: function setMinerMinimumStake(uint256 _minerMinimumStake) returns()
func (_IWorkerHub *IWorkerHubTransactor) SetMinerMinimumStake(opts *bind.TransactOpts, _minerMinimumStake *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "setMinerMinimumStake", _minerMinimumStake)
}

// SetMinerMinimumStake is a paid mutator transaction binding the contract method 0xe69d5b98.
//
// Solidity: function setMinerMinimumStake(uint256 _minerMinimumStake) returns()
func (_IWorkerHub *IWorkerHubSession) SetMinerMinimumStake(_minerMinimumStake *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetMinerMinimumStake(&_IWorkerHub.TransactOpts, _minerMinimumStake)
}

// SetMinerMinimumStake is a paid mutator transaction binding the contract method 0xe69d5b98.
//
// Solidity: function setMinerMinimumStake(uint256 _minerMinimumStake) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) SetMinerMinimumStake(_minerMinimumStake *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetMinerMinimumStake(&_IWorkerHub.TransactOpts, _minerMinimumStake)
}

// SetNewRewardInEpoch is a paid mutator transaction binding the contract method 0xe32bd90c.
//
// Solidity: function setNewRewardInEpoch(uint256 _newRewardAmount) returns()
func (_IWorkerHub *IWorkerHubTransactor) SetNewRewardInEpoch(opts *bind.TransactOpts, _newRewardAmount *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "setNewRewardInEpoch", _newRewardAmount)
}

// SetNewRewardInEpoch is a paid mutator transaction binding the contract method 0xe32bd90c.
//
// Solidity: function setNewRewardInEpoch(uint256 _newRewardAmount) returns()
func (_IWorkerHub *IWorkerHubSession) SetNewRewardInEpoch(_newRewardAmount *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetNewRewardInEpoch(&_IWorkerHub.TransactOpts, _newRewardAmount)
}

// SetNewRewardInEpoch is a paid mutator transaction binding the contract method 0xe32bd90c.
//
// Solidity: function setNewRewardInEpoch(uint256 _newRewardAmount) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) SetNewRewardInEpoch(_newRewardAmount *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetNewRewardInEpoch(&_IWorkerHub.TransactOpts, _newRewardAmount)
}

// SetPenaltyDuration is a paid mutator transaction binding the contract method 0x885b050f.
//
// Solidity: function setPenaltyDuration(uint40 _penaltyDuration) returns()
func (_IWorkerHub *IWorkerHubTransactor) SetPenaltyDuration(opts *bind.TransactOpts, _penaltyDuration *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "setPenaltyDuration", _penaltyDuration)
}

// SetPenaltyDuration is a paid mutator transaction binding the contract method 0x885b050f.
//
// Solidity: function setPenaltyDuration(uint40 _penaltyDuration) returns()
func (_IWorkerHub *IWorkerHubSession) SetPenaltyDuration(_penaltyDuration *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetPenaltyDuration(&_IWorkerHub.TransactOpts, _penaltyDuration)
}

// SetPenaltyDuration is a paid mutator transaction binding the contract method 0x885b050f.
//
// Solidity: function setPenaltyDuration(uint40 _penaltyDuration) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) SetPenaltyDuration(_penaltyDuration *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetPenaltyDuration(&_IWorkerHub.TransactOpts, _penaltyDuration)
}

// SetRevealDuration is a paid mutator transaction binding the contract method 0x1eb9a99a.
//
// Solidity: function setRevealDuration(uint40 _newRevealDuration) returns()
func (_IWorkerHub *IWorkerHubTransactor) SetRevealDuration(opts *bind.TransactOpts, _newRevealDuration *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "setRevealDuration", _newRevealDuration)
}

// SetRevealDuration is a paid mutator transaction binding the contract method 0x1eb9a99a.
//
// Solidity: function setRevealDuration(uint40 _newRevealDuration) returns()
func (_IWorkerHub *IWorkerHubSession) SetRevealDuration(_newRevealDuration *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetRevealDuration(&_IWorkerHub.TransactOpts, _newRevealDuration)
}

// SetRevealDuration is a paid mutator transaction binding the contract method 0x1eb9a99a.
//
// Solidity: function setRevealDuration(uint40 _newRevealDuration) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) SetRevealDuration(_newRevealDuration *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetRevealDuration(&_IWorkerHub.TransactOpts, _newRevealDuration)
}

// SetScoringInfo is a paid mutator transaction binding the contract method 0x0d425ea5.
//
// Solidity: function setScoringInfo(address _workerHubScoring, address _modelScoring) returns()
func (_IWorkerHub *IWorkerHubTransactor) SetScoringInfo(opts *bind.TransactOpts, _workerHubScoring common.Address, _modelScoring common.Address) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "setScoringInfo", _workerHubScoring, _modelScoring)
}

// SetScoringInfo is a paid mutator transaction binding the contract method 0x0d425ea5.
//
// Solidity: function setScoringInfo(address _workerHubScoring, address _modelScoring) returns()
func (_IWorkerHub *IWorkerHubSession) SetScoringInfo(_workerHubScoring common.Address, _modelScoring common.Address) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetScoringInfo(&_IWorkerHub.TransactOpts, _workerHubScoring, _modelScoring)
}

// SetScoringInfo is a paid mutator transaction binding the contract method 0x0d425ea5.
//
// Solidity: function setScoringInfo(address _workerHubScoring, address _modelScoring) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) SetScoringInfo(_workerHubScoring common.Address, _modelScoring common.Address) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetScoringInfo(&_IWorkerHub.TransactOpts, _workerHubScoring, _modelScoring)
}

// SetSubmitDuration is a paid mutator transaction binding the contract method 0x6f643736.
//
// Solidity: function setSubmitDuration(uint40 _newSubmitDuration) returns()
func (_IWorkerHub *IWorkerHubTransactor) SetSubmitDuration(opts *bind.TransactOpts, _newSubmitDuration *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "setSubmitDuration", _newSubmitDuration)
}

// SetSubmitDuration is a paid mutator transaction binding the contract method 0x6f643736.
//
// Solidity: function setSubmitDuration(uint40 _newSubmitDuration) returns()
func (_IWorkerHub *IWorkerHubSession) SetSubmitDuration(_newSubmitDuration *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetSubmitDuration(&_IWorkerHub.TransactOpts, _newSubmitDuration)
}

// SetSubmitDuration is a paid mutator transaction binding the contract method 0x6f643736.
//
// Solidity: function setSubmitDuration(uint40 _newSubmitDuration) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) SetSubmitDuration(_newSubmitDuration *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetSubmitDuration(&_IWorkerHub.TransactOpts, _newSubmitDuration)
}

// SetTreasuryAddress is a paid mutator transaction binding the contract method 0x6605bfda.
//
// Solidity: function setTreasuryAddress(address _treasuryAddress) returns()
func (_IWorkerHub *IWorkerHubTransactor) SetTreasuryAddress(opts *bind.TransactOpts, _treasuryAddress common.Address) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "setTreasuryAddress", _treasuryAddress)
}

// SetTreasuryAddress is a paid mutator transaction binding the contract method 0x6605bfda.
//
// Solidity: function setTreasuryAddress(address _treasuryAddress) returns()
func (_IWorkerHub *IWorkerHubSession) SetTreasuryAddress(_treasuryAddress common.Address) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetTreasuryAddress(&_IWorkerHub.TransactOpts, _treasuryAddress)
}

// SetTreasuryAddress is a paid mutator transaction binding the contract method 0x6605bfda.
//
// Solidity: function setTreasuryAddress(address _treasuryAddress) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) SetTreasuryAddress(_treasuryAddress common.Address) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetTreasuryAddress(&_IWorkerHub.TransactOpts, _treasuryAddress)
}

// SetUnstakDelayTime is a paid mutator transaction binding the contract method 0x351b2b33.
//
// Solidity: function setUnstakDelayTime(uint40 _newUnstakeDelayTime) returns()
func (_IWorkerHub *IWorkerHubTransactor) SetUnstakDelayTime(opts *bind.TransactOpts, _newUnstakeDelayTime *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "setUnstakDelayTime", _newUnstakeDelayTime)
}

// SetUnstakDelayTime is a paid mutator transaction binding the contract method 0x351b2b33.
//
// Solidity: function setUnstakDelayTime(uint40 _newUnstakeDelayTime) returns()
func (_IWorkerHub *IWorkerHubSession) SetUnstakDelayTime(_newUnstakeDelayTime *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetUnstakDelayTime(&_IWorkerHub.TransactOpts, _newUnstakeDelayTime)
}

// SetUnstakDelayTime is a paid mutator transaction binding the contract method 0x351b2b33.
//
// Solidity: function setUnstakDelayTime(uint40 _newUnstakeDelayTime) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) SetUnstakDelayTime(_newUnstakeDelayTime *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SetUnstakDelayTime(&_IWorkerHub.TransactOpts, _newUnstakeDelayTime)
}

// SlashMiner is a paid mutator transaction binding the contract method 0x969ceab4.
//
// Solidity: function slashMiner(address _miner, bool _isFined) returns()
func (_IWorkerHub *IWorkerHubTransactor) SlashMiner(opts *bind.TransactOpts, _miner common.Address, _isFined bool) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "slashMiner", _miner, _isFined)
}

// SlashMiner is a paid mutator transaction binding the contract method 0x969ceab4.
//
// Solidity: function slashMiner(address _miner, bool _isFined) returns()
func (_IWorkerHub *IWorkerHubSession) SlashMiner(_miner common.Address, _isFined bool) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SlashMiner(&_IWorkerHub.TransactOpts, _miner, _isFined)
}

// SlashMiner is a paid mutator transaction binding the contract method 0x969ceab4.
//
// Solidity: function slashMiner(address _miner, bool _isFined) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) SlashMiner(_miner common.Address, _isFined bool) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SlashMiner(&_IWorkerHub.TransactOpts, _miner, _isFined)
}

// StreamData is a paid mutator transaction binding the contract method 0x020e3011.
//
// Solidity: function streamData(uint256 _assignmentId, bytes _data) returns()
func (_IWorkerHub *IWorkerHubTransactor) StreamData(opts *bind.TransactOpts, _assignmentId *big.Int, _data []byte) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "streamData", _assignmentId, _data)
}

// StreamData is a paid mutator transaction binding the contract method 0x020e3011.
//
// Solidity: function streamData(uint256 _assignmentId, bytes _data) returns()
func (_IWorkerHub *IWorkerHubSession) StreamData(_assignmentId *big.Int, _data []byte) (*types.Transaction, error) {
	return _IWorkerHub.Contract.StreamData(&_IWorkerHub.TransactOpts, _assignmentId, _data)
}

// StreamData is a paid mutator transaction binding the contract method 0x020e3011.
//
// Solidity: function streamData(uint256 _assignmentId, bytes _data) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) StreamData(_assignmentId *big.Int, _data []byte) (*types.Transaction, error) {
	return _IWorkerHub.Contract.StreamData(&_IWorkerHub.TransactOpts, _assignmentId, _data)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0xe84dee6b.
//
// Solidity: function submitSolution(uint256 _assigmentId, bytes _data) returns()
func (_IWorkerHub *IWorkerHubTransactor) SubmitSolution(opts *bind.TransactOpts, _assigmentId *big.Int, _data []byte) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "submitSolution", _assigmentId, _data)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0xe84dee6b.
//
// Solidity: function submitSolution(uint256 _assigmentId, bytes _data) returns()
func (_IWorkerHub *IWorkerHubSession) SubmitSolution(_assigmentId *big.Int, _data []byte) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SubmitSolution(&_IWorkerHub.TransactOpts, _assigmentId, _data)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0xe84dee6b.
//
// Solidity: function submitSolution(uint256 _assigmentId, bytes _data) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) SubmitSolution(_assigmentId *big.Int, _data []byte) (*types.Transaction, error) {
	return _IWorkerHub.Contract.SubmitSolution(&_IWorkerHub.TransactOpts, _assigmentId, _data)
}

// TopUpInfer is a paid mutator transaction binding the contract method 0xe9bd0e26.
//
// Solidity: function topUpInfer(uint256 _inferenceId) payable returns()
func (_IWorkerHub *IWorkerHubTransactor) TopUpInfer(opts *bind.TransactOpts, _inferenceId *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "topUpInfer", _inferenceId)
}

// TopUpInfer is a paid mutator transaction binding the contract method 0xe9bd0e26.
//
// Solidity: function topUpInfer(uint256 _inferenceId) payable returns()
func (_IWorkerHub *IWorkerHubSession) TopUpInfer(_inferenceId *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.TopUpInfer(&_IWorkerHub.TransactOpts, _inferenceId)
}

// TopUpInfer is a paid mutator transaction binding the contract method 0xe9bd0e26.
//
// Solidity: function topUpInfer(uint256 _inferenceId) payable returns()
func (_IWorkerHub *IWorkerHubTransactorSession) TopUpInfer(_inferenceId *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.TopUpInfer(&_IWorkerHub.TransactOpts, _inferenceId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IWorkerHub *IWorkerHubTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IWorkerHub *IWorkerHubSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IWorkerHub.Contract.TransferOwnership(&_IWorkerHub.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IWorkerHub.Contract.TransferOwnership(&_IWorkerHub.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IWorkerHub *IWorkerHubTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IWorkerHub *IWorkerHubSession) Unpause() (*types.Transaction, error) {
	return _IWorkerHub.Contract.Unpause(&_IWorkerHub.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IWorkerHub *IWorkerHubTransactorSession) Unpause() (*types.Transaction, error) {
	return _IWorkerHub.Contract.Unpause(&_IWorkerHub.TransactOpts)
}

// UnregisterMiner is a paid mutator transaction binding the contract method 0x656a1b20.
//
// Solidity: function unregisterMiner() returns()
func (_IWorkerHub *IWorkerHubTransactor) UnregisterMiner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "unregisterMiner")
}

// UnregisterMiner is a paid mutator transaction binding the contract method 0x656a1b20.
//
// Solidity: function unregisterMiner() returns()
func (_IWorkerHub *IWorkerHubSession) UnregisterMiner() (*types.Transaction, error) {
	return _IWorkerHub.Contract.UnregisterMiner(&_IWorkerHub.TransactOpts)
}

// UnregisterMiner is a paid mutator transaction binding the contract method 0x656a1b20.
//
// Solidity: function unregisterMiner() returns()
func (_IWorkerHub *IWorkerHubTransactorSession) UnregisterMiner() (*types.Transaction, error) {
	return _IWorkerHub.Contract.UnregisterMiner(&_IWorkerHub.TransactOpts)
}

// UnregisterModel is a paid mutator transaction binding the contract method 0xdb2dab1d.
//
// Solidity: function unregisterModel(address _model) returns()
func (_IWorkerHub *IWorkerHubTransactor) UnregisterModel(opts *bind.TransactOpts, _model common.Address) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "unregisterModel", _model)
}

// UnregisterModel is a paid mutator transaction binding the contract method 0xdb2dab1d.
//
// Solidity: function unregisterModel(address _model) returns()
func (_IWorkerHub *IWorkerHubSession) UnregisterModel(_model common.Address) (*types.Transaction, error) {
	return _IWorkerHub.Contract.UnregisterModel(&_IWorkerHub.TransactOpts, _model)
}

// UnregisterModel is a paid mutator transaction binding the contract method 0xdb2dab1d.
//
// Solidity: function unregisterModel(address _model) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) UnregisterModel(_model common.Address) (*types.Transaction, error) {
	return _IWorkerHub.Contract.UnregisterModel(&_IWorkerHub.TransactOpts, _model)
}

// UnstakeForMiner is a paid mutator transaction binding the contract method 0x73df250d.
//
// Solidity: function unstakeForMiner() returns()
func (_IWorkerHub *IWorkerHubTransactor) UnstakeForMiner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "unstakeForMiner")
}

// UnstakeForMiner is a paid mutator transaction binding the contract method 0x73df250d.
//
// Solidity: function unstakeForMiner() returns()
func (_IWorkerHub *IWorkerHubSession) UnstakeForMiner() (*types.Transaction, error) {
	return _IWorkerHub.Contract.UnstakeForMiner(&_IWorkerHub.TransactOpts)
}

// UnstakeForMiner is a paid mutator transaction binding the contract method 0x73df250d.
//
// Solidity: function unstakeForMiner() returns()
func (_IWorkerHub *IWorkerHubTransactorSession) UnstakeForMiner() (*types.Transaction, error) {
	return _IWorkerHub.Contract.UnstakeForMiner(&_IWorkerHub.TransactOpts)
}

// UpdateModelMinimumFee is a paid mutator transaction binding the contract method 0xb74cd194.
//
// Solidity: function updateModelMinimumFee(address _model, uint256 _minimumFee) returns()
func (_IWorkerHub *IWorkerHubTransactor) UpdateModelMinimumFee(opts *bind.TransactOpts, _model common.Address, _minimumFee *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "updateModelMinimumFee", _model, _minimumFee)
}

// UpdateModelMinimumFee is a paid mutator transaction binding the contract method 0xb74cd194.
//
// Solidity: function updateModelMinimumFee(address _model, uint256 _minimumFee) returns()
func (_IWorkerHub *IWorkerHubSession) UpdateModelMinimumFee(_model common.Address, _minimumFee *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.UpdateModelMinimumFee(&_IWorkerHub.TransactOpts, _model, _minimumFee)
}

// UpdateModelMinimumFee is a paid mutator transaction binding the contract method 0xb74cd194.
//
// Solidity: function updateModelMinimumFee(address _model, uint256 _minimumFee) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) UpdateModelMinimumFee(_model common.Address, _minimumFee *big.Int) (*types.Transaction, error) {
	return _IWorkerHub.Contract.UpdateModelMinimumFee(&_IWorkerHub.TransactOpts, _model, _minimumFee)
}

// UpdateModelTier is a paid mutator transaction binding the contract method 0x0738a9bb.
//
// Solidity: function updateModelTier(address _model, uint32 _tier) returns()
func (_IWorkerHub *IWorkerHubTransactor) UpdateModelTier(opts *bind.TransactOpts, _model common.Address, _tier uint32) (*types.Transaction, error) {
	return _IWorkerHub.contract.Transact(opts, "updateModelTier", _model, _tier)
}

// UpdateModelTier is a paid mutator transaction binding the contract method 0x0738a9bb.
//
// Solidity: function updateModelTier(address _model, uint32 _tier) returns()
func (_IWorkerHub *IWorkerHubSession) UpdateModelTier(_model common.Address, _tier uint32) (*types.Transaction, error) {
	return _IWorkerHub.Contract.UpdateModelTier(&_IWorkerHub.TransactOpts, _model, _tier)
}

// UpdateModelTier is a paid mutator transaction binding the contract method 0x0738a9bb.
//
// Solidity: function updateModelTier(address _model, uint32 _tier) returns()
func (_IWorkerHub *IWorkerHubTransactorSession) UpdateModelTier(_model common.Address, _tier uint32) (*types.Transaction, error) {
	return _IWorkerHub.Contract.UpdateModelTier(&_IWorkerHub.TransactOpts, _model, _tier)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IWorkerHub *IWorkerHubTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWorkerHub.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IWorkerHub *IWorkerHubSession) Receive() (*types.Transaction, error) {
	return _IWorkerHub.Contract.Receive(&_IWorkerHub.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IWorkerHub *IWorkerHubTransactorSession) Receive() (*types.Transaction, error) {
	return _IWorkerHub.Contract.Receive(&_IWorkerHub.TransactOpts)
}

// IWorkerHubBlocksPerEpochIterator is returned from FilterBlocksPerEpoch and is used to iterate over the raw logs and unpacked data for BlocksPerEpoch events raised by the IWorkerHub contract.
type IWorkerHubBlocksPerEpochIterator struct {
	Event *IWorkerHubBlocksPerEpoch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubBlocksPerEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubBlocksPerEpoch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubBlocksPerEpoch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubBlocksPerEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubBlocksPerEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubBlocksPerEpoch represents a BlocksPerEpoch event raised by the IWorkerHub contract.
type IWorkerHubBlocksPerEpoch struct {
	OldBlocks *big.Int
	NewBlocks *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlocksPerEpoch is a free log retrieval operation binding the contract event 0x3179ee2c3011a36d6d80a4b422f208df28ef9493d1d9ce1555b3116bd26ddb3d.
//
// Solidity: event BlocksPerEpoch(uint256 oldBlocks, uint256 newBlocks)
func (_IWorkerHub *IWorkerHubFilterer) FilterBlocksPerEpoch(opts *bind.FilterOpts) (*IWorkerHubBlocksPerEpochIterator, error) {

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "BlocksPerEpoch")
	if err != nil {
		return nil, err
	}
	return &IWorkerHubBlocksPerEpochIterator{contract: _IWorkerHub.contract, event: "BlocksPerEpoch", logs: logs, sub: sub}, nil
}

// WatchBlocksPerEpoch is a free log subscription operation binding the contract event 0x3179ee2c3011a36d6d80a4b422f208df28ef9493d1d9ce1555b3116bd26ddb3d.
//
// Solidity: event BlocksPerEpoch(uint256 oldBlocks, uint256 newBlocks)
func (_IWorkerHub *IWorkerHubFilterer) WatchBlocksPerEpoch(opts *bind.WatchOpts, sink chan<- *IWorkerHubBlocksPerEpoch) (event.Subscription, error) {

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "BlocksPerEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubBlocksPerEpoch)
				if err := _IWorkerHub.contract.UnpackLog(event, "BlocksPerEpoch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlocksPerEpoch is a log parse operation binding the contract event 0x3179ee2c3011a36d6d80a4b422f208df28ef9493d1d9ce1555b3116bd26ddb3d.
//
// Solidity: event BlocksPerEpoch(uint256 oldBlocks, uint256 newBlocks)
func (_IWorkerHub *IWorkerHubFilterer) ParseBlocksPerEpoch(log types.Log) (*IWorkerHubBlocksPerEpoch, error) {
	event := new(IWorkerHubBlocksPerEpoch)
	if err := _IWorkerHub.contract.UnpackLog(event, "BlocksPerEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubCommitDurationIterator is returned from FilterCommitDuration and is used to iterate over the raw logs and unpacked data for CommitDuration events raised by the IWorkerHub contract.
type IWorkerHubCommitDurationIterator struct {
	Event *IWorkerHubCommitDuration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubCommitDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubCommitDuration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubCommitDuration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubCommitDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubCommitDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubCommitDuration represents a CommitDuration event raised by the IWorkerHub contract.
type IWorkerHubCommitDuration struct {
	OldTime *big.Int
	NewTime *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCommitDuration is a free log retrieval operation binding the contract event 0xc9bc20c9ff07142c58c480090e116ebe561a42316260069d619782bb38faf619.
//
// Solidity: event CommitDuration(uint256 oldTime, uint256 newTime)
func (_IWorkerHub *IWorkerHubFilterer) FilterCommitDuration(opts *bind.FilterOpts) (*IWorkerHubCommitDurationIterator, error) {

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "CommitDuration")
	if err != nil {
		return nil, err
	}
	return &IWorkerHubCommitDurationIterator{contract: _IWorkerHub.contract, event: "CommitDuration", logs: logs, sub: sub}, nil
}

// WatchCommitDuration is a free log subscription operation binding the contract event 0xc9bc20c9ff07142c58c480090e116ebe561a42316260069d619782bb38faf619.
//
// Solidity: event CommitDuration(uint256 oldTime, uint256 newTime)
func (_IWorkerHub *IWorkerHubFilterer) WatchCommitDuration(opts *bind.WatchOpts, sink chan<- *IWorkerHubCommitDuration) (event.Subscription, error) {

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "CommitDuration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubCommitDuration)
				if err := _IWorkerHub.contract.UnpackLog(event, "CommitDuration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCommitDuration is a log parse operation binding the contract event 0xc9bc20c9ff07142c58c480090e116ebe561a42316260069d619782bb38faf619.
//
// Solidity: event CommitDuration(uint256 oldTime, uint256 newTime)
func (_IWorkerHub *IWorkerHubFilterer) ParseCommitDuration(log types.Log) (*IWorkerHubCommitDuration, error) {
	event := new(IWorkerHubCommitDuration)
	if err := _IWorkerHub.contract.UnpackLog(event, "CommitDuration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubCommitmentSubmissionIterator is returned from FilterCommitmentSubmission and is used to iterate over the raw logs and unpacked data for CommitmentSubmission events raised by the IWorkerHub contract.
type IWorkerHubCommitmentSubmissionIterator struct {
	Event *IWorkerHubCommitmentSubmission // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubCommitmentSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubCommitmentSubmission)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubCommitmentSubmission)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubCommitmentSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubCommitmentSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubCommitmentSubmission represents a CommitmentSubmission event raised by the IWorkerHub contract.
type IWorkerHubCommitmentSubmission struct {
	Miner       common.Address
	AssigmentId *big.Int
	Commitment  [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCommitmentSubmission is a free log retrieval operation binding the contract event 0x47a3511bbb68bfcf0b476106b3a5a5d5b8d7eb4205a28d6424a40fb19d9afa5b.
//
// Solidity: event CommitmentSubmission(address indexed miner, uint256 indexed assigmentId, bytes32 commitment)
func (_IWorkerHub *IWorkerHubFilterer) FilterCommitmentSubmission(opts *bind.FilterOpts, miner []common.Address, assigmentId []*big.Int) (*IWorkerHubCommitmentSubmissionIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var assigmentIdRule []interface{}
	for _, assigmentIdItem := range assigmentId {
		assigmentIdRule = append(assigmentIdRule, assigmentIdItem)
	}

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "CommitmentSubmission", minerRule, assigmentIdRule)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubCommitmentSubmissionIterator{contract: _IWorkerHub.contract, event: "CommitmentSubmission", logs: logs, sub: sub}, nil
}

// WatchCommitmentSubmission is a free log subscription operation binding the contract event 0x47a3511bbb68bfcf0b476106b3a5a5d5b8d7eb4205a28d6424a40fb19d9afa5b.
//
// Solidity: event CommitmentSubmission(address indexed miner, uint256 indexed assigmentId, bytes32 commitment)
func (_IWorkerHub *IWorkerHubFilterer) WatchCommitmentSubmission(opts *bind.WatchOpts, sink chan<- *IWorkerHubCommitmentSubmission, miner []common.Address, assigmentId []*big.Int) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var assigmentIdRule []interface{}
	for _, assigmentIdItem := range assigmentId {
		assigmentIdRule = append(assigmentIdRule, assigmentIdItem)
	}

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "CommitmentSubmission", minerRule, assigmentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubCommitmentSubmission)
				if err := _IWorkerHub.contract.UnpackLog(event, "CommitmentSubmission", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCommitmentSubmission is a log parse operation binding the contract event 0x47a3511bbb68bfcf0b476106b3a5a5d5b8d7eb4205a28d6424a40fb19d9afa5b.
//
// Solidity: event CommitmentSubmission(address indexed miner, uint256 indexed assigmentId, bytes32 commitment)
func (_IWorkerHub *IWorkerHubFilterer) ParseCommitmentSubmission(log types.Log) (*IWorkerHubCommitmentSubmission, error) {
	event := new(IWorkerHubCommitmentSubmission)
	if err := _IWorkerHub.contract.UnpackLog(event, "CommitmentSubmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubDAOTokenMintedV2Iterator is returned from FilterDAOTokenMintedV2 and is used to iterate over the raw logs and unpacked data for DAOTokenMintedV2 events raised by the IWorkerHub contract.
type IWorkerHubDAOTokenMintedV2Iterator struct {
	Event *IWorkerHubDAOTokenMintedV2 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubDAOTokenMintedV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubDAOTokenMintedV2)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubDAOTokenMintedV2)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubDAOTokenMintedV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubDAOTokenMintedV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubDAOTokenMintedV2 represents a DAOTokenMintedV2 event raised by the IWorkerHub contract.
type IWorkerHubDAOTokenMintedV2 struct {
	ChainId      *big.Int
	InferenceId  *big.Int
	ModelAddress common.Address
	Receivers    []IWorkerHubDAOTokenReceiverInfor
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDAOTokenMintedV2 is a free log retrieval operation binding the contract event 0x2faa16bd9d858bdbd007d15bb6ae06ff3f238c90433800dafb4c0f7e3f07a241.
//
// Solidity: event DAOTokenMintedV2(uint256 chainId, uint256 inferenceId, address modelAddress, (address,uint256,uint8)[] receivers)
func (_IWorkerHub *IWorkerHubFilterer) FilterDAOTokenMintedV2(opts *bind.FilterOpts) (*IWorkerHubDAOTokenMintedV2Iterator, error) {

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "DAOTokenMintedV2")
	if err != nil {
		return nil, err
	}
	return &IWorkerHubDAOTokenMintedV2Iterator{contract: _IWorkerHub.contract, event: "DAOTokenMintedV2", logs: logs, sub: sub}, nil
}

// WatchDAOTokenMintedV2 is a free log subscription operation binding the contract event 0x2faa16bd9d858bdbd007d15bb6ae06ff3f238c90433800dafb4c0f7e3f07a241.
//
// Solidity: event DAOTokenMintedV2(uint256 chainId, uint256 inferenceId, address modelAddress, (address,uint256,uint8)[] receivers)
func (_IWorkerHub *IWorkerHubFilterer) WatchDAOTokenMintedV2(opts *bind.WatchOpts, sink chan<- *IWorkerHubDAOTokenMintedV2) (event.Subscription, error) {

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "DAOTokenMintedV2")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubDAOTokenMintedV2)
				if err := _IWorkerHub.contract.UnpackLog(event, "DAOTokenMintedV2", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDAOTokenMintedV2 is a log parse operation binding the contract event 0x2faa16bd9d858bdbd007d15bb6ae06ff3f238c90433800dafb4c0f7e3f07a241.
//
// Solidity: event DAOTokenMintedV2(uint256 chainId, uint256 inferenceId, address modelAddress, (address,uint256,uint8)[] receivers)
func (_IWorkerHub *IWorkerHubFilterer) ParseDAOTokenMintedV2(log types.Log) (*IWorkerHubDAOTokenMintedV2, error) {
	event := new(IWorkerHubDAOTokenMintedV2)
	if err := _IWorkerHub.contract.UnpackLog(event, "DAOTokenMintedV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubDAOTokenPercentageUpdatedIterator is returned from FilterDAOTokenPercentageUpdated and is used to iterate over the raw logs and unpacked data for DAOTokenPercentageUpdated events raised by the IWorkerHub contract.
type IWorkerHubDAOTokenPercentageUpdatedIterator struct {
	Event *IWorkerHubDAOTokenPercentageUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubDAOTokenPercentageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubDAOTokenPercentageUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubDAOTokenPercentageUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubDAOTokenPercentageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubDAOTokenPercentageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubDAOTokenPercentageUpdated represents a DAOTokenPercentageUpdated event raised by the IWorkerHub contract.
type IWorkerHubDAOTokenPercentageUpdated struct {
	OldValue IWorkerHubDAOTokenPercentage
	NewValue IWorkerHubDAOTokenPercentage
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDAOTokenPercentageUpdated is a free log retrieval operation binding the contract event 0xe217c132ca1c9e392a91d1b2eaeb42b77b8ff74a61c75974d05ffbbe6e00a16a.
//
// Solidity: event DAOTokenPercentageUpdated((uint16,uint16,uint16,uint16,uint16) oldValue, (uint16,uint16,uint16,uint16,uint16) newValue)
func (_IWorkerHub *IWorkerHubFilterer) FilterDAOTokenPercentageUpdated(opts *bind.FilterOpts) (*IWorkerHubDAOTokenPercentageUpdatedIterator, error) {

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "DAOTokenPercentageUpdated")
	if err != nil {
		return nil, err
	}
	return &IWorkerHubDAOTokenPercentageUpdatedIterator{contract: _IWorkerHub.contract, event: "DAOTokenPercentageUpdated", logs: logs, sub: sub}, nil
}

// WatchDAOTokenPercentageUpdated is a free log subscription operation binding the contract event 0xe217c132ca1c9e392a91d1b2eaeb42b77b8ff74a61c75974d05ffbbe6e00a16a.
//
// Solidity: event DAOTokenPercentageUpdated((uint16,uint16,uint16,uint16,uint16) oldValue, (uint16,uint16,uint16,uint16,uint16) newValue)
func (_IWorkerHub *IWorkerHubFilterer) WatchDAOTokenPercentageUpdated(opts *bind.WatchOpts, sink chan<- *IWorkerHubDAOTokenPercentageUpdated) (event.Subscription, error) {

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "DAOTokenPercentageUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubDAOTokenPercentageUpdated)
				if err := _IWorkerHub.contract.UnpackLog(event, "DAOTokenPercentageUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDAOTokenPercentageUpdated is a log parse operation binding the contract event 0xe217c132ca1c9e392a91d1b2eaeb42b77b8ff74a61c75974d05ffbbe6e00a16a.
//
// Solidity: event DAOTokenPercentageUpdated((uint16,uint16,uint16,uint16,uint16) oldValue, (uint16,uint16,uint16,uint16,uint16) newValue)
func (_IWorkerHub *IWorkerHubFilterer) ParseDAOTokenPercentageUpdated(log types.Log) (*IWorkerHubDAOTokenPercentageUpdated, error) {
	event := new(IWorkerHubDAOTokenPercentageUpdated)
	if err := _IWorkerHub.contract.UnpackLog(event, "DAOTokenPercentageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubDAOTokenRewardUpdatedIterator is returned from FilterDAOTokenRewardUpdated and is used to iterate over the raw logs and unpacked data for DAOTokenRewardUpdated events raised by the IWorkerHub contract.
type IWorkerHubDAOTokenRewardUpdatedIterator struct {
	Event *IWorkerHubDAOTokenRewardUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubDAOTokenRewardUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubDAOTokenRewardUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubDAOTokenRewardUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubDAOTokenRewardUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubDAOTokenRewardUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubDAOTokenRewardUpdated represents a DAOTokenRewardUpdated event raised by the IWorkerHub contract.
type IWorkerHubDAOTokenRewardUpdated struct {
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDAOTokenRewardUpdated is a free log retrieval operation binding the contract event 0x454d79b61f30800ce19615c79c4f9a1eb892ed9372cf95ba71cbd2345f8fa9aa.
//
// Solidity: event DAOTokenRewardUpdated(uint256 oldValue, uint256 newValue)
func (_IWorkerHub *IWorkerHubFilterer) FilterDAOTokenRewardUpdated(opts *bind.FilterOpts) (*IWorkerHubDAOTokenRewardUpdatedIterator, error) {

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "DAOTokenRewardUpdated")
	if err != nil {
		return nil, err
	}
	return &IWorkerHubDAOTokenRewardUpdatedIterator{contract: _IWorkerHub.contract, event: "DAOTokenRewardUpdated", logs: logs, sub: sub}, nil
}

// WatchDAOTokenRewardUpdated is a free log subscription operation binding the contract event 0x454d79b61f30800ce19615c79c4f9a1eb892ed9372cf95ba71cbd2345f8fa9aa.
//
// Solidity: event DAOTokenRewardUpdated(uint256 oldValue, uint256 newValue)
func (_IWorkerHub *IWorkerHubFilterer) WatchDAOTokenRewardUpdated(opts *bind.WatchOpts, sink chan<- *IWorkerHubDAOTokenRewardUpdated) (event.Subscription, error) {

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "DAOTokenRewardUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubDAOTokenRewardUpdated)
				if err := _IWorkerHub.contract.UnpackLog(event, "DAOTokenRewardUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDAOTokenRewardUpdated is a log parse operation binding the contract event 0x454d79b61f30800ce19615c79c4f9a1eb892ed9372cf95ba71cbd2345f8fa9aa.
//
// Solidity: event DAOTokenRewardUpdated(uint256 oldValue, uint256 newValue)
func (_IWorkerHub *IWorkerHubFilterer) ParseDAOTokenRewardUpdated(log types.Log) (*IWorkerHubDAOTokenRewardUpdated, error) {
	event := new(IWorkerHubDAOTokenRewardUpdated)
	if err := _IWorkerHub.contract.UnpackLog(event, "DAOTokenRewardUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubDAOTokenUpdatedIterator is returned from FilterDAOTokenUpdated and is used to iterate over the raw logs and unpacked data for DAOTokenUpdated events raised by the IWorkerHub contract.
type IWorkerHubDAOTokenUpdatedIterator struct {
	Event *IWorkerHubDAOTokenUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubDAOTokenUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubDAOTokenUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubDAOTokenUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubDAOTokenUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubDAOTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubDAOTokenUpdated represents a DAOTokenUpdated event raised by the IWorkerHub contract.
type IWorkerHubDAOTokenUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDAOTokenUpdated is a free log retrieval operation binding the contract event 0x518cc1a1508767ac2e92e88727dbf2ace68f44768b3684e0ad2305f6db0cd8da.
//
// Solidity: event DAOTokenUpdated(address oldAddress, address newAddress)
func (_IWorkerHub *IWorkerHubFilterer) FilterDAOTokenUpdated(opts *bind.FilterOpts) (*IWorkerHubDAOTokenUpdatedIterator, error) {

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "DAOTokenUpdated")
	if err != nil {
		return nil, err
	}
	return &IWorkerHubDAOTokenUpdatedIterator{contract: _IWorkerHub.contract, event: "DAOTokenUpdated", logs: logs, sub: sub}, nil
}

// WatchDAOTokenUpdated is a free log subscription operation binding the contract event 0x518cc1a1508767ac2e92e88727dbf2ace68f44768b3684e0ad2305f6db0cd8da.
//
// Solidity: event DAOTokenUpdated(address oldAddress, address newAddress)
func (_IWorkerHub *IWorkerHubFilterer) WatchDAOTokenUpdated(opts *bind.WatchOpts, sink chan<- *IWorkerHubDAOTokenUpdated) (event.Subscription, error) {

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "DAOTokenUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubDAOTokenUpdated)
				if err := _IWorkerHub.contract.UnpackLog(event, "DAOTokenUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDAOTokenUpdated is a log parse operation binding the contract event 0x518cc1a1508767ac2e92e88727dbf2ace68f44768b3684e0ad2305f6db0cd8da.
//
// Solidity: event DAOTokenUpdated(address oldAddress, address newAddress)
func (_IWorkerHub *IWorkerHubFilterer) ParseDAOTokenUpdated(log types.Log) (*IWorkerHubDAOTokenUpdated, error) {
	event := new(IWorkerHubDAOTokenUpdated)
	if err := _IWorkerHub.contract.UnpackLog(event, "DAOTokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubFinePercentageUpdatedIterator is returned from FilterFinePercentageUpdated and is used to iterate over the raw logs and unpacked data for FinePercentageUpdated events raised by the IWorkerHub contract.
type IWorkerHubFinePercentageUpdatedIterator struct {
	Event *IWorkerHubFinePercentageUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubFinePercentageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubFinePercentageUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubFinePercentageUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubFinePercentageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubFinePercentageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubFinePercentageUpdated represents a FinePercentageUpdated event raised by the IWorkerHub contract.
type IWorkerHubFinePercentageUpdated struct {
	OldPercent uint16
	NewPercent uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFinePercentageUpdated is a free log retrieval operation binding the contract event 0xcf2ba21ec685fb1baf4b5e5df96fd2da47ab299e7d95e586c7898f114b6c1269.
//
// Solidity: event FinePercentageUpdated(uint16 oldPercent, uint16 newPercent)
func (_IWorkerHub *IWorkerHubFilterer) FilterFinePercentageUpdated(opts *bind.FilterOpts) (*IWorkerHubFinePercentageUpdatedIterator, error) {

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "FinePercentageUpdated")
	if err != nil {
		return nil, err
	}
	return &IWorkerHubFinePercentageUpdatedIterator{contract: _IWorkerHub.contract, event: "FinePercentageUpdated", logs: logs, sub: sub}, nil
}

// WatchFinePercentageUpdated is a free log subscription operation binding the contract event 0xcf2ba21ec685fb1baf4b5e5df96fd2da47ab299e7d95e586c7898f114b6c1269.
//
// Solidity: event FinePercentageUpdated(uint16 oldPercent, uint16 newPercent)
func (_IWorkerHub *IWorkerHubFilterer) WatchFinePercentageUpdated(opts *bind.WatchOpts, sink chan<- *IWorkerHubFinePercentageUpdated) (event.Subscription, error) {

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "FinePercentageUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubFinePercentageUpdated)
				if err := _IWorkerHub.contract.UnpackLog(event, "FinePercentageUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFinePercentageUpdated is a log parse operation binding the contract event 0xcf2ba21ec685fb1baf4b5e5df96fd2da47ab299e7d95e586c7898f114b6c1269.
//
// Solidity: event FinePercentageUpdated(uint16 oldPercent, uint16 newPercent)
func (_IWorkerHub *IWorkerHubFilterer) ParseFinePercentageUpdated(log types.Log) (*IWorkerHubFinePercentageUpdated, error) {
	event := new(IWorkerHubFinePercentageUpdated)
	if err := _IWorkerHub.contract.UnpackLog(event, "FinePercentageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubFraudulentMinerPenalizedIterator is returned from FilterFraudulentMinerPenalized and is used to iterate over the raw logs and unpacked data for FraudulentMinerPenalized events raised by the IWorkerHub contract.
type IWorkerHubFraudulentMinerPenalizedIterator struct {
	Event *IWorkerHubFraudulentMinerPenalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubFraudulentMinerPenalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubFraudulentMinerPenalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubFraudulentMinerPenalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubFraudulentMinerPenalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubFraudulentMinerPenalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubFraudulentMinerPenalized represents a FraudulentMinerPenalized event raised by the IWorkerHub contract.
type IWorkerHubFraudulentMinerPenalized struct {
	Miner        common.Address
	ModelAddress common.Address
	Treasury     common.Address
	Fine         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFraudulentMinerPenalized is a free log retrieval operation binding the contract event 0x63a49f9cdfcfe1fddc8bd7a881449dc97b664e888be5c2fdee7ca4a70b447e43.
//
// Solidity: event FraudulentMinerPenalized(address indexed miner, address indexed modelAddress, address indexed treasury, uint256 fine)
func (_IWorkerHub *IWorkerHubFilterer) FilterFraudulentMinerPenalized(opts *bind.FilterOpts, miner []common.Address, modelAddress []common.Address, treasury []common.Address) (*IWorkerHubFraudulentMinerPenalizedIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelAddressRule []interface{}
	for _, modelAddressItem := range modelAddress {
		modelAddressRule = append(modelAddressRule, modelAddressItem)
	}
	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "FraudulentMinerPenalized", minerRule, modelAddressRule, treasuryRule)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubFraudulentMinerPenalizedIterator{contract: _IWorkerHub.contract, event: "FraudulentMinerPenalized", logs: logs, sub: sub}, nil
}

// WatchFraudulentMinerPenalized is a free log subscription operation binding the contract event 0x63a49f9cdfcfe1fddc8bd7a881449dc97b664e888be5c2fdee7ca4a70b447e43.
//
// Solidity: event FraudulentMinerPenalized(address indexed miner, address indexed modelAddress, address indexed treasury, uint256 fine)
func (_IWorkerHub *IWorkerHubFilterer) WatchFraudulentMinerPenalized(opts *bind.WatchOpts, sink chan<- *IWorkerHubFraudulentMinerPenalized, miner []common.Address, modelAddress []common.Address, treasury []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelAddressRule []interface{}
	for _, modelAddressItem := range modelAddress {
		modelAddressRule = append(modelAddressRule, modelAddressItem)
	}
	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "FraudulentMinerPenalized", minerRule, modelAddressRule, treasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubFraudulentMinerPenalized)
				if err := _IWorkerHub.contract.UnpackLog(event, "FraudulentMinerPenalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFraudulentMinerPenalized is a log parse operation binding the contract event 0x63a49f9cdfcfe1fddc8bd7a881449dc97b664e888be5c2fdee7ca4a70b447e43.
//
// Solidity: event FraudulentMinerPenalized(address indexed miner, address indexed modelAddress, address indexed treasury, uint256 fine)
func (_IWorkerHub *IWorkerHubFilterer) ParseFraudulentMinerPenalized(log types.Log) (*IWorkerHubFraudulentMinerPenalized, error) {
	event := new(IWorkerHubFraudulentMinerPenalized)
	if err := _IWorkerHub.contract.UnpackLog(event, "FraudulentMinerPenalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubInferenceStatusUpdateIterator is returned from FilterInferenceStatusUpdate and is used to iterate over the raw logs and unpacked data for InferenceStatusUpdate events raised by the IWorkerHub contract.
type IWorkerHubInferenceStatusUpdateIterator struct {
	Event *IWorkerHubInferenceStatusUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubInferenceStatusUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubInferenceStatusUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubInferenceStatusUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubInferenceStatusUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubInferenceStatusUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubInferenceStatusUpdate represents a InferenceStatusUpdate event raised by the IWorkerHub contract.
type IWorkerHubInferenceStatusUpdate struct {
	InferenceId *big.Int
	NewStatus   uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInferenceStatusUpdate is a free log retrieval operation binding the contract event 0xbc645ece538d7606c8ac26de30aef5fbd0ed2ee0c945f4e5d860da3e62781d50.
//
// Solidity: event InferenceStatusUpdate(uint256 indexed inferenceId, uint8 newStatus)
func (_IWorkerHub *IWorkerHubFilterer) FilterInferenceStatusUpdate(opts *bind.FilterOpts, inferenceId []*big.Int) (*IWorkerHubInferenceStatusUpdateIterator, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "InferenceStatusUpdate", inferenceIdRule)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubInferenceStatusUpdateIterator{contract: _IWorkerHub.contract, event: "InferenceStatusUpdate", logs: logs, sub: sub}, nil
}

// WatchInferenceStatusUpdate is a free log subscription operation binding the contract event 0xbc645ece538d7606c8ac26de30aef5fbd0ed2ee0c945f4e5d860da3e62781d50.
//
// Solidity: event InferenceStatusUpdate(uint256 indexed inferenceId, uint8 newStatus)
func (_IWorkerHub *IWorkerHubFilterer) WatchInferenceStatusUpdate(opts *bind.WatchOpts, sink chan<- *IWorkerHubInferenceStatusUpdate, inferenceId []*big.Int) (event.Subscription, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "InferenceStatusUpdate", inferenceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubInferenceStatusUpdate)
				if err := _IWorkerHub.contract.UnpackLog(event, "InferenceStatusUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInferenceStatusUpdate is a log parse operation binding the contract event 0xbc645ece538d7606c8ac26de30aef5fbd0ed2ee0c945f4e5d860da3e62781d50.
//
// Solidity: event InferenceStatusUpdate(uint256 indexed inferenceId, uint8 newStatus)
func (_IWorkerHub *IWorkerHubFilterer) ParseInferenceStatusUpdate(log types.Log) (*IWorkerHubInferenceStatusUpdate, error) {
	event := new(IWorkerHubInferenceStatusUpdate)
	if err := _IWorkerHub.contract.UnpackLog(event, "InferenceStatusUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the IWorkerHub contract.
type IWorkerHubInitializedIterator struct {
	Event *IWorkerHubInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubInitialized represents a Initialized event raised by the IWorkerHub contract.
type IWorkerHubInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IWorkerHub *IWorkerHubFilterer) FilterInitialized(opts *bind.FilterOpts) (*IWorkerHubInitializedIterator, error) {

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &IWorkerHubInitializedIterator{contract: _IWorkerHub.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IWorkerHub *IWorkerHubFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *IWorkerHubInitialized) (event.Subscription, error) {

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubInitialized)
				if err := _IWorkerHub.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IWorkerHub *IWorkerHubFilterer) ParseInitialized(log types.Log) (*IWorkerHubInitialized, error) {
	event := new(IWorkerHubInitialized)
	if err := _IWorkerHub.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubL2OwnerUpdatedIterator is returned from FilterL2OwnerUpdated and is used to iterate over the raw logs and unpacked data for L2OwnerUpdated events raised by the IWorkerHub contract.
type IWorkerHubL2OwnerUpdatedIterator struct {
	Event *IWorkerHubL2OwnerUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubL2OwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubL2OwnerUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubL2OwnerUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubL2OwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubL2OwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubL2OwnerUpdated represents a L2OwnerUpdated event raised by the IWorkerHub contract.
type IWorkerHubL2OwnerUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterL2OwnerUpdated is a free log retrieval operation binding the contract event 0x3cfa9fea14972d7cbbd0fddda517d4467bd2863f1d28e76fa4e0fe230a7bf274.
//
// Solidity: event L2OwnerUpdated(address oldAddress, address newAddress)
func (_IWorkerHub *IWorkerHubFilterer) FilterL2OwnerUpdated(opts *bind.FilterOpts) (*IWorkerHubL2OwnerUpdatedIterator, error) {

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "L2OwnerUpdated")
	if err != nil {
		return nil, err
	}
	return &IWorkerHubL2OwnerUpdatedIterator{contract: _IWorkerHub.contract, event: "L2OwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchL2OwnerUpdated is a free log subscription operation binding the contract event 0x3cfa9fea14972d7cbbd0fddda517d4467bd2863f1d28e76fa4e0fe230a7bf274.
//
// Solidity: event L2OwnerUpdated(address oldAddress, address newAddress)
func (_IWorkerHub *IWorkerHubFilterer) WatchL2OwnerUpdated(opts *bind.WatchOpts, sink chan<- *IWorkerHubL2OwnerUpdated) (event.Subscription, error) {

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "L2OwnerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubL2OwnerUpdated)
				if err := _IWorkerHub.contract.UnpackLog(event, "L2OwnerUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseL2OwnerUpdated is a log parse operation binding the contract event 0x3cfa9fea14972d7cbbd0fddda517d4467bd2863f1d28e76fa4e0fe230a7bf274.
//
// Solidity: event L2OwnerUpdated(address oldAddress, address newAddress)
func (_IWorkerHub *IWorkerHubFilterer) ParseL2OwnerUpdated(log types.Log) (*IWorkerHubL2OwnerUpdated, error) {
	event := new(IWorkerHubL2OwnerUpdated)
	if err := _IWorkerHub.contract.UnpackLog(event, "L2OwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubMinFeeToUseUpdatedIterator is returned from FilterMinFeeToUseUpdated and is used to iterate over the raw logs and unpacked data for MinFeeToUseUpdated events raised by the IWorkerHub contract.
type IWorkerHubMinFeeToUseUpdatedIterator struct {
	Event *IWorkerHubMinFeeToUseUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubMinFeeToUseUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubMinFeeToUseUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubMinFeeToUseUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubMinFeeToUseUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubMinFeeToUseUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubMinFeeToUseUpdated represents a MinFeeToUseUpdated event raised by the IWorkerHub contract.
type IWorkerHubMinFeeToUseUpdated struct {
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMinFeeToUseUpdated is a free log retrieval operation binding the contract event 0x37bba2c63397e7d89baa40e3d0c29e309913eb87b9691bacb16dba509fad523c.
//
// Solidity: event MinFeeToUseUpdated(uint256 oldValue, uint256 newValue)
func (_IWorkerHub *IWorkerHubFilterer) FilterMinFeeToUseUpdated(opts *bind.FilterOpts) (*IWorkerHubMinFeeToUseUpdatedIterator, error) {

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "MinFeeToUseUpdated")
	if err != nil {
		return nil, err
	}
	return &IWorkerHubMinFeeToUseUpdatedIterator{contract: _IWorkerHub.contract, event: "MinFeeToUseUpdated", logs: logs, sub: sub}, nil
}

// WatchMinFeeToUseUpdated is a free log subscription operation binding the contract event 0x37bba2c63397e7d89baa40e3d0c29e309913eb87b9691bacb16dba509fad523c.
//
// Solidity: event MinFeeToUseUpdated(uint256 oldValue, uint256 newValue)
func (_IWorkerHub *IWorkerHubFilterer) WatchMinFeeToUseUpdated(opts *bind.WatchOpts, sink chan<- *IWorkerHubMinFeeToUseUpdated) (event.Subscription, error) {

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "MinFeeToUseUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubMinFeeToUseUpdated)
				if err := _IWorkerHub.contract.UnpackLog(event, "MinFeeToUseUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinFeeToUseUpdated is a log parse operation binding the contract event 0x37bba2c63397e7d89baa40e3d0c29e309913eb87b9691bacb16dba509fad523c.
//
// Solidity: event MinFeeToUseUpdated(uint256 oldValue, uint256 newValue)
func (_IWorkerHub *IWorkerHubFilterer) ParseMinFeeToUseUpdated(log types.Log) (*IWorkerHubMinFeeToUseUpdated, error) {
	event := new(IWorkerHubMinFeeToUseUpdated)
	if err := _IWorkerHub.contract.UnpackLog(event, "MinFeeToUseUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubMinerDeactivatedIterator is returned from FilterMinerDeactivated and is used to iterate over the raw logs and unpacked data for MinerDeactivated events raised by the IWorkerHub contract.
type IWorkerHubMinerDeactivatedIterator struct {
	Event *IWorkerHubMinerDeactivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubMinerDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubMinerDeactivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubMinerDeactivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubMinerDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubMinerDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubMinerDeactivated represents a MinerDeactivated event raised by the IWorkerHub contract.
type IWorkerHubMinerDeactivated struct {
	Miner        common.Address
	ModelAddress common.Address
	ActiveTime   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMinerDeactivated is a free log retrieval operation binding the contract event 0x9335a7723b09748526d22902742e96812ad183ab52d86c2030fe407ff626e50d.
//
// Solidity: event MinerDeactivated(address indexed miner, address indexed modelAddress, uint40 activeTime)
func (_IWorkerHub *IWorkerHubFilterer) FilterMinerDeactivated(opts *bind.FilterOpts, miner []common.Address, modelAddress []common.Address) (*IWorkerHubMinerDeactivatedIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelAddressRule []interface{}
	for _, modelAddressItem := range modelAddress {
		modelAddressRule = append(modelAddressRule, modelAddressItem)
	}

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "MinerDeactivated", minerRule, modelAddressRule)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubMinerDeactivatedIterator{contract: _IWorkerHub.contract, event: "MinerDeactivated", logs: logs, sub: sub}, nil
}

// WatchMinerDeactivated is a free log subscription operation binding the contract event 0x9335a7723b09748526d22902742e96812ad183ab52d86c2030fe407ff626e50d.
//
// Solidity: event MinerDeactivated(address indexed miner, address indexed modelAddress, uint40 activeTime)
func (_IWorkerHub *IWorkerHubFilterer) WatchMinerDeactivated(opts *bind.WatchOpts, sink chan<- *IWorkerHubMinerDeactivated, miner []common.Address, modelAddress []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var modelAddressRule []interface{}
	for _, modelAddressItem := range modelAddress {
		modelAddressRule = append(modelAddressRule, modelAddressItem)
	}

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "MinerDeactivated", minerRule, modelAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubMinerDeactivated)
				if err := _IWorkerHub.contract.UnpackLog(event, "MinerDeactivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinerDeactivated is a log parse operation binding the contract event 0x9335a7723b09748526d22902742e96812ad183ab52d86c2030fe407ff626e50d.
//
// Solidity: event MinerDeactivated(address indexed miner, address indexed modelAddress, uint40 activeTime)
func (_IWorkerHub *IWorkerHubFilterer) ParseMinerDeactivated(log types.Log) (*IWorkerHubMinerDeactivated, error) {
	event := new(IWorkerHubMinerDeactivated)
	if err := _IWorkerHub.contract.UnpackLog(event, "MinerDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubMinerExtraStakeIterator is returned from FilterMinerExtraStake and is used to iterate over the raw logs and unpacked data for MinerExtraStake events raised by the IWorkerHub contract.
type IWorkerHubMinerExtraStakeIterator struct {
	Event *IWorkerHubMinerExtraStake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubMinerExtraStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubMinerExtraStake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubMinerExtraStake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubMinerExtraStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubMinerExtraStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubMinerExtraStake represents a MinerExtraStake event raised by the IWorkerHub contract.
type IWorkerHubMinerExtraStake struct {
	Miner common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerExtraStake is a free log retrieval operation binding the contract event 0x3d236e8f743e932a32c84d3114ce3e7ee0b75225cb3b39f72faac62495fd21c1.
//
// Solidity: event MinerExtraStake(address indexed miner, uint256 value)
func (_IWorkerHub *IWorkerHubFilterer) FilterMinerExtraStake(opts *bind.FilterOpts, miner []common.Address) (*IWorkerHubMinerExtraStakeIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "MinerExtraStake", minerRule)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubMinerExtraStakeIterator{contract: _IWorkerHub.contract, event: "MinerExtraStake", logs: logs, sub: sub}, nil
}

// WatchMinerExtraStake is a free log subscription operation binding the contract event 0x3d236e8f743e932a32c84d3114ce3e7ee0b75225cb3b39f72faac62495fd21c1.
//
// Solidity: event MinerExtraStake(address indexed miner, uint256 value)
func (_IWorkerHub *IWorkerHubFilterer) WatchMinerExtraStake(opts *bind.WatchOpts, sink chan<- *IWorkerHubMinerExtraStake, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "MinerExtraStake", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubMinerExtraStake)
				if err := _IWorkerHub.contract.UnpackLog(event, "MinerExtraStake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinerExtraStake is a log parse operation binding the contract event 0x3d236e8f743e932a32c84d3114ce3e7ee0b75225cb3b39f72faac62495fd21c1.
//
// Solidity: event MinerExtraStake(address indexed miner, uint256 value)
func (_IWorkerHub *IWorkerHubFilterer) ParseMinerExtraStake(log types.Log) (*IWorkerHubMinerExtraStake, error) {
	event := new(IWorkerHubMinerExtraStake)
	if err := _IWorkerHub.contract.UnpackLog(event, "MinerExtraStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubMinerJoinIterator is returned from FilterMinerJoin and is used to iterate over the raw logs and unpacked data for MinerJoin events raised by the IWorkerHub contract.
type IWorkerHubMinerJoinIterator struct {
	Event *IWorkerHubMinerJoin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubMinerJoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubMinerJoin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubMinerJoin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubMinerJoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubMinerJoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubMinerJoin represents a MinerJoin event raised by the IWorkerHub contract.
type IWorkerHubMinerJoin struct {
	Miner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerJoin is a free log retrieval operation binding the contract event 0xb7041987154996ed34981c2bc6fbafd4b1fcab9964486d7cc386f0d8abcc5446.
//
// Solidity: event MinerJoin(address indexed miner)
func (_IWorkerHub *IWorkerHubFilterer) FilterMinerJoin(opts *bind.FilterOpts, miner []common.Address) (*IWorkerHubMinerJoinIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "MinerJoin", minerRule)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubMinerJoinIterator{contract: _IWorkerHub.contract, event: "MinerJoin", logs: logs, sub: sub}, nil
}

// WatchMinerJoin is a free log subscription operation binding the contract event 0xb7041987154996ed34981c2bc6fbafd4b1fcab9964486d7cc386f0d8abcc5446.
//
// Solidity: event MinerJoin(address indexed miner)
func (_IWorkerHub *IWorkerHubFilterer) WatchMinerJoin(opts *bind.WatchOpts, sink chan<- *IWorkerHubMinerJoin, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "MinerJoin", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubMinerJoin)
				if err := _IWorkerHub.contract.UnpackLog(event, "MinerJoin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinerJoin is a log parse operation binding the contract event 0xb7041987154996ed34981c2bc6fbafd4b1fcab9964486d7cc386f0d8abcc5446.
//
// Solidity: event MinerJoin(address indexed miner)
func (_IWorkerHub *IWorkerHubFilterer) ParseMinerJoin(log types.Log) (*IWorkerHubMinerJoin, error) {
	event := new(IWorkerHubMinerJoin)
	if err := _IWorkerHub.contract.UnpackLog(event, "MinerJoin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubMinerRegistrationIterator is returned from FilterMinerRegistration and is used to iterate over the raw logs and unpacked data for MinerRegistration events raised by the IWorkerHub contract.
type IWorkerHubMinerRegistrationIterator struct {
	Event *IWorkerHubMinerRegistration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubMinerRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubMinerRegistration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubMinerRegistration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubMinerRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubMinerRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubMinerRegistration represents a MinerRegistration event raised by the IWorkerHub contract.
type IWorkerHubMinerRegistration struct {
	Miner common.Address
	Tier  uint16
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerRegistration is a free log retrieval operation binding the contract event 0x55e488821080f3f5cdf6088b02793df0d26f40053a70b6154347d2ac313015a1.
//
// Solidity: event MinerRegistration(address indexed miner, uint16 indexed tier, uint256 value)
func (_IWorkerHub *IWorkerHubFilterer) FilterMinerRegistration(opts *bind.FilterOpts, miner []common.Address, tier []uint16) (*IWorkerHubMinerRegistrationIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "MinerRegistration", minerRule, tierRule)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubMinerRegistrationIterator{contract: _IWorkerHub.contract, event: "MinerRegistration", logs: logs, sub: sub}, nil
}

// WatchMinerRegistration is a free log subscription operation binding the contract event 0x55e488821080f3f5cdf6088b02793df0d26f40053a70b6154347d2ac313015a1.
//
// Solidity: event MinerRegistration(address indexed miner, uint16 indexed tier, uint256 value)
func (_IWorkerHub *IWorkerHubFilterer) WatchMinerRegistration(opts *bind.WatchOpts, sink chan<- *IWorkerHubMinerRegistration, miner []common.Address, tier []uint16) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "MinerRegistration", minerRule, tierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubMinerRegistration)
				if err := _IWorkerHub.contract.UnpackLog(event, "MinerRegistration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinerRegistration is a log parse operation binding the contract event 0x55e488821080f3f5cdf6088b02793df0d26f40053a70b6154347d2ac313015a1.
//
// Solidity: event MinerRegistration(address indexed miner, uint16 indexed tier, uint256 value)
func (_IWorkerHub *IWorkerHubFilterer) ParseMinerRegistration(log types.Log) (*IWorkerHubMinerRegistration, error) {
	event := new(IWorkerHubMinerRegistration)
	if err := _IWorkerHub.contract.UnpackLog(event, "MinerRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubMinerRoleSeizedIterator is returned from FilterMinerRoleSeized and is used to iterate over the raw logs and unpacked data for MinerRoleSeized events raised by the IWorkerHub contract.
type IWorkerHubMinerRoleSeizedIterator struct {
	Event *IWorkerHubMinerRoleSeized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubMinerRoleSeizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubMinerRoleSeized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubMinerRoleSeized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubMinerRoleSeizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubMinerRoleSeizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubMinerRoleSeized represents a MinerRoleSeized event raised by the IWorkerHub contract.
type IWorkerHubMinerRoleSeized struct {
	AssignmentId *big.Int
	InferenceId  *big.Int
	Miner        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMinerRoleSeized is a free log retrieval operation binding the contract event 0x3d4f35957f03b76084f29d7c66d573fcec3d2e4bbc2844549e44bc1aed4c6c24.
//
// Solidity: event MinerRoleSeized(uint256 indexed assignmentId, uint256 indexed inferenceId, address indexed miner)
func (_IWorkerHub *IWorkerHubFilterer) FilterMinerRoleSeized(opts *bind.FilterOpts, assignmentId []*big.Int, inferenceId []*big.Int, miner []common.Address) (*IWorkerHubMinerRoleSeizedIterator, error) {

	var assignmentIdRule []interface{}
	for _, assignmentIdItem := range assignmentId {
		assignmentIdRule = append(assignmentIdRule, assignmentIdItem)
	}
	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "MinerRoleSeized", assignmentIdRule, inferenceIdRule, minerRule)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubMinerRoleSeizedIterator{contract: _IWorkerHub.contract, event: "MinerRoleSeized", logs: logs, sub: sub}, nil
}

// WatchMinerRoleSeized is a free log subscription operation binding the contract event 0x3d4f35957f03b76084f29d7c66d573fcec3d2e4bbc2844549e44bc1aed4c6c24.
//
// Solidity: event MinerRoleSeized(uint256 indexed assignmentId, uint256 indexed inferenceId, address indexed miner)
func (_IWorkerHub *IWorkerHubFilterer) WatchMinerRoleSeized(opts *bind.WatchOpts, sink chan<- *IWorkerHubMinerRoleSeized, assignmentId []*big.Int, inferenceId []*big.Int, miner []common.Address) (event.Subscription, error) {

	var assignmentIdRule []interface{}
	for _, assignmentIdItem := range assignmentId {
		assignmentIdRule = append(assignmentIdRule, assignmentIdItem)
	}
	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "MinerRoleSeized", assignmentIdRule, inferenceIdRule, minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubMinerRoleSeized)
				if err := _IWorkerHub.contract.UnpackLog(event, "MinerRoleSeized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinerRoleSeized is a log parse operation binding the contract event 0x3d4f35957f03b76084f29d7c66d573fcec3d2e4bbc2844549e44bc1aed4c6c24.
//
// Solidity: event MinerRoleSeized(uint256 indexed assignmentId, uint256 indexed inferenceId, address indexed miner)
func (_IWorkerHub *IWorkerHubFilterer) ParseMinerRoleSeized(log types.Log) (*IWorkerHubMinerRoleSeized, error) {
	event := new(IWorkerHubMinerRoleSeized)
	if err := _IWorkerHub.contract.UnpackLog(event, "MinerRoleSeized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubMinerUnregistrationIterator is returned from FilterMinerUnregistration and is used to iterate over the raw logs and unpacked data for MinerUnregistration events raised by the IWorkerHub contract.
type IWorkerHubMinerUnregistrationIterator struct {
	Event *IWorkerHubMinerUnregistration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubMinerUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubMinerUnregistration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubMinerUnregistration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubMinerUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubMinerUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubMinerUnregistration represents a MinerUnregistration event raised by the IWorkerHub contract.
type IWorkerHubMinerUnregistration struct {
	Miner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerUnregistration is a free log retrieval operation binding the contract event 0x8f54596d72781f60dbf7dad7e576f06ce17bbda0bdf384463f7734f85f51498e.
//
// Solidity: event MinerUnregistration(address indexed miner)
func (_IWorkerHub *IWorkerHubFilterer) FilterMinerUnregistration(opts *bind.FilterOpts, miner []common.Address) (*IWorkerHubMinerUnregistrationIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "MinerUnregistration", minerRule)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubMinerUnregistrationIterator{contract: _IWorkerHub.contract, event: "MinerUnregistration", logs: logs, sub: sub}, nil
}

// WatchMinerUnregistration is a free log subscription operation binding the contract event 0x8f54596d72781f60dbf7dad7e576f06ce17bbda0bdf384463f7734f85f51498e.
//
// Solidity: event MinerUnregistration(address indexed miner)
func (_IWorkerHub *IWorkerHubFilterer) WatchMinerUnregistration(opts *bind.WatchOpts, sink chan<- *IWorkerHubMinerUnregistration, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "MinerUnregistration", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubMinerUnregistration)
				if err := _IWorkerHub.contract.UnpackLog(event, "MinerUnregistration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinerUnregistration is a log parse operation binding the contract event 0x8f54596d72781f60dbf7dad7e576f06ce17bbda0bdf384463f7734f85f51498e.
//
// Solidity: event MinerUnregistration(address indexed miner)
func (_IWorkerHub *IWorkerHubFilterer) ParseMinerUnregistration(log types.Log) (*IWorkerHubMinerUnregistration, error) {
	event := new(IWorkerHubMinerUnregistration)
	if err := _IWorkerHub.contract.UnpackLog(event, "MinerUnregistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubMinerUnstakeIterator is returned from FilterMinerUnstake and is used to iterate over the raw logs and unpacked data for MinerUnstake events raised by the IWorkerHub contract.
type IWorkerHubMinerUnstakeIterator struct {
	Event *IWorkerHubMinerUnstake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubMinerUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubMinerUnstake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubMinerUnstake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubMinerUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubMinerUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubMinerUnstake represents a MinerUnstake event raised by the IWorkerHub contract.
type IWorkerHubMinerUnstake struct {
	Miner common.Address
	Stake *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinerUnstake is a free log retrieval operation binding the contract event 0x1051154647682075e7cc0645853209e75208cb5acd862fc83f7fd0fcaa9624b4.
//
// Solidity: event MinerUnstake(address indexed miner, uint256 stake)
func (_IWorkerHub *IWorkerHubFilterer) FilterMinerUnstake(opts *bind.FilterOpts, miner []common.Address) (*IWorkerHubMinerUnstakeIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "MinerUnstake", minerRule)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubMinerUnstakeIterator{contract: _IWorkerHub.contract, event: "MinerUnstake", logs: logs, sub: sub}, nil
}

// WatchMinerUnstake is a free log subscription operation binding the contract event 0x1051154647682075e7cc0645853209e75208cb5acd862fc83f7fd0fcaa9624b4.
//
// Solidity: event MinerUnstake(address indexed miner, uint256 stake)
func (_IWorkerHub *IWorkerHubFilterer) WatchMinerUnstake(opts *bind.WatchOpts, sink chan<- *IWorkerHubMinerUnstake, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "MinerUnstake", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubMinerUnstake)
				if err := _IWorkerHub.contract.UnpackLog(event, "MinerUnstake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinerUnstake is a log parse operation binding the contract event 0x1051154647682075e7cc0645853209e75208cb5acd862fc83f7fd0fcaa9624b4.
//
// Solidity: event MinerUnstake(address indexed miner, uint256 stake)
func (_IWorkerHub *IWorkerHubFilterer) ParseMinerUnstake(log types.Log) (*IWorkerHubMinerUnstake, error) {
	event := new(IWorkerHubMinerUnstake)
	if err := _IWorkerHub.contract.UnpackLog(event, "MinerUnstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubMiningTimeLimitUpdateIterator is returned from FilterMiningTimeLimitUpdate and is used to iterate over the raw logs and unpacked data for MiningTimeLimitUpdate events raised by the IWorkerHub contract.
type IWorkerHubMiningTimeLimitUpdateIterator struct {
	Event *IWorkerHubMiningTimeLimitUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubMiningTimeLimitUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubMiningTimeLimitUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubMiningTimeLimitUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubMiningTimeLimitUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubMiningTimeLimitUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubMiningTimeLimitUpdate represents a MiningTimeLimitUpdate event raised by the IWorkerHub contract.
type IWorkerHubMiningTimeLimitUpdate struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMiningTimeLimitUpdate is a free log retrieval operation binding the contract event 0xd223a90576ecd9f418b264c3465ab13fad46f62b72bf17dca91af5dc8b7e55a8.
//
// Solidity: event MiningTimeLimitUpdate(uint40 newValue)
func (_IWorkerHub *IWorkerHubFilterer) FilterMiningTimeLimitUpdate(opts *bind.FilterOpts) (*IWorkerHubMiningTimeLimitUpdateIterator, error) {

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "MiningTimeLimitUpdate")
	if err != nil {
		return nil, err
	}
	return &IWorkerHubMiningTimeLimitUpdateIterator{contract: _IWorkerHub.contract, event: "MiningTimeLimitUpdate", logs: logs, sub: sub}, nil
}

// WatchMiningTimeLimitUpdate is a free log subscription operation binding the contract event 0xd223a90576ecd9f418b264c3465ab13fad46f62b72bf17dca91af5dc8b7e55a8.
//
// Solidity: event MiningTimeLimitUpdate(uint40 newValue)
func (_IWorkerHub *IWorkerHubFilterer) WatchMiningTimeLimitUpdate(opts *bind.WatchOpts, sink chan<- *IWorkerHubMiningTimeLimitUpdate) (event.Subscription, error) {

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "MiningTimeLimitUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubMiningTimeLimitUpdate)
				if err := _IWorkerHub.contract.UnpackLog(event, "MiningTimeLimitUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMiningTimeLimitUpdate is a log parse operation binding the contract event 0xd223a90576ecd9f418b264c3465ab13fad46f62b72bf17dca91af5dc8b7e55a8.
//
// Solidity: event MiningTimeLimitUpdate(uint40 newValue)
func (_IWorkerHub *IWorkerHubFilterer) ParseMiningTimeLimitUpdate(log types.Log) (*IWorkerHubMiningTimeLimitUpdate, error) {
	event := new(IWorkerHubMiningTimeLimitUpdate)
	if err := _IWorkerHub.contract.UnpackLog(event, "MiningTimeLimitUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubModelMinimumFeeUpdateIterator is returned from FilterModelMinimumFeeUpdate and is used to iterate over the raw logs and unpacked data for ModelMinimumFeeUpdate events raised by the IWorkerHub contract.
type IWorkerHubModelMinimumFeeUpdateIterator struct {
	Event *IWorkerHubModelMinimumFeeUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubModelMinimumFeeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubModelMinimumFeeUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubModelMinimumFeeUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubModelMinimumFeeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubModelMinimumFeeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubModelMinimumFeeUpdate represents a ModelMinimumFeeUpdate event raised by the IWorkerHub contract.
type IWorkerHubModelMinimumFeeUpdate struct {
	Model      common.Address
	MinimumFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterModelMinimumFeeUpdate is a free log retrieval operation binding the contract event 0x923b5fe9c9974b3c93e434ae744faaa60ec86513c02614da5c8d9c51eda2bdd7.
//
// Solidity: event ModelMinimumFeeUpdate(address indexed model, uint256 minimumFee)
func (_IWorkerHub *IWorkerHubFilterer) FilterModelMinimumFeeUpdate(opts *bind.FilterOpts, model []common.Address) (*IWorkerHubModelMinimumFeeUpdateIterator, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "ModelMinimumFeeUpdate", modelRule)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubModelMinimumFeeUpdateIterator{contract: _IWorkerHub.contract, event: "ModelMinimumFeeUpdate", logs: logs, sub: sub}, nil
}

// WatchModelMinimumFeeUpdate is a free log subscription operation binding the contract event 0x923b5fe9c9974b3c93e434ae744faaa60ec86513c02614da5c8d9c51eda2bdd7.
//
// Solidity: event ModelMinimumFeeUpdate(address indexed model, uint256 minimumFee)
func (_IWorkerHub *IWorkerHubFilterer) WatchModelMinimumFeeUpdate(opts *bind.WatchOpts, sink chan<- *IWorkerHubModelMinimumFeeUpdate, model []common.Address) (event.Subscription, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "ModelMinimumFeeUpdate", modelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubModelMinimumFeeUpdate)
				if err := _IWorkerHub.contract.UnpackLog(event, "ModelMinimumFeeUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseModelMinimumFeeUpdate is a log parse operation binding the contract event 0x923b5fe9c9974b3c93e434ae744faaa60ec86513c02614da5c8d9c51eda2bdd7.
//
// Solidity: event ModelMinimumFeeUpdate(address indexed model, uint256 minimumFee)
func (_IWorkerHub *IWorkerHubFilterer) ParseModelMinimumFeeUpdate(log types.Log) (*IWorkerHubModelMinimumFeeUpdate, error) {
	event := new(IWorkerHubModelMinimumFeeUpdate)
	if err := _IWorkerHub.contract.UnpackLog(event, "ModelMinimumFeeUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubModelRegistrationIterator is returned from FilterModelRegistration and is used to iterate over the raw logs and unpacked data for ModelRegistration events raised by the IWorkerHub contract.
type IWorkerHubModelRegistrationIterator struct {
	Event *IWorkerHubModelRegistration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubModelRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubModelRegistration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubModelRegistration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubModelRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubModelRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubModelRegistration represents a ModelRegistration event raised by the IWorkerHub contract.
type IWorkerHubModelRegistration struct {
	Model      common.Address
	Tier       uint16
	MinimumFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterModelRegistration is a free log retrieval operation binding the contract event 0x7041913a4cb21c28c931da9d9e4b5ed0ad84e47fcf2a65527f03c438d534ed5c.
//
// Solidity: event ModelRegistration(address indexed model, uint16 indexed tier, uint256 minimumFee)
func (_IWorkerHub *IWorkerHubFilterer) FilterModelRegistration(opts *bind.FilterOpts, model []common.Address, tier []uint16) (*IWorkerHubModelRegistrationIterator, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "ModelRegistration", modelRule, tierRule)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubModelRegistrationIterator{contract: _IWorkerHub.contract, event: "ModelRegistration", logs: logs, sub: sub}, nil
}

// WatchModelRegistration is a free log subscription operation binding the contract event 0x7041913a4cb21c28c931da9d9e4b5ed0ad84e47fcf2a65527f03c438d534ed5c.
//
// Solidity: event ModelRegistration(address indexed model, uint16 indexed tier, uint256 minimumFee)
func (_IWorkerHub *IWorkerHubFilterer) WatchModelRegistration(opts *bind.WatchOpts, sink chan<- *IWorkerHubModelRegistration, model []common.Address, tier []uint16) (event.Subscription, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}
	var tierRule []interface{}
	for _, tierItem := range tier {
		tierRule = append(tierRule, tierItem)
	}

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "ModelRegistration", modelRule, tierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubModelRegistration)
				if err := _IWorkerHub.contract.UnpackLog(event, "ModelRegistration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseModelRegistration is a log parse operation binding the contract event 0x7041913a4cb21c28c931da9d9e4b5ed0ad84e47fcf2a65527f03c438d534ed5c.
//
// Solidity: event ModelRegistration(address indexed model, uint16 indexed tier, uint256 minimumFee)
func (_IWorkerHub *IWorkerHubFilterer) ParseModelRegistration(log types.Log) (*IWorkerHubModelRegistration, error) {
	event := new(IWorkerHubModelRegistration)
	if err := _IWorkerHub.contract.UnpackLog(event, "ModelRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubModelTierUpdateIterator is returned from FilterModelTierUpdate and is used to iterate over the raw logs and unpacked data for ModelTierUpdate events raised by the IWorkerHub contract.
type IWorkerHubModelTierUpdateIterator struct {
	Event *IWorkerHubModelTierUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubModelTierUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubModelTierUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubModelTierUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubModelTierUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubModelTierUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubModelTierUpdate represents a ModelTierUpdate event raised by the IWorkerHub contract.
type IWorkerHubModelTierUpdate struct {
	Model common.Address
	Tier  uint32
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterModelTierUpdate is a free log retrieval operation binding the contract event 0x64905396482bb1067a551077143915c77b512b1cfea5db34c903943c1c2a5a15.
//
// Solidity: event ModelTierUpdate(address indexed model, uint32 tier)
func (_IWorkerHub *IWorkerHubFilterer) FilterModelTierUpdate(opts *bind.FilterOpts, model []common.Address) (*IWorkerHubModelTierUpdateIterator, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "ModelTierUpdate", modelRule)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubModelTierUpdateIterator{contract: _IWorkerHub.contract, event: "ModelTierUpdate", logs: logs, sub: sub}, nil
}

// WatchModelTierUpdate is a free log subscription operation binding the contract event 0x64905396482bb1067a551077143915c77b512b1cfea5db34c903943c1c2a5a15.
//
// Solidity: event ModelTierUpdate(address indexed model, uint32 tier)
func (_IWorkerHub *IWorkerHubFilterer) WatchModelTierUpdate(opts *bind.WatchOpts, sink chan<- *IWorkerHubModelTierUpdate, model []common.Address) (event.Subscription, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "ModelTierUpdate", modelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubModelTierUpdate)
				if err := _IWorkerHub.contract.UnpackLog(event, "ModelTierUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseModelTierUpdate is a log parse operation binding the contract event 0x64905396482bb1067a551077143915c77b512b1cfea5db34c903943c1c2a5a15.
//
// Solidity: event ModelTierUpdate(address indexed model, uint32 tier)
func (_IWorkerHub *IWorkerHubFilterer) ParseModelTierUpdate(log types.Log) (*IWorkerHubModelTierUpdate, error) {
	event := new(IWorkerHubModelTierUpdate)
	if err := _IWorkerHub.contract.UnpackLog(event, "ModelTierUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubModelUnregistrationIterator is returned from FilterModelUnregistration and is used to iterate over the raw logs and unpacked data for ModelUnregistration events raised by the IWorkerHub contract.
type IWorkerHubModelUnregistrationIterator struct {
	Event *IWorkerHubModelUnregistration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubModelUnregistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubModelUnregistration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubModelUnregistration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubModelUnregistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubModelUnregistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubModelUnregistration represents a ModelUnregistration event raised by the IWorkerHub contract.
type IWorkerHubModelUnregistration struct {
	Model common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterModelUnregistration is a free log retrieval operation binding the contract event 0x68180f49300b9177ab3b88d3f909a002abeb9c2f769543a93234ca68333582d7.
//
// Solidity: event ModelUnregistration(address indexed model)
func (_IWorkerHub *IWorkerHubFilterer) FilterModelUnregistration(opts *bind.FilterOpts, model []common.Address) (*IWorkerHubModelUnregistrationIterator, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "ModelUnregistration", modelRule)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubModelUnregistrationIterator{contract: _IWorkerHub.contract, event: "ModelUnregistration", logs: logs, sub: sub}, nil
}

// WatchModelUnregistration is a free log subscription operation binding the contract event 0x68180f49300b9177ab3b88d3f909a002abeb9c2f769543a93234ca68333582d7.
//
// Solidity: event ModelUnregistration(address indexed model)
func (_IWorkerHub *IWorkerHubFilterer) WatchModelUnregistration(opts *bind.WatchOpts, sink chan<- *IWorkerHubModelUnregistration, model []common.Address) (event.Subscription, error) {

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "ModelUnregistration", modelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubModelUnregistration)
				if err := _IWorkerHub.contract.UnpackLog(event, "ModelUnregistration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseModelUnregistration is a log parse operation binding the contract event 0x68180f49300b9177ab3b88d3f909a002abeb9c2f769543a93234ca68333582d7.
//
// Solidity: event ModelUnregistration(address indexed model)
func (_IWorkerHub *IWorkerHubFilterer) ParseModelUnregistration(log types.Log) (*IWorkerHubModelUnregistration, error) {
	event := new(IWorkerHubModelUnregistration)
	if err := _IWorkerHub.contract.UnpackLog(event, "ModelUnregistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubNewAssignmentIterator is returned from FilterNewAssignment and is used to iterate over the raw logs and unpacked data for NewAssignment events raised by the IWorkerHub contract.
type IWorkerHubNewAssignmentIterator struct {
	Event *IWorkerHubNewAssignment // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubNewAssignmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubNewAssignment)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubNewAssignment)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubNewAssignmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubNewAssignmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubNewAssignment represents a NewAssignment event raised by the IWorkerHub contract.
type IWorkerHubNewAssignment struct {
	AssignmentId *big.Int
	InferenceId  *big.Int
	Miner        common.Address
	ExpiredAt    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewAssignment is a free log retrieval operation binding the contract event 0x53cc8b652f33c56dac5f1c97a284cc971e7adcb8abe9454b0853f076c6deb7d5.
//
// Solidity: event NewAssignment(uint256 indexed assignmentId, uint256 indexed inferenceId, address indexed miner, uint40 expiredAt)
func (_IWorkerHub *IWorkerHubFilterer) FilterNewAssignment(opts *bind.FilterOpts, assignmentId []*big.Int, inferenceId []*big.Int, miner []common.Address) (*IWorkerHubNewAssignmentIterator, error) {

	var assignmentIdRule []interface{}
	for _, assignmentIdItem := range assignmentId {
		assignmentIdRule = append(assignmentIdRule, assignmentIdItem)
	}
	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "NewAssignment", assignmentIdRule, inferenceIdRule, minerRule)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubNewAssignmentIterator{contract: _IWorkerHub.contract, event: "NewAssignment", logs: logs, sub: sub}, nil
}

// WatchNewAssignment is a free log subscription operation binding the contract event 0x53cc8b652f33c56dac5f1c97a284cc971e7adcb8abe9454b0853f076c6deb7d5.
//
// Solidity: event NewAssignment(uint256 indexed assignmentId, uint256 indexed inferenceId, address indexed miner, uint40 expiredAt)
func (_IWorkerHub *IWorkerHubFilterer) WatchNewAssignment(opts *bind.WatchOpts, sink chan<- *IWorkerHubNewAssignment, assignmentId []*big.Int, inferenceId []*big.Int, miner []common.Address) (event.Subscription, error) {

	var assignmentIdRule []interface{}
	for _, assignmentIdItem := range assignmentId {
		assignmentIdRule = append(assignmentIdRule, assignmentIdItem)
	}
	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "NewAssignment", assignmentIdRule, inferenceIdRule, minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubNewAssignment)
				if err := _IWorkerHub.contract.UnpackLog(event, "NewAssignment", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewAssignment is a log parse operation binding the contract event 0x53cc8b652f33c56dac5f1c97a284cc971e7adcb8abe9454b0853f076c6deb7d5.
//
// Solidity: event NewAssignment(uint256 indexed assignmentId, uint256 indexed inferenceId, address indexed miner, uint40 expiredAt)
func (_IWorkerHub *IWorkerHubFilterer) ParseNewAssignment(log types.Log) (*IWorkerHubNewAssignment, error) {
	event := new(IWorkerHubNewAssignment)
	if err := _IWorkerHub.contract.UnpackLog(event, "NewAssignment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubNewInferenceIterator is returned from FilterNewInference and is used to iterate over the raw logs and unpacked data for NewInference events raised by the IWorkerHub contract.
type IWorkerHubNewInferenceIterator struct {
	Event *IWorkerHubNewInference // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubNewInferenceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubNewInference)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubNewInference)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubNewInferenceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubNewInferenceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubNewInference represents a NewInference event raised by the IWorkerHub contract.
type IWorkerHubNewInference struct {
	InferenceId       *big.Int
	Model             common.Address
	Creator           common.Address
	Value             *big.Int
	OriginInferenceId *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewInference is a free log retrieval operation binding the contract event 0x08a84d7fb7cd1557f228c827b9280f44d1a157c3256fe453b687a7b9d51c6a5b.
//
// Solidity: event NewInference(uint256 indexed inferenceId, address indexed model, address indexed creator, uint256 value, uint256 originInferenceId)
func (_IWorkerHub *IWorkerHubFilterer) FilterNewInference(opts *bind.FilterOpts, inferenceId []*big.Int, model []common.Address, creator []common.Address) (*IWorkerHubNewInferenceIterator, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "NewInference", inferenceIdRule, modelRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubNewInferenceIterator{contract: _IWorkerHub.contract, event: "NewInference", logs: logs, sub: sub}, nil
}

// WatchNewInference is a free log subscription operation binding the contract event 0x08a84d7fb7cd1557f228c827b9280f44d1a157c3256fe453b687a7b9d51c6a5b.
//
// Solidity: event NewInference(uint256 indexed inferenceId, address indexed model, address indexed creator, uint256 value, uint256 originInferenceId)
func (_IWorkerHub *IWorkerHubFilterer) WatchNewInference(opts *bind.WatchOpts, sink chan<- *IWorkerHubNewInference, inferenceId []*big.Int, model []common.Address, creator []common.Address) (event.Subscription, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "NewInference", inferenceIdRule, modelRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubNewInference)
				if err := _IWorkerHub.contract.UnpackLog(event, "NewInference", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewInference is a log parse operation binding the contract event 0x08a84d7fb7cd1557f228c827b9280f44d1a157c3256fe453b687a7b9d51c6a5b.
//
// Solidity: event NewInference(uint256 indexed inferenceId, address indexed model, address indexed creator, uint256 value, uint256 originInferenceId)
func (_IWorkerHub *IWorkerHubFilterer) ParseNewInference(log types.Log) (*IWorkerHubNewInference, error) {
	event := new(IWorkerHubNewInference)
	if err := _IWorkerHub.contract.UnpackLog(event, "NewInference", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubNewScoringInferenceIterator is returned from FilterNewScoringInference and is used to iterate over the raw logs and unpacked data for NewScoringInference events raised by the IWorkerHub contract.
type IWorkerHubNewScoringInferenceIterator struct {
	Event *IWorkerHubNewScoringInference // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubNewScoringInferenceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubNewScoringInference)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubNewScoringInference)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubNewScoringInferenceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubNewScoringInferenceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubNewScoringInference represents a NewScoringInference event raised by the IWorkerHub contract.
type IWorkerHubNewScoringInference struct {
	InferenceId       *big.Int
	Model             common.Address
	Creator           common.Address
	Value             *big.Int
	OriginInferenceId *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewScoringInference is a free log retrieval operation binding the contract event 0x3ec54c04f8c304e8caa7314d1ac4d34bff1c57151f207745b19e6d8f0a579ea9.
//
// Solidity: event NewScoringInference(uint256 indexed inferenceId, address indexed model, address indexed creator, uint256 value, uint256 originInferenceId)
func (_IWorkerHub *IWorkerHubFilterer) FilterNewScoringInference(opts *bind.FilterOpts, inferenceId []*big.Int, model []common.Address, creator []common.Address) (*IWorkerHubNewScoringInferenceIterator, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "NewScoringInference", inferenceIdRule, modelRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubNewScoringInferenceIterator{contract: _IWorkerHub.contract, event: "NewScoringInference", logs: logs, sub: sub}, nil
}

// WatchNewScoringInference is a free log subscription operation binding the contract event 0x3ec54c04f8c304e8caa7314d1ac4d34bff1c57151f207745b19e6d8f0a579ea9.
//
// Solidity: event NewScoringInference(uint256 indexed inferenceId, address indexed model, address indexed creator, uint256 value, uint256 originInferenceId)
func (_IWorkerHub *IWorkerHubFilterer) WatchNewScoringInference(opts *bind.WatchOpts, sink chan<- *IWorkerHubNewScoringInference, inferenceId []*big.Int, model []common.Address, creator []common.Address) (event.Subscription, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "NewScoringInference", inferenceIdRule, modelRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubNewScoringInference)
				if err := _IWorkerHub.contract.UnpackLog(event, "NewScoringInference", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewScoringInference is a log parse operation binding the contract event 0x3ec54c04f8c304e8caa7314d1ac4d34bff1c57151f207745b19e6d8f0a579ea9.
//
// Solidity: event NewScoringInference(uint256 indexed inferenceId, address indexed model, address indexed creator, uint256 value, uint256 originInferenceId)
func (_IWorkerHub *IWorkerHubFilterer) ParseNewScoringInference(log types.Log) (*IWorkerHubNewScoringInference, error) {
	event := new(IWorkerHubNewScoringInference)
	if err := _IWorkerHub.contract.UnpackLog(event, "NewScoringInference", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IWorkerHub contract.
type IWorkerHubOwnershipTransferredIterator struct {
	Event *IWorkerHubOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubOwnershipTransferred represents a OwnershipTransferred event raised by the IWorkerHub contract.
type IWorkerHubOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IWorkerHub *IWorkerHubFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IWorkerHubOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubOwnershipTransferredIterator{contract: _IWorkerHub.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IWorkerHub *IWorkerHubFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IWorkerHubOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubOwnershipTransferred)
				if err := _IWorkerHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IWorkerHub *IWorkerHubFilterer) ParseOwnershipTransferred(log types.Log) (*IWorkerHubOwnershipTransferred, error) {
	event := new(IWorkerHubOwnershipTransferred)
	if err := _IWorkerHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the IWorkerHub contract.
type IWorkerHubPausedIterator struct {
	Event *IWorkerHubPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubPaused represents a Paused event raised by the IWorkerHub contract.
type IWorkerHubPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IWorkerHub *IWorkerHubFilterer) FilterPaused(opts *bind.FilterOpts) (*IWorkerHubPausedIterator, error) {

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &IWorkerHubPausedIterator{contract: _IWorkerHub.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IWorkerHub *IWorkerHubFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *IWorkerHubPaused) (event.Subscription, error) {

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubPaused)
				if err := _IWorkerHub.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IWorkerHub *IWorkerHubFilterer) ParsePaused(log types.Log) (*IWorkerHubPaused, error) {
	event := new(IWorkerHubPaused)
	if err := _IWorkerHub.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubPenaltyDurationUpdatedIterator is returned from FilterPenaltyDurationUpdated and is used to iterate over the raw logs and unpacked data for PenaltyDurationUpdated events raised by the IWorkerHub contract.
type IWorkerHubPenaltyDurationUpdatedIterator struct {
	Event *IWorkerHubPenaltyDurationUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubPenaltyDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubPenaltyDurationUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubPenaltyDurationUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubPenaltyDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubPenaltyDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubPenaltyDurationUpdated represents a PenaltyDurationUpdated event raised by the IWorkerHub contract.
type IWorkerHubPenaltyDurationUpdated struct {
	OldDuration *big.Int
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPenaltyDurationUpdated is a free log retrieval operation binding the contract event 0xf7a437a25c636d2b29d0ba34f0f6870af14f44478eff2ac852f36030f2e2924e.
//
// Solidity: event PenaltyDurationUpdated(uint40 oldDuration, uint40 newDuration)
func (_IWorkerHub *IWorkerHubFilterer) FilterPenaltyDurationUpdated(opts *bind.FilterOpts) (*IWorkerHubPenaltyDurationUpdatedIterator, error) {

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "PenaltyDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &IWorkerHubPenaltyDurationUpdatedIterator{contract: _IWorkerHub.contract, event: "PenaltyDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchPenaltyDurationUpdated is a free log subscription operation binding the contract event 0xf7a437a25c636d2b29d0ba34f0f6870af14f44478eff2ac852f36030f2e2924e.
//
// Solidity: event PenaltyDurationUpdated(uint40 oldDuration, uint40 newDuration)
func (_IWorkerHub *IWorkerHubFilterer) WatchPenaltyDurationUpdated(opts *bind.WatchOpts, sink chan<- *IWorkerHubPenaltyDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "PenaltyDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubPenaltyDurationUpdated)
				if err := _IWorkerHub.contract.UnpackLog(event, "PenaltyDurationUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePenaltyDurationUpdated is a log parse operation binding the contract event 0xf7a437a25c636d2b29d0ba34f0f6870af14f44478eff2ac852f36030f2e2924e.
//
// Solidity: event PenaltyDurationUpdated(uint40 oldDuration, uint40 newDuration)
func (_IWorkerHub *IWorkerHubFilterer) ParsePenaltyDurationUpdated(log types.Log) (*IWorkerHubPenaltyDurationUpdated, error) {
	event := new(IWorkerHubPenaltyDurationUpdated)
	if err := _IWorkerHub.contract.UnpackLog(event, "PenaltyDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubRestakeIterator is returned from FilterRestake and is used to iterate over the raw logs and unpacked data for Restake events raised by the IWorkerHub contract.
type IWorkerHubRestakeIterator struct {
	Event *IWorkerHubRestake // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubRestakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubRestake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubRestake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubRestakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubRestakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubRestake represents a Restake event raised by the IWorkerHub contract.
type IWorkerHubRestake struct {
	Miner   common.Address
	Restake *big.Int
	Model   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRestake is a free log retrieval operation binding the contract event 0x5f8a19f664e489b0ebcc62ec24b1bde029195fbb4af60118cecf0e16d6d95b2d.
//
// Solidity: event Restake(address indexed miner, uint256 restake, address indexed model)
func (_IWorkerHub *IWorkerHubFilterer) FilterRestake(opts *bind.FilterOpts, miner []common.Address, model []common.Address) (*IWorkerHubRestakeIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "Restake", minerRule, modelRule)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubRestakeIterator{contract: _IWorkerHub.contract, event: "Restake", logs: logs, sub: sub}, nil
}

// WatchRestake is a free log subscription operation binding the contract event 0x5f8a19f664e489b0ebcc62ec24b1bde029195fbb4af60118cecf0e16d6d95b2d.
//
// Solidity: event Restake(address indexed miner, uint256 restake, address indexed model)
func (_IWorkerHub *IWorkerHubFilterer) WatchRestake(opts *bind.WatchOpts, sink chan<- *IWorkerHubRestake, miner []common.Address, model []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	var modelRule []interface{}
	for _, modelItem := range model {
		modelRule = append(modelRule, modelItem)
	}

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "Restake", minerRule, modelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubRestake)
				if err := _IWorkerHub.contract.UnpackLog(event, "Restake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRestake is a log parse operation binding the contract event 0x5f8a19f664e489b0ebcc62ec24b1bde029195fbb4af60118cecf0e16d6d95b2d.
//
// Solidity: event Restake(address indexed miner, uint256 restake, address indexed model)
func (_IWorkerHub *IWorkerHubFilterer) ParseRestake(log types.Log) (*IWorkerHubRestake, error) {
	event := new(IWorkerHubRestake)
	if err := _IWorkerHub.contract.UnpackLog(event, "Restake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubRevealDurationIterator is returned from FilterRevealDuration and is used to iterate over the raw logs and unpacked data for RevealDuration events raised by the IWorkerHub contract.
type IWorkerHubRevealDurationIterator struct {
	Event *IWorkerHubRevealDuration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubRevealDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubRevealDuration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubRevealDuration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubRevealDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubRevealDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubRevealDuration represents a RevealDuration event raised by the IWorkerHub contract.
type IWorkerHubRevealDuration struct {
	OldTime *big.Int
	NewTime *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRevealDuration is a free log retrieval operation binding the contract event 0xacb24019039b4d00193b2be5c85ea8ed6bd6747ed79f7d1e5a6d9384282b4a9d.
//
// Solidity: event RevealDuration(uint256 oldTime, uint256 newTime)
func (_IWorkerHub *IWorkerHubFilterer) FilterRevealDuration(opts *bind.FilterOpts) (*IWorkerHubRevealDurationIterator, error) {

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "RevealDuration")
	if err != nil {
		return nil, err
	}
	return &IWorkerHubRevealDurationIterator{contract: _IWorkerHub.contract, event: "RevealDuration", logs: logs, sub: sub}, nil
}

// WatchRevealDuration is a free log subscription operation binding the contract event 0xacb24019039b4d00193b2be5c85ea8ed6bd6747ed79f7d1e5a6d9384282b4a9d.
//
// Solidity: event RevealDuration(uint256 oldTime, uint256 newTime)
func (_IWorkerHub *IWorkerHubFilterer) WatchRevealDuration(opts *bind.WatchOpts, sink chan<- *IWorkerHubRevealDuration) (event.Subscription, error) {

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "RevealDuration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubRevealDuration)
				if err := _IWorkerHub.contract.UnpackLog(event, "RevealDuration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevealDuration is a log parse operation binding the contract event 0xacb24019039b4d00193b2be5c85ea8ed6bd6747ed79f7d1e5a6d9384282b4a9d.
//
// Solidity: event RevealDuration(uint256 oldTime, uint256 newTime)
func (_IWorkerHub *IWorkerHubFilterer) ParseRevealDuration(log types.Log) (*IWorkerHubRevealDuration, error) {
	event := new(IWorkerHubRevealDuration)
	if err := _IWorkerHub.contract.UnpackLog(event, "RevealDuration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubRevealSubmissionIterator is returned from FilterRevealSubmission and is used to iterate over the raw logs and unpacked data for RevealSubmission events raised by the IWorkerHub contract.
type IWorkerHubRevealSubmissionIterator struct {
	Event *IWorkerHubRevealSubmission // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubRevealSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubRevealSubmission)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubRevealSubmission)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubRevealSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubRevealSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubRevealSubmission represents a RevealSubmission event raised by the IWorkerHub contract.
type IWorkerHubRevealSubmission struct {
	Miner       common.Address
	AssigmentId *big.Int
	Nonce       *big.Int
	Output      []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRevealSubmission is a free log retrieval operation binding the contract event 0xf7e30468a493d9e17158c0dbe51bcfa190627e3fdede3c9284827c22dfc41700.
//
// Solidity: event RevealSubmission(address indexed miner, uint256 indexed assigmentId, uint40 nonce, bytes output)
func (_IWorkerHub *IWorkerHubFilterer) FilterRevealSubmission(opts *bind.FilterOpts, miner []common.Address, assigmentId []*big.Int) (*IWorkerHubRevealSubmissionIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var assigmentIdRule []interface{}
	for _, assigmentIdItem := range assigmentId {
		assigmentIdRule = append(assigmentIdRule, assigmentIdItem)
	}

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "RevealSubmission", minerRule, assigmentIdRule)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubRevealSubmissionIterator{contract: _IWorkerHub.contract, event: "RevealSubmission", logs: logs, sub: sub}, nil
}

// WatchRevealSubmission is a free log subscription operation binding the contract event 0xf7e30468a493d9e17158c0dbe51bcfa190627e3fdede3c9284827c22dfc41700.
//
// Solidity: event RevealSubmission(address indexed miner, uint256 indexed assigmentId, uint40 nonce, bytes output)
func (_IWorkerHub *IWorkerHubFilterer) WatchRevealSubmission(opts *bind.WatchOpts, sink chan<- *IWorkerHubRevealSubmission, miner []common.Address, assigmentId []*big.Int) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var assigmentIdRule []interface{}
	for _, assigmentIdItem := range assigmentId {
		assigmentIdRule = append(assigmentIdRule, assigmentIdItem)
	}

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "RevealSubmission", minerRule, assigmentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubRevealSubmission)
				if err := _IWorkerHub.contract.UnpackLog(event, "RevealSubmission", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevealSubmission is a log parse operation binding the contract event 0xf7e30468a493d9e17158c0dbe51bcfa190627e3fdede3c9284827c22dfc41700.
//
// Solidity: event RevealSubmission(address indexed miner, uint256 indexed assigmentId, uint40 nonce, bytes output)
func (_IWorkerHub *IWorkerHubFilterer) ParseRevealSubmission(log types.Log) (*IWorkerHubRevealSubmission, error) {
	event := new(IWorkerHubRevealSubmission)
	if err := _IWorkerHub.contract.UnpackLog(event, "RevealSubmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubRewardClaimIterator is returned from FilterRewardClaim and is used to iterate over the raw logs and unpacked data for RewardClaim events raised by the IWorkerHub contract.
type IWorkerHubRewardClaimIterator struct {
	Event *IWorkerHubRewardClaim // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubRewardClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubRewardClaim)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubRewardClaim)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubRewardClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubRewardClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubRewardClaim represents a RewardClaim event raised by the IWorkerHub contract.
type IWorkerHubRewardClaim struct {
	Worker common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardClaim is a free log retrieval operation binding the contract event 0x75690555e75b04e280e646889defdcbefd8401507e5394d1173fd84290944c29.
//
// Solidity: event RewardClaim(address indexed worker, uint256 value)
func (_IWorkerHub *IWorkerHubFilterer) FilterRewardClaim(opts *bind.FilterOpts, worker []common.Address) (*IWorkerHubRewardClaimIterator, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "RewardClaim", workerRule)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubRewardClaimIterator{contract: _IWorkerHub.contract, event: "RewardClaim", logs: logs, sub: sub}, nil
}

// WatchRewardClaim is a free log subscription operation binding the contract event 0x75690555e75b04e280e646889defdcbefd8401507e5394d1173fd84290944c29.
//
// Solidity: event RewardClaim(address indexed worker, uint256 value)
func (_IWorkerHub *IWorkerHubFilterer) WatchRewardClaim(opts *bind.WatchOpts, sink chan<- *IWorkerHubRewardClaim, worker []common.Address) (event.Subscription, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "RewardClaim", workerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubRewardClaim)
				if err := _IWorkerHub.contract.UnpackLog(event, "RewardClaim", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardClaim is a log parse operation binding the contract event 0x75690555e75b04e280e646889defdcbefd8401507e5394d1173fd84290944c29.
//
// Solidity: event RewardClaim(address indexed worker, uint256 value)
func (_IWorkerHub *IWorkerHubFilterer) ParseRewardClaim(log types.Log) (*IWorkerHubRewardClaim, error) {
	event := new(IWorkerHubRewardClaim)
	if err := _IWorkerHub.contract.UnpackLog(event, "RewardClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubRewardPerEpochIterator is returned from FilterRewardPerEpoch and is used to iterate over the raw logs and unpacked data for RewardPerEpoch events raised by the IWorkerHub contract.
type IWorkerHubRewardPerEpochIterator struct {
	Event *IWorkerHubRewardPerEpoch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubRewardPerEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubRewardPerEpoch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubRewardPerEpoch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubRewardPerEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubRewardPerEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubRewardPerEpoch represents a RewardPerEpoch event raised by the IWorkerHub contract.
type IWorkerHubRewardPerEpoch struct {
	OldReward *big.Int
	NewReward *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardPerEpoch is a free log retrieval operation binding the contract event 0x3d731857045dfa7982ed8ff308eeda54c7e156ba99609db02c50b4485f64c463.
//
// Solidity: event RewardPerEpoch(uint256 oldReward, uint256 newReward)
func (_IWorkerHub *IWorkerHubFilterer) FilterRewardPerEpoch(opts *bind.FilterOpts) (*IWorkerHubRewardPerEpochIterator, error) {

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "RewardPerEpoch")
	if err != nil {
		return nil, err
	}
	return &IWorkerHubRewardPerEpochIterator{contract: _IWorkerHub.contract, event: "RewardPerEpoch", logs: logs, sub: sub}, nil
}

// WatchRewardPerEpoch is a free log subscription operation binding the contract event 0x3d731857045dfa7982ed8ff308eeda54c7e156ba99609db02c50b4485f64c463.
//
// Solidity: event RewardPerEpoch(uint256 oldReward, uint256 newReward)
func (_IWorkerHub *IWorkerHubFilterer) WatchRewardPerEpoch(opts *bind.WatchOpts, sink chan<- *IWorkerHubRewardPerEpoch) (event.Subscription, error) {

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "RewardPerEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubRewardPerEpoch)
				if err := _IWorkerHub.contract.UnpackLog(event, "RewardPerEpoch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardPerEpoch is a log parse operation binding the contract event 0x3d731857045dfa7982ed8ff308eeda54c7e156ba99609db02c50b4485f64c463.
//
// Solidity: event RewardPerEpoch(uint256 oldReward, uint256 newReward)
func (_IWorkerHub *IWorkerHubFilterer) ParseRewardPerEpoch(log types.Log) (*IWorkerHubRewardPerEpoch, error) {
	event := new(IWorkerHubRewardPerEpoch)
	if err := _IWorkerHub.contract.UnpackLog(event, "RewardPerEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubSolutionSubmissionIterator is returned from FilterSolutionSubmission and is used to iterate over the raw logs and unpacked data for SolutionSubmission events raised by the IWorkerHub contract.
type IWorkerHubSolutionSubmissionIterator struct {
	Event *IWorkerHubSolutionSubmission // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubSolutionSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubSolutionSubmission)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubSolutionSubmission)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubSolutionSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubSolutionSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubSolutionSubmission represents a SolutionSubmission event raised by the IWorkerHub contract.
type IWorkerHubSolutionSubmission struct {
	Miner       common.Address
	AssigmentId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSolutionSubmission is a free log retrieval operation binding the contract event 0x9f669b92b9cbc7611f7ab6c77db07a424051c777433e21bd90f1bdf940096dd9.
//
// Solidity: event SolutionSubmission(address indexed miner, uint256 indexed assigmentId)
func (_IWorkerHub *IWorkerHubFilterer) FilterSolutionSubmission(opts *bind.FilterOpts, miner []common.Address, assigmentId []*big.Int) (*IWorkerHubSolutionSubmissionIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var assigmentIdRule []interface{}
	for _, assigmentIdItem := range assigmentId {
		assigmentIdRule = append(assigmentIdRule, assigmentIdItem)
	}

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "SolutionSubmission", minerRule, assigmentIdRule)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubSolutionSubmissionIterator{contract: _IWorkerHub.contract, event: "SolutionSubmission", logs: logs, sub: sub}, nil
}

// WatchSolutionSubmission is a free log subscription operation binding the contract event 0x9f669b92b9cbc7611f7ab6c77db07a424051c777433e21bd90f1bdf940096dd9.
//
// Solidity: event SolutionSubmission(address indexed miner, uint256 indexed assigmentId)
func (_IWorkerHub *IWorkerHubFilterer) WatchSolutionSubmission(opts *bind.WatchOpts, sink chan<- *IWorkerHubSolutionSubmission, miner []common.Address, assigmentId []*big.Int) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var assigmentIdRule []interface{}
	for _, assigmentIdItem := range assigmentId {
		assigmentIdRule = append(assigmentIdRule, assigmentIdItem)
	}

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "SolutionSubmission", minerRule, assigmentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubSolutionSubmission)
				if err := _IWorkerHub.contract.UnpackLog(event, "SolutionSubmission", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSolutionSubmission is a log parse operation binding the contract event 0x9f669b92b9cbc7611f7ab6c77db07a424051c777433e21bd90f1bdf940096dd9.
//
// Solidity: event SolutionSubmission(address indexed miner, uint256 indexed assigmentId)
func (_IWorkerHub *IWorkerHubFilterer) ParseSolutionSubmission(log types.Log) (*IWorkerHubSolutionSubmission, error) {
	event := new(IWorkerHubSolutionSubmission)
	if err := _IWorkerHub.contract.UnpackLog(event, "SolutionSubmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubStreamedDataIterator is returned from FilterStreamedData and is used to iterate over the raw logs and unpacked data for StreamedData events raised by the IWorkerHub contract.
type IWorkerHubStreamedDataIterator struct {
	Event *IWorkerHubStreamedData // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubStreamedDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubStreamedData)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubStreamedData)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubStreamedDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubStreamedDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubStreamedData represents a StreamedData event raised by the IWorkerHub contract.
type IWorkerHubStreamedData struct {
	AssignmentId *big.Int
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStreamedData is a free log retrieval operation binding the contract event 0x23cfaa418b5f569ff36b152a9fd02ee3ccddaa5f7eed570e777a30353b68dc38.
//
// Solidity: event StreamedData(uint256 indexed assignmentId, bytes data)
func (_IWorkerHub *IWorkerHubFilterer) FilterStreamedData(opts *bind.FilterOpts, assignmentId []*big.Int) (*IWorkerHubStreamedDataIterator, error) {

	var assignmentIdRule []interface{}
	for _, assignmentIdItem := range assignmentId {
		assignmentIdRule = append(assignmentIdRule, assignmentIdItem)
	}

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "StreamedData", assignmentIdRule)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubStreamedDataIterator{contract: _IWorkerHub.contract, event: "StreamedData", logs: logs, sub: sub}, nil
}

// WatchStreamedData is a free log subscription operation binding the contract event 0x23cfaa418b5f569ff36b152a9fd02ee3ccddaa5f7eed570e777a30353b68dc38.
//
// Solidity: event StreamedData(uint256 indexed assignmentId, bytes data)
func (_IWorkerHub *IWorkerHubFilterer) WatchStreamedData(opts *bind.WatchOpts, sink chan<- *IWorkerHubStreamedData, assignmentId []*big.Int) (event.Subscription, error) {

	var assignmentIdRule []interface{}
	for _, assignmentIdItem := range assignmentId {
		assignmentIdRule = append(assignmentIdRule, assignmentIdItem)
	}

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "StreamedData", assignmentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubStreamedData)
				if err := _IWorkerHub.contract.UnpackLog(event, "StreamedData", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStreamedData is a log parse operation binding the contract event 0x23cfaa418b5f569ff36b152a9fd02ee3ccddaa5f7eed570e777a30353b68dc38.
//
// Solidity: event StreamedData(uint256 indexed assignmentId, bytes data)
func (_IWorkerHub *IWorkerHubFilterer) ParseStreamedData(log types.Log) (*IWorkerHubStreamedData, error) {
	event := new(IWorkerHubStreamedData)
	if err := _IWorkerHub.contract.UnpackLog(event, "StreamedData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubSubmitDurationIterator is returned from FilterSubmitDuration and is used to iterate over the raw logs and unpacked data for SubmitDuration events raised by the IWorkerHub contract.
type IWorkerHubSubmitDurationIterator struct {
	Event *IWorkerHubSubmitDuration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubSubmitDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubSubmitDuration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubSubmitDuration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubSubmitDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubSubmitDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubSubmitDuration represents a SubmitDuration event raised by the IWorkerHub contract.
type IWorkerHubSubmitDuration struct {
	OldTime *big.Int
	NewTime *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSubmitDuration is a free log retrieval operation binding the contract event 0x8c0ac957fb32132ec541e9495c4fe8f1d9fdb4dd19a02e7144659d4b382064f3.
//
// Solidity: event SubmitDuration(uint256 oldTime, uint256 newTime)
func (_IWorkerHub *IWorkerHubFilterer) FilterSubmitDuration(opts *bind.FilterOpts) (*IWorkerHubSubmitDurationIterator, error) {

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "SubmitDuration")
	if err != nil {
		return nil, err
	}
	return &IWorkerHubSubmitDurationIterator{contract: _IWorkerHub.contract, event: "SubmitDuration", logs: logs, sub: sub}, nil
}

// WatchSubmitDuration is a free log subscription operation binding the contract event 0x8c0ac957fb32132ec541e9495c4fe8f1d9fdb4dd19a02e7144659d4b382064f3.
//
// Solidity: event SubmitDuration(uint256 oldTime, uint256 newTime)
func (_IWorkerHub *IWorkerHubFilterer) WatchSubmitDuration(opts *bind.WatchOpts, sink chan<- *IWorkerHubSubmitDuration) (event.Subscription, error) {

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "SubmitDuration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubSubmitDuration)
				if err := _IWorkerHub.contract.UnpackLog(event, "SubmitDuration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubmitDuration is a log parse operation binding the contract event 0x8c0ac957fb32132ec541e9495c4fe8f1d9fdb4dd19a02e7144659d4b382064f3.
//
// Solidity: event SubmitDuration(uint256 oldTime, uint256 newTime)
func (_IWorkerHub *IWorkerHubFilterer) ParseSubmitDuration(log types.Log) (*IWorkerHubSubmitDuration, error) {
	event := new(IWorkerHubSubmitDuration)
	if err := _IWorkerHub.contract.UnpackLog(event, "SubmitDuration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubTopUpInferIterator is returned from FilterTopUpInfer and is used to iterate over the raw logs and unpacked data for TopUpInfer events raised by the IWorkerHub contract.
type IWorkerHubTopUpInferIterator struct {
	Event *IWorkerHubTopUpInfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubTopUpInferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubTopUpInfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubTopUpInfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubTopUpInferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubTopUpInferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubTopUpInfer represents a TopUpInfer event raised by the IWorkerHub contract.
type IWorkerHubTopUpInfer struct {
	InferenceId *big.Int
	Creator     common.Address
	Value       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTopUpInfer is a free log retrieval operation binding the contract event 0xe3154336ce264fe53bcfaedafded1428a28ae47b19b3d7a82e5d5ecde0960a57.
//
// Solidity: event TopUpInfer(uint256 indexed inferenceId, address indexed creator, uint256 value)
func (_IWorkerHub *IWorkerHubFilterer) FilterTopUpInfer(opts *bind.FilterOpts, inferenceId []*big.Int, creator []common.Address) (*IWorkerHubTopUpInferIterator, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "TopUpInfer", inferenceIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubTopUpInferIterator{contract: _IWorkerHub.contract, event: "TopUpInfer", logs: logs, sub: sub}, nil
}

// WatchTopUpInfer is a free log subscription operation binding the contract event 0xe3154336ce264fe53bcfaedafded1428a28ae47b19b3d7a82e5d5ecde0960a57.
//
// Solidity: event TopUpInfer(uint256 indexed inferenceId, address indexed creator, uint256 value)
func (_IWorkerHub *IWorkerHubFilterer) WatchTopUpInfer(opts *bind.WatchOpts, sink chan<- *IWorkerHubTopUpInfer, inferenceId []*big.Int, creator []common.Address) (event.Subscription, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "TopUpInfer", inferenceIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubTopUpInfer)
				if err := _IWorkerHub.contract.UnpackLog(event, "TopUpInfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTopUpInfer is a log parse operation binding the contract event 0xe3154336ce264fe53bcfaedafded1428a28ae47b19b3d7a82e5d5ecde0960a57.
//
// Solidity: event TopUpInfer(uint256 indexed inferenceId, address indexed creator, uint256 value)
func (_IWorkerHub *IWorkerHubFilterer) ParseTopUpInfer(log types.Log) (*IWorkerHubTopUpInfer, error) {
	event := new(IWorkerHubTopUpInfer)
	if err := _IWorkerHub.contract.UnpackLog(event, "TopUpInfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubTransferFeeIterator is returned from FilterTransferFee and is used to iterate over the raw logs and unpacked data for TransferFee events raised by the IWorkerHub contract.
type IWorkerHubTransferFeeIterator struct {
	Event *IWorkerHubTransferFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubTransferFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubTransferFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubTransferFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubTransferFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubTransferFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubTransferFee represents a TransferFee event raised by the IWorkerHub contract.
type IWorkerHubTransferFee struct {
	Treasury       common.Address
	TreasuryFee    *big.Int
	L2OwnerAddress common.Address
	L2OwnerFee     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTransferFee is a free log retrieval operation binding the contract event 0x782aada659bac972b342fea00dfc27389e876bece89a9eb635bd5a2c544e8a6b.
//
// Solidity: event TransferFee(address indexed treasury, uint256 treasuryFee, address L2OwnerAddress, uint256 L2OwnerFee)
func (_IWorkerHub *IWorkerHubFilterer) FilterTransferFee(opts *bind.FilterOpts, treasury []common.Address) (*IWorkerHubTransferFeeIterator, error) {

	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "TransferFee", treasuryRule)
	if err != nil {
		return nil, err
	}
	return &IWorkerHubTransferFeeIterator{contract: _IWorkerHub.contract, event: "TransferFee", logs: logs, sub: sub}, nil
}

// WatchTransferFee is a free log subscription operation binding the contract event 0x782aada659bac972b342fea00dfc27389e876bece89a9eb635bd5a2c544e8a6b.
//
// Solidity: event TransferFee(address indexed treasury, uint256 treasuryFee, address L2OwnerAddress, uint256 L2OwnerFee)
func (_IWorkerHub *IWorkerHubFilterer) WatchTransferFee(opts *bind.WatchOpts, sink chan<- *IWorkerHubTransferFee, treasury []common.Address) (event.Subscription, error) {

	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "TransferFee", treasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubTransferFee)
				if err := _IWorkerHub.contract.UnpackLog(event, "TransferFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferFee is a log parse operation binding the contract event 0x782aada659bac972b342fea00dfc27389e876bece89a9eb635bd5a2c544e8a6b.
//
// Solidity: event TransferFee(address indexed treasury, uint256 treasuryFee, address L2OwnerAddress, uint256 L2OwnerFee)
func (_IWorkerHub *IWorkerHubFilterer) ParseTransferFee(log types.Log) (*IWorkerHubTransferFee, error) {
	event := new(IWorkerHubTransferFee)
	if err := _IWorkerHub.contract.UnpackLog(event, "TransferFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubTreasuryAddressUpdatedIterator is returned from FilterTreasuryAddressUpdated and is used to iterate over the raw logs and unpacked data for TreasuryAddressUpdated events raised by the IWorkerHub contract.
type IWorkerHubTreasuryAddressUpdatedIterator struct {
	Event *IWorkerHubTreasuryAddressUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubTreasuryAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubTreasuryAddressUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubTreasuryAddressUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubTreasuryAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubTreasuryAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubTreasuryAddressUpdated represents a TreasuryAddressUpdated event raised by the IWorkerHub contract.
type IWorkerHubTreasuryAddressUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTreasuryAddressUpdated is a free log retrieval operation binding the contract event 0x430359a6d97ced2b6f93c77a91e7ce9dfd43252eb91e916adba170485cd8a6a4.
//
// Solidity: event TreasuryAddressUpdated(address oldAddress, address newAddress)
func (_IWorkerHub *IWorkerHubFilterer) FilterTreasuryAddressUpdated(opts *bind.FilterOpts) (*IWorkerHubTreasuryAddressUpdatedIterator, error) {

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "TreasuryAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &IWorkerHubTreasuryAddressUpdatedIterator{contract: _IWorkerHub.contract, event: "TreasuryAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchTreasuryAddressUpdated is a free log subscription operation binding the contract event 0x430359a6d97ced2b6f93c77a91e7ce9dfd43252eb91e916adba170485cd8a6a4.
//
// Solidity: event TreasuryAddressUpdated(address oldAddress, address newAddress)
func (_IWorkerHub *IWorkerHubFilterer) WatchTreasuryAddressUpdated(opts *bind.WatchOpts, sink chan<- *IWorkerHubTreasuryAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "TreasuryAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubTreasuryAddressUpdated)
				if err := _IWorkerHub.contract.UnpackLog(event, "TreasuryAddressUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTreasuryAddressUpdated is a log parse operation binding the contract event 0x430359a6d97ced2b6f93c77a91e7ce9dfd43252eb91e916adba170485cd8a6a4.
//
// Solidity: event TreasuryAddressUpdated(address oldAddress, address newAddress)
func (_IWorkerHub *IWorkerHubFilterer) ParseTreasuryAddressUpdated(log types.Log) (*IWorkerHubTreasuryAddressUpdated, error) {
	event := new(IWorkerHubTreasuryAddressUpdated)
	if err := _IWorkerHub.contract.UnpackLog(event, "TreasuryAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the IWorkerHub contract.
type IWorkerHubUnpausedIterator struct {
	Event *IWorkerHubUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubUnpaused represents a Unpaused event raised by the IWorkerHub contract.
type IWorkerHubUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IWorkerHub *IWorkerHubFilterer) FilterUnpaused(opts *bind.FilterOpts) (*IWorkerHubUnpausedIterator, error) {

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &IWorkerHubUnpausedIterator{contract: _IWorkerHub.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IWorkerHub *IWorkerHubFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *IWorkerHubUnpaused) (event.Subscription, error) {

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubUnpaused)
				if err := _IWorkerHub.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IWorkerHub *IWorkerHubFilterer) ParseUnpaused(log types.Log) (*IWorkerHubUnpaused, error) {
	event := new(IWorkerHubUnpaused)
	if err := _IWorkerHub.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorkerHubUnstakeDelayTimeIterator is returned from FilterUnstakeDelayTime and is used to iterate over the raw logs and unpacked data for UnstakeDelayTime events raised by the IWorkerHub contract.
type IWorkerHubUnstakeDelayTimeIterator struct {
	Event *IWorkerHubUnstakeDelayTime // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IWorkerHubUnstakeDelayTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorkerHubUnstakeDelayTime)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IWorkerHubUnstakeDelayTime)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IWorkerHubUnstakeDelayTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorkerHubUnstakeDelayTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorkerHubUnstakeDelayTime represents a UnstakeDelayTime event raised by the IWorkerHub contract.
type IWorkerHubUnstakeDelayTime struct {
	OldDelayTime *big.Int
	NewDelayTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUnstakeDelayTime is a free log retrieval operation binding the contract event 0xdf63c46e5024e57c66aafc6698e317c78589c870dca694678c89dd379c5fd490.
//
// Solidity: event UnstakeDelayTime(uint256 oldDelayTime, uint256 newDelayTime)
func (_IWorkerHub *IWorkerHubFilterer) FilterUnstakeDelayTime(opts *bind.FilterOpts) (*IWorkerHubUnstakeDelayTimeIterator, error) {

	logs, sub, err := _IWorkerHub.contract.FilterLogs(opts, "UnstakeDelayTime")
	if err != nil {
		return nil, err
	}
	return &IWorkerHubUnstakeDelayTimeIterator{contract: _IWorkerHub.contract, event: "UnstakeDelayTime", logs: logs, sub: sub}, nil
}

// WatchUnstakeDelayTime is a free log subscription operation binding the contract event 0xdf63c46e5024e57c66aafc6698e317c78589c870dca694678c89dd379c5fd490.
//
// Solidity: event UnstakeDelayTime(uint256 oldDelayTime, uint256 newDelayTime)
func (_IWorkerHub *IWorkerHubFilterer) WatchUnstakeDelayTime(opts *bind.WatchOpts, sink chan<- *IWorkerHubUnstakeDelayTime) (event.Subscription, error) {

	logs, sub, err := _IWorkerHub.contract.WatchLogs(opts, "UnstakeDelayTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorkerHubUnstakeDelayTime)
				if err := _IWorkerHub.contract.UnpackLog(event, "UnstakeDelayTime", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnstakeDelayTime is a log parse operation binding the contract event 0xdf63c46e5024e57c66aafc6698e317c78589c870dca694678c89dd379c5fd490.
//
// Solidity: event UnstakeDelayTime(uint256 oldDelayTime, uint256 newDelayTime)
func (_IWorkerHub *IWorkerHubFilterer) ParseUnstakeDelayTime(log types.Log) (*IWorkerHubUnstakeDelayTime, error) {
	event := new(IWorkerHubUnstakeDelayTime)
	if err := _IWorkerHub.contract.UnpackLog(event, "UnstakeDelayTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
