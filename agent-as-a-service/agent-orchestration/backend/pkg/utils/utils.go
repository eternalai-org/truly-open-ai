package utils

type Environment string

const Production Environment = "production"

func IsEnvProduction(env string) bool {
	return env == string(Production)
}
