package ports

import (
	"context"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
)

type IKnowledgeUsecase interface {
	GetKnowledgeBaseById(context.Context, uint) (*models.KnowledgeBase, error)
	DeleteKnowledgeBaseById(context.Context, uint) error
	CreateKnowledgeBase(context.Context, *serializers.CreateKnowledgeRequest) (*serializers.KnowledgeBase, error)
	ListKnowledgeBase(context.Context, *models.ListKnowledgeBaseRequest) ([]*serializers.KnowledgeBase, error)
	WatchWalletChange(context.Context) error
	GetKnowledgeBaseByStatus(ctx context.Context, status models.KnowledgeBaseStatus, offset, limit int) ([]*models.KnowledgeBase, error)
	UpdateKnowledgeBaseById(ctx context.Context, id uint, updatedFields map[string]interface{}) error
	Webhook(context.Context, *models.RagResponse) (*models.KnowledgeBase, error)
}
