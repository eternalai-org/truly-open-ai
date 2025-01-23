package zkapi

import (
	"context"
	"encoding/hex"
	"errors"
	"math/big"
	"strings"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/erc721"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zksync-sdk/zksync2-go/accounts"
	zktypes "github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
)

func (c *Client) Erc721Transfer(contractAddr string, prkHex string, toAddr string, tokenId *big.Int) (string, error) {
	if contractAddr == "" ||
		!common.IsHexAddress(contractAddr) {
		return "", errors.New("erc20Addr is invalid")
	}
	if toAddr == "" ||
		!common.IsHexAddress(toAddr) {
		return "", errors.New("toAddr is invalid")
	}
	pbkHex, _, err := c.parsePrkAuth(prkHex)
	if err != nil {
		return "", err
	}
	gasPrice, _, err := c.GetCachedGasPriceAndTipCap()
	if err != nil {
		return "", err
	}
	client, err := c.getZkClient()
	if err != nil {
		return "", err
	}
	contractAddress := helpers.HexToAddress(contractAddr)
	instanceABI, err := abi.JSON(strings.NewReader(erc721.Erc721ABI))
	if err != nil {
		return "", err
	}
	dataBytes, err := instanceABI.Pack(
		"safeTransferFrom",
		pbkHex,
		helpers.HexToAddress(toAddr),
		tokenId,
	)
	if err != nil {
		return "", err
	}
	gasNumber, err := client.EstimateGasL2(
		context.Background(),
		zktypes.CallMsg{
			CallMsg: ethereum.CallMsg{
				From: pbkHex,
				To:   &contractAddress,
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
			To:        &contractAddress,
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
	return hash.Hex(), nil
}
