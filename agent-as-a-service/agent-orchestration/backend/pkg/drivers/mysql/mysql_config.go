package mysql

import "github.com/spf13/viper"

// Connection -- mysql connection
type Connection struct {
	Host         string
	Port         int
	Protocol     string
	User         string
	Password     string
	DatabaseName string
	Charset      string
	ParseTime    bool
	Location     string
	Others       string
	EnableLog    bool
	MaxOpenConn  int
	MaxIdleConn  int
	ServiceName  string
}

const (
	// MaxOpenConn --
	MaxOpenConn = 100
	// MaxIdleConn --
	MaxIdleConn = 10
)

// DefaultMysqlConnectionFromConfig -- load connection settings in config with default key
func DefaultMysqlConnectionFromConfig() *Connection {
	maxOpenConn := viper.GetInt("MYSQL_MAX_OPEN_CONN")
	if maxOpenConn <= 0 {
		maxOpenConn = MaxOpenConn
	}
	maxIdleConn := viper.GetInt("MYSQL_MAX_IDLE_CONN")
	if MaxIdleConn <= 0 {
		maxIdleConn = MaxIdleConn
	}
	serviceName := viper.GetString("MYSQL_SERVICE_NAME")
	if serviceName == "" {
		serviceName = "mysql"
	}
	return &Connection{
		Host:         viper.GetString("MYSQL_HOST"),
		Port:         viper.GetInt("MYSQL_PORT"),
		Protocol:     viper.GetString("MYSQL_PROTOCOL"),
		User:         viper.GetString("MYSQL_USER"),
		Password:     viper.GetString("MYSQL_PASSWORD"),
		DatabaseName: viper.GetString("MYSQL_DATABASE_NAME"),
		Charset:      viper.GetString("MYSQL_CHARSET"),
		ParseTime:    viper.GetBool("MYSQL_PARSE_TIME"),
		Location:     viper.GetString("MYSQL_LOCATION"),
		Others:       viper.GetString("MYSQL_OTHERS"),
		EnableLog:    viper.GetBool("LOGGER_ENABLE_DEBUG"),
		ServiceName:  serviceName,
		MaxOpenConn:  maxOpenConn,
		MaxIdleConn:  maxIdleConn,
	}
}
