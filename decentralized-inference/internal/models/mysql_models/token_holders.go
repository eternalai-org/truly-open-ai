package mysql_models

import "gorm.io/gorm"

type TokenHolders struct {
	gorm.Model

	ContractAddress string `json:"contract_address"`
	Address         string `json:"address"`
	Balance         string `json:"balance"` // not 1e18
	LastBlockNumber int    `json:"last_block_number"`
}

func (TokenHolders) TableName() string {
	return "token_holders"
}
