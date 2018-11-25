package main

import (
	"fmt"
)

func toChar() {
	for i := 50; i <= 140; i++ {
		fmt.Println(i, "-", string(i), " - ", []byte(string(i)))
	}
}

func main() {
	fmt.Println(string(76))
	fmt.Println(string("anthony"))
	fmt.Println(string(135))
	var sum int

	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			sum += i

		} else if i%5 == 0 {
			sum += i

		} else {
			continue
		}
	}
	fmt.Println(sum)

	var name string
	fmt.Print("enter your name: ")
	fmt.Scan(&name) //address of name and put what it scan there
	fmt.Printf("Hello %s\n", name)
}
