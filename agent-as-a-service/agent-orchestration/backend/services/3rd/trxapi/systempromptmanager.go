package trxapi

import (
	"context"
	"fmt"
	"math/big"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/systempromptmanager"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
)

func (c *Client) SystemPromptManagerMint(contractAddr string, prkHex string, to common.Address, uri string, data []byte, fee *big.Int) (string, error) {
	c.Conn()
	privateKeyECDSA, err := crypto.HexToECDSA(prkHex)
	if err != nil {
		return "", err
	}
	addr := address.PubkeyToAddress(privateKeyECDSA.PublicKey)
	instanceABI, err := systempromptmanager.SystemPromptManagerMetaData.GetAbi()
	if err != nil {
		return "", err
	}
	dataBytes, err := instanceABI.Pack(
		"mint", to, uri, data, fee,
	)
	if err != nil {
		return "", err
	}
	gasPrice, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	gasFee := common.Big0.Mul(gasPrice, big.NewInt(500000)).Int64()
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
