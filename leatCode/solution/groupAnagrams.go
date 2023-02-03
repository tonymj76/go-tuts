package solution

import "fmt"

func GroupAnagrams(strs []string) [][]string {
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
	return result
}

// groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})

// ProductExceptSelf  Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
// You must write an algorithm that runs in O(n) time and without using the division operation.
func ProductExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	pre, post := 1, 1
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result[i] = pre
		} else {
			result[i] = pre * nums[i-1]
			pre *= nums[i-1]
		}

	}
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= post
		post *= nums[i]
	}
	return result
}
func getMax(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) <= 1 {
		return nums
	}
	var result, kLenght []int
	for end := 0; end < len(nums); end++ {
		kLenght = append(kLenght, nums[end])
		fmt.Println(len(kLenght))
		if len(kLenght) >= k {
			result = append(result, getMax(kLenght))

			kLenght = kLenght[1:]
		}
	}
	return result
}

// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
//
// Example 1:
//
// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]
// Example 2:
//
// Input: nums = [1], k = 1
// Output: [1]
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
