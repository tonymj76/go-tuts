package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// read the input the user insert in the terminal
func readInputUser() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("what is your name: ")
	test, _ := reader.ReadString('\n')
	fmt.Printf("your name is %s", test)
}

// practicing what i learn from learning people program
func scanMyWords() {
	file, err := os.Open("../hidd.txt")
	if err != nil {
		log.Fatalln("No file")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}

//  It helps to process stream of data by splitting it into tokens and removing space between them

func main() {
	input := "foo bar bbax "
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	readInputUser()
	scanMyWords()
}
