package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMySQl(dsn string, debug bool) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}

	if debug {
		db = db.Debug()
	}
	return db, nil
}
