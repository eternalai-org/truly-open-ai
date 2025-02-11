package usecase

import (

"agent-battle/internal/contract/erc20"
"agent-battle/internal/core/port"
"context"
"github.com/ethereum/go-ethereum/common"
"github.com/spf13/viper"
"math/big"
"testing"
)

var (
	erc20Usecase port.IContractErc20Usecase
	toAddress    string
	fromAddress string
	privateKey string
)

func init() {
	viper.SetConfigFile(`../../../../env/local.worker.test.yml`)
	viper.ReadInConfig()
	initVars()
}

// initVars initializes the variables for the tests
func initVars() {
	var err error
	erc20Usecase, err = NewContractErc20Usecase()
	if err != nil {
		panic(err)
	}

	toAddress = "0x976d5565927cf44ee19c346f61fcb37238b426d1"
	privateKey = "bbd4ef749a6174d30490e67b1be7b8c9d1e80fae775ca8dce8d83fedb9ebc247"
	fromAddress = "0xf0c08db5d131f1f77ed96a1d88a31c3f98ed7bdf"
}

func Test_contractErc20Usecase_TransferToken(t *testing.T) {
	t.Skip("Skip Test Transfer Token")
	// // Generate AES Key
    // key, err := encrypt.GenerateAESKey(32)
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }

	// encryptedPrivateKey, fromAddress, err := utils.GenerateAddress(key)
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	//
	// // Decrypt Private Key
	// privateKey, err := encrypt.DecryptToString(encryptedPrivateKey, key)
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }


	type fields struct {
		erc20Usecase port.IContractErc20Usecase
	}
	type args struct {
		ctx         context.Context
		fromAddress string
		toAddress   string
		amount      *big.Int
		privateKey  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test Transfer Token",
			fields: fields{
				erc20Usecase: erc20Usecase,
			},
			args: args{
				ctx:        context.Background(),
				privateKey: privateKey,
				// 1 USDC = 1000000
				amount:    big.NewInt(1000000),
				toAddress: toAddress,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := tt.fields.erc20Usecase
			got, err := uc.TransferToken(tt.args.ctx, tt.args.toAddress, tt.args.amount, tt.args.privateKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransferToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == "" {
				t.Errorf("TransferToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contractErc20Usecase_FilterTransfer(t *testing.T) {
	type fields struct {
		erc20Usecase port.IContractErc20Usecase
	}

	currentBlock, err := erc20Usecase.CurrentBlockNumber(context.Background())
	if err != nil {
		t.Error(err)
		return
	}

	fromBlock := currentBlock - 10_000

	type args struct {
		ctx        context.Context
		startBlock uint64
		endBlock   uint64
		from       []common.Address
		to         []common.Address
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *erc20.Erc20TransferIterator
		wantErr bool
	}{
		{
			name: "Test Filter Transfer",
			fields: fields{
				erc20Usecase: erc20Usecase,
			},
			args: args{
				ctx:        context.Background(),
				startBlock: fromBlock,
				endBlock:   currentBlock,
				from:       []common.Address{
					common.HexToAddress(fromAddress),
				},
				to: []common.Address{
					common.HexToAddress(toAddress),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := tt.fields.erc20Usecase
			got, err := uc.FilterTransfer(tt.args.ctx, tt.args.startBlock, tt.args.endBlock, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("FilterTransfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for got.Next() {
				e := got.Event
				t.Logf("Block Number: %v, txHash: %v, From: %v, To: %v, Value: %v", e.Raw.BlockNumber, e.Raw.TxHash, e.From.Hex(), e.To.Hex(), e.Value)
			}
		})
	}
}

func Test_contractErc20Usecase_TransferETH(t *testing.T) {
	t.Skip("Skip Test Transfer ETH")
	type fields struct {
		erc20Usecase port.IContractErc20Usecase
	}
	type args struct {
		ctx         context.Context
		fromAddress string
		toAddress   string
		amount      *big.Int
		privateKey  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test Transfer ETH",
			fields: fields{
				erc20Usecase: erc20Usecase,
			},
			args: args{
				ctx:        context.Background(),
				privateKey: privateKey,
				// 0.00001 ETH = 100000000000000
				amount:    big.NewInt(100000000000000),
				toAddress: toAddress,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := tt.fields.erc20Usecase
			got, err := uc.TransferETH(tt.args.ctx, tt.args.toAddress, tt.args.amount, tt.args.privateKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransferToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == "" {
				t.Errorf("TransferToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}