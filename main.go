package main

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

var (
	URL string = os.Getenv("URL")
)

func SendMessage(msg string) {
	fmt.Println("Connecting to rabbitmq...")
	conn, err := amqp.Dial(URL)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println("Successfully connected to rabbitmq instance")

	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
		return
	}
	defer ch.Close()

	err = ch.Publish(
		"",
		"Testing bang",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)

	if err != nil {
		log.Println(err)
		return
	}

	q, err := ch.QueueDeclare(
		"Testing bang",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(q)

	fmt.Println("Success publish message to queue")
}
