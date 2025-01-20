package agent_info

import (
	"context"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/internal/core/ports"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/internal/repository"
)

type agentInfoUseCase struct {
	agentInfoRepo repository.IAgentInfoRepo
}

func (uc *agentInfoUseCase) UpdateAgentInfoById(ctx context.Context, id uint, updatedFields map[string]interface{}) error {
	return uc.agentInfoRepo.UpdateById(ctx, id, updatedFields)
}

func NewAgentInfoUseCase(repo repository.IAgentInfoRepo) ports.IAgentInfoUseCase {
	return &agentInfoUseCase{
		agentInfoRepo: repo,
	}
}
