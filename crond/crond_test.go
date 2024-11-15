package crond

import (
	"log"
	"testing"
	"time"
)

type (
	process1 struct{}
	process2 struct{}
)

func (p *process1) GetSpec() string {
	return "@every 1s"
}

func (p *process1) Run() {
	log.Printf("[%d] process1 running ...", time.Now().Unix())
}

func (p *process2) GetSpec() string {
	return "@every 3s"
}

func (p *process2) Run() {
	time.Sleep(5 * time.Second)
	log.Printf("[%d] process2 running ...", time.Now().Unix())
}

func TestNewSchedule(t *testing.T) {
	s := NewSchedule(
		ScheduleProcess(&process1{}),
		ScheduleProcess(&process2{}),
	)

	s.Start()
	select {}
}
