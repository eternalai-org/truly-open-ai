package ethapi

import (
	"context"
	"errors"
	"math/big"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/erc1155"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/erc20"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/erc721"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (c *Client) Erc20Symbol(erc20Addr string) (string, error) {
	if !common.IsHexAddress(erc20Addr) {
		return "", errors.New("erc20Addr is invalid")
	}
	client, err := c.getClient()
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
	client, err := c.getClient()
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
	client, err := c.getClient()
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
	client, err := c.getClient()
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
	client, err := c.getClient()
	if err != nil {
		return nil, err
	}
	balance, err := client.BalanceAt(context.Background(), helpers.HexToAddress(addr), nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (c *Client) Erc1155Balance(erc20Addr string, tokenId string, addr string) (*big.Int, error) {
	if !common.IsHexAddress(erc20Addr) {
		return nil, errors.New("erc20Addr is invalid")
	}
	if !common.IsHexAddress(addr) {
		return nil, errors.New("addr is invalid")
	}
	tokenIdVal, ok := big.NewInt(0).SetString(tokenId, 10)
	if !ok {
		return nil, errors.New("tokenId is invalid")
	}
	client, err := c.getClient()
	if err != nil {
		return nil, err
	}
	instance, err := erc1155.NewERC1155(helpers.HexToAddress(erc20Addr), client)
	if err != nil {
		return nil, err
	}
	balance, err := instance.BalanceOf(&bind.CallOpts{}, helpers.HexToAddress(addr), tokenIdVal)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (c *Client) Erc1155IsApprovedForAll(erc20Addr string, addr string, operator string) (bool, error) {
	if !common.IsHexAddress(erc20Addr) {
		return false, errors.New("erc20Addr is invalid")
	}
	if !common.IsHexAddress(addr) {
		return false, errors.New("addr is invalid")
	}
	client, err := c.getClient()
	if err != nil {
		return false, err
	}
	instance, err := erc1155.NewERC1155(helpers.HexToAddress(erc20Addr), client)
	if err != nil {
		return false, err
	}
	balance, err := instance.IsApprovedForAll(&bind.CallOpts{}, helpers.HexToAddress(addr), helpers.HexToAddress(operator))
	if err != nil {
		return false, err
	}
	return balance, nil
}
