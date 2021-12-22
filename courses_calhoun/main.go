package main

import (
	"fmt"
	"math"
)

func main() {
	n := []int{1, 3, 5, 8, 4}

	n = append([]int{7}, n...)
	fmt.Println(Reverse("demeyor@76"))
	fmt.Println(ReverseList(n))
	fmt.Println([]rune("0"), string(49))
	fmt.Println(finMaxSum([]int{4, 2, 1, 7, 8, 1, 2, 8, 1, 0}, 3))
	fmt.Println()
	fmt.Println(finSmallestSum(8, []int{4, 2, 2, 7, 8, 1, 2, 8, 10}))
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
