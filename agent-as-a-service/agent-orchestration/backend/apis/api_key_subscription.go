package apis

import (
	"net/http"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/gin-gonic/gin"
)

func (s *Server) CreateApiTokenUsage(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.ApiTokenUsageReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	err := s.nls.CreateApiTokenUsage(ctx, req.NetworkID, req.ApiKey, req.Endpoint, req.NumToken)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: true})
}

func (s *Server) RefundApiTokenUsage(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.ApiTokenUsageReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	err := s.nls.RefundApiTokenUsage(ctx, req.ApiKey, req.NumToken)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: true})
}

func (s *Server) GetApiUsage(c *gin.Context) {
	ctx := s.requestContext(c)
	page, limit := s.pagingFromContext(c)
	apiKey := s.stringFromContextQuery(c, "api_key")
	lst, count, err := s.nls.GetApiUsage(ctx, apiKey, page, limit)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewApiSubscriptionUsageLogRespArr(lst), Count: &count})
}

func (s *Server) GetApiPackages(c *gin.Context) {
	ctx := s.requestContext(c)
	packages, err := s.nls.GetApiPackages(ctx)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewApiSubscriptionPackageRespArr(packages)})
}

func (s *Server) GetApiSubscriptionInfo(c *gin.Context) {
	ctx := s.requestContext(c)
	apiKey := s.stringFromContextQuery(c, "api_key")
	address := s.stringFromContextQuery(c, "address")
	obj, err := s.nls.GetApiKeyInfo(ctx, apiKey, address)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewApiSubscriptionKeyRespArr(obj)})
}

func (s *Server) CreateAcctForTest(c *gin.Context) {
	ctx := s.requestContext(c)
	address := s.stringFromContextQuery(c, "address")
	twitterID := s.stringFromContextQuery(c, "twitter_id")
	obj, err := s.nls.CreateApiSubscriptionKeyForTest(ctx, address, twitterID)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewApiSubscriptionKeyResp(obj)})
}
