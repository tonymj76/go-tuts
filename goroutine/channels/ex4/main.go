package main

import (
	"fmt"
)

func getIt(str ...string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for _, str := range str {
			//sends
			ch <- str
		}
	}()
	return ch
}

// recieve only since we are closing the chan when it recieved an item
func recStr(str <-chan string) chan string {
	addstr := make(chan string)
	go func() {
		for str := range str {
			// sends
			addstr <- fmt.Sprintf("hello %s", str)
		}
		close(addstr)
	}()
	return addstr
}
func main() {
	hc := getIt("Anthony", "Bobby")
	re := recStr(hc)
	fmt.Println(<-re)
	fmt.Println(<-re)
}
