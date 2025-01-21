package main

import (
	"flag"
	_ "net/http/pprof"

	"solo/config"
	"solo/pkg/logger"
)

var configFile string

func main() {
	// init flag
	flag.StringVar(&configFile, "config-file", "env/development.yml", "Config file path")
	flag.Parse()

	cnf, err := config.ReadConfig(configFile)
	if err != nil {
		logger.AtLog.Fatal(err)
	}

	_ = cnf
	// TODO - start router here.

}
