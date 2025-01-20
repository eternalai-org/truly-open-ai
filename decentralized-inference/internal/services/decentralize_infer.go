package services

import (
	"context"
	"decentralized-inference/internal/abi"
	"decentralized-inference/internal/client"
	"decentralized-inference/internal/logger"
	"decentralized-inference/internal/models"
	"fmt"
	ethreumAbi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"math/big"
	"strings"
)

func (s *Service) CreateDecentralizeInfer(ctx context.Context, info *models.DecentralizeInferRequest) (*models.DecentralizeInferResponse, error) {
	//@TODO
	agentId, ok := new(big.Int).SetString(info.AgentId, 10)
	if !ok {
		return nil, fmt.Errorf("agentId :%v is not valid", info.AgentId)
	}
	_, pbkHex, err := client.GetAccountInfo(info.InferPriKey)
	if err != nil {
		return nil, fmt.Errorf("get account info error: %v", err)
	}
	client, err := client.NewClient(info.ChainInfo.Rpc, models.ChainTypeEth,
		false,
		"", "")
	if err != nil {
		return nil, fmt.Errorf("init client err: %w", err)
	}

	workerHubContract, err := abi.NewWorkerhubContract(common.HexToAddress(info.WorkerHubAddress), nil)
	if err != nil {
		return nil, err
	}

	hybridModelABI, err := ethreumAbi.JSON(strings.NewReader(abi.AI721ContractMetaData.ABI))
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("error when get abi", zap.Error(err))
		return nil, err
	}

	agentContract, err := abi.NewAI721Contract(common.HexToAddress(info.AgentContractAddress), client.ETHClient)
	if err != nil {
		return nil, err
	}

	agentFee, err := agentContract.GetAgentFee(nil, agentId)
	if err != nil {
		return nil, fmt.Errorf("get agent fee err: %w", err)
	}
	//Infer(opts *bind.TransactOpts, agentId *big.Int, fwdCalldata []byte, externalData string, promptKey string, feeAmount *big.Int)
	dataBytes, err := hybridModelABI.Pack(
		"infer", agentId,
		[]byte(info.Input),
		info.ExternalData,
		"ai721",
		agentFee,
	)

	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("[SubmitInferTaskWorkerHubV1] error when pack data", zap.Error(err))
		return nil, err
	}
	fmt.Println("priKey", info.InferPriKey)
	fmt.Println("contract address", info.AgentContractAddress)
	fmt.Println("agentId", agentId)
	fmt.Println("fwdCalldata", info.Input)
	fmt.Println("externalData", info.ExternalData)
	fmt.Println("promptKey", "ai721")
	fmt.Println("agentFee", agentFee)

	tx, err := client.Transact(info.InferPriKey, *pbkHex, common.HexToAddress(info.AgentContractAddress), big.NewInt(0), dataBytes)
	if err != nil {
		return nil, fmt.Errorf("send transaction with err %v", err)
	}

	logs := tx.Receipt.Logs
	inferId := uint64(0)
	for _, item := range logs {
		inferData, err := workerHubContract.ParseNewInference(*item)
		if err == nil {
			inferId = inferData.InferenceId
			break
		}
	}

	if inferId == 0 {
		return nil, fmt.Errorf("inferId is zero , tx: %v ", tx.TxHash.Hex())
	}

	return &models.DecentralizeInferResponse{
		TxHash:    tx.TxHash.Hex(),
		InferId:   inferId,
		ChainInfo: info.ChainInfo,
	}, nil
}

func (s *Service) GetDecentralizeInferResult(ctx context.Context, info *models.InferResultRequest) (*models.InferResultResponse, error) {

	client, err := client.NewClient(info.ChainInfo.Rpc, models.ChainTypeEth,
		false,
		"", "")
	if err != nil {
		return nil, fmt.Errorf("init client err: %w", err)
	}

	workerHubContract, err := abi.NewWorkerhubContract(common.HexToAddress(info.WorkerHubAddress), client.ETHClient)
	if err != nil {
		return nil, err
	}

	inferInfo, err := workerHubContract.GetInferenceInfo(nil, info.InferId)
	if err != nil {
		return nil, fmt.Errorf("get infer info err: %v", err)
	}

	status := models.InferResultStatusDone
	if len(inferInfo.Output) == 0 {
		currentBlock, err := client.Client.BlockNumber(ctx)
		if err != nil {
			return nil, fmt.Errorf("get block err: %v", err)
		}
		if currentBlock == 0 {
			return nil, fmt.Errorf("get block err: current block is 0")
		}
		if currentBlock > inferInfo.SubmitTimeout.Uint64() {
			status = models.InferResultStatusTimeOut
		} else {
			status = models.InferResultStatusWaitingProcess
		}
	}

	return &models.InferResultResponse{
		ChainInfo:        info.ChainInfo,
		WorkerHubAddress: info.WorkerHubAddress,
		InferId:          info.InferId,
		Input:            string(inferInfo.Input),
		Output:           string(inferInfo.Output),
		Creator:          inferInfo.Creator.String(),
		ProcessedMiner:   inferInfo.ProcessedMiner.String(),
		Status:           status,
		SubmitTimeout:    inferInfo.SubmitTimeout.Uint64(),
	}, nil
}
