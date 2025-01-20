package zkapi

import (
	"context"
	"encoding/hex"
	"errors"
	"math/big"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/ihybridmodel"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zksync-sdk/zksync2-go/accounts"
	zktypes "github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
)

func (c *Client) HybridModelInfer(contractAddress string, prkHex string, inferData string) (string, error) {
	if contractAddress == "" ||
		!common.IsHexAddress(contractAddress) {
		return "", errors.New("contractAddress is invalid")
	}
	pbkHex, _, err := c.parsePrkAuth(prkHex)
	if err != nil {
		return "", err
	}
	gasPrice, err := c.getGasPrice()
	if err != nil {
		return "", err
	}
	client, err := c.getZkClient()
	if err != nil {
		return "", err
	}
	contractAddr := helpers.HexToAddress(contractAddress)
	// EstimateGas
	instanceABI, err := ihybridmodel.IHybridModelMetaData.GetAbi()
	if err != nil {
		return "", err
	}
	dataBytes, err := instanceABI.Pack(
		"infer",
		[]byte(inferData),
	)
	if err != nil {
		return "", err
	}
	gasNumber, err := client.EstimateGasL2(
		context.Background(),
		zktypes.CallMsg{
			CallMsg: ethereum.CallMsg{
				From: pbkHex,
				To:   &contractAddr,
				Data: dataBytes,
			},
			Meta: &zktypes.Eip712Meta{
				GasPerPubdata:   utils.NewBig(utils.DefaultGasPerPubdataLimit.Int64()),
				PaymasterParams: c.PaymasterParams(),
			},
		},
	)
	if err != nil {
		return "", err
	}
	preparedTx, err := c.PopulateTransaction(
		context.Background(),
		pbkHex,
		accounts.Transaction{
			GasFeeCap: gasPrice,
			GasTipCap: gasPrice,
			Gas:       gasNumber,
			To:        &contractAddr,
			Value:     big.NewInt(0),
			Data:      dataBytes,
			Meta: &zktypes.Eip712Meta{
				GasPerPubdata:   utils.NewBig(utils.DefaultGasPerPubdataLimit.Int64()),
				PaymasterParams: c.PaymasterParams(),
			},
		},
	)
	if err != nil {
		return "", err
	}
	chainID, err := c.GetChainID()
	if err != nil {
		return "", err
	}
	prkBytes, err := hex.DecodeString(prkHex)
	if err != nil {
		return "", err
	}
	baseSigner, err := accounts.NewBaseSignerFromRawPrivateKey(prkBytes, int64(chainID))
	if err != nil {
		return "", err
	}
	signer := accounts.Signer(baseSigner)
	rawTx, err := c.SignTransaction(signer, preparedTx)
	if err != nil {
		return "", err
	}
	hash, err := client.SendRawTransaction(context.Background(), rawTx)
	if err != nil {
		return "", err
	}
	// _, err = client.WaitMined(context.Background(), hash)
	// if err != nil {
	// 	return "", err
	// }
	return hash.Hex(), nil
}
