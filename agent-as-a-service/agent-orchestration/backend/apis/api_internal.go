package apis

import (
	"net/http"
	"strconv"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetTwitterUserByID(c *gin.Context) {
	ctx := s.requestContext(c)
	twitterID := s.stringFromContextParam(c, "id")
	user, err := s.nls.GetTwitterUserByID(ctx, twitterID)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) GetTwitterUserByIDByQuery(c *gin.Context) {
	ctx := s.requestContext(c)
	twitterID := s.stringFromContextQuery(c, "id")
	user, err := s.nls.GetTwitterUserByID(ctx, twitterID)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) GetTwitterUserByUserName(c *gin.Context) {
	ctx := s.requestContext(c)
	username := s.stringFromContextParam(c, "username")
	user, err := s.nls.GetTwitterUserByUsername(ctx, username)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) GetTwitterUserByUserNameByQuery(c *gin.Context) {
	ctx := s.requestContext(c)
	username := s.stringFromContextQuery(c, "username")
	user, err := s.nls.GetTwitterUserByUsername(ctx, username)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) GetTwitterUserByQuery(c *gin.Context) {
	ctx := s.requestContext(c)
	usernames := s.stringFromContextQuery(c, "usernames")
	user, err := s.nls.SeachTwitterUserByQuery(ctx, usernames)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) GetTwitterUserFollowing(c *gin.Context) {
	ctx := s.requestContext(c)
	twitterID := s.stringFromContextParam(c, "id")
	paginationToken := s.stringFromContextQuery(c, "pagination_token")
	user, err := s.nls.GetTwitterUserFollowingRapid(ctx, twitterID, paginationToken)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) GetTwitterUserFollowingByQuery(c *gin.Context) {
	ctx := s.requestContext(c)
	twitterID := s.stringFromContextQuery(c, "id")
	paginationToken := s.stringFromContextQuery(c, "pagination_token")
	user, err := s.nls.GetTwitterUserFollowingRapid(ctx, twitterID, paginationToken)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) GetTwitterUserFollowingByUsername(c *gin.Context) {
	ctx := s.requestContext(c)
	username := s.stringFromContextParam(c, "username")
	paginationToken := s.stringFromContextQuery(c, "pagination_token")
	user, err := s.nls.GetTwitterUserFollowingRapidByUsername(ctx, username, paginationToken)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) GetTwitterUserFollowingByUsernameByQuery(c *gin.Context) {
	ctx := s.requestContext(c)
	username := s.stringFromContextQuery(c, "username")
	paginationToken := s.stringFromContextQuery(c, "pagination_token")
	user, err := s.nls.GetTwitterUserFollowingRapidByUsername(ctx, username, paginationToken)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) GetListUserTweets(c *gin.Context) {
	ctx := s.requestContext(c)
	twitterID := s.stringFromContextParam(c, "id")
	paginationToken := s.stringFromContextQuery(c, "pagination_token")
	maxResults := s.maxResultFromContextQuery(c)
	user, err := s.nls.GetListUserTweets(ctx, twitterID, paginationToken, maxResults)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) GetListUserTweetsAll(c *gin.Context) {
	ctx := s.requestContext(c)
	twitterID := s.stringFromContextParam(c, "id")
	paginationToken := s.stringFromContextQuery(c, "pagination_token")
	maxResults := s.maxResultFromContextQuery(c)
	user, err := s.nls.GetListUserTweetsAll(ctx, twitterID, paginationToken, maxResults)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) GetListUserTweetsByQuery(c *gin.Context) {
	ctx := s.requestContext(c)
	twitterID := s.stringFromContextQuery(c, "id")
	paginationToken := s.stringFromContextQuery(c, "pagination_token")
	maxResults := s.maxResultFromContextQuery(c)
	user, err := s.nls.GetListUserTweetsV1(ctx, twitterID, paginationToken, maxResults)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) GetListUserTweetsByUserName(c *gin.Context) {
	ctx := s.requestContext(c)
	username := s.stringFromContextParam(c, "username")
	paginationToken := s.stringFromContextQuery(c, "pagination_token")
	maxResults := s.maxResultFromContextQuery(c)
	user, err := s.nls.GetListUserTweetsByUsername(ctx, username, paginationToken, maxResults)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) GetListUserTweetsByUserNameV1(c *gin.Context) {
	ctx := s.requestContext(c)
	username := s.stringFromContextParam(c, "username")
	paginationToken := s.stringFromContextQuery(c, "pagination_token")
	maxResults := s.maxResultFromContextQuery(c)
	user, err := s.nls.GetListUserTweetsByUsernameV1(ctx, username, paginationToken, maxResults)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) GetListUserTweetsByUserNameByQuery(c *gin.Context) {
	ctx := s.requestContext(c)
	username := s.stringFromContextQuery(c, "username")
	paginationToken := s.stringFromContextQuery(c, "pagination_token")
	maxResults := s.maxResultFromContextQuery(c)
	user, err := s.nls.GetListUserTweetsByUsernameV1(ctx, username, paginationToken, maxResults)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) LookupUserTweets(c *gin.Context) {
	ctx := s.requestContext(c)
	ids := s.stringFromContextQuery(c, "ids")
	user, err := s.nls.LookupUserTweets(ctx, ids)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) LookupUserTweetsV1(c *gin.Context) {
	ctx := s.requestContext(c)
	ids := s.stringFromContextQuery(c, "ids")
	user, err := s.nls.LookupUserTweetsV1(ctx, ids)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) GetListUserMentions(c *gin.Context) {
	ctx := s.requestContext(c)
	twitterID := s.stringFromContextParam(c, "id")
	paginationToken := s.stringFromContextQuery(c, "pagination_token")
	maxResults := s.maxResultFromContextQuery(c)
	user, err := s.nls.GetListUserMentions(ctx, twitterID, paginationToken, maxResults)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) GetListUserMentionsByQuery(c *gin.Context) {
	ctx := s.requestContext(c)
	twitterID := s.stringFromContextQuery(c, "id")
	paginationToken := s.stringFromContextQuery(c, "pagination_token")
	maxResults := s.maxResultFromContextQuery(c)
	user, err := s.nls.GetListUserMentions(ctx, twitterID, paginationToken, maxResults)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) GetListUserMentionsByUsername(c *gin.Context) {
	ctx := s.requestContext(c)
	username := s.stringFromContextParam(c, "username")
	paginationToken := s.stringFromContextQuery(c, "pagination_token")
	maxResults := s.maxResultFromContextQuery(c)
	user, err := s.nls.GetListUserMentionsByUsername(ctx, username, paginationToken, maxResults)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) GetAllUserMentionsByUsername(c *gin.Context) {
	ctx := s.requestContext(c)
	username := s.stringFromContextParam(c, "username")
	paginationToken := s.stringFromContextQuery(c, "pagination_token")
	maxResults := s.maxResultFromContextQuery(c)
	user, err := s.nls.GetAllUserMentionsByUsername(ctx, username, paginationToken, maxResults)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) GetListUserMentionsByUsernameByQuery(c *gin.Context) {
	ctx := s.requestContext(c)
	username := s.stringFromContextQuery(c, "username")
	paginationToken := s.stringFromContextQuery(c, "pagination_token")
	maxResults := s.maxResultFromContextQuery(c)
	user, err := s.nls.GetListUserMentionsByUsername(ctx, username, paginationToken, maxResults)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) SearchRecentTweet(c *gin.Context) {
	ctx := s.requestContext(c)
	query := s.stringFromContextQuery(c, "query")
	paginationToken := s.stringFromContextQuery(c, "pagination_token")
	maxResults := s.maxResultFromContextQuery(c)
	user, err := s.nls.SearchRecentTweet(ctx, query, paginationToken, maxResults)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) SearchTokenTweet(c *gin.Context) {
	ctx := s.requestContext(c)
	query := s.stringFromContextQuery(c, "query")
	paginationToken := s.stringFromContextQuery(c, "pagination_token")
	maxResults := s.maxResultFromContextQuery(c)
	user, err := s.nls.SearchTokenTweet(ctx, query, paginationToken, maxResults)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) SearchUsers(c *gin.Context) {
	ctx := s.requestContext(c)
	query := s.stringFromContextQuery(c, "query")
	paginationToken := s.stringFromContextQuery(c, "pagination_token")
	user, err := s.nls.SearchUsers(ctx, query, paginationToken)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) CreateAgentInternalAction(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.AdminAgentActionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	err := s.nls.CreateAgentInternalAction(ctx, &req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: true})
}

func (s *Server) CreateAgentInternalActionByRefID(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.AdminAgentActionByRefReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	refId := s.stringFromContextQuery(c, "ref_id")
	err := s.nls.CreateAgentInternalActionByRefID(ctx, refId, &req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: true})
}

func (s *Server) AgentWalletCreatePumpFunMeme(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.AdminCreatePumpfunMemeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	err := s.nls.AgentWalletCreatePumpFunMeme(ctx, uint64(s.uintFromContextParam(c, "chain_id")), s.stringFromContextParam(c, "agent_contract_id"), &req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: true})
}

func (s *Server) AgentWalletTradePumpFunMeme(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.AdminTradePumpfunMemeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	err := s.nls.AgentWalletTradePumpFunMeme(ctx, uint64(s.uintFromContextParam(c, "chain_id")), s.stringFromContextParam(c, "agent_contract_id"), &req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: true})
}

func (s *Server) AgentWalletTradeRaydiumToken(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.AdminTradePumpfunMemeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	err := s.nls.AgentWalletTradeRaydiumToken(ctx, uint64(s.uintFromContextParam(c, "chain_id")), s.stringFromContextParam(c, "agent_contract_id"), &req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: true})
}

func (s *Server) AgentWalletGetSolanaTokenBalances(c *gin.Context) {
	ctx := s.requestContext(c)
	rs, err := s.nls.AgentWalletGetSolanaTokenBalances(ctx, uint64(s.uintFromContextParam(c, "chain_id")), s.stringFromContextParam(c, "agent_contract_id"))
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: rs})
}

func (s *Server) GetAgentWalletSolanaTrades(c *gin.Context) {
	ctx := s.requestContext(c)
	page, limit := s.pagingFromContext(c)
	rs, err := s.nls.GetAgentWalletSolanaTrades(
		ctx,
		uint64(s.uintFromContextParam(c, "chain_id")),
		s.stringFromContextParam(c, "agent_contract_id"),
		s.stringFromContextQuery(c, "mint"),
		page, limit,
	)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: rs})
}

func (s *Server) AgentWalletGetSolanaTokenPnls(c *gin.Context) {
	ctx := s.requestContext(c)
	rs, err := s.nls.AgentWalletGetSolanaTokenPnls(
		ctx,
		uint64(s.uintFromContextParam(c, "chain_id")),
		s.stringFromContextParam(c, "agent_contract_id"),
	)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: rs})
}

func (s *Server) GetPumpFunTrades(c *gin.Context) {
	ctx := s.requestContext(c)
	page, limit := s.pagingFromContext(c)
	rs, err := s.nls.GetPumpFunTrades(ctx, s.stringFromContextParam(c, "mint"), page, limit)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: rs})
}

func (s *Server) GetPumpFunTokenPrice(c *gin.Context) {
	ctx := s.requestContext(c)
	rs, err := s.nls.GetPumpFunTokenPrice(ctx, s.stringFromContextParam(c, "mint"))
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: rs})
}

func (s *Server) DexSearchPair(c *gin.Context) {
	ctx := s.requestContext(c)
	rs, err := s.nls.DexSpotPairsLatest(ctx, s.stringFromContextQuery(c, "symbol"), s.stringFromContextQuery(c, "network"))
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: rs})
}

func (s *Server) DexPairsTradeLatest(c *gin.Context) {
	ctx := s.requestContext(c)
	rs, err := s.nls.DexPairsTradeLatest(ctx, s.stringFromContextQuery(c, "contract_address"), s.stringFromContextQuery(c, "network"))
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: rs})
}

func (s *Server) DexScreenInfo(c *gin.Context) {
	ctx := s.requestContext(c)
	rs, err := s.nls.DexScreenInfo(ctx, s.stringFromContextQuery(c, "contract_address"))
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: rs})
}

func (s *Server) GetUser3700Liked(c *gin.Context) {
	ctx := s.requestContext(c)
	page, limit := s.pagingFromContext(c)
	replied, _ := s.boolFromContextQuery(c, "replied")
	resp, err := s.nls.GetUser3700Liked(ctx, replied, page, limit)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewTwitterTweetLikedRespArr(resp)})
}

func (s *Server) GetSolanaDataChart24Hour(c *gin.Context) {
	ctx := s.requestContext(c)
	rs, err := s.nls.GetSolanaDataChart24Hour(ctx, s.stringFromContextParam(c, "mint"))
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: rs})
}

func (s *Server) GetAgentTradeTokens(c *gin.Context) {
	ctx := s.requestContext(c)
	chainID := s.chainFromContextQuery(c)
	ms, err := s.nls.GetAgentTradeTokens(ctx, chainID)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewAgentTradeTokenRespArry(ms)})
}

// ///////
func (s *Server) GetListUserTweetsByUsersForTradeMission(c *gin.Context) {
	ctx := s.requestContext(c)
	twitterIDs := s.stringFromContextQuery(c, "ids")
	user, err := s.nls.GetListUserTweetsByUsersForTradeMission(ctx, twitterIDs)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) GetListUserTweetsByAgentForTradeMission(c *gin.Context) {
	ctx := s.requestContext(c)
	agentID := s.stringFromContextQuery(c, "agent_id")
	user, err := s.nls.GetListUserTweetsByAgentForTradeMission(ctx, agentID)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: user})
}

func (s *Server) GetTokenQuoteLatestForSolana(c *gin.Context) {
	ctx := s.requestContext(c)
	rs, err := s.nls.GetTokenQuoteLatestForSolana(ctx, s.stringFromContextQuery(c, "mint"))
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: rs})
}

func (s *Server) TweetByToken(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.AdminTweetReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	agentInfoIDstr := c.GetHeader("agent-info-id")
	agentInfoID, _ := strconv.ParseUint(agentInfoIDstr, 10, 64)
	err := s.nls.TweetByToken(ctx, uint(agentInfoID), &req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: true})
}

func (s *Server) GetTradeAnalytic(c *gin.Context) {
	ctx := s.requestContext(c)
	rs, err := s.nls.GetTradeAnalytic(ctx, s.stringFromContextQuery(c, "token"))
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: rs})
}

func (s *Server) GetTwitterDataForLaunchpad(c *gin.Context) {
	ctx := s.requestContext(c)
	twitterID := s.stringFromContextQuery(c, "id")
	user, err := s.nls.GetTwitterUserByID(ctx, twitterID)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	tweets, err := s.nls.GetListUserTweetsByUsersForTradeMission(ctx, twitterID)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	if user == nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrBadRequest)})
		return
	}
	resp := map[string]interface{}{
		"followers_count": user.PublicMetrics.Followers,
		"following_count": user.PublicMetrics.Following,
		"tweet_count":     user.PublicMetrics.Tweets,
		"listed_count":    user.PublicMetrics.Listed,
		"blue-checked":    user.Verified,
		"recent-tweets":   tweets,
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: resp})
}
