package daos

import (
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

func (d *DAO) FirstBlockScanInfo(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string) (*models.BlockScanInfo, error) {
	var m models.BlockScanInfo
	if err := d.first(tx, &m, filters, preloads, orders, false); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) FindBlockScanInfo(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, offset int, limit int) ([]*models.BlockScanInfo, error) {
	var ms []*models.BlockScanInfo
	if err := d.find(tx, &ms, filters, preloads, orders, offset, limit, false); err != nil {
		return nil, err
	}
	return ms, nil
}

func (d *DAO) FindJoinSelectBlockScanInfo(tx *gorm.DB, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page int, limit int) ([]*models.BlockScanInfo, error) {
	var ms []*models.BlockScanInfo

	err := d.findJoinSelect(tx, &models.BlockScanInfo{}, &ms, selected, joins, filters, preloads, orders, uint(page), uint(limit), false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return ms, nil
}

func (d *DAO) FirstBlockScanInfoByID(tx *gorm.DB, id uint, preloads map[string][]interface{}, forUpdate bool) (*models.BlockScanInfo, error) {
	var m models.BlockScanInfo
	if err := d.first(tx, &m, map[string][]interface{}{"id = ?": []interface{}{id}}, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DAO) GetTokenMarketPrice(tx *gorm.DB, symbol string) (*numeric.BigFloat, *numeric.BigFloat, error) {
	var rs struct {
		Price        numeric.BigFloat
		Last24hPrice numeric.BigFloat
	}
	query := tx.Raw(`
		select ttp.price, ttp.last24h_price from token_prices ttp where symbol = ?
	`, symbol)

	if err := query.Scan(&rs).Error; err != nil {
		zeroF := numeric.NewBigFloatFromString("0")
		if err == gorm.ErrRecordNotFound {
			return &zeroF, &zeroF, nil
		}
		return &zeroF, &zeroF, errs.NewError(err)
	}
	return &rs.Price, &rs.Last24hPrice, nil
}

func (d *DAO) FirstTokenPrice(tx *gorm.DB, filters map[string][]interface{}, preloads map[string][]interface{}, forUpdate bool) (*models.TokenPrice, error) {
	var m models.TokenPrice
	if err := d.first(tx, &m, filters, preloads, nil, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}
