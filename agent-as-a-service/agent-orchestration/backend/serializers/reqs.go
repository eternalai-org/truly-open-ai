package serializers

type SignatureReq struct {
	NetworkID uint64 `json:"network_id"`
	Address   string `json:"address"`
	Timestamp int64  `json:"timestamp"`
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

type SignatureTimestampReq struct {
	Timestamp int64  `json:"timestamp"`
	Signature string `json:"signature"`
}

type BaseReq struct {
	NetworkID uint64 `json:"network_id"`
	Address   string `json:"address"`
}
