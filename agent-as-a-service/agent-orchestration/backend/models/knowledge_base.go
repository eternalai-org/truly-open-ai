package models

import "gorm.io/gorm"

type (
	KnowledgeBaseStatus     int64
	KnowledgeBaseFileStatus int64
)

const (
	KnowledgeBaseStatusWaitingPayment KnowledgeBaseStatus = iota + 1
	KnowledgeBaseStatusPaymentReceipt
	KnowledgeBaseStatusProcessing
	KnowledgeBaseStatusDone
	KnowledgeBaseStatusMinted
	KnowledgeBaseStatusProcessingFailed
)

const (
	KnowledgeBaseFileStatusPending KnowledgeBaseFileStatus = iota + 1
	KnowledgeBaseFileStatusDone
)

type KnowledgeBase struct {
	gorm.Model
	Status                 KnowledgeBaseStatus  `json:"status"`
	UserAddress            string               `json:"user_address"`
	DepositAddress         string               `json:"deposit_address"`
	DepositPrivKey         string               `json:"-"`
	DepositTxHash          string               `json:"deposit_tx_hash"`
	Name                   string               `json:"name"`
	Description            string               `json:"description"`
	AgentInfoId            string               `json:"agent_info_id"`
	AgentInfo              *AgentInfo           `json:"agent_info" gorm:"foreignkey:AgentInfoId;references:ID"`
	NetworkID              uint64               `json:"network_id"`
	KBTokenContractAddress string               `json:"kb_token_contract_address"`
	KBTokenID              string               `json:"kb_token_id" gorm:"index"`
	KBTokenMintTx          string               `json:"kb_token_mint_tx" gorm:"index"`
	KnowledgeBaseFiles     []*KnowledgeBaseFile `json:"knowledge_base_files"`
	Fee                    float64              `json:"fee"`
	SolanaDepositAddress   string               `json:"solana_deposit_address"`
	SolanaDepositPrivKey   string               `json:"-"`
	FilecoinHash           string               `json:"filecoin_hash"`
	DepositChainId         uint64               `json:"deposit_chain_id"`
	LastErrorMessage       string               `json:"last_error_message"`
	CurrentGroupFileId     int64                `json:"current_group_file_id"`
	KbId                   string               `json:"kb_id"`
	ThumbnailUrl           string               `json:"thumbnail_url"`
}

type KnowledgeBaseFile struct {
	gorm.Model
	KnowledgeBaseId uint                    `json:"knowledge_base_id"`
	FileUrl         string                  `json:"file_url"`
	FileName        string                  `json:"name"`
	FileSize        uint                    `json:"size"`
	GroupFileId     int64                   `json:"group_file_id"`
	Status          KnowledgeBaseFileStatus `json:"status"`
	FilecoinHash    string                  `json:"filecoin_hash"`
}

type ListKnowledgeBaseRequest struct {
	UserAddress string `json:"user_address" form:"-"`
	AgentIds    []uint `json:"agent_ids"`
}

func (m *KnowledgeBase) FileUrls() []string {
	urls := []string{}
	for _, f := range m.KnowledgeBaseFiles {
		urls = append(urls, f.FileUrl)
	}
	return urls
}

type RagResult struct {
	Ref          string `json:"ref"`
	Kb           string `json:"kb"`
	FilecoinHash string `json:"filecoin_hash"`
	Message      string `json:"message"`
}

type RagResponse struct {
	Result *RagResult `json:"result"`
	Error  *string    `json:"error"`
	Status string     `json:"status"`
}

type InsertRagResponse struct {
	Result string  `json:"result"`
	Error  *string `json:"error"`
	Status string  `json:"status"`
}
