package main

import "fmt"

func main() {
	fmt.Println(solve(20))
	fmt.Println(getShake(19, 0))
	if !(2 > 20) {
		fmt.Println("ddkk")
	}
}

func solve(A int) int {
	var (
		i, c int
	)

	for i*i <= A {
		c = i
		i++
	}
	if c*c != A {
		return -1
	}
	return c
}

func getShake(A int, B int) int {
	totalSlice := (A * 3) + B
	glassShake := totalSlice / 2
	return glassShake
}
