package gocron

import (
	"errors"
	valid "github.com/asaskevich/govalidator"
	"strings"
	"time"
)

const (
    //
	expressL = 6
)

var symbol = [5]string{
	"*",
	",",
	"/",
	"-",
	"?",
}

var (
	// ErrExpression defines the error description: the format
	// of the current scheduled task expression is incorrect
	ErrExpression = errors.New("expression format error")

	// ErrExpressSymbol defines the error description: The
	// current area length is 1, but it does not belong to
	// the following symbols: "*", "?"
	ErrExpressSymbol = errors.New("wrong expression symbol or number")
)

type parser struct {
}

func newParser() *parser {
	return &parser{}
}

func (p *parser) parse(expr string, local ...time.Location) (*spec, error) {
	if expr == "" {
		return nil, nil
	}

	fields, err := p.validateExpr(expr)
    if err != nil {
    	return nil, err
	}
	_ = fields

	location := p.getLocation(local...)
	return &spec{
		//Second: ,
		//Minute: ,
		//Hour: ,
		//Day: ,
		//Month: ,
		//Week: ,
		Local: location,
	}, nil
}

func (p *parser) validateExpr(expr string) ([]string, error) {
	fields := strings.Fields(expr)
	if len(fields) == 0 {
		return fields, ErrExpression
	}

	newFields := make([]string, expressL)
	if len(fields) < expressL {
		for i := len(fields); i < expressL; i++ {
			newFields[i] = "*"
		}
	}
	for index, field := range fields {
		if len([]rune(field)) == 1 && (field != symbol[0] && field != symbol[len(symbol) -1] && !valid.IsInt(field)) {
			return newFields, ErrExpressSymbol
		}

		newFields[index] = field
	}

	return newFields, nil
}

func (p *parser) getLocation(local ...time.Location) *time.Location {
	if len(local) > 0 {
		l := local[0]
		return &l
	}

	return time.Local
}
