package services

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

func (s *Service) GetAgentInfra(ctx context.Context, infraId string) (*models.AgentInfra, error) {
	agentInfra, err := s.dao.FirstAgentInfra(daos.GetDBMainCtx(ctx), map[string][]interface{}{"infra_id = ?": {infraId}}, map[string][]interface{}{}, false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	if agentInfra == nil {
		return nil, errs.NewError(errs.ErrBadRequest)
	}
	return agentInfra, nil
}

func (s *Service) ValidateUserInfraFee(ctx context.Context, apiKey string) (*models.AgentInfraInstall, error) {
	agentInfraInstall, err := s.dao.FirstAgentInfraInstall(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"code = ?": {apiKey},
		},
		map[string][]interface{}{
			"User":       {},
			"AgentInfra": {},
		},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	if agentInfraInstall == nil {
		return nil, errs.NewError(errs.ErrBadRequest)
	}
	user := agentInfraInstall.User
	agentInfra := agentInfraInstall.AgentInfra
	if user.EaiBalance.Float.Cmp(&agentInfra.Price.Float) < 0 {
		return nil, errs.NewError(errs.ErrInsufficientBalance)
	}
	return agentInfraInstall, nil
}

func (s *Service) ChargeUserInfraInstall(ctx context.Context, agentInfraInstallID uint, urlPath string, status int) error {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			agentInfraInstall, err := s.dao.FirstAgentInfraInstallByID(
				tx,
				agentInfraInstallID,
				map[string][]interface{}{
					"User":             {},
					"AgentInfra.Owner": {},
				},
				false,
			)
			if err != nil {
				return errs.NewError(err)
			}
			if agentInfraInstall == nil {
				return errs.NewError(errs.ErrBadRequest)
			}
			user := agentInfraInstall.User
			agentInfraLog := &models.AgentInfraLog{
				AgentInfraInstallID: agentInfraInstall.ID,
				UserID:              user.ID,
				AgentInfraID:        agentInfraInstall.AgentInfraID,
				Price:               agentInfraInstall.AgentInfra.Price,
				UrlPath:             urlPath,
			}
			err = s.dao.Create(tx, agentInfraLog)
			if err != nil {
				return errs.NewError(err)
			}
			if agentInfraLog.Status < 300 {
				{
					err = tx.Model(user).
						UpdateColumn("eai_balance", gorm.Expr("eai_balance - ?", agentInfraLog.Price)).
						Error
					if err != nil {
						return errs.NewError(err)
					}
					err = s.dao.Create(
						tx,
						&models.UserTransaction{
							NetworkID: user.NetworkID,
							EventId:   fmt.Sprintf("user_agent_infra_log_%d", agentInfraLog.ID),
							UserID:    user.ID,
							Type:      models.UserTransactionTypeUserAgentInfraFee,
							Amount:    numeric.NewBigFloatFromFloat(models.NegativeBigFloat(&agentInfraLog.Price.Float)),
							Status:    models.UserTransactionStatusDone,
						},
					)
					if err != nil {
						return errs.NewError(err)
					}
				}
				{
					owner := agentInfraInstall.AgentInfra.Owner
					err = tx.Model(owner).
						UpdateColumn("eai_balance", gorm.Expr("eai_balance + ?", agentInfraLog.Price)).
						Error
					if err != nil {
						return errs.NewError(err)
					}
					err = s.dao.Create(
						tx,
						&models.UserTransaction{
							NetworkID: owner.NetworkID,
							EventId:   fmt.Sprintf("creator_agent_infra_log_%d", agentInfraLog.ID),
							UserID:    owner.ID,
							Type:      models.UserTransactionTypeCreatorAgentInfraFee,
							Amount:    agentInfraLog.Price,
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

func (s *Service) CreateOrUpdateAgentInfra(ctx context.Context, userAddress string, req *serializers.AgentInfraReq) (*models.AgentInfra, error) {
	var agentInfra *models.AgentInfra
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			hostURL, err := url.Parse(req.ApiUrl)
			if err != nil {
				return errs.NewError(err)
			}
			if hostURL.Scheme != "https" {
				return errs.NewError(errs.ErrBadRequest)
			}
			if !strings.Contains(hostURL.Host, ".") {
				return errs.NewError(errs.ErrBadRequest)
			}
			err = helpers.CurlURL(req.ApiUrl+"/health", http.MethodGet, map[string]string{}, nil, nil)
			if err != nil {
				return errs.NewError(errs.ErrApiUrlNotHealth)
			}
			user, err := s.GetUser(tx, 0, userAddress, false)
			if err != nil {
				return errs.NewError(err)
			}
			if req.ID > 0 {
				agentInfra, err = s.dao.FirstAgentInfraByID(tx, req.ID, map[string][]interface{}{}, true)
				if err != nil {
					return errs.NewError(err)
				}
				if agentInfra == nil {
					return errs.NewError(errs.ErrBadRequest)
				}
				if agentInfra.OwnerID != user.ID {
					return errs.NewError(errs.ErrBadRequest)
				}
				agentInfra.Name = req.Name
				agentInfra.Description = req.Description
				agentInfra.Icon = req.Icon
				agentInfra.Status = req.Status
				agentInfra.ApiUrl = req.ApiUrl
				agentInfra.Price = req.Price
				agentInfra.Docs = req.Docs
			} else {
				agentInfra = &models.AgentInfra{
					InfraId:      helpers.RandomBigInt(12).Text(16),
					OwnerAddress: user.Address,
					Name:         req.Name,
					Description:  req.Description,
					Icon:         req.Icon,
					OwnerID:      user.ID,
					Status:       req.Status,
					ApiUrl:       req.ApiUrl,
					Price:        req.Price,
					Docs:         req.Docs,
				}
			}
			if err != nil {
				return errs.NewError(err)
			}
			if agentInfra.ID > 0 {
				err = s.dao.Save(tx, agentInfra)
			} else {
				err = s.dao.Create(tx, agentInfra)
			}
			if err != nil {
				return errs.NewError(err)
			}
			return nil
		},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return agentInfra, nil
}

func (s *Service) CreateAgentInfraInstallCode(ctx context.Context, userAddress string, agentInfraID uint) (*models.AgentInfraInstall, error) {
	user, err := s.GetUser(daos.GetDBMainCtx(ctx), 0, userAddress, false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	agentInfra, err := s.dao.FirstAgentInfraByID(daos.GetDBMainCtx(ctx), agentInfraID, map[string][]interface{}{}, true)
	if err != nil {
		return nil, errs.NewError(err)
	}
	if agentInfra == nil {
		return nil, errs.NewError(errs.ErrBadRequest)
	}
	agentInfraInstall, err := s.dao.FirstAgentInfraInstall(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"agent_infra_id = ?": {agentInfra.ID},
			"user_id = ?":        {user.ID},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	if agentInfraInstall == nil {
		agentInfraInstall = &models.AgentInfraInstall{
			Code:         helpers.RandomStringWithLength(64),
			AgentInfraID: agentInfraID,
			UserID:       user.ID,
			Status:       models.AgentInfraInstallStatusNew,
		}
		err = s.dao.Create(daos.GetDBMainCtx(ctx), agentInfraInstall)
		if err != nil {
			return nil, errs.NewError(err)
		}
	}
	return agentInfraInstall, nil
}

func (s *Service) GetListAgentInfra(ctx context.Context, page, limit int) ([]*models.AgentInfra, uint, error) {
	res, count, err := s.dao.FindAgentInfra4Page(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"status = ?": {models.AgentInfraStatusActived},
		},
		map[string][]interface{}{},
		[]string{"id desc"},
		page,
		limit,
	)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	return res, count, nil
}

func (s *Service) GetListAgentInfraByUser(ctx context.Context, userAddress string, page, limit int) ([]*models.AgentInfra, uint, error) {
	res, count, err := s.dao.FindAgentInfra4Page(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"owner_address = ?": {strings.ToLower(userAddress)},
		},
		map[string][]interface{}{},
		[]string{"id desc"},
		page,
		limit,
	)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	return res, count, nil
}

func (s *Service) GetAgentInfraDetail(ctx context.Context, id uint) (*models.AgentInfra, error) {
	res, err := s.dao.FirstAgentInfraByID(daos.GetDBMainCtx(ctx),
		id,
		map[string][]interface{}{}, false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return res, nil
}

func (s *Service) GetListAgentInfraInstallByUser(ctx context.Context, userAddress string, page, limit int) ([]*models.AgentInfraInstall, uint, error) {
	user, err := s.GetUser(daos.GetDBMainCtx(ctx), 0, userAddress, false)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	filter := map[string][]interface{}{
		"user_id = ?": {user.ID},
		"status = ?":  {models.AgentInfraInstallStatusDone},
	}
	res, count, err := s.dao.FindAgentInfraInstall4Page(daos.GetDBMainCtx(ctx),
		filter,
		map[string][]interface{}{
			"AgentInfra": {},
		},
		[]string{"id desc"},
		page,
		limit,
	)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	return res, count, nil
}
