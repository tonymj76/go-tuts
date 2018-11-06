package main

import "fmt"

// the 1st exercise
func diviNum(arg int) (int, bool) {
	if arg <= 1 {
		return 1, false
	}
	if arg%2 == 0 {
		return arg / 2, true
	}
	return arg / 2, false
}

func gather(arg int, fn func(int) (int, bool)) (int, bool) {
	num, ok := fn(arg)
	return num, ok
}

// i failed 1st ex1 correction is
func half(n int) (float64, bool) {
	return float64(n) / 2, n%2 == 0
}

// second exercise
func funcExpresion(arg int) (int, bool) {
	// closure is working here (the arg is visible to the closure function expresion)
	diviNum := func() (int, bool) {
		if arg <= 1 {
			return 1, false
		}
		if arg%2 == 0 {
			return arg / 2, true
		}
		return arg / 2, false
	}
	num, ok := diviNum()
	return num, ok
}

// another way is :
func funcExpresion2(arg int) (int, bool) {
	return func() (int, bool) {
		if arg <= 1 {
			return 1, false
		}
		if arg%2 == 0 {
			return arg / 2, true
		}
		return arg / 2, false
	}()
}

// ex3 variadic func return max of list

func maxList(args ...int) {
	var max int
	for i, arg := range args {
		if max <= arg || i == 0 {
			max = arg
		}
	}
	fmt.Println(max)
}
func minList(args ...int) {
	min := args[0]
	for _, arg := range args {
		if min >= arg {
			min = arg
		}
	}
	fmt.Println(min)
}

// ex4 what is the value of (false && true) || (true && false) || (false && false) || !(false && true) == true

// ex5 make a func foo what will happen if i use a RETURN STATEMENT

func foo(args ...int) {
	if len(args) == 0 {
		fmt.Println("args is empty")
	}
	var sum int
	for _, arg := range args {
		sum += arg
	}
	fmt.Println(sum)
}
func main() {
	fmt.Println(gather(1, diviNum))
	fmt.Println(gather(2, diviNum))
	fmt.Println(funcExpresion(3))
	fmt.Println(funcExpresion(4))
	fmt.Println(funcExpresion2(5))
	fmt.Println(funcExpresion2(6))
	maxList(-200, -700, 1454, 3, 5, 2, 2, 4, 3, 345, 1)
	minList(-200, -700, 1454, 3, 5, 2, 2, 4, 3, 345, 1)
	foo(1, 2)
	foo(1, 3, 2)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
	// corrections
	h, even := half(5)
	fmt.Println(h, even)

}
