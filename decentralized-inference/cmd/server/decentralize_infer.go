package server

import (
	"decentralized-inference/internal/client"
	"decentralized-inference/internal/config"
	"decentralized-inference/internal/models"
	"decentralized-inference/internal/rest"
	"net/http"

	"decentralized-inference/internal/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (rt *Server) CreateDecentralizeInfer(c *gin.Context) {
	rest.ResponseJSON(
		func(ctx *gin.Context) (interface{}, error) {
			request := &models.DecentralizeInferRequest{}
			if err := c.ShouldBindJSON(request); err != nil {
				return nil, rest.NewHttpErr(err, http.StatusBadRequest)
			}

			resp, err := rt.Service.CreateDecentralizeInfer(c.Request.Context(), request)
			if err != nil {
				logger.GetLoggerInstanceFromContext(c.Request.Context()).Error("CreateDecentralizeInfer error", zap.Error(err))
				return nil, rest.NewHttpErr(err, http.StatusInternalServerError)
			}

			logger.GetLoggerInstanceFromContext(c.Request.Context()).Info("CreateDecentralizeInfer success", zap.Any("response", resp))
			return resp, nil
		}, c)
}

func (rt *Server) GetDecentralizeInferResult(c *gin.Context) {
	rest.ResponseJSON(
		func(ctx *gin.Context) (interface{}, error) {
			request := &models.InferResultRequest{}
			if err := c.ShouldBindJSON(request); err != nil {
				return nil, rest.NewHttpErr(err, http.StatusBadRequest)
			}

			return rt.Service.GetDecentralizeInferResult(c.Request.Context(), request)
		}, c)
}

func (rt *Server) InsertChainConfig(c *gin.Context) {
	rest.ResponseJSON(
		func(ctx *gin.Context) (interface{}, error) {
			request := &config.ChatConfig{}
			if err := c.ShouldBindJSON(request); err != nil {
				return nil, rest.NewHttpErr(err, http.StatusBadRequest)
			}
			cli, err := client.NewClient(request.Rpc, models.ChainTypeEth, false, "", "")
			if err != nil {
				return nil, rest.NewHttpErr(err, http.StatusBadRequest)
			}
			chainID, err := cli.Client.ChainID(c.Request.Context())
			if err != nil {
				return nil, rest.NewHttpErr(err, http.StatusBadRequest)
			}
			err = rt.Service.RemoveAllChainConfig(c.Request.Context())
			if err != nil {
				return nil, rest.NewHttpErr(err, http.StatusBadRequest)
			}
			chainConfig := models.ChainConfig{
				ChainID:              chainID.String(),
				ListRPC:              []string{request.Rpc},
				AgentContractAddress: request.Contracts.SystemPromptManagerAddress,
				WorkerHubAddress:     request.Contracts.WorkerHubAddress,
				Type:                 models.ChainTypeEth,
			}
			return nil, rt.Service.InsertChainConfig(c.Request.Context(), &chainConfig)
		}, c)
}
