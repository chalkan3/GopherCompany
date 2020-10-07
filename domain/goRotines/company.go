package goRotines

import (
	"log"

	"o2b.com.br/WhatsAppProcessWorker/domain/sync"

	"o2b.com.br/WhatsAppProcessWorker/infra"
)

// GopherCompany my company for process jobs
type GopherCompany struct {
	TeamWork         *sync.Sync
	HumansManagement *HireWorkers
	TeamLeader       *Job
	Jobs             *infra.RabbitMQ
}

/*

	Private Functions ->
		-> hireTeamLeader
		-> hireHumansManagement
		-> tasksTodo
		-> hireNewEmployee
		-> ceoWorking
*/

// HireTeamLeader create the jobs workers
func (company *GopherCompany) hireTeamLeader() *GopherCompany {
	company.TeamLeader = NewJob(company.Jobs, company.TeamWork)
	return company
}

// HireHumansManagement create new syncronize gophers
func (company *GopherCompany) hireHumansManagement() *GopherCompany {
	company.HumansManagement = NewWorkers(company.TeamWork)
	return company
}

// HireNewEmployee syncronize the workers
func (company *GopherCompany) hireNewEmployee() {
	var hireFirstEmployee bool = true
	company.HumansManagement.ToHireWorkers(hireFirstEmployee)
}

// TasksTodo work on new jobs
func (company *GopherCompany) tasksTodo() *GopherCompany {
	company.TeamLeader.Jobs(company.HumansManagement)
	return company
}

// CeoWorking block main thread
func (company *GopherCompany) ceoWorking() {
	company.TeamWork.WorkForever()
}

/*

	Private Functions ->
		-> GophersWork
		-> tasksTodo
		-> FindTeamLeaders
		-> NewGopherCompany (constructor)

*/

// GophersWork block main thread
func (company *GopherCompany) GophersWork() {
	log.Println("GOPHERS ARE WORKING HERE")
	go company.hireNewEmployee()
	go company.tasksTodo()
	company.ceoWorking()
}

// FindTeamLeaders build gorotines syncronizes
func (company *GopherCompany) FindTeamLeaders() *GopherCompany {
	return company.hireHumansManagement().hireTeamLeader()
}

// NewGopherCompany constructor
func NewGopherCompany() *GopherCompany {
	return &GopherCompany{
		TeamWork: sync.NewSync(),
		Jobs:     infra.NewRabbitMQ(),
	}
}
