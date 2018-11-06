package main

import "fmt"

//import "math"
// General buffer array to create slice
var buffer [200]byte

// this function is pass by value the slice header remains the same
// but the function can have effect on the values
func addOnetoslice(slice []int) {
	for i := range slice {
		//slice[i] = slice[i] + i instend of this
		slice[i]++
	}

}

func addtoslice() {
	slice := make([]int, 10)

	for i := 0; i < len(slice); i++ {
		slice[i] = int(i)
	}
	fmt.Println("before: ", slice)
	addOnetoslice(slice)
	fmt.Println("After: ", slice)
}

// using byte below the up cade use int

func subtractonefromslice(slice []byte) []byte {
	return slice[:len(slice)-1]
}

func sliceHeader() {
	slice := buffer[:10]
	fmt.Printf("before : %T  %v\n", slice, len(slice))
	newSlice := subtractonefromslice(slice)
	fmt.Println("cap newSlice: ", cap(newSlice))
	fmt.Println("cap slice: ", cap(slice))
	fmt.Println(len(newSlice))
	fmt.Println(len(slice))
	fmt.Println(newSlice)
	fmt.Println(slice)
	newSlice[0] = 53
	fmt.Println(newSlice)
	fmt.Println(slice)
}

//* is dereferencing the address to the value,
//(*slice) get the slice(value) but just *slice gets (type *[]bytes)
func modifySliceHeader(slice *[]byte) {
	*slice = (*slice)[:len(*slice)-1]
}

func ptrSliceHeader() {
	slice := buffer[10:20]
	fmt.Println(len(slice))
	modifySliceHeader(&slice)
	fmt.Println(len(slice))
}
