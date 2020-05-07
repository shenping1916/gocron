package gocron

import (
	"time"
)

type bound struct {
	min, max uint
}

var (
	seconds = bound{0, 59}
	minutes = bound{0, 59}
	hours   = bound{0, 23}
	days    = bound{1, 31}
	months  = bound{1, 12}
	weeks   = bound{0, 6}
)

type spec struct {
	Second []uint
	Minute []uint
	Hour   []uint
	Day    []uint
	Month  []uint
	Week   []uint

	Local *time.Location
}

func (s *spec) nexTime(t time.Time) time.Time {
	if t.IsZero() {
		return time.Time{}
	}

	if s.Local == time.Local {
		s.Local = t.Location()
	}

	t = s.timeTailoring(t)

SECOND:
	// second
	for uint(t.Second())^s.ClosedNumber(s.Second, uint(t.Second())) != 0 {
		if t.Second() == 0 {
			break
		}

		t.Add(1 * time.Second)
		goto SECOND
	}

	//// minute
	//for uint64(t.Minute())^s.Minute != 0 {
	//}
	//// hour
	//for uint64(t.Hour())^s.Hour != 0 {
	//}
	//// month
	//for uint64(t.Month())^s.Month != 0 {
	//}

	return t.In(s.Local)
}

//
func (s *spec) ClosedNumber(array []uint, n uint) uint {
	for i := 0; i < len(array); i++ {
		if number := array[i]; number > n {
			return number
		}
	}

	return 0
}

// TimeTailoring provides a complete method of time,
// removing the extra decimal places of time (milliseconds
// and nanoseconds)
func (s *spec) timeTailoring(t time.Time) time.Time {
	return t.Add(-time.Duration(t.Nanosecond()) * time.Nanosecond)
}
