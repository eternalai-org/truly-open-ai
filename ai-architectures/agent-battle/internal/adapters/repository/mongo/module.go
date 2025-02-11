package mongo

import "go.uber.org/fx"

var (
	WorkerRepoModule = fx.Module("mongo", fx.Provide(NewWorkerRepo))
	GameRepoModule   = fx.Module("mongo", fx.Provide(NewGameRepo))
	SettingRepoModule = fx.Module("mongo", fx.Provide(NewSettingRepo))
)
