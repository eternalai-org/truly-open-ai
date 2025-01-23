// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// IAI721TokenMetaData is an auto generated low-level Go binding around an user-defined struct.
type IAI721TokenMetaData struct {
	Fee        *big.Int
	SysPrompts [][]byte
}

// AI721ContractMetaData contains all meta data concerning the AI721Contract contract.
var AI721ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"Authorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentPromptIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentURI\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMintingFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureUsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"sysPrompt\",\"type\":\"bytes[]\"}],\"name\":\"AgentDataAddNew\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"promptIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"oldSysPrompt\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newSysPrompt\",\"type\":\"bytes\"}],\"name\":\"AgentDataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"AgentFeeUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"missions\",\"type\":\"bytes[]\"}],\"name\":\"AgentMissionAddNew\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"missionIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"oldSysMission\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newSysMission\",\"type\":\"bytes\"}],\"name\":\"AgentMissionUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"AgentURIUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"externalData\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"}],\"name\":\"InferencePerformed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ManagerAuthorization\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ManagerDeauthorization\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MintPriceUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"sysPrompt\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"NewToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newValue\",\"type\":\"uint16\"}],\"name\":\"RoyaltyPortionUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"RoyaltyReceiverUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TopUpPoolBalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_sysPrompt\",\"type\":\"bytes\"}],\"name\":\"addNewAgentData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"authorizeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_missionData\",\"type\":\"bytes\"}],\"name\":\"createMission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"}],\"name\":\"dataOf\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"sysPrompts\",\"type\":\"bytes[]\"}],\"internalType\":\"structIAI721.TokenMetaData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"deauthorizeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"}],\"name\":\"getAgentFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getAgentIdByOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"}],\"name\":\"getAgentSystemPrompt\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"}],\"name\":\"getMissionIdsByAgentId\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hybridModel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"_externalData\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_flag\",\"type\":\"bool\"}],\"name\":\"infer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"_externalData\",\"type\":\"string\"}],\"name\":\"infer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_royaltyReceiver\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_royaltyPortion\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_nextTokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_hybridModel\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_workerHub\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"name\":\"poolBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyPortion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"signaturesUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"}],\"name\":\"topUpPoolBalance\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_sysPrompt\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_promptIdx\",\"type\":\"uint256\"}],\"name\":\"updateAgentData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_sysPrompt\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_promptIdx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_randomNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"updateAgentDataWithSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"updateAgentFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"updateAgentURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_randomNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"updateAgentUriWithSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"}],\"name\":\"updateMintPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_missionIdx\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_missionData\",\"type\":\"bytes\"}],\"name\":\"updateMission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_royaltyPortion\",\"type\":\"uint16\"}],\"name\":\"updateRoyaltyPortion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_royaltyReceiver\",\"type\":\"address\"}],\"name\":\"updateRoyaltyReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"workerHub\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// AI721ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use AI721ContractMetaData.ABI instead.
var AI721ContractABI = AI721ContractMetaData.ABI

// AI721Contract is an auto generated Go binding around an Ethereum contract.
type AI721Contract struct {
	AI721ContractCaller     // Read-only binding to the contract
	AI721ContractTransactor // Write-only binding to the contract
	AI721ContractFilterer   // Log filterer for contract events
}

// AI721ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type AI721ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AI721ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AI721ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AI721ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AI721ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AI721ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AI721ContractSession struct {
	Contract     *AI721Contract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AI721ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AI721ContractCallerSession struct {
	Contract *AI721ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AI721ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AI721ContractTransactorSession struct {
	Contract     *AI721ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AI721ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type AI721ContractRaw struct {
	Contract *AI721Contract // Generic contract binding to access the raw methods on
}

// AI721ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AI721ContractCallerRaw struct {
	Contract *AI721ContractCaller // Generic read-only contract binding to access the raw methods on
}

// AI721ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AI721ContractTransactorRaw struct {
	Contract *AI721ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAI721Contract creates a new instance of AI721Contract, bound to a specific deployed contract.
func NewAI721Contract(address common.Address, backend bind.ContractBackend) (*AI721Contract, error) {
	contract, err := bindAI721Contract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AI721Contract{AI721ContractCaller: AI721ContractCaller{contract: contract}, AI721ContractTransactor: AI721ContractTransactor{contract: contract}, AI721ContractFilterer: AI721ContractFilterer{contract: contract}}, nil
}

// NewAI721ContractCaller creates a new read-only instance of AI721Contract, bound to a specific deployed contract.
func NewAI721ContractCaller(address common.Address, caller bind.ContractCaller) (*AI721ContractCaller, error) {
	contract, err := bindAI721Contract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AI721ContractCaller{contract: contract}, nil
}

// NewAI721ContractTransactor creates a new write-only instance of AI721Contract, bound to a specific deployed contract.
func NewAI721ContractTransactor(address common.Address, transactor bind.ContractTransactor) (*AI721ContractTransactor, error) {
	contract, err := bindAI721Contract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AI721ContractTransactor{contract: contract}, nil
}

// NewAI721ContractFilterer creates a new log filterer instance of AI721Contract, bound to a specific deployed contract.
func NewAI721ContractFilterer(address common.Address, filterer bind.ContractFilterer) (*AI721ContractFilterer, error) {
	contract, err := bindAI721Contract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AI721ContractFilterer{contract: contract}, nil
}

// bindAI721Contract binds a generic wrapper to an already deployed contract.
func bindAI721Contract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AI721ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AI721Contract *AI721ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AI721Contract.Contract.AI721ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AI721Contract *AI721ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AI721Contract.Contract.AI721ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AI721Contract *AI721ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AI721Contract.Contract.AI721ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AI721Contract *AI721ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AI721Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AI721Contract *AI721ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AI721Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AI721Contract *AI721ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AI721Contract.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_AI721Contract *AI721ContractCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_AI721Contract *AI721ContractSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _AI721Contract.Contract.BalanceOf(&_AI721Contract.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_AI721Contract *AI721ContractCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _AI721Contract.Contract.BalanceOf(&_AI721Contract.CallOpts, owner)
}

// DataOf is a free data retrieval call binding the contract method 0x871caa98.
//
// Solidity: function dataOf(uint256 _agentId) view returns((uint256,bytes[]))
func (_AI721Contract *AI721ContractCaller) DataOf(opts *bind.CallOpts, _agentId *big.Int) (IAI721TokenMetaData, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "dataOf", _agentId)

	if err != nil {
		return *new(IAI721TokenMetaData), err
	}

	out0 := *abi.ConvertType(out[0], new(IAI721TokenMetaData)).(*IAI721TokenMetaData)

	return out0, err

}

// DataOf is a free data retrieval call binding the contract method 0x871caa98.
//
// Solidity: function dataOf(uint256 _agentId) view returns((uint256,bytes[]))
func (_AI721Contract *AI721ContractSession) DataOf(_agentId *big.Int) (IAI721TokenMetaData, error) {
	return _AI721Contract.Contract.DataOf(&_AI721Contract.CallOpts, _agentId)
}

// DataOf is a free data retrieval call binding the contract method 0x871caa98.
//
// Solidity: function dataOf(uint256 _agentId) view returns((uint256,bytes[]))
func (_AI721Contract *AI721ContractCallerSession) DataOf(_agentId *big.Int) (IAI721TokenMetaData, error) {
	return _AI721Contract.Contract.DataOf(&_AI721Contract.CallOpts, _agentId)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_AI721Contract *AI721ContractCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_AI721Contract *AI721ContractSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _AI721Contract.Contract.Eip712Domain(&_AI721Contract.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_AI721Contract *AI721ContractCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _AI721Contract.Contract.Eip712Domain(&_AI721Contract.CallOpts)
}

// GetAgentFee is a free data retrieval call binding the contract method 0xed96f433.
//
// Solidity: function getAgentFee(uint256 _agentId) view returns(uint256)
func (_AI721Contract *AI721ContractCaller) GetAgentFee(opts *bind.CallOpts, _agentId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "getAgentFee", _agentId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAgentFee is a free data retrieval call binding the contract method 0xed96f433.
//
// Solidity: function getAgentFee(uint256 _agentId) view returns(uint256)
func (_AI721Contract *AI721ContractSession) GetAgentFee(_agentId *big.Int) (*big.Int, error) {
	return _AI721Contract.Contract.GetAgentFee(&_AI721Contract.CallOpts, _agentId)
}

// GetAgentFee is a free data retrieval call binding the contract method 0xed96f433.
//
// Solidity: function getAgentFee(uint256 _agentId) view returns(uint256)
func (_AI721Contract *AI721ContractCallerSession) GetAgentFee(_agentId *big.Int) (*big.Int, error) {
	return _AI721Contract.Contract.GetAgentFee(&_AI721Contract.CallOpts, _agentId)
}

// GetAgentIdByOwner is a free data retrieval call binding the contract method 0xae57a2d3.
//
// Solidity: function getAgentIdByOwner(address _owner) view returns(uint256[])
func (_AI721Contract *AI721ContractCaller) GetAgentIdByOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "getAgentIdByOwner", _owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAgentIdByOwner is a free data retrieval call binding the contract method 0xae57a2d3.
//
// Solidity: function getAgentIdByOwner(address _owner) view returns(uint256[])
func (_AI721Contract *AI721ContractSession) GetAgentIdByOwner(_owner common.Address) ([]*big.Int, error) {
	return _AI721Contract.Contract.GetAgentIdByOwner(&_AI721Contract.CallOpts, _owner)
}

// GetAgentIdByOwner is a free data retrieval call binding the contract method 0xae57a2d3.
//
// Solidity: function getAgentIdByOwner(address _owner) view returns(uint256[])
func (_AI721Contract *AI721ContractCallerSession) GetAgentIdByOwner(_owner common.Address) ([]*big.Int, error) {
	return _AI721Contract.Contract.GetAgentIdByOwner(&_AI721Contract.CallOpts, _owner)
}

// GetAgentSystemPrompt is a free data retrieval call binding the contract method 0xf325f3f1.
//
// Solidity: function getAgentSystemPrompt(uint256 _agentId) view returns(bytes[])
func (_AI721Contract *AI721ContractCaller) GetAgentSystemPrompt(opts *bind.CallOpts, _agentId *big.Int) ([][]byte, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "getAgentSystemPrompt", _agentId)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetAgentSystemPrompt is a free data retrieval call binding the contract method 0xf325f3f1.
//
// Solidity: function getAgentSystemPrompt(uint256 _agentId) view returns(bytes[])
func (_AI721Contract *AI721ContractSession) GetAgentSystemPrompt(_agentId *big.Int) ([][]byte, error) {
	return _AI721Contract.Contract.GetAgentSystemPrompt(&_AI721Contract.CallOpts, _agentId)
}

// GetAgentSystemPrompt is a free data retrieval call binding the contract method 0xf325f3f1.
//
// Solidity: function getAgentSystemPrompt(uint256 _agentId) view returns(bytes[])
func (_AI721Contract *AI721ContractCallerSession) GetAgentSystemPrompt(_agentId *big.Int) ([][]byte, error) {
	return _AI721Contract.Contract.GetAgentSystemPrompt(&_AI721Contract.CallOpts, _agentId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_AI721Contract *AI721ContractCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_AI721Contract *AI721ContractSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _AI721Contract.Contract.GetApproved(&_AI721Contract.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_AI721Contract *AI721ContractCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _AI721Contract.Contract.GetApproved(&_AI721Contract.CallOpts, tokenId)
}

// GetMissionIdsByAgentId is a free data retrieval call binding the contract method 0x96694ad0.
//
// Solidity: function getMissionIdsByAgentId(uint256 _agentId) view returns(bytes[])
func (_AI721Contract *AI721ContractCaller) GetMissionIdsByAgentId(opts *bind.CallOpts, _agentId *big.Int) ([][]byte, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "getMissionIdsByAgentId", _agentId)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetMissionIdsByAgentId is a free data retrieval call binding the contract method 0x96694ad0.
//
// Solidity: function getMissionIdsByAgentId(uint256 _agentId) view returns(bytes[])
func (_AI721Contract *AI721ContractSession) GetMissionIdsByAgentId(_agentId *big.Int) ([][]byte, error) {
	return _AI721Contract.Contract.GetMissionIdsByAgentId(&_AI721Contract.CallOpts, _agentId)
}

// GetMissionIdsByAgentId is a free data retrieval call binding the contract method 0x96694ad0.
//
// Solidity: function getMissionIdsByAgentId(uint256 _agentId) view returns(bytes[])
func (_AI721Contract *AI721ContractCallerSession) GetMissionIdsByAgentId(_agentId *big.Int) ([][]byte, error) {
	return _AI721Contract.Contract.GetMissionIdsByAgentId(&_AI721Contract.CallOpts, _agentId)
}

// HybridModel is a free data retrieval call binding the contract method 0x5eb2364c.
//
// Solidity: function hybridModel() view returns(address)
func (_AI721Contract *AI721ContractCaller) HybridModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "hybridModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HybridModel is a free data retrieval call binding the contract method 0x5eb2364c.
//
// Solidity: function hybridModel() view returns(address)
func (_AI721Contract *AI721ContractSession) HybridModel() (common.Address, error) {
	return _AI721Contract.Contract.HybridModel(&_AI721Contract.CallOpts)
}

// HybridModel is a free data retrieval call binding the contract method 0x5eb2364c.
//
// Solidity: function hybridModel() view returns(address)
func (_AI721Contract *AI721ContractCallerSession) HybridModel() (common.Address, error) {
	return _AI721Contract.Contract.HybridModel(&_AI721Contract.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_AI721Contract *AI721ContractCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_AI721Contract *AI721ContractSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _AI721Contract.Contract.IsApprovedForAll(&_AI721Contract.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_AI721Contract *AI721ContractCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _AI721Contract.Contract.IsApprovedForAll(&_AI721Contract.CallOpts, owner, operator)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_AI721Contract *AI721ContractCaller) IsManager(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "isManager", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_AI721Contract *AI721ContractSession) IsManager(arg0 common.Address) (bool, error) {
	return _AI721Contract.Contract.IsManager(&_AI721Contract.CallOpts, arg0)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_AI721Contract *AI721ContractCallerSession) IsManager(arg0 common.Address) (bool, error) {
	return _AI721Contract.Contract.IsManager(&_AI721Contract.CallOpts, arg0)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_AI721Contract *AI721ContractCaller) MintPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "mintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_AI721Contract *AI721ContractSession) MintPrice() (*big.Int, error) {
	return _AI721Contract.Contract.MintPrice(&_AI721Contract.CallOpts)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_AI721Contract *AI721ContractCallerSession) MintPrice() (*big.Int, error) {
	return _AI721Contract.Contract.MintPrice(&_AI721Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AI721Contract *AI721ContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AI721Contract *AI721ContractSession) Name() (string, error) {
	return _AI721Contract.Contract.Name(&_AI721Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AI721Contract *AI721ContractCallerSession) Name() (string, error) {
	return _AI721Contract.Contract.Name(&_AI721Contract.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_AI721Contract *AI721ContractCaller) NextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "nextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_AI721Contract *AI721ContractSession) NextTokenId() (*big.Int, error) {
	return _AI721Contract.Contract.NextTokenId(&_AI721Contract.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_AI721Contract *AI721ContractCallerSession) NextTokenId() (*big.Int, error) {
	return _AI721Contract.Contract.NextTokenId(&_AI721Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AI721Contract *AI721ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AI721Contract *AI721ContractSession) Owner() (common.Address, error) {
	return _AI721Contract.Contract.Owner(&_AI721Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AI721Contract *AI721ContractCallerSession) Owner() (common.Address, error) {
	return _AI721Contract.Contract.Owner(&_AI721Contract.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_AI721Contract *AI721ContractCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_AI721Contract *AI721ContractSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _AI721Contract.Contract.OwnerOf(&_AI721Contract.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_AI721Contract *AI721ContractCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _AI721Contract.Contract.OwnerOf(&_AI721Contract.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AI721Contract *AI721ContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AI721Contract *AI721ContractSession) Paused() (bool, error) {
	return _AI721Contract.Contract.Paused(&_AI721Contract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AI721Contract *AI721ContractCallerSession) Paused() (bool, error) {
	return _AI721Contract.Contract.Paused(&_AI721Contract.CallOpts)
}

// PoolBalance is a free data retrieval call binding the contract method 0x6a6d964e.
//
// Solidity: function poolBalance(uint256 nftId) view returns(uint256)
func (_AI721Contract *AI721ContractCaller) PoolBalance(opts *bind.CallOpts, nftId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "poolBalance", nftId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolBalance is a free data retrieval call binding the contract method 0x6a6d964e.
//
// Solidity: function poolBalance(uint256 nftId) view returns(uint256)
func (_AI721Contract *AI721ContractSession) PoolBalance(nftId *big.Int) (*big.Int, error) {
	return _AI721Contract.Contract.PoolBalance(&_AI721Contract.CallOpts, nftId)
}

// PoolBalance is a free data retrieval call binding the contract method 0x6a6d964e.
//
// Solidity: function poolBalance(uint256 nftId) view returns(uint256)
func (_AI721Contract *AI721ContractCallerSession) PoolBalance(nftId *big.Int) (*big.Int, error) {
	return _AI721Contract.Contract.PoolBalance(&_AI721Contract.CallOpts, nftId)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _agentId, uint256 _salePrice) view returns(address, uint256)
func (_AI721Contract *AI721ContractCaller) RoyaltyInfo(opts *bind.CallOpts, _agentId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "royaltyInfo", _agentId, _salePrice)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _agentId, uint256 _salePrice) view returns(address, uint256)
func (_AI721Contract *AI721ContractSession) RoyaltyInfo(_agentId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _AI721Contract.Contract.RoyaltyInfo(&_AI721Contract.CallOpts, _agentId, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _agentId, uint256 _salePrice) view returns(address, uint256)
func (_AI721Contract *AI721ContractCallerSession) RoyaltyInfo(_agentId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _AI721Contract.Contract.RoyaltyInfo(&_AI721Contract.CallOpts, _agentId, _salePrice)
}

// RoyaltyPortion is a free data retrieval call binding the contract method 0x11d7beb2.
//
// Solidity: function royaltyPortion() view returns(uint16)
func (_AI721Contract *AI721ContractCaller) RoyaltyPortion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "royaltyPortion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RoyaltyPortion is a free data retrieval call binding the contract method 0x11d7beb2.
//
// Solidity: function royaltyPortion() view returns(uint16)
func (_AI721Contract *AI721ContractSession) RoyaltyPortion() (uint16, error) {
	return _AI721Contract.Contract.RoyaltyPortion(&_AI721Contract.CallOpts)
}

// RoyaltyPortion is a free data retrieval call binding the contract method 0x11d7beb2.
//
// Solidity: function royaltyPortion() view returns(uint16)
func (_AI721Contract *AI721ContractCallerSession) RoyaltyPortion() (uint16, error) {
	return _AI721Contract.Contract.RoyaltyPortion(&_AI721Contract.CallOpts)
}

// RoyaltyReceiver is a free data retrieval call binding the contract method 0x9fbc8713.
//
// Solidity: function royaltyReceiver() view returns(address)
func (_AI721Contract *AI721ContractCaller) RoyaltyReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "royaltyReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoyaltyReceiver is a free data retrieval call binding the contract method 0x9fbc8713.
//
// Solidity: function royaltyReceiver() view returns(address)
func (_AI721Contract *AI721ContractSession) RoyaltyReceiver() (common.Address, error) {
	return _AI721Contract.Contract.RoyaltyReceiver(&_AI721Contract.CallOpts)
}

// RoyaltyReceiver is a free data retrieval call binding the contract method 0x9fbc8713.
//
// Solidity: function royaltyReceiver() view returns(address)
func (_AI721Contract *AI721ContractCallerSession) RoyaltyReceiver() (common.Address, error) {
	return _AI721Contract.Contract.RoyaltyReceiver(&_AI721Contract.CallOpts)
}

// SignaturesUsed is a free data retrieval call binding the contract method 0x757d513b.
//
// Solidity: function signaturesUsed(address nftOwner, bytes signature) view returns(bool)
func (_AI721Contract *AI721ContractCaller) SignaturesUsed(opts *bind.CallOpts, nftOwner common.Address, signature []byte) (bool, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "signaturesUsed", nftOwner, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SignaturesUsed is a free data retrieval call binding the contract method 0x757d513b.
//
// Solidity: function signaturesUsed(address nftOwner, bytes signature) view returns(bool)
func (_AI721Contract *AI721ContractSession) SignaturesUsed(nftOwner common.Address, signature []byte) (bool, error) {
	return _AI721Contract.Contract.SignaturesUsed(&_AI721Contract.CallOpts, nftOwner, signature)
}

// SignaturesUsed is a free data retrieval call binding the contract method 0x757d513b.
//
// Solidity: function signaturesUsed(address nftOwner, bytes signature) view returns(bool)
func (_AI721Contract *AI721ContractCallerSession) SignaturesUsed(nftOwner common.Address, signature []byte) (bool, error) {
	return _AI721Contract.Contract.SignaturesUsed(&_AI721Contract.CallOpts, nftOwner, signature)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_AI721Contract *AI721ContractCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_AI721Contract *AI721ContractSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _AI721Contract.Contract.SupportsInterface(&_AI721Contract.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_AI721Contract *AI721ContractCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _AI721Contract.Contract.SupportsInterface(&_AI721Contract.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AI721Contract *AI721ContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AI721Contract *AI721ContractSession) Symbol() (string, error) {
	return _AI721Contract.Contract.Symbol(&_AI721Contract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AI721Contract *AI721ContractCallerSession) Symbol() (string, error) {
	return _AI721Contract.Contract.Symbol(&_AI721Contract.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_AI721Contract *AI721ContractCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_AI721Contract *AI721ContractSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _AI721Contract.Contract.TokenByIndex(&_AI721Contract.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_AI721Contract *AI721ContractCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _AI721Contract.Contract.TokenByIndex(&_AI721Contract.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_AI721Contract *AI721ContractCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_AI721Contract *AI721ContractSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _AI721Contract.Contract.TokenOfOwnerByIndex(&_AI721Contract.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_AI721Contract *AI721ContractCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _AI721Contract.Contract.TokenOfOwnerByIndex(&_AI721Contract.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _agentId) view returns(string)
func (_AI721Contract *AI721ContractCaller) TokenURI(opts *bind.CallOpts, _agentId *big.Int) (string, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "tokenURI", _agentId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _agentId) view returns(string)
func (_AI721Contract *AI721ContractSession) TokenURI(_agentId *big.Int) (string, error) {
	return _AI721Contract.Contract.TokenURI(&_AI721Contract.CallOpts, _agentId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _agentId) view returns(string)
func (_AI721Contract *AI721ContractCallerSession) TokenURI(_agentId *big.Int) (string, error) {
	return _AI721Contract.Contract.TokenURI(&_AI721Contract.CallOpts, _agentId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AI721Contract *AI721ContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AI721Contract *AI721ContractSession) TotalSupply() (*big.Int, error) {
	return _AI721Contract.Contract.TotalSupply(&_AI721Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AI721Contract *AI721ContractCallerSession) TotalSupply() (*big.Int, error) {
	return _AI721Contract.Contract.TotalSupply(&_AI721Contract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_AI721Contract *AI721ContractCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_AI721Contract *AI721ContractSession) Version() (string, error) {
	return _AI721Contract.Contract.Version(&_AI721Contract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_AI721Contract *AI721ContractCallerSession) Version() (string, error) {
	return _AI721Contract.Contract.Version(&_AI721Contract.CallOpts)
}

// WorkerHub is a free data retrieval call binding the contract method 0x860e9dc6.
//
// Solidity: function workerHub() view returns(address)
func (_AI721Contract *AI721ContractCaller) WorkerHub(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AI721Contract.contract.Call(opts, &out, "workerHub")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WorkerHub is a free data retrieval call binding the contract method 0x860e9dc6.
//
// Solidity: function workerHub() view returns(address)
func (_AI721Contract *AI721ContractSession) WorkerHub() (common.Address, error) {
	return _AI721Contract.Contract.WorkerHub(&_AI721Contract.CallOpts)
}

// WorkerHub is a free data retrieval call binding the contract method 0x860e9dc6.
//
// Solidity: function workerHub() view returns(address)
func (_AI721Contract *AI721ContractCallerSession) WorkerHub() (common.Address, error) {
	return _AI721Contract.Contract.WorkerHub(&_AI721Contract.CallOpts)
}

// AddNewAgentData is a paid mutator transaction binding the contract method 0x9f10cc66.
//
// Solidity: function addNewAgentData(uint256 _agentId, bytes _sysPrompt) returns()
func (_AI721Contract *AI721ContractTransactor) AddNewAgentData(opts *bind.TransactOpts, _agentId *big.Int, _sysPrompt []byte) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "addNewAgentData", _agentId, _sysPrompt)
}

// AddNewAgentData is a paid mutator transaction binding the contract method 0x9f10cc66.
//
// Solidity: function addNewAgentData(uint256 _agentId, bytes _sysPrompt) returns()
func (_AI721Contract *AI721ContractSession) AddNewAgentData(_agentId *big.Int, _sysPrompt []byte) (*types.Transaction, error) {
	return _AI721Contract.Contract.AddNewAgentData(&_AI721Contract.TransactOpts, _agentId, _sysPrompt)
}

// AddNewAgentData is a paid mutator transaction binding the contract method 0x9f10cc66.
//
// Solidity: function addNewAgentData(uint256 _agentId, bytes _sysPrompt) returns()
func (_AI721Contract *AI721ContractTransactorSession) AddNewAgentData(_agentId *big.Int, _sysPrompt []byte) (*types.Transaction, error) {
	return _AI721Contract.Contract.AddNewAgentData(&_AI721Contract.TransactOpts, _agentId, _sysPrompt)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_AI721Contract *AI721ContractTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_AI721Contract *AI721ContractSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AI721Contract.Contract.Approve(&_AI721Contract.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_AI721Contract *AI721ContractTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AI721Contract.Contract.Approve(&_AI721Contract.TransactOpts, to, tokenId)
}

// AuthorizeManager is a paid mutator transaction binding the contract method 0x267c8507.
//
// Solidity: function authorizeManager(address _account) returns()
func (_AI721Contract *AI721ContractTransactor) AuthorizeManager(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "authorizeManager", _account)
}

// AuthorizeManager is a paid mutator transaction binding the contract method 0x267c8507.
//
// Solidity: function authorizeManager(address _account) returns()
func (_AI721Contract *AI721ContractSession) AuthorizeManager(_account common.Address) (*types.Transaction, error) {
	return _AI721Contract.Contract.AuthorizeManager(&_AI721Contract.TransactOpts, _account)
}

// AuthorizeManager is a paid mutator transaction binding the contract method 0x267c8507.
//
// Solidity: function authorizeManager(address _account) returns()
func (_AI721Contract *AI721ContractTransactorSession) AuthorizeManager(_account common.Address) (*types.Transaction, error) {
	return _AI721Contract.Contract.AuthorizeManager(&_AI721Contract.TransactOpts, _account)
}

// ClaimFee is a paid mutator transaction binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() returns()
func (_AI721Contract *AI721ContractTransactor) ClaimFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "claimFee")
}

// ClaimFee is a paid mutator transaction binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() returns()
func (_AI721Contract *AI721ContractSession) ClaimFee() (*types.Transaction, error) {
	return _AI721Contract.Contract.ClaimFee(&_AI721Contract.TransactOpts)
}

// ClaimFee is a paid mutator transaction binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() returns()
func (_AI721Contract *AI721ContractTransactorSession) ClaimFee() (*types.Transaction, error) {
	return _AI721Contract.Contract.ClaimFee(&_AI721Contract.TransactOpts)
}

// CreateMission is a paid mutator transaction binding the contract method 0x6001ad44.
//
// Solidity: function createMission(uint256 _agentId, bytes _missionData) returns()
func (_AI721Contract *AI721ContractTransactor) CreateMission(opts *bind.TransactOpts, _agentId *big.Int, _missionData []byte) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "createMission", _agentId, _missionData)
}

// CreateMission is a paid mutator transaction binding the contract method 0x6001ad44.
//
// Solidity: function createMission(uint256 _agentId, bytes _missionData) returns()
func (_AI721Contract *AI721ContractSession) CreateMission(_agentId *big.Int, _missionData []byte) (*types.Transaction, error) {
	return _AI721Contract.Contract.CreateMission(&_AI721Contract.TransactOpts, _agentId, _missionData)
}

// CreateMission is a paid mutator transaction binding the contract method 0x6001ad44.
//
// Solidity: function createMission(uint256 _agentId, bytes _missionData) returns()
func (_AI721Contract *AI721ContractTransactorSession) CreateMission(_agentId *big.Int, _missionData []byte) (*types.Transaction, error) {
	return _AI721Contract.Contract.CreateMission(&_AI721Contract.TransactOpts, _agentId, _missionData)
}

// DeauthorizeManager is a paid mutator transaction binding the contract method 0x0305ea01.
//
// Solidity: function deauthorizeManager(address _account) returns()
func (_AI721Contract *AI721ContractTransactor) DeauthorizeManager(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "deauthorizeManager", _account)
}

// DeauthorizeManager is a paid mutator transaction binding the contract method 0x0305ea01.
//
// Solidity: function deauthorizeManager(address _account) returns()
func (_AI721Contract *AI721ContractSession) DeauthorizeManager(_account common.Address) (*types.Transaction, error) {
	return _AI721Contract.Contract.DeauthorizeManager(&_AI721Contract.TransactOpts, _account)
}

// DeauthorizeManager is a paid mutator transaction binding the contract method 0x0305ea01.
//
// Solidity: function deauthorizeManager(address _account) returns()
func (_AI721Contract *AI721ContractTransactorSession) DeauthorizeManager(_account common.Address) (*types.Transaction, error) {
	return _AI721Contract.Contract.DeauthorizeManager(&_AI721Contract.TransactOpts, _account)
}

// Infer is a paid mutator transaction binding the contract method 0x3c00cf99.
//
// Solidity: function infer(uint256 _agentId, bytes _calldata, string _externalData, bool _flag) payable returns()
func (_AI721Contract *AI721ContractTransactor) Infer(opts *bind.TransactOpts, _agentId *big.Int, _calldata []byte, _externalData string, _flag bool) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "infer", _agentId, _calldata, _externalData, _flag)
}

// Infer is a paid mutator transaction binding the contract method 0x3c00cf99.
//
// Solidity: function infer(uint256 _agentId, bytes _calldata, string _externalData, bool _flag) payable returns()
func (_AI721Contract *AI721ContractSession) Infer(_agentId *big.Int, _calldata []byte, _externalData string, _flag bool) (*types.Transaction, error) {
	return _AI721Contract.Contract.Infer(&_AI721Contract.TransactOpts, _agentId, _calldata, _externalData, _flag)
}

// Infer is a paid mutator transaction binding the contract method 0x3c00cf99.
//
// Solidity: function infer(uint256 _agentId, bytes _calldata, string _externalData, bool _flag) payable returns()
func (_AI721Contract *AI721ContractTransactorSession) Infer(_agentId *big.Int, _calldata []byte, _externalData string, _flag bool) (*types.Transaction, error) {
	return _AI721Contract.Contract.Infer(&_AI721Contract.TransactOpts, _agentId, _calldata, _externalData, _flag)
}

// Infer0 is a paid mutator transaction binding the contract method 0x566a9951.
//
// Solidity: function infer(uint256 _agentId, bytes _calldata, string _externalData) payable returns()
func (_AI721Contract *AI721ContractTransactor) Infer0(opts *bind.TransactOpts, _agentId *big.Int, _calldata []byte, _externalData string) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "infer0", _agentId, _calldata, _externalData)
}

// Infer0 is a paid mutator transaction binding the contract method 0x566a9951.
//
// Solidity: function infer(uint256 _agentId, bytes _calldata, string _externalData) payable returns()
func (_AI721Contract *AI721ContractSession) Infer0(_agentId *big.Int, _calldata []byte, _externalData string) (*types.Transaction, error) {
	return _AI721Contract.Contract.Infer0(&_AI721Contract.TransactOpts, _agentId, _calldata, _externalData)
}

// Infer0 is a paid mutator transaction binding the contract method 0x566a9951.
//
// Solidity: function infer(uint256 _agentId, bytes _calldata, string _externalData) payable returns()
func (_AI721Contract *AI721ContractTransactorSession) Infer0(_agentId *big.Int, _calldata []byte, _externalData string) (*types.Transaction, error) {
	return _AI721Contract.Contract.Infer0(&_AI721Contract.TransactOpts, _agentId, _calldata, _externalData)
}

// Initialize is a paid mutator transaction binding the contract method 0x1e18fb8c.
//
// Solidity: function initialize(string _name, string _symbol, uint256 _mintPrice, address _royaltyReceiver, uint16 _royaltyPortion, uint256 _nextTokenId, address _hybridModel, address _workerHub) returns()
func (_AI721Contract *AI721ContractTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _mintPrice *big.Int, _royaltyReceiver common.Address, _royaltyPortion uint16, _nextTokenId *big.Int, _hybridModel common.Address, _workerHub common.Address) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "initialize", _name, _symbol, _mintPrice, _royaltyReceiver, _royaltyPortion, _nextTokenId, _hybridModel, _workerHub)
}

// Initialize is a paid mutator transaction binding the contract method 0x1e18fb8c.
//
// Solidity: function initialize(string _name, string _symbol, uint256 _mintPrice, address _royaltyReceiver, uint16 _royaltyPortion, uint256 _nextTokenId, address _hybridModel, address _workerHub) returns()
func (_AI721Contract *AI721ContractSession) Initialize(_name string, _symbol string, _mintPrice *big.Int, _royaltyReceiver common.Address, _royaltyPortion uint16, _nextTokenId *big.Int, _hybridModel common.Address, _workerHub common.Address) (*types.Transaction, error) {
	return _AI721Contract.Contract.Initialize(&_AI721Contract.TransactOpts, _name, _symbol, _mintPrice, _royaltyReceiver, _royaltyPortion, _nextTokenId, _hybridModel, _workerHub)
}

// Initialize is a paid mutator transaction binding the contract method 0x1e18fb8c.
//
// Solidity: function initialize(string _name, string _symbol, uint256 _mintPrice, address _royaltyReceiver, uint16 _royaltyPortion, uint256 _nextTokenId, address _hybridModel, address _workerHub) returns()
func (_AI721Contract *AI721ContractTransactorSession) Initialize(_name string, _symbol string, _mintPrice *big.Int, _royaltyReceiver common.Address, _royaltyPortion uint16, _nextTokenId *big.Int, _hybridModel common.Address, _workerHub common.Address) (*types.Transaction, error) {
	return _AI721Contract.Contract.Initialize(&_AI721Contract.TransactOpts, _name, _symbol, _mintPrice, _royaltyReceiver, _royaltyPortion, _nextTokenId, _hybridModel, _workerHub)
}

// Mint is a paid mutator transaction binding the contract method 0xcc216aca.
//
// Solidity: function mint(address _to, string _uri, bytes _data, uint256 _fee) payable returns(uint256)
func (_AI721Contract *AI721ContractTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _uri string, _data []byte, _fee *big.Int) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "mint", _to, _uri, _data, _fee)
}

// Mint is a paid mutator transaction binding the contract method 0xcc216aca.
//
// Solidity: function mint(address _to, string _uri, bytes _data, uint256 _fee) payable returns(uint256)
func (_AI721Contract *AI721ContractSession) Mint(_to common.Address, _uri string, _data []byte, _fee *big.Int) (*types.Transaction, error) {
	return _AI721Contract.Contract.Mint(&_AI721Contract.TransactOpts, _to, _uri, _data, _fee)
}

// Mint is a paid mutator transaction binding the contract method 0xcc216aca.
//
// Solidity: function mint(address _to, string _uri, bytes _data, uint256 _fee) payable returns(uint256)
func (_AI721Contract *AI721ContractTransactorSession) Mint(_to common.Address, _uri string, _data []byte, _fee *big.Int) (*types.Transaction, error) {
	return _AI721Contract.Contract.Mint(&_AI721Contract.TransactOpts, _to, _uri, _data, _fee)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AI721Contract *AI721ContractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AI721Contract *AI721ContractSession) Pause() (*types.Transaction, error) {
	return _AI721Contract.Contract.Pause(&_AI721Contract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AI721Contract *AI721ContractTransactorSession) Pause() (*types.Transaction, error) {
	return _AI721Contract.Contract.Pause(&_AI721Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AI721Contract *AI721ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AI721Contract *AI721ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _AI721Contract.Contract.RenounceOwnership(&_AI721Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AI721Contract *AI721ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AI721Contract.Contract.RenounceOwnership(&_AI721Contract.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_AI721Contract *AI721ContractTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_AI721Contract *AI721ContractSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AI721Contract.Contract.SafeTransferFrom(&_AI721Contract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_AI721Contract *AI721ContractTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AI721Contract.Contract.SafeTransferFrom(&_AI721Contract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_AI721Contract *AI721ContractTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_AI721Contract *AI721ContractSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _AI721Contract.Contract.SafeTransferFrom0(&_AI721Contract.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_AI721Contract *AI721ContractTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _AI721Contract.Contract.SafeTransferFrom0(&_AI721Contract.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_AI721Contract *AI721ContractTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_AI721Contract *AI721ContractSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _AI721Contract.Contract.SetApprovalForAll(&_AI721Contract.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_AI721Contract *AI721ContractTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _AI721Contract.Contract.SetApprovalForAll(&_AI721Contract.TransactOpts, operator, approved)
}

// TopUpPoolBalance is a paid mutator transaction binding the contract method 0xc82cd94b.
//
// Solidity: function topUpPoolBalance(uint256 _agentId) payable returns()
func (_AI721Contract *AI721ContractTransactor) TopUpPoolBalance(opts *bind.TransactOpts, _agentId *big.Int) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "topUpPoolBalance", _agentId)
}

// TopUpPoolBalance is a paid mutator transaction binding the contract method 0xc82cd94b.
//
// Solidity: function topUpPoolBalance(uint256 _agentId) payable returns()
func (_AI721Contract *AI721ContractSession) TopUpPoolBalance(_agentId *big.Int) (*types.Transaction, error) {
	return _AI721Contract.Contract.TopUpPoolBalance(&_AI721Contract.TransactOpts, _agentId)
}

// TopUpPoolBalance is a paid mutator transaction binding the contract method 0xc82cd94b.
//
// Solidity: function topUpPoolBalance(uint256 _agentId) payable returns()
func (_AI721Contract *AI721ContractTransactorSession) TopUpPoolBalance(_agentId *big.Int) (*types.Transaction, error) {
	return _AI721Contract.Contract.TopUpPoolBalance(&_AI721Contract.TransactOpts, _agentId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_AI721Contract *AI721ContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_AI721Contract *AI721ContractSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AI721Contract.Contract.TransferFrom(&_AI721Contract.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_AI721Contract *AI721ContractTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _AI721Contract.Contract.TransferFrom(&_AI721Contract.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AI721Contract *AI721ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AI721Contract *AI721ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AI721Contract.Contract.TransferOwnership(&_AI721Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AI721Contract *AI721ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AI721Contract.Contract.TransferOwnership(&_AI721Contract.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AI721Contract *AI721ContractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AI721Contract *AI721ContractSession) Unpause() (*types.Transaction, error) {
	return _AI721Contract.Contract.Unpause(&_AI721Contract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AI721Contract *AI721ContractTransactorSession) Unpause() (*types.Transaction, error) {
	return _AI721Contract.Contract.Unpause(&_AI721Contract.TransactOpts)
}

// UpdateAgentData is a paid mutator transaction binding the contract method 0xed82c9e0.
//
// Solidity: function updateAgentData(uint256 _agentId, bytes _sysPrompt, uint256 _promptIdx) returns()
func (_AI721Contract *AI721ContractTransactor) UpdateAgentData(opts *bind.TransactOpts, _agentId *big.Int, _sysPrompt []byte, _promptIdx *big.Int) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "updateAgentData", _agentId, _sysPrompt, _promptIdx)
}

// UpdateAgentData is a paid mutator transaction binding the contract method 0xed82c9e0.
//
// Solidity: function updateAgentData(uint256 _agentId, bytes _sysPrompt, uint256 _promptIdx) returns()
func (_AI721Contract *AI721ContractSession) UpdateAgentData(_agentId *big.Int, _sysPrompt []byte, _promptIdx *big.Int) (*types.Transaction, error) {
	return _AI721Contract.Contract.UpdateAgentData(&_AI721Contract.TransactOpts, _agentId, _sysPrompt, _promptIdx)
}

// UpdateAgentData is a paid mutator transaction binding the contract method 0xed82c9e0.
//
// Solidity: function updateAgentData(uint256 _agentId, bytes _sysPrompt, uint256 _promptIdx) returns()
func (_AI721Contract *AI721ContractTransactorSession) UpdateAgentData(_agentId *big.Int, _sysPrompt []byte, _promptIdx *big.Int) (*types.Transaction, error) {
	return _AI721Contract.Contract.UpdateAgentData(&_AI721Contract.TransactOpts, _agentId, _sysPrompt, _promptIdx)
}

// UpdateAgentDataWithSignature is a paid mutator transaction binding the contract method 0xb8a49a57.
//
// Solidity: function updateAgentDataWithSignature(uint256 _agentId, bytes _sysPrompt, uint256 _promptIdx, uint256 _randomNonce, bytes _signature) returns()
func (_AI721Contract *AI721ContractTransactor) UpdateAgentDataWithSignature(opts *bind.TransactOpts, _agentId *big.Int, _sysPrompt []byte, _promptIdx *big.Int, _randomNonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "updateAgentDataWithSignature", _agentId, _sysPrompt, _promptIdx, _randomNonce, _signature)
}

// UpdateAgentDataWithSignature is a paid mutator transaction binding the contract method 0xb8a49a57.
//
// Solidity: function updateAgentDataWithSignature(uint256 _agentId, bytes _sysPrompt, uint256 _promptIdx, uint256 _randomNonce, bytes _signature) returns()
func (_AI721Contract *AI721ContractSession) UpdateAgentDataWithSignature(_agentId *big.Int, _sysPrompt []byte, _promptIdx *big.Int, _randomNonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AI721Contract.Contract.UpdateAgentDataWithSignature(&_AI721Contract.TransactOpts, _agentId, _sysPrompt, _promptIdx, _randomNonce, _signature)
}

// UpdateAgentDataWithSignature is a paid mutator transaction binding the contract method 0xb8a49a57.
//
// Solidity: function updateAgentDataWithSignature(uint256 _agentId, bytes _sysPrompt, uint256 _promptIdx, uint256 _randomNonce, bytes _signature) returns()
func (_AI721Contract *AI721ContractTransactorSession) UpdateAgentDataWithSignature(_agentId *big.Int, _sysPrompt []byte, _promptIdx *big.Int, _randomNonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AI721Contract.Contract.UpdateAgentDataWithSignature(&_AI721Contract.TransactOpts, _agentId, _sysPrompt, _promptIdx, _randomNonce, _signature)
}

// UpdateAgentFee is a paid mutator transaction binding the contract method 0xb1fd1526.
//
// Solidity: function updateAgentFee(uint256 _agentId, uint256 _fee) returns()
func (_AI721Contract *AI721ContractTransactor) UpdateAgentFee(opts *bind.TransactOpts, _agentId *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "updateAgentFee", _agentId, _fee)
}

// UpdateAgentFee is a paid mutator transaction binding the contract method 0xb1fd1526.
//
// Solidity: function updateAgentFee(uint256 _agentId, uint256 _fee) returns()
func (_AI721Contract *AI721ContractSession) UpdateAgentFee(_agentId *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _AI721Contract.Contract.UpdateAgentFee(&_AI721Contract.TransactOpts, _agentId, _fee)
}

// UpdateAgentFee is a paid mutator transaction binding the contract method 0xb1fd1526.
//
// Solidity: function updateAgentFee(uint256 _agentId, uint256 _fee) returns()
func (_AI721Contract *AI721ContractTransactorSession) UpdateAgentFee(_agentId *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _AI721Contract.Contract.UpdateAgentFee(&_AI721Contract.TransactOpts, _agentId, _fee)
}

// UpdateAgentURI is a paid mutator transaction binding the contract method 0x6b595822.
//
// Solidity: function updateAgentURI(uint256 _agentId, string _uri) returns()
func (_AI721Contract *AI721ContractTransactor) UpdateAgentURI(opts *bind.TransactOpts, _agentId *big.Int, _uri string) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "updateAgentURI", _agentId, _uri)
}

// UpdateAgentURI is a paid mutator transaction binding the contract method 0x6b595822.
//
// Solidity: function updateAgentURI(uint256 _agentId, string _uri) returns()
func (_AI721Contract *AI721ContractSession) UpdateAgentURI(_agentId *big.Int, _uri string) (*types.Transaction, error) {
	return _AI721Contract.Contract.UpdateAgentURI(&_AI721Contract.TransactOpts, _agentId, _uri)
}

// UpdateAgentURI is a paid mutator transaction binding the contract method 0x6b595822.
//
// Solidity: function updateAgentURI(uint256 _agentId, string _uri) returns()
func (_AI721Contract *AI721ContractTransactorSession) UpdateAgentURI(_agentId *big.Int, _uri string) (*types.Transaction, error) {
	return _AI721Contract.Contract.UpdateAgentURI(&_AI721Contract.TransactOpts, _agentId, _uri)
}

// UpdateAgentUriWithSignature is a paid mutator transaction binding the contract method 0xf5888779.
//
// Solidity: function updateAgentUriWithSignature(uint256 _agentId, string _uri, uint256 _randomNonce, bytes _signature) returns()
func (_AI721Contract *AI721ContractTransactor) UpdateAgentUriWithSignature(opts *bind.TransactOpts, _agentId *big.Int, _uri string, _randomNonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "updateAgentUriWithSignature", _agentId, _uri, _randomNonce, _signature)
}

// UpdateAgentUriWithSignature is a paid mutator transaction binding the contract method 0xf5888779.
//
// Solidity: function updateAgentUriWithSignature(uint256 _agentId, string _uri, uint256 _randomNonce, bytes _signature) returns()
func (_AI721Contract *AI721ContractSession) UpdateAgentUriWithSignature(_agentId *big.Int, _uri string, _randomNonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AI721Contract.Contract.UpdateAgentUriWithSignature(&_AI721Contract.TransactOpts, _agentId, _uri, _randomNonce, _signature)
}

// UpdateAgentUriWithSignature is a paid mutator transaction binding the contract method 0xf5888779.
//
// Solidity: function updateAgentUriWithSignature(uint256 _agentId, string _uri, uint256 _randomNonce, bytes _signature) returns()
func (_AI721Contract *AI721ContractTransactorSession) UpdateAgentUriWithSignature(_agentId *big.Int, _uri string, _randomNonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _AI721Contract.Contract.UpdateAgentUriWithSignature(&_AI721Contract.TransactOpts, _agentId, _uri, _randomNonce, _signature)
}

// UpdateMintPrice is a paid mutator transaction binding the contract method 0x00728e46.
//
// Solidity: function updateMintPrice(uint256 _mintPrice) returns()
func (_AI721Contract *AI721ContractTransactor) UpdateMintPrice(opts *bind.TransactOpts, _mintPrice *big.Int) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "updateMintPrice", _mintPrice)
}

// UpdateMintPrice is a paid mutator transaction binding the contract method 0x00728e46.
//
// Solidity: function updateMintPrice(uint256 _mintPrice) returns()
func (_AI721Contract *AI721ContractSession) UpdateMintPrice(_mintPrice *big.Int) (*types.Transaction, error) {
	return _AI721Contract.Contract.UpdateMintPrice(&_AI721Contract.TransactOpts, _mintPrice)
}

// UpdateMintPrice is a paid mutator transaction binding the contract method 0x00728e46.
//
// Solidity: function updateMintPrice(uint256 _mintPrice) returns()
func (_AI721Contract *AI721ContractTransactorSession) UpdateMintPrice(_mintPrice *big.Int) (*types.Transaction, error) {
	return _AI721Contract.Contract.UpdateMintPrice(&_AI721Contract.TransactOpts, _mintPrice)
}

// UpdateMission is a paid mutator transaction binding the contract method 0x8f17098f.
//
// Solidity: function updateMission(uint256 _agentId, uint256 _missionIdx, bytes _missionData) returns()
func (_AI721Contract *AI721ContractTransactor) UpdateMission(opts *bind.TransactOpts, _agentId *big.Int, _missionIdx *big.Int, _missionData []byte) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "updateMission", _agentId, _missionIdx, _missionData)
}

// UpdateMission is a paid mutator transaction binding the contract method 0x8f17098f.
//
// Solidity: function updateMission(uint256 _agentId, uint256 _missionIdx, bytes _missionData) returns()
func (_AI721Contract *AI721ContractSession) UpdateMission(_agentId *big.Int, _missionIdx *big.Int, _missionData []byte) (*types.Transaction, error) {
	return _AI721Contract.Contract.UpdateMission(&_AI721Contract.TransactOpts, _agentId, _missionIdx, _missionData)
}

// UpdateMission is a paid mutator transaction binding the contract method 0x8f17098f.
//
// Solidity: function updateMission(uint256 _agentId, uint256 _missionIdx, bytes _missionData) returns()
func (_AI721Contract *AI721ContractTransactorSession) UpdateMission(_agentId *big.Int, _missionIdx *big.Int, _missionData []byte) (*types.Transaction, error) {
	return _AI721Contract.Contract.UpdateMission(&_AI721Contract.TransactOpts, _agentId, _missionIdx, _missionData)
}

// UpdateRoyaltyPortion is a paid mutator transaction binding the contract method 0x19e93993.
//
// Solidity: function updateRoyaltyPortion(uint16 _royaltyPortion) returns()
func (_AI721Contract *AI721ContractTransactor) UpdateRoyaltyPortion(opts *bind.TransactOpts, _royaltyPortion uint16) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "updateRoyaltyPortion", _royaltyPortion)
}

// UpdateRoyaltyPortion is a paid mutator transaction binding the contract method 0x19e93993.
//
// Solidity: function updateRoyaltyPortion(uint16 _royaltyPortion) returns()
func (_AI721Contract *AI721ContractSession) UpdateRoyaltyPortion(_royaltyPortion uint16) (*types.Transaction, error) {
	return _AI721Contract.Contract.UpdateRoyaltyPortion(&_AI721Contract.TransactOpts, _royaltyPortion)
}

// UpdateRoyaltyPortion is a paid mutator transaction binding the contract method 0x19e93993.
//
// Solidity: function updateRoyaltyPortion(uint16 _royaltyPortion) returns()
func (_AI721Contract *AI721ContractTransactorSession) UpdateRoyaltyPortion(_royaltyPortion uint16) (*types.Transaction, error) {
	return _AI721Contract.Contract.UpdateRoyaltyPortion(&_AI721Contract.TransactOpts, _royaltyPortion)
}

// UpdateRoyaltyReceiver is a paid mutator transaction binding the contract method 0x29dc4d9b.
//
// Solidity: function updateRoyaltyReceiver(address _royaltyReceiver) returns()
func (_AI721Contract *AI721ContractTransactor) UpdateRoyaltyReceiver(opts *bind.TransactOpts, _royaltyReceiver common.Address) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "updateRoyaltyReceiver", _royaltyReceiver)
}

// UpdateRoyaltyReceiver is a paid mutator transaction binding the contract method 0x29dc4d9b.
//
// Solidity: function updateRoyaltyReceiver(address _royaltyReceiver) returns()
func (_AI721Contract *AI721ContractSession) UpdateRoyaltyReceiver(_royaltyReceiver common.Address) (*types.Transaction, error) {
	return _AI721Contract.Contract.UpdateRoyaltyReceiver(&_AI721Contract.TransactOpts, _royaltyReceiver)
}

// UpdateRoyaltyReceiver is a paid mutator transaction binding the contract method 0x29dc4d9b.
//
// Solidity: function updateRoyaltyReceiver(address _royaltyReceiver) returns()
func (_AI721Contract *AI721ContractTransactorSession) UpdateRoyaltyReceiver(_royaltyReceiver common.Address) (*types.Transaction, error) {
	return _AI721Contract.Contract.UpdateRoyaltyReceiver(&_AI721Contract.TransactOpts, _royaltyReceiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _to, uint256 _value) returns()
func (_AI721Contract *AI721ContractTransactor) Withdraw(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _AI721Contract.contract.Transact(opts, "withdraw", _to, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _to, uint256 _value) returns()
func (_AI721Contract *AI721ContractSession) Withdraw(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _AI721Contract.Contract.Withdraw(&_AI721Contract.TransactOpts, _to, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _to, uint256 _value) returns()
func (_AI721Contract *AI721ContractTransactorSession) Withdraw(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _AI721Contract.Contract.Withdraw(&_AI721Contract.TransactOpts, _to, _value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AI721Contract *AI721ContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AI721Contract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AI721Contract *AI721ContractSession) Receive() (*types.Transaction, error) {
	return _AI721Contract.Contract.Receive(&_AI721Contract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AI721Contract *AI721ContractTransactorSession) Receive() (*types.Transaction, error) {
	return _AI721Contract.Contract.Receive(&_AI721Contract.TransactOpts)
}

// AI721ContractAgentDataAddNewIterator is returned from FilterAgentDataAddNew and is used to iterate over the raw logs and unpacked data for AgentDataAddNew events raised by the AI721Contract contract.
type AI721ContractAgentDataAddNewIterator struct {
	Event *AI721ContractAgentDataAddNew // Event containing the contract specifics and raw log

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
func (it *AI721ContractAgentDataAddNewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AI721ContractAgentDataAddNew)
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
		it.Event = new(AI721ContractAgentDataAddNew)
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
func (it *AI721ContractAgentDataAddNewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AI721ContractAgentDataAddNewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AI721ContractAgentDataAddNew represents a AgentDataAddNew event raised by the AI721Contract contract.
type AI721ContractAgentDataAddNew struct {
	AgentId   *big.Int
	SysPrompt [][]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAgentDataAddNew is a free log retrieval operation binding the contract event 0xdebec4c58e3b7c5817893e50cb1f9e65b65978e8c89bb4407eb0109d5887b258.
//
// Solidity: event AgentDataAddNew(uint256 indexed agentId, bytes[] sysPrompt)
func (_AI721Contract *AI721ContractFilterer) FilterAgentDataAddNew(opts *bind.FilterOpts, agentId []*big.Int) (*AI721ContractAgentDataAddNewIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _AI721Contract.contract.FilterLogs(opts, "AgentDataAddNew", agentIdRule)
	if err != nil {
		return nil, err
	}
	return &AI721ContractAgentDataAddNewIterator{contract: _AI721Contract.contract, event: "AgentDataAddNew", logs: logs, sub: sub}, nil
}

// WatchAgentDataAddNew is a free log subscription operation binding the contract event 0xdebec4c58e3b7c5817893e50cb1f9e65b65978e8c89bb4407eb0109d5887b258.
//
// Solidity: event AgentDataAddNew(uint256 indexed agentId, bytes[] sysPrompt)
func (_AI721Contract *AI721ContractFilterer) WatchAgentDataAddNew(opts *bind.WatchOpts, sink chan<- *AI721ContractAgentDataAddNew, agentId []*big.Int) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _AI721Contract.contract.WatchLogs(opts, "AgentDataAddNew", agentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AI721ContractAgentDataAddNew)
				if err := _AI721Contract.contract.UnpackLog(event, "AgentDataAddNew", log); err != nil {
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
func (_AI721Contract *AI721ContractFilterer) ParseAgentDataAddNew(log types.Log) (*AI721ContractAgentDataAddNew, error) {
	event := new(AI721ContractAgentDataAddNew)
	if err := _AI721Contract.contract.UnpackLog(event, "AgentDataAddNew", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AI721ContractAgentDataUpdateIterator is returned from FilterAgentDataUpdate and is used to iterate over the raw logs and unpacked data for AgentDataUpdate events raised by the AI721Contract contract.
type AI721ContractAgentDataUpdateIterator struct {
	Event *AI721ContractAgentDataUpdate // Event containing the contract specifics and raw log

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
func (it *AI721ContractAgentDataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AI721ContractAgentDataUpdate)
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
		it.Event = new(AI721ContractAgentDataUpdate)
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
func (it *AI721ContractAgentDataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AI721ContractAgentDataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AI721ContractAgentDataUpdate represents a AgentDataUpdate event raised by the AI721Contract contract.
type AI721ContractAgentDataUpdate struct {
	AgentId      *big.Int
	PromptIndex  *big.Int
	OldSysPrompt []byte
	NewSysPrompt []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAgentDataUpdate is a free log retrieval operation binding the contract event 0xe42abf7d4a793286da8cc1399cb577a1f5a0e133dfee371bb3a5abbdd77b011e.
//
// Solidity: event AgentDataUpdate(uint256 indexed agentId, uint256 promptIndex, bytes oldSysPrompt, bytes newSysPrompt)
func (_AI721Contract *AI721ContractFilterer) FilterAgentDataUpdate(opts *bind.FilterOpts, agentId []*big.Int) (*AI721ContractAgentDataUpdateIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _AI721Contract.contract.FilterLogs(opts, "AgentDataUpdate", agentIdRule)
	if err != nil {
		return nil, err
	}
	return &AI721ContractAgentDataUpdateIterator{contract: _AI721Contract.contract, event: "AgentDataUpdate", logs: logs, sub: sub}, nil
}

// WatchAgentDataUpdate is a free log subscription operation binding the contract event 0xe42abf7d4a793286da8cc1399cb577a1f5a0e133dfee371bb3a5abbdd77b011e.
//
// Solidity: event AgentDataUpdate(uint256 indexed agentId, uint256 promptIndex, bytes oldSysPrompt, bytes newSysPrompt)
func (_AI721Contract *AI721ContractFilterer) WatchAgentDataUpdate(opts *bind.WatchOpts, sink chan<- *AI721ContractAgentDataUpdate, agentId []*big.Int) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _AI721Contract.contract.WatchLogs(opts, "AgentDataUpdate", agentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AI721ContractAgentDataUpdate)
				if err := _AI721Contract.contract.UnpackLog(event, "AgentDataUpdate", log); err != nil {
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
func (_AI721Contract *AI721ContractFilterer) ParseAgentDataUpdate(log types.Log) (*AI721ContractAgentDataUpdate, error) {
	event := new(AI721ContractAgentDataUpdate)
	if err := _AI721Contract.contract.UnpackLog(event, "AgentDataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AI721ContractAgentFeeUpdateIterator is returned from FilterAgentFeeUpdate and is used to iterate over the raw logs and unpacked data for AgentFeeUpdate events raised by the AI721Contract contract.
type AI721ContractAgentFeeUpdateIterator struct {
	Event *AI721ContractAgentFeeUpdate // Event containing the contract specifics and raw log

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
func (it *AI721ContractAgentFeeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AI721ContractAgentFeeUpdate)
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
		it.Event = new(AI721ContractAgentFeeUpdate)
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
func (it *AI721ContractAgentFeeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AI721ContractAgentFeeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AI721ContractAgentFeeUpdate represents a AgentFeeUpdate event raised by the AI721Contract contract.
type AI721ContractAgentFeeUpdate struct {
	AgentId *big.Int
	Fee     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentFeeUpdate is a free log retrieval operation binding the contract event 0xa08d8197034aee8915921dd8aa7d95cf711690dd77f0b676dded49b3f9a632d1.
//
// Solidity: event AgentFeeUpdate(uint256 indexed agentId, uint256 fee)
func (_AI721Contract *AI721ContractFilterer) FilterAgentFeeUpdate(opts *bind.FilterOpts, agentId []*big.Int) (*AI721ContractAgentFeeUpdateIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _AI721Contract.contract.FilterLogs(opts, "AgentFeeUpdate", agentIdRule)
	if err != nil {
		return nil, err
	}
	return &AI721ContractAgentFeeUpdateIterator{contract: _AI721Contract.contract, event: "AgentFeeUpdate", logs: logs, sub: sub}, nil
}

// WatchAgentFeeUpdate is a free log subscription operation binding the contract event 0xa08d8197034aee8915921dd8aa7d95cf711690dd77f0b676dded49b3f9a632d1.
//
// Solidity: event AgentFeeUpdate(uint256 indexed agentId, uint256 fee)
func (_AI721Contract *AI721ContractFilterer) WatchAgentFeeUpdate(opts *bind.WatchOpts, sink chan<- *AI721ContractAgentFeeUpdate, agentId []*big.Int) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _AI721Contract.contract.WatchLogs(opts, "AgentFeeUpdate", agentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AI721ContractAgentFeeUpdate)
				if err := _AI721Contract.contract.UnpackLog(event, "AgentFeeUpdate", log); err != nil {
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
func (_AI721Contract *AI721ContractFilterer) ParseAgentFeeUpdate(log types.Log) (*AI721ContractAgentFeeUpdate, error) {
	event := new(AI721ContractAgentFeeUpdate)
	if err := _AI721Contract.contract.UnpackLog(event, "AgentFeeUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AI721ContractAgentMissionAddNewIterator is returned from FilterAgentMissionAddNew and is used to iterate over the raw logs and unpacked data for AgentMissionAddNew events raised by the AI721Contract contract.
type AI721ContractAgentMissionAddNewIterator struct {
	Event *AI721ContractAgentMissionAddNew // Event containing the contract specifics and raw log

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
func (it *AI721ContractAgentMissionAddNewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AI721ContractAgentMissionAddNew)
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
		it.Event = new(AI721ContractAgentMissionAddNew)
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
func (it *AI721ContractAgentMissionAddNewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AI721ContractAgentMissionAddNewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AI721ContractAgentMissionAddNew represents a AgentMissionAddNew event raised by the AI721Contract contract.
type AI721ContractAgentMissionAddNew struct {
	AgentId  *big.Int
	Missions [][]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAgentMissionAddNew is a free log retrieval operation binding the contract event 0x12ccdcc9c8e92b22004686225bd3df163c042e77b03eab4566800d40b5047f91.
//
// Solidity: event AgentMissionAddNew(uint256 indexed agentId, bytes[] missions)
func (_AI721Contract *AI721ContractFilterer) FilterAgentMissionAddNew(opts *bind.FilterOpts, agentId []*big.Int) (*AI721ContractAgentMissionAddNewIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _AI721Contract.contract.FilterLogs(opts, "AgentMissionAddNew", agentIdRule)
	if err != nil {
		return nil, err
	}
	return &AI721ContractAgentMissionAddNewIterator{contract: _AI721Contract.contract, event: "AgentMissionAddNew", logs: logs, sub: sub}, nil
}

// WatchAgentMissionAddNew is a free log subscription operation binding the contract event 0x12ccdcc9c8e92b22004686225bd3df163c042e77b03eab4566800d40b5047f91.
//
// Solidity: event AgentMissionAddNew(uint256 indexed agentId, bytes[] missions)
func (_AI721Contract *AI721ContractFilterer) WatchAgentMissionAddNew(opts *bind.WatchOpts, sink chan<- *AI721ContractAgentMissionAddNew, agentId []*big.Int) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _AI721Contract.contract.WatchLogs(opts, "AgentMissionAddNew", agentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AI721ContractAgentMissionAddNew)
				if err := _AI721Contract.contract.UnpackLog(event, "AgentMissionAddNew", log); err != nil {
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
func (_AI721Contract *AI721ContractFilterer) ParseAgentMissionAddNew(log types.Log) (*AI721ContractAgentMissionAddNew, error) {
	event := new(AI721ContractAgentMissionAddNew)
	if err := _AI721Contract.contract.UnpackLog(event, "AgentMissionAddNew", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AI721ContractAgentMissionUpdateIterator is returned from FilterAgentMissionUpdate and is used to iterate over the raw logs and unpacked data for AgentMissionUpdate events raised by the AI721Contract contract.
type AI721ContractAgentMissionUpdateIterator struct {
	Event *AI721ContractAgentMissionUpdate // Event containing the contract specifics and raw log

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
func (it *AI721ContractAgentMissionUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AI721ContractAgentMissionUpdate)
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
		it.Event = new(AI721ContractAgentMissionUpdate)
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
func (it *AI721ContractAgentMissionUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AI721ContractAgentMissionUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AI721ContractAgentMissionUpdate represents a AgentMissionUpdate event raised by the AI721Contract contract.
type AI721ContractAgentMissionUpdate struct {
	AgentId       *big.Int
	MissionIndex  *big.Int
	OldSysMission []byte
	NewSysMission []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAgentMissionUpdate is a free log retrieval operation binding the contract event 0x0a9b80bd675e3f5788f1a5da687efd147dbc4729245a7f300ce1074bbd535127.
//
// Solidity: event AgentMissionUpdate(uint256 indexed agentId, uint256 missionIndex, bytes oldSysMission, bytes newSysMission)
func (_AI721Contract *AI721ContractFilterer) FilterAgentMissionUpdate(opts *bind.FilterOpts, agentId []*big.Int) (*AI721ContractAgentMissionUpdateIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _AI721Contract.contract.FilterLogs(opts, "AgentMissionUpdate", agentIdRule)
	if err != nil {
		return nil, err
	}
	return &AI721ContractAgentMissionUpdateIterator{contract: _AI721Contract.contract, event: "AgentMissionUpdate", logs: logs, sub: sub}, nil
}

// WatchAgentMissionUpdate is a free log subscription operation binding the contract event 0x0a9b80bd675e3f5788f1a5da687efd147dbc4729245a7f300ce1074bbd535127.
//
// Solidity: event AgentMissionUpdate(uint256 indexed agentId, uint256 missionIndex, bytes oldSysMission, bytes newSysMission)
func (_AI721Contract *AI721ContractFilterer) WatchAgentMissionUpdate(opts *bind.WatchOpts, sink chan<- *AI721ContractAgentMissionUpdate, agentId []*big.Int) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _AI721Contract.contract.WatchLogs(opts, "AgentMissionUpdate", agentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AI721ContractAgentMissionUpdate)
				if err := _AI721Contract.contract.UnpackLog(event, "AgentMissionUpdate", log); err != nil {
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
func (_AI721Contract *AI721ContractFilterer) ParseAgentMissionUpdate(log types.Log) (*AI721ContractAgentMissionUpdate, error) {
	event := new(AI721ContractAgentMissionUpdate)
	if err := _AI721Contract.contract.UnpackLog(event, "AgentMissionUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AI721ContractAgentURIUpdateIterator is returned from FilterAgentURIUpdate and is used to iterate over the raw logs and unpacked data for AgentURIUpdate events raised by the AI721Contract contract.
type AI721ContractAgentURIUpdateIterator struct {
	Event *AI721ContractAgentURIUpdate // Event containing the contract specifics and raw log

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
func (it *AI721ContractAgentURIUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AI721ContractAgentURIUpdate)
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
		it.Event = new(AI721ContractAgentURIUpdate)
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
func (it *AI721ContractAgentURIUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AI721ContractAgentURIUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AI721ContractAgentURIUpdate represents a AgentURIUpdate event raised by the AI721Contract contract.
type AI721ContractAgentURIUpdate struct {
	AgentId *big.Int
	Uri     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentURIUpdate is a free log retrieval operation binding the contract event 0x706a4e8eb2f354c7f4d96e5ea1984f36e72482629987edad78c9940ea037c362.
//
// Solidity: event AgentURIUpdate(uint256 indexed agentId, string uri)
func (_AI721Contract *AI721ContractFilterer) FilterAgentURIUpdate(opts *bind.FilterOpts, agentId []*big.Int) (*AI721ContractAgentURIUpdateIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _AI721Contract.contract.FilterLogs(opts, "AgentURIUpdate", agentIdRule)
	if err != nil {
		return nil, err
	}
	return &AI721ContractAgentURIUpdateIterator{contract: _AI721Contract.contract, event: "AgentURIUpdate", logs: logs, sub: sub}, nil
}

// WatchAgentURIUpdate is a free log subscription operation binding the contract event 0x706a4e8eb2f354c7f4d96e5ea1984f36e72482629987edad78c9940ea037c362.
//
// Solidity: event AgentURIUpdate(uint256 indexed agentId, string uri)
func (_AI721Contract *AI721ContractFilterer) WatchAgentURIUpdate(opts *bind.WatchOpts, sink chan<- *AI721ContractAgentURIUpdate, agentId []*big.Int) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _AI721Contract.contract.WatchLogs(opts, "AgentURIUpdate", agentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AI721ContractAgentURIUpdate)
				if err := _AI721Contract.contract.UnpackLog(event, "AgentURIUpdate", log); err != nil {
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
func (_AI721Contract *AI721ContractFilterer) ParseAgentURIUpdate(log types.Log) (*AI721ContractAgentURIUpdate, error) {
	event := new(AI721ContractAgentURIUpdate)
	if err := _AI721Contract.contract.UnpackLog(event, "AgentURIUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AI721ContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AI721Contract contract.
type AI721ContractApprovalIterator struct {
	Event *AI721ContractApproval // Event containing the contract specifics and raw log

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
func (it *AI721ContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AI721ContractApproval)
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
		it.Event = new(AI721ContractApproval)
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
func (it *AI721ContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AI721ContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AI721ContractApproval represents a Approval event raised by the AI721Contract contract.
type AI721ContractApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_AI721Contract *AI721ContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*AI721ContractApprovalIterator, error) {

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

	logs, sub, err := _AI721Contract.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AI721ContractApprovalIterator{contract: _AI721Contract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_AI721Contract *AI721ContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AI721ContractApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _AI721Contract.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AI721ContractApproval)
				if err := _AI721Contract.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_AI721Contract *AI721ContractFilterer) ParseApproval(log types.Log) (*AI721ContractApproval, error) {
	event := new(AI721ContractApproval)
	if err := _AI721Contract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AI721ContractApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the AI721Contract contract.
type AI721ContractApprovalForAllIterator struct {
	Event *AI721ContractApprovalForAll // Event containing the contract specifics and raw log

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
func (it *AI721ContractApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AI721ContractApprovalForAll)
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
		it.Event = new(AI721ContractApprovalForAll)
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
func (it *AI721ContractApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AI721ContractApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AI721ContractApprovalForAll represents a ApprovalForAll event raised by the AI721Contract contract.
type AI721ContractApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_AI721Contract *AI721ContractFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*AI721ContractApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AI721Contract.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &AI721ContractApprovalForAllIterator{contract: _AI721Contract.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_AI721Contract *AI721ContractFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *AI721ContractApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AI721Contract.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AI721ContractApprovalForAll)
				if err := _AI721Contract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_AI721Contract *AI721ContractFilterer) ParseApprovalForAll(log types.Log) (*AI721ContractApprovalForAll, error) {
	event := new(AI721ContractApprovalForAll)
	if err := _AI721Contract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AI721ContractBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the AI721Contract contract.
type AI721ContractBatchMetadataUpdateIterator struct {
	Event *AI721ContractBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *AI721ContractBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AI721ContractBatchMetadataUpdate)
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
		it.Event = new(AI721ContractBatchMetadataUpdate)
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
func (it *AI721ContractBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AI721ContractBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AI721ContractBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the AI721Contract contract.
type AI721ContractBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_AI721Contract *AI721ContractFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*AI721ContractBatchMetadataUpdateIterator, error) {

	logs, sub, err := _AI721Contract.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &AI721ContractBatchMetadataUpdateIterator{contract: _AI721Contract.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_AI721Contract *AI721ContractFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *AI721ContractBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _AI721Contract.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AI721ContractBatchMetadataUpdate)
				if err := _AI721Contract.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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
func (_AI721Contract *AI721ContractFilterer) ParseBatchMetadataUpdate(log types.Log) (*AI721ContractBatchMetadataUpdate, error) {
	event := new(AI721ContractBatchMetadataUpdate)
	if err := _AI721Contract.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AI721ContractEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the AI721Contract contract.
type AI721ContractEIP712DomainChangedIterator struct {
	Event *AI721ContractEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *AI721ContractEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AI721ContractEIP712DomainChanged)
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
		it.Event = new(AI721ContractEIP712DomainChanged)
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
func (it *AI721ContractEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AI721ContractEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AI721ContractEIP712DomainChanged represents a EIP712DomainChanged event raised by the AI721Contract contract.
type AI721ContractEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_AI721Contract *AI721ContractFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*AI721ContractEIP712DomainChangedIterator, error) {

	logs, sub, err := _AI721Contract.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &AI721ContractEIP712DomainChangedIterator{contract: _AI721Contract.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_AI721Contract *AI721ContractFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *AI721ContractEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _AI721Contract.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AI721ContractEIP712DomainChanged)
				if err := _AI721Contract.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_AI721Contract *AI721ContractFilterer) ParseEIP712DomainChanged(log types.Log) (*AI721ContractEIP712DomainChanged, error) {
	event := new(AI721ContractEIP712DomainChanged)
	if err := _AI721Contract.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AI721ContractFeesClaimedIterator is returned from FilterFeesClaimed and is used to iterate over the raw logs and unpacked data for FeesClaimed events raised by the AI721Contract contract.
type AI721ContractFeesClaimedIterator struct {
	Event *AI721ContractFeesClaimed // Event containing the contract specifics and raw log

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
func (it *AI721ContractFeesClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AI721ContractFeesClaimed)
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
		it.Event = new(AI721ContractFeesClaimed)
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
func (it *AI721ContractFeesClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AI721ContractFeesClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AI721ContractFeesClaimed represents a FeesClaimed event raised by the AI721Contract contract.
type AI721ContractFeesClaimed struct {
	Claimer common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeesClaimed is a free log retrieval operation binding the contract event 0x9493e5bbe4e8e0ac67284469a2d677403d0378a85a59e341d3abc433d0d9a209.
//
// Solidity: event FeesClaimed(address indexed claimer, uint256 amount)
func (_AI721Contract *AI721ContractFilterer) FilterFeesClaimed(opts *bind.FilterOpts, claimer []common.Address) (*AI721ContractFeesClaimedIterator, error) {

	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _AI721Contract.contract.FilterLogs(opts, "FeesClaimed", claimerRule)
	if err != nil {
		return nil, err
	}
	return &AI721ContractFeesClaimedIterator{contract: _AI721Contract.contract, event: "FeesClaimed", logs: logs, sub: sub}, nil
}

// WatchFeesClaimed is a free log subscription operation binding the contract event 0x9493e5bbe4e8e0ac67284469a2d677403d0378a85a59e341d3abc433d0d9a209.
//
// Solidity: event FeesClaimed(address indexed claimer, uint256 amount)
func (_AI721Contract *AI721ContractFilterer) WatchFeesClaimed(opts *bind.WatchOpts, sink chan<- *AI721ContractFeesClaimed, claimer []common.Address) (event.Subscription, error) {

	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _AI721Contract.contract.WatchLogs(opts, "FeesClaimed", claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AI721ContractFeesClaimed)
				if err := _AI721Contract.contract.UnpackLog(event, "FeesClaimed", log); err != nil {
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

// ParseFeesClaimed is a log parse operation binding the contract event 0x9493e5bbe4e8e0ac67284469a2d677403d0378a85a59e341d3abc433d0d9a209.
//
// Solidity: event FeesClaimed(address indexed claimer, uint256 amount)
func (_AI721Contract *AI721ContractFilterer) ParseFeesClaimed(log types.Log) (*AI721ContractFeesClaimed, error) {
	event := new(AI721ContractFeesClaimed)
	if err := _AI721Contract.contract.UnpackLog(event, "FeesClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AI721ContractInferencePerformedIterator is returned from FilterInferencePerformed and is used to iterate over the raw logs and unpacked data for InferencePerformed events raised by the AI721Contract contract.
type AI721ContractInferencePerformedIterator struct {
	Event *AI721ContractInferencePerformed // Event containing the contract specifics and raw log

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
func (it *AI721ContractInferencePerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AI721ContractInferencePerformed)
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
		it.Event = new(AI721ContractInferencePerformed)
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
func (it *AI721ContractInferencePerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AI721ContractInferencePerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AI721ContractInferencePerformed represents a InferencePerformed event raised by the AI721Contract contract.
type AI721ContractInferencePerformed struct {
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
func (_AI721Contract *AI721ContractFilterer) FilterInferencePerformed(opts *bind.FilterOpts, tokenId []*big.Int, caller []common.Address) (*AI721ContractInferencePerformedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _AI721Contract.contract.FilterLogs(opts, "InferencePerformed", tokenIdRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &AI721ContractInferencePerformedIterator{contract: _AI721Contract.contract, event: "InferencePerformed", logs: logs, sub: sub}, nil
}

// WatchInferencePerformed is a free log subscription operation binding the contract event 0xcf35460eca25a0549d5eb14c712236d61c9a0bad90c834f996c5f2a98d332719.
//
// Solidity: event InferencePerformed(uint256 indexed tokenId, address indexed caller, bytes data, uint256 fee, string externalData, uint256 inferenceId)
func (_AI721Contract *AI721ContractFilterer) WatchInferencePerformed(opts *bind.WatchOpts, sink chan<- *AI721ContractInferencePerformed, tokenId []*big.Int, caller []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _AI721Contract.contract.WatchLogs(opts, "InferencePerformed", tokenIdRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AI721ContractInferencePerformed)
				if err := _AI721Contract.contract.UnpackLog(event, "InferencePerformed", log); err != nil {
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
func (_AI721Contract *AI721ContractFilterer) ParseInferencePerformed(log types.Log) (*AI721ContractInferencePerformed, error) {
	event := new(AI721ContractInferencePerformed)
	if err := _AI721Contract.contract.UnpackLog(event, "InferencePerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AI721ContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AI721Contract contract.
type AI721ContractInitializedIterator struct {
	Event *AI721ContractInitialized // Event containing the contract specifics and raw log

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
func (it *AI721ContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AI721ContractInitialized)
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
		it.Event = new(AI721ContractInitialized)
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
func (it *AI721ContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AI721ContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AI721ContractInitialized represents a Initialized event raised by the AI721Contract contract.
type AI721ContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AI721Contract *AI721ContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*AI721ContractInitializedIterator, error) {

	logs, sub, err := _AI721Contract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AI721ContractInitializedIterator{contract: _AI721Contract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AI721Contract *AI721ContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AI721ContractInitialized) (event.Subscription, error) {

	logs, sub, err := _AI721Contract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AI721ContractInitialized)
				if err := _AI721Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AI721Contract *AI721ContractFilterer) ParseInitialized(log types.Log) (*AI721ContractInitialized, error) {
	event := new(AI721ContractInitialized)
	if err := _AI721Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AI721ContractManagerAuthorizationIterator is returned from FilterManagerAuthorization and is used to iterate over the raw logs and unpacked data for ManagerAuthorization events raised by the AI721Contract contract.
type AI721ContractManagerAuthorizationIterator struct {
	Event *AI721ContractManagerAuthorization // Event containing the contract specifics and raw log

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
func (it *AI721ContractManagerAuthorizationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AI721ContractManagerAuthorization)
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
		it.Event = new(AI721ContractManagerAuthorization)
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
func (it *AI721ContractManagerAuthorizationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AI721ContractManagerAuthorizationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AI721ContractManagerAuthorization represents a ManagerAuthorization event raised by the AI721Contract contract.
type AI721ContractManagerAuthorization struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerAuthorization is a free log retrieval operation binding the contract event 0x3fbc8e71624117c0dc0fcdbe40685681cecc7c3c43de81a686cda4b61c78e35b.
//
// Solidity: event ManagerAuthorization(address indexed account)
func (_AI721Contract *AI721ContractFilterer) FilterManagerAuthorization(opts *bind.FilterOpts, account []common.Address) (*AI721ContractManagerAuthorizationIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AI721Contract.contract.FilterLogs(opts, "ManagerAuthorization", accountRule)
	if err != nil {
		return nil, err
	}
	return &AI721ContractManagerAuthorizationIterator{contract: _AI721Contract.contract, event: "ManagerAuthorization", logs: logs, sub: sub}, nil
}

// WatchManagerAuthorization is a free log subscription operation binding the contract event 0x3fbc8e71624117c0dc0fcdbe40685681cecc7c3c43de81a686cda4b61c78e35b.
//
// Solidity: event ManagerAuthorization(address indexed account)
func (_AI721Contract *AI721ContractFilterer) WatchManagerAuthorization(opts *bind.WatchOpts, sink chan<- *AI721ContractManagerAuthorization, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AI721Contract.contract.WatchLogs(opts, "ManagerAuthorization", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AI721ContractManagerAuthorization)
				if err := _AI721Contract.contract.UnpackLog(event, "ManagerAuthorization", log); err != nil {
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

// ParseManagerAuthorization is a log parse operation binding the contract event 0x3fbc8e71624117c0dc0fcdbe40685681cecc7c3c43de81a686cda4b61c78e35b.
//
// Solidity: event ManagerAuthorization(address indexed account)
func (_AI721Contract *AI721ContractFilterer) ParseManagerAuthorization(log types.Log) (*AI721ContractManagerAuthorization, error) {
	event := new(AI721ContractManagerAuthorization)
	if err := _AI721Contract.contract.UnpackLog(event, "ManagerAuthorization", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AI721ContractManagerDeauthorizationIterator is returned from FilterManagerDeauthorization and is used to iterate over the raw logs and unpacked data for ManagerDeauthorization events raised by the AI721Contract contract.
type AI721ContractManagerDeauthorizationIterator struct {
	Event *AI721ContractManagerDeauthorization // Event containing the contract specifics and raw log

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
func (it *AI721ContractManagerDeauthorizationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AI721ContractManagerDeauthorization)
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
		it.Event = new(AI721ContractManagerDeauthorization)
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
func (it *AI721ContractManagerDeauthorizationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AI721ContractManagerDeauthorizationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AI721ContractManagerDeauthorization represents a ManagerDeauthorization event raised by the AI721Contract contract.
type AI721ContractManagerDeauthorization struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerDeauthorization is a free log retrieval operation binding the contract event 0x20c29af9eb3de2601188ceae57a4075ba3593ce15d4142aef070ac53d389356c.
//
// Solidity: event ManagerDeauthorization(address indexed account)
func (_AI721Contract *AI721ContractFilterer) FilterManagerDeauthorization(opts *bind.FilterOpts, account []common.Address) (*AI721ContractManagerDeauthorizationIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AI721Contract.contract.FilterLogs(opts, "ManagerDeauthorization", accountRule)
	if err != nil {
		return nil, err
	}
	return &AI721ContractManagerDeauthorizationIterator{contract: _AI721Contract.contract, event: "ManagerDeauthorization", logs: logs, sub: sub}, nil
}

// WatchManagerDeauthorization is a free log subscription operation binding the contract event 0x20c29af9eb3de2601188ceae57a4075ba3593ce15d4142aef070ac53d389356c.
//
// Solidity: event ManagerDeauthorization(address indexed account)
func (_AI721Contract *AI721ContractFilterer) WatchManagerDeauthorization(opts *bind.WatchOpts, sink chan<- *AI721ContractManagerDeauthorization, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AI721Contract.contract.WatchLogs(opts, "ManagerDeauthorization", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AI721ContractManagerDeauthorization)
				if err := _AI721Contract.contract.UnpackLog(event, "ManagerDeauthorization", log); err != nil {
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

// ParseManagerDeauthorization is a log parse operation binding the contract event 0x20c29af9eb3de2601188ceae57a4075ba3593ce15d4142aef070ac53d389356c.
//
// Solidity: event ManagerDeauthorization(address indexed account)
func (_AI721Contract *AI721ContractFilterer) ParseManagerDeauthorization(log types.Log) (*AI721ContractManagerDeauthorization, error) {
	event := new(AI721ContractManagerDeauthorization)
	if err := _AI721Contract.contract.UnpackLog(event, "ManagerDeauthorization", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AI721ContractMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the AI721Contract contract.
type AI721ContractMetadataUpdateIterator struct {
	Event *AI721ContractMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *AI721ContractMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AI721ContractMetadataUpdate)
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
		it.Event = new(AI721ContractMetadataUpdate)
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
func (it *AI721ContractMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AI721ContractMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AI721ContractMetadataUpdate represents a MetadataUpdate event raised by the AI721Contract contract.
type AI721ContractMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_AI721Contract *AI721ContractFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*AI721ContractMetadataUpdateIterator, error) {

	logs, sub, err := _AI721Contract.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &AI721ContractMetadataUpdateIterator{contract: _AI721Contract.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_AI721Contract *AI721ContractFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *AI721ContractMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _AI721Contract.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AI721ContractMetadataUpdate)
				if err := _AI721Contract.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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
func (_AI721Contract *AI721ContractFilterer) ParseMetadataUpdate(log types.Log) (*AI721ContractMetadataUpdate, error) {
	event := new(AI721ContractMetadataUpdate)
	if err := _AI721Contract.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AI721ContractMintPriceUpdateIterator is returned from FilterMintPriceUpdate and is used to iterate over the raw logs and unpacked data for MintPriceUpdate events raised by the AI721Contract contract.
type AI721ContractMintPriceUpdateIterator struct {
	Event *AI721ContractMintPriceUpdate // Event containing the contract specifics and raw log

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
func (it *AI721ContractMintPriceUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AI721ContractMintPriceUpdate)
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
		it.Event = new(AI721ContractMintPriceUpdate)
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
func (it *AI721ContractMintPriceUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AI721ContractMintPriceUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AI721ContractMintPriceUpdate represents a MintPriceUpdate event raised by the AI721Contract contract.
type AI721ContractMintPriceUpdate struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMintPriceUpdate is a free log retrieval operation binding the contract event 0x23050b539195eebd064c1bec834445b7d028a20c345600e868a74d7ca93c5e86.
//
// Solidity: event MintPriceUpdate(uint256 newValue)
func (_AI721Contract *AI721ContractFilterer) FilterMintPriceUpdate(opts *bind.FilterOpts) (*AI721ContractMintPriceUpdateIterator, error) {

	logs, sub, err := _AI721Contract.contract.FilterLogs(opts, "MintPriceUpdate")
	if err != nil {
		return nil, err
	}
	return &AI721ContractMintPriceUpdateIterator{contract: _AI721Contract.contract, event: "MintPriceUpdate", logs: logs, sub: sub}, nil
}

// WatchMintPriceUpdate is a free log subscription operation binding the contract event 0x23050b539195eebd064c1bec834445b7d028a20c345600e868a74d7ca93c5e86.
//
// Solidity: event MintPriceUpdate(uint256 newValue)
func (_AI721Contract *AI721ContractFilterer) WatchMintPriceUpdate(opts *bind.WatchOpts, sink chan<- *AI721ContractMintPriceUpdate) (event.Subscription, error) {

	logs, sub, err := _AI721Contract.contract.WatchLogs(opts, "MintPriceUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AI721ContractMintPriceUpdate)
				if err := _AI721Contract.contract.UnpackLog(event, "MintPriceUpdate", log); err != nil {
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
func (_AI721Contract *AI721ContractFilterer) ParseMintPriceUpdate(log types.Log) (*AI721ContractMintPriceUpdate, error) {
	event := new(AI721ContractMintPriceUpdate)
	if err := _AI721Contract.contract.UnpackLog(event, "MintPriceUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AI721ContractNewTokenIterator is returned from FilterNewToken and is used to iterate over the raw logs and unpacked data for NewToken events raised by the AI721Contract contract.
type AI721ContractNewTokenIterator struct {
	Event *AI721ContractNewToken // Event containing the contract specifics and raw log

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
func (it *AI721ContractNewTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AI721ContractNewToken)
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
		it.Event = new(AI721ContractNewToken)
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
func (it *AI721ContractNewTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AI721ContractNewTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AI721ContractNewToken represents a NewToken event raised by the AI721Contract contract.
type AI721ContractNewToken struct {
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
func (_AI721Contract *AI721ContractFilterer) FilterNewToken(opts *bind.FilterOpts, tokenId []*big.Int, minter []common.Address) (*AI721ContractNewTokenIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _AI721Contract.contract.FilterLogs(opts, "NewToken", tokenIdRule, minterRule)
	if err != nil {
		return nil, err
	}
	return &AI721ContractNewTokenIterator{contract: _AI721Contract.contract, event: "NewToken", logs: logs, sub: sub}, nil
}

// WatchNewToken is a free log subscription operation binding the contract event 0x61beab98a81083e3c0239c33e149bef1316ca78f15b9f29125039f5521a06d06.
//
// Solidity: event NewToken(uint256 indexed tokenId, string uri, bytes sysPrompt, uint256 fee, address indexed minter)
func (_AI721Contract *AI721ContractFilterer) WatchNewToken(opts *bind.WatchOpts, sink chan<- *AI721ContractNewToken, tokenId []*big.Int, minter []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _AI721Contract.contract.WatchLogs(opts, "NewToken", tokenIdRule, minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AI721ContractNewToken)
				if err := _AI721Contract.contract.UnpackLog(event, "NewToken", log); err != nil {
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
func (_AI721Contract *AI721ContractFilterer) ParseNewToken(log types.Log) (*AI721ContractNewToken, error) {
	event := new(AI721ContractNewToken)
	if err := _AI721Contract.contract.UnpackLog(event, "NewToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AI721ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AI721Contract contract.
type AI721ContractOwnershipTransferredIterator struct {
	Event *AI721ContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AI721ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AI721ContractOwnershipTransferred)
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
		it.Event = new(AI721ContractOwnershipTransferred)
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
func (it *AI721ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AI721ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AI721ContractOwnershipTransferred represents a OwnershipTransferred event raised by the AI721Contract contract.
type AI721ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AI721Contract *AI721ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AI721ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AI721Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AI721ContractOwnershipTransferredIterator{contract: _AI721Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AI721Contract *AI721ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AI721ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AI721Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AI721ContractOwnershipTransferred)
				if err := _AI721Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AI721Contract *AI721ContractFilterer) ParseOwnershipTransferred(log types.Log) (*AI721ContractOwnershipTransferred, error) {
	event := new(AI721ContractOwnershipTransferred)
	if err := _AI721Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AI721ContractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AI721Contract contract.
type AI721ContractPausedIterator struct {
	Event *AI721ContractPaused // Event containing the contract specifics and raw log

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
func (it *AI721ContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AI721ContractPaused)
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
		it.Event = new(AI721ContractPaused)
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
func (it *AI721ContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AI721ContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AI721ContractPaused represents a Paused event raised by the AI721Contract contract.
type AI721ContractPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AI721Contract *AI721ContractFilterer) FilterPaused(opts *bind.FilterOpts) (*AI721ContractPausedIterator, error) {

	logs, sub, err := _AI721Contract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AI721ContractPausedIterator{contract: _AI721Contract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AI721Contract *AI721ContractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AI721ContractPaused) (event.Subscription, error) {

	logs, sub, err := _AI721Contract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AI721ContractPaused)
				if err := _AI721Contract.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_AI721Contract *AI721ContractFilterer) ParsePaused(log types.Log) (*AI721ContractPaused, error) {
	event := new(AI721ContractPaused)
	if err := _AI721Contract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AI721ContractRoyaltyPortionUpdateIterator is returned from FilterRoyaltyPortionUpdate and is used to iterate over the raw logs and unpacked data for RoyaltyPortionUpdate events raised by the AI721Contract contract.
type AI721ContractRoyaltyPortionUpdateIterator struct {
	Event *AI721ContractRoyaltyPortionUpdate // Event containing the contract specifics and raw log

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
func (it *AI721ContractRoyaltyPortionUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AI721ContractRoyaltyPortionUpdate)
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
		it.Event = new(AI721ContractRoyaltyPortionUpdate)
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
func (it *AI721ContractRoyaltyPortionUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AI721ContractRoyaltyPortionUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AI721ContractRoyaltyPortionUpdate represents a RoyaltyPortionUpdate event raised by the AI721Contract contract.
type AI721ContractRoyaltyPortionUpdate struct {
	NewValue uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyPortionUpdate is a free log retrieval operation binding the contract event 0xb1f3037624bd2d961f6d56696cc10ccc3a676db381e671b1bc58f0ab1f686dd5.
//
// Solidity: event RoyaltyPortionUpdate(uint16 newValue)
func (_AI721Contract *AI721ContractFilterer) FilterRoyaltyPortionUpdate(opts *bind.FilterOpts) (*AI721ContractRoyaltyPortionUpdateIterator, error) {

	logs, sub, err := _AI721Contract.contract.FilterLogs(opts, "RoyaltyPortionUpdate")
	if err != nil {
		return nil, err
	}
	return &AI721ContractRoyaltyPortionUpdateIterator{contract: _AI721Contract.contract, event: "RoyaltyPortionUpdate", logs: logs, sub: sub}, nil
}

// WatchRoyaltyPortionUpdate is a free log subscription operation binding the contract event 0xb1f3037624bd2d961f6d56696cc10ccc3a676db381e671b1bc58f0ab1f686dd5.
//
// Solidity: event RoyaltyPortionUpdate(uint16 newValue)
func (_AI721Contract *AI721ContractFilterer) WatchRoyaltyPortionUpdate(opts *bind.WatchOpts, sink chan<- *AI721ContractRoyaltyPortionUpdate) (event.Subscription, error) {

	logs, sub, err := _AI721Contract.contract.WatchLogs(opts, "RoyaltyPortionUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AI721ContractRoyaltyPortionUpdate)
				if err := _AI721Contract.contract.UnpackLog(event, "RoyaltyPortionUpdate", log); err != nil {
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
func (_AI721Contract *AI721ContractFilterer) ParseRoyaltyPortionUpdate(log types.Log) (*AI721ContractRoyaltyPortionUpdate, error) {
	event := new(AI721ContractRoyaltyPortionUpdate)
	if err := _AI721Contract.contract.UnpackLog(event, "RoyaltyPortionUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AI721ContractRoyaltyReceiverUpdateIterator is returned from FilterRoyaltyReceiverUpdate and is used to iterate over the raw logs and unpacked data for RoyaltyReceiverUpdate events raised by the AI721Contract contract.
type AI721ContractRoyaltyReceiverUpdateIterator struct {
	Event *AI721ContractRoyaltyReceiverUpdate // Event containing the contract specifics and raw log

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
func (it *AI721ContractRoyaltyReceiverUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AI721ContractRoyaltyReceiverUpdate)
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
		it.Event = new(AI721ContractRoyaltyReceiverUpdate)
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
func (it *AI721ContractRoyaltyReceiverUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AI721ContractRoyaltyReceiverUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AI721ContractRoyaltyReceiverUpdate represents a RoyaltyReceiverUpdate event raised by the AI721Contract contract.
type AI721ContractRoyaltyReceiverUpdate struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyReceiverUpdate is a free log retrieval operation binding the contract event 0xec6b72b10aed766af02b35918b55be261c89aaaa4c8add826471ce35ec7f97b3.
//
// Solidity: event RoyaltyReceiverUpdate(address newAddress)
func (_AI721Contract *AI721ContractFilterer) FilterRoyaltyReceiverUpdate(opts *bind.FilterOpts) (*AI721ContractRoyaltyReceiverUpdateIterator, error) {

	logs, sub, err := _AI721Contract.contract.FilterLogs(opts, "RoyaltyReceiverUpdate")
	if err != nil {
		return nil, err
	}
	return &AI721ContractRoyaltyReceiverUpdateIterator{contract: _AI721Contract.contract, event: "RoyaltyReceiverUpdate", logs: logs, sub: sub}, nil
}

// WatchRoyaltyReceiverUpdate is a free log subscription operation binding the contract event 0xec6b72b10aed766af02b35918b55be261c89aaaa4c8add826471ce35ec7f97b3.
//
// Solidity: event RoyaltyReceiverUpdate(address newAddress)
func (_AI721Contract *AI721ContractFilterer) WatchRoyaltyReceiverUpdate(opts *bind.WatchOpts, sink chan<- *AI721ContractRoyaltyReceiverUpdate) (event.Subscription, error) {

	logs, sub, err := _AI721Contract.contract.WatchLogs(opts, "RoyaltyReceiverUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AI721ContractRoyaltyReceiverUpdate)
				if err := _AI721Contract.contract.UnpackLog(event, "RoyaltyReceiverUpdate", log); err != nil {
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
func (_AI721Contract *AI721ContractFilterer) ParseRoyaltyReceiverUpdate(log types.Log) (*AI721ContractRoyaltyReceiverUpdate, error) {
	event := new(AI721ContractRoyaltyReceiverUpdate)
	if err := _AI721Contract.contract.UnpackLog(event, "RoyaltyReceiverUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AI721ContractTopUpPoolBalanceIterator is returned from FilterTopUpPoolBalance and is used to iterate over the raw logs and unpacked data for TopUpPoolBalance events raised by the AI721Contract contract.
type AI721ContractTopUpPoolBalanceIterator struct {
	Event *AI721ContractTopUpPoolBalance // Event containing the contract specifics and raw log

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
func (it *AI721ContractTopUpPoolBalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AI721ContractTopUpPoolBalance)
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
		it.Event = new(AI721ContractTopUpPoolBalance)
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
func (it *AI721ContractTopUpPoolBalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AI721ContractTopUpPoolBalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AI721ContractTopUpPoolBalance represents a TopUpPoolBalance event raised by the AI721Contract contract.
type AI721ContractTopUpPoolBalance struct {
	AgentId *big.Int
	Caller  common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTopUpPoolBalance is a free log retrieval operation binding the contract event 0xf7ee57effd30f2ab842c1d65fdcfa7d20c2fb2f967e9ac30acafecabf013ea4c.
//
// Solidity: event TopUpPoolBalance(uint256 agentId, address caller, uint256 amount)
func (_AI721Contract *AI721ContractFilterer) FilterTopUpPoolBalance(opts *bind.FilterOpts) (*AI721ContractTopUpPoolBalanceIterator, error) {

	logs, sub, err := _AI721Contract.contract.FilterLogs(opts, "TopUpPoolBalance")
	if err != nil {
		return nil, err
	}
	return &AI721ContractTopUpPoolBalanceIterator{contract: _AI721Contract.contract, event: "TopUpPoolBalance", logs: logs, sub: sub}, nil
}

// WatchTopUpPoolBalance is a free log subscription operation binding the contract event 0xf7ee57effd30f2ab842c1d65fdcfa7d20c2fb2f967e9ac30acafecabf013ea4c.
//
// Solidity: event TopUpPoolBalance(uint256 agentId, address caller, uint256 amount)
func (_AI721Contract *AI721ContractFilterer) WatchTopUpPoolBalance(opts *bind.WatchOpts, sink chan<- *AI721ContractTopUpPoolBalance) (event.Subscription, error) {

	logs, sub, err := _AI721Contract.contract.WatchLogs(opts, "TopUpPoolBalance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AI721ContractTopUpPoolBalance)
				if err := _AI721Contract.contract.UnpackLog(event, "TopUpPoolBalance", log); err != nil {
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
func (_AI721Contract *AI721ContractFilterer) ParseTopUpPoolBalance(log types.Log) (*AI721ContractTopUpPoolBalance, error) {
	event := new(AI721ContractTopUpPoolBalance)
	if err := _AI721Contract.contract.UnpackLog(event, "TopUpPoolBalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AI721ContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AI721Contract contract.
type AI721ContractTransferIterator struct {
	Event *AI721ContractTransfer // Event containing the contract specifics and raw log

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
func (it *AI721ContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AI721ContractTransfer)
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
		it.Event = new(AI721ContractTransfer)
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
func (it *AI721ContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AI721ContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AI721ContractTransfer represents a Transfer event raised by the AI721Contract contract.
type AI721ContractTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_AI721Contract *AI721ContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*AI721ContractTransferIterator, error) {

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

	logs, sub, err := _AI721Contract.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AI721ContractTransferIterator{contract: _AI721Contract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_AI721Contract *AI721ContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AI721ContractTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _AI721Contract.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AI721ContractTransfer)
				if err := _AI721Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_AI721Contract *AI721ContractFilterer) ParseTransfer(log types.Log) (*AI721ContractTransfer, error) {
	event := new(AI721ContractTransfer)
	if err := _AI721Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AI721ContractUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AI721Contract contract.
type AI721ContractUnpausedIterator struct {
	Event *AI721ContractUnpaused // Event containing the contract specifics and raw log

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
func (it *AI721ContractUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AI721ContractUnpaused)
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
		it.Event = new(AI721ContractUnpaused)
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
func (it *AI721ContractUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AI721ContractUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AI721ContractUnpaused represents a Unpaused event raised by the AI721Contract contract.
type AI721ContractUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AI721Contract *AI721ContractFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AI721ContractUnpausedIterator, error) {

	logs, sub, err := _AI721Contract.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AI721ContractUnpausedIterator{contract: _AI721Contract.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AI721Contract *AI721ContractFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AI721ContractUnpaused) (event.Subscription, error) {

	logs, sub, err := _AI721Contract.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AI721ContractUnpaused)
				if err := _AI721Contract.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_AI721Contract *AI721ContractFilterer) ParseUnpaused(log types.Log) (*AI721ContractUnpaused, error) {
	event := new(AI721ContractUnpaused)
	if err := _AI721Contract.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
