package initializers

import (
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectToRedis() {
	redisHost := os.Getenv("REDIS_HOST")
	redisPass := os.Getenv("REDIS_PASSWORD")

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPass, // no password set
		DB:       0,         // use default DB
	})
}
