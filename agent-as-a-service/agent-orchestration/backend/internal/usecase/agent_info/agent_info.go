package agent_info

import (
	"context"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/internal/core/ports"
)

type agentInfoUseCase struct {
	agentInfoRepo ports.IAgentInfoUseCase
}

func (uc *agentInfoUseCase) UpdateAgentInfoById(ctx context.Context, id uint, updatedFields map[string]interface{}) error {
	return uc.agentInfoRepo.UpdateAgentInfoById(ctx, id, updatedFields)
}

func NewAgentInfoUseCase(repo ports.IAgentInfoUseCase) ports.IAgentInfoUseCase {
	return &agentInfoUseCase{
		agentInfoRepo: repo,
	}
}
