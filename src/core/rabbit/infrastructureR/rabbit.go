package infrastructureR

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

type RabbitMQConfig struct {
	URL       string
	QueueName string
}

func NewRabbitMQ() (*RabbitMQ, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("no se pueden cargar datos del archivo .env", err)
	}

	rabbitURL := os.Getenv("RABBIT_URL")
	queueName := os.Getenv("QUEUE_TWO")

	if rabbitURL == "" || queueName == "" {
		return nil, fmt.Errorf("valores indefinidos")
	}
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		return nil, fmt.Errorf("no se puede conctar al servidor: %w", err)
	}
	fmt.Println("conectado a rabbit...")

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("no se puede abrir el canal en rabbit", err)
	}
	_, err = ch.QueueDeclare(
		queueName,
		true,  //durable
		false, //auto-delete
		false, //exclusive
		false, //nowait
		nil,   //argumentos adicionales
	)
	if err != nil {
		return nil, fmt.Errorf("no se puede declarar la cola: %w", err)
	}
	return &RabbitMQ{
		connection: conn,
		channel:    ch,
	}, nil
}

func (client *RabbitMQ) SendNotification(queueName string, msg []byte) error {
	err := client.channel.Publish(
		"",
		queueName,
		false, //obligatorio
		false, //inmediato
		amqp.Publishing{
			ContentType: "applitation/json",
			Body:        msg,
		},
	)
	if err != nil {
		return fmt.Errorf("no se puede notificar %w", err)
	}
	return nil
}

func (client *RabbitMQ) Close() error {
	if err := client.channel.Close(); err != nil {
		return err
	}
	return client.connection.Close()
}
