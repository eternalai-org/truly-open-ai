package models

import "gorm.io/gorm"

type AgentInfoKnowledgeBase struct {
	gorm.Model
	AgentInfoId     uint `json:"agent_info_id" gorm:"index"`
	KnowledgeBaseId uint `json:"knowledge_base_id" gorm:"index"`

	AgentInfo     *AgentInfo
	KnowledgeBase *KnowledgeBase
}
