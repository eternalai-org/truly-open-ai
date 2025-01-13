package ports

import "context"

type IAppConfigUseCase interface {
	GetAllNameValueInAppConfig(ctx context.Context, networkId string) (map[string]string, error)
}
