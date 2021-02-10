package sum

// Sum two numbers
func Sum(x, y int) int {
	return x + y
}

// List of integer
func List(args []int) int {
	var sum int
	for _, x := range args {
		sum += x
	}
	return sum
}
