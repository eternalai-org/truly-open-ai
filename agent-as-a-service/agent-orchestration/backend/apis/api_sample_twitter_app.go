package apis

import (
	"net/http"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/gin-gonic/gin"
)

func (s *Server) SampleTwitterAppAuthenInstall(c *gin.Context) {
	ctx := s.requestContext(c)
	authUrl, err := s.nls.SampleTwitterAppAuthenInstall(ctx, s.stringFromContextQuery(c, "install_code"), s.stringFromContextQuery(c, "install_uri"))
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	c.Redirect(http.StatusFound, authUrl)
}

func (s *Server) SampleTwitterAppAuthenCallback(c *gin.Context) {
	ctx := s.requestContext(c)
	returnUri, err := s.nls.SampleTwitterAppAuthenCallback(ctx, s.stringFromContextQuery(c, "install_code"), s.stringFromContextQuery(c, "install_uri"), s.stringFromContextQuery(c, "code"))
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	c.Redirect(http.StatusFound, returnUri)
}

func (s *Server) SampleTwitterAppGetBTCPrice(c *gin.Context) {
	ctx := s.requestContext(c)
	btcPrice := s.nls.SampleTwitterAppGetBTCPrice(ctx)
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: btcPrice})
}

func (s *Server) SampleTwitterAppTweetMessage(c *gin.Context) {
	ctx := s.requestContext(c)
	var req struct {
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	tweetId, err := s.nls.SampleTwitterAppTweetMessage(ctx, c.GetHeader("api-key"), req.Content)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: tweetId})
}
