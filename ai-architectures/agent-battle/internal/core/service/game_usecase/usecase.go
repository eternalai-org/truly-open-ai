package game_usecase

import (

"agent-battle/internal/contract/erc20/usecase"
"agent-battle/pkg/cryptoamount"
"context"
"encoding/base64"
"errors"
"math/big"
"strings"
"time"

	"agent-battle/internal/adapters/repository/mongo"
	"agent-battle/internal/core/model"
	"agent-battle/internal/core/port"
	"agent-battle/pkg/drivers/mongodb"
	"agent-battle/pkg/encrypt"
	"agent-battle/pkg/logger"
	"agent-battle/pkg/secret_manager"
	"agent-battle/pkg/utils"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type gameUsecase struct {
	gameRepo     mongo.IGameRepo
	settingRepo  mongo.ISettingRepo
	erc20Usecase port.IContractErc20Usecase
	secretKey    string
	setting      *model.Setting
}

func (uc *gameUsecase) ListGame(ctx context.Context, req *model.ListGameRequest) (*model.ListGameResponse, error) {
	result := []*model.Game{}
	filters := make(bson.M)
	if len(req.TweetIds) != 0 {
		filters["tweet_id"] = bson.M{"$in": req.TweetIds}
	}
	match := bson.M{"$match": filters}
	skip, limit := req.BuildPipeline()
	pipeline := make(bson.A, 0)
	sort := bson.M{
		"$sort": bson.D{
			{Key: "_id", Value: -1},
		},
	}
	pipeline = append(pipeline, match, sort, skip, limit)
	opts := mongodb.OptionsAggregate()
	if err := uc.gameRepo.Aggregations(ctx, &result, pipeline, opts); err != nil {
		return nil, err
	}

	total, err := uc.gameRepo.CountDocuments(ctx, filters, mongodb.OptionsCount())
	if err != nil {
		return nil, err
	}

	return &model.ListGameResponse{Games: result, TotalRecords: total}, nil
}

func (uc *gameUsecase) GameResult(ctx context.Context, req *model.GameResultRequest) (*model.Game, error) {
	game, err := uc.gameByTweetId(ctx, req.TweetId)
	if err != nil {
		return nil, err
	}

	if game.Status == model.GameStatusResultUpdated {
		return nil, errors.New("a game result has been determined")
	}

	valid := false
	for _, u := range game.AgentWallets {
		if strings.EqualFold(u.Username, req.Username) {
			valid = true
			break
		}
	}

	if req.Username != "" && !valid {
		return nil, errors.New("the declared winner does not correspond to any of the registered agents")
	}

	game.Winner = req.Username
	game.Status = model.GameStatusResultUpdated

	if err := uc.updateGame(ctx, game); err != nil {
		return nil, err
	}
	return game, nil
}

func (uc *gameUsecase) EndGame(ctx context.Context, tweetId string) (*model.Game, error) {
	game, err := uc.gameByTweetId(ctx, tweetId)
	if err != nil {
		return nil, err
	}

	game.EndTime = time.Now()
	if err := uc.updateGame(ctx, game); err != nil {
		return nil, err
	}

	if err := uc.markEndGame(ctx, game); err != nil {
		return nil, err
	}

	return game, nil
}

func (uc *gameUsecase) gameByTweetId(ctx context.Context, tweetId string) (*model.Game, error) {
	game := &model.Game{}
	if err := uc.gameRepo.FindOne(ctx, bson.M{"tweet_id": tweetId}, game); err != nil {
		return nil, err
	}

	return game, nil
}

func (uc *gameUsecase) DetailGame(ctx context.Context, tweetId string) (*model.Game, error) {
	return uc.gameByTweetId(ctx, tweetId)
}

func (uc *gameUsecase) WatchGameState(ctx context.Context) error {
	games := []*model.Game{}
	filters := make(bson.M)

	/*Actual Mongo Query:
	{
	 "status": { "$in": [1, 3] },
	 "address": { "$ne": "", "$exists": true },
	 "tweet_id": { "$ne": "", "$exists": true }
	}
	*/
	filters["status"] = bson.M{"$in": []model.GameStatus{
		model.GameStatusRunning, model.GameStatusResultUpdated,
	}}
	filters["address"] = bson.M{"$ne": "", "$exists": true}
	filters["tweet_id"] = bson.M{"$ne": "", "$exists": true}
	err := uc.gameRepo.Find(ctx, filters, &games)
	if err != nil {
		return err
	}

	for _, game := range games {
		// If game result has been determined, then prize to winners
		if game.Status == model.GameStatusResultUpdated {
			if err := uc.prizeToWinners(ctx, game); err != nil {
				logger.GetLoggerInstanceFromContext(ctx).Error("update_game", zap.Error(err))
			}
			continue
		}

		// If game end time has passed, then mark end game
		if game.EndTime.Before(time.Now()) {
			// check game balance and update players the last time,
			// to make sure that the game balance is correct
			if err := uc.checkGameBalance(ctx, game); err != nil {
				logger.GetLoggerInstanceFromContext(ctx).Error("check_game_balance err", zap.Error(err))
			}

			if err := uc.markEndGame(ctx, game); err != nil {
				logger.GetLoggerInstanceFromContext(ctx).Error("make_end_game", zap.Error(err))
			}
			continue
		}

		// If game is running, then check game balance and update players
		if err := uc.checkGameBalance(ctx, game); err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Error("check_game_balance err", zap.Error(err))
		}
	}
	return nil
}

func (uc *gameUsecase) markEndGame(ctx context.Context, game *model.Game) error {
	if uc.setting == nil {
		return errors.New("please configure the application setting first")
	}

	logger.GetLoggerInstanceFromContext(ctx).Info("mark_end_game", zap.String("tweet_id", game.TweetId))
	// transfer token from agents to wallet game
	for _, a := range game.AgentWallets {
		if a.CannotTransferFundsToGameWallet() {
			continue
		}

		if err := uc.handlerTransferTokenFromAgentToGame(ctx, game.Address, a); err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Error(
				"[markEndGame] transfer token from agent to game address failed",
				zap.Error(err),
				zap.String("fromAddress", game.Address),
				zap.String("toAddress", game.Address),
				zap.String("tweetId", game.TweetId),
			)
			return err
		}

		// update current transfer result to game
		if err := uc.updateGame(ctx, game); err != nil {
			return err
		}

		// wait for 1 second for the next transfer
		time.Sleep(1 * time.Second)
	}

	// update game status to ended
	game.Status = model.GameStatusEnded
	return uc.updateGame(ctx, game)
}

// handlerTransferTokenFromAgentToGame transfer token from agent to game address
func (uc *gameUsecase) handlerTransferTokenFromAgentToGame(
	ctx context.Context,
	gameAddress string,
	agent *model.AgentWallet,
) error {
	amount := agent.Amount.ToBigInt()

	agentPrivateKey, err := encrypt.DecryptToString(agent.PrivKey, uc.secretKey)
	if err != nil {
		return err
	}

	operationPrivateKey, err := encrypt.DecryptToString(uc.setting.OperationPrivKey, uc.secretKey)
	if err != nil {
		return err
	}

	// estimate gas fee for transfer token then transfer eth fee from operation address to from address
	gasFee, err := uc.erc20Usecase.EstimateGasFee(ctx, agent.Address, gameAddress, amount)
	if err != nil {
		return err
	}

	// transfer eth fee from operation address to from address
	ethTx, err := uc.erc20Usecase.TransferETH(ctx, agent.Address, gasFee, operationPrivateKey)
	if err != nil {
		return err
	}
	logger.GetLoggerInstanceFromContext(ctx).Info("transfer_eth_fee", zap.String("tx", ethTx))

	// transfer token from agent to game address
	tx, err := uc.erc20Usecase.TransferToken(ctx, gameAddress, amount, agentPrivateKey)
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error(
			"transfer_token",
			zap.Error(err),
			zap.String("gameAddress", gameAddress),
			zap.String("agentAddress", agent.Address),
		)
		return err
	}

	agent.TxHash = tx
	agent.TransferAmount = agent.Amount
	return nil
}

// faucetGasToGameWallet transfer eth fee from operation address to game address
func (uc *gameUsecase) faucetGasToGameWallet(
	ctx context.Context,
	game *model.Game,
) error {
	firstWinner := game.GetFirstPlayerWinner()
	// estimate gas fee for transfer token from game to player winner
	winnerGasFee, err := uc.erc20Usecase.EstimateGasFee(
		ctx,
		game.Address,
		firstWinner.Address,
		game.MaxPrizeAmount.ToBigInt(),
	)
	if err != nil {
		return err
	}

	// estimate gas fee for transfer token from game to setting treasurer address
	treasurerGasFee, err := uc.erc20Usecase.EstimateGasFee(
		ctx,
		game.Address,
		uc.setting.TreasurerAddress,
		game.GameFee.ToBigInt(),
	)
	if err != nil {
		return err
	}

	// (pre-calculate needed transfers and request sufficient gas + buffer for all transfers)
	// (winnerGasFee*game.TotalPlayerWinners + treasurerGasFee) * 110 / 100
	// 1.1 means 10% buffer to avoid out of gas

	totalPlayerWinners := big.NewInt(game.TotalPlayerWinners)
	totalPlayerGasFee := winnerGasFee.Mul(winnerGasFee, totalPlayerWinners)
	totalGasFee := totalPlayerGasFee.Add(totalPlayerGasFee, treasurerGasFee)

	// totalGasFeeWithBuffer = totalGasFee * 110 / 100
	totalGasFeeWithBuffer := totalGasFee.Mul(totalGasFee, big.NewInt(110)).
		Div(totalGasFee, big.NewInt(100))

	// transfer eth fee from operation address to game address
	operationPrivateKey, err := encrypt.DecryptToString(uc.setting.OperationPrivKey, uc.secretKey)
	if err != nil {
		return err
	}

	ethTx, err := uc.erc20Usecase.TransferETH(ctx, game.Address, totalGasFeeWithBuffer, operationPrivateKey)
	if err != nil {
		return err
	}
	logger.GetLoggerInstanceFromContext(ctx).
		Info("estimateGasFeeForWinnersTransfer:transfer_eth_fee", zap.String("tx", ethTx))

	return nil
}

func (uc *gameUsecase) prizeToWinners(ctx context.Context, game *model.Game) error {
	if uc.setting == nil {
		return errors.New("please configure the application setting first")
	}

	logger.GetLoggerInstanceFromContext(ctx).Info("prize_to_winners", zap.String("tweet_id", game.TweetId))
	canCompleteGame := true
	defer func() error {
		if canCompleteGame {
			game.Status = model.GameStatusCompleted
		}
		return uc.updateGame(ctx, game)
	}()

	// there is no participants, then skip
	if game.HasNoParticipants() {
		return nil
	}

	// determine player winner and calculate prize per player winner
	game.DeterminePlayerWinner(uc.setting.GasFeePercentage).
		CalculatePrizePerPlayerWinner()

	// there is no player winners, then refund to players
	if game.HasNoPlayerWinners() {
		err := uc.refundToPlayers(ctx, game)
		if err != nil {
			// if refund to players failed, then can not complete game
			canCompleteGame = false
			return err
		}
		return nil
	}

	err := uc.faucetGasToGameWallet(ctx, game)
	if err != nil {
		canCompleteGame = false
		return err
	}

	// transfer token from game to treasurer address
	priKey, err := encrypt.DecryptToString(game.PrivKey, uc.secretKey)
	if err != nil {
		canCompleteGame = false
		logger.GetLoggerInstanceFromContext(ctx).Error("decrypt_to_string", zap.Error(err))
		return err
	}

	// transfer token from game to treasurer address
	if game.GameFeeTxHash == "" {
		tx, err := uc.erc20Usecase.TransferToken(ctx, uc.setting.TreasurerAddress, game.GameFee.ToBigInt(), priKey)
		if err != nil {
			canCompleteGame = false
			logger.GetLoggerInstanceFromContext(ctx).Error(
				"[prizeToWinners] transfer token from game to treasurer address failed",
				zap.Error(err),
				zap.String("fromAddress", game.Address),
				zap.String("toAddress", uc.setting.TreasurerAddress),
				zap.String("tweetId", game.TweetId),
			)
			return err
		}

		// update game fee tx hash
		game.GameFeeTxHash = tx
		// update current prize result to game
		if err := uc.updateGame(ctx, game); err != nil {
			canCompleteGame = false
			logger.GetLoggerInstanceFromContext(ctx).Error("update_game", zap.Error(err))
			return err
		}
	}

	// wait for 1 second for the next transfer
	time.Sleep(1 * time.Second)

	for _, p := range game.Players {
		if p.PrizeTxHash != "" {
			continue
		}

		if !p.Win {
			continue
		}

		tx, err := uc.erc20Usecase.TransferToken(ctx, p.Address, p.PrizeAmount.ToBigInt(), priKey)
		if err != nil {
			canCompleteGame = false
			logger.GetLoggerInstanceFromContext(ctx).Error(
				"[prizeToWinners] transfer token from game to player winner failed",
				zap.Error(err),
				zap.String("fromAddress", game.Address),
				zap.String("toAddress", p.Address),
				zap.String("tweetId", game.TweetId),
			)
			return err
		}

		p.Win = true
		p.PrizeTxHash = tx

		// update current prize result to game
		if err := uc.updateGame(ctx, game); err != nil {
			canCompleteGame = false
			logger.GetLoggerInstanceFromContext(ctx).Error("[prizeToWinners] update_game failed", zap.Error(err))
			return err
		}

		// wait for 1 second for the next transfer
		time.Sleep(1 * time.Second)
	}

	return nil
}

func (uc *gameUsecase) refundToPlayers(ctx context.Context, game *model.Game) error {
	refundAmountPerPlayer := game.CalculateRefundAmountPerPlayer()
	for _, p := range refundAmountPerPlayer {
		if p.RefundTxHash != "" {
			continue
		}

		err := uc.handlerRefundTokenFromGameToPlayer(ctx, game, p)
		if err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Error(
				"[refundToPlayers] refund token from game to player failed",
				zap.Error(err),
				zap.String("fromAddress", game.Address),
				zap.String("toAddress", p.Address),
				zap.String("tweetId", game.TweetId),
			)
			return err
		}

		// update current refund result to game
		if err := uc.updateGame(ctx, game); err != nil {
			return err
		}
	}

	return nil
}

// handlerRefundTokenFromGameToPlayer transfer token from game to player address
func (uc *gameUsecase) handlerRefundTokenFromGameToPlayer(
	ctx context.Context,
	game *model.Game,
	player *model.Player,
) error {
	amount := player.RefundAmount.ToBigInt()

	// decrypt game private key
	gamePrivateKey, err := encrypt.DecryptToString(game.PrivKey, uc.secretKey)
	if err != nil {
		return err
	}

	// decrypt operation private key
	operationPrivateKey, err := encrypt.DecryptToString(uc.setting.OperationPrivKey, uc.secretKey)
	if err != nil {
		return err
	}

	// estimate gas fee for transfer token then transfer eth fee from operation address to from address
	gasFee, err := uc.erc20Usecase.EstimateGasFee(ctx, game.Address, player.Address, amount)
	if err != nil {
		return err
	}

	// transfer eth fee from operation address to from address
	ethTx, err := uc.erc20Usecase.TransferETH(ctx, game.Address, gasFee, operationPrivateKey)
	if err != nil {
		return err
	}
	logger.GetLoggerInstanceFromContext(ctx).Info("transfer_eth_fee", zap.String("tx", ethTx))

	// transfer token from agent to game address
	tx, err := uc.erc20Usecase.TransferToken(ctx, player.Address, amount, gamePrivateKey)
	if err != nil {
		return err
	}

	// update player refund tx hash
	player.RefundTxHash = tx
	return nil
}

func (uc *gameUsecase) updateGame(ctx context.Context, game *model.Game) error {
	lastUpdatedAt := game.DateModified
	game.DateModified = time.Now()
	return uc.gameRepo.Update(ctx, game, game.Id, lastUpdatedAt)
}

func (uc *gameUsecase) checkGameBalance(ctx context.Context, game *model.Game) error {
	logger.GetLoggerInstanceFromContext(ctx).Info("check_game_balance", zap.String("tweet_id", game.TweetId))
	currentBlock, err := uc.erc20Usecase.CurrentBlockNumber(ctx)
	if err != nil {
		return err
	}

	fromBlock := currentBlock - 10_000
	if game.CurrentBlock != 0 {
		fromBlock = game.CurrentBlock
	}

	iter, err := uc.erc20Usecase.FilterTransfer(ctx, fromBlock, currentBlock, nil, game.AgentAddresses())
	if err != nil {
		return err
	}

	for iter.Next() {
		isSkip := false
		e := iter.Event
		if len(game.Players) == 0 {
			game.Players = make([]*model.Player, 0)
		}

		for _, p := range game.Players {
			// If player is exist, then skipping
			// not checking the case-insensitivity of the address and txHash
			if strings.EqualFold(e.From.Hex(), p.Address) &&
				strings.EqualFold(e.Raw.TxHash.Hex(), p.TxHash) {
				isSkip = true
				break
			}
		}

		if isSkip {
			continue
		}

		game.Players = append(game.Players, &model.Player{
			Address:           utils.BeatifyWalletAddress(e.From.Hex()),
			BetToAgentAddress: utils.BeatifyWalletAddress(e.To.Hex()),
			Amount:            cryptoamount.NewCryptoAmountFromBigInt(e.Value),
			TxHash:            e.Raw.TxHash.Hex(),
		})
	}

	// why we need to for loops here?
	game.CurrentBlock = currentBlock
	agentAmounts := make(map[string]cryptoamount.CryptoAmount)
	for _, p := range game.Players {
		agentAmounts[p.BetToAgentAddress] += p.Amount
	}

	game.TotalAmount = 0
	for _, a := range game.AgentWallets {
		a.Amount = agentAmounts[a.Address]
		game.TotalAmount += a.Amount
	}

	return uc.updateGame(ctx, game)
}

func (uc *gameUsecase) StartGame(ctx context.Context, request *model.StartGameRequest) (*model.Game, error) {
	logger.GetLoggerInstanceFromContext(ctx).Info("start_game", zap.Any("request", request))
	game, err := uc.gameByTweetId(ctx, request.TweetId)
	if err != nil && !utils.IsErrNoDocuments(err) {
		return nil, err
	}

	if game != nil && !game.Id.IsZero() {
		return game, nil
	}

	game = &model.Game{
		TweetId: request.TweetId,
		EndTime: time.Now().Add(time.Duration(request.TimeOut) * time.Second),
	}

	encryptedTempKey, tempAddr, err := utils.GenerateAddress(uc.secretKey)
	if err != nil {
		return nil, err
	}

	game.Address = tempAddr
	game.PrivKey = encryptedTempKey
	game.Status = model.GameStatusRunning

	for _, u := range request.Usernames {
		encryptedTempKey, tempAddr, err := utils.GenerateAddress(uc.secretKey)
		if err != nil {
			return nil, err
		}
		game.AgentWallets = append(
			game.AgentWallets,
			&model.AgentWallet{Username: u, Address: tempAddr, PrivKey: encryptedTempKey},
		)
	}
	game.Id = primitive.NewObjectID()
	game.DateCreated = time.Now()
	game.DateModified = game.DateCreated
	game.StartTime = game.DateCreated

	if _, err := uc.gameRepo.Create(ctx, game); err != nil {
		return nil, err
	}

	return game, nil
}

func (uc *gameUsecase) createSetting(ctx context.Context) {
	uc.setApplicationSetting(ctx)
	if uc.setting != nil {
		return
	}

	// generate operation address, operation private key
	operationPrivKey, operationAddress, err := utils.GenerateAddress(uc.secretKey)
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Fatal("generate_address", zap.Error(err))
	}

	treasurerAddress := viper.GetString("TREASURER_ADDRESS")
	if treasurerAddress == "" {
		logger.GetLoggerInstanceFromContext(ctx).Fatal("TREASURER_ADDRESS is required")
	}

	gameFeePercentage := viper.GetFloat64("GAME_GAS_FEE_PERCENTAGE")
	appSetting := model.NewSetting(treasurerAddress).
		SetOperationAddress(operationAddress).
		SetOperationPrivKey(operationPrivKey).
		SetGasFeePercentage(gameFeePercentage)

	if _, err := uc.settingRepo.Create(ctx, appSetting); err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Fatal("create_setting", zap.Error(err))
	}

	// re-set application setting
	uc.setApplicationSetting(ctx)
}

func (uc *gameUsecase) findSetting(ctx context.Context) (*model.Setting, error) {
	if uc.setting != nil {
		return uc.setting, nil
	}

	setting := &model.Setting{}
	if err := uc.settingRepo.FindOne(ctx, bson.M{}, setting); err != nil {
		return nil, err
	}

	// load some settings from env
	treasureAddress := viper.GetString("TREASURER_ADDRESS")
	gasFeePercentage := viper.GetFloat64("GAME_GAS_FEE_PERCENTAGE")
	if treasureAddress != "" {
		setting.SetTreasurerAddress(treasureAddress)
	}
	if gasFeePercentage != 0 {
		setting.SetGasFeePercentage(gasFeePercentage)
	}

	uc.setting = setting
	return setting, nil
}

func (uc *gameUsecase) setApplicationSetting(ctx context.Context) {
	setting, err := uc.findSetting(ctx)
	if err != nil {
		logger.AtLog.Debug("Can not get setting %v", err)
	} else {
		uc.setting = setting
	}
}

var Module = fx.Module(
	"game_usecase",
	mongo.GameRepoModule,
	mongo.SettingRepoModule,
	usecase.Module,
	fx.Provide(NewGameUsecase),
)

func NewGameUsecase(
	gameRepo mongo.IGameRepo,
	settingRepo mongo.ISettingRepo,
	erc20Usecase port.IContractErc20Usecase,
) port.IGameUsecase {
	ctx := context.Background()
	secretKey := viper.GetString("SECRET_KEY")
	var googleSecretKey string
	if utils.IsEnvProduction() {
		key, err := secret_manager.GetGoogleSecretKey(ctx, secretKey)
		googleSecretKey = key
		if err != nil {
			logger.GetLoggerInstanceFromContext(ctx).Fatal("GetGoogleSecretKey", zap.Error(err))
		}
	} else {
		decodedBase64, err := base64.StdEncoding.DecodeString(secretKey)
		if err != nil {
			logger.AtLog.Fatalf("Can not decode secret key %v", err)
		}
		googleSecretKey = string(decodedBase64)
	}

	uc := &gameUsecase{
		gameRepo:     gameRepo,
		settingRepo:  settingRepo,
		secretKey:    googleSecretKey,
		erc20Usecase: erc20Usecase,
	}
	uc.createSetting(ctx)
	return uc
}
