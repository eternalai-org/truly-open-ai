package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/logger"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/pkg/utils"
	"go.uber.org/zap"
)

// CreateMysqlConnection --
func CreateMysqlConnection(conn *Connection) *AtMysql {
	if conn == nil {
		conn = DefaultMysqlConnectionFromConfig()
	}
	db, err := NewConnection(conn)
	if err != nil {
		logger.Panic("CreateMysqlConnection", zap.Error(err))
	}
	return db
}

// AtMysql --
type AtMysql struct {
	*sql.DB
}

// NewConnection -- open connection to db
func NewConnection(conn *Connection) (*AtMysql, error) {
	var err error

	settings := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=%s&parseTime=%t",
		conn.User, conn.Password,
		conn.Protocol,
		conn.Host, conn.Port,
		conn.DatabaseName,
		conn.Charset, conn.ParseTime,
	)
	if utils.IsStringNotEmpty(conn.Others) {
		settings = fmt.Sprintf("%s&%s", settings, conn.Others)
	}

	logSettings := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=%s&parseTime=%t",
		conn.User, utils.CensorString(conn.Password),
		conn.Protocol,
		conn.Host, conn.Port,
		conn.DatabaseName,
		conn.Charset, conn.ParseTime,
	)

	logger.Info("[mysql] Connecting with this configuration: ", logSettings)

	db, err := sql.Open("mysql", settings)
	if err != nil {
		logger.Error("mysql", "[mysql] Could not connect database, details: ", zap.Error(err))
		return nil, err
	}
	if conn.MaxOpenConn > 0 {
		db.SetMaxOpenConns(conn.MaxOpenConn)
	}
	if conn.MaxIdleConn > 0 && conn.MaxIdleConn < conn.MaxOpenConn {
		db.SetMaxIdleConns(conn.MaxIdleConn)
	}
	db.SetConnMaxLifetime(3600 * time.Second)

	err = db.Ping()
	if err != nil {
		logger.Error("mysql", "[mysql] Could not ping to database, details: ", zap.Error(err))
		return nil, err
	}
	logger.Info("[mysql] Connected, config:", logSettings)

	return &AtMysql{db}, nil
}

// Close -- close connection
func (c *AtMysql) Close() {
	if c == nil {
		return
	}
	err := c.DB.Close()
	if err != nil {
		logger.Error("mysql", "[mysql] Could not close connect database, details: ", zap.Error(err))
	}
}

// GetDB -- get db
func (c *AtMysql) GetDB() *sql.DB {
	return c.DB
}
