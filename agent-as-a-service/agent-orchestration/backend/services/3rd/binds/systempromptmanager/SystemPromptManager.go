// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package systempromptmanager

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

// ISystemPromptManagerTokenMetaData is an auto generated low-level Go binding around an user-defined struct.
type ISystemPromptManagerTokenMetaData struct {
	Fee        *big.Int
	SysPrompts [][]byte
}

// SystemPromptManagerMetaData contains all meta data concerning the SystemPromptManager contract.
var SystemPromptManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyMinted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Authorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNFTData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"sysPrompt\",\"type\":\"bytes[]\"}],\"name\":\"AgentDataAddNew\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"promptIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"oldSysPrompt\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newSysPrompt\",\"type\":\"bytes\"}],\"name\":\"AgentDataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"AgentFeeUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"AgentURIUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"externalData\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"}],\"name\":\"InferencePerformed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ManagerAuthorization\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ManagerDeauthorization\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MintPriceUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"sysPrompt\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"NewToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newValue\",\"type\":\"uint16\"}],\"name\":\"RoyaltyPortionUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"RoyaltyReceiverUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TopUpPoolBalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_sysPrompt\",\"type\":\"bytes\"}],\"name\":\"addNewAgentData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"authorizeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"}],\"name\":\"dataOf\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"sysPrompts\",\"type\":\"bytes[]\"}],\"internalType\":\"structISystemPromptManager.TokenMetaData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"deauthorizeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftOwner\",\"type\":\"address\"}],\"name\":\"earnedFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"}],\"name\":\"getAgentFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"}],\"name\":\"getAgentSystemPrompt\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"getHashToSign\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_randomNonce\",\"type\":\"uint256\"}],\"name\":\"getHashToSign\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_sysPrompt\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_promptIdx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_randomNonce\",\"type\":\"uint256\"}],\"name\":\"getHashToSign\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hybridModel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"_externalData\",\"type\":\"string\"}],\"name\":\"infer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_royaltyReceiver\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_royaltyPortion\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_nextTokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_hybridModel\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_workerHub\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"mintBySignature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftOwner\",\"type\":\"address\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"name\":\"poolBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyPortion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_hybridModel\",\"type\":\"address\"}],\"name\":\"setHybridModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftOwner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"signaturesUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"}],\"name\":\"topUpPoolBalance\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_sysPrompt\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_promptIdx\",\"type\":\"uint256\"}],\"name\":\"updateAgentData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_sysPrompt\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_promptIdx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_randomNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"updateAgentDataWithSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"updateAgentFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"updateAgentURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_agentId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_randomNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"updateAgentUriWithSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"}],\"name\":\"updateMintPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_royaltyPortion\",\"type\":\"uint16\"}],\"name\":\"updateRoyaltyPortion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_royaltyReceiver\",\"type\":\"address\"}],\"name\":\"updateRoyaltyReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// SystemPromptManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemPromptManagerMetaData.ABI instead.
var SystemPromptManagerABI = SystemPromptManagerMetaData.ABI

// SystemPromptManager is an auto generated Go binding around an Ethereum contract.
type SystemPromptManager struct {
	SystemPromptManagerCaller     // Read-only binding to the contract
	SystemPromptManagerTransactor // Write-only binding to the contract
	SystemPromptManagerFilterer   // Log filterer for contract events
}

// SystemPromptManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemPromptManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemPromptManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemPromptManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemPromptManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemPromptManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemPromptManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemPromptManagerSession struct {
	Contract     *SystemPromptManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SystemPromptManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemPromptManagerCallerSession struct {
	Contract *SystemPromptManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// SystemPromptManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemPromptManagerTransactorSession struct {
	Contract     *SystemPromptManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// SystemPromptManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemPromptManagerRaw struct {
	Contract *SystemPromptManager // Generic contract binding to access the raw methods on
}

// SystemPromptManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemPromptManagerCallerRaw struct {
	Contract *SystemPromptManagerCaller // Generic read-only contract binding to access the raw methods on
}

// SystemPromptManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemPromptManagerTransactorRaw struct {
	Contract *SystemPromptManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemPromptManager creates a new instance of SystemPromptManager, bound to a specific deployed contract.
func NewSystemPromptManager(address common.Address, backend bind.ContractBackend) (*SystemPromptManager, error) {
	contract, err := bindSystemPromptManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemPromptManager{SystemPromptManagerCaller: SystemPromptManagerCaller{contract: contract}, SystemPromptManagerTransactor: SystemPromptManagerTransactor{contract: contract}, SystemPromptManagerFilterer: SystemPromptManagerFilterer{contract: contract}}, nil
}

// NewSystemPromptManagerCaller creates a new read-only instance of SystemPromptManager, bound to a specific deployed contract.
func NewSystemPromptManagerCaller(address common.Address, caller bind.ContractCaller) (*SystemPromptManagerCaller, error) {
	contract, err := bindSystemPromptManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemPromptManagerCaller{contract: contract}, nil
}

// NewSystemPromptManagerTransactor creates a new write-only instance of SystemPromptManager, bound to a specific deployed contract.
func NewSystemPromptManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemPromptManagerTransactor, error) {
	contract, err := bindSystemPromptManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemPromptManagerTransactor{contract: contract}, nil
}

// NewSystemPromptManagerFilterer creates a new log filterer instance of SystemPromptManager, bound to a specific deployed contract.
func NewSystemPromptManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemPromptManagerFilterer, error) {
	contract, err := bindSystemPromptManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemPromptManagerFilterer{contract: contract}, nil
}

// bindSystemPromptManager binds a generic wrapper to an already deployed contract.
func bindSystemPromptManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SystemPromptManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemPromptManager *SystemPromptManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemPromptManager.Contract.SystemPromptManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemPromptManager *SystemPromptManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.SystemPromptManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemPromptManager *SystemPromptManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.SystemPromptManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemPromptManager *SystemPromptManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemPromptManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemPromptManager *SystemPromptManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemPromptManager *SystemPromptManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SystemPromptManager.Contract.BalanceOf(&_SystemPromptManager.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SystemPromptManager.Contract.BalanceOf(&_SystemPromptManager.CallOpts, owner)
}

// DataOf is a free data retrieval call binding the contract method 0x871caa98.
//
// Solidity: function dataOf(uint256 _agentId) view returns((uint256,bytes[]))
func (_SystemPromptManager *SystemPromptManagerCaller) DataOf(opts *bind.CallOpts, _agentId *big.Int) (ISystemPromptManagerTokenMetaData, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "dataOf", _agentId)

	if err != nil {
		return *new(ISystemPromptManagerTokenMetaData), err
	}

	out0 := *abi.ConvertType(out[0], new(ISystemPromptManagerTokenMetaData)).(*ISystemPromptManagerTokenMetaData)

	return out0, err

}

// DataOf is a free data retrieval call binding the contract method 0x871caa98.
//
// Solidity: function dataOf(uint256 _agentId) view returns((uint256,bytes[]))
func (_SystemPromptManager *SystemPromptManagerSession) DataOf(_agentId *big.Int) (ISystemPromptManagerTokenMetaData, error) {
	return _SystemPromptManager.Contract.DataOf(&_SystemPromptManager.CallOpts, _agentId)
}

// DataOf is a free data retrieval call binding the contract method 0x871caa98.
//
// Solidity: function dataOf(uint256 _agentId) view returns((uint256,bytes[]))
func (_SystemPromptManager *SystemPromptManagerCallerSession) DataOf(_agentId *big.Int) (ISystemPromptManagerTokenMetaData, error) {
	return _SystemPromptManager.Contract.DataOf(&_SystemPromptManager.CallOpts, _agentId)
}

// EarnedFees is a free data retrieval call binding the contract method 0xfeb7219d.
//
// Solidity: function earnedFees(address nftOwner) view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerCaller) EarnedFees(opts *bind.CallOpts, nftOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "earnedFees", nftOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EarnedFees is a free data retrieval call binding the contract method 0xfeb7219d.
//
// Solidity: function earnedFees(address nftOwner) view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerSession) EarnedFees(nftOwner common.Address) (*big.Int, error) {
	return _SystemPromptManager.Contract.EarnedFees(&_SystemPromptManager.CallOpts, nftOwner)
}

// EarnedFees is a free data retrieval call binding the contract method 0xfeb7219d.
//
// Solidity: function earnedFees(address nftOwner) view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerCallerSession) EarnedFees(nftOwner common.Address) (*big.Int, error) {
	return _SystemPromptManager.Contract.EarnedFees(&_SystemPromptManager.CallOpts, nftOwner)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_SystemPromptManager *SystemPromptManagerCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "eip712Domain")

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
func (_SystemPromptManager *SystemPromptManagerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _SystemPromptManager.Contract.Eip712Domain(&_SystemPromptManager.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_SystemPromptManager *SystemPromptManagerCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _SystemPromptManager.Contract.Eip712Domain(&_SystemPromptManager.CallOpts)
}

// GetAgentFee is a free data retrieval call binding the contract method 0xed96f433.
//
// Solidity: function getAgentFee(uint256 _agentId) view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerCaller) GetAgentFee(opts *bind.CallOpts, _agentId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "getAgentFee", _agentId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAgentFee is a free data retrieval call binding the contract method 0xed96f433.
//
// Solidity: function getAgentFee(uint256 _agentId) view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerSession) GetAgentFee(_agentId *big.Int) (*big.Int, error) {
	return _SystemPromptManager.Contract.GetAgentFee(&_SystemPromptManager.CallOpts, _agentId)
}

// GetAgentFee is a free data retrieval call binding the contract method 0xed96f433.
//
// Solidity: function getAgentFee(uint256 _agentId) view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerCallerSession) GetAgentFee(_agentId *big.Int) (*big.Int, error) {
	return _SystemPromptManager.Contract.GetAgentFee(&_SystemPromptManager.CallOpts, _agentId)
}

// GetAgentSystemPrompt is a free data retrieval call binding the contract method 0xf325f3f1.
//
// Solidity: function getAgentSystemPrompt(uint256 _agentId) view returns(bytes[])
func (_SystemPromptManager *SystemPromptManagerCaller) GetAgentSystemPrompt(opts *bind.CallOpts, _agentId *big.Int) ([][]byte, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "getAgentSystemPrompt", _agentId)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetAgentSystemPrompt is a free data retrieval call binding the contract method 0xf325f3f1.
//
// Solidity: function getAgentSystemPrompt(uint256 _agentId) view returns(bytes[])
func (_SystemPromptManager *SystemPromptManagerSession) GetAgentSystemPrompt(_agentId *big.Int) ([][]byte, error) {
	return _SystemPromptManager.Contract.GetAgentSystemPrompt(&_SystemPromptManager.CallOpts, _agentId)
}

// GetAgentSystemPrompt is a free data retrieval call binding the contract method 0xf325f3f1.
//
// Solidity: function getAgentSystemPrompt(uint256 _agentId) view returns(bytes[])
func (_SystemPromptManager *SystemPromptManagerCallerSession) GetAgentSystemPrompt(_agentId *big.Int) ([][]byte, error) {
	return _SystemPromptManager.Contract.GetAgentSystemPrompt(&_SystemPromptManager.CallOpts, _agentId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SystemPromptManager *SystemPromptManagerCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SystemPromptManager *SystemPromptManagerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SystemPromptManager.Contract.GetApproved(&_SystemPromptManager.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SystemPromptManager *SystemPromptManagerCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SystemPromptManager.Contract.GetApproved(&_SystemPromptManager.CallOpts, tokenId)
}

// GetHashToSign is a free data retrieval call binding the contract method 0x68cfcc9a.
//
// Solidity: function getHashToSign(address _to, string _uri, bytes _data, uint256 _fee, address _manager) view returns(bytes32)
func (_SystemPromptManager *SystemPromptManagerCaller) GetHashToSign(opts *bind.CallOpts, _to common.Address, _uri string, _data []byte, _fee *big.Int, _manager common.Address) ([32]byte, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "getHashToSign", _to, _uri, _data, _fee, _manager)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHashToSign is a free data retrieval call binding the contract method 0x68cfcc9a.
//
// Solidity: function getHashToSign(address _to, string _uri, bytes _data, uint256 _fee, address _manager) view returns(bytes32)
func (_SystemPromptManager *SystemPromptManagerSession) GetHashToSign(_to common.Address, _uri string, _data []byte, _fee *big.Int, _manager common.Address) ([32]byte, error) {
	return _SystemPromptManager.Contract.GetHashToSign(&_SystemPromptManager.CallOpts, _to, _uri, _data, _fee, _manager)
}

// GetHashToSign is a free data retrieval call binding the contract method 0x68cfcc9a.
//
// Solidity: function getHashToSign(address _to, string _uri, bytes _data, uint256 _fee, address _manager) view returns(bytes32)
func (_SystemPromptManager *SystemPromptManagerCallerSession) GetHashToSign(_to common.Address, _uri string, _data []byte, _fee *big.Int, _manager common.Address) ([32]byte, error) {
	return _SystemPromptManager.Contract.GetHashToSign(&_SystemPromptManager.CallOpts, _to, _uri, _data, _fee, _manager)
}

// GetHashToSign0 is a free data retrieval call binding the contract method 0x937a91d9.
//
// Solidity: function getHashToSign(uint256 _agentId, string _uri, uint256 _randomNonce) view returns(bytes32)
func (_SystemPromptManager *SystemPromptManagerCaller) GetHashToSign0(opts *bind.CallOpts, _agentId *big.Int, _uri string, _randomNonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "getHashToSign0", _agentId, _uri, _randomNonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHashToSign0 is a free data retrieval call binding the contract method 0x937a91d9.
//
// Solidity: function getHashToSign(uint256 _agentId, string _uri, uint256 _randomNonce) view returns(bytes32)
func (_SystemPromptManager *SystemPromptManagerSession) GetHashToSign0(_agentId *big.Int, _uri string, _randomNonce *big.Int) ([32]byte, error) {
	return _SystemPromptManager.Contract.GetHashToSign0(&_SystemPromptManager.CallOpts, _agentId, _uri, _randomNonce)
}

// GetHashToSign0 is a free data retrieval call binding the contract method 0x937a91d9.
//
// Solidity: function getHashToSign(uint256 _agentId, string _uri, uint256 _randomNonce) view returns(bytes32)
func (_SystemPromptManager *SystemPromptManagerCallerSession) GetHashToSign0(_agentId *big.Int, _uri string, _randomNonce *big.Int) ([32]byte, error) {
	return _SystemPromptManager.Contract.GetHashToSign0(&_SystemPromptManager.CallOpts, _agentId, _uri, _randomNonce)
}

// GetHashToSign1 is a free data retrieval call binding the contract method 0xe340d79f.
//
// Solidity: function getHashToSign(uint256 _agentId, bytes _sysPrompt, uint256 _promptIdx, uint256 _randomNonce) view returns(bytes32)
func (_SystemPromptManager *SystemPromptManagerCaller) GetHashToSign1(opts *bind.CallOpts, _agentId *big.Int, _sysPrompt []byte, _promptIdx *big.Int, _randomNonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "getHashToSign1", _agentId, _sysPrompt, _promptIdx, _randomNonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHashToSign1 is a free data retrieval call binding the contract method 0xe340d79f.
//
// Solidity: function getHashToSign(uint256 _agentId, bytes _sysPrompt, uint256 _promptIdx, uint256 _randomNonce) view returns(bytes32)
func (_SystemPromptManager *SystemPromptManagerSession) GetHashToSign1(_agentId *big.Int, _sysPrompt []byte, _promptIdx *big.Int, _randomNonce *big.Int) ([32]byte, error) {
	return _SystemPromptManager.Contract.GetHashToSign1(&_SystemPromptManager.CallOpts, _agentId, _sysPrompt, _promptIdx, _randomNonce)
}

// GetHashToSign1 is a free data retrieval call binding the contract method 0xe340d79f.
//
// Solidity: function getHashToSign(uint256 _agentId, bytes _sysPrompt, uint256 _promptIdx, uint256 _randomNonce) view returns(bytes32)
func (_SystemPromptManager *SystemPromptManagerCallerSession) GetHashToSign1(_agentId *big.Int, _sysPrompt []byte, _promptIdx *big.Int, _randomNonce *big.Int) ([32]byte, error) {
	return _SystemPromptManager.Contract.GetHashToSign1(&_SystemPromptManager.CallOpts, _agentId, _sysPrompt, _promptIdx, _randomNonce)
}

// HybridModel is a free data retrieval call binding the contract method 0x5eb2364c.
//
// Solidity: function hybridModel() view returns(address)
func (_SystemPromptManager *SystemPromptManagerCaller) HybridModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "hybridModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HybridModel is a free data retrieval call binding the contract method 0x5eb2364c.
//
// Solidity: function hybridModel() view returns(address)
func (_SystemPromptManager *SystemPromptManagerSession) HybridModel() (common.Address, error) {
	return _SystemPromptManager.Contract.HybridModel(&_SystemPromptManager.CallOpts)
}

// HybridModel is a free data retrieval call binding the contract method 0x5eb2364c.
//
// Solidity: function hybridModel() view returns(address)
func (_SystemPromptManager *SystemPromptManagerCallerSession) HybridModel() (common.Address, error) {
	return _SystemPromptManager.Contract.HybridModel(&_SystemPromptManager.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SystemPromptManager *SystemPromptManagerCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SystemPromptManager *SystemPromptManagerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SystemPromptManager.Contract.IsApprovedForAll(&_SystemPromptManager.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SystemPromptManager *SystemPromptManagerCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SystemPromptManager.Contract.IsApprovedForAll(&_SystemPromptManager.CallOpts, owner, operator)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_SystemPromptManager *SystemPromptManagerCaller) IsManager(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "isManager", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_SystemPromptManager *SystemPromptManagerSession) IsManager(arg0 common.Address) (bool, error) {
	return _SystemPromptManager.Contract.IsManager(&_SystemPromptManager.CallOpts, arg0)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_SystemPromptManager *SystemPromptManagerCallerSession) IsManager(arg0 common.Address) (bool, error) {
	return _SystemPromptManager.Contract.IsManager(&_SystemPromptManager.CallOpts, arg0)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerCaller) MintPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "mintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerSession) MintPrice() (*big.Int, error) {
	return _SystemPromptManager.Contract.MintPrice(&_SystemPromptManager.CallOpts)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerCallerSession) MintPrice() (*big.Int, error) {
	return _SystemPromptManager.Contract.MintPrice(&_SystemPromptManager.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SystemPromptManager *SystemPromptManagerCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SystemPromptManager *SystemPromptManagerSession) Name() (string, error) {
	return _SystemPromptManager.Contract.Name(&_SystemPromptManager.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SystemPromptManager *SystemPromptManagerCallerSession) Name() (string, error) {
	return _SystemPromptManager.Contract.Name(&_SystemPromptManager.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerCaller) NextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "nextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerSession) NextTokenId() (*big.Int, error) {
	return _SystemPromptManager.Contract.NextTokenId(&_SystemPromptManager.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerCallerSession) NextTokenId() (*big.Int, error) {
	return _SystemPromptManager.Contract.NextTokenId(&_SystemPromptManager.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address nftOwner) view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerCaller) Nonce(opts *bind.CallOpts, nftOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "nonce", nftOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address nftOwner) view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerSession) Nonce(nftOwner common.Address) (*big.Int, error) {
	return _SystemPromptManager.Contract.Nonce(&_SystemPromptManager.CallOpts, nftOwner)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address nftOwner) view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerCallerSession) Nonce(nftOwner common.Address) (*big.Int, error) {
	return _SystemPromptManager.Contract.Nonce(&_SystemPromptManager.CallOpts, nftOwner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemPromptManager *SystemPromptManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemPromptManager *SystemPromptManagerSession) Owner() (common.Address, error) {
	return _SystemPromptManager.Contract.Owner(&_SystemPromptManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemPromptManager *SystemPromptManagerCallerSession) Owner() (common.Address, error) {
	return _SystemPromptManager.Contract.Owner(&_SystemPromptManager.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SystemPromptManager *SystemPromptManagerCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SystemPromptManager *SystemPromptManagerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SystemPromptManager.Contract.OwnerOf(&_SystemPromptManager.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SystemPromptManager *SystemPromptManagerCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SystemPromptManager.Contract.OwnerOf(&_SystemPromptManager.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SystemPromptManager *SystemPromptManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SystemPromptManager *SystemPromptManagerSession) Paused() (bool, error) {
	return _SystemPromptManager.Contract.Paused(&_SystemPromptManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SystemPromptManager *SystemPromptManagerCallerSession) Paused() (bool, error) {
	return _SystemPromptManager.Contract.Paused(&_SystemPromptManager.CallOpts)
}

// PoolBalance is a free data retrieval call binding the contract method 0x6a6d964e.
//
// Solidity: function poolBalance(uint256 nftId) view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerCaller) PoolBalance(opts *bind.CallOpts, nftId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "poolBalance", nftId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolBalance is a free data retrieval call binding the contract method 0x6a6d964e.
//
// Solidity: function poolBalance(uint256 nftId) view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerSession) PoolBalance(nftId *big.Int) (*big.Int, error) {
	return _SystemPromptManager.Contract.PoolBalance(&_SystemPromptManager.CallOpts, nftId)
}

// PoolBalance is a free data retrieval call binding the contract method 0x6a6d964e.
//
// Solidity: function poolBalance(uint256 nftId) view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerCallerSession) PoolBalance(nftId *big.Int) (*big.Int, error) {
	return _SystemPromptManager.Contract.PoolBalance(&_SystemPromptManager.CallOpts, nftId)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _agentId, uint256 _salePrice) view returns(address, uint256)
func (_SystemPromptManager *SystemPromptManagerCaller) RoyaltyInfo(opts *bind.CallOpts, _agentId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "royaltyInfo", _agentId, _salePrice)

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
func (_SystemPromptManager *SystemPromptManagerSession) RoyaltyInfo(_agentId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _SystemPromptManager.Contract.RoyaltyInfo(&_SystemPromptManager.CallOpts, _agentId, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _agentId, uint256 _salePrice) view returns(address, uint256)
func (_SystemPromptManager *SystemPromptManagerCallerSession) RoyaltyInfo(_agentId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _SystemPromptManager.Contract.RoyaltyInfo(&_SystemPromptManager.CallOpts, _agentId, _salePrice)
}

// RoyaltyPortion is a free data retrieval call binding the contract method 0x11d7beb2.
//
// Solidity: function royaltyPortion() view returns(uint16)
func (_SystemPromptManager *SystemPromptManagerCaller) RoyaltyPortion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "royaltyPortion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RoyaltyPortion is a free data retrieval call binding the contract method 0x11d7beb2.
//
// Solidity: function royaltyPortion() view returns(uint16)
func (_SystemPromptManager *SystemPromptManagerSession) RoyaltyPortion() (uint16, error) {
	return _SystemPromptManager.Contract.RoyaltyPortion(&_SystemPromptManager.CallOpts)
}

// RoyaltyPortion is a free data retrieval call binding the contract method 0x11d7beb2.
//
// Solidity: function royaltyPortion() view returns(uint16)
func (_SystemPromptManager *SystemPromptManagerCallerSession) RoyaltyPortion() (uint16, error) {
	return _SystemPromptManager.Contract.RoyaltyPortion(&_SystemPromptManager.CallOpts)
}

// RoyaltyReceiver is a free data retrieval call binding the contract method 0x9fbc8713.
//
// Solidity: function royaltyReceiver() view returns(address)
func (_SystemPromptManager *SystemPromptManagerCaller) RoyaltyReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "royaltyReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoyaltyReceiver is a free data retrieval call binding the contract method 0x9fbc8713.
//
// Solidity: function royaltyReceiver() view returns(address)
func (_SystemPromptManager *SystemPromptManagerSession) RoyaltyReceiver() (common.Address, error) {
	return _SystemPromptManager.Contract.RoyaltyReceiver(&_SystemPromptManager.CallOpts)
}

// RoyaltyReceiver is a free data retrieval call binding the contract method 0x9fbc8713.
//
// Solidity: function royaltyReceiver() view returns(address)
func (_SystemPromptManager *SystemPromptManagerCallerSession) RoyaltyReceiver() (common.Address, error) {
	return _SystemPromptManager.Contract.RoyaltyReceiver(&_SystemPromptManager.CallOpts)
}

// SignaturesUsed is a free data retrieval call binding the contract method 0x757d513b.
//
// Solidity: function signaturesUsed(address nftOwner, bytes signature) view returns(bool)
func (_SystemPromptManager *SystemPromptManagerCaller) SignaturesUsed(opts *bind.CallOpts, nftOwner common.Address, signature []byte) (bool, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "signaturesUsed", nftOwner, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SignaturesUsed is a free data retrieval call binding the contract method 0x757d513b.
//
// Solidity: function signaturesUsed(address nftOwner, bytes signature) view returns(bool)
func (_SystemPromptManager *SystemPromptManagerSession) SignaturesUsed(nftOwner common.Address, signature []byte) (bool, error) {
	return _SystemPromptManager.Contract.SignaturesUsed(&_SystemPromptManager.CallOpts, nftOwner, signature)
}

// SignaturesUsed is a free data retrieval call binding the contract method 0x757d513b.
//
// Solidity: function signaturesUsed(address nftOwner, bytes signature) view returns(bool)
func (_SystemPromptManager *SystemPromptManagerCallerSession) SignaturesUsed(nftOwner common.Address, signature []byte) (bool, error) {
	return _SystemPromptManager.Contract.SignaturesUsed(&_SystemPromptManager.CallOpts, nftOwner, signature)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_SystemPromptManager *SystemPromptManagerCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_SystemPromptManager *SystemPromptManagerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _SystemPromptManager.Contract.SupportsInterface(&_SystemPromptManager.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_SystemPromptManager *SystemPromptManagerCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _SystemPromptManager.Contract.SupportsInterface(&_SystemPromptManager.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SystemPromptManager *SystemPromptManagerCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SystemPromptManager *SystemPromptManagerSession) Symbol() (string, error) {
	return _SystemPromptManager.Contract.Symbol(&_SystemPromptManager.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SystemPromptManager *SystemPromptManagerCallerSession) Symbol() (string, error) {
	return _SystemPromptManager.Contract.Symbol(&_SystemPromptManager.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _SystemPromptManager.Contract.TokenByIndex(&_SystemPromptManager.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _SystemPromptManager.Contract.TokenByIndex(&_SystemPromptManager.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _SystemPromptManager.Contract.TokenOfOwnerByIndex(&_SystemPromptManager.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _SystemPromptManager.Contract.TokenOfOwnerByIndex(&_SystemPromptManager.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _agentId) view returns(string)
func (_SystemPromptManager *SystemPromptManagerCaller) TokenURI(opts *bind.CallOpts, _agentId *big.Int) (string, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "tokenURI", _agentId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _agentId) view returns(string)
func (_SystemPromptManager *SystemPromptManagerSession) TokenURI(_agentId *big.Int) (string, error) {
	return _SystemPromptManager.Contract.TokenURI(&_SystemPromptManager.CallOpts, _agentId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _agentId) view returns(string)
func (_SystemPromptManager *SystemPromptManagerCallerSession) TokenURI(_agentId *big.Int) (string, error) {
	return _SystemPromptManager.Contract.TokenURI(&_SystemPromptManager.CallOpts, _agentId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerSession) TotalSupply() (*big.Int, error) {
	return _SystemPromptManager.Contract.TotalSupply(&_SystemPromptManager.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SystemPromptManager *SystemPromptManagerCallerSession) TotalSupply() (*big.Int, error) {
	return _SystemPromptManager.Contract.TotalSupply(&_SystemPromptManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_SystemPromptManager *SystemPromptManagerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SystemPromptManager.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_SystemPromptManager *SystemPromptManagerSession) Version() (string, error) {
	return _SystemPromptManager.Contract.Version(&_SystemPromptManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_SystemPromptManager *SystemPromptManagerCallerSession) Version() (string, error) {
	return _SystemPromptManager.Contract.Version(&_SystemPromptManager.CallOpts)
}

// AddNewAgentData is a paid mutator transaction binding the contract method 0x9f10cc66.
//
// Solidity: function addNewAgentData(uint256 _agentId, bytes _sysPrompt) returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) AddNewAgentData(opts *bind.TransactOpts, _agentId *big.Int, _sysPrompt []byte) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "addNewAgentData", _agentId, _sysPrompt)
}

// AddNewAgentData is a paid mutator transaction binding the contract method 0x9f10cc66.
//
// Solidity: function addNewAgentData(uint256 _agentId, bytes _sysPrompt) returns()
func (_SystemPromptManager *SystemPromptManagerSession) AddNewAgentData(_agentId *big.Int, _sysPrompt []byte) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.AddNewAgentData(&_SystemPromptManager.TransactOpts, _agentId, _sysPrompt)
}

// AddNewAgentData is a paid mutator transaction binding the contract method 0x9f10cc66.
//
// Solidity: function addNewAgentData(uint256 _agentId, bytes _sysPrompt) returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) AddNewAgentData(_agentId *big.Int, _sysPrompt []byte) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.AddNewAgentData(&_SystemPromptManager.TransactOpts, _agentId, _sysPrompt)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SystemPromptManager *SystemPromptManagerSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.Approve(&_SystemPromptManager.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.Approve(&_SystemPromptManager.TransactOpts, to, tokenId)
}

// AuthorizeManager is a paid mutator transaction binding the contract method 0x267c8507.
//
// Solidity: function authorizeManager(address _account) returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) AuthorizeManager(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "authorizeManager", _account)
}

// AuthorizeManager is a paid mutator transaction binding the contract method 0x267c8507.
//
// Solidity: function authorizeManager(address _account) returns()
func (_SystemPromptManager *SystemPromptManagerSession) AuthorizeManager(_account common.Address) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.AuthorizeManager(&_SystemPromptManager.TransactOpts, _account)
}

// AuthorizeManager is a paid mutator transaction binding the contract method 0x267c8507.
//
// Solidity: function authorizeManager(address _account) returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) AuthorizeManager(_account common.Address) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.AuthorizeManager(&_SystemPromptManager.TransactOpts, _account)
}

// ClaimFee is a paid mutator transaction binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) ClaimFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "claimFee")
}

// ClaimFee is a paid mutator transaction binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() returns()
func (_SystemPromptManager *SystemPromptManagerSession) ClaimFee() (*types.Transaction, error) {
	return _SystemPromptManager.Contract.ClaimFee(&_SystemPromptManager.TransactOpts)
}

// ClaimFee is a paid mutator transaction binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) ClaimFee() (*types.Transaction, error) {
	return _SystemPromptManager.Contract.ClaimFee(&_SystemPromptManager.TransactOpts)
}

// DeauthorizeManager is a paid mutator transaction binding the contract method 0x0305ea01.
//
// Solidity: function deauthorizeManager(address _account) returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) DeauthorizeManager(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "deauthorizeManager", _account)
}

// DeauthorizeManager is a paid mutator transaction binding the contract method 0x0305ea01.
//
// Solidity: function deauthorizeManager(address _account) returns()
func (_SystemPromptManager *SystemPromptManagerSession) DeauthorizeManager(_account common.Address) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.DeauthorizeManager(&_SystemPromptManager.TransactOpts, _account)
}

// DeauthorizeManager is a paid mutator transaction binding the contract method 0x0305ea01.
//
// Solidity: function deauthorizeManager(address _account) returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) DeauthorizeManager(_account common.Address) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.DeauthorizeManager(&_SystemPromptManager.TransactOpts, _account)
}

// Infer is a paid mutator transaction binding the contract method 0x566a9951.
//
// Solidity: function infer(uint256 _agentId, bytes _calldata, string _externalData) payable returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) Infer(opts *bind.TransactOpts, _agentId *big.Int, _calldata []byte, _externalData string) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "infer", _agentId, _calldata, _externalData)
}

// Infer is a paid mutator transaction binding the contract method 0x566a9951.
//
// Solidity: function infer(uint256 _agentId, bytes _calldata, string _externalData) payable returns()
func (_SystemPromptManager *SystemPromptManagerSession) Infer(_agentId *big.Int, _calldata []byte, _externalData string) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.Infer(&_SystemPromptManager.TransactOpts, _agentId, _calldata, _externalData)
}

// Infer is a paid mutator transaction binding the contract method 0x566a9951.
//
// Solidity: function infer(uint256 _agentId, bytes _calldata, string _externalData) payable returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) Infer(_agentId *big.Int, _calldata []byte, _externalData string) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.Infer(&_SystemPromptManager.TransactOpts, _agentId, _calldata, _externalData)
}

// Initialize is a paid mutator transaction binding the contract method 0x1e18fb8c.
//
// Solidity: function initialize(string _name, string _symbol, uint256 _mintPrice, address _royaltyReceiver, uint16 _royaltyPortion, uint256 _nextTokenId, address _hybridModel, address _workerHub) returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _mintPrice *big.Int, _royaltyReceiver common.Address, _royaltyPortion uint16, _nextTokenId *big.Int, _hybridModel common.Address, _workerHub common.Address) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "initialize", _name, _symbol, _mintPrice, _royaltyReceiver, _royaltyPortion, _nextTokenId, _hybridModel, _workerHub)
}

// Initialize is a paid mutator transaction binding the contract method 0x1e18fb8c.
//
// Solidity: function initialize(string _name, string _symbol, uint256 _mintPrice, address _royaltyReceiver, uint16 _royaltyPortion, uint256 _nextTokenId, address _hybridModel, address _workerHub) returns()
func (_SystemPromptManager *SystemPromptManagerSession) Initialize(_name string, _symbol string, _mintPrice *big.Int, _royaltyReceiver common.Address, _royaltyPortion uint16, _nextTokenId *big.Int, _hybridModel common.Address, _workerHub common.Address) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.Initialize(&_SystemPromptManager.TransactOpts, _name, _symbol, _mintPrice, _royaltyReceiver, _royaltyPortion, _nextTokenId, _hybridModel, _workerHub)
}

// Initialize is a paid mutator transaction binding the contract method 0x1e18fb8c.
//
// Solidity: function initialize(string _name, string _symbol, uint256 _mintPrice, address _royaltyReceiver, uint16 _royaltyPortion, uint256 _nextTokenId, address _hybridModel, address _workerHub) returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) Initialize(_name string, _symbol string, _mintPrice *big.Int, _royaltyReceiver common.Address, _royaltyPortion uint16, _nextTokenId *big.Int, _hybridModel common.Address, _workerHub common.Address) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.Initialize(&_SystemPromptManager.TransactOpts, _name, _symbol, _mintPrice, _royaltyReceiver, _royaltyPortion, _nextTokenId, _hybridModel, _workerHub)
}

// Mint is a paid mutator transaction binding the contract method 0xcc216aca.
//
// Solidity: function mint(address _to, string _uri, bytes _data, uint256 _fee) payable returns(uint256)
func (_SystemPromptManager *SystemPromptManagerTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _uri string, _data []byte, _fee *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "mint", _to, _uri, _data, _fee)
}

// Mint is a paid mutator transaction binding the contract method 0xcc216aca.
//
// Solidity: function mint(address _to, string _uri, bytes _data, uint256 _fee) payable returns(uint256)
func (_SystemPromptManager *SystemPromptManagerSession) Mint(_to common.Address, _uri string, _data []byte, _fee *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.Mint(&_SystemPromptManager.TransactOpts, _to, _uri, _data, _fee)
}

// Mint is a paid mutator transaction binding the contract method 0xcc216aca.
//
// Solidity: function mint(address _to, string _uri, bytes _data, uint256 _fee) payable returns(uint256)
func (_SystemPromptManager *SystemPromptManagerTransactorSession) Mint(_to common.Address, _uri string, _data []byte, _fee *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.Mint(&_SystemPromptManager.TransactOpts, _to, _uri, _data, _fee)
}

// MintBySignature is a paid mutator transaction binding the contract method 0x53ec32b7.
//
// Solidity: function mintBySignature(address _to, string _uri, bytes _data, uint256 _fee, address _manager, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_SystemPromptManager *SystemPromptManagerTransactor) MintBySignature(opts *bind.TransactOpts, _to common.Address, _uri string, _data []byte, _fee *big.Int, _manager common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "mintBySignature", _to, _uri, _data, _fee, _manager, v, r, s)
}

// MintBySignature is a paid mutator transaction binding the contract method 0x53ec32b7.
//
// Solidity: function mintBySignature(address _to, string _uri, bytes _data, uint256 _fee, address _manager, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_SystemPromptManager *SystemPromptManagerSession) MintBySignature(_to common.Address, _uri string, _data []byte, _fee *big.Int, _manager common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.MintBySignature(&_SystemPromptManager.TransactOpts, _to, _uri, _data, _fee, _manager, v, r, s)
}

// MintBySignature is a paid mutator transaction binding the contract method 0x53ec32b7.
//
// Solidity: function mintBySignature(address _to, string _uri, bytes _data, uint256 _fee, address _manager, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_SystemPromptManager *SystemPromptManagerTransactorSession) MintBySignature(_to common.Address, _uri string, _data []byte, _fee *big.Int, _manager common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.MintBySignature(&_SystemPromptManager.TransactOpts, _to, _uri, _data, _fee, _manager, v, r, s)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SystemPromptManager *SystemPromptManagerSession) Pause() (*types.Transaction, error) {
	return _SystemPromptManager.Contract.Pause(&_SystemPromptManager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) Pause() (*types.Transaction, error) {
	return _SystemPromptManager.Contract.Pause(&_SystemPromptManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemPromptManager *SystemPromptManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _SystemPromptManager.Contract.RenounceOwnership(&_SystemPromptManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SystemPromptManager.Contract.RenounceOwnership(&_SystemPromptManager.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SystemPromptManager *SystemPromptManagerSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.SafeTransferFrom(&_SystemPromptManager.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.SafeTransferFrom(&_SystemPromptManager.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_SystemPromptManager *SystemPromptManagerSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.SafeTransferFrom0(&_SystemPromptManager.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.SafeTransferFrom0(&_SystemPromptManager.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SystemPromptManager *SystemPromptManagerSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.SetApprovalForAll(&_SystemPromptManager.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.SetApprovalForAll(&_SystemPromptManager.TransactOpts, operator, approved)
}

// SetHybridModel is a paid mutator transaction binding the contract method 0xe645f296.
//
// Solidity: function setHybridModel(address _hybridModel) returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) SetHybridModel(opts *bind.TransactOpts, _hybridModel common.Address) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "setHybridModel", _hybridModel)
}

// SetHybridModel is a paid mutator transaction binding the contract method 0xe645f296.
//
// Solidity: function setHybridModel(address _hybridModel) returns()
func (_SystemPromptManager *SystemPromptManagerSession) SetHybridModel(_hybridModel common.Address) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.SetHybridModel(&_SystemPromptManager.TransactOpts, _hybridModel)
}

// SetHybridModel is a paid mutator transaction binding the contract method 0xe645f296.
//
// Solidity: function setHybridModel(address _hybridModel) returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) SetHybridModel(_hybridModel common.Address) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.SetHybridModel(&_SystemPromptManager.TransactOpts, _hybridModel)
}

// TopUpPoolBalance is a paid mutator transaction binding the contract method 0xc82cd94b.
//
// Solidity: function topUpPoolBalance(uint256 _agentId) payable returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) TopUpPoolBalance(opts *bind.TransactOpts, _agentId *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "topUpPoolBalance", _agentId)
}

// TopUpPoolBalance is a paid mutator transaction binding the contract method 0xc82cd94b.
//
// Solidity: function topUpPoolBalance(uint256 _agentId) payable returns()
func (_SystemPromptManager *SystemPromptManagerSession) TopUpPoolBalance(_agentId *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.TopUpPoolBalance(&_SystemPromptManager.TransactOpts, _agentId)
}

// TopUpPoolBalance is a paid mutator transaction binding the contract method 0xc82cd94b.
//
// Solidity: function topUpPoolBalance(uint256 _agentId) payable returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) TopUpPoolBalance(_agentId *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.TopUpPoolBalance(&_SystemPromptManager.TransactOpts, _agentId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SystemPromptManager *SystemPromptManagerSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.TransferFrom(&_SystemPromptManager.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.TransferFrom(&_SystemPromptManager.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemPromptManager *SystemPromptManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.TransferOwnership(&_SystemPromptManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.TransferOwnership(&_SystemPromptManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SystemPromptManager *SystemPromptManagerSession) Unpause() (*types.Transaction, error) {
	return _SystemPromptManager.Contract.Unpause(&_SystemPromptManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _SystemPromptManager.Contract.Unpause(&_SystemPromptManager.TransactOpts)
}

// UpdateAgentData is a paid mutator transaction binding the contract method 0xed82c9e0.
//
// Solidity: function updateAgentData(uint256 _agentId, bytes _sysPrompt, uint256 _promptIdx) returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) UpdateAgentData(opts *bind.TransactOpts, _agentId *big.Int, _sysPrompt []byte, _promptIdx *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "updateAgentData", _agentId, _sysPrompt, _promptIdx)
}

// UpdateAgentData is a paid mutator transaction binding the contract method 0xed82c9e0.
//
// Solidity: function updateAgentData(uint256 _agentId, bytes _sysPrompt, uint256 _promptIdx) returns()
func (_SystemPromptManager *SystemPromptManagerSession) UpdateAgentData(_agentId *big.Int, _sysPrompt []byte, _promptIdx *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.UpdateAgentData(&_SystemPromptManager.TransactOpts, _agentId, _sysPrompt, _promptIdx)
}

// UpdateAgentData is a paid mutator transaction binding the contract method 0xed82c9e0.
//
// Solidity: function updateAgentData(uint256 _agentId, bytes _sysPrompt, uint256 _promptIdx) returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) UpdateAgentData(_agentId *big.Int, _sysPrompt []byte, _promptIdx *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.UpdateAgentData(&_SystemPromptManager.TransactOpts, _agentId, _sysPrompt, _promptIdx)
}

// UpdateAgentDataWithSignature is a paid mutator transaction binding the contract method 0xb8a49a57.
//
// Solidity: function updateAgentDataWithSignature(uint256 _agentId, bytes _sysPrompt, uint256 _promptIdx, uint256 _randomNonce, bytes _signature) returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) UpdateAgentDataWithSignature(opts *bind.TransactOpts, _agentId *big.Int, _sysPrompt []byte, _promptIdx *big.Int, _randomNonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "updateAgentDataWithSignature", _agentId, _sysPrompt, _promptIdx, _randomNonce, _signature)
}

// UpdateAgentDataWithSignature is a paid mutator transaction binding the contract method 0xb8a49a57.
//
// Solidity: function updateAgentDataWithSignature(uint256 _agentId, bytes _sysPrompt, uint256 _promptIdx, uint256 _randomNonce, bytes _signature) returns()
func (_SystemPromptManager *SystemPromptManagerSession) UpdateAgentDataWithSignature(_agentId *big.Int, _sysPrompt []byte, _promptIdx *big.Int, _randomNonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.UpdateAgentDataWithSignature(&_SystemPromptManager.TransactOpts, _agentId, _sysPrompt, _promptIdx, _randomNonce, _signature)
}

// UpdateAgentDataWithSignature is a paid mutator transaction binding the contract method 0xb8a49a57.
//
// Solidity: function updateAgentDataWithSignature(uint256 _agentId, bytes _sysPrompt, uint256 _promptIdx, uint256 _randomNonce, bytes _signature) returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) UpdateAgentDataWithSignature(_agentId *big.Int, _sysPrompt []byte, _promptIdx *big.Int, _randomNonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.UpdateAgentDataWithSignature(&_SystemPromptManager.TransactOpts, _agentId, _sysPrompt, _promptIdx, _randomNonce, _signature)
}

// UpdateAgentFee is a paid mutator transaction binding the contract method 0xb1fd1526.
//
// Solidity: function updateAgentFee(uint256 _agentId, uint256 _fee) returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) UpdateAgentFee(opts *bind.TransactOpts, _agentId *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "updateAgentFee", _agentId, _fee)
}

// UpdateAgentFee is a paid mutator transaction binding the contract method 0xb1fd1526.
//
// Solidity: function updateAgentFee(uint256 _agentId, uint256 _fee) returns()
func (_SystemPromptManager *SystemPromptManagerSession) UpdateAgentFee(_agentId *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.UpdateAgentFee(&_SystemPromptManager.TransactOpts, _agentId, _fee)
}

// UpdateAgentFee is a paid mutator transaction binding the contract method 0xb1fd1526.
//
// Solidity: function updateAgentFee(uint256 _agentId, uint256 _fee) returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) UpdateAgentFee(_agentId *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.UpdateAgentFee(&_SystemPromptManager.TransactOpts, _agentId, _fee)
}

// UpdateAgentURI is a paid mutator transaction binding the contract method 0x6b595822.
//
// Solidity: function updateAgentURI(uint256 _agentId, string _uri) returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) UpdateAgentURI(opts *bind.TransactOpts, _agentId *big.Int, _uri string) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "updateAgentURI", _agentId, _uri)
}

// UpdateAgentURI is a paid mutator transaction binding the contract method 0x6b595822.
//
// Solidity: function updateAgentURI(uint256 _agentId, string _uri) returns()
func (_SystemPromptManager *SystemPromptManagerSession) UpdateAgentURI(_agentId *big.Int, _uri string) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.UpdateAgentURI(&_SystemPromptManager.TransactOpts, _agentId, _uri)
}

// UpdateAgentURI is a paid mutator transaction binding the contract method 0x6b595822.
//
// Solidity: function updateAgentURI(uint256 _agentId, string _uri) returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) UpdateAgentURI(_agentId *big.Int, _uri string) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.UpdateAgentURI(&_SystemPromptManager.TransactOpts, _agentId, _uri)
}

// UpdateAgentUriWithSignature is a paid mutator transaction binding the contract method 0xf5888779.
//
// Solidity: function updateAgentUriWithSignature(uint256 _agentId, string _uri, uint256 _randomNonce, bytes _signature) returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) UpdateAgentUriWithSignature(opts *bind.TransactOpts, _agentId *big.Int, _uri string, _randomNonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "updateAgentUriWithSignature", _agentId, _uri, _randomNonce, _signature)
}

// UpdateAgentUriWithSignature is a paid mutator transaction binding the contract method 0xf5888779.
//
// Solidity: function updateAgentUriWithSignature(uint256 _agentId, string _uri, uint256 _randomNonce, bytes _signature) returns()
func (_SystemPromptManager *SystemPromptManagerSession) UpdateAgentUriWithSignature(_agentId *big.Int, _uri string, _randomNonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.UpdateAgentUriWithSignature(&_SystemPromptManager.TransactOpts, _agentId, _uri, _randomNonce, _signature)
}

// UpdateAgentUriWithSignature is a paid mutator transaction binding the contract method 0xf5888779.
//
// Solidity: function updateAgentUriWithSignature(uint256 _agentId, string _uri, uint256 _randomNonce, bytes _signature) returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) UpdateAgentUriWithSignature(_agentId *big.Int, _uri string, _randomNonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.UpdateAgentUriWithSignature(&_SystemPromptManager.TransactOpts, _agentId, _uri, _randomNonce, _signature)
}

// UpdateMintPrice is a paid mutator transaction binding the contract method 0x00728e46.
//
// Solidity: function updateMintPrice(uint256 _mintPrice) returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) UpdateMintPrice(opts *bind.TransactOpts, _mintPrice *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "updateMintPrice", _mintPrice)
}

// UpdateMintPrice is a paid mutator transaction binding the contract method 0x00728e46.
//
// Solidity: function updateMintPrice(uint256 _mintPrice) returns()
func (_SystemPromptManager *SystemPromptManagerSession) UpdateMintPrice(_mintPrice *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.UpdateMintPrice(&_SystemPromptManager.TransactOpts, _mintPrice)
}

// UpdateMintPrice is a paid mutator transaction binding the contract method 0x00728e46.
//
// Solidity: function updateMintPrice(uint256 _mintPrice) returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) UpdateMintPrice(_mintPrice *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.UpdateMintPrice(&_SystemPromptManager.TransactOpts, _mintPrice)
}

// UpdateRoyaltyPortion is a paid mutator transaction binding the contract method 0x19e93993.
//
// Solidity: function updateRoyaltyPortion(uint16 _royaltyPortion) returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) UpdateRoyaltyPortion(opts *bind.TransactOpts, _royaltyPortion uint16) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "updateRoyaltyPortion", _royaltyPortion)
}

// UpdateRoyaltyPortion is a paid mutator transaction binding the contract method 0x19e93993.
//
// Solidity: function updateRoyaltyPortion(uint16 _royaltyPortion) returns()
func (_SystemPromptManager *SystemPromptManagerSession) UpdateRoyaltyPortion(_royaltyPortion uint16) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.UpdateRoyaltyPortion(&_SystemPromptManager.TransactOpts, _royaltyPortion)
}

// UpdateRoyaltyPortion is a paid mutator transaction binding the contract method 0x19e93993.
//
// Solidity: function updateRoyaltyPortion(uint16 _royaltyPortion) returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) UpdateRoyaltyPortion(_royaltyPortion uint16) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.UpdateRoyaltyPortion(&_SystemPromptManager.TransactOpts, _royaltyPortion)
}

// UpdateRoyaltyReceiver is a paid mutator transaction binding the contract method 0x29dc4d9b.
//
// Solidity: function updateRoyaltyReceiver(address _royaltyReceiver) returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) UpdateRoyaltyReceiver(opts *bind.TransactOpts, _royaltyReceiver common.Address) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "updateRoyaltyReceiver", _royaltyReceiver)
}

// UpdateRoyaltyReceiver is a paid mutator transaction binding the contract method 0x29dc4d9b.
//
// Solidity: function updateRoyaltyReceiver(address _royaltyReceiver) returns()
func (_SystemPromptManager *SystemPromptManagerSession) UpdateRoyaltyReceiver(_royaltyReceiver common.Address) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.UpdateRoyaltyReceiver(&_SystemPromptManager.TransactOpts, _royaltyReceiver)
}

// UpdateRoyaltyReceiver is a paid mutator transaction binding the contract method 0x29dc4d9b.
//
// Solidity: function updateRoyaltyReceiver(address _royaltyReceiver) returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) UpdateRoyaltyReceiver(_royaltyReceiver common.Address) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.UpdateRoyaltyReceiver(&_SystemPromptManager.TransactOpts, _royaltyReceiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _to, uint256 _value) returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) Withdraw(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.contract.Transact(opts, "withdraw", _to, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _to, uint256 _value) returns()
func (_SystemPromptManager *SystemPromptManagerSession) Withdraw(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.Withdraw(&_SystemPromptManager.TransactOpts, _to, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _to, uint256 _value) returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) Withdraw(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SystemPromptManager.Contract.Withdraw(&_SystemPromptManager.TransactOpts, _to, _value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SystemPromptManager *SystemPromptManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemPromptManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SystemPromptManager *SystemPromptManagerSession) Receive() (*types.Transaction, error) {
	return _SystemPromptManager.Contract.Receive(&_SystemPromptManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SystemPromptManager *SystemPromptManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _SystemPromptManager.Contract.Receive(&_SystemPromptManager.TransactOpts)
}

// SystemPromptManagerAgentDataAddNewIterator is returned from FilterAgentDataAddNew and is used to iterate over the raw logs and unpacked data for AgentDataAddNew events raised by the SystemPromptManager contract.
type SystemPromptManagerAgentDataAddNewIterator struct {
	Event *SystemPromptManagerAgentDataAddNew // Event containing the contract specifics and raw log

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
func (it *SystemPromptManagerAgentDataAddNewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemPromptManagerAgentDataAddNew)
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
		it.Event = new(SystemPromptManagerAgentDataAddNew)
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
func (it *SystemPromptManagerAgentDataAddNewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemPromptManagerAgentDataAddNewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemPromptManagerAgentDataAddNew represents a AgentDataAddNew event raised by the SystemPromptManager contract.
type SystemPromptManagerAgentDataAddNew struct {
	AgentId   *big.Int
	SysPrompt [][]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAgentDataAddNew is a free log retrieval operation binding the contract event 0xdebec4c58e3b7c5817893e50cb1f9e65b65978e8c89bb4407eb0109d5887b258.
//
// Solidity: event AgentDataAddNew(uint256 indexed agentId, bytes[] sysPrompt)
func (_SystemPromptManager *SystemPromptManagerFilterer) FilterAgentDataAddNew(opts *bind.FilterOpts, agentId []*big.Int) (*SystemPromptManagerAgentDataAddNewIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _SystemPromptManager.contract.FilterLogs(opts, "AgentDataAddNew", agentIdRule)
	if err != nil {
		return nil, err
	}
	return &SystemPromptManagerAgentDataAddNewIterator{contract: _SystemPromptManager.contract, event: "AgentDataAddNew", logs: logs, sub: sub}, nil
}

// WatchAgentDataAddNew is a free log subscription operation binding the contract event 0xdebec4c58e3b7c5817893e50cb1f9e65b65978e8c89bb4407eb0109d5887b258.
//
// Solidity: event AgentDataAddNew(uint256 indexed agentId, bytes[] sysPrompt)
func (_SystemPromptManager *SystemPromptManagerFilterer) WatchAgentDataAddNew(opts *bind.WatchOpts, sink chan<- *SystemPromptManagerAgentDataAddNew, agentId []*big.Int) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _SystemPromptManager.contract.WatchLogs(opts, "AgentDataAddNew", agentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemPromptManagerAgentDataAddNew)
				if err := _SystemPromptManager.contract.UnpackLog(event, "AgentDataAddNew", log); err != nil {
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
func (_SystemPromptManager *SystemPromptManagerFilterer) ParseAgentDataAddNew(log types.Log) (*SystemPromptManagerAgentDataAddNew, error) {
	event := new(SystemPromptManagerAgentDataAddNew)
	if err := _SystemPromptManager.contract.UnpackLog(event, "AgentDataAddNew", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemPromptManagerAgentDataUpdateIterator is returned from FilterAgentDataUpdate and is used to iterate over the raw logs and unpacked data for AgentDataUpdate events raised by the SystemPromptManager contract.
type SystemPromptManagerAgentDataUpdateIterator struct {
	Event *SystemPromptManagerAgentDataUpdate // Event containing the contract specifics and raw log

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
func (it *SystemPromptManagerAgentDataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemPromptManagerAgentDataUpdate)
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
		it.Event = new(SystemPromptManagerAgentDataUpdate)
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
func (it *SystemPromptManagerAgentDataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemPromptManagerAgentDataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemPromptManagerAgentDataUpdate represents a AgentDataUpdate event raised by the SystemPromptManager contract.
type SystemPromptManagerAgentDataUpdate struct {
	AgentId      *big.Int
	PromptIndex  *big.Int
	OldSysPrompt []byte
	NewSysPrompt []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAgentDataUpdate is a free log retrieval operation binding the contract event 0xe42abf7d4a793286da8cc1399cb577a1f5a0e133dfee371bb3a5abbdd77b011e.
//
// Solidity: event AgentDataUpdate(uint256 indexed agentId, uint256 promptIndex, bytes oldSysPrompt, bytes newSysPrompt)
func (_SystemPromptManager *SystemPromptManagerFilterer) FilterAgentDataUpdate(opts *bind.FilterOpts, agentId []*big.Int) (*SystemPromptManagerAgentDataUpdateIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _SystemPromptManager.contract.FilterLogs(opts, "AgentDataUpdate", agentIdRule)
	if err != nil {
		return nil, err
	}
	return &SystemPromptManagerAgentDataUpdateIterator{contract: _SystemPromptManager.contract, event: "AgentDataUpdate", logs: logs, sub: sub}, nil
}

// WatchAgentDataUpdate is a free log subscription operation binding the contract event 0xe42abf7d4a793286da8cc1399cb577a1f5a0e133dfee371bb3a5abbdd77b011e.
//
// Solidity: event AgentDataUpdate(uint256 indexed agentId, uint256 promptIndex, bytes oldSysPrompt, bytes newSysPrompt)
func (_SystemPromptManager *SystemPromptManagerFilterer) WatchAgentDataUpdate(opts *bind.WatchOpts, sink chan<- *SystemPromptManagerAgentDataUpdate, agentId []*big.Int) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _SystemPromptManager.contract.WatchLogs(opts, "AgentDataUpdate", agentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemPromptManagerAgentDataUpdate)
				if err := _SystemPromptManager.contract.UnpackLog(event, "AgentDataUpdate", log); err != nil {
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
func (_SystemPromptManager *SystemPromptManagerFilterer) ParseAgentDataUpdate(log types.Log) (*SystemPromptManagerAgentDataUpdate, error) {
	event := new(SystemPromptManagerAgentDataUpdate)
	if err := _SystemPromptManager.contract.UnpackLog(event, "AgentDataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemPromptManagerAgentFeeUpdateIterator is returned from FilterAgentFeeUpdate and is used to iterate over the raw logs and unpacked data for AgentFeeUpdate events raised by the SystemPromptManager contract.
type SystemPromptManagerAgentFeeUpdateIterator struct {
	Event *SystemPromptManagerAgentFeeUpdate // Event containing the contract specifics and raw log

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
func (it *SystemPromptManagerAgentFeeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemPromptManagerAgentFeeUpdate)
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
		it.Event = new(SystemPromptManagerAgentFeeUpdate)
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
func (it *SystemPromptManagerAgentFeeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemPromptManagerAgentFeeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemPromptManagerAgentFeeUpdate represents a AgentFeeUpdate event raised by the SystemPromptManager contract.
type SystemPromptManagerAgentFeeUpdate struct {
	AgentId *big.Int
	Fee     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentFeeUpdate is a free log retrieval operation binding the contract event 0xa08d8197034aee8915921dd8aa7d95cf711690dd77f0b676dded49b3f9a632d1.
//
// Solidity: event AgentFeeUpdate(uint256 indexed agentId, uint256 fee)
func (_SystemPromptManager *SystemPromptManagerFilterer) FilterAgentFeeUpdate(opts *bind.FilterOpts, agentId []*big.Int) (*SystemPromptManagerAgentFeeUpdateIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _SystemPromptManager.contract.FilterLogs(opts, "AgentFeeUpdate", agentIdRule)
	if err != nil {
		return nil, err
	}
	return &SystemPromptManagerAgentFeeUpdateIterator{contract: _SystemPromptManager.contract, event: "AgentFeeUpdate", logs: logs, sub: sub}, nil
}

// WatchAgentFeeUpdate is a free log subscription operation binding the contract event 0xa08d8197034aee8915921dd8aa7d95cf711690dd77f0b676dded49b3f9a632d1.
//
// Solidity: event AgentFeeUpdate(uint256 indexed agentId, uint256 fee)
func (_SystemPromptManager *SystemPromptManagerFilterer) WatchAgentFeeUpdate(opts *bind.WatchOpts, sink chan<- *SystemPromptManagerAgentFeeUpdate, agentId []*big.Int) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _SystemPromptManager.contract.WatchLogs(opts, "AgentFeeUpdate", agentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemPromptManagerAgentFeeUpdate)
				if err := _SystemPromptManager.contract.UnpackLog(event, "AgentFeeUpdate", log); err != nil {
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
func (_SystemPromptManager *SystemPromptManagerFilterer) ParseAgentFeeUpdate(log types.Log) (*SystemPromptManagerAgentFeeUpdate, error) {
	event := new(SystemPromptManagerAgentFeeUpdate)
	if err := _SystemPromptManager.contract.UnpackLog(event, "AgentFeeUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemPromptManagerAgentURIUpdateIterator is returned from FilterAgentURIUpdate and is used to iterate over the raw logs and unpacked data for AgentURIUpdate events raised by the SystemPromptManager contract.
type SystemPromptManagerAgentURIUpdateIterator struct {
	Event *SystemPromptManagerAgentURIUpdate // Event containing the contract specifics and raw log

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
func (it *SystemPromptManagerAgentURIUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemPromptManagerAgentURIUpdate)
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
		it.Event = new(SystemPromptManagerAgentURIUpdate)
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
func (it *SystemPromptManagerAgentURIUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemPromptManagerAgentURIUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemPromptManagerAgentURIUpdate represents a AgentURIUpdate event raised by the SystemPromptManager contract.
type SystemPromptManagerAgentURIUpdate struct {
	AgentId *big.Int
	Uri     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentURIUpdate is a free log retrieval operation binding the contract event 0x706a4e8eb2f354c7f4d96e5ea1984f36e72482629987edad78c9940ea037c362.
//
// Solidity: event AgentURIUpdate(uint256 indexed agentId, string uri)
func (_SystemPromptManager *SystemPromptManagerFilterer) FilterAgentURIUpdate(opts *bind.FilterOpts, agentId []*big.Int) (*SystemPromptManagerAgentURIUpdateIterator, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _SystemPromptManager.contract.FilterLogs(opts, "AgentURIUpdate", agentIdRule)
	if err != nil {
		return nil, err
	}
	return &SystemPromptManagerAgentURIUpdateIterator{contract: _SystemPromptManager.contract, event: "AgentURIUpdate", logs: logs, sub: sub}, nil
}

// WatchAgentURIUpdate is a free log subscription operation binding the contract event 0x706a4e8eb2f354c7f4d96e5ea1984f36e72482629987edad78c9940ea037c362.
//
// Solidity: event AgentURIUpdate(uint256 indexed agentId, string uri)
func (_SystemPromptManager *SystemPromptManagerFilterer) WatchAgentURIUpdate(opts *bind.WatchOpts, sink chan<- *SystemPromptManagerAgentURIUpdate, agentId []*big.Int) (event.Subscription, error) {

	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _SystemPromptManager.contract.WatchLogs(opts, "AgentURIUpdate", agentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemPromptManagerAgentURIUpdate)
				if err := _SystemPromptManager.contract.UnpackLog(event, "AgentURIUpdate", log); err != nil {
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
func (_SystemPromptManager *SystemPromptManagerFilterer) ParseAgentURIUpdate(log types.Log) (*SystemPromptManagerAgentURIUpdate, error) {
	event := new(SystemPromptManagerAgentURIUpdate)
	if err := _SystemPromptManager.contract.UnpackLog(event, "AgentURIUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemPromptManagerApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SystemPromptManager contract.
type SystemPromptManagerApprovalIterator struct {
	Event *SystemPromptManagerApproval // Event containing the contract specifics and raw log

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
func (it *SystemPromptManagerApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemPromptManagerApproval)
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
		it.Event = new(SystemPromptManagerApproval)
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
func (it *SystemPromptManagerApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemPromptManagerApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemPromptManagerApproval represents a Approval event raised by the SystemPromptManager contract.
type SystemPromptManagerApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SystemPromptManager *SystemPromptManagerFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*SystemPromptManagerApprovalIterator, error) {

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

	logs, sub, err := _SystemPromptManager.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SystemPromptManagerApprovalIterator{contract: _SystemPromptManager.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SystemPromptManager *SystemPromptManagerFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SystemPromptManagerApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _SystemPromptManager.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemPromptManagerApproval)
				if err := _SystemPromptManager.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_SystemPromptManager *SystemPromptManagerFilterer) ParseApproval(log types.Log) (*SystemPromptManagerApproval, error) {
	event := new(SystemPromptManagerApproval)
	if err := _SystemPromptManager.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemPromptManagerApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the SystemPromptManager contract.
type SystemPromptManagerApprovalForAllIterator struct {
	Event *SystemPromptManagerApprovalForAll // Event containing the contract specifics and raw log

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
func (it *SystemPromptManagerApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemPromptManagerApprovalForAll)
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
		it.Event = new(SystemPromptManagerApprovalForAll)
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
func (it *SystemPromptManagerApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemPromptManagerApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemPromptManagerApprovalForAll represents a ApprovalForAll event raised by the SystemPromptManager contract.
type SystemPromptManagerApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SystemPromptManager *SystemPromptManagerFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*SystemPromptManagerApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SystemPromptManager.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &SystemPromptManagerApprovalForAllIterator{contract: _SystemPromptManager.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SystemPromptManager *SystemPromptManagerFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *SystemPromptManagerApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SystemPromptManager.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemPromptManagerApprovalForAll)
				if err := _SystemPromptManager.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_SystemPromptManager *SystemPromptManagerFilterer) ParseApprovalForAll(log types.Log) (*SystemPromptManagerApprovalForAll, error) {
	event := new(SystemPromptManagerApprovalForAll)
	if err := _SystemPromptManager.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemPromptManagerBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the SystemPromptManager contract.
type SystemPromptManagerBatchMetadataUpdateIterator struct {
	Event *SystemPromptManagerBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *SystemPromptManagerBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemPromptManagerBatchMetadataUpdate)
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
		it.Event = new(SystemPromptManagerBatchMetadataUpdate)
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
func (it *SystemPromptManagerBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemPromptManagerBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemPromptManagerBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the SystemPromptManager contract.
type SystemPromptManagerBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_SystemPromptManager *SystemPromptManagerFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*SystemPromptManagerBatchMetadataUpdateIterator, error) {

	logs, sub, err := _SystemPromptManager.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &SystemPromptManagerBatchMetadataUpdateIterator{contract: _SystemPromptManager.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_SystemPromptManager *SystemPromptManagerFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *SystemPromptManagerBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _SystemPromptManager.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemPromptManagerBatchMetadataUpdate)
				if err := _SystemPromptManager.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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
func (_SystemPromptManager *SystemPromptManagerFilterer) ParseBatchMetadataUpdate(log types.Log) (*SystemPromptManagerBatchMetadataUpdate, error) {
	event := new(SystemPromptManagerBatchMetadataUpdate)
	if err := _SystemPromptManager.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemPromptManagerEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the SystemPromptManager contract.
type SystemPromptManagerEIP712DomainChangedIterator struct {
	Event *SystemPromptManagerEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *SystemPromptManagerEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemPromptManagerEIP712DomainChanged)
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
		it.Event = new(SystemPromptManagerEIP712DomainChanged)
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
func (it *SystemPromptManagerEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemPromptManagerEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemPromptManagerEIP712DomainChanged represents a EIP712DomainChanged event raised by the SystemPromptManager contract.
type SystemPromptManagerEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_SystemPromptManager *SystemPromptManagerFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*SystemPromptManagerEIP712DomainChangedIterator, error) {

	logs, sub, err := _SystemPromptManager.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &SystemPromptManagerEIP712DomainChangedIterator{contract: _SystemPromptManager.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_SystemPromptManager *SystemPromptManagerFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *SystemPromptManagerEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _SystemPromptManager.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemPromptManagerEIP712DomainChanged)
				if err := _SystemPromptManager.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_SystemPromptManager *SystemPromptManagerFilterer) ParseEIP712DomainChanged(log types.Log) (*SystemPromptManagerEIP712DomainChanged, error) {
	event := new(SystemPromptManagerEIP712DomainChanged)
	if err := _SystemPromptManager.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemPromptManagerFeesClaimedIterator is returned from FilterFeesClaimed and is used to iterate over the raw logs and unpacked data for FeesClaimed events raised by the SystemPromptManager contract.
type SystemPromptManagerFeesClaimedIterator struct {
	Event *SystemPromptManagerFeesClaimed // Event containing the contract specifics and raw log

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
func (it *SystemPromptManagerFeesClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemPromptManagerFeesClaimed)
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
		it.Event = new(SystemPromptManagerFeesClaimed)
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
func (it *SystemPromptManagerFeesClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemPromptManagerFeesClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemPromptManagerFeesClaimed represents a FeesClaimed event raised by the SystemPromptManager contract.
type SystemPromptManagerFeesClaimed struct {
	Claimer common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeesClaimed is a free log retrieval operation binding the contract event 0x9493e5bbe4e8e0ac67284469a2d677403d0378a85a59e341d3abc433d0d9a209.
//
// Solidity: event FeesClaimed(address indexed claimer, uint256 amount)
func (_SystemPromptManager *SystemPromptManagerFilterer) FilterFeesClaimed(opts *bind.FilterOpts, claimer []common.Address) (*SystemPromptManagerFeesClaimedIterator, error) {

	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _SystemPromptManager.contract.FilterLogs(opts, "FeesClaimed", claimerRule)
	if err != nil {
		return nil, err
	}
	return &SystemPromptManagerFeesClaimedIterator{contract: _SystemPromptManager.contract, event: "FeesClaimed", logs: logs, sub: sub}, nil
}

// WatchFeesClaimed is a free log subscription operation binding the contract event 0x9493e5bbe4e8e0ac67284469a2d677403d0378a85a59e341d3abc433d0d9a209.
//
// Solidity: event FeesClaimed(address indexed claimer, uint256 amount)
func (_SystemPromptManager *SystemPromptManagerFilterer) WatchFeesClaimed(opts *bind.WatchOpts, sink chan<- *SystemPromptManagerFeesClaimed, claimer []common.Address) (event.Subscription, error) {

	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _SystemPromptManager.contract.WatchLogs(opts, "FeesClaimed", claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemPromptManagerFeesClaimed)
				if err := _SystemPromptManager.contract.UnpackLog(event, "FeesClaimed", log); err != nil {
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
func (_SystemPromptManager *SystemPromptManagerFilterer) ParseFeesClaimed(log types.Log) (*SystemPromptManagerFeesClaimed, error) {
	event := new(SystemPromptManagerFeesClaimed)
	if err := _SystemPromptManager.contract.UnpackLog(event, "FeesClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemPromptManagerInferencePerformedIterator is returned from FilterInferencePerformed and is used to iterate over the raw logs and unpacked data for InferencePerformed events raised by the SystemPromptManager contract.
type SystemPromptManagerInferencePerformedIterator struct {
	Event *SystemPromptManagerInferencePerformed // Event containing the contract specifics and raw log

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
func (it *SystemPromptManagerInferencePerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemPromptManagerInferencePerformed)
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
		it.Event = new(SystemPromptManagerInferencePerformed)
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
func (it *SystemPromptManagerInferencePerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemPromptManagerInferencePerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemPromptManagerInferencePerformed represents a InferencePerformed event raised by the SystemPromptManager contract.
type SystemPromptManagerInferencePerformed struct {
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
func (_SystemPromptManager *SystemPromptManagerFilterer) FilterInferencePerformed(opts *bind.FilterOpts, tokenId []*big.Int, caller []common.Address) (*SystemPromptManagerInferencePerformedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _SystemPromptManager.contract.FilterLogs(opts, "InferencePerformed", tokenIdRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &SystemPromptManagerInferencePerformedIterator{contract: _SystemPromptManager.contract, event: "InferencePerformed", logs: logs, sub: sub}, nil
}

// WatchInferencePerformed is a free log subscription operation binding the contract event 0xcf35460eca25a0549d5eb14c712236d61c9a0bad90c834f996c5f2a98d332719.
//
// Solidity: event InferencePerformed(uint256 indexed tokenId, address indexed caller, bytes data, uint256 fee, string externalData, uint256 inferenceId)
func (_SystemPromptManager *SystemPromptManagerFilterer) WatchInferencePerformed(opts *bind.WatchOpts, sink chan<- *SystemPromptManagerInferencePerformed, tokenId []*big.Int, caller []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _SystemPromptManager.contract.WatchLogs(opts, "InferencePerformed", tokenIdRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemPromptManagerInferencePerformed)
				if err := _SystemPromptManager.contract.UnpackLog(event, "InferencePerformed", log); err != nil {
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
func (_SystemPromptManager *SystemPromptManagerFilterer) ParseInferencePerformed(log types.Log) (*SystemPromptManagerInferencePerformed, error) {
	event := new(SystemPromptManagerInferencePerformed)
	if err := _SystemPromptManager.contract.UnpackLog(event, "InferencePerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemPromptManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SystemPromptManager contract.
type SystemPromptManagerInitializedIterator struct {
	Event *SystemPromptManagerInitialized // Event containing the contract specifics and raw log

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
func (it *SystemPromptManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemPromptManagerInitialized)
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
		it.Event = new(SystemPromptManagerInitialized)
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
func (it *SystemPromptManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemPromptManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemPromptManagerInitialized represents a Initialized event raised by the SystemPromptManager contract.
type SystemPromptManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SystemPromptManager *SystemPromptManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*SystemPromptManagerInitializedIterator, error) {

	logs, sub, err := _SystemPromptManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SystemPromptManagerInitializedIterator{contract: _SystemPromptManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SystemPromptManager *SystemPromptManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SystemPromptManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _SystemPromptManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemPromptManagerInitialized)
				if err := _SystemPromptManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SystemPromptManager *SystemPromptManagerFilterer) ParseInitialized(log types.Log) (*SystemPromptManagerInitialized, error) {
	event := new(SystemPromptManagerInitialized)
	if err := _SystemPromptManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemPromptManagerManagerAuthorizationIterator is returned from FilterManagerAuthorization and is used to iterate over the raw logs and unpacked data for ManagerAuthorization events raised by the SystemPromptManager contract.
type SystemPromptManagerManagerAuthorizationIterator struct {
	Event *SystemPromptManagerManagerAuthorization // Event containing the contract specifics and raw log

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
func (it *SystemPromptManagerManagerAuthorizationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemPromptManagerManagerAuthorization)
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
		it.Event = new(SystemPromptManagerManagerAuthorization)
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
func (it *SystemPromptManagerManagerAuthorizationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemPromptManagerManagerAuthorizationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemPromptManagerManagerAuthorization represents a ManagerAuthorization event raised by the SystemPromptManager contract.
type SystemPromptManagerManagerAuthorization struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerAuthorization is a free log retrieval operation binding the contract event 0x3fbc8e71624117c0dc0fcdbe40685681cecc7c3c43de81a686cda4b61c78e35b.
//
// Solidity: event ManagerAuthorization(address indexed account)
func (_SystemPromptManager *SystemPromptManagerFilterer) FilterManagerAuthorization(opts *bind.FilterOpts, account []common.Address) (*SystemPromptManagerManagerAuthorizationIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SystemPromptManager.contract.FilterLogs(opts, "ManagerAuthorization", accountRule)
	if err != nil {
		return nil, err
	}
	return &SystemPromptManagerManagerAuthorizationIterator{contract: _SystemPromptManager.contract, event: "ManagerAuthorization", logs: logs, sub: sub}, nil
}

// WatchManagerAuthorization is a free log subscription operation binding the contract event 0x3fbc8e71624117c0dc0fcdbe40685681cecc7c3c43de81a686cda4b61c78e35b.
//
// Solidity: event ManagerAuthorization(address indexed account)
func (_SystemPromptManager *SystemPromptManagerFilterer) WatchManagerAuthorization(opts *bind.WatchOpts, sink chan<- *SystemPromptManagerManagerAuthorization, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SystemPromptManager.contract.WatchLogs(opts, "ManagerAuthorization", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemPromptManagerManagerAuthorization)
				if err := _SystemPromptManager.contract.UnpackLog(event, "ManagerAuthorization", log); err != nil {
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
func (_SystemPromptManager *SystemPromptManagerFilterer) ParseManagerAuthorization(log types.Log) (*SystemPromptManagerManagerAuthorization, error) {
	event := new(SystemPromptManagerManagerAuthorization)
	if err := _SystemPromptManager.contract.UnpackLog(event, "ManagerAuthorization", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemPromptManagerManagerDeauthorizationIterator is returned from FilterManagerDeauthorization and is used to iterate over the raw logs and unpacked data for ManagerDeauthorization events raised by the SystemPromptManager contract.
type SystemPromptManagerManagerDeauthorizationIterator struct {
	Event *SystemPromptManagerManagerDeauthorization // Event containing the contract specifics and raw log

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
func (it *SystemPromptManagerManagerDeauthorizationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemPromptManagerManagerDeauthorization)
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
		it.Event = new(SystemPromptManagerManagerDeauthorization)
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
func (it *SystemPromptManagerManagerDeauthorizationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemPromptManagerManagerDeauthorizationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemPromptManagerManagerDeauthorization represents a ManagerDeauthorization event raised by the SystemPromptManager contract.
type SystemPromptManagerManagerDeauthorization struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerDeauthorization is a free log retrieval operation binding the contract event 0x20c29af9eb3de2601188ceae57a4075ba3593ce15d4142aef070ac53d389356c.
//
// Solidity: event ManagerDeauthorization(address indexed account)
func (_SystemPromptManager *SystemPromptManagerFilterer) FilterManagerDeauthorization(opts *bind.FilterOpts, account []common.Address) (*SystemPromptManagerManagerDeauthorizationIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SystemPromptManager.contract.FilterLogs(opts, "ManagerDeauthorization", accountRule)
	if err != nil {
		return nil, err
	}
	return &SystemPromptManagerManagerDeauthorizationIterator{contract: _SystemPromptManager.contract, event: "ManagerDeauthorization", logs: logs, sub: sub}, nil
}

// WatchManagerDeauthorization is a free log subscription operation binding the contract event 0x20c29af9eb3de2601188ceae57a4075ba3593ce15d4142aef070ac53d389356c.
//
// Solidity: event ManagerDeauthorization(address indexed account)
func (_SystemPromptManager *SystemPromptManagerFilterer) WatchManagerDeauthorization(opts *bind.WatchOpts, sink chan<- *SystemPromptManagerManagerDeauthorization, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SystemPromptManager.contract.WatchLogs(opts, "ManagerDeauthorization", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemPromptManagerManagerDeauthorization)
				if err := _SystemPromptManager.contract.UnpackLog(event, "ManagerDeauthorization", log); err != nil {
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
func (_SystemPromptManager *SystemPromptManagerFilterer) ParseManagerDeauthorization(log types.Log) (*SystemPromptManagerManagerDeauthorization, error) {
	event := new(SystemPromptManagerManagerDeauthorization)
	if err := _SystemPromptManager.contract.UnpackLog(event, "ManagerDeauthorization", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemPromptManagerMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the SystemPromptManager contract.
type SystemPromptManagerMetadataUpdateIterator struct {
	Event *SystemPromptManagerMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *SystemPromptManagerMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemPromptManagerMetadataUpdate)
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
		it.Event = new(SystemPromptManagerMetadataUpdate)
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
func (it *SystemPromptManagerMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemPromptManagerMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemPromptManagerMetadataUpdate represents a MetadataUpdate event raised by the SystemPromptManager contract.
type SystemPromptManagerMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_SystemPromptManager *SystemPromptManagerFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*SystemPromptManagerMetadataUpdateIterator, error) {

	logs, sub, err := _SystemPromptManager.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &SystemPromptManagerMetadataUpdateIterator{contract: _SystemPromptManager.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_SystemPromptManager *SystemPromptManagerFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *SystemPromptManagerMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _SystemPromptManager.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemPromptManagerMetadataUpdate)
				if err := _SystemPromptManager.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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
func (_SystemPromptManager *SystemPromptManagerFilterer) ParseMetadataUpdate(log types.Log) (*SystemPromptManagerMetadataUpdate, error) {
	event := new(SystemPromptManagerMetadataUpdate)
	if err := _SystemPromptManager.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemPromptManagerMintPriceUpdateIterator is returned from FilterMintPriceUpdate and is used to iterate over the raw logs and unpacked data for MintPriceUpdate events raised by the SystemPromptManager contract.
type SystemPromptManagerMintPriceUpdateIterator struct {
	Event *SystemPromptManagerMintPriceUpdate // Event containing the contract specifics and raw log

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
func (it *SystemPromptManagerMintPriceUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemPromptManagerMintPriceUpdate)
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
		it.Event = new(SystemPromptManagerMintPriceUpdate)
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
func (it *SystemPromptManagerMintPriceUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemPromptManagerMintPriceUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemPromptManagerMintPriceUpdate represents a MintPriceUpdate event raised by the SystemPromptManager contract.
type SystemPromptManagerMintPriceUpdate struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMintPriceUpdate is a free log retrieval operation binding the contract event 0x23050b539195eebd064c1bec834445b7d028a20c345600e868a74d7ca93c5e86.
//
// Solidity: event MintPriceUpdate(uint256 newValue)
func (_SystemPromptManager *SystemPromptManagerFilterer) FilterMintPriceUpdate(opts *bind.FilterOpts) (*SystemPromptManagerMintPriceUpdateIterator, error) {

	logs, sub, err := _SystemPromptManager.contract.FilterLogs(opts, "MintPriceUpdate")
	if err != nil {
		return nil, err
	}
	return &SystemPromptManagerMintPriceUpdateIterator{contract: _SystemPromptManager.contract, event: "MintPriceUpdate", logs: logs, sub: sub}, nil
}

// WatchMintPriceUpdate is a free log subscription operation binding the contract event 0x23050b539195eebd064c1bec834445b7d028a20c345600e868a74d7ca93c5e86.
//
// Solidity: event MintPriceUpdate(uint256 newValue)
func (_SystemPromptManager *SystemPromptManagerFilterer) WatchMintPriceUpdate(opts *bind.WatchOpts, sink chan<- *SystemPromptManagerMintPriceUpdate) (event.Subscription, error) {

	logs, sub, err := _SystemPromptManager.contract.WatchLogs(opts, "MintPriceUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemPromptManagerMintPriceUpdate)
				if err := _SystemPromptManager.contract.UnpackLog(event, "MintPriceUpdate", log); err != nil {
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
func (_SystemPromptManager *SystemPromptManagerFilterer) ParseMintPriceUpdate(log types.Log) (*SystemPromptManagerMintPriceUpdate, error) {
	event := new(SystemPromptManagerMintPriceUpdate)
	if err := _SystemPromptManager.contract.UnpackLog(event, "MintPriceUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemPromptManagerNewTokenIterator is returned from FilterNewToken and is used to iterate over the raw logs and unpacked data for NewToken events raised by the SystemPromptManager contract.
type SystemPromptManagerNewTokenIterator struct {
	Event *SystemPromptManagerNewToken // Event containing the contract specifics and raw log

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
func (it *SystemPromptManagerNewTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemPromptManagerNewToken)
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
		it.Event = new(SystemPromptManagerNewToken)
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
func (it *SystemPromptManagerNewTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemPromptManagerNewTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemPromptManagerNewToken represents a NewToken event raised by the SystemPromptManager contract.
type SystemPromptManagerNewToken struct {
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
func (_SystemPromptManager *SystemPromptManagerFilterer) FilterNewToken(opts *bind.FilterOpts, tokenId []*big.Int, minter []common.Address) (*SystemPromptManagerNewTokenIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _SystemPromptManager.contract.FilterLogs(opts, "NewToken", tokenIdRule, minterRule)
	if err != nil {
		return nil, err
	}
	return &SystemPromptManagerNewTokenIterator{contract: _SystemPromptManager.contract, event: "NewToken", logs: logs, sub: sub}, nil
}

// WatchNewToken is a free log subscription operation binding the contract event 0x61beab98a81083e3c0239c33e149bef1316ca78f15b9f29125039f5521a06d06.
//
// Solidity: event NewToken(uint256 indexed tokenId, string uri, bytes sysPrompt, uint256 fee, address indexed minter)
func (_SystemPromptManager *SystemPromptManagerFilterer) WatchNewToken(opts *bind.WatchOpts, sink chan<- *SystemPromptManagerNewToken, tokenId []*big.Int, minter []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _SystemPromptManager.contract.WatchLogs(opts, "NewToken", tokenIdRule, minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemPromptManagerNewToken)
				if err := _SystemPromptManager.contract.UnpackLog(event, "NewToken", log); err != nil {
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
func (_SystemPromptManager *SystemPromptManagerFilterer) ParseNewToken(log types.Log) (*SystemPromptManagerNewToken, error) {
	event := new(SystemPromptManagerNewToken)
	if err := _SystemPromptManager.contract.UnpackLog(event, "NewToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemPromptManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SystemPromptManager contract.
type SystemPromptManagerOwnershipTransferredIterator struct {
	Event *SystemPromptManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SystemPromptManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemPromptManagerOwnershipTransferred)
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
		it.Event = new(SystemPromptManagerOwnershipTransferred)
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
func (it *SystemPromptManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemPromptManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemPromptManagerOwnershipTransferred represents a OwnershipTransferred event raised by the SystemPromptManager contract.
type SystemPromptManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SystemPromptManager *SystemPromptManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SystemPromptManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SystemPromptManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SystemPromptManagerOwnershipTransferredIterator{contract: _SystemPromptManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SystemPromptManager *SystemPromptManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SystemPromptManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SystemPromptManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemPromptManagerOwnershipTransferred)
				if err := _SystemPromptManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SystemPromptManager *SystemPromptManagerFilterer) ParseOwnershipTransferred(log types.Log) (*SystemPromptManagerOwnershipTransferred, error) {
	event := new(SystemPromptManagerOwnershipTransferred)
	if err := _SystemPromptManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemPromptManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SystemPromptManager contract.
type SystemPromptManagerPausedIterator struct {
	Event *SystemPromptManagerPaused // Event containing the contract specifics and raw log

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
func (it *SystemPromptManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemPromptManagerPaused)
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
		it.Event = new(SystemPromptManagerPaused)
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
func (it *SystemPromptManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemPromptManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemPromptManagerPaused represents a Paused event raised by the SystemPromptManager contract.
type SystemPromptManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SystemPromptManager *SystemPromptManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*SystemPromptManagerPausedIterator, error) {

	logs, sub, err := _SystemPromptManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SystemPromptManagerPausedIterator{contract: _SystemPromptManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SystemPromptManager *SystemPromptManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SystemPromptManagerPaused) (event.Subscription, error) {

	logs, sub, err := _SystemPromptManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemPromptManagerPaused)
				if err := _SystemPromptManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_SystemPromptManager *SystemPromptManagerFilterer) ParsePaused(log types.Log) (*SystemPromptManagerPaused, error) {
	event := new(SystemPromptManagerPaused)
	if err := _SystemPromptManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemPromptManagerRoyaltyPortionUpdateIterator is returned from FilterRoyaltyPortionUpdate and is used to iterate over the raw logs and unpacked data for RoyaltyPortionUpdate events raised by the SystemPromptManager contract.
type SystemPromptManagerRoyaltyPortionUpdateIterator struct {
	Event *SystemPromptManagerRoyaltyPortionUpdate // Event containing the contract specifics and raw log

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
func (it *SystemPromptManagerRoyaltyPortionUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemPromptManagerRoyaltyPortionUpdate)
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
		it.Event = new(SystemPromptManagerRoyaltyPortionUpdate)
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
func (it *SystemPromptManagerRoyaltyPortionUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemPromptManagerRoyaltyPortionUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemPromptManagerRoyaltyPortionUpdate represents a RoyaltyPortionUpdate event raised by the SystemPromptManager contract.
type SystemPromptManagerRoyaltyPortionUpdate struct {
	NewValue uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyPortionUpdate is a free log retrieval operation binding the contract event 0xb1f3037624bd2d961f6d56696cc10ccc3a676db381e671b1bc58f0ab1f686dd5.
//
// Solidity: event RoyaltyPortionUpdate(uint16 newValue)
func (_SystemPromptManager *SystemPromptManagerFilterer) FilterRoyaltyPortionUpdate(opts *bind.FilterOpts) (*SystemPromptManagerRoyaltyPortionUpdateIterator, error) {

	logs, sub, err := _SystemPromptManager.contract.FilterLogs(opts, "RoyaltyPortionUpdate")
	if err != nil {
		return nil, err
	}
	return &SystemPromptManagerRoyaltyPortionUpdateIterator{contract: _SystemPromptManager.contract, event: "RoyaltyPortionUpdate", logs: logs, sub: sub}, nil
}

// WatchRoyaltyPortionUpdate is a free log subscription operation binding the contract event 0xb1f3037624bd2d961f6d56696cc10ccc3a676db381e671b1bc58f0ab1f686dd5.
//
// Solidity: event RoyaltyPortionUpdate(uint16 newValue)
func (_SystemPromptManager *SystemPromptManagerFilterer) WatchRoyaltyPortionUpdate(opts *bind.WatchOpts, sink chan<- *SystemPromptManagerRoyaltyPortionUpdate) (event.Subscription, error) {

	logs, sub, err := _SystemPromptManager.contract.WatchLogs(opts, "RoyaltyPortionUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemPromptManagerRoyaltyPortionUpdate)
				if err := _SystemPromptManager.contract.UnpackLog(event, "RoyaltyPortionUpdate", log); err != nil {
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
func (_SystemPromptManager *SystemPromptManagerFilterer) ParseRoyaltyPortionUpdate(log types.Log) (*SystemPromptManagerRoyaltyPortionUpdate, error) {
	event := new(SystemPromptManagerRoyaltyPortionUpdate)
	if err := _SystemPromptManager.contract.UnpackLog(event, "RoyaltyPortionUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemPromptManagerRoyaltyReceiverUpdateIterator is returned from FilterRoyaltyReceiverUpdate and is used to iterate over the raw logs and unpacked data for RoyaltyReceiverUpdate events raised by the SystemPromptManager contract.
type SystemPromptManagerRoyaltyReceiverUpdateIterator struct {
	Event *SystemPromptManagerRoyaltyReceiverUpdate // Event containing the contract specifics and raw log

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
func (it *SystemPromptManagerRoyaltyReceiverUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemPromptManagerRoyaltyReceiverUpdate)
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
		it.Event = new(SystemPromptManagerRoyaltyReceiverUpdate)
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
func (it *SystemPromptManagerRoyaltyReceiverUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemPromptManagerRoyaltyReceiverUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemPromptManagerRoyaltyReceiverUpdate represents a RoyaltyReceiverUpdate event raised by the SystemPromptManager contract.
type SystemPromptManagerRoyaltyReceiverUpdate struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyReceiverUpdate is a free log retrieval operation binding the contract event 0xec6b72b10aed766af02b35918b55be261c89aaaa4c8add826471ce35ec7f97b3.
//
// Solidity: event RoyaltyReceiverUpdate(address newAddress)
func (_SystemPromptManager *SystemPromptManagerFilterer) FilterRoyaltyReceiverUpdate(opts *bind.FilterOpts) (*SystemPromptManagerRoyaltyReceiverUpdateIterator, error) {

	logs, sub, err := _SystemPromptManager.contract.FilterLogs(opts, "RoyaltyReceiverUpdate")
	if err != nil {
		return nil, err
	}
	return &SystemPromptManagerRoyaltyReceiverUpdateIterator{contract: _SystemPromptManager.contract, event: "RoyaltyReceiverUpdate", logs: logs, sub: sub}, nil
}

// WatchRoyaltyReceiverUpdate is a free log subscription operation binding the contract event 0xec6b72b10aed766af02b35918b55be261c89aaaa4c8add826471ce35ec7f97b3.
//
// Solidity: event RoyaltyReceiverUpdate(address newAddress)
func (_SystemPromptManager *SystemPromptManagerFilterer) WatchRoyaltyReceiverUpdate(opts *bind.WatchOpts, sink chan<- *SystemPromptManagerRoyaltyReceiverUpdate) (event.Subscription, error) {

	logs, sub, err := _SystemPromptManager.contract.WatchLogs(opts, "RoyaltyReceiverUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemPromptManagerRoyaltyReceiverUpdate)
				if err := _SystemPromptManager.contract.UnpackLog(event, "RoyaltyReceiverUpdate", log); err != nil {
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
func (_SystemPromptManager *SystemPromptManagerFilterer) ParseRoyaltyReceiverUpdate(log types.Log) (*SystemPromptManagerRoyaltyReceiverUpdate, error) {
	event := new(SystemPromptManagerRoyaltyReceiverUpdate)
	if err := _SystemPromptManager.contract.UnpackLog(event, "RoyaltyReceiverUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemPromptManagerTopUpPoolBalanceIterator is returned from FilterTopUpPoolBalance and is used to iterate over the raw logs and unpacked data for TopUpPoolBalance events raised by the SystemPromptManager contract.
type SystemPromptManagerTopUpPoolBalanceIterator struct {
	Event *SystemPromptManagerTopUpPoolBalance // Event containing the contract specifics and raw log

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
func (it *SystemPromptManagerTopUpPoolBalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemPromptManagerTopUpPoolBalance)
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
		it.Event = new(SystemPromptManagerTopUpPoolBalance)
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
func (it *SystemPromptManagerTopUpPoolBalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemPromptManagerTopUpPoolBalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemPromptManagerTopUpPoolBalance represents a TopUpPoolBalance event raised by the SystemPromptManager contract.
type SystemPromptManagerTopUpPoolBalance struct {
	AgentId *big.Int
	Caller  common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTopUpPoolBalance is a free log retrieval operation binding the contract event 0xf7ee57effd30f2ab842c1d65fdcfa7d20c2fb2f967e9ac30acafecabf013ea4c.
//
// Solidity: event TopUpPoolBalance(uint256 agentId, address caller, uint256 amount)
func (_SystemPromptManager *SystemPromptManagerFilterer) FilterTopUpPoolBalance(opts *bind.FilterOpts) (*SystemPromptManagerTopUpPoolBalanceIterator, error) {

	logs, sub, err := _SystemPromptManager.contract.FilterLogs(opts, "TopUpPoolBalance")
	if err != nil {
		return nil, err
	}
	return &SystemPromptManagerTopUpPoolBalanceIterator{contract: _SystemPromptManager.contract, event: "TopUpPoolBalance", logs: logs, sub: sub}, nil
}

// WatchTopUpPoolBalance is a free log subscription operation binding the contract event 0xf7ee57effd30f2ab842c1d65fdcfa7d20c2fb2f967e9ac30acafecabf013ea4c.
//
// Solidity: event TopUpPoolBalance(uint256 agentId, address caller, uint256 amount)
func (_SystemPromptManager *SystemPromptManagerFilterer) WatchTopUpPoolBalance(opts *bind.WatchOpts, sink chan<- *SystemPromptManagerTopUpPoolBalance) (event.Subscription, error) {

	logs, sub, err := _SystemPromptManager.contract.WatchLogs(opts, "TopUpPoolBalance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemPromptManagerTopUpPoolBalance)
				if err := _SystemPromptManager.contract.UnpackLog(event, "TopUpPoolBalance", log); err != nil {
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
func (_SystemPromptManager *SystemPromptManagerFilterer) ParseTopUpPoolBalance(log types.Log) (*SystemPromptManagerTopUpPoolBalance, error) {
	event := new(SystemPromptManagerTopUpPoolBalance)
	if err := _SystemPromptManager.contract.UnpackLog(event, "TopUpPoolBalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemPromptManagerTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SystemPromptManager contract.
type SystemPromptManagerTransferIterator struct {
	Event *SystemPromptManagerTransfer // Event containing the contract specifics and raw log

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
func (it *SystemPromptManagerTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemPromptManagerTransfer)
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
		it.Event = new(SystemPromptManagerTransfer)
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
func (it *SystemPromptManagerTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemPromptManagerTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemPromptManagerTransfer represents a Transfer event raised by the SystemPromptManager contract.
type SystemPromptManagerTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SystemPromptManager *SystemPromptManagerFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*SystemPromptManagerTransferIterator, error) {

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

	logs, sub, err := _SystemPromptManager.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SystemPromptManagerTransferIterator{contract: _SystemPromptManager.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SystemPromptManager *SystemPromptManagerFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SystemPromptManagerTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _SystemPromptManager.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemPromptManagerTransfer)
				if err := _SystemPromptManager.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_SystemPromptManager *SystemPromptManagerFilterer) ParseTransfer(log types.Log) (*SystemPromptManagerTransfer, error) {
	event := new(SystemPromptManagerTransfer)
	if err := _SystemPromptManager.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemPromptManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SystemPromptManager contract.
type SystemPromptManagerUnpausedIterator struct {
	Event *SystemPromptManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *SystemPromptManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemPromptManagerUnpaused)
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
		it.Event = new(SystemPromptManagerUnpaused)
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
func (it *SystemPromptManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemPromptManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemPromptManagerUnpaused represents a Unpaused event raised by the SystemPromptManager contract.
type SystemPromptManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SystemPromptManager *SystemPromptManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SystemPromptManagerUnpausedIterator, error) {

	logs, sub, err := _SystemPromptManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SystemPromptManagerUnpausedIterator{contract: _SystemPromptManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SystemPromptManager *SystemPromptManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SystemPromptManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _SystemPromptManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemPromptManagerUnpaused)
				if err := _SystemPromptManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_SystemPromptManager *SystemPromptManagerFilterer) ParseUnpaused(log types.Log) (*SystemPromptManagerUnpaused, error) {
	event := new(SystemPromptManagerUnpaused)
	if err := _SystemPromptManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
