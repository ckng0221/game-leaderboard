package utils

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func IncrementUserScore(client *redis.Client, user_id uint, score int) {
	ctx := context.Background()

	err := client.ZAddArgsIncr(ctx, "leaderboard", redis.ZAddArgs{Members: []redis.Z{{Score: float64(score), Member: user_id}}}).Err()

	if err != nil {
		log.Println("error adding score")
		return
	}
	log.Printf("Added score: %v tor UserID: %v.\n", score, user_id)
}

func GetTopNLeaderboard(client *redis.Client, n int) ([]redis.Z, error) {
	ctx := context.Background()

	results, err := client.ZRevRangeWithScores(ctx, "leaderboard", 0, int64(n-1)).Result()
	if err != nil {
		return nil, errors.New("error retrieving leaderboard")
	}

	fmt.Println("Elements in sorted set (in reverse order):", results)
	return results, nil
}
