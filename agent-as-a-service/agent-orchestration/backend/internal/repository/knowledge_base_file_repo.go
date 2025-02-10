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
	UpdateByKnowledgeBaseId(ctx context.Context, kbId uint, updatedFields map[string]interface{}) error
	ListByKbId(ctx context.Context, kbId uint) ([]*models.KnowledgeBaseFile, error)
	DeleteByIds(ctx context.Context, ids []uint) error
}

func (r *knowledgeBaseFileRepo) DeleteByIds(ctx context.Context, ids []uint) error {
	files := []*models.KnowledgeBaseFile{}
	err := r.db.WithContext(ctx).Delete(&files, ids).Error
	return err
}

func (r *knowledgeBaseFileRepo) ListByKbId(ctx context.Context, kbId uint) ([]*models.KnowledgeBaseFile, error) {
	files := []*models.KnowledgeBaseFile{}
	err := r.db.WithContext(ctx).Where("knowledge_base_id = ?", kbId).Find(&files).Error
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (r *knowledgeBaseFileRepo) UpdateByKnowledgeBaseId(ctx context.Context, kbFileId uint, updatedFields map[string]interface{}) error {
	result := r.db.WithContext(ctx).
		Model(&models.KnowledgeBaseFile{}).
		Where("id = ?", kbFileId).Updates(updatedFields)
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
