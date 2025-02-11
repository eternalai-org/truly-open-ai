package setting

import (
	"agent-battle/internal/core/port"
	"agent-battle/internal/core/service/game_usecase"
	"agent-battle/internal/core/worker/game"
	"agent-battle/pkg/drivers/mongodb"

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
