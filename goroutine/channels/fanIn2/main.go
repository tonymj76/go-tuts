package main

import (
	"fmt"
)

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func fanout(in <-chan int, n int, out chan<- int) {
	for i := 0; i < n; i++ {
		factorial(in, out)
	}
}

func factorial(in <-chan int, out chan<- int) {
	go func() {
		for n := range in {
			out <- fact(n)
		}
	}()
}

func fact(num int) int {
	if num == 1 {
		return num
	}
	return num * fact(num-1)
}

func main() {
	in := gen()
	out := make(chan int)
	fanout(in, 10, out)
	go func() {
		for n := range out {
			fmt.Println(n)
		}

	}()

	var a string
	fmt.Scanln(&a)
}
