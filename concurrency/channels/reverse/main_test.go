package main

import (
	"strings"
	"testing"
)

func Benchmark_walker(b *testing.B) {

	for i := 0; i < b.N; i++ {
		benchmarkWalker()
	}

}

func Benchmark_walker2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkWalker2()
	}

}
func Benchmark_walker3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkWalker3()
	}

}

func benchmarkWalker() {
	str := "this just send the data to the channel and main exits without waithing for the goroutine"
	s := strings.Fields(str)
	reverseStrStream := make(chan string) // Q. how can you do this with buffer channels
	newr := []string{}

	walker(s, reverseStrStream)
	for i := 0; i < len(s); i++ {
		newr = append(newr, <-reverseStrStream)
	}
}

func benchmarkWalker2() {
	str := "this just send the data to the channel and main exits without waithing for the goroutine"
	s := strings.Fields(str)
	newr := []string{}
	for i := 0; i < len(s); i++ {
		newr = append(newr, walker2(s)[i])
	}
}
func benchmarkWalker3() {
	str := "this just send the data to the channel and main exits without waithing for the goroutine"
	s := strings.Fields(str)
	newr := []string{}
	for r := range walker3(s) {
		newr = append(newr, r)
	}
}
