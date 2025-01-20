package models

type ChainInfoRequest struct {
	Rpc    string `json:"rpc"`
	ZkSync bool   `json:"zk_sync"`
}
type DecentralizeInferRequest struct {
	ChainInfo            ChainInfoRequest `json:"chain_info"`
	AgentContractAddress string           `json:"agent_contract_address"`
	InferPriKey          string           `json:"infer_pri_key"`
	Input                string           `json:"input"`
}

type DecentralizeInferResponse struct {
	ChainInfo ChainInfoRequest `json:"chain_info"`
	InferId   string           `json:"infer_id"`
	TxHash    string           `json:"tx_hash"`
}

type InferResultRequest struct {
	ChainInfo                ChainInfoRequest `json:"chain_info"`
	WorkerHubContractAddress string           `json:"worker_hub_contract_address"`
	InferId                  string           `json:"infer_id"`
}

type InferResultResponse struct {
	ChainInfo                ChainInfoRequest `json:"chain_info"`
	WorkerHubContractAddress string           `json:"worker_hub_contract_address"`
	InferId                  string           `json:"infer_id"`
	Input                    string           `json:"input"`
	Output                   string           `json:"output"`
}
