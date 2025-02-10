package types

type DecentralizeInferRequest struct {
	ChainID               string `json:"chain_id"`
	AgentID               string `json:"agent_id"`
	AgentContractAddress  string `json:"agent_contract_address"`
	InferWalletPrivateKey string `json:"infer_wallet_private_key"`

	Prompt string `json:"prompt"`
}

type DecentralizeInferResponse struct {
	InferID     string `json:"infer_id"`
	InferTxHash string `json:"infer_tx_hash"`
}

type StreamData struct {
	Data     []byte
	Stop     bool
	StreamID int
	Err      error
}
