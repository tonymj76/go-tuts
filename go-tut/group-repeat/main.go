package main

import "fmt"

type Result struct {
	C rune // character
	L int  // count
}

func longestRepetition(text string) Result {
	if text == "" {
		return Result{}
	}
	var holder [][]rune
	var count, max int32
	var index int
	char := rune(text[0])
	for _, str := range text {
		if char == str {
			count++
		} else {
			holder = append(holder, []rune{char, count})
			char = str
			count = 1
		}
	}
	holder = append(holder, []rune{char, count})
	for i := 0; i < len(holder); i++ {
		if holder[i][1] > max {
			max = holder[i][1]
			index = i
		}
	}

	return Result{C: holder[index][0], L: int(max)}
}

var nums = []int{3, 2, 3}
var target = 6

func main() {
	fmt.Println(twoSum2(nums, target))
}

func twoSum(nums []int, target int) []int {
	var currIdx, currentNum int
	for idx, num := range nums {
		if (num + currentNum) == target {
			return []int{currIdx, idx}
		}
		currIdx = idx
		currentNum = num
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// func twoSum3(nums []int, target int) []int {
// 	mapInts := make(map[int]int)
// 	for inx, num := range nums {
// 		mapInts[] = num
// 	}
// 	for
// }

func LongestRepetition(text string) Result {
	var char, res rune
	var count, maxCount int
	for _, c := range text {
		if c == char {
			count++
			continue
		}

		if count > maxCount {
			res = char
			maxCount = count
		}
		char = c
		count = 1
	}
	if count > maxCount {
		res = char
		maxCount = count
	}
	return Result{res, maxCount}
}
