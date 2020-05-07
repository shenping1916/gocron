package gocron

import (
	"testing"
	"time"
)

func Test_spec_nexTime(t *testing.T) {
	now := time.Now()
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
			}{Second: []uint{5}, Minute: []uint{0}, Hour: []uint{0}, Day: []uint{0}, Month: []uint{0}, Week: []uint{0}, Local: time.Local},
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
			got := s.nexTime(tt.args.t)
			t.Log(got)
		})
	}
}
