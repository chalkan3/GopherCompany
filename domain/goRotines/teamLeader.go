package goRotines

import (
	"encoding/json"

	"github.com/Pallinder/go-randomdata"
	"github.com/streadway/amqp"
	"o2b.com.br/WhatsAppProcessWorker/domain/entities"
	"o2b.com.br/WhatsAppProcessWorker/domain/sync"
	"o2b.com.br/WhatsAppProcessWorker/infra"
)

type TeamLeader struct {
	TeamTalk         *sync.Sync
	WorkPile         *infra.RabbitMQ
	HumansManagement *HumansManagement
}

func (teamLeader *TeamLeader) getNewTasks() <-chan amqp.Delivery {
	return teamLeader.WorkPile.Consume()
}

func (teamLeader *TeamLeader) working(tasks <-chan amqp.Delivery) {
	for task := range tasks {
		var message entities.Message
		json.Unmarshal(task.Body, &message)
		go teamLeader.HumansManagement.ToWork(randomdata.FullName(randomdata.Female), teamLeader.TeamTalk.Done, &message)
	}
}

func (teamLeader *TeamLeader) AskHumansManagement(humansManagement *HumansManagement) {
	teamLeader.HumansManagement = humansManagement
}
func (teamLeader *TeamLeader) planning() {
	teamLeader.working(teamLeader.getNewTasks())
}

func (teamLeader *TeamLeader) HireHumansManagement() *TeamLeader {
	teamLeader.HumansManagement = NewHumansManagement(teamLeader.TeamTalk)
	return teamLeader
}
func (teamLeader *TeamLeader) Jobs() {
	for {
		select {
		case <-teamLeader.TeamTalk.GoWork:
			teamLeader.planning()
		default:
		}
	}
}

func NewTeamLeader(_sync *sync.Sync) *TeamLeader {
	return &TeamLeader{
		TeamTalk: _sync,
		WorkPile: infra.NewRabbitMQ(),
	}
}
