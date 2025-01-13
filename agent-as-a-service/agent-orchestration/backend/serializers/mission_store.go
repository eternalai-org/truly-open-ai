package serializers

import (
	"encoding/json"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
)

type MissionParam struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
type MissionStoreReq struct {
	ID           uint            `json:"id"`
	OwnerAddress string          `json:"owner_address"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	Prompt       string          `json:"prompt"`
	Price        uint            `json:"price"`
	DurationDay  uint            `json:"duration_day"`
	ToolList     string          `json:"tool_list"`
	Icon         string          `json:"icon"`
	OutputType   string          `json:"output_type"`
	Params       []*MissionParam `json:"params"`
}

type MissionStoreRatingReq struct {
	HistoryID   uint    `json:"history_id"`
	UserAddress string  `json:"user_address"`
	Rating      float64 `json:"rating"`
	Comment     string  `json:"comment"`
}

type MissionStoreResp struct {
	ID             uint            `json:"id"`
	CreatedAt      time.Time       `json:"created_at"`
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	UserPrompt     string          `json:"user_prompt"`
	Price          uint            `json:"price"`
	OwnerAddress   string          `json:"owner_address"`
	ToolList       string          `json:"tool_list"`
	DepositAddress string          `json:"deposit_address"`
	DurationDay    uint            `json:"duration_day"`
	Rating         float64         `json:"rating"`
	NumRating      uint            `json:"num_rating"`
	NumUsed        uint            `json:"num_used"`
	Icon           string          `json:"icon"`
	OutputType     string          `json:"output_type"`
	Params         []*MissionParam `json:"params"`
}

type MissionStoreHistoryResp struct {
	ID             uint              `json:"id"`
	CreatedAt      time.Time         `json:"created_at"`
	UserAddress    string            `json:"user_address"`
	MissionStoreID uint              `json:"mission_store_id"`
	MissionStore   *MissionStoreResp `json:"mission_store"`
	TxHash         string            `json:"tx_hash"`
	EventId        string            `json:"event_id"`
	StartedAt      *time.Time        `json:"started_at"`
	ExpiresAt      *time.Time        `json:"expires_at"`
	IsRated        bool              `json:"is_rated"`
}

type MissionStoreRatingResp struct {
	ID             uint              `json:"id"`
	CreatedAt      time.Time         `json:"created_at"`
	UserAddress    string            `json:"user_address"`
	MissionStoreID uint              `json:"mission_store_id"`
	MissionStore   *MissionStoreResp `json:"mission_store"`
	Rating         float64           `json:"rating"`
	Comment        string            `json:"comment"`
}

func NewMissionStoreResp(m *models.MissionStore) *MissionStoreResp {
	if m == nil {
		return nil
	}
	params := []*MissionParam{}
	if m.Params != "" {
		json.Unmarshal([]byte(m.Params), &params)
	}
	return &MissionStoreResp{
		ID:           m.ID,
		CreatedAt:    m.CreatedAt,
		Name:         m.Name,
		Description:  m.Description,
		UserPrompt:   m.UserPrompt,
		Price:        m.Price,
		OwnerAddress: m.OwnerAddress,
		ToolList:     m.ToolList,
		Icon:         m.Icon,
		Rating:       m.Rating,
		NumRating:    m.NumRating,
		NumUsed:      m.NumUsed,
		OutputType:   string(m.OutputType),
		Params:       params,
	}
}

func NewMissionStoreHistoryResp(h *models.MissionStoreHistory) *MissionStoreHistoryResp {
	if h == nil {
		return nil
	}
	return &MissionStoreHistoryResp{
		ID:             h.ID,
		CreatedAt:      h.CreatedAt,
		UserAddress:    h.UserAddress,
		MissionStoreID: h.MissionStoreID,
		MissionStore:   NewMissionStoreResp(h.MissionStore),
		TxHash:         h.TxHash,
		EventId:        h.EventId,
		StartedAt:      h.StartedAt,
		ExpiresAt:      h.ExpiresAt,
		IsRated:        h.IsRated,
	}
}

func NewMissionStoreRatingResp(r *models.MissionStoreRating) *MissionStoreRatingResp {
	if r == nil {
		return nil
	}
	return &MissionStoreRatingResp{
		ID:             r.ID,
		CreatedAt:      r.CreatedAt,
		UserAddress:    r.UserAddress,
		MissionStoreID: r.MissionStoreID,
		MissionStore:   NewMissionStoreResp(r.MissionStore),
		Rating:         r.Rating,
		Comment:        r.Comment,
	}
}

func NewMissionStoreRespArray(arr []*models.MissionStore) []*MissionStoreResp {
	resps := []*MissionStoreResp{}
	for _, m := range arr {
		resps = append(resps, NewMissionStoreResp(m))
	}
	return resps
}

func NewMissionStoreHistoryRespArray(arr []*models.MissionStoreHistory) []*MissionStoreHistoryResp {
	resps := []*MissionStoreHistoryResp{}
	for _, h := range arr {
		resps = append(resps, NewMissionStoreHistoryResp(h))
	}
	return resps
}

func NewMissionStoreRatingRespArray(arr []*models.MissionStoreRating) []*MissionStoreRatingResp {
	resps := []*MissionStoreRatingResp{}
	for _, r := range arr {
		resps = append(resps, NewMissionStoreRatingResp(r))
	}
	return resps
}
