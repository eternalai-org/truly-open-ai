package zkapi

import (
	"context"
	"encoding/hex"
	"errors"
	"math/big"
	"strings"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/erc20"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/erc721"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zksync-sdk/zksync2-go/accounts"
	zktypes "github.com/zksync-sdk/zksync2-go/types"
	"github.com/zksync-sdk/zksync2-go/utils"
)

func (c *Client) Erc20Symbol(erc20Addr string) (string, error) {
	if !common.IsHexAddress(erc20Addr) {
		return "", errors.New("erc20Addr is invalid")
	}
	client, err := c.getZkClient()
	if err != nil {
		return "", err
	}
	instance, err := erc20.NewErc20(helpers.HexToAddress(erc20Addr), client)
	if err != nil {
		return "", err
	}
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		return "", err
	}
	return symbol, nil
}

type Erc20InfoResp struct {
	Symbol     string   `json:"symbol"`
	Name       string   `json:"name"`
	TotalSuply *big.Int `json:"total_suply"`
	Decimals   int      `json:"decimals"`
}

func (c *Client) Erc20Info(erc20Addr string) (*Erc20InfoResp, error) {
	if !common.IsHexAddress(erc20Addr) {
		return nil, errors.New("erc20Addr is invalid")
	}
	client, err := c.getZkClient()
	if err != nil {
		return nil, err
	}
	instance, err := erc20.NewErc20(helpers.HexToAddress(erc20Addr), client)
	if err != nil {
		return nil, err
	}
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	totalSupply, err := instance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	return &Erc20InfoResp{
		Symbol:     symbol,
		Name:       name,
		TotalSuply: totalSupply,
		Decimals:   int(decimals),
	}, nil
}

func (c *Client) Erc20Balance(erc20Addr string, addr string) (*big.Int, error) {
	if !common.IsHexAddress(erc20Addr) {
		return nil, errors.New("erc20Addr is invalid")
	}
	if !common.IsHexAddress(addr) {
		return nil, errors.New("addr is invalid")
	}
	client, err := c.getZkClient()
	if err != nil {
		return nil, err
	}
	instance, err := erc20.NewErc20(helpers.HexToAddress(erc20Addr), client)
	if err != nil {
		return nil, err
	}
	balance, err := instance.BalanceOf(&bind.CallOpts{}, helpers.HexToAddress(addr))
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (c *Client) NftOwnerOf(nftAddr string, tokenID string) (string, error) {
	client, err := c.getZkClient()
	if err != nil {
		return "", err
	}

	nftContract, err := erc721.NewErc721(helpers.HexToAddress(nftAddr), client)
	if err != nil {
		return "", err
	}
	tokenIDInt, ok := big.NewInt(0).SetString(tokenID, 10)
	if !ok {
		return "", errors.New("bad token id")
	}
	res, err := nftContract.OwnerOf(&bind.CallOpts{}, tokenIDInt)
	if err != nil {
		return "", err
	}
	return res.Hex(), nil
}

func (c *Client) Balance(addr string) (*big.Int, error) {
	if !common.IsHexAddress(addr) {
		return nil, errors.New("addr is invalid")
	}
	client, err := c.getZkClient()
	if err != nil {
		return nil, err
	}
	balance, err := client.BalanceAt(context.Background(), helpers.HexToAddress(addr), nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (c *Client) Erc20Allowance(erc20Addr string, addr string, spender string) (*big.Int, error) {
	if !common.IsHexAddress(erc20Addr) {
		return nil, errors.New("erc20Addr is invalid")
	}
	if !common.IsHexAddress(addr) {
		return nil, errors.New("addr is invalid")
	}
	if !common.IsHexAddress(spender) {
		return nil, errors.New("spender is invalid")
	}
	client, err := c.getZkClient()
	if err != nil {
		return nil, err
	}
	instance, err := erc20.NewErc20(helpers.HexToAddress(erc20Addr), client)
	if err != nil {
		return nil, err
	}
	balance, err := instance.Allowance(&bind.CallOpts{}, helpers.HexToAddress(addr), helpers.HexToAddress(spender))
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (c *Client) Erc20ApproveMax(erc20Addr string, prkHex string, toAddr string, wait bool) (string, error) {
	if erc20Addr == "" ||
		!common.IsHexAddress(erc20Addr) {
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
	instanceABI, err := abi.JSON(strings.NewReader(erc20.Erc20ABI))
	if err != nil {
		return "", err
	}
	amount, _ := new(big.Int).SetString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 16)
	dataBytes, err := instanceABI.Pack(
		"approve",
		helpers.HexToAddress(toAddr),
		amount,
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

func (c *Client) Erc20Transfer(erc20Addr string, prkHex string, toAddr string, amount *big.Int) (string, error) {
	if erc20Addr == "" ||
		!common.IsHexAddress(erc20Addr) {
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
	instanceABI, err := abi.JSON(strings.NewReader(erc20.Erc20ABI))
	if err != nil {
		return "", err
	}
	dataBytes, err := instanceABI.Pack(
		"transfer",
		helpers.HexToAddress(toAddr),
		amount,
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
