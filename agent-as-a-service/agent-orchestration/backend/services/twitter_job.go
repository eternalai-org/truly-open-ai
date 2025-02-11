package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/twitter"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

func (s *Service) ScanTwitterTweetLiked(ctx context.Context) error {
	twitterInfos, err := s.dao.FindTwitterInfo(daos.GetDBMainCtx(ctx), map[string][]interface{}{
		"twitter_id in ('1443830739372417024', '95760717') ": {}}, map[string][]interface{}{}, []string{}, 0, 100)
	if err != nil {
		return err
	}
	for _, agent := range twitterInfos {
		likedTweets, err := helpers.GetLikedTweets(agent.TwitterID, agent.AccessToken, "")
		if err != nil {
			return err
		}
		tweetIDs := []string{}
		for _, v := range likedTweets.Data {
			tweetIDs = append(tweetIDs, v.ID)
		}
		listTweets, err := s.twitterAPI.GetTweetDetails(tweetIDs)
		if err != nil {
			return err
		}
		daos.WithTransaction(
			daos.GetDBMainCtx(ctx),
			func(tx *gorm.DB) error {
				for _, item := range listTweets {
					tweet, err := s.dao.FirstTwitterTweetLiked(tx, map[string][]interface{}{
						"tweet_id = ?":      {item.ID},
						"liked_user_id = ?": {agent.TwitterID},
					}, map[string][]interface{}{}, []string{})
					if err != nil {
						continue
					}
					if tweet != nil {
						//update
						tweet.LikedUserID = agent.TwitterID
						tweet.LikeCount = item.LikeCount
						tweet.RetweetCount = item.RetweetCount
						tweet.ReplyCount = item.ReplyCount
						tweet.LikeCount = item.LikeCount
						tweet.QuoteCount = item.QuoteCount
						tweet.ImpressionCount = item.ImpressionCount

						tweet.InReplyToUserID = item.InReplyToUserID
						tweet.InReplyToTweetID = item.InReplyToTweetID
						tweet.IsReply = item.IsReply
						tweet.OriginalText = item.OriginalText

						tweet.IsRetweet = item.IsRetweet
						tweet.RepostTweetID = item.RepostTweetID
						tweet.RepostText = item.RepostText
						tweet.IsQuote = item.IsQuote
						tweet.QuoteTweetID = item.QuoteTweetID
						tweet.QuoteText = item.QuoteText

						err = s.dao.Save(tx, tweet)
						if err != nil {
							continue
						}

					} else {
						err = s.dao.Create(tx, &models.TwitterTweetLiked{
							LikedUserID:      agent.TwitterID,
							TweetID:          item.ID,
							TwitterID:        item.AuthorID,
							LikeCount:        item.LikeCount,
							RetweetCount:     item.RetweetCount,
							ReplyCount:       item.ReplyCount,
							QuoteCount:       item.QuoteCount,
							ImpressionCount:  item.ImpressionCount,
							FullText:         item.Text,
							PostedAt:         item.CreatedAt,
							InReplyToUserID:  item.InReplyToUserID,
							InReplyToTweetID: item.InReplyToTweetID,
							IsReply:          item.IsReply,
							OriginalText:     item.OriginalText,
							IsRetweet:        item.IsRetweet,
							RepostTweetID:    item.RepostTweetID,
							RepostText:       item.RepostText,
							IsQuote:          item.IsQuote,
							QuoteTweetID:     item.QuoteTweetID,
							QuoteText:        item.QuoteText,
						})
						if err != nil {
							continue
						}

					}
				}
				return nil
			},
		)
	}
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) JobScanTwitterLiked(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx, "JobScanTwitterLiked",
		func() error {
			return s.ScanTwitterTweetLiked(ctx)
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) ScanTwitterTweetByParentID(ctx context.Context, launchpad *models.Launchpad) (*twitter.TweetRecentSearch, error) {
	lst, err := s.SearchRecentTweetV1(ctx, fmt.Sprintf("in_reply_to_tweet_id:%s", launchpad.StartTweetId), launchpad.LastScanID, 50)
	if err != nil {
		return nil, errs.NewError(err)
	}
	if lst != nil {
		for _, v := range lst.LookUps {
			newMissionID := uint(0)
			daos.WithTransaction(
				daos.GetDBMainCtx(ctx),
				func(tx *gorm.DB) error {

					tmp := helpers.ParseStringToDateTimeTwitter(v.Tweet.CreatedAt)
					err = s.dao.Create(tx, &models.TwitterTweet{
						TwitterID: v.Tweet.AuthorID,
						TweetID:   v.Tweet.ID,
						FullText:  v.Tweet.Text,
						PostedAt:  *tmp,
					})
					if err != nil {
						return errs.NewError(err)
					}
					address := helpers.ExtractEVMAddress(v.Tweet.Text)
					if address != "" && address != launchpad.Address {
						//check join
						member, err := s.dao.FirstLaunchpadMember(tx, map[string][]interface{}{
							"twitter_id = ?":   {v.Tweet.AuthorID},
							"launchpad_id = ?": {launchpad.ID},
						}, map[string][]interface{}{}, []string{})
						if err != nil {
							return errs.NewError(err)
						}
						if member == nil {
							tier3 := models.MulBigFloats(&launchpad.MaxFundBalance.Float, numeric.NewFloatFromString("0.005"))
							member = &models.LaunchpadMember{
								NetworkID:      launchpad.NetworkID,
								UserAddress:    strings.ToLower(address),
								TwitterID:      v.Tweet.AuthorID,
								LaunchpadID:    launchpad.ID,
								TweetID:        v.Tweet.ID,
								TweetContent:   v.Tweet.Text,
								Tier:           models.LaunchpadTier3,
								MaxFundBalance: numeric.BigFloat{*tier3},
								Status:         models.LaunchpadMemberStatusNew,
							}
							err = s.dao.Create(tx, member)
							if err != nil {
								return errs.NewError(err)
							}
						} else {
							//one X join one time per launchpad
							return nil
						}
						toolList := fmt.Sprintf(launchpad.AgentSnapshotMission.ToolList, v.Tweet.AuthorID, s.conf.InternalApiKey, launchpad.ID, member.ID, s.conf.InternalApiKey)
						newMission := &models.AgentSnapshotMission{}
						err = copier.Copy(newMission, launchpad.AgentSnapshotMission)
						if err != nil {
							return errs.NewError(err)
						}
						newMission.ID = 0
						newMission.ToolList = toolList
						newMission.LaunchpadMemberID = member.ID
						newMission.ReactMaxSteps = 2
						err = s.dao.Create(tx, newMission)
						if err != nil {
							return errs.NewError(err)
						}
						newMissionID = newMission.ID

					}
					return nil
				},
			)
			if newMissionID > 0 {
				s.AgentSnapshotPostCreate(ctx, newMissionID, "", "")
			}

		}
	}
	return lst, nil
}
