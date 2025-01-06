// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wbvm

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

// WBVMMetaData contains all meta data concerning the WBVM contract.
var WBVMMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x0002000000000002000200000000000200000000030100190000006003300270000000df0330019700010000003103550000008004000039000000400040043f0000000100200190000000300000c13d000000040030008c000000670000413d000000000201043b000000e002200270000000e40020009c0000006c0000a13d000000e50020009c0000009a0000213d000000e90020009c0000016b0000613d000000ea0020009c000001280000613d000000eb0020009c000002260000c13d0000000001000416000000000001004b000002260000c13d0000000103000039000000000203041a000000010520019000000001012002700000007f0410018f00000000010460190000001f0010008c00000000060000190000000106002039000000000662013f0000000100600190000000610000c13d000000800010043f000000000005004b000001af0000c13d0000010001200197000000a00010043f000000000004004b000000c001000039000000a001006039000001bf0000013d0000000001000416000000000001004b000002260000c13d000000000100041a000000010210019000000001031002700000007f0330618f0000001f0030008c00000000010000190000000101002039000000000012004b000000610000c13d000000200030008c000000540000413d000200000003001d00000000000004350000000001000414000000df0010009c000000df01008041000000c001100210000000e0011001c70000801002000039037703720000040f0000000100200190000002260000613d000000000101043b00000002020000290000001f0220003900000005022002700000000002210019000000000021004b000000540000813d000000000001041b0000000101100039000000000021004b000000500000413d000000e101000041000000000010041b0000000103000039000000000103041a000000010010019000000001041002700000007f0440618f0000001f0040008c00000000020000190000000102002039000000000121013f0000000100100190000000ba0000613d000000f90100004100000000001004350000002201000039000000040010043f000000fa010000410000037900010430000000000003004b000002260000c13d037702550000040f0000000001000019000003780001042e000000ec0020009c000000df0000a13d000000ed0020009c0000013b0000613d000000ee0020009c000001430000613d000000ef0020009c000002260000c13d000000240030008c000002260000413d0000000002000416000000000002004b000002260000c13d0000000401100370000000000101043b000200000001001d000000000100041100000000001004350000000301000039000000200010043f0000000001000414000000df0010009c000000df01008041000000c001100210000000f6011001c70000801002000039037703720000040f0000000100200190000002260000613d000000000101043b000000000101041a000000020010006c000001890000813d000000400100043d000000fc02000041000000000021043500000004021000390000002003000039000000000032043500000024021000390000000000020435000000df0010009c000000df010080410000004001100210000000fd011001c70000037900010430000000e60020009c000001740000613d000000e70020009c000000690000613d000000e80020009c000002260000c13d000000440030008c000002260000413d0000000002000416000000000002004b000002260000c13d0000000402100370000000000202043b000000f20020009c000002260000213d0000002401100370000000000101043b000000f20010009c000002260000213d0000000000200435000200000001001d0000000401000039000000200010043f00000040020000390000000001000019037703490000040f00000002020000290000000000200435000000200010043f00000000010000190000004002000039000001360000013d0000001f0040008c000000d30000a13d000200000004001d00000000003004350000000001000414000000df0010009c000000df01008041000000c001100210000000e0011001c70000801002000039037703720000040f0000000100200190000002260000613d000000000101043b00000002020000290000001f0220003900000005022002700000000002210019000000000021004b0000000103000039000000d30000813d000000000001041b0000000101100039000000000021004b000000cf0000413d000000e201000041000000000013041b0000000201000039000000000201041a000001000220019700000012022001bf000000000021041b000000200100003900000100001004430000012000000443000000e301000041000003780001042e000000f00020009c000001540000613d000000f10020009c000002260000c13d000000440030008c000002260000413d0000000002000416000000000002004b000002260000c13d0000000402100370000000000202043b000200000002001d000000f20020009c000002260000213d0000002401100370000000000101043b000100000001001d000000000100041100000000001004350000000401000039000000200010043f0000000001000414000000df0010009c000000df01008041000000c001100210000000f6011001c70000801002000039037703720000040f0000000100200190000002260000613d000000000101043b00000002020000290000000000200435000000200010043f0000000001000414000000df0010009c000000df01008041000000c001100210000000f6011001c70000801002000039037703720000040f0000000100200190000002260000613d000000000101043b0000000102000029000000000021041b000000400100043d0000000000210435000000df0010009c000000df0100804100000040011002100000000002000414000000df0020009c000000df02008041000000c002200210000000000112019f000000e0011001c70000800d020000390000000303000039000000fe04000041000000000500041100000002060000290377036d0000040f0000000100200190000002260000613d000000400100043d00000001020000390000000000210435000000df0010009c000000df010080410000004001100210000000f4011001c7000003780001042e000000240030008c000002260000413d0000000002000416000000000002004b000002260000c13d0000000401100370000000000101043b000000f20010009c000002260000213d00000000001004350000000301000039000000200010043f00000040020000390000000001000019037703490000040f000000000101041a000000800010043f000000f301000041000003780001042e0000000001000416000000000001004b000002260000c13d00000000010004100377035e0000040f000000800010043f000000f301000041000003780001042e000000640030008c000002260000413d0000000002000416000000000002004b000002260000c13d0000000402100370000000000402043b000000f20040009c000002260000213d0000002402100370000000000202043b000000f20020009c000002260000213d0000004401100370000000000301043b0000000001040019000001800000013d0000000001000416000000000001004b000002260000c13d000000000200041a000000010420019000000001012002700000007f0310018f00000000010360190000001f0010008c00000000050000190000000105002039000000000552013f0000000100500190000000610000c13d000000800010043f000000000004004b000001a10000c13d0000010001200197000000a00010043f000000000003004b000000c001000039000000a001006039000001bf0000013d0000000001000416000000000001004b000002260000c13d0000000201000039000000000101041a000000ff0110018f000000800010043f000000f301000041000003780001042e000000440030008c000002260000413d0000000002000416000000000002004b000002260000c13d0000000402100370000000000202043b000000f20020009c000002260000213d0000002401100370000000000301043b0000000001000411037702840000040f0000000101000039000000400200043d0000000000120435000000df0020009c000000df020080410000004001200210000000f4011001c7000003780001042e000000000100041100000000001004350000000301000039000000200010043f0000000001000414000000df0010009c000000df01008041000000c001100210000000f6011001c70000801002000039037703720000040f0000000100200190000002260000613d000000000101043b000000000201041a0000000203000029000000000232004b000001d00000813d000000f90100004100000000001004350000001101000039000000040010043f000000fa0100004100000379000104300000000000000435000000020020008c000001b20000413d000000ff0200004100000000040000190000000003040019000000000402041a000000a005300039000000000045043500000001022000390000002004300039000000000014004b000001a60000413d000001be0000013d0000000000300435000000020020008c000001b40000813d000000a001000039000001bf0000013d000000f50200004100000000040000190000000003040019000000000402041a000000a005300039000000000045043500000001022000390000002004300039000000000014004b000001b60000413d000000c001300039000000800210008a00000080010000390377022e0000040f000000400100043d000200000001001d0000008002000039037702400000040f00000002020000290000000001210049000000df0010009c000000df010080410000006001100210000000df0020009c000000df020080410000004002200210000000000121019f000003780001042e000000000021041b00000000010004140000000002000411000000040020008c000001d70000c13d0000000002000031000001e70000013d000000df0010009c000000df01008041000000c001100210000000000003004b000001e00000613d000000f7011001c70000800902000039000000000400041100000000050000190377036d0000040f000000020300002900000000020100190000006002200270000000df0020019d000000df022001970001000000010355000000000002004b000002120000613d0000001f0120003900000101011001970000003f011000390000010101100197000000400500043d0000000001150019000000000051004b00000000040000190000000104004039000000f80010009c000002280000213d0000000100400190000002280000c13d000000400010043f0000001f0120018f0000000009250436000000010400036700000005022002720000000502200210000002040000613d0000000005290019000000000604034f0000000007090019000000006806043c0000000007870436000000000057004b000002000000c13d000000000001004b000002120000613d000000000424034f00000000022900190000000301100210000000000502043300000000051501cf000000000515022f000000000404043b0000010001100089000000000414022f00000000011401cf000000000151019f0000000000120435000000400100043d0000000000310435000000df0010009c000000df0100804100000040011002100000000002000414000000df0020009c000000df02008041000000c002200210000000000112019f000000e0011001c70000800d020000390000000203000039000000fb0400004100000000050004110377036d0000040f0000000100200190000002260000613d0000000001000019000003780001042e00000000010000190000037900010430000000f90100004100000000001004350000004101000039000000040010043f000000fa0100004100000379000104300000001f0220003900000101022001970000000001120019000000000021004b00000000020000190000000102004039000000f80010009c0000023a0000213d00000001002001900000023a0000c13d000000400010043f000000000001042d000000f90100004100000000001004350000004101000039000000040010043f000000fa01000041000003790001043000000020030000390000000004310436000000003202043400000000002404350000004001100039000000000002004b0000024f0000613d000000000400001900000000054100190000000006430019000000000606043300000000006504350000002004400039000000000024004b000002480000413d000000000321001900000000000304350000001f0220003900000101022001970000000001210019000000000001042d000000000100041100000000001004350000000301000039000000200010043f0000000001000414000000df0010009c000000df01008041000000c001100210000000f6011001c70000801002000039037703720000040f00000001002001900000027c0000613d000000000101043b000000000301041a0000000002000416000000000023001a0000027e0000413d0000000003230019000000000031041b000000400100043d0000000000210435000000df0010009c000000df0100804100000040011002100000000002000414000000df0020009c000000df02008041000000c002200210000000000112019f000000e0011001c70000800d020000390000000203000039000001020400004100000000050004110377036d0000040f00000001002001900000027c0000613d000000000001042d00000000010000190000037900010430000000f90100004100000000001004350000001101000039000000040010043f000000fa0100004100000379000104300003000000000002000300000003001d000100000002001d000000f201100197000200000001001d00000000001004350000000301000039000000200010043f0000000001000414000000df0010009c000000df01008041000000c001100210000000f6011001c70000801002000039037703720000040f0000000100200190000003330000613d000000000101043b000000000101041a000000030010006c000003350000413d00000000020004110000000201000029000000000021004b000002f80000613d00000000001004350000000401000039000000200010043f0000000001000414000000df0010009c000000df01008041000000c001100210000000f6011001c70000801002000039037703720000040f0000000100200190000003330000613d000000000101043b00000000020004110000000000200435000000200010043f0000000001000414000000df0010009c000000df01008041000000c001100210000000f6011001c70000801002000039037703720000040f0000000100200190000003330000613d000000000101043b000000000101041a000001040010009c0000000201000029000002f80000613d00000000001004350000000401000039000000200010043f0000000001000414000000df0010009c000000df01008041000000c001100210000000f6011001c70000801002000039037703720000040f0000000100200190000003330000613d000000000101043b00000000020004110000000000200435000000200010043f0000000001000414000000df0010009c000000df01008041000000c001100210000000f6011001c70000801002000039037703720000040f0000000100200190000003330000613d000000000101043b000000000101041a000000030010006c000003350000413d000000020100002900000000001004350000000401000039000000200010043f0000000001000414000000df0010009c000000df01008041000000c001100210000000f6011001c70000801002000039037703720000040f0000000100200190000003330000613d000000000101043b00000000020004110000000000200435000000200010043f0000000001000414000000df0010009c000000df01008041000000c001100210000000f6011001c70000801002000039037703720000040f0000000100200190000003330000613d000000000101043b000000000201041a000000030220006c000003420000413d000000000021041b000000020100002900000000001004350000000301000039000000200010043f0000000001000414000000df0010009c000000df01008041000000c001100210000000f6011001c70000801002000039037703720000040f0000000100200190000003330000613d000000000101043b000000000201041a000000030220006c000003420000413d000000000021041b0000000101000029000000f201100197000100000001001d00000000001004350000000301000039000000200010043f0000000001000414000000df0010009c000000df01008041000000c001100210000000f6011001c70000801002000039037703720000040f0000000100200190000003330000613d000000000101043b000000000201041a0000000303000029000000000032001a000003420000413d0000000002320019000000000021041b000000400100043d0000000000310435000000df0010009c000000df0100804100000040011002100000000002000414000000df0020009c000000df02008041000000c002200210000000000112019f000000e0011001c70000800d0200003900000003030000390000010304000041000000020500002900000001060000290377036d0000040f0000000100200190000003330000613d000000000001042d00000000010000190000037900010430000000400100043d000000fc02000041000000000021043500000004021000390000002003000039000000000032043500000024021000390000000000020435000000df0010009c000000df010080410000004001100210000000fd011001c70000037900010430000000f90100004100000000001004350000001101000039000000040010043f000000fa010000410000037900010430000000000001042f000000df0010009c000000df010080410000004001100210000000df0020009c000000df020080410000006002200210000000000112019f0000000002000414000000df0020009c000000df02008041000000c002200210000000000112019f000000f7011001c70000801002000039037703720000040f00000001002001900000035c0000613d000000000101043b000000000001042d000000000100001900000379000104300000010502000041000000000020043900000004001004430000000001000414000000df0010009c000000df01008041000000c00110021000000106011001c70000800a02000039037703720000040f00000001002001900000036c0000613d000000000101043b000000000001042d000000000001042f00000370002104210000000102000039000000000001042d0000000002000019000000000001042d00000375002104230000000102000039000000000001042d0000000002000019000000000001042d0000037700000432000003780001042e00000379000104300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffff0200000000000000000000000000000000000020000000000000000000000000577261707065642042564d0000000000000000000000000000000000000000165742564d00000000000000000000000000000000000000000000000000000008000000020000000000000000000000000000004000000100000000000000000000000000000000000000000000000000000000000000000000000000313ce56600000000000000000000000000000000000000000000000000000000a9059cba00000000000000000000000000000000000000000000000000000000a9059cbb00000000000000000000000000000000000000000000000000000000d0e30db000000000000000000000000000000000000000000000000000000000dd62ed3e00000000000000000000000000000000000000000000000000000000313ce5670000000000000000000000000000000000000000000000000000000070a082310000000000000000000000000000000000000000000000000000000095d89b410000000000000000000000000000000000000000000000000000000018160ddc0000000000000000000000000000000000000000000000000000000018160ddd0000000000000000000000000000000000000000000000000000000023b872dd000000000000000000000000000000000000000000000000000000002e1a7d4d0000000000000000000000000000000000000000000000000000000006fdde0300000000000000000000000000000000000000000000000000000000095ea7b3000000000000000000000000ffffffffffffffffffffffffffffffffffffffff00000000000000000000000000000000000000200000008000000000000000000000000000000000000000000000000000000020000000000000000000000000b10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf602000000000000000000000000000000000000400000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffff4e487b710000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000240000000000000000000000007fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b6508c379a00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000440000000000000000000000008c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0e1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109cddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9cc7f708afc65944829bd487b90b72536b1951864fbfc14e125fc972a6507f3902000002000000000000000000000000000000240000000000000000000000000000000000000000000000000000000000000000000000000000000000000000df38709dc40c99c8c317fcf754c75acd1d6005a5a0d105f2d796ae36da5c3ebf",
}

// WBVMABI is the input ABI used to generate the binding from.
// Deprecated: Use WBVMMetaData.ABI instead.
var WBVMABI = WBVMMetaData.ABI

// WBVMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WBVMMetaData.Bin instead.
var WBVMBin = WBVMMetaData.Bin

// DeployWBVM deploys a new Ethereum contract, binding an instance of WBVM to it.
func DeployWBVM(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WBVM, error) {
	parsed, err := WBVMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WBVMBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WBVM{WBVMCaller: WBVMCaller{contract: contract}, WBVMTransactor: WBVMTransactor{contract: contract}, WBVMFilterer: WBVMFilterer{contract: contract}}, nil
}

// WBVM is an auto generated Go binding around an Ethereum contract.
type WBVM struct {
	WBVMCaller     // Read-only binding to the contract
	WBVMTransactor // Write-only binding to the contract
	WBVMFilterer   // Log filterer for contract events
}

// WBVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type WBVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WBVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WBVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WBVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WBVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WBVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WBVMSession struct {
	Contract     *WBVM             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WBVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WBVMCallerSession struct {
	Contract *WBVMCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WBVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WBVMTransactorSession struct {
	Contract     *WBVMTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WBVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type WBVMRaw struct {
	Contract *WBVM // Generic contract binding to access the raw methods on
}

// WBVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WBVMCallerRaw struct {
	Contract *WBVMCaller // Generic read-only contract binding to access the raw methods on
}

// WBVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WBVMTransactorRaw struct {
	Contract *WBVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWBVM creates a new instance of WBVM, bound to a specific deployed contract.
func NewWBVM(address common.Address, backend bind.ContractBackend) (*WBVM, error) {
	contract, err := bindWBVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WBVM{WBVMCaller: WBVMCaller{contract: contract}, WBVMTransactor: WBVMTransactor{contract: contract}, WBVMFilterer: WBVMFilterer{contract: contract}}, nil
}

// NewWBVMCaller creates a new read-only instance of WBVM, bound to a specific deployed contract.
func NewWBVMCaller(address common.Address, caller bind.ContractCaller) (*WBVMCaller, error) {
	contract, err := bindWBVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WBVMCaller{contract: contract}, nil
}

// NewWBVMTransactor creates a new write-only instance of WBVM, bound to a specific deployed contract.
func NewWBVMTransactor(address common.Address, transactor bind.ContractTransactor) (*WBVMTransactor, error) {
	contract, err := bindWBVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WBVMTransactor{contract: contract}, nil
}

// NewWBVMFilterer creates a new log filterer instance of WBVM, bound to a specific deployed contract.
func NewWBVMFilterer(address common.Address, filterer bind.ContractFilterer) (*WBVMFilterer, error) {
	contract, err := bindWBVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WBVMFilterer{contract: contract}, nil
}

// bindWBVM binds a generic wrapper to an already deployed contract.
func bindWBVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WBVMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WBVM *WBVMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WBVM.Contract.WBVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WBVM *WBVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WBVM.Contract.WBVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WBVM *WBVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WBVM.Contract.WBVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WBVM *WBVMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WBVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WBVM *WBVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WBVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WBVM *WBVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WBVM.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_WBVM *WBVMCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WBVM.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_WBVM *WBVMSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _WBVM.Contract.Allowance(&_WBVM.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_WBVM *WBVMCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _WBVM.Contract.Allowance(&_WBVM.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_WBVM *WBVMCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WBVM.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_WBVM *WBVMSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _WBVM.Contract.BalanceOf(&_WBVM.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_WBVM *WBVMCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _WBVM.Contract.BalanceOf(&_WBVM.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WBVM *WBVMCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WBVM.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WBVM *WBVMSession) Decimals() (uint8, error) {
	return _WBVM.Contract.Decimals(&_WBVM.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WBVM *WBVMCallerSession) Decimals() (uint8, error) {
	return _WBVM.Contract.Decimals(&_WBVM.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WBVM *WBVMCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WBVM.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WBVM *WBVMSession) Name() (string, error) {
	return _WBVM.Contract.Name(&_WBVM.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WBVM *WBVMCallerSession) Name() (string, error) {
	return _WBVM.Contract.Name(&_WBVM.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WBVM *WBVMCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WBVM.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WBVM *WBVMSession) Symbol() (string, error) {
	return _WBVM.Contract.Symbol(&_WBVM.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WBVM *WBVMCallerSession) Symbol() (string, error) {
	return _WBVM.Contract.Symbol(&_WBVM.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WBVM *WBVMCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WBVM.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WBVM *WBVMSession) TotalSupply() (*big.Int, error) {
	return _WBVM.Contract.TotalSupply(&_WBVM.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WBVM *WBVMCallerSession) TotalSupply() (*big.Int, error) {
	return _WBVM.Contract.TotalSupply(&_WBVM.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_WBVM *WBVMTransactor) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WBVM.contract.Transact(opts, "approve", guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_WBVM *WBVMSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WBVM.Contract.Approve(&_WBVM.TransactOpts, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_WBVM *WBVMTransactorSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WBVM.Contract.Approve(&_WBVM.TransactOpts, guy, wad)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WBVM *WBVMTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WBVM.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WBVM *WBVMSession) Deposit() (*types.Transaction, error) {
	return _WBVM.Contract.Deposit(&_WBVM.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WBVM *WBVMTransactorSession) Deposit() (*types.Transaction, error) {
	return _WBVM.Contract.Deposit(&_WBVM.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_WBVM *WBVMTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WBVM.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_WBVM *WBVMSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WBVM.Contract.Transfer(&_WBVM.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_WBVM *WBVMTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WBVM.Contract.Transfer(&_WBVM.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_WBVM *WBVMTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WBVM.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_WBVM *WBVMSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WBVM.Contract.TransferFrom(&_WBVM.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_WBVM *WBVMTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WBVM.Contract.TransferFrom(&_WBVM.TransactOpts, src, dst, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WBVM *WBVMTransactor) Withdraw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _WBVM.contract.Transact(opts, "withdraw", wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WBVM *WBVMSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WBVM.Contract.Withdraw(&_WBVM.TransactOpts, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WBVM *WBVMTransactorSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WBVM.Contract.Withdraw(&_WBVM.TransactOpts, wad)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WBVM *WBVMTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WBVM.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WBVM *WBVMSession) Receive() (*types.Transaction, error) {
	return _WBVM.Contract.Receive(&_WBVM.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WBVM *WBVMTransactorSession) Receive() (*types.Transaction, error) {
	return _WBVM.Contract.Receive(&_WBVM.TransactOpts)
}

// WBVMApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WBVM contract.
type WBVMApprovalIterator struct {
	Event *WBVMApproval // Event containing the contract specifics and raw log

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
func (it *WBVMApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WBVMApproval)
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
		it.Event = new(WBVMApproval)
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
func (it *WBVMApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WBVMApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WBVMApproval represents a Approval event raised by the WBVM contract.
type WBVMApproval struct {
	Src common.Address
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_WBVM *WBVMFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, guy []common.Address) (*WBVMApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _WBVM.contract.FilterLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return &WBVMApprovalIterator{contract: _WBVM.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_WBVM *WBVMFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WBVMApproval, src []common.Address, guy []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _WBVM.contract.WatchLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WBVMApproval)
				if err := _WBVM.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_WBVM *WBVMFilterer) ParseApproval(log types.Log) (*WBVMApproval, error) {
	event := new(WBVMApproval)
	if err := _WBVM.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WBVMDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the WBVM contract.
type WBVMDepositIterator struct {
	Event *WBVMDeposit // Event containing the contract specifics and raw log

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
func (it *WBVMDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WBVMDeposit)
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
		it.Event = new(WBVMDeposit)
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
func (it *WBVMDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WBVMDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WBVMDeposit represents a Deposit event raised by the WBVM contract.
type WBVMDeposit struct {
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_WBVM *WBVMFilterer) FilterDeposit(opts *bind.FilterOpts, dst []common.Address) (*WBVMDepositIterator, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WBVM.contract.FilterLogs(opts, "Deposit", dstRule)
	if err != nil {
		return nil, err
	}
	return &WBVMDepositIterator{contract: _WBVM.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_WBVM *WBVMFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WBVMDeposit, dst []common.Address) (event.Subscription, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WBVM.contract.WatchLogs(opts, "Deposit", dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WBVMDeposit)
				if err := _WBVM.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_WBVM *WBVMFilterer) ParseDeposit(log types.Log) (*WBVMDeposit, error) {
	event := new(WBVMDeposit)
	if err := _WBVM.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WBVMTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WBVM contract.
type WBVMTransferIterator struct {
	Event *WBVMTransfer // Event containing the contract specifics and raw log

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
func (it *WBVMTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WBVMTransfer)
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
		it.Event = new(WBVMTransfer)
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
func (it *WBVMTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WBVMTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WBVMTransfer represents a Transfer event raised by the WBVM contract.
type WBVMTransfer struct {
	Src common.Address
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_WBVM *WBVMFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*WBVMTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WBVM.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &WBVMTransferIterator{contract: _WBVM.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_WBVM *WBVMFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WBVMTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WBVM.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WBVMTransfer)
				if err := _WBVM.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_WBVM *WBVMFilterer) ParseTransfer(log types.Log) (*WBVMTransfer, error) {
	event := new(WBVMTransfer)
	if err := _WBVM.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WBVMWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the WBVM contract.
type WBVMWithdrawalIterator struct {
	Event *WBVMWithdrawal // Event containing the contract specifics and raw log

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
func (it *WBVMWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WBVMWithdrawal)
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
		it.Event = new(WBVMWithdrawal)
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
func (it *WBVMWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WBVMWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WBVMWithdrawal represents a Withdrawal event raised by the WBVM contract.
type WBVMWithdrawal struct {
	Src common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_WBVM *WBVMFilterer) FilterWithdrawal(opts *bind.FilterOpts, src []common.Address) (*WBVMWithdrawalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}

	logs, sub, err := _WBVM.contract.FilterLogs(opts, "Withdrawal", srcRule)
	if err != nil {
		return nil, err
	}
	return &WBVMWithdrawalIterator{contract: _WBVM.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_WBVM *WBVMFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *WBVMWithdrawal, src []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}

	logs, sub, err := _WBVM.contract.WatchLogs(opts, "Withdrawal", srcRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WBVMWithdrawal)
				if err := _WBVM.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_WBVM *WBVMFilterer) ParseWithdrawal(log types.Log) (*WBVMWithdrawal, error) {
	event := new(WBVMWithdrawal)
	if err := _WBVM.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
