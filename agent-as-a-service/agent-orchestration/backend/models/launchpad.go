package models

import (
	"github.com/jinzhu/gorm"
)

type Launchpad struct {
	gorm.Model
	TweetId      string `gorm:"unique_index"`
	ReplyTweetId string
	Name         string
	Description  string `gorm:"type:text"`
	TwitterId    string
	Address      string
	ReplyContent string `gorm:"type:text"`
}

type LaunchpadMember struct {
	gorm.Model
	UserAddress string
	LaunchpadID uint
	Launchpad   *Launchpad
}
