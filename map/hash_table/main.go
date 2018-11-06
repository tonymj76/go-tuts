package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")

	check(err)
	// why res.Body.Close() not res.Close() becus res.Body is io.reader type which means we are reading a file that we need to close

	defer res.Body.Close()
	reader := bufio.NewReader(res.Body)
	// peek, err := reader.Peek(reader.Size())
	// check(err)
	// fmt.Println(string(peek))
	// fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++")

	// r, err := reader.ReadString('\n')
	// check(err)
	// fmt.Println(r, reader.Size())

	// the use of string.Newreader(f string) *Reader

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
