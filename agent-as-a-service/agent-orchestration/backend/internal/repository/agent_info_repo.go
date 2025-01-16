package repository

import (
	"context"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/internal/core/ports"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"gorm.io/gorm"
)

type agentInfoRepo struct {
	db *gorm.DB
}

type AgentInfoRepo interface {
	UpdateById(ctx context.Context, id uint, updatedFields map[string]interface{}) error
}

func (r *agentInfoRepo) UpdateAgentInfoById(ctx context.Context, id uint, updatedFields map[string]interface{}) error {
	result := r.db.WithContext(ctx).Model(&models.AgentInfo{}).Where("id = ?", id).Updates(updatedFields)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewAgentInfoRepository(db *gorm.DB) ports.IAgentInfoUseCase {
	return &agentInfoRepo{db}
}
