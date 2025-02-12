package apis

import (
	"net/http"
	"strconv"

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
	err := s.nls.UploadMisstionStore(ctx, &req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: true})
}

func (s *Server) RateMissionStore(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.MissionStoreRatingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	err := s.nls.RatingMisstionStore(ctx, &req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: true})
}

func (s *Server) GetListMissionStore(c *gin.Context) {
	ctx := s.requestContext(c)
	page, limit := s.pagingFromContext(c)
	search := s.stringFromContextQuery(c, "search")
	res, count, err := s.nls.GetListMisstionStore(ctx, search, page, limit)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewMissionStoreRespArray(res), Count: &count})
}

func (s *Server) GetMissionStoreDetail(c *gin.Context) {
	ctx := s.requestContext(c)
	idStr := s.stringFromContextParam(c, "id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	res, err := s.nls.GetMisstionStoreDetail(ctx, uint(id))
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewMissionStoreResp(res)})
}

func (s *Server) GetMissionStoreRating(c *gin.Context) {
	ctx := s.requestContext(c)
	missionStoreIDstr := s.stringFromContextParam(c, "id")
	missionStoreID, _ := strconv.ParseUint(missionStoreIDstr, 10, 64)
	page, limit := s.pagingFromContext(c)
	res, count, err := s.nls.GetMissionStoreRating(ctx, uint(missionStoreID), page, limit)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewMissionStoreRatingRespArray(res), Count: &count})

}

func (s *Server) ClaimFeeMissionStore(c *gin.Context) {
	ctx := s.requestContext(c)
	s.nls.ClaimFeeMisstionStore(ctx)
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: true})
}

func (s *Server) GetMissionStoreHistory(c *gin.Context) {
	ctx := s.requestContext(c)
	missionStoreIDstr := s.stringFromContextParam(c, "id")
	missionStoreID, _ := strconv.ParseUint(missionStoreIDstr, 10, 64)
	page, limit := s.pagingFromContext(c)
	res, count, err := s.nls.GetMissionStoreHistory(ctx, uint(missionStoreID), page, limit)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewMissionStoreHistoryRespArray(res), Count: &count})

}

func (s *Server) SaveAgentStore(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.AgentStoreReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil || userAddress == "" {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
		return
	}
	obj, err := s.nls.SaveAgentStore(ctx, userAddress, &req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewAgentStoreResp(obj)})
}

func (s *Server) GetListAgentStore(c *gin.Context) {
	ctx := s.requestContext(c)
	page, limit := s.pagingFromContext(c)
	search := s.stringFromContextQuery(c, "search")
	res, count, err := s.nls.GetListAgentStore(ctx, search, s.stringFromContextQuery(c, "type"), page, limit)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewAgentStoreRespArray(res), Count: &count})
}

func (s *Server) GetListAgentStoreByOwner(c *gin.Context) {
	ctx := s.requestContext(c)
	page, limit := s.pagingFromContext(c)
	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	res, count, err := s.nls.GetListAgentStoreByOwner(ctx, userAddress, page, limit)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewAgentStoreRespArray(res), Count: &count})
}

func (s *Server) GetAgentStoreDetail(c *gin.Context) {
	ctx := s.requestContext(c)
	res, err := s.nls.GetAgentStoreDetail(ctx, s.uintFromContextParam(c, "id"))
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewAgentStoreResp(res)})
}

func (s *Server) SaveMissionStore(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.AgentStoreMissionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil || userAddress == "" {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
		return
	}
	err = s.nls.SaveMissionStore(ctx, userAddress, s.uintFromContextParam(c, "id"), &req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: true})
}

func (s *Server) AuthenAgentStoreCallback(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.AuthenAgentStoreCallback
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	err := s.nls.SaveAgentStoreCallback(ctx, &req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: true})
}

func (s *Server) GetListAgentStoreInstall(c *gin.Context) {
	ctx := s.requestContext(c)
	page, limit := s.pagingFromContext(c)
	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	res, count, err := s.nls.GetListAgentStoreInstall(ctx, userAddress, s.uintFromContextQuery(c, "agent_info_id"), page, limit)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewAgentStoreRespArrayFromInstall(res), Count: &count})
}

func (s *Server) GetAgentStoreInstallCode(c *gin.Context) {
	ctx := s.requestContext(c)
	agentStoreID := s.uintFromContextParam(c, "id")
	agentInfoID := s.uintFromContextParam(c, "agent_info_id")
	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil || userAddress == "" {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
		return
	}
	res, err := s.nls.CreateAgentStoreInstallCode(ctx, userAddress, agentStoreID, agentInfoID)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: res.Code})
}

func (s *Server) RunMission(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.AgentStoreMissionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil || userAddress == "" {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
		return
	}
	obj, err := s.nls.AgentSnapshotPostCreateForUser(ctx, req.NetworkID, userAddress, req.Prompt, req.Model, req.ID)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: obj.ResponseId})
}

func (s *Server) MissionStoreResult(c *gin.Context) {
	ctx := s.requestContext(c)
	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil || userAddress == "" {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
		return
	}
	responseId := s.stringFromContextQuery(c, "id")
	resp, err := s.nls.GetMissionStoreResult(ctx, userAddress, responseId)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxSTRING(c, http.StatusOK, resp)
}
