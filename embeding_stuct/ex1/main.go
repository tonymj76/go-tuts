package main

import "fmt"

type namer struct {
	name string
}

type people struct {
	namer
	age int
}

type company struct {
	namer
	job []string
}

// giveName is abstract type
type giveName interface {
	myName() (string, bool)
}

func (n namer) myName() (string, bool) {
	return n.name, len(n.name) > 0
}

func showName(arg giveName) {
	if name, ok := arg.myName(); ok {
		fmt.Println(name)
	}

}
func main() {
	p1 := people{
		namer: namer{name: "jerry"},
		age:   43,
	}

	c := company{
		namer: namer{name: "Build2day"},
		job:   []string{"computer op", "dish washer", "something"},
	}

	fmt.Println(c.job[2])
	showName(c)
	fmt.Println("++++++++++++++++++++++++++++++")
	showName(p1)
}
