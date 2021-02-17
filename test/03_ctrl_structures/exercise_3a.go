package main

import (
	"fmt"
)

func average(one, two, three float64) float64 {
	return (one + two + three) / 3
}

func average2(numbers ...float64) float64 {
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

func main() {
	for x, y := 'a', 0; x < 'a'+26; x++ {
		if y%2 != 0 {
			fmt.Println(string(x))
		}
		y++
	}
	fmt.Println(average2(3.3, 5.2, 3))
}
