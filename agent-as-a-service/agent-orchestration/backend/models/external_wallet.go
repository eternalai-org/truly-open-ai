package models

import (
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

type ExternalWalletType string

const (
	ExternalWalletTypeSOL ExternalWalletType = "SOL"
)

type ExternalWallet struct {
	gorm.Model
	APIKey  string `gorm:"unique_index"`
	Type    ExternalWalletType
	Address string `gorm:"unique_index"`
}

type ExternalWalletOrderType string
type ExternalWalletOrderStatus string

const (
	ExternalWalletOrderTypeBuy      ExternalWalletOrderType = "buy"
	ExternalWalletOrderTypeSell     ExternalWalletOrderType = "sell"
	ExternalWalletOrderTypeWithdraw ExternalWalletOrderType = "withdraw"

	ExternalWalletOrderStatusNew   ExternalWalletOrderStatus = "new"
	ExternalWalletOrderStatusDone  ExternalWalletOrderStatus = "done"
	ExternalWalletOrderStatusError ExternalWalletOrderStatus = "error"
)

type ExternalWalletOrder struct {
	gorm.Model
	ExternalWalletID uint
	ExternalWallet   *ExternalWallet
	Type             ExternalWalletOrderType
	TokenAddress     string
	TokenName        string
	TokenSymbol      string
	Destination      string
	AmountIn         numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	AmountOut        numeric.BigFloat `gorm:"type:decimal(36,18);default:0"`
	TxHash           string
	Status           ExternalWalletOrderStatus
	Error            string `gorm:"type:longtext"`
}

type ExternalWalletToken struct {
	gorm.Model
	Symbol        string
	Name          string
	TokenAddress  string
	Price         float64 `gorm:"default:0"`
	Enabled       bool    `gorm:"default:0"`
	Decimals      int     `gorm:"default:0"`
	CoingeckoSlug string
}
