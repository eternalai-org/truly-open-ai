package repository

import (
	"context"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"gorm.io/gorm"
)

type knowledgeBaseRepo struct {
	db *gorm.DB
}

type KnowledgeBaseRepo interface {
	GetKnowledgeBaseById(ctx context.Context, id uint) (*models.KnowledgeBase, error)
	DeleteKnowledgeBaseById(ctx context.Context, id uint) error
	CreateKnowledgeBase(ctx context.Context, model *models.KnowledgeBase) (*models.KnowledgeBase, error)
	GetKnowledgeBaseByStatus(ctx context.Context, status models.KnowledgeBaseStatus, offset, limit int) ([]*models.KnowledgeBase, error)
	ListKnowledgeBaseByAddress(ctx context.Context, address string) ([]*models.KnowledgeBase, error)
	UpdateStatus(ctx context.Context, model *models.KnowledgeBase) error
	UpdateKnowledgeBaseById(ctx context.Context, id uint, updatedFields map[string]interface{}) error
}

func (r *knowledgeBaseRepo) UpdateStatus(ctx context.Context, model *models.KnowledgeBase) error {
	return r.db.Model(model).Update("status", model.Status).Error
}

func (r *knowledgeBaseRepo) ListKnowledgeBaseByAddress(ctx context.Context, address string) ([]*models.KnowledgeBase, error) {
	knowledges := []*models.KnowledgeBase{}
	err := r.db.WithContext(ctx).
		Preload("KnowledgeBaseFiles").
		Where("user_address = ?", address).
		Find(&knowledges).Error
	if err != nil {
		return nil, err
	}
	return knowledges, nil
}

func (r *knowledgeBaseRepo) CreateKnowledgeBase(ctx context.Context, model *models.KnowledgeBase) (*models.KnowledgeBase, error) {
	result := r.db.WithContext(ctx).Create(model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}

func (r *knowledgeBaseRepo) GetKnowledgeBaseById(ctx context.Context, id uint) (*models.KnowledgeBase, error) {
	knowledge := &models.KnowledgeBase{}
	err := r.db.WithContext(ctx).
		Preload("KnowledgeBaseFiles").
		Where("id = ?", id).
		First(knowledge).Error
	if err != nil {
		return nil, err
	}
	return knowledge, nil
}

func (r *knowledgeBaseRepo) DeleteKnowledgeBaseById(ctx context.Context, id uint) error {
	tx := r.db.WithContext(ctx)
	return tx.Transaction(func(tx *gorm.DB) error {
		data := []*models.KnowledgeBaseFile{}
		if err := tx.Where("knowledge_base_id = ?", id).Delete(&data).Error; err != nil {
			return err
		}

		return tx.Delete(&models.KnowledgeBase{}, id).Error
	})
}

func (r *knowledgeBaseRepo) GetKnowledgeBaseByStatus(ctx context.Context, status models.KnowledgeBaseStatus, offset, limit int) ([]*models.KnowledgeBase, error) {
	var data []*models.KnowledgeBase
	err := r.db.WithContext(ctx).
		Preload("KnowledgeBaseFiles").
		Where("status = ?", status).
		Offset(offset).
		Limit(limit).
		Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *knowledgeBaseRepo) UpdateKnowledgeBaseById(ctx context.Context, id uint, updatedFields map[string]interface{}) error {
	result := r.db.WithContext(ctx).Model(&models.KnowledgeBase{}).Where("id = ?", id).Updates(updatedFields)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewKnowledgeBaseRepository(db *gorm.DB) KnowledgeBaseRepo {
	return &knowledgeBaseRepo{db}
}
