package goRotines

import "o2b.com.br/WhatsAppProcessWorker/infra"

type TodoPile struct {
	Name string
}

type DonePile struct {
	Name string
}
type WorkPile struct {
	Pile     *infra.RabbitMQ
	TodoPile *TodoPile
	DonePile *DonePile
}

func NewWorkPile(donePileName string, todoPileName string) *WorkPile {
	return &WorkPile{
		Pile: infra.NewRabbitMQ(),
		TodoPile: &TodoPile{
			Name: todoPileName,
		},
		DonePile: &DonePile{
			Name: donePileName,
		},
	}
}
