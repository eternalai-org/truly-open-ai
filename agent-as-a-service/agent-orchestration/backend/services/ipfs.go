package services

import (
	"context"
	"fmt"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/lighthouse"
	"github.com/google/uuid"
)

func (s *Service) IpfsUploadData(ctx context.Context, ext string, dataBytes []byte) (string, error) {
	hash, err := lighthouse.UploadData(s.conf.Lighthouse.Apikey, fmt.Sprintf("%s.%s", uuid.New(), ext), dataBytes)
	if err != nil {
		return "", errs.NewError(err)
	}
	return hash, nil
}

func (s *Service) IpfsUploadDataForName(ctx context.Context, fileName string, dataBytes []byte) (string, error) {
	hash, err := lighthouse.UploadData(s.conf.Lighthouse.Apikey, fileName, dataBytes)
	if err != nil {
		return "", errs.NewError(err)
	}
	return hash, nil
}
