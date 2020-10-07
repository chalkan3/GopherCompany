package goRotines

import (
	"fmt"
	"log"
	"time"

	"o2b.com.br/WhatsAppProcessWorker/domain"
	"o2b.com.br/WhatsAppProcessWorker/domain/entities"
	"o2b.com.br/WhatsAppProcessWorker/domain/sync"
)

type HireWorkers struct {
	Sync        *sync.Sync
	CompanySize int
	Workers     int
}

func (c *HireWorkers) hiredGetACoffe() {
	c.Workers = c.Workers + 1
	c.Sync.GoWork <- true
}
func (c *HireWorkers) TimeToHire() {
	for {
		fmt.Println("dasd")
		time.Sleep(5 * time.Second)
		c.Sync.ToHire <- true
		fmt.Println("tohire")

	}

}
func (c *HireWorkers) ToWork(workerName string, done chan int, message *entities.Message) {
	log.Printf("******************* WORKER " + workerName + " START TO WORK ******************************")

	worker := domain.NewWorker()
	worker.Message = message

	worker.Process()
	log.Printf("******************* WORKER IS RESTING A LITTLE ******************************")
	done <- 1

}
func (c *HireWorkers) handShakeFirstEmployee(firstEmployee bool) bool {
	c.hiredGetACoffe()
	return false
}

func (c *HireWorkers) ToHireWorkers(firstEmployee bool) {
	for {
		select {
		case <-c.Sync.ToHire:
			select {
			case <-c.Sync.Done:
				if c.Workers < c.CompanySize {
					c.Sync.Hired <- true
				}
			default:

			}
		case <-c.Sync.Hired:
			c.hiredGetACoffe()
		default:
			firstEmployee = c.handShakeFirstEmployee(firstEmployee)
		}
	}

}

func NewWorkers(_sync *sync.Sync) *HireWorkers {
	return &HireWorkers{
		Sync:        _sync,
		CompanySize: 3,
		Workers:     0,
	}
}