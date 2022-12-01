package main

import "fmt"

//Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
//
//
//
//Example 1:
//
//Input: nums = [1,1,1,2,2,3], k = 2
//Output: [1,2]
//Example 2:
//
//Input: nums = [1], k = 1
//Output: [1]

func main() {
	fmt.Println(topKFrequent([]int{1}, 1))
	unsorted := []int{12, 11, 13, 5, 6}
	insertionSort(unsorted)
	fmt.Println(unsorted)
}

func topKFrequent(nums []int, k int) []int {
	var result []int
	hashMap := make(map[int]int)
	for _, num := range nums {
		hashMap[num]++
	}

	bucketSort := make([][]int, len(nums)+1)

	for k, v := range hashMap {
		bucketSort[v] = append(bucketSort[v], k)
	}

	for i := len(bucketSort) - 1; i > 0; i-- {
		if len(bucketSort[i]) != 0 {
			for _, value := range bucketSort[i] {
				if len(result) == k {
					return result
				}
				result = append(result, value)
			}
		}
	}
	return result
}

func insertionSort(list []int) {
	for i := 0; i < len(list); i++ {
		for x := 0; x < len(list)-1-i; x++ {
			if list[x] > list[x+1] {
				list[x], list[x+1] = list[x+1], list[x]
			}
		}
	}
}
