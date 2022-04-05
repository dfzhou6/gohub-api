package verifycode

import (
	"gohub/pkg/app"
	"gohub/pkg/config"
	"gohub/pkg/redis"
	"time"
)

type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

func (rs *RedisStore) Set(key string, value string) bool {
	Expire := time.Minute * time.Duration(config.GetInt64("verifycode.expire_time"))
	if app.IsLocal() {
		Expire = time.Minute * time.Duration(config.GetInt64("verifycode.debug_expire_time"))
	}
	return rs.RedisClient.Set(rs.KeyPrefix+key, value, Expire)
}

func (rs *RedisStore) Get(key string, clear bool) string {
	key = rs.KeyPrefix + key
	val := rs.RedisClient.Get(key)
	if clear {
		rs.RedisClient.Del(key)
	}
	return val
}

func (rs *RedisStore) Verify(key, ans string, clear bool) bool {
	val := rs.Get(key, clear)
	return val == ans
}
