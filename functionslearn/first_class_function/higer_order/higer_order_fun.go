package main

import "fmt"

func simple(a func(one, two int) int, o, t int) int {
	return a(o, t)
}
func test1() {
	f := func(arg1, arg2 int) int {
		return arg1 + arg2
	}

	fmt.Println(simple(f, 2, 3))
}

func simple2(a func(arg1, arg2 int) func() int, x, b int) func() int {
	return a(x, b)
}

func a(arg1, arg2 int) func() int {
	return func() int {
		return (arg1 + arg2)
	}
}

func mainSimple2() {
	g := simple2(a, 2, 3)
	fmt.Println(g())
}

func greatings() func(arg string) string {
	h := "hello"
	f := func(arg string) string{
		h = h + " " + arg
		return h
	}
	return f
}

func main() {
	a := greatings()
	b := greatings()
	fmt.Println(a("world"))
	fmt.Println(b("tony"))

	fmt.Println(a("!!"))
	fmt.Println(b("and jerry"))
	a("something")
	fmt.Println(a("else"))
}