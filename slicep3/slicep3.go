package main

import "fmt"

var buffer [200]byte

func addOnetoslice(slice []int){
	for i := range slice {
		//slice[i] = slice[i] + i instand of this
		slice[i]++
	}

}

func addtoslice() {
	slice := make([]int, 10)
	
	for i:=0; i<len(slice); i++{
		slice[i] = int(i)}
	fmt.Println("before: ", slice)
	addOnetoslice(slice)
	fmt.Println("After: ", slice)
}
// using byte below the up cade use int

func subtractonefromslice(slice []byte) []byte{
	return slice[:len(slice)-1]
}

func main() {
	slice := buffer[:10]
	fmt.Printf("before : %T  %v\n", slice, len(slice))
	newSlice := subtractonefromslice(slice)
	fmt.Println("cap newSlice: ",cap(newSlice))
	fmt.Println("cap slice: ",cap(slice))
	fmt.Println(len(newSlice))
	fmt.Println(len(slice))
	fmt.Println(newSlice)
	fmt.Println(slice)
	newSlice[0] = 53
	fmt.Println(newSlice)
	fmt.Println(slice)
}



























