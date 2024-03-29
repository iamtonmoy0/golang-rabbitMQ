package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("ampq://guest:guest@localhost:5672/")
	failOnError(err, "failed to connect rabbitmq")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to Open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	failOnError(err, "failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "hello world"
	err = PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(body)})
	failOnError(err, "failed  to publish a message")
	log.Printf("[x] sent %s\n", body)
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf(msg, err)
	}
}
