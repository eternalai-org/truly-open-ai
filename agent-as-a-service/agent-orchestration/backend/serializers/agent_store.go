package serializers

import (
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
)

type AgentStoreReq struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AuthenUrl   string `json:"authen_url"`
}

type AgentStoreResp struct {
	ID            uint                `json:"id"`
	CreatedAt     time.Time           `json:"created_at"`
	Name          string              `json:"name"`
	Description   string              `json:"description"`
	AuthenUrl     string              `json:"authen_url"`
	MissionStores []*MissionStoreResp `json:"mission_stores"`
}

func NewAgentStoreResp(m *models.AgentStore) *AgentStoreResp {
	if m == nil {
		return nil
	}
	return &AgentStoreResp{
		ID:            m.ID,
		CreatedAt:     m.CreatedAt,
		Name:          m.Name,
		Description:   m.Description,
		AuthenUrl:     m.AuthenUrl,
		MissionStores: NewMissionStoreRespArray(m.MissionStores),
	}
}
func NewAgentStoreRespArray(arr []*models.AgentStore) []*AgentStoreResp {
	resps := []*AgentStoreResp{}
	for _, r := range arr {
		resps = append(resps, NewAgentStoreResp(r))
	}
	return resps
}
