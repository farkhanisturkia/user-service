package redis

import (
	"context"
	"fmt"
	"learn-microservices/user-service/config"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client
var Ctx = context.Background()

func Init() {
	host := config.GetEnv("REDIS_HOST", "")
	port := config.GetEnv("REDIS_PORT", "")
	password := config.GetEnv("REDIS_PASS", "")
	db := 0

	Client = redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
		Password:     password,
		DB:           db,
		PoolSize:     10,
		MinIdleConns: 5,
	})

	_, err := Client.Ping(Ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("redis connection failed: %v", err))
	}

	fmt.Println("Redis connected")
}
