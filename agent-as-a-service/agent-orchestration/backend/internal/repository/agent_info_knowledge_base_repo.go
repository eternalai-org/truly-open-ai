package repository

import (
	"context"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	gorm1 "github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

type agentInfoKnowledgeBaseRepo struct {
	db *gorm.DB
}

type IAgentInfoKnowledgeBaseRepo interface {
	Create(ctx context.Context, model *models.AgentInfoKnowledgeBase) (*models.AgentInfoKnowledgeBase, error)
	ListByAgentIds(ctx context.Context, ids []uint) ([]*models.AgentInfoKnowledgeBase, error)
	GetByAgentId(ctx context.Context, id uint) (*models.AgentInfoKnowledgeBase, error)
	GetKBAgentsUsedOfSocialAgent(tx *gorm1.DB, socialAgentId uint) ([]*models.KnowledgeBase, error)
}

func (r *agentInfoKnowledgeBaseRepo) GetByAgentId(ctx context.Context, id uint) (*models.AgentInfoKnowledgeBase, error) {
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

func (r *agentInfoKnowledgeBaseRepo) GetKBAgentsUsedOfSocialAgent(tx *gorm1.DB, socialAgentId uint) ([]*models.KnowledgeBase, error) {
	knowledge := []*models.KnowledgeBase{}
	err := tx.
		Preload("AgentInfo").
		Where("kb_id <> '' and id IN (SELECT knowledge_base_id FROM agent_info_knowledge_bases WHERE agent_info_id = ?)",
			socialAgentId).
		Find(&knowledge).Error
	if err != nil {
		return nil, err
	}

	return knowledge, nil
}

func (r *agentInfoKnowledgeBaseRepo) ListByAgentIds(ctx context.Context, ids []uint) ([]*models.AgentInfoKnowledgeBase, error) {
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
