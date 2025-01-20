package daos

import (
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/jinzhu/gorm"
)

func (d *DAO) FirstAgentTradeTokenByID(tx *gorm.DB, id uint, preloads map[string][]interface{}, forUpdate bool) (*models.AgentTradeToken, error) {
	var m models.AgentTradeToken
	if err := d.first(tx, &m, map[string][]interface{}{"id = ?": []interface{}{id}}, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FirstAgentTradeToken(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string) (*models.AgentTradeToken, error) {
	var m models.AgentTradeToken
	if err := d.first(tx, &m, filters, preloads, orders, false); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FindAgentTradeToken(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, offset int, limit int) ([]*models.AgentTradeToken, error) {
	var ms []*models.AgentTradeToken
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, err
	}
	return ms, nil
}

func (d *DAO) FindAgentTradeToken4Page(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.AgentTradeToken, uint, error) {
	var (
		offset = (page - 1) * limit
	)
	var ms []*models.AgentTradeToken
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, 0, errs.NewError(err)
	}
	c, err := d.count(tx, &models.AgentTradeToken{}, filters)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	return ms, c, nil
}

func (d *DAO) FindAgentTradeTokenJoinSelect(tx *gorm.DB, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.AgentTradeToken, error) {
	var ms []*models.AgentTradeToken
	err := d.findJoinSelect(tx, &models.AgentTradeToken{}, &ms, selected, joins, filters, preloads, orders, uint(page), uint(limit), false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return ms, nil
}

func (d *DAO) FindAgentTradeTokenJoin(tx *gorm.DB, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.AgentTradeToken, error) {
	var (
		offset = (page - 1) * limit
	)
	var ms []*models.AgentTradeToken
	err := d.findJoin(tx, &ms, joins, filters, preloads, orders, offset, limit, false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return ms, nil
}
