package mongo

import (
	"agent-battle/internal/core/model"
	"agent-battle/pkg/drivers/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
)

type ISettingRepo interface {
	mongodb.Repository[*model.Setting]
}

type settingRepo struct {
	mongodb.Repository[*model.Setting]
}

func NewSettingRepo(db *mongo.Database, secondaryDb ...*mongo.Database) ISettingRepo {
	return &settingRepo{
		Repository: mongodb.NewBaseRepository(&model.Setting{}, db, secondaryDb...),
	}
}
