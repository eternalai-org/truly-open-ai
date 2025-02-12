package apis

import (
	"net/http"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/gin-gonic/gin"
)

func (s *Server) InfraTwitterAppAuthenInstall(c *gin.Context) {
	ctx := s.requestContext(c)
	authUrl, err := s.nls.InfraTwitterAppAuthenInstall(
		ctx,
		s.stringFromContextQuery(c, "address"),
		s.stringFromContextQuery(c, "install_uri"),
		s.stringFromContextQuery(c, "install_code"),
		s.stringFromContextQuery(c, "signature"),
	)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	c.Redirect(http.StatusFound, authUrl)
}

func (s *Server) InfraTwitterAppAuthenCallback(c *gin.Context) {
	ctx := s.requestContext(c)
	returnUri, err := s.nls.InfraTwitterAppAuthenCallback(ctx, s.stringFromContextQuery(c, "install_uri"), s.stringFromContextQuery(c, "install_code"), s.stringFromContextQuery(c, "code"))
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	c.Redirect(http.StatusFound, returnUri)
}
