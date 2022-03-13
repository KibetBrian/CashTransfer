package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func Connect() context.Context {
	ctx := context.Background()
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalln(err, "Error pinging redis")
	}
	fmt.Println(pong, "Redis connected...")
	return ctx
}

func SetRefreshToken(key string, value string, time time.Duration) (bool, error) {
	ctx := Connect()
	err := client.Set(ctx, key, value, time).Err()
	if err != nil {
		return false, fmt.Errorf("failed to  set value. err: %v", err)
	}
	return true, nil
}

func GetRefreshToken(k string) (string, error) {
	ctx := Connect()
	s, err := client.Get(ctx, k).Result()
	if err == nil {
		return "", fmt.Errorf("failed to  get value of key err: %v", err)
	}
	return s, nil
}
