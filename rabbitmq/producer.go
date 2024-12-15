package rabbitmq

import (
	"log"
	"product-management/config"

	"github.com/streadway/amqp"
)

func PublishToQueue(message string) error {
	conn, err := amqp.Dial(config.AppConfig.RabbitMQ)
	if err != nil {
		return err
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	err = channel.Publish("", "imageQueue", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})
	if err != nil {
		log.Println("Failed to publish message:", err)
		return err
	}
	return nil
}
