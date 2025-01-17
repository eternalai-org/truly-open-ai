package ports

import (
	"context"
)

type IAgentInfoUseCase interface {
	UpdateAgentInfoById(ctx context.Context, id uint, updatedFields map[string]interface{}) error
}
