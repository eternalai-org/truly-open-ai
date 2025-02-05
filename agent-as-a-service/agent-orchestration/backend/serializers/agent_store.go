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

type AgentStoreMissionReq struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Prompt      string `json:"prompt"`
	Price       uint   `json:"price"`
	ToolList    string `json:"tool_list"`
	Icon        string `json:"icon"`
	OutputType  string `json:"output_type"`
}

type AuthenAgentStoreCallback struct {
	AgentStoreID       uint              `json:"agent_store_id"`
	InstallAgentInfoID uint              `json:"install_agent_info_id"`
	CallbackParams     map[string]string `json:"callback_params"`
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
