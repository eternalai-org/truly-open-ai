package game

import (
	"context"
	"time"

	"agent-battle/internal/adapters/repository/mongo"
	"agent-battle/internal/core/model"
	"agent-battle/internal/core/port"
	"agent-battle/pkg/logger"
	"agent-battle/pkg/telegram"

	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const WatchGameState string = "WatchGameState"

type gameWorker struct {
	gameUsecase    port.IGameUsecase
	workerRepo     mongo.IWorkerRepo
	telegramClient telegram.ITelegramClient
}

func (w *gameWorker) Watch(ctx context.Context, actionId string) error {
	logger.GetLoggerInstanceFromContext(ctx).Info("watch_game_start", zap.Any("action_id", actionId))

	worker := &model.Worker{}
	defer func() {
		if err := recover(); err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Error("watch error", zap.Any("error", err))
		}

		worker.State.IsRunning = false
		_ = w.updateWorker(ctx, worker)
		w.Watch(ctx, actionId)
		time.Sleep(time.Duration(3) * time.Second)
	}()

	for {
		worker = &model.Worker{}
		if err := w.workerRepo.FindOne(ctx, bson.M{"name": actionId}, worker); err != nil {
			return err
		}

		if worker.Interval == 0 {
			worker.Interval = 3
		}

		time.Sleep(time.Duration(worker.Interval) * time.Second)
		if worker.State == nil {
			worker.State = &model.State{}
		}

		if worker.Status == model.WorkerStatusDisable ||
			(worker.State.IsRunning && worker.State.LastRunDatetime.Add(1*time.Minute).After(time.Now())) {
			continue
		}

		worker.State.IsRunning = true
		if err := w.updateWorker(ctx, worker); err != nil {
			continue
		}

		if actionId == WatchGameState {
			logger.GetLoggerInstanceFromContext(ctx).Info("START", zap.Any("actionId", actionId))
			if err := w.gameUsecase.WatchGameState(ctx); err != nil {
				logger.GetLoggerInstanceFromContext(ctx).Error("WatchGameState", zap.Error(err))
			}

			worker.State.IsRunning = false
			worker.State.LastRunDatetime = time.Now()
			_ = w.updateWorker(ctx, worker)

			logger.GetLoggerInstanceFromContext(ctx).Info("DONE")
			continue
		}
	}
}

func (w *gameWorker) updateWorker(ctx context.Context, worker *model.Worker) error {
	lastUpdateAt := worker.DateModified
	worker.DateModified = time.Now()
	return w.workerRepo.Update(ctx, worker, worker.Id, lastUpdateAt)
}

func (w *gameWorker) FindWorkerByName(ctx context.Context, name string) (*model.Worker, error) {
	return w.workerRepo.FindOrCreateWorkerByName(ctx, name)
}

func NewGameWorker(
	gameUsecase port.IGameUsecase,
	workerRepo mongo.IWorkerRepo,
) port.IGameWorker {
	telegramClient := telegram.New()
	return &gameWorker{
		gameUsecase:    gameUsecase,
		telegramClient: telegramClient,
		workerRepo:     workerRepo,
	}
}

var Module = fx.Module("game_worker", mongo.WorkerRepoModule, fx.Provide(NewGameWorker))
