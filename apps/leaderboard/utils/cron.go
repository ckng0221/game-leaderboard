package utils

import (
	"api/initializers"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func ClearLeaderboardCron() {
	// Cron job, to clear the Redis leaderboard sorted set
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 * * * * *", func() {
		currentTime := time.Now()

		// log.Println("Current time:", time.Now().Format("2006-01-02 15:04"))

		// Check if it's 00:00 of the first day of the month
		if currentTime.Hour() == 0 && currentTime.Minute() == 0 && currentTime.Day() == 1 {
			log.Println("It's 00:00 of a new month")
			ClearLeaderboard(initializers.RedisClient)
			log.Println("Done clearing leaderboard")
		}

	})

	c.Start()
}
