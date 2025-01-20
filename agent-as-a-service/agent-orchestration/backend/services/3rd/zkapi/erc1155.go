package zkapi

import (
	"context"
	"encoding/hex"
	"errors"
	"math/big"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/erc1155"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zksync-sdk/zksync2-go/accounts"
	zktypes "github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
)

func (c *Client) Erc1155SetApprovalForAll(erc1155Addr string, prkHex string, toAddr string, wait bool) (string, error) {
	if erc1155Addr == "" ||
		!common.IsHexAddress(erc1155Addr) {
		return "", errors.New("erc1155Addr is invalid")
	}
	if toAddr == "" ||
		!common.IsHexAddress(toAddr) {
		return "", errors.New("toAddr is invalid")
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
	contractAddress := helpers.HexToAddress(erc1155Addr)
	// EstimateGas
	instanceABI, err := erc1155.ERC1155MetaData.GetAbi()
	if err != nil {
		return "", err
	}
	dataBytes, err := instanceABI.Pack(
		"setApprovalForAll",
		helpers.HexToAddress(toAddr),
		true,
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
	if wait {
		_, err = client.WaitMined(context.Background(), hash)
		if err != nil {
			return "", err
		}
	}
	return hash.Hex(), nil
}

func (c *Client) Erc1155Transfer(erc20Addr string, prkHex string, toAddr string, tokenId string, amount *big.Int) (string, error) {
	if erc20Addr == "" ||
		!common.IsHexAddress(erc20Addr) {
		return "", errors.New("erc20Addr is invalid")
	}
	if toAddr == "" ||
		!common.IsHexAddress(toAddr) {
		return "", errors.New("toAddr is invalid")
	}
	tokenIdVal, ok := big.NewInt(0).SetString(tokenId, 10)
	if !ok {
		return "", errors.New("tokenId is invalid")
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
	contractAddress := helpers.HexToAddress(erc20Addr)
	// EstimateGas
	instanceABI, err := erc1155.ERC1155MetaData.GetAbi()
	if err != nil {
		return "", err
	}
	dataBytes, err := instanceABI.Pack(
		"safeTransferFrom",
		pbkHex,
		helpers.HexToAddress(toAddr),
		tokenIdVal,
		amount,
		[]byte{},
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
	// _, err = client.WaitMined(context.Background(), hash)
	// if err != nil {
	// 	return "", err
	// }
	return hash.Hex(), nil
}
