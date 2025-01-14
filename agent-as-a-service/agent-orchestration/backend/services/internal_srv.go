package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/rapid"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/twitter"
)

func (s *Service) GetTwitterUserByID(ctx context.Context, twitterID string) (*twitter.UserObj, error) {
	twitterUser, err := s.twitterWrapAPI.GetTwitterByID(twitterID)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return twitterUser, nil
}

func (s *Service) GetTwitterUserByUsername(ctx context.Context, username string) (*twitter.UserObj, error) {
	twitterUser, err := s.twitterWrapAPI.GetTwitterByUserName(username)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return twitterUser, nil
}

func (s *Service) SeachTwitterUserByQuery(ctx context.Context, username string) (*twitter.UserLookups, error) {
	usernameArry := strings.Split(username, ",")
	twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"twitter_id = ?": {s.conf.TokenTwiterIdForInternal},
		},
		map[string][]interface{}{},
		false,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}

	if twitterInfo != nil {
		twitterUser, err := s.twitterWrapAPI.LookupUsername(twitterInfo.AccessToken, usernameArry)
		if err != nil {
			return nil, errs.NewTwitterError(err)
		}
		return twitterUser, nil
	}

	return nil, errs.NewError(errs.ErrBadRequest)
}

func (s *Service) GetTwitterUserFollowing(ctx context.Context, twitterID, paginationToken string) (*twitter.UserFollowLookup, error) {
	twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"twitter_id = ?": {s.conf.TokenTwiterIdForInternal},
		},
		map[string][]interface{}{},
		false,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}

	if twitterInfo != nil {
		twitterUser, err := s.twitterWrapAPI.GetListFollowing(twitterID, paginationToken, twitterInfo.AccessToken)
		if err != nil {
			return nil, errs.NewTwitterError(err)
		}
		return twitterUser, nil
	}
	return nil, errs.NewError(errs.ErrBadRequest)
}

func (s *Service) GetTwitterUserFollowingRapid(ctx context.Context, twitterID, paginationToken string) ([]rapid.Following, error) {
	twitterUser, err := s.rapid.GetTwitterFollowings(twitterID)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return twitterUser, nil
}

func (s *Service) GetTwitterUserFollowingV1(ctx context.Context, twitterID, paginationToken string) ([]rapid.Following, error) {
	twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"twitter_id = ?": {s.conf.TokenTwiterIdForInternal},
		},
		map[string][]interface{}{},
		false,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}

	var twitterUser []rapid.Following
	if twitterInfo != nil {
		twitterUserFollowing, err := s.twitterWrapAPI.GetListFollowing(twitterID, paginationToken, twitterInfo.AccessToken)
		if err != nil {
			return nil, errs.NewTwitterError(err)
		}
		for _, item := range twitterUserFollowing.Lookups {
			twitterUser = append(twitterUser, rapid.Following{
				ID:       item.User.ID,
				Username: item.User.UserName,
				Name:     item.User.Name,
			})
		}
	}
	return nil, errs.NewError(errs.ErrBadRequest)
}

func (s *Service) GetTwitterUserFollowingRapidByUsername(ctx context.Context, username, paginationToken string) ([]rapid.Following, error) {
	if username == "" {
		return nil, errs.NewError(errs.ErrBadRequest)
	}
	var twitterUser []rapid.Following
	err := s.RedisCached(
		fmt.Sprintf("GetTwitterUserFollowingRapidByUsername_%s", username),
		true,
		3*time.Minute,
		&twitterUser,
		func() (interface{}, error) {
			user, err := s.SyncGetTwitterUserByUsername(ctx, username)
			if err != nil {
				return nil, errs.NewError(err)
			}

			if user == nil {
				return nil, errs.NewError(errs.ErrBadRequest)
			}

			twitterUser, err = s.rapid.GetTwitterFollowings(user.TwitterID)
			if err != nil {
				return nil, errs.NewError(err)
			}
			return twitterUser, nil
		},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return twitterUser, nil
}

func (s *Service) GetTwitterUserFollowingByUsername(ctx context.Context, username, paginationToken string) ([]rapid.Following, error) {
	if username == "" {
		return nil, errs.NewError(errs.ErrBadRequest)
	}
	var twitterUser []rapid.Following
	err := s.RedisCached(
		fmt.Sprintf("GetTwitterUserFollowingByUsername_%s", username),
		true,
		3*time.Minute,
		&twitterUser,
		func() (interface{}, error) {
			user, err := s.SyncGetTwitterUserByUsername(ctx, username)
			if err != nil {
				return nil, errs.NewError(err)
			}

			if user == nil {
				return nil, errs.NewError(errs.ErrBadRequest)
			}
			twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"twitter_id = ?": {s.conf.TokenTwiterIdForInternal},
				},
				map[string][]interface{}{},
				false,
			)
			if err != nil {
				return nil, errs.NewError(err)
			}

			if twitterInfo != nil {
				twitterUserFollowing, err := s.twitterWrapAPI.GetListFollowing(user.TwitterID, paginationToken, twitterInfo.AccessToken)
				if err != nil {
					return nil, errs.NewTwitterError(err)
				}
				for _, item := range twitterUserFollowing.Lookups {
					twitterUser = append(twitterUser, rapid.Following{
						ID:       item.User.ID,
						Username: item.User.UserName,
						Name:     item.User.Name,
					})
				}
			}

			// twitterUser, err = s.rapid.GetTwitterFollowings(user.TwitterID)
			// if err != nil {
			// 	return nil, errs.NewError(err)
			// }
			return twitterUser, nil
		},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return twitterUser, nil
}

func (s *Service) GetListUserTweets(ctx context.Context, twitterID, paginationToken string, maxResults int) (*twitter.UserTimeline, error) {
	twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"twitter_id = ?": {s.conf.TokenTwiterIdForInternal},
		},
		map[string][]interface{}{},
		false,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}

	if twitterInfo != nil {
		return s.GetListUserTweetsFromTwitterInfoToken(ctx, twitterInfo, twitterID, paginationToken, maxResults)
	}
	return nil, errs.NewError(errs.ErrBadRequest)
}

func (s *Service) GetListUserTweetsAll(ctx context.Context, twitterID, paginationToken string, maxResults int) (*twitter.UserTimeline, error) {
	twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"twitter_id = ?": {s.conf.TokenTwiterIdForInternal},
		},
		map[string][]interface{}{},
		false,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}

	if twitterInfo != nil {
		return s.GetAllUserTweetsFromTwitterInfoToken(ctx, twitterInfo, twitterID, paginationToken, maxResults)
	}
	return nil, errs.NewError(errs.ErrBadRequest)
}

func (s *Service) GetListUserTweetsV1(ctx context.Context, twitterID, paginationToken string, maxResults int) (*twitter.UserTimeline, error) {
	twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"twitter_id = ?": {s.conf.TokenTwiterIdForInternal},
		},
		map[string][]interface{}{},
		false,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}

	if twitterInfo != nil {
		return s.GetListUserTweetsFromTwitterInfoTokenV1(ctx, twitterInfo, twitterID, paginationToken, maxResults)
	}
	return nil, errs.NewError(errs.ErrBadRequest)
}

func (s *Service) GetListUserTweetsFromTwitterInfoTokenV1(ctx context.Context, twitterInfo *models.TwitterInfo, twitterID, paginationToken string, maxResults int) (*twitter.UserTimeline, error) {

	if twitterInfo != nil {
		twitterUser, err := s.twitterWrapAPI.GetListUserTweets(twitterID, paginationToken, twitterInfo.AccessToken, maxResults)
		if err != nil {
			return nil, errs.NewTwitterError(err)
		}
		return twitterUser, nil
	}
	return nil, errs.NewError(errs.ErrBadRequest)
}

func (s *Service) GetListUserTweetsFromTwitterInfoToken(ctx context.Context, twitterInfo *models.TwitterInfo, twitterID, paginationToken string, maxResults int) (*twitter.UserTimeline, error) {

	if twitterInfo != nil {
		twitterUser, err := s.twitterWrapAPI.GetListUserTweets(twitterID, paginationToken, twitterInfo.AccessToken, maxResults)
		if err != nil {
			return nil, errs.NewTwitterError(err)
		}
		return twitterUser, nil
	}
	return nil, errs.NewError(errs.ErrBadRequest)
}

func (s *Service) GetAllUserTweetsFromTwitterInfoToken(ctx context.Context, twitterInfo *models.TwitterInfo, twitterID, paginationToken string, maxResults int) (*twitter.UserTimeline, error) {

	if twitterInfo != nil {
		twitterUser, err := s.twitterWrapAPI.GetAllUserTweets(twitterID, paginationToken, twitterInfo.AccessToken, maxResults)
		if err != nil {
			return nil, errs.NewTwitterError(err)
		}
		return twitterUser, nil
	}
	return nil, errs.NewError(errs.ErrBadRequest)
}

func (s *Service) GetListUserTweetsByUsername(ctx context.Context, username, paginationToken string, maxResults int) (*twitter.UserTimeline, error) {
	if username == "" {
		return nil, errs.NewError(errs.ErrTwitterUsernameNotFound)
	}
	var twitterUser *twitter.UserTimeline
	// var twitterUser *twitter.UserTimelineV1
	err := s.RedisCached(
		fmt.Sprintf("GetListUserTweetsByUsername_%s_%s", username, paginationToken),
		true,
		1*time.Minute,
		&twitterUser,
		func() (interface{}, error) {
			twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"twitter_id = ?": {s.conf.TokenTwiterIdForInternal},
				},
				map[string][]interface{}{},
				false,
			)
			if err != nil {
				return nil, errs.NewError(err)
			}

			if twitterInfo != nil {
				user, err := s.SyncGetTwitterUserByUsername(ctx, username)
				if err != nil {
					return nil, errs.NewError(err)
				}

				if user == nil || (user != nil && user.TwitterID == "") {
					return nil, errs.NewError(errs.ErrTwitterUsernameNotFound)
				}

				twitterUser, err = s.twitterWrapAPI.GetListUserTweets(user.TwitterID, paginationToken, twitterInfo.AccessToken, maxResults)
				if err != nil {
					return nil, errs.NewTwitterError(err)
				}
				return twitterUser, nil
			}
			return nil, errs.NewError(errs.ErrBadRequest)
		},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return twitterUser, nil
}

func (s *Service) GetListUserTweetsByUsernameV1(ctx context.Context, username, paginationToken string, maxResults int) (*twitter.UserTimeline, error) {
	if username == "" {
		return nil, errs.NewError(errs.ErrTwitterUsernameNotFound)
	}
	var twitterUser *twitter.UserTimeline
	err := s.RedisCached(
		fmt.Sprintf("GetListUserTweetsByUsernameV1_%s_%s", username, paginationToken),
		true,
		1*time.Minute,
		&twitterUser,
		func() (interface{}, error) {
			twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"twitter_id = ?": {s.conf.TokenTwiterIdForInternal},
				},
				map[string][]interface{}{},
				false,
			)
			if err != nil {
				return nil, errs.NewError(err)
			}

			if twitterInfo != nil {
				user, err := s.SyncGetTwitterUserByUsername(ctx, username)
				if err != nil {
					return nil, errs.NewError(err)
				}

				if user == nil || (user != nil && user.TwitterID == "") {
					return nil, errs.NewError(errs.ErrTwitterUsernameNotFound)
				}

				twitterUser, err = s.twitterWrapAPI.GetListUserTweets(user.TwitterID, paginationToken, twitterInfo.AccessToken, maxResults)
				if err != nil {
					return nil, errs.NewTwitterError(err)
				}
				return twitterUser, nil
			}
			return nil, errs.NewError(errs.ErrBadRequest)
		},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return twitterUser, nil
}

func (s *Service) LookupUserTweets(ctx context.Context, tweetIDs string) (*twitter.TweetLookups, error) {
	twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"twitter_id = ?": {s.conf.TokenTwiterIdForInternal},
		},
		map[string][]interface{}{},
		false,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}

	if twitterInfo != nil {
		tweetIdArr := strings.Split(tweetIDs, ",")
		twitterUser, err := s.twitterWrapAPI.LookupUserTweets(twitterInfo.AccessToken, tweetIdArr)
		if err != nil {
			return nil, errs.NewTwitterError(err)
		}
		return twitterUser, nil
	}
	return nil, errs.NewError(errs.ErrBadRequest)
}

func (s *Service) LookupUserTweetsV1(ctx context.Context, tweetIDs string) (*twitter.TweetLookups, error) {
	twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"twitter_id = ?": {s.conf.TokenTwiterIdForInternal},
		},
		map[string][]interface{}{},
		false,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}

	if twitterInfo != nil {
		tweetIdArr := strings.Split(tweetIDs, ",")
		twitterUser, err := s.twitterWrapAPI.LookupUserTweets(twitterInfo.AccessToken, tweetIdArr)
		if err != nil {
			return nil, errs.NewError(err)
		}
		return twitterUser, nil
	}
	return nil, errs.NewError(errs.ErrBadRequest)
}

func (s *Service) GetListUserMentions(ctx context.Context, twitterID, paginationToken string, maxResults int) (*twitter.UserTimeline, error) {
	twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"twitter_id = ?": {s.conf.TokenTwiterIdForInternal},
		},
		map[string][]interface{}{},
		false,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}

	if twitterInfo != nil {
		tweetUser, err := s.CreateUpdateUserTwitter(daos.GetDBMainCtx(ctx), twitterID)
		if err != nil {
			return nil, errs.NewError(err)
		}

		tweetMentions, err := s.twitterWrapAPI.GetListUserMentions(twitterID, paginationToken, twitterInfo.AccessToken, maxResults)
		if err != nil {
			return nil, errs.NewTwitterError(err)
		}
		listRealMentions := []twitter.TweetObj{}
		for _, tweets := range tweetMentions.Tweets {
			if s.TweetIsMention(tweets, tweetUser.TwitterUsername) {
				listRealMentions = append(listRealMentions, tweets)
			}
		}
		tweetMentions.Tweets = listRealMentions
		return tweetMentions, nil
	}
	return nil, errs.NewError(errs.ErrBadRequest)
}

func (s *Service) TweetIsMention(tweets twitter.TweetObj, username string) bool {
	if len(tweets.Entities.Mentions) > 0 {
		lastMentions := tweets.Entities.Mentions[len(tweets.Entities.Mentions)-1]
		if strings.EqualFold(lastMentions.UserName, username) {
			return true
		}

		if len(tweets.Entities.Mentions) > 1 {
			lastMentions = tweets.Entities.Mentions[len(tweets.Entities.Mentions)-2]
			if strings.EqualFold(lastMentions.UserName, username) {
				return true
			}
		}

		if len(tweets.Entities.Mentions) > 2 {
			lastMentions = tweets.Entities.Mentions[len(tweets.Entities.Mentions)-3]
			if strings.EqualFold(lastMentions.UserName, username) {
				return true
			}
		}
	}
	return false
}

func (s *Service) SyncGetTwitterUserByUsername(ctx context.Context, username string) (*models.TwitterUser, error) {
	tweetUser, err := s.dao.FirstTwitterUser(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"twitter_username = ?": {username},
		}, map[string][]interface{}{},
		[]string{},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}

	if tweetUser == nil {
		info, err := s.twitterWrapAPI.GetTwitterByUserName(username)
		if err != nil {
			return nil, errs.NewError(err)
		}

		if info != nil {
			if info.ID == "" {
				return nil, errs.NewError(errs.ErrTwitterUsernameNotFound)
			}

			tweetUser, err = s.CreateUpdateUserTwitter(daos.GetDBMainCtx(ctx), info.ID)
			if err != nil {
				return nil, errs.NewError(err)
			}
		}
	}
	return tweetUser, nil
}

func (s *Service) GetListUserMentionsByUsername(ctx context.Context, username, paginationToken string, maxResults int) (*twitter.UserTimeline, error) {
	var tweetMentions *twitter.UserTimeline
	err := s.RedisCached(
		fmt.Sprintf("GetListUserMentionsByUsername_%s", username),
		true,
		1*time.Minute,
		&tweetMentions,
		func() (interface{}, error) {
			twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"twitter_id = ?": {s.conf.TokenTwiterIdForInternal},
				},
				map[string][]interface{}{},
				false,
			)
			if err != nil {
				return nil, errs.NewError(err)
			}

			if twitterInfo != nil {
				user, err := s.SyncGetTwitterUserByUsername(ctx, username)
				if err != nil {
					return nil, errs.NewError(err)
				}

				if user == nil {
					return nil, errs.NewError(errs.ErrBadRequest)
				}

				if user.TwitterID == "" {
					return nil, errs.NewError(errs.ErrTwitterIdNotFound)
				}

				mapID := map[string]string{}
				listRealMentions := []twitter.TweetObj{}
				loop := 1
				paginationToken := ""
				for {
					tweetMentions, err = s.twitterWrapAPI.GetListUserMentions(user.TwitterID, paginationToken, twitterInfo.AccessToken, maxResults)
					if err != nil {
						return nil, errs.NewTwitterError(err)
					}

					for _, tweets := range tweetMentions.Tweets {
						if _, ok := mapID[tweets.ID]; !ok {
							if s.TweetIsMention(tweets, user.TwitterUsername) && !strings.EqualFold(tweets.AuthorID, user.TwitterID) {
								replied, _ := s.CacheCheckIsTweetReplied(ctx, tweets.ID)
								if !replied {
									listRealMentions = append(listRealMentions, tweets)
									mapID[tweets.ID] = "1"
								}
							}
						}

					}

					if len(listRealMentions) > 10 {
						break
					}
					paginationToken = tweetMentions.Meta.NextToken
					loop += 1
					if loop > 3 {
						break
					}
				}
				tweetMentions.Tweets = listRealMentions
				return tweetMentions, nil
			}
			return nil, errs.NewError(errs.ErrBadRequest)
		},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return tweetMentions, nil

}

func (s *Service) GetAllUserMentionsByUsername(ctx context.Context, username, paginationToken string, maxResults int) (*twitter.UserTimeline, error) {
	var tweetMentions *twitter.UserTimeline
	err := s.RedisCached(
		fmt.Sprintf("GetAllUserMentionsByUsername_%s_%s", username, paginationToken),
		true,
		1*time.Minute,
		&tweetMentions,
		func() (interface{}, error) {
			twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"twitter_id = ?": {s.conf.TokenTwiterIdForInternal},
				},
				map[string][]interface{}{},
				false,
			)
			if err != nil {
				return nil, errs.NewError(err)
			}

			if twitterInfo != nil {
				user, err := s.SyncGetTwitterUserByUsername(ctx, username)
				if err != nil {
					return nil, errs.NewError(err)
				}

				if user == nil {
					return nil, errs.NewError(errs.ErrBadRequest)
				}

				if user.TwitterID == "" {
					return nil, errs.NewError(errs.ErrTwitterIdNotFound)
				}

				tweetMentions, err = s.twitterWrapAPI.GetListUserMentions(user.TwitterID, paginationToken, twitterInfo.AccessToken, maxResults)
				if err != nil {
					return nil, errs.NewTwitterError(err)
				}
				return tweetMentions, nil
			}
			return nil, errs.NewError(errs.ErrBadRequest)
		},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return tweetMentions, nil

}

func (s *Service) CacheCheckIsTweetReplied(ctx context.Context, twitterID string) (bool, error) {
	var replied bool
	err := s.RedisCached(
		fmt.Sprintf("CacheCheckIsTweetReplied_%s", twitterID),
		true,
		24*time.Hour,
		&replied,
		func() (interface{}, error) {
			var err error
			replied, err = s.dao.IsTweetReplied(daos.GetDBMainCtx(ctx), twitterID)
			if err != nil {
				return false, errs.NewError(err)
			}
			return replied, nil
		},
	)
	if err != nil {
		return false, errs.NewError(err)
	}
	return replied, nil
}

func (s *Service) SearchRecentTweet(ctx context.Context, query, paginationToken string, maxResults int) (*twitter.TweetRecentSearch, error) {
	var tweetRecentSearch twitter.TweetRecentSearch
	var lookUps map[string]twitter.TweetLookup
	var meta twitter.TweetRecentSearchMeta
	cacheKey := fmt.Sprintf(`CacheAgentTerminalLatestLookUps_%s_%d`, query, maxResults)
	cacheKey1 := fmt.Sprintf(`CacheAgentTerminalLatestMeta_%s_%d`, query, maxResults)
	err := s.GetRedisCachedWithKey(cacheKey, &lookUps)
	if err != nil {
		twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
			map[string][]interface{}{
				"twitter_id = ?": {s.conf.TokenTwiterIdForInternal},
			},
			map[string][]interface{}{},
			false,
		)
		if err != nil {
			return nil, errs.NewError(err)
		}

		if twitterInfo != nil {
			// query = url.QueryEscape(query)
			tweetRecentSearch, err := s.twitterWrapAPI.SearchRecentTweet(query, paginationToken, twitterInfo.AccessToken, maxResults)
			if err != nil {
				return nil, errs.NewTwitterError(err)
			}
			lookUps = tweetRecentSearch.LookUps
			meta = tweetRecentSearch.Meta
			_ = s.SetRedisCachedWithKey(cacheKey, tweetRecentSearch.LookUps, 5*time.Minute)
			_ = s.SetRedisCachedWithKey(cacheKey1, tweetRecentSearch.Meta, 5*time.Minute)
			return tweetRecentSearch, nil
		}
	}

	_ = s.GetRedisCachedWithKey(cacheKey1, &meta)

	tweetRecentSearch.LookUps = lookUps
	tweetRecentSearch.Meta = meta
	return &tweetRecentSearch, nil
}

func (s *Service) SearchRecentTweetV1(ctx context.Context, query, sinceID string, maxResults int) (*twitter.TweetRecentSearch, error) {
	var tweetRecentSearch twitter.TweetRecentSearch
	var lookUps map[string]twitter.TweetLookup
	var meta twitter.TweetRecentSearchMeta
	cacheKey := fmt.Sprintf(`CacheAgentTerminalLatestLookUps_%s_%d`, query, maxResults)
	cacheKey1 := fmt.Sprintf(`CacheAgentTerminalLatestMeta_%s_%d`, query, maxResults)
	err := s.GetRedisCachedWithKey(cacheKey, &lookUps)
	if err != nil {
		twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
			map[string][]interface{}{
				"twitter_id = ?": {s.conf.TokenTwiterIdForInternal},
			},
			map[string][]interface{}{},
			false,
		)
		if err != nil {
			return nil, errs.NewError(err)
		}

		if twitterInfo != nil {
			// query = url.QueryEscape(query)
			tweetRecentSearch, err := s.twitterWrapAPI.SearchRecentTweetV1(query, sinceID, twitterInfo.AccessToken, maxResults)
			if err != nil {
				return nil, errs.NewTwitterError(err)
			}
			lookUps = tweetRecentSearch.LookUps
			meta = tweetRecentSearch.Meta
			_ = s.SetRedisCachedWithKey(cacheKey, tweetRecentSearch.LookUps, 5*time.Minute)
			_ = s.SetRedisCachedWithKey(cacheKey1, tweetRecentSearch.Meta, 5*time.Minute)
			return tweetRecentSearch, nil
		}
	}

	_ = s.GetRedisCachedWithKey(cacheKey1, &meta)

	tweetRecentSearch.LookUps = lookUps
	tweetRecentSearch.Meta = meta
	return &tweetRecentSearch, nil
}

func (s *Service) SearchTokenTweet(ctx context.Context, query, paginationToken string, maxResults int) (*twitter.TweetRecentSearch, error) {
	var tweetRecentSearch twitter.TweetRecentSearch
	var lookUps map[string]twitter.TweetLookup
	var meta twitter.TweetRecentSearchMeta
	cacheKey := fmt.Sprintf(`CacheSearchTokenTweetLatestLookUps_%s`, query)
	cacheKey1 := fmt.Sprintf(`CacheSearchTokenTweetLatestMeta_%s`, query)
	err := s.GetRedisCachedWithKey(cacheKey, &lookUps)
	if err != nil {
		twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
			map[string][]interface{}{
				"twitter_id = ?": {s.conf.TokenTwiterIdForInternal},
			},
			map[string][]interface{}{},
			false,
		)
		if err != nil {
			return nil, errs.NewError(err)
		}

		if twitterInfo != nil {
			tweetRecentSearch, err := s.twitterWrapAPI.SearchRecentTweet(query, paginationToken, twitterInfo.AccessToken, maxResults)
			if err != nil {
				return nil, errs.NewTwitterError(err)
			}
			lookUps = tweetRecentSearch.LookUps
			meta = tweetRecentSearch.Meta
			_ = s.SetRedisCachedWithKey(cacheKey, tweetRecentSearch.LookUps, 5*time.Minute)
			_ = s.SetRedisCachedWithKey(cacheKey1, tweetRecentSearch.Meta, 5*time.Minute)
			return tweetRecentSearch, nil
		}
	}

	_ = s.GetRedisCachedWithKey(cacheKey1, &meta)

	tweetRecentSearch.LookUps = lookUps
	tweetRecentSearch.Meta = meta
	return &tweetRecentSearch, nil
}

func (s *Service) SearchUsers(ctx context.Context, query, paginationToken string) ([]*twitter.UserObj, error) {
	var twitterUser []*twitter.UserObj
	err := s.RedisCached(
		fmt.Sprintf("SearchUsers_%s", query),
		true,
		3*time.Minute,
		&twitterUser,
		func() (interface{}, error) {
			twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"twitter_id = ?": {s.conf.TokenTwiterIdForInternal},
				},
				map[string][]interface{}{},
				false,
			)
			if err != nil {
				return nil, errs.NewError(err)
			}

			if twitterInfo != nil {
				twitterUser, err = s.twitterWrapAPI.SearchUsers(query, paginationToken, twitterInfo.AccessToken)
				if err != nil {
					return nil, errs.NewTwitterError(err)
				}
				return twitterUser, nil
			}
			return nil, errs.NewError(errs.ErrBadRequest)
		},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return twitterUser, nil

}

func (s *Service) GetUser3700Liked(ctx context.Context, replied *bool, page, limit int) ([]*models.TwitterTweetLiked, error) {
	joinFilters := map[string][]interface{}{}
	filters := map[string][]interface{}{
		"twitter_tweet_likeds.liked_user_id = ?": {"1443830739372417024"},
	}
	if replied != nil {
		if *replied {
			filters[`
				tweet_id in (
					select tweetid from agent_snapshot_post_actions aspa 
					where 1=1
					and status ='done'
					and type = 'reply'
				)
			`] = []interface{}{}
		} else {
			filters[`
				tweet_id not in (
					select tweetid from agent_snapshot_post_actions aspa 
					where 1=1
					and status ='done'
					and type = 'reply'
				)
			`] = []interface{}{}
		}

	}
	selected := []string{}

	tweets, err := s.dao.FindTwitterTweetLikedJoinSelect(
		daos.GetDBMainCtx(ctx),
		selected,
		joinFilters,
		filters,
		map[string][]interface{}{},
		[]string{"id desc"}, page, limit,
	)
	if err != nil {
		return nil, errs.NewError(err)
	}

	return tweets, nil
}

func (s *Service) GetAgentTradeTokens(ctx context.Context, networkID uint64) ([]*models.AgentTradeToken, error) {
	filters := map[string][]interface{}{
		`enabled = ?`: {true},
	}

	if networkID > 0 {
		filters[`network_id = ?`] = []interface{}{networkID}
	}
	posts, err := s.dao.FindAgentTradeToken(daos.GetDBMainCtx(ctx),
		filters,
		map[string][]interface{}{},
		[]string{}, 0, 100,
	)

	if err != nil {
		return nil, errs.NewError(err)
	}

	return posts, nil
}

func (s *Service) GetListUserTweetsByUsersForTradeMission(ctx context.Context, userTwitterIds string) ([]twitter.TweetObj, error) {
	if userTwitterIds == "" {
		return []twitter.TweetObj{}, nil
	}
	tweets := []twitter.TweetObj{}
	err := s.RedisCached(
		fmt.Sprintf("GetListUserTweetsByUsersForTradeMission_%s", userTwitterIds),
		true,
		1*time.Minute,
		&tweets,
		func() (interface{}, error) {
			twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"twitter_id = ?": {s.conf.TokenTwiterIdForInternal},
				},
				map[string][]interface{}{},
				false,
			)
			if err != nil {
				return nil, errs.NewError(err)
			}

			if twitterInfo != nil {
				listTwitterIDs := strings.Split(userTwitterIds, ",")
				for _, twitterID := range listTwitterIDs {
					twitterUser, err := s.twitterWrapAPI.GetListUserTweets(twitterID, "", twitterInfo.AccessToken, 50)
					if err != nil {
						return nil, errs.NewTwitterError(err)
					}
					if twitterUser != nil && len(twitterUser.Tweets) > 0 {
						tweets = append(tweets, twitterUser.Tweets...)
					}
				}
			}
			return tweets, nil
		},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return tweets, nil
}

func (s *Service) GetListUserTweetsByAgentForTradeMission(ctx context.Context, refID string) ([]twitter.TweetObj, error) {
	if refID == "" {
		return []twitter.TweetObj{}, nil
	}
	tweets := []twitter.TweetObj{}
	err := s.RedisCached(
		fmt.Sprintf("GetListUserTweetsByAgentForTradeMission_%s", refID),
		true,
		1*time.Minute,
		&tweets,
		func() (interface{}, error) {
			twitterInfo, err := s.dao.FirstTwitterInfo(daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"twitter_id = ?": {s.conf.TokenTwiterIdForInternal},
				},
				map[string][]interface{}{},
				false,
			)
			if err != nil {
				return nil, errs.NewError(err)
			}

			missionPost, err := s.dao.FirstAgentSnapshotPost(daos.GetDBMainCtx(ctx),
				map[string][]interface{}{
					"infer_tx_hash = ?": {refID},
				}, map[string][]interface{}{}, []string{})

			if err != nil {
				return nil, errs.NewError(err)
			}

			if missionPost != nil {
				mission, err := s.dao.FirstAgentSnapshotMission(daos.GetDBMainCtx(ctx),
					map[string][]interface{}{
						"id = ?":       {missionPost.AgentSnapshotMissionID},
						"tool_set = ?": {models.ToolsetTypeTrading},
					},
					map[string][]interface{}{}, []string{})
				if err != nil {
					return nil, errs.NewError(err)
				}

				if twitterInfo != nil && mission != nil && mission.UserTwitterIds != "" {
					listTwitterIDs := strings.Split(mission.UserTwitterIds, ",")
					for _, twitterID := range listTwitterIDs {
						twitterUser, err := s.twitterWrapAPI.GetListUserTweets(twitterID, "", twitterInfo.AccessToken, 50)
						if err != nil {
							return nil, errs.NewTwitterError(err)
						}
						if twitterUser != nil && len(twitterUser.Tweets) > 0 {
							tweets = append(tweets, twitterUser.Tweets...)
						}
					}
				}
			}

			return tweets, nil
		},
	)
	if err != nil {
		return nil, errs.NewError(err)
	}
	return tweets, nil
}
