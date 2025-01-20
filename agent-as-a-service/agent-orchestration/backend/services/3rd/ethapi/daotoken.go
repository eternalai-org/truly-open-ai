package ethapi

import (
	"math/big"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/daotoken"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (c *Client) DeployDAOTToken(prkHex string, name string, symbol string) (string, string, error) {
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
	address, tx, _, err := daotoken.DeployDAOToken(auth, client, name, symbol)
	if err != nil {
		return "", "", err
	}
	return address.Hex(), tx.Hash().Hex(), nil
}
