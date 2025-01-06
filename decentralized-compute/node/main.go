package main

import (
	"context"
	"flag"
	_ "net/http/pprof"
	"time"

	"solo/pkg"

	"solo/config"
	"solo/pkg/logger"
	"solo/pkg/miner"
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

	taskWatcher, err := miner.NewMiner(cnf)
	if err != nil {
		logger.AtLog.Fatal(err)
	}

goto_here:
	verifed := taskWatcher.Verify()
	if !verifed {
		err := taskWatcher.MakeVerify()
		if err != nil {
			logger.AtLog.Error(err)
		}

		time.Sleep(time.Second * pkg.TimeToWating)
		goto goto_here
	}

	done := make(chan bool)

	// get and process tasks
	// taskWatcher.RejoinForMinting(ctx)
	ctx := context.Background()
	go taskWatcher.GetPendingTasks(ctx)

	go taskWatcher.ExecueteTasks(ctx)

	<-done
}
