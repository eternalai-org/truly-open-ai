package services

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/logger"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func (s *Service) AgentTwinTrain(ctx context.Context, agentInfoID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("AgentTwinTrain_%d", agentInfoID),
		func() error {
			agent, err := s.dao.FirstAgentInfoByID(
				daos.GetDBMainCtx(ctx),
				agentInfoID,
				map[string][]interface{}{},
				false,
			)
			if err != nil {
				return errs.NewError(err)
			}

			_, isReadyRunTwin, err := s.CheckAgentIsReadyToRunTwinTraining(agent)
			if !isReadyRunTwin {
				return nil
			}

			if agent.TwinTwitterUsernames == "" {
				return nil
			}

			type CallTwinTrainingRequest struct {
				AgentID    string   `json:"agent_id"`
				TweeterIds []string `json:"twitter_ids"`
			}

			if agent.TwinStatus == models.TwinStatusPending {
				var twinFee *big.Float
				arr := strings.Split(agent.TwinTwitterUsernames, ",")

				twinFee = numeric.NewFloatFromString(fmt.Sprintf("%v", float64(len(arr))*300))

				updateAgentFields := map[string]interface{}{
					"twin_status": models.TwinStatusRunning,
				}
				arrTwitterIds := []string{}
				for _, twitterUsername := range arr {
					twitterInfo, err := s.GetTwitterUserByUsername(ctx, twitterUsername)
					if err != nil || twitterInfo == nil {
						logger.Error("twin_train_error", "s.GetTwitterUserByUsername got error",
							zap.Any("twitter_username", twitterUsername), zap.Any("error", err),
							zap.Any("twitterInfo", twitterInfo),
						)
						continue
					}

					arrTwitterIds = append(arrTwitterIds, twitterInfo.ID)
				}
				if len(arrTwitterIds) == 0 {
					err = daos.GetDBMainCtx(ctx).Model(agent).
						Updates(map[string]interface{}{
							//"twin_status":               models.TwinStatusDoneError,
							"twin_call_process_request": "twitter_id list is empty",
						}).Error
					if err != nil {
						return errs.NewError(err)
					}

					return nil
				}

				twinTrainRequest := CallTwinTrainingRequest{
					AgentID:    fmt.Sprintf("%v", agent.ID),
					TweeterIds: arrTwitterIds,
				}
				requestBytes, _ := json.Marshal(twinTrainRequest)
				updateAgentFields["twin_call_process_request"] = string(requestBytes)

				body, err := helpers.CurlURLString(s.conf.AgentOffchainUrl+"/v1/twin/submit", "POST", map[string]string{
					"Content-Type": "application/json",
				}, &twinTrainRequest)
				if err != nil {
					return err
				}
				updateAgentFields["twin_call_process_response"] = body

				twinStartTrainingAt := time.Now()
				agent.TwinStartTrainingAt = &twinStartTrainingAt
				updateAgentFields["twin_start_training_at"] = *agent.TwinStartTrainingAt

				agent.TwinStatus = models.TwinStatusRunning
				agent.TwinCallProcessRequest = string(requestBytes)
				agent.TwinCallProcessResponse = body

				estimateDoneTime := time.Now().Add(20 * time.Minute)
				agent.EstimateTwinDoneTimestamp = &estimateDoneTime
				updateAgentFields["estimate_twin_done_timestamp"] = *agent.TwinStartTrainingAt

				err = daos.GetDBMainCtx(ctx).Model(agent).
					Updates(updateAgentFields).Error
				if err != nil {
					return errs.NewError(err)
				}

				err = daos.GetDBMainCtx(ctx).
					Model(agent).
					Updates(
						map[string]interface{}{
							"eai_balance": gorm.Expr("eai_balance - ?", numeric.NewBigFloatFromFloat(twinFee)),
							"twin_fee":    numeric.NewBigFloatFromFloat(twinFee),
						},
					).
					Error
				if err != nil {
					return errs.NewError(err)
				}
				if twinFee.Cmp(big.NewFloat(0)) > 0 {
					_ = s.dao.Create(
						daos.GetDBMainCtx(ctx),
						&models.AgentEaiTopup{
							NetworkID:      agent.NetworkID,
							EventId:        fmt.Sprintf("twin_train_fee_%d", agent.ID),
							AgentInfoID:    agent.ID,
							Type:           models.AgentEaiTopupTypeSpent,
							Amount:         numeric.NewBigFloatFromFloat(twinFee),
							Status:         models.AgentEaiTopupStatusDone,
							DepositAddress: agent.ETHAddress,
							ToAddress:      agent.ETHAddress,
							Toolset:        "twin_train_fee",
						},
					)
				}
			}

			return nil
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}
