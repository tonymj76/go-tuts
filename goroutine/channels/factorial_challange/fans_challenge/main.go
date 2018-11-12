package main

import (
	"fmt"
	"sync"
)

func main() {
	//out := make(chan int)
	// for i:=0; i<500; i++{
	// 	out <- <-factorize(gen(3, 13))
	// 	//out <- <-factorize(gen(13, 23))
	// }
	// this one works
	// for n := range factorize(gen(3, 13)){
	// 	fmt.Println(n)
	// }

	//defer close(out)
	/* for i:=0; i<500; i++{
		for n:= range factorize(gen(3, 13)) {
			out <- n
		}
	}
	for n:= range out {
		fmt.Println(n)
	} */

	ch1 := factorize(gen(3, 13))
	ch2 := factorize(gen(13, 23))
	ch3 := factorize(gen(13, 23))

	for n := range merge(ch1, ch2, ch3) {
		fmt.Println(n)
	}
}

func gen(f1, f2 int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for j := 0; j < 10000; j++ {
			for i := f1; i < f2; i++ {
				out <- i
			}
		}
	}()
	return out
}

func fact(n int) int {
	if n == 1 {
		return n
	}
	return n * fact(n-1)
}

func factorize(ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range ch {
			out <- fact(n)
		}
	}()
	return out
}

func merge(ch ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	c := make(chan int)

	outPut := func(cho <-chan int) {
		for n := range cho {
			c <- n
		}
		defer wg.Done()
	}
	wg.Add(len(ch))
	for _, x := range ch {
		go outPut(x)
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	return c
}

// this won't work because we need to loop tro the chan
/* func factorize(start, stop int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		out <- fact(<-gen(start, stop))
	}()
	return out

	the main works start at the merge function where 3 go routine was created to excute each
	function declared, i was getting errors because there isn't any go routine that will keep receiving the input until is empty
} */
