package main

import (
	"solo/internal/delivery/http"
	"solo/internal/usecase"
)

func main() {
	uc, _ := usecase.NewAPILocalChainV1()
	_cmd, _ := http.NewHttp(uc, 80)
	_cmd.Run()
}
