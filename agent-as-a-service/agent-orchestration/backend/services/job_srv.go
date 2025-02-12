package services

import (
	"context"
	"fmt"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jasonlvhit/gocron"
)

func (s *Service) JobUpdateMarketPrice(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx,
		"JobUpdateMarketPrice",
		func() error {
			eaiPrice, err := s.cmc.GetQuotesLatest([]string{"31401"})
			if err != nil {
				return err
			}
			_ = s.CreateOrUpdateTokenPrice(ctx, "EAI", numeric.NewBigFloatFromFloat(&eaiPrice["31401"].Quote.USD.Price.Float))

			// BTC - ETH price
			mapSymbol := map[string]string{
				"BTCUSDT": "BTC",
				"ETHUSDT": "ETH",
				"SOLUSDT": "SOL",
				"BNBUSDT": "BNB",
			}

			symbol := `["BTCUSDT","ETHUSDT","SOLUSDT","BNBUSDT"]`
			prices, err := helpers.GetListExternalPrice(symbol)
			if err != nil {
				return err
			}

			for _, price := range prices {
				err = s.CreateOrUpdateTokenPrice(ctx, mapSymbol[price.Symbol], numeric.NewBigFloatFromString(price.Price))
				if err != nil {
					fmt.Println(err)
				}

			}

			return nil
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) CreateOrUpdateTokenPrice(ctx context.Context, symbol string, price numeric.BigFloat) error {
	filters := map[string][]interface{}{}
	filters["symbol = ?"] = []interface{}{symbol}
	tokenPrice, err := s.dao.FirstTokenPrice(daos.GetDBMainCtx(ctx), filters, map[string][]interface{}{}, true)
	if err != nil {
		return errs.NewError(err)
	}

	isCreated := false
	if tokenPrice == nil {
		tokenPrice = &models.TokenPrice{}
		tokenPrice.Symbol = symbol
		isCreated = true
	}
	tokenPrice.NetworkID = 0
	tokenPrice.Price = price
	tokenPrice.UpdatedAt = time.Now()

	if isCreated {
		err = s.dao.Create(daos.GetDBMainCtx(ctx), tokenPrice)
		if err != nil {
			return errs.NewError(err)
		}
	} else {
		err = s.dao.Save(daos.GetDBMainCtx(ctx), tokenPrice)
		if err != nil {
			return errs.NewError(err)
		}
	}

	return nil
}

func (s *Service) BlockchainDurations(ctx context.Context) ([]uint, error) {
	chains, err := s.dao.FindBlockScanInfo(
		daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"duration > ?": {0},
		},
		map[string][]interface{}{},
		[]string{},
		0,
		100,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	durationMap := map[uint]bool{}
	for _, v := range chains {
		durationMap[v.Duration] = true
	}
	durations := []uint{}
	for duration := range durationMap {
		durations = append(durations, duration)
	}
	return durations, nil
}

func (s *Service) JobEnabledDB(ctx context.Context) error {
	err := daos.GetDBMainCtx(ctx).
		Model(&models.AppConfig{}).
		Where("network_id = ?", models.GENERTAL_NETWORK_ID).
		Where("name = ?", "job_enabled").
		UpdateColumn("value", "1").Error
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) DisableJobs() {
	s.jobMutex.Lock()
	s.jobDisabled = true
	s.jobMutex.Unlock()
}

func (s *Service) RunJobs(ctx context.Context) error {
	gocron.Every(180).Second().Do(func() {
		s.KnowledgeUsecase.WatchWalletChange(context.Background())
	})
	gocron.Every(30).Second().Do(func() {
		s.KnowledgeUsecase.ScanKnowledgeBaseStatusPaymentReceipt(context.Background())
	})

	gocron.Every(30).Second().Do(func() {
		s.JobScanEventsByChain(context.Background())
	})
	//

	gocron.Every(1).Minute().Do(
		func() error {
			s.JobAgentMintNft(context.Background())
			s.JobRetryAgentMintNft(context.Background())
			s.JobRetryAgentMintNftError(context.Background())
			s.JobAgentStart(context.Background())
			return nil
		},
	)
	gocron.Every(1).Minute().Do(
		func() error {
			s.JobAgentTwinTrain(context.Background())
			return nil
		},
	)

	gocron.Every(1).Minute().Do(s.JobUpdateTwitterAccessToken, context.Background())
	gocron.Every(1).Minute().Do(s.JobCreateTokenInfo, context.Background())
	gocron.Every(7).Minute().Do(s.JobUpdateTokenPriceInfo, context.Background())

	// create agent
	// gocron.Every(5).Minute().Do(s.JobScanAgentTwitterPostForCreateAgent, context.Background())
	// gocron.Every(1).Minute().Do(s.JobAgentTwitterPostCreateAgent, context.Background())

	// trading analyze
	gocron.Every(5).Minute().Do(s.JobScanAgentTwitterPostForTA, context.Background())
	gocron.Every(1).Minute().Do(s.JobAgentTwitterPostTA, context.Background())

	// lucky moneys
	gocron.Every(5).Minute().Do(s.JobLuckyMoneyActionExecuted, context.Background())
	gocron.Every(5).Minute().Do(s.JobLuckyMoneyCollectPost, context.Background())
	gocron.Every(5).Minute().Do(s.JobLuckyMoneyProcessUserReward, context.Background())

	gocron.Every(5).Minute().Do(s.JobUpdateOffchainAutoOutput, context.Background())
	gocron.Every(30).Minute().Do(s.JobUpdateOffchainAutoOutput3Hour, context.Background())
	gocron.Every(5).Minute().Do(s.JobAgentSnapshotPostStatusInferRefund, context.Background())

	gocron.Every(1).Minute().Do(
		func() {
			s.JobAgentSnapshotPostCreate(context.Background())
		},
	)
	gocron.Every(10).Second().Do(
		func() {
			s.JobAgentSnapshotPostActionExecuted(context.Background())
		},
	)
	gocron.Every(5).Minute().Do(
		func() {
			s.JobAgentSnapshotPostActionDupplicated(context.Background())
			s.JobAgentSnapshotPostActionCancelled(context.Background())
		},
	)
	//
	gocron.Every(5).Minute().Do(
		func() {
			s.JobUpdateMarketPrice(context.Background())
			s.JobUpdateMemeUsdPrice(context.Background())
		},
	)
	gocron.Every(1).Day().
		From(helpers.GetNextScheduleTime(24*time.Hour, 0*time.Hour)).
		Do(s.AgentDailyReport, context.Background())

	gocron.Every(5).Minute().
		From(helpers.GetNextScheduleTime(5*time.Minute, 0*time.Hour)).
		Do(s.JobAgentTeleAlertTopupNotAction, context.Background())

	gocron.Every(15).Minute().Do(
		func() {
			s.JobScanTwitterLiked(context.Background())
		},
	)

	gocron.Every(30).Second().Do(
		func() {
			s.JobRunCheck(
				context.Background(),
				"JobAgentLiquidity",
				func() error {
					s.JobAgentDeployToken(context.Background())
					s.JobMemeAddPositionInternal(context.Background())
					s.JobMemeRemovePositionInternal(context.Background())
					s.JobCheckMemeReachMarketCap(context.Background())
					s.JobMemeAddPositionUniswap(context.Background())
					s.JobRetryAddPool1(context.Background())
					s.JobRetryAddPool2(context.Background())
					s.JobRetryAgentDeployToken(context.Background())
					s.JobMemeBurnPositionUniswap(context.Background())
					return nil
				},
			)
		},
	)

	gocron.Every(1).Minutes().Do(
		func() {
			s.JobCreateAgentKnowledgeBase(context.Background())
		},
	)

	// create launchpad
	// gocron.Every(5).Minute().Do(s.JobScanAgentTwitterPostForCreateLaunchpad, context.Background())
	// gocron.Every(30).Second().Do(
	// 	func() {
	// 		s.JobRunCheck(
	// 			context.Background(),
	// 			"JobAgentLaunchpad",
	// 			func() error {
	// 				s.JobAgentTwitterPostCreateLaunchpad(context.Background())
	// 				s.JobAgentLaunchpadEnd(context.Background())
	// 				s.JobAgentDeployDAOToken(context.Background())
	// 				s.JobAgentSettleDAOToken(context.Background())
	// 				s.JobAgentTgeTransferDAOToken(context.Background())
	// 				s.JobAgentAddLiquidityDAOToken(context.Background())
	// 				s.JobAgentTgeRefundBaseToken(context.Background())
	// 				return nil
	// 			},
	// 		)
	// 	},
	// )
	// gocron.Every(1).Minutes().Do(
	// 	func() {
	// 		s.JobScanRepliesByLaunchpadTweetID(context.Background())
	// 	},
	// )

	<-gocron.Start()
	return nil
}
