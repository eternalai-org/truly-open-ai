package serializers

import (
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
)

type AgentStoreReq struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AuthenUrl   string `json:"authen_url"`
}

type AgentStoreMissionReq struct {
	ID          uint             `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Prompt      string           `json:"prompt"`
	Price       numeric.BigFloat `json:"price"`
	ToolList    string           `json:"tool_list"`
	Icon        string           `json:"icon"`
}

type AuthenAgentStoreCallback struct {
	Code           string            `json:"code"`
	CallbackParams map[string]string `json:"callback_params"`
}

type AgentStoreResp struct {
	ID                 uint                     `json:"id"`
	CreatedAt          time.Time                `json:"created_at"`
	Name               string                   `json:"name"`
	Description        string                   `json:"description"`
	AuthenUrl          string                   `json:"authen_url"`
	AgentStoreMissions []*AgentStoreMissionResp `json:"agent_store_missions"`
}

func NewAgentStoreResp(m *models.AgentStore) *AgentStoreResp {
	if m == nil {
		return nil
	}
	return &AgentStoreResp{
		ID:                 m.ID,
		CreatedAt:          m.CreatedAt,
		Name:               m.Name,
		Description:        m.Description,
		AuthenUrl:          m.AuthenUrl,
		AgentStoreMissions: NewAgentStoreMissionRespArray(m.AgentStoreMissions),
	}
}
func NewAgentStoreRespArray(arr []*models.AgentStore) []*AgentStoreResp {
	resps := []*AgentStoreResp{}
	for _, r := range arr {
		resps = append(resps, NewAgentStoreResp(r))
	}
	return resps
}

type AgentStoreMissionResp struct {
	AgentStoreID uint             `json:"agent_store_id"`
	ID           uint             `json:"id"`
	CreatedAt    time.Time        `json:"created_at"`
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	UserPrompt   string           `json:"user_prompt"`
	Price        numeric.BigFloat `json:"price"`
	ToolList     string           `json:"tool_list"`
	Icon         string           `json:"icon"`
}

func NewAgentStoreMissionResp(m *models.AgentStoreMission) *AgentStoreMissionResp {
	if m == nil {
		return nil
	}
	return &AgentStoreMissionResp{
		AgentStoreID: m.AgentStoreID,
		ID:           m.ID,
		CreatedAt:    m.CreatedAt,
		Name:         m.Name,
		Description:  m.Description,
		UserPrompt:   m.UserPrompt,
		Price:        m.Price,
		ToolList:     m.ToolList,
		Icon:         m.Icon,
	}
}
func NewAgentStoreMissionRespArray(arr []*models.AgentStoreMission) []*AgentStoreMissionResp {
	resps := []*AgentStoreMissionResp{}
	for _, r := range arr {
		resps = append(resps, NewAgentStoreMissionResp(r))
	}
	return resps
}

func NewAgentStoreRespArrayFromInstall(arr []*models.AgentStoreInstall) []*AgentStoreResp {
	resps := []*AgentStoreResp{}
	for _, r := range arr {
		resps = append(resps, NewAgentStoreResp(r.AgentStore))
	}
	return resps
}
