package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type MissionStore struct {
	gorm.Model
	Name           string
	Description    string `gorm:"type:text"`
	UserPrompt     string `gorm:"type:longtext"`
	Price          uint
	OwnerAddress   string
	ToolList       string `gorm:"type:longtext"`
	DepositAddress string `gorm:"index"`
	DurationDay    uint
	Rating         float64 `gorm:"type:decimal(5,2);default:0"`
	NumRating      uint
	NumUsed        uint
}

type MissionStoreHistory struct {
	gorm.Model
	UserAddress    string
	MissionStoreID uint
	MissionStore   *MissionStore
	TxHash         string
	EventId        string `gorm:"unique_index"`
	StartedAt      *time.Time
	ExpiresAt      *time.Time
	IsRated        bool
}

type MissionStoreRating struct {
	gorm.Model
	UserAddress           string
	MissionStoreID        uint `gorm:"unique_index:mission_store_rating_main_uidx"`
	MissionStore          *MissionStore
	MissionStoreHistoryID uint `gorm:"unique_index:mission_store_rating_main_uidx"`
	MissionStoreHistory   MissionStoreHistory
	Rating                float64 `gorm:"type:decimal(5,2);default:0"`
	Comment               string  `gorm:"type:longtext"`
}
