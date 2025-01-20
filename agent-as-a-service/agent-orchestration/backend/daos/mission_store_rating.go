package daos

import (
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/jinzhu/gorm"
)

func (d *DAO) FindMissionStoreRatingJoinSelect(tx *gorm.DB, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.MissionStoreRating, error) {
	var ms []*models.MissionStoreRating
	err := d.findJoinSelect(tx, &models.MissionStoreRating{}, &ms, selected, joins, filters, preloads, orders, uint(page), uint(limit), false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return ms, nil
}

func (d *DAO) FindMissionStoreRatingJoinSelect4Page(tx *gorm.DB, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.MissionStoreRating, uint, error) {
	var ms []*models.MissionStoreRating
	c, err := d.findJoinSelect4Page(tx, &models.MissionStoreRating{}, &ms, selected, joins, filters, preloads, orders, uint(page), uint(limit), false)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	return ms, c, nil
}

func (d *DAO) FirstMissionStoreRatingJoinSelect(tx *gorm.DB, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, forUpdate bool) (*models.MissionStoreRating, error) {
	var m models.MissionStoreRating
	if err := d.firstJoinSelect(tx, &m, selected, joins, filters, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FirstMissionStoreRatingByID(tx *gorm.DB, id uint, preloads map[string][]interface{}, forUpdate bool) (*models.MissionStoreRating, error) {
	var m models.MissionStoreRating
	if err := d.first(tx, &m, map[string][]interface{}{"id = ?": []interface{}{id}}, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FirstMissionStoreRating(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, forUpdate bool) (*models.MissionStoreRating, error) {
	var m models.MissionStoreRating
	if err := d.first(tx, &m, filters, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FindMissionStoreRating(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, offset int, limit int) ([]*models.MissionStoreRating, error) {
	var ms []*models.MissionStoreRating
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, err
	}
	return ms, nil
}

func (d *DAO) FindMissionStoreRating4Page(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.MissionStoreRating, uint, error) {
	var (
		offset = (page - 1) * limit
	)
	var ms []*models.MissionStoreRating
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, 0, errs.NewError(err)
	}
	c, err := d.count(tx, &models.MissionStoreRating{}, filters)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	return ms, c, nil
}
