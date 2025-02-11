package game_usecase

import (

"agent-battle/internal/adapters/repository/mongo"
"agent-battle/internal/contract/erc20/usecase"
"agent-battle/internal/core/model"
"agent-battle/internal/core/port"
"agent-battle/pkg/cryptoamount"
"agent-battle/pkg/drivers/mongodb"
"agent-battle/pkg/logger"
"context"
"fmt"
"github.com/spf13/viper"
"testing"
"time"
)

var (
	erc20Usecase port.IContractErc20Usecase
	gameRepo     mongo.IGameRepo
	settingRepo  mongo.ISettingRepo
	gameUseCase  port.IGameUsecase
)

func init() {
	viper.SetConfigFile(`../../../../env/local.worker.test.yml`)
	viper.ReadInConfig()
	initVars()
}

func initVars() {
	var err error
	erc20Usecase, err = usecase.NewContractErc20Usecase()
	if err != nil {
		panic(err)
	}

	db, err := mongodb.Init()
	if err != nil {
		panic(err)
	}

	gameRepo = mongo.NewGameRepo(db)
	settingRepo = mongo.NewSettingRepo(db)
	gameUseCase = NewGameUsecase(gameRepo, settingRepo, erc20Usecase)
}

/*Test_gameUsecase_GameFullFlow tests the full flow of the game

This test will start a game, bet to the game, watch the game state to update the player's status,
end the game, update the result and prize to winners.
 - There are two agents and three players in the game
 - The first agent wins the game
 - Each player bets 1 EAI token
 - The total amount is equal to 3 EAI tokens
 - The game fee is 5% of the total amount -> 0.15 EAI tokens
 - The total prize amount is 2.85 EAI tokens
 - The total bet amount winners is 2 EAI tokens
 - The first winner player will get 1.425 EAI tokens
 - The second winner player will get 1.425 EAI tokens
 */
func Test_gameUsecase_GameFullFlow(t *testing.T) {
	ctx := context.Background()

	// Start game
	game, err := startGame()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fmt.Sprintf("Game started: %s", game.TweetId))

	// Bet to game
	err = betToGame(game)
	if err != nil {
		t.Fatal(err)
	}

	// Watch game state to update the player's status
	err = gameUseCase.WatchGameState(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// check if there are three players in the game
	latestGame, err := gameUseCase.DetailGame(ctx, game.TweetId)
	if err != nil {
		t.Fatal(err)
	}

	if len(latestGame.Players) != 3 {
		t.Fatal("There should be three players in the game")
	}

	// End game
	endedGame, err := handlerEndGame(latestGame)
	if err != nil {
		t.Fatal(err)
	}
	for _, agent := range endedGame.AgentWallets {
		if agent.TxHash == "" {
			t.Fatal(fmt.Sprintf("Agent %s should have a transaction hash", agent.Username))
		}

		if agent.TransferAmount == 0 {
			t.Fatal(fmt.Sprintf("Agent %s should have a transfer amount", agent.Username))
		}
	}
	t.Log(fmt.Sprintf("Game ended: %s", endedGame.TweetId))


	// Update result and prize to winners
	updatedGame, err := updateResultAndPrizeToWinners(endedGame)
	if err != nil {
		t.Fatal(err)
	}

	// total amount is equal to 3 EAI tokens
	if updatedGame.TotalAmount != cryptoamount.CryptoAmount(3e18) {
		t.Fatal("Total amount should be equal to 3 EAI tokens")
	}

	// the game fee is 5% of the total amount
	if updatedGame.GameFee != cryptoamount.CryptoAmount(0.15e18) {
		t.Fatal("Game fee should be equal to 0.15 EAI tokens")
	}

	// the total prize amount is 2.85 EAI tokens
	if updatedGame.TotalPrizeAmount != cryptoamount.CryptoAmount(2.85e18) {
		t.Fatal("Total prize amount should be equal to 2.85 EAI tokens")
	}

	// because the first agent wins the game, the total bet amount winners is 2 EAI tokens
	if updatedGame.TotalBetAmountWinners != cryptoamount.CryptoAmount(2e18) {
		t.Fatal("Total bet amount winners should be equal to 2 EAI tokens")
	}

	for _, player := range updatedGame.Players {
		if !player.Win {
			continue
		}

		if player.PrizeTxHash == "" {
			t.Fatal("The winner should have a prize transaction hash")
		}

		// winner will get (1/2) * 2.85 = 1.425 EAI tokens
		if player.PrizeAmount != cryptoamount.CryptoAmount(1.425e18) {
			t.Fatal("The winner should get 1.425 EAI tokens")
		}
	}
}

// Start game
func startGame() (*model.Game, error){
	randTweetId := fmt.Sprintf("%d", time.Now().Unix())

	game, err := gameUseCase.StartGame(context.Background(), &model.StartGameRequest{
		TimeOut: 6000, // 6000 seconds = 100 minutes
		TweetId: randTweetId,
		Usernames: []string{
			"test1",
			"test2",
		},
	})
	if err != nil {
		return nil, err
	}

	return game, nil
}

// Bet to game
func betToGame(game *model.Game) error {
	// Bet to game
	firstAgent := game.AgentWallets[0]
	secondAgent := game.AgentWallets[1]

	// Transfer funds from player to agent
	// The first agent has 2 EAI tokens
	// The second agent has 1 EAI token
	err := transferFundsFromPlayerToAgent(firstAgent, "FIRST_")
	if err != nil {
		return err
	}

	err = transferFundsFromPlayerToAgent(firstAgent, "SECOND_")
	if err != nil {
		return err
	}

	err = transferFundsFromPlayerToAgent(secondAgent, "SECOND_")
	if err != nil {
		return err
	}

	return nil
}

// Transfer funds from player to agent
func transferFundsFromPlayerToAgent(agent *model.AgentWallet, playerPrefixEnvKey string) error {
	encryptedPrivKey := viper.GetString(playerPrefixEnvKey + "PLAYER_ENCRYPTED_PRIVATE_KEY")
	transferAmount := cryptoamount.CryptoAmount(1e18) // 1 EAI token

	// Transfer funds from player to agent
	signature, err := erc20Usecase.TransferToken(context.Background(), agent.Address, transferAmount.ToBigInt(), encryptedPrivKey)
	if err != nil {
		return err
	}

	logger.GetLoggerInstanceFromContext(context.Background()).
		Info(fmt.Sprintf("Transfer funds from player to agent: %s", signature))
	return nil
}

// End game
func handlerEndGame(game *model.Game) (*model.Game, error) {
	for {
		ctx := context.Background()
		// End game
		updatedGame, _ := gameUseCase.EndGame(ctx, game.TweetId)
		if updatedGame != nil && updatedGame.Status == model.GameStatusEnded {
			return updatedGame, nil
		}
		time.Sleep(10 * time.Second)
	}
}

// Update result and prize to winners
func updateResultAndPrizeToWinners(game *model.Game) (*model.Game, error) {
	ctx := context.Background()

	// Update result and prize to winners
	gameResultRequest := &model.GameResultRequest{
		TweetId: game.TweetId,
		Username: "test1", // The first agent wins the game
	}
	_, err := gameUseCase.GameResult(ctx, gameResultRequest)
	if err != nil {
		return nil, err
	}

	for {
		// watch game state
		err = gameUseCase.WatchGameState(ctx)
		if err != nil {
			return nil, err
		}
		
		// get the latest game
		latestGame, err := gameUseCase.DetailGame(ctx, game.TweetId)
		if err != nil {
			return nil, err
		}

		if latestGame.TotalPlayerWinners > 0 && latestGame.Status == model.GameStatusCompleted {
			return latestGame, nil
		}
		
		time.Sleep(10 * time.Second)
	}
}