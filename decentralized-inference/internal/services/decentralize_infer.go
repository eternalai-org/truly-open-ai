package services

import (
	"bytes"
	"context"
	"decentralized-inference/internal/abi"
	"decentralized-inference/internal/client"
	"decentralized-inference/internal/config"
	"decentralized-inference/internal/libs/http_client"
	"decentralized-inference/internal/logger"
	"decentralized-inference/internal/models"
	"decentralized-inference/internal/types"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	ethreumAbi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
)

type AssistantResp struct {
	ContractAgentID string `bson:"contract_agent_id" json:"contract_agent_id"`
	SystemContent   string `bson:"system_content" json:"system_content"`
}

func (s *Service) GetSystemPromptFromAAAS(agentId string) string {
	url := "http://localhost:8480/api/agent/list-local-agent"
	resp, statusCode, err := http_client.RequestHttp(url, http.MethodGet, map[string]string{}, nil, 10)
	if err != nil {
		return ""
	}
	if statusCode != http.StatusOK {
		return ""
	}
	var response struct {
		Result []*AssistantResp `json:"result"`
		Data   interface{}      `json:"data"`
		Error  error            `json:"error"`
		Count  *uint            `json:"count,omitempty"`
	}

	if err := json.Unmarshal(resp, &response); err != nil {
		return ""
	}

	for _, item := range response.Result {
		if item.ContractAgentID == agentId {
			return item.SystemContent
		}
	}

	return ""

}

func (s *Service) CreateDecentralizeInferV2(ctx context.Context, info *models.DecentralizeInferRequest) (interface{}, error) {
	fmt.Println(fmt.Sprintf("Client is chatting with agent: %v, contractAddress: %v", info.AgentId, info.AgentContractAddress))
	var systemPromptStr = "You are a helpful assistant."
	var err error
	systemPromptStrFromAAAS := s.GetSystemPromptFromAAAS(info.AgentId)
	if systemPromptStrFromAAAS != "" {
		systemPromptStr = systemPromptStrFromAAAS
	} else {
		systemPromptStr, err = func() (string, error) {
			agentId, ok := new(big.Int).SetString(info.AgentId, 10)
			if !ok {
				return "", fmt.Errorf("agentId :%v is not valid", info.AgentId)
			}

			ethClient, _err := client.NewClient(info.ChainInfo.Rpc, models.ChainTypeEth,
				false,
				"", "")
			if _err != nil {
				return "", fmt.Errorf("init ethClient err: %w", err)
			}

			agentContract, _err := abi.NewAI721Contract(common.HexToAddress(info.AgentContractAddress), ethClient.ETHClient)
			if _err != nil {
				return "", _err
			}

			systemPromptContract, _err := agentContract.GetAgentSystemPrompt(nil, agentId)
			if _err != nil {
				return "", fmt.Errorf("get agent system prompt err: %w", err)
			}

			if len(systemPromptContract) > 0 {
				systemPromptStr = string(systemPromptContract[0])
			}

			return systemPromptStr, nil
		}()
	}

	if err != nil {
		return nil, fmt.Errorf("get system prompt err: %w", err)
	}

	fmt.Println("Agent system prompt:", systemPromptStr)

	if systemPromptStr == "" {
		systemPromptStr = "You are a helpful assistant."
	}

	chatRequest := &openai.ChatCompletionRequest{
		Model: info.Model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: systemPromptStr,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: info.Input,
			},
		},
		MaxTokens: 4096,
	}

	fullUrl := "http://localhost:8004/v1/chat/completions"
	input, _ := json.Marshal(chatRequest)

	fmt.Println("Full request to LLM :", string(input))
	chatCompletionResp, statusCode, err := http_client.RequestHttp(fullUrl, http.MethodPost, map[string]string{}, bytes.NewBuffer(input), 10)
	if err != nil {
		return nil, err
	}
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("call api %v , status code %v != 200 , body:%v", fullUrl, statusCode, string(chatCompletionResp))
	}
	var response struct {
		Data openai.ChatCompletionResponse `json:"data"`
	}
	if err := json.Unmarshal(chatCompletionResp, &response); err != nil {
		return nil, err
	}
	return &response.Data, nil
	//var submitData = info.Input
	//
	//if s.conf.SubmitFilePath {
	//	fileName, err := s.WriteInput(strings.ToLower((*pbkHex).Hex()), []byte(info.Input))
	//	if err != nil {
	//		return nil, fmt.Errorf("write input file err: %w", err)
	//	}
	//	submitData = fmt.Sprintf("%v%v", config.FilePrefix, fileName)
	//}
	//
	////Infer(opts *bind.TransactOpts, agentId *big.Int, fwdCalldata []byte, externalData string, promptKey string, feeAmount *big.Int)
	//dataBytes, err := agentContractABI.Pack(
	//	"infer", agentId,
	//	[]byte(submitData),
	//	info.ExternalData,
	//	"ai721",
	//	agentFee,
	//)
	//
	//if err != nil {
	//	logger.GetLoggerInstanceFromContext(ctx).Error("[SubmitInferTaskWorkerHubV1] error when pack data", zap.Error(err))
	//	return nil, err
	//}
	//
	//tx, err := ethClient.Transact(info.InferPriKey, *pbkHex, common.HexToAddress(info.AgentContractAddress), big.NewInt(0), dataBytes)
	//if err != nil {
	//	return nil, fmt.Errorf("send transaction with err %v", err)
	//}
	//
	//logs := tx.Receipt.Logs
	//var inferId *big.Int
	//for _, item := range logs {
	//	inferData, err := workerHubContract.ParseNewInference(*item)
	//	if err == nil {
	//		inferId = inferData.InferenceId
	//		break
	//	}
	//}
	//
	//if inferId == nil || inferId.Cmp(big.NewInt(0)) == 0 {
	//	return nil, fmt.Errorf("inferId is zero , tx: %v ", tx.TxHash.Hex())
	//}
	//inferIdResp := inferId.Uint64()
	//
	//return &models.DecentralizeInferResponse{
	//	TxHash:    tx.TxHash.Hex(),
	//	InferId:   inferIdResp,
	//	ChainInfo: info.ChainInfo,
	//}, nil
}

func (s *Service) CreateDecentralizeInfer(ctx context.Context, info *models.DecentralizeInferRequest) (*models.DecentralizeInferResponse, error) {
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

	agentContractABI, err := ethreumAbi.JSON(strings.NewReader(abi.AI721ContractMetaData.ABI))
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

	var submitData = info.Input

	if s.conf.SubmitFilePath {
		fileName, err := s.WriteInput(strings.ToLower((*pbkHex).Hex()), []byte(info.Input))
		if err != nil {
			return nil, fmt.Errorf("write input file err: %w", err)
		}
		submitData = fmt.Sprintf("%v%v", config.FilePrefix, fileName)
	}

	//Infer(opts *bind.TransactOpts, agentId *big.Int, fwdCalldata []byte, externalData string, promptKey string, feeAmount *big.Int)
	dataBytes, err := agentContractABI.Pack(
		"infer", agentId,
		[]byte(submitData),
		info.ExternalData,
		"ai721",
		agentFee,
	)

	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("[SubmitInferTaskWorkerHubV1] error when pack data", zap.Error(err))
		return nil, err
	}

	tx, err := client.Transact(info.InferPriKey, *pbkHex, common.HexToAddress(info.AgentContractAddress), big.NewInt(0), dataBytes)
	if err != nil {
		return nil, fmt.Errorf("send transaction with err %v", err)
	}

	logs := tx.Receipt.Logs
	var inferId *big.Int
	for _, item := range logs {
		inferData, err := workerHubContract.ParseNewInference(*item)
		if err == nil {
			inferId = inferData.InferenceId
			break
		}
	}

	if inferId == nil || inferId.Cmp(big.NewInt(0)) == 0 {
		return nil, fmt.Errorf("inferId is zero , tx: %v ", tx.TxHash.Hex())
	}
	inferIdResp := inferId.Uint64()

	return &models.DecentralizeInferResponse{
		TxHash:    tx.TxHash.Hex(),
		InferId:   inferIdResp,
		ChainInfo: info.ChainInfo,
	}, nil
}

func (s *Service) GetDecentralizeInferResult(ctx context.Context, info *models.InferResultRequest) (*models.InferResultResponse, error) {
	return nil, nil

	//client, err := client.NewClient(info.ChainInfo.Rpc, models.ChainTypeEth,
	//	false,
	//	"", "")
	//if err != nil {
	//	return nil, fmt.Errorf("init client err: %v", err)
	//}
	//chainId, err := client.Client.ChainID(ctx)
	//
	//if err != nil {
	//	return nil, fmt.Errorf("init client err: %w", err)
	//}
	//
	//workerHubContract, err := abi.NewWorkerhubContract(common.HexToAddress(info.WorkerHubAddress), client.ETHClient)
	//if err != nil {
	//	return nil, err
	//}
	//
	//inferInfo, err := workerHubContract.GetInferenceInfo(nil, info.InferId)
	//if err != nil {
	//	return nil, fmt.Errorf("get infer info err: %v", err)
	//}
	//
	//status := models.InferResultStatusDone
	//txSubmitSolution := ""
	//output := []byte("")
	//input, err := s.GetData(inferInfo.Input)
	//if err != nil {
	//	return nil, fmt.Errorf("get input err: %v", err)
	//}
	//if len(inferInfo.Output) == 0 {
	//	currentBlock, err := client.Client.BlockNumber(ctx)
	//	if err != nil {
	//		return nil, fmt.Errorf("get block err: %v", err)
	//	}
	//	if currentBlock == 0 {
	//		return nil, fmt.Errorf("get block err: current block is 0")
	//	}
	//	if currentBlock > inferInfo.SubmitTimeout.Uint64() {
	//		status = models.InferResultStatusTimeOut
	//	} else {
	//		status = models.InferResultStatusWaitingProcess
	//	}
	//} else {
	//	inferResultInfo, err := s.GetModelWorkerProcessHistoryByFilter(ctx, bson.M{
	//		"inference_id":   info.InferId,
	//		"worker_address": strings.ToLower(inferInfo.ProcessedMiner.String()),
	//		"chain_id":       chainId,
	//	})
	//	if err == nil && inferResultInfo != nil {
	//		txSubmitSolution = inferResultInfo.TxHash
	//	}
	//	output, err = s.GetData(inferInfo.Output)
	//	if err != nil {
	//		return nil, fmt.Errorf("get data err: %v", err)
	//	}
	//}
	//return &models.InferResultResponse{
	//	ChainInfo:        info.ChainInfo,
	//	WorkerHubAddress: info.WorkerHubAddress,
	//	InferId:          info.InferId,
	//	Input:            string(input),
	//	Output:           string(output),
	//	Creator:          inferInfo.Creator.String(),
	//	ProcessedMiner:   inferInfo.ProcessedMiner.String(),
	//	Status:           status,
	//	SubmitTimeout:    inferInfo.SubmitTimeout.Uint64(),
	//	TxSubmitSolution: txSubmitSolution,
	//}, nil
}

func (s *Service) CreateDecentralizeInferNoAgent(ctx context.Context, info *models.DecentralizeInferNoAgentRequest) (*models.DecentralizeInferResponse, error) {
	modelId, ok := new(big.Int).SetString(info.ModelId, 10)
	if !ok {
		return nil, fmt.Errorf("modelId :%v is not valid", info.ModelId)
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

	contractABI, err := ethreumAbi.JSON(strings.NewReader(abi.WorkerhubContractMetaData.ABI))
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("error when get abi", zap.Error(err))
		return nil, err
	}

	var submitData = info.Input

	if s.conf.SubmitFilePath {
		fileName, err := s.WriteInput(strings.ToLower((*pbkHex).Hex()), []byte(info.Input))
		if err != nil {
			return nil, fmt.Errorf("write input file err: %w", err)
		}
		submitData = fmt.Sprintf("%v%v", config.FilePrefix, fileName)
	}

	//Infer(opts *bind.TransactOpts, modelId uint32, input []byte, creator common.Address, flag bool)
	dataBytes, err := contractABI.Pack(
		"infer", uint32(modelId.Uint64()),
		[]byte(submitData),
		*pbkHex,
		false,
	)

	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("[SubmitInferTaskWorkerHubV1] error when pack data", zap.Error(err))
		return nil, err
	}

	tx, err := client.Transact(info.InferPriKey, *pbkHex, common.HexToAddress(info.WorkerHubAddress), big.NewInt(0), dataBytes)
	if err != nil {
		return nil, fmt.Errorf("send transaction with err %v", err)
	}

	logs := tx.Receipt.Logs
	inferId := uint64(0)
	for _, item := range logs {
		inferData, err := workerHubContract.ParseNewInference(*item)
		if err == nil {
			inferId = inferData.InferenceId.Uint64()
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

func (s *Service) CreateDecentralizeInferV2WithStream(ctx context.Context, info *models.DecentralizeInferRequest, printChan chan types.StreamData) (interface{}, error) {

	fmt.Println(fmt.Sprintf("Client is chatting with agent: %v, contractAddress: %v", info.AgentId, info.AgentContractAddress))
	var systemPromptStr = "You are a helpful assistant."
	var err error
	systemPromptStrFromAAAS := s.GetSystemPromptFromAAAS(info.AgentId)
	if systemPromptStrFromAAAS != "" {
		systemPromptStr = systemPromptStrFromAAAS
	} else {
		systemPromptStr, err = func() (string, error) {
			agentId, ok := new(big.Int).SetString(info.AgentId, 10)
			if !ok {
				return "", fmt.Errorf("agentId :%v is not valid", info.AgentId)
			}

			ethClient, _err := client.NewClient(info.ChainInfo.Rpc, models.ChainTypeEth,
				false,
				"", "")
			if _err != nil {
				return "", fmt.Errorf("init ethClient err: %w", err)
			}

			agentContract, _err := abi.NewAI721Contract(common.HexToAddress(info.AgentContractAddress), ethClient.ETHClient)
			if _err != nil {
				return "", _err
			}

			systemPromptContract, _err := agentContract.GetAgentSystemPrompt(nil, agentId)
			if _err != nil {
				return "", fmt.Errorf("get agent system prompt err: %w", err)
			}

			if len(systemPromptContract) > 0 {
				systemPromptStr = string(systemPromptContract[0])
			}

			return systemPromptStr, nil
		}()
	}

	if err != nil {
		return nil, fmt.Errorf("get system prompt err: %w", err)
	}

	fmt.Println("Agent system prompt:", systemPromptStr)

	if systemPromptStr == "" {
		systemPromptStr = "You are a helpful assistant."
	}

	chatRequest := &openai.ChatCompletionRequest{
		Model: info.Model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: systemPromptStr,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: info.Input,
			},
		},
		MaxTokens: 4096,
		Stream:    true,
	}

	fullUrl := "http://localhost:8004/v1/chat/completions"
	input, _ := json.Marshal(chatRequest)

	fmt.Println("Full request to LLM :", string(input))

	go http_client.RequestHttpWithStream(fullUrl, http.MethodPost, map[string]string{}, bytes.NewBuffer(input), 10, printChan)

	/*
		if err != nil {
			return nil, err
		}


		if statusCode != http.StatusOK {
			return nil, fmt.Errorf("call api %v , status code %v != 200 , body:%v", fullUrl, statusCode, string(chatCompletionResp))
		}
		var response struct {
			Data openai.ChatCompletionResponse `json:"data"`
		}
		if err := json.Unmarshal(chatCompletionResp, &response); err != nil {
			return nil, err
		}
	*/

	return nil, nil
}
