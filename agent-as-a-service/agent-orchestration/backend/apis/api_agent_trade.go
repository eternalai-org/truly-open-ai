package apis

import (
	"net/http"
	"strings"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetMemeConfigs(c *gin.Context) {
	ctx := s.requestContext(c)
	configs, err := s.nls.GetAllConfigs(ctx)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	fractalConfigs := configs[models.BASE_CHAIN_ID]
	res := map[string]string{}
	for k, v := range fractalConfigs {
		if k == "meme_token_fee_address" ||
			k == "meme_token_fee_admin_address" ||
			k == "btc_contract_addr" ||
			k == "meme_mint_whitelist" {
			res[k] = v
		}
	}

	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: res})
}

func (s *Server) GetListMemeByAddress(c *gin.Context) {
	ctx := s.requestContext(c)
	address := s.stringFromContextQuery(c, "address")
	page, limit := s.pagingFromContext(c)
	txs, count, err := s.nls.GetListMemes(ctx,
		address, page, limit)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewMemeRespArray(txs), Count: &count})
}

func (s *Server) GetListMemeReport(c *gin.Context) {
	ctx := s.requestContext(c)
	address := s.stringFromContextQuery(c, "address")
	page, limit := s.pagingFromContext(c)
	search := s.stringFromContextQuery(c, "search")
	status := s.stringFromContextQuery(c, "status")
	sortStr := s.sortListFromContext(c)
	sortCol := s.stringFromContextQuery(c, "sort_col")
	following := false
	if strings.Contains(sortCol, "following") {
		following = true
	}
	chain := s.chainFromContextQuery(c)
	ms, count, err := s.nls.GetListMemeReport(ctx, chain, address, search, status, following,
		sortStr, page, limit,
	)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewMemeRespArray(ms), Count: &count})
}

func (s *Server) GetFeedMemeReport(c *gin.Context) {
	ctx := s.requestContext(c)
	address := s.stringFromContextQuery(c, "address")
	page, limit := s.pagingFromContext(c)
	search := s.stringFromContextQuery(c, "search")
	sortType := s.stringFromContextQuery(c, "sort_type")
	ms, count, err := s.nls.GetFeedMemeReport(ctx, address, search, sortType, page, limit)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewMemeRespArray(ms), Count: &count})
}

func (s *Server) GetMemeDetail(c *gin.Context) {
	ctx := s.requestContext(c)
	address := s.stringFromContextQuery(c, "address")
	tokenAddress := s.stringFromContextParam(c, "id")
	resp, err := s.nls.GetMemeDetail(ctx, address, tokenAddress)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	// ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewMemeResp(ms)})
	ctxSTRING(c, http.StatusOK, resp)
}

func (s *Server) SeenMemeDetail(c *gin.Context) {
	ctx := s.requestContext(c)
	address := s.stringFromContextQuery(c, "address")
	tokenAddress := s.stringFromContextParam(c, "id")
	go s.nls.CheckedSeenMeme(ctx, address, tokenAddress)
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: true})
}

func (s *Server) ShareMeme(c *gin.Context) {
	ctx := s.requestContext(c)
	address := s.stringFromContextQuery(c, "address")
	tokenAddress := s.stringFromContextParam(c, "id")
	go s.nls.ShareMeme(ctx, address, tokenAddress)
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: true})
}

func (s *Server) GetMemeCandleChart(c *gin.Context) {
	ctx := s.requestContext(c)
	day := s.uintFromContextQuery(c, "day")
	if day == 0 {
		day = 90
	} else if day > 90 {
		day = 90
	}

	tokenAddress := s.stringFromContextParam(c, "id")
	chartType := s.stringFromContextQuery(c, "type")

	resp, err := s.nls.GetMemeChartCandleData(ctx,
		tokenAddress, day, chartType,
	)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	ctxSTRING(c, http.StatusOK, resp)
}

func (s *Server) GetMemeTradeHistory(c *gin.Context) {
	ctx := s.requestContext(c)
	address := s.stringFromContextQuery(c, "address")
	page, limit := s.pagingFromContext(c)
	tokenAddress := s.stringFromContextQuery(c, "token_address")
	ms, count, err := s.nls.GetMemeTradeHistory(ctx, address, tokenAddress, page, limit)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewTradeHistoryRespArry(ms), Count: &count})
}

func (s *Server) GetMemeTradeHistoryLatest(c *gin.Context) {
	ctx := s.requestContext(c)
	tokenAddress := s.stringFromContextParam(c, "id")
	resp, err := s.nls.GetMemeTradeHistoryLatest(ctx, tokenAddress)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxSTRING(c, http.StatusOK, resp)
}

func (s *Server) GetMemeNotification(c *gin.Context) {
	ctx := s.requestContext(c)
	address := s.stringFromContextQuery(c, "address")
	page, limit := s.pagingFromContext(c)
	ms, count, err := s.nls.GetMemeNotifications(ctx, address, page, limit)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewMemeNotificationRespArry(ms), Count: &count})
}

func (s *Server) GetMemeNotificationLatest(c *gin.Context) {
	ctx := s.requestContext(c)
	address := s.stringFromContextQuery(c, "address")
	resp, err := s.nls.GetMemeNotificationLatest(ctx, address)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxSTRING(c, http.StatusOK, resp)
}

func (s *Server) GetMemeListHolders(c *gin.Context) {
	ctx := s.requestContext(c)
	tokenAddress := s.stringFromContextQuery(c, "token_address")
	resp, err := s.nls.GetTokenHolders(ctx, tokenAddress)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxSTRING(c, http.StatusOK, resp)
}

func (s *Server) GetMemeListHolding(c *gin.Context) {
	ctx := s.requestContext(c)
	address := s.stringFromContextQuery(c, "address")
	page, limit := s.pagingFromContext(c)
	ms, count, err := s.nls.GetTokenHolding(ctx, address, page, limit)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewTokenHolderRespArray(ms), Count: &count})
}

// /thread
// func (s *Server) CreateMemeThread(c *gin.Context) {
// 	ctx := s.requestContext(c)
// 	address := s.stringFromContextQuery(c, "address")
// 	var req serializers.MemeThreadReq
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
// 		return
// 	}

// 	meme, err := s.nls.CreateMemeThread(ctx, address, &req)
// 	if err != nil {
// 		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
// 		return
// 	}
// 	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewMemeThreadResp(meme)})
// }

// func (s *Server) LikeMemeThread(c *gin.Context) {
// 	ctx := s.requestContext(c)
// 	address := s.stringFromContextQuery(c, "address")
// 	var req serializers.MemeThreadReq
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
// 		return
// 	}

// 	resp, err := s.nls.LikeMemeThread(ctx, address, &req)
// 	if err != nil {
// 		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
// 		return
// 	}
// 	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: resp})
// }

// func (s *Server) UnLikeMemeThread(c *gin.Context) {
// 	ctx := s.requestContext(c)
// 	address := s.stringFromContextQuery(c, "address")
// 	var req serializers.MemeThreadReq
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
// 		return
// 	}

// 	resp, err := s.nls.UnLikeMemeThread(ctx, address, &req)
// 	if err != nil {
// 		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
// 		return
// 	}
// 	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: resp})
// }

func (s *Server) GetListMemeThread(c *gin.Context) {
	ctx := s.requestContext(c)
	address := s.stringFromContextQuery(c, "address")
	page, limit := s.pagingFromContext(c)
	tokenAddress := s.stringFromContextParam(c, "id")

	ms, count, err := s.nls.GetListMemeThread(ctx, address, tokenAddress, page, limit)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewMemeThreadRespArry(ms), Count: &count})
}

func (s *Server) GetListMemeThreadLatest(c *gin.Context) {
	ctx := s.requestContext(c)
	tokenAddress := s.stringFromContextParam(c, "id")

	resp, err := s.nls.GetListMemeThreadLatest(ctx, tokenAddress)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxSTRING(c, http.StatusOK, resp)
}

// //// User
func (s *Server) GetMemeUserProfile(c *gin.Context) {
	ctx := s.requestContext(c)
	address := s.stringFromContextQuery(c, "address")
	ms, err := s.nls.GetMemeUserProfile(ctx, 0, address)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewUserResp(ms)})
}

func (s *Server) MemeValidatedFollow(c *gin.Context) {
	ctx := s.requestContext(c)
	address := s.stringFromContextQuery(c, "address")
	userAddress := s.stringFromContextQuery(c, "user_address")
	valid, err := s.nls.ValidatedFollowed(ctx, address, userAddress)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: valid})
}

// func (s *Server) UpdateMemeUserProfile(c *gin.Context) {
// 	ctx := s.requestContext(c)
// 	address := s.stringFromContextQuery(c, "address")
// 	var req serializers.UserProfileReq
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
// 		return
// 	}

// 	user, err := s.nls.UpdateUserProfile(ctx, address, &req)
// 	if err != nil {
// 		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
// 		return
// 	}
// 	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewUserResp(user)})
// }

// /
func (s *Server) FollowUser(c *gin.Context) {
	ctx := s.requestContext(c)
	address := s.stringFromContextQuery(c, "address")
	var req serializers.MemeThreadReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	resp, err := s.nls.FollowUsers(ctx, address, &req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: resp})
}

func (s *Server) UnFollowUser(c *gin.Context) {
	ctx := s.requestContext(c)
	address := s.stringFromContextQuery(c, "address")
	var req serializers.MemeThreadReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	resp, err := s.nls.UnFollowUsers(ctx, address, &req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: resp})
}

func (s *Server) GetListFollowers(c *gin.Context) {
	ctx := s.requestContext(c)
	address := s.stringFromContextQuery(c, "address")
	page, limit := s.pagingFromContext(c)
	ms, count, err := s.nls.GetListFollowers(ctx, address, page, limit)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewMemeFollowersRespArray(ms), Count: &count})
}

func (s *Server) GetListFollowings(c *gin.Context) {
	ctx := s.requestContext(c)
	address := s.stringFromContextQuery(c, "address")
	page, limit := s.pagingFromContext(c)
	ms, count, err := s.nls.GetListFollowings(ctx, address, page, limit)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewMemeFollowersRespArray(ms), Count: &count})
}

// CrossChain
func (s *Server) HideMemeThread(c *gin.Context) {
	ctx := s.requestContext(c)
	address := s.stringFromContextQuery(c, "address")

	threadID := s.uintFromContextParam(c, "id")
	resp, err := s.nls.HideMemeThread(ctx, address, threadID)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: resp})
}

func (s *Server) GetMemeWhiteListAddress(c *gin.Context) {
	ctx := s.requestContext(c)
	resp, err := s.nls.GetMemeWhiteListAddress(ctx)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: resp})
}

func (s *Server) GetMemeBurnHistory(c *gin.Context) {
	ctx := s.requestContext(c)
	address := s.stringFromContextQuery(c, "address")
	page, limit := s.pagingFromContext(c)
	tokenAddress := s.stringFromContextQuery(c, "token_address")
	ms, count, err := s.nls.GetMemeBurnHistory(ctx, 0, address, tokenAddress, page, limit)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewMemeBurnHistoryRespArry(ms), Count: &count})
}

func (s *Server) UserSeenMemeNotification(c *gin.Context) {
	ctx := s.requestContext(c)
	address := s.stringFromContextQuery(c, "address")

	notiID := s.uintFromContextParam(c, "id")
	resp, err := s.nls.UserSeenMemeNotification(ctx, address, notiID)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: resp})
}

func (s *Server) GenerateMemeStory(c *gin.Context) {
	ctx := s.requestContext(c)
	name := s.stringFromContextQuery(c, "name")
	story, err := s.nls.GenerateMemeStory(ctx, name)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: story})
}
