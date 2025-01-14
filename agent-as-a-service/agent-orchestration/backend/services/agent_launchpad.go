package services

import (
	"context"
	"fmt"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/jinzhu/gorm"
)

func (s *Service) AgentTwitterPostCreateLaunchpad(ctx context.Context, twitterPostID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("AgentTwitterPostCreateAgent_%d", twitterPostID),
		func() error {
			err := daos.WithTransaction(
				daos.GetDBMainCtx(ctx),
				func(tx *gorm.DB) error {
					twitterPost, err := s.dao.FirstAgentTwitterPostByID(
						tx,
						twitterPostID,
						map[string][]interface{}{
							"AgentInfo":             {},
							"AgentInfo.TwitterInfo": {},
						},
						false,
					)
					if err != nil {
						return errs.NewError(err)
					}
					if twitterPost.Status == models.AgentTwitterPostStatusNew &&
						twitterPost.PostType == models.AgentSnapshotPostActionTypeCreateLaunchpad {
						twitterPost.Status = models.AgentTwitterPostStatusReplied
						err = s.dao.Save(tx, twitterPost)
						if err != nil {
							return errs.NewError(err)
						}
					}
					return nil
				},
			)
			if err != nil {
				return errs.NewError(err)
			}
			return nil
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}
