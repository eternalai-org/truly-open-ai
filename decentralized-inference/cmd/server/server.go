package server

import (
	"context"
	"decentralized-inference/internal/config"
	"decentralized-inference/internal/database"
	"decentralized-inference/internal/logger"
	"decentralized-inference/internal/services"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

type Server struct {
	Cfg     *config.Config
	Service *services.Service
}

func NewServer() (*Server, error) {

	cfg := config.GetConfig()
	mongoDB, err := database.InitMongo(cfg.Mongodb.Db, cfg.Mongodb.Uri)
	if err != nil {
		return nil, err
	}

	svc := services.NewService()
	svc.WithOptions(
		services.WithDatabase(mongoDB),
		services.WithConfig(cfg),
	)

	err = svc.StartService()
	if err != nil {
		return nil, err
	}
	return &Server{
		Cfg:     cfg,
		Service: svc,
	}, nil
}

func (s *Server) Start() {
	app := s.startRouter()

	port := s.Cfg.Server.Port
	if port == 0 {
		port = 8484
	}

	go func() {
		if err := app.Run(fmt.Sprintf(":%d", port)); err != nil {
			logger.AtLog.Fatalf("server start error: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	logger.GetLoggerInstanceFromContext(context.TODO()).Info("Shutting down...", zap.Any("signal", sig))
	logger.GetLoggerInstanceFromContext(context.TODO()).Info("Server stopped.")
}
