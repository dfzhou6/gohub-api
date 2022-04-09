package cache

import (
	"gohub/pkg/config"
	"gohub/pkg/redis"
	"time"
)

type RedisStore struct {
	RedisClient *redis.RedisClient
	KerPrefix   string
}

var _ Store = (*RedisStore)(nil)

func NewRedisStore(address string, username string, password string, db int) *RedisStore {
	rs := &RedisStore{}
	rs.RedisClient = redis.NewClient(address, username, password, db)
	rs.KerPrefix = config.GetString("app.name") + ":cache:"
	return rs
}

func (rs *RedisStore) Set(key string, value string, expire time.Duration) {
	rs.RedisClient.Set(rs.KerPrefix+key, value, expire)
}

func (rs *RedisStore) Get(key string) string {
	return rs.RedisClient.Get(rs.KerPrefix + key)
}

func (rs *RedisStore) Has(key string) bool {
	return rs.RedisClient.Has(rs.KerPrefix + key)
}

func (rs *RedisStore) Forget(key string) {
	rs.RedisClient.Del(rs.KerPrefix + key)
}

func (rs *RedisStore) Forever(key string, value string) {
	rs.RedisClient.Set(rs.KerPrefix+key, value, 0)
}

func (rs *RedisStore) Flush() {
	rs.RedisClient.FlushDB()
}

func (rs *RedisStore) Increment(params ...interface{}) {
	rs.RedisClient.Increment(params...)
}

func (rs *RedisStore) Decrement(params ...interface{}) {
	rs.RedisClient.Decrement(params...)
}

func (rs *RedisStore) IsAlive() error {
	return rs.RedisClient.Ping()
}
