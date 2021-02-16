package main

import (
	"fmt"
)

func main() {
	for x, y := 'a', 0; x < 'a'+26; x++ {
		if y%2 != 0 {
			fmt.Println(string(x))
		}
		y++
	}
}
