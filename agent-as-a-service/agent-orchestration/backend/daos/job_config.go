package daos

import (
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/jinzhu/gorm"
)

func (d *DAO) FindJobConfigJoinSelect(tx *gorm.DB, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.JobConfig, error) {
	var ms []*models.JobConfig
	err := d.findJoinSelect(tx, &models.JobConfig{}, &ms, selected, joins, filters, preloads, orders, uint(page), uint(limit), false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return ms, nil
}

func (d *DAO) FindJobConfigJoinSelect4Page(tx *gorm.DB, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.JobConfig, uint, error) {
	var ms []*models.JobConfig
	c, err := d.findJoinSelect4Page(tx, &models.JobConfig{}, &ms, selected, joins, filters, preloads, orders, uint(page), uint(limit), false)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	return ms, c, nil
}

func (d *DAO) FirstJobConfigJoinSelect(tx *gorm.DB, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, forUpdate bool) (*models.JobConfig, error) {
	var m models.JobConfig
	if err := d.firstJoinSelect(tx, &m, selected, joins, filters, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FirstJobConfigByID(tx *gorm.DB, id uint, preloads map[string][]interface{}, forUpdate bool) (*models.JobConfig, error) {
	var m models.JobConfig
	if err := d.first(tx, &m, map[string][]interface{}{"id = ?": []interface{}{id}}, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FirstJobConfig(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, forUpdate bool) (*models.JobConfig, error) {
	var m models.JobConfig
	if err := d.first(tx, &m, filters, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FindJobConfig(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, offset int, limit int) ([]*models.JobConfig, error) {
	var ms []*models.JobConfig
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, err
	}
	return ms, nil
}

func (d *DAO) FindJobConfig4Page(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.JobConfig, uint, error) {
	var (
		offset = (page - 1) * limit
	)
	var ms []*models.JobConfig
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, 0, errs.NewError(err)
	}
	c, err := d.count(tx, &models.JobConfig{}, filters)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	return ms, c, nil
}
