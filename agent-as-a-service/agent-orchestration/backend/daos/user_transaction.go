package daos

import (
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/jinzhu/gorm"
)

func (d *DAO) FirstUserTransactionByID(tx *gorm.DB, id uint, preloads map[string][]interface{}, forUpdate bool) (*models.UserTransaction, error) {
	var m models.UserTransaction
	if err := d.first(tx, &m, map[string][]interface{}{"id = ?": []interface{}{id}}, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FirstUserTransaction(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string) (*models.UserTransaction, error) {
	var m models.UserTransaction
	if err := d.first(tx, &m, filters, preloads, orders, false); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FindUserTransaction(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, offset int, limit int) ([]*models.UserTransaction, error) {
	var ms []*models.UserTransaction
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, err
	}
	return ms, nil
}

func (d *DAO) FindUserTransaction4Page(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.UserTransaction, uint, error) {
	var (
		offset = (page - 1) * limit
	)
	var ms []*models.UserTransaction
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, 0, errs.NewError(err)
	}
	c, err := d.count(tx, &models.UserTransaction{}, filters)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	return ms, c, nil
}

func (d *DAO) FindUserTransactionJoinSelect(tx *gorm.DB, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.UserTransaction, error) {
	var ms []*models.UserTransaction
	err := d.findJoinSelect(tx, &models.UserTransaction{}, &ms, selected, joins, filters, preloads, orders, uint(page), uint(limit), false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return ms, nil
}
