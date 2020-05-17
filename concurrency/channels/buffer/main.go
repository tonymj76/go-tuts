package main

import (
	"bytes"
	"fmt"
	"os"
	"sync"
)

func main() {
	var storeBuffer bytes.Buffer
	var wg sync.WaitGroup

	buffStream := make(chan int, 4)
	defer storeBuffer.WriteTo(os.Stdout)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(buffStream)
		for x := 0; x < 8; x++ {
			fmt.Fprintf(&storeBuffer, "Sending %d\n", x)
			buffStream <- x
		}
		fmt.Fprintln(&storeBuffer, "Producer Done")
	}()

	for intergers := range buffStream {
		fmt.Fprintf(&storeBuffer, "Receiveing %d\n", intergers)
	}
}
