package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func failOnError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %s", message, err)
	}
}

func Receive(ch *amqp.Channel) {
	msgs, err := ch.Consume("hello", "", false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	// empty struct help ping with no data value (save data memory)
	var forever chan struct{}
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	Receive(ch)
}
