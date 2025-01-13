package appconfig

import (
	"context"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/internal/core/ports"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/internal/repository"
)

type appConfigUseCase struct {
	appConfigRepo repository.AppConfigRepo
}

func (uc *appConfigUseCase) GetAllNameValueInAppConfig(ctx context.Context, networkId string) (map[string]string, error) {
	return uc.appConfigRepo.GetAllNameValueInAppConfig(ctx, networkId)
}

func NewAppConfigUseCase(appConfigRepo repository.AppConfigRepo) ports.IAppConfigUseCase {
	return &appConfigUseCase{
		appConfigRepo: appConfigRepo,
	}
}
