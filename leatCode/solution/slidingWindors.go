package solution

import (
	"fmt"
	"math"
)

// LengthOfLongestSubstring Given a string s, find the length of the longest substring without repeating characters.
func LengthOfLongestSubstring(s string) int {
	maxValue := math.Inf(-1)
	hashMap := make(map[rune]struct{})
	var list []rune
	for _, ru := range s {

		if _, ok := hashMap[ru]; ok {
			maxValue = math.Max(maxValue, float64(len(list)))
			first := list[0]
			list = list[1:]
			delete(hashMap, first)
		}

		list = append(list, ru)
		hashMap[ru] = struct{}{}
	}
	fmt.Println(hashMap, list, maxValue)
	return int(maxValue)
}

func TestCaseSubSgtring(fn func(s string) int) {
	testCase := []struct {
		want int
		cs   string
	}{
		{
			cs:   "abcabcbb",
			want: 3,
		},
		{
			cs:   "bbbbb",
			want: 1,
		},
		{
			cs:   "pwwkew",
			want: 3,
		},
	}
	for _, tc := range testCase {
		got := fn(tc.cs)
		if got == tc.want {
			fmt.Println("PASS")
		} else {
			fmt.Println("Faild")
		}
	}
}
