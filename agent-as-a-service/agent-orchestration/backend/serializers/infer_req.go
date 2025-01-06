package serializers

import "github.com/sashabaranov/go-openai"

type KnowledgeBaseInfoType string

var KnowledgeBaseInfoTypePublic KnowledgeBaseInfoType = "public"
var KnowledgeBaseInfoTypePrivate KnowledgeBaseInfoType = "private"

type KnowledgeBaseInfo struct {
	Type            KnowledgeBaseInfoType `json:"type"`
	Title           string                `json:"title"`
	FileDescription string                `json:"file_description"`
	Creator         string                `json:"creator"`
}

type WakeupRequestMetadata struct {
	TwitterId         string               `json:"twitter_id"`
	TwitterUsername   string               `json:"twitter_username"`
	AgentContractId   string               `json:"agent_contract_id"`
	ChainId           string               `json:"chain_id"`
	SystemReminder    string               `json:"system_reminder"`
	Params            ParamWakeupRequest   `json:"params"`
	RefID             string               `json:"ref_id"`
	ListKnowledgeBase []*KnowledgeBaseInfo `json:"list_knowledge_base"`
}

type ParamWakeupRequest struct {
	QuoteUsername string `json:"quote_username"`
}

type ChatCompletionRequest struct {
	openai.ChatCompletionRequest `json:",inline"`
	ChainId                      uint64                 `json:"chain_id"`
	UserAddress                  string                 `json:"user_address"`
	InternalServer               bool                   `json:"internal_server"`
	MetaData                     *WakeupRequestMetadata `json:"meta_data"`
}
