// Given a collection of intervals, merge all overlapping intervals.
// For example,
// Given [1,3],[8,10],[2,6],[15,18]
// return [8,10],[15,18], [1,6]

// [1,3], [2,6], [8,10], [15,18]
// [1,6], [8,10], [15,18]
// https://app.coderpad.io/24PWFERP
package main

func main() {
	input := [][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}
	store := make(map[int][]int)
	result := [][]int{}

	for x := 0; x < len(input); x++ {
		store[x] = input[x]
	}

	for _, v := range store {
		for x := 0; x < len(input); x++ {
			sec := v[1]
			if sec <= input[x][0] && sec >= input[x][1] {
				if v[0] <= input[x][0] {
					r := [][]int{[]int{v[0], input[x][1]}}
					result = append(result, r...)
				}
			}
		}
	}

}
