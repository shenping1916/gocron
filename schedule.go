package gocron

import "time"

type taskFunc func() error

type scheduler struct {
	Name         string
	Fc           taskFunc
	Pre          time.Time
	Next         time.Time
	SuccessCount uint64
	FailureCount uint64
	Running      bool
}

func newSchedule(name string, f taskFunc) *scheduler {
	return new(scheduler).init(name, f)
}

func (s *scheduler) init(name string, f taskFunc) *scheduler {
	s.Name = name
	s.Fc = f
	return s
}

func (s *scheduler) Reset() {
	s.Name = ""
	s.Fc = nil
	s.Pre = time.Time{}
	s.Next = time.Time{}
	s.SuccessCount = 0
	s.FailureCount = 0
	s.Running = false
}
