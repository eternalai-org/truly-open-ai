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
				missionStore.Name = req.Name
				missionStore.Description = req.Description
				missionStore.UserPrompt = req.Prompt
				missionStore.Price = req.Price
				missionStore.ToolList = req.ToolList
				missionStore.DurationDay = req.DurationDay
			} else {
				depositAddress, err := s.CreateETHAddress(ctx)
				if err != nil {
					return errs.NewError(err)
				}
				missionStore = &models.MissionStore{
					Name:           req.Name,
					Description:    req.Description,
					UserPrompt:     req.Prompt,
					Price:          req.Price,
					OwnerAddress:   req.OwnerAddress,
					ToolList:       req.ToolList,
					DepositAddress: depositAddress,
					DurationDay:    req.DurationDay,
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
				UserAddress:           history.UserAddress,
				MissionStoreID:        missionStore.ID,
				MissionStoreHistoryID: history.ID,
				Rating:                req.Rating,
				Comment:               req.Comment,
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

func (s *Service) GetListMisstionStore(ctx context.Context, page, limit int) ([]*models.MissionStore, uint, error) {
	res, count, err := s.dao.FindMissionStore4Page(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{},
		map[string][]interface{}{}, []string{"rating desc"}, page, limit)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	return res, count, nil
}

func (s *Service) ClaimFeeMisstionStore(ctx context.Context) error {

	return nil
}
