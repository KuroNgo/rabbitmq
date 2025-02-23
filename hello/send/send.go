package main

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func failOnError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %s", message, err)
	}
}
func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	data := "Phong nè"
	sent, byteData, err := Sent(ch, context.Background(), []byte(data))
	if err != nil {
		return
	}

	log.Printf(" [x] Sent %s %s", sent, byteData)
}

func Sent(ch *amqp.Channel, ctx context.Context, body []byte) (string, []byte, error) {
	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx, "", q.Name, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")

	return q.Name, body, nil
}
