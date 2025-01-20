package trxapi

import (
	"context"
	"fmt"
	"math/big"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/erc20"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
)

func (c *Client) Trc20Balance(erc20Addr string, addr string) (*big.Int, error) {
	c.Conn()
	rs, err := c.conn.TRC20ContractBalance(addr, erc20Addr)
	if err != nil {
		if err.Error() == "account not found" {
			return nil, nil
		}
		return nil, err
	}
	return rs, nil
}

func (c *Client) TRC20ApproveMax(contractAddr string, prkHex string, toAddr string) (string, error) {
	c.Conn()
	toAddr = AddrTronToEvm(toAddr)
	privateKeyECDSA, err := crypto.HexToECDSA(prkHex)
	if err != nil {
		return "", err
	}
	addr := address.PubkeyToAddress(privateKeyECDSA.PublicKey)
	instanceABI, err := erc20.Erc20MetaData.GetAbi()
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
	gasPrice, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	gasFee := common.Big0.Mul(gasPrice, big.NewInt(200000)).Int64()
	err = c.CheckBalance(addr.String(), big.NewInt(gasFee))
	if err != nil {
		return "", err
	}
	tx, err := c.conn.TRC20Call(addr.String(), contractAddr, common.Bytes2Hex(dataBytes), false, gasFee)
	if err != nil {
		return "", err
	}
	signedTx, err := c.SignTx(privateKeyECDSA, tx.Transaction)
	if err != nil {
		return "", err
	}
	result, err := c.conn.Broadcast(signedTx)
	if err != nil {
		return "", err
	}
	if result.Code != 0 {
		return "", fmt.Errorf("bad transaction: %v", string(result.GetMessage()))
	}
	return common.BytesToHash(tx.GetTxid()).Hex(), nil
}
