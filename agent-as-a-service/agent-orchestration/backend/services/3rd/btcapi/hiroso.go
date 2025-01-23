package btcapi

import (
	"fmt"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
)

type WalletBalance struct {
	Results []*WalletInscription `json:"results"`
}

type WalletInscription struct {
	ID     string `json:"id"`
	TxId   string `json:"tx_id"`
	Output string `json:"output"`
	Value  string `json:"value"`
	Offset string `json:"offset"`
}

func (c *Client) GetBTCAddressInscriptionsAll(address string) ([]*WalletInscription, error) {
	rs := []*WalletInscription{}
	offset := uint(0)
	for i := 0; i < 25; i++ {
		rsT, err := c.GetBTCAddressInscriptions(address, offset, 20)
		if err != nil {
			return nil, errs.NewError(err)
		}
		rs = append(rs, rsT...)
		if len(rsT) < 20 {
			break
		}
		offset = offset + 20
		time.Sleep(100 * time.Millisecond)
	}
	return rs, nil
}

func (c *Client) GetBTCAddressInscriptions(address string, offset uint, limit uint) ([]*WalletInscription, error) {
	rs := WalletBalance{}
	_, err := c.getJSON(
		fmt.Sprintf("%s/ordinals/v1/inscriptions?address=%s&offset=%d&limit=%d", c.HirosoUrl, address, offset, limit),
		map[string]string{
			"x-hiro-api-key": "ce6ddb97db5c94f6a8bac16d8fb803cb",
		},
		&rs,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return rs.Results, nil
}
