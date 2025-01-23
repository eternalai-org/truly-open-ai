package btcapi

import (
	"errors"
	"fmt"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
)

type RawaddrTXResp struct {
	Hash   string `json:"hash"`
	Result int64  `json:"result"`
	Fee    uint64 `json:"fee"`
}

type RawaddrResp struct {
	NTx uint64           `json:"n_tx"`
	Txs []*RawaddrTXResp `json:"txs"`
}

func (c *Client) GetBTCAddressAllTxs(address string) ([]*RawaddrTXResp, error) {
	rs := []*RawaddrTXResp{}
	offset := 0
	for {
		time.Sleep(11 * time.Second)
		txs, err := c.GetBTCAddressTxs(address, offset, 50)
		if err != nil {
			return nil, errs.NewError(err)
		}
		rs = append(rs, txs...)
		if len(txs) < 50 {
			break
		}
		offset += 50
	}
	return rs, nil
}

func (c *Client) GetBTCAddressTxs(address string, offset int, limit int) ([]*RawaddrTXResp, error) {
	rs := RawaddrResp{}
	_, err := c.getJSON(
		fmt.Sprintf("%s/rawaddr/%s?offset=%d&limit=%d", c.BlockchainInfoUrl, address, offset, limit),
		map[string]string{},
		&rs,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return rs.Txs, nil
}

type MempoolTXResp struct {
	Txid string `json:"txid"`
	Fee  uint64 `json:"fee"`
	Vin  []*struct {
		Prevout *struct {
			ScriptpubkeyAddress string `json:"scriptpubkey_address"`
			Value               uint64 `json:"value"`
		} `json:"prevout"`
	} `json:"vin"`
}

func (c *Client) GetBTCAddressTxsV2(address string, lastTxID string) ([]*MempoolTXResp, error) {
	rs := []*MempoolTXResp{}
	for {
		rs1 := []*MempoolTXResp{}
		_, err := c.getJSON(
			fmt.Sprintf("https://mempool.space/api/address/%s/txs/chain%s", address, fmt.Sprintf("/%s", lastTxID)),
			map[string]string{},
			&rs1,
		)
		if err != nil {
			return nil, errs.NewError(err)
		}
		time.Sleep(time.Duration(0.5 * float64(time.Second)))
		rs = append(rs, rs1...)
		if len(rs1) < 25 || len(rs) >= 500 {
			break
		}
		lastTxID = rs1[len(rs1)-1].Txid
	}
	return rs, nil
}

func (c *Client) GetLatestBlockHash() (string, int64, error) {
	var rs struct {
		Hash       string `json:"hash"`
		BlockIndex int64  `json:"block_index"`
	}
	_, err := c.getJSON(
		fmt.Sprintf("%s/latestblock", c.BlockchainInfoUrl),
		map[string]string{},
		&rs,
	)
	if err != nil {
		return "", 0, errs.NewError(err)
	}
	return rs.Hash, rs.BlockIndex, nil
}

type BlockInfoResp struct {
	Hash string `json:"hash"`
	Tx   []*struct {
		Hash   string `json:"hash"`
		Inputs []*struct {
			PrevOut *struct {
				Spent bool   `json:"spent"`
				Value uint64 `json:"value"`
				Addr  string `json:"addr"`
			} `json:"prev_out"`
		} `json:"inputs"`
		Out []*struct {
			Spent bool   `json:"spent"`
			Value uint64 `json:"value"`
			Addr  string `json:"addr"`
		} `json:"out"`
	} `json:"tx"`
}

type OutResp struct {
	TxHash      string
	Index       uint
	FromAddress string
	ToAddrress  string
	Amount      uint64
}

func (c *Client) GetOutAddrsByHeight(height int64) ([]string, []*OutResp, error) {
	var rs struct {
		Blocks []*BlockInfoResp `json:"blocks"`
	}
	_, err := c.getJSON(
		fmt.Sprintf("%s/block-height/%d", c.BlockchainInfoUrl, height),
		map[string]string{},
		&rs,
	)
	if err != nil {
		return nil, nil, errs.NewError(err)
	}
	if len(rs.Blocks) <= 0 {
		return nil, nil, errs.NewError(errors.New("not found"))
	}
	addrMap := map[string]bool{}
	outArr := []*OutResp{}
	for _, r := range rs.Blocks {
		for _, tx := range r.Tx {
			var fromAddress string
			for _, in := range tx.Inputs {
				if in.PrevOut != nil && in.PrevOut.Addr != "" {
					fromAddress = in.PrevOut.Addr
				}
			}
			for idx, out := range tx.Out {
				if out.Addr != "" {
					_, ok := addrMap[out.Addr]
					if !ok {
						addrMap[out.Addr] = true
					}
					outArr = append(outArr, &OutResp{
						TxHash:      tx.Hash,
						Index:       uint(idx),
						FromAddress: fromAddress,
						ToAddrress:  out.Addr,
						Amount:      out.Value,
					})
				}
			}
		}
	}
	addrs := []string{}
	for addr := range addrMap {
		addrs = append(addrs, addr)
	}
	return addrs, outArr, nil
}
