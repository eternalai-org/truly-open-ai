package daos

import (
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/jinzhu/gorm"
)

func (dao *DAO) FirstAgentInfoKnowledgeBaseByAgentInfoID(tx *gorm.DB, agentInfoID uint) (*models.AgentInfoKnowledgeBase, error) {
	var agentInfoKnowledgeBase models.AgentInfoKnowledgeBase
	if err := dao.first(tx, &agentInfoKnowledgeBase, map[string][]interface{}{"agent_info_id = ?": []interface{}{agentInfoID}}, nil, nil, false); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &agentInfoKnowledgeBase, nil
}
