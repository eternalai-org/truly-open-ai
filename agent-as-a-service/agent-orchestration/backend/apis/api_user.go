package apis

import (
	"net/http"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/gin-gonic/gin"
)

func (s *Server) VerifyLoginUserByWeb3(c *gin.Context) {
	ctx := s.requestContext(c)
	var req struct {
		Address   string `json:"address"`
		Message   string `json:"message"`
		Signature string `json:"signature"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	tkAuth, err := s.nls.VerifyLoginUserByWeb3(ctx, s.getRequestIP(c), s.getUserAgent(c), req.Address, req.Message, req.Signature)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Data: tkAuth})
}

func (s *Server) GetUserProfileWithAuth(c *gin.Context) {
	ctx := s.requestContext(c)
	userAddress, err := s.getUserAddressFromTK1Token(c)
	user, err := s.nls.GetUserProfile(ctx, userAddress)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: serializers.NewUserResp(user)})
}

func (s *Server) UserUploadFile(c *gin.Context) {
	ctx := s.requestContext(c)
	userAddress, err := s.getUserAddressFromTK1Token(c)
	fileHeader, err := c.FormFile("file")
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	path, err := s.nls.UserUploadFile(ctx, userAddress, fileHeader)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: path})
}

func (s *Server) AgentRequestTwitterShareCode(c *gin.Context) {
	ctx := s.requestContext(c)
	topupAddress := s.stringFromContextQuery(c, "address")
	authSecretCode, authPublicCode, err := s.nls.AgentRequestTwitterShareCode(
		ctx, topupAddress,
	)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c,
		http.StatusOK,
		&serializers.Resp{
			Result: gin.H{
				"public_code": authPublicCode,
				"secret_code": authSecretCode,
			},
		},
	)
}

func (s *Server) AgentVerifyShareTwitter(c *gin.Context) {
	ctx := s.requestContext(c)
	var req struct {
		SecretCode string `json:"secret_code"`
		Link       string `json:"link"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	valid, err := s.nls.AgentVerifyShareTwitter(ctx, req.SecretCode, req.Link)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: valid})
}
