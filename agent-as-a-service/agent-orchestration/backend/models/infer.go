package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

const (
	MillisecondText = "millisecond"
	SecondText      = "second"
	MinuteText      = "minute"
	HourText        = "hour"
)

type BatchInferHistoryStatus string
type ContractInferenceStatus int
type ModelPredictHistoryStatus int
type ModelPredictHistoryVerifyStatus string
type TrainingRequestExportStatus string
type TrainingRequestStatus string
type TrainingRequestType int

const (
	TrainingRequestTypeUserDefine TrainingRequestType = iota + 1
	TrainingRequestTypeHuggingFace
)
const (
	BatchInferHistoryStatusPending       BatchInferHistoryStatus = "pending"
	BatchInferHistoryStatusAgentInferred BatchInferHistoryStatus = "agent-inferred"
	BatchInferHistoryStatusQueueHandled  BatchInferHistoryStatus = "queue-handled"
	BatchInferHistoryStatusCompleted     BatchInferHistoryStatus = "completed"
	BatchInferHistoryStatusFailed        BatchInferHistoryStatus = "failed"
)

const (
	ContractInferenceStatusNil         ContractInferenceStatus = 0
	ContractInferenceStatusSolving     ContractInferenceStatus = 1
	ContractInferenceStatusCommit      ContractInferenceStatus = 2
	ContractInferenceStatusReveal      ContractInferenceStatus = 3
	ContractInferenceStatusProcessed   ContractInferenceStatus = 4
	ContractInferenceStatusKilled      ContractInferenceStatus = 5
	ContractInferenceStatusTransferred ContractInferenceStatus = 6
)

const (
	ModelPredictHistoryStatusSolving   ModelPredictHistoryStatus = 0
	ModelPredictHistoryStatusProcessed ModelPredictHistoryStatus = 1
	// zk
	ModelPredictHistoryStatusCommit ModelPredictHistoryStatus = 2
	ModelPredictHistoryStatusReveal ModelPredictHistoryStatus = 3
)

const (
	ModelPredictHistoryVerifyStatusPending ModelPredictHistoryVerifyStatus = "pending"
	ModelPredictHistoryVerifyStatusPassed  ModelPredictHistoryVerifyStatus = "passed"
	ModelPredictHistoryVerifyStatusFailed  ModelPredictHistoryVerifyStatus = "failed"
)

const (
	TrainingRequestExportStatusWaiting   TrainingRequestExportStatus = "waiting"
	TrainingRequestExportStatusRunning   TrainingRequestExportStatus = "running"
	TrainingRequestExportStatusCompleted TrainingRequestExportStatus = "completed"
	TrainingRequestExportStatusFailed    TrainingRequestExportStatus = "failed"
)

const (
	TrainingRequestStatusWaiting          TrainingRequestStatus = "waiting"
	TrainingRequestStatusRunning          TrainingRequestStatus = "running"
	TrainingRequestStatusCompleted        TrainingRequestStatus = "completed"
	TrainingRequestStatusFailed           TrainingRequestStatus = "failed"
	TrainingRequestStatusWaitMint         TrainingRequestStatus = "wait-mint"
	TrainingRequestStatusChecking         TrainingRequestStatus = "checking"
	TrainingRequestStatusDeploying        TrainingRequestStatus = "deploying"
	TrainingRequestStatusDeployed         TrainingRequestStatus = "deployed"
	TrainingRequestStatusWaitingToPayment TrainingRequestStatus = "wait-payment"
	TrainingRequestStatusPaymentSuccess   TrainingRequestStatus = "payment-success"
)

type BatchInferHistory struct {
	gorm.Model
	UserID                          string                  `json:"user_id"`
	UserAddress                     string                  `json:"user_address"`
	Toolset                         ToolsetType             `json:"toolset"`
	AgentContractAddress            string                  `json:"agent_contract_address"`
	ContractAgentID                 string                  `json:"contract_agent_id"`
	ChainID                         uint64                  `json:"chain_id"`
	AssistantID                     string                  `json:"assistant_id"`
	PromptInput                     string                  `json:"prompt_input"`
	SystemPrompt                    string                  `json:"system_prompt"`
	PromptInputHash                 string                  `json:"prompt_input_hash"`
	AgentType                       string                  `json:"agent_type"`
	TwitterSnapshot                 string                  `json:"twitter_snapshot"`
	UserInfoSnapshot                string                  `json:"user_info_snapshot"`
	OutputMaxCharacter              uint                    `json:"output_max_character"`
	PromptOutput                    string                  `json:"prompt_output"`
	PromptOutputHash                string                  `json:"prompt_output_hash"`
	SubmitPromptOutputAt            time.Time               `json:"submit_prompt_output_at"`
	Status                          BatchInferHistoryStatus `json:"status"`
	Log                             string                  `json:"log"`
	InferID                         string                  `json:"infer_id"`
	ModelID                         string                  `json:"model_id"`
	InscribeTxHash                  string                  `json:"inscribe_tx_hash"`
	SubmitSolutionInscribeTxHash    string                  `json:"submit_solution_tx_hash"`
	BtcInscribeTxHash               string                  `json:"btc_inscribe_tx_hash"`
	BtcSubmitSolutionInscribeTxHash string                  `json:"btc_submit_solution_inscribe_tx_hash"`
	InferWalletAddress              string                  `json:"infer_wallet_address"`
	SubmitInferAt                   time.Time               `json:"submit_infer_at"`
	AssignmentAddresses             string                  `json:"assignment_addresses" gorm:"type:json"` //[]string
	SubmitSolutionAddress           string                  `json:"submit_solution_address"`
	CommitTxHash                    string                  `json:"commit_tx_hash" gorm:"type:json"`              //[]string
	RevealTxHash                    string                  `json:"reveal_tx_hash" gorm:"type:json"`              //[]string
	BtcCommitInscribeTxHash         string                  `json:"btc_commit_inscribe_tx_hash" gorm:"type:json"` //[]string
	BtcRevealInscribeTxHash         string                  `json:"btc_reveal_inscribe_tx_hash" gorm:"type:json"` //[]string
}
type ChainConfig struct {
	gorm.Model
	ChainID              string             `json:"chain_id" gorm:"column:chain_id"`
	RPC                  string             `json:"rpc" gorm:"column:rpc"`
	Name                 string             `json:"name" gorm:"column:name"`
	Explorer             string             `json:"explorer" gorm:"column:explorer"`
	EAIErc20             string             `json:"eai_erc20" gorm:"column:eai_erc20"`
	NFTAddress           string             `json:"nft_address" gorm:"column:nft_address"` // or collect address
	PaymasterAddress     string             `json:"paymaster_address" gorm:"column:paymaster_address"`
	PaymasterFeeZero     bool               `json:"paymaster_fee_zero" gorm:"column:paymaster_fee_zero"`
	PaymasterToken       string             `json:"paymaster_token" gorm:"column:paymaster_token"`
	WorkerhubAddress     string             `json:"workerhub_address" gorm:"column:workerhub_address"`
	StakinghubAddress    string             `json:"stakinghub_address" gorm:"column:stakinghub_address"`
	ZkSync               bool               `json:"zk_sync" gorm:"column:zk_sync"`
	EAINative            bool               `json:"eai_native" gorm:"column:eai_native"`
	Thumbnail            string             `json:"thumbnail" gorm:"column:thumbnail"`
	DaoToken             string             `json:"dao_token" gorm:"column:dao_token"`
	ModelIds             string             `json:"model_ids" gorm:"column:model_ids;type:json"` //[]string
	ModelDetails         []*TrainingRequest `json:"model_details" gorm:"-"`
	AgentContractAddress string             `json:"agent_contract_address" gorm:"column:agent_contract_address"`
	NeedInscribeTx       bool               `json:"need_inscribe_tx" gorm:"column:need_inscribe_tx"`
	WebStatUri           string             `json:"web_stat_uri" gorm:"column:web_stat_uri"` // e.g., https://webstat.shard-ai.l2aas.com
	BlockNumberDelay     uint64             `json:"block_number_delay" gorm:"column:block_number_delay"`
	ListRPC              string             `json:"-" gorm:"column:list_rpc;type:json"` //[]string
	GasLimit             uint64             `json:"gas_limit" gorm:"column:gas_limit"`
	SupportModelNames    string             `json:"support_model_names" gorm:"column:support_model_names;type:json"` //map[string]string  Serialize as JSON
	SupportStoreRaw      bool               `json:"support_store_raw" gorm:"column:support_store_raw"`
	IsWorkerHubV4        bool               `json:"is_workerhub_v4" gorm:"column:is_workerhub_v4"`
}

type TrainingRequestERC20Info struct {
	gorm.Model
	ERC20Address         string    `json:"erc20_address" gorm:"column:erc20_address"`
	ERC20IssueTime       time.Time `json:"erc20_issue_time" gorm:"column:erc20_issue_time"`
	ERC20Description     string    `json:"erc20_description" gorm:"column:erc20_description"`
	ERC20Name            string    `json:"erc20_name" gorm:"column:erc20_name"`
	ERC20Symbol          string    `json:"erc20_symbol" gorm:"column:erc20_symbol"`
	ERC20Supply          string    `json:"erc20_supply" gorm:"column:erc20_supply"`
	TokenPrice           string    `json:"token_price" gorm:"column:token_price"`
	TokenPriceNumber     float64   `json:"token_price_number" gorm:"column:token_price_number"`
	TokenMarketCap       string    `json:"token_market_cap" gorm:"column:token_market_cap"`
	TokenMarketCapNumber float64   `json:"token_market_cap_number" gorm:"column:token_market_cap_number"`
	CirculatingSupply    string    `json:"circulating_supply" gorm:"column:circulating_supply"`
	CirculatingSupplyNum float64   `json:"circulating_supply_num" gorm:"column:circulating_supply_num"`
	CurrentDonate        string    `json:"current_donate" gorm:"column:current_donate"`
	TotalInvestment      string    `json:"total_investment" gorm:"-"`
	TotalBackers         int       `json:"total_backers" gorm:"-"`
	IsERC20Deployed      bool      `json:"is_erc20_deployed" gorm:"column:is_erc20_deployed"`
	ERC20Owner           string    `json:"erc20_owner" gorm:"column:erc20_owner"`
}

type HGModelCheckResult struct {
	EstimatedSize   float64 `json:"estimated_size"` // GB
	MinRequirements struct {
		VRAM string `json:"VRAM"`
		RAM  string `json:"RAM"`
	} `json:"min_requirements"`
}

type ZkSyncNetwork struct {
	gorm.Model
	RPC                  string `json:"rpc" gorm:"column:rpc"`
	ChainId              string `json:"chain_id" gorm:"column:chain_id"`
	WorkerHubAddress     string `json:"worker_hub_address" gorm:"column:worker_hub_address"`
	PaymasterToken       string `json:"paymaster_token" gorm:"column:paymaster_token"`
	PaymasterAddress     string `json:"paymaster_address" gorm:"column:paymaster_address"`
	NFTAddress           string `json:"nft_address" gorm:"column:nft_address"` // or collectionAddress
	PaymasterFeeZero     bool   `json:"paymaster_fee_zero" gorm:"column:paymaster_fee_zero"`
	DaoToken             string `json:"dao_token" gorm:"column:dao_token"`
	EAIERC20             string `json:"eaierc_20" gorm:"column:eaierc_20"`
	Name                 string `json:"name" gorm:"column:name"`
	Explorer             string `json:"explorer" gorm:"column:explorer"`
	AgentContractAddress string `json:"agent_contract_address" gorm:"column:agent_contract_address"`
	NeedInscribeTx       bool   `json:"need_inscribe_tx" gorm:"column:need_inscribe_tx"`
	StakingHubAddress    string `json:"staking_hub_address" gorm:"column:staking_hub_address"`
	BlockNumberDelay     uint64 `json:"block_number_delay" gorm:"column:block_number_delay"`
	ListRPC              string `json:"list_rpc" gorm:"type:json;column:list_rpc"` // []string
	GasLimit             uint64 `json:"gas_limit" gorm:"column:gas_limit"`
}

type TrainingRequest struct {
	gorm.Model
	ModelID                    string                `json:"model_id" gorm:"column:model_id"`
	Params                     string                `json:"params" gorm:"column:params"`
	Status                     TrainingRequestStatus `json:"status" gorm:"column:status"`
	ModelCheckingLog           string                `json:"model_checking_log" gorm:"column:model_checking_log"`
	Result                     string                `json:"result" gorm:"column:result"`
	Creator                    string                `json:"creator" gorm:"column:creator"`
	IsOnchain                  bool                  `json:"is_onchain" gorm:"column:is_onchain"`
	Description                string                `json:"description" gorm:"column:description"`
	OutputUUID                 string                `json:"output_uuid" gorm:"column:output_uuid"`
	OutputLink                 string                `json:"output_link" gorm:"column:output_link"`
	Progress                   int                   `json:"progress" gorm:"column:progress"`
	ExecutedAt                 int64                 `json:"executed_at" gorm:"column:executed_at"`
	CompletedAt                int64                 `json:"completed_at" gorm:"column:completed_at"`
	Error                      string                `json:"error" gorm:"column:error"`
	Logs                       string                `json:"logs" gorm:"column:logs"`
	ErrLogs                    string                `json:"err_logs" gorm:"column:err_logs"`
	Datasets                   string                `json:"datasets" gorm:"type:json"` // []string
	PredictNumber              int                   `json:"predict_number" gorm:"column:predict_number"`
	LastPredict                *ModelPredictHistory  `json:"last_predict" gorm:"-"`
	Category                   string                `json:"category" gorm:"column:category"`
	CurrentListing             *ModelMarket          `json:"current_listing" gorm:"-"`
	Owner                      string                `json:"owner" gorm:"column:owner"`
	LastUpdatedBlock           uint64                `json:"last_updated_block" gorm:"column:last_updated_block"`
	Thumbnail                  string                `json:"thumbnail" gorm:"column:thumbnail"`
	ClaimedDeployReward        bool                  `json:"claimed_deploy_reward" gorm:"column:claimed_deploy_reward"`
	TrainingType               string                `json:"training_type" gorm:"column:training_type"`
	TrainingRequestERC20InfoID uint
	TrainingRequestERC20Info   *TrainingRequestERC20Info
	Type                       TrainingRequestType         `json:"type" gorm:"column:type"`
	ExportStatus               TrainingRequestExportStatus `json:"export_status" gorm:"column:export_status"`
	ExportResult               string                      `json:"export_result" gorm:"column:export_result"`
	ExportLog                  string                      `json:"export_log" gorm:"column:export_log"`
	ModelMetadata              string                      `json:"model_metadata" gorm:"column:model_metadata"`
	ModelAddressUserDefine     string                      `json:"model_address_user_define" gorm:"column:model_address_user_define"`
	FileSize                   int64                       `json:"file_size" gorm:"column:file_size"`
	ChargeFee                  string                      `json:"charge_fee" gorm:"column:charge_fee"`
	ChargeFeeNumber            float64                     `json:"charge_fee_number" gorm:"column:charge_fee_number"`
	ChargeFeeTokenSymbol       string                      `json:"charge_fee_token_symbol" gorm:"column:charge_fee_token_symbol"`
	ChargeReceiveAddress       string                      `json:"charge_receive_address" gorm:"column:charge_receive_address"`
	ChargePaidTxHash           string                      `json:"charge_paid_tx_hash" gorm:"column:charge_paid_tx_hash"`
	SortOrder                  int                         `json:"sort_order" gorm:"column:sort_order"`
	HuggingFaceId              string                      `json:"hugging_face_id" gorm:"column:hugging_face_id"`
	DataTypeExporter           string                      `json:"data_type_exporter" gorm:"column:data_type_exporter"`
	CheckResult                string                      `json:"check_result" gorm:"type:json"` //json HGModelCheckResult
	ZkSync                     bool                        `json:"zk_sync" gorm:"column:zk_sync"`
	ZkSyncNetworkID            uint
	ZkSyncNetwork              *ZkSyncNetwork
	ModelStorageUrl            string `json:"model_storage_url" gorm:"column:model_storage_url"`
	Slug                       string `json:"slug" gorm:"column:slug"`
}

type ModelPredictHistory struct {
	gorm.Model
	RequestID               string                          `json:"requestId" gorm:"column:request_id"`
	ModelID                 string                          `json:"modelId" gorm:"column:model_id"`
	InferTxHash             string                          `json:"infer_tx_hash" gorm:"column:infer_tx_hash"`
	IsAgentInfer            bool                            `json:"is_agent_infer" gorm:"column:is_agent_infer"`
	Requester               string                          `json:"requester" gorm:"column:requester"`
	ClassName               string                          `json:"className" gorm:"column:class_name"`
	Status                  ModelPredictHistoryStatus       `json:"status" gorm:"column:status"` // 0: pending, 1: completed
	CdnURL                  string                          `json:"cdnURL" gorm:"column:cdn_url"`
	Result                  string                          `json:"result" gorm:"column:result"`
	Category                string                          `json:"category" gorm:"column:category"`
	Hashtag                 string                          `json:"hashtag" gorm:"column:hashtag"` // unique generate hashtag to share on Twitter
	LikeCount               int64                           `json:"likeCount" gorm:"column:like_count"`
	ViewCount               int64                           `json:"viewCount" gorm:"column:view_count"`
	ClaimedView             int64                           `json:"claimedView" gorm:"column:claimed_view"`
	ClaimedLike             int64                           `json:"claimedLike" gorm:"column:claimed_like"`
	SharedBy                string                          `json:"sharedBy" gorm:"type:json"` //    []string
	ClaimedPredictReward    bool                            `json:"claimedPredictReward" gorm:"column:claimed_predict_reward"`
	Metadata                string                          `json:"metadata" gorm:"column:metadata"`
	Description             string                          `json:"description" gorm:"column:description"`
	VerifyStatus            ModelPredictHistoryVerifyStatus `json:"verifyStatus" gorm:"column:verify_status"`
	Fee                     string                          `json:"fee" gorm:"column:fee"`
	ProcessedBys            string                          `json:"processedBys" gorm:"type:json"` //   []string
	WinnerMiner             string                          `json:"winnerMiner" gorm:"column:winner_miner"`
	WinnerFee               string                          `json:"winnerFee" gorm:"column:winner_fee"`
	WinnerTx                string                          `json:"winnerTx" gorm:"column:winner_tx"`
	ModelData               *TrainingRequest                `json:"model_data" gorm:"-"`
	RefundTx                string                          `json:"refundTx" gorm:"column:refund_tx"`
	ResolveInference        bool                            `json:"resolve_inference" gorm:"column:resolve_inference"`
	ContractInferenceStatus ContractInferenceStatus         `json:"contract_inference_status" gorm:"column:contract_inference_status"`
	SubmitTimeout           uint64                          `json:"submit_timeout" gorm:"column:submit_timeout"`
	CommitTimeout           uint64                          `json:"commit_timeout" gorm:"column:commit_timeout"`
	RevealTimeout           uint64                          `json:"reveal_timeout" gorm:"column:reveal_timeout"`
	InferenceMiner          string                          `json:"inference_miner" gorm:"column:inference_miner"`
	DaoTxTransferToUser     string                          `json:"dao_tx_transfer_to_user" gorm:"column:dao_tx_transfer_to_user"`
	DaoTxTransferToReferer  string                          `json:"dao_tx_transfer_to_referer" gorm:"column:dao_tx_transfer_to_referer"`
}

type ModelMarket struct {
	gorm.Model
	ModelID       string `json:"model_id" gorm:"column:model_id"`
	ModelAddress  string `json:"model_address" gorm:"column:model_address"`
	IsListing     bool   `json:"is_listing" gorm:"column:is_listing"`
	UserAddress   string `json:"user_address" gorm:"column:user_address"`
	Price         string `json:"price" gorm:"column:price"`
	Matched       bool   `json:"matched" gorm:"column:matched"`
	Cancelled     bool   `json:"cancelled" gorm:"column:cancelled"`
	OfferID       string `json:"offer_id" gorm:"column:offer_id"`
	Deadline      int64  `json:"deadline" gorm:"column:deadline"`
	AcceptAsset   string `json:"accept_asset" gorm:"column:accept_asset"`
	IsInvalid     bool   `json:"is_invalid" gorm:"column:is_invalid"`
	AtBlock       uint64 `json:"at_block" gorm:"column:at_block"`
	BuyerAddress  string `json:"buyer_address" gorm:"column:buyer_address"`
	SellerAddress string `json:"seller_address" gorm:"column:seller_address"`
}

type JobConfig struct {
	gorm.Model
	JobName      string     `json:"job_name"`
	Enable       bool       `json:"enable"`
	Interval     int64      `json:"interval"`
	LastRun      *time.Time `json:"last_run"`
	IntervalUnit string     `json:"interval_unit"`
}
