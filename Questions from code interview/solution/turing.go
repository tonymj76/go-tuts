package main

import (
	"fmt"
)

func findLucky(arr []int) int {
	hashMap := make(map[int]int)
	bubble := make([][]int, len(arr))
	if len(arr) < 2 {
		return -1
	}
	for _, a := range arr {
		hashMap[a]++
	}
	for k, v := range hashMap {
		bubble[v] = append(bubble[v], k)
	}
	for i := len(bubble) - 1; i >= 0; i-- {
		if len(bubble[i]) != 0 {
			for _, n := range bubble[i] {
				if i == n {
					return n
				}
			}
		}
	}
	return -1
}

func main() {
	fmt.Println()
	testCase(findLucky)
}

func testCase(fn func(arr []int) int) {
	cases := []struct {
		arr  []int
		want int
	}{
		{
			arr:  []int{1, 2, 2, 3, 3, 3},
			want: 3,
		},
		{
			arr:  []int{2, 2, 3, 4},
			want: 2,
		},
		{
			arr:  []int{2, 2, 2, 3, 3},
			want: -1,
		},
		{
			arr:  []int{5},
			want: -1,
		},
	}

	for _, tc := range cases {
		got := fn(tc.arr)
		if got == tc.want {
			fmt.Println("PASS")
		} else {
			fmt.Println("FAIL")
		}
	}
}
