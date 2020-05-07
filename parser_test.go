package gocron

import (
	"testing"
)

func Test_parser_validateExpr(t *testing.T) {
	type args struct {
		expr string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "", args: struct{ expr string }{expr: "0/30 "}, want: nil, wantErr: false},
		{name: "", args: struct{ expr string }{expr: "*/5"}, want: nil, wantErr: false},
		{name: "", args: struct{ expr string }{expr: " 10,20,30,40,50 "}, want: nil, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{}
			got, err := p.validateExpr(tt.args.expr)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateExpr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			t.Log(got)
		})
	}
}

func Test_parser_parse(t *testing.T) {
	type args struct {
		expr string
		name []string
	}
	tests := []struct {
		name    string
		args    args
		want    *spec
		wantErr bool
	}{
		{name: "", args: struct {
			expr string
			name []string
			//}{expr: "1,3-6/21 0,5,10,15-18,20,25,55-59,35,40,45-46,50,30 20-23 1,2,3-6,7,8", name: nil}, want: nil, wantErr: false},
		}{expr: "* 20,5,10,50,15,25,55-58,35,40,0,45,30 6-12,15,18-22 1-3,7,8-9/2 1,5-7/4 1-4", name: nil}, want: nil, wantErr: false},
		//}{expr: "0 0,10 17-22 * * 0,2-5/1", name: nil}, want: nil, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{}
			got, err := p.parse(tt.args.expr, tt.args.name...)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			t.Log(got)
		})
	}
}
