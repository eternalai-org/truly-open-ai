package server

import (
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
