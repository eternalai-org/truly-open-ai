package main

import (
	"context"
	"decentralized-inference/cmd/other/chat"
	"decentralized-inference/internal/config"
	"decentralized-inference/internal/models"
	"decentralized-inference/internal/services"
	"fmt"
	"testing"
)

var request = models.DecentralizeInferRequest{
	Input:                "Hello World",
	WorkerHubAddress:     "0x963691C0b25a8d0866EA17CefC1bfBDb6Ec27894",
	AgentContractAddress: "0x458bE45957F8f29bBf597d5a953097c4095D9231",
	InferPriKey:          "", //update
	ExternalData:         "",
	AgentId:              "1",
	ChainInfo: models.ChainInfoRequest{
		Rpc:    "https://base.llamarpc.com",
		ZkSync: false,
	},
}

var requestNoAgent = models.DecentralizeInferNoAgentRequest{
	Input:            "Hello World",
	WorkerHubAddress: "0x963691C0b25a8d0866EA17CefC1bfBDb6Ec27894",
	InferPriKey:      "", //update
	ModelId:          "70050",
	ChainInfo: models.ChainInfoRequest{
		Rpc:    "https://base.llamarpc.com",
		ZkSync: false,
	},
}
var s = services.Service{}

func init() {
	chatConfig, _ := chat.LoadChatConfig()
	request.InferPriKey = chatConfig.InferWalletKey
	requestNoAgent.InferPriKey = chatConfig.InferWalletKey
	cfg := config.GetConfig()
	s.WithOptions(
		services.WithConfig(cfg),
	)
}
func TestCreateInfer(t *testing.T) {
	response, err := s.CreateDecentralizeInfer(context.Background(), &request)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func TestCreateInferNoAgent(t *testing.T) {
	response, err := s.CreateDecentralizeInferNoAgent(context.Background(), &requestNoAgent)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func TestGetInferInfo(t *testing.T) {
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
