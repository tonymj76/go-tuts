package module1

import "testing"

func TestEuclideanGCD(t *testing.T) {
	testCase := []struct {
		a, b, want int
	}{
		{252, 105, 21},
		{10, 5, 5},
		{25, 5, 5},
		{30, 15, 15},
		{30, 9, 3},
		{100, 9, 1},
	}
	for _, tc := range testCase {
		t.Run("euclideanGCD", func(t *testing.T) {
			if got := EuclideanGCD(tc.a, tc.b); got != tc.want {
				t.Fatalf("want %d; got %d", tc.want, got)
			}
		})
	}
}
