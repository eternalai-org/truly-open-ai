package daos

import (
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/jinzhu/gorm"
)

func (d *DAO) FirstApiSubscriptionKeyByID(tx *gorm.DB, id uint, preloads map[string][]interface{}, forUpdate bool) (*models.ApiSubscriptionKey, error) {
	var m models.ApiSubscriptionKey
	if err := d.first(tx, &m, map[string][]interface{}{"id = ?": []interface{}{id}}, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FirstApiSubscriptionKey(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string) (*models.ApiSubscriptionKey, error) {
	var m models.ApiSubscriptionKey
	if err := d.first(tx, &m, filters, preloads, orders, false); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FindApiSubscriptionKey(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, offset int, limit int) ([]*models.ApiSubscriptionKey, error) {
	var ms []*models.ApiSubscriptionKey
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, err
	}
	return ms, nil
}

func (d *DAO) FindApiSubscriptionKey4Page(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.ApiSubscriptionKey, uint, error) {
	var (
		offset = (page - 1) * limit
	)
	var ms []*models.ApiSubscriptionKey
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, 0, errs.NewError(err)
	}
	c, err := d.count(tx, &models.ApiSubscriptionKey{}, filters)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	return ms, c, nil
}

func (d *DAO) FindApiSubscriptionKeyJoinSelect(tx *gorm.DB, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.ApiSubscriptionKey, error) {
	var ms []*models.ApiSubscriptionKey
	err := d.findJoinSelect(tx, &models.ApiSubscriptionKey{}, &ms, selected, joins, filters, preloads, orders, uint(page), uint(limit), false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return ms, nil
}

func (d *DAO) FindApiSubscriptionKeyJoin(tx *gorm.DB, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.ApiSubscriptionKey, error) {
	var (
		offset = (page - 1) * limit
	)
	var ms []*models.ApiSubscriptionKey
	err := d.findJoin(tx, &ms, joins, filters, preloads, orders, offset, limit, false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return ms, nil
}
