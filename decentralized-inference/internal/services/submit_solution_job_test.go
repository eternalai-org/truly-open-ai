package services

import (
	"decentralized-inference/internal/models"
	"testing"
)

func TestScanSubmitSolution(t *testing.T) {
	s := Service{}
	chain := models.ChainConfig{
		ChainID:              "8453",
		ListRPC:              []string{request.ChainInfo.Rpc},
		Type:                 models.ChainTypeEth,
		WorkerHubAddress:     request.WorkerHubAddress,
		AgentContractAddress: request.AgentContractAddress,
	}
	s.JobWatchSubmitSolution(&chain)
}
