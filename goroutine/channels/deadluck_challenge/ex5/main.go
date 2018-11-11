package main

import (
	"strconv"
	"fmt"
)

func increment(n int) <-chan string {
	out := make(chan string)
	done := make(chan bool)
	for i:=0; i<n; i++{
		go func(i int) {
			for j:=0; j<20; j++{
				out <- fmt.Sprint("count: NO"+strconv.Itoa(i)+" "+strconv.Itoa(j))
			}
			done <- true
		}(i)
	}
	go func() {
		for i:=0; i<n; i++{
			<-done
		}
		close(out)
	}()
	return out
}

func main() {
	var count int
	for n:= range increment(2){
		fmt.Println(n)
		count ++
	}
	fmt.Println(count)
}