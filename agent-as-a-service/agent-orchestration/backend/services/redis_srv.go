package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
)

func (s *Service) RedisCached(
	cachedKey string,
	enabled bool,
	expiration time.Duration,
	resp interface{},
	dataFetchedFunc func() (interface{}, error),
) error {
	cachedKey = s.hashRedisKey(fmt.Sprintf("%s_ai_agent_%s", s.conf.Redis.Prefix, cachedKey))
	if enabled {
		val, err := s.rdb.Get(cachedKey).Result()
		if err == nil && val != "" {
			err = json.Unmarshal([]byte(val), resp)
			if err != nil {
				return errs.NewError(err)
			}
			return nil
		}
	}
	//
	data, err := dataFetchedFunc()
	if err != nil {
		return errs.NewError(err)
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return errs.NewError(err)
	}
	//
	if enabled {
		err = s.rdb.Set(cachedKey, string(dataBytes), expiration).Err()
		if err != nil {
			return errs.NewError(err)
		}
	}
	//
	err = json.Unmarshal(dataBytes, resp)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) RedisCachedWithoutHashKey(
	cachedKey string,
	enabled bool,
	expiration time.Duration,
	resp interface{},
	dataFetchedFunc func() (interface{}, error),
) error {
	cachedKey = fmt.Sprintf("%s_ai_agent_%s", s.conf.Redis.Prefix, cachedKey)
	if enabled {
		val, err := s.rdb.Get(cachedKey).Result()
		if err == nil && val != "" {
			err = json.Unmarshal([]byte(val), resp)
			if err != nil {
				return errs.NewError(err)
			}
			return nil
		}
	}
	//
	data, err := dataFetchedFunc()
	if err != nil {
		return errs.NewError(err)
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return errs.NewError(err)
	}
	//
	if enabled {
		err = s.rdb.Set(cachedKey, string(dataBytes), expiration).Err()
		if err != nil {
			return errs.NewError(err)
		}
	}
	//
	err = json.Unmarshal(dataBytes, resp)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) hashRedisKey(key string) string {
	return helpers.Sha256ToNonce(key).Text(16)
}

func (s *Service) GetRedisCachedWithKey(cachedKey string, resp interface{}) error {
	cachedKey = s.hashRedisKey(fmt.Sprintf("%s_ai_agent_%s", s.conf.Redis.Prefix, cachedKey))
	val, err := s.rdb.Get(cachedKey).Result()
	if err != nil {
		return errs.NewError(err)
	}
	if val == "" {
		return errs.NewError(errs.ErrBadContent)
	}
	err = json.Unmarshal([]byte(val), resp)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) SetRedisCachedWithKey(cachedKey string, req interface{}, expiration time.Duration) error {
	cachedKey = s.hashRedisKey(fmt.Sprintf("%s_ai_agent_%s", s.conf.Redis.Prefix, cachedKey))
	dataBytes, err := json.Marshal(req)
	if err != nil {
		return errs.NewError(err)
	}
	err = s.rdb.Set(cachedKey, string(dataBytes), expiration).Err()
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) DeleteRedisCachedWithKey(cachedKey string) error {
	cachedKey = s.hashRedisKey(fmt.Sprintf("%s_ai_agent_%s", s.conf.Redis.Prefix, cachedKey))
	err := s.rdb.Del(cachedKey).Err()
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) DeleteRedisCachedWithPrefix(cachedKey string) error {
	iter := s.rdb.Scan(0, fmt.Sprintf("%s_ai_agent_%s*", s.conf.Redis.Prefix, cachedKey), 0).Iterator()
	for iter.Next() {
		err := s.rdb.Del(iter.Val()).Err()
		if err != nil {
			return errs.NewError(err)
		}
	}
	return nil
}

func (s *Service) RedisFlushAll(ctx context.Context) error {
	err := s.rdb.FlushAll().Err()
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}
