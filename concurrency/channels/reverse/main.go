package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

func reverse(s string) string {
	runeStr := []rune(s)
	for i, j := 0, len(runeStr)-1; i < j; i, j = i+1, j-1 {
		runeStr[i], runeStr[j] = runeStr[j], runeStr[i]
	}
	return string(runeStr)
}

func walker(s []string, rStrStream chan<- string) {
	for _, data := range s {
		go func(d string) {
			rStrStream <- reverse(d)
		}(data)
	}
	// defer close(rStrStream)
}

func walker2(s []string) (h []string) {
	for _, str := range s {
		h = append(h, reverse(str))
	}
	return
}

func walker3(strs []string) <-chan string {
	strStreem := make(chan string, len(strs))
	go func() {
		for i := 0; i < len(strs); i++ {
			strStreem <- reverse(strs[i])
		}
		close(strStreem)
	}()
	return strStreem
}

//We can't close here because we want the walkers to finish all the works giving to them
func walker4(strStream <-chan string, reverseStrStream chan<- string) {
	for str := range strStream {
		reverseStrStream <- reverse(str)
	}
}

func main() {
	var wg sync.WaitGroup
	str := "this just send the data to the channel and main exits without waithing for the goroutine"
	strSlice := strings.Fields(str)

	// Q. how can you do this with buffer channels
	reverseStrStream := make(chan string)
	strStream := make(chan string, len(strSlice))

	// create a walker
	for i := 0; i < (len(strSlice) / 2); i++ {
		go walker4(strStream, reverseStrStream)
	}

	go func() {
		for _, s := range strSlice {
			strStream <- s
		}
	}()

	// walker(s, reverseStrStream)
	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(<-reverseStrStream)
	// }

	//walker 3
	// for reverseStr := range walker3(strSlice) {
	// 	fmt.Println(reverseStr)
	// }

	//walker 4
	for i := 0; i < len(strSlice); i++ {
		fmt.Println(<-reverseStrStream)
	}
	close(reverseStrStream)
	close(strStream)

	//this just shows that closing a chan will unblock all goroutines
	begin := make(chan interface{})
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(x int) { //this func will not run unless we close the chann or send to it
			defer wg.Done()

			fmt.Printf("%v has begun %v\n", x, <-begin)
		}(i)
	}
	fmt.Println("unblocking all goroutines")
	close(begin)
	wg.Wait()
	fmt.Println(runtime.NumCPU())

}
