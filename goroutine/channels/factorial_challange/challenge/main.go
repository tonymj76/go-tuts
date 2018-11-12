package main

import (
	"fmt"
)

func factorial(num int, s string) chan int {
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
	fmt.Printf("%s No:%d fac == ", s, num)
	return out
}

func main() {
	for i := 0; i < 100; i++ {
		for n := range factorial(i+1, "foo") {
			fmt.Println(n)
		}

	}
}
