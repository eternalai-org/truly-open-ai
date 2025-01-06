package ethapi

import (
	"errors"
	"math/big"
	"strings"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/iagenttokendeployer"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (c *Client) IAGENTTokenDeployerGetToken(contractAddress string, salt string) (string, error) {
	if !common.IsHexAddress(contractAddress) {
		return "", errors.New("contractAddress is invalid")
	}
	client, err := c.getClient()
	if err != nil {
		return "", err
	}
	instance, err := iagenttokendeployer.NewIAGENTTokenDeployer(helpers.HexToAddress(contractAddress), client)
	if err != nil {
		return "", err
	}
	tokenAddress, err := instance.GetToken(&bind.CallOpts{}, [32]byte(common.HexToHash(salt)))
	if err != nil {
		return "", err
	}
	return tokenAddress.Hex(), nil
}

func (c *Client) IAGENTTokenDeployerCreateTokenData(salt string, name string, symbol string, supply *big.Int) ([]byte, error) {
	instanceABI, err := abi.JSON(strings.NewReader(iagenttokendeployer.IAGENTTokenDeployerABI))
	if err != nil {
		return nil, err
	}
	dataBytes, err := instanceABI.Pack(
		"createToken",
		[32]byte(common.HexToHash(salt)),
		name,
		symbol,
		supply,
	)
	if err != nil {
		return nil, err
	}
	return dataBytes, nil
}

func (c *Client) IAGENTTokenDeployerCallWithValueData(target string, data []byte, value *big.Int) ([]byte, error) {
	instanceABI, err := abi.JSON(strings.NewReader(iagenttokendeployer.IAGENTTokenDeployerABI))
	if err != nil {
		return nil, err
	}
	dataBytes, err := instanceABI.Pack(
		"callWithValue",
		helpers.HexToAddress(target),
		data,
		value,
	)
	if err != nil {
		return nil, err
	}
	return dataBytes, nil
}

func (c *Client) IAGENTTokenDeployerMulticallData(data [][]byte) ([]byte, error) {
	instanceABI, err := abi.JSON(strings.NewReader(iagenttokendeployer.IAGENTTokenDeployerABI))
	if err != nil {
		return nil, err
	}
	dataBytes, err := instanceABI.Pack(
		"multicall",
		data,
	)
	if err != nil {
		return nil, err
	}
	return dataBytes, nil
}
