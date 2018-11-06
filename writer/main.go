package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	by := []byte("hello\nworld")
	err := ioutil.WriteFile("go_writer", by, 0644)
	check(err)

	f, err := os.Create("os_create")
	defer f.Close()
	check(err)
	n1, err := f.Write([]byte("how you doing\n"))
	check(err)
	fmt.Println("how many bytes: ", n1)
	n2, err := f.WriteString("this uses WriteString\n")
	// to save the file properly
	check(err)
	fmt.Println("how many bytes: ", n2)
	f.Sync()

	writer := bufio.NewWriter(f)
	n3, err := writer.WriteString("am using writestring to write")
	check(err)
	fmt.Println("how many bytes: ", n3)
	writer.Flush()

}
