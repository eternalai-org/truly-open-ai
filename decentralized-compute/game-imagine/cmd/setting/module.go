package setting

import (
	"game-imagine/internal/core/port"
	"game-imagine/internal/core/service/game_usecase"
	"game-imagine/internal/core/worker/game"
	"game-imagine/pkg/drivers/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
)

var Module = fx.Module("setting",
	mongodb.Module,
	game_usecase.Module,
	game.Module,
	fx.Provide(
		func(
			db *mongo.Database,
			gameUsecase port.IGameUsecase,
			gameWorker port.IGameWorker,
		) *Setting {
			return Init(
				WithDB(db),
				WithGameUsecase(gameUsecase),
				WithGameWorker(gameWorker),
			)
		},
	),
)
