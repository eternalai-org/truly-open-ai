package models

import (
	"github.com/jinzhu/gorm"
)

type Launchpad struct {
	gorm.Model
	Name         string
	Description  string `gorm:"type:text"`
	OwnerAddress string
	AdminAddress string
}

type LaunchpadMember struct {
	gorm.Model
	UserAddress string
	LaunchpadID uint
	Launchpad   *Launchpad
}
