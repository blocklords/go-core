package crond

import "github.com/robfig/cron/v3"

type (
	Process interface {
		GetSpec() string
		Run()
	}
	Schedule struct {
		cron       *cron.Cron
		entryIdMap []cron.EntryID
		jobWrapper []cron.JobWrapper
		processes  []Process
	}

	ScheduleFn func(s *Schedule)
)

func ScheduleProcess(process Process) ScheduleFn {
	return func(s *Schedule) {
		s.processes = append(s.processes, process)
	}
}

func ScheduleCron(crond *cron.Cron) ScheduleFn {
	return func(s *Schedule) {
		s.cron = crond
	}
}

func NewSchedule(fns ...ScheduleFn) *Schedule {
	s := &Schedule{
		cron: cron.New(
			cron.WithSeconds(),
			cron.WithChain(
				cron.SkipIfStillRunning(cron.DefaultLogger),
				cron.Recover(cron.DefaultLogger),
			),
		),
		entryIdMap: make([]cron.EntryID, 0),
		processes:  make([]Process, 0),
	}

	for _, fn := range fns {
		fn(s)
	}

	return s
}

func (s *Schedule) Start() {
	for _, process := range s.processes {
		spec := process.GetSpec()
		if spec == "" {
			continue
		}
		entryID, err := s.cron.AddFunc(spec, process.Run)
		if err != nil {
			panic(err)
		}
		s.entryIdMap = append(s.entryIdMap, entryID)
	}
	s.cron.Start()
}

func (s *Schedule) Stop() {
	for _, entryId := range s.entryIdMap {
		s.cron.Remove(entryId)
	}
}
