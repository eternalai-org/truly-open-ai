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
	CreateList(ctx context.Context, models []*models.AgentInfoKnowledgeBase, agentInfoId uint) ([]*models.AgentInfoKnowledgeBase, error)
	ListByAgentIds(ctx context.Context, ids []uint) ([]*models.AgentInfoKnowledgeBase, error)
	GetByAgentId(ctx context.Context, id uint) (*models.AgentInfoKnowledgeBase, error)
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

func (r *agentInfoKnowledgeBaseRepo) CreateList(ctx context.Context, addModels []*models.AgentInfoKnowledgeBase, agentInfoId uint) ([]*models.AgentInfoKnowledgeBase, error) {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		removeModels := []*models.AgentInfoKnowledgeBase{}
		if err := tx.Where("agent_info_id = ?", agentInfoId).Delete(&removeModels).Error; err != nil {
			return err
		}

		for _, m := range addModels {
			result := tx.Create(m)
			if result.Error != nil {
				return result.Error
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return addModels, nil
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
