package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func factorial(c chan int) chan int {
	out := make(chan int)
	total := 1
	go func() {
		defer wg.Done()
		for i := <-c; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}

func main() {
	wg.Add(1)
	c := make(chan int)
	g := factorial(c)
	c <- 4
	for n := range g {
		fmt.Println(n)
	}
	wg.Wait()
}
