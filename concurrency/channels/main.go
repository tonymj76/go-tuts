package main

import (
	"fmt"
)

func reverse(s string) string {
	runeStr := []rune(s)
	for i, j := 0, len(runeStr)-1; i < j; i, j = i+1, j-1 {
		runeStr[i], runeStr[j] = runeStr[j], runeStr[i]
	}
	return string(runeStr)
}

func main() {
	dataStream := make(chan string)
	alphaStream := make(chan string)

	// go func() {
	// 	fmt.Println(<-dataStream)
	// }()

	// dataStream <- "Hello Tony "// this just send the data to the channel and main exits without waithing for the goroutine

	go func() {
		//if 0 != 1 { // take this out to remove deadlock
		//	return
		//}
		dataStream <- reverse("Hello Tony")
	}()

	fmt.Println(<-dataStream) // this blocks main function and wait to receive data
	go send(alphaStream)
	read(alphaStream)
}

func read(r <-chan string) {
	for msg := range r {
		fmt.Println(msg)
	}
}

func send(s chan<- string) {
	for i := 0; i < 10; i++ {
		al := rune('a' + i)
		s <- string(al)
	}
	close(s)
}
