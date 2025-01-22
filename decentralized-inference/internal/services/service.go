package services

import (
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
	go s.JobWatchSubmitSolution()
	return nil
}
