package apis

import (
	"net/http"
	"strconv"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/gin-gonic/gin"
)

func (s *Server) AsyncBatchPrompt(c *gin.Context) {
	ctx := s.requestContext(c)
	var req serializers.ChatCompletionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil || userAddress == "" {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
		return
	}

	response, err := s.nls.BatchChatCompletionPrompt(ctx, &req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrUnAuthorization)})
		return
	}

	ctxJSON(c, http.StatusOK, &serializers.APIResponse{
		Status:  1,
		Message: "OK",
		Data:    response,
	})
}
func (s *Server) GetBatchItem(c *gin.Context) {
	ctx := s.requestContext(c)
	idStr := s.stringFromContextParam(c, "id")
	id, _ := strconv.ParseUint(idStr, 10, 64)

	resp, err := s.nls.GetBatchItemDetail(ctx, uint(id))
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(errs.ErrBadRequest)})
		return
	}

	ctxJSON(c, http.StatusOK, &serializers.APIResponse{
		Status:  1,
		Message: "OK",
		Data:    resp,
	})
}
