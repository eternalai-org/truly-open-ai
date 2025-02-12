package port

import (
	"context"

	"agent-battle/internal/core/model"
)

type IGameWorker interface {
	Watch(context.Context, string) error
	FindWorkerByName(context.Context, string) (*model.Worker, error)
}

type IGameUsecase interface {
	StartGame(context.Context, *model.StartGameRequest) (*model.Game, error)
	EndGame(context.Context, string) (*model.Game, error)
	DetailGame(context.Context, string) (*model.Game, error)
	ListGame(context.Context, *model.ListGameRequest) (*model.ListGameResponse, error)
	GameResult(context.Context, *model.GameResultRequest) (*model.Game, error)
	WatchGameState(context.Context) error
	RefundsExpiredPlayers(context.Context, string) error
}
