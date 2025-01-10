package trxapi

import (
	"context"
	"fmt"

	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
)

func (c *Client) triggerConstantContract(ct *core.TriggerSmartContract) (*api.TransactionExtention, error) {
	c.Conn()
	return c.conn.Client.TriggerConstantContract(context.Background(), ct)
}

func (c *Client) triggerContract(ct *core.TriggerSmartContract, feeLimit int64) (*api.TransactionExtention, error) {
	c.Conn()
	tx, err := c.conn.Client.TriggerContract(context.Background(), ct)
	if err != nil {
		return nil, err
	}
	if tx.Result.Code > 0 {
		return nil, fmt.Errorf("%s", string(tx.Result.Message))
	}
	if feeLimit > 0 {
		tx.Transaction.RawData.FeeLimit = feeLimit
		// update hash
		c.conn.UpdateHash(tx)
	}
	return tx, err
}

func (c *Client) TRC20Call(from, contractAddress, data string, constant bool, value, feeLimit int64) (*api.TransactionExtention, error) {
	c.Conn()
	var err error
	fromDesc := address.HexToAddress("410000000000000000000000000000000000000000")
	if len(from) > 0 {
		fromDesc, err = address.Base58ToAddress(from)
		if err != nil {
			return nil, err
		}
	}
	contractDesc, err := address.Base58ToAddress(contractAddress)
	if err != nil {
		return nil, err
	}
	dataBytes, err := common.FromHex(data)
	if err != nil {
		return nil, err
	}
	ct := &core.TriggerSmartContract{
		OwnerAddress:    fromDesc.Bytes(),
		ContractAddress: contractDesc.Bytes(),
		Data:            dataBytes,
		CallValue:       value,
	}
	var result *api.TransactionExtention
	if constant {
		result, err = c.triggerConstantContract(ct)
	} else {
		result, err = c.triggerContract(ct, feeLimit)
	}
	if err != nil {
		return nil, err
	}
	if result.Result.Code > 0 {
		return result, fmt.Errorf(string(result.Result.Message))
	}
	return result, nil

}
