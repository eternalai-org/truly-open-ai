package services

import (
	"context"
	"strings"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/jinzhu/gorm"
)

func (s *Service) SaveAgentStore(ctx context.Context, req *serializers.AgentStoreReq) error {

	var agentStore *models.AgentStore
	var err error
	err = daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			if req.ID > 0 {
				agentStore, err = s.dao.FirstAgentStoreByID(tx, req.ID, map[string][]interface{}{}, true)
				if err != nil {
					return errs.NewError(err)
				}
				if agentStore == nil {
					return errs.NewError(errs.ErrBadRequest)
				}
				//update name vs description only
				agentStore.Name = req.Name
				agentStore.Description = req.Description
				agentStore.AuthenUrl = req.AuthenUrl
			} else {
				agentStore = &models.AgentStore{
					Name:        req.Name,
					Description: req.Description,
					AuthenUrl:   req.AuthenUrl,
				}
			}
			if err != nil {
				return errs.NewError(err)
			}
			if agentStore.ID > 0 {
				err = s.dao.Save(tx, agentStore)
			} else {
				err = s.dao.Create(tx, agentStore)
			}

			if err != nil {
				return errs.NewError(err)
			}
			return nil
		},
	)
	if err != nil {
		return errs.NewError(err)
	}

	return nil
}

func (s *Service) GetListAgentStore(ctx context.Context, userAddress string, page, limit int) ([]*models.AgentStore, uint, error) {
	res, count, err := s.dao.FindAgentStore4Page(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"user_address = ?": {strings.ToLower(userAddress)},
		},
		map[string][]interface{}{}, []string{"id desc"}, page, limit)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	return res, count, nil
}

func (s *Service) GetAgentStoreDetail(ctx context.Context, id uint) (*models.AgentStore, error) {
	res, err := s.dao.FirstAgentStoreByID(daos.GetDBMainCtx(ctx),
		id,
		map[string][]interface{}{}, false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return res, nil
}

func (s *Service) AddMissionStore(ctx context.Context, agentStoreID uint, missionStoreID uint) error {
	agentStoreMission := &models.AgentStoreMission{
		AgentStoreID:   agentStoreID,
		MissionStoreID: missionStoreID,
	}
	err := s.dao.Create(daos.GetDBMainCtx(ctx), agentStoreMission)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}
