package main

import "fmt"

func main() {
	//prime := []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 10, 9, 15}
	//for _, n := range prime {
	//	fmt.Println(isPrime(n))
	//}
	fmt.Println(solveCount(25))
	fmt.Println(naturalNumber(6))
}

func isPrime(n int) int {
	if n == 2 {
		return 1
	}
	if n%2 == 0 {
		return 0
	}

	var i int
	for i*i <= n {
		if i*i == n {
			return 0
		}
		i++

	}
	return 1
}

func solveCount(A int) int {
	var (
		i, k, sqrt, count, g = 1, 1, 1, 0, 1
	)

	for i*i <= A {
		sqrt = i
		i++
	}

	fmt.Println(sqrt)

	for k <= sqrt {
		if A%k == 0 {
			count += 2
		}
		if !(count > 2) {
			g++
		}
		k++
	}

	return g
}

func naturalNumber(A int) int {

	var (
		i, k, sqrt, sum = 1, 2, 1, 1
	)

	for i*i <= A {
		sqrt = i
		i++
	}

	for k <= sqrt {
		if A%k == 0 {
			if k == A/k {
				sum = sum + k
			} else {
				sum = sum + k + A/k
			}
		}
		k++
	}

	if sum == A {
		return 1
	}
	return 0
}
