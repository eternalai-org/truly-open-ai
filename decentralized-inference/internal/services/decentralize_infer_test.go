package services

import (
	"context"
	"decentralized-inference/cmd/other/chat"
	"decentralized-inference/internal/models"
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
	Input:                "Hello World",
	WorkerHubAddress:     "0x963691C0b25a8d0866EA17CefC1bfBDb6Ec27894",
	AgentContractAddress: "0x458bE45957F8f29bBf597d5a953097c4095D9231",
	InferPriKey:          "", //update
	ModelId:              "",
	ChainInfo: models.ChainInfoRequest{
		Rpc:    "https://base.llamarpc.com",
		ZkSync: false,
	},
}

func init() {
	chatConfig, _ := chat.LoadChatConfig()
	request.InferPriKey = chatConfig.InferWalletKey
	requestNoAgent.InferPriKey = chatConfig.InferWalletKey
}
func TestCreateInfer(t *testing.T) {
	s := Service{}
	response, err := s.CreateDecentralizeInfer(context.Background(), &request)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func TestCreateInferNoAgent(t *testing.T) {
	s := Service{}
	response, err := s.CreateDecentralizeInferNoAgent(context.Background(), &requestNoAgent)
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
