package daos

import (
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/jinzhu/gorm"
)

func (d *DAO) FindTrainingRequestJoinSelect(tx *gorm.DB, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.TrainingRequest, error) {
	var ms []*models.TrainingRequest
	err := d.findJoinSelect(tx, &models.TrainingRequest{}, &ms, selected, joins, filters, preloads, orders, uint(page), uint(limit), false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return ms, nil
}

func (d *DAO) FindTrainingRequestJoinSelect4Page(tx *gorm.DB, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.TrainingRequest, uint, error) {
	var ms []*models.TrainingRequest
	c, err := d.findJoinSelect4Page(tx, &models.TrainingRequest{}, &ms, selected, joins, filters, preloads, orders, uint(page), uint(limit), false)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	return ms, c, nil
}

func (d *DAO) FirstTrainingRequestJoinSelect(tx *gorm.DB, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, forUpdate bool) (*models.TrainingRequest, error) {
	var m models.TrainingRequest
	if err := d.firstJoinSelect(tx, &m, selected, joins, filters, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FirstTrainingRequestByID(tx *gorm.DB, id uint, preloads map[string][]interface{}, forUpdate bool) (*models.TrainingRequest, error) {
	var m models.TrainingRequest
	if err := d.first(tx, &m, map[string][]interface{}{"id = ?": []interface{}{id}}, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FirstTrainingRequest(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, forUpdate bool) (*models.TrainingRequest, error) {
	var m models.TrainingRequest
	if err := d.first(tx, &m, filters, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FindTrainingRequest(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, offset int, limit int) ([]*models.TrainingRequest, error) {
	var ms []*models.TrainingRequest
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, err
	}
	return ms, nil
}

func (d *DAO) FindTrainingRequest4Page(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.TrainingRequest, uint, error) {
	var (
		offset = (page - 1) * limit
	)
	var ms []*models.TrainingRequest
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, 0, errs.NewError(err)
	}
	c, err := d.count(tx, &models.TrainingRequest{}, filters)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	return ms, c, nil
}
