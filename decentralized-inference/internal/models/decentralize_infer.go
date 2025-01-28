package models

type ChainInfoRequest struct {
	Rpc    string `json:"rpc"`
	ZkSync bool   `json:"zk_sync"`
}
type DecentralizeInferRequest struct {
	ChainInfo            ChainInfoRequest `json:"chain_info"`
	AgentContractAddress string           `json:"agent_contract_address"`
	WorkerHubAddress     string           `json:"worker_hub_address"`
	InferPriKey          string           `json:"infer_pri_key"`
	ExternalData         string           `json:"external_data"`
	Input                string           `json:"input"`
	AgentId              string           `json:"agent_id"`
	Model                string           `json:"model"`
}

type DecentralizeInferNoAgentRequest struct {
	ChainInfo        ChainInfoRequest `json:"chain_info"`
	WorkerHubAddress string           `json:"worker_hub_address"`
	InferPriKey      string           `json:"infer_pri_key"`
	Input            string           `json:"input"`
	ModelId          string           `json:"model_id"`
}

type DecentralizeInferResponse struct {
	ChainInfo ChainInfoRequest `json:"chain_info"`
	InferId   uint64           `json:"infer_id"`
	TxHash    string           `json:"tx_hash"`
}

type InferResultRequest struct {
	ChainInfo        ChainInfoRequest `json:"chain_info"`
	WorkerHubAddress string           `json:"worker_hub_address"`
	InferId          uint64           `json:"infer_id"`
}

type InferResultStatus string

var InferResultStatusDone = InferResultStatus("done")
var InferResultStatusWaitingProcess = InferResultStatus("waiting_process")
var InferResultStatusTimeOut = InferResultStatus("timeout")

type InferResultResponse struct {
	ChainInfo        ChainInfoRequest  `json:"chain_info"`
	WorkerHubAddress string            `json:"worker_hub_address"`
	ModelId          uint32            `json:"model_id"`
	InferId          uint64            `json:"infer_id"`
	Input            string            `json:"input"`
	Output           string            `json:"output"`
	Status           InferResultStatus `json:"status"`
	SubmitTimeout    uint64            `json:"submit_timeout"`
	Creator          string            `json:"creator"`
	ProcessedMiner   string            `json:"processed_miner"`
	TxSubmitSolution string            `json:"tx_submit_solution"`
}
