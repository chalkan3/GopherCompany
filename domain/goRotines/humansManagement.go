package goRotines

import (
	"log"
	"strconv"

	"o2b.com.br/WhatsAppProcessWorker/domain/entities"
	"o2b.com.br/WhatsAppProcessWorker/domain/sync"
)

type HumansManagement struct {
	Talk        *sync.Sync
	CompanySize int
	Workers     int
}

func (c *HumansManagement) hiredGetACoffe() {
	c.Workers = c.Workers + 1
	c.Talk.GoWork <- true
}

func (c *HumansManagement) ToWork(workerName string, done chan int, message *entities.Message, workPile *WorkPile) {

	messageID := strconv.FormatInt(message.ID, 10)
	log.Printf("Gopher [" + workerName + "] says: I'll process your order number[" + messageID + "]")
	NewWorker(message, workPile).Process()
	log.Printf("Gopher [" + workerName + "] says: The order number[" + messageID + "] done")
	done <- 1

}
func (c *HumansManagement) handShakeFirstEmployee(firstEmployee bool) bool {
	c.hiredGetACoffe()
	return false
}

func (c *HumansManagement) ToHireWorkers(firstEmployee bool) {
	for {
		select {
		case <-c.Talk.ToHire:
			select {
			case <-c.Talk.Done:
				if c.Workers < c.CompanySize {
					c.Talk.Hired <- true
				}
			default:

			}
		case <-c.Talk.Hired:
			c.hiredGetACoffe()
		default:
			firstEmployee = c.handShakeFirstEmployee(firstEmployee)
		}
	}

}

func NewHumansManagement(_sync *sync.Sync) *HumansManagement {
	return &HumansManagement{
		Talk:        _sync,
		CompanySize: 3,
		Workers:     0,
	}
}
