package mysql_models

import (
	"gorm.io/gorm"
)

type BuyOrderStatus string

const (
	BuyOrderStatusNew          BuyOrderStatus = "new"
	BuyOrderStatusTransferring BuyOrderStatus = "transferring"
	BuyOrderStatusBuying       BuyOrderStatus = "buying"
	BuyOrderStatusDone         BuyOrderStatus = "done"
	BuyOrderStatusError        BuyOrderStatus = "error"
	BuyOrderStatusReserved     BuyOrderStatus = "reserved"
)

type BuyOrder struct {
	gorm.Model
	BuyWalletID     uint   `gorm:"index"`
	Address         string `gorm:"index"`
	EthAddress      string
	EthAmount       float64 `gorm:"type:decimal(36,18);default:0"`
	EthGasAmount    float64 `gorm:"type:decimal(36,18);default:0"`
	EaiAmount       float64 `gorm:"type:decimal(36,18);default:0"`
	TransferEthHash string
	BuyEaiHash      string
	TransferEaiHash string
	Status          BuyOrderStatus
	ChainId         int64 `gorm:"default:0"`
}
