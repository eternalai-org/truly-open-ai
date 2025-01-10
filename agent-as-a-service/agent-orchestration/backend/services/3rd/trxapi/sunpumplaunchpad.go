package trxapi

import (
	"context"
	"fmt"
	"math/big"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/sunpumplaunchpad"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
)

func (c *Client) SunpumpLaunchpadCreateAndInitPurchase(contractAddr string, prkHex string, symbol string, name string) (string, error) {
	c.Conn()
	privateKeyECDSA, err := crypto.HexToECDSA(prkHex)
	if err != nil {
		return "", err
	}
	addr := address.PubkeyToAddress(privateKeyECDSA.PublicKey)
	instanceABI, err := sunpumplaunchpad.SunpumplaunchpadMetaData.GetAbi()
	if err != nil {
		return "", err
	}
	dataBytes, err := instanceABI.Pack(
		"createAndInitPurchase", name, symbol,
	)
	if err != nil {
		return "", err
	}
	gasPrice, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	gasFee := common.Big0.Mul(gasPrice, big.NewInt(100000)).Int64()
	err = c.CheckBalance(addr.String(), big.NewInt(gasFee))
	if err != nil {
		return "", err
	}
	tx, err := c.TRC20Call(addr.String(), contractAddr, common.Bytes2Hex(dataBytes), false, 20000000, gasFee)
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
