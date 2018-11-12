package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func runs(baton chan int) {
	var newRunner int
	runner, ok := <-baton
	if !ok {
		fmt.Println("We reach the end of the race")

	}

	if runner != 4 {
		newRunner = runner
		newRunner++
		fmt.Printf("The baton has been pass to %d and the race has started\n", newRunner)
		// this runner is waiting for the baton or chan to start running
		go runs(baton)
	}
	time.Sleep(1e3 * time.Millisecond)

	if runner == 4 {
		fmt.Printf("the race is over the winner is %d\n", runner)
		wg.Done()
		return
	}
	fmt.Printf("runner %d has pass the baton to new runner %d\n", runner, newRunner)
	// send
	baton <- newRunner

}

func main() {
	wg.Add(1)
	ch := make(chan int)
	go runs(ch)
	ch <- 1
	wg.Wait()

}

/*
what i learn about this game is that:
1) when you think about channels also think go routines functions
2) Note when you want to end the game with wg.Done,
3) also note where you send item to the channel also matters*/
