package rabbit

import (
	"github.com/rkiminius/carbon-based-life-forms/helper"
	log "github.com/sirupsen/logrus"
)

// StartConsumer -
func (conn Conn) StartConsumer(queueName string, handler func(body []byte)) {

	q, err := conn.Channel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	helper.FailOnError(err, "Failed to declare a queue")

	msgs, err := conn.Channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	helper.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			//log.Printf("Received a message: %s", d.Body)
			handler(d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
