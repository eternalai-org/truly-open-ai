package apis

import (
	"net/http"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/gin-gonic/gin"
)

func (s *Server) ExternalWalletCreateSOL(c *gin.Context) {
	ctx := s.requestContext(c)
	m, err := s.nls.ExternalWalletCreateSOL(ctx)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: m})
}

func (s *Server) ExternalWalletGet(c *gin.Context) {
	ctx := s.requestContext(c)
	m, err := s.nls.ExternalWalletGet(ctx, c.GetHeader("api-key"))
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: m})
}

func (s *Server) ExternalWalletBalances(c *gin.Context) {
	ctx := s.requestContext(c)
	m, err := s.nls.ExternalWalletBalances(ctx, c.GetHeader("api-key"))
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: m})
}

func (s *Server) ExternalWalletComputeOrder(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.ExternalWalletOrderReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	m, err := s.nls.ExternalWalletComputeOrder(ctx, c.GetHeader("api-key"), &req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: m})
}

func (s *Server) ExternalWalletCreateOrder(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.ExternalWalletOrderReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	m, err := s.nls.ExternalWalletCreateOrder(ctx, c.GetHeader("api-key"), &req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewExternalWalletOrderResp(m)})
}

func (s *Server) ExternalWalletGetOrders(c *gin.Context) {
	ctx := s.requestContext(c)
	page, limit := s.pagingFromContext(c)
	ms, err := s.nls.ExternalWalletGetOrders(ctx, c.GetHeader("api-key"), page, limit)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewExternalWalletOrderRespArr(ms)})
}

func (s *Server) ExternalWalletGetTokens(c *gin.Context) {
	ctx := s.requestContext(c)
	ms, err := s.nls.ExternalWalletGetTokens(ctx, c.GetHeader("api-key"))
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewExternalWalletTokenRespArr(ms)})
}
