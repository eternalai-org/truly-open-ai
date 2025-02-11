package mongodb

import "go.uber.org/fx"

var Module = fx.Module("mongodb", fx.Provide(Init))
