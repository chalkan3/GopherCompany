package goRotines

import (
	"o2b.com.br/WhatsAppProcessWorker/domain"

	"o2b.com.br/WhatsAppProcessWorker/domain/entities"
)

// Worker is my worker
type Worker struct {
	Message  *entities.Message
	WorkPile *WorkPile
}

func (w *Worker) putOnDonePile() {
	doneMessage := domain.JSONStringfy(&entities.Message{
		ID: w.Message.ID,
	})

	w.WorkPile.Pile.Publish(doneMessage, w.WorkPile.DonePile.Name)
}

// Process is my worker process
func (w *Worker) Process() {

	w.putOnDonePile()
}

// NewWorker constructor
func NewWorker(_message *entities.Message, _workPile *WorkPile) *Worker {
	return &Worker{
		Message:  _message,
		WorkPile: _workPile,
	}
}
