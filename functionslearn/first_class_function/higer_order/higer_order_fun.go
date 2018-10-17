package main

import "fmt"


func simple(a func(one , two int) int, o, t int) int {
	return a(o, t)
}
func test1(){
	f := func(arg1, arg2 int) int{
		return arg1 + arg2
	}

	fmt.Println(simple(f, 2, 3))
}


func simple2(a func(arg1, arg2 int) func() int, x,b int ) (func() int) {
	return a(x, b)
}

func a(arg1, arg2 int) func() int{
	return func() int{
		return (arg1 + arg2)
	}
}

func main(){
	g := simple2(a, 2, 3)
	fmt.Println(g())
}