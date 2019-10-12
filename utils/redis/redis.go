package redis

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"strconv"
	"unname/utils/conf"
	"unname/utils/log"
)

var rds *redis.Client

func init() {
	addr := conf.Getenv("REDIS_ADDR", "localhost:6379")
	password := conf.Getenv("REDIS_PASSWORD", "")
	DB, _ := strconv.Atoi(conf.Getenv("GUARD_REDIS_INDEX", "5"))
	rds = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       DB,
		PoolSize: 5,
	})

	pong, err := rds.Ping().Result()
	if err != nil {
		log.Logger.Error("REDIS", zap.String("e", err.Error()))
	} else {
		log.Logger.Info("REDIS", zap.String("url", addr), zap.String("ping", pong))
	}
}

func HashAdd(member, key, value string) {
	rds.HSet(member, key, value)
}

func HashExist(member, key string) bool {
	_, err := rds.HGet(member, key).Result()
	if err == nil {
		log.Logger.Warn("redis exist")
		return true
	}
	return false
}

func HashGet(member, key string) string {
	res, err := rds.HGet(member, key).Result()
	if err != nil {
		log.Logger.Warn("redis hashGet failed",
			zap.String("err", err.Error()))
		return ""
	}
	return res
}
