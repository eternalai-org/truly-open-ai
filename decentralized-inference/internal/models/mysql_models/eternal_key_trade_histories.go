package mysql_models

import (
	"time"

	"gorm.io/gorm"
)

type EternalKeyTradeHistory struct {
	gorm.Model
	ContractAddress   string    `gorm:"index"`
	EventId           string    `gorm:"unique_index"`
	TxAt              time.Time `gorm:"index"`
	TxHash            string    `gorm:"index"`
	UserAddress       string    `gorm:"index"`
	CreatorAddress    string    `gorm:"index"`
	EternalKeyAddress string    `gorm:"index"`
	IsBuy             bool
	ShareAmount       float64 `gorm:"type:decimal(60,18);default:0"`
	BuyAmount         float64 `gorm:"type:decimal(36,18);default:0"`
	ProtocolFeeAmount float64 `gorm:"type:decimal(36,18);default:0"`
	CreatorFeeAmount  float64 `gorm:"type:decimal(36,18);default:0"`
	HolderFeeAmount   float64 `gorm:"type:decimal(36,18);default:0"`
	Supply            float64 `gorm:"type:decimal(60,18);default:0"`
	BuyPrice          float64 `gorm:"type:decimal(36,18);default:0"`
	BaseTokenSymbol   string  `gorm:"default:'EAI'"`
	BaseTokenPrice    float64 `gorm:"type:decimal(36,18);default:0"`
}

func (EternalKeyTradeHistory) TableName() string {
	return "eternal_key_trade_histories"
}
