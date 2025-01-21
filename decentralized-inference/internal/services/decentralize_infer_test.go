package services

import (
	"context"
	"decentralized-inference/internal/abi"
	client2 "decentralized-inference/internal/client"
	"decentralized-inference/internal/models"
	"fmt"
	ethreumAbi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
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

func TestCreateAgent(t *testing.T) {
	client, err := client2.NewClient(request.ChainInfo.Rpc, models.ChainTypeEth, false, "", "")
	if err != nil {
		t.Fatal(err)
	}
	_, pbkHex, err := client2.GetAccountInfo(request.InferPriKey)
	if err != nil {
		panic(err)
	}

	hybridModelABI, err := ethreumAbi.JSON(strings.NewReader(abi.AI721ContractMetaData.ABI))
	if err != nil {
		panic(err)
	}

	//Mint(opts *bind.TransactOpts, to common.Address, uri string, data []byte, fee *big.Int, promptKey string, promptScheduler common.Address, modelId uint32)
	dataBytes, err := hybridModelABI.Pack(
		"mint", common.HexToAddress(request.WorkerHubAddress),
		"aaa",
		[]byte("bbb"),
		big.NewInt(0),
		"ai721",
		common.HexToAddress(request.WorkerHubAddress),
		uint32(0),
	)

	if err != nil {
		panic(err)
	}

	tx, err := client.Transact(request.InferPriKey, *pbkHex, common.HexToAddress(request.AgentContractAddress), big.NewInt(0), dataBytes)
	if err != nil {
		panic(err)
	}
	fmt.Println(tx.TxHash)

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
		InferId:          42,
		ChainInfo:        request.ChainInfo,
	}
	response, err := s.GetDecentralizeInferResult(context.Background(), &requestResult)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
