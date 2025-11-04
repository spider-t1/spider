package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"spider/internal/config"
)

const (
	DefaultRDB = "default"
	SystemRDB  = "system"
)

type redisClient struct {
	client *redis.Client
}

var rdbMap map[string]*redisClient

func Init() {

	rdbMap = make(map[string]*redisClient)
	rdbs := config.Cfg.Redis
	for _, r := range rdbs {
		rdb := &redisClient{}
		// 初始化redis
		rdb.client = redis.NewClient(&redis.Options{
			Addr:           fmt.Sprintf("%s:%s", r.Host, r.Port),
			Password:       r.Password, // no password set
			DB:             r.DB,       // use default DB
			MaxIdleConns:   r.MaxIdle,  // 连接池最大连接数
			MaxActiveConns: r.MaxActive,
		})
		_, err := rdb.client.Ping(context.Background()).Result()
		if err != nil {
			panic(err)
		}
		rdbMap[r.DBName] = rdb
	}
}

func GetClient(rdbNames ...string) *redisClient {
	if len(rdbNames) == 0 {
		return rdbMap[DefaultRDB]
	}
	return rdbMap[rdbNames[0]]
}
