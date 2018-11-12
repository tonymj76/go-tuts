package main

import (
	"fmt"
)

func factorial(num int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := num; i > 0; i-- {
			total *= i
		}
		// send
		out <- total
		close(out)
	}()
	return out
}

func main() {
	// this take off the item in the chan and also block main from existing
	for n := range factorial(5) {
		fmt.Println(n)
	}
}
