package usecase

import (
	"context"
	"fmt"
	"os"

	"github.com/eternalai-org/eternal-ai/agent-orchestration/core/domain"
	"github.com/eternalai-org/eternal-ai/agent-orchestration/core/utils"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type walletUsecase struct {
	walletRepo domain.IWalletRepository
}

// NewWalletUsecase ...
func NewWalletUsecase(wr domain.IWalletRepository) domain.IWalletUsecase {
	return &walletUsecase{
		walletRepo: wr,
	}
}

func (wu *walletUsecase) GenerateUpdate(ctx context.Context, req domain.WalletUpdateRequest) (*domain.Wallet, error) {
	var err error
	if req.Address == "" {
		return nil, domain.ErrBadParamInput
	}
	if req.PrivateKey == "" {
		return nil, domain.ErrBadParamInput
	}
	address := req.Address
	privateKey := req.PrivateKey
	walletCipherKey := viper.GetString(`wallet_salt`)
	if walletCipherKey == "" {
		walletCipherKey = os.Getenv("WALLET_SALT")
	}
	if walletCipherKey == "" {
		return nil, domain.ErrBadParamInput
	}
	encryptedText, err := utils.EncryptToString(privateKey, walletCipherKey)
	if err != nil {
		return nil, domain.ErrorWithMessage(domain.ErrSystemError, errors.Wrap(err, "Encrypted error").Error())
	}
	decryptedText, err := utils.DecryptToString(encryptedText, walletCipherKey)
	if err != nil {
		return nil, domain.ErrorWithMessage(domain.ErrSystemError, errors.Wrap(err, "Decrypted error").Error())
	}
	if decryptedText != privateKey {
		return nil, domain.ErrorWithMessage(domain.ErrSystemError, "check decrypted failed")
	}
	w := domain.Wallet{
		WalletId:   req.WalletId,
		Address:    address,
		PrivateKey: encryptedText,
		Type:       req.WalletType,
	}
	err = wu.walletRepo.Create(ctx, &w)
	if err != nil {
		return nil, domain.ErrorWithMessage(domain.ErrSystemError, errors.Wrap(err, "Encrypted error").Error())
	}
	return &w, nil
}

// GetPrivateKey ...
func (wu *walletUsecase) GetPrivateKey(ctx context.Context, req domain.WalletGetPrivateKeyRequest) (string, error) {
	fmt.Println(req.Address)
	w, err := wu.walletRepo.FindByAddress(ctx, req.Address)
	if err != nil {
		return "", domain.ErrorWithMessage(domain.ErrSystemError, "uu.walletRepo.FindByAddress "+err.Error())
	}
	walletCipherKey := viper.GetString(`wallet_salt`)
	if walletCipherKey == "" {
		walletCipherKey = os.Getenv("WALLET_SALT")
	}
	if walletCipherKey == "" {
		return "", domain.ErrorWithMessage(domain.ErrSystemError, "walletCipherKey empty")
	}
	decryptedText, err := utils.DecryptToString(w.PrivateKey, walletCipherKey)
	if err != nil {
		return "", domain.ErrorWithMessage(domain.ErrSystemError, errors.Wrap(err, "Decrypted error").Error())
	}

	return string(decryptedText), nil
}
