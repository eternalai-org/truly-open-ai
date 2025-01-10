// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aikb20

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

// EternalAIKB20MetaData contains all meta data concerning the EternalAIKB20 contract.
var EternalAIKB20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgentPromptIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidData\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"sysPrompt\",\"type\":\"bytes[]\"}],\"name\":\"AgentDataAddNew\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"promptIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"oldSysPrompt\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newSysPrompt\",\"type\":\"bytes\"}],\"name\":\"AgentDataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"AgentFeeUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"AgentURIUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gpuManager\",\"type\":\"address\"}],\"name\":\"GPUManagerUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"externalData\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inferenceId\",\"type\":\"uint256\"}],\"name\":\"InferencePerformed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"}],\"name\":\"ModelIdUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"promptScheduler\",\"type\":\"address\"}],\"name\":\"PromptSchedulerUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenFee\",\"type\":\"address\"}],\"name\":\"TokenFeeUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TopUpPoolBalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_gpuManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_modelId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_poolBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_promptScheduler\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signature\",\"type\":\"bytes32\"}],\"name\":\"_signaturesUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"promptKey\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"addNewAgentData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"promptKey\",\"type\":\"string\"}],\"name\":\"getAgentSystemPrompt\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inferData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"promptScheduler_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gpuManager_\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"modelId_\",\"type\":\"uint32\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"tokenFee_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fwdCalldata\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"externalData\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"promptKey\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"retrieve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fwdCalldata\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"externalData\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"promptKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"retrieve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"topUpPoolBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"updateAgentFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gpuManager\",\"type\":\"address\"}],\"name\":\"updateGPUManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"}],\"name\":\"updateModelId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"promptScheduler\",\"type\":\"address\"}],\"name\":\"updatePromptScheduler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080806040523461001657613149908161001c8239f35b600080fdfe608080604052600436101561001357600080fd5b60003560e01c90816306fdde03146120e15750806308c147fd1461208f578063095ea7b31461204b57806318160ddd1461200f5780631c29f41814611fbf57806323b872dd14611e9e57806325abc00214611e4c578063313ce56714611e125780633950935114611d955780633f4ba83a14611cfa5780635c975abb14611cb95780636097ff9014611c2e57806364058dc014611b7c57806364ff68f514611a1057806370a08231146119ab578063715018a61461190d5780638456cb591461186f57806388ee5fb2146117bd5780638da5cb5b1461176b57806395d89b4114611649578063971030a514611604578063a457c2d714611500578063a55d5b5d146113b2578063a9059cbb14611363578063c5dee8ce14610b34578063c82cd94b14610a8b578063c9fb89fe14610a4f578063dd0a9d9714610892578063dd62ed3e14610815578063eaf01882146107d8578063ed98f621146106ce578063f2fde38b146105e6578063fd9be522146104e45763fe465a9d1461019557600080fd5b346104df5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df5767ffffffffffffffff6004358181116104df576101e5903690600401612246565b6024929192358281116104df57610200903690600401612246565b93610209612274565b84156104b55760405193838286378484810160fc815260209687910301902095865468010000000000000000811015610457576001978882018082558210156104865760005286600020019181116104575785936102678354612360565b601f811161041c575b50600090601f831160011461037c576000919083610371575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82891b9260031b1c19161790555b826040519384928337810160fc81520301902091604051918083018184528454809152604084019160408260051b8601019560005280600020926000905b838210610327577fba498a6e4c1996a9a7e64f07fe99506267fde36ac9398c866036bd6cdcfbab6187890388a1005b909192939483816103608a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08b849d03018652896123b3565b9997019594939190910191016102f8565b013590503880610289565b9082917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08a94169185600052876000209260005b818110610404575084116103cc575b505050811b0190556102ba565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c199101351690553880806103bf565b8383013585558b998d979095019492830192016103b0565b610447908460005286600020601f850160051c81019188861061044d575b601f0160051c0190612971565b38610270565b909150819061043a565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60046040517fe1507814000000000000000000000000000000000000000000000000000000008152fd5b600080fd5b346104df5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df5761051b612200565b60243590610527612274565b610102908154928084106000146105e0575082915b8261054357005b610550836105de95612a5d565b905560ff546040517fa9059cbb00000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff92831660248201526044810193909352166105d982606481015b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101845283612468565b612a6a565b005b9161053c565b346104df5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df5761061d612200565b610625612274565b73ffffffffffffffffffffffffffffffffffffffff81161561064a576105de906122f3565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152fd5b346104df5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df5763ffffffff600435818116918282036104df57610718612274565b82159081156107c7575b5061079d577f7a55cba18b26688d885d840641dae0c8c4842f3cd33131a902de436afe2013ee916020917fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff77ffffffff000000000000000000000000000000000000000060fe549260a01b1691161760fe55604051908152a1005b60046040517f5cb045db000000000000000000000000000000000000000000000000000000008152fd5b905060fe5460a01c16821483610722565b346104df5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df57602061010054604051908152f35b346104df5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df5761084c612200565b610854612223565b9073ffffffffffffffffffffffffffffffffffffffff8091166000526034602052604060002091166000526020526020604060002054604051908152f35b346104df5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df5767ffffffffffffffff6004358181116104df576108e2903690600401612246565b91906024358281116104df576108fc903690600401612246565b939091604435918483116104df5761091b610925933690600401612246565b9160643593612da9565b92905060fe5492602060405180957fde1ce2bb00000000000000000000000000000000000000000000000000000000825263ffffffff8160a01c1660048301526060602483015281600073ffffffffffffffffffffffffffffffffffffffff8261099260648201896121a2565b3360448301520393165af1938415610a43576000946109ed575b506109e8907fd6f4605d9629a8ab6dc9e171ea3027456f2bb13d67203afd50552488bc90a3f3949560fb54946040519586953399169386612d44565b0390a2005b7fd6f4605d9629a8ab6dc9e171ea3027456f2bb13d67203afd50552488bc90a3f3945090610a346109e89260203d8111610a3c575b610a2c8183612468565b810190612d24565b9450906109ac565b503d610a22565b6040513d6000823e3d90fd5b346104df5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df57602060fb54604051908152f35b346104df5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df577f52d53b7ae2f89bb8b4d82cfa1c120c7a724ed4f14d15862f38b2841f2bbbbbaa610b2f600435610b078173ffffffffffffffffffffffffffffffffffffffff60ff541630903390612ca6565b610100610b158282546124a9565b905560408051338152602081019290925290918291820190565b0390a1005b346104df5760e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df5760043567ffffffffffffffff81116104df57610b83903690600401612246565b60243567ffffffffffffffff81116104df57610ba3903690600401612246565b9060643573ffffffffffffffffffffffffffffffffffffffff9283821682036104df576084359484861686036104df5763ffffffff60a4351660a435036104df5760c4359385851685036104df5760005460ff8160081c16159889809a611356575b801561133f575b156112bb57610c5893828b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00610c5096161760005561128c575b5036916128af565b9236916128af565b90610c7360ff60005460081c16610c6e816128e6565b6128e6565b80519067ffffffffffffffff8211610457578190610c92603654612360565b601f811161123d575b50602090601f83116001146111595760009261114e575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916176036555b80519067ffffffffffffffff8211610457578190610d04603754612360565b601f81116110ff575b50602090601f831160011461101b57600092611010575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916176037555b610da660ff60005460081c16610d6c816128e6565b610d75816128e6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060655416606555610c6e816128e6565b610daf336122f3565b610dc060ff60005460081c166128e6565b828116158015611006575b8015610ffc575b61079d5760fe54837fffffffffffffffffffffffff000000000000000000000000000000000000000095168560fd54161760fd557fffffffffffffffff0000000000000000000000000000000000000000000000008477ffffffff000000000000000000000000000000000000000060a43560a01b1693169116171760fe55169060ff54161760ff553315610f9e5760ff60655416610f1a57604435610e7a816035546124a9565b603555336000526033602052604060002081815401905560405190815260007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60203393a3610ec557005b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff600054166000557f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498602060405160018152a1005b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f45524332305061757361626c653a20746f6b656e207472616e7366657220776860448201527f696c6520706175736564000000000000000000000000000000000000000000006064820152fd5b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152fd5b5082821615610dd2565b5082841615610dcb565b015190508780610d24565b925060376000527f42a7b7dd785cd69714a189dffb3fd7d7174edc9ece837694ce50f7078f7c31ae906000935b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0841685106110e45760019450837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08116106110ad575b505050811b01603755610d57565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c1916905587808061109f565b81810151835560209485019460019093019290910190611048565b6111489060376000527f42a7b7dd785cd69714a189dffb3fd7d7174edc9ece837694ce50f7078f7c31ae601f850160051c8101916020861061044d57601f0160051c0190612971565b88610d0d565b015190508880610cb2565b925060366000527f4a11f94e20a93c79f6ec743a1954ec4fc2c08429ae2122118bf234b2185c81b8906000935b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0841685106112225760019450837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08116106111eb575b505050811b01603655610ce5565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558880806111dd565b81810151835560209485019460019093019290910190611186565b6112869060366000527f4a11f94e20a93c79f6ec743a1954ec4fc2c08429ae2122118bf234b2185c81b8601f850160051c8101916020861061044d57601f0160051c0190612971565b89610c9b565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016610101176000558b610c48565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152fd5b50303b158015610c0c5750600160ff831614610c0c565b50600160ff831610610c05565b346104df5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df576113a761139d612200565b60243590336124e5565b602060405160018152f35b346104df576020807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df5767ffffffffffffffff6004358181116104df5761140483913690600401612246565b9190826040519384928337810160fc815203019020805491821161045757829060405192611437838260051b0185612468565b8084528284018092600052836000206000915b8383106114cf5750505050604051918083019381845251809452604083019360408160051b85010192916000955b8287106114855785850386f35b9091929382806114bf837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08a6001960301865288516121a2565b9601920196019592919092611478565b60018681926040999899516114ef816114e881896123b3565b0382612468565b81520192019201919095949561144a565b346104df5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df57611537612200565b60243590336000526034602052604060002073ffffffffffffffffffffffffffffffffffffffff821660005260205260406000205491808310611580576113a792039033612700565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f0000000000000000000000000000000000000000000000000000006064820152fd5b346104df5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df57602063ffffffff60fe5460a01c16604051908152f35b346104df5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df57604051600060375461168981612360565b8084529060019081811690811561172657506001146116cb575b6116c7846116b381860382612468565b6040519182916020835260208301906121a2565b0390f35b6037600090815292507f42a7b7dd785cd69714a189dffb3fd7d7174edc9ece837694ce50f7078f7c31ae5b82841061170e5750505081016020016116b3826116a3565b805460208587018101919091529093019281016116f6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208087019190915292151560051b850190920192506116b391508390506116a3565b346104df5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df57602073ffffffffffffffffffffffffffffffffffffffff60c95416604051908152f35b346104df5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df5773ffffffffffffffffffffffffffffffffffffffff611809612200565b611811612274565b16801561079d576020817f752bc81ffc4d7e1886a090f40c08324c06359626e920d446f8deafb05622782a927fffffffffffffffffffffffff000000000000000000000000000000000000000060fd54161760fd55604051908152a1005b346104df5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df576118a6612274565b6118ae612988565b6118b6612988565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060655416176065557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b346104df5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df57611944612274565b600073ffffffffffffffffffffffffffffffffffffffff60c9547fffffffffffffffffffffffff0000000000000000000000000000000000000000811660c955167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346104df5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df5773ffffffffffffffffffffffffffffffffffffffff6119f7612200565b1660005260336020526020604060002054604051908152f35b346104df5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df5767ffffffffffffffff6004358181116104df57611a60903690600401612246565b6024358381116104df57611a78903690600401612246565b9390916044358281116104df57611a93903690600401612246565b90606435928315158094036104df5760009660209373ffffffffffffffffffffffffffffffffffffffff93611acb9360843593612da9565b93905060fe54906040519788938480937f5cc6873100000000000000000000000000000000000000000000000000000000825263ffffffff8660a01c16600483015260806024830152611b21608483018a6121a2565b9033604484015260648301520393165af1938415610a43576000946109ed57506109e8907fd6f4605d9629a8ab6dc9e171ea3027456f2bb13d67203afd50552488bc90a3f3949560fb54946040519586953399169386612d44565b346104df5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df5773ffffffffffffffffffffffffffffffffffffffff611bc8612200565b611bd0612274565b16801561079d576020817f667557d852582e84e7de441f650ea0aacbb7de26e3485436e0c27ba8d19a79f1927fffffffffffffffffffffffff000000000000000000000000000000000000000060fe54161760fe55604051908152a1005b346104df5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df577f34ae73c4d1b1bf2ec45030153dda7665cc119bcf159982735207f5e992cc98466020600435611c8b612274565b8060fb5403611c9e575b604051908152a1005b6fffffffffffffffffffffffffffffffff811660fb55611c95565b346104df5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df57602060ff606554166040519015158152f35b346104df5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df57611d31612274565b611d396129f2565b611d416129f2565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00606554166065557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b346104df5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df576113a7611dcf612200565b336000526034602052604060002073ffffffffffffffffffffffffffffffffffffffff8216600052602052611e0b6024356040600020546124a9565b9033612700565b346104df5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df57602060405160128152f35b346104df5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df57602073ffffffffffffffffffffffffffffffffffffffff60fe5416604051908152f35b346104df5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df57611ed5612200565b611edd612223565b6044359073ffffffffffffffffffffffffffffffffffffffff83166000526034602052604060002033600052602052604060002054927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8403611f45575b6113a793506124e5565b828410611f6157611f5c836113a795033383612700565b611f3b565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152fd5b346104df5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df57600435600052610101602052602060ff604060002054166040519015158152f35b346104df5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df576020603554604051908152f35b346104df5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df576113a7612085612200565b6024359033612700565b346104df5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df57602073ffffffffffffffffffffffffffffffffffffffff60fd5416604051908152f35b346104df5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104df57600060365461211e81612360565b808452906001908181169081156117265750600114612147576116c7846116b381860382612468565b6036600090815292507f4a11f94e20a93c79f6ec743a1954ec4fc2c08429ae2122118bf234b2185c81b85b82841061218a5750505081016020016116b3826116a3565b80546020858701810191909152909301928101612172565b919082519283825260005b8481106121ec5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b6020818301810151848301820152016121ad565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036104df57565b6024359073ffffffffffffffffffffffffffffffffffffffff821682036104df57565b9181601f840112156104df5782359167ffffffffffffffff83116104df57602083818601950101116104df57565b73ffffffffffffffffffffffffffffffffffffffff60c95416330361229557565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b60c9549073ffffffffffffffffffffffffffffffffffffffff80911691827fffffffffffffffffffffffff000000000000000000000000000000000000000082161760c955167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b90600182811c921680156123a9575b602083101461237a57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f169161236f565b90600092918054916123c483612360565b91828252600193848116908160001461242657506001146123e6575b50505050565b90919394506000526020928360002092846000945b8386106124125750505050010190388080806123e0565b8054858701830152940193859082016123fb565b91505060209495507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff009193501683830152151560051b010190388080806123e0565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761045757604052565b919082018092116124b657565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff80911691821561267c57169182156125f85760ff60655416610f1a576000828152603360205260408120549180831061257457604082827fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef95876020965260338652038282205586815220818154019055604051908152a3565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152fd5b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152fd5b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152fd5b73ffffffffffffffffffffffffffffffffffffffff8091169182156127f2571691821561276e5760207f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925918360005260348252604060002085600052825280604060002055604051908152a3565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152fd5b60846040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152fd5b67ffffffffffffffff811161045757601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b9291926128bb82612875565b916128c96040519384612468565b8294818452818301116104df578281602093846000960137010152565b156128ed57565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152fd5b81811061297c575050565b60008155600101612971565b60ff6065541661299457565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152fd5b60ff60655416156129ff57565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152fd5b919082039182116124b657565b73ffffffffffffffffffffffffffffffffffffffff1690604051604081019080821067ffffffffffffffff83111761045757612b08916040526020938482527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564858301526000808587829751910182855af13d15612bd2573d91612aed83612875565b92612afb6040519485612468565b83523d868885013e612bd6565b805191821591848315612ba7575b505050905015612b235750565b608490604051907f08c379a00000000000000000000000000000000000000000000000000000000082526004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b919381809450010312612bce57820151908115158203612bcb575080388084612b16565b80fd5b5080fd5b6060915b91929015612c515750815115612bea575090565b3b15612bf35790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015612c645750805190602001fd5b612ca2906040519182917f08c379a00000000000000000000000000000000000000000000000000000000083526020600484015260248301906121a2565b0390fd5b9290604051927f23b872dd00000000000000000000000000000000000000000000000000000000602085015273ffffffffffffffffffffffffffffffffffffffff809216602485015216604483015260648201526064815260a081019181831067ffffffffffffffff84111761045757612d2292604052612a6a565b565b908160209103126104df575167ffffffffffffffff811681036104df5790565b959493601f817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe092606096612d846020979660808d5260808d01906121a2565b95878c01528a860360408c015281865286860137600085828601015201160101930152565b9194929390936040958187519283928337810160fc815260209283910301902054156130ea5760fb5482106130c15773ffffffffffffffffffffffffffffffffffffffff90612e00838360ff541630903390612ca6565b8160fd5416928163ffffffff60fe5460a01c1660248a51809781937f963a027800000000000000000000000000000000000000000000000000000000835260048301525afa9384156130b657600094613087575b50838110808061307a575b1561302157506101008481540390558061300b575b505b8160ff54169160fe541683158015612f83575b15612f0057916105d97f095ea7b300000000000000000000000000000000000000000000000000000000926105ad86612ef596612efd9a9b9c51968794850152602484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b9336916128af565b90565b6084828951907f08c379a00000000000000000000000000000000000000000000000000000000082526004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e6365000000000000000000006064820152fd5b5087517fdd62ed3e0000000000000000000000000000000000000000000000000000000081523060048201528160248201528281604481875afa90811561300057600091612fd3575b5015612e89565b908382813d8311612ff9575b612fe98183612468565b81010312612bcb57505138612fcc565b503d612fdf565b89513d6000823e3d90fd5b6130196101029182546124a9565b905538612e74565b613051578361302f91612a5d565b8061303b575b50612e76565b6130496101029182546124a9565b905538613035565b600488517f356680b7000000000000000000000000000000000000000000000000000000008152fd5b5084610100541015612e5f565b90938282813d83116130af575b61309e8183612468565b81010312612bcb5750519238612e54565b503d613094565b88513d6000823e3d90fd5b600486517fbf94a257000000000000000000000000000000000000000000000000000000008152fd5b600486517fe1507814000000000000000000000000000000000000000000000000000000008152fdfea264697066735822122080bfbde80a334d61847137824587082935b32aad31f1cf2c9a9ea1866af2cb7564736f6c63430008140033",
}

// EternalAIKB20ABI is the input ABI used to generate the binding from.
// Deprecated: Use EternalAIKB20MetaData.ABI instead.
var EternalAIKB20ABI = EternalAIKB20MetaData.ABI

// EternalAIKB20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EternalAIKB20MetaData.Bin instead.
var EternalAIKB20Bin = EternalAIKB20MetaData.Bin

// DeployEternalAIKB20 deploys a new Ethereum contract, binding an instance of EternalAIKB20 to it.
func DeployEternalAIKB20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EternalAIKB20, error) {
	parsed, err := EternalAIKB20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EternalAIKB20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EternalAIKB20{EternalAIKB20Caller: EternalAIKB20Caller{contract: contract}, EternalAIKB20Transactor: EternalAIKB20Transactor{contract: contract}, EternalAIKB20Filterer: EternalAIKB20Filterer{contract: contract}}, nil
}

// EternalAIKB20 is an auto generated Go binding around an Ethereum contract.
type EternalAIKB20 struct {
	EternalAIKB20Caller     // Read-only binding to the contract
	EternalAIKB20Transactor // Write-only binding to the contract
	EternalAIKB20Filterer   // Log filterer for contract events
}

// EternalAIKB20Caller is an auto generated read-only Go binding around an Ethereum contract.
type EternalAIKB20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EternalAIKB20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type EternalAIKB20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EternalAIKB20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EternalAIKB20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EternalAIKB20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EternalAIKB20Session struct {
	Contract     *EternalAIKB20    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EternalAIKB20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EternalAIKB20CallerSession struct {
	Contract *EternalAIKB20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EternalAIKB20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EternalAIKB20TransactorSession struct {
	Contract     *EternalAIKB20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EternalAIKB20Raw is an auto generated low-level Go binding around an Ethereum contract.
type EternalAIKB20Raw struct {
	Contract *EternalAIKB20 // Generic contract binding to access the raw methods on
}

// EternalAIKB20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EternalAIKB20CallerRaw struct {
	Contract *EternalAIKB20Caller // Generic read-only contract binding to access the raw methods on
}

// EternalAIKB20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EternalAIKB20TransactorRaw struct {
	Contract *EternalAIKB20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewEternalAIKB20 creates a new instance of EternalAIKB20, bound to a specific deployed contract.
func NewEternalAIKB20(address common.Address, backend bind.ContractBackend) (*EternalAIKB20, error) {
	contract, err := bindEternalAIKB20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EternalAIKB20{EternalAIKB20Caller: EternalAIKB20Caller{contract: contract}, EternalAIKB20Transactor: EternalAIKB20Transactor{contract: contract}, EternalAIKB20Filterer: EternalAIKB20Filterer{contract: contract}}, nil
}

// NewEternalAIKB20Caller creates a new read-only instance of EternalAIKB20, bound to a specific deployed contract.
func NewEternalAIKB20Caller(address common.Address, caller bind.ContractCaller) (*EternalAIKB20Caller, error) {
	contract, err := bindEternalAIKB20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EternalAIKB20Caller{contract: contract}, nil
}

// NewEternalAIKB20Transactor creates a new write-only instance of EternalAIKB20, bound to a specific deployed contract.
func NewEternalAIKB20Transactor(address common.Address, transactor bind.ContractTransactor) (*EternalAIKB20Transactor, error) {
	contract, err := bindEternalAIKB20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EternalAIKB20Transactor{contract: contract}, nil
}

// NewEternalAIKB20Filterer creates a new log filterer instance of EternalAIKB20, bound to a specific deployed contract.
func NewEternalAIKB20Filterer(address common.Address, filterer bind.ContractFilterer) (*EternalAIKB20Filterer, error) {
	contract, err := bindEternalAIKB20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EternalAIKB20Filterer{contract: contract}, nil
}

// bindEternalAIKB20 binds a generic wrapper to an already deployed contract.
func bindEternalAIKB20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EternalAIKB20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EternalAIKB20 *EternalAIKB20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EternalAIKB20.Contract.EternalAIKB20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EternalAIKB20 *EternalAIKB20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.EternalAIKB20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EternalAIKB20 *EternalAIKB20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.EternalAIKB20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EternalAIKB20 *EternalAIKB20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EternalAIKB20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EternalAIKB20 *EternalAIKB20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EternalAIKB20 *EternalAIKB20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.contract.Transact(opts, method, params...)
}

// GpuManager is a free data retrieval call binding the contract method 0x08c147fd.
//
// Solidity: function _gpuManager() view returns(address)
func (_EternalAIKB20 *EternalAIKB20Caller) GpuManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EternalAIKB20.contract.Call(opts, &out, "_gpuManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GpuManager is a free data retrieval call binding the contract method 0x08c147fd.
//
// Solidity: function _gpuManager() view returns(address)
func (_EternalAIKB20 *EternalAIKB20Session) GpuManager() (common.Address, error) {
	return _EternalAIKB20.Contract.GpuManager(&_EternalAIKB20.CallOpts)
}

// GpuManager is a free data retrieval call binding the contract method 0x08c147fd.
//
// Solidity: function _gpuManager() view returns(address)
func (_EternalAIKB20 *EternalAIKB20CallerSession) GpuManager() (common.Address, error) {
	return _EternalAIKB20.Contract.GpuManager(&_EternalAIKB20.CallOpts)
}

// ModelId is a free data retrieval call binding the contract method 0x971030a5.
//
// Solidity: function _modelId() view returns(uint32)
func (_EternalAIKB20 *EternalAIKB20Caller) ModelId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _EternalAIKB20.contract.Call(opts, &out, "_modelId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ModelId is a free data retrieval call binding the contract method 0x971030a5.
//
// Solidity: function _modelId() view returns(uint32)
func (_EternalAIKB20 *EternalAIKB20Session) ModelId() (uint32, error) {
	return _EternalAIKB20.Contract.ModelId(&_EternalAIKB20.CallOpts)
}

// ModelId is a free data retrieval call binding the contract method 0x971030a5.
//
// Solidity: function _modelId() view returns(uint32)
func (_EternalAIKB20 *EternalAIKB20CallerSession) ModelId() (uint32, error) {
	return _EternalAIKB20.Contract.ModelId(&_EternalAIKB20.CallOpts)
}

// PoolBalance is a free data retrieval call binding the contract method 0xeaf01882.
//
// Solidity: function _poolBalance() view returns(uint256)
func (_EternalAIKB20 *EternalAIKB20Caller) PoolBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EternalAIKB20.contract.Call(opts, &out, "_poolBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolBalance is a free data retrieval call binding the contract method 0xeaf01882.
//
// Solidity: function _poolBalance() view returns(uint256)
func (_EternalAIKB20 *EternalAIKB20Session) PoolBalance() (*big.Int, error) {
	return _EternalAIKB20.Contract.PoolBalance(&_EternalAIKB20.CallOpts)
}

// PoolBalance is a free data retrieval call binding the contract method 0xeaf01882.
//
// Solidity: function _poolBalance() view returns(uint256)
func (_EternalAIKB20 *EternalAIKB20CallerSession) PoolBalance() (*big.Int, error) {
	return _EternalAIKB20.Contract.PoolBalance(&_EternalAIKB20.CallOpts)
}

// PromptScheduler is a free data retrieval call binding the contract method 0x25abc002.
//
// Solidity: function _promptScheduler() view returns(address)
func (_EternalAIKB20 *EternalAIKB20Caller) PromptScheduler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EternalAIKB20.contract.Call(opts, &out, "_promptScheduler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PromptScheduler is a free data retrieval call binding the contract method 0x25abc002.
//
// Solidity: function _promptScheduler() view returns(address)
func (_EternalAIKB20 *EternalAIKB20Session) PromptScheduler() (common.Address, error) {
	return _EternalAIKB20.Contract.PromptScheduler(&_EternalAIKB20.CallOpts)
}

// PromptScheduler is a free data retrieval call binding the contract method 0x25abc002.
//
// Solidity: function _promptScheduler() view returns(address)
func (_EternalAIKB20 *EternalAIKB20CallerSession) PromptScheduler() (common.Address, error) {
	return _EternalAIKB20.Contract.PromptScheduler(&_EternalAIKB20.CallOpts)
}

// SignaturesUsed is a free data retrieval call binding the contract method 0x1c29f418.
//
// Solidity: function _signaturesUsed(bytes32 signature) view returns(bool)
func (_EternalAIKB20 *EternalAIKB20Caller) SignaturesUsed(opts *bind.CallOpts, signature [32]byte) (bool, error) {
	var out []interface{}
	err := _EternalAIKB20.contract.Call(opts, &out, "_signaturesUsed", signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SignaturesUsed is a free data retrieval call binding the contract method 0x1c29f418.
//
// Solidity: function _signaturesUsed(bytes32 signature) view returns(bool)
func (_EternalAIKB20 *EternalAIKB20Session) SignaturesUsed(signature [32]byte) (bool, error) {
	return _EternalAIKB20.Contract.SignaturesUsed(&_EternalAIKB20.CallOpts, signature)
}

// SignaturesUsed is a free data retrieval call binding the contract method 0x1c29f418.
//
// Solidity: function _signaturesUsed(bytes32 signature) view returns(bool)
func (_EternalAIKB20 *EternalAIKB20CallerSession) SignaturesUsed(signature [32]byte) (bool, error) {
	return _EternalAIKB20.Contract.SignaturesUsed(&_EternalAIKB20.CallOpts, signature)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_EternalAIKB20 *EternalAIKB20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EternalAIKB20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_EternalAIKB20 *EternalAIKB20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _EternalAIKB20.Contract.Allowance(&_EternalAIKB20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_EternalAIKB20 *EternalAIKB20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _EternalAIKB20.Contract.Allowance(&_EternalAIKB20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EternalAIKB20 *EternalAIKB20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EternalAIKB20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EternalAIKB20 *EternalAIKB20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _EternalAIKB20.Contract.BalanceOf(&_EternalAIKB20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EternalAIKB20 *EternalAIKB20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _EternalAIKB20.Contract.BalanceOf(&_EternalAIKB20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EternalAIKB20 *EternalAIKB20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _EternalAIKB20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EternalAIKB20 *EternalAIKB20Session) Decimals() (uint8, error) {
	return _EternalAIKB20.Contract.Decimals(&_EternalAIKB20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EternalAIKB20 *EternalAIKB20CallerSession) Decimals() (uint8, error) {
	return _EternalAIKB20.Contract.Decimals(&_EternalAIKB20.CallOpts)
}

// GetAgentSystemPrompt is a free data retrieval call binding the contract method 0xa55d5b5d.
//
// Solidity: function getAgentSystemPrompt(string promptKey) view returns(bytes[])
func (_EternalAIKB20 *EternalAIKB20Caller) GetAgentSystemPrompt(opts *bind.CallOpts, promptKey string) ([][]byte, error) {
	var out []interface{}
	err := _EternalAIKB20.contract.Call(opts, &out, "getAgentSystemPrompt", promptKey)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetAgentSystemPrompt is a free data retrieval call binding the contract method 0xa55d5b5d.
//
// Solidity: function getAgentSystemPrompt(string promptKey) view returns(bytes[])
func (_EternalAIKB20 *EternalAIKB20Session) GetAgentSystemPrompt(promptKey string) ([][]byte, error) {
	return _EternalAIKB20.Contract.GetAgentSystemPrompt(&_EternalAIKB20.CallOpts, promptKey)
}

// GetAgentSystemPrompt is a free data retrieval call binding the contract method 0xa55d5b5d.
//
// Solidity: function getAgentSystemPrompt(string promptKey) view returns(bytes[])
func (_EternalAIKB20 *EternalAIKB20CallerSession) GetAgentSystemPrompt(promptKey string) ([][]byte, error) {
	return _EternalAIKB20.Contract.GetAgentSystemPrompt(&_EternalAIKB20.CallOpts, promptKey)
}

// InferData is a free data retrieval call binding the contract method 0xc9fb89fe.
//
// Solidity: function inferData() view returns(uint256)
func (_EternalAIKB20 *EternalAIKB20Caller) InferData(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EternalAIKB20.contract.Call(opts, &out, "inferData")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InferData is a free data retrieval call binding the contract method 0xc9fb89fe.
//
// Solidity: function inferData() view returns(uint256)
func (_EternalAIKB20 *EternalAIKB20Session) InferData() (*big.Int, error) {
	return _EternalAIKB20.Contract.InferData(&_EternalAIKB20.CallOpts)
}

// InferData is a free data retrieval call binding the contract method 0xc9fb89fe.
//
// Solidity: function inferData() view returns(uint256)
func (_EternalAIKB20 *EternalAIKB20CallerSession) InferData() (*big.Int, error) {
	return _EternalAIKB20.Contract.InferData(&_EternalAIKB20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EternalAIKB20 *EternalAIKB20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EternalAIKB20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EternalAIKB20 *EternalAIKB20Session) Name() (string, error) {
	return _EternalAIKB20.Contract.Name(&_EternalAIKB20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EternalAIKB20 *EternalAIKB20CallerSession) Name() (string, error) {
	return _EternalAIKB20.Contract.Name(&_EternalAIKB20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EternalAIKB20 *EternalAIKB20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EternalAIKB20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EternalAIKB20 *EternalAIKB20Session) Owner() (common.Address, error) {
	return _EternalAIKB20.Contract.Owner(&_EternalAIKB20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EternalAIKB20 *EternalAIKB20CallerSession) Owner() (common.Address, error) {
	return _EternalAIKB20.Contract.Owner(&_EternalAIKB20.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EternalAIKB20 *EternalAIKB20Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EternalAIKB20.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EternalAIKB20 *EternalAIKB20Session) Paused() (bool, error) {
	return _EternalAIKB20.Contract.Paused(&_EternalAIKB20.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EternalAIKB20 *EternalAIKB20CallerSession) Paused() (bool, error) {
	return _EternalAIKB20.Contract.Paused(&_EternalAIKB20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EternalAIKB20 *EternalAIKB20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EternalAIKB20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EternalAIKB20 *EternalAIKB20Session) Symbol() (string, error) {
	return _EternalAIKB20.Contract.Symbol(&_EternalAIKB20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EternalAIKB20 *EternalAIKB20CallerSession) Symbol() (string, error) {
	return _EternalAIKB20.Contract.Symbol(&_EternalAIKB20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EternalAIKB20 *EternalAIKB20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EternalAIKB20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EternalAIKB20 *EternalAIKB20Session) TotalSupply() (*big.Int, error) {
	return _EternalAIKB20.Contract.TotalSupply(&_EternalAIKB20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EternalAIKB20 *EternalAIKB20CallerSession) TotalSupply() (*big.Int, error) {
	return _EternalAIKB20.Contract.TotalSupply(&_EternalAIKB20.CallOpts)
}

// AddNewAgentData is a paid mutator transaction binding the contract method 0xfe465a9d.
//
// Solidity: function addNewAgentData(string promptKey, bytes data) returns()
func (_EternalAIKB20 *EternalAIKB20Transactor) AddNewAgentData(opts *bind.TransactOpts, promptKey string, data []byte) (*types.Transaction, error) {
	return _EternalAIKB20.contract.Transact(opts, "addNewAgentData", promptKey, data)
}

// AddNewAgentData is a paid mutator transaction binding the contract method 0xfe465a9d.
//
// Solidity: function addNewAgentData(string promptKey, bytes data) returns()
func (_EternalAIKB20 *EternalAIKB20Session) AddNewAgentData(promptKey string, data []byte) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.AddNewAgentData(&_EternalAIKB20.TransactOpts, promptKey, data)
}

// AddNewAgentData is a paid mutator transaction binding the contract method 0xfe465a9d.
//
// Solidity: function addNewAgentData(string promptKey, bytes data) returns()
func (_EternalAIKB20 *EternalAIKB20TransactorSession) AddNewAgentData(promptKey string, data []byte) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.AddNewAgentData(&_EternalAIKB20.TransactOpts, promptKey, data)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_EternalAIKB20 *EternalAIKB20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_EternalAIKB20 *EternalAIKB20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.Approve(&_EternalAIKB20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_EternalAIKB20 *EternalAIKB20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.Approve(&_EternalAIKB20.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_EternalAIKB20 *EternalAIKB20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_EternalAIKB20 *EternalAIKB20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.DecreaseAllowance(&_EternalAIKB20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_EternalAIKB20 *EternalAIKB20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.DecreaseAllowance(&_EternalAIKB20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_EternalAIKB20 *EternalAIKB20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_EternalAIKB20 *EternalAIKB20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.IncreaseAllowance(&_EternalAIKB20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_EternalAIKB20 *EternalAIKB20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.IncreaseAllowance(&_EternalAIKB20.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xc5dee8ce.
//
// Solidity: function initialize(string name_, string symbol_, uint256 mintAmount_, address promptScheduler_, address gpuManager_, uint32 modelId_, address tokenFee_) returns()
func (_EternalAIKB20 *EternalAIKB20Transactor) Initialize(opts *bind.TransactOpts, name_ string, symbol_ string, mintAmount_ *big.Int, promptScheduler_ common.Address, gpuManager_ common.Address, modelId_ uint32, tokenFee_ common.Address) (*types.Transaction, error) {
	return _EternalAIKB20.contract.Transact(opts, "initialize", name_, symbol_, mintAmount_, promptScheduler_, gpuManager_, modelId_, tokenFee_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc5dee8ce.
//
// Solidity: function initialize(string name_, string symbol_, uint256 mintAmount_, address promptScheduler_, address gpuManager_, uint32 modelId_, address tokenFee_) returns()
func (_EternalAIKB20 *EternalAIKB20Session) Initialize(name_ string, symbol_ string, mintAmount_ *big.Int, promptScheduler_ common.Address, gpuManager_ common.Address, modelId_ uint32, tokenFee_ common.Address) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.Initialize(&_EternalAIKB20.TransactOpts, name_, symbol_, mintAmount_, promptScheduler_, gpuManager_, modelId_, tokenFee_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc5dee8ce.
//
// Solidity: function initialize(string name_, string symbol_, uint256 mintAmount_, address promptScheduler_, address gpuManager_, uint32 modelId_, address tokenFee_) returns()
func (_EternalAIKB20 *EternalAIKB20TransactorSession) Initialize(name_ string, symbol_ string, mintAmount_ *big.Int, promptScheduler_ common.Address, gpuManager_ common.Address, modelId_ uint32, tokenFee_ common.Address) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.Initialize(&_EternalAIKB20.TransactOpts, name_, symbol_, mintAmount_, promptScheduler_, gpuManager_, modelId_, tokenFee_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_EternalAIKB20 *EternalAIKB20Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EternalAIKB20.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_EternalAIKB20 *EternalAIKB20Session) Pause() (*types.Transaction, error) {
	return _EternalAIKB20.Contract.Pause(&_EternalAIKB20.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_EternalAIKB20 *EternalAIKB20TransactorSession) Pause() (*types.Transaction, error) {
	return _EternalAIKB20.Contract.Pause(&_EternalAIKB20.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EternalAIKB20 *EternalAIKB20Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EternalAIKB20.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EternalAIKB20 *EternalAIKB20Session) RenounceOwnership() (*types.Transaction, error) {
	return _EternalAIKB20.Contract.RenounceOwnership(&_EternalAIKB20.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EternalAIKB20 *EternalAIKB20TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EternalAIKB20.Contract.RenounceOwnership(&_EternalAIKB20.TransactOpts)
}

// Retrieve is a paid mutator transaction binding the contract method 0x64ff68f5.
//
// Solidity: function retrieve(bytes fwdCalldata, string externalData, string promptKey, bool flag, uint256 feeAmount) returns()
func (_EternalAIKB20 *EternalAIKB20Transactor) Retrieve(opts *bind.TransactOpts, fwdCalldata []byte, externalData string, promptKey string, flag bool, feeAmount *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.contract.Transact(opts, "retrieve", fwdCalldata, externalData, promptKey, flag, feeAmount)
}

// Retrieve is a paid mutator transaction binding the contract method 0x64ff68f5.
//
// Solidity: function retrieve(bytes fwdCalldata, string externalData, string promptKey, bool flag, uint256 feeAmount) returns()
func (_EternalAIKB20 *EternalAIKB20Session) Retrieve(fwdCalldata []byte, externalData string, promptKey string, flag bool, feeAmount *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.Retrieve(&_EternalAIKB20.TransactOpts, fwdCalldata, externalData, promptKey, flag, feeAmount)
}

// Retrieve is a paid mutator transaction binding the contract method 0x64ff68f5.
//
// Solidity: function retrieve(bytes fwdCalldata, string externalData, string promptKey, bool flag, uint256 feeAmount) returns()
func (_EternalAIKB20 *EternalAIKB20TransactorSession) Retrieve(fwdCalldata []byte, externalData string, promptKey string, flag bool, feeAmount *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.Retrieve(&_EternalAIKB20.TransactOpts, fwdCalldata, externalData, promptKey, flag, feeAmount)
}

// Retrieve0 is a paid mutator transaction binding the contract method 0xdd0a9d97.
//
// Solidity: function retrieve(bytes fwdCalldata, string externalData, string promptKey, uint256 feeAmount) returns()
func (_EternalAIKB20 *EternalAIKB20Transactor) Retrieve0(opts *bind.TransactOpts, fwdCalldata []byte, externalData string, promptKey string, feeAmount *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.contract.Transact(opts, "retrieve0", fwdCalldata, externalData, promptKey, feeAmount)
}

// Retrieve0 is a paid mutator transaction binding the contract method 0xdd0a9d97.
//
// Solidity: function retrieve(bytes fwdCalldata, string externalData, string promptKey, uint256 feeAmount) returns()
func (_EternalAIKB20 *EternalAIKB20Session) Retrieve0(fwdCalldata []byte, externalData string, promptKey string, feeAmount *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.Retrieve0(&_EternalAIKB20.TransactOpts, fwdCalldata, externalData, promptKey, feeAmount)
}

// Retrieve0 is a paid mutator transaction binding the contract method 0xdd0a9d97.
//
// Solidity: function retrieve(bytes fwdCalldata, string externalData, string promptKey, uint256 feeAmount) returns()
func (_EternalAIKB20 *EternalAIKB20TransactorSession) Retrieve0(fwdCalldata []byte, externalData string, promptKey string, feeAmount *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.Retrieve0(&_EternalAIKB20.TransactOpts, fwdCalldata, externalData, promptKey, feeAmount)
}

// TopUpPoolBalance is a paid mutator transaction binding the contract method 0xc82cd94b.
//
// Solidity: function topUpPoolBalance(uint256 amount) returns()
func (_EternalAIKB20 *EternalAIKB20Transactor) TopUpPoolBalance(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.contract.Transact(opts, "topUpPoolBalance", amount)
}

// TopUpPoolBalance is a paid mutator transaction binding the contract method 0xc82cd94b.
//
// Solidity: function topUpPoolBalance(uint256 amount) returns()
func (_EternalAIKB20 *EternalAIKB20Session) TopUpPoolBalance(amount *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.TopUpPoolBalance(&_EternalAIKB20.TransactOpts, amount)
}

// TopUpPoolBalance is a paid mutator transaction binding the contract method 0xc82cd94b.
//
// Solidity: function topUpPoolBalance(uint256 amount) returns()
func (_EternalAIKB20 *EternalAIKB20TransactorSession) TopUpPoolBalance(amount *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.TopUpPoolBalance(&_EternalAIKB20.TransactOpts, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_EternalAIKB20 *EternalAIKB20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_EternalAIKB20 *EternalAIKB20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.Transfer(&_EternalAIKB20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_EternalAIKB20 *EternalAIKB20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.Transfer(&_EternalAIKB20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_EternalAIKB20 *EternalAIKB20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_EternalAIKB20 *EternalAIKB20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.TransferFrom(&_EternalAIKB20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_EternalAIKB20 *EternalAIKB20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.TransferFrom(&_EternalAIKB20.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EternalAIKB20 *EternalAIKB20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EternalAIKB20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EternalAIKB20 *EternalAIKB20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.TransferOwnership(&_EternalAIKB20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EternalAIKB20 *EternalAIKB20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.TransferOwnership(&_EternalAIKB20.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_EternalAIKB20 *EternalAIKB20Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EternalAIKB20.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_EternalAIKB20 *EternalAIKB20Session) Unpause() (*types.Transaction, error) {
	return _EternalAIKB20.Contract.Unpause(&_EternalAIKB20.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_EternalAIKB20 *EternalAIKB20TransactorSession) Unpause() (*types.Transaction, error) {
	return _EternalAIKB20.Contract.Unpause(&_EternalAIKB20.TransactOpts)
}

// UpdateAgentFee is a paid mutator transaction binding the contract method 0x6097ff90.
//
// Solidity: function updateAgentFee(uint256 newFee) returns()
func (_EternalAIKB20 *EternalAIKB20Transactor) UpdateAgentFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.contract.Transact(opts, "updateAgentFee", newFee)
}

// UpdateAgentFee is a paid mutator transaction binding the contract method 0x6097ff90.
//
// Solidity: function updateAgentFee(uint256 newFee) returns()
func (_EternalAIKB20 *EternalAIKB20Session) UpdateAgentFee(newFee *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.UpdateAgentFee(&_EternalAIKB20.TransactOpts, newFee)
}

// UpdateAgentFee is a paid mutator transaction binding the contract method 0x6097ff90.
//
// Solidity: function updateAgentFee(uint256 newFee) returns()
func (_EternalAIKB20 *EternalAIKB20TransactorSession) UpdateAgentFee(newFee *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.UpdateAgentFee(&_EternalAIKB20.TransactOpts, newFee)
}

// UpdateGPUManager is a paid mutator transaction binding the contract method 0x88ee5fb2.
//
// Solidity: function updateGPUManager(address gpuManager) returns()
func (_EternalAIKB20 *EternalAIKB20Transactor) UpdateGPUManager(opts *bind.TransactOpts, gpuManager common.Address) (*types.Transaction, error) {
	return _EternalAIKB20.contract.Transact(opts, "updateGPUManager", gpuManager)
}

// UpdateGPUManager is a paid mutator transaction binding the contract method 0x88ee5fb2.
//
// Solidity: function updateGPUManager(address gpuManager) returns()
func (_EternalAIKB20 *EternalAIKB20Session) UpdateGPUManager(gpuManager common.Address) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.UpdateGPUManager(&_EternalAIKB20.TransactOpts, gpuManager)
}

// UpdateGPUManager is a paid mutator transaction binding the contract method 0x88ee5fb2.
//
// Solidity: function updateGPUManager(address gpuManager) returns()
func (_EternalAIKB20 *EternalAIKB20TransactorSession) UpdateGPUManager(gpuManager common.Address) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.UpdateGPUManager(&_EternalAIKB20.TransactOpts, gpuManager)
}

// UpdateModelId is a paid mutator transaction binding the contract method 0xed98f621.
//
// Solidity: function updateModelId(uint32 modelId) returns()
func (_EternalAIKB20 *EternalAIKB20Transactor) UpdateModelId(opts *bind.TransactOpts, modelId uint32) (*types.Transaction, error) {
	return _EternalAIKB20.contract.Transact(opts, "updateModelId", modelId)
}

// UpdateModelId is a paid mutator transaction binding the contract method 0xed98f621.
//
// Solidity: function updateModelId(uint32 modelId) returns()
func (_EternalAIKB20 *EternalAIKB20Session) UpdateModelId(modelId uint32) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.UpdateModelId(&_EternalAIKB20.TransactOpts, modelId)
}

// UpdateModelId is a paid mutator transaction binding the contract method 0xed98f621.
//
// Solidity: function updateModelId(uint32 modelId) returns()
func (_EternalAIKB20 *EternalAIKB20TransactorSession) UpdateModelId(modelId uint32) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.UpdateModelId(&_EternalAIKB20.TransactOpts, modelId)
}

// UpdatePromptScheduler is a paid mutator transaction binding the contract method 0x64058dc0.
//
// Solidity: function updatePromptScheduler(address promptScheduler) returns()
func (_EternalAIKB20 *EternalAIKB20Transactor) UpdatePromptScheduler(opts *bind.TransactOpts, promptScheduler common.Address) (*types.Transaction, error) {
	return _EternalAIKB20.contract.Transact(opts, "updatePromptScheduler", promptScheduler)
}

// UpdatePromptScheduler is a paid mutator transaction binding the contract method 0x64058dc0.
//
// Solidity: function updatePromptScheduler(address promptScheduler) returns()
func (_EternalAIKB20 *EternalAIKB20Session) UpdatePromptScheduler(promptScheduler common.Address) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.UpdatePromptScheduler(&_EternalAIKB20.TransactOpts, promptScheduler)
}

// UpdatePromptScheduler is a paid mutator transaction binding the contract method 0x64058dc0.
//
// Solidity: function updatePromptScheduler(address promptScheduler) returns()
func (_EternalAIKB20 *EternalAIKB20TransactorSession) UpdatePromptScheduler(promptScheduler common.Address) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.UpdatePromptScheduler(&_EternalAIKB20.TransactOpts, promptScheduler)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xfd9be522.
//
// Solidity: function withdrawFee(address recipient, uint256 amount) returns()
func (_EternalAIKB20 *EternalAIKB20Transactor) WithdrawFee(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.contract.Transact(opts, "withdrawFee", recipient, amount)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xfd9be522.
//
// Solidity: function withdrawFee(address recipient, uint256 amount) returns()
func (_EternalAIKB20 *EternalAIKB20Session) WithdrawFee(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.WithdrawFee(&_EternalAIKB20.TransactOpts, recipient, amount)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xfd9be522.
//
// Solidity: function withdrawFee(address recipient, uint256 amount) returns()
func (_EternalAIKB20 *EternalAIKB20TransactorSession) WithdrawFee(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EternalAIKB20.Contract.WithdrawFee(&_EternalAIKB20.TransactOpts, recipient, amount)
}

// EternalAIKB20AgentDataAddNewIterator is returned from FilterAgentDataAddNew and is used to iterate over the raw logs and unpacked data for AgentDataAddNew events raised by the EternalAIKB20 contract.
type EternalAIKB20AgentDataAddNewIterator struct {
	Event *EternalAIKB20AgentDataAddNew // Event containing the contract specifics and raw log

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
func (it *EternalAIKB20AgentDataAddNewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EternalAIKB20AgentDataAddNew)
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
		it.Event = new(EternalAIKB20AgentDataAddNew)
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
func (it *EternalAIKB20AgentDataAddNewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EternalAIKB20AgentDataAddNewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EternalAIKB20AgentDataAddNew represents a AgentDataAddNew event raised by the EternalAIKB20 contract.
type EternalAIKB20AgentDataAddNew struct {
	SysPrompt [][]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAgentDataAddNew is a free log retrieval operation binding the contract event 0xba498a6e4c1996a9a7e64f07fe99506267fde36ac9398c866036bd6cdcfbab61.
//
// Solidity: event AgentDataAddNew(bytes[] sysPrompt)
func (_EternalAIKB20 *EternalAIKB20Filterer) FilterAgentDataAddNew(opts *bind.FilterOpts) (*EternalAIKB20AgentDataAddNewIterator, error) {

	logs, sub, err := _EternalAIKB20.contract.FilterLogs(opts, "AgentDataAddNew")
	if err != nil {
		return nil, err
	}
	return &EternalAIKB20AgentDataAddNewIterator{contract: _EternalAIKB20.contract, event: "AgentDataAddNew", logs: logs, sub: sub}, nil
}

// WatchAgentDataAddNew is a free log subscription operation binding the contract event 0xba498a6e4c1996a9a7e64f07fe99506267fde36ac9398c866036bd6cdcfbab61.
//
// Solidity: event AgentDataAddNew(bytes[] sysPrompt)
func (_EternalAIKB20 *EternalAIKB20Filterer) WatchAgentDataAddNew(opts *bind.WatchOpts, sink chan<- *EternalAIKB20AgentDataAddNew) (event.Subscription, error) {

	logs, sub, err := _EternalAIKB20.contract.WatchLogs(opts, "AgentDataAddNew")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EternalAIKB20AgentDataAddNew)
				if err := _EternalAIKB20.contract.UnpackLog(event, "AgentDataAddNew", log); err != nil {
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

// ParseAgentDataAddNew is a log parse operation binding the contract event 0xba498a6e4c1996a9a7e64f07fe99506267fde36ac9398c866036bd6cdcfbab61.
//
// Solidity: event AgentDataAddNew(bytes[] sysPrompt)
func (_EternalAIKB20 *EternalAIKB20Filterer) ParseAgentDataAddNew(log types.Log) (*EternalAIKB20AgentDataAddNew, error) {
	event := new(EternalAIKB20AgentDataAddNew)
	if err := _EternalAIKB20.contract.UnpackLog(event, "AgentDataAddNew", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EternalAIKB20AgentDataUpdateIterator is returned from FilterAgentDataUpdate and is used to iterate over the raw logs and unpacked data for AgentDataUpdate events raised by the EternalAIKB20 contract.
type EternalAIKB20AgentDataUpdateIterator struct {
	Event *EternalAIKB20AgentDataUpdate // Event containing the contract specifics and raw log

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
func (it *EternalAIKB20AgentDataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EternalAIKB20AgentDataUpdate)
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
		it.Event = new(EternalAIKB20AgentDataUpdate)
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
func (it *EternalAIKB20AgentDataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EternalAIKB20AgentDataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EternalAIKB20AgentDataUpdate represents a AgentDataUpdate event raised by the EternalAIKB20 contract.
type EternalAIKB20AgentDataUpdate struct {
	PromptIndex  *big.Int
	OldSysPrompt []byte
	NewSysPrompt []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAgentDataUpdate is a free log retrieval operation binding the contract event 0x46c2322429cd7009e56aa485e0ffad13b6f8973e3e3b38168d4074d9a528eabe.
//
// Solidity: event AgentDataUpdate(uint256 promptIndex, bytes oldSysPrompt, bytes newSysPrompt)
func (_EternalAIKB20 *EternalAIKB20Filterer) FilterAgentDataUpdate(opts *bind.FilterOpts) (*EternalAIKB20AgentDataUpdateIterator, error) {

	logs, sub, err := _EternalAIKB20.contract.FilterLogs(opts, "AgentDataUpdate")
	if err != nil {
		return nil, err
	}
	return &EternalAIKB20AgentDataUpdateIterator{contract: _EternalAIKB20.contract, event: "AgentDataUpdate", logs: logs, sub: sub}, nil
}

// WatchAgentDataUpdate is a free log subscription operation binding the contract event 0x46c2322429cd7009e56aa485e0ffad13b6f8973e3e3b38168d4074d9a528eabe.
//
// Solidity: event AgentDataUpdate(uint256 promptIndex, bytes oldSysPrompt, bytes newSysPrompt)
func (_EternalAIKB20 *EternalAIKB20Filterer) WatchAgentDataUpdate(opts *bind.WatchOpts, sink chan<- *EternalAIKB20AgentDataUpdate) (event.Subscription, error) {

	logs, sub, err := _EternalAIKB20.contract.WatchLogs(opts, "AgentDataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EternalAIKB20AgentDataUpdate)
				if err := _EternalAIKB20.contract.UnpackLog(event, "AgentDataUpdate", log); err != nil {
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

// ParseAgentDataUpdate is a log parse operation binding the contract event 0x46c2322429cd7009e56aa485e0ffad13b6f8973e3e3b38168d4074d9a528eabe.
//
// Solidity: event AgentDataUpdate(uint256 promptIndex, bytes oldSysPrompt, bytes newSysPrompt)
func (_EternalAIKB20 *EternalAIKB20Filterer) ParseAgentDataUpdate(log types.Log) (*EternalAIKB20AgentDataUpdate, error) {
	event := new(EternalAIKB20AgentDataUpdate)
	if err := _EternalAIKB20.contract.UnpackLog(event, "AgentDataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EternalAIKB20AgentFeeUpdateIterator is returned from FilterAgentFeeUpdate and is used to iterate over the raw logs and unpacked data for AgentFeeUpdate events raised by the EternalAIKB20 contract.
type EternalAIKB20AgentFeeUpdateIterator struct {
	Event *EternalAIKB20AgentFeeUpdate // Event containing the contract specifics and raw log

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
func (it *EternalAIKB20AgentFeeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EternalAIKB20AgentFeeUpdate)
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
		it.Event = new(EternalAIKB20AgentFeeUpdate)
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
func (it *EternalAIKB20AgentFeeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EternalAIKB20AgentFeeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EternalAIKB20AgentFeeUpdate represents a AgentFeeUpdate event raised by the EternalAIKB20 contract.
type EternalAIKB20AgentFeeUpdate struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAgentFeeUpdate is a free log retrieval operation binding the contract event 0x34ae73c4d1b1bf2ec45030153dda7665cc119bcf159982735207f5e992cc9846.
//
// Solidity: event AgentFeeUpdate(uint256 fee)
func (_EternalAIKB20 *EternalAIKB20Filterer) FilterAgentFeeUpdate(opts *bind.FilterOpts) (*EternalAIKB20AgentFeeUpdateIterator, error) {

	logs, sub, err := _EternalAIKB20.contract.FilterLogs(opts, "AgentFeeUpdate")
	if err != nil {
		return nil, err
	}
	return &EternalAIKB20AgentFeeUpdateIterator{contract: _EternalAIKB20.contract, event: "AgentFeeUpdate", logs: logs, sub: sub}, nil
}

// WatchAgentFeeUpdate is a free log subscription operation binding the contract event 0x34ae73c4d1b1bf2ec45030153dda7665cc119bcf159982735207f5e992cc9846.
//
// Solidity: event AgentFeeUpdate(uint256 fee)
func (_EternalAIKB20 *EternalAIKB20Filterer) WatchAgentFeeUpdate(opts *bind.WatchOpts, sink chan<- *EternalAIKB20AgentFeeUpdate) (event.Subscription, error) {

	logs, sub, err := _EternalAIKB20.contract.WatchLogs(opts, "AgentFeeUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EternalAIKB20AgentFeeUpdate)
				if err := _EternalAIKB20.contract.UnpackLog(event, "AgentFeeUpdate", log); err != nil {
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

// ParseAgentFeeUpdate is a log parse operation binding the contract event 0x34ae73c4d1b1bf2ec45030153dda7665cc119bcf159982735207f5e992cc9846.
//
// Solidity: event AgentFeeUpdate(uint256 fee)
func (_EternalAIKB20 *EternalAIKB20Filterer) ParseAgentFeeUpdate(log types.Log) (*EternalAIKB20AgentFeeUpdate, error) {
	event := new(EternalAIKB20AgentFeeUpdate)
	if err := _EternalAIKB20.contract.UnpackLog(event, "AgentFeeUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EternalAIKB20AgentURIUpdateIterator is returned from FilterAgentURIUpdate and is used to iterate over the raw logs and unpacked data for AgentURIUpdate events raised by the EternalAIKB20 contract.
type EternalAIKB20AgentURIUpdateIterator struct {
	Event *EternalAIKB20AgentURIUpdate // Event containing the contract specifics and raw log

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
func (it *EternalAIKB20AgentURIUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EternalAIKB20AgentURIUpdate)
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
		it.Event = new(EternalAIKB20AgentURIUpdate)
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
func (it *EternalAIKB20AgentURIUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EternalAIKB20AgentURIUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EternalAIKB20AgentURIUpdate represents a AgentURIUpdate event raised by the EternalAIKB20 contract.
type EternalAIKB20AgentURIUpdate struct {
	Uri string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAgentURIUpdate is a free log retrieval operation binding the contract event 0x9b8d9731c54e4a06c1dae7b8a0ad1c08e9f164508165a830f84de0d9810062bb.
//
// Solidity: event AgentURIUpdate(string uri)
func (_EternalAIKB20 *EternalAIKB20Filterer) FilterAgentURIUpdate(opts *bind.FilterOpts) (*EternalAIKB20AgentURIUpdateIterator, error) {

	logs, sub, err := _EternalAIKB20.contract.FilterLogs(opts, "AgentURIUpdate")
	if err != nil {
		return nil, err
	}
	return &EternalAIKB20AgentURIUpdateIterator{contract: _EternalAIKB20.contract, event: "AgentURIUpdate", logs: logs, sub: sub}, nil
}

// WatchAgentURIUpdate is a free log subscription operation binding the contract event 0x9b8d9731c54e4a06c1dae7b8a0ad1c08e9f164508165a830f84de0d9810062bb.
//
// Solidity: event AgentURIUpdate(string uri)
func (_EternalAIKB20 *EternalAIKB20Filterer) WatchAgentURIUpdate(opts *bind.WatchOpts, sink chan<- *EternalAIKB20AgentURIUpdate) (event.Subscription, error) {

	logs, sub, err := _EternalAIKB20.contract.WatchLogs(opts, "AgentURIUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EternalAIKB20AgentURIUpdate)
				if err := _EternalAIKB20.contract.UnpackLog(event, "AgentURIUpdate", log); err != nil {
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

// ParseAgentURIUpdate is a log parse operation binding the contract event 0x9b8d9731c54e4a06c1dae7b8a0ad1c08e9f164508165a830f84de0d9810062bb.
//
// Solidity: event AgentURIUpdate(string uri)
func (_EternalAIKB20 *EternalAIKB20Filterer) ParseAgentURIUpdate(log types.Log) (*EternalAIKB20AgentURIUpdate, error) {
	event := new(EternalAIKB20AgentURIUpdate)
	if err := _EternalAIKB20.contract.UnpackLog(event, "AgentURIUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EternalAIKB20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the EternalAIKB20 contract.
type EternalAIKB20ApprovalIterator struct {
	Event *EternalAIKB20Approval // Event containing the contract specifics and raw log

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
func (it *EternalAIKB20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EternalAIKB20Approval)
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
		it.Event = new(EternalAIKB20Approval)
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
func (it *EternalAIKB20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EternalAIKB20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EternalAIKB20Approval represents a Approval event raised by the EternalAIKB20 contract.
type EternalAIKB20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EternalAIKB20 *EternalAIKB20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*EternalAIKB20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EternalAIKB20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &EternalAIKB20ApprovalIterator{contract: _EternalAIKB20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EternalAIKB20 *EternalAIKB20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EternalAIKB20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EternalAIKB20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EternalAIKB20Approval)
				if err := _EternalAIKB20.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EternalAIKB20 *EternalAIKB20Filterer) ParseApproval(log types.Log) (*EternalAIKB20Approval, error) {
	event := new(EternalAIKB20Approval)
	if err := _EternalAIKB20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EternalAIKB20GPUManagerUpdateIterator is returned from FilterGPUManagerUpdate and is used to iterate over the raw logs and unpacked data for GPUManagerUpdate events raised by the EternalAIKB20 contract.
type EternalAIKB20GPUManagerUpdateIterator struct {
	Event *EternalAIKB20GPUManagerUpdate // Event containing the contract specifics and raw log

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
func (it *EternalAIKB20GPUManagerUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EternalAIKB20GPUManagerUpdate)
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
		it.Event = new(EternalAIKB20GPUManagerUpdate)
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
func (it *EternalAIKB20GPUManagerUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EternalAIKB20GPUManagerUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EternalAIKB20GPUManagerUpdate represents a GPUManagerUpdate event raised by the EternalAIKB20 contract.
type EternalAIKB20GPUManagerUpdate struct {
	GpuManager common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGPUManagerUpdate is a free log retrieval operation binding the contract event 0x752bc81ffc4d7e1886a090f40c08324c06359626e920d446f8deafb05622782a.
//
// Solidity: event GPUManagerUpdate(address gpuManager)
func (_EternalAIKB20 *EternalAIKB20Filterer) FilterGPUManagerUpdate(opts *bind.FilterOpts) (*EternalAIKB20GPUManagerUpdateIterator, error) {

	logs, sub, err := _EternalAIKB20.contract.FilterLogs(opts, "GPUManagerUpdate")
	if err != nil {
		return nil, err
	}
	return &EternalAIKB20GPUManagerUpdateIterator{contract: _EternalAIKB20.contract, event: "GPUManagerUpdate", logs: logs, sub: sub}, nil
}

// WatchGPUManagerUpdate is a free log subscription operation binding the contract event 0x752bc81ffc4d7e1886a090f40c08324c06359626e920d446f8deafb05622782a.
//
// Solidity: event GPUManagerUpdate(address gpuManager)
func (_EternalAIKB20 *EternalAIKB20Filterer) WatchGPUManagerUpdate(opts *bind.WatchOpts, sink chan<- *EternalAIKB20GPUManagerUpdate) (event.Subscription, error) {

	logs, sub, err := _EternalAIKB20.contract.WatchLogs(opts, "GPUManagerUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EternalAIKB20GPUManagerUpdate)
				if err := _EternalAIKB20.contract.UnpackLog(event, "GPUManagerUpdate", log); err != nil {
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

// ParseGPUManagerUpdate is a log parse operation binding the contract event 0x752bc81ffc4d7e1886a090f40c08324c06359626e920d446f8deafb05622782a.
//
// Solidity: event GPUManagerUpdate(address gpuManager)
func (_EternalAIKB20 *EternalAIKB20Filterer) ParseGPUManagerUpdate(log types.Log) (*EternalAIKB20GPUManagerUpdate, error) {
	event := new(EternalAIKB20GPUManagerUpdate)
	if err := _EternalAIKB20.contract.UnpackLog(event, "GPUManagerUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EternalAIKB20InferencePerformedIterator is returned from FilterInferencePerformed and is used to iterate over the raw logs and unpacked data for InferencePerformed events raised by the EternalAIKB20 contract.
type EternalAIKB20InferencePerformedIterator struct {
	Event *EternalAIKB20InferencePerformed // Event containing the contract specifics and raw log

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
func (it *EternalAIKB20InferencePerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EternalAIKB20InferencePerformed)
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
		it.Event = new(EternalAIKB20InferencePerformed)
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
func (it *EternalAIKB20InferencePerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EternalAIKB20InferencePerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EternalAIKB20InferencePerformed represents a InferencePerformed event raised by the EternalAIKB20 contract.
type EternalAIKB20InferencePerformed struct {
	Caller       common.Address
	Data         []byte
	Fee          *big.Int
	ExternalData string
	InferenceId  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInferencePerformed is a free log retrieval operation binding the contract event 0xd6f4605d9629a8ab6dc9e171ea3027456f2bb13d67203afd50552488bc90a3f3.
//
// Solidity: event InferencePerformed(address indexed caller, bytes data, uint256 fee, string externalData, uint256 inferenceId)
func (_EternalAIKB20 *EternalAIKB20Filterer) FilterInferencePerformed(opts *bind.FilterOpts, caller []common.Address) (*EternalAIKB20InferencePerformedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _EternalAIKB20.contract.FilterLogs(opts, "InferencePerformed", callerRule)
	if err != nil {
		return nil, err
	}
	return &EternalAIKB20InferencePerformedIterator{contract: _EternalAIKB20.contract, event: "InferencePerformed", logs: logs, sub: sub}, nil
}

// WatchInferencePerformed is a free log subscription operation binding the contract event 0xd6f4605d9629a8ab6dc9e171ea3027456f2bb13d67203afd50552488bc90a3f3.
//
// Solidity: event InferencePerformed(address indexed caller, bytes data, uint256 fee, string externalData, uint256 inferenceId)
func (_EternalAIKB20 *EternalAIKB20Filterer) WatchInferencePerformed(opts *bind.WatchOpts, sink chan<- *EternalAIKB20InferencePerformed, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _EternalAIKB20.contract.WatchLogs(opts, "InferencePerformed", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EternalAIKB20InferencePerformed)
				if err := _EternalAIKB20.contract.UnpackLog(event, "InferencePerformed", log); err != nil {
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

// ParseInferencePerformed is a log parse operation binding the contract event 0xd6f4605d9629a8ab6dc9e171ea3027456f2bb13d67203afd50552488bc90a3f3.
//
// Solidity: event InferencePerformed(address indexed caller, bytes data, uint256 fee, string externalData, uint256 inferenceId)
func (_EternalAIKB20 *EternalAIKB20Filterer) ParseInferencePerformed(log types.Log) (*EternalAIKB20InferencePerformed, error) {
	event := new(EternalAIKB20InferencePerformed)
	if err := _EternalAIKB20.contract.UnpackLog(event, "InferencePerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EternalAIKB20InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EternalAIKB20 contract.
type EternalAIKB20InitializedIterator struct {
	Event *EternalAIKB20Initialized // Event containing the contract specifics and raw log

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
func (it *EternalAIKB20InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EternalAIKB20Initialized)
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
		it.Event = new(EternalAIKB20Initialized)
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
func (it *EternalAIKB20InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EternalAIKB20InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EternalAIKB20Initialized represents a Initialized event raised by the EternalAIKB20 contract.
type EternalAIKB20Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EternalAIKB20 *EternalAIKB20Filterer) FilterInitialized(opts *bind.FilterOpts) (*EternalAIKB20InitializedIterator, error) {

	logs, sub, err := _EternalAIKB20.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EternalAIKB20InitializedIterator{contract: _EternalAIKB20.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EternalAIKB20 *EternalAIKB20Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EternalAIKB20Initialized) (event.Subscription, error) {

	logs, sub, err := _EternalAIKB20.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EternalAIKB20Initialized)
				if err := _EternalAIKB20.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_EternalAIKB20 *EternalAIKB20Filterer) ParseInitialized(log types.Log) (*EternalAIKB20Initialized, error) {
	event := new(EternalAIKB20Initialized)
	if err := _EternalAIKB20.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EternalAIKB20ModelIdUpdateIterator is returned from FilterModelIdUpdate and is used to iterate over the raw logs and unpacked data for ModelIdUpdate events raised by the EternalAIKB20 contract.
type EternalAIKB20ModelIdUpdateIterator struct {
	Event *EternalAIKB20ModelIdUpdate // Event containing the contract specifics and raw log

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
func (it *EternalAIKB20ModelIdUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EternalAIKB20ModelIdUpdate)
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
		it.Event = new(EternalAIKB20ModelIdUpdate)
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
func (it *EternalAIKB20ModelIdUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EternalAIKB20ModelIdUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EternalAIKB20ModelIdUpdate represents a ModelIdUpdate event raised by the EternalAIKB20 contract.
type EternalAIKB20ModelIdUpdate struct {
	ModelId uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterModelIdUpdate is a free log retrieval operation binding the contract event 0x7a55cba18b26688d885d840641dae0c8c4842f3cd33131a902de436afe2013ee.
//
// Solidity: event ModelIdUpdate(uint32 modelId)
func (_EternalAIKB20 *EternalAIKB20Filterer) FilterModelIdUpdate(opts *bind.FilterOpts) (*EternalAIKB20ModelIdUpdateIterator, error) {

	logs, sub, err := _EternalAIKB20.contract.FilterLogs(opts, "ModelIdUpdate")
	if err != nil {
		return nil, err
	}
	return &EternalAIKB20ModelIdUpdateIterator{contract: _EternalAIKB20.contract, event: "ModelIdUpdate", logs: logs, sub: sub}, nil
}

// WatchModelIdUpdate is a free log subscription operation binding the contract event 0x7a55cba18b26688d885d840641dae0c8c4842f3cd33131a902de436afe2013ee.
//
// Solidity: event ModelIdUpdate(uint32 modelId)
func (_EternalAIKB20 *EternalAIKB20Filterer) WatchModelIdUpdate(opts *bind.WatchOpts, sink chan<- *EternalAIKB20ModelIdUpdate) (event.Subscription, error) {

	logs, sub, err := _EternalAIKB20.contract.WatchLogs(opts, "ModelIdUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EternalAIKB20ModelIdUpdate)
				if err := _EternalAIKB20.contract.UnpackLog(event, "ModelIdUpdate", log); err != nil {
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

// ParseModelIdUpdate is a log parse operation binding the contract event 0x7a55cba18b26688d885d840641dae0c8c4842f3cd33131a902de436afe2013ee.
//
// Solidity: event ModelIdUpdate(uint32 modelId)
func (_EternalAIKB20 *EternalAIKB20Filterer) ParseModelIdUpdate(log types.Log) (*EternalAIKB20ModelIdUpdate, error) {
	event := new(EternalAIKB20ModelIdUpdate)
	if err := _EternalAIKB20.contract.UnpackLog(event, "ModelIdUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EternalAIKB20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EternalAIKB20 contract.
type EternalAIKB20OwnershipTransferredIterator struct {
	Event *EternalAIKB20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EternalAIKB20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EternalAIKB20OwnershipTransferred)
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
		it.Event = new(EternalAIKB20OwnershipTransferred)
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
func (it *EternalAIKB20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EternalAIKB20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EternalAIKB20OwnershipTransferred represents a OwnershipTransferred event raised by the EternalAIKB20 contract.
type EternalAIKB20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EternalAIKB20 *EternalAIKB20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EternalAIKB20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EternalAIKB20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EternalAIKB20OwnershipTransferredIterator{contract: _EternalAIKB20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EternalAIKB20 *EternalAIKB20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EternalAIKB20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EternalAIKB20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EternalAIKB20OwnershipTransferred)
				if err := _EternalAIKB20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EternalAIKB20 *EternalAIKB20Filterer) ParseOwnershipTransferred(log types.Log) (*EternalAIKB20OwnershipTransferred, error) {
	event := new(EternalAIKB20OwnershipTransferred)
	if err := _EternalAIKB20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EternalAIKB20PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the EternalAIKB20 contract.
type EternalAIKB20PausedIterator struct {
	Event *EternalAIKB20Paused // Event containing the contract specifics and raw log

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
func (it *EternalAIKB20PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EternalAIKB20Paused)
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
		it.Event = new(EternalAIKB20Paused)
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
func (it *EternalAIKB20PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EternalAIKB20PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EternalAIKB20Paused represents a Paused event raised by the EternalAIKB20 contract.
type EternalAIKB20Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EternalAIKB20 *EternalAIKB20Filterer) FilterPaused(opts *bind.FilterOpts) (*EternalAIKB20PausedIterator, error) {

	logs, sub, err := _EternalAIKB20.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EternalAIKB20PausedIterator{contract: _EternalAIKB20.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EternalAIKB20 *EternalAIKB20Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EternalAIKB20Paused) (event.Subscription, error) {

	logs, sub, err := _EternalAIKB20.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EternalAIKB20Paused)
				if err := _EternalAIKB20.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_EternalAIKB20 *EternalAIKB20Filterer) ParsePaused(log types.Log) (*EternalAIKB20Paused, error) {
	event := new(EternalAIKB20Paused)
	if err := _EternalAIKB20.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EternalAIKB20PromptSchedulerUpdateIterator is returned from FilterPromptSchedulerUpdate and is used to iterate over the raw logs and unpacked data for PromptSchedulerUpdate events raised by the EternalAIKB20 contract.
type EternalAIKB20PromptSchedulerUpdateIterator struct {
	Event *EternalAIKB20PromptSchedulerUpdate // Event containing the contract specifics and raw log

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
func (it *EternalAIKB20PromptSchedulerUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EternalAIKB20PromptSchedulerUpdate)
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
		it.Event = new(EternalAIKB20PromptSchedulerUpdate)
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
func (it *EternalAIKB20PromptSchedulerUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EternalAIKB20PromptSchedulerUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EternalAIKB20PromptSchedulerUpdate represents a PromptSchedulerUpdate event raised by the EternalAIKB20 contract.
type EternalAIKB20PromptSchedulerUpdate struct {
	PromptScheduler common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPromptSchedulerUpdate is a free log retrieval operation binding the contract event 0x667557d852582e84e7de441f650ea0aacbb7de26e3485436e0c27ba8d19a79f1.
//
// Solidity: event PromptSchedulerUpdate(address promptScheduler)
func (_EternalAIKB20 *EternalAIKB20Filterer) FilterPromptSchedulerUpdate(opts *bind.FilterOpts) (*EternalAIKB20PromptSchedulerUpdateIterator, error) {

	logs, sub, err := _EternalAIKB20.contract.FilterLogs(opts, "PromptSchedulerUpdate")
	if err != nil {
		return nil, err
	}
	return &EternalAIKB20PromptSchedulerUpdateIterator{contract: _EternalAIKB20.contract, event: "PromptSchedulerUpdate", logs: logs, sub: sub}, nil
}

// WatchPromptSchedulerUpdate is a free log subscription operation binding the contract event 0x667557d852582e84e7de441f650ea0aacbb7de26e3485436e0c27ba8d19a79f1.
//
// Solidity: event PromptSchedulerUpdate(address promptScheduler)
func (_EternalAIKB20 *EternalAIKB20Filterer) WatchPromptSchedulerUpdate(opts *bind.WatchOpts, sink chan<- *EternalAIKB20PromptSchedulerUpdate) (event.Subscription, error) {

	logs, sub, err := _EternalAIKB20.contract.WatchLogs(opts, "PromptSchedulerUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EternalAIKB20PromptSchedulerUpdate)
				if err := _EternalAIKB20.contract.UnpackLog(event, "PromptSchedulerUpdate", log); err != nil {
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

// ParsePromptSchedulerUpdate is a log parse operation binding the contract event 0x667557d852582e84e7de441f650ea0aacbb7de26e3485436e0c27ba8d19a79f1.
//
// Solidity: event PromptSchedulerUpdate(address promptScheduler)
func (_EternalAIKB20 *EternalAIKB20Filterer) ParsePromptSchedulerUpdate(log types.Log) (*EternalAIKB20PromptSchedulerUpdate, error) {
	event := new(EternalAIKB20PromptSchedulerUpdate)
	if err := _EternalAIKB20.contract.UnpackLog(event, "PromptSchedulerUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EternalAIKB20TokenFeeUpdateIterator is returned from FilterTokenFeeUpdate and is used to iterate over the raw logs and unpacked data for TokenFeeUpdate events raised by the EternalAIKB20 contract.
type EternalAIKB20TokenFeeUpdateIterator struct {
	Event *EternalAIKB20TokenFeeUpdate // Event containing the contract specifics and raw log

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
func (it *EternalAIKB20TokenFeeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EternalAIKB20TokenFeeUpdate)
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
		it.Event = new(EternalAIKB20TokenFeeUpdate)
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
func (it *EternalAIKB20TokenFeeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EternalAIKB20TokenFeeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EternalAIKB20TokenFeeUpdate represents a TokenFeeUpdate event raised by the EternalAIKB20 contract.
type EternalAIKB20TokenFeeUpdate struct {
	TokenFee common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenFeeUpdate is a free log retrieval operation binding the contract event 0x951060957c3c47944aafba621b12dfb70a865f7a4dd1bf1087975f19856e56a8.
//
// Solidity: event TokenFeeUpdate(address tokenFee)
func (_EternalAIKB20 *EternalAIKB20Filterer) FilterTokenFeeUpdate(opts *bind.FilterOpts) (*EternalAIKB20TokenFeeUpdateIterator, error) {

	logs, sub, err := _EternalAIKB20.contract.FilterLogs(opts, "TokenFeeUpdate")
	if err != nil {
		return nil, err
	}
	return &EternalAIKB20TokenFeeUpdateIterator{contract: _EternalAIKB20.contract, event: "TokenFeeUpdate", logs: logs, sub: sub}, nil
}

// WatchTokenFeeUpdate is a free log subscription operation binding the contract event 0x951060957c3c47944aafba621b12dfb70a865f7a4dd1bf1087975f19856e56a8.
//
// Solidity: event TokenFeeUpdate(address tokenFee)
func (_EternalAIKB20 *EternalAIKB20Filterer) WatchTokenFeeUpdate(opts *bind.WatchOpts, sink chan<- *EternalAIKB20TokenFeeUpdate) (event.Subscription, error) {

	logs, sub, err := _EternalAIKB20.contract.WatchLogs(opts, "TokenFeeUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EternalAIKB20TokenFeeUpdate)
				if err := _EternalAIKB20.contract.UnpackLog(event, "TokenFeeUpdate", log); err != nil {
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

// ParseTokenFeeUpdate is a log parse operation binding the contract event 0x951060957c3c47944aafba621b12dfb70a865f7a4dd1bf1087975f19856e56a8.
//
// Solidity: event TokenFeeUpdate(address tokenFee)
func (_EternalAIKB20 *EternalAIKB20Filterer) ParseTokenFeeUpdate(log types.Log) (*EternalAIKB20TokenFeeUpdate, error) {
	event := new(EternalAIKB20TokenFeeUpdate)
	if err := _EternalAIKB20.contract.UnpackLog(event, "TokenFeeUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EternalAIKB20TopUpPoolBalanceIterator is returned from FilterTopUpPoolBalance and is used to iterate over the raw logs and unpacked data for TopUpPoolBalance events raised by the EternalAIKB20 contract.
type EternalAIKB20TopUpPoolBalanceIterator struct {
	Event *EternalAIKB20TopUpPoolBalance // Event containing the contract specifics and raw log

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
func (it *EternalAIKB20TopUpPoolBalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EternalAIKB20TopUpPoolBalance)
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
		it.Event = new(EternalAIKB20TopUpPoolBalance)
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
func (it *EternalAIKB20TopUpPoolBalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EternalAIKB20TopUpPoolBalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EternalAIKB20TopUpPoolBalance represents a TopUpPoolBalance event raised by the EternalAIKB20 contract.
type EternalAIKB20TopUpPoolBalance struct {
	Caller common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTopUpPoolBalance is a free log retrieval operation binding the contract event 0x52d53b7ae2f89bb8b4d82cfa1c120c7a724ed4f14d15862f38b2841f2bbbbbaa.
//
// Solidity: event TopUpPoolBalance(address caller, uint256 amount)
func (_EternalAIKB20 *EternalAIKB20Filterer) FilterTopUpPoolBalance(opts *bind.FilterOpts) (*EternalAIKB20TopUpPoolBalanceIterator, error) {

	logs, sub, err := _EternalAIKB20.contract.FilterLogs(opts, "TopUpPoolBalance")
	if err != nil {
		return nil, err
	}
	return &EternalAIKB20TopUpPoolBalanceIterator{contract: _EternalAIKB20.contract, event: "TopUpPoolBalance", logs: logs, sub: sub}, nil
}

// WatchTopUpPoolBalance is a free log subscription operation binding the contract event 0x52d53b7ae2f89bb8b4d82cfa1c120c7a724ed4f14d15862f38b2841f2bbbbbaa.
//
// Solidity: event TopUpPoolBalance(address caller, uint256 amount)
func (_EternalAIKB20 *EternalAIKB20Filterer) WatchTopUpPoolBalance(opts *bind.WatchOpts, sink chan<- *EternalAIKB20TopUpPoolBalance) (event.Subscription, error) {

	logs, sub, err := _EternalAIKB20.contract.WatchLogs(opts, "TopUpPoolBalance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EternalAIKB20TopUpPoolBalance)
				if err := _EternalAIKB20.contract.UnpackLog(event, "TopUpPoolBalance", log); err != nil {
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

// ParseTopUpPoolBalance is a log parse operation binding the contract event 0x52d53b7ae2f89bb8b4d82cfa1c120c7a724ed4f14d15862f38b2841f2bbbbbaa.
//
// Solidity: event TopUpPoolBalance(address caller, uint256 amount)
func (_EternalAIKB20 *EternalAIKB20Filterer) ParseTopUpPoolBalance(log types.Log) (*EternalAIKB20TopUpPoolBalance, error) {
	event := new(EternalAIKB20TopUpPoolBalance)
	if err := _EternalAIKB20.contract.UnpackLog(event, "TopUpPoolBalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EternalAIKB20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the EternalAIKB20 contract.
type EternalAIKB20TransferIterator struct {
	Event *EternalAIKB20Transfer // Event containing the contract specifics and raw log

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
func (it *EternalAIKB20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EternalAIKB20Transfer)
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
		it.Event = new(EternalAIKB20Transfer)
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
func (it *EternalAIKB20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EternalAIKB20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EternalAIKB20Transfer represents a Transfer event raised by the EternalAIKB20 contract.
type EternalAIKB20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EternalAIKB20 *EternalAIKB20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EternalAIKB20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EternalAIKB20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EternalAIKB20TransferIterator{contract: _EternalAIKB20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EternalAIKB20 *EternalAIKB20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EternalAIKB20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EternalAIKB20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EternalAIKB20Transfer)
				if err := _EternalAIKB20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EternalAIKB20 *EternalAIKB20Filterer) ParseTransfer(log types.Log) (*EternalAIKB20Transfer, error) {
	event := new(EternalAIKB20Transfer)
	if err := _EternalAIKB20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EternalAIKB20UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the EternalAIKB20 contract.
type EternalAIKB20UnpausedIterator struct {
	Event *EternalAIKB20Unpaused // Event containing the contract specifics and raw log

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
func (it *EternalAIKB20UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EternalAIKB20Unpaused)
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
		it.Event = new(EternalAIKB20Unpaused)
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
func (it *EternalAIKB20UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EternalAIKB20UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EternalAIKB20Unpaused represents a Unpaused event raised by the EternalAIKB20 contract.
type EternalAIKB20Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EternalAIKB20 *EternalAIKB20Filterer) FilterUnpaused(opts *bind.FilterOpts) (*EternalAIKB20UnpausedIterator, error) {

	logs, sub, err := _EternalAIKB20.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EternalAIKB20UnpausedIterator{contract: _EternalAIKB20.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EternalAIKB20 *EternalAIKB20Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EternalAIKB20Unpaused) (event.Subscription, error) {

	logs, sub, err := _EternalAIKB20.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EternalAIKB20Unpaused)
				if err := _EternalAIKB20.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_EternalAIKB20 *EternalAIKB20Filterer) ParseUnpaused(log types.Log) (*EternalAIKB20Unpaused, error) {
	event := new(EternalAIKB20Unpaused)
	if err := _EternalAIKB20.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
