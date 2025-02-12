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
			if req.Type == "" {
				req.Type = models.AgentStoreTypeStore
			}
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
			} else {
				agentStore = &models.AgentStore{
					Type:         req.Type,
					StoreId:      helpers.RandomBigInt(12).Text(16),
					OwnerID:      user.ID,
					OwnerAddress: user.Address,
					Status:       models.AgentStoreStatusActived,
				}
			}
			agentStore.Name = req.Name
			agentStore.Description = req.Description
			agentStore.AuthenUrl = req.AuthenUrl
			agentStore.Icon = req.Icon
			agentStore.Docs = req.Docs
			agentStore.Price = req.Price
			if req.Status != "" {
				agentStore.Status = req.Status
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

func (s *Service) GetListAgentStore(ctx context.Context, search, types string, page, limit int) ([]*models.AgentStore, uint, error) {
	filters := map[string][]interface{}{
		"status = ?": {models.AgentStoreStatusActived},
	}
	if types != "" {
		filters["type in (?)"] = []interface{}{strings.Split(types, ",")}
	}
	if search != "" {
		filters["name like ?"] = []interface{}{"%" + search + "%"}
	}
	res, count, err := s.dao.FindAgentStore4Page(
		daos.GetDBMainCtx(ctx),
		filters,
		map[string][]interface{}{
			"AgentStoreMissions": {},
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

func (s *Service) GetListAgentStoreByOwner(ctx context.Context, userAddress string, page, limit int) ([]*models.AgentStore, uint, error) {
	user, err := s.GetUser(daos.GetDBMainCtx(ctx), 0, userAddress, false)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	res, count, err := s.dao.FindAgentStore4Page(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"owner_id = ?": {user.ID},
		},
		map[string][]interface{}{
			"AgentStoreMissions": {},
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

func (s *Service) GetAgentStoreDetail(ctx context.Context, id uint) (*models.AgentStore, error) {
	res, err := s.dao.FirstAgentStoreByID(daos.GetDBMainCtx(ctx),
		id,
		map[string][]interface{}{
			"AgentStoreMissions": {},
		},
		false,
	)
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
			agentStore, err := s.dao.FirstAgentStoreByID(tx, agentStoreID, map[string][]interface{}{}, true)
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
				if req.Status != "" {
					mission.Status = models.AgentStoreStatus(req.Status)
				}
			} else {
				mission = &models.AgentStoreMission{
					AgentStoreID: agentStoreID,
					Name:         req.Name,
					Description:  req.Description,
					UserPrompt:   req.Prompt,
					Price:        req.Price,
					ToolList:     req.ToolList,
					Icon:         req.Icon,
					Status:       models.AgentStoreStatusActived,
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
			agentStore, err := s.dao.FirstAgentStoreByID(tx, obj.AgentStoreID, map[string][]interface{}{}, true)
			if err != nil {
				return errs.NewError(err)
			}
			if agentStore == nil {
				return errs.NewError(errs.ErrBadRequest)
			}
			params, _ := json.Marshal(req.CallbackParams)
			if obj == nil {
				return errs.NewError(errs.ErrBadRequest)
			} else {
				obj.CallbackParams = string(params)
				obj.Status = models.AgentStoreInstallStatusDone
			}

			err = s.dao.Save(tx, obj)
			if err != nil {
				return errs.NewError(err)
			}
			//
			agentStore.NumInstall = agentStore.NumInstall + 1
			err = s.dao.Save(tx, agentStore)
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
	user, err := s.GetUser(daos.GetDBMainCtx(ctx), 0, userAddress, false)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	filter := map[string][]interface{}{
		"status = ?": {models.AgentStoreInstallStatusDone},
	}
	if agentInfoID > 0 {
		filter["agent_info_id = ?"] = []interface{}{agentInfoID}
	} else {
		filter["user_id = ?"] = []interface{}{user.ID}
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
	}
	user, err := s.GetUser(daos.GetDBMainCtx(ctx), 0, userAddress, false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	agentStoreInstall, err := s.dao.FirstAgentStoreInstall(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"user_id = ?":        {user.ID},
			"agent_store_id = ?": {agentStoreID},
			"agent_info_id = ?":  {agentInfoID},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	if agentStoreInstall == nil {
		agentStoreInstall = &models.AgentStoreInstall{
			Code:         helpers.RandomStringWithLength(64),
			AgentStoreID: agentStoreID,
			AgentInfoID:  agentInfoID,
			Status:       models.AgentStoreInstallStatusNew,
			Type:         models.AgentStoreInstallTypeAgent,
			UserID:       user.ID,
		}
		if agentInfoID == 0 {
			agentStoreInstall.Type = models.AgentStoreInstallTypeUser
		}
		err = s.dao.Create(daos.GetDBMainCtx(ctx), agentStoreInstall)
		if err != nil {
			return nil, errs.NewError(err)
		}
	}
	return agentStoreInstall, nil
}

func (s *Service) GetMissionStoreResult(ctx context.Context, userAddress string, responseID string) (string, error) {
	user, err := s.GetUser(daos.GetDBMainCtx(ctx), 0, userAddress, false)
	if err != nil {
		return "", errs.NewError(err)
	}
	snapshotPost, err := s.dao.FirstAgentSnapshotPost(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"response_id = ?": {responseID},
		},
		map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return "", errs.NewError(err)
	}
	if user.ID != snapshotPost.UserID {
		return "", errs.NewError(errs.ErrBadRequest)
	}
	err = s.UpdateOffchainAutoOutputV2ForId(ctx, snapshotPost.ID)
	if err != nil {
		return "", errs.NewError(err)
	}
	snapshotPost, err = s.dao.FirstAgentSnapshotPost(
		daos.GetDBMainCtx(ctx),
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
		return "", errs.NewError(err)
	}
	return helpers.ConvertJsonString(&serializers.Resp{Result: serializers.NewAgentSnapshotPostResp(snapshotPost)}), nil
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
