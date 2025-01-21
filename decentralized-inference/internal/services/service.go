package services

import (
	"context"
	"decentralized-inference/internal/config"
	"decentralized-inference/internal/database"
)

type Service struct {
	db   *database.Database
	conf *config.Config
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) WithOptions(options ...ServiceOption) {
	for _, opt := range options {
		opt(s)
	}
}

type ServiceOption func(*Service)

func WithDatabase(db *database.Database) ServiceOption {
	return func(s *Service) {
		s.db = db
	}
}

func WithConfig(conf *config.Config) ServiceOption {
	return func(s *Service) {
		s.conf = conf
	}
}

func (s *Service) StartService() error {
	list, err := s.ListChainConfig(context.Background())
	if err != nil {
		return err
	}

	for _, item := range list {
		config := item.MakeCopy()
		if len(config.WorkerHubAddress) == 0 {
			go s.JobWatchSubmitSolution(config)
		}
	}
	return nil
}
