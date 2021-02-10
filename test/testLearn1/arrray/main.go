package main

import (
	"fmt"
	"math"
)

func mapping(node int, from []int, to []int) {
	var newArray []int
	for x := 1; x <= node; x++ {
		newArray = append(newArray, x)
	}
	newMap := make(map[int][]int)
	for index, n := range from {
		if _, ok := newMap[n]; ok {
			newMap[n] = append(newMap[n], to[index])
		} else {
			newMap[n] = append(newMap[n], to[index])
		}
	}
	for key, value := range newMap {
		for v := 0; v < len(value); v++ {
			if h, ok := newMap[value[v]]; ok {
				newMap[key] = append(newMap[key], h...)
			}
		}
	}

	for _, key := range to {
		for k, _ := range newMap {
			if key == k {
				delete(newMap, k)
			}
		}
	}

	cells := []int{}
	for key, value := range newMap {
		cells = append(cells, value...)
		cells = append(cells, key)
	}
	//intersections

	resultSlice := []int{}
	checkMap := map[int]struct{}{}
	for _, arr := range cells {
		checkMap[arr] = struct{}{}
	}
	for _, sArr := range newArray {
		if _, ok := checkMap[sArr]; !ok {
			resultSlice = append(resultSlice, sArr)
		}
	}
	fmt.Println(newMap)
	fmt.Println()
	fmt.Println(resultSlice)
	sum := 0
	for _, value := range newMap {
		sum += int(math.Ceil(math.Sqrt(float64(len(value) + 1))))
	}
	fmt.Println(sum + len(resultSlice))
}

func main() {
	mapping(8, []int{8, 5}, []int{1, 8})
	mapping(16, []int{6, 11, 9, 5, 11, 9, 15, 9}, []int{13, 15, 12, 14, 15, 16, 1, 16})
}
