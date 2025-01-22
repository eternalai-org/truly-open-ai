package main

import (
	_ "net/http/pprof"
	"solo/internal/delivery/cmd"
)

var configFile string

func main() {
	_cmd, _ := cmd.NewCMD()
	_cmd.Run()
}
