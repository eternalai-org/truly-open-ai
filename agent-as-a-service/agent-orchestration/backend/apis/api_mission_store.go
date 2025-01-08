package apis

import (
	"net/http"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/gin-gonic/gin"
)

func (s *Server) UploadMissionStore(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.MissionStoreReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	s.nls.UploadMisstionStore(ctx, &req)
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: true})
}

func (s *Server) RateMissionStore(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.MissionStoreRatingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	s.nls.RatingMisstionStore(ctx, &req)
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: true})
}

func (s *Server) GetListMissionStore(c *gin.Context) {
	ctx := s.requestContext(c)
	page, limit := s.pagingFromContext(c)
	s.nls.GetListMisstionStore(ctx, page, limit)
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: true})
}

func (s *Server) ClaimFeeMissionStore(c *gin.Context) {
	ctx := s.requestContext(c)
	s.nls.ClaimFeeMisstionStore(ctx)
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: true})
}
