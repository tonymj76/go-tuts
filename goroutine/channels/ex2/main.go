package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i:= 0; i<10; i++{
			c <- i
		}
		// once you close channel that means it can't receive any input again
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}