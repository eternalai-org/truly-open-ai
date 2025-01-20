package repository

import (
	"context"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/models"
	"gorm.io/gorm"
)

type agentInfoRepo struct {
	db *gorm.DB
}

type IAgentInfoRepo interface {
	UpdateById(ctx context.Context, id uint, updatedFields map[string]interface{}) error
}

func (r *agentInfoRepo) UpdateById(ctx context.Context, id uint, updatedFields map[string]interface{}) error {
	result := r.db.WithContext(ctx).Model(&models.AgentInfo{}).Where("id = ?", id).Updates(updatedFields)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewAgentInfoRepository(db *gorm.DB) IAgentInfoRepo {
	return &agentInfoRepo{db}
}
