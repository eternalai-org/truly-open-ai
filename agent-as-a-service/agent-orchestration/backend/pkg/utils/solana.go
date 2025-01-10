package utils

import (
	"context"
	"math/big"
	"strings"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/pkg/encrypt"

	"github.com/gagliardetto/solana-go/rpc"

	"github.com/gagliardetto/solana-go"
)

func GenerateSolanaAddress(secretKey string) (string, string, error) {
	account := solana.NewWallet()

	// Get the public key
	pubKey := account.PublicKey().String()

	privKey, err := encrypt.EncryptToString(account.PrivateKey.String(), secretKey)
	if err != nil {
		return "", "", err
	}
	return privKey, strings.ToLower(pubKey), nil
}

func GetBalanceOnSolanaChain(ctx context.Context, address string) (*big.Int, error) {
	pubKey, err := solana.PublicKeyFromBase58(address)
	if err != nil {
		return nil, err
	}
	endpoint := rpc.MainNetBeta_RPC
	client := rpc.New(endpoint)
	out, err := client.GetBalance(
		ctx,
		pubKey,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		return nil, err
	}
	lamportsOnAccount := new(big.Int).SetUint64(out.Value)
	return lamportsOnAccount, nil
}
