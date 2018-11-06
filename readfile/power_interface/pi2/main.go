package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var b bytes.Buffer

	b.Write([]byte("Hello \nWorld "))
	//check(err)

	// use Fprintf to CONCATENATE a string to the buffer
	fmt.Fprintf(&b, "Hope you last\n")
	io.Copy(os.Stdout, &b)
}
