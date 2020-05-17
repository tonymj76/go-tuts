package main

import "fmt"

func leakGoroutine(strings <-chan string) <-chan string {
	done := make(chan string)
	go func() {
		defer close(done)
		for s := range strings {
			fmt.Println(s)
		}
	}()

	return done
}

func main() {
	leakGoroutine(nil) // Note that is nil value is pass into a chan which will block the go routine because is waiting to recieve a value
	// which makes this go routine stay in the running process until the main funcion exits cause leak of go routine
}
