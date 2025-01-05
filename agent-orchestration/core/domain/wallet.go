package domain

import (
	"context"

	"github.com/jinzhu/gorm"
)

// WalletType ...
type WalletType int

const (
	// WalletTypeEthereum ...
	WalletTypeEthereum WalletType = iota
	// WalletTypeBitcoin ...
	WalletTypeBitcoin WalletType = 1
	// WalletTypeBEP2 ...
	WalletTypeBEP2 WalletType = 2
	// WalletTypeBeam ...
	WalletTypeBeam WalletType = 3
	// WalletTypeIncognito ...
	WalletTypeIncognito WalletType = 4
	// WalletTypeTron ...
	WalletTypeTron WalletType = 6
	// WalletTypeAda ...
	WalletTypeAda WalletType = 7
	// WalletTypeSOL ...
	WalletTypeSOL WalletType = 8
)

// Wallet ...
type Wallet struct {
	gorm.Model
	WalletId   string
	Address    string
	PrivateKey string `gorm:"type:text"`
	Type       WalletType
}

// func (*Wallet) TableName() string {
// 	return "wallet.wallets"
// }

// IWalletRepository ...
type IWalletRepository interface {
	Create(ctx context.Context, u *Wallet) error
	Update(ctx context.Context, u *Wallet) error
	Delete(ctx context.Context, u *Wallet) error
	FindByAddress(ctx context.Context, address string) (*Wallet, error)
}

// IWalletUsecase ...
type IWalletUsecase interface {
	GenerateUpdate(ctx context.Context, req WalletUpdateRequest) (*Wallet, error)
	GetPrivateKey(ctx context.Context, req WalletGetPrivateKeyRequest) (string, error)
}

// WalletGenerateRequest ...
type WalletGenerateRequest struct {
	WalletType WalletType `json:"WalletType"`
}

type WalletUpdateRequest struct {
	WalletId   string            `json:"WalletId"`
	WalletType WalletType        `json:"WalletType"`
	Address    string            `json:"Address"`
	PrivateKey string            `json:"PrivateKey"`
	Batch      map[string]string `json:"Batch"`
}

// WalletGetPrivateKeyRequest ...
type WalletGetPrivateKeyRequest struct {
	Address string `json:"Address"`
}

// WalletCheckRequest ...
type WalletCheckRequest struct {
	WalletType WalletType `json:"WalletType"`
	Address    string     `json:"Address"`
}
