package database

import (
	"github.com/eternalai-org/eternal-ai/agent-orchestration/core/domain"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// Init : config
func Init(dbURL string, migrateFunc func(db *gorm.DB) error, idleNum int, openNum int, debug bool) (*gorm.DB, error) {
	dbConn, err := gorm.Open("mysql", dbURL)
	if debug {
		dbConn.LogMode(true)
	}
	if err != nil {
		return nil, errors.Wrap(err, "gorm.Open")
	}

	dbConn = dbConn.Set("gorm:save_associations", false)
	dbConn = dbConn.Set("gorm:association_save_reference", false)
	dbConn.DB().SetMaxIdleConns(idleNum)
	dbConn.DB().SetMaxOpenConns(openNum)

	if migrateFunc != nil {
		err = migrateFunc(dbConn)
		if err != nil {
			return dbConn, err
		}
	}

	return dbConn, nil
}

func MigrateCore(db *gorm.DB) error {
	allTables := []interface{}{
		(*domain.Wallet)(nil),
	}
	if err := db.AutoMigrate(allTables...).Error; err != nil {
		return errors.Wrap(err, "db.AutoMigrate")
	}
	return nil
}
