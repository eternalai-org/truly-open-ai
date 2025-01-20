package services

import (
	"context"
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func (s *Service) GetUser(tx *gorm.DB, networkID uint64, address string, forUpdate bool) (*models.User, error) {
	address = strings.TrimSpace(address)
	if address == "" {
		return nil, errs.NewError(errs.ErrBadRequest)
	}
	user, err := s.dao.FirstUser(
		tx,
		map[string][]interface{}{
			"network_id = ?": {networkID},
			"address = ?":    {strings.ToLower(address)},
		},
		map[string][]interface{}{},
		false,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	if user == nil {
		user = &models.User{
			NetworkID: networkID,
			Address:   strings.ToLower(address),
			Username:  uuid.NewString(),
		}
		err = s.dao.Create(tx, user)
		if err != nil {
			return nil, errs.NewError(err)
		}
		user.Username = fmt.Sprintf("user%d", user.ID)
		err = s.dao.Save(tx, user)
		if err != nil {
			return nil, errs.NewError(err)
		}
	}
	if forUpdate {
		user, err = s.dao.FirstUserByID(
			tx, user.ID,
			map[string][]interface{}{},
			true,
		)
		if err != nil {
			return nil, errs.NewError(err)
		}
	}
	return user, nil
}

func (s *Service) GetUserProfile(ctx context.Context, address string) (*models.User, error) {
	user, err := s.GetUser(daos.GetDBMainCtx(ctx), models.GENERTAL_NETWORK_ID, address, false)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return user, nil
}

func (s *Service) UserUploadFile(ctx context.Context, address string, fileHeader *multipart.FileHeader) (string, error) {
	fileToImport, err := fileHeader.Open()
	if err != nil {
		return "", errs.NewError(err)
	}
	defer fileToImport.Close()

	filename := fmt.Sprintf("%s.%s", uuid.NewString(), helpers.GetFileExtension(fileHeader.Filename))
	urlPath, err := s.gsClient.UploadPublicMultipartFile("imagine", filename, fileToImport, fileHeader)
	if err != nil {
		return "", errs.NewError(err)
	}
	return fmt.Sprintf("%s%s", s.conf.GsStorage.Url, urlPath), nil
}

func (s *Service) VerifyLoginUserByWeb3(ctx context.Context, userIP, userAgent, address, message, signature string) (string, error) {
	authTk := ""
	if !strings.HasPrefix(signature, "0x") {
		signature = "0x" + signature
	}
	err := helpers.VerifySignature(address, signature, message)
	if err != nil {
		return authTk, errs.NewError(err)
	}

	expireTime := int64(0)
	session := fmt.Sprintf("%s_%s_%s_%s_%s", userAgent, userIP, message, address, fmt.Sprintf("%d", expireTime))
	sessionID := crypto.Keccak256Hash([]byte(session))
	authToken := &helpers.AuthToken{
		Address:   address,
		Exp:       expireTime,
		SessionID: sessionID.Hex(),
	}
	encryptedAuthToken, err := helpers.EncryptAndSignAuthToken(*authToken, s.conf.EncryptAuthenKey)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("TK1 %s", encryptedAuthToken), nil
}
