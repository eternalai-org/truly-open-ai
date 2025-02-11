package repository

import (
	"context"
	"math"

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
	CalcTotalFee(ctx context.Context, kbId uint) (float64, error)
	UpdateTransferHash(ctx context.Context, kbFileIds []uint, transferHash string) error
}

func round(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func (r *knowledgeBaseFileRepo) UpdateTransferHash(ctx context.Context, kbFileIds []uint, transferHash string) error {
	if err := r.db.WithContext(ctx).Model(&models.KnowledgeBaseFile{}).Where("id IN (?)", kbFileIds).Update("transfer_hash", transferHash).Error; err != nil {
		return err
	}
	return nil
}

func (r *knowledgeBaseFileRepo) CalcTotalFee(ctx context.Context, kbId uint) (float64, error) {
	files := []*models.KnowledgeBaseFile{}
	err := r.db.WithContext(ctx).Unscoped().Where("knowledge_base_id = ?", kbId).Find(&files).Error
	if err != nil {
		return 0, err
	}

	unitPrice := 10
	total := float64(0)
	for _, r := range files {
		if r.DeletedAt.Valid && r.Status == models.KnowledgeBaseFileStatusPending {
			continue
		}
		total += float64(r.FileSize)
	}

	price := total / 1_000_000 // 1 Megabyte is equal to 1000000 bytes (decimal).
	price = round(price, 0)
	if price == 0 {
		price = 1
	}
	return price * float64(unitPrice), nil
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
