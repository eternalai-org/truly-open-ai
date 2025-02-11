package mongo

import (
	"agent-battle/internal/core/model"
	"agent-battle/pkg/drivers/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
)

type IGameRepo interface {
	mongodb.Repository[*model.Game]
}

type gameRepo struct {
	mongodb.Repository[*model.Game]
}

func NewGameRepo(db *mongo.Database, secondaryDb ...*mongo.Database) IGameRepo {
	return &gameRepo{
		Repository: mongodb.NewBaseRepository(&model.Game{}, db, secondaryDb...),
	}
}
