package main

import (
	"fmt"
	"math"
)

func main() {
	//n := []int{1, 3, 5, 8, 4}
	//
	//n = append([]int{7}, n...)
	//fmt.Println(Reverse("demeyor@76"))
	//fmt.Println(ReverseList(n))
	//fmt.Println([]rune("0"), string(49))
	//fmt.Println(finMaxSum([]int{4, 2, 1, 7, 8, 1, 2, 8, 1, 0}, 3))
	//fmt.Println(longestSubstringKDistinct("aaahhibc", 2))
	fmt.Println(longestSubstringKDistinct("ytsseeeteewq", 3))
	//fmt.Println(twoSum([]int{3, 2, 3}, 6))
	//fmt.Println(finSmallestSum(8, []int{4, 2, 2, 7, 8, 1, 2, 8, 10}))
}

func Reverse(word string) string {
	var rev string
	for _, r := range word {
		rev = string(r) + rev
	}
	return rev
}

func ReverseList(list []int) []int {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	return list
}

func finMaxSum(n []int, k int) int {
	maxValue := math.Inf(-1)
	currentRunningSum := 0
	for i := 0; i < len(n); i++ {
		currentRunningSum += n[i]
		if i >= k-1 {
			maxValue = math.Max(float64(currentRunningSum), maxValue)
			currentRunningSum -= n[i-(k-1)]
		}
	}
	return int(maxValue)
}

func finSmallestSum(target int, list []int) int {
	minNumber := math.Inf(1)
	var currListSum int
	var startingPoint int
	for endingPoint := 0; endingPoint < len(list); endingPoint++ {
		currListSum += list[endingPoint]
		for currListSum >= target {
			minNumber = math.Min(minNumber, float64(endingPoint-startingPoint+1))
			currListSum -= list[startingPoint]
			startingPoint++
		}
	}
	return int(minNumber)
}

func longestSubstringKDistinct(letters string, k int) int {
	var windowStart int
	var maxLength float64
	hashLetters := make(map[rune]int)
	for windowEnd, letter := range letters {
		rightChar := letter
		hashLetters[rightChar]++
		for len(hashLetters) > k {
			leftChar := rune(letters[windowStart])
			hashLetters[leftChar]--
			if value, ok := hashLetters[leftChar]; ok && value == 0 {
				delete(hashLetters, leftChar)
			}
			windowStart++
		}
		maxLength = math.Max(maxLength, float64(windowEnd-windowStart+1))
	}
	return int(maxLength)
}

func twoSum(list []int, k int) []int {
	ans := make([]int, 2)
	var windowStart int
	currSum := 0
	for endingPoint := 0; endingPoint < len(list); endingPoint++ {
		currSum += list[endingPoint]
		for currSum >= k {
			if currSum == k {
				ans[0], ans[1] = windowStart, endingPoint
				return ans
			}
			currSum -= list[windowStart]
			windowStart++
		}

	}
	return ans
}

func twoSum3(nums []int, target int) []int {
	m := make(map[int]int) // [num][index]
	for i, v := range nums {
		index, ok := m[target-v]
		if ok {
			return []int{i, index}
		}
		m[v] = i
	}
	return []int{}
}

func twoSum2(list []int, k int) []int {
	index := make(map[int]int)
	for i, item := range list {
		target := k - item
		if idx, ok := index[target]; ok {
			return []int{i, idx}
		}
		index[item] = i
	}
	return nil
}

func TestCase(fn func(letters string, k int) int) {
	testCase := []struct {
		l       string
		want, k int
	}{
		{"jsfeekh", 3, 2},
		{"aaahhibc", 5, 2},
		{"ytsseeeteewq", 6, 2},
		{"ytsseeeteewq", 9, 3},
	}
	for _, tc := range testCase {
		got := fn(tc.l, tc.k)
		if got == tc.want {
			fmt.Println("PASS")
		} else {
			fmt.Println("FAIL")
		}
	}
}

const holder = `
	Find the longest substring length with K distinct characters in K
find the longest contiguous sequence of character such that the number of characters that are distinct
does not exceed the value of 2 so there can be one character of a really long length or there are
two characters of a really long length but once we had a third character into the mix we are violating the
criteria
`
