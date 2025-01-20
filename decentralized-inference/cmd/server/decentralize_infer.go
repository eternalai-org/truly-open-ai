package server

import (
	"decentralized-inference/internal/models"
	"decentralized-inference/internal/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (rt *Server) CreateDecentralizeInfer(c *gin.Context) {
	rest.ResponseJSON(
		func(ctx *gin.Context) (interface{}, error) {
			request := &models.DecentralizeInferRequest{}
			if err := c.ShouldBindQuery(request); err != nil {
				return nil, rest.NewHttpErr(err, http.StatusBadRequest)
			}

			return rt.Service.CreateDecentralizeInfer(c.Request.Context(), request)
		}, c)
}

func (rt *Server) GetDecentralizeInferResult(c *gin.Context) {
	rest.ResponseJSON(
		func(ctx *gin.Context) (interface{}, error) {
			request := &models.InferResultRequest{}
			if err := c.ShouldBindQuery(request); err != nil {
				return nil, rest.NewHttpErr(err, http.StatusBadRequest)
			}

			return rt.Service.GetDecentralizeInferResult(c.Request.Context(), request)
		}, c)
}
