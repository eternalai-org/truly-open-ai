package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
)

func (s *Service) BatchChatCompletionPrompt(ctx context.Context, request *serializers.ChatCompletionRequest) (*models.BatchInferHistory, error) {
	chainConfig, err := s.dao.FirstChainConfig(daos.GetDBMainCtx(ctx), map[string][]interface{}{
		"chain_id = ?": {request.ChainId},
	}, map[string][]interface{}{}, false)
	if err != nil || chainConfig == nil {
		return nil, errs.NewError(errs.ErrBadRequest)
	}

	if request.Model == "" {
		request.Model = s.GetModelDefaultByChainID(request.ChainId)
	}

	promptInputByte, err := json.Marshal(request)
	if err != nil {
		return nil, errs.NewError(err)
	}

	batchInferHistory := &models.BatchInferHistory{
		UserAddress:          request.UserAddress,
		AgentContractAddress: chainConfig.AgentContractAddress,
		ContractAgentID:      "",
		ChainID:              request.ChainId,
		PromptInput:          string(promptInputByte),
		Status:               models.BatchInferHistoryStatusPending,
	}

	if batchInferHistory.Toolset == "" {
		batchInferHistory.Toolset = models.ToolsetTypeDefault
	}
	if request.MetaData != nil && len(request.MetaData.AgentContractId) > 0 {
		//assistant, err := s.GetAgentByID(ctx, request.MetaData.AgentContractId, chainConfig.AgentContractAddress, request.ChainId)
		//if err != nil || assistant == nil {
		//	return nil, errors.New(fmt.Sprintf("agent is not found on chain_id: %v", request.ChainId))
		//}
		batchInferHistory.ContractAgentID = request.MetaData.AgentContractId
		//batchInferHistory.AssistantID = assistant.ID.Hex()
	}

	err = s.dao.Create(daos.GetDBMainCtx(ctx), batchInferHistory)
	if err != nil {
		return nil, errs.NewError(err)
	}

	return batchInferHistory, nil
}

func (s *Service) GetBatchItemDetail(ctx context.Context, id uint) (*models.BatchInferHistory, error) {
	batchInferHistory, err := s.dao.FirstBatchInferHistoryByID(daos.GetDBMainCtx(ctx), id, map[string][]interface{}{}, false)
	if err != nil {
		return nil, errs.NewError(err)
	}

	return batchInferHistory, nil
}

func (s *Service) JobNameSubmitBatchInferFullPrompt(ctx context.Context, chainConfig *models.ChainConfig) string {
	return fmt.Sprintf("JobSubmitBatchInferFullPrompt-%v", chainConfig.ChainID)
}

func (s *Service) GetJobDuration(durationUnit string) time.Duration {
	switch strings.ToLower(durationUnit) {
	case models.HourText:
		return time.Hour
	case models.MinuteText:
		return time.Minute
	case models.SecondText:
		return time.Second
	case models.MillisecondText:
		return time.Millisecond
	default:
		return time.Minute
	}
}

// func (s *Service) JobSubmitBatchInferFullPrompt(ctx context.Context, chainConfig *models.ChainConfig) error {
// 	interval := 2 * time.Second // minimum job config interval
// 	jobName := s.JobNameSubmitBatchInferFullPrompt(ctx, chainConfig)

// 	for {
// 		time.Sleep(interval)
// 		jobCfg, err := s.dao.FirstJobConfig(daos.GetDBMainCtx(ctx), map[string][]interface{}{
// 			"job_name = ?": {jobName},
// 		}, map[string][]interface{}{}, false)
// 		if err != nil {
// 			return errs.NewError(err)
// 		} else if jobCfg == nil {
// 			now := time.Now()
// 			jobCfg = &models.JobConfig{
// 				JobName:      jobName,
// 				LastRun:      &now,
// 				Interval:     5, // unit: minute
// 				Enable:       false,
// 				IntervalUnit: "minute",
// 			}
// 			err = s.dao.Create(daos.GetDBMainCtx(ctx), jobCfg)
// 			if err != nil {
// 				return errs.NewError(err)
// 			}
// 		}

// 		if jobCfg == nil {
// 			return errs.ErrBadRequest
// 		}
// 		if !jobCfg.Enable {
// 			continue
// 		}

// 		timeInterval := s.GetJobDuration(jobCfg.IntervalUnit)
// 		if time.Since(*jobCfg.LastRun) < time.Duration(jobCfg.Interval)*timeInterval {
// 			continue
// 		}

// 		chainConfig, err := s.dao.FirstChainConfig(daos.GetDBMainCtx(ctx), map[string][]interface{}{
// 			"chain_id = ?": {chainConfig.ChainID},
// 		}, map[string][]interface{}{}, false)

// 		if err != nil {
// 			continue
// 		}
// 		if err := s.DoLogicSubmitBatchInferFullPrompt(ctx, chainConfig); err != nil {
// 			continue
// 		}

// 		now := time.Now()
// 		jobCfg.LastRun = &now
// 		err = s.dao.Save(daos.GetDBMainCtx(ctx), jobCfg)
// 		if err != nil {
// 			return errs.NewError(err)
// 		}
// 	}
// }

// func (s *Service) GetListModelSupportsByChainId(ctx context.Context, chainId uint64) (map[string]string, error) {
// 	chainConfig, err := s.dao.FirstChainConfig(daos.GetDBMainCtx(ctx), map[string][]interface{}{
// 		"chain_id = ?": {chainId},
// 	}, map[string][]interface{}{}, false)

// 	if err != nil {
// 		return nil, errs.NewError(err)
// 	}

// 	if len(chainConfig.SupportModelNames) > 0 {
// 		return chainConfig.SupportModelNames, nil
// 	}

// 	if len(models.MapChainIDToLLMModelAndModelID[chainId]) > 0 {
// 		return models.MapChainIDToLLMModelAndModelID[chainId], nil
// 	}

// 	return map[string]string{}, nil
// }

// func (s *Service) DoLogicSubmitBatchInferFullPrompt(ctx context.Context, chainConfig *models.ChainConfig) error {

// 	filter := map[string][]interface{}{
// 		"status = ?":   {models.BatchInferHistoryStatusPending},
// 		"chain_id = ?": {chainConfig.ChainID},
// 	}

// 	batchInfers, err := s.dao.FindBatchInferHistory(daos.GetDBMainCtx(ctx), filter, map[string][]interface{}{}, []string{"id asc"}, 0, 10000)
// 	if err != nil {
// 		return err
// 	}

// 	var mapModelNameToBatchItem = make(map[string][]*models.BatchInferHistory)
// 	batchInferFailed := make([]*models.BatchInferHistory, 0)
// 	for index, batchInfer := range batchInfers {
// 		chatCompletionRequest := &serializers.ChatCompletionRequest{}
// 		err := json.Unmarshal([]byte(batchInfer.PromptInput), chatCompletionRequest)
// 		if err != nil {
// 			batchInfer.Log = err.Error()
// 			batchInferFailed = append(batchInferFailed, batchInfer)
// 			continue
// 		}

// 		if _, ok := mapModelNameToBatchItem[chatCompletionRequest.Model]; !ok {
// 			mapModelNameToBatchItem[chatCompletionRequest.Model] = make([]*models.BatchInferHistory, 0)
// 		}

// 		mapModelNameToBatchItem[chatCompletionRequest.Model] = append(mapModelNameToBatchItem[chatCompletionRequest.Model], batchInfers[index])
// 	}

// 	for _, batchItem := range batchInferFailed {
// 		s.dao.Save(daos.GetDBMainCtx(ctx), batchItem)
// 	}

// 	for modelName, batchInfersByModelName := range mapModelNameToBatchItem {
// 		if len(batchInfersByModelName) == 0 {
// 			continue
// 		}

// 		modelSupport, err := s.GetListModelSupportsByChainId(ctx, chainConfig.ChainID)
// 		if err != nil {
// 			logger.GetLoggerInstanceFromContext(ctx).Error("JobSubmitBatchInferFullPrompt GetListModelSupportsByChainId",
// 				zap.Any("chainConfig", chainConfig), zap.Any("err", err))
// 			return err
// 		}
// 		modelID := modelSupport[modelName]
// 		if len(modelID) == 0 {
// 			logger.GetLoggerInstanceFromContext(ctx).Error("JobSubmitBatchInferFullPrompt modelID not found in list model support",
// 				zap.Any("chainConfig", chainConfig), zap.Any("modelName", modelName), zap.Any("modelSupport", modelSupport))
// 			continue
// 		}

// 		trainingRequestEntity, err := s.GetTrainingRequestByModelId(ctx, modelID)
// 		if err != nil {
// 			logger.GetLoggerInstanceFromContext(ctx).Error("JobSubmitBatchInferFullPrompt GetTrainingRequestByModelId",
// 				zap.Any("chainConfig", chainConfig), zap.Any("modelID", modelID), zap.Any("err", err))
// 			return err
// 		}

// 		data, _ := json.Marshal(batchInfersByModelName)
// 		hash, err := lighthouse.UploadDataWithRetry(s.GetLightHouseApiKey(), "", data)
// 		if err != nil {
// 			logger.GetLoggerInstanceFromContext(ctx).Error("JobSubmitBatchInferFullPrompt UploadDataWithRetry",
// 				zap.Any("chainConfig", chainConfig),
// 				zap.Any("err", err))
// 			return err
// 		}

// 		inferRequest := &InferRequest{
// 			Prompt:  fmt.Sprintf("ipfs://%v", hash),
// 			ModelId: trainingRequestEntity.ModelID,
// 		}

// 		createInferResp, err := s.UserCreateInfer(ctx, &models.User{}, trainingRequestEntity, inferRequest)
// 		if err != nil {
// 			logger.GetLoggerInstanceFromContext(ctx).Error("[JobSubmitBatchInferFullPrompt UserCreateInfer]",
// 				zap.Any("error", err),
// 				zap.Any("chainConfig", chainConfig))
// 			return err
// 		}

// 		submitInferAt := time.Now().UTC()

// 		for _, infer := range batchInfersByModelName {
// 			err = s.UpdateOneBatchInferHistoryByFilter(ctx, bson.M{
// 				"_id": infer.ID,
// 			}, bson.M{
// 				"status":               models.BatchInferHistoryStatusAgentInferred,
// 				"inscribe_tx_hash":     createInferResp.TxHash,
// 				"infer_id":             fmt.Sprintf("%v", createInferResp.InferID),
// 				"model_id":             modelID,
// 				"prompt_input_hash":    fmt.Sprintf("ipfs://%v", hash),
// 				"infer_wallet_address": createInferResp.UsedWalletAddress,
// 				"submit_infer_at":      submitInferAt,
// 			})
// 			if err != nil {
// 				logger.GetLoggerInstanceFromContext(ctx).Error("JobSubmitBatchInferFullPrompt UploadDataWithRetry",
// 					zap.Any("chainConfig", chainConfig),
// 					zap.Any("infer", infer),
// 					zap.Any("err", err))
// 				return err
// 			}
// 		}
// 	}

// 	return nil
// }
