package main

import (
	"fmt"
)

func fatorial(x int) int {
	if x == 1 {
		return x
	}
	return x * fatorial(x-1)

}

func resultFatorial() {
	var num int
	fmt.Print("what addedNums do you want to find: ")
	fmt.Scan(&num)
	fmt.Println(fatorial(num))
	fmt.Printf("%T\n", num)
}

func main() {
	dict := make([]int, 5)
	dict2 := []int{}
	dict[3] = 55
	fmt.Println(dict)

	fmt.Println(dict2)

}
