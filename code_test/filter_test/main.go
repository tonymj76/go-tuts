package main

import (
	"fmt"
	"strings"
)

func filterDoubleSlice(slice []string) []string {
	var newslice []string
	for i := 0; i < len(slice); i++ {
		if strings.Index(strings.Join(newslice, " "), slice[i]) == -1 {
			newslice = append(newslice, slice[i])
		}
	}
	return newslice
}

func main() {
	fmt.Println(filterDoubleSlice([]string{"name", "name", "set", "set"}))
}

func filterDouble(str string) string {
	var newStr string
	for i := 0; i < len(str); i++ {
		if strings.Index(newStr, string(str[i])) == -1 {
			newStr += string(str[i])
		}
	}
	return newStr
}

/**
Find Lucky Integer in an Array
Given an array of integers arr, a lucky integer is an integer which has a frequency (how many times an element occurs within the array) in the array equal to its value.

Return a lucky integer in the array. If there are multiple lucky integers return the largest of them. If there is no lucky integer return -1.

Example 1:

Input: arr = [2,2,3,4]
Output: 2
Explanation:

The only lucky number in the array is 2 because it occurs twice within the array which is equal to its value 2.

Example 2:

Input: arr = [1,2,2,3,3,3]
Output: 3
Explanation:

1, 2 and 3 are all lucky numbers and the largest is returned.

Example 3:

Input: arr = [2,2,2,3,3]
Output: -1
Explanation:

There are no lucky numbers in the array because 2 occurs 3 times and 3 occurs twice in the array.

Example 4:

Input: arr = [5]
Output: -1
**/

func findLucky(arr []int) int {
	var result int = 0
	// code here
	var windowStart int
	hashMap := make(map[int]int)
	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		currentNum := arr[windowStart]
		if currentNum == arr[windowEnd+1] {
			hashMap[currentNum]++
		}
	}
	return result
}

//func main() {
//	nums := []int{1, 2, 2, 3, 3, 3}
//
//	fmt.Println(findLucky(nums))
//}
