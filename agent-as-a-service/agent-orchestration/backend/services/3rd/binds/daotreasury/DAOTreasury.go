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
	ABI: "[{\"inputs\":[],\"name\":\"R\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"T\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_daoToken\",\"type\":\"address\"}],\"name\":\"addLiqidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseFundBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelFund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_positionManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_baseToken\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolV3\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"position0TokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"position1TokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"positionManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseFundBalanceBeforeFee\",\"type\":\"uint256\"}],\"name\":\"settleFund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608080604052346100165761212a908161001c8239f35b600080fdfe6080604052600436101561001b575b361561001957600080fd5b005b60003560e01c80630c340a241461014b5780630c9827df1461014657806314a8fc9714610141578063150b7a021461013c5780633408e47014610137578063485cc9551461013257806348ea59a81461012d5780634914b0301461012857806360b71d4e14610123578063715018a61461011e578063791b98bc146101195780638da5cb5b14610114578063ac9650d81461010f578063c55dae631461010a578063c7d5120214610105578063e914fa1214610100578063f2fde38b146100fb578063f68391d6146100f65763fd90f6d40361000e57610d43565b610d25565b610c94565b610bee565b6107db565b6107b2565b6106ab565b6105d8565b6105af565b61054e565b610530565b610507565b6104e9565b6103f0565b6103d5565b61033e565b610197565b610179565b346101745760003660031901126101745760c9546040516001600160a01b039091168152602090f35b600080fd5b3461017457600036600319011261017457602060cd54604051908152f35b34610174576020366003190112610174576004356101b3610d6c565b6101bf60d15415611288565b60ce546040516370a0823160e01b8152306004820152916001600160a01b0390911690602083602481855afa9081156102925761001993600092610256575b5061023561022482610215610250948610156112c9565b61023061022b61022487611321565b6064900490565b60d155565b611337565b61023e8160d055565b6033546001600160a01b031692611314565b91611c2a565b6102509192506102246102826102359260203d811161028b575b61027a81836102df565b8101906112ba565b939250506101fe565b503d610270565b611037565b6001600160a01b0381160361017457565b634e487b7160e01b600052604160045260246000fd5b60a0810190811067ffffffffffffffff8211176102da57604052565b6102a8565b90601f8019910116810190811067ffffffffffffffff8211176102da57604052565b60405190610160820182811067ffffffffffffffff8211176102da57604052565b67ffffffffffffffff81116102da57601f01601f191660200190565b346101745760803660031901126101745761035a600435610297565b610365602435610297565b60643567ffffffffffffffff811161017457366023820112156101745780600401359061039182610322565b9161039f60405193846102df565b8083523660248284010111610174576000928160246020940184830137010152604051630a85bd0160e11b8152602090f35b0390f35b34610174576000366003190112610174576020604051468152f35b346101745760403660031901126101745760043561040d81610297565b61045d60243561041c81610297565b6000549261044160ff8560081c1615809581966104db575b81156104bb575b50610fbb565b83610454600160ff196000541617600055565b6104a2576110c2565b61046357005b61047361ff001960005416600055565b604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602090a1005b6104b661010061ff00196000541617600055565b6110c2565b303b159150816104cd575b503861043b565b6001915060ff1614386104c6565b600160ff8216109150610434565b3461017457600036600319011261017457602060d054604051908152f35b346101745760003660031901126101745760cf546040516001600160a01b039091168152602090f35b3461017457600036600319011261017457602060d154604051908152f35b34610174576000806003193601126105ac57610568610d6c565b603380546001600160a01b0319811690915581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b80fd5b346101745760003660031901126101745760ca546040516001600160a01b039091168152602090f35b34610174576000366003190112610174576033546040516001600160a01b039091168152602090f35b60005b8381106106145750506000910152565b8181015183820152602001610604565b9060209161063d81518092818552858086019101610601565b601f01601f1916010190565b602080820190808352835180925260408301928160408460051b8301019501936000915b84831061067d5750505050505090565b909192939495848061069b600193603f198682030187528a51610624565b980193019301919493929061066d565b6020366003190112610174576004803567ffffffffffffffff918282116101745736602383011215610174578181013592831161017457602490818301928236918660051b0101116101745761070084610e25565b9360005b81811061071957604051806103d18882610649565b600080610727838589610eaf565b60409391610739855180938193610ef6565b0390305af490610747610f04565b9182901561077657505090610771916107608289610fa7565b5261076b8188610fa7565b50610e85565b610704565b86838792604482511061017457826107ae93856107999401518301019101610f34565b925162461bcd60e51b81529283928301610f96565b0390fd5b346101745760003660031901126101745760ce546040516001600160a01b039091168152602090f35b3461017457602080600319360112610174576004356107f981610297565b610801610d6c565b60d154151580610be3575b610815906113d0565b60cf546001600160a01b03929061082e90841615611402565b60ce546001600160a01b031660cf80546001600160a01b0319166001600160a01b03851617905592610861303385611b6b565b61086f61022460d05461134d565b8482168483161015610b845761093994916109349185849182876108ae6108a961089885611363565b6a52b7d2dcc80cd2e4000000900490565b611d3f565b95808716966108f76108dc6108d76108d26108c88c61189f565b603c9060020b0590565b611449565b611434565b986108e7878261206a565b906108f18b61152f565b906120a1565b9c8d959a6109066000976112fb565b975b60ca5461092b9061091f906001600160a01b031681565b6001600160a01b031690565b9b8c8093611c87565b611c87565b6040516309f56ab160e11b81526001600160a01b038d811660048301529485166024820152610bb86044820152911690921660648301528160848160008a5af18015610292576109a991600091610b57575b5060018060a01b03166001600160601b0360a01b60cb54161760cb55565b6109b1610301565b6001600160a01b0389168152956001600160a01b03841687890152610bb86040880152620d89b3196060880152608099858b8901906109f2919060020b9052565b60a088015260c0870152600060e0870181905261010080880191909152306101208089019190915292909190600019956101409587878b01526040519a8d8c80634418b22b60e11b9d8e82526004820190610a4c91611496565b03818d5a90600091f19b8c15610292578e9c610a72918e600092610b37575b505060cc55565b610a7a610301565b6001600160a01b03909d168d526001600160a01b03909116908c0152610bb860408c015260020b60608b0152620d89b48a8a015260a08a015260c0890152600060e08901819052908801523090870152850152604051808095819482526004820190610ae591611496565b03915a90600091f19081156102925761001992600092610b07575b505060cd55565b610b269250803d10610b30575b610b1e81836102df565b810190611464565b5050503880610b00565b503d610b14565b610b4d9250803d10610b3057610b1e81836102df565b505050388e610a6b565b610b779150893d8b11610b7d575b610b6f81836102df565b810190611078565b3861098b565b503d610b65565b849391610934918390610b996108a9826113a8565b90828785841695610bcc610bb26108d76108c88a61189f565b97610bbd8682611edc565b90610bc78a61152f565b6120a1565b939461093960009d610bdd876112fb565b9b610908565b5060d054151561080c565b34610174576000806003193601126105ac57610c08610d6c565b610c1460d15415611288565b60ce546040516370a0823160e01b8152306004820152906001600160a01b0316602082602481845afa90811561029257610c71928492610c74575b50610c5f61022b61022484611321565b6033546001600160a01b031690611c2a565b80f35b610c8d91925060203d811161028b5761027a81836102df565b9038610c4f565b3461017457602036600319011261017457600435610cb181610297565b610cb9610d6c565b6001600160a01b03811615610cd15761001990610dc4565b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608490fd5b3461017457600036600319011261017457602060cc54604051908152f35b346101745760003660031901126101745760cb546040516001600160a01b039091168152602090f35b6033546001600160a01b03163303610d8057565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b603380546001600160a01b039283166001600160a01b0319821681179092559091167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3565b67ffffffffffffffff81116102da5760051b60200190565b90610e2f82610e0d565b610e3c60405191826102df565b8281528092610e4d601f1991610e0d565b019060005b828110610e5e57505050565b806060602080938501015201610e52565b634e487b7160e01b600052601160045260246000fd5b6000198114610e945760010190565b610e6f565b634e487b7160e01b600052603260045260246000fd5b9190811015610ef15760051b81013590601e198136030182121561017457019081359167ffffffffffffffff8311610174576020018236038113610174579190565b610e99565b908092918237016000815290565b3d15610f2f573d90610f1582610322565b91610f2360405193846102df565b82523d6000602084013e565b606090565b6020818303126101745780519067ffffffffffffffff8211610174570181601f82011215610174578051610f6781610322565b92610f7560405194856102df565b8184526020828401011161017457610f939160208085019101610601565b90565b906020610f93928181520190610624565b8051821015610ef15760209160051b010190565b15610fc257565b60405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b90816020910312610174575160ff811681036101745790565b6040513d6000823e3d90fd5b1561104a57565b60405162461bcd60e51b8152602060048201526006602482015265084a8889c62760d31b6044820152606490fd5b908160209103126101745751610f9381610297565b1561109457565b60405162461bcd60e51b81526020600482015260066024820152652826aba2272d60d11b6044820152606490fd5b906110cb611249565b6110d3611238565b60405163313ce56760e01b81526001600160a01b039260209182816004818789165afa80156102925760ff601291611115936000916111ab575b501614611043565b60405163c45a015560e01b81529180836004818589165afa9485156102925761118c9561116f946111529360009261118e575b505016151561108d565b60018060a01b03166001600160601b0360a01b60ca54161760ca55565b60018060a01b03166001600160601b0360a01b60ce54161760ce55565b565b6111a49250803d10610b7d57610b6f81836102df565b3880611148565b6111cb9150863d88116111d1575b6111c381836102df565b81019061101e565b3861110d565b503d6111b9565b156111df57565b60405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608490fd5b61118c60ff60005460081c166111d8565b61126360ff60005460081c1661125e816111d8565b6111d8565b61126c33610dc4565b61128160ff60005460081c1661125e816111d8565b6001606555565b1561128f57565b606460405162461bcd60e51b81526020600482015260046024820152632321272d60e11b6044820152fd5b90816020910312610174575190565b156112d057565b60405162461bcd60e51b81526020600482015260036024820152622722a360e91b6044820152606490fd5b906a52b7d2dcc80cd2e4000000918203918211610e9457565b91908203918211610e9457565b90600582029180830460051490151715610e9457565b90605f820291808304605f1490151715610e9457565b90600a820291808304600a1490151715610e9457565b90670de0b6b3a764000091828102928184041490151715610e9457565b634e487b7160e01b600052601260045260246000fd5b80156113a3576000190490565b611380565b80156113a35772047bf19673df52e37f2410011d1000000000000490565b81156113a3570490565b156113d757565b606460405162461bcd60e51b8152602060048201526004602482015263232124ad60e11b6044820152fd5b1561140957565b606460405162461bcd60e51b8152602060048201526004602482015263222a272d60e11b6044820152fd5b603c9060020b02908160020b918203610e9457565b60020b60010190627fffff8213627fffff19831217610e9457565b91908260809103126101745781519160208101516001600160801b038116810361017457916060604083015192015190565b81516001600160a01b03168152610160810192916020818101516001600160a01b03169083015260408181015162ffffff169083015260608181015160020b9083015260808181015160020b9083015260a0818101519083015260c0808201519083015260e080820151908301526101008082015190830152610120808201516001600160a01b03169083015261014080910151910152565b60020b60008112156118995780600003905b620d89e88211611887576001821615611875576001600160881b036ffffcb933bd6fad37aa2d162d1a5940015b169160028116611859575b6004811661183d575b60088116611821575b60108116611805575b602081166117e9575b604081166117cd575b6080908181166117b2575b6101008116611797575b610200811661177c575b6104008116611761575b6108008116611746575b611000811661172b575b6120008116611710575b61400081166116f5575b61800081166116da575b6201000081166116bf575b6202000081166116a5575b62040000811661168b575b6208000016611670575b50600012611662575b63ffffffff811661165a576000905b60201c60ff91909116016001600160a01b031690565b600190611644565b61166b90611396565b611635565b6b048a170391f7dc42444e8fa26000929302901c919061162c565b6d2216e584f5fa1ea926041bedfe98909302811c92611622565b926e5d6af8dedb81196699c329225ee60402811c92611617565b926f09aa508b5b7a84e1c677de54f3e99bc902811c9261160c565b926f31be135f97d08fd981231505542fcfa602811c92611601565b926f70d869a156d2a1b890bb3df62baf32f702811c926115f7565b926fa9f746462d870fdf8a65dc1f90e061e502811c926115ed565b926fd097f3bdfd2022b8845ad8f792aa582502811c926115e3565b926fe7159475a2c29b7443b29c7fa6e889d902811c926115d9565b926ff3392b0822b70005940c7a398e4b70f302811c926115cf565b926ff987a7253ac413176f2b074cf7815e5402811c926115c5565b926ffcbe86c7900a88aedcffc83b479aa3a402811c926115bb565b926ffe5dee046a99a2a811c461f1969c305302811c926115b1565b916fff2ea16466c96a3843ec78b326b528610260801c916115a6565b916fff973b41fa98c081472e6896dfb254c00260801c9161159d565b916fffcb9843d60f6159c9db58835c9266440260801c91611594565b916fffe5caca7e10e4e61c3624eaa0941cd00260801c9161158b565b916ffff2e50f5f656932ef12357cf3c7fdcc0260801c91611582565b916ffff97272373d413259a46990580e213a0260801c91611579565b6001600160881b03600160801b61156e565b6040516315e4079d60e11b8152600490fd5b80611541565b6001600160a01b038116906401000276a382101580611b36575b15611b2457640100000000600160c01b039060201b16806001600160801b03811160071b9181831c9267ffffffffffffffff841160061b93841c9363ffffffff851160051b94851c9461ffff861160041b95861c60ff9687821160031b91821c92600f841160021b93841c94600160038711811b96871c119617171717171717916080831015600014611b185750607e1982011c5b8002607f928392828493841c81841c1c800280851c81851c1c800280861c81861c1c800280871c81871c1c80029081881c82881c1c80029283891c84891c1c800294858a1c868a1c1c800296878b1c888b1c1c800298898c1c8a8c1c1c80029a8b8d1c8c821c1c8002809d1c8d821c1c8002809e81901c90821c1c80029e8f80911c911c1c800260cd1c6604000000000000169d60cc1c6608000000000000169c60cb1c6610000000000000169b60ca1c6620000000000000169a60c91c6640000000000000169960c81c6680000000000000169860c71c670100000000000000169760c61c670200000000000000169660c51c670400000000000000169560c41c670800000000000000169460c31c671000000000000000169360c21c672000000000000000169260c11c674000000000000000169160c01c6780000000000000001690607f190160401b1717171717171717171717171717693627a301d71055774c85026f028f6481ab7f045a5af012a19d003aa919810160801d60020b906fdb2df09e81959a81455e260799a0632f0160801d60020b91600090838314600014611afc575050905090565b611b0861091f8561152f565b119050611b13575090565b905090565b905081607f031b61194e565b6040516324c070df60e11b8152600490fd5b5073fffd8963efd1fc6a506488495d951d5263988d2682106118b9565b90816020910312610174575180151581036101745790565b6000928380936040519060208201936323b872dd60e01b855260018060a01b0380921660248401521660448201526aa56fa5b99019a5c8000000606482015260648152611bb7816102be565b51925af1611bc3610f04565b81611bfb575b5015611bd157565b60405162461bcd60e51b81526020600482015260026024820152612a2360f11b6044820152606490fd5b8051801592508215611c10575b505038611bc9565b611c239250602080918301019101611b53565b3880611c08565b60405163a9059cbb60e01b602082019081526001600160a01b03909316602482015260448082019490945292835260808301929167ffffffffffffffff8411838510176102da576000809493819460405251925af1611bc3610f04565b60405163095ea7b360e01b602082019081526001600160a01b0393909316602482015260001960448083019190915281526000928392918390611ccb6064826102df565b51925af1611cd7610f04565b81611d10575b5015611ce557565b60405162461bcd60e51b815260206004820152600360248201526229a0a360e91b6044820152606490fd5b8051801592508215611d25575b505038611cdd565b611d389250602080918301019101611b53565b3880611d1d565b8060601b600160601b91808204831490151715610e9457670de0b6b3a764000090048060601b918183041490151715610e9457610f93908015611eb05780611e49611e42611e38611e2e611e24611e1a611e10611e066001610f939a6000908b60801c80611ea4575b508060401c80611e97575b508060201c80611e8a575b508060101c80611e7d575b508060081c80611e70575b508060041c80611e63575b508060021c80611e56575b50821c611e4f575b811c1b611dff818b6113c6565b0160011c90565b611dff818a6113c6565b611dff81896113c6565b611dff81886113c6565b611dff81876113c6565b611dff81866113c6565b611dff81856113c6565b80926113c6565b90611eb6565b8101611df2565b6002915091019038611dea565b6004915091019038611ddf565b6008915091019038611dd4565b6010915091019038611dc9565b6020915091019038611dbe565b6040915091019038611db3565b91505060809038611da8565b50600090565b9080821015611b13575090565b6001600160a01b039182169082160391908211610e9457565b610f9391611f2c9173ff53611968f1e5ca45cfca7918447e7f5776f6d4906001600160a01b03908181168310611f31575b611f2590611f1f838516848316611f54565b93611ec3565b169161202b565b612056565b915073ff53611968f1e5ca45cfca7918447e7f5776f6d4611f0d565b1561017457565b6000198282099082810292838084109303928084039314611f9057600160601b9183831115610174570990828211900360a01b910360601c1790565b50505060601c90565b600160601b91600019838309928260601b9283808610950394808603951461201857908291611fc9868411611f4d565b0981806000031680920460028082600302188083028203028083028203028083028203028083028203028083028203028092029003029360018380600003040190848311900302920304170290565b50509150612027821515611f4d565b0490565b9091600019838309928083029283808610950394808603951461201857908291611fc9868411611f4d565b906001600160801b03821691820361017457565b610f9391611f2c916001600160a01b039061209490640100ad139c90838116821161209b57611ec3565b1690611f99565b90611ec3565b916001600160a01b03916120e19190808416848616116120ee575b83806120c88784611ec3565b169116916001600160801b0360601b9060601b1661202b565b91169081156113a3570490565b936120bc56fea2646970667358221220cc87fb2c192559c02998fa94c82e4929518f2427b7a56a4029089a841f0e01d764736f6c63430008130033",
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
// Solidity: function settleFund(uint256 _baseFundBalanceBeforeFee) returns()
func (_DAOTreasury *DAOTreasuryTransactor) SettleFund(opts *bind.TransactOpts, _baseFundBalanceBeforeFee *big.Int) (*types.Transaction, error) {
	return _DAOTreasury.contract.Transact(opts, "settleFund", _baseFundBalanceBeforeFee)
}

// SettleFund is a paid mutator transaction binding the contract method 0x14a8fc97.
//
// Solidity: function settleFund(uint256 _baseFundBalanceBeforeFee) returns()
func (_DAOTreasury *DAOTreasurySession) SettleFund(_baseFundBalanceBeforeFee *big.Int) (*types.Transaction, error) {
	return _DAOTreasury.Contract.SettleFund(&_DAOTreasury.TransactOpts, _baseFundBalanceBeforeFee)
}

// SettleFund is a paid mutator transaction binding the contract method 0x14a8fc97.
//
// Solidity: function settleFund(uint256 _baseFundBalanceBeforeFee) returns()
func (_DAOTreasury *DAOTreasuryTransactorSession) SettleFund(_baseFundBalanceBeforeFee *big.Int) (*types.Transaction, error) {
	return _DAOTreasury.Contract.SettleFund(&_DAOTreasury.TransactOpts, _baseFundBalanceBeforeFee)
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
