package main

import (
		"fmt"
		"bufio"
		 "io/ioutil"
		 "os"
		// "io"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// file return the a []byte
	file, err := ioutil.ReadFile("hidd.txt")
	check(err)
	fmt.Println(string(file))
	fmt.Println("-------------------------------------------")

	// f here return the address to the file stored
	f, err := os.Open("hidd.txt")
	check(err)
	defer f.Close()
	//fmt.Println(f)

	b1 := make([]byte, 100)
	ni, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", ni, string(b1))
	fmt.Println("-------------------------------------------")

	o2, err := f.Seek(50, 0)
	check(err)
	b2 := make([]byte, 100)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))
	fmt.Println("-------------------------------------------")

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(100)
	check(err)
	fmt.Printf("10 bytes: %s\n", b4)


}