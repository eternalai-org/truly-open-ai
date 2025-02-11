package setting

import (
	"agent-battle/internal/core/port"

	mongoDriver "go.mongodb.org/mongo-driver/mongo"
)

type Setting struct {
	mongoDb     *mongoDriver.Database
	GameUsecase port.IGameUsecase
	GameWorker  port.IGameWorker
}

type SettingOption func(*Setting)

func WithGameUsecase(uc port.IGameUsecase) SettingOption {
	return func(s *Setting) {
		s.GameUsecase = uc
	}
}

func WithGameWorker(uc port.IGameWorker) SettingOption {
	return func(s *Setting) {
		s.GameWorker = uc
	}
}

func WithDB(db *mongoDriver.Database) SettingOption {
	return func(s *Setting) {
		s.mongoDb = db
	}
}

func Init(options ...SettingOption) *Setting {
	s := &Setting{}
	for _, opt := range options {
		opt(s)
	}

	return s
}
