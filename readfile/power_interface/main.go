package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func init() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: ./example2 <url> ", os.Args)
		os.Exit(-1)
	}
}
func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func main() {
	// get response from the web server
	res, err := http.Get(os.Args[0])
	check(err)

	// copies from Body to Stdout
	// os.Stdout implement the io.Writer and res.Body imple io.Reader
	io.Copy(os.Stdout, res.Body)

	if e := res.Body.Close(); e != nil {
		fmt.Println("did you print? ", e)
	}

}
