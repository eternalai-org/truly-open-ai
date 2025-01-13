package services

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/jinzhu/gorm"
)

func (s *Service) UploadMisstionStore(ctx context.Context, req *serializers.MissionStoreReq) error {

	var missionStore *models.MissionStore
	var err error
	err = daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			if req.ID > 0 {
				missionStore, err = s.dao.FirstMissionStoreByID(tx, req.ID, map[string][]interface{}{}, true)
				if err != nil {
					return errs.NewError(err)
				}
				if missionStore == nil {
					return errs.NewError(errs.ErrBadRequest)
				}
				//update name vs description only
				missionStore.Name = req.Name
				missionStore.Description = req.Description
			} else {
				params, _ := json.Marshal(req.Params)
				missionStore = &models.MissionStore{
					Name:         req.Name,
					Description:  req.Description,
					UserPrompt:   req.Prompt,
					Price:        req.Price,
					OwnerAddress: req.OwnerAddress,
					ToolList:     req.ToolList,
					Icon:         req.Icon,
					OutputType:   models.OutputType(req.OutputType),
					Params:       string(params),
				}
			}
			if err != nil {
				return errs.NewError(err)
			}
			if missionStore.ID > 0 {
				err = s.dao.Save(tx, missionStore)
			} else {
				err = s.dao.Create(tx, missionStore)
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

func (s *Service) RatingMisstionStore(ctx context.Context, req *serializers.MissionStoreRatingReq) error {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			if req.HistoryID == 0 {
				return errs.NewError(errs.ErrBadRequest)

			}
			history, err := s.dao.FirstMissionStoreHistoryByID(tx, req.HistoryID, map[string][]interface{}{}, true)
			if err != nil {
				return errs.NewError(err)
			}
			if history == nil || history.IsRated || !strings.EqualFold(strings.ToLower(history.UserAddress), strings.ToLower(req.UserAddress)) {
				return errs.NewError(errs.ErrBadRequest)
			}

			missionStore, err := s.dao.FirstMissionStoreByID(tx, history.MissionStoreID, map[string][]interface{}{}, true)
			if err != nil {
				return errs.NewError(err)
			}
			if missionStore == nil {
				return errs.NewError(errs.ErrBadRequest)
			}
			rating := &models.MissionStoreRating{
				UserAddress:         history.UserAddress,
				MissionStoreID:      missionStore.ID,
				AgentSnapshotPostID: history.ID,
				Rating:              req.Rating,
				Comment:             req.Comment,
			}

			err = s.dao.Create(tx, rating)
			if err != nil {
				return errs.NewError(err)
			}
			//update history
			history.IsRated = true
			err = s.dao.Save(tx, history)
			if err != nil {
				return errs.NewError(err)
			}
			//update missionstore
			newRate := (missionStore.Rating*float64(missionStore.NumRating) + req.Rating) / (float64(missionStore.NumRating) + 1)
			missionStore.Rating = newRate
			missionStore.NumRating += 1
			err = s.dao.Save(tx, missionStore)
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

func (s *Service) GetListMisstionStore(ctx context.Context, search string, page, limit int) ([]*models.MissionStore, uint, error) {
	res, count, err := s.dao.FindMissionStore4Page(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"name like ?": {"%" + search + "%"},
		},
		map[string][]interface{}{}, []string{"rating desc"}, page, limit)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	return res, count, nil
}

func (s *Service) GetMisstionStoreDetail(ctx context.Context, id uint) (*models.MissionStore, error) {
	res, err := s.dao.FirstMissionStoreByID(daos.GetDBMainCtx(ctx),
		id,
		map[string][]interface{}{}, false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return res, nil
}

func (s *Service) ClaimFeeMisstionStore(ctx context.Context) error {

	return nil
}

func (s *Service) GetMissionStoreRating(ctx context.Context, id uint, page, limit int) ([]*models.MissionStoreRating, uint, error) {
	res, count, err := s.dao.FindMissionStoreRating4Page(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"mission_store_id = ?": {id},
		},
		map[string][]interface{}{}, []string{"rating desc"}, page, limit)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	return res, count, nil
}

func (s *Service) GetMissionStoreHistory(ctx context.Context, id uint, page, limit int) ([]*models.MissionStoreHistory, uint, error) {
	res, count, err := s.dao.FindMissionStoreHistory4Page(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"mission_store_id = ?": {id},
		},
		map[string][]interface{}{}, []string{"id desc"}, page, limit)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	return res, count, nil
}

func (s *Service) addToolPostTwitter(toollist string, appendTool string) (string, error) {
	var initialData []map[string]interface{}
	if err := json.Unmarshal([]byte(toollist), &initialData); err != nil {
		return toollist, errs.NewError(err)
	}

	var appendData map[string]interface{}
	if err := json.Unmarshal([]byte(appendTool), &appendData); err != nil {
		return toollist, errs.NewError(err)
	}

	initialData = append(initialData, appendData)

	updatedJSON, err := json.Marshal(initialData)
	if err != nil {
		return toollist, errs.NewError(err)
	}
	return string(updatedJSON), nil
}
