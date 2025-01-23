package mysql

import (
	"time"

	mysqlDriver "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/logger"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
	gormtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorm.io/gorm.v1"
	mysqlgorm "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

// NewDefaultMysqlGormConn --
func NewDefaultMysqlGormConn(conn *Connection, dbUrl string, debug bool) *gorm.DB {
	if conn == nil {
		conn = DefaultMysqlConnectionFromConfig()
	}
	conn.EnableLog = debug
	settings := dbUrl
	// settings := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
	// 	conn.User, conn.Password,
	// 	conn.Protocol,
	// 	conn.Host, conn.Port,
	// 	conn.DatabaseName,
	// 	conn.Charset,
	// 	conn.ParseTime,
	// 	conn.Location,
	// )
	// logSettings := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
	// 	conn.User, utils.CensorString(conn.Password),
	// 	conn.Protocol,
	// 	conn.Host, conn.Port,
	// 	conn.DatabaseName,
	// 	conn.Charset,
	// 	conn.ParseTime,
	// 	conn.Location,
	// )
	logSettings := dbUrl
	logger.Info("[mysql] Gorm connecting with this configuration: ", logSettings)
	zapLogger := zapgorm2.New(logger.Logger())
	zapLogger.SlowThreshold = 10 * time.Second
	zapLogger.LogLevel = gormlogger.Silent
	if conn.EnableLog {
		zapLogger.LogLevel = gormlogger.Info
	}
	sqltrace.Register("mysql", &mysqlDriver.MySQLDriver{}, sqltrace.WithServiceName(conn.ServiceName))
	sqlDB, err := sqltrace.Open("mysql", settings)
	if err != nil {
		logger.Fatal("NewDefaultMysqlGormConn", zap.Error(err))
	}
	if conn.MaxOpenConn > 0 {
		sqlDB.SetMaxOpenConns(conn.MaxOpenConn)
	}
	if conn.MaxIdleConn > 0 && conn.MaxIdleConn < conn.MaxOpenConn {
		sqlDB.SetMaxIdleConns(conn.MaxIdleConn)
	}
	gormConfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   zapLogger,
	}
	db, err := gormtrace.Open(mysqlgorm.New(mysqlgorm.Config{Conn: sqlDB}),
		gormConfig,
		gormtrace.WithServiceName(conn.ServiceName),
		gormtrace.WithErrorCheck(gormTraceErrCheck),
	)
	if err != nil {
		logger.Fatal("NewDefaultMysqlGormConn", zap.Error(err))
	}

	logger.Info("[mysql] Gorm connected, config:", logSettings)

	return db
}

// gormTraceErrCheck check error of gorm
func gormTraceErrCheck(err error) bool {
	if err == nil {
		return false
	}

	return err != gorm.ErrRecordNotFound
}
