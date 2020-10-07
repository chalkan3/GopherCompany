package main

import (
	"o2b.com.br/WhatsAppProcessWorker/domain/goRotines"
	"o2b.com.br/WhatsAppProcessWorker/domain/sync"

	"o2b.com.br/WhatsAppProcessWorker/infra"
)

func main() {
	sync := sync.NewSync()
	job := goRotines.NewJob(infra.NewRabbitMQ(), sync)
	hireWorkers := goRotines.NewWorkers(sync)

	go hireWorkers.ToHireWorkers(true)
	go job.Jobs(hireWorkers)
	sync.WorkForever()
}
