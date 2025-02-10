package services

import (
	"context"
	"fmt"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

func (s *Service) GetAgentStore(ctx context.Context, storeId string) (*models.AgentStore, error) {
	agentStore, err := s.dao.FirstAgentStore(daos.GetDBMainCtx(ctx), map[string][]interface{}{"store_id = ?": {storeId}}, map[string][]interface{}{}, false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	if agentStore == nil {
		return nil, errs.NewError(errs.ErrBadRequest)
	}
	return agentStore, nil
}

func (s *Service) ValidateUserStoreFee(ctx context.Context, apiKey string) (*models.AgentStoreInstall, error) {
	agentStoreInstall, err := s.dao.FirstAgentStoreInstall(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"code = ?": {apiKey},
		},
		map[string][]interface{}{
			"User":       {},
			"AgentStore": {},
		},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	if agentStoreInstall == nil {
		return nil, errs.NewError(errs.ErrBadRequest)
	}
	user := agentStoreInstall.User
	agentStore := agentStoreInstall.AgentStore
	if user.EaiBalance.Float.Cmp(&agentStore.Price.Float) < 0 {
		return nil, errs.NewError(errs.ErrInsufficientBalance)
	}
	return agentStoreInstall, nil
}

func (s *Service) ChargeUserStoreInstall(ctx context.Context, agentStoreInstallID uint, urlPath string, status int) error {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			agentStoreInstall, err := s.dao.FirstAgentStoreInstallByID(
				tx,
				agentStoreInstallID,
				map[string][]interface{}{
					"User":             {},
					"AgentStore.Owner": {},
				},
				false,
			)
			if err != nil {
				return errs.NewError(err)
			}
			if agentStoreInstall == nil {
				return errs.NewError(errs.ErrBadRequest)
			}
			user := agentStoreInstall.User
			agentStoreLog := &models.AgentStoreLog{
				AgentStoreInstallID: agentStoreInstall.ID,
				UserID:              user.ID,
				AgentStoreID:        agentStoreInstall.AgentStoreID,
				Price:               agentStoreInstall.AgentStore.Price,
				UrlPath:             urlPath,
				Status:              status,
			}
			err = s.dao.Create(tx, agentStoreLog)
			if err != nil {
				return errs.NewError(err)
			}
			if agentStoreLog.Status < 300 {
				{
					err = tx.Model(user).
						UpdateColumn("eai_balance", gorm.Expr("eai_balance - ?", agentStoreLog.Price)).
						Error
					if err != nil {
						return errs.NewError(err)
					}
					err = s.dao.Create(
						tx,
						&models.UserTransaction{
							NetworkID: user.NetworkID,
							EventId:   fmt.Sprintf("user_agent_store_log_%d", agentStoreLog.ID),
							UserID:    user.ID,
							Type:      models.UserTransactionTypeUserAgentInfraFee,
							Amount:    numeric.NewBigFloatFromFloat(models.NegativeBigFloat(&agentStoreLog.Price.Float)),
							Status:    models.UserTransactionStatusDone,
						},
					)
					if err != nil {
						return errs.NewError(err)
					}
				}
				{
					owner := agentStoreInstall.AgentStore.Owner
					err = tx.Model(owner).
						UpdateColumn("eai_balance", gorm.Expr("eai_balance + ?", agentStoreLog.Price)).
						Error
					if err != nil {
						return errs.NewError(err)
					}
					err = s.dao.Create(
						tx,
						&models.UserTransaction{
							NetworkID: owner.NetworkID,
							EventId:   fmt.Sprintf("creator_agent_store_log_%d", agentStoreLog.ID),
							UserID:    owner.ID,
							Type:      models.UserTransactionTypeCreatorAgentInfraFee,
							Amount:    agentStoreLog.Price,
							Status:    models.UserTransactionStatusDone,
						},
					)
					if err != nil {
						return errs.NewError(err)
					}
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
