package apis

import (
	"errors"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/pkg/utils"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/gin-gonic/gin"
	openai2 "github.com/sashabaranov/go-openai"
	"io"
	"net/http"
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

func (s *Server) listKnowledgeByAgent(c *gin.Context) {
	ctx := s.requestContext(c)
	ms, err := s.nls.SyncAgentInfoDetailByAgentID(ctx, s.stringFromContextParam(c, "id"))
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	data, err := s.nls.KnowledgeUsecase.MapKnowledgeBaseByAgentIds(ctx, []uint{ms.ID})
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	resp := []*serializers.KnowledgeBase{}
	if kbs, ok := data[ms.ID]; ok {
		for _, kb := range kbs {
			kbR := &serializers.KnowledgeBase{}
			if err := utils.Copy(kbR, kb); err != nil {
				ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
				return
			}
			resp = append(resp, kbR)
		}
	}

	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: resp})
	return
}

func (s *Server) listKnowledge(c *gin.Context) {
	ctx := s.requestContext(c)
	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: err})
		return
	}

	req := &models.ListKnowledgeBaseRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	req.UserAddress = userAddress

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

func (s *Server) AgentUseKnowledgeBase(c *gin.Context) {
	ctx := s.requestContext(c)
	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: err})
		return
	}

	req := &serializers.AgentUseKnowledgeBaseRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	req.UserAddress = userAddress

	resp, err := s.nls.AgentUseKnowledgeBase(ctx, req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: resp})
}

func (s *Server) updateKnowledgeBaseInContractWithSignature(c *gin.Context) {
	ctx := s.requestContext(c)
	userAddress, err := s.getUserAddressFromTK1Token(c)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: err})
		return
	}

	req := &serializers.UpdateKnowledgeBaseWithSignatureRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	info, err := s.nls.KnowledgeUsecase.GetKnowledgeBaseByKBTokenId(ctx, req.KnowledgeBaseId)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: err})
		return
	}

	if info.UserAddress != userAddress {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errors.New("you are not owner")})
		return
	}

	info, err = s.nls.UpdateKnowledgeBaseInContractWithSignature(ctx, info, req)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: err})
		return
	}

	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: info})
}

func (s *Server) retrieveKnowledge(c *gin.Context) {
	req := &serializers.RetrieveKnowledgeRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}
	if req.Prompt == "" || req.KbId == "" {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errors.New("Prompt and KbId are required")})
		return
	}

	chatTopK := 5
	if s.conf.KnowledgeBaseConfig.KbChatTopK > 0 {
		chatTopK = s.conf.KnowledgeBaseConfig.KbChatTopK
	}
	if req.TopK > 0 {
		chatTopK = req.TopK
	}
	var threshold *float64
	if req.Threshold > 0 {
		threshold = &req.Threshold
	}

	resp, err := s.nls.RetrieveKnowledge("", []openai2.ChatCompletionMessage{{
		Content: req.Prompt,
		Role:    openai2.ChatMessageRoleUser,
	}}, []*models.KnowledgeBase{{
		KbId: req.KbId,
	}}, &chatTopK, threshold)
	if err != nil {
		ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errs.NewError(err)})
		return
	}

	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: resp})
}

func (s *Server) checkBalance(c *gin.Context) {
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
	if resp.Status == models.KnowledgeBaseStatusWaitingPayment {
		if err := s.nls.KnowledgeUsecase.CheckBalance(ctx, resp); err != nil {
			ctxAbortWithStatusJSON(c, http.StatusBadRequest, &serializers.Resp{Error: errors.New("You not owner")})
			return
		}
	}
	ctxJSON(c, http.StatusOK, &serializers.Resp{Result: resp})
}
