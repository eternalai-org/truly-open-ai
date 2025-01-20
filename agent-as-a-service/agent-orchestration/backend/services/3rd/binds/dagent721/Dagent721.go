// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dagent721

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

// Dagent721MetaData contains all meta data concerning the Dagent721 contract.
var Dagent721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentPromptIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentURI\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureUsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"sysPrompt\",\"type\":\"bytes[]\"}],\"name\":\"AgentDataAddNew\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"promptIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"oldSysPrompt\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newSysPrompt\",\"type\":\"bytes\"}],\"name\":\"AgentDataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"AgentFeeUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"missions\",\"type\":\"bytes[]\"}],\"name\":\"AgentMissionAddNew\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"missionIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"oldSysMission\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newSysMission\",\"type\":\"bytes\"}],\"name\":\"AgentMissionUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldModelId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newModelId\",\"type\":\"uint256\"}],\"name\":\"AgentModelIdUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPromptScheduler\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOldPromptScheduler\",\"type\":\"address\"}],\"name\":\"AgentPromptSchedulerdUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"AgentURIUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"externalData\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"}],\"name\":\"InferencePerformed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MintPriceUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"sysPrompt\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"NewToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newValue\",\"type\":\"uint16\"}],\"name\":\"RoyaltyPortionUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"RoyaltyReceiverUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TopUpPoolBalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_gpuManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"name\":\"_poolBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"signature\",\"type\":\"bytes32\"}],\"name\":\"_signaturesUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"promptKey\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"sysPrompt\",\"type\":\"bytes\"}],\"name\":\"addNewAgentData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"missionData\",\"type\":\"bytes\"}],\"name\":\"createMission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"}],\"name\":\"dataOf\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"}],\"name\":\"getAgentFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getAgentIdByOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"promptKey\",\"type\":\"string\"}],\"name\":\"getAgentSystemPrompt\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"}],\"name\":\"getMissionIdsByAgentId\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fwdCalldata\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"externalData\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"promptKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"infer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fwdCalldata\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"externalData\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"promptKey\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"infer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"mintPrice_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"royaltyReceiver_\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"royaltyPortion_\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"nextTokenId_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gpuManager_\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenFee_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"promptKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"promptScheduler\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyPortion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"topUpPoolBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sysPrompt\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"promptKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"promptIdx\",\"type\":\"uint256\"}],\"name\":\"updateAgentData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sysPrompt\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"promptKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"promptIdx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"randomNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"updateAgentDataWithSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"updateAgentFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"newModelId\",\"type\":\"uint32\"}],\"name\":\"updateAgentModelId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"updateAgentURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"randomNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"updateAgentUriWithSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gpuManager\",\"type\":\"address\"}],\"name\":\"updateGPUManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintPrice\",\"type\":\"uint256\"}],\"name\":\"updateMintPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"royaltyPortion\",\"type\":\"uint16\"}],\"name\":\"updateRoyaltyPortion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"royaltyReceiver\",\"type\":\"address\"}],\"name\":\"updateRoyaltyReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newPromptScheduler\",\"type\":\"address\"}],\"name\":\"updateSchedulePrompt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Dagent721ABI is the input ABI used to generate the binding from.
// Deprecated: Use Dagent721MetaData.ABI instead.
var Dagent721ABI = Dagent721MetaData.ABI

// Dagent721 is an auto generated Go binding around an Ethereum contract.
type Dagent721 struct {
	Dagent721Caller     // Read-only binding to the contract
	Dagent721Transactor // Write-only binding to the contract
	Dagent721Filterer   // Log filterer for contract events
}

// Dagent721Caller is an auto generated read-only Go binding around an Ethereum contract.
type Dagent721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Dagent721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Dagent721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Dagent721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Dagent721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Dagent721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Dagent721Session struct {
	Contract     *Dagent721        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Dagent721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Dagent721CallerSession struct {
	Contract *Dagent721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Dagent721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Dagent721TransactorSession struct {
	Contract     *Dagent721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Dagent721Raw is an auto generated low-level Go binding around an Ethereum contract.
type Dagent721Raw struct {
	Contract *Dagent721 // Generic contract binding to access the raw methods on
}

// Dagent721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Dagent721CallerRaw struct {
	Contract *Dagent721Caller // Generic read-only contract binding to access the raw methods on
}

// Dagent721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Dagent721TransactorRaw struct {
	Contract *Dagent721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDagent721 creates a new instance of Dagent721, bound to a specific deployed contract.
func NewDagent721(address common.Address, backend bind.ContractBackend) (*Dagent721, error) {
	contract, err := bindDagent721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dagent721{Dagent721Caller: Dagent721Caller{contract: contract}, Dagent721Transactor: Dagent721Transactor{contract: contract}, Dagent721Filterer: Dagent721Filterer{contract: contract}}, nil
}

// NewDagent721Caller creates a new read-only instance of Dagent721, bound to a specific deployed contract.
func NewDagent721Caller(address common.Address, caller bind.ContractCaller) (*Dagent721Caller, error) {
	contract, err := bindDagent721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Dagent721Caller{contract: contract}, nil
}

// NewDagent721Transactor creates a new write-only instance of Dagent721, bound to a specific deployed contract.
func NewDagent721Transactor(address common.Address, transactor bind.ContractTransactor) (*Dagent721Transactor, error) {
	contract, err := bindDagent721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Dagent721Transactor{contract: contract}, nil
}

// NewDagent721Filterer creates a new log filterer instance of Dagent721, bound to a specific deployed contract.
func NewDagent721Filterer(address common.Address, filterer bind.ContractFilterer) (*Dagent721Filterer, error) {
	contract, err := bindDagent721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Dagent721Filterer{contract: contract}, nil
}

// bindDagent721 binds a generic wrapper to an already deployed contract.
func bindDagent721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Dagent721MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dagent721 *Dagent721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dagent721.Contract.Dagent721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dagent721 *Dagent721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagent721.Contract.Dagent721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dagent721 *Dagent721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dagent721.Contract.Dagent721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dagent721 *Dagent721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dagent721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dagent721 *Dagent721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagent721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dagent721 *Dagent721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dagent721.Contract.contract.Transact(opts, method, params...)
}

// GpuManager is a free data retrieval call binding the contract method 0x08c147fd.
//
// Solidity: function _gpuManager() view returns(address)
func (_Dagent721 *Dagent721Caller) GpuManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dagent721.contract.Call(opts, &out, "_gpuManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GpuManager is a free data retrieval call binding the contract method 0x08c147fd.
//
// Solidity: function _gpuManager() view returns(address)
func (_Dagent721 *Dagent721Session) GpuManager() (common.Address, error) {
	return _Dagent721.Contract.GpuManager(&_Dagent721.CallOpts)
}

// GpuManager is a free data retrieval call binding the contract method 0x08c147fd.
//
// Solidity: function _gpuManager() view returns(address)
func (_Dagent721 *Dagent721CallerSession) GpuManager() (common.Address, error) {
	return _Dagent721.Contract.GpuManager(&_Dagent721.CallOpts)
}

// PoolBalance is a free data retrieval call binding the contract method 0xf121000e.
//
// Solidity: function _poolBalance(uint256 nftId) view returns(uint256)
func (_Dagent721 *Dagent721Caller) PoolBalance(opts *bind.CallOpts, nftId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Dagent721.contract.Call(opts, &out, "_poolBalance", nftId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolBalance is a free data retrieval call binding the contract method 0xf121000e.
//
// Solidity: function _poolBalance(uint256 nftId) view returns(uint256)
func (_Dagent721 *Dagent721Session) PoolBalance(nftId *big.Int) (*big.Int, error) {
	return _Dagent721.Contract.PoolBalance(&_Dagent721.CallOpts, nftId)
}

// PoolBalance is a free data retrieval call binding the contract method 0xf121000e.
//
// Solidity: function _poolBalance(uint256 nftId) view returns(uint256)
func (_Dagent721 *Dagent721CallerSession) PoolBalance(nftId *big.Int) (*big.Int, error) {
	return _Dagent721.Contract.PoolBalance(&_Dagent721.CallOpts, nftId)
}

// SignaturesUsed is a free data retrieval call binding the contract method 0xc9070a92.
//
// Solidity: function _signaturesUsed(address nftId, bytes32 signature) view returns(bool)
func (_Dagent721 *Dagent721Caller) SignaturesUsed(opts *bind.CallOpts, nftId common.Address, signature [32]byte) (bool, error) {
	var out []interface{}
	err := _Dagent721.contract.Call(opts, &out, "_signaturesUsed", nftId, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SignaturesUsed is a free data retrieval call binding the contract method 0xc9070a92.
//
// Solidity: function _signaturesUsed(address nftId, bytes32 signature) view returns(bool)
func (_Dagent721 *Dagent721Session) SignaturesUsed(nftId common.Address, signature [32]byte) (bool, error) {
	return _Dagent721.Contract.SignaturesUsed(&_Dagent721.CallOpts, nftId, signature)
}

// SignaturesUsed is a free data retrieval call binding the contract method 0xc9070a92.
//
// Solidity: function _signaturesUsed(address nftId, bytes32 signature) view returns(bool)
func (_Dagent721 *Dagent721CallerSession) SignaturesUsed(nftId common.Address, signature [32]byte) (bool, error) {
	return _Dagent721.Contract.SignaturesUsed(&_Dagent721.CallOpts, nftId, signature)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Dagent721 *Dagent721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dagent721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Dagent721 *Dagent721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Dagent721.Contract.BalanceOf(&_Dagent721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Dagent721 *Dagent721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Dagent721.Contract.BalanceOf(&_Dagent721.CallOpts, owner)
}

// DataOf is a free data retrieval call binding the contract method 0x871caa98.
//
// Solidity: function dataOf(uint256 agentId) view returns(uint128, bool)
func (_Dagent721 *Dagent721Caller) DataOf(opts *bind.CallOpts, agentId *big.Int) (*big.Int, bool, error) {
	var out []interface{}
	err := _Dagent721.contract.Call(opts, &out, "dataOf", agentId)

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// DataOf is a free data retrieval call binding the contract method 0x871caa98.
//
// Solidity: function dataOf(uint256 agentId) view returns(uint128, bool)
func (_Dagent721 *Dagent721Session) DataOf(agentId *big.Int) (*big.Int, bool, error) {
	return _Dagent721.Contract.DataOf(&_Dagent721.CallOpts, agentId)
}

// DataOf is a free data retrieval call binding the contract method 0x871caa98.
//
// Solidity: function dataOf(uint256 agentId) view returns(uint128, bool)
func (_Dagent721 *Dagent721CallerSession) DataOf(agentId *big.Int) (*big.Int, bool, error) {
	return _Dagent721.Contract.DataOf(&_Dagent721.CallOpts, agentId)
}

// GetAgentFee is a free data retrieval call binding the contract method 0xed96f433.
//
// Solidity: function getAgentFee(uint256 agentId) view returns(uint256)
func (_Dagent721 *Dagent721Caller) GetAgentFee(opts *bind.CallOpts, agentId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Dagent721.contract.Call(opts, &out, "getAgentFee", agentId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAgentFee is a free data retrieval call binding the contract method 0xed96f433.
//
// Solidity: function getAgentFee(uint256 agentId) view returns(uint256)
func (_Dagent721 *Dagent721Session) GetAgentFee(agentId *big.Int) (*big.Int, error) {
	return _Dagent721.Contract.GetAgentFee(&_Dagent721.CallOpts, agentId)
}

// GetAgentFee is a free data retrieval call binding the contract method 0xed96f433.
//
// Solidity: function getAgentFee(uint256 agentId) view returns(uint256)
func (_Dagent721 *Dagent721CallerSession) GetAgentFee(agentId *big.Int) (*big.Int, error) {
	return _Dagent721.Contract.GetAgentFee(&_Dagent721.CallOpts, agentId)
}

// GetAgentIdByOwner is a free data retrieval call binding the contract method 0xae57a2d3.
//
// Solidity: function getAgentIdByOwner(address owner) view returns(uint256[])
func (_Dagent721 *Dagent721Caller) GetAgentIdByOwner(opts *bind.CallOpts, owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Dagent721.contract.Call(opts, &out, "getAgentIdByOwner", owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAgentIdByOwner is a free data retrieval call binding the contract method 0xae57a2d3.
//
// Solidity: function getAgentIdByOwner(address owner) view returns(uint256[])
func (_Dagent721 *Dagent721Session) GetAgentIdByOwner(owner common.Address) ([]*big.Int, error) {
	return _Dagent721.Contract.GetAgentIdByOwner(&_Dagent721.CallOpts, owner)
}

// GetAgentIdByOwner is a free data retrieval call binding the contract method 0xae57a2d3.
//
// Solidity: function getAgentIdByOwner(address owner) view returns(uint256[])
func (_Dagent721 *Dagent721CallerSession) GetAgentIdByOwner(owner common.Address) ([]*big.Int, error) {
	return _Dagent721.Contract.GetAgentIdByOwner(&_Dagent721.CallOpts, owner)
}

// GetAgentSystemPrompt is a free data retrieval call binding the contract method 0x5a88fc7c.
//
// Solidity: function getAgentSystemPrompt(uint256 agentId, string promptKey) view returns(bytes[])
func (_Dagent721 *Dagent721Caller) GetAgentSystemPrompt(opts *bind.CallOpts, agentId *big.Int, promptKey string) ([][]byte, error) {
	var out []interface{}
	err := _Dagent721.contract.Call(opts, &out, "getAgentSystemPrompt", agentId, promptKey)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetAgentSystemPrompt is a free data retrieval call binding the contract method 0x5a88fc7c.
//
// Solidity: function getAgentSystemPrompt(uint256 agentId, string promptKey) view returns(bytes[])
func (_Dagent721 *Dagent721Session) GetAgentSystemPrompt(agentId *big.Int, promptKey string) ([][]byte, error) {
	return _Dagent721.Contract.GetAgentSystemPrompt(&_Dagent721.CallOpts, agentId, promptKey)
}

// GetAgentSystemPrompt is a free data retrieval call binding the contract method 0x5a88fc7c.
//
// Solidity: function getAgentSystemPrompt(uint256 agentId, string promptKey) view returns(bytes[])
func (_Dagent721 *Dagent721CallerSession) GetAgentSystemPrompt(agentId *big.Int, promptKey string) ([][]byte, error) {
	return _Dagent721.Contract.GetAgentSystemPrompt(&_Dagent721.CallOpts, agentId, promptKey)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Dagent721 *Dagent721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Dagent721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Dagent721 *Dagent721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Dagent721.Contract.GetApproved(&_Dagent721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Dagent721 *Dagent721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Dagent721.Contract.GetApproved(&_Dagent721.CallOpts, tokenId)
}

// GetMissionIdsByAgentId is a free data retrieval call binding the contract method 0x96694ad0.
//
// Solidity: function getMissionIdsByAgentId(uint256 agentId) view returns(bytes[])
func (_Dagent721 *Dagent721Caller) GetMissionIdsByAgentId(opts *bind.CallOpts, agentId *big.Int) ([][]byte, error) {
	var out []interface{}
	err := _Dagent721.contract.Call(opts, &out, "getMissionIdsByAgentId", agentId)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetMissionIdsByAgentId is a free data retrieval call binding the contract method 0x96694ad0.
//
// Solidity: function getMissionIdsByAgentId(uint256 agentId) view returns(bytes[])
func (_Dagent721 *Dagent721Session) GetMissionIdsByAgentId(agentId *big.Int) ([][]byte, error) {
	return _Dagent721.Contract.GetMissionIdsByAgentId(&_Dagent721.CallOpts, agentId)
}

// GetMissionIdsByAgentId is a free data retrieval call binding the contract method 0x96694ad0.
//
// Solidity: function getMissionIdsByAgentId(uint256 agentId) view returns(bytes[])
func (_Dagent721 *Dagent721CallerSession) GetMissionIdsByAgentId(agentId *big.Int) ([][]byte, error) {
	return _Dagent721.Contract.GetMissionIdsByAgentId(&_Dagent721.CallOpts, agentId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Dagent721 *Dagent721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Dagent721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Dagent721 *Dagent721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Dagent721.Contract.IsApprovedForAll(&_Dagent721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Dagent721 *Dagent721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Dagent721.Contract.IsApprovedForAll(&_Dagent721.CallOpts, owner, operator)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_Dagent721 *Dagent721Caller) MintPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dagent721.contract.Call(opts, &out, "mintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_Dagent721 *Dagent721Session) MintPrice() (*big.Int, error) {
	return _Dagent721.Contract.MintPrice(&_Dagent721.CallOpts)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_Dagent721 *Dagent721CallerSession) MintPrice() (*big.Int, error) {
	return _Dagent721.Contract.MintPrice(&_Dagent721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dagent721 *Dagent721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dagent721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dagent721 *Dagent721Session) Name() (string, error) {
	return _Dagent721.Contract.Name(&_Dagent721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dagent721 *Dagent721CallerSession) Name() (string, error) {
	return _Dagent721.Contract.Name(&_Dagent721.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Dagent721 *Dagent721Caller) NextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dagent721.contract.Call(opts, &out, "nextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Dagent721 *Dagent721Session) NextTokenId() (*big.Int, error) {
	return _Dagent721.Contract.NextTokenId(&_Dagent721.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Dagent721 *Dagent721CallerSession) NextTokenId() (*big.Int, error) {
	return _Dagent721.Contract.NextTokenId(&_Dagent721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dagent721 *Dagent721Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dagent721.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dagent721 *Dagent721Session) Owner() (common.Address, error) {
	return _Dagent721.Contract.Owner(&_Dagent721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dagent721 *Dagent721CallerSession) Owner() (common.Address, error) {
	return _Dagent721.Contract.Owner(&_Dagent721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Dagent721 *Dagent721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Dagent721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Dagent721 *Dagent721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Dagent721.Contract.OwnerOf(&_Dagent721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Dagent721 *Dagent721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Dagent721.Contract.OwnerOf(&_Dagent721.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Dagent721 *Dagent721Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Dagent721.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Dagent721 *Dagent721Session) Paused() (bool, error) {
	return _Dagent721.Contract.Paused(&_Dagent721.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Dagent721 *Dagent721CallerSession) Paused() (bool, error) {
	return _Dagent721.Contract.Paused(&_Dagent721.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 agentId, uint256 salePrice) view returns(address, uint256)
func (_Dagent721 *Dagent721Caller) RoyaltyInfo(opts *bind.CallOpts, agentId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _Dagent721.contract.Call(opts, &out, "royaltyInfo", agentId, salePrice)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 agentId, uint256 salePrice) view returns(address, uint256)
func (_Dagent721 *Dagent721Session) RoyaltyInfo(agentId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	return _Dagent721.Contract.RoyaltyInfo(&_Dagent721.CallOpts, agentId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 agentId, uint256 salePrice) view returns(address, uint256)
func (_Dagent721 *Dagent721CallerSession) RoyaltyInfo(agentId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	return _Dagent721.Contract.RoyaltyInfo(&_Dagent721.CallOpts, agentId, salePrice)
}

// RoyaltyPortion is a free data retrieval call binding the contract method 0x11d7beb2.
//
// Solidity: function royaltyPortion() view returns(uint16)
func (_Dagent721 *Dagent721Caller) RoyaltyPortion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Dagent721.contract.Call(opts, &out, "royaltyPortion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RoyaltyPortion is a free data retrieval call binding the contract method 0x11d7beb2.
//
// Solidity: function royaltyPortion() view returns(uint16)
func (_Dagent721 *Dagent721Session) RoyaltyPortion() (uint16, error) {
	return _Dagent721.Contract.RoyaltyPortion(&_Dagent721.CallOpts)
}

// RoyaltyPortion is a free data retrieval call binding the contract method 0x11d7beb2.
//
// Solidity: function royaltyPortion() view returns(uint16)
func (_Dagent721 *Dagent721CallerSession) RoyaltyPortion() (uint16, error) {
	return _Dagent721.Contract.RoyaltyPortion(&_Dagent721.CallOpts)
}

// RoyaltyReceiver is a free data retrieval call binding the contract method 0x9fbc8713.
//
// Solidity: function royaltyReceiver() view returns(address)
func (_Dagent721 *Dagent721Caller) RoyaltyReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dagent721.contract.Call(opts, &out, "royaltyReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoyaltyReceiver is a free data retrieval call binding the contract method 0x9fbc8713.
//
// Solidity: function royaltyReceiver() view returns(address)
func (_Dagent721 *Dagent721Session) RoyaltyReceiver() (common.Address, error) {
	return _Dagent721.Contract.RoyaltyReceiver(&_Dagent721.CallOpts)
}

// RoyaltyReceiver is a free data retrieval call binding the contract method 0x9fbc8713.
//
// Solidity: function royaltyReceiver() view returns(address)
func (_Dagent721 *Dagent721CallerSession) RoyaltyReceiver() (common.Address, error) {
	return _Dagent721.Contract.RoyaltyReceiver(&_Dagent721.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dagent721 *Dagent721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Dagent721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dagent721 *Dagent721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Dagent721.Contract.SupportsInterface(&_Dagent721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Dagent721 *Dagent721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Dagent721.Contract.SupportsInterface(&_Dagent721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dagent721 *Dagent721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dagent721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dagent721 *Dagent721Session) Symbol() (string, error) {
	return _Dagent721.Contract.Symbol(&_Dagent721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dagent721 *Dagent721CallerSession) Symbol() (string, error) {
	return _Dagent721.Contract.Symbol(&_Dagent721.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Dagent721 *Dagent721Caller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Dagent721.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Dagent721 *Dagent721Session) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Dagent721.Contract.TokenByIndex(&_Dagent721.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Dagent721 *Dagent721CallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Dagent721.Contract.TokenByIndex(&_Dagent721.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Dagent721 *Dagent721Caller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Dagent721.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Dagent721 *Dagent721Session) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Dagent721.Contract.TokenOfOwnerByIndex(&_Dagent721.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Dagent721 *Dagent721CallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Dagent721.Contract.TokenOfOwnerByIndex(&_Dagent721.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _agentId) view returns(string)
func (_Dagent721 *Dagent721Caller) TokenURI(opts *bind.CallOpts, _agentId *big.Int) (string, error) {
	var out []interface{}
	err := _Dagent721.contract.Call(opts, &out, "tokenURI", _agentId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _agentId) view returns(string)
func (_Dagent721 *Dagent721Session) TokenURI(_agentId *big.Int) (string, error) {
	return _Dagent721.Contract.TokenURI(&_Dagent721.CallOpts, _agentId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _agentId) view returns(string)
func (_Dagent721 *Dagent721CallerSession) TokenURI(_agentId *big.Int) (string, error) {
	return _Dagent721.Contract.TokenURI(&_Dagent721.CallOpts, _agentId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dagent721 *Dagent721Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dagent721.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dagent721 *Dagent721Session) TotalSupply() (*big.Int, error) {
	return _Dagent721.Contract.TotalSupply(&_Dagent721.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dagent721 *Dagent721CallerSession) TotalSupply() (*big.Int, error) {
	return _Dagent721.Contract.TotalSupply(&_Dagent721.CallOpts)
}

// AddNewAgentData is a paid mutator transaction binding the contract method 0x413b0efa.
//
// Solidity: function addNewAgentData(uint256 agentId, string promptKey, bytes sysPrompt) returns()
func (_Dagent721 *Dagent721Transactor) AddNewAgentData(opts *bind.TransactOpts, agentId *big.Int, promptKey string, sysPrompt []byte) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "addNewAgentData", agentId, promptKey, sysPrompt)
}

// AddNewAgentData is a paid mutator transaction binding the contract method 0x413b0efa.
//
// Solidity: function addNewAgentData(uint256 agentId, string promptKey, bytes sysPrompt) returns()
func (_Dagent721 *Dagent721Session) AddNewAgentData(agentId *big.Int, promptKey string, sysPrompt []byte) (*types.Transaction, error) {
	return _Dagent721.Contract.AddNewAgentData(&_Dagent721.TransactOpts, agentId, promptKey, sysPrompt)
}

// AddNewAgentData is a paid mutator transaction binding the contract method 0x413b0efa.
//
// Solidity: function addNewAgentData(uint256 agentId, string promptKey, bytes sysPrompt) returns()
func (_Dagent721 *Dagent721TransactorSession) AddNewAgentData(agentId *big.Int, promptKey string, sysPrompt []byte) (*types.Transaction, error) {
	return _Dagent721.Contract.AddNewAgentData(&_Dagent721.TransactOpts, agentId, promptKey, sysPrompt)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Dagent721 *Dagent721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Dagent721 *Dagent721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dagent721.Contract.Approve(&_Dagent721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Dagent721 *Dagent721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dagent721.Contract.Approve(&_Dagent721.TransactOpts, to, tokenId)
}

// CreateMission is a paid mutator transaction binding the contract method 0x6001ad44.
//
// Solidity: function createMission(uint256 agentId, bytes missionData) returns()
func (_Dagent721 *Dagent721Transactor) CreateMission(opts *bind.TransactOpts, agentId *big.Int, missionData []byte) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "createMission", agentId, missionData)
}

// CreateMission is a paid mutator transaction binding the contract method 0x6001ad44.
//
// Solidity: function createMission(uint256 agentId, bytes missionData) returns()
func (_Dagent721 *Dagent721Session) CreateMission(agentId *big.Int, missionData []byte) (*types.Transaction, error) {
	return _Dagent721.Contract.CreateMission(&_Dagent721.TransactOpts, agentId, missionData)
}

// CreateMission is a paid mutator transaction binding the contract method 0x6001ad44.
//
// Solidity: function createMission(uint256 agentId, bytes missionData) returns()
func (_Dagent721 *Dagent721TransactorSession) CreateMission(agentId *big.Int, missionData []byte) (*types.Transaction, error) {
	return _Dagent721.Contract.CreateMission(&_Dagent721.TransactOpts, agentId, missionData)
}

// Infer is a paid mutator transaction binding the contract method 0x3b8117fa.
//
// Solidity: function infer(uint256 agentId, bytes fwdCalldata, string externalData, string promptKey, uint256 feeAmount) returns()
func (_Dagent721 *Dagent721Transactor) Infer(opts *bind.TransactOpts, agentId *big.Int, fwdCalldata []byte, externalData string, promptKey string, feeAmount *big.Int) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "infer", agentId, fwdCalldata, externalData, promptKey, feeAmount)
}

// Infer is a paid mutator transaction binding the contract method 0x3b8117fa.
//
// Solidity: function infer(uint256 agentId, bytes fwdCalldata, string externalData, string promptKey, uint256 feeAmount) returns()
func (_Dagent721 *Dagent721Session) Infer(agentId *big.Int, fwdCalldata []byte, externalData string, promptKey string, feeAmount *big.Int) (*types.Transaction, error) {
	return _Dagent721.Contract.Infer(&_Dagent721.TransactOpts, agentId, fwdCalldata, externalData, promptKey, feeAmount)
}

// Infer is a paid mutator transaction binding the contract method 0x3b8117fa.
//
// Solidity: function infer(uint256 agentId, bytes fwdCalldata, string externalData, string promptKey, uint256 feeAmount) returns()
func (_Dagent721 *Dagent721TransactorSession) Infer(agentId *big.Int, fwdCalldata []byte, externalData string, promptKey string, feeAmount *big.Int) (*types.Transaction, error) {
	return _Dagent721.Contract.Infer(&_Dagent721.TransactOpts, agentId, fwdCalldata, externalData, promptKey, feeAmount)
}

// Infer0 is a paid mutator transaction binding the contract method 0xfa064b19.
//
// Solidity: function infer(uint256 agentId, bytes fwdCalldata, string externalData, string promptKey, bool flag, uint256 feeAmount) returns()
func (_Dagent721 *Dagent721Transactor) Infer0(opts *bind.TransactOpts, agentId *big.Int, fwdCalldata []byte, externalData string, promptKey string, flag bool, feeAmount *big.Int) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "infer0", agentId, fwdCalldata, externalData, promptKey, flag, feeAmount)
}

// Infer0 is a paid mutator transaction binding the contract method 0xfa064b19.
//
// Solidity: function infer(uint256 agentId, bytes fwdCalldata, string externalData, string promptKey, bool flag, uint256 feeAmount) returns()
func (_Dagent721 *Dagent721Session) Infer0(agentId *big.Int, fwdCalldata []byte, externalData string, promptKey string, flag bool, feeAmount *big.Int) (*types.Transaction, error) {
	return _Dagent721.Contract.Infer0(&_Dagent721.TransactOpts, agentId, fwdCalldata, externalData, promptKey, flag, feeAmount)
}

// Infer0 is a paid mutator transaction binding the contract method 0xfa064b19.
//
// Solidity: function infer(uint256 agentId, bytes fwdCalldata, string externalData, string promptKey, bool flag, uint256 feeAmount) returns()
func (_Dagent721 *Dagent721TransactorSession) Infer0(agentId *big.Int, fwdCalldata []byte, externalData string, promptKey string, flag bool, feeAmount *big.Int) (*types.Transaction, error) {
	return _Dagent721.Contract.Infer0(&_Dagent721.TransactOpts, agentId, fwdCalldata, externalData, promptKey, flag, feeAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x1e18fb8c.
//
// Solidity: function initialize(string name_, string symbol_, uint256 mintPrice_, address royaltyReceiver_, uint16 royaltyPortion_, uint256 nextTokenId_, address gpuManager_, address tokenFee_) returns()
func (_Dagent721 *Dagent721Transactor) Initialize(opts *bind.TransactOpts, name_ string, symbol_ string, mintPrice_ *big.Int, royaltyReceiver_ common.Address, royaltyPortion_ uint16, nextTokenId_ *big.Int, gpuManager_ common.Address, tokenFee_ common.Address) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "initialize", name_, symbol_, mintPrice_, royaltyReceiver_, royaltyPortion_, nextTokenId_, gpuManager_, tokenFee_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1e18fb8c.
//
// Solidity: function initialize(string name_, string symbol_, uint256 mintPrice_, address royaltyReceiver_, uint16 royaltyPortion_, uint256 nextTokenId_, address gpuManager_, address tokenFee_) returns()
func (_Dagent721 *Dagent721Session) Initialize(name_ string, symbol_ string, mintPrice_ *big.Int, royaltyReceiver_ common.Address, royaltyPortion_ uint16, nextTokenId_ *big.Int, gpuManager_ common.Address, tokenFee_ common.Address) (*types.Transaction, error) {
	return _Dagent721.Contract.Initialize(&_Dagent721.TransactOpts, name_, symbol_, mintPrice_, royaltyReceiver_, royaltyPortion_, nextTokenId_, gpuManager_, tokenFee_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1e18fb8c.
//
// Solidity: function initialize(string name_, string symbol_, uint256 mintPrice_, address royaltyReceiver_, uint16 royaltyPortion_, uint256 nextTokenId_, address gpuManager_, address tokenFee_) returns()
func (_Dagent721 *Dagent721TransactorSession) Initialize(name_ string, symbol_ string, mintPrice_ *big.Int, royaltyReceiver_ common.Address, royaltyPortion_ uint16, nextTokenId_ *big.Int, gpuManager_ common.Address, tokenFee_ common.Address) (*types.Transaction, error) {
	return _Dagent721.Contract.Initialize(&_Dagent721.TransactOpts, name_, symbol_, mintPrice_, royaltyReceiver_, royaltyPortion_, nextTokenId_, gpuManager_, tokenFee_)
}

// Mint is a paid mutator transaction binding the contract method 0xb8ea1693.
//
// Solidity: function mint(address to, string uri, bytes data, uint256 fee, string promptKey, address promptScheduler, uint32 modelId) returns(uint256)
func (_Dagent721 *Dagent721Transactor) Mint(opts *bind.TransactOpts, to common.Address, uri string, data []byte, fee *big.Int, promptKey string, promptScheduler common.Address, modelId uint32) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "mint", to, uri, data, fee, promptKey, promptScheduler, modelId)
}

// Mint is a paid mutator transaction binding the contract method 0xb8ea1693.
//
// Solidity: function mint(address to, string uri, bytes data, uint256 fee, string promptKey, address promptScheduler, uint32 modelId) returns(uint256)
func (_Dagent721 *Dagent721Session) Mint(to common.Address, uri string, data []byte, fee *big.Int, promptKey string, promptScheduler common.Address, modelId uint32) (*types.Transaction, error) {
	return _Dagent721.Contract.Mint(&_Dagent721.TransactOpts, to, uri, data, fee, promptKey, promptScheduler, modelId)
}

// Mint is a paid mutator transaction binding the contract method 0xb8ea1693.
//
// Solidity: function mint(address to, string uri, bytes data, uint256 fee, string promptKey, address promptScheduler, uint32 modelId) returns(uint256)
func (_Dagent721 *Dagent721TransactorSession) Mint(to common.Address, uri string, data []byte, fee *big.Int, promptKey string, promptScheduler common.Address, modelId uint32) (*types.Transaction, error) {
	return _Dagent721.Contract.Mint(&_Dagent721.TransactOpts, to, uri, data, fee, promptKey, promptScheduler, modelId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Dagent721 *Dagent721Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Dagent721 *Dagent721Session) Pause() (*types.Transaction, error) {
	return _Dagent721.Contract.Pause(&_Dagent721.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Dagent721 *Dagent721TransactorSession) Pause() (*types.Transaction, error) {
	return _Dagent721.Contract.Pause(&_Dagent721.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dagent721 *Dagent721Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dagent721 *Dagent721Session) RenounceOwnership() (*types.Transaction, error) {
	return _Dagent721.Contract.RenounceOwnership(&_Dagent721.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dagent721 *Dagent721TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dagent721.Contract.RenounceOwnership(&_Dagent721.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Dagent721 *Dagent721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Dagent721 *Dagent721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dagent721.Contract.SafeTransferFrom(&_Dagent721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Dagent721 *Dagent721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dagent721.Contract.SafeTransferFrom(&_Dagent721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Dagent721 *Dagent721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Dagent721 *Dagent721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Dagent721.Contract.SafeTransferFrom0(&_Dagent721.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Dagent721 *Dagent721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Dagent721.Contract.SafeTransferFrom0(&_Dagent721.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Dagent721 *Dagent721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Dagent721 *Dagent721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Dagent721.Contract.SetApprovalForAll(&_Dagent721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Dagent721 *Dagent721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Dagent721.Contract.SetApprovalForAll(&_Dagent721.TransactOpts, operator, approved)
}

// TopUpPoolBalance is a paid mutator transaction binding the contract method 0x67b611a0.
//
// Solidity: function topUpPoolBalance(uint256 agentId, uint256 amount) returns()
func (_Dagent721 *Dagent721Transactor) TopUpPoolBalance(opts *bind.TransactOpts, agentId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "topUpPoolBalance", agentId, amount)
}

// TopUpPoolBalance is a paid mutator transaction binding the contract method 0x67b611a0.
//
// Solidity: function topUpPoolBalance(uint256 agentId, uint256 amount) returns()
func (_Dagent721 *Dagent721Session) TopUpPoolBalance(agentId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Dagent721.Contract.TopUpPoolBalance(&_Dagent721.TransactOpts, agentId, amount)
}

// TopUpPoolBalance is a paid mutator transaction binding the contract method 0x67b611a0.
//
// Solidity: function topUpPoolBalance(uint256 agentId, uint256 amount) returns()
func (_Dagent721 *Dagent721TransactorSession) TopUpPoolBalance(agentId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Dagent721.Contract.TopUpPoolBalance(&_Dagent721.TransactOpts, agentId, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Dagent721 *Dagent721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Dagent721 *Dagent721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dagent721.Contract.TransferFrom(&_Dagent721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Dagent721 *Dagent721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Dagent721.Contract.TransferFrom(&_Dagent721.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dagent721 *Dagent721Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dagent721 *Dagent721Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dagent721.Contract.TransferOwnership(&_Dagent721.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dagent721 *Dagent721TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dagent721.Contract.TransferOwnership(&_Dagent721.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Dagent721 *Dagent721Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Dagent721 *Dagent721Session) Unpause() (*types.Transaction, error) {
	return _Dagent721.Contract.Unpause(&_Dagent721.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Dagent721 *Dagent721TransactorSession) Unpause() (*types.Transaction, error) {
	return _Dagent721.Contract.Unpause(&_Dagent721.TransactOpts)
}

// UpdateAgentData is a paid mutator transaction binding the contract method 0xe96177c4.
//
// Solidity: function updateAgentData(uint256 agentId, bytes sysPrompt, string promptKey, uint256 promptIdx) returns()
func (_Dagent721 *Dagent721Transactor) UpdateAgentData(opts *bind.TransactOpts, agentId *big.Int, sysPrompt []byte, promptKey string, promptIdx *big.Int) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "updateAgentData", agentId, sysPrompt, promptKey, promptIdx)
}

// UpdateAgentData is a paid mutator transaction binding the contract method 0xe96177c4.
//
// Solidity: function updateAgentData(uint256 agentId, bytes sysPrompt, string promptKey, uint256 promptIdx) returns()
func (_Dagent721 *Dagent721Session) UpdateAgentData(agentId *big.Int, sysPrompt []byte, promptKey string, promptIdx *big.Int) (*types.Transaction, error) {
	return _Dagent721.Contract.UpdateAgentData(&_Dagent721.TransactOpts, agentId, sysPrompt, promptKey, promptIdx)
}

// UpdateAgentData is a paid mutator transaction binding the contract method 0xe96177c4.
//
// Solidity: function updateAgentData(uint256 agentId, bytes sysPrompt, string promptKey, uint256 promptIdx) returns()
func (_Dagent721 *Dagent721TransactorSession) UpdateAgentData(agentId *big.Int, sysPrompt []byte, promptKey string, promptIdx *big.Int) (*types.Transaction, error) {
	return _Dagent721.Contract.UpdateAgentData(&_Dagent721.TransactOpts, agentId, sysPrompt, promptKey, promptIdx)
}

// UpdateAgentDataWithSignature is a paid mutator transaction binding the contract method 0x1c83fb2c.
//
// Solidity: function updateAgentDataWithSignature(uint256 agentId, bytes sysPrompt, string promptKey, uint256 promptIdx, uint256 randomNonce, bytes signature) returns()
func (_Dagent721 *Dagent721Transactor) UpdateAgentDataWithSignature(opts *bind.TransactOpts, agentId *big.Int, sysPrompt []byte, promptKey string, promptIdx *big.Int, randomNonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "updateAgentDataWithSignature", agentId, sysPrompt, promptKey, promptIdx, randomNonce, signature)
}

// UpdateAgentDataWithSignature is a paid mutator transaction binding the contract method 0x1c83fb2c.
//
// Solidity: function updateAgentDataWithSignature(uint256 agentId, bytes sysPrompt, string promptKey, uint256 promptIdx, uint256 randomNonce, bytes signature) returns()
func (_Dagent721 *Dagent721Session) UpdateAgentDataWithSignature(agentId *big.Int, sysPrompt []byte, promptKey string, promptIdx *big.Int, randomNonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _Dagent721.Contract.UpdateAgentDataWithSignature(&_Dagent721.TransactOpts, agentId, sysPrompt, promptKey, promptIdx, randomNonce, signature)
}

// UpdateAgentDataWithSignature is a paid mutator transaction binding the contract method 0x1c83fb2c.
//
// Solidity: function updateAgentDataWithSignature(uint256 agentId, bytes sysPrompt, string promptKey, uint256 promptIdx, uint256 randomNonce, bytes signature) returns()
func (_Dagent721 *Dagent721TransactorSession) UpdateAgentDataWithSignature(agentId *big.Int, sysPrompt []byte, promptKey string, promptIdx *big.Int, randomNonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _Dagent721.Contract.UpdateAgentDataWithSignature(&_Dagent721.TransactOpts, agentId, sysPrompt, promptKey, promptIdx, randomNonce, signature)
}

// UpdateAgentFee is a paid mutator transaction binding the contract method 0xb1fd1526.
//
// Solidity: function updateAgentFee(uint256 agentId, uint256 fee) returns()
func (_Dagent721 *Dagent721Transactor) UpdateAgentFee(opts *bind.TransactOpts, agentId *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "updateAgentFee", agentId, fee)
}

// UpdateAgentFee is a paid mutator transaction binding the contract method 0xb1fd1526.
//
// Solidity: function updateAgentFee(uint256 agentId, uint256 fee) returns()
func (_Dagent721 *Dagent721Session) UpdateAgentFee(agentId *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Dagent721.Contract.UpdateAgentFee(&_Dagent721.TransactOpts, agentId, fee)
}

// UpdateAgentFee is a paid mutator transaction binding the contract method 0xb1fd1526.
//
// Solidity: function updateAgentFee(uint256 agentId, uint256 fee) returns()
func (_Dagent721 *Dagent721TransactorSession) UpdateAgentFee(agentId *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Dagent721.Contract.UpdateAgentFee(&_Dagent721.TransactOpts, agentId, fee)
}

// UpdateAgentModelId is a paid mutator transaction binding the contract method 0x0ffc8cf4.
//
// Solidity: function updateAgentModelId(uint256 agentId, uint32 newModelId) returns()
func (_Dagent721 *Dagent721Transactor) UpdateAgentModelId(opts *bind.TransactOpts, agentId *big.Int, newModelId uint32) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "updateAgentModelId", agentId, newModelId)
}

// UpdateAgentModelId is a paid mutator transaction binding the contract method 0x0ffc8cf4.
//
// Solidity: function updateAgentModelId(uint256 agentId, uint32 newModelId) returns()
func (_Dagent721 *Dagent721Session) UpdateAgentModelId(agentId *big.Int, newModelId uint32) (*types.Transaction, error) {
	return _Dagent721.Contract.UpdateAgentModelId(&_Dagent721.TransactOpts, agentId, newModelId)
}

// UpdateAgentModelId is a paid mutator transaction binding the contract method 0x0ffc8cf4.
//
// Solidity: function updateAgentModelId(uint256 agentId, uint32 newModelId) returns()
func (_Dagent721 *Dagent721TransactorSession) UpdateAgentModelId(agentId *big.Int, newModelId uint32) (*types.Transaction, error) {
	return _Dagent721.Contract.UpdateAgentModelId(&_Dagent721.TransactOpts, agentId, newModelId)
}

// UpdateAgentURI is a paid mutator transaction binding the contract method 0x6b595822.
//
// Solidity: function updateAgentURI(uint256 agentId, string uri) returns()
func (_Dagent721 *Dagent721Transactor) UpdateAgentURI(opts *bind.TransactOpts, agentId *big.Int, uri string) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "updateAgentURI", agentId, uri)
}

// UpdateAgentURI is a paid mutator transaction binding the contract method 0x6b595822.
//
// Solidity: function updateAgentURI(uint256 agentId, string uri) returns()
func (_Dagent721 *Dagent721Session) UpdateAgentURI(agentId *big.Int, uri string) (*types.Transaction, error) {
	return _Dagent721.Contract.UpdateAgentURI(&_Dagent721.TransactOpts, agentId, uri)
}

// UpdateAgentURI is a paid mutator transaction binding the contract method 0x6b595822.
//
// Solidity: function updateAgentURI(uint256 agentId, string uri) returns()
func (_Dagent721 *Dagent721TransactorSession) UpdateAgentURI(agentId *big.Int, uri string) (*types.Transaction, error) {
	return _Dagent721.Contract.UpdateAgentURI(&_Dagent721.TransactOpts, agentId, uri)
}

// UpdateAgentUriWithSignature is a paid mutator transaction binding the contract method 0xf5888779.
//
// Solidity: function updateAgentUriWithSignature(uint256 agentId, string uri, uint256 randomNonce, bytes signature) returns()
func (_Dagent721 *Dagent721Transactor) UpdateAgentUriWithSignature(opts *bind.TransactOpts, agentId *big.Int, uri string, randomNonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "updateAgentUriWithSignature", agentId, uri, randomNonce, signature)
}

// UpdateAgentUriWithSignature is a paid mutator transaction binding the contract method 0xf5888779.
//
// Solidity: function updateAgentUriWithSignature(uint256 agentId, string uri, uint256 randomNonce, bytes signature) returns()
func (_Dagent721 *Dagent721Session) UpdateAgentUriWithSignature(agentId *big.Int, uri string, randomNonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _Dagent721.Contract.UpdateAgentUriWithSignature(&_Dagent721.TransactOpts, agentId, uri, randomNonce, signature)
}

// UpdateAgentUriWithSignature is a paid mutator transaction binding the contract method 0xf5888779.
//
// Solidity: function updateAgentUriWithSignature(uint256 agentId, string uri, uint256 randomNonce, bytes signature) returns()
func (_Dagent721 *Dagent721TransactorSession) UpdateAgentUriWithSignature(agentId *big.Int, uri string, randomNonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _Dagent721.Contract.UpdateAgentUriWithSignature(&_Dagent721.TransactOpts, agentId, uri, randomNonce, signature)
}

// UpdateGPUManager is a paid mutator transaction binding the contract method 0x88ee5fb2.
//
// Solidity: function updateGPUManager(address gpuManager) returns()
func (_Dagent721 *Dagent721Transactor) UpdateGPUManager(opts *bind.TransactOpts, gpuManager common.Address) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "updateGPUManager", gpuManager)
}

// UpdateGPUManager is a paid mutator transaction binding the contract method 0x88ee5fb2.
//
// Solidity: function updateGPUManager(address gpuManager) returns()
func (_Dagent721 *Dagent721Session) UpdateGPUManager(gpuManager common.Address) (*types.Transaction, error) {
	return _Dagent721.Contract.UpdateGPUManager(&_Dagent721.TransactOpts, gpuManager)
}

// UpdateGPUManager is a paid mutator transaction binding the contract method 0x88ee5fb2.
//
// Solidity: function updateGPUManager(address gpuManager) returns()
func (_Dagent721 *Dagent721TransactorSession) UpdateGPUManager(gpuManager common.Address) (*types.Transaction, error) {
	return _Dagent721.Contract.UpdateGPUManager(&_Dagent721.TransactOpts, gpuManager)
}

// UpdateMintPrice is a paid mutator transaction binding the contract method 0x00728e46.
//
// Solidity: function updateMintPrice(uint256 mintPrice) returns()
func (_Dagent721 *Dagent721Transactor) UpdateMintPrice(opts *bind.TransactOpts, mintPrice *big.Int) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "updateMintPrice", mintPrice)
}

// UpdateMintPrice is a paid mutator transaction binding the contract method 0x00728e46.
//
// Solidity: function updateMintPrice(uint256 mintPrice) returns()
func (_Dagent721 *Dagent721Session) UpdateMintPrice(mintPrice *big.Int) (*types.Transaction, error) {
	return _Dagent721.Contract.UpdateMintPrice(&_Dagent721.TransactOpts, mintPrice)
}

// UpdateMintPrice is a paid mutator transaction binding the contract method 0x00728e46.
//
// Solidity: function updateMintPrice(uint256 mintPrice) returns()
func (_Dagent721 *Dagent721TransactorSession) UpdateMintPrice(mintPrice *big.Int) (*types.Transaction, error) {
	return _Dagent721.Contract.UpdateMintPrice(&_Dagent721.TransactOpts, mintPrice)
}

// UpdateRoyaltyPortion is a paid mutator transaction binding the contract method 0x19e93993.
//
// Solidity: function updateRoyaltyPortion(uint16 royaltyPortion) returns()
func (_Dagent721 *Dagent721Transactor) UpdateRoyaltyPortion(opts *bind.TransactOpts, royaltyPortion uint16) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "updateRoyaltyPortion", royaltyPortion)
}

// UpdateRoyaltyPortion is a paid mutator transaction binding the contract method 0x19e93993.
//
// Solidity: function updateRoyaltyPortion(uint16 royaltyPortion) returns()
func (_Dagent721 *Dagent721Session) UpdateRoyaltyPortion(royaltyPortion uint16) (*types.Transaction, error) {
	return _Dagent721.Contract.UpdateRoyaltyPortion(&_Dagent721.TransactOpts, royaltyPortion)
}

// UpdateRoyaltyPortion is a paid mutator transaction binding the contract method 0x19e93993.
//
// Solidity: function updateRoyaltyPortion(uint16 royaltyPortion) returns()
func (_Dagent721 *Dagent721TransactorSession) UpdateRoyaltyPortion(royaltyPortion uint16) (*types.Transaction, error) {
	return _Dagent721.Contract.UpdateRoyaltyPortion(&_Dagent721.TransactOpts, royaltyPortion)
}

// UpdateRoyaltyReceiver is a paid mutator transaction binding the contract method 0x29dc4d9b.
//
// Solidity: function updateRoyaltyReceiver(address royaltyReceiver) returns()
func (_Dagent721 *Dagent721Transactor) UpdateRoyaltyReceiver(opts *bind.TransactOpts, royaltyReceiver common.Address) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "updateRoyaltyReceiver", royaltyReceiver)
}

// UpdateRoyaltyReceiver is a paid mutator transaction binding the contract method 0x29dc4d9b.
//
// Solidity: function updateRoyaltyReceiver(address royaltyReceiver) returns()
func (_Dagent721 *Dagent721Session) UpdateRoyaltyReceiver(royaltyReceiver common.Address) (*types.Transaction, error) {
	return _Dagent721.Contract.UpdateRoyaltyReceiver(&_Dagent721.TransactOpts, royaltyReceiver)
}

// UpdateRoyaltyReceiver is a paid mutator transaction binding the contract method 0x29dc4d9b.
//
// Solidity: function updateRoyaltyReceiver(address royaltyReceiver) returns()
func (_Dagent721 *Dagent721TransactorSession) UpdateRoyaltyReceiver(royaltyReceiver common.Address) (*types.Transaction, error) {
	return _Dagent721.Contract.UpdateRoyaltyReceiver(&_Dagent721.TransactOpts, royaltyReceiver)
}

// UpdateSchedulePrompt is a paid mutator transaction binding the contract method 0x1ddbc69a.
//
// Solidity: function updateSchedulePrompt(uint256 agentId, address newPromptScheduler) returns()
func (_Dagent721 *Dagent721Transactor) UpdateSchedulePrompt(opts *bind.TransactOpts, agentId *big.Int, newPromptScheduler common.Address) (*types.Transaction, error) {
	return _Dagent721.contract.Transact(opts, "updateSchedulePrompt", agentId, newPromptScheduler)
}

// UpdateSchedulePrompt is a paid mutator transaction binding the contract method 0x1ddbc69a.
//
// Solidity: function updateSchedulePrompt(uint256 agentId, address newPromptScheduler) returns()
func (_Dagent721 *Dagent721Session) UpdateSchedulePrompt(agentId *big.Int, newPromptScheduler common.Address) (*types.Transaction, error) {
	return _Dagent721.Contract.UpdateSchedulePrompt(&_Dagent721.TransactOpts, agentId, newPromptScheduler)
}

// UpdateSchedulePrompt is a paid mutator transaction binding the contract method 0x1ddbc69a.
//
// Solidity: function updateSchedulePrompt(uint256 agentId, address newPromptScheduler) returns()
func (_Dagent721 *Dagent721TransactorSession) UpdateSchedulePrompt(agentId *big.Int, newPromptScheduler common.Address) (*types.Transaction, error) {
	return _Dagent721.Contract.UpdateSchedulePrompt(&_Dagent721.TransactOpts, agentId, newPromptScheduler)
}

// Dagent721AgentDataAddNewIterator is returned from FilterAgentDataAddNew and is used to iterate over the raw logs and unpacked data for AgentDataAddNew events raised by the Dagent721 contract.
type Dagent721AgentDataAddNewIterator struct {
	Event *Dagent721AgentDataAddNew // Event containing the contract specifics and raw log

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
func (it *Dagent721AgentDataAddNewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dagent721AgentDataAddNew)
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
		it.Event = new(Dagent721AgentDataAddNew)
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
func (it *Dagent721AgentDataAddNewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dagent721AgentDataAddNewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dagent721AgentDataAddNew represents a AgentDataAddNew event raised by the Dagent721 contract.
type Dagent721AgentDataAddNew struct {
	AgentId   *big.Int
	SysPrompt [][]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAgentDataAddNew is a free log retrieval operation binding the contract event 0xdebec4c58e3b7c5817893e50cb1f9e65b65978e8c89bb4407eb0109d5887b258.
//
// Solidity: event AgentDataAddNew(uint256 indexed agentId, bytes[] sysPrompt)
func (_Dagent721 *Dagent721Filterer) FilterAgentDataAddNew(opts *bind.FilterOpts, agentId []*big.Int) (*Dagent721AgentDataAddNewIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _Dagent721.contract.FilterLogs(opts, "AgentDataAddNew", agentIdRule)
	if err != nil {
		return nil, err
	}
	return &Dagent721AgentDataAddNewIterator{contract: _Dagent721.contract, event: "AgentDataAddNew", logs: logs, sub: sub}, nil
}

// WatchAgentDataAddNew is a free log subscription operation binding the contract event 0xdebec4c58e3b7c5817893e50cb1f9e65b65978e8c89bb4407eb0109d5887b258.
//
// Solidity: event AgentDataAddNew(uint256 indexed agentId, bytes[] sysPrompt)
func (_Dagent721 *Dagent721Filterer) WatchAgentDataAddNew(opts *bind.WatchOpts, sink chan<- *Dagent721AgentDataAddNew, agentId []*big.Int) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _Dagent721.contract.WatchLogs(opts, "AgentDataAddNew", agentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dagent721AgentDataAddNew)
				if err := _Dagent721.contract.UnpackLog(event, "AgentDataAddNew", log); err != nil {
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

// ParseAgentDataAddNew is a log parse operation binding the contract event 0xdebec4c58e3b7c5817893e50cb1f9e65b65978e8c89bb4407eb0109d5887b258.
//
// Solidity: event AgentDataAddNew(uint256 indexed agentId, bytes[] sysPrompt)
func (_Dagent721 *Dagent721Filterer) ParseAgentDataAddNew(log types.Log) (*Dagent721AgentDataAddNew, error) {
	event := new(Dagent721AgentDataAddNew)
	if err := _Dagent721.contract.UnpackLog(event, "AgentDataAddNew", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Dagent721AgentDataUpdateIterator is returned from FilterAgentDataUpdate and is used to iterate over the raw logs and unpacked data for AgentDataUpdate events raised by the Dagent721 contract.
type Dagent721AgentDataUpdateIterator struct {
	Event *Dagent721AgentDataUpdate // Event containing the contract specifics and raw log

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
func (it *Dagent721AgentDataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dagent721AgentDataUpdate)
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
		it.Event = new(Dagent721AgentDataUpdate)
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
func (it *Dagent721AgentDataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dagent721AgentDataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dagent721AgentDataUpdate represents a AgentDataUpdate event raised by the Dagent721 contract.
type Dagent721AgentDataUpdate struct {
	AgentId      *big.Int
	PromptIndex  *big.Int
	OldSysPrompt []byte
	NewSysPrompt []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAgentDataUpdate is a free log retrieval operation binding the contract event 0xe42abf7d4a793286da8cc1399cb577a1f5a0e133dfee371bb3a5abbdd77b011e.
//
// Solidity: event AgentDataUpdate(uint256 indexed agentId, uint256 promptIndex, bytes oldSysPrompt, bytes newSysPrompt)
func (_Dagent721 *Dagent721Filterer) FilterAgentDataUpdate(opts *bind.FilterOpts, agentId []*big.Int) (*Dagent721AgentDataUpdateIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _Dagent721.contract.FilterLogs(opts, "AgentDataUpdate", agentIdRule)
	if err != nil {
		return nil, err
	}
	return &Dagent721AgentDataUpdateIterator{contract: _Dagent721.contract, event: "AgentDataUpdate", logs: logs, sub: sub}, nil
}

// WatchAgentDataUpdate is a free log subscription operation binding the contract event 0xe42abf7d4a793286da8cc1399cb577a1f5a0e133dfee371bb3a5abbdd77b011e.
//
// Solidity: event AgentDataUpdate(uint256 indexed agentId, uint256 promptIndex, bytes oldSysPrompt, bytes newSysPrompt)
func (_Dagent721 *Dagent721Filterer) WatchAgentDataUpdate(opts *bind.WatchOpts, sink chan<- *Dagent721AgentDataUpdate, agentId []*big.Int) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _Dagent721.contract.WatchLogs(opts, "AgentDataUpdate", agentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dagent721AgentDataUpdate)
				if err := _Dagent721.contract.UnpackLog(event, "AgentDataUpdate", log); err != nil {
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

// ParseAgentDataUpdate is a log parse operation binding the contract event 0xe42abf7d4a793286da8cc1399cb577a1f5a0e133dfee371bb3a5abbdd77b011e.
//
// Solidity: event AgentDataUpdate(uint256 indexed agentId, uint256 promptIndex, bytes oldSysPrompt, bytes newSysPrompt)
func (_Dagent721 *Dagent721Filterer) ParseAgentDataUpdate(log types.Log) (*Dagent721AgentDataUpdate, error) {
	event := new(Dagent721AgentDataUpdate)
	if err := _Dagent721.contract.UnpackLog(event, "AgentDataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Dagent721AgentFeeUpdateIterator is returned from FilterAgentFeeUpdate and is used to iterate over the raw logs and unpacked data for AgentFeeUpdate events raised by the Dagent721 contract.
type Dagent721AgentFeeUpdateIterator struct {
	Event *Dagent721AgentFeeUpdate // Event containing the contract specifics and raw log

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
func (it *Dagent721AgentFeeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dagent721AgentFeeUpdate)
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
		it.Event = new(Dagent721AgentFeeUpdate)
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
func (it *Dagent721AgentFeeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dagent721AgentFeeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dagent721AgentFeeUpdate represents a AgentFeeUpdate event raised by the Dagent721 contract.
type Dagent721AgentFeeUpdate struct {
	AgentId *big.Int
	Fee     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentFeeUpdate is a free log retrieval operation binding the contract event 0xa08d8197034aee8915921dd8aa7d95cf711690dd77f0b676dded49b3f9a632d1.
//
// Solidity: event AgentFeeUpdate(uint256 indexed agentId, uint256 fee)
func (_Dagent721 *Dagent721Filterer) FilterAgentFeeUpdate(opts *bind.FilterOpts, agentId []*big.Int) (*Dagent721AgentFeeUpdateIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _Dagent721.contract.FilterLogs(opts, "AgentFeeUpdate", agentIdRule)
	if err != nil {
		return nil, err
	}
	return &Dagent721AgentFeeUpdateIterator{contract: _Dagent721.contract, event: "AgentFeeUpdate", logs: logs, sub: sub}, nil
}

// WatchAgentFeeUpdate is a free log subscription operation binding the contract event 0xa08d8197034aee8915921dd8aa7d95cf711690dd77f0b676dded49b3f9a632d1.
//
// Solidity: event AgentFeeUpdate(uint256 indexed agentId, uint256 fee)
func (_Dagent721 *Dagent721Filterer) WatchAgentFeeUpdate(opts *bind.WatchOpts, sink chan<- *Dagent721AgentFeeUpdate, agentId []*big.Int) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _Dagent721.contract.WatchLogs(opts, "AgentFeeUpdate", agentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dagent721AgentFeeUpdate)
				if err := _Dagent721.contract.UnpackLog(event, "AgentFeeUpdate", log); err != nil {
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

// ParseAgentFeeUpdate is a log parse operation binding the contract event 0xa08d8197034aee8915921dd8aa7d95cf711690dd77f0b676dded49b3f9a632d1.
//
// Solidity: event AgentFeeUpdate(uint256 indexed agentId, uint256 fee)
func (_Dagent721 *Dagent721Filterer) ParseAgentFeeUpdate(log types.Log) (*Dagent721AgentFeeUpdate, error) {
	event := new(Dagent721AgentFeeUpdate)
	if err := _Dagent721.contract.UnpackLog(event, "AgentFeeUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Dagent721AgentMissionAddNewIterator is returned from FilterAgentMissionAddNew and is used to iterate over the raw logs and unpacked data for AgentMissionAddNew events raised by the Dagent721 contract.
type Dagent721AgentMissionAddNewIterator struct {
	Event *Dagent721AgentMissionAddNew // Event containing the contract specifics and raw log

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
func (it *Dagent721AgentMissionAddNewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dagent721AgentMissionAddNew)
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
		it.Event = new(Dagent721AgentMissionAddNew)
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
func (it *Dagent721AgentMissionAddNewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dagent721AgentMissionAddNewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dagent721AgentMissionAddNew represents a AgentMissionAddNew event raised by the Dagent721 contract.
type Dagent721AgentMissionAddNew struct {
	AgentId  *big.Int
	Missions [][]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAgentMissionAddNew is a free log retrieval operation binding the contract event 0x12ccdcc9c8e92b22004686225bd3df163c042e77b03eab4566800d40b5047f91.
//
// Solidity: event AgentMissionAddNew(uint256 indexed agentId, bytes[] missions)
func (_Dagent721 *Dagent721Filterer) FilterAgentMissionAddNew(opts *bind.FilterOpts, agentId []*big.Int) (*Dagent721AgentMissionAddNewIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _Dagent721.contract.FilterLogs(opts, "AgentMissionAddNew", agentIdRule)
	if err != nil {
		return nil, err
	}
	return &Dagent721AgentMissionAddNewIterator{contract: _Dagent721.contract, event: "AgentMissionAddNew", logs: logs, sub: sub}, nil
}

// WatchAgentMissionAddNew is a free log subscription operation binding the contract event 0x12ccdcc9c8e92b22004686225bd3df163c042e77b03eab4566800d40b5047f91.
//
// Solidity: event AgentMissionAddNew(uint256 indexed agentId, bytes[] missions)
func (_Dagent721 *Dagent721Filterer) WatchAgentMissionAddNew(opts *bind.WatchOpts, sink chan<- *Dagent721AgentMissionAddNew, agentId []*big.Int) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _Dagent721.contract.WatchLogs(opts, "AgentMissionAddNew", agentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dagent721AgentMissionAddNew)
				if err := _Dagent721.contract.UnpackLog(event, "AgentMissionAddNew", log); err != nil {
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

// ParseAgentMissionAddNew is a log parse operation binding the contract event 0x12ccdcc9c8e92b22004686225bd3df163c042e77b03eab4566800d40b5047f91.
//
// Solidity: event AgentMissionAddNew(uint256 indexed agentId, bytes[] missions)
func (_Dagent721 *Dagent721Filterer) ParseAgentMissionAddNew(log types.Log) (*Dagent721AgentMissionAddNew, error) {
	event := new(Dagent721AgentMissionAddNew)
	if err := _Dagent721.contract.UnpackLog(event, "AgentMissionAddNew", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Dagent721AgentMissionUpdateIterator is returned from FilterAgentMissionUpdate and is used to iterate over the raw logs and unpacked data for AgentMissionUpdate events raised by the Dagent721 contract.
type Dagent721AgentMissionUpdateIterator struct {
	Event *Dagent721AgentMissionUpdate // Event containing the contract specifics and raw log

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
func (it *Dagent721AgentMissionUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dagent721AgentMissionUpdate)
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
		it.Event = new(Dagent721AgentMissionUpdate)
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
func (it *Dagent721AgentMissionUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dagent721AgentMissionUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dagent721AgentMissionUpdate represents a AgentMissionUpdate event raised by the Dagent721 contract.
type Dagent721AgentMissionUpdate struct {
	AgentId       *big.Int
	MissionIndex  *big.Int
	OldSysMission []byte
	NewSysMission []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAgentMissionUpdate is a free log retrieval operation binding the contract event 0x0a9b80bd675e3f5788f1a5da687efd147dbc4729245a7f300ce1074bbd535127.
//
// Solidity: event AgentMissionUpdate(uint256 indexed agentId, uint256 missionIndex, bytes oldSysMission, bytes newSysMission)
func (_Dagent721 *Dagent721Filterer) FilterAgentMissionUpdate(opts *bind.FilterOpts, agentId []*big.Int) (*Dagent721AgentMissionUpdateIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _Dagent721.contract.FilterLogs(opts, "AgentMissionUpdate", agentIdRule)
	if err != nil {
		return nil, err
	}
	return &Dagent721AgentMissionUpdateIterator{contract: _Dagent721.contract, event: "AgentMissionUpdate", logs: logs, sub: sub}, nil
}

// WatchAgentMissionUpdate is a free log subscription operation binding the contract event 0x0a9b80bd675e3f5788f1a5da687efd147dbc4729245a7f300ce1074bbd535127.
//
// Solidity: event AgentMissionUpdate(uint256 indexed agentId, uint256 missionIndex, bytes oldSysMission, bytes newSysMission)
func (_Dagent721 *Dagent721Filterer) WatchAgentMissionUpdate(opts *bind.WatchOpts, sink chan<- *Dagent721AgentMissionUpdate, agentId []*big.Int) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _Dagent721.contract.WatchLogs(opts, "AgentMissionUpdate", agentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dagent721AgentMissionUpdate)
				if err := _Dagent721.contract.UnpackLog(event, "AgentMissionUpdate", log); err != nil {
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

// ParseAgentMissionUpdate is a log parse operation binding the contract event 0x0a9b80bd675e3f5788f1a5da687efd147dbc4729245a7f300ce1074bbd535127.
//
// Solidity: event AgentMissionUpdate(uint256 indexed agentId, uint256 missionIndex, bytes oldSysMission, bytes newSysMission)
func (_Dagent721 *Dagent721Filterer) ParseAgentMissionUpdate(log types.Log) (*Dagent721AgentMissionUpdate, error) {
	event := new(Dagent721AgentMissionUpdate)
	if err := _Dagent721.contract.UnpackLog(event, "AgentMissionUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Dagent721AgentModelIdUpdateIterator is returned from FilterAgentModelIdUpdate and is used to iterate over the raw logs and unpacked data for AgentModelIdUpdate events raised by the Dagent721 contract.
type Dagent721AgentModelIdUpdateIterator struct {
	Event *Dagent721AgentModelIdUpdate // Event containing the contract specifics and raw log

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
func (it *Dagent721AgentModelIdUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dagent721AgentModelIdUpdate)
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
		it.Event = new(Dagent721AgentModelIdUpdate)
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
func (it *Dagent721AgentModelIdUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dagent721AgentModelIdUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dagent721AgentModelIdUpdate represents a AgentModelIdUpdate event raised by the Dagent721 contract.
type Dagent721AgentModelIdUpdate struct {
	AgentId    *big.Int
	OldModelId *big.Int
	NewModelId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAgentModelIdUpdate is a free log retrieval operation binding the contract event 0xe8662b9bac978d6b361a9cc824ecf5a8ea4cfb61ccbdd24dec6237ee9b7d7fa7.
//
// Solidity: event AgentModelIdUpdate(uint256 indexed agentId, uint256 oldModelId, uint256 newModelId)
func (_Dagent721 *Dagent721Filterer) FilterAgentModelIdUpdate(opts *bind.FilterOpts, agentId []*big.Int) (*Dagent721AgentModelIdUpdateIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _Dagent721.contract.FilterLogs(opts, "AgentModelIdUpdate", agentIdRule)
	if err != nil {
		return nil, err
	}
	return &Dagent721AgentModelIdUpdateIterator{contract: _Dagent721.contract, event: "AgentModelIdUpdate", logs: logs, sub: sub}, nil
}

// WatchAgentModelIdUpdate is a free log subscription operation binding the contract event 0xe8662b9bac978d6b361a9cc824ecf5a8ea4cfb61ccbdd24dec6237ee9b7d7fa7.
//
// Solidity: event AgentModelIdUpdate(uint256 indexed agentId, uint256 oldModelId, uint256 newModelId)
func (_Dagent721 *Dagent721Filterer) WatchAgentModelIdUpdate(opts *bind.WatchOpts, sink chan<- *Dagent721AgentModelIdUpdate, agentId []*big.Int) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _Dagent721.contract.WatchLogs(opts, "AgentModelIdUpdate", agentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dagent721AgentModelIdUpdate)
				if err := _Dagent721.contract.UnpackLog(event, "AgentModelIdUpdate", log); err != nil {
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

// ParseAgentModelIdUpdate is a log parse operation binding the contract event 0xe8662b9bac978d6b361a9cc824ecf5a8ea4cfb61ccbdd24dec6237ee9b7d7fa7.
//
// Solidity: event AgentModelIdUpdate(uint256 indexed agentId, uint256 oldModelId, uint256 newModelId)
func (_Dagent721 *Dagent721Filterer) ParseAgentModelIdUpdate(log types.Log) (*Dagent721AgentModelIdUpdate, error) {
	event := new(Dagent721AgentModelIdUpdate)
	if err := _Dagent721.contract.UnpackLog(event, "AgentModelIdUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Dagent721AgentPromptSchedulerdUpdateIterator is returned from FilterAgentPromptSchedulerdUpdate and is used to iterate over the raw logs and unpacked data for AgentPromptSchedulerdUpdate events raised by the Dagent721 contract.
type Dagent721AgentPromptSchedulerdUpdateIterator struct {
	Event *Dagent721AgentPromptSchedulerdUpdate // Event containing the contract specifics and raw log

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
func (it *Dagent721AgentPromptSchedulerdUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dagent721AgentPromptSchedulerdUpdate)
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
		it.Event = new(Dagent721AgentPromptSchedulerdUpdate)
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
func (it *Dagent721AgentPromptSchedulerdUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dagent721AgentPromptSchedulerdUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dagent721AgentPromptSchedulerdUpdate represents a AgentPromptSchedulerdUpdate event raised by the Dagent721 contract.
type Dagent721AgentPromptSchedulerdUpdate struct {
	AgentId               *big.Int
	OldPromptScheduler    common.Address
	NewOldPromptScheduler common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterAgentPromptSchedulerdUpdate is a free log retrieval operation binding the contract event 0x668af5d324df41df4dbf51978e1caa591bcf48468550da656523572a47d9abbd.
//
// Solidity: event AgentPromptSchedulerdUpdate(uint256 indexed agentId, address oldPromptScheduler, address newOldPromptScheduler)
func (_Dagent721 *Dagent721Filterer) FilterAgentPromptSchedulerdUpdate(opts *bind.FilterOpts, agentId []*big.Int) (*Dagent721AgentPromptSchedulerdUpdateIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _Dagent721.contract.FilterLogs(opts, "AgentPromptSchedulerdUpdate", agentIdRule)
	if err != nil {
		return nil, err
	}
	return &Dagent721AgentPromptSchedulerdUpdateIterator{contract: _Dagent721.contract, event: "AgentPromptSchedulerdUpdate", logs: logs, sub: sub}, nil
}

// WatchAgentPromptSchedulerdUpdate is a free log subscription operation binding the contract event 0x668af5d324df41df4dbf51978e1caa591bcf48468550da656523572a47d9abbd.
//
// Solidity: event AgentPromptSchedulerdUpdate(uint256 indexed agentId, address oldPromptScheduler, address newOldPromptScheduler)
func (_Dagent721 *Dagent721Filterer) WatchAgentPromptSchedulerdUpdate(opts *bind.WatchOpts, sink chan<- *Dagent721AgentPromptSchedulerdUpdate, agentId []*big.Int) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _Dagent721.contract.WatchLogs(opts, "AgentPromptSchedulerdUpdate", agentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dagent721AgentPromptSchedulerdUpdate)
				if err := _Dagent721.contract.UnpackLog(event, "AgentPromptSchedulerdUpdate", log); err != nil {
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

// ParseAgentPromptSchedulerdUpdate is a log parse operation binding the contract event 0x668af5d324df41df4dbf51978e1caa591bcf48468550da656523572a47d9abbd.
//
// Solidity: event AgentPromptSchedulerdUpdate(uint256 indexed agentId, address oldPromptScheduler, address newOldPromptScheduler)
func (_Dagent721 *Dagent721Filterer) ParseAgentPromptSchedulerdUpdate(log types.Log) (*Dagent721AgentPromptSchedulerdUpdate, error) {
	event := new(Dagent721AgentPromptSchedulerdUpdate)
	if err := _Dagent721.contract.UnpackLog(event, "AgentPromptSchedulerdUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Dagent721AgentURIUpdateIterator is returned from FilterAgentURIUpdate and is used to iterate over the raw logs and unpacked data for AgentURIUpdate events raised by the Dagent721 contract.
type Dagent721AgentURIUpdateIterator struct {
	Event *Dagent721AgentURIUpdate // Event containing the contract specifics and raw log

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
func (it *Dagent721AgentURIUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dagent721AgentURIUpdate)
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
		it.Event = new(Dagent721AgentURIUpdate)
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
func (it *Dagent721AgentURIUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dagent721AgentURIUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dagent721AgentURIUpdate represents a AgentURIUpdate event raised by the Dagent721 contract.
type Dagent721AgentURIUpdate struct {
	AgentId *big.Int
	Uri     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentURIUpdate is a free log retrieval operation binding the contract event 0x706a4e8eb2f354c7f4d96e5ea1984f36e72482629987edad78c9940ea037c362.
//
// Solidity: event AgentURIUpdate(uint256 indexed agentId, string uri)
func (_Dagent721 *Dagent721Filterer) FilterAgentURIUpdate(opts *bind.FilterOpts, agentId []*big.Int) (*Dagent721AgentURIUpdateIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _Dagent721.contract.FilterLogs(opts, "AgentURIUpdate", agentIdRule)
	if err != nil {
		return nil, err
	}
	return &Dagent721AgentURIUpdateIterator{contract: _Dagent721.contract, event: "AgentURIUpdate", logs: logs, sub: sub}, nil
}

// WatchAgentURIUpdate is a free log subscription operation binding the contract event 0x706a4e8eb2f354c7f4d96e5ea1984f36e72482629987edad78c9940ea037c362.
//
// Solidity: event AgentURIUpdate(uint256 indexed agentId, string uri)
func (_Dagent721 *Dagent721Filterer) WatchAgentURIUpdate(opts *bind.WatchOpts, sink chan<- *Dagent721AgentURIUpdate, agentId []*big.Int) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _Dagent721.contract.WatchLogs(opts, "AgentURIUpdate", agentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dagent721AgentURIUpdate)
				if err := _Dagent721.contract.UnpackLog(event, "AgentURIUpdate", log); err != nil {
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

// ParseAgentURIUpdate is a log parse operation binding the contract event 0x706a4e8eb2f354c7f4d96e5ea1984f36e72482629987edad78c9940ea037c362.
//
// Solidity: event AgentURIUpdate(uint256 indexed agentId, string uri)
func (_Dagent721 *Dagent721Filterer) ParseAgentURIUpdate(log types.Log) (*Dagent721AgentURIUpdate, error) {
	event := new(Dagent721AgentURIUpdate)
	if err := _Dagent721.contract.UnpackLog(event, "AgentURIUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Dagent721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Dagent721 contract.
type Dagent721ApprovalIterator struct {
	Event *Dagent721Approval // Event containing the contract specifics and raw log

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
func (it *Dagent721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dagent721Approval)
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
		it.Event = new(Dagent721Approval)
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
func (it *Dagent721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dagent721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dagent721Approval represents a Approval event raised by the Dagent721 contract.
type Dagent721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Dagent721 *Dagent721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*Dagent721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Dagent721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Dagent721ApprovalIterator{contract: _Dagent721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Dagent721 *Dagent721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Dagent721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Dagent721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dagent721Approval)
				if err := _Dagent721.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Dagent721 *Dagent721Filterer) ParseApproval(log types.Log) (*Dagent721Approval, error) {
	event := new(Dagent721Approval)
	if err := _Dagent721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Dagent721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Dagent721 contract.
type Dagent721ApprovalForAllIterator struct {
	Event *Dagent721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Dagent721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dagent721ApprovalForAll)
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
		it.Event = new(Dagent721ApprovalForAll)
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
func (it *Dagent721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dagent721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dagent721ApprovalForAll represents a ApprovalForAll event raised by the Dagent721 contract.
type Dagent721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Dagent721 *Dagent721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*Dagent721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Dagent721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Dagent721ApprovalForAllIterator{contract: _Dagent721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Dagent721 *Dagent721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Dagent721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Dagent721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dagent721ApprovalForAll)
				if err := _Dagent721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Dagent721 *Dagent721Filterer) ParseApprovalForAll(log types.Log) (*Dagent721ApprovalForAll, error) {
	event := new(Dagent721ApprovalForAll)
	if err := _Dagent721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Dagent721BatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the Dagent721 contract.
type Dagent721BatchMetadataUpdateIterator struct {
	Event *Dagent721BatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *Dagent721BatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dagent721BatchMetadataUpdate)
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
		it.Event = new(Dagent721BatchMetadataUpdate)
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
func (it *Dagent721BatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dagent721BatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dagent721BatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Dagent721 contract.
type Dagent721BatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Dagent721 *Dagent721Filterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*Dagent721BatchMetadataUpdateIterator, error) {

	logs, sub, err := _Dagent721.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &Dagent721BatchMetadataUpdateIterator{contract: _Dagent721.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Dagent721 *Dagent721Filterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *Dagent721BatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Dagent721.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dagent721BatchMetadataUpdate)
				if err := _Dagent721.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Dagent721 *Dagent721Filterer) ParseBatchMetadataUpdate(log types.Log) (*Dagent721BatchMetadataUpdate, error) {
	event := new(Dagent721BatchMetadataUpdate)
	if err := _Dagent721.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Dagent721InferencePerformedIterator is returned from FilterInferencePerformed and is used to iterate over the raw logs and unpacked data for InferencePerformed events raised by the Dagent721 contract.
type Dagent721InferencePerformedIterator struct {
	Event *Dagent721InferencePerformed // Event containing the contract specifics and raw log

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
func (it *Dagent721InferencePerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dagent721InferencePerformed)
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
		it.Event = new(Dagent721InferencePerformed)
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
func (it *Dagent721InferencePerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dagent721InferencePerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dagent721InferencePerformed represents a InferencePerformed event raised by the Dagent721 contract.
type Dagent721InferencePerformed struct {
	TokenId      *big.Int
	Caller       common.Address
	Data         []byte
	Fee          *big.Int
	ExternalData string
	InferenceId  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInferencePerformed is a free log retrieval operation binding the contract event 0xcf35460eca25a0549d5eb14c712236d61c9a0bad90c834f996c5f2a98d332719.
//
// Solidity: event InferencePerformed(uint256 indexed tokenId, address indexed caller, bytes data, uint256 fee, string externalData, uint256 inferenceId)
func (_Dagent721 *Dagent721Filterer) FilterInferencePerformed(opts *bind.FilterOpts, tokenId []*big.Int, caller []common.Address) (*Dagent721InferencePerformedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Dagent721.contract.FilterLogs(opts, "InferencePerformed", tokenIdRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &Dagent721InferencePerformedIterator{contract: _Dagent721.contract, event: "InferencePerformed", logs: logs, sub: sub}, nil
}

// WatchInferencePerformed is a free log subscription operation binding the contract event 0xcf35460eca25a0549d5eb14c712236d61c9a0bad90c834f996c5f2a98d332719.
//
// Solidity: event InferencePerformed(uint256 indexed tokenId, address indexed caller, bytes data, uint256 fee, string externalData, uint256 inferenceId)
func (_Dagent721 *Dagent721Filterer) WatchInferencePerformed(opts *bind.WatchOpts, sink chan<- *Dagent721InferencePerformed, tokenId []*big.Int, caller []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Dagent721.contract.WatchLogs(opts, "InferencePerformed", tokenIdRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dagent721InferencePerformed)
				if err := _Dagent721.contract.UnpackLog(event, "InferencePerformed", log); err != nil {
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

// ParseInferencePerformed is a log parse operation binding the contract event 0xcf35460eca25a0549d5eb14c712236d61c9a0bad90c834f996c5f2a98d332719.
//
// Solidity: event InferencePerformed(uint256 indexed tokenId, address indexed caller, bytes data, uint256 fee, string externalData, uint256 inferenceId)
func (_Dagent721 *Dagent721Filterer) ParseInferencePerformed(log types.Log) (*Dagent721InferencePerformed, error) {
	event := new(Dagent721InferencePerformed)
	if err := _Dagent721.contract.UnpackLog(event, "InferencePerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Dagent721InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Dagent721 contract.
type Dagent721InitializedIterator struct {
	Event *Dagent721Initialized // Event containing the contract specifics and raw log

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
func (it *Dagent721InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dagent721Initialized)
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
		it.Event = new(Dagent721Initialized)
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
func (it *Dagent721InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dagent721InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dagent721Initialized represents a Initialized event raised by the Dagent721 contract.
type Dagent721Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Dagent721 *Dagent721Filterer) FilterInitialized(opts *bind.FilterOpts) (*Dagent721InitializedIterator, error) {

	logs, sub, err := _Dagent721.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Dagent721InitializedIterator{contract: _Dagent721.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Dagent721 *Dagent721Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Dagent721Initialized) (event.Subscription, error) {

	logs, sub, err := _Dagent721.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dagent721Initialized)
				if err := _Dagent721.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Dagent721 *Dagent721Filterer) ParseInitialized(log types.Log) (*Dagent721Initialized, error) {
	event := new(Dagent721Initialized)
	if err := _Dagent721.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Dagent721MetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the Dagent721 contract.
type Dagent721MetadataUpdateIterator struct {
	Event *Dagent721MetadataUpdate // Event containing the contract specifics and raw log

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
func (it *Dagent721MetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dagent721MetadataUpdate)
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
		it.Event = new(Dagent721MetadataUpdate)
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
func (it *Dagent721MetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dagent721MetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dagent721MetadataUpdate represents a MetadataUpdate event raised by the Dagent721 contract.
type Dagent721MetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Dagent721 *Dagent721Filterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*Dagent721MetadataUpdateIterator, error) {

	logs, sub, err := _Dagent721.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &Dagent721MetadataUpdateIterator{contract: _Dagent721.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Dagent721 *Dagent721Filterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *Dagent721MetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Dagent721.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dagent721MetadataUpdate)
				if err := _Dagent721.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Dagent721 *Dagent721Filterer) ParseMetadataUpdate(log types.Log) (*Dagent721MetadataUpdate, error) {
	event := new(Dagent721MetadataUpdate)
	if err := _Dagent721.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Dagent721MintPriceUpdateIterator is returned from FilterMintPriceUpdate and is used to iterate over the raw logs and unpacked data for MintPriceUpdate events raised by the Dagent721 contract.
type Dagent721MintPriceUpdateIterator struct {
	Event *Dagent721MintPriceUpdate // Event containing the contract specifics and raw log

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
func (it *Dagent721MintPriceUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dagent721MintPriceUpdate)
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
		it.Event = new(Dagent721MintPriceUpdate)
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
func (it *Dagent721MintPriceUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dagent721MintPriceUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dagent721MintPriceUpdate represents a MintPriceUpdate event raised by the Dagent721 contract.
type Dagent721MintPriceUpdate struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMintPriceUpdate is a free log retrieval operation binding the contract event 0x23050b539195eebd064c1bec834445b7d028a20c345600e868a74d7ca93c5e86.
//
// Solidity: event MintPriceUpdate(uint256 newValue)
func (_Dagent721 *Dagent721Filterer) FilterMintPriceUpdate(opts *bind.FilterOpts) (*Dagent721MintPriceUpdateIterator, error) {

	logs, sub, err := _Dagent721.contract.FilterLogs(opts, "MintPriceUpdate")
	if err != nil {
		return nil, err
	}
	return &Dagent721MintPriceUpdateIterator{contract: _Dagent721.contract, event: "MintPriceUpdate", logs: logs, sub: sub}, nil
}

// WatchMintPriceUpdate is a free log subscription operation binding the contract event 0x23050b539195eebd064c1bec834445b7d028a20c345600e868a74d7ca93c5e86.
//
// Solidity: event MintPriceUpdate(uint256 newValue)
func (_Dagent721 *Dagent721Filterer) WatchMintPriceUpdate(opts *bind.WatchOpts, sink chan<- *Dagent721MintPriceUpdate) (event.Subscription, error) {

	logs, sub, err := _Dagent721.contract.WatchLogs(opts, "MintPriceUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dagent721MintPriceUpdate)
				if err := _Dagent721.contract.UnpackLog(event, "MintPriceUpdate", log); err != nil {
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

// ParseMintPriceUpdate is a log parse operation binding the contract event 0x23050b539195eebd064c1bec834445b7d028a20c345600e868a74d7ca93c5e86.
//
// Solidity: event MintPriceUpdate(uint256 newValue)
func (_Dagent721 *Dagent721Filterer) ParseMintPriceUpdate(log types.Log) (*Dagent721MintPriceUpdate, error) {
	event := new(Dagent721MintPriceUpdate)
	if err := _Dagent721.contract.UnpackLog(event, "MintPriceUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Dagent721NewTokenIterator is returned from FilterNewToken and is used to iterate over the raw logs and unpacked data for NewToken events raised by the Dagent721 contract.
type Dagent721NewTokenIterator struct {
	Event *Dagent721NewToken // Event containing the contract specifics and raw log

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
func (it *Dagent721NewTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dagent721NewToken)
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
		it.Event = new(Dagent721NewToken)
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
func (it *Dagent721NewTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dagent721NewTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dagent721NewToken represents a NewToken event raised by the Dagent721 contract.
type Dagent721NewToken struct {
	TokenId   *big.Int
	Uri       string
	SysPrompt []byte
	Fee       *big.Int
	Minter    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewToken is a free log retrieval operation binding the contract event 0x61beab98a81083e3c0239c33e149bef1316ca78f15b9f29125039f5521a06d06.
//
// Solidity: event NewToken(uint256 indexed tokenId, string uri, bytes sysPrompt, uint256 fee, address indexed minter)
func (_Dagent721 *Dagent721Filterer) FilterNewToken(opts *bind.FilterOpts, tokenId []*big.Int, minter []common.Address) (*Dagent721NewTokenIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _Dagent721.contract.FilterLogs(opts, "NewToken", tokenIdRule, minterRule)
	if err != nil {
		return nil, err
	}
	return &Dagent721NewTokenIterator{contract: _Dagent721.contract, event: "NewToken", logs: logs, sub: sub}, nil
}

// WatchNewToken is a free log subscription operation binding the contract event 0x61beab98a81083e3c0239c33e149bef1316ca78f15b9f29125039f5521a06d06.
//
// Solidity: event NewToken(uint256 indexed tokenId, string uri, bytes sysPrompt, uint256 fee, address indexed minter)
func (_Dagent721 *Dagent721Filterer) WatchNewToken(opts *bind.WatchOpts, sink chan<- *Dagent721NewToken, tokenId []*big.Int, minter []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _Dagent721.contract.WatchLogs(opts, "NewToken", tokenIdRule, minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dagent721NewToken)
				if err := _Dagent721.contract.UnpackLog(event, "NewToken", log); err != nil {
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

// ParseNewToken is a log parse operation binding the contract event 0x61beab98a81083e3c0239c33e149bef1316ca78f15b9f29125039f5521a06d06.
//
// Solidity: event NewToken(uint256 indexed tokenId, string uri, bytes sysPrompt, uint256 fee, address indexed minter)
func (_Dagent721 *Dagent721Filterer) ParseNewToken(log types.Log) (*Dagent721NewToken, error) {
	event := new(Dagent721NewToken)
	if err := _Dagent721.contract.UnpackLog(event, "NewToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Dagent721OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dagent721 contract.
type Dagent721OwnershipTransferredIterator struct {
	Event *Dagent721OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Dagent721OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dagent721OwnershipTransferred)
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
		it.Event = new(Dagent721OwnershipTransferred)
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
func (it *Dagent721OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dagent721OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dagent721OwnershipTransferred represents a OwnershipTransferred event raised by the Dagent721 contract.
type Dagent721OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dagent721 *Dagent721Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Dagent721OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dagent721.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Dagent721OwnershipTransferredIterator{contract: _Dagent721.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dagent721 *Dagent721Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Dagent721OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dagent721.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dagent721OwnershipTransferred)
				if err := _Dagent721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Dagent721 *Dagent721Filterer) ParseOwnershipTransferred(log types.Log) (*Dagent721OwnershipTransferred, error) {
	event := new(Dagent721OwnershipTransferred)
	if err := _Dagent721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Dagent721PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Dagent721 contract.
type Dagent721PausedIterator struct {
	Event *Dagent721Paused // Event containing the contract specifics and raw log

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
func (it *Dagent721PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dagent721Paused)
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
		it.Event = new(Dagent721Paused)
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
func (it *Dagent721PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dagent721PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dagent721Paused represents a Paused event raised by the Dagent721 contract.
type Dagent721Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Dagent721 *Dagent721Filterer) FilterPaused(opts *bind.FilterOpts) (*Dagent721PausedIterator, error) {

	logs, sub, err := _Dagent721.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &Dagent721PausedIterator{contract: _Dagent721.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Dagent721 *Dagent721Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *Dagent721Paused) (event.Subscription, error) {

	logs, sub, err := _Dagent721.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dagent721Paused)
				if err := _Dagent721.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Dagent721 *Dagent721Filterer) ParsePaused(log types.Log) (*Dagent721Paused, error) {
	event := new(Dagent721Paused)
	if err := _Dagent721.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Dagent721RoyaltyPortionUpdateIterator is returned from FilterRoyaltyPortionUpdate and is used to iterate over the raw logs and unpacked data for RoyaltyPortionUpdate events raised by the Dagent721 contract.
type Dagent721RoyaltyPortionUpdateIterator struct {
	Event *Dagent721RoyaltyPortionUpdate // Event containing the contract specifics and raw log

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
func (it *Dagent721RoyaltyPortionUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dagent721RoyaltyPortionUpdate)
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
		it.Event = new(Dagent721RoyaltyPortionUpdate)
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
func (it *Dagent721RoyaltyPortionUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dagent721RoyaltyPortionUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dagent721RoyaltyPortionUpdate represents a RoyaltyPortionUpdate event raised by the Dagent721 contract.
type Dagent721RoyaltyPortionUpdate struct {
	NewValue uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyPortionUpdate is a free log retrieval operation binding the contract event 0xb1f3037624bd2d961f6d56696cc10ccc3a676db381e671b1bc58f0ab1f686dd5.
//
// Solidity: event RoyaltyPortionUpdate(uint16 newValue)
func (_Dagent721 *Dagent721Filterer) FilterRoyaltyPortionUpdate(opts *bind.FilterOpts) (*Dagent721RoyaltyPortionUpdateIterator, error) {

	logs, sub, err := _Dagent721.contract.FilterLogs(opts, "RoyaltyPortionUpdate")
	if err != nil {
		return nil, err
	}
	return &Dagent721RoyaltyPortionUpdateIterator{contract: _Dagent721.contract, event: "RoyaltyPortionUpdate", logs: logs, sub: sub}, nil
}

// WatchRoyaltyPortionUpdate is a free log subscription operation binding the contract event 0xb1f3037624bd2d961f6d56696cc10ccc3a676db381e671b1bc58f0ab1f686dd5.
//
// Solidity: event RoyaltyPortionUpdate(uint16 newValue)
func (_Dagent721 *Dagent721Filterer) WatchRoyaltyPortionUpdate(opts *bind.WatchOpts, sink chan<- *Dagent721RoyaltyPortionUpdate) (event.Subscription, error) {

	logs, sub, err := _Dagent721.contract.WatchLogs(opts, "RoyaltyPortionUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dagent721RoyaltyPortionUpdate)
				if err := _Dagent721.contract.UnpackLog(event, "RoyaltyPortionUpdate", log); err != nil {
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

// ParseRoyaltyPortionUpdate is a log parse operation binding the contract event 0xb1f3037624bd2d961f6d56696cc10ccc3a676db381e671b1bc58f0ab1f686dd5.
//
// Solidity: event RoyaltyPortionUpdate(uint16 newValue)
func (_Dagent721 *Dagent721Filterer) ParseRoyaltyPortionUpdate(log types.Log) (*Dagent721RoyaltyPortionUpdate, error) {
	event := new(Dagent721RoyaltyPortionUpdate)
	if err := _Dagent721.contract.UnpackLog(event, "RoyaltyPortionUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Dagent721RoyaltyReceiverUpdateIterator is returned from FilterRoyaltyReceiverUpdate and is used to iterate over the raw logs and unpacked data for RoyaltyReceiverUpdate events raised by the Dagent721 contract.
type Dagent721RoyaltyReceiverUpdateIterator struct {
	Event *Dagent721RoyaltyReceiverUpdate // Event containing the contract specifics and raw log

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
func (it *Dagent721RoyaltyReceiverUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dagent721RoyaltyReceiverUpdate)
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
		it.Event = new(Dagent721RoyaltyReceiverUpdate)
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
func (it *Dagent721RoyaltyReceiverUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dagent721RoyaltyReceiverUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dagent721RoyaltyReceiverUpdate represents a RoyaltyReceiverUpdate event raised by the Dagent721 contract.
type Dagent721RoyaltyReceiverUpdate struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyReceiverUpdate is a free log retrieval operation binding the contract event 0xec6b72b10aed766af02b35918b55be261c89aaaa4c8add826471ce35ec7f97b3.
//
// Solidity: event RoyaltyReceiverUpdate(address newAddress)
func (_Dagent721 *Dagent721Filterer) FilterRoyaltyReceiverUpdate(opts *bind.FilterOpts) (*Dagent721RoyaltyReceiverUpdateIterator, error) {

	logs, sub, err := _Dagent721.contract.FilterLogs(opts, "RoyaltyReceiverUpdate")
	if err != nil {
		return nil, err
	}
	return &Dagent721RoyaltyReceiverUpdateIterator{contract: _Dagent721.contract, event: "RoyaltyReceiverUpdate", logs: logs, sub: sub}, nil
}

// WatchRoyaltyReceiverUpdate is a free log subscription operation binding the contract event 0xec6b72b10aed766af02b35918b55be261c89aaaa4c8add826471ce35ec7f97b3.
//
// Solidity: event RoyaltyReceiverUpdate(address newAddress)
func (_Dagent721 *Dagent721Filterer) WatchRoyaltyReceiverUpdate(opts *bind.WatchOpts, sink chan<- *Dagent721RoyaltyReceiverUpdate) (event.Subscription, error) {

	logs, sub, err := _Dagent721.contract.WatchLogs(opts, "RoyaltyReceiverUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dagent721RoyaltyReceiverUpdate)
				if err := _Dagent721.contract.UnpackLog(event, "RoyaltyReceiverUpdate", log); err != nil {
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

// ParseRoyaltyReceiverUpdate is a log parse operation binding the contract event 0xec6b72b10aed766af02b35918b55be261c89aaaa4c8add826471ce35ec7f97b3.
//
// Solidity: event RoyaltyReceiverUpdate(address newAddress)
func (_Dagent721 *Dagent721Filterer) ParseRoyaltyReceiverUpdate(log types.Log) (*Dagent721RoyaltyReceiverUpdate, error) {
	event := new(Dagent721RoyaltyReceiverUpdate)
	if err := _Dagent721.contract.UnpackLog(event, "RoyaltyReceiverUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Dagent721TopUpPoolBalanceIterator is returned from FilterTopUpPoolBalance and is used to iterate over the raw logs and unpacked data for TopUpPoolBalance events raised by the Dagent721 contract.
type Dagent721TopUpPoolBalanceIterator struct {
	Event *Dagent721TopUpPoolBalance // Event containing the contract specifics and raw log

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
func (it *Dagent721TopUpPoolBalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dagent721TopUpPoolBalance)
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
		it.Event = new(Dagent721TopUpPoolBalance)
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
func (it *Dagent721TopUpPoolBalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dagent721TopUpPoolBalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dagent721TopUpPoolBalance represents a TopUpPoolBalance event raised by the Dagent721 contract.
type Dagent721TopUpPoolBalance struct {
	AgentId *big.Int
	Caller  common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTopUpPoolBalance is a free log retrieval operation binding the contract event 0xf7ee57effd30f2ab842c1d65fdcfa7d20c2fb2f967e9ac30acafecabf013ea4c.
//
// Solidity: event TopUpPoolBalance(uint256 agentId, address caller, uint256 amount)
func (_Dagent721 *Dagent721Filterer) FilterTopUpPoolBalance(opts *bind.FilterOpts) (*Dagent721TopUpPoolBalanceIterator, error) {

	logs, sub, err := _Dagent721.contract.FilterLogs(opts, "TopUpPoolBalance")
	if err != nil {
		return nil, err
	}
	return &Dagent721TopUpPoolBalanceIterator{contract: _Dagent721.contract, event: "TopUpPoolBalance", logs: logs, sub: sub}, nil
}

// WatchTopUpPoolBalance is a free log subscription operation binding the contract event 0xf7ee57effd30f2ab842c1d65fdcfa7d20c2fb2f967e9ac30acafecabf013ea4c.
//
// Solidity: event TopUpPoolBalance(uint256 agentId, address caller, uint256 amount)
func (_Dagent721 *Dagent721Filterer) WatchTopUpPoolBalance(opts *bind.WatchOpts, sink chan<- *Dagent721TopUpPoolBalance) (event.Subscription, error) {

	logs, sub, err := _Dagent721.contract.WatchLogs(opts, "TopUpPoolBalance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dagent721TopUpPoolBalance)
				if err := _Dagent721.contract.UnpackLog(event, "TopUpPoolBalance", log); err != nil {
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

// ParseTopUpPoolBalance is a log parse operation binding the contract event 0xf7ee57effd30f2ab842c1d65fdcfa7d20c2fb2f967e9ac30acafecabf013ea4c.
//
// Solidity: event TopUpPoolBalance(uint256 agentId, address caller, uint256 amount)
func (_Dagent721 *Dagent721Filterer) ParseTopUpPoolBalance(log types.Log) (*Dagent721TopUpPoolBalance, error) {
	event := new(Dagent721TopUpPoolBalance)
	if err := _Dagent721.contract.UnpackLog(event, "TopUpPoolBalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Dagent721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Dagent721 contract.
type Dagent721TransferIterator struct {
	Event *Dagent721Transfer // Event containing the contract specifics and raw log

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
func (it *Dagent721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dagent721Transfer)
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
		it.Event = new(Dagent721Transfer)
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
func (it *Dagent721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dagent721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dagent721Transfer represents a Transfer event raised by the Dagent721 contract.
type Dagent721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Dagent721 *Dagent721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*Dagent721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Dagent721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Dagent721TransferIterator{contract: _Dagent721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Dagent721 *Dagent721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Dagent721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Dagent721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dagent721Transfer)
				if err := _Dagent721.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Dagent721 *Dagent721Filterer) ParseTransfer(log types.Log) (*Dagent721Transfer, error) {
	event := new(Dagent721Transfer)
	if err := _Dagent721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Dagent721UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Dagent721 contract.
type Dagent721UnpausedIterator struct {
	Event *Dagent721Unpaused // Event containing the contract specifics and raw log

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
func (it *Dagent721UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dagent721Unpaused)
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
		it.Event = new(Dagent721Unpaused)
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
func (it *Dagent721UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dagent721UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dagent721Unpaused represents a Unpaused event raised by the Dagent721 contract.
type Dagent721Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Dagent721 *Dagent721Filterer) FilterUnpaused(opts *bind.FilterOpts) (*Dagent721UnpausedIterator, error) {

	logs, sub, err := _Dagent721.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &Dagent721UnpausedIterator{contract: _Dagent721.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Dagent721 *Dagent721Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *Dagent721Unpaused) (event.Subscription, error) {

	logs, sub, err := _Dagent721.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dagent721Unpaused)
				if err := _Dagent721.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Dagent721 *Dagent721Filterer) ParseUnpaused(log types.Log) (*Dagent721Unpaused, error) {
	event := new(Dagent721Unpaused)
	if err := _Dagent721.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
