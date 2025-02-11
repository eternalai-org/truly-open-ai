package port

import (
	"context"
	"math/big"

	"agent-battle/internal/contract/erc20"

	"github.com/ethereum/go-ethereum/common"
)

type IContractErc20Usecase interface {
	BalanceOfAddress(ctx context.Context, address string) (*big.Int, error)
	FilterTransfer(ctx context.Context, startBlock, endBlock uint64, from, to []common.Address) (*erc20.Erc20TransferIterator, error)
	CurrentBlockNumber(ctx context.Context) (uint64, error)
	TransferToken(ctx context.Context, toAddress string, amount *big.Int, privateKey string) (string, error)
	TransferETH(ctx context.Context, toAddress string, amount *big.Int, privateKey string) (string, error)
	EstimateGasFee(
		ctx context.Context,
		fromAddress string,
		toAddress string,
		amount *big.Int,
	) (*big.Int, error)
}
