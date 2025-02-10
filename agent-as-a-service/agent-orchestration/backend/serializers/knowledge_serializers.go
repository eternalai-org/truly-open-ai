package serializers

import (
	"time"
)

type CreateKnowledgeRequest struct {
	Name                 string  `json:"name" form:"name"`
	Description          string  `json:"description" form:"description"`
	NetworkID            uint64  `json:"network_id" form:"network_id"`
	AgentInfoId          uint    `json:"agent_info_id" form:"-"`
	Files                []*File `json:"files" form:"files"`
	UserAddress          string  `json:"user_address" form:"-"`
	DepositAddress       string  `json:"-" form:"-"`
	ThumbnailUrl         string  `json:"thumbnail_url" form:"thumbnail_url"`
	SolanaDepositAddress string  `json:"-" form:"-"`
}

type UpdateKnowledgeRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	NetworkID   uint64 `json:"network_id"`
	UserAddress string `json:"user_address" form:"-"`
}

type File struct {
	Url  string `json:"url" form:"url"`
	Name string `json:"name" form:"name"`
	Size uint   `json:"size" form:"size"`
}

type KnowledgeBase struct {
	ID                     uint                 `json:"id"`
	CreatedAt              time.Time            `json:"created_at"`
	UpdatedAt              time.Time            `json:"updated_at"`
	Status                 int64                `json:"status"`
	UserAddress            string               `json:"user_address"`
	DepositAddress         string               `json:"deposit_address"`
	SolanaDepositAddress   string               `json:"solana_deposit_address"`
	DepositTxHash          string               `json:"deposit_tx_hash"`
	Name                   string               `json:"name"`
	Description            string               `json:"description"`
	AgentInfoId            uint                 `json:"agent_info_id"`
	ResultUrl              string               `json:"result_url"`
	NetworkID              uint64               `json:"network_id"`
	Fee                    float64              `json:"fee"`
	KnowledgeBaseFiles     []*KnowledgeBaseFile `json:"knowledge_base_files"`
	KBTokenContractAddress string               `json:"kb_token_contract_address"`
	KBTokenID              string               `json:"kb_token_id"`
	KbId                   string               `json:"kb_id"`
	ThumbnailUrl           string               `json:"thumbnail_url"`
	LastErrorMessage       string               `json:"last_error_message"`
	UsageFee               float64              `json:"usage_fee"`
	UserCount              int64                `json:"user_count"`
	UsageCount             int64                `json:"usage_count"`
}

type KnowledgeBaseFile struct {
	KnowledgeBaseId uint   `json:"knowledge_base_id"`
	FileUrl         string `json:"file_url"`
	FileName        string `json:"name"`
	FileSize        uint   `json:"size"`
}

type AgentUseKnowledgeBaseRequest struct {
	AgentID         string `json:"agent_id" form:"agent_id"`
	KnowledgeBaseID uint   `json:"knowledge_base_id" form:"knowledge_base_id"`
	UserAddress     string `json:"-"`
}

type RetrieveKnowledgeBaseRequest struct {
	Query     string   `json:"query"`
	TopK      int      `json:"top_k"`
	Kb        []string `json:"kb"`
	Threshold float64  `json:"threshold"`
}

type RetrieveKnowledgeBaseResponse struct {
	Result []struct {
		Content   string  `json:"content"`
		Score     float64 `json:"score"`
		Reference string  `json:"reference"`
	} `json:"result"`
	Error  interface{} `json:"error"`
	Status string      `json:"status"`
}

type UpdateKnowledgeBaseWithSignatureRequest struct {
	KnowledgeId uint   `json:"knowledge_id"`
	NetworkID   string `json:"network_id"`
	// sysPrompt []byte, promptKey string, promptIdx *big.Int, randomNonce *big.Int, signature []byte
	HashData        string `json:"hash_data"`
	PromptKeyData   string `json:"prompt_key_data"`
	RandomNonceData string `json:"random_nonce_data"`
	SignatureData   string `json:"signature_data"`

	//uint256 agentId, string uri, uint256 randomNonce, bytes signature
	/*HashUri        string `json:"hash_uri"`
	RandomNonceUri string `json:"random_nonce_uri"`
	SignatureUri   string `json:"signature_uri"`*/
}

type RetrieveKnowledgeRequest struct {
	Prompt    string  `json:"prompt"`
	KbId      string  `json:"kb_id"`
	TopK      int     `json:"top_k"`
	Threshold float64 `json:"threshold"`
}
