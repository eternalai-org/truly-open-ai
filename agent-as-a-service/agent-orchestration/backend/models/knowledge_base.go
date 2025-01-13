package models

import "gorm.io/gorm"

type KnowledgeBaseStatus int64

const (
	KnowledgeBaseStatusWaitingPayment KnowledgeBaseStatus = iota + 1
	KnowledgeBaseStatusPaymentReceipt
	KnowledgeBaseStatusProcessing
	KnowledgeBaseStatusDone
	KnowledgeBaseStatusMinted
	KnowledgeBaseStatusProcessingFailed
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
	AgentId                string               `json:"agent_id"`
	ResultUrl              string               `json:"result_url"`
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
}

type AgentInfoKnowledgeBase struct {
	gorm.Model
	AgentInfoId     uint `json:"agent_info_id" gorm:"index"`
	KnowledgeBaseId uint `json:"knowledge_base_id" gorm:"index"`

	AgentInfo     *AgentInfo
	KnowledgeBase *KnowledgeBase
}

type KnowledgeBaseFile struct {
	gorm.Model
	KnowledgeBaseId uint   `json:"knowledge_base_id"`
	FileUrl         string `json:"file_url"`
	FileName        string `json:"name"`
	FileSize        uint   `json:"size"`
}

type ListKnowledgeBaseRequest struct {
	UserAddress string `json:"user_address" form:"-"`
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
