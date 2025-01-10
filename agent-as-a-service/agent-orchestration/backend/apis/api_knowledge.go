package apis

import (
	"errors"
	"io"
	"net/http"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/pkg/utils"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/gin-gonic/gin"
)

func (s *Server) createKnowledge(c *gin.Context) {
	ctx := s.requestContext(c)
	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	req := &serializers.CreateKnowledgeRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	req.UserAddress = userAddress

	resp, err := s.nls.KnowledgeUsecase.CreateKnowledgeBase(ctx, req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: resp})
}

func (s *Server) webhookKnowledge(c *gin.Context) {
	ctx := s.requestContext(c)

	req := &models.RagResponse{}
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	_, err := s.nls.KnowledgeUsecase.Webhook(ctx, req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: "success"})
}

func (s *Server) webhookKnowledgeFile(c *gin.Context) {
	ctx := s.requestContext(c)

	fileHeader, err := c.FormFile("file")
	file, err := fileHeader.Open()
	if err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	id := s.uintFromContextParam(c, "id")
	_, err = s.nls.KnowledgeUsecase.WebhookFile(ctx, fileHeader.Filename, fileBytes, id)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: "success"})
}

func (s *Server) listKnowledge(c *gin.Context) {
	ctx := s.requestContext(c)
	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: err})
		return
	}

	req := &models.ListKnowledgeBaseRequest{
		UserAddress: userAddress,
	}

	resp, err := s.nls.KnowledgeUsecase.ListKnowledgeBase(ctx, req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: err})
		return
	}

	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: resp})
}

func (s *Server) deleteKnowledge(c *gin.Context) {
	ctx := s.requestContext(c)
	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: err})
		return
	}

	id := s.uintFromContextParam(c, "id")
	resp, err := s.nls.KnowledgeUsecase.GetKnowledgeBaseById(ctx, id)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: err})
		return
	}

	if resp.UserAddress != userAddress {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errors.New("You not owner")})
		return
	}

	if err := s.nls.KnowledgeUsecase.DeleteKnowledgeBaseById(ctx, id); err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: err})
		return
	}

	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: resp.ID})
}

func (s *Server) detailKnowledge(c *gin.Context) {
	ctx := s.requestContext(c)
	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: err})
		return
	}

	id := s.uintFromContextParam(c, "id")
	resp, err := s.nls.KnowledgeUsecase.GetKnowledgeBaseById(ctx, id)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: err})
		return
	}

	if resp.UserAddress != userAddress {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errors.New("You not owner")})
		return
	}

	r := &serializers.KnowledgeBase{}
	if err := utils.Copy(r, resp); err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: err})
		return
	}

	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: r})
}

func (s *Server) updateKnowledge(c *gin.Context) {
	ctx := s.requestContext(c)
	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: err})
		return
	}

	id := s.uintFromContextParam(c, "id")
	resp, err := s.nls.KnowledgeUsecase.GetKnowledgeBaseById(ctx, id)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: err})
		return
	}

	if resp.UserAddress != userAddress {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errors.New("You not owner")})
		return
	}

	req := &serializers.UpdateKnowledgeRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	updateMap := make(map[string]interface{})
	if req.Name != "" {
		updateMap["name"] = req.Name
	}

	if req.Description != "" {
		updateMap["description"] = req.Description
	}

	if req.NetworkID != 0 {
		updateMap["network_id"] = req.NetworkID
	}

	if err := s.nls.KnowledgeUsecase.UpdateKnowledgeBaseById(c, resp.ID, updateMap); err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: err})
		return
	}

	resp, _ = s.nls.KnowledgeUsecase.GetKnowledgeBaseById(ctx, id)
	r := &serializers.KnowledgeBase{}
	if err := utils.Copy(r, resp); err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: err})
		return
	}

	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: r})
}
