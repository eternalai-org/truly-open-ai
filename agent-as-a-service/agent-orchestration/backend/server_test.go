package main_test

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/configs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/databases"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/logger"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services"
)

var ts *services.Service

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	conf := configs.GetConfig()
	logger.NewLogger("agents-ai-api", conf.Env, "", true)
	defer logger.Sync()
	dbMain, err := databases.Init(
		conf.DbURL,
		nil,
		1,
		20,
		conf.Debug,
	)
	if err != nil {
		panic(err)
	}
	daos.InitDBConn(
		dbMain,
	)
	var (
		s = services.NewService(
			conf,
		)
	)
	ts = s
}

func Test_JOB(t *testing.T) {
	// ts.JobScanAgentTwitterPostForTA(context.Background())
	// ts.RetryAgentDeployToken(context.Background(), 51265)
	// ts.JobUpdateOffchainAutoOutputForMission(context.Background())
	// ts.JobAgentTwitterPostTA(context.Background())
	// ts.JobLuckyMoneyActionExecuted(context.Background())
	fmt.Println(
	// ts.DeployDAOTreasuryLogic(context.Background(), models.BASE_CHAIN_ID),
	// ts.DeployDAOTreasuryAddress(context.Background(), models.BASE_CHAIN_ID),
	// ts.AgentAddLiquidityDAOToken(context.Background(), 1),
	// ts.CreateSOLAddress(context.Background()),
	// ts.CreateETHAddress(context.Background()),
	)
	select {}
}

func Test_UTIL(t *testing.T) {
	ts.TestUtil()
}

func Test_SRV(t *testing.T) {
	// ts.ScanAgentTwitterPostFroCreateLaunchpad(context.Background())
	// ts.AgentTwitterPostCreateLaunchpad(context.Background(), 34285)
}
