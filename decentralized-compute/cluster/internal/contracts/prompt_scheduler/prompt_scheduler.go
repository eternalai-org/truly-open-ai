// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package prompt_scheduler

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

// ISchedulerInference is an auto generated low-level Go binding around an user-defined struct.
type ISchedulerInference struct {
	Value          *big.Int
	ModelId        uint32
	SubmitTimeout  *big.Int
	Status         uint8
	Creator        common.Address
	ProcessedMiner common.Address
	Input          []byte
	Output         []byte
}

// PromptSchedulerMetaData contains all meta data concerning the PromptScheduler contract.
var PromptSchedulerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadySubmitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInferenceStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAssignedWorker\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubmitTimeout\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Uint256Set_DuplicatedValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"batchId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"inferId\",\"type\":\"uint64\"}],\"name\":\"AppendToBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"inferenceId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumIScheduler.InferenceStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"InferenceStatusUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"inferenceId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"expiredAt\",\"type\":\"uint40\"}],\"name\":\"NewAssignment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"inferenceId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"NewInference\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"inferId\",\"type\":\"uint256\"}],\"name\":\"SolutionSubmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assignmentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"StreamedData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_batchPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gpuManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_inferenceCounter\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_lastBatchTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minerRequirement\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minerValidatorFeeRatio\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_submitDuration\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_wEAIToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchId\",\"type\":\"uint64\"}],\"name\":\"getBatchInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"getInferenceByMiner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"inferId\",\"type\":\"uint64\"}],\"name\":\"getInferenceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"submitTimeout\",\"type\":\"uint40\"},{\"internalType\":\"enumIScheduler.InferenceStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"processedMiner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"output\",\"type\":\"bytes\"}],\"internalType\":\"structIScheduler.Inference\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinerRequirement\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"infer\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"modelId\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"infer\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wEAIToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gpuManager_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"minerRequirement_\",\"type\":\"uint8\"},{\"internalType\":\"uint40\",\"name\":\"submitDuration_\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"minerValidatorFeeRatio_\",\"type\":\"uint16\"},{\"internalType\":\"uint40\",\"name\":\"batchPeriod_\",\"type\":\"uint40\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"submitDuration\",\"type\":\"uint40\"}],\"name\":\"setSubmitDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wEAIToken\",\"type\":\"address\"}],\"name\":\"setWEAIAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"inferId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"solution\",\"type\":\"bytes\"}],\"name\":\"submitSolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080806040523461001657612eba908161001c8239f35b600080fdfe608080604052600436101561001d575b50361561001b57600080fd5b005b600090813560e01c90816308c147fd1461283657508063187179381461267e57806334b96ee414611f7d5780633f4ba83a14611ee15780634872926214611ccb57806348751e5014610f065780635037011114611c8f57806354fd4d5014611c0f5780635630180614611b325780635c975abb14611af15780635cc68731146112db578063627a04b8146112925780636f643736146111df578063715018a6146111405780637362323c146110ba5780637a80e13e146110745780637f8f29fc146110385780638456cb5914610f99578063871c15b114610f4857806387b97f1c14610f065780638da5cb5b14610eb4578063a1e0a42914610e74578063a50d860014610af8578063de1ce2bb146102275763f2fde38b0361000f57346102245760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102245761017061293c565b610178612982565b73ffffffffffffffffffffffffffffffffffffffff8116156101a05761019d90612a01565b80f35b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152fd5b80fd5b50346102245760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102245761025f612886565b906024359067ffffffffffffffff8211610224575061028460009136906004016128b0565b9061028d61295f565b91610296612bab565b604073ffffffffffffffffffffffffffffffffffffffff60015416604460ff60045460881c16835197889384927fe13f220e00000000000000000000000000000000000000000000000000000000845263ffffffff8c16600485015260248401525af1918215610aec576000948593610a90575b506001549467ffffffffffffffff8660a01c1667ffffffffffffffff8114610a6157600101957fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7bffffffffffffffff00000000000000000000000000000000000000008860a01b1691161760015567ffffffffffffffff8616600052600260205260406000208481556001810180547dffffffffffffffffffffffffffffffffffffffff000000000000000000008860501b16907fffff0000000000000000000000000000000000000000ffffffffffff0000000063ffffffff8c1691161717905567ffffffffffffffff84116108b2576104096003820154612c80565b601f8111610a1a575b50600084601f81116001146109565760039160009161094b575b508560011b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff87841b1c1916179101555b61047464ffffffffff60045460101c1643612cd3565b9067ffffffffffffffff871660005260026020526002604060002060018101690100000000000000000081547fffffffffffffffffffffffffffffffffffffffffffff000000000000ffffffff68ffffffffff000000008860201b169116171790550173ffffffffffffffffffffffffffffffffffffffff82167fffffffffffffffffffffffff000000000000000000000000000000000000000082541617905573ffffffffffffffffffffffffffffffffffffffff8116600052600360205260406000209067ffffffffffffffff88166000526001820160205260406000205461091057815491680100000000000000008310156108b257600183018082558310156108e15773ffffffffffffffffffffffffffffffffffffffff928160005267ffffffffffffffff8a169060206000200155600181549167ffffffffffffffff8b166000520160205260406000205564ffffffffff6040519316835216907fce7f171f53023fc0778971a0fe3f4067468c0505e2485b2716e44c4487215e79602067ffffffffffffffff891692a367ffffffffffffffff61062561061c60065442612ce0565b60075490612ced565b1663ffffffff8716600052600560205260406000208160005260205260036040600020018054680100000000000000008110156108b25761066b91600182018155612d26565b81549060031b9067ffffffffffffffff808a16831b921b191617905567ffffffffffffffff86169063ffffffff8816907fb51ff9d6e56ade059ce28900d43a2402f306a6f04f1ce592fe2e24b817d5c8cd600080a460008073ffffffffffffffffffffffffffffffffffffffff8154167fffffffff00000000000000000000000000000000000000000000000000000000602560405161070a81612b32565b8181527f7432353629000000000000000000000000000000000000000000000000000000604060208301927f7472616e7366657246726f6d28616464726573732c616464726573732c75696e84520152201682604051602081019283523360248201523060448201528860648201526064815261078681612b4e565b51925af1610792612e0e565b9015908115610882575b506108585773ffffffffffffffffffffffffffffffffffffffff63ffffffff602097847f964ad1b4133d111aec4a94ec6525e35e7e6793cdb76b507a298ac273ef85c0179460405197885260608b890152816060890152608088013760006080868801015260006040870152169416926080817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f67ffffffffffffffff8a169601168101030190a467ffffffffffffffff60405191168152f35b60046040517fbfa871c5000000000000000000000000000000000000000000000000000000008152fd5b8051801515925082610897575b50503861079c565b6108aa9250602080918301019101612e6c565b15388061088f565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60248867ffffffffffffffff604051917f346c4a0e000000000000000000000000000000000000000000000000000000008352166004820152fd5b90508401353861042c565b506003820181526020812090805b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe087168110610a025750857fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08116106109ca575b50506003600185811b0191015561045e565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88760031b161c199085013516905538806109b8565b90916020600181928589013581550193019101610964565b600382016000526020600020601f860160051c810160208710610a5a575b601f830160051c82018110610a4e575050610412565b60008155600101610a38565b5080610a38565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9092506040813d604011610ae4575b81610aac60409383612b6a565b81010312610ae05780519473ffffffffffffffffffffffffffffffffffffffff86168603610224575060200151913861030a565b8480fd5b3d9150610a9f565b6040513d6000823e3d90fd5b50346102245760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022457610b3061293c565b6024359073ffffffffffffffffffffffffffffffffffffffff808316809303610e675760ff916044358381168103610e70576064359064ffffffffff928383168303610e675760843561ffff8116809103610e6c5760a435948516809503610e6757606c97885492888460081c161597888099610e5b575b8015610e45575b15610dc1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0094896001878316178d55610d93575b5082158015610d89575b610d5f578715610d3557610c5b71ff000000000000000000000000000000000095610c278c8e5460081c16610c2281612a6e565b612a6e565b610c3033612a01565b8c549b8c60081c1690610c4282612a6e565b610c4b82612a6e565b60d1541660d155610c2281612a6e565b6001610103557fffffffffffffffffffffffff00000000000000000000000000000000000000009116818c5416178b5560015416176001557fffffffffffffffffffffffffffff00ffffffffffffffffffff0000000000000066ffffffffff00006004549560101b169416179160881b16171760045542600655600755610ce0578280f35b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690557f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498602060405160018152a138808280f35b60046040517faa7feadc000000000000000000000000000000000000000000000000000000008152fd5b60046040517fe6c4247b000000000000000000000000000000000000000000000000000000008152fd5b5080821615610bee565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016610101178b5538610be4565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152fd5b50303b158015610baf575060018a861614610baf565b5060018a861610610ba8565b600080fd5b8880fd5b8580fd5b503461022457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022457602061ffff60045416604051908152f35b503461022457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022457602073ffffffffffffffffffffffffffffffffffffffff609f5416604051908152f35b503461022457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022457602060ff60045460881c16604051908152f35b503461022457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102245773ffffffffffffffffffffffffffffffffffffffff6020915416604051908152f35b503461022457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022457610fd0612982565b610fd8612bab565b610fe0612bab565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060d154161760d1557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a180f35b503461022457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610224576020600654604051908152f35b503461022457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022457602064ffffffffff60045460101c16604051908152f35b50346102245760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102245773ffffffffffffffffffffffffffffffffffffffff61110761293c565b61110f612982565b168015610d5f577fffffffffffffffffffffffff000000000000000000000000000000000000000082541617815580f35b503461022457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022457611177612982565b600073ffffffffffffffffffffffffffffffffffffffff609f547fffffffffffffffffffffffff00000000000000000000000000000000000000008116609f55167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b50346102245760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102245760043564ffffffffff8116808203610e6757611229612982565b15611268577fffffffffffffffffffffffffffffffffffffffffffffffffff0000000000ffff66ffffffffff00006004549260101b1691161760045580f35b60046040517f5cb045db000000000000000000000000000000000000000000000000000000008152fd5b503461022457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022457602067ffffffffffffffff60015460a01c16604051908152f35b50346102245760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022457611313612886565b60243567ffffffffffffffff8111611aed576113339036906004016128b0565b909161133d61295f565b91606435918215158303610e70576000949550611358612bab565b604073ffffffffffffffffffffffffffffffffffffffff60015416604460ff60045460881c16835198899384927fe13f220e00000000000000000000000000000000000000000000000000000000845263ffffffff8816600485015260248401525af1928315610aec576000958694611a95575b506001549567ffffffffffffffff8760a01c1667ffffffffffffffff8114610a6157600101967fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7bffffffffffffffff00000000000000000000000000000000000000008960a01b1691161760015567ffffffffffffffff8716600052600260205260406000208581556001810180547dffffffffffffffffffffffffffffffffffffffff000000000000000000008960501b16907fffff0000000000000000000000000000000000000000ffffffffffff0000000063ffffffff881691161717905567ffffffffffffffff85116108b2576114cb6003820154612c80565b601f8111611a4e575b50600085601f81116001146119885760039160009161197d575b508660011b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff88841b1c1916179101555b61153664ffffffffff60045460101c1643612cd3565b9067ffffffffffffffff881660005260026020526002604060002060018101690100000000000000000081547fffffffffffffffffffffffffffffffffffffffffffff000000000000ffffffff68ffffffffff000000008860201b169116171790550173ffffffffffffffffffffffffffffffffffffffff82167fffffffffffffffffffffffff000000000000000000000000000000000000000082541617905573ffffffffffffffffffffffffffffffffffffffff8116600052600360205260406000209067ffffffffffffffff89166000526001820160205260406000205461194257815491680100000000000000008310156108b257600183018082558310156108e15773ffffffffffffffffffffffffffffffffffffffff928160005267ffffffffffffffff8b169060206000200155600181549167ffffffffffffffff8c166000520160205260406000205564ffffffffff6040519316835216907fce7f171f53023fc0778971a0fe3f4067468c0505e2485b2716e44c4487215e79602067ffffffffffffffff8a1692a367ffffffffffffffff6116de61061c60065442612ce0565b1663ffffffff8316600052600560205260406000208160005260205260036040600020018054680100000000000000008110156108b25761172491600182018155612d26565b81549060031b9067ffffffffffffffff808b16831b921b191617905567ffffffffffffffff87169063ffffffff8416907fb51ff9d6e56ade059ce28900d43a2402f306a6f04f1ce592fe2e24b817d5c8cd600080a460008073ffffffffffffffffffffffffffffffffffffffff8154167fffffffff0000000000000000000000000000000000000000000000000000000060256040516117c381612b32565b8181527f7432353629000000000000000000000000000000000000000000000000000000604060208301927f7472616e7366657246726f6d28616464726573732c616464726573732c75696e84520152201682604051602081019283523360248201523060448201528960648201526064815261183f81612b4e565b51925af161184b612e0e565b9015908115611912575b506108585763ffffffff7f964ad1b4133d111aec4a94ec6525e35e7e6793cdb76b507a298ac273ef85c0179273ffffffffffffffffffffffffffffffffffffffff928560209a60405198895260608c8a01528160608a0152608089013760006080878901015215156040870152169416926080817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f67ffffffffffffffff8a169601168101030190a467ffffffffffffffff60405191168152f35b8051801515925082611927575b505038611855565b61193a9250602080918301019101612e6c565b15388061191f565b60248967ffffffffffffffff604051917f346c4a0e000000000000000000000000000000000000000000000000000000008352166004820152fd5b90508a0135386114ee565b506003820181526020812090805b8b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe089168210611a36575050867fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08116106119fe575b50506003600186811b01910155611520565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19908b013516905538806119ec565b60018394602093948493013581550193019101611996565b600382016000526020600020601f870160051c810160208810611a8e575b601f830160051c82018110611a825750506114d4565b60008155600101611a6c565b5080611a6c565b9093506040813d604011611ae5575b81611ab160409383612b6a565b81010312610e705780519573ffffffffffffffffffffffffffffffffffffffff8716870361022457506020015192386113cc565b3d9150611aa4565b8280fd5b503461022457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022457602060ff60d154166040519015158152f35b5034610224576020807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611c0b5773ffffffffffffffffffffffffffffffffffffffff611b8061293c565b168252600381526040822060405192838383549182815201908193835284832090835b818110611bf75750505084611bb9910385612b6a565b60405193838594850191818652518092526040850193925b828110611be057505050500390f35b835185528695509381019392810192600101611bd1565b825484529286019260019283019201611ba3565b5080fd5b503461022457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022457611c8b604051611c4d81612b16565b600681527f76302e302e32000000000000000000000000000000000000000000000000000060208201526040519182916020835260208301906128de565b0390f35b503461022457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610224576020600754604051908152f35b5034610224576020807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611c0b5767ffffffffffffffff611d0d612899565b606060e0604051611d1d81612af9565b868152868682015286604082015286838201528660808201528660a08201528260c08201520152168252600281526040822060405191611d5c83612af9565b815483526001820154908084019163ffffffff80821684526040860164ffffffffff908184861c16815260ff8460481c169260608901936006811015611eb45790849392918a969552608086019773ffffffffffffffffffffffffffffffffffffffff809660501c168952611df460048760028d0154169b60a08a019c8d5260c0611de960038301612d49565b9a01998a5201612d49565b9760e08c019889526040519b818d5251908c0152511660408a01525116606088015251946006861015611e855781611c8b9588976080890152511660a0870152511660c085015251611e54610100918260e08701526101208601906128de565b9151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085840301908501526128de565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60248b7f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b503461022457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022457611f18612982565b611f20612c15565b611f28612c15565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060d1541660d1557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a180f35b50346102245760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022457611fb5612899565b6024919067ffffffffffffffff833581811161267a57611fd99036906004016128b0565b919092611fe4612bab565b82156112685781169283855260029060209282845273ffffffffffffffffffffffffffffffffffffffff91828460408a2001541633036126505786885283855264ffffffffff90600191808360408c200154881c16904316116126265787895284865260ff8260408b20015460481c1660068110156125fa5782036125d057878952848652612079600460408b200154612c80565b6125a65783825416803b156125a2578980918c604051809481937fdfecce6f0000000000000000000000000000000000000000000000000000000083523360048401525af1801561259757612557575b508789528486526040892096600488019184821161252b5781906120ed8454612c80565b601f81116124da575b508b90601f8311600114612440578c92612435575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82851b9260031b1c19161790555b85019069020000000000000000007fffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffffff835416178255855461ffff6004541690818102918183041490151715612409576127109004928854167f7472616e7366657228616464726573732c75696e743235362900000000000000866040516121c281612b16565b601981520152604051868101917fa9059cbb000000000000000000000000000000000000000000000000000000008352338c830152856044830152604482526080820192828410858511176123db578b809493819460405251925af1612226612e0e565b90159081156123aa575b506108585763ffffffff8161224a61061c60065442612ce0565b16925416928389526005865260408920836000528652600360406000200180541561237c579988999a60009998995282876000205416881061230a575b506122b5907f79e6e7d96ce1a52bc5d92a4a2458c9647a3dd14e184fddfef0d2e8204f05a582979854612ce0565b9289526005855260408920911660005283526122d982604060002001918254612cd3565b9055604051908152a2337f9f669b92b9cbc7611f7ab6c77db07a424051c777433e21bd90f1bdf940096dd98380a380f35b90928092949698959750156123515750929587959094909390927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff909101916122b5612287565b897f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b8a7f4e487b710000000000000000000000000000000000000000000000000000000060005260326004526000fd5b805180151592508790836123c2575b50505038612230565b6123d29350820181019101612e6c565b153886816123b9565b8c7f4e487b710000000000000000000000000000000000000000000000000000000060005260416004526000fd5b89897f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b01359050388061210b565b8c917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08795168684528b8420935b8c8282106124c4575050841161248c575b505050811b01905561213c565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c1991013516905538808061247f565b838501358655899790950194928301920161246e565b909150838c52888c20601f840160051c8101918a8510612521575b84939291601f88920160051c01915b8281106125125750506120f6565b60008155859450879101612504565b90915081906124f5565b8b8b7f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b83819a929a1161256b5760405297386120c9565b8a827f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b6040513d8c823e3d90fd5b8980fd5b60046040517f9fbfc589000000000000000000000000000000000000000000000000000000008152fd5b60046040517fef084b59000000000000000000000000000000000000000000000000000000008152fd5b8a8a7f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b60046040517f487c58e5000000000000000000000000000000000000000000000000000000008152fd5b60046040517f2ef424cc000000000000000000000000000000000000000000000000000000008152fd5b8380fd5b50346102245760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610224576126b6612886565b906024359067ffffffffffffffff92838316809303611c0b5763ffffffff168082526020926005845260408320818452845260026040842001549183526005845260408320908352835282600392836040822001936040519081968791818854948581520190819886528286209486915b8c82828501106127f4575050612769955491848d8383106127e2575b8383106127cc575b8383106127b6575b5050106127a8575b509050949392940386612b6a565b6040519460408601928652604082870152518092526060850193925b8281106127925785850386f35b8351871685529381019392810192600101612785565b60c01c81520184903861275b565b94600192958560801c168152019301848d612753565b94600192958560401c168152019301848d61274b565b8416855290930192600101848d612743565b87548181168652604081811c83168c880152608082811c9093169087015260c01c60608601526001909701968c96508995509093019260049290920191612727565b905034611c0b57817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611c0b5760209073ffffffffffffffffffffffffffffffffffffffff600154168152f35b6004359063ffffffff82168203610e6757565b6004359067ffffffffffffffff82168203610e6757565b9181601f84011215610e675782359167ffffffffffffffff8311610e675760208381860195010111610e6757565b919082519283825260005b8481106129285750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b6020818301810151848301820152016128e9565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203610e6757565b6044359073ffffffffffffffffffffffffffffffffffffffff82168203610e6757565b73ffffffffffffffffffffffffffffffffffffffff609f541633036129a357565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b609f549073ffffffffffffffffffffffffffffffffffffffff80911691827fffffffffffffffffffffffff0000000000000000000000000000000000000000821617609f55167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b15612a7557565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152fd5b610100810190811067ffffffffffffffff8211176108b257604052565b6040810190811067ffffffffffffffff8211176108b257604052565b6060810190811067ffffffffffffffff8211176108b257604052565b60a0810190811067ffffffffffffffff8211176108b257604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176108b257604052565b60ff60d15416612bb757565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152fd5b60ff60d1541615612c2257565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152fd5b90600182811c92168015612cc9575b6020831014612c9a57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691612c8f565b91908201809211610a6157565b91908203918211610a6157565b8115612cf7570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b91909180548310156108e157600052601860206000208360021c019260031b1690565b9060405191826000825492612d5d84612c80565b908184526001948581169081600014612dcc5750600114612d89575b5050612d8792500383612b6a565b565b9093915060005260209081600020936000915b818310612db4575050612d8793508201013880612d79565b85548884018501529485019487945091830191612d9c565b9050612d879550602093507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b8201013880612d79565b3d15612e67573d9067ffffffffffffffff82116108b25760405191612e5b60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8401160184612b6a565b82523d6000602084013e565b606090565b90816020910312610e6757518015158103610e67579056fea26469706673582212208d1a392674f136b26d5d6363df369b77aef099d2dd338b2dc6b30d542ec9ed7f64736f6c63430008140033",
}

// PromptSchedulerABI is the input ABI used to generate the binding from.
// Deprecated: Use PromptSchedulerMetaData.ABI instead.
var PromptSchedulerABI = PromptSchedulerMetaData.ABI

// PromptSchedulerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PromptSchedulerMetaData.Bin instead.
var PromptSchedulerBin = PromptSchedulerMetaData.Bin

// DeployPromptScheduler deploys a new Ethereum contract, binding an instance of PromptScheduler to it.
func DeployPromptScheduler(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PromptScheduler, error) {
	parsed, err := PromptSchedulerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PromptSchedulerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PromptScheduler{PromptSchedulerCaller: PromptSchedulerCaller{contract: contract}, PromptSchedulerTransactor: PromptSchedulerTransactor{contract: contract}, PromptSchedulerFilterer: PromptSchedulerFilterer{contract: contract}}, nil
}

// PromptScheduler is an auto generated Go binding around an Ethereum contract.
type PromptScheduler struct {
	PromptSchedulerCaller     // Read-only binding to the contract
	PromptSchedulerTransactor // Write-only binding to the contract
	PromptSchedulerFilterer   // Log filterer for contract events
}

// PromptSchedulerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PromptSchedulerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PromptSchedulerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PromptSchedulerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PromptSchedulerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PromptSchedulerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PromptSchedulerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PromptSchedulerSession struct {
	Contract     *PromptScheduler  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PromptSchedulerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PromptSchedulerCallerSession struct {
	Contract *PromptSchedulerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PromptSchedulerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PromptSchedulerTransactorSession struct {
	Contract     *PromptSchedulerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PromptSchedulerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PromptSchedulerRaw struct {
	Contract *PromptScheduler // Generic contract binding to access the raw methods on
}

// PromptSchedulerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PromptSchedulerCallerRaw struct {
	Contract *PromptSchedulerCaller // Generic read-only contract binding to access the raw methods on
}

// PromptSchedulerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PromptSchedulerTransactorRaw struct {
	Contract *PromptSchedulerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPromptScheduler creates a new instance of PromptScheduler, bound to a specific deployed contract.
func NewPromptScheduler(address common.Address, backend bind.ContractBackend) (*PromptScheduler, error) {
	contract, err := bindPromptScheduler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PromptScheduler{PromptSchedulerCaller: PromptSchedulerCaller{contract: contract}, PromptSchedulerTransactor: PromptSchedulerTransactor{contract: contract}, PromptSchedulerFilterer: PromptSchedulerFilterer{contract: contract}}, nil
}

// NewPromptSchedulerCaller creates a new read-only instance of PromptScheduler, bound to a specific deployed contract.
func NewPromptSchedulerCaller(address common.Address, caller bind.ContractCaller) (*PromptSchedulerCaller, error) {
	contract, err := bindPromptScheduler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PromptSchedulerCaller{contract: contract}, nil
}

// NewPromptSchedulerTransactor creates a new write-only instance of PromptScheduler, bound to a specific deployed contract.
func NewPromptSchedulerTransactor(address common.Address, transactor bind.ContractTransactor) (*PromptSchedulerTransactor, error) {
	contract, err := bindPromptScheduler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PromptSchedulerTransactor{contract: contract}, nil
}

// NewPromptSchedulerFilterer creates a new log filterer instance of PromptScheduler, bound to a specific deployed contract.
func NewPromptSchedulerFilterer(address common.Address, filterer bind.ContractFilterer) (*PromptSchedulerFilterer, error) {
	contract, err := bindPromptScheduler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PromptSchedulerFilterer{contract: contract}, nil
}

// bindPromptScheduler binds a generic wrapper to an already deployed contract.
func bindPromptScheduler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PromptSchedulerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PromptScheduler *PromptSchedulerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PromptScheduler.Contract.PromptSchedulerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PromptScheduler *PromptSchedulerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PromptScheduler.Contract.PromptSchedulerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PromptScheduler *PromptSchedulerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PromptScheduler.Contract.PromptSchedulerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PromptScheduler *PromptSchedulerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PromptScheduler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PromptScheduler *PromptSchedulerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PromptScheduler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PromptScheduler *PromptSchedulerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PromptScheduler.Contract.contract.Transact(opts, method, params...)
}

// BatchPeriod is a free data retrieval call binding the contract method 0x50370111.
//
// Solidity: function _batchPeriod() view returns(uint256)
func (_PromptScheduler *PromptSchedulerCaller) BatchPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "_batchPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchPeriod is a free data retrieval call binding the contract method 0x50370111.
//
// Solidity: function _batchPeriod() view returns(uint256)
func (_PromptScheduler *PromptSchedulerSession) BatchPeriod() (*big.Int, error) {
	return _PromptScheduler.Contract.BatchPeriod(&_PromptScheduler.CallOpts)
}

// BatchPeriod is a free data retrieval call binding the contract method 0x50370111.
//
// Solidity: function _batchPeriod() view returns(uint256)
func (_PromptScheduler *PromptSchedulerCallerSession) BatchPeriod() (*big.Int, error) {
	return _PromptScheduler.Contract.BatchPeriod(&_PromptScheduler.CallOpts)
}

// GpuManager is a free data retrieval call binding the contract method 0x08c147fd.
//
// Solidity: function _gpuManager() view returns(address)
func (_PromptScheduler *PromptSchedulerCaller) GpuManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "_gpuManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GpuManager is a free data retrieval call binding the contract method 0x08c147fd.
//
// Solidity: function _gpuManager() view returns(address)
func (_PromptScheduler *PromptSchedulerSession) GpuManager() (common.Address, error) {
	return _PromptScheduler.Contract.GpuManager(&_PromptScheduler.CallOpts)
}

// GpuManager is a free data retrieval call binding the contract method 0x08c147fd.
//
// Solidity: function _gpuManager() view returns(address)
func (_PromptScheduler *PromptSchedulerCallerSession) GpuManager() (common.Address, error) {
	return _PromptScheduler.Contract.GpuManager(&_PromptScheduler.CallOpts)
}

// InferenceCounter is a free data retrieval call binding the contract method 0x627a04b8.
//
// Solidity: function _inferenceCounter() view returns(uint64)
func (_PromptScheduler *PromptSchedulerCaller) InferenceCounter(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "_inferenceCounter")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// InferenceCounter is a free data retrieval call binding the contract method 0x627a04b8.
//
// Solidity: function _inferenceCounter() view returns(uint64)
func (_PromptScheduler *PromptSchedulerSession) InferenceCounter() (uint64, error) {
	return _PromptScheduler.Contract.InferenceCounter(&_PromptScheduler.CallOpts)
}

// InferenceCounter is a free data retrieval call binding the contract method 0x627a04b8.
//
// Solidity: function _inferenceCounter() view returns(uint64)
func (_PromptScheduler *PromptSchedulerCallerSession) InferenceCounter() (uint64, error) {
	return _PromptScheduler.Contract.InferenceCounter(&_PromptScheduler.CallOpts)
}

// LastBatchTimestamp is a free data retrieval call binding the contract method 0x7f8f29fc.
//
// Solidity: function _lastBatchTimestamp() view returns(uint256)
func (_PromptScheduler *PromptSchedulerCaller) LastBatchTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "_lastBatchTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBatchTimestamp is a free data retrieval call binding the contract method 0x7f8f29fc.
//
// Solidity: function _lastBatchTimestamp() view returns(uint256)
func (_PromptScheduler *PromptSchedulerSession) LastBatchTimestamp() (*big.Int, error) {
	return _PromptScheduler.Contract.LastBatchTimestamp(&_PromptScheduler.CallOpts)
}

// LastBatchTimestamp is a free data retrieval call binding the contract method 0x7f8f29fc.
//
// Solidity: function _lastBatchTimestamp() view returns(uint256)
func (_PromptScheduler *PromptSchedulerCallerSession) LastBatchTimestamp() (*big.Int, error) {
	return _PromptScheduler.Contract.LastBatchTimestamp(&_PromptScheduler.CallOpts)
}

// MinerRequirement is a free data retrieval call binding the contract method 0x87b97f1c.
//
// Solidity: function _minerRequirement() view returns(uint8)
func (_PromptScheduler *PromptSchedulerCaller) MinerRequirement(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "_minerRequirement")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MinerRequirement is a free data retrieval call binding the contract method 0x87b97f1c.
//
// Solidity: function _minerRequirement() view returns(uint8)
func (_PromptScheduler *PromptSchedulerSession) MinerRequirement() (uint8, error) {
	return _PromptScheduler.Contract.MinerRequirement(&_PromptScheduler.CallOpts)
}

// MinerRequirement is a free data retrieval call binding the contract method 0x87b97f1c.
//
// Solidity: function _minerRequirement() view returns(uint8)
func (_PromptScheduler *PromptSchedulerCallerSession) MinerRequirement() (uint8, error) {
	return _PromptScheduler.Contract.MinerRequirement(&_PromptScheduler.CallOpts)
}

// MinerValidatorFeeRatio is a free data retrieval call binding the contract method 0xa1e0a429.
//
// Solidity: function _minerValidatorFeeRatio() view returns(uint16)
func (_PromptScheduler *PromptSchedulerCaller) MinerValidatorFeeRatio(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "_minerValidatorFeeRatio")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MinerValidatorFeeRatio is a free data retrieval call binding the contract method 0xa1e0a429.
//
// Solidity: function _minerValidatorFeeRatio() view returns(uint16)
func (_PromptScheduler *PromptSchedulerSession) MinerValidatorFeeRatio() (uint16, error) {
	return _PromptScheduler.Contract.MinerValidatorFeeRatio(&_PromptScheduler.CallOpts)
}

// MinerValidatorFeeRatio is a free data retrieval call binding the contract method 0xa1e0a429.
//
// Solidity: function _minerValidatorFeeRatio() view returns(uint16)
func (_PromptScheduler *PromptSchedulerCallerSession) MinerValidatorFeeRatio() (uint16, error) {
	return _PromptScheduler.Contract.MinerValidatorFeeRatio(&_PromptScheduler.CallOpts)
}

// SubmitDuration is a free data retrieval call binding the contract method 0x7a80e13e.
//
// Solidity: function _submitDuration() view returns(uint40)
func (_PromptScheduler *PromptSchedulerCaller) SubmitDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "_submitDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubmitDuration is a free data retrieval call binding the contract method 0x7a80e13e.
//
// Solidity: function _submitDuration() view returns(uint40)
func (_PromptScheduler *PromptSchedulerSession) SubmitDuration() (*big.Int, error) {
	return _PromptScheduler.Contract.SubmitDuration(&_PromptScheduler.CallOpts)
}

// SubmitDuration is a free data retrieval call binding the contract method 0x7a80e13e.
//
// Solidity: function _submitDuration() view returns(uint40)
func (_PromptScheduler *PromptSchedulerCallerSession) SubmitDuration() (*big.Int, error) {
	return _PromptScheduler.Contract.SubmitDuration(&_PromptScheduler.CallOpts)
}

// WEAIToken is a free data retrieval call binding the contract method 0x871c15b1.
//
// Solidity: function _wEAIToken() view returns(address)
func (_PromptScheduler *PromptSchedulerCaller) WEAIToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "_wEAIToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WEAIToken is a free data retrieval call binding the contract method 0x871c15b1.
//
// Solidity: function _wEAIToken() view returns(address)
func (_PromptScheduler *PromptSchedulerSession) WEAIToken() (common.Address, error) {
	return _PromptScheduler.Contract.WEAIToken(&_PromptScheduler.CallOpts)
}

// WEAIToken is a free data retrieval call binding the contract method 0x871c15b1.
//
// Solidity: function _wEAIToken() view returns(address)
func (_PromptScheduler *PromptSchedulerCallerSession) WEAIToken() (common.Address, error) {
	return _PromptScheduler.Contract.WEAIToken(&_PromptScheduler.CallOpts)
}

// GetBatchInfo is a free data retrieval call binding the contract method 0x18717938.
//
// Solidity: function getBatchInfo(uint32 modelId, uint64 batchId) view returns(uint256, uint64[])
func (_PromptScheduler *PromptSchedulerCaller) GetBatchInfo(opts *bind.CallOpts, modelId uint32, batchId uint64) (*big.Int, []uint64, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "getBatchInfo", modelId, batchId)

	if err != nil {
		return *new(*big.Int), *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]uint64)).(*[]uint64)

	return out0, out1, err

}

// GetBatchInfo is a free data retrieval call binding the contract method 0x18717938.
//
// Solidity: function getBatchInfo(uint32 modelId, uint64 batchId) view returns(uint256, uint64[])
func (_PromptScheduler *PromptSchedulerSession) GetBatchInfo(modelId uint32, batchId uint64) (*big.Int, []uint64, error) {
	return _PromptScheduler.Contract.GetBatchInfo(&_PromptScheduler.CallOpts, modelId, batchId)
}

// GetBatchInfo is a free data retrieval call binding the contract method 0x18717938.
//
// Solidity: function getBatchInfo(uint32 modelId, uint64 batchId) view returns(uint256, uint64[])
func (_PromptScheduler *PromptSchedulerCallerSession) GetBatchInfo(modelId uint32, batchId uint64) (*big.Int, []uint64, error) {
	return _PromptScheduler.Contract.GetBatchInfo(&_PromptScheduler.CallOpts, modelId, batchId)
}

// GetInferenceByMiner is a free data retrieval call binding the contract method 0x56301806.
//
// Solidity: function getInferenceByMiner(address miner) view returns(uint256[])
func (_PromptScheduler *PromptSchedulerCaller) GetInferenceByMiner(opts *bind.CallOpts, miner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "getInferenceByMiner", miner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetInferenceByMiner is a free data retrieval call binding the contract method 0x56301806.
//
// Solidity: function getInferenceByMiner(address miner) view returns(uint256[])
func (_PromptScheduler *PromptSchedulerSession) GetInferenceByMiner(miner common.Address) ([]*big.Int, error) {
	return _PromptScheduler.Contract.GetInferenceByMiner(&_PromptScheduler.CallOpts, miner)
}

// GetInferenceByMiner is a free data retrieval call binding the contract method 0x56301806.
//
// Solidity: function getInferenceByMiner(address miner) view returns(uint256[])
func (_PromptScheduler *PromptSchedulerCallerSession) GetInferenceByMiner(miner common.Address) ([]*big.Int, error) {
	return _PromptScheduler.Contract.GetInferenceByMiner(&_PromptScheduler.CallOpts, miner)
}

// GetInferenceInfo is a free data retrieval call binding the contract method 0x48729262.
//
// Solidity: function getInferenceInfo(uint64 inferId) view returns((uint256,uint32,uint40,uint8,address,address,bytes,bytes))
func (_PromptScheduler *PromptSchedulerCaller) GetInferenceInfo(opts *bind.CallOpts, inferId uint64) (ISchedulerInference, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "getInferenceInfo", inferId)

	if err != nil {
		return *new(ISchedulerInference), err
	}

	out0 := *abi.ConvertType(out[0], new(ISchedulerInference)).(*ISchedulerInference)

	return out0, err

}

// GetInferenceInfo is a free data retrieval call binding the contract method 0x48729262.
//
// Solidity: function getInferenceInfo(uint64 inferId) view returns((uint256,uint32,uint40,uint8,address,address,bytes,bytes))
func (_PromptScheduler *PromptSchedulerSession) GetInferenceInfo(inferId uint64) (ISchedulerInference, error) {
	return _PromptScheduler.Contract.GetInferenceInfo(&_PromptScheduler.CallOpts, inferId)
}

// GetInferenceInfo is a free data retrieval call binding the contract method 0x48729262.
//
// Solidity: function getInferenceInfo(uint64 inferId) view returns((uint256,uint32,uint40,uint8,address,address,bytes,bytes))
func (_PromptScheduler *PromptSchedulerCallerSession) GetInferenceInfo(inferId uint64) (ISchedulerInference, error) {
	return _PromptScheduler.Contract.GetInferenceInfo(&_PromptScheduler.CallOpts, inferId)
}

// GetMinerRequirement is a free data retrieval call binding the contract method 0x48751e50.
//
// Solidity: function getMinerRequirement() view returns(uint8)
func (_PromptScheduler *PromptSchedulerCaller) GetMinerRequirement(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "getMinerRequirement")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetMinerRequirement is a free data retrieval call binding the contract method 0x48751e50.
//
// Solidity: function getMinerRequirement() view returns(uint8)
func (_PromptScheduler *PromptSchedulerSession) GetMinerRequirement() (uint8, error) {
	return _PromptScheduler.Contract.GetMinerRequirement(&_PromptScheduler.CallOpts)
}

// GetMinerRequirement is a free data retrieval call binding the contract method 0x48751e50.
//
// Solidity: function getMinerRequirement() view returns(uint8)
func (_PromptScheduler *PromptSchedulerCallerSession) GetMinerRequirement() (uint8, error) {
	return _PromptScheduler.Contract.GetMinerRequirement(&_PromptScheduler.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PromptScheduler *PromptSchedulerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PromptScheduler *PromptSchedulerSession) Owner() (common.Address, error) {
	return _PromptScheduler.Contract.Owner(&_PromptScheduler.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PromptScheduler *PromptSchedulerCallerSession) Owner() (common.Address, error) {
	return _PromptScheduler.Contract.Owner(&_PromptScheduler.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PromptScheduler *PromptSchedulerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PromptScheduler *PromptSchedulerSession) Paused() (bool, error) {
	return _PromptScheduler.Contract.Paused(&_PromptScheduler.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PromptScheduler *PromptSchedulerCallerSession) Paused() (bool, error) {
	return _PromptScheduler.Contract.Paused(&_PromptScheduler.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_PromptScheduler *PromptSchedulerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PromptScheduler.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_PromptScheduler *PromptSchedulerSession) Version() (string, error) {
	return _PromptScheduler.Contract.Version(&_PromptScheduler.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_PromptScheduler *PromptSchedulerCallerSession) Version() (string, error) {
	return _PromptScheduler.Contract.Version(&_PromptScheduler.CallOpts)
}

// Infer is a paid mutator transaction binding the contract method 0x5cc68731.
//
// Solidity: function infer(uint32 modelId, bytes input, address creator, bool flag) returns(uint64)
func (_PromptScheduler *PromptSchedulerTransactor) Infer(opts *bind.TransactOpts, modelId uint32, input []byte, creator common.Address, flag bool) (*types.Transaction, error) {
	return _PromptScheduler.contract.Transact(opts, "infer", modelId, input, creator, flag)
}

// Infer is a paid mutator transaction binding the contract method 0x5cc68731.
//
// Solidity: function infer(uint32 modelId, bytes input, address creator, bool flag) returns(uint64)
func (_PromptScheduler *PromptSchedulerSession) Infer(modelId uint32, input []byte, creator common.Address, flag bool) (*types.Transaction, error) {
	return _PromptScheduler.Contract.Infer(&_PromptScheduler.TransactOpts, modelId, input, creator, flag)
}

// Infer is a paid mutator transaction binding the contract method 0x5cc68731.
//
// Solidity: function infer(uint32 modelId, bytes input, address creator, bool flag) returns(uint64)
func (_PromptScheduler *PromptSchedulerTransactorSession) Infer(modelId uint32, input []byte, creator common.Address, flag bool) (*types.Transaction, error) {
	return _PromptScheduler.Contract.Infer(&_PromptScheduler.TransactOpts, modelId, input, creator, flag)
}

// Infer0 is a paid mutator transaction binding the contract method 0xde1ce2bb.
//
// Solidity: function infer(uint32 modelId, bytes input, address creator) returns(uint64)
func (_PromptScheduler *PromptSchedulerTransactor) Infer0(opts *bind.TransactOpts, modelId uint32, input []byte, creator common.Address) (*types.Transaction, error) {
	return _PromptScheduler.contract.Transact(opts, "infer0", modelId, input, creator)
}

// Infer0 is a paid mutator transaction binding the contract method 0xde1ce2bb.
//
// Solidity: function infer(uint32 modelId, bytes input, address creator) returns(uint64)
func (_PromptScheduler *PromptSchedulerSession) Infer0(modelId uint32, input []byte, creator common.Address) (*types.Transaction, error) {
	return _PromptScheduler.Contract.Infer0(&_PromptScheduler.TransactOpts, modelId, input, creator)
}

// Infer0 is a paid mutator transaction binding the contract method 0xde1ce2bb.
//
// Solidity: function infer(uint32 modelId, bytes input, address creator) returns(uint64)
func (_PromptScheduler *PromptSchedulerTransactorSession) Infer0(modelId uint32, input []byte, creator common.Address) (*types.Transaction, error) {
	return _PromptScheduler.Contract.Infer0(&_PromptScheduler.TransactOpts, modelId, input, creator)
}

// Initialize is a paid mutator transaction binding the contract method 0xa50d8600.
//
// Solidity: function initialize(address wEAIToken_, address gpuManager_, uint8 minerRequirement_, uint40 submitDuration_, uint16 minerValidatorFeeRatio_, uint40 batchPeriod_) returns()
func (_PromptScheduler *PromptSchedulerTransactor) Initialize(opts *bind.TransactOpts, wEAIToken_ common.Address, gpuManager_ common.Address, minerRequirement_ uint8, submitDuration_ *big.Int, minerValidatorFeeRatio_ uint16, batchPeriod_ *big.Int) (*types.Transaction, error) {
	return _PromptScheduler.contract.Transact(opts, "initialize", wEAIToken_, gpuManager_, minerRequirement_, submitDuration_, minerValidatorFeeRatio_, batchPeriod_)
}

// Initialize is a paid mutator transaction binding the contract method 0xa50d8600.
//
// Solidity: function initialize(address wEAIToken_, address gpuManager_, uint8 minerRequirement_, uint40 submitDuration_, uint16 minerValidatorFeeRatio_, uint40 batchPeriod_) returns()
func (_PromptScheduler *PromptSchedulerSession) Initialize(wEAIToken_ common.Address, gpuManager_ common.Address, minerRequirement_ uint8, submitDuration_ *big.Int, minerValidatorFeeRatio_ uint16, batchPeriod_ *big.Int) (*types.Transaction, error) {
	return _PromptScheduler.Contract.Initialize(&_PromptScheduler.TransactOpts, wEAIToken_, gpuManager_, minerRequirement_, submitDuration_, minerValidatorFeeRatio_, batchPeriod_)
}

// Initialize is a paid mutator transaction binding the contract method 0xa50d8600.
//
// Solidity: function initialize(address wEAIToken_, address gpuManager_, uint8 minerRequirement_, uint40 submitDuration_, uint16 minerValidatorFeeRatio_, uint40 batchPeriod_) returns()
func (_PromptScheduler *PromptSchedulerTransactorSession) Initialize(wEAIToken_ common.Address, gpuManager_ common.Address, minerRequirement_ uint8, submitDuration_ *big.Int, minerValidatorFeeRatio_ uint16, batchPeriod_ *big.Int) (*types.Transaction, error) {
	return _PromptScheduler.Contract.Initialize(&_PromptScheduler.TransactOpts, wEAIToken_, gpuManager_, minerRequirement_, submitDuration_, minerValidatorFeeRatio_, batchPeriod_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PromptScheduler *PromptSchedulerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PromptScheduler.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PromptScheduler *PromptSchedulerSession) Pause() (*types.Transaction, error) {
	return _PromptScheduler.Contract.Pause(&_PromptScheduler.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PromptScheduler *PromptSchedulerTransactorSession) Pause() (*types.Transaction, error) {
	return _PromptScheduler.Contract.Pause(&_PromptScheduler.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PromptScheduler *PromptSchedulerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PromptScheduler.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PromptScheduler *PromptSchedulerSession) RenounceOwnership() (*types.Transaction, error) {
	return _PromptScheduler.Contract.RenounceOwnership(&_PromptScheduler.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PromptScheduler *PromptSchedulerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PromptScheduler.Contract.RenounceOwnership(&_PromptScheduler.TransactOpts)
}

// SetSubmitDuration is a paid mutator transaction binding the contract method 0x6f643736.
//
// Solidity: function setSubmitDuration(uint40 submitDuration) returns()
func (_PromptScheduler *PromptSchedulerTransactor) SetSubmitDuration(opts *bind.TransactOpts, submitDuration *big.Int) (*types.Transaction, error) {
	return _PromptScheduler.contract.Transact(opts, "setSubmitDuration", submitDuration)
}

// SetSubmitDuration is a paid mutator transaction binding the contract method 0x6f643736.
//
// Solidity: function setSubmitDuration(uint40 submitDuration) returns()
func (_PromptScheduler *PromptSchedulerSession) SetSubmitDuration(submitDuration *big.Int) (*types.Transaction, error) {
	return _PromptScheduler.Contract.SetSubmitDuration(&_PromptScheduler.TransactOpts, submitDuration)
}

// SetSubmitDuration is a paid mutator transaction binding the contract method 0x6f643736.
//
// Solidity: function setSubmitDuration(uint40 submitDuration) returns()
func (_PromptScheduler *PromptSchedulerTransactorSession) SetSubmitDuration(submitDuration *big.Int) (*types.Transaction, error) {
	return _PromptScheduler.Contract.SetSubmitDuration(&_PromptScheduler.TransactOpts, submitDuration)
}

// SetWEAIAddress is a paid mutator transaction binding the contract method 0x7362323c.
//
// Solidity: function setWEAIAddress(address wEAIToken) returns()
func (_PromptScheduler *PromptSchedulerTransactor) SetWEAIAddress(opts *bind.TransactOpts, wEAIToken common.Address) (*types.Transaction, error) {
	return _PromptScheduler.contract.Transact(opts, "setWEAIAddress", wEAIToken)
}

// SetWEAIAddress is a paid mutator transaction binding the contract method 0x7362323c.
//
// Solidity: function setWEAIAddress(address wEAIToken) returns()
func (_PromptScheduler *PromptSchedulerSession) SetWEAIAddress(wEAIToken common.Address) (*types.Transaction, error) {
	return _PromptScheduler.Contract.SetWEAIAddress(&_PromptScheduler.TransactOpts, wEAIToken)
}

// SetWEAIAddress is a paid mutator transaction binding the contract method 0x7362323c.
//
// Solidity: function setWEAIAddress(address wEAIToken) returns()
func (_PromptScheduler *PromptSchedulerTransactorSession) SetWEAIAddress(wEAIToken common.Address) (*types.Transaction, error) {
	return _PromptScheduler.Contract.SetWEAIAddress(&_PromptScheduler.TransactOpts, wEAIToken)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0x34b96ee4.
//
// Solidity: function submitSolution(uint64 inferId, bytes solution) returns()
func (_PromptScheduler *PromptSchedulerTransactor) SubmitSolution(opts *bind.TransactOpts, inferId uint64, solution []byte) (*types.Transaction, error) {
	return _PromptScheduler.contract.Transact(opts, "submitSolution", inferId, solution)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0x34b96ee4.
//
// Solidity: function submitSolution(uint64 inferId, bytes solution) returns()
func (_PromptScheduler *PromptSchedulerSession) SubmitSolution(inferId uint64, solution []byte) (*types.Transaction, error) {
	return _PromptScheduler.Contract.SubmitSolution(&_PromptScheduler.TransactOpts, inferId, solution)
}

// SubmitSolution is a paid mutator transaction binding the contract method 0x34b96ee4.
//
// Solidity: function submitSolution(uint64 inferId, bytes solution) returns()
func (_PromptScheduler *PromptSchedulerTransactorSession) SubmitSolution(inferId uint64, solution []byte) (*types.Transaction, error) {
	return _PromptScheduler.Contract.SubmitSolution(&_PromptScheduler.TransactOpts, inferId, solution)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PromptScheduler *PromptSchedulerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PromptScheduler.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PromptScheduler *PromptSchedulerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PromptScheduler.Contract.TransferOwnership(&_PromptScheduler.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PromptScheduler *PromptSchedulerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PromptScheduler.Contract.TransferOwnership(&_PromptScheduler.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PromptScheduler *PromptSchedulerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PromptScheduler.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PromptScheduler *PromptSchedulerSession) Unpause() (*types.Transaction, error) {
	return _PromptScheduler.Contract.Unpause(&_PromptScheduler.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PromptScheduler *PromptSchedulerTransactorSession) Unpause() (*types.Transaction, error) {
	return _PromptScheduler.Contract.Unpause(&_PromptScheduler.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PromptScheduler *PromptSchedulerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PromptScheduler.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PromptScheduler *PromptSchedulerSession) Receive() (*types.Transaction, error) {
	return _PromptScheduler.Contract.Receive(&_PromptScheduler.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PromptScheduler *PromptSchedulerTransactorSession) Receive() (*types.Transaction, error) {
	return _PromptScheduler.Contract.Receive(&_PromptScheduler.TransactOpts)
}

// PromptSchedulerAppendToBatchIterator is returned from FilterAppendToBatch and is used to iterate over the raw logs and unpacked data for AppendToBatch events raised by the PromptScheduler contract.
type PromptSchedulerAppendToBatchIterator struct {
	Event *PromptSchedulerAppendToBatch // Event containing the contract specifics and raw log

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
func (it *PromptSchedulerAppendToBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptSchedulerAppendToBatch)
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
		it.Event = new(PromptSchedulerAppendToBatch)
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
func (it *PromptSchedulerAppendToBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptSchedulerAppendToBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptSchedulerAppendToBatch represents a AppendToBatch event raised by the PromptScheduler contract.
type PromptSchedulerAppendToBatch struct {
	BatchId uint64
	ModelId uint32
	InferId uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAppendToBatch is a free log retrieval operation binding the contract event 0xb51ff9d6e56ade059ce28900d43a2402f306a6f04f1ce592fe2e24b817d5c8cd.
//
// Solidity: event AppendToBatch(uint64 indexed batchId, uint32 indexed modelId, uint64 indexed inferId)
func (_PromptScheduler *PromptSchedulerFilterer) FilterAppendToBatch(opts *bind.FilterOpts, batchId []uint64, modelId []uint32, inferId []uint64) (*PromptSchedulerAppendToBatchIterator, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}
	var inferIdRule []interface{}
	for _, inferIdItem := range inferId {
		inferIdRule = append(inferIdRule, inferIdItem)
	}

	logs, sub, err := _PromptScheduler.contract.FilterLogs(opts, "AppendToBatch", batchIdRule, modelIdRule, inferIdRule)
	if err != nil {
		return nil, err
	}
	return &PromptSchedulerAppendToBatchIterator{contract: _PromptScheduler.contract, event: "AppendToBatch", logs: logs, sub: sub}, nil
}

// WatchAppendToBatch is a free log subscription operation binding the contract event 0xb51ff9d6e56ade059ce28900d43a2402f306a6f04f1ce592fe2e24b817d5c8cd.
//
// Solidity: event AppendToBatch(uint64 indexed batchId, uint32 indexed modelId, uint64 indexed inferId)
func (_PromptScheduler *PromptSchedulerFilterer) WatchAppendToBatch(opts *bind.WatchOpts, sink chan<- *PromptSchedulerAppendToBatch, batchId []uint64, modelId []uint32, inferId []uint64) (event.Subscription, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}
	var inferIdRule []interface{}
	for _, inferIdItem := range inferId {
		inferIdRule = append(inferIdRule, inferIdItem)
	}

	logs, sub, err := _PromptScheduler.contract.WatchLogs(opts, "AppendToBatch", batchIdRule, modelIdRule, inferIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptSchedulerAppendToBatch)
				if err := _PromptScheduler.contract.UnpackLog(event, "AppendToBatch", log); err != nil {
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

// ParseAppendToBatch is a log parse operation binding the contract event 0xb51ff9d6e56ade059ce28900d43a2402f306a6f04f1ce592fe2e24b817d5c8cd.
//
// Solidity: event AppendToBatch(uint64 indexed batchId, uint32 indexed modelId, uint64 indexed inferId)
func (_PromptScheduler *PromptSchedulerFilterer) ParseAppendToBatch(log types.Log) (*PromptSchedulerAppendToBatch, error) {
	event := new(PromptSchedulerAppendToBatch)
	if err := _PromptScheduler.contract.UnpackLog(event, "AppendToBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromptSchedulerInferenceStatusUpdateIterator is returned from FilterInferenceStatusUpdate and is used to iterate over the raw logs and unpacked data for InferenceStatusUpdate events raised by the PromptScheduler contract.
type PromptSchedulerInferenceStatusUpdateIterator struct {
	Event *PromptSchedulerInferenceStatusUpdate // Event containing the contract specifics and raw log

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
func (it *PromptSchedulerInferenceStatusUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptSchedulerInferenceStatusUpdate)
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
		it.Event = new(PromptSchedulerInferenceStatusUpdate)
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
func (it *PromptSchedulerInferenceStatusUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptSchedulerInferenceStatusUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptSchedulerInferenceStatusUpdate represents a InferenceStatusUpdate event raised by the PromptScheduler contract.
type PromptSchedulerInferenceStatusUpdate struct {
	InferenceId uint64
	NewStatus   uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInferenceStatusUpdate is a free log retrieval operation binding the contract event 0x79e6e7d96ce1a52bc5d92a4a2458c9647a3dd14e184fddfef0d2e8204f05a582.
//
// Solidity: event InferenceStatusUpdate(uint64 indexed inferenceId, uint8 newStatus)
func (_PromptScheduler *PromptSchedulerFilterer) FilterInferenceStatusUpdate(opts *bind.FilterOpts, inferenceId []uint64) (*PromptSchedulerInferenceStatusUpdateIterator, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}

	logs, sub, err := _PromptScheduler.contract.FilterLogs(opts, "InferenceStatusUpdate", inferenceIdRule)
	if err != nil {
		return nil, err
	}
	return &PromptSchedulerInferenceStatusUpdateIterator{contract: _PromptScheduler.contract, event: "InferenceStatusUpdate", logs: logs, sub: sub}, nil
}

// WatchInferenceStatusUpdate is a free log subscription operation binding the contract event 0x79e6e7d96ce1a52bc5d92a4a2458c9647a3dd14e184fddfef0d2e8204f05a582.
//
// Solidity: event InferenceStatusUpdate(uint64 indexed inferenceId, uint8 newStatus)
func (_PromptScheduler *PromptSchedulerFilterer) WatchInferenceStatusUpdate(opts *bind.WatchOpts, sink chan<- *PromptSchedulerInferenceStatusUpdate, inferenceId []uint64) (event.Subscription, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}

	logs, sub, err := _PromptScheduler.contract.WatchLogs(opts, "InferenceStatusUpdate", inferenceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptSchedulerInferenceStatusUpdate)
				if err := _PromptScheduler.contract.UnpackLog(event, "InferenceStatusUpdate", log); err != nil {
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

// ParseInferenceStatusUpdate is a log parse operation binding the contract event 0x79e6e7d96ce1a52bc5d92a4a2458c9647a3dd14e184fddfef0d2e8204f05a582.
//
// Solidity: event InferenceStatusUpdate(uint64 indexed inferenceId, uint8 newStatus)
func (_PromptScheduler *PromptSchedulerFilterer) ParseInferenceStatusUpdate(log types.Log) (*PromptSchedulerInferenceStatusUpdate, error) {
	event := new(PromptSchedulerInferenceStatusUpdate)
	if err := _PromptScheduler.contract.UnpackLog(event, "InferenceStatusUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromptSchedulerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PromptScheduler contract.
type PromptSchedulerInitializedIterator struct {
	Event *PromptSchedulerInitialized // Event containing the contract specifics and raw log

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
func (it *PromptSchedulerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptSchedulerInitialized)
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
		it.Event = new(PromptSchedulerInitialized)
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
func (it *PromptSchedulerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptSchedulerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptSchedulerInitialized represents a Initialized event raised by the PromptScheduler contract.
type PromptSchedulerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PromptScheduler *PromptSchedulerFilterer) FilterInitialized(opts *bind.FilterOpts) (*PromptSchedulerInitializedIterator, error) {

	logs, sub, err := _PromptScheduler.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PromptSchedulerInitializedIterator{contract: _PromptScheduler.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PromptScheduler *PromptSchedulerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PromptSchedulerInitialized) (event.Subscription, error) {

	logs, sub, err := _PromptScheduler.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptSchedulerInitialized)
				if err := _PromptScheduler.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PromptScheduler *PromptSchedulerFilterer) ParseInitialized(log types.Log) (*PromptSchedulerInitialized, error) {
	event := new(PromptSchedulerInitialized)
	if err := _PromptScheduler.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromptSchedulerNewAssignmentIterator is returned from FilterNewAssignment and is used to iterate over the raw logs and unpacked data for NewAssignment events raised by the PromptScheduler contract.
type PromptSchedulerNewAssignmentIterator struct {
	Event *PromptSchedulerNewAssignment // Event containing the contract specifics and raw log

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
func (it *PromptSchedulerNewAssignmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptSchedulerNewAssignment)
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
		it.Event = new(PromptSchedulerNewAssignment)
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
func (it *PromptSchedulerNewAssignmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptSchedulerNewAssignmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptSchedulerNewAssignment represents a NewAssignment event raised by the PromptScheduler contract.
type PromptSchedulerNewAssignment struct {
	InferenceId uint64
	Miner       common.Address
	ExpiredAt   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewAssignment is a free log retrieval operation binding the contract event 0xce7f171f53023fc0778971a0fe3f4067468c0505e2485b2716e44c4487215e79.
//
// Solidity: event NewAssignment(uint64 indexed inferenceId, address indexed miner, uint40 expiredAt)
func (_PromptScheduler *PromptSchedulerFilterer) FilterNewAssignment(opts *bind.FilterOpts, inferenceId []uint64, miner []common.Address) (*PromptSchedulerNewAssignmentIterator, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _PromptScheduler.contract.FilterLogs(opts, "NewAssignment", inferenceIdRule, minerRule)
	if err != nil {
		return nil, err
	}
	return &PromptSchedulerNewAssignmentIterator{contract: _PromptScheduler.contract, event: "NewAssignment", logs: logs, sub: sub}, nil
}

// WatchNewAssignment is a free log subscription operation binding the contract event 0xce7f171f53023fc0778971a0fe3f4067468c0505e2485b2716e44c4487215e79.
//
// Solidity: event NewAssignment(uint64 indexed inferenceId, address indexed miner, uint40 expiredAt)
func (_PromptScheduler *PromptSchedulerFilterer) WatchNewAssignment(opts *bind.WatchOpts, sink chan<- *PromptSchedulerNewAssignment, inferenceId []uint64, miner []common.Address) (event.Subscription, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _PromptScheduler.contract.WatchLogs(opts, "NewAssignment", inferenceIdRule, minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptSchedulerNewAssignment)
				if err := _PromptScheduler.contract.UnpackLog(event, "NewAssignment", log); err != nil {
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

// ParseNewAssignment is a log parse operation binding the contract event 0xce7f171f53023fc0778971a0fe3f4067468c0505e2485b2716e44c4487215e79.
//
// Solidity: event NewAssignment(uint64 indexed inferenceId, address indexed miner, uint40 expiredAt)
func (_PromptScheduler *PromptSchedulerFilterer) ParseNewAssignment(log types.Log) (*PromptSchedulerNewAssignment, error) {
	event := new(PromptSchedulerNewAssignment)
	if err := _PromptScheduler.contract.UnpackLog(event, "NewAssignment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromptSchedulerNewInferenceIterator is returned from FilterNewInference and is used to iterate over the raw logs and unpacked data for NewInference events raised by the PromptScheduler contract.
type PromptSchedulerNewInferenceIterator struct {
	Event *PromptSchedulerNewInference // Event containing the contract specifics and raw log

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
func (it *PromptSchedulerNewInferenceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptSchedulerNewInference)
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
		it.Event = new(PromptSchedulerNewInference)
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
func (it *PromptSchedulerNewInferenceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptSchedulerNewInferenceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptSchedulerNewInference represents a NewInference event raised by the PromptScheduler contract.
type PromptSchedulerNewInference struct {
	InferenceId uint64
	Creator     common.Address
	ModelId     uint32
	Value       *big.Int
	Input       []byte
	Flag        bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewInference is a free log retrieval operation binding the contract event 0x964ad1b4133d111aec4a94ec6525e35e7e6793cdb76b507a298ac273ef85c017.
//
// Solidity: event NewInference(uint64 indexed inferenceId, address indexed creator, uint32 indexed modelId, uint256 value, bytes input, bool flag)
func (_PromptScheduler *PromptSchedulerFilterer) FilterNewInference(opts *bind.FilterOpts, inferenceId []uint64, creator []common.Address, modelId []uint32) (*PromptSchedulerNewInferenceIterator, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _PromptScheduler.contract.FilterLogs(opts, "NewInference", inferenceIdRule, creatorRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return &PromptSchedulerNewInferenceIterator{contract: _PromptScheduler.contract, event: "NewInference", logs: logs, sub: sub}, nil
}

// WatchNewInference is a free log subscription operation binding the contract event 0x964ad1b4133d111aec4a94ec6525e35e7e6793cdb76b507a298ac273ef85c017.
//
// Solidity: event NewInference(uint64 indexed inferenceId, address indexed creator, uint32 indexed modelId, uint256 value, bytes input, bool flag)
func (_PromptScheduler *PromptSchedulerFilterer) WatchNewInference(opts *bind.WatchOpts, sink chan<- *PromptSchedulerNewInference, inferenceId []uint64, creator []common.Address, modelId []uint32) (event.Subscription, error) {

	var inferenceIdRule []interface{}
	for _, inferenceIdItem := range inferenceId {
		inferenceIdRule = append(inferenceIdRule, inferenceIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _PromptScheduler.contract.WatchLogs(opts, "NewInference", inferenceIdRule, creatorRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptSchedulerNewInference)
				if err := _PromptScheduler.contract.UnpackLog(event, "NewInference", log); err != nil {
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

// ParseNewInference is a log parse operation binding the contract event 0x964ad1b4133d111aec4a94ec6525e35e7e6793cdb76b507a298ac273ef85c017.
//
// Solidity: event NewInference(uint64 indexed inferenceId, address indexed creator, uint32 indexed modelId, uint256 value, bytes input, bool flag)
func (_PromptScheduler *PromptSchedulerFilterer) ParseNewInference(log types.Log) (*PromptSchedulerNewInference, error) {
	event := new(PromptSchedulerNewInference)
	if err := _PromptScheduler.contract.UnpackLog(event, "NewInference", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromptSchedulerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PromptScheduler contract.
type PromptSchedulerOwnershipTransferredIterator struct {
	Event *PromptSchedulerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PromptSchedulerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptSchedulerOwnershipTransferred)
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
		it.Event = new(PromptSchedulerOwnershipTransferred)
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
func (it *PromptSchedulerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptSchedulerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptSchedulerOwnershipTransferred represents a OwnershipTransferred event raised by the PromptScheduler contract.
type PromptSchedulerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PromptScheduler *PromptSchedulerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PromptSchedulerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PromptScheduler.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PromptSchedulerOwnershipTransferredIterator{contract: _PromptScheduler.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PromptScheduler *PromptSchedulerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PromptSchedulerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PromptScheduler.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptSchedulerOwnershipTransferred)
				if err := _PromptScheduler.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PromptScheduler *PromptSchedulerFilterer) ParseOwnershipTransferred(log types.Log) (*PromptSchedulerOwnershipTransferred, error) {
	event := new(PromptSchedulerOwnershipTransferred)
	if err := _PromptScheduler.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromptSchedulerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PromptScheduler contract.
type PromptSchedulerPausedIterator struct {
	Event *PromptSchedulerPaused // Event containing the contract specifics and raw log

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
func (it *PromptSchedulerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptSchedulerPaused)
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
		it.Event = new(PromptSchedulerPaused)
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
func (it *PromptSchedulerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptSchedulerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptSchedulerPaused represents a Paused event raised by the PromptScheduler contract.
type PromptSchedulerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PromptScheduler *PromptSchedulerFilterer) FilterPaused(opts *bind.FilterOpts) (*PromptSchedulerPausedIterator, error) {

	logs, sub, err := _PromptScheduler.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PromptSchedulerPausedIterator{contract: _PromptScheduler.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PromptScheduler *PromptSchedulerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PromptSchedulerPaused) (event.Subscription, error) {

	logs, sub, err := _PromptScheduler.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptSchedulerPaused)
				if err := _PromptScheduler.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_PromptScheduler *PromptSchedulerFilterer) ParsePaused(log types.Log) (*PromptSchedulerPaused, error) {
	event := new(PromptSchedulerPaused)
	if err := _PromptScheduler.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromptSchedulerSolutionSubmissionIterator is returned from FilterSolutionSubmission and is used to iterate over the raw logs and unpacked data for SolutionSubmission events raised by the PromptScheduler contract.
type PromptSchedulerSolutionSubmissionIterator struct {
	Event *PromptSchedulerSolutionSubmission // Event containing the contract specifics and raw log

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
func (it *PromptSchedulerSolutionSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptSchedulerSolutionSubmission)
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
		it.Event = new(PromptSchedulerSolutionSubmission)
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
func (it *PromptSchedulerSolutionSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptSchedulerSolutionSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptSchedulerSolutionSubmission represents a SolutionSubmission event raised by the PromptScheduler contract.
type PromptSchedulerSolutionSubmission struct {
	Miner   common.Address
	InferId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSolutionSubmission is a free log retrieval operation binding the contract event 0x9f669b92b9cbc7611f7ab6c77db07a424051c777433e21bd90f1bdf940096dd9.
//
// Solidity: event SolutionSubmission(address indexed miner, uint256 indexed inferId)
func (_PromptScheduler *PromptSchedulerFilterer) FilterSolutionSubmission(opts *bind.FilterOpts, miner []common.Address, inferId []*big.Int) (*PromptSchedulerSolutionSubmissionIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var inferIdRule []interface{}
	for _, inferIdItem := range inferId {
		inferIdRule = append(inferIdRule, inferIdItem)
	}

	logs, sub, err := _PromptScheduler.contract.FilterLogs(opts, "SolutionSubmission", minerRule, inferIdRule)
	if err != nil {
		return nil, err
	}
	return &PromptSchedulerSolutionSubmissionIterator{contract: _PromptScheduler.contract, event: "SolutionSubmission", logs: logs, sub: sub}, nil
}

// WatchSolutionSubmission is a free log subscription operation binding the contract event 0x9f669b92b9cbc7611f7ab6c77db07a424051c777433e21bd90f1bdf940096dd9.
//
// Solidity: event SolutionSubmission(address indexed miner, uint256 indexed inferId)
func (_PromptScheduler *PromptSchedulerFilterer) WatchSolutionSubmission(opts *bind.WatchOpts, sink chan<- *PromptSchedulerSolutionSubmission, miner []common.Address, inferId []*big.Int) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var inferIdRule []interface{}
	for _, inferIdItem := range inferId {
		inferIdRule = append(inferIdRule, inferIdItem)
	}

	logs, sub, err := _PromptScheduler.contract.WatchLogs(opts, "SolutionSubmission", minerRule, inferIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptSchedulerSolutionSubmission)
				if err := _PromptScheduler.contract.UnpackLog(event, "SolutionSubmission", log); err != nil {
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
// Solidity: event SolutionSubmission(address indexed miner, uint256 indexed inferId)
func (_PromptScheduler *PromptSchedulerFilterer) ParseSolutionSubmission(log types.Log) (*PromptSchedulerSolutionSubmission, error) {
	event := new(PromptSchedulerSolutionSubmission)
	if err := _PromptScheduler.contract.UnpackLog(event, "SolutionSubmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromptSchedulerStreamedDataIterator is returned from FilterStreamedData and is used to iterate over the raw logs and unpacked data for StreamedData events raised by the PromptScheduler contract.
type PromptSchedulerStreamedDataIterator struct {
	Event *PromptSchedulerStreamedData // Event containing the contract specifics and raw log

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
func (it *PromptSchedulerStreamedDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptSchedulerStreamedData)
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
		it.Event = new(PromptSchedulerStreamedData)
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
func (it *PromptSchedulerStreamedDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptSchedulerStreamedDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptSchedulerStreamedData represents a StreamedData event raised by the PromptScheduler contract.
type PromptSchedulerStreamedData struct {
	AssignmentId *big.Int
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStreamedData is a free log retrieval operation binding the contract event 0x23cfaa418b5f569ff36b152a9fd02ee3ccddaa5f7eed570e777a30353b68dc38.
//
// Solidity: event StreamedData(uint256 indexed assignmentId, bytes data)
func (_PromptScheduler *PromptSchedulerFilterer) FilterStreamedData(opts *bind.FilterOpts, assignmentId []*big.Int) (*PromptSchedulerStreamedDataIterator, error) {

	var assignmentIdRule []interface{}
	for _, assignmentIdItem := range assignmentId {
		assignmentIdRule = append(assignmentIdRule, assignmentIdItem)
	}

	logs, sub, err := _PromptScheduler.contract.FilterLogs(opts, "StreamedData", assignmentIdRule)
	if err != nil {
		return nil, err
	}
	return &PromptSchedulerStreamedDataIterator{contract: _PromptScheduler.contract, event: "StreamedData", logs: logs, sub: sub}, nil
}

// WatchStreamedData is a free log subscription operation binding the contract event 0x23cfaa418b5f569ff36b152a9fd02ee3ccddaa5f7eed570e777a30353b68dc38.
//
// Solidity: event StreamedData(uint256 indexed assignmentId, bytes data)
func (_PromptScheduler *PromptSchedulerFilterer) WatchStreamedData(opts *bind.WatchOpts, sink chan<- *PromptSchedulerStreamedData, assignmentId []*big.Int) (event.Subscription, error) {

	var assignmentIdRule []interface{}
	for _, assignmentIdItem := range assignmentId {
		assignmentIdRule = append(assignmentIdRule, assignmentIdItem)
	}

	logs, sub, err := _PromptScheduler.contract.WatchLogs(opts, "StreamedData", assignmentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptSchedulerStreamedData)
				if err := _PromptScheduler.contract.UnpackLog(event, "StreamedData", log); err != nil {
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
func (_PromptScheduler *PromptSchedulerFilterer) ParseStreamedData(log types.Log) (*PromptSchedulerStreamedData, error) {
	event := new(PromptSchedulerStreamedData)
	if err := _PromptScheduler.contract.UnpackLog(event, "StreamedData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PromptSchedulerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PromptScheduler contract.
type PromptSchedulerUnpausedIterator struct {
	Event *PromptSchedulerUnpaused // Event containing the contract specifics and raw log

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
func (it *PromptSchedulerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PromptSchedulerUnpaused)
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
		it.Event = new(PromptSchedulerUnpaused)
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
func (it *PromptSchedulerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PromptSchedulerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PromptSchedulerUnpaused represents a Unpaused event raised by the PromptScheduler contract.
type PromptSchedulerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PromptScheduler *PromptSchedulerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PromptSchedulerUnpausedIterator, error) {

	logs, sub, err := _PromptScheduler.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PromptSchedulerUnpausedIterator{contract: _PromptScheduler.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PromptScheduler *PromptSchedulerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PromptSchedulerUnpaused) (event.Subscription, error) {

	logs, sub, err := _PromptScheduler.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PromptSchedulerUnpaused)
				if err := _PromptScheduler.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_PromptScheduler *PromptSchedulerFilterer) ParseUnpaused(log types.Log) (*PromptSchedulerUnpaused, error) {
	event := new(PromptSchedulerUnpaused)
	if err := _PromptScheduler.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
