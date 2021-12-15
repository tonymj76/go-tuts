package module1

import "testing"

func TestBubbleSort(t *testing.T) {
	testCase := []struct {
		a, want []int
	}{
		{[]int{7, 3, 4, 23, 2}, []int{2, 3, 4, 7, 23}},
	}

	for _, tc := range testCase {
		t.Run("", func(t *testing.T) {
			got := BubbleSort(tc.a)
			t.Log(got)
			for i := 0; i < len(got); i++ {
				if got[i] != tc.want[i] {
					t.Fatalf("want %v; got %v", tc.want[i], got[i])
				}
			}
		})
	}
}
