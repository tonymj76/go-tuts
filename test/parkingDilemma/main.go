package main

import (
	"fmt"
	"sort"
)

func parkingDilemma(slice []int, k int) int {
	sort.Ints(slice)
	var st int
	for x:=0; x < k; x++ {
		if len(slice[])
	}
	return 3
}

func getDive(slice []int) int {
	return slice[len(slice)] - slice[0]
}

func parkingDilemma2(slice []int, k int) int {
	sort.Ints(slice)
	var sum int
	var st int
	newSlice := []int{}
	if st == len(slice) {
		return sum
	}
	newSlice = append(newSlice, getDive(slice[st:2]))
	fmt.Println(newSlice)
	st++
	return parkingDilemma2(slice, k)
}

func main() {
	fmt.Println(parkingDilemma2([]int{2, 6, 7, 12}, 2))
}
