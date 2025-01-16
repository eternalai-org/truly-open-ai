// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package daotreasury

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

// DAOTreasuryMetaData contains all meta data concerning the DAOTreasury contract.
var DAOTreasuryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"R\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"T\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_daoToken\",\"type\":\"address\"}],\"name\":\"addLiqidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseFundBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelFund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_positionManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_baseToken\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolV3\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"position0TokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"position1TokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"positionManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_settledFundBalance\",\"type\":\"uint256\"}],\"name\":\"settleFund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settledFundBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080806040523461001657611f19908161001c8239f35b600080fdfe6080604052600436101561001b575b361561001957600080fd5b005b60003560e01c80630c340a241461015b5780630c9827df1461015657806314a8fc9714610151578063150b7a021461014c5780633408e47014610147578063485cc9551461014257806348ea59a81461013d5780634914b0301461013857806360b71d4e14610133578063715018a61461012e578063791b98bc146101295780638da5cb5b14610124578063ac9650d81461011f578063bb5df5541461011a578063c55dae6314610115578063c7d5120214610110578063e914fa121461010b578063f2fde38b14610106578063f68391d6146101015763fd90f6d40361000e57610d61565b610d43565b610cb2565b610c0c565b610812565b6107e9565b6107cb565b6106c4565b6105f1565b6105c8565b610567565b610549565b610520565b610502565b610409565b6103ee565b610357565b6101a7565b610189565b346101845760003660031901126101845760c9546040516001600160a01b039091168152602090f35b600080fd5b3461018457600036600319011261018457602060cd54604051908152f35b34610184576020366003190112610184576004356101c3610d8a565b6101cf60d254156112a6565b60ce546040516370a0823160e01b8152306004820152916001600160a01b0390911690602083602481855afa9081156102ab576100199360009261026f575b5061024e61023482610225610269948610156112e7565b61024061023b61023487611326565b6064900490565b60d255565b6102498160d055565b61133c565b6102578160d155565b6033546001600160a01b031692611319565b91611c4a565b61026991925061023461029b61024e9260203d81116102a4575b61029381836102f8565b8101906112d8565b9392505061020e565b503d610289565b611055565b6001600160a01b0381160361018457565b634e487b7160e01b600052604160045260246000fd5b60a0810190811067ffffffffffffffff8211176102f357604052565b6102c1565b90601f8019910116810190811067ffffffffffffffff8211176102f357604052565b60405190610160820182811067ffffffffffffffff8211176102f357604052565b67ffffffffffffffff81116102f357601f01601f191660200190565b34610184576080366003190112610184576103736004356102b0565b61037e6024356102b0565b60643567ffffffffffffffff81116101845736602382011215610184578060040135906103aa8261033b565b916103b860405193846102f8565b8083523660248284010111610184576000928160246020940184830137010152604051630a85bd0160e11b8152602090f35b0390f35b34610184576000366003190112610184576020604051468152f35b3461018457604036600319011261018457600435610426816102b0565b610476602435610435816102b0565b6000549261045a60ff8560081c1615809581966104f4575b81156104d4575b50610fd9565b8361046d600160ff196000541617600055565b6104bb576110e0565b61047c57005b61048c61ff001960005416600055565b604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602090a1005b6104cf61010061ff00196000541617600055565b6110e0565b303b159150816104e6575b5038610454565b6001915060ff1614386104df565b600160ff821610915061044d565b3461018457600036600319011261018457602060d154604051908152f35b346101845760003660031901126101845760cf546040516001600160a01b039091168152602090f35b3461018457600036600319011261018457602060d254604051908152f35b34610184576000806003193601126105c557610581610d8a565b603380546001600160a01b0319811690915581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b80fd5b346101845760003660031901126101845760ca546040516001600160a01b039091168152602090f35b34610184576000366003190112610184576033546040516001600160a01b039091168152602090f35b60005b83811061062d5750506000910152565b818101518382015260200161061d565b906020916106568151809281855285808601910161061a565b601f01601f1916010190565b602080820190808352835180925260408301928160408460051b8301019501936000915b8483106106965750505050505090565b90919293949584806106b4600193603f198682030187528a5161063d565b9801930193019194939290610686565b6020366003190112610184576004803567ffffffffffffffff918282116101845736602383011215610184578181013592831161018457602490818301928236918660051b0101116101845761071984610e43565b9360005b81811061073257604051806103ea8882610662565b600080610740838589610ecd565b60409391610752855180938193610f14565b0390305af490610760610f22565b9182901561078f5750509061078a916107798289610fc5565b526107848188610fc5565b50610ea3565b61071d565b86838792604482511061018457826107c793856107b29401518301019101610f52565b925162461bcd60e51b81529283928301610fb4565b0390fd5b3461018457600036600319011261018457602060d054604051908152f35b346101845760003660031901126101845760ce546040516001600160a01b039091168152602090f35b346101845760208060031936011261018457600435610830816102b0565b610838610d8a565b60d254151580610c01575b61084c906113d5565b60cf546001600160a01b03929061086590841615611407565b60ce546001600160a01b031660cf80546001600160a01b0319166001600160a01b03851617905591610898303383611b8b565b61096d60d054936108ad61023460d154611352565b9481871684881610918215610be2578496866109046108ff6108f58c6108ef6108ea6108d9899a611368565b6a52b7d2dcc80cd2e4000000900490565b611d5f565b166118bf565b603c9060020b0590565b611439565b998a9561093061092461091f6109198a61144e565b99611469565b61154f565b6001600160a01b031690565b9015610bc8576109726109b594956a52b7d2dcc80cd2e4000000969a5b60ca5461096490610924906001600160a01b031681565b9a8b8093611ca7565b611ca7565b6040516309f56ab160e11b81526001600160a01b03808d1660048301528088166024830152610bb860448301529190921616606482015291829081906084820190565b03816000895af180156102ab576109ec91600091610b9b575b5060018060a01b03166001600160601b0360a01b60cb54161760cb55565b6109f461031a565b6001600160a01b0388168152946001600160a01b03831686880152610bb86040870152620d89b3196060870152610a326080998a88019060020b9052565b600060a087015260c0860152600060e086015261010060008187015261012091610a66308489019060018060a01b03169052565b600019946101409486868a0152604051998c8b80610a92634418b22b60e11b9d8e8352600483016114b6565b038160008d5af19a8b156102ab578d9b610ab491600091610b7b575b5060cc55565b610abc61031a565b6001600160a01b03909c168c526001600160a01b03909116908b0152610bb860408b015260020b60608a0152620d89b48989015260a0890152600060c0890181905260e08901819052908801523090870152850152604051808095819482526004820190610b29916114b6565b03915a90600091f19081156102ab5761001992600092610b4b575b505060cd55565b610b6a9250803d10610b74575b610b6281836102f8565b810190611484565b5050503880610b44565b503d610b58565b610b9291508d803d10610b7457610b6281836102f8565b50505038610aae565b610bbb9150883d8a11610bc1575b610bb381836102f8565b810190611096565b386109ce565b503d610ba9565b6109726109b594956a52b7d2dcc80cd2e40000009a61094d565b8096866109046108ff6108f58c610bfc6108ea8c996113ad565b6108ef565b5060d1541515610843565b34610184576000806003193601126105c557610c26610d8a565b610c3260d254156112a6565b60ce546040516370a0823160e01b8152306004820152906001600160a01b0316602082602481845afa9081156102ab57610c8f928492610c92575b50610c7d61023b61023484611326565b6033546001600160a01b031690611c4a565b80f35b610cab91925060203d81116102a45761029381836102f8565b9038610c6d565b3461018457602036600319011261018457600435610ccf816102b0565b610cd7610d8a565b6001600160a01b03811615610cef5761001990610de2565b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608490fd5b3461018457600036600319011261018457602060cc54604051908152f35b346101845760003660031901126101845760cb546040516001600160a01b039091168152602090f35b6033546001600160a01b03163303610d9e57565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b603380546001600160a01b039283166001600160a01b0319821681179092559091167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b67ffffffffffffffff81116102f35760051b60200190565b90610e4d82610e2b565b610e5a60405191826102f8565b8281528092610e6b601f1991610e2b565b019060005b828110610e7c57505050565b806060602080938501015201610e70565b634e487b7160e01b600052601160045260246000fd5b6000198114610eb25760010190565b610e8d565b634e487b7160e01b600052603260045260246000fd5b9190811015610f0f5760051b81013590601e198136030182121561018457019081359167ffffffffffffffff8311610184576020018236038113610184579190565b610eb7565b908092918237016000815290565b3d15610f4d573d90610f338261033b565b91610f4160405193846102f8565b82523d6000602084013e565b606090565b6020818303126101845780519067ffffffffffffffff8211610184570181601f82011215610184578051610f858161033b565b92610f9360405194856102f8565b8184526020828401011161018457610fb1916020808501910161061a565b90565b906020610fb192818152019061063d565b8051821015610f0f5760209160051b010190565b15610fe057565b60405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b90816020910312610184575160ff811681036101845790565b6040513d6000823e3d90fd5b1561106857565b60405162461bcd60e51b8152602060048201526006602482015265084a8889c62760d31b6044820152606490fd5b908160209103126101845751610fb1816102b0565b156110b257565b60405162461bcd60e51b81526020600482015260066024820152652826aba2272d60d11b6044820152606490fd5b906110e9611267565b6110f1611256565b60405163313ce56760e01b81526001600160a01b039260209182816004818789165afa80156102ab5760ff601291611133936000916111c9575b501614611061565b60405163c45a015560e01b81529180836004818589165afa9485156102ab576111aa9561118d94611170936000926111ac575b50501615156110ab565b60018060a01b03166001600160601b0360a01b60ca54161760ca55565b60018060a01b03166001600160601b0360a01b60ce54161760ce55565b565b6111c29250803d10610bc157610bb381836102f8565b3880611166565b6111e99150863d88116111ef575b6111e181836102f8565b81019061103c565b3861112b565b503d6111d7565b156111fd57565b60405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608490fd5b6111aa60ff60005460081c166111f6565b61128160ff60005460081c1661127c816111f6565b6111f6565b61128a33610de2565b61129f60ff60005460081c1661127c816111f6565b6001606555565b156112ad57565b606460405162461bcd60e51b81526020600482015260046024820152632321272d60e11b6044820152fd5b90816020910312610184575190565b156112ee57565b60405162461bcd60e51b81526020600482015260036024820152622722a360e91b6044820152606490fd5b91908203918211610eb257565b90600582029180830460051490151715610eb257565b90605f820291808304605f1490151715610eb257565b90600a820291808304600a1490151715610eb257565b90670de0b6b3a764000091828102928184041490151715610eb257565b634e487b7160e01b600052601260045260246000fd5b80156113a8576000190490565b611385565b80156113a85772047bf19673df52e37f2410011d1000000000000490565b81156113a8570490565b156113dc57565b606460405162461bcd60e51b8152602060048201526004602482015263232124ad60e11b6044820152fd5b1561140e57565b606460405162461bcd60e51b8152602060048201526004602482015263222a272d60e11b6044820152fd5b603c9060020b02908160020b918203610eb257565b60020b603c0190627fffff8213627fffff19831217610eb257565b60020b601e0190627fffff8213627fffff19831217610eb257565b91908260809103126101845781519160208101516001600160801b038116810361018457916060604083015192015190565b81516001600160a01b03168152610160810192916020818101516001600160a01b03169083015260408181015162ffffff169083015260608181015160020b9083015260808181015160020b9083015260a0818101519083015260c0808201519083015260e080820151908301526101008082015190830152610120808201516001600160a01b03169083015261014080910151910152565b60020b60008112156118b95780600003905b620d89e882116118a7576001821615611895576001600160881b036ffffcb933bd6fad37aa2d162d1a5940015b169160028116611879575b6004811661185d575b60088116611841575b60108116611825575b60208116611809575b604081166117ed575b6080908181166117d2575b61010081166117b7575b610200811661179c575b6104008116611781575b6108008116611766575b611000811661174b575b6120008116611730575b6140008116611715575b61800081166116fa575b6201000081166116df575b6202000081166116c5575b6204000081166116ab575b6208000016611690575b50600012611682575b63ffffffff811661167a576000905b60201c60ff91909116016001600160a01b031690565b600190611664565b61168b9061139b565b611655565b6b048a170391f7dc42444e8fa26000929302901c919061164c565b6d2216e584f5fa1ea926041bedfe98909302811c92611642565b926e5d6af8dedb81196699c329225ee60402811c92611637565b926f09aa508b5b7a84e1c677de54f3e99bc902811c9261162c565b926f31be135f97d08fd981231505542fcfa602811c92611621565b926f70d869a156d2a1b890bb3df62baf32f702811c92611617565b926fa9f746462d870fdf8a65dc1f90e061e502811c9261160d565b926fd097f3bdfd2022b8845ad8f792aa582502811c92611603565b926fe7159475a2c29b7443b29c7fa6e889d902811c926115f9565b926ff3392b0822b70005940c7a398e4b70f302811c926115ef565b926ff987a7253ac413176f2b074cf7815e5402811c926115e5565b926ffcbe86c7900a88aedcffc83b479aa3a402811c926115db565b926ffe5dee046a99a2a811c461f1969c305302811c926115d1565b916fff2ea16466c96a3843ec78b326b528610260801c916115c6565b916fff973b41fa98c081472e6896dfb254c00260801c916115bd565b916fffcb9843d60f6159c9db58835c9266440260801c916115b4565b916fffe5caca7e10e4e61c3624eaa0941cd00260801c916115ab565b916ffff2e50f5f656932ef12357cf3c7fdcc0260801c916115a2565b916ffff97272373d413259a46990580e213a0260801c91611599565b6001600160881b03600160801b61158e565b6040516315e4079d60e11b8152600490fd5b80611561565b6001600160a01b038116906401000276a382101580611b56575b15611b4457640100000000600160c01b039060201b16806001600160801b03811160071b9181831c9267ffffffffffffffff841160061b93841c9363ffffffff851160051b94851c9461ffff861160041b95861c60ff9687821160031b91821c92600f841160021b93841c94600160038711811b96871c119617171717171717916080831015600014611b385750607e1982011c5b8002607f928392828493841c81841c1c800280851c81851c1c800280861c81861c1c800280871c81871c1c80029081881c82881c1c80029283891c84891c1c800294858a1c868a1c1c800296878b1c888b1c1c800298898c1c8a8c1c1c80029a8b8d1c8c821c1c8002809d1c8d821c1c8002809e81901c90821c1c80029e8f80911c911c1c800260cd1c6604000000000000169d60cc1c6608000000000000169c60cb1c6610000000000000169b60ca1c6620000000000000169a60c91c6640000000000000169960c81c6680000000000000169860c71c670100000000000000169760c61c670200000000000000169660c51c670400000000000000169560c41c670800000000000000169460c31c671000000000000000169360c21c672000000000000000169260c11c674000000000000000169160c01c6780000000000000001690607f190160401b1717171717171717171717171717693627a301d71055774c85026f028f6481ab7f045a5af012a19d003aa919810160801d60020b906fdb2df09e81959a81455e260799a0632f0160801d60020b91600090838314600014611b1c575050905090565b611b286109248561154f565b119050611b33575090565b905090565b905081607f031b61196e565b6040516324c070df60e11b8152600490fd5b5073fffd8963efd1fc6a506488495d951d5263988d2682106118d9565b90816020910312610184575180151581036101845790565b6000928380936040519060208201936323b872dd60e01b855260018060a01b0380921660248401521660448201526aa56fa5b99019a5c8000000606482015260648152611bd7816102d7565b51925af1611be3610f22565b81611c1b575b5015611bf157565b60405162461bcd60e51b81526020600482015260026024820152612a2360f11b6044820152606490fd5b8051801592508215611c30575b505038611be9565b611c439250602080918301019101611b73565b3880611c28565b60405163a9059cbb60e01b602082019081526001600160a01b03909316602482015260448082019490945292835260808301929167ffffffffffffffff8411838510176102f3576000809493819460405251925af1611be3610f22565b60405163095ea7b360e01b602082019081526001600160a01b0393909316602482015260001960448083019190915281526000928392918390611ceb6064826102f8565b51925af1611cf7610f22565b81611d30575b5015611d0557565b60405162461bcd60e51b815260206004820152600360248201526229a0a360e91b6044820152606490fd5b8051801592508215611d45575b505038611cfd565b611d589250602080918301019101611b73565b3880611d3d565b8060601b600160601b91808204831490151715610eb257670de0b6b3a764000090048060601b918183041490151715610eb257610fb1908015611ed05780611e69611e62611e58611e4e611e44611e3a611e30611e266001610fb19a6000908b60801c80611ec4575b508060401c80611eb7575b508060201c80611eaa575b508060101c80611e9d575b508060081c80611e90575b508060041c80611e83575b508060021c80611e76575b50821c611e6f575b811c1b611e1f818b6113cb565b0160011c90565b611e1f818a6113cb565b611e1f81896113cb565b611e1f81886113cb565b611e1f81876113cb565b611e1f81866113cb565b611e1f81856113cb565b80926113cb565b90611ed6565b8101611e12565b6002915091019038611e0a565b6004915091019038611dff565b6008915091019038611df4565b6010915091019038611de9565b6020915091019038611dde565b6040915091019038611dd3565b91505060809038611dc8565b50600090565b9080821015611b3357509056fea2646970667358221220a3ccb794dbb7372ce03fd2e07ea573c86254ce76683263bae6258d366b5a806864736f6c63430008130033",
}

// DAOTreasuryABI is the input ABI used to generate the binding from.
// Deprecated: Use DAOTreasuryMetaData.ABI instead.
var DAOTreasuryABI = DAOTreasuryMetaData.ABI

// DAOTreasuryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DAOTreasuryMetaData.Bin instead.
var DAOTreasuryBin = DAOTreasuryMetaData.Bin

// DeployDAOTreasury deploys a new Ethereum contract, binding an instance of DAOTreasury to it.
func DeployDAOTreasury(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DAOTreasury, error) {
	parsed, err := DAOTreasuryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DAOTreasuryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DAOTreasury{DAOTreasuryCaller: DAOTreasuryCaller{contract: contract}, DAOTreasuryTransactor: DAOTreasuryTransactor{contract: contract}, DAOTreasuryFilterer: DAOTreasuryFilterer{contract: contract}}, nil
}

// DAOTreasury is an auto generated Go binding around an Ethereum contract.
type DAOTreasury struct {
	DAOTreasuryCaller     // Read-only binding to the contract
	DAOTreasuryTransactor // Write-only binding to the contract
	DAOTreasuryFilterer   // Log filterer for contract events
}

// DAOTreasuryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DAOTreasuryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAOTreasuryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DAOTreasuryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAOTreasuryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DAOTreasuryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAOTreasurySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DAOTreasurySession struct {
	Contract     *DAOTreasury      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DAOTreasuryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DAOTreasuryCallerSession struct {
	Contract *DAOTreasuryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DAOTreasuryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DAOTreasuryTransactorSession struct {
	Contract     *DAOTreasuryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DAOTreasuryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DAOTreasuryRaw struct {
	Contract *DAOTreasury // Generic contract binding to access the raw methods on
}

// DAOTreasuryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DAOTreasuryCallerRaw struct {
	Contract *DAOTreasuryCaller // Generic read-only contract binding to access the raw methods on
}

// DAOTreasuryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DAOTreasuryTransactorRaw struct {
	Contract *DAOTreasuryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDAOTreasury creates a new instance of DAOTreasury, bound to a specific deployed contract.
func NewDAOTreasury(address common.Address, backend bind.ContractBackend) (*DAOTreasury, error) {
	contract, err := bindDAOTreasury(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DAOTreasury{DAOTreasuryCaller: DAOTreasuryCaller{contract: contract}, DAOTreasuryTransactor: DAOTreasuryTransactor{contract: contract}, DAOTreasuryFilterer: DAOTreasuryFilterer{contract: contract}}, nil
}

// NewDAOTreasuryCaller creates a new read-only instance of DAOTreasury, bound to a specific deployed contract.
func NewDAOTreasuryCaller(address common.Address, caller bind.ContractCaller) (*DAOTreasuryCaller, error) {
	contract, err := bindDAOTreasury(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DAOTreasuryCaller{contract: contract}, nil
}

// NewDAOTreasuryTransactor creates a new write-only instance of DAOTreasury, bound to a specific deployed contract.
func NewDAOTreasuryTransactor(address common.Address, transactor bind.ContractTransactor) (*DAOTreasuryTransactor, error) {
	contract, err := bindDAOTreasury(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DAOTreasuryTransactor{contract: contract}, nil
}

// NewDAOTreasuryFilterer creates a new log filterer instance of DAOTreasury, bound to a specific deployed contract.
func NewDAOTreasuryFilterer(address common.Address, filterer bind.ContractFilterer) (*DAOTreasuryFilterer, error) {
	contract, err := bindDAOTreasury(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DAOTreasuryFilterer{contract: contract}, nil
}

// bindDAOTreasury binds a generic wrapper to an already deployed contract.
func bindDAOTreasury(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DAOTreasuryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DAOTreasury *DAOTreasuryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DAOTreasury.Contract.DAOTreasuryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DAOTreasury *DAOTreasuryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAOTreasury.Contract.DAOTreasuryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DAOTreasury *DAOTreasuryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DAOTreasury.Contract.DAOTreasuryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DAOTreasury *DAOTreasuryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DAOTreasury.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DAOTreasury *DAOTreasuryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAOTreasury.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DAOTreasury *DAOTreasuryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DAOTreasury.Contract.contract.Transact(opts, method, params...)
}

// BaseFundBalance is a free data retrieval call binding the contract method 0x48ea59a8.
//
// Solidity: function baseFundBalance() view returns(uint256)
func (_DAOTreasury *DAOTreasuryCaller) BaseFundBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAOTreasury.contract.Call(opts, &out, "baseFundBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseFundBalance is a free data retrieval call binding the contract method 0x48ea59a8.
//
// Solidity: function baseFundBalance() view returns(uint256)
func (_DAOTreasury *DAOTreasurySession) BaseFundBalance() (*big.Int, error) {
	return _DAOTreasury.Contract.BaseFundBalance(&_DAOTreasury.CallOpts)
}

// BaseFundBalance is a free data retrieval call binding the contract method 0x48ea59a8.
//
// Solidity: function baseFundBalance() view returns(uint256)
func (_DAOTreasury *DAOTreasuryCallerSession) BaseFundBalance() (*big.Int, error) {
	return _DAOTreasury.Contract.BaseFundBalance(&_DAOTreasury.CallOpts)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_DAOTreasury *DAOTreasuryCaller) BaseToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DAOTreasury.contract.Call(opts, &out, "baseToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_DAOTreasury *DAOTreasurySession) BaseToken() (common.Address, error) {
	return _DAOTreasury.Contract.BaseToken(&_DAOTreasury.CallOpts)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_DAOTreasury *DAOTreasuryCallerSession) BaseToken() (common.Address, error) {
	return _DAOTreasury.Contract.BaseToken(&_DAOTreasury.CallOpts)
}

// DaoToken is a free data retrieval call binding the contract method 0x4914b030.
//
// Solidity: function daoToken() view returns(address)
func (_DAOTreasury *DAOTreasuryCaller) DaoToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DAOTreasury.contract.Call(opts, &out, "daoToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaoToken is a free data retrieval call binding the contract method 0x4914b030.
//
// Solidity: function daoToken() view returns(address)
func (_DAOTreasury *DAOTreasurySession) DaoToken() (common.Address, error) {
	return _DAOTreasury.Contract.DaoToken(&_DAOTreasury.CallOpts)
}

// DaoToken is a free data retrieval call binding the contract method 0x4914b030.
//
// Solidity: function daoToken() view returns(address)
func (_DAOTreasury *DAOTreasuryCallerSession) DaoToken() (common.Address, error) {
	return _DAOTreasury.Contract.DaoToken(&_DAOTreasury.CallOpts)
}

// FeeBalance is a free data retrieval call binding the contract method 0x60b71d4e.
//
// Solidity: function feeBalance() view returns(uint256)
func (_DAOTreasury *DAOTreasuryCaller) FeeBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAOTreasury.contract.Call(opts, &out, "feeBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeBalance is a free data retrieval call binding the contract method 0x60b71d4e.
//
// Solidity: function feeBalance() view returns(uint256)
func (_DAOTreasury *DAOTreasurySession) FeeBalance() (*big.Int, error) {
	return _DAOTreasury.Contract.FeeBalance(&_DAOTreasury.CallOpts)
}

// FeeBalance is a free data retrieval call binding the contract method 0x60b71d4e.
//
// Solidity: function feeBalance() view returns(uint256)
func (_DAOTreasury *DAOTreasuryCallerSession) FeeBalance() (*big.Int, error) {
	return _DAOTreasury.Contract.FeeBalance(&_DAOTreasury.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainId)
func (_DAOTreasury *DAOTreasuryCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAOTreasury.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainId)
func (_DAOTreasury *DAOTreasurySession) GetChainId() (*big.Int, error) {
	return _DAOTreasury.Contract.GetChainId(&_DAOTreasury.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainId)
func (_DAOTreasury *DAOTreasuryCallerSession) GetChainId() (*big.Int, error) {
	return _DAOTreasury.Contract.GetChainId(&_DAOTreasury.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_DAOTreasury *DAOTreasuryCaller) Governor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DAOTreasury.contract.Call(opts, &out, "governor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_DAOTreasury *DAOTreasurySession) Governor() (common.Address, error) {
	return _DAOTreasury.Contract.Governor(&_DAOTreasury.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_DAOTreasury *DAOTreasuryCallerSession) Governor() (common.Address, error) {
	return _DAOTreasury.Contract.Governor(&_DAOTreasury.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DAOTreasury *DAOTreasuryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DAOTreasury.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DAOTreasury *DAOTreasurySession) Owner() (common.Address, error) {
	return _DAOTreasury.Contract.Owner(&_DAOTreasury.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DAOTreasury *DAOTreasuryCallerSession) Owner() (common.Address, error) {
	return _DAOTreasury.Contract.Owner(&_DAOTreasury.CallOpts)
}

// PoolV3 is a free data retrieval call binding the contract method 0xfd90f6d4.
//
// Solidity: function poolV3() view returns(address)
func (_DAOTreasury *DAOTreasuryCaller) PoolV3(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DAOTreasury.contract.Call(opts, &out, "poolV3")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolV3 is a free data retrieval call binding the contract method 0xfd90f6d4.
//
// Solidity: function poolV3() view returns(address)
func (_DAOTreasury *DAOTreasurySession) PoolV3() (common.Address, error) {
	return _DAOTreasury.Contract.PoolV3(&_DAOTreasury.CallOpts)
}

// PoolV3 is a free data retrieval call binding the contract method 0xfd90f6d4.
//
// Solidity: function poolV3() view returns(address)
func (_DAOTreasury *DAOTreasuryCallerSession) PoolV3() (common.Address, error) {
	return _DAOTreasury.Contract.PoolV3(&_DAOTreasury.CallOpts)
}

// Position0TokenId is a free data retrieval call binding the contract method 0xf68391d6.
//
// Solidity: function position0TokenId() view returns(uint256)
func (_DAOTreasury *DAOTreasuryCaller) Position0TokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAOTreasury.contract.Call(opts, &out, "position0TokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Position0TokenId is a free data retrieval call binding the contract method 0xf68391d6.
//
// Solidity: function position0TokenId() view returns(uint256)
func (_DAOTreasury *DAOTreasurySession) Position0TokenId() (*big.Int, error) {
	return _DAOTreasury.Contract.Position0TokenId(&_DAOTreasury.CallOpts)
}

// Position0TokenId is a free data retrieval call binding the contract method 0xf68391d6.
//
// Solidity: function position0TokenId() view returns(uint256)
func (_DAOTreasury *DAOTreasuryCallerSession) Position0TokenId() (*big.Int, error) {
	return _DAOTreasury.Contract.Position0TokenId(&_DAOTreasury.CallOpts)
}

// Position1TokenId is a free data retrieval call binding the contract method 0x0c9827df.
//
// Solidity: function position1TokenId() view returns(uint256)
func (_DAOTreasury *DAOTreasuryCaller) Position1TokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAOTreasury.contract.Call(opts, &out, "position1TokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Position1TokenId is a free data retrieval call binding the contract method 0x0c9827df.
//
// Solidity: function position1TokenId() view returns(uint256)
func (_DAOTreasury *DAOTreasurySession) Position1TokenId() (*big.Int, error) {
	return _DAOTreasury.Contract.Position1TokenId(&_DAOTreasury.CallOpts)
}

// Position1TokenId is a free data retrieval call binding the contract method 0x0c9827df.
//
// Solidity: function position1TokenId() view returns(uint256)
func (_DAOTreasury *DAOTreasuryCallerSession) Position1TokenId() (*big.Int, error) {
	return _DAOTreasury.Contract.Position1TokenId(&_DAOTreasury.CallOpts)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_DAOTreasury *DAOTreasuryCaller) PositionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DAOTreasury.contract.Call(opts, &out, "positionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_DAOTreasury *DAOTreasurySession) PositionManager() (common.Address, error) {
	return _DAOTreasury.Contract.PositionManager(&_DAOTreasury.CallOpts)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_DAOTreasury *DAOTreasuryCallerSession) PositionManager() (common.Address, error) {
	return _DAOTreasury.Contract.PositionManager(&_DAOTreasury.CallOpts)
}

// SettledFundBalance is a free data retrieval call binding the contract method 0xbb5df554.
//
// Solidity: function settledFundBalance() view returns(uint256)
func (_DAOTreasury *DAOTreasuryCaller) SettledFundBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAOTreasury.contract.Call(opts, &out, "settledFundBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SettledFundBalance is a free data retrieval call binding the contract method 0xbb5df554.
//
// Solidity: function settledFundBalance() view returns(uint256)
func (_DAOTreasury *DAOTreasurySession) SettledFundBalance() (*big.Int, error) {
	return _DAOTreasury.Contract.SettledFundBalance(&_DAOTreasury.CallOpts)
}

// SettledFundBalance is a free data retrieval call binding the contract method 0xbb5df554.
//
// Solidity: function settledFundBalance() view returns(uint256)
func (_DAOTreasury *DAOTreasuryCallerSession) SettledFundBalance() (*big.Int, error) {
	return _DAOTreasury.Contract.SettledFundBalance(&_DAOTreasury.CallOpts)
}

// AddLiqidity is a paid mutator transaction binding the contract method 0xc7d51202.
//
// Solidity: function addLiqidity(address _daoToken) returns()
func (_DAOTreasury *DAOTreasuryTransactor) AddLiqidity(opts *bind.TransactOpts, _daoToken common.Address) (*types.Transaction, error) {
	return _DAOTreasury.contract.Transact(opts, "addLiqidity", _daoToken)
}

// AddLiqidity is a paid mutator transaction binding the contract method 0xc7d51202.
//
// Solidity: function addLiqidity(address _daoToken) returns()
func (_DAOTreasury *DAOTreasurySession) AddLiqidity(_daoToken common.Address) (*types.Transaction, error) {
	return _DAOTreasury.Contract.AddLiqidity(&_DAOTreasury.TransactOpts, _daoToken)
}

// AddLiqidity is a paid mutator transaction binding the contract method 0xc7d51202.
//
// Solidity: function addLiqidity(address _daoToken) returns()
func (_DAOTreasury *DAOTreasuryTransactorSession) AddLiqidity(_daoToken common.Address) (*types.Transaction, error) {
	return _DAOTreasury.Contract.AddLiqidity(&_DAOTreasury.TransactOpts, _daoToken)
}

// CancelFund is a paid mutator transaction binding the contract method 0xe914fa12.
//
// Solidity: function cancelFund() returns()
func (_DAOTreasury *DAOTreasuryTransactor) CancelFund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAOTreasury.contract.Transact(opts, "cancelFund")
}

// CancelFund is a paid mutator transaction binding the contract method 0xe914fa12.
//
// Solidity: function cancelFund() returns()
func (_DAOTreasury *DAOTreasurySession) CancelFund() (*types.Transaction, error) {
	return _DAOTreasury.Contract.CancelFund(&_DAOTreasury.TransactOpts)
}

// CancelFund is a paid mutator transaction binding the contract method 0xe914fa12.
//
// Solidity: function cancelFund() returns()
func (_DAOTreasury *DAOTreasuryTransactorSession) CancelFund() (*types.Transaction, error) {
	return _DAOTreasury.Contract.CancelFund(&_DAOTreasury.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _positionManager, address _baseToken) returns()
func (_DAOTreasury *DAOTreasuryTransactor) Initialize(opts *bind.TransactOpts, _positionManager common.Address, _baseToken common.Address) (*types.Transaction, error) {
	return _DAOTreasury.contract.Transact(opts, "initialize", _positionManager, _baseToken)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _positionManager, address _baseToken) returns()
func (_DAOTreasury *DAOTreasurySession) Initialize(_positionManager common.Address, _baseToken common.Address) (*types.Transaction, error) {
	return _DAOTreasury.Contract.Initialize(&_DAOTreasury.TransactOpts, _positionManager, _baseToken)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _positionManager, address _baseToken) returns()
func (_DAOTreasury *DAOTreasuryTransactorSession) Initialize(_positionManager common.Address, _baseToken common.Address) (*types.Transaction, error) {
	return _DAOTreasury.Contract.Initialize(&_DAOTreasury.TransactOpts, _positionManager, _baseToken)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_DAOTreasury *DAOTreasuryTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _DAOTreasury.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_DAOTreasury *DAOTreasurySession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _DAOTreasury.Contract.Multicall(&_DAOTreasury.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_DAOTreasury *DAOTreasuryTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _DAOTreasury.Contract.Multicall(&_DAOTreasury.TransactOpts, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_DAOTreasury *DAOTreasuryTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _DAOTreasury.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_DAOTreasury *DAOTreasurySession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _DAOTreasury.Contract.OnERC721Received(&_DAOTreasury.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_DAOTreasury *DAOTreasuryTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _DAOTreasury.Contract.OnERC721Received(&_DAOTreasury.TransactOpts, arg0, arg1, arg2, arg3)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DAOTreasury *DAOTreasuryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAOTreasury.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DAOTreasury *DAOTreasurySession) RenounceOwnership() (*types.Transaction, error) {
	return _DAOTreasury.Contract.RenounceOwnership(&_DAOTreasury.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DAOTreasury *DAOTreasuryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DAOTreasury.Contract.RenounceOwnership(&_DAOTreasury.TransactOpts)
}

// SettleFund is a paid mutator transaction binding the contract method 0x14a8fc97.
//
// Solidity: function settleFund(uint256 _settledFundBalance) returns()
func (_DAOTreasury *DAOTreasuryTransactor) SettleFund(opts *bind.TransactOpts, _settledFundBalance *big.Int) (*types.Transaction, error) {
	return _DAOTreasury.contract.Transact(opts, "settleFund", _settledFundBalance)
}

// SettleFund is a paid mutator transaction binding the contract method 0x14a8fc97.
//
// Solidity: function settleFund(uint256 _settledFundBalance) returns()
func (_DAOTreasury *DAOTreasurySession) SettleFund(_settledFundBalance *big.Int) (*types.Transaction, error) {
	return _DAOTreasury.Contract.SettleFund(&_DAOTreasury.TransactOpts, _settledFundBalance)
}

// SettleFund is a paid mutator transaction binding the contract method 0x14a8fc97.
//
// Solidity: function settleFund(uint256 _settledFundBalance) returns()
func (_DAOTreasury *DAOTreasuryTransactorSession) SettleFund(_settledFundBalance *big.Int) (*types.Transaction, error) {
	return _DAOTreasury.Contract.SettleFund(&_DAOTreasury.TransactOpts, _settledFundBalance)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DAOTreasury *DAOTreasuryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DAOTreasury.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DAOTreasury *DAOTreasurySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DAOTreasury.Contract.TransferOwnership(&_DAOTreasury.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DAOTreasury *DAOTreasuryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DAOTreasury.Contract.TransferOwnership(&_DAOTreasury.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DAOTreasury *DAOTreasuryTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAOTreasury.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DAOTreasury *DAOTreasurySession) Receive() (*types.Transaction, error) {
	return _DAOTreasury.Contract.Receive(&_DAOTreasury.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DAOTreasury *DAOTreasuryTransactorSession) Receive() (*types.Transaction, error) {
	return _DAOTreasury.Contract.Receive(&_DAOTreasury.TransactOpts)
}

// DAOTreasuryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DAOTreasury contract.
type DAOTreasuryInitializedIterator struct {
	Event *DAOTreasuryInitialized // Event containing the contract specifics and raw log

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
func (it *DAOTreasuryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOTreasuryInitialized)
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
		it.Event = new(DAOTreasuryInitialized)
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
func (it *DAOTreasuryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOTreasuryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOTreasuryInitialized represents a Initialized event raised by the DAOTreasury contract.
type DAOTreasuryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DAOTreasury *DAOTreasuryFilterer) FilterInitialized(opts *bind.FilterOpts) (*DAOTreasuryInitializedIterator, error) {

	logs, sub, err := _DAOTreasury.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DAOTreasuryInitializedIterator{contract: _DAOTreasury.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DAOTreasury *DAOTreasuryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DAOTreasuryInitialized) (event.Subscription, error) {

	logs, sub, err := _DAOTreasury.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOTreasuryInitialized)
				if err := _DAOTreasury.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_DAOTreasury *DAOTreasuryFilterer) ParseInitialized(log types.Log) (*DAOTreasuryInitialized, error) {
	event := new(DAOTreasuryInitialized)
	if err := _DAOTreasury.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOTreasuryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DAOTreasury contract.
type DAOTreasuryOwnershipTransferredIterator struct {
	Event *DAOTreasuryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DAOTreasuryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOTreasuryOwnershipTransferred)
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
		it.Event = new(DAOTreasuryOwnershipTransferred)
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
func (it *DAOTreasuryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOTreasuryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOTreasuryOwnershipTransferred represents a OwnershipTransferred event raised by the DAOTreasury contract.
type DAOTreasuryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DAOTreasury *DAOTreasuryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DAOTreasuryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DAOTreasury.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DAOTreasuryOwnershipTransferredIterator{contract: _DAOTreasury.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DAOTreasury *DAOTreasuryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DAOTreasuryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DAOTreasury.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOTreasuryOwnershipTransferred)
				if err := _DAOTreasury.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DAOTreasury *DAOTreasuryFilterer) ParseOwnershipTransferred(log types.Log) (*DAOTreasuryOwnershipTransferred, error) {
	event := new(DAOTreasuryOwnershipTransferred)
	if err := _DAOTreasury.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
