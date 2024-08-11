package utils

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

var RabbitMqConn *amqp.Connection

type RabbitMq struct {
	conn         *amqp.Connection
	channel      *amqp.Channel
	queue        amqp.Queue
	exchangeName string
}

func RabbitMQConsumer(amqpURL, exchangeName string) (*RabbitMq, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	err = channel.ExchangeDeclare(
		exchangeName, // name
		"fanout",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare an exchange: %w", err)
	}

	q, err := channel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare an queue: %w", err)
	}

	err = channel.QueueBind(
		q.Name,       // queue name
		"",           // routing key
		exchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to bind a queue: %w", err)
	}

	return &RabbitMq{
		conn:         conn,
		channel:      channel,
		exchangeName: exchangeName,
		queue:        q,
	}, nil
}

func (r *RabbitMq) Close() {
	r.channel.Close()
	r.conn.Close()
}

// Consume starts consuming messages from the queue.
func (r *RabbitMq) Consume() (<-chan amqp.Delivery, error) {
	msgs, err := r.channel.Consume(
		r.queue.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		return nil, fmt.Errorf("failed to register a consumer: %w", err)
	}
	return msgs, nil
}
