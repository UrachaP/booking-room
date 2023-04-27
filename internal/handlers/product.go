package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"bookingrooms/internal/models"
	"github.com/labstack/echo/v4"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (h Handlers) BuyProduct(c echo.Context) error {
	var requestProduct models.RequestProduct

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//amount of data
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	q, err := ch.QueueDeclare(
		"1",   // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	go func() {
		for i := 0; i < 20; i++ {
			body, err := json.Marshal(models.RequestProduct{ID: 1, Amount: 1})
			if err != nil {
				log.Println(err.Error())
				return
			}
			err = ch.PublishWithContext(c.Request().Context(),
				"",     // exchange
				q.Name, // routing key
				false,  // mandatory
				false,  // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        body,
				})
			failOnError(err, "Failed to publish a message")
			log.Printf(" [x] Sent %s\n", body)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan error

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			b := d.Body
			err := json.Unmarshal(b, &requestProduct)
			if err != nil {
				log.Println(err.Error())
				//return
			}
			err = h.productService.CutProductStock(requestProduct)
			if err != nil {
				log.Println(err.Error())
				//return
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return c.JSON(http.StatusOK, "buy product success")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
