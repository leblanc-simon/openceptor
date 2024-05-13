package handler

import (
	"context"
	"log"
	"time"

	"openceptor.eu/connection"
	"openceptor.eu/request"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func SendToQueue(r *request.Request) {
	conn := connection.GetRabbitMqInstance(nil)

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"request", // name
		true,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := r.ToJson()
	failOnError(err, "Failed to encoding request")

	err = ch.PublishWithContext(ctx, 
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
			Headers: amqp.Table{
				"type": "\\App\\Domain\\Request\\Request",
			},
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
