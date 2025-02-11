package worker

import (
	"context"

	"agent-battle/cmd/setting"
	"agent-battle/internal/core/model"
	"agent-battle/internal/core/worker/game"
	"agent-battle/pkg/logger"
	"agent-battle/pkg/utils"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type CronServer struct {
	App *cron.Cron
}

func Init(s *setting.Setting) *CronServer {
	cr := cron.New(cron.WithLogger(cron.VerbosePrintfLogger(logger.AtLog)))
	if !utils.IsWorker() {
		return &CronServer{App: cr}
	}

	ctx := context.Background()
	if worker, err := s.GameWorker.FindWorkerByName(ctx, game.WatchGameState); err == nil {
		logger.AtLog.Infof("WatchGameState: %s", worker.Crontab)
		if worker.Crontab == "" || worker.Status == model.WorkerStatusDisable {
			return nil
		}

		crontab := worker.Crontab
		if _, err := cr.AddFunc(crontab, func() {
			if err := s.GameWorker.Watch(ctx, game.WatchGameState); err != nil {
				logger.GetLoggerInstanceFromContext(ctx).Error("ExecFunc#WatchDepositWalletAmountChange", zap.Error(err))
			}
		}); err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Error("AddFunc#WatchDepositWalletAmountChange", zap.Error(err))
		}
	}

	return &CronServer{App: cr}
}
