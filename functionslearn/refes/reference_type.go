package main

import  "fmt"
// shows you don't need pointers for slice, map and chan due to the are pass by value references
func changeMe(letter []string){
	letter[0] = "tony"
	fmt.Println(letter)
}

func main() {
	le := make([]string, 1, 25)
	fmt.Println(le)
	changeMe(le)
	fmt.Println(le)
}