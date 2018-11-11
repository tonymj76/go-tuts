package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		var num int
		num = <-c
		fmt.Println(num)
	}()
	c <- 1
	close(c)
	// the best way is this
	// make a chan in the main function
	ch := make(chan int) 
	// create a go routine that sends to the chan created in main
	go func() {
		// send value
		ch <- 1
	}()
	// receive the value in the main function which also luck the main
	// function from exiting because of <-ch in print statement
	fmt.Println(<-ch)
}

// This results in a deadlock.
// Can you determine why?
// at line 9 i is send to c(channel) now at this state is waiting for
// another go routine to receive the value
// And what would you do to fix it?
// introduce a goroutine that will recieve the chan int and wait deligently for the chan to send in an int to work with
