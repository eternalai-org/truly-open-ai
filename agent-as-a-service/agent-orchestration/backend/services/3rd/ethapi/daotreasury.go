package ethapi

import (
	"math/big"
	"strings"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/daotreasury"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (c *Client) DeployDAOTreasury(prkHex string) (string, string, error) {
	_, prk, err := c.parsePrkAuth(prkHex)
	if err != nil {
		return "", "", err
	}
	chainID, err := c.GetChainID()
	if err != nil {
		return "", "", err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(prk, big.NewInt(int64(chainID)))
	if err != nil {
		return "", "", err
	}
	client, err := c.getClient()
	if err != nil {
		return "", "", err
	}
	gasPrice, err := c.getGasPrice()
	if err != nil {
		return "", "", err
	}
	auth.GasPrice = gasPrice
	address, tx, _, err := daotreasury.DeployDAOTreasury(auth, client)
	if err != nil {
		return "", "", err
	}
	return address.Hex(), tx.Hash().Hex(), nil
}

func (c *Client) DAOTreasuryInitializeData(positionManager common.Address, baseToken common.Address) ([]byte, error) {
	instanceABI, err := abi.JSON(strings.NewReader(daotreasury.DAOTreasuryABI))
	if err != nil {
		return nil, err
	}
	dataBytes, err := instanceABI.Pack(
		"initialize",
		positionManager,
		baseToken,
	)
	if err != nil {
		return nil, err
	}
	return dataBytes, nil
}
