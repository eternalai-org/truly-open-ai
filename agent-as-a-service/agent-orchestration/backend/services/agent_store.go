package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/jinzhu/gorm"
)

func (s *Service) SaveAgentStore(ctx context.Context, userAddress string, req *serializers.AgentStoreReq) (*models.AgentStore, error) {
	var agentStore *models.AgentStore
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			user, err := s.GetUser(tx, 0, userAddress, false)
			if err != nil {
				return errs.NewError(err)
			}
			if req.ID > 0 {
				agentStore, err = s.dao.FirstAgentStoreByID(tx, req.ID, map[string][]interface{}{}, true)
				if err != nil {
					return errs.NewError(err)
				}
				if agentStore == nil {
					return errs.NewError(errs.ErrBadRequest)
				}
				if agentStore.OwnerID != user.ID {
					return errs.NewError(errs.ErrBadRequest)
				}
				//update name vs description only
				agentStore.Name = req.Name
				agentStore.Description = req.Description
				agentStore.AuthenUrl = req.AuthenUrl
				agentStore.Icon = req.Icon
			} else {
				agentStore = &models.AgentStore{
					OwnerAddress: userAddress,
					Name:         req.Name,
					Description:  req.Description,
					AuthenUrl:    req.AuthenUrl,
					Icon:         req.Icon,
					OwnerID:      user.ID,
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
		return nil, errs.NewError(err)
	}
	return agentStore, nil
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
		map[string][]interface{}{
			"AgentStoreMissions": {},
		}, false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return res, nil
}

func (s *Service) SaveMissionStore(ctx context.Context, userAddress string, agentStoreID uint, req *serializers.AgentStoreMissionReq) error {
	var mission *models.AgentStoreMission
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			agentStore, err := s.dao.FirstAgentStoreByID(tx, req.ID, map[string][]interface{}{}, true)
			if err != nil {
				return errs.NewError(err)
			}
			if agentStore == nil {
				return errs.NewError(errs.ErrBadRequest)
			}
			if !strings.EqualFold(agentStore.OwnerAddress, userAddress) {
				return errs.NewError(errs.ErrBadRequest)
			}
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
			obj, err := s.dao.FirstAgentStoreInstall(
				tx,
				map[string][]interface{}{
					"code = ?": {req.Code},
				},
				map[string][]interface{}{},
				[]string{},
			)
			if err != nil {
				return errs.NewError(err)
			}
			params, _ := json.Marshal(req.CallbackParams)
			if obj == nil {
				return errs.NewError(errs.ErrBadRequest)
			} else {
				obj.CallbackParams = string(params)
				obj.Status = models.InstallStatusDone
			}

			err = s.dao.Save(tx, obj)

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

func (s *Service) GetListAgentStoreInstall(ctx context.Context, userAddress string, agentInfoID uint, page, limit int) ([]*models.AgentStoreInstall, uint, error) {
	filter := map[string][]interface{}{
		"status = ?": {models.InstallStatusDone},
	}
	if agentInfoID > 0 {
		filter["agent_info_id = ?"] = []interface{}{agentInfoID}
	} else {
		filter["user_address = ?"] = []interface{}{strings.ToLower(userAddress)}
	}
	res, count, err := s.dao.FindAgentStoreInstall4Page(daos.GetDBMainCtx(ctx),
		filter,
		map[string][]interface{}{
			"AgentStore":                    {},
			"AgentStore.AgentStoreMissions": {},
		}, []string{"id desc"}, page, limit)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	return res, count, nil
}

func (s *Service) CreateAgentStoreInstallCode(ctx context.Context, userAddress string, agentStoreID, agentInfoID uint) (*models.AgentStoreInstall, error) {
	if agentInfoID > 0 {
		agentInfo, err := s.dao.FirstAgentInfoByID(daos.GetDBMainCtx(ctx), agentInfoID, map[string][]interface{}{}, true)
		if err != nil {
			return nil, errs.NewError(err)
		}
		if agentInfo == nil {
			return nil, errs.NewError(errs.ErrBadRequest)
		}

		if !strings.EqualFold(agentInfo.Creator, userAddress) {
			return nil, errs.NewError(errs.ErrBadRequest)
		}
	}
	obj := &models.AgentStoreInstall{
		Code:         helpers.RandomReferralCode(32),
		AgentStoreID: agentStoreID,
		AgentInfoID:  agentInfoID,
		Status:       models.InstallStatusNew,
		Type:         models.InstallTypeAgent,
	}
	if agentInfoID == 0 {
		obj.UserAddress = strings.ToLower(userAddress)
		obj.Type = models.InstallTypeUser
	}
	err := s.dao.Create(daos.GetDBMainCtx(ctx), obj)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return obj, nil
}

func (s *Service) GetMissionStoreResult(ctx context.Context, responseID string) (string, error) {
	var resp string
	cacheKey := fmt.Sprintf(`CacheAgentSnapshotPost_%s`, strings.ToLower(responseID))
	err := s.GetRedisCachedWithKey(cacheKey, &resp)
	if err != nil {
		s.CacheMissionStoreResult(daos.GetDBMainCtx(ctx), responseID)
		s.GetRedisCachedWithKey(cacheKey, &resp)
	}

	return resp, nil
}

func (s *Service) CacheMissionStoreResult(tx *gorm.DB, responseID string) error {
	snapshotPost, err := s.dao.FirstAgentSnapshotPost(
		tx,
		map[string][]interface{}{
			"response_id = ?": {responseID},
		},
		map[string][]interface{}{
			"AgentSnapshotMission": {},
			"AgentInfo":            {},
		},
		[]string{},
	)
	if err != nil {
		return errs.NewError(err)
	}

	err = s.CacheAgentSnapshotPost(snapshotPost)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) CacheAgentSnapshotPost(snapshotPost *models.AgentSnapshotPost) error {
	if snapshotPost != nil {
		cacheData, err := json.Marshal(&serializers.Resp{Result: serializers.NewAgentSnapshotPostResp(snapshotPost)})
		if err != nil {
			errs.NewError(err)
		}
		err = s.SetRedisCachedWithKey(
			fmt.Sprintf(`CacheAgentSnapshotPost_%s`, strings.ToLower(snapshotPost.ResponseId)),
			string(cacheData),
			1*time.Hour,
		)

		if err != nil {
			return errs.NewError(err)
		}
	}
	return nil
}
