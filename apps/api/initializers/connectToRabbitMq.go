package initializers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

var RabbitMqConn *amqp.Connection

type RabbitMq struct {
	conn         *amqp.Connection
	channel      *amqp.Channel
	exchangeName string
}

func NewRabbitMQ(amqpURL, exchangeName string) (*RabbitMq, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	return &RabbitMq{
		conn:         conn,
		channel:      channel,
		exchangeName: exchangeName,
	}, nil
}

func (r *RabbitMq) Close() {
	r.channel.Close()
	r.conn.Close()
}

func (r *RabbitMq) PublishScore(userId uint, score int) error {
	err := r.channel.ExchangeDeclare(
		r.exchangeName, // name
		"fanout",       // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare an exchange: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := map[string]any{
		"user_id": userId,
		"score":   score,
	}
	bodyByte, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error parsing body: %w", err)
	}

	err = r.channel.PublishWithContext(ctx,
		r.exchangeName, // exchange
		"",             // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        bodyByte,
		})

	if err != nil {
		return fmt.Errorf("failed to publish a message:%w", err)
	}

	log.Printf(" [x] Published %s", bodyByte)

	return nil
}

var RabbitMqObj *RabbitMq

func ConnectToRabbitMq() {
	rabbitMqHost := os.Getenv("RABBIT_MQ_HOST")
	exchangeName := os.Getenv("EXCHANGE_NAME")

	var err error
	RabbitMqObj, err = NewRabbitMQ(rabbitMqHost, exchangeName)
	if err != nil {
		log.Fatalf("failed to initialize RabbitMQ: %v", err)
	}
	// defer RabbitMqObj.Close()
}
