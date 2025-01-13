package repository

import (
	"context"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"gorm.io/gorm"
)

type agentInfoKnowledgeBaseRepo struct {
	db *gorm.DB
}

type IAgentInfoKnowledgeBaseRepo interface {
	Create(ctx context.Context, model *models.AgentInfoKnowledgeBase) (*models.AgentInfoKnowledgeBase, error)
}

func (r *agentInfoKnowledgeBaseRepo) Create(ctx context.Context, model *models.AgentInfoKnowledgeBase) (*models.AgentInfoKnowledgeBase, error) {
	result := r.db.WithContext(ctx).Create(model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}

func NewAgentInfoKnowledgeBaseRepository(db *gorm.DB) IAgentInfoKnowledgeBaseRepo {
	return &agentInfoKnowledgeBaseRepo{db}
}
