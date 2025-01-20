package repository

import (
	"context"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/models"
	"gorm.io/gorm"
)

type knowledgeBaseFileRepo struct {
	db *gorm.DB
}

type KnowledgeBaseFileRepo interface {
	Create(ctx context.Context, model *models.KnowledgeBaseFile) (*models.KnowledgeBaseFile, error)
	UpdateByKnowledgeBaseId(ctx context.Context, kbId uint, updatedFields map[string]interface{}) error
}

func (r *knowledgeBaseFileRepo) UpdateByKnowledgeBaseId(ctx context.Context, kbId uint, updatedFields map[string]interface{}) error {
	result := r.db.WithContext(ctx).
		Model(&models.KnowledgeBaseFile{}).
		Where("knowledge_base_id = ?", kbId).Updates(updatedFields)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *knowledgeBaseFileRepo) Create(ctx context.Context, model *models.KnowledgeBaseFile) (*models.KnowledgeBaseFile, error) {
	result := r.db.WithContext(ctx).Create(model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}

func NewKnowledgeBaseFileRepository(db *gorm.DB) KnowledgeBaseFileRepo {
	return &knowledgeBaseFileRepo{db}
}
