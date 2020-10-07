package domain

import (
	"fmt"

	"o2b.com.br/WhatsAppProcessWorker/domain/entities"
)

// Worker is my worker
type Worker struct {
	Message *entities.Message
}

// Process is my worker process
func (w *Worker) Process() {
	fmt.Println(*w.Message.Body)
}

// NewWorker constructor
func NewWorker() *Worker {
	return &Worker{}
}
