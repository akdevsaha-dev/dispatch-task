package storage

import "github.com/redis/go-redis/v9"

func RedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
