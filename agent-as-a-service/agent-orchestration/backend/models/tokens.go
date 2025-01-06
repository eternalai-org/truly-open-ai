package models

import (
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

type Erc20Holder struct {
	gorm.Model
	NetworkID           uint64           `gorm:"unique_index:token_holder_main_uidx"`
	ContractAddress     string           `gorm:"unique_index:token_holder_main_uidx"`
	Address             string           `gorm:"unique_index:token_holder_main_uidx"`
	Balance             string           `gorm:"default:0"`
	UserName            string           `gorm:"-"`
	ImageURL            string           `gorm:"-"`
	MemeName            string           `gorm:"-"`
	MemeTicker          string           `gorm:"-"`
	MemeImage           string           `gorm:"-"`
	MemePrice           numeric.BigFloat `gorm:"-"`
	MemePriceUsd        numeric.BigFloat `gorm:"-"`
	MemeBaseTokenSymbol string           `gorm:"-"`
}

type Erc721Holder struct {
	gorm.Model
	NetworkID       uint64 `gorm:"unique_index:nft_holders_main_idx"`
	ContractAddress string `gorm:"unique_index:nft_holders_main_idx"`
	TokenID         uint   `gorm:"unique_index:nft_holders_main_idx"`
	OwnerAddress    string `gorm:"index"`
}
type Erc1155Holder struct {
	gorm.Model
	NetworkID       uint64 `gorm:"unique_index:nft_holders_main_idx"`
	ContractAddress string `gorm:"unique_index:nft_holders_main_idx"`
	TokenID         uint   `gorm:"unique_index:nft_holders_main_idx"`
	Address         string `gorm:"unique_index:nft_holders_main_idx"`
	Balance         string `gorm:"default:0"`
}
