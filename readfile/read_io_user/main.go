package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)
// read the input the user insert in the terminal
func readInputUser() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("what is your name: ")
	test, _ := reader.ReadString('\n')
	fmt.Println(test)
}

//  It helps to process stream of data by splitting it into tokens and removing space between them

func main() {
	input := "foo bar bbax "
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}