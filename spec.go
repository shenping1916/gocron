package gocron

import (
	"fmt"
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

// delayTime calculates the specific time for delayed execution (next execution)
// based on the incoming time parameters
func (s *spec) delayTime(t time.Time) time.Time {
	if t.IsZero() {
		return time.Time{}
	}

	if s.Local == time.Local {
		s.Local = t.Location()
	}

	// ok is used to indicate whether the time unit needs to be incremented
	ok := false

	t = s.timeTailoring(t)
	fmt.Println(t)

SECOND:
	// second
	for uint(t.Second())^s.ClosedNumber(s.Second, uint(t.Second())) != 0 {
		t = t.Add(1 * time.Second)
		if t.Second() == 0 {
			ok = true
		}

		goto SECOND
	}

MINUTE:
	// minute
	for uint(t.Minute())^s.ClosedNumber(s.Minute, uint(t.Minute())) != 0 {
		if t.Minute() == 0 {
			ok = true
		}

		if ok {
			t = t.Add(1 * time.Minute)
		}
		goto MINUTE
	}

HOUR:
	// hour
	for uint(t.Hour())^s.ClosedNumber(s.Hour, uint(t.Hour())) != 0 {
		if t.Hour() == 0 {
			ok = true
		}

		if ok {
			t = t.Add(1 * time.Hour)
		}
		goto HOUR
	}

DAY:
	// day or week
	for uint(t.Day())^s.ClosedNumber(s.Day, uint(t.Day())) != 0 ||
		uint(t.Weekday())^s.ClosedNumber(s.Week, uint(t.Weekday())) != 0 {
		if t.Day() == 1 || t.Weekday() == time.Sunday {
			ok = true
		}

		if ok {
			t = t.AddDate(0, 0, 1)
		}
		goto DAY
	}

MONTH:
	for uint(t.Month())^s.ClosedNumber(s.Month, uint(t.Month())) != 0 {
		if t.Month() == 1 {
			ok = true
		}

		if ok {
			t = t.AddDate(0, 1, 0)
		}
		goto MONTH
	}

	return t.In(s.Local)
}

// ClosedNumber is used to get the closest number of the specified
// number in the array. If the specified number crosses the boundary,
// the first element of the array is returned
func (s *spec) ClosedNumber(array []uint, n uint) uint {
	for i := 0; i < len(array); i++ {
		if number := array[i]; number >= n {
			return number
		}
	}

	return array[0]
}

// TimeTailoring provides a complete method of time,
// removing the extra decimal places of time (milliseconds
// and nanoseconds)
func (s *spec) timeTailoring(t time.Time) time.Time {
	return t.Add(-time.Duration(t.Nanosecond()) * time.Nanosecond)
}
