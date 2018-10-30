package main

import "fmt"


func main() {
	myGreeting := map[int]string{
		0: "Good morning",
		1: "Bonjour!",
		2: "soslsoso",
		3: "hello you",
	}

	fmt.Println(myGreeting)
	fmt.Println("-----------------------------------")
	if value, ok := myGreeting[4]; ok {
		delete(myGreeting, 2)
		fmt.Println("val: ", value)
		fmt.Println("oks: ", ok)
	}else {
		fmt.Println("That value doesn't ok..")
		fmt.Println("val: ", value)
		fmt.Println("oks: ", ok)
	}
	// just testing string to byte and back to string
	fmt.Println(string("Anthony"[2]))
}