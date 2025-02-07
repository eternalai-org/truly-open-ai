package serializers

import (
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
)

type AgentInfraReq struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AuthenUrl   string `json:"authen_url"`
	Icon        string `json:"icon"`
}

type AgentInfraMissionReq struct {
	ID          uint             `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Prompt      string           `json:"prompt"`
	Price       numeric.BigFloat `json:"price"`
	ToolList    string           `json:"tool_list"`
	Icon        string           `json:"icon"`
}

type AuthenAgentInfraCallback struct {
	Code           string            `json:"code"`
	CallbackParams map[string]string `json:"callback_params"`
}

type AgentInfraResp struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	AuthenUrl   string    `json:"authen_url"`
	Icon        string    `json:"icon"`
}
