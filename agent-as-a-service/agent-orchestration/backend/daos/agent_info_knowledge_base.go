package daos

import (
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/jinzhu/gorm"
)

func (dao *DAO) FirstAgentInfoKnowledgeBaseByAgentInfoID(tx *gorm.DB, agentInfoID uint, preload map[string][]interface{}, orders []string) (*models.AgentInfoKnowledgeBase, error) {
	var agentInfoKnowledgeBase models.AgentInfoKnowledgeBase
	if err := dao.first(tx, &agentInfoKnowledgeBase, map[string][]interface{}{"agent_info_id = ?": []interface{}{agentInfoID}},
		preload, orders, false); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &agentInfoKnowledgeBase, nil
}

func (dao *DAO) FirstAgentInfoKnowledgeBaseByAgentInfoIDAndKnowledgeBaseID(tx *gorm.DB, agentInfoID uint,
	knowledgeBaseID uint) (*models.AgentInfoKnowledgeBase, error) {
	var agentInfoKnowledgeBase models.AgentInfoKnowledgeBase
	if err := dao.first(tx, &agentInfoKnowledgeBase, map[string][]interface{}{
		"agent_info_id = ?":     {agentInfoID},
		"knowledge_base_id = ?": {knowledgeBaseID},
	}, nil, nil, false); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &agentInfoKnowledgeBase, nil
}

func (dao *DAO) CreateAgentInfoKnowledgeBase(tx *gorm.DB, agentInfoKnowledgeBase *models.AgentInfoKnowledgeBase) error {
	return dao.Create(tx, agentInfoKnowledgeBase)
}

// ========== KNOWLEDGE BASE ==========
func (dao *DAO) FirstKnowledgeBase(tx *gorm.DB, filters map[string][]interface{}, preload map[string][]interface{}, order []string, forUpdate bool) (*models.KnowledgeBase, error) {
	var knowledgeBase models.KnowledgeBase
	if err := dao.first(tx, &knowledgeBase, filters, preload, order, forUpdate); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &knowledgeBase, nil
}
