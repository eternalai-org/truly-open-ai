package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"agent-battle/cmd/api"
	"agent-battle/cmd/setting"
	"agent-battle/cmd/worker"

	"agent-battle/pkg/logger"
	"agent-battle/pkg/utils"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"

	"go.uber.org/fx"
)

var (
	envFlag    string
	configFile string
)

// @title Agent Battle API
// @version 1.0.0
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// init flag
	flag.StringVar(&envFlag, "env", "development", "Config env: development, production")
	flag.StringVar(&configFile, "config-file", "env/development.yml", "Config file path")
	flag.Parse()

	if utils.Environment(envFlag) == utils.Production {
		viper.AutomaticEnv()
	} else {
		viper.SetConfigFile(configFile)
		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}
	}

	// init logger
	if viper.GetString("LOGGER_FORMAT") == "json" {
		logger.InitLoggerDefault(viper.GetBool("LOGGER_ENABLE_DEBUG"))
	} else {
		logger.InitLoggerDefaultDev()
	}
	logger.AtLog.Info("Server is starting...")

	httpServer := fx.New(
		setting.Module,
		api.Module,
		worker.Module,
		fx.Invoke(NewApp),
	)
	httpServer.Run()

	defer func() {
		logger.AtLog.Info("Server exiting")
	}()
}

func NewApp(lc fx.Lifecycle, apiServer *api.APIServer, cronServer *worker.CronServer, mongoDb *mongo.Database) {
	app := apiServer.App
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := app.Listen(fmt.Sprintf(":%d", viper.GetInt(`SERVER_HTTP_PORT`))); err != nil && err != http.ErrServerClosed {
					logger.AtLog.Fatalf("listen error: %v", err)
				}
			}()
			if utils.IsWorker() {
				cronServer.App.Start()
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.AtLog.Info("Shutting down server...")
			if err := app.ShutdownWithTimeout(30 * time.Second); err != nil {
				logger.AtLog.Fatalf("Server forced to shutdown: %v", err)
			}
			if utils.IsWorker() {
				cronServer.App.Stop()
			}

			_ = mongoDb.Client().Disconnect(context.Background())
			return nil
		},
	})
}
