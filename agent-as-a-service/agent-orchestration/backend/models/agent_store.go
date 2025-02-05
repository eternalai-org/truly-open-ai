package models

import (
	"github.com/jinzhu/gorm"
)

type AgentStore struct {
	gorm.Model
	Name          string
	Description   string `gorm:"type:text"`
	OwnerAddress  string
	AuthenUrl     string `gorm:"type:longtext"`
	MissionStores []*MissionStore
}

type AgentStoreMission struct {
	gorm.Model
	AgentStoreID uint `gorm:"index"`
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
}

type AgentStoreInstall struct {
	gorm.Model
	AgentStoreID   uint   `gorm:"index"`
	AgentInfoID    uint   `gorm:"index"`
	CallbackParams string `gorm:"type:longtext"` //{"user_id" : "123", "authen_token" : "xxx",...}
}
