package gocron

import (
	"strconv"
	"testing"
)

func mustDigital(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func mustDigital1(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

var number = "109287329423893863"

func BenchmarkMustDigital(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mustDigital(number)
	}
}

func BenchmarkMustDigital1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mustDigital1(number)
	}
}
