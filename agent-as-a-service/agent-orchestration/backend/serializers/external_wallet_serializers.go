package serializers

import (
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
)

type ExternalWalletResp struct {
	ApiKey  string `json:"api_key"`
	Address string `json:"address"`
}

type ExternalWalletOrderReq struct {
	Action      models.ExternalWalletOrderType `json:"action"`
	Mint        string                         `json:"mint"`
	Amount      float64                        `json:"amount"`
	Destination string                         `json:"destination"`
}

type ExternalWalletOrderResp struct {
	ID               uint                             `json:"id"`
	CreatedAt        time.Time                        `json:"created_at"`
	UpdatedAt        time.Time                        `json:"updated_at"`
	ExternalWalletID uint                             `json:"external_wallet_id"`
	Type             models.ExternalWalletOrderType   `json:"type"`
	TokenAddress     string                           `json:"token_address"`
	Destination      string                           `json:"destination"`
	AmountIn         numeric.BigFloat                 `json:"amount_in"`
	AmountOut        numeric.BigFloat                 `json:"amount_out"`
	TxHash           string                           `json:"tx_hash"`
	Status           models.ExternalWalletOrderStatus `json:"status"`
	Error            string                           `json:"error"`
}

func NewExternalWalletOrderResp(m *models.ExternalWalletOrder) *ExternalWalletOrderResp {
	if m == nil {
		return nil
	}
	resp := &ExternalWalletOrderResp{
		ID:               m.ID,
		CreatedAt:        m.CreatedAt,
		UpdatedAt:        m.UpdatedAt,
		ExternalWalletID: m.ExternalWalletID,
		Type:             m.Type,
		TokenAddress:     m.TokenAddress,
		Destination:      m.Destination,
		AmountIn:         m.AmountIn,
		AmountOut:        m.AmountOut,
		TxHash:           m.TxHash,
		Status:           m.Status,
		Error:            m.Error,
	}
	return resp
}

func NewExternalWalletOrderRespArr(arr []*models.ExternalWalletOrder) []*ExternalWalletOrderResp {
	resps := []*ExternalWalletOrderResp{}
	for _, m := range arr {
		resps = append(resps, NewExternalWalletOrderResp(m))
	}
	return resps
}

type ExternalWalletTokenResp struct {
	ID            uint      `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Symbol        string    `json:"symbol"`
	Name          string    `json:"name"`
	TokenAddress  string    `json:"token_address"`
	Enabled       bool      `json:"enabled"`
	Decimals      int       `json:"decimals"`
	CoingeckoSlug string    `json:"coingecko_slug"`
}

func NewExternalWalletTokenResp(m *models.ExternalWalletToken) *ExternalWalletTokenResp {
	if m == nil {
		return nil
	}
	resp := &ExternalWalletTokenResp{
		ID:            m.ID,
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
		Symbol:        m.Symbol,
		Name:          m.Name,
		TokenAddress:  m.TokenAddress,
		Enabled:       m.Enabled,
		Decimals:      m.Decimals,
		CoingeckoSlug: m.CoingeckoSlug,
	}
	return resp
}

func NewExternalWalletTokenRespArr(arr []*models.ExternalWalletToken) []*ExternalWalletTokenResp {
	resps := []*ExternalWalletTokenResp{}
	for _, m := range arr {
		resps = append(resps, NewExternalWalletTokenResp(m))
	}
	return resps
}
