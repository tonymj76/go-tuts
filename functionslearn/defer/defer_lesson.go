package main

import "fmt"

func first() {
	fmt.Println("Am First")
}

func second() {
	fmt.Println("Am second")
}

func main() {
	defer first()
	second()
	learnDefer()
}

//we can also use defer in panic function

func learnDefer() {
	defer func() {
		fmt.Println(recover())
	}()
	panic("pass in word you want.")
}
