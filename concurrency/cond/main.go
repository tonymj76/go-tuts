package main

import (
	"fmt"
	"sync"
	"time"
)

//Cond is a rendezvous point where goroutines are waiting for or announcing the occurence of an event
//before they move
func main() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromSlice := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removing from the queue")
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("Adding to queue and len is ", len(queue))
		queue = append(queue, struct{}{})
		fmt.Println("After adding to queue len is ", len(queue))
		go removeFromSlice(1 * time.Second)
		c.L.Unlock()
	}
	fmt.Println(queue, len(queue), cap(queue))
}

//This can be useful whan you have a generated file let say log file that needs to be clean after it reach some
//certain size
