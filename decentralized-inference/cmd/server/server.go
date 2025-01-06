package server

import (
	"context"
	"decentralized-inference/config"
	"decentralized-inference/internal/logger"
	"decentralized-inference/service"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

type Server struct {
	Cfg     *config.Config
	Service *service.Service
}

func NewServer() (*Server, error) {
	return &Server{
		Cfg:     config.GetConfig(),
		Service: service.NewService(),
	}, nil
}

func (s *Server) Start() {
	// Start the server

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	logger.GetLoggerInstanceFromContext(context.TODO()).Info("Shutting down...", zap.Any("signal", sig))
	logger.GetLoggerInstanceFromContext(context.TODO()).Info("Server stopped.")
}
