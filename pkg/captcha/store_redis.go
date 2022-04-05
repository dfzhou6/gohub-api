package captcha

import (
	"errors"
	"gohub/pkg/app"
	"gohub/pkg/config"
	"gohub/pkg/redis"
	"time"
)

type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

func (s *RedisStore) Set(key string, val string) error {
	Expire := time.Minute * time.Duration(config.GetInt64("captcha.expire_time"))
	if app.IsLocal() {
		Expire = time.Minute * time.Duration(config.GetInt64("captcha.debug_expire_time"))
	}

	if ok := s.RedisClient.Set(s.KeyPrefix+key, val, Expire); !ok {
		return errors.New("无法存储验证码")
	}
	return nil
}

func (s *RedisStore) Get(key string, clear bool) string {
	key = s.KeyPrefix + key
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return val
}

func (s *RedisStore) Verify(key, ans string, clear bool) bool {
	val := s.Get(key, clear)
	return val == ans
}
