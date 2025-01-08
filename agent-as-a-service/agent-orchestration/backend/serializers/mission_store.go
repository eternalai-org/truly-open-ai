package serializers

type MissionStoreReq struct {
	ID           uint   `json:"id"`
	OwnerAddress string `json:"owner_address"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Prompt       string `json:"prompt"`
	Price        uint   `json:"price"`
	DurationDay  uint   `json:"duration_day"`
	ToolList     string `json:"tool_list"`
}

type MissionStoreRatingReq struct {
	HistoryID   uint    `json:"history_id"`
	UserAddress string  `json:"user_address"`
	Rating      float64 `json:"rating"`
	Comment     string  `json:"comment"`
}
