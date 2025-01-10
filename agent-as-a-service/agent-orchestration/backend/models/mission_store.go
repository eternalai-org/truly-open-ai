package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type (
	OutputType string
)

const (
	OutputTypeTwitter OutputType = "twitter"
	OutputTypeOthers  OutputType = "others"
)

type MissionStore struct {
	gorm.Model
	Name         string
	Description  string `gorm:"type:text"`
	UserPrompt   string `gorm:"type:longtext"`
	Price        uint
	OwnerAddress string
	ToolList     string  `gorm:"type:longtext"`
	Rating       float64 `gorm:"type:decimal(5,2);default:0"`
	NumRating    uint
	NumUsed      uint
	Icon         string `gorm:"type:text"`
	OutputType   OutputType
	Params       string `gorm:"type:longtext"` //[{"name" : "token", "type" : "text", "description" : "this token is used to authenticate in xxx server side"}]

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
