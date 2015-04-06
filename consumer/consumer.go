package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/andreagrandi/go-amqp-example/contracts"
	"github.com/streadway/amqp"
	"log"
)

var (
	amqpURI = flag.String("amqp", "amqp://guest:guest@localhost:5672/", "AMQP URI")
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func init() {
	flag.Parse()
	initAmqp()
}

var conn *amqp.Connection
var ch *amqp.Channel
var replies <-chan amqp.Delivery

func initAmqp() {
	var err error
	var q amqp.Queue

	conn, err = amqp.Dial(*amqpURI)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")

	q, err = ch.QueueDeclare(
		"go-amqp-example", // name, leave empty to generate a unique name
		false,             // durable
		false,             // delete when usused
		true,              // exclusive
		false,             // noWait
		nil,               // arguments
	)
	failOnError(err, "Error declaring the Queue")

	replies, err = ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Error consuming the Queue")
}

func main() {
	log.Println("Start consuming the Queue...")
	for {
		for r := range replies {
			log.Println("New replies to Consume...")
			user := contracts.User{}
			json.Unmarshal(r.Body, &user)
			fmt.Printf("FirstName: %s, LastName: %s\n", user.FirstName, user.LastName)
		}
	}
}
