package models

import (
	"github.com/jinzhu/gorm"
)

type Launchpad struct {
	gorm.Model
	TwitterPostID   uint   `gorm:"unique_index"`
	TweetId         string `gorm:"unique_index"`
	Name            string
	Description     string `gorm:"type:text"`
	TwitterId       string
	TwitterUsername string
	TwitterName     string
	Address         string
}

type LaunchpadMember struct {
	gorm.Model
	UserAddress string
	LaunchpadID uint
	Launchpad   *Launchpad
}
