package services

import (
	"decentralized-inference/internal/database"
)

type Service struct {
	db *database.Database
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
