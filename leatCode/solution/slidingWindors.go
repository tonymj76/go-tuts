package solution

import (
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

/*
You are given a string s and an integer k.
You can choose any character of the string and change it to
any other uppercase English character. You can perform this
operation at most k times.

Return the length of the longest substring containing the same
letter you can get after performing the above operations.
*/

func characterReplacement(s string, k int) int {
	var startW int
	result := 0.0
	hashMap := make(map[rune]int)
	for endW, ru := range s {
		hashMap[ru]++
		for (endW - startW - getKeyMax(hashMap) + 1) > k {
			hashMap[rune(s[startW])]--
			startW++
		}
		result = math.Max(result, float64(endW-startW+1))
	}
	return int(result)
}

func getKeyMax(hash map[rune]int) int {
	var max int
	for _, v := range hash {
		if max < v {
			max = v
		}
	}
	return max
}
