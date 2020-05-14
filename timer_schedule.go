package gocron

import "time"

type taskFunc func(args ...interface{}) error

type scheduler struct {
	name         string
	fc           taskFunc
	pre          time.Time
	next         time.Time
	successCount uint64
	failureCount uint64
	running      bool
}

func newSchedule(name string, f taskFunc) *scheduler {
	return new(scheduler).init(name, f)
}

func (s *scheduler) init(name string, f taskFunc) *scheduler {
	s.name = name
	s.fc = f
	return s
}

func (s *scheduler) Reset() {
	s.name = ""
	s.fc = nil
	s.pre = time.Time{}
	s.next = time.Time{}
	s.successCount = 0
	s.failureCount = 0
	s.running = false
}
