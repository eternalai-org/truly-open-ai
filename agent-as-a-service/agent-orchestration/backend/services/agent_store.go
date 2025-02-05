package services

import (
	"context"
	"encoding/json"

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

func (s *Service) GetListAgentStore(ctx context.Context, page, limit int) ([]*models.AgentStore, uint, error) {
	res, count, err := s.dao.FindAgentStore4Page(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{},
		map[string][]interface{}{
			"AgentStoreMissions": {},
		}, []string{"id desc"}, page, limit)
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

func (s *Service) SaveMissionStore(ctx context.Context, agentStoreID uint, req *serializers.AgentStoreMissionReq) error {
	var mission *models.AgentStoreMission
	var err error
	err = daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			if req.ID > 0 {
				mission, err = s.dao.FirstAgentStoreMissionByID(tx, req.ID, map[string][]interface{}{}, true)
				if err != nil {
					return errs.NewError(err)
				}
				if mission == nil {
					return errs.NewError(errs.ErrBadRequest)
				}
				mission.Name = req.Name
				mission.Description = req.Description
				mission.UserPrompt = req.Prompt
				mission.ToolList = req.ToolList
			} else {
				mission = &models.AgentStoreMission{
					AgentStoreID: agentStoreID,
					Name:         req.Name,
					Description:  req.Description,
					UserPrompt:   req.Prompt,
					Price:        req.Price,
					ToolList:     req.ToolList,
					Icon:         req.Icon,
				}
			}
			if err != nil {
				return errs.NewError(err)
			}
			if mission.ID > 0 {
				err = s.dao.Save(tx, mission)
			} else {
				err = s.dao.Create(tx, mission)
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

func (s *Service) SaveAgentStoreCallback(ctx context.Context, req *serializers.AuthenAgentStoreCallback) error {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			obj, err := s.dao.FirstAgentStoreInstall(tx, map[string][]interface{}{
				"agent_store_id = ?": {req.AgentStoreID},
				"agent_info_id = ?":  {req.InstallAgentInfoID},
			}, map[string][]interface{}{}, true)
			if err != nil {
				return errs.NewError(err)
			}
			params, _ := json.Marshal(req.CallbackParams)
			if obj == nil {
				obj = &models.AgentStoreInstall{
					AgentStoreID:   req.AgentStoreID,
					AgentInfoID:    req.InstallAgentInfoID,
					CallbackParams: string(params),
				}
			} else {
				obj.CallbackParams = string(params)
			}

			if err != nil {
				return errs.NewError(err)
			}
			if obj.ID > 0 {
				err = s.dao.Save(tx, obj)
			} else {
				err = s.dao.Create(tx, obj)
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

func (s *Service) GetListAgentStoreInstall(ctx context.Context, agentInfoID uint, page, limit int) ([]*models.AgentStoreInstall, uint, error) {
	res, count, err := s.dao.FindAgentStoreInstall4Page(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{},
		map[string][]interface{}{
			"AgentStore":                    {},
			"AgentStore.AgentStoreMissions": {},
		}, []string{"id desc"}, page, limit)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	return res, count, nil
}
