package usecase

import (

"agent-battle/internal/contract/erc20"
"agent-battle/internal/core/port"
	"agent-battle/pkg/cryptoamount"
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

	toAddress = viper.GetString("TO_ADDRESS")
	privateKey = viper.GetString("FROM_ADDRESS_PRIVATE_KEY")
	fromAddress = viper.GetString("FROM_ADDRESS")
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
	currentBlock = 26238476

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
				cryptoAmount := cryptoamount.NewCryptoAmountFromBigInt(e.Value)

				t.Logf("Block Number: %v, txHash: %v, From: %v, To: %v, Value: %v, cryptoAmount: %v", e.Raw.BlockNumber, e.Raw.TxHash, e.From.Hex(), e.To.Hex(), e.Value, cryptoAmount)
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