package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/eternalai-org/eternal-ai/agent-orchestration/core/logger"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/spf13/viper"

	"github.com/eternalai-org/eternal-ai/agent-orchestration/core/common/database"
	"github.com/eternalai-org/eternal-ai/agent-orchestration/core/server"

	_walletHttpDeliver "github.com/eternalai-org/eternal-ai/agent-orchestration/core/modules/wallet/delivery/http"
	_walletRepo "github.com/eternalai-org/eternal-ai/agent-orchestration/core/modules/wallet/repository/mysql"
	_walletUcase "github.com/eternalai-org/eternal-ai/agent-orchestration/core/modules/wallet/usecase"
)

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	viper.SetConfigFile(`config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode...")
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			errCaptured := fmt.Errorf("%s\n%s", err, string(debug.Stack()))
			log.Println(errCaptured)
		}
	}()
	// new db connection
	dbURL := viper.GetString(`db_url`)
	if dbURL == "" {
		dbURL = os.Getenv("BD_URL")
	}
	// new db connection
	var migrateCoreFunc func(db *gorm.DB) error
	if os.Getenv("DEV") != "true" {
		migrateCoreFunc = database.MigrateCore
	}
	dbConn, err := database.Init(dbURL, migrateCoreFunc, 1, 5, viper.GetBool(`debug`))
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %s", err.Error()))
	}
	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	defer func() {
		if err := recover(); err != nil {
			errCaptured := fmt.Errorf("%s\n%s", err, string(debug.Stack()))
			log.Println(errCaptured)
			return
		}
	}()
	// repositories
	// -- Wallet
	wr := _walletRepo.NewMysqlWalletRepository(dbConn)
	// usecases
	// -- Wallet
	wu := _walletUcase.NewWalletUsecase(wr)

	s := server.NewHTTPServer(logger.Logger())
	_walletHttpDeliver.NewWalletHandler(s, wu)

	serverAddress := viper.GetString(`port`)
	if err := s.Engine().Run(serverAddress); err != nil {
		log.Fatal(err)
	}
}
