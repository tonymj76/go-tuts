package module1

func addNums(num int) int {
	return (num * (num + 1)) / 2
}

// Fibonacci this re calculates the same number over and over again
//func Fibonacci(num int) int{
//	if num <= 1 {
//		return num
//	}
//	return Fibonacci(num - 1) + Fibonacci(num -2)
//}

func Fibonacci(num int) int {
	if num <= 1 {
		return num
	}

	var cur int
	firstSeq := 0
	secondSeq := 1

	for x := 2; x <= num; x++ {
		cur = firstSeq + secondSeq
		firstSeq = secondSeq
		secondSeq = cur
	}
	return cur

	// Note: There are many possible solutions to this problem.
	// We will end up exploring most of them in this course. Below
	// is another viable solution using a technique called recursion.
	// return Fibonacci(n-1) + Fibonacci(n-2)
}
