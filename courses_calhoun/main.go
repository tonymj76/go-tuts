package main

import "fmt"

func main() {
	n := []int{1, 3, 5, 3, 4}

	n = append([]int{7}, n...)
	fmt.Println(n, 1%2)

}

func Reverse(word string) string {
	var rev string
	for _, r := range word {
		rev = string(r) + rev
	}
	return rev
}

func ReverseList(list []int) []int {
	return []int{}
}
