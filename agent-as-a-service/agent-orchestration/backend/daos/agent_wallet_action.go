package daos

import (
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/jinzhu/gorm"
)

func (d *DAO) FirstAgentWalletActionByID(tx *gorm.DB, id uint, preloads map[string][]interface{}, forUpdate bool) (*models.AgentWalletAction, error) {
	var m models.AgentWalletAction
	if err := d.first(tx, &m, map[string][]interface{}{"id = ?": []interface{}{id}}, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FirstAgentWalletAction(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string) (*models.AgentWalletAction, error) {
	var m models.AgentWalletAction
	if err := d.first(tx, &m, filters, preloads, orders, false); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FindAgentWalletAction(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, offset int, limit int) ([]*models.AgentWalletAction, error) {
	var ms []*models.AgentWalletAction
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, err
	}
	return ms, nil
}

func (d *DAO) FindAgentWalletAction4Page(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.AgentWalletAction, uint, error) {
	var (
		offset = (page - 1) * limit
	)
	var ms []*models.AgentWalletAction
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, 0, errs.NewError(err)
	}
	c, err := d.count(tx, &models.AgentWalletAction{}, filters)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	return ms, c, nil
}

func (d *DAO) FindAgentWalletActionJoinSelect(tx *gorm.DB, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.AgentWalletAction, error) {
	var ms []*models.AgentWalletAction
	err := d.findJoinSelect(tx, &models.AgentWalletAction{}, &ms, selected, joins, filters, preloads, orders, uint(page), uint(limit), false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return ms, nil
}

func (d *DAO) GetWalletActionTradeSum(tx *gorm.DB, agentInfoID uint) ([]*models.AgentWalletTradeSum, error) {
	var rs []*models.AgentWalletTradeSum
	query := tx.Raw(`
		select mint,
			sum(
               if(side = 'buy', amount_in, 0)
				) buy_amount,
			sum(
					if(side = 'sell', amount_out, 0)
				) sell_amount
		from agent_wallet_actions
		where action_type = 'trade_raydium'
			and status = 'done'
			and agent_info_id = ?
		group by mint;
	`, agentInfoID)

	if err := query.Scan(&rs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errs.NewError(err)
	}
	return rs, nil
}
