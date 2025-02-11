package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type (
	Environment string
	Mode        string
)

const (
	Develop    Environment = "develop"
	Staging    Environment = "staging"
	Production Environment = "production"

	Api    Mode = "api"
	Worker Mode = "worker"
)

// HandleSigterm -- Handles Ctrl+C or most other means of "controlled" shutdown gracefully.
// Invokes the supplied func before exiting.
func HandleSigterm(handleExit func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		handleExit()
		os.Exit(1)
	}()
}

func IsEnvProduction() bool {
	return viper.GetString("ENV") == string(Production)
}

func IsWorker() bool {
	return viper.GetString("MODE") == string(Worker)
}

func IsApi() bool {
	return viper.GetString("MODE") == string(Api)
}

func ToPointer[E any](s E) *E {
	return &s
}

func GenUniqueInt64() int64 {
	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
	uniqueID := uuid.New().ID()
	return currentTimestamp + int64(uniqueID)
}

func StructToMap(obj interface{}) (map[string]string, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	err = json.Unmarshal(bytes, &m)
	if err != nil {
		return nil, err
	}

	strMap := make(map[string]string)
	for k, v := range m {
		strMap[k] = fmt.Sprintf("%v", v)
	}

	return strMap, nil
}
