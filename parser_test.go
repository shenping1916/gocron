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
		{name: "", args: struct{ expr string }{expr: "0/30 ?"}, want: nil, wantErr: false},
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