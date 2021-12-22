package main

import "fmt"

func main() {
	n := []int{1, 3, 5, 8, 4}

	n = append([]int{7}, n...)
	fmt.Println(Reverse("demeyor@76"))
	fmt.Println(ReverseList(n))

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
