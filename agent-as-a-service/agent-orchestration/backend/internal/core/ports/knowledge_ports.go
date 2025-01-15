package ports

import (
	"context"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/jinzhu/gorm"
)

type IKnowledgeUsecase interface {
	CreateAgentInfoKnowledgeBase(context.Context, *models.AgentInfoKnowledgeBase) (*models.AgentInfoKnowledgeBase, error)
	GetAgentInfoKnowledgeBaseByAgentId(context.Context, uint) (*models.AgentInfoKnowledgeBase, error)
	GetKnowledgeBaseById(context.Context, uint) (*models.KnowledgeBase, error)
	DeleteKnowledgeBaseById(context.Context, uint) error
	CreateKnowledgeBase(context.Context, *serializers.CreateKnowledgeRequest) (*serializers.KnowledgeBase, error)
	ListKnowledgeBase(context.Context, *models.ListKnowledgeBaseRequest) ([]*serializers.KnowledgeBase, error)
	WatchWalletChange(context.Context) error
	GetKnowledgeBaseByStatus(ctx context.Context, status models.KnowledgeBaseStatus, offset, limit int) ([]*models.KnowledgeBase, error)
	UpdateKnowledgeBaseById(ctx context.Context, id uint, updatedFields map[string]interface{}) error
	Webhook(context.Context, *models.RagResponse) (*models.KnowledgeBase, error)
	WebhookFile(context.Context, string, []byte, uint) (*models.KnowledgeBase, error)
	MapKnowledgeBaseByAgentIds(ctx context.Context, ids []uint) (map[uint]*models.KnowledgeBase, error)
	GetKnowledgeBaseByKBId(context.Context, string) (*models.KnowledgeBase, error)
	GetKBAgentsUsedOfSocialAgent(tx *gorm.DB, socialAgentId uint) ([]*models.KnowledgeBase, error)
	GetManyKnowledgeBaseByQuery(context.Context, string, string, int, int) ([]*models.KnowledgeBase, error)
}
