package daos

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var dbMain *gorm.DB

func InitDBConn(dbMainConn *gorm.DB) {
	dbMain = dbMainConn
}

func GetDBMainCtx(ctx context.Context) *gorm.DB {
	return dbMain.New()
}

func WithTransaction(dbConn *gorm.DB, callback func(*gorm.DB) error) (err error) {
	tx := dbConn.Begin()
	defer func() {
		if rval := recover(); rval != nil {
			tx.Rollback()
			err = errs.NewError(errors.New(fmt.Sprint(rval)))
		}
	}()
	if err = callback(tx); err != nil {
		tx.Rollback()
		return err
	}
	if err = tx.Commit().Error; err != nil {
		return errs.NewError(err)
	}
	return nil
}

type DAO struct {
	mtx       sync.Mutex
	configMap map[string]string
}

func (d *DAO) Create(tx *gorm.DB, m interface{}) error {
	if err := tx.Create(m).Error; err != nil {
		return err
	}
	return nil
}

func (d *DAO) Save(tx *gorm.DB, m interface{}) error {
	if err := tx.Save(m).Error; err != nil {
		return err
	}
	return nil
}

func (d *DAO) Delete(tx *gorm.DB, m interface{}) error {
	if err := tx.Delete(m).Error; err != nil {
		return err
	}
	return nil
}

func (d *DAO) DeleteUnscoped(tx *gorm.DB, m interface{}) error {
	if err := tx.Unscoped().Delete(m).Error; err != nil {
		return err
	}
	return nil
}

func (d *DAO) first(tx *gorm.DB, m interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, forUpdate bool) error {
	query := tx
	for k, v := range filters {
		if v != nil {
			query = query.Where(k, v...)
		} else {
			query = query.Where(k)
		}
	}
	for k, v := range preloads {
		if v != nil {
			query = query.Preload(k, v...)
		} else {
			query = query.Preload(k)
		}
	}
	if orders != nil && len(orders) > 0 {
		for _, v := range orders {
			query = query.Order(v)
		}
	}
	if forUpdate {
		query = query.Set("gorm:query_option", "FOR UPDATE")
	}
	if err := query.First(m).Error; err != nil {
		return err
	}
	return nil
}

func (d *DAO) find(tx *gorm.DB, ms interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, offset int, limit int, forUpdate bool) error {
	query := tx
	for k, v := range filters {
		if v != nil {
			query = query.Where(k, v...)
		} else {
			query = query.Where(k)
		}
	}
	for k, v := range preloads {
		if v != nil {
			query = query.Preload(k, v...)
		} else {
			query = query.Preload(k)
		}
	}
	if orders != nil && len(orders) > 0 {
		for _, v := range orders {
			query = query.Order(v)
		}
	}
	if offset >= 0 {
		query = query.Offset(offset)
	}
	if limit >= 0 {
		query = query.Limit(limit)
	}
	if forUpdate {
		query = query.Set("gorm:query_option", "FOR UPDATE")
	}
	if err := query.Find(ms).Error; err != nil {
		return err
	}
	return nil
}

func (d *DAO) count(tx *gorm.DB, m interface{}, filters map[string][]interface{}) (uint, error) {
	query := tx
	for k, v := range filters {
		if v != nil {
			query = query.Where(k, v...)
		} else {
			query = query.Where(k)
		}
	}
	var count uint
	if err := query.Model(m).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (d *DAO) findJoin(tx *gorm.DB, ms interface{}, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, offset int, limit int, forUpdate bool) error {
	query := tx
	for k, v := range joins {
		if v != nil {
			query = query.Joins(k, v...)
		} else {
			query = query.Joins(k)
		}
	}
	for k, v := range filters {
		if v != nil {
			query = query.Where(k, v...)
		} else {
			query = query.Where(k)
		}
	}
	for k, v := range preloads {
		if v != nil {
			query = query.Preload(k, v...)
		} else {
			query = query.Preload(k)
		}
	}
	for _, v := range orders {
		query = query.Order(v)
	}
	if offset >= 0 {
		query = query.Offset(offset)
	}
	if limit >= 0 {
		query = query.Limit(limit)
	}
	if forUpdate {
		query = query.Set("gorm:query_option", "FOR UPDATE")
	}
	if err := query.Find(ms).Error; err != nil {
		return err
	}
	return nil
}

func (d *DAO) countJoin(tx *gorm.DB, m interface{}, joins map[string][]interface{}, filters map[string][]interface{}) (uint, error) {
	query := tx
	for k, v := range joins {
		if v != nil {
			query = query.Joins(k, v...)
		} else {
			query = query.Joins(k)
		}
	}
	for k, v := range filters {
		if v != nil {
			query = query.Where(k, v...)
		} else {
			query = query.Where(k)
		}
	}
	var count uint
	if err := query.Model(m).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (d *DAO) findJoin4Page(tx *gorm.DB, m interface{}, ms interface{}, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page uint, limit uint, forUpdate bool) (uint, error) {
	var count uint
	offset := page*limit - limit
	query := tx
	for k, v := range joins {
		if v != nil {
			query = query.Joins(k, v...)
		} else {
			query = query.Joins(k)
		}
	}
	for k, v := range filters {
		if v != nil {
			query = query.Where(k, v...)
		} else {
			query = query.Where(k)
		}
	}
	for k, v := range preloads {
		if v != nil {
			query = query.Preload(k, v...)
		} else {
			query = query.Preload(k)
		}
	}
	for _, v := range orders {
		query = query.Order(v)
	}
	if err := query.Model(m).Count(&count).Error; err != nil {
		return 0, err
	}
	if forUpdate {
		query = query.Set("gorm:query_option", "FOR UPDATE")
	}
	query = query.Limit(limit).Offset(offset)
	if err := query.Find(ms).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (d *DAO) findAll(tx *gorm.DB, ms interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, forUpdate bool) error {
	query := tx
	for k, v := range filters {
		if v != nil {
			query = query.Where(k, v...)
		} else {
			query = query.Where(k)
		}
	}
	for k, v := range preloads {
		if v != nil {
			query = query.Preload(k, v...)
		} else {
			query = query.Preload(k)
		}
	}
	if len(orders) > 0 {
		for _, v := range orders {
			query = query.Order(v)
		}
	}
	if forUpdate {
		query = query.Set("gorm:query_option", "FOR UPDATE")
	}
	if err := query.Find(ms).Error; err != nil {
		return err
	}
	return nil
}

func (d *DAO) findJoinSelect(tx *gorm.DB, m interface{}, ms interface{}, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page uint, limit uint, forUpdate bool) error {
	offset := page*limit - limit
	query := tx
	if len(selected) > 0 {
		query = query.Select(strings.Join(selected, ", "))
	}
	for k, v := range joins {
		if v != nil {
			query = query.Joins(k, v...)
		} else {
			query = query.Joins(k)
		}
	}
	for k, v := range filters {
		if v != nil {
			query = query.Where(k, v...)
		} else {
			query = query.Where(k)
		}
	}
	for k, v := range preloads {
		if v != nil {
			query = query.Preload(k, v...)
		} else {
			query = query.Preload(k)
		}
	}
	if len(orders) > 0 {
		for _, v := range orders {
			query = query.Order(v)
		}
	}
	if forUpdate {
		query = query.Set("gorm:query_option", "FOR UPDATE")
	}
	query = query.Limit(limit).Offset(offset)
	if err := query.Find(ms).Error; err != nil {
		return err
	}
	return nil
}

func (d *DAO) firstJoinSelect(tx *gorm.DB, m interface{}, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, forUpdate bool) error {
	query := tx
	if len(selected) > 0 {
		query = query.Select(strings.Join(selected, ", "))
	}

	for k, v := range joins {
		if v != nil {
			query = query.Joins(k, v...)
		} else {
			query = query.Joins(k)
		}
	}

	for k, v := range filters {
		if v != nil {
			query = query.Where(k, v...)
		} else {
			query = query.Where(k)
		}
	}
	for k, v := range preloads {
		if v != nil {
			query = query.Preload(k, v...)
		} else {
			query = query.Preload(k)
		}
	}
	if orders != nil && len(orders) > 0 {
		for _, v := range orders {
			query = query.Order(v)
		}
	}
	if forUpdate {
		query = query.Set("gorm:query_option", "FOR UPDATE")
	}
	if err := query.First(m).Error; err != nil {
		return err
	}
	return nil
}

func (d *DAO) findJoinSelect4Page(tx *gorm.DB, m interface{}, ms interface{}, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page uint, limit uint, forUpdate bool) (uint, error) {
	var count uint
	offset := page*limit - limit
	query := tx
	if len(selected) > 0 {
		query = query.Select(strings.Join(selected, ", "))
	}
	for k, v := range joins {
		if v != nil {
			query = query.Joins(k, v...)
		} else {
			query = query.Joins(k)
		}
	}
	for k, v := range filters {
		if v != nil {
			query = query.Where(k, v...)
		} else {
			query = query.Where(k)
		}
	}
	for k, v := range preloads {
		if v != nil {
			query = query.Preload(k, v...)
		} else {
			query = query.Preload(k)
		}
	}
	if len(orders) > 0 {
		for _, v := range orders {
			query = query.Order(v)
		}
	}
	if forUpdate {
		query = query.Set("gorm:query_option", "FOR UPDATE")
	}

	if err := query.Model(m).Count(&count).Error; err != nil {
		return 0, err
	}

	query = query.Limit(limit).Offset(offset)
	if err := query.Find(ms).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (d *DAO) findJoinSelect4PageNoCount(tx *gorm.DB, m interface{}, ms interface{}, selected []string, joins map[string][]interface{}, filters map[string][]interface{}, preloads map[string][]interface{}, orders []string, page uint, limit uint, forUpdate bool) (uint, error) {
	offset := page*limit - limit
	query := tx
	if len(selected) > 0 {
		query = query.Select(strings.Join(selected, ", "))
	}
	for k, v := range joins {
		if v != nil {
			query = query.Joins(k, v...)
		} else {
			query = query.Joins(k)
		}
	}
	for k, v := range filters {
		if v != nil {
			query = query.Where(k, v...)
		} else {
			query = query.Where(k)
		}
	}
	for k, v := range preloads {
		if v != nil {
			query = query.Preload(k, v...)
		} else {
			query = query.Preload(k)
		}
	}
	if len(orders) > 0 {
		for _, v := range orders {
			query = query.Order(v)
		}
	}
	if forUpdate {
		query = query.Set("gorm:query_option", "FOR UPDATE")
	}

	query = query.Limit(limit).Offset(offset)
	if err := query.Find(ms).Error; err != nil {
		return 0, err
	}
	return 0, nil
}
