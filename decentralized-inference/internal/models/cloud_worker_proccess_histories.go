package models

import (
	"decentralized-inference/internal/eaimodel"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WorkerProcessStatus int

const (
	CLOUD_PROCESSING_STATUS_SUBMIT_SOLUTIONL_CONTRACT_EXPIRED WorkerProcessStatus = -6
	CLOUD_PROCESSING_STATUS_REVEAL_CONTRACT_EXPIRED           WorkerProcessStatus = -5 // task is expired (contract)
	CLOUD_PROCESSING_STATUS_COMMIT_CONTRACT_EXPIRED           WorkerProcessStatus = -4 // task is expired (contract)

	CLOUD_PROCESSING_STATUS_PROCESSING_EXPIRED WorkerProcessStatus = -3 // task is expired (processing_time)
	CLOUD_PROCESSING_STATUS_CONTRACT_EXPIRED   WorkerProcessStatus = -2 // task is expired (contract)
	CLOUD_PROCESSING_STATUS_ERROR              WorkerProcessStatus = -1 //error is negative
	CLOUD_PROCESSING_STATUS_PENDING            WorkerProcessStatus = 0
	CLOUD_PROCESSING_STATUS_PROCESSING         WorkerProcessStatus = 1
	CLOUD_PROCESSING_STATUS_DONE               WorkerProcessStatus = 2

	CLOUD_PROCESSING_STATUS_SEIZE_MINER      WorkerProcessStatus = 3
	CLOUD_PROCESSING_STATUS_SEIZE_MINER_DONE WorkerProcessStatus = 4
	// CLOUD_PROCESSING_STATUS_SEIZE_MINER -> CLOUD_PROCESSING_STATUS_SEIZE_MINER_DONE
	// event -> CLOUD_PROCESSING_STATUS_PENDING

	// CLOUD_PROCESSING_STATUS_PROCESSING -> ... CLOUD_PROCESSING_STATUS_SUBMIT_SOLUTION,CLOUD_PROCESSING_STATUS_COMMIT
	CLOUD_PROCESSING_STATUS_WAIT_SUBMIT_SOLUTION WorkerProcessStatus = 5
	CLOUD_PROCESSING_STATUS_COMMIT_DONE          WorkerProcessStatus = 6
	CLOUD_PROCESSING_STATUS_REVEAL               WorkerProcessStatus = 7
	CLOUD_PROCESSING_STATUS_REVEAL_DONE          WorkerProcessStatus = 8
)

type ModelWorkerProcessHistories struct {
	mgm.DefaultModel             `bson:",inline"`
	WorkerID                     primitive.ObjectID   `bson:"worker_id" json:"worker_id"`
	AssignmentId                 string               `bson:"assignment_id" json:"assignment_id"`
	InferenceId                  string               `bson:"inference_id" json:"inference_id"`
	ModelAddress                 string               `bson:"model_address" json:"model_address"`
	ModelID                      string               `bson:"model_id" json:"model_id"`
	WorkerAddress                string               `bson:"worker_address" json:"worker_address"`
	InferenceInput               string               `bson:"inference_input" json:"inference_input"`
	InferenceInputBytes          int64                `bson:"inference_input_bytes" json:"inference_input_bytes"`
	RayClusterOutputBytes        int64                `bson:"ray_cluster_output_bytes" json:"ray_cluster_output_bytes"`
	ModelType                    string               `bson:"model_type" json:"model_type"`
	CID                          string               `bson:"cid" json:"cid"`
	ResultLink                   string               `bson:"result_link" json:"result_link"` // link to download result for all model: TEXT AND IMAGE
	Nonce                        uint64               `bson:"nonce" json:"nonce"`
	ChainID                      string               `bson:"chain_id" json:"chain_id"`
	GasPrice                     string               `bson:"gas_price" json:"gas_price"`
	TxHash                       string               `bson:"tx_hash" json:"tx_hash"`
	TxInscribeHash               string               `bson:"tx_inscribe_hash" json:"tx_inscribe_hash"`
	Err                          string               `bson:"error" json:"error"`
	Status                       WorkerProcessStatus  `bson:"status" json:"status"`
	StorageType                  eaimodel.StorageType `bson:"storage_type" json:"storage_type"`
	Reward                       string               `bson:"reward" json:"reward"`
	IsClaimed                    bool                 `bson:"is_claimed" json:"is_claimed"`
	ProcessingExpired            *time.Time           `bson:"processing_expired" json:"processing_expired"`
	ExpiredAt                    *time.Time           `bson:"expired_at" json:"expired_at"` // expired_at from contract
	ZkSync                       bool                 `bson:"zk_sync" json:"zk_sync"`
	Data                         []byte               `bson:"data" json:"data"`
	AssignmentRole               AssignmentRoleType   `bson:"assignment_role" json:"assignment_role"`
	SeizeMinerRoleTxHash         string               `bson:"seize_miner_role_tx_hash" json:"seize_miner_role_tx_hash"`
	SeizeMinerRoleTxInscribeHash string               `bson:"seize_miner_role_tx_inscribe_hash" json:"seize_miner_role_tx_inscribe_hash"`
	CommitTxHash                 string               `bson:"commit_tx_hash" json:"commit_tx_hash"`
	CommitTxInscribeHash         string               `bson:"commit_tx_inscribe_hash" json:"commit_tx_inscribe_hash"`
	RevealTxHash                 string               `bson:"reveal_tx_hash" json:"reveal_tx_hash"`
	RevealTxInscribeHash         string               `bson:"reveal_tx_inscribe_hash" json:"reveal_tx_inscribe_hash"`
	ResolveTxHash                string               `bson:"resolve_tx_hash" json:"resolve_tx_hash"`
	ExecuteTaskDone              bool                 `bson:"execute_task_done" json:"execute_task_done"`
	UpdateExecuteTaskDoneAt      time.Time            `bson:"update_execute_task_done_at" json:"update_execute_task_done_at"`
	OriginInferenceID            string               `bson:"origin_inference_id" json:"origin_inference_id"` // inferID need scoring
	IsAgentInfer                 bool                 `bson:"is_agent_infer" json:"is_agent_infer"`
	BatchInfers                  []*BatchInferHistory `bson:"batch_infers" json:"batch_infers"`
	StoreRawFlag                 bool                 `bson:"store_raw_flag" json:"store_raw_flag"`
	SeizeMinerRoleAgain          bool                 `bson:"seize_miner_role_again" json:"seize_miner_role_again"`
	UpdateSubmitTimeoutAt        *time.Time           `bson:"update_submit_timeout_at" json:"update_submit_timeout_at"`
	WorkerHubType                WorkerHubType        `bson:"worker_hub_type" json:"worker_hub_type"`
}

type AssignmentRoleType int

const (
	AssignmentRoleValidator AssignmentRoleType = 0
	AssignmentRoleMiner     AssignmentRoleType = 1
)

type WorkerHubType int

const (
	WorkerHubTypeNormal        WorkerHubType = 0
	WorkerHubTypeKnowledgeBase WorkerHubType = 1
)

func (ModelWorkerProcessHistories) CollectionName() string {
	return "model_worker_process_histories"
}
