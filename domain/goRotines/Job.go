package goRotines

import (
	"encoding/json"

	"github.com/Pallinder/go-randomdata"
	"github.com/streadway/amqp"
	"o2b.com.br/WhatsAppProcessWorker/domain/entities"
	"o2b.com.br/WhatsAppProcessWorker/domain/sync"
	"o2b.com.br/WhatsAppProcessWorker/infra"
)

type Job struct {
	Sync     *sync.Sync
	RabbitMQ *infra.RabbitMQ
}

func (q *Job) working(tasks <-chan amqp.Delivery, hireWorkers *HireWorkers) {
	for task := range tasks {
		var message entities.Message
		json.Unmarshal(task.Body, &message)
		go hireWorkers.ToWork(randomdata.FullName(randomdata.Female), q.Sync.Done, &message)
	}
}
func (q *Job) Jobs(hireWorkers *HireWorkers) {

	for {
		select {
		case <-q.Sync.GoWork:
			tasksToDo := q.RabbitMQ.Consume()
			q.working(tasksToDo, hireWorkers)
		default:

		}

	}

}

func NewJob(_rabbitmq *infra.RabbitMQ, _sync *sync.Sync) *Job {
	return &Job{
		Sync:     _sync,
		RabbitMQ: _rabbitmq,
	}
}
