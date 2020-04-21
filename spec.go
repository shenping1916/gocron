package gocron

import "time"

type bound struct {
	min, max uint
}

var (
	second = bound{0, 59}
	minute = bound{0, 59}
	hour = bound{0, 23}
	day = bound{1, 31}
	month = bound{1, 12}
	week = bound{0, 6}
)

type spec struct {
	Second uint64
	Minute uint64
	Hour uint64
	Day uint64
	Month uint64
	Week uint64

	Local *time.Location
}

func (s *spec) NexTime(t time.Time) time.Time {
	if t.IsZero() {
		return time.Time{}
	}

	if s.Local == time.Local {
		s.Local = t.Location()
	}

	//// second
	//for 1<<uint(t.Second())&s.Second == 0 {
	//
	//}

	return t.In(s.Local)
}

