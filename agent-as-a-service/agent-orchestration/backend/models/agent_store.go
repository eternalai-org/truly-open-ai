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
	AgentStoreID   uint `gorm:"index"`
	MissionStoreID uint `gorm:"index"`
}

type AgentStoreInstall struct {
	gorm.Model
	AgentStoreID   uint `gorm:"index"`
	UserAddress    string
	CallbackParams string `gorm:"type:longtext"` //{"user_id" : "123", "authen_token" : "xxx",...}
}
