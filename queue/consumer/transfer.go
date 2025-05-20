package consumer

import (
	"achmadardian/test-ewallet/queue"
	"achmadardian/test-ewallet/services"
	"encoding/json"
	"fmt"
	"log"
)

type TransferConsumer struct {
	rabbitmq        *queue.RabbitMqClient
	transferService *services.TransferService
}

func NewTransferConsumer(rabbitmq *queue.RabbitMqClient, transferService *services.TransferService) *TransferConsumer {
	return &TransferConsumer{
		rabbitmq:        rabbitmq,
		transferService: transferService,
	}
}

func (t *TransferConsumer) ConsumeTransfer(queueName string) error {
	q, err := t.rabbitmq.Chann.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return fmt.Errorf("failed to declare when consume: %s", err)
	}

	msgs, err := t.rabbitmq.Chann.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return fmt.Errorf("failed to consume queue %s", err)
	}

	go func() {
		for d := range msgs {
			var payload queue.Transfer

			err := json.Unmarshal(d.Body, &payload)
			if err != nil {
				log.Printf("error unmarshal msg: %s", err)
				continue
			}

			err = t.transferService.HandleTransfer(&payload)
			if err != nil {
				log.Printf("process: %s", err)
				continue
			}
		}
	}()

	return nil
}
