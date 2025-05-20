package queue

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMqClient struct {
	Conn  *amqp.Connection
	Chann *amqp.Channel
}

func Init() *RabbitMqClient {
	url := os.Getenv("QUEUE_URL")
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("failed to connect rabbitmq: %s", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open rabbitmq channel: %s", err)
	}

	return &RabbitMqClient{
		Conn:  conn,
		Chann: ch,
	}
}

type Transfer struct {
	Id                     uuid.UUID
	SenderId               uuid.UUID
	ReceiverId             uuid.UUID
	SenderBalanceBefore    int64
	SenderBalanceAfter     int64
	ReceiverBalancerBefore int64
	ReceiverBalanceAfter   int64
	Amount                 int64
	Remarks                string
}

func (p *RabbitMqClient) PublishTransfer(tr Transfer, queueName string) error {
	q, err := p.Chann.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("failed to declare queue: %s", err)
	}

	body, err := json.Marshal(tr)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %s", err)
	}

	err = p.Chann.Publish(
		"",
		q.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		return fmt.Errorf("failed to publish message: %s", err)
	}

	return nil
}
