package mysql

import (
	"context"
	"math/rand"

	"gorm.io/gorm"
)

type IBaseRepository interface {
	buildCondition(action Action, args ...interface{}) *gorm.DB
	setMasterDB(db *gorm.DB) *baseRepository
	setReplicaDb(db ...*gorm.DB) *baseRepository
	getReplicaDB() *gorm.DB

	GetDB(actions ...Action) *gorm.DB
	FindOneHandler(ctx context.Context, result interface{}, args ...interface{}) error
	FindByIDHandler(ctx context.Context, ID int, result interface{}, args ...interface{}) error
	FindsHandler(ctx context.Context, result interface{}, args ...interface{}) error
}

func NewBaseRepository(masterDB *gorm.DB, replicaDB ...*gorm.DB) IBaseRepository {
	result := &baseRepository{}
	result = result.setMasterDB(masterDB).setReplicaDb(replicaDB...)
	return result
}

type baseRepository struct {
	masterDB  *gorm.DB
	replicaDB []*gorm.DB
}

func (m *baseRepository) buildCondition(action Action, args ...interface{}) *gorm.DB {
	db := m.GetDB(action)

	limit := ZeroLimit
	orderBy := ZeroOrderBy

	for _, arg := range args {
		switch param := arg.(type) {
		case Filters:
			db = param.setCondition(db)
		case *FilterItem:
			db = param.setCondition(db)
		case FilterItem:
			db = param.setCondition(db)
		case *PreloadItem:
			db = param.setCondition(db)
		case PreloadItem:
			db = param.setCondition(db)
		case Preloads:
			db = param.setCondition(db)
		case Limit:
			limit = param
		case OrderBy:
			orderBy = param
		}
	}

	if !limit.IsZero() {
		db = db.Limit(limit.toInt())
	}

	if !orderBy.IsZero() {
		db = orderBy.SetCondition(db)
	}

	return db
}

func (m *baseRepository) setMasterDB(db *gorm.DB) *baseRepository {
	m.masterDB = db
	return m
}

func (m *baseRepository) setReplicaDb(db ...*gorm.DB) *baseRepository {
	m.replicaDB = db
	return m
}

func (m *baseRepository) GetDB(actions ...Action) *gorm.DB {
	databaseAction := WriteOrRead
	if len(actions) > 0 {
		databaseAction = actions[0]
	}

	switch databaseAction {
	case ReadOnly:
		return m.getReplicaDB()
	case WriteOrRead:
		return m.masterDB
	}

	return m.masterDB
}

func (m *baseRepository) getReplicaDB() *gorm.DB {
	// if has no replicaDB then get masterDB
	lenReplicaDB := len(m.replicaDB)
	if lenReplicaDB == 0 {
		return m.masterDB
	}

	return m.replicaDB[rand.Intn(lenReplicaDB)]
}

func (m *baseRepository) FindOneHandler(ctx context.Context, result interface{}, args ...interface{}) error {
	db := m.buildCondition(ReadOnly, args...)
	err := db.First(result).Error
	return err
}

func (m *baseRepository) FindByIDHandler(ctx context.Context, ID int, result interface{}, args ...interface{}) error {
	db := m.buildCondition(ReadOnly, args...)
	err := db.First(result, ID).Error
	return err
}

func (m *baseRepository) FindsHandler(ctx context.Context, result interface{}, args ...interface{}) error {
	db := m.buildCondition(ReadOnly, args...)
	err := db.Find(result).Error
	return err
}
