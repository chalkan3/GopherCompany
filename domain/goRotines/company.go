package goRotines

import (
	"log"

	"o2b.com.br/WhatsAppProcessWorker/domain/sync"

	"o2b.com.br/WhatsAppProcessWorker/infra"
)

// GopherCompany my company for process jobs
type GopherCompany struct {
	TeamWork   *sync.Sync
	TeamLeader *TeamLeader
	Jobs       *infra.RabbitMQ
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
func (company *GopherCompany) hireTeamLeaderHumansManagement() *GopherCompany {
	company.TeamLeader = NewTeamLeader(company.TeamWork).HireHumansManagement()
	return company
}

// HireNewEmployee syncronize the workers
func (company *GopherCompany) findEmployee() {
	var hireFirstEmployee bool = true
	company.TeamLeader.HumansManagement.ToHireWorkers(hireFirstEmployee)
}

// TasksTodo work on new jobs
func (company *GopherCompany) tasksTodo() *GopherCompany {
	company.TeamLeader.Jobs()
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
	go company.findEmployee()
	go company.tasksTodo()
	company.ceoWorking()
}

// FindTeamLeaders costructor of my syncronizer e o job
func (company *GopherCompany) FindTeamLeaders() *GopherCompany {
	return company.hireTeamLeaderHumansManagement()
}

// FoundGopherCompany constructor
func FoundGopherCompany() *GopherCompany {
	return &GopherCompany{
		TeamWork: sync.NewSync(),
		Jobs:     infra.NewRabbitMQ(),
	}
}
