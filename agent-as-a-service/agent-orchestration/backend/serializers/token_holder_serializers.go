package serializers

import (
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
)

type TokenHolderResp struct {
	ContractAddress     string           `json:"token_address"`
	Address             string           `json:"user_address"`
	UserName            string           `json:"user_name"`
	UserImageURL        string           `json:"user_image_url"`
	Balance             string           `json:"balance"`
	MemeName            string           `json:"meme_name"`
	MemeTicker          string           `json:"meme_ticker"`
	MemeImage           string           `json:"meme_image"`
	MemePrice           numeric.BigFloat `json:"meme_price"`
	MemePriceUsd        numeric.BigFloat `json:"meme_price_usd"`
	MemeBaseTokenSymbol string           `json:"meme_base_token_symbol"`
}

func NewTokenHolderResp(m *models.Erc20Holder) *TokenHolderResp {
	if m == nil {
		return nil
	}
	resp := &TokenHolderResp{
		Address:             m.Address,
		Balance:             m.Balance,
		ContractAddress:     m.ContractAddress,
		UserName:            m.UserName,
		UserImageURL:        m.ImageURL,
		MemeName:            m.MemeName,
		MemeTicker:          m.MemeTicker,
		MemeImage:           m.MemeImage,
		MemePrice:           m.MemePrice,
		MemePriceUsd:        m.MemePriceUsd,
		MemeBaseTokenSymbol: m.MemeBaseTokenSymbol,
	}
	return resp
}

func NewTokenHolderRespArray(arr []*models.Erc20Holder) []*TokenHolderResp {
	resps := []*TokenHolderResp{}
	for _, m := range arr {
		resps = append(resps, NewTokenHolderResp(m))
	}
	return resps
}

type ShareHolderResp struct {
	NetworkID       uint64           `json:"network_id"`
	ContractAddress string           `json:"contract_address"`
	TokenID         string           `json:"token_id"`
	Address         string           `json:"address"`
	Balance         string           `json:"balance"`
	TotalBalance    numeric.BigFloat `json:"total_balance"`
	TwitterID       string           `json:"twitter_id"`
	TwitterName     string           `json:"twitter_name"`
	TwitterUsername string           `json:"twitter_username"`
	TwitterAvatar   string           `json:"twitter_avatar"`
	UserName        string           `json:"user_name"`
	Description     string           `json:"description"`
	ImageURL        string           `json:"image_url"`
}
