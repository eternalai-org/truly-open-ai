package serializers

import (
	"time"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
)

type ApiTokenUsageReq struct {
	ApiKey    string `json:"api_key"`
	NumToken  uint   `json:"num_token"`
	Endpoint  string `json:"endpoint"`
	NetworkID uint64 `json:"network_id"`
}

type ApiSubscriptionPackageResp struct {
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       numeric.BigFloat   `json:"price"`
	NetworkID   uint64             `json:"network_id"`
	NumToken    uint64             `json:"num_token"`
	Type        models.PackageType `json:"type"`
	DurationDay uint               `json:"duration_day"`
}

func NewApiSubscriptionPackageResp(m *models.ApiSubscriptionPackage) *ApiSubscriptionPackageResp {
	if m == nil {
		return nil
	}
	return &ApiSubscriptionPackageResp{
		Name:        m.Name,
		Description: m.Description,
		Price:       m.Price,
		NetworkID:   m.NetworkID,
		NumToken:    m.NumToken,
		Type:        m.Type,
		DurationDay: m.DurationDay,
	}
}

func NewApiSubscriptionPackageRespArr(arr []*models.ApiSubscriptionPackage) []*ApiSubscriptionPackageResp {
	resps := []*ApiSubscriptionPackageResp{}
	for _, m := range arr {
		resps = append(resps, NewApiSubscriptionPackageResp(m))
	}
	return resps
}

type ApiSubscriptionKeyResp struct {
	NetworkID      uint64                      `json:"network_id"`
	UserAddress    string                      `json:"user_address"`
	TwitterID      string                      `json:"twitter_id"`
	TwitterInfoID  uint                        `json:"twitter_info_id"`
	ApiKey         string                      `json:"api_key"`
	PackageID      uint                        `json:"package_id"`
	Package        *ApiSubscriptionPackageResp `json:"package"`
	QuotaRemaining uint64                      `json:"quota_remaining"`
	StartedAt      *time.Time                  `json:"started_at"`
	ExpiresAt      *time.Time                  `json:"expires_at"`
	DepositAddress string                      `json:"deposit_address"`
}

func NewApiSubscriptionKeyResp(m *models.ApiSubscriptionKey) *ApiSubscriptionKeyResp {
	if m == nil {
		return nil
	}
	return &ApiSubscriptionKeyResp{
		NetworkID:      m.NetworkID,
		UserAddress:    m.UserAddress,
		TwitterID:      m.TwitterID,
		TwitterInfoID:  m.TwitterInfoID,
		ApiKey:         m.ApiKey,
		PackageID:      m.PackageID,
		Package:        NewApiSubscriptionPackageResp(&m.Package),
		QuotaRemaining: m.QuotaRemaining,
		StartedAt:      m.StartedAt,
		ExpiresAt:      m.ExpiresAt,
		DepositAddress: m.DepositAddress,
	}
}

func NewApiSubscriptionKeyRespArr(arr []*models.ApiSubscriptionKey) []*ApiSubscriptionKeyResp {
	resps := []*ApiSubscriptionKeyResp{}
	for _, m := range arr {
		resps = append(resps, NewApiSubscriptionKeyResp(m))
	}
	return resps
}

type ApiSubscriptionHistoryResp struct {
	NetworkID      uint64                      `json:"network_id"`
	UserAddress    string                      `json:"user_address"`
	ApiKey         string                      `json:"api_key"`
	PackageID      uint                        `json:"package_id"`
	Package        *ApiSubscriptionPackageResp `json:"package"`
	DepositAddress string                      `json:"deposit_address"`
	DepositStatus  models.DepositStatus        `json:"deposit_status"`
	TxHash         string                      `json:"tx_hash"`
	NumToken       uint64                      `json:"num_token"`
	StartedAt      *time.Time                  `json:"started_at"`
	ExpiresAt      *time.Time                  `json:"expires_at"`
}

func NewApiSubscriptionHistoryResp(m *models.ApiSubscriptionHistory) *ApiSubscriptionHistoryResp {
	if m == nil {
		return nil
	}
	return &ApiSubscriptionHistoryResp{
		NetworkID:      m.NetworkID,
		UserAddress:    m.UserAddress,
		ApiKey:         m.ApiKey,
		PackageID:      m.PackageID,
		Package:        NewApiSubscriptionPackageResp(&m.Package),
		DepositAddress: m.DepositAddress,
		DepositStatus:  m.DepositStatus,
		TxHash:         m.TxHash,
		NumToken:       m.NumToken,
		StartedAt:      m.StartedAt,
		ExpiresAt:      m.ExpiresAt,
	}
}

func NewApiSubscriptionHistoryRespArr(arr []*models.ApiSubscriptionHistory) []*ApiSubscriptionHistoryResp {
	resps := []*ApiSubscriptionHistoryResp{}
	for _, m := range arr {
		resps = append(resps, NewApiSubscriptionHistoryResp(m))
	}
	return resps
}

type ApiSubscriptionUsageLogResp struct {
	ApiKey   string `json:"api_key"`
	Endpoint string `json:"endpoint"`
	NumToken uint64 `json:"num_token"`
}

func NewApiSubscriptionUsageLogResp(m *models.ApiSubscriptionUsageLog) *ApiSubscriptionUsageLogResp {
	if m == nil {
		return nil
	}
	return &ApiSubscriptionUsageLogResp{
		ApiKey:   m.ApiKey,
		Endpoint: m.Endpoint,
		NumToken: m.NumToken,
	}
}

func NewApiSubscriptionUsageLogRespArr(arr []*models.ApiSubscriptionUsageLog) []*ApiSubscriptionUsageLogResp {
	resps := []*ApiSubscriptionUsageLogResp{}
	for _, m := range arr {
		resps = append(resps, NewApiSubscriptionUsageLogResp(m))
	}
	return resps
}
