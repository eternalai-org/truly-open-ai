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
	ListAgentInfoKnowledgeBaseByAgentIds(ctx context.Context, ids []uint) ([]*models.AgentInfoKnowledgeBase, error)
	GetAgentInfoKnowledgeBaseByAgentId(ctx context.Context, id uint) (*models.AgentInfoKnowledgeBase, error)
}

func (r *agentInfoKnowledgeBaseRepo) GetAgentInfoKnowledgeBaseByAgentId(ctx context.Context, id uint) (*models.AgentInfoKnowledgeBase, error) {
	knowledge := &models.AgentInfoKnowledgeBase{}
	err := r.db.WithContext(ctx).
		Preload("KnowledgeBase").
		Preload("AgentInfo").
		Preload("KnowledgeBase.KnowledgeBaseFiles").
		Where("agent_info_id = ?", id).
		First(knowledge).Error
	if err != nil {
		return nil, err
	}
	return knowledge, nil
}

func (r *agentInfoKnowledgeBaseRepo) ListAgentInfoKnowledgeBaseByAgentIds(ctx context.Context, ids []uint) ([]*models.AgentInfoKnowledgeBase, error) {
	resp := []*models.AgentInfoKnowledgeBase{}
	err := r.db.WithContext(ctx).
		Preload("KnowledgeBase").
		Preload("KnowledgeBase.KnowledgeBaseFiles").
		Where("agent_info_id IN (?)", ids).
		Find(&resp).Error
	if err != nil {
		return nil, err
	}
	return resp, nil
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
