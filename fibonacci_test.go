package fibonacci

import (
	"fmt"
	"testing"
)

func TestFibonacciByTable(t *testing.T) {
	var tests = []struct {
		n, r int
	}{
		{1, 1},
		{10, 55},
		{20, 6765},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.n, tt.r)
		t.Run(testname, func(t *testing.T) {
			r := Fibonacci(tt.n)

			if r != tt.r {
				t.Errorf("Result is incorrect. Received %d but expected %d", r, tt.r)
			}
		})
	}
}
