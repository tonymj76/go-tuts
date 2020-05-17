package main

import (
	"fmt"
	"time"
)

func closeLeak(done, strings <-chan string) <-chan interface{} {
	terminate := make(chan interface{})
	go func() {
		defer fmt.Println("closing closeLeak")
		defer close(terminate)
		for {
			select {
			case s := <-strings:
				fmt.Print("wait for ... ")
				fmt.Println(s)
				return
			case <-done:
				return // dont use break here
			}
		}
	}()
	return terminate
}

func writeString() <-chan string {
	value := make(chan string)
	go func() {
		value <- "yes am the second value"
		close(value) //should always be in goroutine
	}()
	return value
}

func main() {
	done := make(chan string)
	go func() {
		time.Sleep(4 * time.Microsecond)
		fmt.Println("canceling closeLeak insidous")
		close(done)
	}()
	// <-closeLeak(done, nil)
	fmt.Println("holding main function")
	<-closeLeak(done, writeString())
	fmt.Println("done")
	time.Sleep(2 * time.Second)
}
