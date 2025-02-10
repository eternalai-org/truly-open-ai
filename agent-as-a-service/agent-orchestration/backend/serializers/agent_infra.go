package serializers

import (
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
)

type AgentInfraReq struct {
	ID          uint                    `json:"id"`
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Docs        string                  `json:"docs"`
	Icon        string                  `json:"icon"`
	Status      models.AgentInfraStatus `json:"status"`
	ApiUrl      string                  `json:"api_url"`
	Price       numeric.BigFloat        `json:"price"`
}

type AgentInfraResp struct {
	ID           uint                    `json:"id"`
	CreatedAt    time.Time               `json:"created_at"`
	InfraId      string                  `json:"infra_id"`
	OwnerID      uint                    `json:"owner_id"`
	OwnerAddress string                  `json:"owner_address"`
	Name         string                  `json:"name"`
	Description  string                  `json:"description"`
	Docs         string                  `json:"docs"`
	Icon         string                  `json:"icon"`
	Status       models.AgentInfraStatus `json:"status"`
	ApiUrl       string                  `json:"api_url"`
	Price        numeric.BigFloat        `json:"price"`
}

func NewAgentInfraResp(m *models.AgentInfra) *AgentInfraResp {
	if m == nil {
		return nil
	}
	return &AgentInfraResp{
		ID:           m.ID,
		CreatedAt:    m.CreatedAt,
		InfraId:      m.InfraId,
		OwnerID:      m.OwnerID,
		OwnerAddress: m.OwnerAddress,
		Name:         m.Name,
		Description:  m.Description,
		Docs:         m.Docs,
		Icon:         m.Icon,
		Status:       m.Status,
		ApiUrl:       m.ApiUrl,
		Price:        m.Price,
	}
}

func NewAgentInfraRespArray(arr []*models.AgentInfra) []*AgentInfraResp {
	resps := []*AgentInfraResp{}
	for _, r := range arr {
		resps = append(resps, NewAgentInfraResp(r))
	}
	return resps
}
