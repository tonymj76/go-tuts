package main

import (
	"fmt"
	"time"
)

func test1(s chan<- string) {
	time.Sleep(2 * time.Second)
	s <- "hello test1"
}

func test2(s chan<- string) {
	time.Sleep(5 * time.Second)
	s <- "hello test2"
}

func main() {
	stream2 := make(chan string)
	stream1 := make(chan string)
	go test1(stream1)
	go test2(stream2)

loop:
	for {
		select {
		case val := <-stream1:
			fmt.Println("stream 1 returns ", val)
			break loop
		case <-stream2:
			fmt.Println("stream 2 return")
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("default returns")
		}
	}

	// s1 := []int{1, 2, 3}
	// s2 := []int{4, 5, 6, 7}
	// fmt.Println(s1)

	// n := copy(s1, s2)
	// fmt.Println(n)
	// fmt.Println(s1, s2)
	fmt.Println([]int{} == nil)
}
