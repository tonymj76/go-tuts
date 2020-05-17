package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	c := make(chan int)

	go func() {
		defer wg.Done()
		// receive item
		rec := <-c
		fmt.Println(rec)
	}()
	c <- 1 // this means send item
	wg.Wait()

	away()
}

func away() {
	c := make(chan int)

	go func() {
		c <- 1999
		close(c)
	}()
	fmt.Println(<-c)
}
