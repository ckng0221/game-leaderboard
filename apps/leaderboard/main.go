package main

import (
	"api/initializers"
	"encoding/json"
	"leaderboard/utils"
	"log"
	"os"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToRedis()
}

type ScoreEvent struct {
	UserId uint `json:"user_id"`
	Score  int  `json:"score"`
}

func main() {
	rabbitMqHost := os.Getenv("RABBIT_MQ_HOST")
	exchangeName := os.Getenv("EXCHANGE_NAME")

	rabbitMQ, err := utils.RabbitMQConsumer(rabbitMqHost, exchangeName)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}
	defer rabbitMQ.Close()

	// Consume messages
	msgs, err := rabbitMQ.Consume()
	if err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}

	// Process messages
	go func() {
		for msg := range msgs {
			log.Printf("Received a message: %s", msg.Body)

			// Update Redis
			var scoreEvent ScoreEvent
			err := json.Unmarshal(msg.Body, &scoreEvent)
			if err != nil {
				log.Println("failed to unmarshal body: %w", err)
				return
			}

			utils.IncrementUserScore(initializers.RedisClient, scoreEvent.UserId, scoreEvent.Score)
		}
	}()

	// Block main thread
	log.Println("Listening score event...")
	select {}
}
