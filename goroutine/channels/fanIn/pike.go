package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/* fi := fanIn(boring("ann"), boring("Ernest"))
	for i:=0; i<10; i++{
		fmt.Println(<-fi)
	} */
	for n := range fanIn(boring("ann"), boring("Ernest")) {
		fmt.Println(n)
	}
	fmt.Println("You guys are booling")
}

func boring(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for j := 0; ; j++ {
			ch <- fmt.Sprintf("%s  %d", msg, j)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return ch
}

func fanIn(b1, b2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for {
			ch <- <-b1
		}
	}()
	go func() {
		defer close(ch)
		for {
			ch <- <-b2
		}
	}()
	return ch
}
