package config

import (
	"encoding/json"
	"os"
)

var config *Config

func GetConfig() *Config {
	return config
}

type Config struct {
	MongoDB MongoConfig `json:"mongodb"`
}

type MongoConfig struct {
	Uri string `json:"uri"`
}

func Load() (*Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		panic("config.json not found, please copy config.json.example to config.json and fill in the values")
	}
	decoder := json.NewDecoder(file)
	v := Config{}
	err = decoder.Decode(&v)
	if err != nil {
		panic("config.json is invalid, please check the values")
	}
	config = &v
	return config, nil
}

func init() {
	if _, err := Load(); err != nil {
		panic(err)
	}
}
