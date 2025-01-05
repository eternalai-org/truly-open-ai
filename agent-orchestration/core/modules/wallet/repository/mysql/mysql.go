package mysql

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/eternalai-org/eternal-ai/agent-orchestration/core/domain"
)

type mysqlWalletRepository struct {
	DB *gorm.DB
}

// NewMysqlWalletRepository ...
func NewMysqlWalletRepository(db *gorm.DB) domain.IWalletRepository {
	return &mysqlWalletRepository{
		DB: db,
	}
}

func (m *mysqlWalletRepository) Create(ctx context.Context, w *domain.Wallet) error {
	if err := m.DB.Create(w).Error; err != nil {
		return errors.Wrap(err, "m.DB.Create")
	}
	return nil
}

func (m *mysqlWalletRepository) Update(ctx context.Context, w *domain.Wallet) error {
	if err := m.DB.Save(w).Error; err != nil {
		return errors.Wrap(err, "m.DB.Update")
	}
	return nil
}

func (m *mysqlWalletRepository) Delete(ctx context.Context, w *domain.Wallet) error {
	if err := m.DB.Delete(w).Error; err != nil {
		return errors.Wrap(err, "m.DB.Delete")
	}
	return nil
}

func (m *mysqlWalletRepository) FindByAddress(ctx context.Context, address string) (*domain.Wallet, error) {
	var model domain.Wallet
	if err := m.DB.Where("address = ?", address).First(&model).Error; err != nil {
		return nil, err
	}
	return &model, nil
}
