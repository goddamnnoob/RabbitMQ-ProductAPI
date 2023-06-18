package services

import (
	"goddamnnoob/RabbitMQ-ProductAPI/integrations"
	"log"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

const (
	queuename = "products"
)

func AddProductToMessageQueue(productid uuid.UUID) (err error) {

	connection, err := integrations.GetNewRabbitMQConnection()
	if err != nil {
		return err
	}
	defer connection.Close()
	ch, err := connection.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(queuename, false, false, false, false, nil)

	if err != nil {
		return err
	}

	log.Println("Queue Declared ")

	err = ch.Publish("", queuename, false, false, amqp.Publishing{
		ContentType: "document/text",
		Body:        []byte(productid.String()),
	})

	if err != nil {
		return err
	}

	log.Println("Pushed product to Queue")

	return nil

}
