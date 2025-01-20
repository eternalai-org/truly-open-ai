package repository

import (
	"context"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/models"
	"gorm.io/gorm"
)

type appConfigRepo struct {
	db *gorm.DB
}

type AppConfigRepo interface {
	GetAllNameValueInAppConfig(ctx context.Context, networkId string) (map[string]string, error)
}

func (r *appConfigRepo) GetAllNameValueInAppConfig(ctx context.Context, networkId string) (map[string]string, error) {
	var data []*models.AppConfig
	err := r.db.WithContext(ctx).
		Where("network_id = ?", networkId).
		Find(&data).Error
	if err != nil {
		return nil, err
	}
	info := make(map[string]string)
	for _, config := range data {
		info[config.Name] = config.Value
	}
	return info, nil
}

func NewAppConfigRepository(db *gorm.DB) AppConfigRepo {
	return &appConfigRepo{db}
}
