package pumfun

import (
	"fmt"
	"net/http"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/helpers"
)

type Client struct {
	BaseUrl string
}

type PumpFunTradeResp struct {
	Signature   string `json:"signature"`
	Mint        string `json:"mint"`
	SolAmount   int64  `json:"sol_amount"`
	TokenAmount int64  `json:"token_amount"`
	IsBuy       bool   `json:"is_buy"`
	Timestamp   int64  `json:"timestamp"`
}

func (e *Client) GetPumpFunTrades(mint string, page int, limit int) ([]*PumpFunTradeResp, error) {
	var rs []*PumpFunTradeResp
	err := helpers.CurlURL(
		fmt.Sprintf(
			"%s/trades/all/%s?offset=%d&limit=%d&minimumSize=0",
			e.BaseUrl,
			mint,
			(page-1)*limit,
			limit,
		),
		http.MethodGet,
		make(map[string]string),
		nil,
		&rs,
	)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

type PumpFunCoinInfoResp struct {
	Mint         string  `json:"mint"`
	Name         string  `json:"name"`
	Symbol       string  `json:"symbol"`
	TotalSupply  int64   `json:"total_supply"`
	UsdMarketCap float64 `json:"usd_market_cap"`
	RaydiumPool  string  `json:"raydium_pool"`
}

func (e *Client) GetPumpFunCoinInfo(mint string) (*PumpFunCoinInfoResp, error) {
	var rs PumpFunCoinInfoResp
	err := helpers.CurlURL(
		fmt.Sprintf(
			"%s/coins/%s",
			e.BaseUrl,
			mint,
		),
		http.MethodGet,
		make(map[string]string),
		nil,
		&rs,
	)
	if err != nil {
		return nil, err
	}
	return &rs, nil
}
