package gocron

import (
	"testing"
	"time"
)

func Test_spec_nexTime(t *testing.T) {
	//now := time.Now()
	now := time.Unix(1592150401, 0)
	type fields struct {
		Second []uint
		Minute []uint
		Hour   []uint
		Day    []uint
		Month  []uint
		Week   []uint
		Local  *time.Location
	}
	type args struct {
		t time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   time.Time
	}{
		{
			name: "",
			fields: struct {
				Second []uint
				Minute []uint
				Hour   []uint
				Day    []uint
				Month  []uint
				Week   []uint
				Local  *time.Location
			}{
				//Second: []uint{0, 6, 12, 18, 24, 30, 36, 42, 48, 54},
				//Minute: []uint{0, 7, 14, 21, 28, 35, 42, 49, 56},
				//Hour:   []uint{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22},
				//Day:    []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
				//Month:  []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				//Week:   []uint{0, 1, 2, 3, 4, 5, 6},
				Second: []uint{0},
				Minute: []uint{0},
				Hour:   []uint{0},
				Day:    []uint{1, 15},
				Month:  []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				Week:   []uint{1},
				Local:  time.Local},
			args: struct{ t time.Time }{now},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &spec{
				Second: tt.fields.Second,
				Minute: tt.fields.Minute,
				Hour:   tt.fields.Hour,
				Day:    tt.fields.Day,
				Month:  tt.fields.Month,
				Week:   tt.fields.Week,
				Local:  tt.fields.Local,
			}
			got := s.delayTime(tt.args.t)
			t.Log(got)
		})
	}
}
