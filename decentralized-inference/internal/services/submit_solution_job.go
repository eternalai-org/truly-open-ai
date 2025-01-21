package services

import (
	"context"
	"decentralized-inference/internal/abi"
	"decentralized-inference/internal/client"
	"decentralized-inference/internal/config"
	"decentralized-inference/internal/eaimodel"
	"decentralized-inference/internal/lighthouse"
	"decentralized-inference/internal/logger"
	"decentralized-inference/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (s *Service) GetData(data []byte) ([]byte, error) {
	dataString := string(data)
	dataByte := data
	var err error
	if strings.HasPrefix(dataString, config.IPFSPrefix) {
		data, _, err = lighthouse.DownloadDataSimpleWithRetry(dataString)
		if err != nil {
			return nil, err
		}
	} else if strings.HasPrefix(dataString, config.FilePrefix) {
		fileName := strings.TrimPrefix(dataString, config.FilePrefix)
		dataByte, err = os.ReadFile(fmt.Sprintf("%v/%v", config.GetConfig().FilePathInfer, fileName))
		if err != nil {
			return nil, err
		}
	}
	return dataByte, nil
}

func (s *Service) WriteInput(address string, data []byte) (string, error) {
	filePath := config.GetConfig().FilePathInfer
	filePath = strings.TrimSuffix(filePath, "/")
	fileName := fmt.Sprintf("%v_%v", address, time.Now().Unix())
	fileFullDir := fmt.Sprintf("%v/%v", filePath, fileName)
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(filePath, 0755)
		if err != nil {
			return "", err
		}
	}
	err = os.WriteFile(fileFullDir, data, 0644)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func (s *Service) JobWatchSubmitSolution(chainConfig *models.ChainConfig) {
	ctx := context.Background()
	logger.GetLoggerInstanceFromContext(ctx).Info("Start JobWatchSubmitSolution", zap.Any("chainConfig", chainConfig))

	defer func() {
		if err := recover(); err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Error("JobWatchSubmitSolution panic", zap.Any("err", err))
		}
		time.Sleep(2 * time.Second)
		s.JobWatchSubmitSolution(chainConfig)
	}()

	interval := 500 * time.Millisecond
	jobName := fmt.Sprintf("JobWatchSubmitSolution-%v", chainConfig.ChainID)

	for {
		time.Sleep(interval)

		jobCfg, err := s.GetJobConfig(jobName)
		if err != nil {
			if err.Error() != mongo.ErrNoDocuments.Error() {
				panic(err)
			} else {
				jobCfg = &models.JobConfig{
					JobName:  jobName,
					LastRun:  time.Now(),
					Interval: 2000, // unit: millisecond
					Enable:   true,
				}
				err = s.addJobConfig(context.Background(), jobCfg)
				if err != nil {
					panic(err)
				}
			}
		} else {
			jobCfg.LastRun = time.Now()
			err = s.updateJobLastRun(ctx, jobCfg.JobName, jobCfg.LastRun)
			if err != nil {
				panic(err)
			}
		}
		if !jobCfg.Enable {
			continue
		}
		interval = time.Duration(jobCfg.Interval) * time.Millisecond

		chainConfig, err = s.FindChainConfig(ctx, chainConfig.ChainID)
		if err != nil {
			continue
		}
		rpc := chainConfig.GetRPC()
		c, err := client.NewClient(rpc, chainConfig.Type,
			chainConfig.PaymasterFeeZero,
			chainConfig.PaymasterAddress,
			chainConfig.PaymasterToken)

		if err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Error("[JobWatchSubmitSolution] NewEthClient", zap.String("chainId", chainConfig.ChainID), zap.Error(err))
			return
		}

		contract, err := abi.NewWorkerhubContract(common.HexToAddress(chainConfig.WorkerHubAddress), c.Client)
		if err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Error("[JobWatchSubmitSolution] NewWorkerHub", zap.String("chainId", chainConfig.ChainID), zap.Error(err))
			return
		}

		currentBlock, err := c.Client.BlockNumber(ctx)
		if err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Error("[JobWatchSubmitSolution] BlockNumber", zap.Error(err))
			continue
		}

		state, err := s.GetContractSyncState(chainConfig.WorkerHubAddress, jobName)
		if err != nil {
			if err.Error() != mongo.ErrNoDocuments.Error() {
				logger.GetLoggerInstanceFromContext(ctx).Error("[JobWatchSubmitSolution] GetContractSyncState", zap.Error(err))
				continue
			}
			if state == nil {
				state = &models.ContractSyncState{
					Job:             jobName,
					ContractAddress: strings.ToLower(chainConfig.WorkerHubAddress),
					LastSyncedBlock: currentBlock - 1,
					ResyncFromBlock: 0,
				}
				err = s.AddContractSyncState(state)
				if err != nil {
					logger.GetLoggerInstanceFromContext(ctx).Error("[JobWatchSubmitSolution] AddContractSyncState", zap.Error(err))
					time.Sleep(2 * time.Second)
					continue
				}
			}
		}

		if state.ClearDataAndSync {
			state.ResyncToBlock = state.LastSyncedBlock
			state.ClearDataAndSync = false
			state.LastSyncedBlock = state.ResyncFromBlock
			err = s.AddOrUpdateContractSyncState(state)
			if err != nil {
				logger.GetLoggerInstanceFromContext(ctx).Error("[JobWatchSubmitSolution] AddOrUpdateContractSyncState", zap.Error(err))
				continue
			}

			logger.GetLoggerInstanceFromContext(ctx).Info("[JobWatchSubmitSolution] ClearDataAndSync", zap.Any("state", state))
		}

		startBlock := state.LastSyncedBlock - chainConfig.BackwardBlockNumber
		var endBlock uint64
		for {
			if startBlock > currentBlock {
				break
			}
			if currentBlock-startBlock > 1000 {
				endBlock = startBlock + 1000
			} else {
				endBlock = currentBlock
			}

			err = s.filterEventSolutionSubmission(ctx, contract, chainConfig, c, startBlock, endBlock)
			if err != nil {
				logger.GetLoggerInstanceFromContext(ctx).Error("[JobWatchSubmitSolution] filterEventSolutionSubmission",
					zap.Any("network", chainConfig),
					zap.Error(err))
				break
			}
			state.LastSyncedBlock = endBlock
			err = s.UpdateContractSyncStateByAddressAndJob(state)
			if err != nil {
				logger.GetLoggerInstanceFromContext(ctx).Error("[JobWatchSubmitSolution] UpdateContractSyncStateByAddressAndJob", zap.Error(err))
				break
			}

			startBlock = endBlock + 1
		}
	}
}

func (s *Service) filterEventSolutionSubmission(ctx context.Context, whContract *abi.WorkerhubContract,
	chain *models.ChainConfig, client *client.Client, startBlock uint64, endBlock uint64,
) error {
	iter, err := whContract.FilterSolutionSubmission(&bind.FilterOpts{
		Start:   startBlock,
		End:     &endBlock,
		Context: ctx,
	}, nil, nil)
	if err != nil {
		return err
	}
	for iter.Next() {
		inferId := iter.Event.InferId.String()
		workerAddress := strings.ToLower(iter.Event.Miner.Hex())

		receipt, err := client.Client.TransactionReceipt(ctx, iter.Event.Raw.TxHash)
		if err != nil {
			return errors.Join(err, errors.New("error while getting tx receipt"))
		}

		if receipt.Status != ethtypes.ReceiptStatusSuccessful {
			return errors.New("tx failed")
		}

		assignmentEntity, err := s.GetModelWorkerProcessHistoryByFilter(ctx,
			bson.M{
				"inference_id":   inferId,
				"worker_address": workerAddress,
				"chain_id":       chain.ChainID,
			})
		if err != nil {
			return err
		}
		if assignmentEntity == nil {
			continue
		}
		txHash := strings.ToLower(iter.Event.Raw.TxHash.Hex())

		requestInfo, err := whContract.GetInferenceInfo(nil, iter.Event.InferId.Uint64())
		if err != nil {
			return err
		}

		/*fee := requestInfo.Value
		requester := requestInfo.Creator*/

		var predictResult eaimodel.TaskResult

		err = json.Unmarshal(requestInfo.Output, &predictResult)
		if err != nil || assignmentEntity.StoreRawFlag {
			predictResult = eaimodel.TaskResult{
				ResultURI: "",
				Storage:   eaimodel.EaiChainStorageType,
				Data:      requestInfo.Output,
			}
		}
		if len(predictResult.ResultURI) > 0 {
			predictResult.Data, err = s.GetData([]byte(predictResult.ResultURI))
			if err != nil {
				return err
			}
		}

		var batchInfers []*models.BatchInferHistory

		inputBytes, err := s.GetData(requestInfo.Input)
		if err != nil {
			return err
		}
		var batchFullPrompts []*models.BatchInferHistory
		err = json.Unmarshal(inputBytes, &batchFullPrompts)
		if err == nil && len(batchFullPrompts) > 0 {
			batchInfers = batchFullPrompts
		}
		assignmentHistory := &models.ModelWorkerProcessHistories{
			AssignmentId:    inferId,
			WorkerAddress:   strings.ToLower(requestInfo.ProcessedMiner.Hex()),
			ModelAddress:    "",
			ModelID:         strconv.Itoa(int(requestInfo.ModelId)),
			InferenceInput:  string(requestInfo.Input),
			InferenceId:     inferId,
			Status:          models.CLOUD_PROCESSING_STATUS_DONE,
			AssignmentRole:  models.AssignmentRoleMiner,
			ZkSync:          true,
			ChainID:         chain.ChainID,
			ExecuteTaskDone: false,
			IsAgentInfer:    len(batchInfers) > 0,
			BatchInfers:     batchInfers,
			TxHash:          txHash,
		}
		if !strings.HasPrefix(assignmentHistory.InferenceInput, config.IPFSPrefix) && chain.SupportStoreRaw {
			assignmentHistory.StoreRawFlag = true
		}
		assignmentHistory.CreatedAt = time.Now().UTC()
		assignmentHistory.UpdatedAt = assignmentHistory.CreatedAt

		err = s.InsertModelWorkerProcessHistories(ctx, assignmentHistory)
		if err != nil && !mongo.IsDuplicateKeyError(err) {
			logger.GetLoggerInstanceFromContext(ctx).Error("[JobWatchWorkerHubNewInferZKChain] InsertModelWorkerProcessHistories",
				zap.Error(err), zap.Any("assignment", assignmentHistory))
			return err
		}

	}
	return nil
}
