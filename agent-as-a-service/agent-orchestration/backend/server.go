package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"strings"
	"syscall"
	"time"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/apis"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/configs"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/databases"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/logger"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/services"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"go.uber.org/zap"
)

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func main() {
	conf := configs.GetConfig()
	logger.NewLogger("agents-ai-api", conf.Env, "", true)
	defer logger.Sync()

	defer func() {
		if err := recover(); err != nil {
			panicErr := errors.Wrap(errors.New("panic start server"), string(debug.Stack()))
			logger.Info(
				logger.LOGGER_API_APP_PANIC,
				"panic start server",
				zap.Error(panicErr),
			)
			fmt.Println(err)
			fmt.Println(panicErr)
			return
		}
	}()

	jobEnabled := strings.ToLower(os.Getenv("JOB")) == "true" || os.Getenv("JOB") == "1" || conf.Job
	migrateDBMain := databases.MigrateDBMain
	if os.Getenv("DEV") == "true" || !jobEnabled {
		migrateDBMain = nil
	}

	dbMain, err := databases.Init(conf.DbURL, migrateDBMain, 5, 20, conf.Debug)
	if err != nil {
		logger.Fatal("databases.Init", zap.Error(err))
	}

	daos.InitDBConn(
		dbMain,
	)
	s := services.NewService(conf)
	r := gin.New()
	srv := apis.NewServer(
		r,
		conf,
		s,
	)
	srv.Routers()
	if conf.Port == 0 {
		conf.Port = 8080
	}
	if jobEnabled {
		srv.RunJobs()
		// srv.RunTeleBotJob()
	}
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Port),
		Handler: r,
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				debug.PrintStack()
			}
		}()
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	srv.DisableJobs()

	delayTs := 10 * time.Second
	if jobEnabled {
		delayTs = 30 * time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), delayTs)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	<-ctx.Done()
	log.Println("Server is down")
}
