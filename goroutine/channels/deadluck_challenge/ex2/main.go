package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			//send
			c <- i
		}
		// Note while using n range you have to close the chan to 
		// signal that the chan is close so that the n range expression won't 
		// wait for more items while the loop func has stop
		close(c) 
	}()
	// receive
	/* fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c )*/

	for n :=range c {
		fmt.Println(n)
	}

	// if you don't want to use the close() func then use a loop that will terminate after it reach the max or end put of the go routine loop func
	
	/* for i:=0; i<10; i++{
		fmt.Println(<-c)
	} */
}

// Why does this only print zero?
// Ans Because we are just printing the 1st element in the chan
// And what can you do to get it to print all 0 - 9 numbers?
// we can either dublicate the print statement up to ten or add in n range of the chan sends