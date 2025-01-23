package services

import (
	"context"
	"errors"
	"strings"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
)

func (s *Service) AgentUseKnowledgeBase(ctx context.Context, request *serializers.AgentUseKnowledgeBaseRequest) (*models.AgentInfoKnowledgeBase, error) {
	// currently only use 1 knowledge base for 1 agent, so insert will by multiple, but when use only 1 knowledge base order by id desc
	agentInfo, err := s.dao.FirstAgentInfo(daos.GetDBMainCtx(ctx), map[string][]interface{}{
		"agent_id = ?": {request.AgentID},
	}, nil, []string{"id desc"})
	if err != nil {
		return nil, err
	}
	if agentInfo == nil {
		return nil, errors.New("agent not found")
	}
	if !strings.EqualFold(request.UserAddress, agentInfo.Creator) {
		return nil, errors.New("only agent owner can use knowledge base")
	}

	// check knowledge base exist
	knowledgeBase, err := s.dao.FirstKnowledgeBase(daos.GetDBMainCtx(ctx), map[string][]interface{}{
		"id = ?": {request.KnowledgeBaseID},
	}, nil, []string{"id desc"}, false)
	if err != nil {
		return nil, err
	}
	if knowledgeBase == nil {
		return nil, errors.New("knowledge base not found")
	}

	if knowledgeBase.Status != models.KnowledgeBaseStatusMinted {
		return nil, errors.New("knowledge base is not minted")
	}

	agentInfoKnowledgeBase, err := s.dao.FirstAgentInfoKnowledgeBaseByAgentInfoIDAndKnowledgeBaseID(daos.GetDBMainCtx(ctx), agentInfo.ID, request.KnowledgeBaseID)
	if err != nil {
		return nil, err
	}
	if agentInfoKnowledgeBase != nil {
		return agentInfoKnowledgeBase, nil
	}

	newKnowledgeBase := &models.AgentInfoKnowledgeBase{
		AgentInfoId:     agentInfo.ID,
		KnowledgeBaseId: request.KnowledgeBaseID,
	}
	if err := s.dao.CreateAgentInfoKnowledgeBase(daos.GetDBMainCtx(ctx), newKnowledgeBase); err != nil {
		return nil, err
	}

	return newKnowledgeBase, nil
}
