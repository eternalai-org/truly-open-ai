package services

import (
	"context"
	"decentralized-inference/internal/models"
	"fmt"
	"testing"
)

var request = models.DecentralizeInferRequest{
	Input:                "Hello World",
	WorkerHubAddress:     "0x963691C0b25a8d0866EA17CefC1bfBDb6Ec27894",
	AgentContractAddress: "0x458bE45957F8f29bBf597d5a953097c4095D9231",
	InferPriKey:          "ea88d751a20c8d6f84b2941d990c6c1d439e08aa4b46bbc4e315dc03865345ae", //update
	ExternalData:         "",
	AgentId:              "1",
	ChainInfo: models.ChainInfoRequest{
		Rpc:    "https://base.llamarpc.com",
		ZkSync: false,
	},
}

func TestCreateInfer(t *testing.T) {
	s := Service{}
	response, err := s.CreateDecentralizeInfer(context.Background(), &request)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func TestGetInferInfo(t *testing.T) {
	s := Service{}
	requestResult := models.InferResultRequest{
		WorkerHubAddress: request.WorkerHubAddress,
		InferId:          43,
		ChainInfo:        request.ChainInfo,
	}
	response, err := s.GetDecentralizeInferResult(context.Background(), &requestResult)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
