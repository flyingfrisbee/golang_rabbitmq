package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	SendMessage("Hello")
	SendMessage("Rabbit")
	SendMessage("MQ")
	conn, err := amqp.Dial(URL)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
		return
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"Testing bang",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
		return
	}

	go func() {
		for d := range msgs {
			fmt.Printf("Received message: %s\n", d.Body)
		}
	}()

	fmt.Println("successfully connected to our rabbitmq instance")
	fmt.Println("waiting for messages...")
	forever := make(chan bool)
	<-forever
}
