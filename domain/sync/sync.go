package sync

import "log"

// Sync is a struct to sync my gorotines
type Sync struct {
	Done    chan int
	GoWork  chan bool
	ToHire  chan bool
	Forever chan bool
	Hired   chan bool
	Start   chan bool
}

func (s *Sync) WorkForever() {
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-s.Forever
}

// NewSync is my contructor
func NewSync() *Sync {
	return &Sync{
		Done:    make(chan int, 3),
		GoWork:  make(chan bool),
		ToHire:  make(chan bool),
		Forever: make(chan bool),
		Hired:   make(chan bool),
	}
}
