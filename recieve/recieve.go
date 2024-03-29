package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "failed to connect rabbitmq")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to Open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	failOnError(err, "failed to declare a queue")

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	failOnError(err, "failed to register consumer")
	var forever chan struct{}
	go func() {
		for d := range msgs {
			log.Printf("message recieved", string(d.Body))
		}
	}()
	log.Printf("[*] waiting for messages,to exit press CTRL +C ")
	<-forever
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf(msg, err)
	}
}
