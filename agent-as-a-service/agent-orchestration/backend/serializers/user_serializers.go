package serializers

import (
	"encoding/json"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
)

type UserProfileReq struct {
	Username    string                 `json:"username"`
	Description string                 `json:"description"`
	ImageURL    string                 `json:"image_url"`
	Social      map[string]interface{} `json:"social"`
}

type UserResp struct {
	ID              uint                   `json:"id"`
	CreatedAt       time.Time              `json:"created_at"`
	NetworkID       uint64                 `json:"network_id"`
	Address         string                 `json:"address"`
	Username        string                 `json:"username"`
	SubscriptionNum uint                   `json:"subscription_num"`
	Description     string                 `json:"description"`
	ImageURL        string                 `json:"image_url"`
	Social          map[string]interface{} `json:"social"`
	Price30d        numeric.BigFloat       `json:"price30d"`
	Price90d        numeric.BigFloat       `json:"price90d"`
	Subscribed      bool                   `json:"subscribed"`
	TotalLike       uint                   `json:"total_like"`
	TotalPost       uint                   `json:"total_post"`
	TotalMessage    uint                   `json:"total_message"`
	TipPayment      numeric.BigFloat       `json:"tip_payment"`
	TipReceive      numeric.BigFloat       `json:"tip_receive"`
	TwitterID       string                 `json:"twitter_id"`
	TwitterAvatar   string                 `json:"twitter_avatar"`
	TwitterUsername string                 `json:"twitter_username"`
	TwitterName     string                 `json:"twitter_name"`
}

func NewUserResp(m *models.User) *UserResp {
	if m == nil {
		return nil
	}

	social := make(map[string]interface{})
	json.Unmarshal([]byte(m.Social), &social)

	resp := &UserResp{
		ID:              m.ID,
		CreatedAt:       m.CreatedAt,
		NetworkID:       m.NetworkID,
		Address:         m.Address,
		Username:        m.Username,
		Description:     m.Description,
		ImageURL:        m.ImageURL,
		Social:          social,
		Price30d:        m.Price30d,
		Price90d:        m.Price90d,
		Subscribed:      m.Subscribed,
		SubscriptionNum: m.SubscriptionNum,
		TotalLike:       m.TotalLike,
		TotalPost:       m.TotalPost,
		TotalMessage:    m.TotalMessage,
		TipPayment:      m.TipPayment,
		TipReceive:      m.TipReceive,
		TwitterID:       m.TwitterID,
		TwitterAvatar:   m.TwitterAvatar,
		TwitterUsername: m.TwitterUsername,
		TwitterName:     m.TwitterName,
	}
	return resp
}

func NewUserRespArr(arr []*models.User) []*UserResp {
	resps := []*UserResp{}
	for _, m := range arr {
		resps = append(resps, NewUserResp(m))
	}
	return resps
}
