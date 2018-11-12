package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		// close the chan once we are done
		defer close(c)
	}()

	// this keep printing 0 once i close the function
	/* for {
		fmt.Println(<-c)
	} */
	for n := range c {
		fmt.Println(n)
	}
}

// This prints the number, but we have received this error:
// fatal error: all goroutines are asleep - deadlock!
// Where is the deadlock?
// Why are all goroutines asleep?
// How can we fix this?
