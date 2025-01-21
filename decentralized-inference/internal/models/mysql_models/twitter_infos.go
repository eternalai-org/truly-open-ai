package mysql_models

import (
	"gorm.io/gorm"
)

type TwitterInfos struct {
	gorm.Model

	TwitterID                  string `json:"twitter_id"`
	TwitterName                string `json:"twitter_name"`
	TwitterUserName            string `gorm:"column:twitter_username" json:"twitter_user_name"`
	TwitterAvatar              string `json:"twitter_avatar"`
	TwitterDescription         string `json:"twitter_description"`
	TwitterDescriptionEntities string `json:"twitter_description_entities"`
}

func (TwitterInfos) TableName() string {
	return "twitter_infos"
}
