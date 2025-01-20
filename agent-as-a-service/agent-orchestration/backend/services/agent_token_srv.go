package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/jinzhu/gorm"
)

func (s *Service) GetListMemes(ctx context.Context, address string, page, limit int) ([]*models.Meme, uint, error) {
	filters := map[string][]interface{}{}
	if address != "" {
		filters[`memes.owner_address = ?`] = []interface{}{strings.ToLower(address)}
	}

	joinFilters := map[string][]interface{}{}

	selected := []string{
		`memes.*`,
	}

	keys, count, err := s.dao.FindMemeJoinSelect4Page(daos.GetDBMainCtx(ctx),
		selected,
		joinFilters,
		filters,
		map[string][]interface{}{},
		[]string{}, page, limit,
	)

	if err != nil {
		return nil, count, errs.NewError(err)
	}

	return keys, count, nil
}

func (s *Service) CreateMemeThread(ctx context.Context, address string, req *serializers.MemeThreadReq) (*models.MemeThreads, error) {
	var thread *models.MemeThreads
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			owner, err := s.GetUser(tx, models.BASE_CHAIN_ID, address, false)
			if err != nil {
				return errs.NewError(err)
			}
			thread = &models.MemeThreads{
				UserID:         owner.ID,
				MemeID:         req.MemeID,
				Text:           req.Text,
				ImageUrl:       req.ImageUrl,
				ParentThreadID: req.ParentThreadID,
			}
			err = s.dao.Create(tx, thread)
			if err != nil {
				return errs.NewError(err)
			}
			meme, _ := s.dao.FirstMemeByID(tx, req.MemeID, map[string][]interface{}{}, true)
			meme.ReplyCount = meme.ReplyCount + 1
			meme.LastReply = helpers.TimeNow()
			err = s.dao.Save(tx, meme)
			if err != nil {
				return errs.NewError(err)
			}
			if req.ParentThreadID > 0 {
				parentThread, err := s.dao.FirstMemeThreadByID(tx, req.ParentThreadID, map[string][]interface{}{}, false)
				if err != nil {
					return errs.NewError(err)
				}
				if parentThread != nil {
					parentUser, err := s.dao.FirstUserByID(tx, parentThread.UserID, map[string][]interface{}{}, true)
					if err != nil {
						return errs.NewError(err)
					}
					if parentUser != nil {
						parentUser.Mentions = parentUser.Mentions + 1
						updateFields := map[string]interface{}{
							"mentions": parentUser.Mentions,
						}
						err = tx.Model(parentUser).Updates(
							updateFields,
						).Error
						if err != nil {
							return errs.NewError(err)
						}
					}
				}
			}
			return nil
		},
	)

	if err != nil {
		return nil, errs.NewError(err)
	}

	//cache latest
	if thread != nil {
		go func() error {
			meme, err := s.dao.FirstMemeByID(daos.GetDBMainCtx(ctx), req.MemeID, map[string][]interface{}{}, false)
			if err != nil {
				return errs.NewError(err)
			}

			if meme != nil {
				_ = s.CacheListMemeThreadLatest(daos.GetDBMainCtx(ctx), meme.TokenAddress)
			}
			return nil
		}()
	}

	return thread, nil
}

func (s *Service) LikeMemeThread(ctx context.Context, address string, req *serializers.MemeThreadReq) (bool, error) {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			owner, err := s.GetUser(tx, models.BASE_CHAIN_ID, address, false)
			if err != nil {
				return errs.NewError(err)
			}

			threadLike := &models.MemeThreadLike{
				UserID:   owner.ID,
				ThreadID: req.ThreadID,
			}

			err = s.dao.Create(tx, threadLike)
			if err != nil {
				return errs.NewError(err)
			}

			thread, _ := s.dao.FirstMemeThreadByID(tx, threadLike.ThreadID, map[string][]interface{}{}, true)
			thread.Likes = thread.Likes + 1
			updateFields := map[string]interface{}{
				"likes": thread.Likes,
			}
			err = tx.Model(thread).Updates(
				updateFields,
			).Error
			if err != nil {
				return errs.NewError(err)
			}

			threadOwner, err := s.dao.FirstUserByID(tx, thread.UserID, map[string][]interface{}{}, true)
			if err != nil {
				return errs.NewError(err)
			}

			if threadOwner != nil {
				threadOwner.Likes = threadOwner.Likes + 1

				updateFields = map[string]interface{}{
					"likes": threadOwner.Likes,
				}

				err = tx.Model(threadOwner).Updates(
					updateFields,
				).Error

				if err != nil {
					return errs.NewError(err)
				}
			}
			return nil
		},
	)

	if err != nil {
		return false, errs.NewError(err)
	}
	return true, nil
}

func (s *Service) UnLikeMemeThread(ctx context.Context, address string, req *serializers.MemeThreadReq) (bool, error) {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			owner, err := s.GetUser(tx, models.BASE_CHAIN_ID, address, false)
			if err != nil {
				return errs.NewError(err)
			}

			threadLike, err := s.dao.FirstMemeThreadLike(tx,
				map[string][]interface{}{
					"user_id = ?":   {owner.ID},
					"thread_id = ?": {req.ThreadID},
				},
				map[string][]interface{}{}, []string{},
			)

			if err != nil {
				return errs.NewError(err)
			}

			if threadLike != nil {
				err := s.dao.DeleteUnscoped(tx, threadLike)
				if err != nil {
					return errs.NewError(err)
				}

				thread, _ := s.dao.FirstMemeThreadByID(tx, threadLike.ThreadID, map[string][]interface{}{}, true)
				thread.Likes = thread.Likes - 1
				if thread.Likes < 0 {
					thread.Likes = 0
				}

				updateFields := map[string]interface{}{
					"likes": thread.Likes,
				}
				err = tx.Model(thread).Updates(
					updateFields,
				).Error
				if err != nil {
					return errs.NewError(err)
				}

				threadOwner, err := s.dao.FirstUserByID(tx, thread.UserID, map[string][]interface{}{}, true)
				if err != nil {
					return errs.NewError(err)
				}

				if threadOwner != nil {
					threadOwner.Likes = threadOwner.Likes - 1
					if threadOwner.Likes < 0 {
						threadOwner.Likes = 0
					}

					updateFields = map[string]interface{}{
						"likes": threadOwner.Likes,
					}
					err = tx.Model(threadOwner).Updates(
						updateFields,
					).Error
					if err != nil {
						return errs.NewError(err)
					}

				}
			}

			return nil
		},
	)

	if err != nil {
		return false, errs.NewError(err)
	}
	return true, nil
}

func (s *Service) GetListMemeThread(ctx context.Context, userAddress, tokenAddress string, page, limit int) ([]*models.MemeThreads, uint, error) {
	filters := map[string][]interface{}{
		"hidden = 0": {},
	}
	joinFilters := map[string][]interface{}{}

	selected := []string{
		"meme_threads.*",
	}

	if tokenAddress != "" {
		joinFilters[`join memes on memes.id=meme_threads.meme_id and  memes.token_address = ?`] = []interface{}{strings.ToLower(tokenAddress)}
	}

	if userAddress != "" {
		joinFilters[`join users on users.id=meme_threads.user_id`] = []interface{}{}
		joinFilters[`left join meme_thread_likes on meme_thread_likes.user_id = users.id and meme_threads.id = meme_thread_likes.thread_id`] = []interface{}{}
		selected = append(selected, "(case when meme_thread_likes.id is null then false else true end) liked")
	}

	keys, err := s.dao.FindMemeThreadJoinSelect(daos.GetDBMainCtx(ctx),
		selected,
		joinFilters,
		filters,
		map[string][]interface{}{
			"User": []interface{}{},
		},
		[]string{"created_at desc"}, page, limit,
	)

	if err != nil {
		return nil, 0, errs.NewError(err)
	}

	return keys, 0, nil
}

func (s *Service) GetListMemeThreadLatest(ctx context.Context, tokenAddress string) (string, error) {
	var resp string
	cacheKey := fmt.Sprintf(`GetListMemeThreadLatest_%s`, strings.ToLower(tokenAddress))
	err := s.GetRedisCachedWithKey(cacheKey, &resp)
	if err != nil {
		s.CacheListMemeThreadLatest(daos.GetDBMainCtx(ctx), tokenAddress)
		s.GetRedisCachedWithKey(cacheKey, &resp)
	}

	return resp, nil
}

func (s *Service) CacheListMemeThreadLatest(tx *gorm.DB, tokenAddress string) error {
	filters := map[string][]interface{}{
		"hidden = 0": {},
	}
	joinFilters := map[string][]interface{}{
		`join memes on memes.id = meme_threads.meme_id and memes.token_address = ?`: {strings.ToLower(tokenAddress)},
	}

	selected := []string{"meme_threads.*"}
	threads, err := s.dao.FindMemeThreadJoinSelect(tx,
		selected,
		joinFilters,
		filters,
		map[string][]interface{}{
			"User": []interface{}{},
		},
		[]string{"id desc"}, 1, 20,
	)

	if err != nil {
		return errs.NewError(err)
	}

	cacheData, err := json.Marshal(&serializers.Resp{Result: serializers.NewMemeThreadRespArry(threads)})
	if err != nil {
		errs.NewError(err)
	}

	err = s.SetRedisCachedWithKey(
		fmt.Sprintf(`GetListMemeThreadLatest_%s`, strings.ToLower(tokenAddress)),
		string(cacheData),
		1*time.Hour,
	)
	if err != nil {
		return errs.NewError(err)
	}

	return nil
}

// ///user

func (s *Service) FollowUsers(ctx context.Context, address string, req *serializers.MemeThreadReq) (bool, error) {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			owner, err := s.GetUser(tx, models.BASE_CHAIN_ID, address, false)
			if err != nil {
				return errs.NewError(err)
			}

			toUser, err := s.GetUser(tx, models.BASE_CHAIN_ID, req.UserAddress, false)
			if err != nil {
				return errs.NewError(err)
			}

			userFollow := &models.MemeFollowers{
				UserID:       owner.ID,
				FollowUserID: toUser.ID,
			}

			err = s.dao.Create(tx, userFollow)
			if err != nil {
				return errs.NewError(err)
			}

			toUser, _ = s.dao.FirstUserByID(tx, toUser.ID, map[string][]interface{}{}, true)
			toUser.Followers = toUser.Followers + 1
			err = s.dao.Save(tx, toUser)
			if err != nil {
				return errs.NewError(err)
			}

			owner, _ = s.dao.FirstUserByID(tx, owner.ID, map[string][]interface{}{}, true)
			owner.Following = owner.Following + 1
			err = s.dao.Save(tx, owner)
			if err != nil {
				return errs.NewError(err)
			}

			_ = s.CreateMemeNotifications(daos.GetDBMainCtx(ctx), toUser.ID, 0, owner.ID, models.NotiTypeNewFollower, fmt.Sprintf("%s_%d", models.NotiTypeNewFollower, owner.ID))
			return nil
		},
	)

	if err != nil {
		return false, errs.NewError(err)
	}
	return true, nil
}

func (s *Service) UnFollowUsers(ctx context.Context, address string, req *serializers.MemeThreadReq) (bool, error) {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			owner, err := s.GetUser(tx, models.BASE_CHAIN_ID, address, false)
			if err != nil {
				return errs.NewError(err)
			}

			toUser, err := s.GetUser(tx, models.BASE_CHAIN_ID, req.UserAddress, false)
			if err != nil {
				return errs.NewError(err)
			}

			userFollow, err := s.dao.FirstMemeFollowers(tx,
				map[string][]interface{}{
					"user_id = ?":        {owner.ID},
					"follow_user_id = ?": {toUser.ID},
				},
				map[string][]interface{}{}, []string{},
			)

			if err != nil {
				return errs.NewError(err)
			}

			if userFollow != nil {
				err := s.dao.DeleteUnscoped(tx, userFollow)
				if err != nil {
					return errs.NewError(err)
				}

				toUser, _ = s.dao.FirstUserByID(tx, toUser.ID, map[string][]interface{}{}, true)
				toUser.Followers = toUser.Followers - 1
				if toUser.Followers < 0 {
					toUser.Followers = 0
				}

				updateFields := map[string]interface{}{
					"followers": toUser.Followers,
				}
				err = tx.Model(toUser).Updates(updateFields).Error
				if err != nil {
					return errs.NewError(err)
				}

				owner, _ = s.dao.FirstUserByID(tx, owner.ID, map[string][]interface{}{}, true)
				owner.Following = owner.Following - 1
				if toUser.Following < 0 {
					toUser.Following = 0
				}

				updateFields = map[string]interface{}{
					"following": owner.Following,
				}
				err = tx.Model(owner).Updates(updateFields).Error
				if err != nil {
					return errs.NewError(err)
				}

				return nil
			}
			return nil
		},
	)

	if err != nil {
		return false, errs.NewError(err)
	}
	return true, nil
}

func (s *Service) ValidatedFollowed(ctx context.Context, address, userAddress string) (bool, error) {
	filters := map[string][]interface{}{}
	joinFilters := map[string][]interface{}{
		`join users u on u.id=meme_followers.follow_user_id and u.address = ?`: {strings.ToLower(userAddress)},
		`join users f on f.id=meme_followers.user_id and f.address = ?`:        {strings.ToLower(address)},
	}

	selected := []string{
		"meme_followers.*",
	}

	follow, err := s.dao.FirstMemeFollowersJoinSelect(daos.GetDBMainCtx(ctx),
		selected,
		joinFilters,
		filters,
		map[string][]interface{}{},
		[]string{},
	)

	if err != nil {
		return false, errs.NewError(err)
	}

	if follow != nil {
		return true, nil
	}

	return false, nil
}

func (s *Service) GetListFollowings(ctx context.Context, userAddress string, page, limit int) ([]*models.MemeFollowers, uint, error) {
	filters := map[string][]interface{}{}
	joinFilters := map[string][]interface{}{
		`join users on users.id=meme_followers.user_id and users.address = ?`: {strings.ToLower(userAddress)},
	}

	selected := []string{
		"meme_followers.*",
	}
	keys, err := s.dao.FindMemeFollowersJoinSelect(daos.GetDBMainCtx(ctx),
		selected,
		joinFilters,
		filters,
		map[string][]interface{}{
			"FollowUser": []interface{}{},
		},
		[]string{"created_at desc"}, page, limit,
	)
	if err != nil {
		return nil, 0, errs.NewError(err)
	}
	return keys, 0, nil
}

func (s *Service) CreateMemeNotifications(tx *gorm.DB, userID, memeID, followerID uint, notiType models.NotiType, eventID string) error {
	isNew := false
	var inst *models.MemeNotification
	var err error
	inst, err = s.dao.FirstMemeNotification(tx,
		map[string][]interface{}{
			"event_id = ?": {strings.ToLower(eventID)},
		},
		map[string][]interface{}{}, []string{},
	)

	if err != nil {
		return errs.NewError(err)
	}

	if inst == nil {
		inst = &models.MemeNotification{
			EventId:    strings.ToLower(eventID),
			UserID:     userID,
			MemeID:     memeID,
			FollowerID: followerID,
			NotiType:   notiType,
		}
		err = s.dao.Create(tx, inst)
		if err != nil {
			return errs.NewError(err)
		}
		isNew = true
		return nil
	}

	if err != nil {
		return errs.NewError(err)
	}

	if isNew && userID > 0 {
		_ = s.CacheNotificationsLatest(tx, userID)
	}
	return nil
}

func (s *Service) CacheNotificationsLatest(tx *gorm.DB, userID uint) error {
	filters := map[string][]interface{}{
		"user_id = ? or user_id = 0": {userID},
	}
	joinFilters := map[string][]interface{}{}
	selected := []string{}
	noties, err := s.dao.FindMemeNotificationJoinSelect(tx,
		selected,
		joinFilters,
		filters,
		map[string][]interface{}{
			"Meme":     []interface{}{},
			"Follower": []interface{}{},
		},
		[]string{"created_at desc"}, 1, 20,
	)

	if err != nil {
		errs.NewError(err)
	}

	cacheData, err := json.Marshal(&serializers.Resp{Result: serializers.NewMemeNotificationRespArry(noties)})
	if err != nil {
		errs.NewError(err)
	}

	err = s.SetRedisCachedWithKey(
		fmt.Sprintf(`CacheNotificationsLatest_%d`, userID),
		string(cacheData),
		1*time.Hour,
	)
	if err != nil {
		return errs.NewError(err)
	}

	return nil
}

// /
func (s *Service) GetMemeNotifications(ctx context.Context, userAddress string, page, limit int) ([]*models.MemeNotification, uint, error) {
	filters := map[string][]interface{}{}
	joinFilters := map[string][]interface{}{
		`
			left join users on users.id=meme_notifications.user_id
			left join meme_notification_seens mns on mns.notification_id = meme_notifications.id and mns.user_id = users.id
		`: {},
	}

	selected := []string{
		"(case when meme_notifications.seen = 0 then ifnull(mns.seen, 0) else meme_notifications.seen end) seen",
		`meme_notifications.*`,
	}

	if userAddress != "" {
		filters["users.address = ? or meme_notifications.user_id = 0"] = []interface{}{strings.ToLower(userAddress)}
	}

	keys, err := s.dao.FindMemeNotificationJoinSelect(daos.GetDBMainCtx(ctx),
		selected,
		joinFilters,
		filters,
		map[string][]interface{}{
			"Meme":     []interface{}{},
			"Follower": []interface{}{},
		},
		[]string{"created_at desc"}, page, limit,
	)

	if err != nil {
		return nil, 0, errs.NewError(err)
	}

	return keys, 0, nil
}

func (s *Service) GetMemeNotificationLatest(ctx context.Context, userAddress string) (string, error) {
	var resp string
	user, err := s.GetUser(daos.GetDBMainCtx(ctx), models.BASE_CHAIN_ID, userAddress, false)
	if err != nil {
		return "", errs.NewError(err)
	}

	cacheKey := fmt.Sprintf(`CacheNotificationsLatest_%d`, user.ID)
	err = s.GetRedisCachedWithKey(cacheKey, &resp)
	if err != nil {
		s.CacheNotificationsLatest(daos.GetDBMainCtx(ctx), user.ID)
		s.GetRedisCachedWithKey(cacheKey, &resp)
	}

	return resp, nil
}

func (s *Service) HideMemeThread(ctx context.Context, address string, threadID uint) (bool, error) {
	inst, err := s.dao.FirstMemeWhiteListAddress(daos.GetDBMainCtx(ctx),
		map[string][]interface{}{
			"address = ?": {strings.ToLower(address)},
		},
		map[string][]interface{}{}, []string{},
	)

	if err != nil {
		return false, errs.NewError(err)
	}

	if inst == nil {
		return false, errs.NewError(errs.ErrBadRequest)
	}

	thread, _ := s.dao.FirstMemeThreadByID(daos.GetDBMainCtx(ctx), threadID,
		map[string][]interface{}{
			"Meme": {},
		}, false)
	if err != nil {
		return false, errs.NewError(err)
	}

	updateFields := map[string]interface{}{
		"hidden": true,
	}
	err = daos.GetDBMainCtx(ctx).Model(thread).Updates(
		updateFields,
	).Error

	if thread.Meme != nil {
		s.CacheListMemeThreadLatest(daos.GetDBMainCtx(ctx), thread.Meme.TokenAddress)
	}

	if err != nil {
		return false, errs.NewError(err)
	}
	return true, nil
}

func (s *Service) UserSeenMemeNotification(ctx context.Context, address string, notiID uint) (bool, error) {
	user, err := s.GetUser(daos.GetDBMainCtx(ctx), models.BASE_CHAIN_ID, address, false)
	if err != nil {
		return false, errs.NewError(err)
	}

	inst, err := s.dao.FirstMemeNotificationByID(daos.GetDBMainCtx(ctx), notiID,
		map[string][]interface{}{}, false,
	)

	if err != nil {
		return false, errs.NewError(err)
	}

	if inst == nil {
		return false, errs.NewError(errs.ErrBadRequest)
	}

	if inst.UserID > 0 {
		if inst.UserID == user.ID && !inst.Seen {
			updateFields := map[string]interface{}{
				"seen": true,
			}
			err = daos.GetDBMainCtx(ctx).Model(inst).Updates(
				updateFields,
			).Error

			if err != nil {
				return false, errs.NewError(err)
			}
		}
	} else {
		instSeen, err := s.dao.FirstMemeNotificationSeen(daos.GetDBMainCtx(ctx),
			map[string][]interface{}{
				"notification_id = ?": {inst.ID},
				"user_id = ?":         {user.ID},
			},
			map[string][]interface{}{}, []string{},
		)

		if err != nil {
			return false, errs.NewError(err)
		}

		if instSeen == nil {
			instSeen = &models.MemeNotificationSeen{
				NotificationID: inst.ID,
				UserID:         user.ID,
				Seen:           true,
			}

			err = s.dao.Create(daos.GetDBMainCtx(ctx), instSeen)
			if err != nil {
				return false, errs.NewError(err)
			}
		}
	}
	return true, nil
}

func (s *Service) GenerateMemeStory(ctx context.Context, memeName string) (string, error) {
	aiMsg := fmt.Sprintf(`
		I have a token called %s
		It's a meme token, and I need a meme story for marketing purposes.
		Please create a narrative about the token that is under 500 characters.
	`, memeName,
	)
	whitePaperStr, err := s.openais["Lama"].ChatMessage(aiMsg)
	if err != nil {
		return "", errs.NewError(err)
	}
	return whitePaperStr, nil
}

func (s *Service) CalculateFunnyWeightMeme(ctx context.Context, memeID uint) error {
	err := daos.WithTransaction(
		daos.GetDBMainCtx(ctx),
		func(tx *gorm.DB) error {
			meme, err := s.dao.FirstMemeByID(tx, memeID, map[string][]interface{}{}, false)
			if err != nil {
				return errs.NewError(err)
			}
			if meme == nil {
				return errs.NewError(errs.ErrBadRequest)
			}
			aiMsg := fmt.Sprintf(`
Name: %s
Symbol: %s
Story: %s
Number of comments: %d

This is information of a meme token, including its name, symbol, the story information and number of community discussion.
Please evaluate the humor, appeal, vibrant, and creativity of the story for each meme and return a weight from 1 to 100 corresponding to the ID this meme.
Please just return number of weight.
			`, meme.Name, meme.Ticker, meme.Description, meme.ReplyCount)
			aiStr, err := s.openais["Lama"].ChatMessage(aiMsg)
			if err != nil {
				return errs.NewError(err)
			}
			if aiStr != "" {
				memeWeight := helpers.GetNumberFromString(aiStr)
				if memeWeight > 0 {
					err = daos.GetDBMainCtx(ctx).
						Model(meme).
						Updates(
							map[string]interface{}{
								"weight": memeWeight,
							},
						).Error
					if err != nil {
						return errs.NewError(err)
					}
				}

			}
			return nil
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}
