package models

import "github.com/sashabaranov/go-openai"

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
	Status                    InferResultStatus `json:"status"`
	openai.CompletionResponse `json:",inline"`
	ChainInfo                 ChainInfoRequest      `json:"chain_info"`
	WorkerHubAddress          string                `json:"worker_hub_address"`
	OnChainData               CompletionOnChainData `json:"on_chain_data"`
}

type CompletionOnChainData struct {
	InferID             uint64   `json:"infer_id"`
	AssignmentAddresses []string `json:"pbft_committee"`
	SubmitAddress       string   `json:"proposer"`
	InferTx             string   `json:"infer_tx"`
	SubmitTx            string   `json:"propose_tx"`
	SeizeMinerTx        string   `json:"-"`
	InputCid            string   `json:"input_cid"`
	OutputCid           string   `json:"output_cid"`
}
