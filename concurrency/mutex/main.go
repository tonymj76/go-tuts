package main

import (
	"fmt"
	"sync"
)

func main() {
	var lock sync.Mutex
	var count int
	var wg sync.WaitGroup
	loop := 5

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Println("Increment = ", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Println("decrement = ", count)
	}
	wg.Add(10)
	for i := 0; i < loop; i++ {
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	for i := 0; i < loop; i++ {
		go func() {
			defer wg.Done()
			decrement()
		}()
	}
	wg.Wait()
}
