package main

import (
	"os"
	"slice_code_test/goroutine/go_action2/concurrency/pattern/runner"
	"log"
	"time"
)
// timeout is the number of second the program has to finish.
const timeout = 7 * time.Second

// main is the entry point for the progrma
func main() {
	log.Println("Starting work.")
	
	// create a new timer value for this run
	r := runner.New(timeout)

	//Add the tasks to be run
	r.Add(createTask(), createTask(), createTask())

	// Run the tasks and handle the result.
	if err := r.Start(); err != nil {
		switch err{
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}
	log.Println("Process ended.")
}

// createTask returns an example task that sleeps for the specified
// number of seconds based on the id.
func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}