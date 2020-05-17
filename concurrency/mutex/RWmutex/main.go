package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

func producer(wg *sync.WaitGroup, lock sync.Locker) {
	defer wg.Done()
	for i := 5; i > 0; i-- {
		lock.Lock()
		lock.Unlock()
		time.Sleep(1)
	}
}

func observer(wg *sync.WaitGroup, lock sync.Locker) {
	lock.Lock()
	lock.Unlock()
	wg.Done()
}

func test(count int, mutex, RWmutex sync.Locker) time.Duration {
	var wg sync.WaitGroup
	wg.Add(count + 1)
	defer wg.Wait()
	beginTestTime := time.Now()
	go producer(&wg, mutex)
	for i := count; i > 0; i-- {
		go observer(&wg, RWmutex)
	}
	return time.Since(beginTestTime)
}

func main() {
	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()
	var m sync.RWMutex
	fmt.Fprintln(tw, "Readers\tRWMutex\tMutex")
	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(tw, "%d\t%v\t%v\n", count, test(count, &m, m.RLocker()), test(count, &m, &m))
	}
}
