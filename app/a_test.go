package app

import (
	"testing"
)

func TestFib(t *testing.T) {
	var (
		in       = 7
		expected = 13
	)
	actual := Fib(in)
	if actual != expected {
		t.Errorf("Fib(%d) = %d; expected %d", in, actual, expected)
	}
}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func BenchmarkFib10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}
