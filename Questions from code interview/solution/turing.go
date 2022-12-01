package main

import (
	"fmt"
	"math"
)

func findLucky(arr []int) int {
	result := math.Inf(-1)
	var windowStart int
	//hashMap := make(map[int]int)
	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		count := 0
		for arr[windowEnd] == arr[windowStart+1] {
			count++
			windowStart++
		}
	}
	return int(result)
}

func groupAnagrams(strs []string) [][]string {
	hashmap := make(map[[26]int32][]string)
	var result [][]string
	for _, str := range strs {
		var count [26]int32

		for _, s := range str {
			count[s-'a']++
		}
		hashmap[count] = append(hashmap[count], str)
	}
	for _, v := range hashmap {
		result = append(result, v)
	}
	return nil
}

func main() {
	//nums := []int{1, 2, 2, 3, 3, 3}
	//
	//fmt.Println(findLucky(nums))
	fmt.Println(rune(97))
	groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
}
