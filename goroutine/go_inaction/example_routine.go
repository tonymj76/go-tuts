package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// how many logical processor we want to use
	runtime.GOMAXPROCS(1)
	//schedular that create two goroutine and wait for them to finish
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")
	// goroutine
	go func() {
		// defer done untill the goroutine finishes
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for l := 'A'; l < 'A'+26; l++ {
				fmt.Printf("%q", l)
			}

		}
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}

		}
	}()
	// main wait for the 2 goroutine to run finish
	wg.Wait()
	fmt.Println("\nTerminating program")
}
