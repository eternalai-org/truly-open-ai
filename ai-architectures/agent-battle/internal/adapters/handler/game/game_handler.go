package game

import (
	"context"
	"time"

	"agent-battle/internal/core/model"
	"agent-battle/internal/core/port"
	"agent-battle/pkg/constants"
	"agent-battle/pkg/rest"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type gameHandler struct {
	validator   *validator.Validate
	gameUsecase port.IGameUsecase
}

func NewGameHandler(router fiber.Router, gameUsecase port.IGameUsecase) {
	handler := gameHandler{
		validator:   validator.New(),
		gameUsecase: gameUsecase,
	}

	r := router.Group("/game")
	{
		r.Get("", handler.listGame)
		r.Post("/:tweet_id/end", handler.endGame)
		r.Post("/:tweet_id/result", handler.gameResult)
		r.Post("/:tweet_id/refund-expired-players", handler.refundExpiredPlayers)
		r.Post("/start", handler.startGame)
		r.Get("/:tweet_id", handler.detailGame)
	}
}

// @Summary List game
// @Tags Game
// @Accept json
// @Produce json
// @Param tweet_id query string false "Tweet ID"
// @Success 200 {object} model.ListGameResponse
// @Router /api/v1/game [get]
func (h gameHandler) listGame(ctx *fiber.Ctx) error {
	return rest.NewFiberHandlerTemplate(
		func(ctx *fiber.Ctx, _ string) (interface{}, error) {
			c, cancel := context.WithTimeout(ctx.UserContext(), constants.FiberRequestTimeoutInSec*time.Second)
			defer cancel()

			req := &model.ListGameRequest{}
			if err := ctx.QueryParser(req); err != nil {
				return nil, err
			}

			return h.gameUsecase.ListGame(c, req)
		}).ResponseJSON(ctx)
}

// @Summary Detail game
// @Tags Game
// @Accept json
// @Produce json
// @Param tweet_id path string true "Tweet ID"
// @Success 200 {object} model.Game
// @Router /api/v1/game/{tweet_id} [get]
func (h gameHandler) detailGame(ctx *fiber.Ctx) error {
	return rest.NewFiberHandlerTemplate(
		func(ctx *fiber.Ctx, _ string) (interface{}, error) {
			c, cancel := context.WithTimeout(ctx.UserContext(), constants.FiberRequestTimeoutInSec*time.Second)
			defer cancel()

			return h.gameUsecase.DetailGame(c, ctx.Params("tweet_id"))
		}).ResponseJSON(ctx)
}

// @Summary Game result
// @Tags Game
// @Accept json
// @Produce json
// @Param tweet_id path string true "Tweet ID"
// @Param game_result body model.GameResultRequest true "Game Result"
// @Success 200 {object} model.Game
// @Router /api/v1/game/{tweet_id}/result [post]
func (h gameHandler) gameResult(ctx *fiber.Ctx) error {
	return rest.NewFiberHandlerTemplate(
		func(ctx *fiber.Ctx, _ string) (interface{}, error) {
			c, cancel := context.WithTimeout(ctx.UserContext(), constants.FiberRequestTimeoutInSec*time.Second)
			defer cancel()

			req := &model.GameResultRequest{}
			if err := ctx.BodyParser(req); err != nil {
				return nil, err
			}
			req.TweetId = ctx.Params("tweet_id")
			return h.gameUsecase.GameResult(c, req)
		}).ResponseJSON(ctx)
}

// @Summary End game
// @Tags Game
// @Accept json
// @Produce json
// @Param tweet_id path string true "Tweet ID"
// @Success 200 {object} model.Game
// @Router /api/v1/game/{tweet_id}/end [post]
func (h gameHandler) endGame(ctx *fiber.Ctx) error {
	return rest.NewFiberHandlerTemplate(
		func(ctx *fiber.Ctx, _ string) (interface{}, error) {
			c, cancel := context.WithTimeout(ctx.UserContext(), constants.FiberRequestTimeoutInSec*time.Second)
			defer cancel()

			return h.gameUsecase.EndGame(c, ctx.Params("tweet_id"))
		}).ResponseJSON(ctx)
}

// @Summary Start game
// @Tags Game
// @Accept json
// @Produce json
// @Param start_game body model.StartGameRequest true "Start Game"
// @Success 200 {object} model.Game
// @Router /api/v1/game/start [post]
func (h gameHandler) startGame(ctx *fiber.Ctx) error {
	return rest.NewFiberHandlerTemplate(
		func(ctx *fiber.Ctx, _ string) (interface{}, error) {
			requestCtx, cancel := context.WithTimeout(ctx.UserContext(), constants.FiberRequestTimeoutInSec*time.Second)
			defer cancel()

			req := &model.StartGameRequest{}
			if err := ctx.BodyParser(req); err != nil {
				return nil, err
			}

			return h.gameUsecase.StartGame(requestCtx, req)
		}).ResponseJSON(ctx)
}

// @Summary Refund expired players
// @Tags Game
// @Accept json
// @Produce json
// @Param tweet_id path string true "Tweet ID"
// @Success 200
// @Router /api/v1/game/{tweet_id}/refund-expired-players [post]
func (h gameHandler) refundExpiredPlayers(ctx *fiber.Ctx) error {
	return rest.NewFiberHandlerTemplate(
		func(ctx *fiber.Ctx, _ string) (interface{}, error) {
			c, cancel := context.WithTimeout(ctx.UserContext(), constants.FiberRequestTimeoutInSec*time.Second)
			defer cancel()

			return nil, h.gameUsecase.RefundsExpiredPlayers(c, ctx.Params("tweet_id"))
		}).ResponseJSON(ctx)
}