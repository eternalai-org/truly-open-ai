package repository

import (
	"context"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"gorm.io/gorm"
)

type knowledgeBaseFileRepo struct {
	db *gorm.DB
}

type KnowledgeBaseFileRepo interface {
	Create(ctx context.Context, model *models.KnowledgeBaseFile) (*models.KnowledgeBaseFile, error)
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
