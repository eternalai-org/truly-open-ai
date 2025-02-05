package model

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type BatchInferHistory struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id" bson:"user_id"`
	UserAddress string `json:"user_address" bson:"user_address"`
	Toolset     string `json:"toolset" bson:"toolset"`

	AgentContractAddress string `json:"agent_contract_address" bson:"agent_contract_address"`
	ContractAgentID      string `json:"contract_agent_id" bson:"contract_agent_id"`
	ChainID              string `json:"chain_id" bson:"chain_id"`

	AssistantID     string `json:"assistant_id" bson:"assistant_id"`
	PromptInput     string `json:"prompt_input" bson:"prompt_input"`
	SystemPrompt    string `bson:"system_prompt" json:"system_prompt"`
	PromptInputHash string `bson:"prompt_input_hash" json:"prompt_input_hash"`

	AgentType        int    `bson:"agent_type" json:"agent_type"`
	TwitterSnapshot  string `bson:"twitter_snapshot" json:"twitter_snapshot"` // file coin hash
	UserInfoSnapshot string `json:"user_info_snapshot" bson:"user_info_snapshot"`

	OutputMaxCharacter uint   `json:"output_max_character" bson:"output_max_character"`
	PromptOutput       string `json:"prompt_output" bson:"prompt_output"`
	PromptOutputHash   string `bson:"prompt_output_hash" json:"prompt_output_hash"`

	Status string `json:"status" bson:"status"`
	Log    string `json:"log" bson:"log"`

	InferID string `json:"infer_id" bson:"infer_id"` // when call to agent contract to create infer
	ModelID string `json:"model_id" bson:"model_id"`

	InscribeTxHash               string `json:"inscribe_tx_hash" bson:"inscribe_tx_hash"`
	SubmitSolutionInscribeTxHash string `json:"submit_solution_tx_hash" bson:"submit_solution_tx_hash"`

	BtcInscribeTxHash               string    `json:"btc_inscribe_tx_hash" bson:"btc_inscribe_tx_hash"`
	BtcSubmitSolutionInscribeTxHash string    `bson:"btc_submit_solution_inscribe_tx_hash" json:"btc_submit_solution_inscribe_tx_hash"`
	InferWalletAddress              string    `json:"infer_wallet_address" bson:"infer_wallet_address"`
	SubmitInferAt                   time.Time `json:"submit_infer_at" bson:"submit_infer_at"`

	AssignmentAddresses   []string `json:"assignment_addresses" bson:"assignment_addresses"`
	SubmitSolutionAddress string   `json:"submit_solution_address" bson:"submit_solution_address"`

	CommitTxHash []string `json:"commit_tx_hash" bson:"commit_tx_hash"`
	RevealTxHash []string `json:"reveal_tx_hash" bson:"reveal_tx_hash"`

	BtcCommitInscribeTxHash []string `bson:"btc_commit_inscribe_tx_hash" json:"btc_commit_inscribe_tx_hash"`
	BtcRevealInscribeTxHash []string `bson:"btc_reveal_inscribe_tx_hash" json:"btc_reveal_inscribe_tx_hash"`
}

type ExternalDataRequestType string

var ExternalDataRequestTypeBatch ExternalDataRequestType = "batch"

type AgentInferExternalData struct {
	RoomId             string                   `json:"room_id,omitempty"`
	AgentID            string                   `json:"agent_id,omitempty"`
	OutputMaxCharacter *uint                    `json:"output_max_character,omitempty"`
	Type               *ExternalDataRequestType `json:"type,omitempty"`
}

type StorageType string

type LLMInferRequest struct {
	Messages    []LLMInferMessage `json:"messages"`
	Model       string            `json:"model"`
	Seed        uint64            `json:"seed"`
	MaxToken    uint64            `json:"max_tokens"`
	Temperature float32           `json:"temperature"`
	Stream      bool              `json:"stream"`
}

type LLMInferMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type LLMInferResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role      string        `json:"role"`
			Content   string        `json:"content"`
			ToolCalls []interface{} `json:"tool_calls"`
		} `json:"message"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
		StopReason   interface{} `json:"stop_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens        int         `json:"prompt_tokens"`
		TotalTokens         int         `json:"total_tokens"`
		CompletionTokens    int         `json:"completion_tokens"`
		PromptTokensDetails interface{} `json:"prompt_tokens_details"`
	} `json:"usage"`
	PromptLogprobs interface{} `json:"prompt_logprobs"`
	OnchainData    struct {
		InferId       uint64   `json:"infer_id"`
		PbftCommittee []string `json:"pbft_committee"`
		Proposer      string   `json:"proposer"`
		InferTx       string   `json:"infer_tx"`
		ProposeTx     string   `json:"propose_tx"`
	} `json:"onchain_data"`
}

type Response struct {
	ResultURI string `json:"result_uri"`
	Storage   string `json:"storage"`
	Data      string `json:"data"`
}

type Assiment struct {
	InferenceId *big.Int
	Commitment  [32]byte
	Digest      [32]byte
	RevealNonce *big.Int
	Worker      common.Address
	Role        uint8
	Vote        uint8
	Output      []byte
}

type AssimentChan struct {
	Err         error
	Data        *Assiment
	AssismentID *big.Int
}

type TaskResult struct {
	ResultURI string      `json:"result_uri"`
	Storage   StorageType `json:"storage"`
	Data      []byte      `json:"data"`
}

type Event struct {
	InferenceId       *big.Int
	Model             common.Address
	Creator           common.Address
	Value             *big.Int
	OriginInferenceId *big.Int
	Input             []byte
	Flag              bool
	Raw               types.Log // Blockchain specific contextual infos
}

type InferInfo struct {
	Value          *big.Int
	ModelId        uint32
	SubmitTimeout  *big.Int
	Status         uint8
	Creator        common.Address
	ProcessedMiner common.Address
	Input          []byte
	Output         []byte
}

type InferInfoChan struct {
	Data *InferInfo
	Err  error
}

type MinerInfo struct {
	Address        string `json:"address"`
	Tasks          int    `json:"tasks"`
	ProcessedTasks int    `json:"processed_tasks"`
	ModelName      string `json:"model_name"`
	Balance        string `json:"balance"`
	Reward         string `json:"reward"`
	CurrentBlock   uint64 `json:"current_block"`
	ClusterID      string `json:"cluster_id"`

	//TODO - soon
	Status       string `json:"status"`
	StakedAmount string `json:"staked_amount"`
}

type Task struct {
	TaskID string `json:"task_id"`
	// ModelAddress     string `json:"model_name"`
	// ModelID       string `json:"model_id"`
	AssignmentID   string `json:"assignment_id"`
	ModelContract  string `json:"model_contract"`
	Params         string `json:"params"`
	Value          string `json:"value"`
	AssignmentRole string `json:"assignment_role"`

	ZKSync      bool   `json:"zk_sync"`
	Requestor   string `json:"requestor"`
	InferenceID string `json:"inference_id"`
	// TaskResult   *eaimodel.TaskResult `json:"task_result"`
	Status       uint8 `json:"status"`
	Retry        int
	BatchInfers  []*BatchInferHistory
	ExternalData *AgentInferExternalData
	IsBatch      bool
}

type EventPromptSchedulerNewInference struct {
	InferenceId uint64
	Creator     common.Address
	ModelId     uint32
	Value       *big.Int
	Input       []byte
	Flag        bool
	Raw         types.Log // Blockchain specific contextual infos
}
