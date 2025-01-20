package services

import (
	"context"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/orderpayment"
)

func (s *Service) OrderpaymentOrderPaidEvent(ctx context.Context, networkID uint64, event *orderpayment.OrderpaymentOrderPaid) error {
	return nil
}
