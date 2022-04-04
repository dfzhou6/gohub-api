package redis

import (
	"context"
	"gohub/pkg/logger"
	"sync"
	"time"

	redis "github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client  *redis.Client
	Context context.Context
}

var once sync.Once

var Redis *RedisClient

func ConnectRedis(address string, username string, password string, db int) {
	once.Do(func() {
		Redis = NewClient(address, username, password, db)
	})
}

func NewClient(address string, username string, password string, db int) *RedisClient {
	rds := &RedisClient{}
	rds.Context = context.Background()
	rds.Client = redis.NewClient(&redis.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       db,
	})

	err := rds.Ping()
	logger.LogIf(err)

	return rds
}

func (r RedisClient) Ping() error {
	_, err := r.Client.Ping(r.Context).Result()
	return err
}

func (r RedisClient) Set(key string, value interface{}, expire time.Duration) bool {
	if err := r.Client.Set(r.Context, key, value, expire).Err(); err != nil {
		logger.ErrorString("Redis", "Set", err.Error())
		return false
	}
	return true
}

func (r RedisClient) Get(key string) string {
	result, err := r.Client.Get(r.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Get", err.Error())
		}
		return ""
	}
	return result
}

func (r RedisClient) Has(key string) bool {
	_, err := r.Client.Get(r.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Has", err.Error())
		}
		return false
	}
	return true
}

func (r RedisClient) Del(keys ...string) bool {
	if err := r.Client.Del(r.Context, keys...).Err(); err != nil {
		logger.ErrorString("Redis", "Del", err.Error())
		return false
	}
	return true
}

func (r RedisClient) FlushDB() bool {
	if err := r.Client.FlushAll(r.Context).Err(); err != nil {
		logger.ErrorString("Redis", "FlushDB", err.Error())
		return false
	}
	return true
}

func (r RedisClient) Increment(params ...interface{}) bool {
	switch len(params) {
	case 1:
		key := params[0].(string)
		if err := r.Client.Incr(r.Context, key).Err(); err != nil {
			logger.ErrorString("Redis", "Incr", err.Error())
			return false
		}
	case 2:
		key := params[0].(string)
		value := params[1].(int64)
		if err := r.Client.IncrBy(r.Context, key, value).Err(); err != nil {
			logger.ErrorString("Redis", "IncrBy", err.Error())
			return false
		}
	default:
		logger.ErrorString("Redis", "Incr", "参数过多")
		return false
	}

	return true
}

func (r RedisClient) Decrement(params ...interface{}) bool {
	switch len(params) {
	case 1:
		key := params[0].(string)
		if err := r.Client.Decr(r.Context, key).Err(); err != nil {
			logger.ErrorString("Redis", "Decr", err.Error())
			return false
		}
	case 2:
		key := params[0].(string)
		val := params[1].(int64)
		if err := r.Client.DecrBy(r.Context, key, val).Err(); err != nil {
			logger.ErrorString("Redis", "DecrBy", err.Error())
			return false
		}
	default:
		logger.ErrorString("Reids", "Decr", "参数过多")
		return false
	}

	return true
}
