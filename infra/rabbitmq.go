package infra

import (
	"fmt"

	linq "github.com/ahmetb/go-linq"
	"o2b.com.br/WhatsAppProcessWorker/domain"

	"github.com/streadway/amqp"
)

// Queue rabbitmq queue
type Queue struct {
	Name             string `json:"name,omitempty"`
	Durable          bool   `json:"durable,omitempty"`
	DeleteWhenUnused bool   `json:"delete_when_unused,omitempty"`
	Exclusive        bool   `json:"exclusive,omitempty"`
	NoWait           bool   `json:"no_wait,omitempty"`
}

// RabbitMQ rabbit mq
type RabbitMQ struct {
	User     string  `json:"user,omitempty"`
	Password string  `json:"password,omitempty"`
	Port     string  `json:"Port,omitempty"`
	Host     string  `json:"Host,omitempty"`
	Queue    []Queue `json:"queue,omitempty"`
}

func (mq *RabbitMQ) connectionChannel() *amqp.Channel {
	conn := mq.GetConnection()
	ch, err := conn.Channel()
	domain.FailOnError(err, "Failed to connect to chanel")

	return ch

}
func (mq *RabbitMQ) declareQueue(queueName string) (*amqp.Queue, *amqp.Channel) {
	processQueue := linq.From(mq.Queue).Where(func(c interface{}) bool {
		return c.(Queue).Name == queueName
	}).First().(Queue)

	ch := mq.connectionChannel()

	q, err := ch.QueueDeclare(
		processQueue.Name,
		processQueue.Durable,
		processQueue.DeleteWhenUnused,
		processQueue.Exclusive,
		processQueue.NoWait,
		nil, // arguments
	)

	domain.FailOnError(err, "Falied create queue")

	return &q, ch

}
func (mq *RabbitMQ) formatAmpqURI() string {
	fmt.Println("amqp://" + mq.User + ":" + mq.Password + "@" + mq.Host + ":" + mq.Port)
	return "amqp://" + mq.User + ":" + mq.Password + "@" + mq.Host + ":" + mq.Port

}

// Publish publish to rabbitmq queue
func (mq *RabbitMQ) Publish(payload string, queueName string) {

	q, ch := mq.declareQueue(queueName)

	err := ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(payload),
		})
	domain.FailOnError(err, "Failed to publish a message")
}

// Consume publish to rabbitmq queue
func (mq *RabbitMQ) Consume(queueName string) <-chan amqp.Delivery {

	q, ch := mq.declareQueue(queueName)
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	domain.FailOnError(err, "Failed to publish a message")

	return msgs
}

// GetConnection get RabbitMQ connection
func (mq *RabbitMQ) GetConnection() *amqp.Connection {
	conn, err := amqp.Dial(mq.formatAmpqURI())
	domain.FailOnError(err, "Failed to connect to RabbitMQ")

	return conn
}

// NewRabbitMQ is my IoC
func NewRabbitMQ() *RabbitMQ {
	return GetManagerAppConfig().RabbitMQ
}
