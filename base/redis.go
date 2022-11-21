package base

import (
	"time"

	"github.com/go-redis/redis/v9"
)

var Redis *redis.Client

func RedisInit() {
	if Redis != nil { // 接続済みのコネクションはcloseする
		Redis.Close()
	}
	// Redisの初期化
	Redis = redis.NewClient(&redis.Options{
		Addr:            GetEnv("REDIS_URL", "redis:6379"),
		Password:        GetEnv("REDIS_PASSWORD", ""),
		DB:              GetEnvInt("REDIS_DB", 0),
		MaxIdleConns:    GetEnvInt("REDIS_MAXCONN", 3),
		MinIdleConns:    GetEnvInt("REDIS_MINCONN", 0),
		ConnMaxIdleTime: 240 * time.Second,
		PoolSize:        GetEnvInt("REDIS_POOLSIZE", 3),
		PoolTimeout:     240 * time.Second,
	})
}
