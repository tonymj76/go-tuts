package main

import "fmt"

func first() {
	fmt.Println("Am First")
}

func second() {
	fmt.Println("Am second")
}


func main(){
	defer  first()
	second()
}