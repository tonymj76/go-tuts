package solution

import (
	"fmt"
	"math"
)

// LengthOfLongestSubstring Given a string s, find the length of the longest substring without repeating characters.
func LengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	}
	maxValue := math.Inf(-1)
	hashMap := make(map[rune]bool)
	var startW int
	for endW, ru := range s {
		for hashMap[ru] {
			delete(hashMap, rune(s[startW]))
			startW++
		}
		hashMap[ru] = true
		maxValue = math.Max(maxValue, float64(endW-startW+1))
	}

	return int(maxValue)
}

func TestCaseSubSgtring(fn func(s string) int) {
	testCase := []struct {
		want int
		cs   string
	}{
		{cs: "abcabcbb", want: 3},
		{cs: "bbbbb", want: 1},
		{cs: "pwwkew", want: 3},
		{cs: "", want: 0},
		{cs: " ", want: 1},
		{cs: "avrd", want: 4},
	}
	for _, tc := range testCase {
		got := fn(tc.cs)
		if got == tc.want {
			fmt.Println("PASS")
		} else {
			fmt.Printf("Faild: want %d ->> got %d\n", tc.want, got)
		}
	}
}
