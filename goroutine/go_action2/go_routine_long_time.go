package main

import ("fmt"
	"sync"
	"runtime")

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)
	wg.Add(2)

	go primeNumber("A")
	go primeNumber("B")

	fmt.Println("waiting to finish")
	wg.Wait()
	fmt.Println("terminate the program")
}

func primeNumber(p string) {
		defer wg.Done()
// seriously what is label next in golang
next:
		for outer:=2; outer < 50000; outer++{
			for inner :=2 ; inner < outer; inner++ {
				if outer%inner == 0{
					continue next  
				}
			}
			fmt.Printf("%s:%d\n", p, outer)
		}
		fmt.Println("Completed", p)
	}