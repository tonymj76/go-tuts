package main

import (
	"sync"
	"fmt"
)
var wg sync.WaitGroup
func main() {
	wg.Add(1)
	c := make(chan int)

	go func(c <-chan int){
		defer wg.Done()
		// receive item
		rec := <-c
		fmt.Println(rec)
	}(c)
	c <- 1 // this means send item
	wg.Wait()
}
