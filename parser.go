package gocron

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	// expressL defines the default expression length
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

type Slice []uint

func (s Slice) Len() int           { return len(s) }
func (s Slice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Slice) Less(i, j int) bool { return s[i] < s[j] }

type parser struct {
}

func newParser() *parser {
	return &parser{}
}

// parse is used to parse expressions, the required field is expr,
// and the optional field is the location name (like: America/New_York,
// Asia/Shanghai, America/Los_Angeles, Europe/Berlin)
func (p *parser) parse(expr string, name ...string) (*spec, error) {
	if expr == "" {
		return nil, nil
	}

	fields, err := p.validateExpr(expr)
	if err != nil {
		return nil, err
	}

	var second, minute, hour, day, month, week []uint
	if err := func() error {
		var err error
		for index, field := range fields {
			if err != nil {
				break
			}
			switch index {
			case 0:
				second, err = p.fieldParse(field, seconds)
			case 1:
				minute, err = p.fieldParse(field, minutes)
			case 2:
				hour, err = p.fieldParse(field, hours)
			case 3:
				day, err = p.fieldParse(field, days)
			case 4:
				month, err = p.fieldParse(field, months)
			case 5:
				week, err = p.fieldParse(field, weeks)
			}
		}
		return err
	}(); err != nil {
		return nil, err
	}

	location, err := p.getLocation(name...)
	if err != nil {
		return nil, err
	}
	return &spec{
		Second: second,
		Minute: minute,
		Hour:   hour,
		Day:    day,
		Month:  month,
		Week:   week,
		Local:  location,
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
			newFields[i] = symbol[0]
		}
	}
	for index, field := range fields {
		if len([]rune(field)) == 1 && field != symbol[0] && field != symbol[len(symbol)-1] && !p.mustDigital(field) {
			return newFields, ErrExpressSymbol
		}

		newFields[index] = field
	}

	return newFields, nil
}

func (p *parser) fieldParse(block string, r bound) ([]uint, error) {
	var u Slice
	switch r.min {
	case 0:
		u = make([]uint, 0, r.max+1)
	case 1:
		u = make([]uint, 0, r.max)
	}

	// determine whether block is a number
	if p.mustDigital(block) {
		number, err := p.digitalBound(block, block, r.min, r.max)
		if err != nil {
			return u, err
		}

		u = append(u, number)
	} else if block == symbol[0] || block == symbol[4] {
		// "*" or "?"
		for i := r.min; i <= r.max; i++ {
			u = append(u, i)
		}
	} else {
		// "/"
		if strings.Contains(block, symbol[2]) {
			if s := strings.Split(block, symbol[2]); len(s) == 2 {
				if number := s[1]; p.mustDigital(number) {
					step, err := p.digitalBound(block, number, r.min, r.max)
					if err != nil {
						return u, err
					}

					if s[0] == symbol[0] {
						// like: "*/10"
						for i := r.min; i <= r.max; i += step {
							u = append(u, i)
						}
					} else if number := s[0]; p.mustDigital(number) {
						// like: "4/7"
						start, err := p.digitalBound(block, number, r.min, r.max)
						if err != nil {
							return u, err
						}

						for i := start; i <= r.max; i += step {
							u = append(u, i)
						}
					} else if number := s[1]; p.mustDigital(number) {
						// like: "1-3,7,8-9/2"  |  1,5-7/4
						step, err := p.digitalBound(block, number, r.min, r.max)
						if err != nil {
							return u, err
						}

						if ss := strings.Split(s[0], symbol[1]); len(ss) > 0 {
							for _, value := range ss {
								if p.mustDigital(value) {
									number, err := p.digitalBound(block, value, r.min, r.max)
									if err != nil {
										return []uint{}, err
									}

									number += step
									if number < r.min || number > r.max {
										return []uint{}, fmt.Errorf("field: %s plus step value: %d, must be in the range [%d, %d], but now it is: %d", ss[0], step, r.min, r.max, number)
									}

									u = append(u, number)
								} else if ss := strings.Split(value, symbol[3]); len(ss) == 2 {
									d, err := p.baseAssemble(block, ss[0], ss[1], r.min, r.max, step)
									if err != nil {
										return d, err
									}

									u = append(u, d...)
								}
							}
						}
					}
				}
			}
		}

		// ","
		// like: 0,10,20,30,40,50
		if strings.Contains(block, symbol[1]) {
			if s := strings.Split(block, symbol[1]); len(s) > 0 && !strings.Contains(block, symbol[2]) {
				for i := 0; i < len(s); i++ {
					e := s[i]
					if p.mustDigital(e) {
						number, err := p.digitalBound(block, e, r.min, r.max)
						if err != nil {
							return u, err
						}

						u = append(u, number)
					} else if ss := strings.Split(e, symbol[3]); len(ss) > 0 {
						d, err := p.baseAssemble(block, ss[0], ss[1], r.min, r.max)
						if err != nil {
							return d, err
						}

						u = append(u, d...)
					}
				}
			}
		}

		// "-"
		// like: 8-20
		if strings.Contains(block, symbol[3]) {
			if s := strings.Split(block, symbol[3]); len(s) == 2 && !strings.Contains(block, symbol[2]) {
				d, err := p.baseAssemble(block, s[0], s[1], r.min, r.max)
				if err != nil {
					return d, err
				}

				u = append(u, d...)
			}
		}

		sort.Sort(u)
	}

	return u, nil
}

func (p *parser) mustDigital(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func (p *parser) parseStringUint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

func (p *parser) baseAssemble(block, start, end string, min, max uint, step ...uint) ([]uint, error) {
	data := make([]uint, 0, max+1)
	if p.mustDigital(start) && p.mustDigital(end) {
		ns, err := p.digitalBound(block, start, min, max)
		if err != nil {
			return data, err
		}
		temp := ns

		ne, err := p.digitalBound(block, end, min, max)
		if err != nil {
			return data, err
		}

		if ns > ne {
			return data, fmt.Errorf("express: %s ==> start: [%d] must be less than or equal to end: [%d]", block, ns, ne)
		}

		if len(step) > 0 {
			ns += step[0]
			ne += step[0]
		}
		for i := ns; i <= ne; i++ {
			if i < min || i > max {
				return data, fmt.Errorf("field: \"%s\" of character: \"%d\"(%d+%d) out of bounds, must be in the range [%d, %d]", block, i, temp, step[0], min, max)
			}
			data = append(data, i)
		}
	}

	return data, nil
}

func (p *parser) digitalBound(field, character string, min, max uint) (uint, error) {
	number, err := p.parseStringUint64(character)
	if err != nil {
		return 0, err
	}

	if numberUint := uint(number); numberUint >= min && numberUint <= max {
		return numberUint, nil
	}

	return 0, fmt.Errorf("field: \"%s\" of character: \"%s\" out of bounds, must be in range [%d, %d]", field, character, min, max)
}

func (p *parser) getLocation(name ...string) (location *time.Location, err error) {
	if len(name) > 0 {
		location, err = time.LoadLocation(name[0])
	}

	location = time.Local
	return
}
