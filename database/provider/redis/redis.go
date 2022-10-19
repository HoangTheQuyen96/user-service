package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
)

func CreateRedisClient(address string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: address,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Sprintln("Redis connection failed, error: [%v]", err)
		panic(err)
	}

	if pong == "PONG" {
		fmt.Println("Redis connected.")
	}

	return client
}
