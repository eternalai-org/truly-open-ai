package model

import (
	"agent-battle/pkg/cryptoamount"
	"math"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type GameStatus int64

const (
	GameStatusRunning GameStatus = iota + 1
	GameStatusEnded
	GameStatusResultUpdated
	GameStatusCompleted
)

type ListGameRequest struct {
	Pagination `query:",inline" json:",inline"`
	TweetIds   []string `json:"tweet_ids" query:"tweet_ids"`
}

type ListGameResponse struct {
	Games        []*Game `json:"games"`
	TotalRecords int64   `json:"total_records"`
}

type Game struct {
	Model                 `bson:",inline" json:",inline"`
	TweetId               string                    `json:"tweet_id" bson:"tweet_id"`
	StartTime             time.Time                 `json:"start_time" bson:"start_time"`
	EndTime               time.Time                 `json:"end_time" bson:"end_time"`
	BetEndTime            time.Time                 `json:"bet_end_time" bson:"bet_end_time"`
	AgentWallets          []*AgentWallet            `json:"agent_wallets" bson:"agent_wallets"`
	Status                GameStatus                `json:"status" bson:"status"`
	Players               []*Player                 `json:"players" bson:"players"`
	Address               string                    `json:"-" bson:"address,omitempty"`
	PrivKey               string                    `json:"-" bson:"priv_key,omitempty"`
	Winner                string                    `json:"winner" bson:"winner,omitempty"`
	TotalAmount           cryptoamount.CryptoAmount `json:"-" bson:"total_amount,omitempty"`
	TotalPlayerWinners    int64                     `json:"total_player_winners" bson:"total_player_winners,omitempty"`
	GameFee               cryptoamount.CryptoAmount `json:"-" bson:"game_fee,omitempty"`
	TotalPrizeAmount      cryptoamount.CryptoAmount `json:"-" bson:"total_prize_amount,omitempty"`
	MaxPrizeAmount        cryptoamount.CryptoAmount `json:"-" bson:"max_prize_amount,omitempty"`
	TotalBetAmountWinners cryptoamount.CryptoAmount `json:"-" bson:"total_bet_amount_winners,omitempty"`
	GameFeeTxHash         string                    `json:"-" bson:"game_fee_tx_hash,omitempty"`
	CurrentBlock          uint64                    `json:"-" bson:"current_block,omitempty"`

	// ExpiredPlayers stores the players who have not bet within the bet time out
	ExpiredPlayers []*Player `json:"-" bson:"expired_players,omitempty"`
}

func (Game) CollectionName() string {
	return "game"
}

func (a Game) GetWinner() *AgentWallet {
	var winner *AgentWallet
	for _, agent := range a.AgentWallets {
		if agent.Username == a.Winner {
			winner = agent
			break
		}
	}

	return winner
}

// CalculatePrizePerPlayerWinner calculates the prize per player winner and game fee
func (a *Game) CalculatePrizePerPlayerWinner() *Game {
	if a.TotalAmount == 0 && a.TotalPlayerWinners == 0 {
		return a
	}

	// calculate prize amount per player winner
	// Tỉ lệ chia theo bet porpotion
	// F(w) = Tỉ lệ bet của những đứa win / tổng số bet win
	// Giá trị token nhận đc thực tế
	// F(r) = F(w) *( total bet - game fee)
	scannedPlayers := int64(0)
	remainingPrizeAmount := a.TotalPrizeAmount
	for _, p := range a.Players {
		if !p.Win {
			continue
		}

		// update prize amount based on the portion of the bet amount
		p.PrizeAmount = (p.Amount * a.TotalPrizeAmount / a.TotalBetAmountWinners).Round(0)
		scannedPlayers++
		if scannedPlayers == a.TotalPlayerWinners {
			// if the last player, update the prize amount with the remaining prize amount to avoid rounding errors
			if p.PrizeAmount > remainingPrizeAmount {
				p.PrizeAmount = remainingPrizeAmount
			}
		} else {
			// update remaining prize amount
			remainingPrizeAmount -= p.PrizeAmount
			remainingPrizeAmount = remainingPrizeAmount.Round(0)
		}

		// update the max prize amount, this is used to estimate gas fee for the prize transaction
		a.MaxPrizeAmount = cryptoamount.CryptoAmount(math.Max(
			a.MaxPrizeAmount.ToFloat64(),
			p.PrizeAmount.ToFloat64()),
		).Round(0)
	}
	return a
}

// CalculateRefundAmountPerPlayer calculates the refund amount per player address
// Incase the game has no winner, the refund amount is the same as the bet amount
func (a *Game) CalculateRefundAmountPerPlayer() map[string]*Player {
	refundAmountPerPlayer := make(map[string]*Player)
	for _, p := range a.Players {
		if player, ok := refundAmountPerPlayer[p.Address]; ok {
			player.RefundAmount += p.Amount
		} else {
			p.RefundAmount = p.Amount
			refundAmountPerPlayer[p.Address] = p
		}
	}
	return refundAmountPerPlayer
}

func (a *Game) CalculateRefundAmountPerExpiredPlayer() map[string]*Player {
	refundAmountPerPlayer := make(map[string]*Player)
	for _, p := range a.ExpiredPlayers {
		if player, ok := refundAmountPerPlayer[p.Address]; ok {
			player.RefundAmount += p.Amount
		} else {
			p.RefundAmount = p.Amount
			refundAmountPerPlayer[p.Address] = p
		}
	}
	return refundAmountPerPlayer
}

// DeterminePlayerWinner determines the player winner based on the agent winner
func (a *Game) DeterminePlayerWinner(gasFeePercentage float64) *Game {
	winner := a.GetWinner()
	if winner == nil {
		return a
	}

	totalWinners := 0
	totalBetAmountWinners := cryptoamount.CryptoAmount(0)
	for _, p := range a.Players {
		if p.BetToAgentAddress == winner.Address {
			p.Win = true
			totalWinners++
			totalBetAmountWinners += p.Amount
		}
	}
	a.TotalPlayerWinners = int64(totalWinners)
	a.TotalBetAmountWinners = totalBetAmountWinners

	// calculate game fee and total prize amount
	gameFee := (a.TotalAmount * cryptoamount.CryptoAmount(gasFeePercentage)).Round(0)
	a.GameFee = gameFee
	a.TotalPrizeAmount = a.TotalAmount - gameFee

	return a
}

// HasNoParticipants returns true if the game has no participants
func (a Game) HasNoParticipants() bool {
	return len(a.Players) == 0
}

// HasNoPlayerWinners returns true if the game has no player winners
func (a Game) HasNoPlayerWinners() bool {
	return a.TotalPlayerWinners == 0
}

// GetFirstPlayerWinner returns the first player who wins the game
func (a Game) GetFirstPlayerWinner() *Player {
	for _, p := range a.Players {
		if p.Win {
			return p
		}
	}
	return nil
}

type Player struct {
	Address            string                    `json:"address" bson:"address"`
	Amount             cryptoamount.CryptoAmount `json:"amount" bson:"amount"`
	TxHash             string                    `json:"tx_hash" bson:"tx_hash,omitempty"`
	BetToAgentAddress  string                    `bson:"bet_to_agent_address" json:"bet_to_agent_address"`
	BetToAgentUsername string                    `json:"bet_to_agent_username" bson:"bet_to_agent_username,omitempty"`
	PrizeTxHash        string                    `json:"prize_tx_hash" bson:"prize_tx_hash,omitempty"`
	PrizeAmount        cryptoamount.CryptoAmount `json:"prize_amount" bson:"prize_amount,omitempty"`
	RefundTxHash       string                    `json:"refund_tx_hash" bson:"refund_tx_hash,omitempty"`
	RefundAmount       cryptoamount.CryptoAmount `json:"refund_amount" bson:"refund_amount,omitempty"`
	Win                bool                      `json:"win" bson:"win,omitempty"`
}

type AgentWallet struct {
	Username       string                    `json:"username" bson:"username"`
	Address        string                    `json:"address" bson:"address"`
	PrivKey        string                    `json:"-" bson:"priv_key"`
	Amount         cryptoamount.CryptoAmount `json:"amount" bson:"amount"`
	TxHash         string                    `json:"tx_hash" bson:"tx_hash"`
	TransferAmount cryptoamount.CryptoAmount `json:"transfer_amount" bson:"transfer_amount"`
}

func (a *AgentWallet) CannotTransferFundsToGameWallet() bool {
	// If the agent wallet has enough funds to transfer to the game wallet, return false
	if a.Amount == 0 {
		return true
	}

	// if funds already transferred to the game wallet, return true
	return a.TransferAmount > 0
}

type StartGameRequest struct {
	TweetId    string   `json:"tweet_id" query:"tweet_id" bson:"tweet_id"`
	TimeOut    uint64   `json:"time_out" bson:"time_out"`
	BetTimeOut uint64   `json:"bet_time_out" bson:"bet_time_out"`
	Usernames  []string `json:"usernames" query:"usernames" bson:"usernames"`
}

type GameResultRequest struct {
	TweetId  string `json:"tweet_id" query:"-"`
	Username string `json:"username" query:"username"`
}

func (a *Game) AgentAddresses() []common.Address {
	var addresses []common.Address
	for _, w := range a.AgentWallets {
		addresses = append(addresses, common.HexToAddress(w.Address))
	}
	return addresses
}
