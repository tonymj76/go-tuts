package module1

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	// go test -run=Fibonacci -v
	testCase := []struct {
		n, want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{14, 377},
		{22, 17711},
		{25, 75025},
		// This test case may be much slower depending on your
		// solution. We will look at how to speed it up in a future
		// module. Feel free to comment it out if it is too slow.
		{41, 165580141},
	}
	for _, test := range testCase {
		t.Run(fmt.Sprintf("n=%v", test.n), func(t *testing.T) {
			if got := Fibonacci(test.n); got != test.want {
				t.Fatalf("fibonacci(%d)  got = %d; want = %d\n", test.n, got, test.want)
			}
		})
	}
}
