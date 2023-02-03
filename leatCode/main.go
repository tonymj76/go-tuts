package main

import (
	"fmt"
	"slice_code_test/leatCode/solution"
)

func main() {
	fmt.Println(solution.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println()
	solution.ProductExceptSelf([]int{1, 2, 3, 4})
	fmt.Println()
	fmt.Println(solution.MaxSlidingWindow([]int{9, 11}, 2))
	solution.TestCaseSubSgtring(solution.LengthOfLongestSubstring)
}
