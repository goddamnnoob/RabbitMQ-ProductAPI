package integrations

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func GetNewRabbitMQConnection() (connection *amqp.Connection, err error) {
	rabbitmqendp := os.Getenv("RABBITMQ_ENDPOINT")
	connection, err = amqp.Dial(rabbitmqendp)
	if err != nil {
		log.Println("Connection to RabbitMQ failed")
		return nil, err
	}
	return connection, nil

}
